package binance_connector

import (
	"context"
	"encoding/json"
)

type StartUserDataStreamService struct {
	websocketAPI *WebsocketAPIClient
}

func (s *StartUserDataStreamService) Do(ctx context.Context) (*StartUserDataStreamResponse, error) {
	parameters := map[string]interface{}{
		"apiKey": s.websocketAPI.APIKey,
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "userDataStream.start",
		"params": parameters,
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
		var startResponse StartUserDataStreamResponse
		err = json.Unmarshal(response, &startResponse)
		if err != nil {
			return nil, err
		}
		return &startResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type StartUserDataStreamResponse struct {
	ID     string              `json:"id"`
	Status int                 `json:"status"`
	Error  *WsAPIErrorResponse `json:"error,omitempty"`
	Result struct {
		ListenKey string `json:"listenKey,omitempty"`
	} `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit `json:"rateLimits,omitempty"`
}

type PingUserDataStreamService struct {
	websocketAPI *WebsocketAPIClient
	listenKey    string
}

func (s *PingUserDataStreamService) ListenKey(listenKey string) *PingUserDataStreamService {
	s.listenKey = listenKey
	return s
}

func (s *PingUserDataStreamService) Do(ctx context.Context) (*PingUserDataStreamResponse, error) {
	parameters := map[string]interface{}{
		"listenKey": s.listenKey,
		"apiKey":    s.websocketAPI.APIKey,
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "userDataStream.ping",
		"params": parameters,
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
		var pingResponse PingUserDataStreamResponse
		err = json.Unmarshal(response, &pingResponse)
		if err != nil {
			return nil, err
		}
		return &pingResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type PingUserDataStreamResponse struct {
	ID         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Response   struct{}            `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits,omitempty"`
}

type StopUserDataStreamService struct {
	websocketAPI *WebsocketAPIClient
	listenKey    string
}

func (s *StopUserDataStreamService) ListenKey(listenKey string) *StopUserDataStreamService {
	s.listenKey = listenKey
	return s
}

func (s *StopUserDataStreamService) Do(ctx context.Context) (*StopUserDataStreamResponse, error) {
	parameters := map[string]interface{}{
		"listenKey": s.listenKey,
		"apiKey":    s.websocketAPI.APIKey,
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "userDataStream.close",
		"params": parameters,
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
		var stopResponse StopUserDataStreamResponse
		err = json.Unmarshal(response, &stopResponse)
		if err != nil {
			return nil, err
		}
		return &stopResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type StopUserDataStreamResponse struct {
	ID         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Response   struct{}            `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits,omitempty"`
}
