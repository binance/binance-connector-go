package binance_connector

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

type WebsocketAPIClient struct {
	APIKey         string
	APISecret      string
	Endpoint       string
	Conn           *websocket.Conn
	Dialer         *websocket.Dialer
	ReqResponseMap map[string]chan []byte
}

type WsAPIRateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

type WsAPIErrorResponse struct {
	Code    int    `json:"code"`
	ID      string `json:"id"`
	Message string `json:"msg"`
}

var (
	// WebsocketAPITimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketAPITimeout = time.Second * 60
	// WebsocketAPIKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketAPIKeepalive = true
)

func NewWebsocketAPIClient(apiKey string, apiSecret string, baseURL ...string) *WebsocketAPIClient {
	// Set default base URL to production WS URL
	url := "wss://ws-api.binance.com:443/ws-api/v3"

	if len(baseURL) > 0 {
		url = baseURL[0]
	}

	return &WebsocketAPIClient{
		APIKey:    apiKey,
		APISecret: apiSecret,
		Endpoint:  url,
		Dialer: &websocket.Dialer{
			Proxy:             http.ProxyFromEnvironment,
			HandshakeTimeout:  45 * time.Second,
			EnableCompression: false,
		},
	}
}

func (c *WebsocketAPIClient) Connect() error {
	if c.Dialer == nil {
		return fmt.Errorf("dialer not initialized")
	}
	headers := http.Header{}
	headers.Add("User-Agent", fmt.Sprintf("%s/%s", Name, Version))
	conn, _, err := c.Dialer.Dial(c.Endpoint, headers)
	if err != nil {
		return err
	}

	fmt.Println("Connected to Binance Websocket API")
	c.Conn = conn

	c.ReqResponseMap = make(map[string]chan []byte)
	c.startReader() // start reader again
	return nil
}

func (c *WebsocketAPIClient) startReader() {
	go func() {
		for {
			_, message, err := c.Conn.ReadMessage()
			if err != nil {
				log.Println("Error reading:", err)
				return
			}
			c.Handler(message)
		}
	}()
}

// Handler function to handle responses
func (c *WebsocketAPIClient) Handler(message []byte) {
	var response WsAPIErrorResponse
	err := json.Unmarshal(message, &response)
	if err != nil {
		log.Println("Error unmarshaling:", err)
		return
	}
	// Send the message to the corresponding request
	if channel, ok := c.ReqResponseMap[response.ID]; ok {
		channel <- message
	}
}

func (c *WebsocketAPIClient) WaitForCloseSignal() {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)
	<-stopCh
}

func (c *WebsocketAPIClient) Close() error {
	return c.Conn.Close()
}

func (c *WebsocketAPIClient) SendMessage(msg interface{}) error {
	return c.Conn.WriteJSON(msg)
}

func (c *WebsocketAPIClient) RequestHandler(req interface{}, handler WsHandler, errHandler ErrHandler) (stopCh chan struct{}, err error) {
	err = c.SendMessage(req)
	if err != nil {
		return nil, err
	}
	stopCh, err = wsApiServe(c.Conn, handler, errHandler)
	if err != nil {
		return nil, err
	}
	return stopCh, nil
}

func wsApiServe(c *websocket.Conn, handler WsHandler, errHandler ErrHandler) (stopCh chan struct{}, err error) {
	stopCh = make(chan struct{})
	go func() {
		if WebsocketAPIKeepalive {
			keepAlive(c, WebsocketAPITimeout)
		}
		silent := false
		for {
			select {
			case <-stopCh:
				return
			default:
				_, message, err := c.ReadMessage()
				if err != nil {
					if !silent {
						fmt.Println(err)
						errHandler(err)
					}
					continue
				}
				handler(message)
			}
		}
	}()
	return stopCh, nil
}

