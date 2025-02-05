package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	getFuturesLeadTraderStatusEndpoint = "/sapi/v1/copyTrading/futures/userStatus"
)

type GetFuturesLeadTraderStatusService struct {
	c *Client
}

func (s *GetFuturesLeadTraderStatusService) Do(ctx context.Context, opts ...RequestOption) (*GetFuturesLeadTraderStatusResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getFuturesLeadTraderStatusEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(GetFuturesLeadTraderStatusResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetFuturesLeadTraderStatusResponse struct {
	IsLeadTrader bool  `json:"isLeadTrader"`
	Time         int64 `json:"time"`
}

const (
	getFuturesLeadTradingSymbolWhitelist = "/sapi/v1/copyTrading/futures/leadSymbol"
)

type GetFuturesLeadTradingSymbolWhitelistService struct {
	c *Client
}

func (s *GetFuturesLeadTradingSymbolWhitelistService) Do(ctx context.Context, opts ...RequestOption) (res *GetFuturesLeadTradingSymbolWhitelistResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getFuturesLeadTradingSymbolWhitelist,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetFuturesLeadTradingSymbolWhitelistResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetFuturesLeadTradingSymbolWhitelistResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Symbol     string `json:"symbol"`
		BaseAsset  string `json:"baseAsset"`
		QuoteAsset string `json:"quoteAsset"`
	} `json:"data"`
}
