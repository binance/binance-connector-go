package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// Create a Virtual Sub-account(For Master Account)
const (
	enableSubAccountEndpoint = "/sapi/v1/sub-account/virtualSubAccount"
)

type CreateSubAccountService struct {
	c                *Client
	subAccountString string
}

func (s *CreateSubAccountService) SubAccountString(subAccountString string) *CreateSubAccountService {
	s.subAccountString = subAccountString
	return s
}

func (s *CreateSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *CreateSubAccountResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: enableSubAccountEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("subAccountString", s.subAccountString)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateSubAccountResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type CreateSubAccountResp struct {
	Email string `json:"email"`
}

// Query Sub-account List (For Master Account)
const (
	querySubAccountListEndpoint = "/sapi/v1/sub-account/list"
)

type QuerySubAccountListService struct {
	c        *Client
	email    *string
	isFreeze *string
	page     *int
	limit    *int
}

func (s *QuerySubAccountListService) Email(email string) *QuerySubAccountListService {
	s.email = &email
	return s
}

func (s *QuerySubAccountListService) IsFreeze(isFreeze string) *QuerySubAccountListService {
	s.isFreeze = &isFreeze
	return s
}

func (s *QuerySubAccountListService) Page(page int) *QuerySubAccountListService {
	s.page = &page
	return s
}

func (s *QuerySubAccountListService) Limit(limit int) *QuerySubAccountListService {
	s.limit = &limit
	return s
}

func (s *QuerySubAccountListService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountListResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: querySubAccountListEndpoint,
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", s.email)
	}
	if s.isFreeze != nil {
		r.setParam("isFreeze", s.isFreeze)
	}
	if s.page != nil {
		r.setParam("page", s.page)
	}
	if s.limit != nil {
		r.setParam("limit", s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountListResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountListResp struct {
	SubAccounts []SubAccount `json:"subAccounts"`
}

type SubAccount struct {
	Email                       string `json:"email"`
	IsFreeze                    bool   `json:"isFreeze"`
	CreateTime                  uint64 `json:"createTime"`
	IsManagedSubAccount         bool   `json:"isManagedSubAccount"`
	IsAssetManagementSubAccount bool   `json:"isAssetManagementSubAccount"`
}

// Query Sub-account Spot Asset Transfer History (For Master Account)
const (
	querySubAccountSpotAssetTransferHistoryEndpoint = "/sapi/v1/sub-account/sub/transfer/history"
)

type QuerySubAccountSpotAssetTransferHistoryService struct {
	c         *Client
	fromEmail *string
	toEmail   *string
	startTime *uint64
	endTime   *uint64
	page      *int
	limit     *int
}

func (s *QuerySubAccountSpotAssetTransferHistoryService) FromEmail(fromEmail string) *QuerySubAccountSpotAssetTransferHistoryService {
	s.fromEmail = &fromEmail
	return s
}

func (s *QuerySubAccountSpotAssetTransferHistoryService) ToEmail(toEmail string) *QuerySubAccountSpotAssetTransferHistoryService {
	s.toEmail = &toEmail
	return s
}

func (s *QuerySubAccountSpotAssetTransferHistoryService) StartTime(startTime uint64) *QuerySubAccountSpotAssetTransferHistoryService {
	s.startTime = &startTime
	return s
}

func (s *QuerySubAccountSpotAssetTransferHistoryService) EndTime(endTime uint64) *QuerySubAccountSpotAssetTransferHistoryService {
	s.endTime = &endTime
	return s
}

func (s *QuerySubAccountSpotAssetTransferHistoryService) Page(page int) *QuerySubAccountSpotAssetTransferHistoryService {
	s.page = &page
	return s
}

func (s *QuerySubAccountSpotAssetTransferHistoryService) Limit(limit int) *QuerySubAccountSpotAssetTransferHistoryService {
	s.limit = &limit
	return s
}

func (s *QuerySubAccountSpotAssetTransferHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *QuerySubAccountSpotAssetTransferHistoryResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: querySubAccountSpotAssetTransferHistoryEndpoint,
		secType:  secTypeSigned,
	}
	if s.fromEmail != nil {
		r.setParam("fromEmail", s.fromEmail)
	}
	if s.toEmail != nil {
		r.setParam("toEmail", s.toEmail)
	}
	if s.startTime != nil {
		r.setParam("startTime", s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", s.endTime)
	}
	if s.page != nil {
		r.setParam("page", s.page)
	}
	if s.limit != nil {
		r.setParam("limit", s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QuerySubAccountSpotAssetTransferHistoryResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type QuerySubAccountSpotAssetTransferHistoryResp struct {
	Rows []struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Asset  string `json:"asset"`
		Qty    string `json:"qty"`
		Status string `json:"status"`
		TranId int64  `json:"tranId"`
		Time   uint64 `json:"time"`
	}
}

// Query Sub-account Futures Asset Transfer History (For Master Account)
const (
	querySubAccountFuturesAssetTransferHistoryEndpoint = "/sapi/v1/sub-account/futures/internalTransfer"
)

type QuerySubAccountFuturesAssetTransferHistoryService struct {
	c           *Client
	email       string
	futuresType int
	startTime   *uint64
	endTime     *uint64
	page        *int
	limit       *int
}

func (s *QuerySubAccountFuturesAssetTransferHistoryService) Email(email string) *QuerySubAccountFuturesAssetTransferHistoryService {
	s.email = email
	return s
}

func (s *QuerySubAccountFuturesAssetTransferHistoryService) FuturesType(futuresType int) *QuerySubAccountFuturesAssetTransferHistoryService {
	s.futuresType = futuresType
	return s
}

func (s *QuerySubAccountFuturesAssetTransferHistoryService) StartTime(startTime uint64) *QuerySubAccountFuturesAssetTransferHistoryService {
	s.startTime = &startTime
	return s
}

func (s *QuerySubAccountFuturesAssetTransferHistoryService) EndTime(endTime uint64) *QuerySubAccountFuturesAssetTransferHistoryService {
	s.endTime = &endTime
	return s
}

func (s *QuerySubAccountFuturesAssetTransferHistoryService) Page(page int) *QuerySubAccountFuturesAssetTransferHistoryService {
	s.page = &page
	return s
}

func (s *QuerySubAccountFuturesAssetTransferHistoryService) Limit(limit int) *QuerySubAccountFuturesAssetTransferHistoryService {
	s.limit = &limit
	return s
}

func (s *QuerySubAccountFuturesAssetTransferHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *QuerySubAccountFuturesAssetTransferHistoryResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: querySubAccountFuturesAssetTransferHistoryEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("futuresType", s.futuresType)
	if s.startTime != nil {
		r.setParam("startTime", s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", s.endTime)
	}
	if s.page != nil {
		r.setParam("page", s.page)
	}
	if s.limit != nil {
		r.setParam("limit", s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QuerySubAccountFuturesAssetTransferHistoryResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type QuerySubAccountFuturesAssetTransferHistoryResp struct {
	Success     bool  `json:"success"`
	FuturesType int64 `json:"futuresType"`
	Transfers   []struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Asset  string `json:"asset"`
		Qty    string `json:"qty"`
		TranId int64  `json:"tranId"`
		Time   uint64 `json:"time"`
	} `json:"transfers"`
}

// Sub-account Futures Asset Transfer (For Master Account)
const (
	subAccountFuturesAssetTransferEndpoint = "/sapi/v1/sub-account/futures/internalTransfer"
)

type SubAccountFuturesAssetTransferService struct {
	c           *Client
	fromEmail   string
	toEmail     string
	futuresType int64
	asset       string
	amount      float32
}

func (s *SubAccountFuturesAssetTransferService) FromEmail(fromEmail string) *SubAccountFuturesAssetTransferService {
	s.fromEmail = fromEmail
	return s
}

func (s *SubAccountFuturesAssetTransferService) ToEmail(toEmail string) *SubAccountFuturesAssetTransferService {
	s.toEmail = toEmail
	return s
}

func (s *SubAccountFuturesAssetTransferService) FuturesType(futuresType int64) *SubAccountFuturesAssetTransferService {
	s.futuresType = futuresType
	return s
}

func (s *SubAccountFuturesAssetTransferService) Asset(asset string) *SubAccountFuturesAssetTransferService {
	s.asset = asset
	return s
}

func (s *SubAccountFuturesAssetTransferService) Amount(amount float32) *SubAccountFuturesAssetTransferService {
	s.amount = amount
	return s
}

func (s *SubAccountFuturesAssetTransferService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFuturesAssetTransferResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: subAccountFuturesAssetTransferEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("fromEmail", s.fromEmail)
	r.setParam("toEmail", s.toEmail)
	r.setParam("futuresType", s.futuresType)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountFuturesAssetTransferResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountFuturesAssetTransferResp struct {
	Success bool   `json:"success"`
	TxnId   string `json:"txnId"`
}

// Query Sub-account Assets (For Master Account)
const (
	querySubAccountAssetsEndpoint = "/sapi/v3/sub-account/assets"
)

type QuerySubAccountAssetsService struct {
	c     *Client
	email string
}

func (s *QuerySubAccountAssetsService) Email(email string) *QuerySubAccountAssetsService {
	s.email = email
	return s
}

func (s *QuerySubAccountAssetsService) Do(ctx context.Context, opts ...RequestOption) (res *QuerySubAccountAssetsResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: querySubAccountAssetsEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QuerySubAccountAssetsResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type QuerySubAccountAssetsResp struct {
	Balances []struct {
		Asset  string `json:"asset"`
		Free   string `json:"free"`
		Locked string `json:"locked"`
	} `json:"balances"`
}

// Query Sub-account Spot Assets Summary (For Master Account)
const (
	querySubAccountSpotAssetsSummaryEndpoint = "/sapi/v1/sub-account/spotSummary"
)

type QuerySubAccountSpotAssetsSummaryService struct {
	c     *Client
	email *string
	page  *int
	size  *int
}

func (s *QuerySubAccountSpotAssetsSummaryService) Email(email string) *QuerySubAccountSpotAssetsSummaryService {
	s.email = &email
	return s
}

func (s *QuerySubAccountSpotAssetsSummaryService) Page(page int) *QuerySubAccountSpotAssetsSummaryService {
	s.page = &page
	return s
}

func (s *QuerySubAccountSpotAssetsSummaryService) Size(size int) *QuerySubAccountSpotAssetsSummaryService {
	s.size = &size
	return s
}

func (s *QuerySubAccountSpotAssetsSummaryService) Do(ctx context.Context, opts ...RequestOption) (res *QuerySubAccountSpotAssetsSummaryResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: querySubAccountSpotAssetsSummaryEndpoint,
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", *s.email)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QuerySubAccountSpotAssetsSummaryResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type QuerySubAccountSpotAssetsSummaryResp struct {
	TotalCount                int64  `json:"totalCount"`
	MasterAccountTotalAsset   string `json:"masterAccountTotalAsset"`
	SpotSubUserAssetBtcVoList []struct {
		Email   string `json:"email"`
		ToAsset string `json:"toAsset"`
	} `json:"spotSubUserAssetBtcVoList"`
}

// Get Sub-account Deposit Address (For Master Account)
const (
	getSubAccountDepositAddressEndpoint = "/sapi/v1/capital/deposit/subAddress"
)

type GetSubAccountDepositAddressService struct {
	c       *Client
	email   string
	coin    string
	network *string
}

func (s *GetSubAccountDepositAddressService) Email(email string) *GetSubAccountDepositAddressService {
	s.email = email
	return s
}

func (s *GetSubAccountDepositAddressService) Coin(coin string) *GetSubAccountDepositAddressService {
	s.coin = coin
	return s
}

func (s *GetSubAccountDepositAddressService) Network(network string) *GetSubAccountDepositAddressService {
	s.network = &network
	return s
}

func (s *GetSubAccountDepositAddressService) Do(ctx context.Context, opts ...RequestOption) (res *GetSubAccountDepositAddressResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getSubAccountDepositAddressEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("coin", s.coin)
	if s.network != nil {
		r.setParam("network", *s.network)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetSubAccountDepositAddressResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetSubAccountDepositAddressResp struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}

// Get Sub-account Deposit History (For Master Account)
const (
	getSubAccountDepositHistoryEndpoint = "/sapi/v1/capital/deposit/subHisrec"
)

type GetSubAccountDepositHistoryService struct {
	c         *Client
	email     string
	coin      string
	status    *int64
	startTime *uint64
	endTime   *uint64
	limit     *int
	offset    *int64
}

func (s *GetSubAccountDepositHistoryService) Email(email string) *GetSubAccountDepositHistoryService {
	s.email = email
	return s
}

func (s *GetSubAccountDepositHistoryService) Coin(coin string) *GetSubAccountDepositHistoryService {
	s.coin = coin
	return s
}

func (s *GetSubAccountDepositHistoryService) Status(status int64) *GetSubAccountDepositHistoryService {
	s.status = &status
	return s
}

func (s *GetSubAccountDepositHistoryService) StartTime(startTime uint64) *GetSubAccountDepositHistoryService {
	s.startTime = &startTime
	return s
}

func (s *GetSubAccountDepositHistoryService) EndTime(endTime uint64) *GetSubAccountDepositHistoryService {
	s.endTime = &endTime
	return s
}

func (s *GetSubAccountDepositHistoryService) Limit(limit int) *GetSubAccountDepositHistoryService {
	s.limit = &limit
	return s
}

func (s *GetSubAccountDepositHistoryService) Offset(offset int64) *GetSubAccountDepositHistoryService {
	s.offset = &offset
	return s
}

func (s *GetSubAccountDepositHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *GetSubAccountDepositHistoryResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getSubAccountDepositHistoryEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("coin", s.coin)
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetSubAccountDepositHistoryResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetSubAccountDepositHistoryResp struct {
	DepositList []struct {
		Id            int64  `json:"id"`
		Amount        string `json:"amount"`
		Coin          string `json:"coin"`
		Network       string `json:"network"`
		Status        int64  `json:"status"`
		Address       string `json:"address"`
		AddressTag    string `json:"addressTag"`
		TxId          string `json:"txId"`
		InsertTime    uint64 `json:"insertTime"`
		TransferType  int64  `json:"transferType"`
		ConfirmTimes  string `json:"confirmTimes"`
		UnlockConfirm int64  `json:"unlockConfirm"`
		WalletType    int    `json:"walletType"`
	}
}

// Get Sub-account's Status on Margin/Futures (For Master Account)
const (
	getSubAccountStatusEndpoint = "/sapi/v1/sub-account/status"
)

type GetSubAccountStatusService struct {
	c     *Client
	email *string
}

func (s *GetSubAccountStatusService) Email(email string) *GetSubAccountStatusService {
	s.email = &email
	return s
}

func (s *GetSubAccountStatusService) Do(ctx context.Context, opts ...RequestOption) (res *GetSubAccountStatusResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getSubAccountStatusEndpoint,
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", *s.email)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetSubAccountStatusResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetSubAccountStatusResp struct {
	Email            string `json:"email"`
	IsSubUserEnabled bool   `json:"isSubUserEnabled"`
	IsUserActive     bool   `json:"isUserActive"`
	InsertTime       uint64 `json:"insertTime"`
	IsMarginEnabled  bool   `json:"isMarginEnabled"`
	IsFuturesEnabled bool   `json:"isFuturesEnabled"`
	Mobile           int64  `json:"mobile"`
}

// Enable Margin for Sub-account (For Master Account)
const (
	enableMarginForSubAccountEndpoint = "/sapi/v1/sub-account/margin/enable"
)

type EnableMarginForSubAccountService struct {
	c     *Client
	email string
}

func (s *EnableMarginForSubAccountService) Email(email string) *EnableMarginForSubAccountService {
	s.email = email
	return s
}

func (s *EnableMarginForSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *EnableMarginForSubAccountResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: enableMarginForSubAccountEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(EnableMarginForSubAccountResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type EnableMarginForSubAccountResp struct {
	Email           string `json:"email"`
	IsMarginEnabled bool   `json:"isMarginEnabled"`
}

// Get Detail on Sub-account's Margin Account (For Master Account)
const (
	getDetailOnSubAccountMarginAccountEndpoint = "/sapi/v1/sub-account/margin/account"
)

type GetDetailOnSubAccountMarginAccountService struct {
	c     *Client
	email string
}

func (s *GetDetailOnSubAccountMarginAccountService) Email(email string) *GetDetailOnSubAccountMarginAccountService {
	s.email = email
	return s
}

func (s *GetDetailOnSubAccountMarginAccountService) Do(ctx context.Context, opts ...RequestOption) (res *GetDetailOnSubAccountMarginAccountResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getDetailOnSubAccountMarginAccountEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetDetailOnSubAccountMarginAccountResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetDetailOnSubAccountMarginAccountResp struct {
	Email               string `json:"email"`
	MarginLevel         string `json:"marginLevel"`
	TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`
	MarginTradeCoeffVo  struct {
		ForceLiquidationRate string `json:"forceLiquidationRate"`
		MarginCallBar        string `json:"marginCallBar"`
		NormalBar            string `json:"normalBar"`
	} `json:"marginTradeCoeffVo"`
	MarginUserAssetVoList []struct {
		Asset    string `json:"asset"`
		Borrowed string `json:"borrowed"`
		Free     string `json:"free"`
		Interest string `json:"interest"`
		Locked   string `json:"locked"`
		NetAsset string `json:"netAsset"`
	}
}

// Get Summary of Sub-account's Margin Account (For Master Account)
const (
	getSummaryOfSubAccountMarginAccountEndpoint = "/sapi/v1/sub-account/margin/accountSummary"
)

type GetSummaryOfSubAccountMarginAccountService struct {
	c *Client
}

func (s *GetSummaryOfSubAccountMarginAccountService) Do(ctx context.Context, opts ...RequestOption) (res *GetSummaryOfSubAccountMarginAccountResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getSummaryOfSubAccountMarginAccountEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetSummaryOfSubAccountMarginAccountResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetSummaryOfSubAccountMarginAccountResp struct {
	TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`
	SubAccountList      []struct {
		Email               string `json:"email"`
		TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
		TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
		TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`
	} `json:"subAccountList"`
}

// Enable Futures for Sub-account (For Master Account)
const (
	enableFuturesForSubAccountEndpoint = "/sapi/v1/sub-account/futures/enable"
)

type EnableFuturesForSubAccountService struct {
	c     *Client
	email string
}

func (s *EnableFuturesForSubAccountService) Email(email string) *EnableFuturesForSubAccountService {
	s.email = email
	return s
}

func (s *EnableFuturesForSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *EnableFuturesForSubAccountResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: enableFuturesForSubAccountEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(EnableFuturesForSubAccountResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type EnableFuturesForSubAccountResp struct {
	Email            string `json:"email"`
	IsFuturesEnabled bool   `json:"isFuturesEnabled"`
}

// Get Detail on Sub-account's Futures Account (For Master Account)
const (
	getDetailOnSubAccountFuturesAccountEndpoint = "/sapi/v1/sub-account/futures/account"
)

type GetDetailOnSubAccountFuturesAccountService struct {
	c     *Client
	email string
}

func (s *GetDetailOnSubAccountFuturesAccountService) Email(email string) *GetDetailOnSubAccountFuturesAccountService {
	s.email = email
	return s
}

func (s *GetDetailOnSubAccountFuturesAccountService) Do(ctx context.Context, opts ...RequestOption) (res *GetDetailOnSubAccountFuturesAccountResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getDetailOnSubAccountFuturesAccountEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetDetailOnSubAccountFuturesAccountResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetDetailOnSubAccountFuturesAccountResp struct {
	Email  string `json:"email"`
	Asset  string `json:"asset"`
	Assets []struct {
		Asset                  string `json:"asset"`
		InitialMargin          string `json:"initialMargin"`
		MaintenanceMargin      string `json:"maintenanceMargin"`
		MarginBalance          string `json:"marginBalance"`
		MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
		PositionInitialMargin  string `json:"positionInitialMargin"`
		UnrealizedProfit       string `json:"unrealizedProfit"`
		WalletBalance          string `json:"walletBalance"`
	} `json:"assets"`
	CanDeposit                  bool   `json:"canDeposit"`
	CanTrade                    bool   `json:"canTrade"`
	CanWithdraw                 bool   `json:"canWithdraw"`
	FeeTier                     int    `json:"feeTier"`
	MaxWithdrawAmount           string `json:"maxWithdrawAmount"`
	TotalInitialMargin          string `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string `json:"totalWalletBalance"`
	UpdateTime                  uint64 `json:"updateTime"`
}

// Get Summary of Sub-account's Futures Account (For Master Account)
const (
	getSummaryOfSubAccountFuturesAccountEndpoint = "/sapi/v1/sub-account/futures/accountSummary"
)

type GetSummaryOfSubAccountFuturesAccountService struct {
	c *Client
}

func (s *GetSummaryOfSubAccountFuturesAccountService) Do(ctx context.Context, opts ...RequestOption) (res *GetSummaryOfSubAccountFuturesAccountResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getSummaryOfSubAccountFuturesAccountEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetSummaryOfSubAccountFuturesAccountResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetSummaryOfSubAccountFuturesAccountResp struct {
	TotalInitialMargin          string `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string `json:"totalWalletBalance"`
	Asset                       string `json:"asset"`
	SubAccountList              []struct {
		Email                       string `json:"email"`
		TotalInitialMargin          string `json:"totalInitialMargin"`
		TotalMaintenanceMargin      string `json:"totalMaintenanceMargin"`
		TotalMarginBalance          string `json:"totalMarginBalance"`
		TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
		TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
		TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
		TotalWalletBalance          string `json:"totalWalletBalance"`
		Asset                       string `json:"asset"`
	} `json:"subAccountList"`
}

// Get Futures Position-Risk of Sub-account (For Master Account)
const (
	getFuturesPositionRiskOfSubAccountEndpoint = "/sapi/v1/sub-account/futures/positionRisk"
)

type GetFuturesPositionRiskOfSubAccountService struct {
	c     *Client
	email string
}

func (s *GetFuturesPositionRiskOfSubAccountService) Email(email string) *GetFuturesPositionRiskOfSubAccountService {
	s.email = email
	return s
}

func (s *GetFuturesPositionRiskOfSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *GetFuturesPositionRiskOfSubAccountResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getFuturesPositionRiskOfSubAccountEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetFuturesPositionRiskOfSubAccountResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetFuturesPositionRiskOfSubAccountResp struct {
	EntryPrice       string `json:"entryPrice"`
	Leverage         string `json:"leverage"`
	MaxNotional      string `json:"maxNotional"`
	LiquidationPrice string `json:"liquidationPrice"`
	MarkPrice        string `json:"markPrice"`
	PositionAmount   string `json:"positionAmount"`
	Symbol           string `json:"symbol"`
	UnrealizedProfit string `json:"unrealizedProfit"`
}

// Futures Transfer for Sub-account (For Master Account)
const (
	futuresTransferForSubAccountEndpoint = "/sapi/v1/sub-account/futures/transfer"
)

type FuturesTransferForSubAccountService struct {
	c            *Client
	email        string
	asset        string
	amount       float64
	transferType int
}

func (s *FuturesTransferForSubAccountService) Email(email string) *FuturesTransferForSubAccountService {
	s.email = email
	return s
}

func (s *FuturesTransferForSubAccountService) Asset(asset string) *FuturesTransferForSubAccountService {
	s.asset = asset
	return s
}

func (s *FuturesTransferForSubAccountService) Amount(amount float64) *FuturesTransferForSubAccountService {
	s.amount = amount
	return s
}

func (s *FuturesTransferForSubAccountService) TransferType(transferType int) *FuturesTransferForSubAccountService {
	s.transferType = transferType
	return s
}

func (s *FuturesTransferForSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *FuturesTransferForSubAccountResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: futuresTransferForSubAccountEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	r.setParam("type", s.transferType)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(FuturesTransferForSubAccountResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type FuturesTransferForSubAccountResp struct {
	TxnId int `json:"txnId"`
}

// Margin Transfer for Sub-account (For Master Account)
const (
	marginTransferForSubAccountEndpoint = "/sapi/v1/sub-account/margin/transfer"
)

type MarginTransferForSubAccountService struct {
	c            *Client
	email        string
	asset        string
	amount       float32
	transferType int
}

func (s *MarginTransferForSubAccountService) Email(email string) *MarginTransferForSubAccountService {
	s.email = email
	return s
}

func (s *MarginTransferForSubAccountService) Asset(asset string) *MarginTransferForSubAccountService {
	s.asset = asset
	return s
}

func (s *MarginTransferForSubAccountService) Amount(amount float32) *MarginTransferForSubAccountService {
	s.amount = amount
	return s
}

func (s *MarginTransferForSubAccountService) TransferType(transferType int) *MarginTransferForSubAccountService {
	s.transferType = transferType
	return s
}

func (s *MarginTransferForSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *MarginTransferForSubAccountResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: marginTransferForSubAccountEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	r.setParam("type", s.transferType)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginTransferForSubAccountResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type MarginTransferForSubAccountResp struct {
	TxnId int `json:"txnId"`
}

// Transfer to Sub-account of Same Master (For Sub-account)
const (
	transferToSubAccountOfSameMasterEndpoint = "/sapi/v1/sub-account/transfer/subToSub"
)

type TransferToSubAccountOfSameMasterService struct {
	c       *Client
	toEmail string
	asset   string
	amount  float64
}

func (s *TransferToSubAccountOfSameMasterService) ToEmail(toEmail string) *TransferToSubAccountOfSameMasterService {
	s.toEmail = toEmail
	return s
}

func (s *TransferToSubAccountOfSameMasterService) Asset(asset string) *TransferToSubAccountOfSameMasterService {
	s.asset = asset
	return s
}

func (s *TransferToSubAccountOfSameMasterService) Amount(amount float64) *TransferToSubAccountOfSameMasterService {
	s.amount = amount
	return s
}

func (s *TransferToSubAccountOfSameMasterService) Do(ctx context.Context, opts ...RequestOption) (res *TransferToSubAccountOfSameMasterResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: transferToSubAccountOfSameMasterEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("toEmail", s.toEmail)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(TransferToSubAccountOfSameMasterResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type TransferToSubAccountOfSameMasterResp struct {
	TxnId int `json:"txnId"`
}

// Transfer to Master (For Sub-account)
const (
	transferToMasterEndpoint = "/sapi/v1/sub-account/transfer/subToMaster"
)

type TransferToMasterService struct {
	c      *Client
	asset  string
	amount float64
}

func (s *TransferToMasterService) Asset(asset string) *TransferToMasterService {
	s.asset = asset
	return s
}

func (s *TransferToMasterService) Amount(amount float64) *TransferToMasterService {
	s.amount = amount
	return s
}

func (s *TransferToMasterService) Do(ctx context.Context, opts ...RequestOption) (res *TransferToMasterResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: transferToMasterEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(TransferToMasterResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type TransferToMasterResp struct {
	TxnId int `json:"txnId"`
}

// Sub-account Transfer History (For Sub-account)
const (
	subAccountTransferHistoryEndpoint = "/sapi/v1/sub-account/transfer/subUserHistory"
)

type SubAccountTransferHistoryService struct {
	c            *Client
	asset        *string
	transferType *int
	startTime    *uint64
	endTime      *uint64
	limit        *int
}

func (s *SubAccountTransferHistoryService) Asset(asset string) *SubAccountTransferHistoryService {
	s.asset = &asset
	return s
}

func (s *SubAccountTransferHistoryService) TransferType(transferType int) *SubAccountTransferHistoryService {
	s.transferType = &transferType
	return s
}

func (s *SubAccountTransferHistoryService) StartTime(startTime uint64) *SubAccountTransferHistoryService {
	s.startTime = &startTime
	return s
}

func (s *SubAccountTransferHistoryService) EndTime(endTime uint64) *SubAccountTransferHistoryService {
	s.endTime = &endTime
	return s
}

func (s *SubAccountTransferHistoryService) Limit(limit int) *SubAccountTransferHistoryService {
	s.limit = &limit
	return s
}

func (s *SubAccountTransferHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountTransferHistoryResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: subAccountTransferHistoryEndpoint,
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.transferType != nil {
		r.setParam("type", *s.transferType)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountTransferHistoryResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountTransferHistoryResp struct {
	CounterParty    string `json:"counterParty"`
	Email           string `json:"email"`
	Type            int    `json:"type"`
	Asset           string `json:"asset"`
	Qty             string `json:"qty"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Status          string `json:"status"`
	TranId          int    `json:"tranId"`
	Time            uint64 `json:"time"`
}

// Universal Transfer (For Master Account)
const (
	universalTransferEndpoint = "/sapi/v1/asset/universalTransfer"
)

type UniversalTransferService struct {
	c               *Client
	fromEmail       *string
	toEmail         *string
	fromAccountType string
	toAccountType   string
	clientTranId    *string
	symbol          *string
	asset           string
	amount          float64
}

func (s *UniversalTransferService) FromEmail(fromEmail string) *UniversalTransferService {
	s.fromEmail = &fromEmail
	return s
}

func (s *UniversalTransferService) ToEmail(toEmail string) *UniversalTransferService {
	s.toEmail = &toEmail
	return s
}

func (s *UniversalTransferService) FromAccountType(fromAccountType string) *UniversalTransferService {
	s.fromAccountType = fromAccountType
	return s
}

func (s *UniversalTransferService) ToAccountType(toAccountType string) *UniversalTransferService {
	s.toAccountType = toAccountType
	return s
}

func (s *UniversalTransferService) ClientTranId(clientTranId string) *UniversalTransferService {
	s.clientTranId = &clientTranId
	return s
}

func (s *UniversalTransferService) Symbol(symbol string) *UniversalTransferService {
	s.symbol = &symbol
	return s
}

func (s *UniversalTransferService) Asset(asset string) *UniversalTransferService {
	s.asset = asset
	return s
}

func (s *UniversalTransferService) Amount(amount float64) *UniversalTransferService {
	s.amount = amount
	return s
}

func (s *UniversalTransferService) Do(ctx context.Context, opts ...RequestOption) (res *UniversalTransferResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: universalTransferEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("fromAccountType", s.fromAccountType)
	r.setParam("toAccountType", s.toAccountType)
	if s.fromEmail != nil {
		r.setParam("fromEmail", *s.fromEmail)
	}
	if s.toEmail != nil {
		r.setParam("toEmail", *s.toEmail)
	}
	if s.clientTranId != nil {
		r.setParam("clientTranId", *s.clientTranId)
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UniversalTransferResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type UniversalTransferResp struct {
	TranId       int    `json:"tranId"`
	ClientTranId string `json:"clientTranId"`
}

// Query Universal Transfer History (For Master Account)
const (
	queryUniversalTransferHistoryEndpoint = "/sapi/v1/asset/universalTransfer"
)

type QueryUniversalTransferHistoryService struct {
	c            *Client
	fromEmail    *string
	toEmail      *string
	clientTranId *string
	startTime    *uint64
	endTime      *uint64
	page         *int
	limit        *int
}

func (s *QueryUniversalTransferHistoryService) FromEmail(fromEmail string) *QueryUniversalTransferHistoryService {
	s.fromEmail = &fromEmail
	return s
}

func (s *QueryUniversalTransferHistoryService) ToEmail(toEmail string) *QueryUniversalTransferHistoryService {
	s.toEmail = &toEmail
	return s
}

func (s *QueryUniversalTransferHistoryService) ClientTranId(clientTranId string) *QueryUniversalTransferHistoryService {
	s.clientTranId = &clientTranId
	return s
}

func (s *QueryUniversalTransferHistoryService) StartTime(startTime uint64) *QueryUniversalTransferHistoryService {
	s.startTime = &startTime
	return s
}

func (s *QueryUniversalTransferHistoryService) EndTime(endTime uint64) *QueryUniversalTransferHistoryService {
	s.endTime = &endTime
	return s
}

func (s *QueryUniversalTransferHistoryService) Page(page int) *QueryUniversalTransferHistoryService {
	s.page = &page
	return s
}

func (s *QueryUniversalTransferHistoryService) Limit(limit int) *QueryUniversalTransferHistoryService {
	s.limit = &limit
	return s
}

func (s *QueryUniversalTransferHistoryService) Do(ctx context.Context, opts ...RequestOption) (res QueryUniversalTransferHistoryResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: queryUniversalTransferHistoryEndpoint,
		secType:  secTypeSigned,
	}
	if s.fromEmail != nil {
		r.setParam("fromEmail", *s.fromEmail)
	}
	if s.toEmail != nil {
		r.setParam("toEmail", *s.toEmail)
	}
	if s.clientTranId != nil {
		r.setParam("clientTranId", *s.clientTranId)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res.Result = make([]*InternalUniversalTransfer, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

type QueryUniversalTransferHistoryResp struct {
	Result     []*InternalUniversalTransfer `json:"result"`
	TotalCount int                          `json:"totalCount"`
}

type InternalUniversalTransfer struct {
	TranId          int64  `json:"tranId"`
	ClientTranId    string `json:"clientTranId"`
	FromEmail       string `json:"fromEmail"`
	ToEmail         string `json:"toEmail"`
	Asset           string `json:"asset"`
	Amount          string `json:"amount"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Status          string `json:"status"`
	CreateTimeStamp uint64 `json:"createTimeStamp"`
}

// Get Detail on Sub-account's Futures Account V2 (For Master Account)
const (
	getDetailOnSubAccountFuturesAccountV2Endpoint = "/sapi/v2/sub-account/futures/account"
)

type GetDetailOnSubAccountFuturesAccountV2Service struct {
	c           *Client
	email       string
	futuresType int
}

func (s *GetDetailOnSubAccountFuturesAccountV2Service) Email(email string) *GetDetailOnSubAccountFuturesAccountV2Service {
	s.email = email
	return s
}

func (s *GetDetailOnSubAccountFuturesAccountV2Service) FuturesType(futuresType int) *GetDetailOnSubAccountFuturesAccountV2Service {
	s.futuresType = futuresType
	return s
}

func (s *GetDetailOnSubAccountFuturesAccountV2Service) Do(ctx context.Context, opts ...RequestOption) (res interface{}, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getDetailOnSubAccountFuturesAccountV2Endpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("futuresType", s.futuresType)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	if s.futuresType == 1 {
		res = new(GetDetailOnSubAccountFuturesAccountV2USDTResp)
	} else {
		res = new(GetDetailOnSubAccountFuturesAccountV2COINResp)
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetDetailOnSubAccountFuturesAccountV2USDTResp struct {
	FutureAccountResp struct {
		Email  string `json:"email"`
		Assets []struct {
			Asset                  string `json:"asset"`
			InitialMargin          string `json:"initialMargin"`
			MaintenanceMargin      string `json:"maintenanceMargin"`
			MarginBalance          string `json:"marginBalance"`
			MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
			OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
			PositionInitialMargin  string `json:"positionInitialMargin"`
			UnrealizedProfit       string `json:"unrealizedProfit"`
			WalletBalance          string `json:"walletBalance"`
		} `json:"assets"`
		CanDeposit                  bool   `json:"canDeposit"`
		CanTrade                    bool   `json:"canTrade"`
		CanWithdraw                 bool   `json:"canWithdraw"`
		FeeTier                     int    `json:"feeTier"`
		MaxWithdrawAmount           string `json:"maxWithdrawAmount"`
		TotalInitialMargin          string `json:"totalInitialMargin"`
		TotalMaintenanceMargin      string `json:"totalMaintenanceMargin"`
		TotalMarginBalance          string `json:"totalMarginBalance"`
		TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
		TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
		TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
		TotalWalletBalance          string `json:"totalWalletBalance"`
		UpdateTime                  uint64 `json:"updateTime"`
	} `json:"futureAccountResp"`
}

type GetDetailOnSubAccountFuturesAccountV2COINResp struct {
	DeliveryAccountResp struct {
		Email  string `json:"email"`
		Assets []struct {
			Asset                  string `json:"asset"`
			InitialMargin          string `json:"initialMargin"`
			MaintenanceMargin      string `json:"maintenanceMargin"`
			MarginBalance          string `json:"marginBalance"`
			MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
			OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
			PositionInitialMargin  string `json:"positionInitialMargin"`
			UnrealizedProfit       string `json:"unrealizedProfit"`
			WalletBalance          string `json:"walletBalance"`
		} `json:"assets"`
		CanDeposit  bool   `json:"canDeposit"`
		CanTrade    bool   `json:"canTrade"`
		CanWithdraw bool   `json:"canWithdraw"`
		FeeTier     int    `json:"feeTier"`
		UpdateTime  uint64 `json:"updateTime"`
	} `json:"deliveryAccountResp"`
}

// Get Summary of Sub-account's Futures Account V2 (For Master Account)
const (
	getSummaryOfSubAccountFuturesAccountV2Endpoint = "/sapi/v2/sub-account/futures/accountSummary"
)

type GetSummaryOfSubAccountFuturesAccountV2Service struct {
	c           *Client
	futuresType int
	page        *int
	limit       *int
}

func (s *GetSummaryOfSubAccountFuturesAccountV2Service) FuturesType(futuresType int) *GetSummaryOfSubAccountFuturesAccountV2Service {
	s.futuresType = futuresType
	return s
}

func (s *GetSummaryOfSubAccountFuturesAccountV2Service) Page(page int) *GetSummaryOfSubAccountFuturesAccountV2Service {
	s.page = &page
	return s
}

func (s *GetSummaryOfSubAccountFuturesAccountV2Service) Limit(limit int) *GetSummaryOfSubAccountFuturesAccountV2Service {
	s.limit = &limit
	return s
}

func (s *GetSummaryOfSubAccountFuturesAccountV2Service) Do(ctx context.Context, opts ...RequestOption) (res interface{}, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getSummaryOfSubAccountFuturesAccountV2Endpoint,
		secType:  secTypeSigned,
	}
	r.setParam("futuresType", s.futuresType)
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	if s.futuresType == 1 {
		res = new(GetSummaryOfSubAccountFuturesAccountV2USDTResp)
	} else {
		res = new(GetSummaryOfSubAccountFuturesAccountV2COINResp)
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetSummaryOfSubAccountFuturesAccountV2USDTResp struct {
	FutureAccountSummaryResp struct {
		TotalInitialMargin          string `json:"totalInitialMargin"`
		TotalMaintenanceMargin      string `json:"totalMaintenanceMargin"`
		TotalMarginBalance          string `json:"totalMarginBalance"`
		TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
		TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
		TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
		TotalWalletBalance          string `json:"totalWalletBalance"`
		Asset                       string `json:"asset"`
		SubAccountList              []struct {
			Email                       string `json:"email"`
			TotalInitialMargin          string `json:"totalInitialMargin"`
			TotalMaintenanceMargin      string `json:"totalMaintenanceMargin"`
			TotalMarginBalance          string `json:"totalMarginBalance"`
			TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
			TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
			TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
			TotalWalletBalance          string `json:"totalWalletBalance"`
			Asset                       string `json:"asset"`
		} `json:"subAccountList"`
	} `json:"futureAccountSummaryResp"`
}

type GetSummaryOfSubAccountFuturesAccountV2COINResp struct {
	DeliveryAccountSummaryResp struct {
		TotalMarginBalanceOfBTC    string `json:"totalMarginBalanceOfBTC"`
		TotalUnrealizedProfitOfBTC string `json:"totalUnrealizedProfitOfBTC"`
		TotalWalletBalanceOfBTC    string `json:"totalWalletBalanceOfBTC"`
		Asset                      string `json:"asset"`
		SubAccountList             []struct {
			Email                 string `json:"email"`
			TotalMarginBalance    string `json:"totalMarginBalance"`
			TotalUnrealizedProfit string `json:"totalUnrealizedProfit"`
			TotalWalletBalance    string `json:"totalWalletBalance"`
			Asset                 string `json:"asset"`
		} `json:"subAccountList"`
	} `json:"deliveryAccountSummaryResp"`
}

// Get Futures Position-Risk of Sub-account V2 (For Master Account)
const (
	getFuturesPositionRiskOfSubAccountV2Endpoint = "/sapi/v2/sub-account/futures/positionRisk"
)

type GetFuturesPositionRiskOfSubAccountV2Service struct {
	c           *Client
	email       string
	futuresType int
}

func (s *GetFuturesPositionRiskOfSubAccountV2Service) Email(email string) *GetFuturesPositionRiskOfSubAccountV2Service {
	s.email = email
	return s
}

func (s *GetFuturesPositionRiskOfSubAccountV2Service) FuturesType(futuresType int) *GetFuturesPositionRiskOfSubAccountV2Service {
	s.futuresType = futuresType
	return s
}

func (s *GetFuturesPositionRiskOfSubAccountV2Service) Do(ctx context.Context, opts ...RequestOption) (res interface{}, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getFuturesPositionRiskOfSubAccountV2Endpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("futuresType", s.futuresType)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	if s.futuresType == 1 {
		res = new(GetFuturesPositionRiskOfSubAccountV2USDTResp)
	} else {
		res = new(GetFuturesPositionRiskOfSubAccountV2COINResp)
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetFuturesPositionRiskOfSubAccountV2USDTResp struct {
	FuturePositionRiskVos []struct {
		EntryPrice       string `json:"entryPrice"`
		Leverage         string `json:"leverage"`
		MaxNotional      string `json:"maxNotional"`
		LiquidationPrice string `json:"liquidationPrice"`
		MarkPrice        string `json:"markPrice"`
		PositionAmount   string `json:"positionAmount"`
		Symbol           string `json:"symbol"`
		UnrealizedProfit string `json:"unrealizedProfit"`
	} `json:"futurePositionRiskVos"`
}

type GetFuturesPositionRiskOfSubAccountV2COINResp struct {
	DeliveryPositionRiskVos []struct {
		EntryPrice       string `json:"entryPrice"`
		MarkPrice        string `json:"markPrice"`
		Leverage         string `json:"leverage"`
		Isolated         string `json:"isolated"`
		IsolatedWallet   string `json:"isolatedWallet"`
		IsolatedMargin   string `json:"isolatedMargin"`
		IsAutoAddMargin  string `json:"isAutoAddMargin"`
		PositionSide     string `json:"positionSide"`
		PositionAmount   string `json:"positionAmount"`
		Symbol           string `json:"symbol"`
		UnrealizedProfit string `json:"unrealizedProfit"`
	} `json:"deliveryPositionRiskVos"`
}

// Enable Leverage Token for Sub-account (For Master Account)
const (
	enableLeverageTokenForSubAccountEndpoint = "/sapi/v1/sub-account/blvt/enable"
)

type EnableLeverageTokenForSubAccountService struct {
	c          *Client
	email      string
	enableBlvt bool
}

func (s *EnableLeverageTokenForSubAccountService) Email(email string) *EnableLeverageTokenForSubAccountService {
	s.email = email
	return s
}

func (s *EnableLeverageTokenForSubAccountService) EnableBlvt(enableBlvt bool) *EnableLeverageTokenForSubAccountService {
	s.enableBlvt = enableBlvt
	return s
}

func (s *EnableLeverageTokenForSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *EnableLeverageTokenForSubAccountResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: enableLeverageTokenForSubAccountEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("enableBlvt", s.enableBlvt)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(EnableLeverageTokenForSubAccountResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type EnableLeverageTokenForSubAccountResp struct {
	Email      string `json:"email"`
	EnableBlvt bool   `json:"enableBlvt"`
}

// Get IP Restriction for a Sub-account API Key (For Master Account)

const (
	getIPRestrictionForSubAccountAPIKeyEndpoint = "/sapi/v1/sub-account/subaccountApi/ipRestriction"
)

type GetIPRestrictionForSubAccountAPIKeyService struct {
	c                *Client
	email            string
	subAccountApiKey string
}

func (s *GetIPRestrictionForSubAccountAPIKeyService) Email(email string) *GetIPRestrictionForSubAccountAPIKeyService {
	s.email = email
	return s
}

func (s *GetIPRestrictionForSubAccountAPIKeyService) SubAccountApiKey(subAccountApiKey string) *GetIPRestrictionForSubAccountAPIKeyService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *GetIPRestrictionForSubAccountAPIKeyService) Do(ctx context.Context, opts ...RequestOption) (res *GetIPRestrictionForSubAccountAPIKeyResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getIPRestrictionForSubAccountAPIKeyEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("subAccountApiKey", s.subAccountApiKey)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetIPRestrictionForSubAccountAPIKeyResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetIPRestrictionForSubAccountAPIKeyResp struct {
	IpRestrict string `json:"ipRestrict"`
	IpList     []struct {
		Ip string `json:"ip"`
	} `json:"ipList"`
	UpdateTime uint64 `json:"updateTime"`
	ApiKey     string `json:"apiKey"`
}

// Delete IP List For a Sub-account API Key (For Master Account)

const (
	deleteIPListForSubAccountAPIKeyEndpoint = "/sapi/v1/sub-account/subaccountApi/ipRestriction/ipList"
)

type DeleteIPListForSubAccountAPIKeyService struct {
	c                *Client
	email            string
	subAccountApiKey string
	ipAddress        *string
}

func (s *DeleteIPListForSubAccountAPIKeyService) Email(email string) *DeleteIPListForSubAccountAPIKeyService {
	s.email = email
	return s
}

func (s *DeleteIPListForSubAccountAPIKeyService) SubAccountApiKey(subAccountApiKey string) *DeleteIPListForSubAccountAPIKeyService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *DeleteIPListForSubAccountAPIKeyService) IpAddress(ipAddress string) *DeleteIPListForSubAccountAPIKeyService {
	s.ipAddress = &ipAddress
	return s
}

func (s *DeleteIPListForSubAccountAPIKeyService) Do(ctx context.Context, opts ...RequestOption) (res *DeleteIPListForSubAccountAPIKeyResp, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: deleteIPListForSubAccountAPIKeyEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("subAccountApiKey", s.subAccountApiKey)
	if s.ipAddress != nil {
		r.setParam("ipAddress", *s.ipAddress)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(DeleteIPListForSubAccountAPIKeyResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type DeleteIPListForSubAccountAPIKeyResp struct {
	IpRestrict string `json:"ipRestrict"`
	IpList     []struct {
		Ip string `json:"ip"`
	} `json:"ipList"`
	UpdateTime uint64 `json:"updateTime"`
	ApiKey     string `json:"apiKey"`
}

// Update IP Restriction for Sub-Account API key (For Master Account)

const (
	updateIPRestrictionForSubAccountAPIKeyEndpoint = "/sapi/v2/sub-account/subaccountApi/ipRestriction"
)

type UpdateIPRestrictionForSubAccountAPIKeyService struct {
	c                *Client
	email            string
	subAccountApiKey string
	status           string
	ipAddress        *string
}

func (s *UpdateIPRestrictionForSubAccountAPIKeyService) Email(email string) *UpdateIPRestrictionForSubAccountAPIKeyService {
	s.email = email
	return s
}

func (s *UpdateIPRestrictionForSubAccountAPIKeyService) SubAccountApiKey(subAccountApiKey string) *UpdateIPRestrictionForSubAccountAPIKeyService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *UpdateIPRestrictionForSubAccountAPIKeyService) Status(status string) *UpdateIPRestrictionForSubAccountAPIKeyService {
	s.status = status
	return s
}

func (s *UpdateIPRestrictionForSubAccountAPIKeyService) IpAddress(ipAddress string) *UpdateIPRestrictionForSubAccountAPIKeyService {
	s.ipAddress = &ipAddress
	return s
}

func (s *UpdateIPRestrictionForSubAccountAPIKeyService) Do(ctx context.Context, opts ...RequestOption) (res *UpdateIPRestrictionForSubAccountAPIKeyResp, err error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: updateIPRestrictionForSubAccountAPIKeyEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("subAccountApiKey", s.subAccountApiKey)
	r.setParam("status", s.status)
	if s.ipAddress != nil {
		r.setParam("ipAddress", *s.ipAddress)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UpdateIPRestrictionForSubAccountAPIKeyResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type UpdateIPRestrictionForSubAccountAPIKeyResp struct {
	Status string `json:"status"`
	IpList []struct {
		Ip string `json:"ip"`
	} `json:"ipList"`
	UpdateTime uint64 `json:"updateTime"`
	ApiKey     string `json:"apiKey"`
}

// Deposit Assets Into The Managed Sub-account（For Investor Master Account）

const (
	depositAssetsIntoTheManagedSubAccountEndpoint = "/sapi/v1/managed-subaccount/deposit"
)

type DepositAssetsIntoTheManagedSubAccountService struct {
	c       *Client
	toEmail string
	asset   string
	amount  float64
}

func (s *DepositAssetsIntoTheManagedSubAccountService) ToEmail(toEmail string) *DepositAssetsIntoTheManagedSubAccountService {
	s.toEmail = toEmail
	return s
}

func (s *DepositAssetsIntoTheManagedSubAccountService) Asset(asset string) *DepositAssetsIntoTheManagedSubAccountService {
	s.asset = asset
	return s
}

func (s *DepositAssetsIntoTheManagedSubAccountService) Amount(amount float64) *DepositAssetsIntoTheManagedSubAccountService {
	s.amount = amount
	return s
}

func (s *DepositAssetsIntoTheManagedSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *DepositAssetsIntoTheManagedSubAccountResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: depositAssetsIntoTheManagedSubAccountEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("toEmail", s.toEmail)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(DepositAssetsIntoTheManagedSubAccountResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type DepositAssetsIntoTheManagedSubAccountResp struct {
	TranId int64 `json:"tranId"`
}

// Query Managed Sub-account Asset Details（For Investor Master Account）

const (
	queryManagedSubAccountAssetDetailsEndpoint = "/sapi/v1/sub-account/managed-subaccount/asset"
)

type QueryManagedSubAccountAssetDetailsService struct {
	c     *Client
	email string
}

func (s *QueryManagedSubAccountAssetDetailsService) Email(email string) *QueryManagedSubAccountAssetDetailsService {
	s.email = email
	return s
}

func (s *QueryManagedSubAccountAssetDetailsService) Do(ctx context.Context, opts ...RequestOption) (res *QueryManagedSubAccountAssetDetailsResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: queryManagedSubAccountAssetDetailsEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QueryManagedSubAccountAssetDetailsResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type QueryManagedSubAccountAssetDetailsResp struct {
	AssetDetail []struct {
		Coin             string `json:"coin"`
		Name             string `json:"name"`
		TotalBalance     string `json:"totalBalance"`
		AvailableBalance string `json:"availableBalance"`
		InOrder          string `json:"inOrder"`
		BtcValue         string `json:"btcValue"`
	} `json:"assetDetail"`
}

// Withdrawl Assets From The Managed Sub-account（For Investor Master Account）

const (
	withdrawAssetsFromTheManagedSubAccountEndpoint = "/sapi/v1/sub-account/managed-subaccount/withdraw"
)

type WithdrawAssetsFromTheManagedSubAccountService struct {
	c            *Client
	fromEmail    string
	asset        string
	amount       float32
	transferDate *int64
}

func (s *WithdrawAssetsFromTheManagedSubAccountService) FromEmail(fromEmail string) *WithdrawAssetsFromTheManagedSubAccountService {
	s.fromEmail = fromEmail
	return s
}

func (s *WithdrawAssetsFromTheManagedSubAccountService) Asset(asset string) *WithdrawAssetsFromTheManagedSubAccountService {
	s.asset = asset
	return s
}

func (s *WithdrawAssetsFromTheManagedSubAccountService) Amount(amount float32) *WithdrawAssetsFromTheManagedSubAccountService {
	s.amount = amount
	return s
}

func (s *WithdrawAssetsFromTheManagedSubAccountService) TransferDate(transferDate int64) *WithdrawAssetsFromTheManagedSubAccountService {
	s.transferDate = &transferDate
	return s
}

func (s *WithdrawAssetsFromTheManagedSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *WithdrawAssetsFromTheManagedSubAccountResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: withdrawAssetsFromTheManagedSubAccountEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("fromEmail", s.fromEmail)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	if s.transferDate != nil {
		r.setParam("transferDate", *s.transferDate)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(WithdrawAssetsFromTheManagedSubAccountResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type WithdrawAssetsFromTheManagedSubAccountResp struct {
	TranId int64 `json:"tranId"`
}

// Query Managed Sub-account Snapshot（For Investor Master Account）

const (
	queryManagedSubAccountSnapshotEndpoint = "/sapi/v1/managed-subaccount/accountSnapshot"
)

type QueryManagedSubAccountSnapshotService struct {
	c         *Client
	email     string
	subType   string
	startTime *uint64
	endTime   *uint64
	limit     *int
}

func (s *QueryManagedSubAccountSnapshotService) Email(email string) *QueryManagedSubAccountSnapshotService {
	s.email = email
	return s
}

func (s *QueryManagedSubAccountSnapshotService) SubType(subType string) *QueryManagedSubAccountSnapshotService {
	s.subType = subType
	return s
}

func (s *QueryManagedSubAccountSnapshotService) StartTime(startTime uint64) *QueryManagedSubAccountSnapshotService {
	s.startTime = &startTime
	return s
}

func (s *QueryManagedSubAccountSnapshotService) EndTime(endTime uint64) *QueryManagedSubAccountSnapshotService {
	s.endTime = &endTime
	return s
}

func (s *QueryManagedSubAccountSnapshotService) Limit(limit int) *QueryManagedSubAccountSnapshotService {
	s.limit = &limit
	return s
}

func (s *QueryManagedSubAccountSnapshotService) Do(ctx context.Context, opts ...RequestOption) (res *QueryManagedSubAccountSnapshotResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: queryManagedSubAccountSnapshotEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("type", s.subType)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QueryManagedSubAccountSnapshotResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type QueryManagedSubAccountSnapshotResp struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	SnapshotVos []struct {
		Data []struct {
			Balances []struct {
				Asset  string `json:"asset"`
				Free   string `json:"free"`
				Locked string `json:"locked"`
			} `json:"balances"`
			TotalAssetOfBtc string `json:"totalAssetOfBtc"`
		} `json:"data"`
		Type       string `json:"type"`
		UpdateTime uint64 `json:"updateTime"`
	} `json:"snapshotVos"`
}

// Query Managed Sub Account Transfer Log (Investor) (USER_DATA)
const (
	queryManagedSubAccountTransferLogEndpoint = "/sapi/v1/managed-subaccount/queryTransLogForInvestor"
)

type QueryManagedSubAccountTransferLogService struct {
	c                           *Client
	email                       string
	startTime                   uint64
	endTime                     uint64
	page                        int
	limit                       int
	transfers                   *string
	transferFunctionAccountType *string
}

func (s *QueryManagedSubAccountTransferLogService) Email(email string) *QueryManagedSubAccountTransferLogService {
	s.email = email
	return s
}

func (s *QueryManagedSubAccountTransferLogService) StartTime(startTime uint64) *QueryManagedSubAccountTransferLogService {
	s.startTime = startTime
	return s
}

func (s *QueryManagedSubAccountTransferLogService) EndTime(endTime uint64) *QueryManagedSubAccountTransferLogService {
	s.endTime = endTime
	return s
}

func (s *QueryManagedSubAccountTransferLogService) Page(page int) *QueryManagedSubAccountTransferLogService {
	s.page = page
	return s
}

func (s *QueryManagedSubAccountTransferLogService) Limit(limit int) *QueryManagedSubAccountTransferLogService {
	s.limit = limit
	return s
}

func (s *QueryManagedSubAccountTransferLogService) Transfers(transfers string) *QueryManagedSubAccountTransferLogService {
	s.transfers = &transfers
	return s
}

func (s *QueryManagedSubAccountTransferLogService) TransferFunctionAccountType(transferFunctionAccountType string) *QueryManagedSubAccountTransferLogService {
	s.transferFunctionAccountType = &transferFunctionAccountType
	return s
}

func (s *QueryManagedSubAccountTransferLogService) Do(ctx context.Context, opts ...RequestOption) (res *QueryManagedSubAccountTransferLogResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: queryManagedSubAccountTransferLogEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("startTime", s.startTime)
	r.setParam("endTime", s.endTime)
	r.setParam("page", s.page)
	r.setParam("limit", s.limit)
	if s.transfers != nil {
		r.setParam("transfers", *s.transfers)
	}
	if s.transferFunctionAccountType != nil {
		r.setParam("transferFunctionAccountType", *s.transferFunctionAccountType)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QueryManagedSubAccountTransferLogResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type QueryManagedSubAccountTransferLogResp struct {
	ManagerSubTransferHistoryVos []struct {
		FromEmail       string `json:"fromEmail"`
		FromAccountType string `json:"fromAccountType"`
		ToEmail         string `json:"toEmail"`
		ToAccountType   string `json:"toAccountType"`
		Asset           string `json:"asset"`
		Amount          int    `json:"amount"`
		ScheduledData   int    `json:"scheduledData"`
		CreateTime      uint64 `json:"createTime"`
		Status          string `json:"status"`
	} `json:"managerSubTransferHistoryVos"`
	Count int `json:"count"`
}

// Query Managed Sub-account Futures Asset Details（For Investor Master Account）(USER_DATA)
const (
	queryManagedSubAccountFuturesAssetDetailsEndpoint = "/sapi/v1/managed-subaccount/fetch-future-asset"
)

type QueryManagedSubAccountFuturesAssetDetailsService struct {
	c     *Client
	email string
}

func (s *QueryManagedSubAccountFuturesAssetDetailsService) Email(email string) *QueryManagedSubAccountFuturesAssetDetailsService {
	s.email = email
	return s
}

func (s *QueryManagedSubAccountFuturesAssetDetailsService) Do(ctx context.Context, opts ...RequestOption) (res *QueryManagedSubAccountFuturesAssetDetailsResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: queryManagedSubAccountFuturesAssetDetailsEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QueryManagedSubAccountFuturesAssetDetailsResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type QueryManagedSubAccountFuturesAssetDetailsResp struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	SnapshotVos []struct {
		Type       string `json:"type"`
		UpdateTime uint64 `json:"updateTime"`
		Data       struct {
			Assets []struct {
				Asset         string `json:"asset"`
				MarginBalance int64  `json:"marginBalance"`
				WalletBalance int64  `json:"walletBalance"`
			} `json:"assets"`
			Position []struct {
				Symbol      string `json:"symbol"`
				EntryPrice  int64  `json:"entryPrice"`
				MarkPrice   int64  `json:"markPrice"`
				PositionAmt int64  `json:"positionAmt"`
			} `json:"position"`
		} `json:"data"`
	} `json:"snapshotVos"`
}

// Query Managed Sub-account Margin Asset Details (For Investor Master Account) (USER_DATA)
const (
	queryManagedSubAccountMarginAssetDetailsEndpoint = "/sapi/v1/managed-subaccount/marginAsset"
)

type QueryManagedSubAccountMarginAssetDetailsService struct {
	c     *Client
	email string
}

func (s *QueryManagedSubAccountMarginAssetDetailsService) Email(email string) *QueryManagedSubAccountMarginAssetDetailsService {
	s.email = email
	return s
}

func (s *QueryManagedSubAccountMarginAssetDetailsService) Do(ctx context.Context, opts ...RequestOption) (res *QueryManagedSubAccountMarginAssetDetailsResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: queryManagedSubAccountMarginAssetDetailsEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QueryManagedSubAccountMarginAssetDetailsResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type QueryManagedSubAccountMarginAssetDetailsResp struct {
	MarginLevel         string `json:"marginLevel"`
	TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`
	UserAssets          []struct {
		Asset    string `json:"asset"`
		Borrowed string `json:"borrowed"`
		Free     string `json:"free"`
		Interest string `json:"interest"`
		Locked   string `json:"locked"`
		NetAsset string `json:"netAsset"`
	} `json:"userAssets"`
}

// Query Managed Sub Account Transfer Log (Trading Team) (USER_DATA)
const (
	queryManagedSubAccountTransferLogForTradingTeamEndpoint = "/sapi/v1/managed-subaccount/queryTransLogForTradeParent"
)

type QueryManagedSubAccountTransferLogForTradingTeamService struct {
	c                           *Client
	email                       string
	startTime                   uint64
	endTime                     uint64
	page                        int
	limit                       int
	transfers                   *string
	transferFunctionAccountType *string
}

func (s *QueryManagedSubAccountTransferLogForTradingTeamService) Email(email string) *QueryManagedSubAccountTransferLogForTradingTeamService {
	s.email = email
	return s
}

func (s *QueryManagedSubAccountTransferLogForTradingTeamService) StartTime(startTime uint64) *QueryManagedSubAccountTransferLogForTradingTeamService {
	s.startTime = startTime
	return s
}

func (s *QueryManagedSubAccountTransferLogForTradingTeamService) EndTime(endTime uint64) *QueryManagedSubAccountTransferLogForTradingTeamService {
	s.endTime = endTime
	return s
}

func (s *QueryManagedSubAccountTransferLogForTradingTeamService) Page(page int) *QueryManagedSubAccountTransferLogForTradingTeamService {
	s.page = page
	return s
}

func (s *QueryManagedSubAccountTransferLogForTradingTeamService) Limit(limit int) *QueryManagedSubAccountTransferLogForTradingTeamService {
	s.limit = limit
	return s
}

func (s *QueryManagedSubAccountTransferLogForTradingTeamService) Transfers(transfers string) *QueryManagedSubAccountTransferLogForTradingTeamService {
	s.transfers = &transfers
	return s
}

func (s *QueryManagedSubAccountTransferLogForTradingTeamService) TransferFunctionAccountType(transferFunctionAccountType string) *QueryManagedSubAccountTransferLogForTradingTeamService {
	s.transferFunctionAccountType = &transferFunctionAccountType
	return s
}

func (s *QueryManagedSubAccountTransferLogForTradingTeamService) Do(ctx context.Context, opts ...RequestOption) (res *QueryManagedSubAccountTransferLogForTradingTeamResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: queryManagedSubAccountTransferLogForTradingTeamEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("startTime", s.startTime)
	r.setParam("endTime", s.endTime)
	r.setParam("page", s.page)
	r.setParam("limit", s.limit)
	if s.transfers != nil {
		r.setParam("transfers", *s.transfers)
	}
	if s.transferFunctionAccountType != nil {
		r.setParam("transferFunctionAccountType", *s.transferFunctionAccountType)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QueryManagedSubAccountTransferLogForTradingTeamResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type QueryManagedSubAccountTransferLogForTradingTeamResp struct {
	ManagerSubTransferHistoryVos []struct {
		FromEmail       string `json:"fromEmail"`
		FromAccountType string `json:"fromAccountType"`
		ToEmail         string `json:"toEmail"`
		ToAccountType   string `json:"toAccountType"`
		Asset           string `json:"asset"`
		Amount          string `json:"amount"`
		ScheduledData   int64  `json:"scheduledData"`
		CreateTime      uint64 `json:"createTime"`
		Status          string `json:"status"`
	} `json:"managerSubTransferHistoryVos"`
	Count int `json:"count"`
}

// Query Sub-account Assets (For Master Account)(USER_DATA)
const (
	querySubAccountAssetsForMasterAccountEndpoint = "/sapi/v4/sub-account/assets"
)

type QuerySubAccountAssetsForMasterAccountService struct {
	c     *Client
	email string
}

func (s *QuerySubAccountAssetsForMasterAccountService) Email(email string) *QuerySubAccountAssetsForMasterAccountService {
	s.email = email
	return s
}

func (s *QuerySubAccountAssetsForMasterAccountService) Do(ctx context.Context, opts ...RequestOption) (res *QuerySubAccountAssetsForMasterAccountResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: querySubAccountAssetsForMasterAccountEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QuerySubAccountAssetsForMasterAccountResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type QuerySubAccountAssetsForMasterAccountResp struct {
	Balances []struct {
		Asset  string `json:"asset"`
		Free   string `json:"free"`
		Locked string `json:"locked"`
	} `json:"balances"`
}

// Query Managed Sub-account List
const (
	queryManagedSubAccountListEndpoint = "/sapi/v1/managed-subaccount/info"
)

type QueryManagedSubAccountList struct {
	c     *Client
	email *string
	page  *int
	limit *int
}

func (s *QueryManagedSubAccountList) Email(email string) *QueryManagedSubAccountList {
	s.email = &email
	return s
}

func (s *QueryManagedSubAccountList) Page(page int) *QueryManagedSubAccountList {
	s.page = &page
	return s
}

func (s *QueryManagedSubAccountList) Limit(limit int) *QueryManagedSubAccountList {
	s.limit = &limit
	return s
}

func (s *QueryManagedSubAccountList) Do(ctx context.Context, opts ...RequestOption) (res *QueryManagedSubAccountListResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: queryManagedSubAccountListEndpoint,
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", *s.email)
	}
	if s.page != nil {
		r.setParam("page", strconv.Itoa(*s.page))
	}
	if s.limit != nil {
		r.setParam("limit", strconv.Itoa(*s.limit))
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QueryManagedSubAccountListResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type QueryManagedSubAccountListResp struct {
	Total                    int `json:"total"`
	ManagerSubUserInfoVoList []struct {
		RootUserId               int    `json:"rootUserId"`
		ManagersubUserId         int    `json:"managersubUserId"`
		BindParentUserId         int    `json:"bindParentUserId"`
		Email                    string `json:"email"`
		InsertTimestamp          uint64 `json:"insertTimestamp"`
		BindParentEmail          string `json:"bindParentEmail"`
		IsSubUserEnabled         bool   `json:"isSubUserEnabled"`
		IsUserActive             bool   `json:"isUserActive"`
		IsMarginEnabled          bool   `json:"isMarginEnabled"`
		IsFutureEnabled          bool   `json:"isFutureEnabled"`
		IsSignedLVTRiskAgreement bool   `json:"isSignedLVTRiskAgreement"`
	} `json:"managerSubUserInfoVoList"`
}

// Query Sub-account Transaction Tatistics (For Master Account) (USER_DATA)
const (
	QuerySubAccountTransactionTatisticsEndpoint = "/sapi/v1/sub-account/transaction-tatistics"
)

type QuerySubAccountTransactionTatistics struct {
	c     *Client
	email string
}

func (s *QuerySubAccountTransactionTatistics) Email(email string) *QuerySubAccountTransactionTatistics {
	s.email = email
	return s
}

func (s *QuerySubAccountTransactionTatistics) Do(ctx context.Context, opts ...RequestOption) (res *QuerySubAccountTransactionTatisticsResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: QuerySubAccountTransactionTatisticsEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QuerySubAccountTransactionTatisticsResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type QuerySubAccountTransactionTatisticsResp struct {
	Recent30BtcTotal         string `json:"recent30BtcTotal"`
	Recent30BtcFuturesTotal  string `json:"recent30BtcFuturesTotal"`
	Recent30BtcMarginTotal   string `json:"recent30BtcMarginTotal"`
	Recent30BusdTotal        string `json:"recent30BusdTotal"`
	Recent30BusdFuturesTotal string `json:"recent30BusdFuturesTotal"`
	Recent30BusdMarginTotal  string `json:"recent30BusdMarginTotal"`
	TradeInfoVos             []struct {
		UserId      int64 `json:"userId"`
		Btc         int   `json:"btc"`
		BtcFutures  int   `json:"btcFutures"`
		BtcMargin   int   `json:"btcMargin"`
		Busd        int   `json:"busd"`
		BusdFutures int   `json:"busdFutures"`
		BusdMargin  int   `json:"busdMargin"`
		Date        int64 `json:"date"`
	} `json:"tradeInfoVos"`
}

// Get Managed Sub-account Deposit Address (For Investor Master Account) (USER_DATA)
const (
	getManagedSubAccountDepositAddressEndpoint = "/sapi/v1/managed-subaccount/deposit/address"
)

type GetManagedSubAccountDepositAddressService struct {
	c       *Client
	email   string
	coin    string
	network *string
}

func (s *GetManagedSubAccountDepositAddressService) Email(email string) *GetManagedSubAccountDepositAddressService {
	s.email = email
	return s
}

func (s *GetManagedSubAccountDepositAddressService) Coin(coin string) *GetManagedSubAccountDepositAddressService {
	s.coin = coin
	return s
}

func (s *GetManagedSubAccountDepositAddressService) Network(network string) *GetManagedSubAccountDepositAddressService {
	s.network = &network
	return s
}

func (s *GetManagedSubAccountDepositAddressService) Do(ctx context.Context, opts ...RequestOption) (res *GetManagedSubAccountDepositAddressResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getManagedSubAccountDepositAddressEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("email", s.email)
	r.setParam("coin", s.coin)
	if s.network != nil {
		r.setParam("network", *s.network)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetManagedSubAccountDepositAddressResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetManagedSubAccountDepositAddressResp struct {
	Coin    string `json:"coin"`
	Address string `json:"address"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}