func websocketAPISignature(apiKey string, apiSecret string, parameters map[string]string) (map[string]string, error) {
	if apiKey == "" || apiSecret == "" {
		return nil, &WebsocketClientError{
			Message: "api_key and api_secret are required for websocket API signature",
		}
	}

	parameters["timestamp"] = strconv.FormatInt(time.Now().Unix()*1000, 10)
	parameters["apiKey"] = apiKey

	// Sort parameters by key
	keys := make([]string, 0, len(parameters))
	for key := range parameters {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Build sorted query string
	var sortedParams []string
	for _, key := range keys {
		sortedParams = append(sortedParams, key+"="+parameters[key])
	}

	// Calculate signature
	queryString := strings.Join(sortedParams, "&")
	signature := hmacHashing(apiSecret, queryString)

	parameters["signature"] = signature

	return parameters, nil
}

func hmacHashing(apiSecret string, data string) string {
	mac := hmac.New(sha256.New, []byte(apiSecret))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

func getUUID() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s", randomHex(8), randomHex(4), randomHex(4), randomHex(4), randomHex(12))
}

func randomHex(n int) string {
	rand.Seed(time.Now().UnixNano())
	hexChars := "0123456789abcdef"
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = hexChars[rand.Intn(len(hexChars))]
	}
	return string(bytes)
}

type WebsocketClientError struct {
	Message string
}

func (e *WebsocketClientError) Error() string {
	return e.Message
}

type TestConnectivityService struct {
	websocketAPI *WebsocketAPIClient
}

