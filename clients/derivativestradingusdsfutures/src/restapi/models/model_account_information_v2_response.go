/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountInformationV2Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountInformationV2Response{}

// AccountInformationV2Response struct for AccountInformationV2Response
type AccountInformationV2Response struct {
	FeeTier                     *int64                                       `json:"feeTier,omitempty"`
	FeeBurn                     *bool                                        `json:"feeBurn,omitempty"`
	CanDeposit                  *bool                                        `json:"canDeposit,omitempty"`
	CanWithdraw                 *bool                                        `json:"canWithdraw,omitempty"`
	UpdateTime                  *int64                                       `json:"updateTime,omitempty"`
	MultiAssetsMargin           *bool                                        `json:"multiAssetsMargin,omitempty"`
	TradeGroupId                *int64                                       `json:"tradeGroupId,omitempty"`
	TotalInitialMargin          *string                                      `json:"totalInitialMargin,omitempty"`
	TotalMaintMargin            *string                                      `json:"totalMaintMargin,omitempty"`
	TotalWalletBalance          *string                                      `json:"totalWalletBalance,omitempty"`
	TotalUnrealizedProfit       *string                                      `json:"totalUnrealizedProfit,omitempty"`
	TotalMarginBalance          *string                                      `json:"totalMarginBalance,omitempty"`
	TotalPositionInitialMargin  *string                                      `json:"totalPositionInitialMargin,omitempty"`
	TotalOpenOrderInitialMargin *string                                      `json:"totalOpenOrderInitialMargin,omitempty"`
	TotalCrossWalletBalance     *string                                      `json:"totalCrossWalletBalance,omitempty"`
	TotalCrossUnPnl             *string                                      `json:"totalCrossUnPnl,omitempty"`
	AvailableBalance            *string                                      `json:"availableBalance,omitempty"`
	MaxWithdrawAmount           *string                                      `json:"maxWithdrawAmount,omitempty"`
	Assets                      []AccountInformationV2ResponseAssetsInner    `json:"assets,omitempty"`
	Positions                   []AccountInformationV2ResponsePositionsInner `json:"positions,omitempty"`
	CanTrade                    *bool                                        `json:"canTrade,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _AccountInformationV2Response AccountInformationV2Response

// NewAccountInformationV2Response instantiates a new AccountInformationV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInformationV2Response() *AccountInformationV2Response {
	this := AccountInformationV2Response{}
	return &this
}

// NewAccountInformationV2ResponseWithDefaults instantiates a new AccountInformationV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInformationV2ResponseWithDefaults() *AccountInformationV2Response {
	this := AccountInformationV2Response{}
	return &this
}

// GetFeeTier returns the FeeTier field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetFeeTier() int64 {
	if o == nil || common.IsNil(o.FeeTier) {
		var ret int64
		return ret
	}
	return *o.FeeTier
}

// GetFeeTierOk returns a tuple with the FeeTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetFeeTierOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FeeTier) {
		return nil, false
	}
	return o.FeeTier, true
}

// HasFeeTier returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasFeeTier() bool {
	if o != nil && !common.IsNil(o.FeeTier) {
		return true
	}

	return false
}

// SetFeeTier gets a reference to the given int64 and assigns it to the FeeTier field.
func (o *AccountInformationV2Response) SetFeeTier(v int64) {
	o.FeeTier = &v
}

// GetFeeBurn returns the FeeBurn field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetFeeBurn() bool {
	if o == nil || common.IsNil(o.FeeBurn) {
		var ret bool
		return ret
	}
	return *o.FeeBurn
}

// GetFeeBurnOk returns a tuple with the FeeBurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetFeeBurnOk() (*bool, bool) {
	if o == nil || common.IsNil(o.FeeBurn) {
		return nil, false
	}
	return o.FeeBurn, true
}

// HasFeeBurn returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasFeeBurn() bool {
	if o != nil && !common.IsNil(o.FeeBurn) {
		return true
	}

	return false
}

// SetFeeBurn gets a reference to the given bool and assigns it to the FeeBurn field.
func (o *AccountInformationV2Response) SetFeeBurn(v bool) {
	o.FeeBurn = &v
}

// GetCanDeposit returns the CanDeposit field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetCanDeposit() bool {
	if o == nil || common.IsNil(o.CanDeposit) {
		var ret bool
		return ret
	}
	return *o.CanDeposit
}

// GetCanDepositOk returns a tuple with the CanDeposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetCanDepositOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanDeposit) {
		return nil, false
	}
	return o.CanDeposit, true
}

// HasCanDeposit returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasCanDeposit() bool {
	if o != nil && !common.IsNil(o.CanDeposit) {
		return true
	}

	return false
}

// SetCanDeposit gets a reference to the given bool and assigns it to the CanDeposit field.
func (o *AccountInformationV2Response) SetCanDeposit(v bool) {
	o.CanDeposit = &v
}

// GetCanWithdraw returns the CanWithdraw field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetCanWithdraw() bool {
	if o == nil || common.IsNil(o.CanWithdraw) {
		var ret bool
		return ret
	}
	return *o.CanWithdraw
}

// GetCanWithdrawOk returns a tuple with the CanWithdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetCanWithdrawOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanWithdraw) {
		return nil, false
	}
	return o.CanWithdraw, true
}

// HasCanWithdraw returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasCanWithdraw() bool {
	if o != nil && !common.IsNil(o.CanWithdraw) {
		return true
	}

	return false
}

// SetCanWithdraw gets a reference to the given bool and assigns it to the CanWithdraw field.
func (o *AccountInformationV2Response) SetCanWithdraw(v bool) {
	o.CanWithdraw = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AccountInformationV2Response) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetMultiAssetsMargin returns the MultiAssetsMargin field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetMultiAssetsMargin() bool {
	if o == nil || common.IsNil(o.MultiAssetsMargin) {
		var ret bool
		return ret
	}
	return *o.MultiAssetsMargin
}

// GetMultiAssetsMarginOk returns a tuple with the MultiAssetsMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetMultiAssetsMarginOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MultiAssetsMargin) {
		return nil, false
	}
	return o.MultiAssetsMargin, true
}

// HasMultiAssetsMargin returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasMultiAssetsMargin() bool {
	if o != nil && !common.IsNil(o.MultiAssetsMargin) {
		return true
	}

	return false
}

// SetMultiAssetsMargin gets a reference to the given bool and assigns it to the MultiAssetsMargin field.
func (o *AccountInformationV2Response) SetMultiAssetsMargin(v bool) {
	o.MultiAssetsMargin = &v
}

// GetTradeGroupId returns the TradeGroupId field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetTradeGroupId() int64 {
	if o == nil || common.IsNil(o.TradeGroupId) {
		var ret int64
		return ret
	}
	return *o.TradeGroupId
}

// GetTradeGroupIdOk returns a tuple with the TradeGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetTradeGroupIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeGroupId) {
		return nil, false
	}
	return o.TradeGroupId, true
}

// HasTradeGroupId returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasTradeGroupId() bool {
	if o != nil && !common.IsNil(o.TradeGroupId) {
		return true
	}

	return false
}

// SetTradeGroupId gets a reference to the given int64 and assigns it to the TradeGroupId field.
func (o *AccountInformationV2Response) SetTradeGroupId(v int64) {
	o.TradeGroupId = &v
}

// GetTotalInitialMargin returns the TotalInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetTotalInitialMargin() string {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalInitialMargin
}

// GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetTotalInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		return nil, false
	}
	return o.TotalInitialMargin, true
}

// HasTotalInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasTotalInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalInitialMargin) {
		return true
	}

	return false
}

// SetTotalInitialMargin gets a reference to the given string and assigns it to the TotalInitialMargin field.
func (o *AccountInformationV2Response) SetTotalInitialMargin(v string) {
	o.TotalInitialMargin = &v
}

// GetTotalMaintMargin returns the TotalMaintMargin field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetTotalMaintMargin() string {
	if o == nil || common.IsNil(o.TotalMaintMargin) {
		var ret string
		return ret
	}
	return *o.TotalMaintMargin
}

// GetTotalMaintMarginOk returns a tuple with the TotalMaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetTotalMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMaintMargin) {
		return nil, false
	}
	return o.TotalMaintMargin, true
}

// HasTotalMaintMargin returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasTotalMaintMargin() bool {
	if o != nil && !common.IsNil(o.TotalMaintMargin) {
		return true
	}

	return false
}

// SetTotalMaintMargin gets a reference to the given string and assigns it to the TotalMaintMargin field.
func (o *AccountInformationV2Response) SetTotalMaintMargin(v string) {
	o.TotalMaintMargin = &v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetTotalWalletBalance() string {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		return nil, false
	}
	return o.TotalWalletBalance, true
}

// HasTotalWalletBalance returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasTotalWalletBalance() bool {
	if o != nil && !common.IsNil(o.TotalWalletBalance) {
		return true
	}

	return false
}

// SetTotalWalletBalance gets a reference to the given string and assigns it to the TotalWalletBalance field.
func (o *AccountInformationV2Response) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = &v
}

// GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetTotalUnrealizedProfit() string {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.TotalUnrealizedProfit
}

// GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetTotalUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		return nil, false
	}
	return o.TotalUnrealizedProfit, true
}

// HasTotalUnrealizedProfit returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasTotalUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.TotalUnrealizedProfit) {
		return true
	}

	return false
}

// SetTotalUnrealizedProfit gets a reference to the given string and assigns it to the TotalUnrealizedProfit field.
func (o *AccountInformationV2Response) SetTotalUnrealizedProfit(v string) {
	o.TotalUnrealizedProfit = &v
}

// GetTotalMarginBalance returns the TotalMarginBalance field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetTotalMarginBalance() string {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		var ret string
		return ret
	}
	return *o.TotalMarginBalance
}

// GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetTotalMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		return nil, false
	}
	return o.TotalMarginBalance, true
}

// HasTotalMarginBalance returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasTotalMarginBalance() bool {
	if o != nil && !common.IsNil(o.TotalMarginBalance) {
		return true
	}

	return false
}

// SetTotalMarginBalance gets a reference to the given string and assigns it to the TotalMarginBalance field.
func (o *AccountInformationV2Response) SetTotalMarginBalance(v string) {
	o.TotalMarginBalance = &v
}

// GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetTotalPositionInitialMargin() string {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalPositionInitialMargin
}

// GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetTotalPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		return nil, false
	}
	return o.TotalPositionInitialMargin, true
}

// HasTotalPositionInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasTotalPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalPositionInitialMargin) {
		return true
	}

	return false
}

// SetTotalPositionInitialMargin gets a reference to the given string and assigns it to the TotalPositionInitialMargin field.
func (o *AccountInformationV2Response) SetTotalPositionInitialMargin(v string) {
	o.TotalPositionInitialMargin = &v
}

// GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetTotalOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalOpenOrderInitialMargin
}

// GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetTotalOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		return nil, false
	}
	return o.TotalOpenOrderInitialMargin, true
}

// HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasTotalOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalOpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetTotalOpenOrderInitialMargin gets a reference to the given string and assigns it to the TotalOpenOrderInitialMargin field.
func (o *AccountInformationV2Response) SetTotalOpenOrderInitialMargin(v string) {
	o.TotalOpenOrderInitialMargin = &v
}

// GetTotalCrossWalletBalance returns the TotalCrossWalletBalance field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetTotalCrossWalletBalance() string {
	if o == nil || common.IsNil(o.TotalCrossWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalCrossWalletBalance
}

// GetTotalCrossWalletBalanceOk returns a tuple with the TotalCrossWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetTotalCrossWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalCrossWalletBalance) {
		return nil, false
	}
	return o.TotalCrossWalletBalance, true
}

// HasTotalCrossWalletBalance returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasTotalCrossWalletBalance() bool {
	if o != nil && !common.IsNil(o.TotalCrossWalletBalance) {
		return true
	}

	return false
}

// SetTotalCrossWalletBalance gets a reference to the given string and assigns it to the TotalCrossWalletBalance field.
func (o *AccountInformationV2Response) SetTotalCrossWalletBalance(v string) {
	o.TotalCrossWalletBalance = &v
}

// GetTotalCrossUnPnl returns the TotalCrossUnPnl field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetTotalCrossUnPnl() string {
	if o == nil || common.IsNil(o.TotalCrossUnPnl) {
		var ret string
		return ret
	}
	return *o.TotalCrossUnPnl
}

// GetTotalCrossUnPnlOk returns a tuple with the TotalCrossUnPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetTotalCrossUnPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalCrossUnPnl) {
		return nil, false
	}
	return o.TotalCrossUnPnl, true
}

// HasTotalCrossUnPnl returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasTotalCrossUnPnl() bool {
	if o != nil && !common.IsNil(o.TotalCrossUnPnl) {
		return true
	}

	return false
}

// SetTotalCrossUnPnl gets a reference to the given string and assigns it to the TotalCrossUnPnl field.
func (o *AccountInformationV2Response) SetTotalCrossUnPnl(v string) {
	o.TotalCrossUnPnl = &v
}

// GetAvailableBalance returns the AvailableBalance field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetAvailableBalance() string {
	if o == nil || common.IsNil(o.AvailableBalance) {
		var ret string
		return ret
	}
	return *o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetAvailableBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvailableBalance) {
		return nil, false
	}
	return o.AvailableBalance, true
}

// HasAvailableBalance returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasAvailableBalance() bool {
	if o != nil && !common.IsNil(o.AvailableBalance) {
		return true
	}

	return false
}

// SetAvailableBalance gets a reference to the given string and assigns it to the AvailableBalance field.
func (o *AccountInformationV2Response) SetAvailableBalance(v string) {
	o.AvailableBalance = &v
}

// GetMaxWithdrawAmount returns the MaxWithdrawAmount field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetMaxWithdrawAmount() string {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.MaxWithdrawAmount
}

// GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetMaxWithdrawAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		return nil, false
	}
	return o.MaxWithdrawAmount, true
}

// HasMaxWithdrawAmount returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasMaxWithdrawAmount() bool {
	if o != nil && !common.IsNil(o.MaxWithdrawAmount) {
		return true
	}

	return false
}

// SetMaxWithdrawAmount gets a reference to the given string and assigns it to the MaxWithdrawAmount field.
func (o *AccountInformationV2Response) SetMaxWithdrawAmount(v string) {
	o.MaxWithdrawAmount = &v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetAssets() []AccountInformationV2ResponseAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []AccountInformationV2ResponseAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetAssetsOk() ([]AccountInformationV2ResponseAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []AccountInformationV2ResponseAssetsInner and assigns it to the Assets field.
func (o *AccountInformationV2Response) SetAssets(v []AccountInformationV2ResponseAssetsInner) {
	o.Assets = v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetPositions() []AccountInformationV2ResponsePositionsInner {
	if o == nil || common.IsNil(o.Positions) {
		var ret []AccountInformationV2ResponsePositionsInner
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetPositionsOk() ([]AccountInformationV2ResponsePositionsInner, bool) {
	if o == nil || common.IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasPositions() bool {
	if o != nil && !common.IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given []AccountInformationV2ResponsePositionsInner and assigns it to the Positions field.
func (o *AccountInformationV2Response) SetPositions(v []AccountInformationV2ResponsePositionsInner) {
	o.Positions = v
}

// GetCanTrade returns the CanTrade field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetCanTrade() bool {
	if o == nil || common.IsNil(o.CanTrade) {
		var ret bool
		return ret
	}
	return *o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetCanTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanTrade) {
		return nil, false
	}
	return o.CanTrade, true
}

// HasCanTrade returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasCanTrade() bool {
	if o != nil && !common.IsNil(o.CanTrade) {
		return true
	}

	return false
}

// SetCanTrade gets a reference to the given bool and assigns it to the CanTrade field.
func (o *AccountInformationV2Response) SetCanTrade(v bool) {
	o.CanTrade = &v
}

func (o AccountInformationV2Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInformationV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FeeTier) {
		toSerialize["feeTier"] = o.FeeTier
	}
	if !common.IsNil(o.FeeBurn) {
		toSerialize["feeBurn"] = o.FeeBurn
	}
	if !common.IsNil(o.CanDeposit) {
		toSerialize["canDeposit"] = o.CanDeposit
	}
	if !common.IsNil(o.CanWithdraw) {
		toSerialize["canWithdraw"] = o.CanWithdraw
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.MultiAssetsMargin) {
		toSerialize["multiAssetsMargin"] = o.MultiAssetsMargin
	}
	if !common.IsNil(o.TradeGroupId) {
		toSerialize["tradeGroupId"] = o.TradeGroupId
	}
	if !common.IsNil(o.TotalInitialMargin) {
		toSerialize["totalInitialMargin"] = o.TotalInitialMargin
	}
	if !common.IsNil(o.TotalMaintMargin) {
		toSerialize["totalMaintMargin"] = o.TotalMaintMargin
	}
	if !common.IsNil(o.TotalWalletBalance) {
		toSerialize["totalWalletBalance"] = o.TotalWalletBalance
	}
	if !common.IsNil(o.TotalUnrealizedProfit) {
		toSerialize["totalUnrealizedProfit"] = o.TotalUnrealizedProfit
	}
	if !common.IsNil(o.TotalMarginBalance) {
		toSerialize["totalMarginBalance"] = o.TotalMarginBalance
	}
	if !common.IsNil(o.TotalPositionInitialMargin) {
		toSerialize["totalPositionInitialMargin"] = o.TotalPositionInitialMargin
	}
	if !common.IsNil(o.TotalOpenOrderInitialMargin) {
		toSerialize["totalOpenOrderInitialMargin"] = o.TotalOpenOrderInitialMargin
	}
	if !common.IsNil(o.TotalCrossWalletBalance) {
		toSerialize["totalCrossWalletBalance"] = o.TotalCrossWalletBalance
	}
	if !common.IsNil(o.TotalCrossUnPnl) {
		toSerialize["totalCrossUnPnl"] = o.TotalCrossUnPnl
	}
	if !common.IsNil(o.AvailableBalance) {
		toSerialize["availableBalance"] = o.AvailableBalance
	}
	if !common.IsNil(o.MaxWithdrawAmount) {
		toSerialize["maxWithdrawAmount"] = o.MaxWithdrawAmount
	}
	if !common.IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !common.IsNil(o.Positions) {
		toSerialize["positions"] = o.Positions
	}
	if !common.IsNil(o.CanTrade) {
		toSerialize["canTrade"] = o.CanTrade
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountInformationV2Response) UnmarshalJSON(data []byte) (err error) {
	varAccountInformationV2Response := _AccountInformationV2Response{}

	err = json.Unmarshal(data, &varAccountInformationV2Response)

	if err != nil {
		return err
	}

	*o = AccountInformationV2Response(varAccountInformationV2Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "feeTier")
		delete(additionalProperties, "feeBurn")
		delete(additionalProperties, "canDeposit")
		delete(additionalProperties, "canWithdraw")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "multiAssetsMargin")
		delete(additionalProperties, "tradeGroupId")
		delete(additionalProperties, "totalInitialMargin")
		delete(additionalProperties, "totalMaintMargin")
		delete(additionalProperties, "totalWalletBalance")
		delete(additionalProperties, "totalUnrealizedProfit")
		delete(additionalProperties, "totalMarginBalance")
		delete(additionalProperties, "totalPositionInitialMargin")
		delete(additionalProperties, "totalOpenOrderInitialMargin")
		delete(additionalProperties, "totalCrossWalletBalance")
		delete(additionalProperties, "totalCrossUnPnl")
		delete(additionalProperties, "availableBalance")
		delete(additionalProperties, "maxWithdrawAmount")
		delete(additionalProperties, "assets")
		delete(additionalProperties, "positions")
		delete(additionalProperties, "canTrade")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountInformationV2Response struct {
	value *AccountInformationV2Response
	isSet bool
}

func (v NullableAccountInformationV2Response) Get() *AccountInformationV2Response {
	return v.value
}

func (v *NullableAccountInformationV2Response) Set(val *AccountInformationV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInformationV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInformationV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInformationV2Response(val *AccountInformationV2Response) *NullableAccountInformationV2Response {
	return &NullableAccountInformationV2Response{value: val, isSet: true}
}

func (v NullableAccountInformationV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInformationV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
