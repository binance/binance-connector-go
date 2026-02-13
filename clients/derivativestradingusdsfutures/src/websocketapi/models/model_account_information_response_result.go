/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountInformationResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountInformationResponseResult{}

// AccountInformationResponseResult struct for AccountInformationResponseResult
type AccountInformationResponseResult struct {
	FeeTier                     *int64                                           `json:"feeTier,omitempty"`
	CanTrade                    *bool                                            `json:"canTrade,omitempty"`
	CanDeposit                  *bool                                            `json:"canDeposit,omitempty"`
	CanWithdraw                 *bool                                            `json:"canWithdraw,omitempty"`
	UpdateTime                  *int64                                           `json:"updateTime,omitempty"`
	MultiAssetsMargin           *bool                                            `json:"multiAssetsMargin,omitempty"`
	TradeGroupId                *int64                                           `json:"tradeGroupId,omitempty"`
	TotalInitialMargin          *string                                          `json:"totalInitialMargin,omitempty"`
	TotalMaintMargin            *string                                          `json:"totalMaintMargin,omitempty"`
	TotalWalletBalance          *string                                          `json:"totalWalletBalance,omitempty"`
	TotalUnrealizedProfit       *string                                          `json:"totalUnrealizedProfit,omitempty"`
	TotalMarginBalance          *string                                          `json:"totalMarginBalance,omitempty"`
	TotalPositionInitialMargin  *string                                          `json:"totalPositionInitialMargin,omitempty"`
	TotalOpenOrderInitialMargin *string                                          `json:"totalOpenOrderInitialMargin,omitempty"`
	TotalCrossWalletBalance     *string                                          `json:"totalCrossWalletBalance,omitempty"`
	TotalCrossUnPnl             *string                                          `json:"totalCrossUnPnl,omitempty"`
	AvailableBalance            *string                                          `json:"availableBalance,omitempty"`
	MaxWithdrawAmount           *string                                          `json:"maxWithdrawAmount,omitempty"`
	Assets                      []AccountInformationResponseResultAssetsInner    `json:"assets,omitempty"`
	Positions                   []AccountInformationResponseResultPositionsInner `json:"positions,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _AccountInformationResponseResult AccountInformationResponseResult

// NewAccountInformationResponseResult instantiates a new AccountInformationResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInformationResponseResult() *AccountInformationResponseResult {
	this := AccountInformationResponseResult{}
	return &this
}

// NewAccountInformationResponseResultWithDefaults instantiates a new AccountInformationResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInformationResponseResultWithDefaults() *AccountInformationResponseResult {
	this := AccountInformationResponseResult{}
	return &this
}

// GetFeeTier returns the FeeTier field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetFeeTier() int64 {
	if o == nil || common.IsNil(o.FeeTier) {
		var ret int64
		return ret
	}
	return *o.FeeTier
}

// GetFeeTierOk returns a tuple with the FeeTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetFeeTierOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FeeTier) {
		return nil, false
	}
	return o.FeeTier, true
}

// HasFeeTier returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasFeeTier() bool {
	if o != nil && !common.IsNil(o.FeeTier) {
		return true
	}

	return false
}

// SetFeeTier gets a reference to the given int64 and assigns it to the FeeTier field.
func (o *AccountInformationResponseResult) SetFeeTier(v int64) {
	o.FeeTier = &v
}

// GetCanTrade returns the CanTrade field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetCanTrade() bool {
	if o == nil || common.IsNil(o.CanTrade) {
		var ret bool
		return ret
	}
	return *o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetCanTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanTrade) {
		return nil, false
	}
	return o.CanTrade, true
}

// HasCanTrade returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasCanTrade() bool {
	if o != nil && !common.IsNil(o.CanTrade) {
		return true
	}

	return false
}

// SetCanTrade gets a reference to the given bool and assigns it to the CanTrade field.
func (o *AccountInformationResponseResult) SetCanTrade(v bool) {
	o.CanTrade = &v
}

// GetCanDeposit returns the CanDeposit field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetCanDeposit() bool {
	if o == nil || common.IsNil(o.CanDeposit) {
		var ret bool
		return ret
	}
	return *o.CanDeposit
}

// GetCanDepositOk returns a tuple with the CanDeposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetCanDepositOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanDeposit) {
		return nil, false
	}
	return o.CanDeposit, true
}

// HasCanDeposit returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasCanDeposit() bool {
	if o != nil && !common.IsNil(o.CanDeposit) {
		return true
	}

	return false
}

// SetCanDeposit gets a reference to the given bool and assigns it to the CanDeposit field.
func (o *AccountInformationResponseResult) SetCanDeposit(v bool) {
	o.CanDeposit = &v
}

// GetCanWithdraw returns the CanWithdraw field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetCanWithdraw() bool {
	if o == nil || common.IsNil(o.CanWithdraw) {
		var ret bool
		return ret
	}
	return *o.CanWithdraw
}

// GetCanWithdrawOk returns a tuple with the CanWithdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetCanWithdrawOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanWithdraw) {
		return nil, false
	}
	return o.CanWithdraw, true
}

// HasCanWithdraw returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasCanWithdraw() bool {
	if o != nil && !common.IsNil(o.CanWithdraw) {
		return true
	}

	return false
}

// SetCanWithdraw gets a reference to the given bool and assigns it to the CanWithdraw field.
func (o *AccountInformationResponseResult) SetCanWithdraw(v bool) {
	o.CanWithdraw = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AccountInformationResponseResult) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetMultiAssetsMargin returns the MultiAssetsMargin field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetMultiAssetsMargin() bool {
	if o == nil || common.IsNil(o.MultiAssetsMargin) {
		var ret bool
		return ret
	}
	return *o.MultiAssetsMargin
}

// GetMultiAssetsMarginOk returns a tuple with the MultiAssetsMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetMultiAssetsMarginOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MultiAssetsMargin) {
		return nil, false
	}
	return o.MultiAssetsMargin, true
}

// HasMultiAssetsMargin returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasMultiAssetsMargin() bool {
	if o != nil && !common.IsNil(o.MultiAssetsMargin) {
		return true
	}

	return false
}

// SetMultiAssetsMargin gets a reference to the given bool and assigns it to the MultiAssetsMargin field.
func (o *AccountInformationResponseResult) SetMultiAssetsMargin(v bool) {
	o.MultiAssetsMargin = &v
}

// GetTradeGroupId returns the TradeGroupId field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetTradeGroupId() int64 {
	if o == nil || common.IsNil(o.TradeGroupId) {
		var ret int64
		return ret
	}
	return *o.TradeGroupId
}

// GetTradeGroupIdOk returns a tuple with the TradeGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetTradeGroupIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeGroupId) {
		return nil, false
	}
	return o.TradeGroupId, true
}

// HasTradeGroupId returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasTradeGroupId() bool {
	if o != nil && !common.IsNil(o.TradeGroupId) {
		return true
	}

	return false
}

// SetTradeGroupId gets a reference to the given int64 and assigns it to the TradeGroupId field.
func (o *AccountInformationResponseResult) SetTradeGroupId(v int64) {
	o.TradeGroupId = &v
}

// GetTotalInitialMargin returns the TotalInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetTotalInitialMargin() string {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalInitialMargin
}

// GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetTotalInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		return nil, false
	}
	return o.TotalInitialMargin, true
}

// HasTotalInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasTotalInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalInitialMargin) {
		return true
	}

	return false
}

// SetTotalInitialMargin gets a reference to the given string and assigns it to the TotalInitialMargin field.
func (o *AccountInformationResponseResult) SetTotalInitialMargin(v string) {
	o.TotalInitialMargin = &v
}

// GetTotalMaintMargin returns the TotalMaintMargin field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetTotalMaintMargin() string {
	if o == nil || common.IsNil(o.TotalMaintMargin) {
		var ret string
		return ret
	}
	return *o.TotalMaintMargin
}

// GetTotalMaintMarginOk returns a tuple with the TotalMaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetTotalMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMaintMargin) {
		return nil, false
	}
	return o.TotalMaintMargin, true
}

// HasTotalMaintMargin returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasTotalMaintMargin() bool {
	if o != nil && !common.IsNil(o.TotalMaintMargin) {
		return true
	}

	return false
}

// SetTotalMaintMargin gets a reference to the given string and assigns it to the TotalMaintMargin field.
func (o *AccountInformationResponseResult) SetTotalMaintMargin(v string) {
	o.TotalMaintMargin = &v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetTotalWalletBalance() string {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		return nil, false
	}
	return o.TotalWalletBalance, true
}

// HasTotalWalletBalance returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasTotalWalletBalance() bool {
	if o != nil && !common.IsNil(o.TotalWalletBalance) {
		return true
	}

	return false
}

// SetTotalWalletBalance gets a reference to the given string and assigns it to the TotalWalletBalance field.
func (o *AccountInformationResponseResult) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = &v
}

// GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetTotalUnrealizedProfit() string {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.TotalUnrealizedProfit
}

// GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetTotalUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		return nil, false
	}
	return o.TotalUnrealizedProfit, true
}

// HasTotalUnrealizedProfit returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasTotalUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.TotalUnrealizedProfit) {
		return true
	}

	return false
}

// SetTotalUnrealizedProfit gets a reference to the given string and assigns it to the TotalUnrealizedProfit field.
func (o *AccountInformationResponseResult) SetTotalUnrealizedProfit(v string) {
	o.TotalUnrealizedProfit = &v
}

// GetTotalMarginBalance returns the TotalMarginBalance field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetTotalMarginBalance() string {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		var ret string
		return ret
	}
	return *o.TotalMarginBalance
}

// GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetTotalMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		return nil, false
	}
	return o.TotalMarginBalance, true
}

// HasTotalMarginBalance returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasTotalMarginBalance() bool {
	if o != nil && !common.IsNil(o.TotalMarginBalance) {
		return true
	}

	return false
}

// SetTotalMarginBalance gets a reference to the given string and assigns it to the TotalMarginBalance field.
func (o *AccountInformationResponseResult) SetTotalMarginBalance(v string) {
	o.TotalMarginBalance = &v
}

// GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetTotalPositionInitialMargin() string {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalPositionInitialMargin
}

// GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetTotalPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		return nil, false
	}
	return o.TotalPositionInitialMargin, true
}

// HasTotalPositionInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasTotalPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalPositionInitialMargin) {
		return true
	}

	return false
}

// SetTotalPositionInitialMargin gets a reference to the given string and assigns it to the TotalPositionInitialMargin field.
func (o *AccountInformationResponseResult) SetTotalPositionInitialMargin(v string) {
	o.TotalPositionInitialMargin = &v
}

// GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetTotalOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalOpenOrderInitialMargin
}

// GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetTotalOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		return nil, false
	}
	return o.TotalOpenOrderInitialMargin, true
}

// HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasTotalOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalOpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetTotalOpenOrderInitialMargin gets a reference to the given string and assigns it to the TotalOpenOrderInitialMargin field.
func (o *AccountInformationResponseResult) SetTotalOpenOrderInitialMargin(v string) {
	o.TotalOpenOrderInitialMargin = &v
}

// GetTotalCrossWalletBalance returns the TotalCrossWalletBalance field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetTotalCrossWalletBalance() string {
	if o == nil || common.IsNil(o.TotalCrossWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalCrossWalletBalance
}

// GetTotalCrossWalletBalanceOk returns a tuple with the TotalCrossWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetTotalCrossWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalCrossWalletBalance) {
		return nil, false
	}
	return o.TotalCrossWalletBalance, true
}

// HasTotalCrossWalletBalance returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasTotalCrossWalletBalance() bool {
	if o != nil && !common.IsNil(o.TotalCrossWalletBalance) {
		return true
	}

	return false
}

// SetTotalCrossWalletBalance gets a reference to the given string and assigns it to the TotalCrossWalletBalance field.
func (o *AccountInformationResponseResult) SetTotalCrossWalletBalance(v string) {
	o.TotalCrossWalletBalance = &v
}

// GetTotalCrossUnPnl returns the TotalCrossUnPnl field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetTotalCrossUnPnl() string {
	if o == nil || common.IsNil(o.TotalCrossUnPnl) {
		var ret string
		return ret
	}
	return *o.TotalCrossUnPnl
}

// GetTotalCrossUnPnlOk returns a tuple with the TotalCrossUnPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetTotalCrossUnPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalCrossUnPnl) {
		return nil, false
	}
	return o.TotalCrossUnPnl, true
}

// HasTotalCrossUnPnl returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasTotalCrossUnPnl() bool {
	if o != nil && !common.IsNil(o.TotalCrossUnPnl) {
		return true
	}

	return false
}

// SetTotalCrossUnPnl gets a reference to the given string and assigns it to the TotalCrossUnPnl field.
func (o *AccountInformationResponseResult) SetTotalCrossUnPnl(v string) {
	o.TotalCrossUnPnl = &v
}

// GetAvailableBalance returns the AvailableBalance field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetAvailableBalance() string {
	if o == nil || common.IsNil(o.AvailableBalance) {
		var ret string
		return ret
	}
	return *o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetAvailableBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvailableBalance) {
		return nil, false
	}
	return o.AvailableBalance, true
}

// HasAvailableBalance returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasAvailableBalance() bool {
	if o != nil && !common.IsNil(o.AvailableBalance) {
		return true
	}

	return false
}

// SetAvailableBalance gets a reference to the given string and assigns it to the AvailableBalance field.
func (o *AccountInformationResponseResult) SetAvailableBalance(v string) {
	o.AvailableBalance = &v
}

// GetMaxWithdrawAmount returns the MaxWithdrawAmount field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetMaxWithdrawAmount() string {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.MaxWithdrawAmount
}

// GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetMaxWithdrawAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		return nil, false
	}
	return o.MaxWithdrawAmount, true
}

// HasMaxWithdrawAmount returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasMaxWithdrawAmount() bool {
	if o != nil && !common.IsNil(o.MaxWithdrawAmount) {
		return true
	}

	return false
}

// SetMaxWithdrawAmount gets a reference to the given string and assigns it to the MaxWithdrawAmount field.
func (o *AccountInformationResponseResult) SetMaxWithdrawAmount(v string) {
	o.MaxWithdrawAmount = &v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetAssets() []AccountInformationResponseResultAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []AccountInformationResponseResultAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetAssetsOk() ([]AccountInformationResponseResultAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []AccountInformationResponseResultAssetsInner and assigns it to the Assets field.
func (o *AccountInformationResponseResult) SetAssets(v []AccountInformationResponseResultAssetsInner) {
	o.Assets = v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetPositions() []AccountInformationResponseResultPositionsInner {
	if o == nil || common.IsNil(o.Positions) {
		var ret []AccountInformationResponseResultPositionsInner
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetPositionsOk() ([]AccountInformationResponseResultPositionsInner, bool) {
	if o == nil || common.IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasPositions() bool {
	if o != nil && !common.IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given []AccountInformationResponseResultPositionsInner and assigns it to the Positions field.
func (o *AccountInformationResponseResult) SetPositions(v []AccountInformationResponseResultPositionsInner) {
	o.Positions = v
}

func (o AccountInformationResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInformationResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FeeTier) {
		toSerialize["feeTier"] = o.FeeTier
	}
	if !common.IsNil(o.CanTrade) {
		toSerialize["canTrade"] = o.CanTrade
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountInformationResponseResult) UnmarshalJSON(data []byte) (err error) {
	varAccountInformationResponseResult := _AccountInformationResponseResult{}

	err = json.Unmarshal(data, &varAccountInformationResponseResult)

	if err != nil {
		return err
	}

	*o = AccountInformationResponseResult(varAccountInformationResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "feeTier")
		delete(additionalProperties, "canTrade")
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountInformationResponseResult struct {
	value *AccountInformationResponseResult
	isSet bool
}

func (v NullableAccountInformationResponseResult) Get() *AccountInformationResponseResult {
	return v.value
}

func (v *NullableAccountInformationResponseResult) Set(val *AccountInformationResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInformationResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInformationResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInformationResponseResult(val *AccountInformationResponseResult) *NullableAccountInformationResponseResult {
	return &NullableAccountInformationResponseResult{value: val, isSet: true}
}

func (v NullableAccountInformationResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInformationResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
