package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

// Get Staking Product List (USER_DATA)
type GetStakingProductListService struct {
	c       *Client
	product string
	asset   *string
	current *int64
	size    *int64
}

// Product set product
func (s *GetStakingProductListService) Product(product string) *GetStakingProductListService {
	s.product = product
	return s
}

// Asset set asset
func (s *GetStakingProductListService) Asset(asset string) *GetStakingProductListService {
	s.asset = &asset
	return s
}

// Current set current
func (s *GetStakingProductListService) Current(current int64) *GetStakingProductListService {
	s.current = &current
	return s
}

// Size set size
func (s *GetStakingProductListService) Size(size int64) *GetStakingProductListService {
	s.size = &size
	return s
}

// Do send request
func (s *GetStakingProductListService) Do(ctx context.Context, opts ...RequestOption) (res []*StakingProductListResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/staking/productList",
		secType:  secTypeSigned,
	}
	r.setParam("product", s.product)
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*StakingProductListResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// StakingProductListResponse define staking product list response
type StakingProductListResponse struct {
	ProjectId string `json:"projectId"`
	Detail    struct {
		Asset       string `json:"asset"`
		RewardAsset string `json:"rewardAsset"`
		Duration    int64  `json:"duration"`
		Renewable   bool   `json:"renewable"`
		Apy         string `json:"apy"`
	} `json:"detail"`
	Quota struct {
		TotalPersonalQuota string `json:"totalPersonalQuota"`
		Minimum            string `json:"minimum"`
	} `json:"quota"`
}

// Purchase Staking Product(USER_DATA)
type PurchaseStakingProductService struct {
	c         *Client
	product   string
	productId string
	amount    float64
	renewable *string
}

// Product set product
func (s *PurchaseStakingProductService) Product(product string) *PurchaseStakingProductService {
	s.product = product
	return s
}

// ProductId set productId
func (s *PurchaseStakingProductService) ProductId(productId string) *PurchaseStakingProductService {
	s.productId = productId
	return s
}

// Amount set amount
func (s *PurchaseStakingProductService) Amount(amount float64) *PurchaseStakingProductService {
	s.amount = amount
	return s
}

// Renewable set renewable
func (s *PurchaseStakingProductService) Renewable(renewable string) *PurchaseStakingProductService {
	s.renewable = &renewable
	return s
}