func (s *TestConnectivityService) Do(ctx context.Context) (*TestConnectivityResponse, error) {
	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "ping",
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err := s.websocketAPI.SendMessage(payload)
	if err != nil {
		return nil, err
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var pingResponse TestConnectivityResponse
		err = json.Unmarshal(response, &pingResponse)
		if err != nil {
			return nil, err
		}
		return &pingResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type TestConnectivityResponse struct {
	ID         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     struct{}            `json:"result,omitempty"`
	RateLimits WsAPIRateLimit      `json:"rateLimits,omitempty"`
}

type CheckServerTimeService struct {
	websocketAPI *WebsocketAPIClient
}

func (s *CheckServerTimeService) Do(ctx context.Context) (*CheckServerTimeResponse, error) {
	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "time",
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err := s.websocketAPI.SendMessage(payload)
	if err != nil {
		return nil, err
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var timeResponse CheckServerTimeResponse
		err = json.Unmarshal(response, &timeResponse)
		if err != nil {
			return nil, err
		}
		return &timeResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type CheckServerTimeResponse struct {
	ID     string              `json:"id"`
	Status int                 `json:"status"`
	Error  *WsAPIErrorResponse `json:"error,omitempty"`
	Result struct {
		ServerTime uint64 `json:"serverTime"`
	} `json:"result,omitempty"`
	RateLimits WsAPIRateLimit `json:"rateLimits,omitempty"`
}

type ExchangeInformationService struct {
	websocketAPI *WebsocketAPIClient
	symbol       *string
	symbols      *[]string
	permissions  *[]string
}

func (s *ExchangeInformationService) Symbol(symbol string) *ExchangeInformationService {
	s.symbol = &symbol
	return s
}

func (s *ExchangeInformationService) Symbols(symbols []string) *ExchangeInformationService {
	s.symbols = &symbols
	return s
}

func (s *ExchangeInformationService) Permissions(permissions []string) *ExchangeInformationService {
	s.permissions = &permissions
	return s
}

func (s *ExchangeInformationService) Do(ctx context.Context) (*ExchangeInformationResponse, error) {
	payload := map[string]interface{}{
		"id":     getUUID(),
		"method": "exchangeInfo",
	}

	// Only one of the following parameters can be set: symbol, symbols, permissions
	if s.symbol != nil {
		payload["params"] = map[string]interface{}{
			"symbol": *s.symbol,
		}
	} else if s.symbols != nil {
		payload["params"] = map[string]interface{}{
			"symbols": *s.symbols,
		}
	} else if s.permissions != nil {
		payload["params"] = map[string]interface{}{
			"permissions": *s.permissions,
		}
	}

	responseCh := make(chan *ExchangeInformationResponse)

	handler := func(message []byte) {
		var response ExchangeInformationResponse
		err := json.Unmarshal(message, &response)
		if err != nil {
			fmt.Println("Error unmarshaling:", err)
			return
		}
		responseCh <- &response
	}

	errHandler := func(err error) {
		fmt.Println("Error:", err)
	}

	doneCh, err := s.websocketAPI.RequestHandler(payload, handler, errHandler)
	if err != nil {
		return nil, err
	}

	select {
	case <-doneCh:
		return nil, ctx.Err()
	case response := <-responseCh:
		return response, nil
	}
}

type ExchangeInformationResponse struct {
	ID         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     json.RawMessage     `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits,omitempty"`
}

// General Websocket API Endpoints:
func (w *WebsocketAPIClient) NewTestConnectivityService() *TestConnectivityService {
	return &TestConnectivityService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewCheckServerTimeService() *CheckServerTimeService {
	return &CheckServerTimeService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewExchangeInformationService() *ExchangeInformationService {
	return &ExchangeInformationService{websocketAPI: w}
}

// Account Websocket API Endpoints:
func (w *WebsocketAPIClient) NewAccountInformationService() *AccountInformationService {
	return &AccountInformationService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewAccountOrderRateLimitsService() *AccountOrderRateLimitsService {
	return &AccountOrderRateLimitsService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewAccountOrderHistoryService() *AccountOrderHistoryService {
	return &AccountOrderHistoryService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewAccountOCOHistoryService() *AccountOCOHistoryService {
	return &AccountOCOHistoryService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewAccountTradeHistoryService() *AccountTradeHistoryService {
	return &AccountTradeHistoryService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewAccountPreventedMatchesService() *AccountPreventedMatchesService {
	return &AccountPreventedMatchesService{websocketAPI: w}
}

// Market Websocket API Endpoints:
func (w *WebsocketAPIClient) NewDepthService() *DepthService {
	return &DepthService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewRecentTradesService() *RecentTradesService {
	return &RecentTradesService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewHistoricalTradesService() *HistoricalTradesService {
	return &HistoricalTradesService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewAggTradesService() *AggregateTradesService {
	return &AggregateTradesService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewKlinesService() *KlinesService {
	return &KlinesService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewAvgPriceService() *AvgPriceService {
	return &AvgPriceService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewTicker24hrService() *Ticker24hrService {
	return &Ticker24hrService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewTickerService() *TickerService {
	return &TickerService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewTickerPriceService() *TickerPriceService {
	return &TickerPriceService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewTickerBookService() *TickerBookService {
	return &TickerBookService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewUiKlinesService() *UIKlinesService {
	return &UIKlinesService{websocketAPI: w}
}

// Trading Websocket API Endpoints:
func (w *WebsocketAPIClient) NewPlaceNewOrderService() *OrderPlacementService {
	return &OrderPlacementService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewTestPlaceOrderService() *TestOrderPlacementService {
	return &TestOrderPlacementService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewQueryOrderService() *OrderStatusService {
	return &OrderStatusService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewCancelOrderService() *OrderCancelService {
	return &OrderCancelService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewCancelReplaceOrderService() *OrderCancelReplaceService {
	return &OrderCancelReplaceService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewCurrentOpenOrdersService() *OpenOrdersStatusService {
	return &OpenOrdersStatusService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewCancelOpenOrdersService() *OpenOrdersCancelAllService {
	return &OpenOrdersCancelAllService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewPlaceOCOService() *OrderListPlaceService {
	return &OrderListPlaceService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewQueryOCOService() *OrderListStatusService {
	return &OrderListStatusService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewCurrentOpenOCOService() *OpenOrderListsStatusService {
	return &OpenOrderListsStatusService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewCancelOCOService() *OrderListCancelService {
	return &OrderListCancelService{websocketAPI: w}
}

// User Data Websocket API Endpoints:
func (w *WebsocketAPIClient) NewStartUserDataStreamService() *StartUserDataStreamService {
	return &StartUserDataStreamService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewPingUserDataStreamService() *PingUserDataStreamService {
	return &PingUserDataStreamService{websocketAPI: w}
}

func (w *WebsocketAPIClient) NewStopUserDataStreamService() *StopUserDataStreamService {
	return &StopUserDataStreamService{websocketAPI: w}
}
