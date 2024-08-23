package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	fiatDepositWithdrawHistory = "/sapi/v1/fiat/orders"
)

type GetFiatDepositWithdrawHistoryService struct {
	c               *Client
	transactionType *string
	beginTime       *uint64
	endTime         *uint64
	page            *int
	rows            *int
}

func (s *GetFiatDepositWithdrawHistoryService) TransactionType(transactionType string) *GetFiatDepositWithdrawHistoryService {
	s.transactionType = &transactionType
	return s
}

func (s *GetFiatDepositWithdrawHistoryService) BeginTime(beginTime uint64) *GetFiatDepositWithdrawHistoryService {
	s.beginTime = &beginTime
	return s
}

func (s *GetFiatDepositWithdrawHistoryService) EndTime(endTime uint64) *GetFiatDepositWithdrawHistoryService {
	s.endTime = &endTime
	return s
}

func (s *GetFiatDepositWithdrawHistoryService) Page(page int) *GetFiatDepositWithdrawHistoryService {
	s.page = &page
	return s
}

func (s *GetFiatDepositWithdrawHistoryService) Rows(rows int) *GetFiatDepositWithdrawHistoryService {
	s.rows = &rows
	return s
}

func (s *GetFiatDepositWithdrawHistoryService) Do(ctx context.Context) (res *GetFiatDepositWithdrawHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fiatDepositWithdrawHistory,
		secType:  secTypeSigned,
	}
	r.setParam("transactionType", *s.transactionType)
	if s.beginTime != nil {
		r.setParam("beginTime", *s.beginTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.rows != nil {
		r.setParam("rows", *s.rows)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(GetFiatDepositWithdrawHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetFiatDepositWithdrawHistoryResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		OrderNo         string `json:"orderNo"`
		FiatCurrency    string `json:"fiatCurrency"`
		IndicatedAmount string `json:"indicatedAmount"`
		Amount          string `json:"amount"`
		TotalFee        string `json:"totalFee"`
		Method          string `json:"method"`
		Status          string `json:"status"`
		CreateTime      uint64 `json:"createTime"`
		UpdateTime      uint64 `json:"updateTime"`
	}
	Total   int64 `json:"total"`
	Success bool  `json:"success"`
}

const (
	fiatPaymentHistory = "/sapi/v1/fiat/payments"
)

type GetFiatPaymentHistoryService struct {
	c               *Client
	transactionType *string
	beginTime       *uint64
	endTime         *uint64
	page            *int
	rows            *int
}

func (s *GetFiatPaymentHistoryService) TransactionType(transactionType string) *GetFiatPaymentHistoryService {
	s.transactionType = &transactionType
	return s
}

func (s *GetFiatPaymentHistoryService) BeginTime(beginTime uint64) *GetFiatPaymentHistoryService {
	s.beginTime = &beginTime
	return s
}

func (s *GetFiatPaymentHistoryService) EndTime(endTime uint64) *GetFiatPaymentHistoryService {
	s.endTime = &endTime
	return s
}

func (s *GetFiatPaymentHistoryService) Page(page int) *GetFiatPaymentHistoryService {
	s.page = &page
	return s
}

func (s *GetFiatPaymentHistoryService) Rows(rows int) *GetFiatPaymentHistoryService {
	s.rows = &rows
	return s
}

func (s *GetFiatPaymentHistoryService) Do(ctx context.Context) (res *GetFiatPaymentHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fiatPaymentHistory,
		secType:  secTypeSigned,
	}
	r.setParam("transactionType", *s.transactionType)
	if s.beginTime != nil {
		r.setParam("beginTime", *s.beginTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.rows != nil {
		r.setParam("rows", *s.rows)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(GetFiatPaymentHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetFiatPaymentHistoryResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		OrderNo        string `json:"orderNo"`
		SourceAmount   string `json:"sourceAmount"`
		FiatCurrency   string `json:"fiatCurrency"`
		ObtainAmount   string `json:"obtainAmount"`
		CryptoCurrency string `json:"cryptoCurrency"`
		TotalFee       string `json:"totalFee"`
		Price          string `json:"price"`
		Status         string `json:"status"`
		PaymentMethod  string `json:"paymentMethod"`
		CreateTime     uint64 `json:"createTime"`
		UpdateTime     uint64 `json:"updateTime"`
	}
	Total   int64 `json:"total"`
	Success bool  `json:"success"`
}