// Do send request
func (s *PurchaseStakingProductService) Do(ctx context.Context, opts ...RequestOption) (res *PurchaseStakingProductResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/staking/purchase",
		secType:  secTypeSigned,
	}
	r.setParam("product", s.product)
	r.setParam("productId", s.productId)
	r.setParam("amount", s.amount)
	if s.renewable != nil {
		r.setParam("renewable", *s.renewable)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(PurchaseStakingProductResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PurchaseStakingProductResponse define purchase staking product response
type PurchaseStakingProductResponse struct {
	PositionId string `json:"positionId"`
	Success    bool   `json:"success"`
}

// Redeem Staking Product(USER_DATA)
type RedeemStakingProductService struct {
	c          *Client
	product    string
	positionId *string
	productId  string
	amount     *float64
}

// Product set product
func (s *RedeemStakingProductService) Product(product string) *RedeemStakingProductService {
	s.product = product
	return s
}

// PositionId set positionId
func (s *RedeemStakingProductService) PositionId(positionId string) *RedeemStakingProductService {
	s.positionId = &positionId
	return s
}

// ProductId set productId
func (s *RedeemStakingProductService) ProductId(productId string) *RedeemStakingProductService {
	s.productId = productId
	return s
}

// Amount set amount
func (s *RedeemStakingProductService) Amount(amount float64) *RedeemStakingProductService {
	s.amount = &amount
	return s
}

// Do send request
func (s *RedeemStakingProductService) Do(ctx context.Context, opts ...RequestOption) (res *RedeemStakingProductResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/staking/redeem",
		secType:  secTypeSigned,
	}
	r.setParam("product", s.product)
	if s.positionId != nil {
		r.setParam("positionId", *s.positionId)
	}
	r.setParam("productId", s.productId)
	if s.amount != nil {
		r.setParam("amount", *s.amount)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(RedeemStakingProductResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RedeemStakingProductResponse define redeem staking product response
type RedeemStakingProductResponse struct {
	Success bool `json:"success"`
}

// Get Staking Product Position(USER_DATA)
type GetStakingProductPositionService struct {
	c         *Client
	product   string
	productId *string
	asset     *string
	current   *int64
	size      *int64
}

// Product set product
func (s *GetStakingProductPositionService) Product(product string) *GetStakingProductPositionService {
	s.product = product
	return s
}

// ProductId set productId
func (s *GetStakingProductPositionService) ProductId(productId string) *GetStakingProductPositionService {
	s.productId = &productId
	return s
}

// Asset set asset
func (s *GetStakingProductPositionService) Asset(asset string) *GetStakingProductPositionService {
	s.asset = &asset
	return s
}

// Current set current
func (s *GetStakingProductPositionService) Current(current int64) *GetStakingProductPositionService {
	s.current = &current
	return s
}

// Size set size
func (s *GetStakingProductPositionService) Size(size int64) *GetStakingProductPositionService {
	s.size = &size
	return s
}

// Do send request
func (s *GetStakingProductPositionService) Do(ctx context.Context, opts ...RequestOption) (res []*GetStakingProductPositionResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/staking/position",
		secType:  secTypeSigned,
	}
	r.setParam("product", s.product)
	if s.productId != nil {
		r.setParam("productId", *s.productId)
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*GetStakingProductPositionResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetStakingProductPositionResponse define get staking product position response
type GetStakingProductPositionResponse struct {
	PositionId            string `json:"positionId"`
	ProjectId             string `json:"projectId"`
	Asset                 string `json:"asset"`
	Amount                string `json:"amount"`
	PurchaseTime          string `json:"purchaseTime"`
	Duration              string `json:"duration"`
	AccrualDays           string `json:"accrualDays"`
	RewardAsset           string `json:"rewardAsset"`
	APY                   string `json:"APY"`
	RewardAmt             string `json:"rewardAmt"`
	ExtraRewardAsset      string `json:"extraRewardAsset"`
	ExtraRewardAPY        string `json:"extraRewardAPY"`
	EstExtraRewardAmt     string `json:"estExtraRewardAmt"`
	NextInterestPay       string `json:"nextInterestPay"`
	NextInterestPayDate   string `json:"nextInterestPayDate"`
	PayInterestPeriod     string `json:"payInterestPeriod"`
	RedeemAmountEarly     string `json:"redeemAmountEarly"`
	InterestEndDate       string `json:"interestEndDate"`
	DeliverDate           string `json:"deliverDate"`
	RedeemPeriod          string `json:"redeemPeriod"`
	RedeemingAmt          string `json:"redeemingAmt"`
	PartialAmtDeliverDate string `json:"partialAmtDeliverDate"`
	CanRedeemEarly        bool   `json:"canRedeemEarly"`
	Renewable             bool   `json:"renewable"`
	Type                  string `json:"type"`
	Status                string `json:"status"`
}

// Get Staking History(USER_DATA)
type GetStakingHistoryService struct {
	c         *Client
	product   string
	txnType   string
	asset     *string
	startTime *uint64
	endTime   *uint64
	current   *int64
	size      *int64
}

// Product set product
func (s *GetStakingHistoryService) Product(product string) *GetStakingHistoryService {
	s.product = product
	return s
}

// TxnType set txnType
func (s *GetStakingHistoryService) TxnType(txnType string) *GetStakingHistoryService {
	s.txnType = txnType
	return s
}

// Asset set asset
func (s *GetStakingHistoryService) Asset(asset string) *GetStakingHistoryService {
	s.asset = &asset
	return s
}

// StartTime set startTime
func (s *GetStakingHistoryService) StartTime(startTime uint64) *GetStakingHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetStakingHistoryService) EndTime(endTime uint64) *GetStakingHistoryService {
	s.endTime = &endTime
	return s
}

// Current set current
func (s *GetStakingHistoryService) Current(current int64) *GetStakingHistoryService {
	s.current = &current
	return s
}

// Size set size
func (s *GetStakingHistoryService) Size(size int64) *GetStakingHistoryService {
	s.size = &size
	return s
}

// Do send request
func (s *GetStakingHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*GetStakingHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/staking/stakingRecord",
		secType:  secTypeSigned,
	}
	r.setParam("product", s.product)
	r.setParam("type", s.txnType)
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*GetStakingHistoryResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetStakingHistoryResponse define get staking history response
type GetStakingHistoryResponse struct {
	PositionId  string `json:"positionId"`
	Time        uint64 `json:"time"`
	Asset       string `json:"asset"`
	Project     string `json:"project"`
	Amount      string `json:"amount"`
	LockPeriod  string `json:"lockPeriod"`
	DeliverDate string `json:"deliverDate"`
	Type        string `json:"type"`
	Status      string `json:"status"`
}

// Set Auto Staking(USER_DATA)
type SetAutoStakingService struct {
	c          *Client
	product    string
	positionId string
	renewable  string
}

// Product set product
func (s *SetAutoStakingService) Product(product string) *SetAutoStakingService {
	s.product = product
	return s
}

// PositionId set positionId
func (s *SetAutoStakingService) PositionId(positionId string) *SetAutoStakingService {
	s.positionId = positionId
	return s
}

// Renewable set renewable
func (s *SetAutoStakingService) Renewable(renewable string) *SetAutoStakingService {
	s.renewable = renewable
	return s
}

// Do send request
func (s *SetAutoStakingService) Do(ctx context.Context, opts ...RequestOption) (res *SetAutoStakingResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/staking/setAutoStaking",
		secType:  secTypeSigned,
	}
	r.setParam("product", s.product)
	r.setParam("positionId", s.positionId)
	r.setParam("renewable", s.renewable)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SetAutoStakingResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SetAutoStakingResponse define set auto staking response
type SetAutoStakingResponse struct {
	Success bool `json:"success"`
}

// Get Personal Left Quota of Staking Product(USER_DATA)
type PersonalLeftQuotaService struct {
	c         *Client
	product   string
	productId string
}

// Product set product
func (s *PersonalLeftQuotaService) Product(product string) *PersonalLeftQuotaService {
	s.product = product
	return s
}

// ProductId set productId
func (s *PersonalLeftQuotaService) ProductId(productId string) *PersonalLeftQuotaService {
	s.productId = productId
	return s
}

// Do send request
func (s *PersonalLeftQuotaService) Do(ctx context.Context, opts ...RequestOption) (res []*PersonalLeftQuotaResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/staking/personalLeftQuota",
		secType:  secTypeSigned,
	}
	r.setParam("product", s.product)
	r.setParam("productId", s.productId)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*PersonalLeftQuotaResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PersonalLeftQuotaResponse define get staking asset response
type PersonalLeftQuotaResponse struct {
	LeftPersonalQuota string `json:"leftPersonalQuota"`
}
