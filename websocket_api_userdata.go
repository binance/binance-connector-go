package binance_connector

import (
	"context"
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

	var resp StartUserDataStreamResponse
	if err := s.websocketAPI.Do(ctx, id, payload, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

	var resp PingUserDataStreamResponse
	if err := s.websocketAPI.Do(ctx, id, payload, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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
		"method": "userDataStream.stop",
		"params": parameters,
	}

	var resp StopUserDataStreamResponse
	if err := s.websocketAPI.Do(ctx, id, payload, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StopUserDataStreamResponse struct {
	ID         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Response   struct{}            `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits,omitempty"`
}
