package binance_connector

import (
	"context"
	"net/http"
)

// Create Listen Key
type CreateListenKey struct {
	c *Client
}

// Do send request
func (s *CreateListenKey) Do(ctx context.Context, opts ...RequestOption) (listenKey string, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/api/v3/userDataStream",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}
	j, err := newJSON(data)
	if err != nil {
		return "", err
	}
	listenKey = j.Get("listenKey").MustString()
	return listenKey, nil
}

// Keep Alive/Ping User Stream
type PingUserStream struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *PingUserStream) ListenKey(listenKey string) *PingUserStream {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *PingUserStream) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: "/api/v3/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setParam("listenKey", s.listenKey)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// CloseUserStream delete listen key
type CloseUserStream struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *CloseUserStream) ListenKey(listenKey string) *CloseUserStream {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *CloseUserStream) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/api/v3/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setParam("listenKey", s.listenKey)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}
