package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	getC2CTradeHistory = "/sapi/v1/c2c/orderMatch/listUserOrderHistory"
)

// Request requirements
// The max interval between startTime and endTime is 30 days.
// If startTime and endTime are not sent, the recent 7 days' data will be returned.
// The earliest startTime is supported on June 10, 2020
// Return up to 200 records per request.

type GetC2CTradeHistoryService struct {
	c              *Client
	tradeType      *string
	startTimestamp *uint64
	endTimestamp   *uint64
	page           *int
	recvWindow     *uint64
	timestamp      *uint64
}

func (s *GetC2CTradeHistoryService) TradeType(tradeType string) *GetC2CTradeHistoryService {
	s.tradeType = &tradeType
	return s
}

func (s *GetC2CTradeHistoryService) StartTimestamp(startTimestamp uint64) *GetC2CTradeHistoryService {
	s.startTimestamp = &startTimestamp
	return s
}

func (s *GetC2CTradeHistoryService) EndTimestamp(endTimestamp uint64) *GetC2CTradeHistoryService {
	s.endTimestamp = &endTimestamp
	return s
}

func (s *GetC2CTradeHistoryService) Page(page int) *GetC2CTradeHistoryService {
	s.page = &page
	return s
}

func (s *GetC2CTradeHistoryService) RecvWindow(recvWindow uint64) *GetC2CTradeHistoryService {
	s.recvWindow = &recvWindow
	return s
}

func (s *GetC2CTradeHistoryService) Timestamp(timestamp uint64) *GetC2CTradeHistoryService {
	s.timestamp = &timestamp
	return s
}

func (s *GetC2CTradeHistoryService) Do(ctx context.Context) (res *GetC2CTradeHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getC2CTradeHistory,
		secType:  secTypeSigned,
	}

	r.setParam("timestamp", *s.timestamp)
	if s.tradeType != nil {
		r.setParam("tradeType", *s.tradeType)
	}
	if s.startTimestamp != nil {
		r.setParam("startTimestamp", *s.startTimestamp)
	}
	if s.endTimestamp != nil {
		r.setParam("endTimestamp", *s.endTimestamp)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = &GetC2CTradeHistoryResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type GetC2CTradeHistoryResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		OrderNumber         string `json:"orderNumber"`
		AdvNo               string `json:"advNo"`
		TradeType           string `json:"tradeType"`
		Asset               string `json:"asset"`
		Fiat                string `json:"fiat"`
		FiatSymbol          string `json:"fiatSymbol"`
		Amount              string `json:"amount"` //amount in crypto
		TotalPrice          string `json:"totalPrice"`
		UnitPrice           string `json:"unitPrice"`   //in fiat
		OrderStatus         string `json:"orderStatus"` // PENDING, TRADING, BUYER_PAYED, DISTRIBUTING, COMPLETED, IN_APPEAL, CANCELLED, CANCELLED_BY_SYSTEM
		CreateTime          uint64 `json:"createTime"`
		Commission          string `json:"commission"`
		CounterPartNickName string `json:"counterPartNickName"`
		AdvertisementRole   string `json:"advertisementRole"`
	}
	Total   int64 `json:"total"`
	Success bool  `json:"success"`
}
