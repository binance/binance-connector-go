package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

// System Status (System)
const (
	systemStatusEndpoint = "/sapi/v1/system/status"
)

// GetSystemStatusService get account info
type GetSystemStatusService struct {
	c *Client
}

func (s *GetSystemStatusService) Do(ctx context.Context, opts ...RequestOption) (res []*SystemStatusResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: systemStatusEndpoint,
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*SystemStatusResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SystemStatusResponse define response of GetSystemStatusService
type SystemStatusResponse struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

// All Coins' Information (USER_DATA)
const (
	allCoinsInfoEndpoint = "/sapi/v1/capital/config/getall"
)

// GetAllCoinsInfoService get all coins' information
type GetAllCoinsInfoService struct {
	c *Client
}

func (s *GetAllCoinsInfoService) Do(ctx context.Context) (res []*CoinInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: allCoinsInfoEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = make([]*CoinInfo, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CoinInfo define response of GetAllCoinsInfoService
type CoinInfo struct {
	Coin             string `json:"coin"`
	DepositAllEnable bool   `json:"depositAllEnable"`
	Free             string `json:"free"`
	Freeze           string `json:"freeze"`
	Ipoable          string `json:"ipoable"`
	Ipoing           string `json:"ipoing"`
	IsLegalMoney     bool   `json:"isLegalMoney"`
	Locked           string `json:"locked"`
	Name             string `json:"name"`
	NetworkList      []struct {
		AddressRegex            string `json:"addressRegex"`
		Coin                    string `json:"coin"`
		DepositDesc             string `json:"depositDesc"`
		DepositEnable           bool   `json:"depositEnable"`
		IsDefault               bool   `json:"isDefault"`
		MemoRegex               string `json:"memoRegex"`
		MinConfirm              int    `json:"minConfirm"`
		Name                    string `json:"name"`
		Network                 string `json:"network"`
		ResetAddressStatus      bool   `json:"resetAddressStatus"`
		SpecialTips             string `json:"specialTips"`
		UnLockConfirm           int    `json:"unLockConfirm"`
		WithdrawDesc            string `json:"withdrawDesc"`
		WithdrawEnable          bool   `json:"withdrawEnable"`
		WithdrawFee             string `json:"withdrawFee"`
		WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple"`
		WithdrawMax             string `json:"withdrawMax"`
		WithdrawMin             string `json:"withdrawMin"`
		SameAddress             bool   `json:"sameAddress"`
		EstimatedArrivalTime    uint64 `json:"estimatedArrivalTime"`
		Busy                    bool   `json:"busy"`
	} `json:"networkList"`
	Storage           string `json:"storage"`
	Trading           bool   `json:"trading"`
	WithdrawAllEnable bool   `json:"withdrawAllEnable"`
	Withdrawing       string `json:"withdrawing"`
}

// Daily Account Snapshot (USER_DATA)
const (
	accountSnapshotEndpoint = "/sapi/v1/accountSnapshot"
)

// GetAccountSnapshotService get all orders from account
type GetAccountSnapshotService struct {
	c          *Client
	marketType string
	startTime  *uint64
	endTime    *uint64
	limit      *int
}

// MarketType set market type
func (s *GetAccountSnapshotService) MarketType(marketType string) *GetAccountSnapshotService {
	s.marketType = marketType
	return s
}

// StartTime set start time
func (s *GetAccountSnapshotService) StartTime(startTime uint64) *GetAccountSnapshotService {
	s.startTime = &startTime
	return s
}

// EndTime set end time
func (s *GetAccountSnapshotService) EndTime(endTime uint64) *GetAccountSnapshotService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetAccountSnapshotService) Limit(limit int) *GetAccountSnapshotService {
	s.limit = &limit
	return s
}

func (s *GetAccountSnapshotService) Do(ctx context.Context) (res *AccountSnapshotResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: accountSnapshotEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("type", s.marketType)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AccountSnapshotResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AccountSnapshotResponse define response of GetAccountSnapshotService
type AccountSnapshotResponse struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	SnapshotVos []struct {
		Data struct {
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

// Disable Fast Withdraw Switch (USER_DATA)
const (
	disableFastWithdrawSwitchEndpoint = "/sapi/v1/account/disableFastWithdrawSwitch"
)

// DisableFastWithdrawSwitchService disable fast withdraw switch
type DisableFastWithdrawSwitchService struct {
	c *Client
}

func (s *DisableFastWithdrawSwitchService) Do(ctx context.Context) (res *DisableFastWithdrawSwitchResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: disableFastWithdrawSwitchEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(DisableFastWithdrawSwitchResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DisableFastWithdrawSwitchResponse define response of DisableFastWithdrawSwitchService
// This endpoint has empty response
type DisableFastWithdrawSwitchResponse struct {
}

// Enable Fast Withdraw Switch (USER_DATA)
const (
	enableFastWithdrawSwitchEndpoint = "/sapi/v1/account/enableFastWithdrawSwitch"
)

// EnableFastWithdrawSwitchService enable fast withdraw switch
type EnableFastWithdrawSwitchService struct {
	c *Client
}

func (s *EnableFastWithdrawSwitchService) Do(ctx context.Context) (res *EnableFastWithdrawSwitchResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: enableFastWithdrawSwitchEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(EnableFastWithdrawSwitchResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// EnableFastWithdrawSwitchResponse define response of EnableFastWithdrawSwitchService
// This endpoint has empty response
type EnableFastWithdrawSwitchResponse struct {
}

// Withdraw(USER_DATA)
const (
	withdrawEndpoint = "/sapi/v1/capital/withdraw/apply"
)

// WithdrawService withdraw
type WithdrawService struct {
	c                  *Client
	coin               string
	withdrawOrderId    *string
	network            *string
	address            string
	addressTag         *string
	amount             float64
	transactionFeeFlag *bool
	name               *string
	walletType         *int
}

// Coin set coin
func (s *WithdrawService) Coin(coin string) *WithdrawService {
	s.coin = coin
	return s
}

// WithdrawOrderId set withdrawOrderId
func (s *WithdrawService) WithdrawOrderId(withdrawOrderId string) *WithdrawService {
	s.withdrawOrderId = &withdrawOrderId
	return s
}

// Network set network
func (s *WithdrawService) Network(network string) *WithdrawService {
	s.network = &network
	return s
}

// Address set address
func (s *WithdrawService) Address(address string) *WithdrawService {
	s.address = address
	return s
}

// AddressTag set addressTag
func (s *WithdrawService) AddressTag(addressTag string) *WithdrawService {
	s.addressTag = &addressTag
	return s
}

// Amount set amount
func (s *WithdrawService) Amount(amount float64) *WithdrawService {
	s.amount = amount
	return s
}

// TransactionFeeFlag set transactionFeeFlag
func (s *WithdrawService) TransactionFeeFlag(transactionFeeFlag bool) *WithdrawService {
	s.transactionFeeFlag = &transactionFeeFlag
	return s
}

// Name set name
func (s *WithdrawService) Name(name string) *WithdrawService {
	s.name = &name
	return s
}

// WalletType set walletType
func (s *WithdrawService) WalletType(walletType int) *WithdrawService {
	s.walletType = &walletType
	return s
}

func (s *WithdrawService) Do(ctx context.Context) (res *WithdrawResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: withdrawEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("coin", s.coin)
	r.setParam("address", s.address)
	r.setParam("amount", s.amount)
	if s.withdrawOrderId != nil {
		r.setParam("withdrawOrderId", *s.withdrawOrderId)
	}
	if s.network != nil {
		r.setParam("network", *s.network)
	}
	if s.addressTag != nil {
		r.setParam("addressTag", *s.addressTag)
	}
	if s.transactionFeeFlag != nil {
		r.setParam("transactionFeeFlag", *s.transactionFeeFlag)
	}
	if s.name != nil {
		r.setParam("name", *s.name)
	}
	if s.walletType != nil {
		r.setParam("walletType", *s.walletType)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(WithdrawResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// WithdrawResponse define response of WithdrawService
type WithdrawResponse struct {
	Id string `json:"id"`
}

// Deposit History (supporting network) (USER_DATA)
const (
	depositHistoryEndpoint = "/sapi/v1/capital/deposit/hisrec"
)

// DepositHistoryService deposit history
type DepositHistoryService struct {
	c         *Client
	coin      *string
	status    *int
	startTime *uint64
	endTime   *uint64
	offset    *int
	limit     *int
	txid      *string
}

// Coin set coin
func (s *DepositHistoryService) Coin(coin string) *DepositHistoryService {
	s.coin = &coin
	return s
}

// Status set status
func (s *DepositHistoryService) Status(status int) *DepositHistoryService {
	s.status = &status
	return s
}

// StartTime set startTime
func (s *DepositHistoryService) StartTime(startTime uint64) *DepositHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *DepositHistoryService) EndTime(endTime uint64) *DepositHistoryService {
	s.endTime = &endTime
	return s
}

// Offset set offset
func (s *DepositHistoryService) Offset(offset int) *DepositHistoryService {
	s.offset = &offset
	return s
}

// Limit set limit
func (s *DepositHistoryService) Limit(limit int) *DepositHistoryService {
	s.limit = &limit
	return s
}

// TxId set txid
func (s *DepositHistoryService) TxId(txid string) *DepositHistoryService {
	s.txid = &txid
	return s
}

func (s *DepositHistoryService) Do(ctx context.Context) (res []*DepositHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: depositHistoryEndpoint,
		secType:  secTypeSigned,
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.txid != nil {
		r.setParam("txId", *s.txid)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = make([]*DepositHistoryResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DepositHistoryResponse define response of DepositHistoryService
type DepositHistoryResponse struct {
	Id            string `json:"id"`
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	Network       string `json:"network"`
	Status        int    `json:"status"`
	Address       string `json:"address"`
	AddressTag    string `json:"addressTag"`
	TxId          string `json:"txId"`
	InsertTime    uint64 `json:"insertTime"`
	TransferType  int    `json:"transferType"`
	ConfirmTimes  string `json:"confirmTimes"`
	UnlockConfirm int    `json:"unlockConfirm"`
	WalletType    int    `json:"walletType"`
}

// Withdraw History (supporting network) (USER_DATA)
const (
	withdrawHistoryEndpoint = "/sapi/v1/capital/withdraw/history"
)

// WithdrawHistoryService withdraw history
type WithdrawHistoryService struct {
	c               *Client
	coin            *string
	withdrawOrderId *string
	status          *int
	offset          *int
	limit           *int
	startTime       *uint64
	endTime         *uint64
}

// Coin set coin
func (s *WithdrawHistoryService) Coin(coin string) *WithdrawHistoryService {
	s.coin = &coin
	return s
}

// WithdrawOrderId set withdrawOrderId
func (s *WithdrawHistoryService) WithdrawOrderId(withdrawOrderId string) *WithdrawHistoryService {
	s.withdrawOrderId = &withdrawOrderId
	return s
}

// Status set status
func (s *WithdrawHistoryService) Status(status int) *WithdrawHistoryService {
	s.status = &status
	return s
}

// Offset set offset
func (s *WithdrawHistoryService) Offset(offset int) *WithdrawHistoryService {
	s.offset = &offset
	return s
}

// Limit set limit
func (s *WithdrawHistoryService) Limit(limit int) *WithdrawHistoryService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *WithdrawHistoryService) StartTime(startTime uint64) *WithdrawHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *WithdrawHistoryService) EndTime(endTime uint64) *WithdrawHistoryService {
	s.endTime = &endTime
	return s
}

func (s *WithdrawHistoryService) Do(ctx context.Context) (res []*WithdrawHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: withdrawHistoryEndpoint,
		secType:  secTypeSigned,
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.withdrawOrderId != nil {
		r.setParam("withdrawOrderId", *s.withdrawOrderId)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = make([]*WithdrawHistoryResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// WithdrawHistoryResponse define response of WithdrawHistoryService
type WithdrawHistoryResponse struct {
	Id              string `json:"id"`
	Amount          string `json:"amount"`
	TransactionFee  string `json:"transactionFee"`
	Coin            string `json:"coin"`
	Status          int    `json:"status"`
	Address         string `json:"address"`
	TxId            string `json:"txId"`
	ApplyTime       string `json:"applyTime"`
	Network         string `json:"network"`
	TransferType    int    `json:"transferType"`
	WithdrawOrderId string `json:"withdrawOrderId"`
	Info            string `json:"info"`
	ConfirmNo       int    `json:"confirmNo"`
	WalletType      int    `json:"walletType"`
	TxKey           string `json:"txKey"`
}

// Deposit Address (supporting network) (USER_DATA)
const (
	depositAddressEndpoint = "/sapi/v1/capital/deposit/address"
)

// DepositAddressService deposit address
type DepositAddressService struct {
	c       *Client
	coin    string
	network *string
}

// Coin set coin
func (s *DepositAddressService) Coin(coin string) *DepositAddressService {
	s.coin = coin
	return s
}

// Network set network
func (s *DepositAddressService) Network(network string) *DepositAddressService {
	s.network = &network
	return s
}

func (s *DepositAddressService) Do(ctx context.Context) (res *DepositAddressResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: depositAddressEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("coin", s.coin)
	if s.network != nil {
		r.setParam("network", *s.network)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(DepositAddressResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DepositAddressResponse define response of DepositAddressService
type DepositAddressResponse struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}

// Account Status (USER_DATA)
const (
	accountStatusEndpoint = "/sapi/v1/account/status"
)

// AccountStatusService account status
type AccountStatusService struct {
	c *Client
}

func (s *AccountStatusService) Do(ctx context.Context) (res *AccountStatusResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: accountStatusEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AccountStatusResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AccountStatusResponse define response of AccountStatusService
type AccountStatusResponse struct {
	Data string `json:"data"`
}

// Account API Trading Status (USER_DATA)
const (
	accountApiTradingStatusEndpoint = "/sapi/v1/account/apiTradingStatus"
)

// AccountApiTradingStatusService account api trading status
type AccountApiTradingStatusService struct {
	c *Client
}

func (s *AccountApiTradingStatusService) Do(ctx context.Context) (res *AccountApiTradingStatusResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: accountApiTradingStatusEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AccountApiTradingStatusResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AccountApiTradingStatusResponse define response of AccountApiTradingStatusService
type AccountApiTradingStatusResponse struct {
	Data struct {
		IsLocked           bool  `json:"isLocked"`
		PlannedRecoverTime int64 `json:"plannedRecoverTime"`
		TriggerCondition   struct {
			GCR  int `json:"GCR"`
			IFER int `json:"IFER"`
			UFR  int `json:"UFR"`
		} `json:"triggerCondition"`
		UpdateTime uint64 `json:"updateTime"`
	} `json:"data"`
}

// DustLog(USER_DATA)
const (
	dustLogEndpoint = "/sapi/v1/asset/dribblet"
)

// DustLogService dust log
type DustLogService struct {
	c         *Client
	startTime *uint64
	endTime   *uint64
}

// StartTime set startTime
func (s *DustLogService) StartTime(startTime uint64) *DustLogService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *DustLogService) EndTime(endTime uint64) *DustLogService {
	s.endTime = &endTime
	return s
}

func (s *DustLogService) Do(ctx context.Context) (res *DustLogResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: dustLogEndpoint,
		secType:  secTypeSigned,
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(DustLogResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DustLogResponse define response of DustLogService
type DustLogResponse struct {
	Total              int `json:"total"`
	UserAssetDribblets []struct {
		OperateTime              uint64 `json:"operateTime"`
		TotalTransferedAmount    string `json:"totalTransferedAmount"`
		TotalServiceChargeAmount string `json:"totalServiceChargeAmount"`
		TransId                  int64  `json:"transId"`
		UserAssetDribbletDetails []struct {
			TransId             int64  `json:"transId"`
			ServiceChargeAmount string `json:"serviceChargeAmount"`
			Amount              string `json:"amount"`
			OperateTime         uint64 `json:"operateTime"`
			TransferedAmount    string `json:"transferedAmount"`
			FromAsset           string `json:"fromAsset"`
		} `json:"userAssetDribbletDetails"`
	} `json:"userAssetDribblets"`
}

// Get Assets That Can Be Converted Into BNB (USER_DATA)
const (
	assetDetailEndpoint = "/sapi/v1/asset/dust-btc"
)

// AssetDetailService asset detail
type AssetDetailService struct {
	c *Client
}

func (s *AssetDetailService) Do(ctx context.Context) (res *AssetDetailResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: assetDetailEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AssetDetailResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AssetDetailResponse define response of AssetDetailService
type AssetDetailResponse struct {
	Details []struct {
		Asset            string `json:"asset"`
		AssetFullName    string `json:"assetFullName"`
		AmountFree       string `json:"amountFree"`
		ToBTC            string `json:"toBTC"`
		ToBNB            string `json:"toBNB"`
		ToBNBOffExchange string `json:"toBNBOffExchange"`
		Exchange         string `json:"exchange"`
	} `json:"details"`
	TotalTransferBtc   string `json:"totalTransferBtc"`
	TotalTransferBnb   string `json:"totalTransferBnb"`
	DribbletPercentage string `json:"dribbletPercentage"`
}

// Dust Transfer (USER_DATA)
const (
	dustTransferEndpoint = "/sapi/v1/asset/dust"
)

// DustTransferService dust transfer
type DustTransferService struct {
	c     *Client
	asset []string
}

// Asset set asset
func (s *DustTransferService) Asset(asset []string) *DustTransferService {
	s.asset = asset
	return s
}

func (s *DustTransferService) Do(ctx context.Context) (res *DustTransferResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: dustTransferEndpoint,
		secType:  secTypeSigned,
	}
	for _, a := range s.asset {
		r.addParam("asset", a)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(DustTransferResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DustTransferResponse define response of DustTransferService
type DustTransferResponse struct {
	TotalServiceCharge string                `json:"totalServiceCharge"`
	TotalTransfered    string                `json:"totalTransfered"`
	TransferResult     []*DustTransferResult `json:"transferResult"`
}

// DustTransferResult represents the result of a dust transfer.
type DustTransferResult struct {
	Amount              string `json:"amount"`
	FromAsset           string `json:"fromAsset"`
	OperateTime         int64  `json:"operateTime"`
	ServiceChargeAmount string `json:"serviceChargeAmount"`
	TranID              int64  `json:"tranId"`
	TransferedAmount    string `json:"transferedAmount"`
}

// Asset Dividend Record (USER_DATA)
const (
	assetDividendRecordEndpoint = "/sapi/v1/asset/assetDividend"
)

// AssetDividendRecordService asset dividend record
type AssetDividendRecordService struct {
	c         *Client
	asset     *string
	startTime *uint64
	endTime   *uint64
	limit     *int
}

// Asset set asset
func (s *AssetDividendRecordService) Asset(asset string) *AssetDividendRecordService {
	s.asset = &asset
	return s
}

// StartTime set startTime
func (s *AssetDividendRecordService) StartTime(startTime uint64) *AssetDividendRecordService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *AssetDividendRecordService) EndTime(endTime uint64) *AssetDividendRecordService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *AssetDividendRecordService) Limit(limit int) *AssetDividendRecordService {
	s.limit = &limit
	return s
}

func (s *AssetDividendRecordService) Do(ctx context.Context) (res *AssetDividendRecordResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: assetDividendRecordEndpoint,
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
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
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AssetDividendRecordResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AssetDividendRecordResponse define response of AssetDividendRecordService
type AssetDividendRecordResponse struct {
	Rows []struct {
		Id      int64  `json:"id"`
		Amount  string `json:"amount"`
		Asset   string `json:"asset"`
		DivTime uint64 `json:"divTime"`
		EnInfo  string `json:"enInfo"`
		TranId  int64  `json:"tranId"`
	} `json:"rows"`
	Total int64 `json:"total"`
}

// Asset Detail (USER_DATA)
const (
	assetDetailV2Endpoint = "/sapi/v1/asset/assetDetail"
)

// AssetDetailV2Service asset detail v2
type AssetDetailV2Service struct {
	c     *Client
	asset *string
}

// Asset set asset
func (s *AssetDetailV2Service) Asset(asset string) *AssetDetailV2Service {
	s.asset = &asset
	return s
}

func (s *AssetDetailV2Service) Do(ctx context.Context) (res *AssetDetailV2Response, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: assetDetailV2Endpoint,
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AssetDetailV2Response)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AssetDetailV2Response define response of AssetDetailV2Service
type AssetDetailV2Response struct {
	AssetDetail struct {
		MinWithdrawAmount string `json:"minWithdrawAmount"`
		DepositStatus     bool   `json:"depositStatus"`
		WithdrawFee       string `json:"withdrawFee"`
		WithdrawStatus    bool   `json:"withdrawStatus"`
		DepositTip        string `json:"depositTip"`
	} `json:"assetDetail"`
}

// Trade Fee (USER_DATA)
const (
	tradeFeeEndpoint = "/sapi/v1/asset/tradeFee"
)

// TradeFeeService trade fee
type TradeFeeService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *TradeFeeService) Symbol(symbol string) *TradeFeeService {
	s.symbol = &symbol
	return s
}

func (s *TradeFeeService) Do(ctx context.Context) (res []*TradeFeeResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: tradeFeeEndpoint,
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = make([]*TradeFeeResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TradeFeeResponse define response of TradeFeeService
type TradeFeeResponse struct {
	Symbol          string `json:"symbol"`
	MakerCommission string `json:"makerCommission"`
	TakerCommission string `json:"takerCommission"`
}

// User Universal Transfer (USER_DATA)
const (
	userUniversalTransferEndpoint = "/sapi/v1/asset/transfer"
)

// UserUniversalTransferService user universal transfer
type UserUniversalTransferService struct {
	c            *Client
	transferType string
	asset        string
	amount       float64
	fromSymbol   *string
	toSymbol     *string
}

// TransferType set transferType
func (s *UserUniversalTransferService) TransferType(transferType string) *UserUniversalTransferService {
	s.transferType = transferType
	return s
}

// Asset set asset
func (s *UserUniversalTransferService) Asset(asset string) *UserUniversalTransferService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *UserUniversalTransferService) Amount(amount float64) *UserUniversalTransferService {
	s.amount = amount
	return s
}

// FromSymbol set fromSymbol
func (s *UserUniversalTransferService) FromSymbol(fromSymbol string) *UserUniversalTransferService {
	s.fromSymbol = &fromSymbol
	return s
}

// ToSymbol set toSymbol
func (s *UserUniversalTransferService) ToSymbol(toSymbol string) *UserUniversalTransferService {
	s.toSymbol = &toSymbol
	return s
}

func (s *UserUniversalTransferService) Do(ctx context.Context) (res *UserUniversalTransferResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: userUniversalTransferEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("type", s.transferType)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	if s.fromSymbol != nil {
		r.setParam("fromSymbol", *s.fromSymbol)
	}
	if s.toSymbol != nil {
		r.setParam("toSymbol", *s.toSymbol)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(UserUniversalTransferResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UserUniversalTransferResponse define response of UserUniversalTransferService
type UserUniversalTransferResponse struct {
	TranId int64 `json:"tranId"`
}

// Query User Universal Transfer History (USER_DATA)
const (
	userUniversalTransferHistoryEndpoint = "/sapi/v1/asset/transfer"
)

// UserUniversalTransferHistoryService user universal transfer history
type UserUniversalTransferHistoryService struct {
	c            *Client
	transferType string
	startTime    *uint64
	endTime      *uint64
	current      *int
	size         *int
	fromSymbol   *string
	toSymbol     *string
}

// TransferType set transferType
func (s *UserUniversalTransferHistoryService) TransferType(transferType string) *UserUniversalTransferHistoryService {
	s.transferType = transferType
	return s
}

// StartTime set startTime
func (s *UserUniversalTransferHistoryService) StartTime(startTime uint64) *UserUniversalTransferHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *UserUniversalTransferHistoryService) EndTime(endTime uint64) *UserUniversalTransferHistoryService {
	s.endTime = &endTime
	return s
}

// Current set current
func (s *UserUniversalTransferHistoryService) Current(current int) *UserUniversalTransferHistoryService {
	s.current = &current
	return s
}

// Size set size
func (s *UserUniversalTransferHistoryService) Size(size int) *UserUniversalTransferHistoryService {
	s.size = &size
	return s
}

// FromSymbol set fromSymbol
func (s *UserUniversalTransferHistoryService) FromSymbol(fromSymbol string) *UserUniversalTransferHistoryService {
	s.fromSymbol = &fromSymbol
	return s
}

// ToSymbol set toSymbol
func (s *UserUniversalTransferHistoryService) ToSymbol(toSymbol string) *UserUniversalTransferHistoryService {
	s.toSymbol = &toSymbol
	return s
}

func (s *UserUniversalTransferHistoryService) Do(ctx context.Context) (res *UserUniversalTransferHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: userUniversalTransferHistoryEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("type", s.transferType)
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
	if s.fromSymbol != nil {
		r.setParam("fromSymbol", *s.fromSymbol)
	}
	if s.toSymbol != nil {
		r.setParam("toSymbol", *s.toSymbol)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(UserUniversalTransferHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UserUniversalTransferHistoryResponse define response of UserUniversalTransferHistoryService
type UserUniversalTransferHistoryResponse struct {
	Total int64 `json:"total"`
	Rows  []struct {
		Asset     string `json:"asset"`
		Amount    string `json:"amount"`
		Type      string `json:"type"`
		Status    string `json:"status"`
		TranId    int64  `json:"tranId"`
		Timestamp uint64 `json:"timestamp"`
	} `json:"rows"`
}

// Funding Wallet (USER_DATA)
const (
	fundingWalletEndpoint = "/sapi/v1/asset/get-funding-asset"
)

// FundingWalletService funding wallet
type FundingWalletService struct {
	c                *Client
	asset            *string
	needBtcValuation *string
}

// Asset set asset
func (s *FundingWalletService) Asset(asset string) *FundingWalletService {
	s.asset = &asset
	return s
}

// NeedBtcValuation set needBtcValuation
func (s *FundingWalletService) NeedBtcValuation(needBtcValuation string) *FundingWalletService {
	s.needBtcValuation = &needBtcValuation
	return s
}

func (s *FundingWalletService) Do(ctx context.Context) (res []*FundingWalletResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: fundingWalletEndpoint,
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.needBtcValuation != nil {
		r.setParam("needBtcValuation", *s.needBtcValuation)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = make([]*FundingWalletResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// FundingWalletResponse define response of FundingWalletService
type FundingWalletResponse struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	BtcValuation string `json:"btcValuation"`
}

// User Asset (USER_DATA)
const (
	userAssetEndpoint = "/sapi/v3/asset/getUserAsset"
)

// UserAssetService user asset
type UserAssetService struct {
	c                *Client
	asset            *string
	needBtcValuation *bool
}

// Asset set asset
func (s *UserAssetService) Asset(asset string) *UserAssetService {
	s.asset = &asset
	return s
}

// NeedBtcValuation set needBtcValuation
func (s *UserAssetService) NeedBtcValuation(needBtcValuation bool) *UserAssetService {
	s.needBtcValuation = &needBtcValuation
	return s
}

func (s *UserAssetService) Do(ctx context.Context) (res []*UserAssetResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: userAssetEndpoint,
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.needBtcValuation != nil {
		r.setParam("needBtcValuation", *s.needBtcValuation)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = make([]*UserAssetResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UserAssetResponse define response of UserAssetService
type UserAssetResponse struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	Ipoable      string `json:"ipoable"`
	BtcValuation string `json:"btcValuation"`
}

// BUSD Convert (TRADE)
const (
	bUSDConvertEndpoint = "/sapi/v1/asset/convert-transfer"
)

// BUSDConvertService BUSD convert
type BUSDConvertService struct {
	c            *Client
	clientTranId string
	asset        string
	amount       float64
	targetAsset  string
	accountType  *string
}

// ClientTranId set clientTranId
func (s *BUSDConvertService) ClientTranId(clientTranId string) *BUSDConvertService {
	s.clientTranId = clientTranId
	return s
}

// Asset set asset
func (s *BUSDConvertService) Asset(asset string) *BUSDConvertService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *BUSDConvertService) Amount(amount float64) *BUSDConvertService {
	s.amount = amount
	return s
}

// TargetAsset set targetAsset
func (s *BUSDConvertService) TargetAsset(targetAsset string) *BUSDConvertService {
	s.targetAsset = targetAsset
	return s
}

// AccountType set accountType
func (s *BUSDConvertService) AccountType(accountType string) *BUSDConvertService {
	s.accountType = &accountType
	return s
}

func (s *BUSDConvertService) Do(ctx context.Context) (res *BUSDConvertResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: bUSDConvertEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("clientTranId", s.clientTranId)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	r.setParam("targetAsset", s.targetAsset)
	if s.accountType != nil {
		r.setParam("accountType", *s.accountType)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(BUSDConvertResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// BUSDConvertResponse define response of BUSDConvertService
type BUSDConvertResponse struct {
	TranId int64  `json:"tranId"`
	Status string `json:"status"`
}

// BUSD Convert History (USER_DATA)
const (
	bUSDConvertHistoryEndpoint = "/sapi/v1/asset/convert-transfer/queryByPage"
)

// BUSDConvertHistoryService BUSD convert history
type BUSDConvertHistoryService struct {
	c            *Client
	tranId       *int64
	clientTranId *string
	asset        *string
	startTime    uint64
	endTime      uint64
	accountType  *string
	current      *int
	size         *int
}

// TranId set tranId
func (s *BUSDConvertHistoryService) TranId(tranId int64) *BUSDConvertHistoryService {
	s.tranId = &tranId
	return s
}

// ClientTranId set clientTranId
func (s *BUSDConvertHistoryService) ClientTranId(clientTranId string) *BUSDConvertHistoryService {
	s.clientTranId = &clientTranId
	return s
}

// Asset set asset
func (s *BUSDConvertHistoryService) Asset(asset string) *BUSDConvertHistoryService {
	s.asset = &asset
	return s
}

// StartTime set startTime
func (s *BUSDConvertHistoryService) StartTime(startTime uint64) *BUSDConvertHistoryService {
	s.startTime = startTime
	return s
}

// EndTime set endTime
func (s *BUSDConvertHistoryService) EndTime(endTime uint64) *BUSDConvertHistoryService {
	s.endTime = endTime
	return s
}

// AccountType set accountType
func (s *BUSDConvertHistoryService) AccountType(accountType string) *BUSDConvertHistoryService {
	s.accountType = &accountType
	return s
}

// Current set current
func (s *BUSDConvertHistoryService) Current(current int) *BUSDConvertHistoryService {
	s.current = &current
	return s
}

// Size set size
func (s *BUSDConvertHistoryService) Size(size int) *BUSDConvertHistoryService {
	s.size = &size
	return s
}

func (s *BUSDConvertHistoryService) Do(ctx context.Context) (res *BUSDConvertHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: bUSDConvertHistoryEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("startTime", s.startTime)
	r.setParam("endTime", s.endTime)
	if s.tranId != nil {
		r.setParam("tranId", *s.tranId)
	}
	if s.clientTranId != nil {
		r.setParam("clientTranId", *s.clientTranId)
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.accountType != nil {
		r.setParam("accountType", *s.accountType)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(BUSDConvertHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// BUSDConvertHistoryResponse define response of BUSDConvertHistoryService
type BUSDConvertHistoryResponse struct {
	Total int32 `json:"total"`
	Rows  []struct {
		TranId         int64  `json:"tranId"`
		Type           int32  `json:"type"`
		Time           uint64 `json:"time"`
		DeductedAsset  string `json:"deductedAsset"`
		DeductedAmount string `json:"deductedAmount"`
		TargetAsset    string `json:"targetAsset"`
		TargetAmount   string `json:"targetAmount"`
		Status         string `json:"status"`
		AccountType    string `json:"accountType"`
	} `json:"rows"`
}

// Get Cloud-Mining payment and refund history (USER_DATA)
const (
	cloudMiningPaymentHistoryEndpoint = "/sapi/v1/asset/ledger-transfer/cloud-mining/queryByPage"
)

// CloudMiningPaymentHistoryService cloud mining payment history
type CloudMiningPaymentHistoryService struct {
	c            *Client
	tranid       *int64
	clientTranId *string
	asset        *string
	startTime    uint64
	endTime      uint64
	current      *int
	size         *int
}

// Tranid set tranid
func (s *CloudMiningPaymentHistoryService) Tranid(tranid int64) *CloudMiningPaymentHistoryService {
	s.tranid = &tranid
	return s
}

// ClientTranId set clientTranId
func (s *CloudMiningPaymentHistoryService) ClientTranId(clientTranId string) *CloudMiningPaymentHistoryService {
	s.clientTranId = &clientTranId
	return s
}

// Asset set asset
func (s *CloudMiningPaymentHistoryService) Asset(asset string) *CloudMiningPaymentHistoryService {
	s.asset = &asset
	return s
}

// StartTime set startTime
func (s *CloudMiningPaymentHistoryService) StartTime(startTime uint64) *CloudMiningPaymentHistoryService {
	s.startTime = startTime
	return s
}

// EndTime set endTime
func (s *CloudMiningPaymentHistoryService) EndTime(endTime uint64) *CloudMiningPaymentHistoryService {
	s.endTime = endTime
	return s
}

// Current set current
func (s *CloudMiningPaymentHistoryService) Current(current int) *CloudMiningPaymentHistoryService {
	s.current = &current
	return s
}

// Size set size
func (s *CloudMiningPaymentHistoryService) Size(size int) *CloudMiningPaymentHistoryService {
	s.size = &size
	return s
}

func (s *CloudMiningPaymentHistoryService) Do(ctx context.Context) (res *CloudMiningPaymentHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: cloudMiningPaymentHistoryEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("startTime", s.startTime)
	r.setParam("endTime", s.endTime)
	if s.tranid != nil {
		r.setParam("tranId", *s.tranid)
	}
	if s.clientTranId != nil {
		r.setParam("clientTranId", *s.clientTranId)
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
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(CloudMiningPaymentHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CloudMiningPaymentHistoryResponse define response of CloudMiningPaymentHistoryService
type CloudMiningPaymentHistoryResponse struct {
	Total int32 `json:"total"`
	Rows  []struct {
		CreateTime uint64 `json:"createTime"`
		TranId     int64  `json:"tranId"`
		Type       int32  `json:"type"`
		Asset      string `json:"asset"`
		Amount     string `json:"amount"`
		Status     string `json:"status"`
	} `json:"rows"`
}

// Get API Key Permission (USER_DATA)
const (
	apiKeyPermissionEndpoint = "/sapi/v1/account/apiRestrictions"
)

// APIKeyPermissionService get api key permission
type APIKeyPermissionService struct {
	c *Client
}

func (s *APIKeyPermissionService) Do(ctx context.Context) (res *APIKeyPermissionResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: apiKeyPermissionEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(APIKeyPermissionResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// APIKeyPermissionResponse define response of APIKeyPermissionService
type APIKeyPermissionResponse struct {
	IPRestrict                     bool   `json:"ipRestrict"`
	CreateTime                     uint64 `json:"createTime"`
	EnableWithdrawals              bool   `json:"enableWithdrawals"`
	EnableInternalTransfer         bool   `json:"enableInternalTransfer"`
	PermitsUniversalTransfer       bool   `json:"permitsUniversalTransfer"`
	EnableVanillaOptions           bool   `json:"enableVanillaOptions"`
	EnableReading                  bool   `json:"enableReading"`
	EnableFutures                  bool   `json:"enableFutures"`
	EnableMargin                   bool   `json:"enableMargin"`
	EnableSpotAndMarginTrading     bool   `json:"enableSpotAndMarginTrading"`
	TradingAuthorityExpirationTime uint64 `json:"tradingAuthorityExpirationTime"`
}

// Query auto-converting stable coins (USER_DATA)
const (
	autoConvertStableCoinEndpoint = "/sapi/v1/capital/contract/convertible-coins"
)

// AutoConvertStableCoinService auto convert stable coin
type AutoConvertStableCoinService struct {
	c *Client
}

func (s *AutoConvertStableCoinService) Do(ctx context.Context) (res *AutoConvertStableCoinResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: autoConvertStableCoinEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AutoConvertStableCoinResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AutoConvertStableCoinResponse define response of AutoConvertStableCoinService
type AutoConvertStableCoinResponse struct {
	ConvertEnabled bool `json:"convertEnabled"`
	Coins          []struct {
		Asset string `json:"coin"`
	} `json:"coins"`
	ExchangeRates []struct {
		Asset string `json:"coin"`
	} `json:"exchangeRates"`
}
