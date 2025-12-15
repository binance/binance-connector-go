/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp{}

// GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp struct for GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp
type GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp struct {
	Email                       *string                                                                      `json:"email,omitempty"`
	Assets                      []GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner `json:"assets,omitempty"`
	CanDeposit                  *bool                                                                        `json:"canDeposit,omitempty"`
	CanTrade                    *bool                                                                        `json:"canTrade,omitempty"`
	CanWithdraw                 *bool                                                                        `json:"canWithdraw,omitempty"`
	FeeTier                     *int64                                                                       `json:"feeTier,omitempty"`
	MaxWithdrawAmount           *string                                                                      `json:"maxWithdrawAmount,omitempty"`
	TotalInitialMargin          *string                                                                      `json:"totalInitialMargin,omitempty"`
	TotalMaintenanceMargin      *string                                                                      `json:"totalMaintenanceMargin,omitempty"`
	TotalMarginBalance          *string                                                                      `json:"totalMarginBalance,omitempty"`
	TotalOpenOrderInitialMargin *string                                                                      `json:"totalOpenOrderInitialMargin,omitempty"`
	TotalPositionInitialMargin  *string                                                                      `json:"totalPositionInitialMargin,omitempty"`
	TotalUnrealizedProfit       *string                                                                      `json:"totalUnrealizedProfit,omitempty"`
	TotalWalletBalance          *string                                                                      `json:"totalWalletBalance,omitempty"`
	UpdateTime                  *int64                                                                       `json:"updateTime,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp

// NewGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp instantiates a new GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp() *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp {
	this := GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp{}
	return &this
}

// NewGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespWithDefaults instantiates a new GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespWithDefaults() *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp {
	this := GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetEmail(v string) {
	o.Email = &v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetAssets() []GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetAssetsOk() ([]GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner and assigns it to the Assets field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetAssets(v []GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) {
	o.Assets = v
}

// GetCanDeposit returns the CanDeposit field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetCanDeposit() bool {
	if o == nil || common.IsNil(o.CanDeposit) {
		var ret bool
		return ret
	}
	return *o.CanDeposit
}

// GetCanDepositOk returns a tuple with the CanDeposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetCanDepositOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanDeposit) {
		return nil, false
	}
	return o.CanDeposit, true
}

// HasCanDeposit returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasCanDeposit() bool {
	if o != nil && !common.IsNil(o.CanDeposit) {
		return true
	}

	return false
}

// SetCanDeposit gets a reference to the given bool and assigns it to the CanDeposit field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetCanDeposit(v bool) {
	o.CanDeposit = &v
}

// GetCanTrade returns the CanTrade field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetCanTrade() bool {
	if o == nil || common.IsNil(o.CanTrade) {
		var ret bool
		return ret
	}
	return *o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetCanTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanTrade) {
		return nil, false
	}
	return o.CanTrade, true
}

// HasCanTrade returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasCanTrade() bool {
	if o != nil && !common.IsNil(o.CanTrade) {
		return true
	}

	return false
}

// SetCanTrade gets a reference to the given bool and assigns it to the CanTrade field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetCanTrade(v bool) {
	o.CanTrade = &v
}

// GetCanWithdraw returns the CanWithdraw field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetCanWithdraw() bool {
	if o == nil || common.IsNil(o.CanWithdraw) {
		var ret bool
		return ret
	}
	return *o.CanWithdraw
}

// GetCanWithdrawOk returns a tuple with the CanWithdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetCanWithdrawOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanWithdraw) {
		return nil, false
	}
	return o.CanWithdraw, true
}

// HasCanWithdraw returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasCanWithdraw() bool {
	if o != nil && !common.IsNil(o.CanWithdraw) {
		return true
	}

	return false
}

// SetCanWithdraw gets a reference to the given bool and assigns it to the CanWithdraw field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetCanWithdraw(v bool) {
	o.CanWithdraw = &v
}

// GetFeeTier returns the FeeTier field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetFeeTier() int64 {
	if o == nil || common.IsNil(o.FeeTier) {
		var ret int64
		return ret
	}
	return *o.FeeTier
}

// GetFeeTierOk returns a tuple with the FeeTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetFeeTierOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FeeTier) {
		return nil, false
	}
	return o.FeeTier, true
}

// HasFeeTier returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasFeeTier() bool {
	if o != nil && !common.IsNil(o.FeeTier) {
		return true
	}

	return false
}

// SetFeeTier gets a reference to the given int64 and assigns it to the FeeTier field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetFeeTier(v int64) {
	o.FeeTier = &v
}

// GetMaxWithdrawAmount returns the MaxWithdrawAmount field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetMaxWithdrawAmount() string {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.MaxWithdrawAmount
}

// GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetMaxWithdrawAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		return nil, false
	}
	return o.MaxWithdrawAmount, true
}

// HasMaxWithdrawAmount returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasMaxWithdrawAmount() bool {
	if o != nil && !common.IsNil(o.MaxWithdrawAmount) {
		return true
	}

	return false
}

// SetMaxWithdrawAmount gets a reference to the given string and assigns it to the MaxWithdrawAmount field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetMaxWithdrawAmount(v string) {
	o.MaxWithdrawAmount = &v
}

// GetTotalInitialMargin returns the TotalInitialMargin field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalInitialMargin() string {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalInitialMargin
}

// GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		return nil, false
	}
	return o.TotalInitialMargin, true
}

// HasTotalInitialMargin returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasTotalInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalInitialMargin) {
		return true
	}

	return false
}

// SetTotalInitialMargin gets a reference to the given string and assigns it to the TotalInitialMargin field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetTotalInitialMargin(v string) {
	o.TotalInitialMargin = &v
}

// GetTotalMaintenanceMargin returns the TotalMaintenanceMargin field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalMaintenanceMargin() string {
	if o == nil || common.IsNil(o.TotalMaintenanceMargin) {
		var ret string
		return ret
	}
	return *o.TotalMaintenanceMargin
}

// GetTotalMaintenanceMarginOk returns a tuple with the TotalMaintenanceMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalMaintenanceMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMaintenanceMargin) {
		return nil, false
	}
	return o.TotalMaintenanceMargin, true
}

// HasTotalMaintenanceMargin returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasTotalMaintenanceMargin() bool {
	if o != nil && !common.IsNil(o.TotalMaintenanceMargin) {
		return true
	}

	return false
}

// SetTotalMaintenanceMargin gets a reference to the given string and assigns it to the TotalMaintenanceMargin field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetTotalMaintenanceMargin(v string) {
	o.TotalMaintenanceMargin = &v
}

// GetTotalMarginBalance returns the TotalMarginBalance field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalMarginBalance() string {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		var ret string
		return ret
	}
	return *o.TotalMarginBalance
}

// GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		return nil, false
	}
	return o.TotalMarginBalance, true
}

// HasTotalMarginBalance returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasTotalMarginBalance() bool {
	if o != nil && !common.IsNil(o.TotalMarginBalance) {
		return true
	}

	return false
}

// SetTotalMarginBalance gets a reference to the given string and assigns it to the TotalMarginBalance field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetTotalMarginBalance(v string) {
	o.TotalMarginBalance = &v
}

// GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalOpenOrderInitialMargin
}

// GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		return nil, false
	}
	return o.TotalOpenOrderInitialMargin, true
}

// HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasTotalOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalOpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetTotalOpenOrderInitialMargin gets a reference to the given string and assigns it to the TotalOpenOrderInitialMargin field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetTotalOpenOrderInitialMargin(v string) {
	o.TotalOpenOrderInitialMargin = &v
}

// GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalPositionInitialMargin() string {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalPositionInitialMargin
}

// GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		return nil, false
	}
	return o.TotalPositionInitialMargin, true
}

// HasTotalPositionInitialMargin returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasTotalPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalPositionInitialMargin) {
		return true
	}

	return false
}

// SetTotalPositionInitialMargin gets a reference to the given string and assigns it to the TotalPositionInitialMargin field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetTotalPositionInitialMargin(v string) {
	o.TotalPositionInitialMargin = &v
}

// GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalUnrealizedProfit() string {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.TotalUnrealizedProfit
}

// GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		return nil, false
	}
	return o.TotalUnrealizedProfit, true
}

// HasTotalUnrealizedProfit returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasTotalUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.TotalUnrealizedProfit) {
		return true
	}

	return false
}

// SetTotalUnrealizedProfit gets a reference to the given string and assigns it to the TotalUnrealizedProfit field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetTotalUnrealizedProfit(v string) {
	o.TotalUnrealizedProfit = &v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalWalletBalance() string {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		return nil, false
	}
	return o.TotalWalletBalance, true
}

// HasTotalWalletBalance returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasTotalWalletBalance() bool {
	if o != nil && !common.IsNil(o.TotalWalletBalance) {
		return true
	}

	return false
}

// SetTotalWalletBalance gets a reference to the given string and assigns it to the TotalWalletBalance field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !common.IsNil(o.CanDeposit) {
		toSerialize["canDeposit"] = o.CanDeposit
	}
	if !common.IsNil(o.CanTrade) {
		toSerialize["canTrade"] = o.CanTrade
	}
	if !common.IsNil(o.CanWithdraw) {
		toSerialize["canWithdraw"] = o.CanWithdraw
	}
	if !common.IsNil(o.FeeTier) {
		toSerialize["feeTier"] = o.FeeTier
	}
	if !common.IsNil(o.MaxWithdrawAmount) {
		toSerialize["maxWithdrawAmount"] = o.MaxWithdrawAmount
	}
	if !common.IsNil(o.TotalInitialMargin) {
		toSerialize["totalInitialMargin"] = o.TotalInitialMargin
	}
	if !common.IsNil(o.TotalMaintenanceMargin) {
		toSerialize["totalMaintenanceMargin"] = o.TotalMaintenanceMargin
	}
	if !common.IsNil(o.TotalMarginBalance) {
		toSerialize["totalMarginBalance"] = o.TotalMarginBalance
	}
	if !common.IsNil(o.TotalOpenOrderInitialMargin) {
		toSerialize["totalOpenOrderInitialMargin"] = o.TotalOpenOrderInitialMargin
	}
	if !common.IsNil(o.TotalPositionInitialMargin) {
		toSerialize["totalPositionInitialMargin"] = o.TotalPositionInitialMargin
	}
	if !common.IsNil(o.TotalUnrealizedProfit) {
		toSerialize["totalUnrealizedProfit"] = o.TotalUnrealizedProfit
	}
	if !common.IsNil(o.TotalWalletBalance) {
		toSerialize["totalWalletBalance"] = o.TotalWalletBalance
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) UnmarshalJSON(data []byte) (err error) {
	varGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp := _GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp{}

	err = json.Unmarshal(data, &varGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp)

	if err != nil {
		return err
	}

	*o = GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp(varGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "assets")
		delete(additionalProperties, "canDeposit")
		delete(additionalProperties, "canTrade")
		delete(additionalProperties, "canWithdraw")
		delete(additionalProperties, "feeTier")
		delete(additionalProperties, "maxWithdrawAmount")
		delete(additionalProperties, "totalInitialMargin")
		delete(additionalProperties, "totalMaintenanceMargin")
		delete(additionalProperties, "totalMarginBalance")
		delete(additionalProperties, "totalOpenOrderInitialMargin")
		delete(additionalProperties, "totalPositionInitialMargin")
		delete(additionalProperties, "totalUnrealizedProfit")
		delete(additionalProperties, "totalWalletBalance")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp struct {
	value *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp
	isSet bool
}

func (v NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) Get() *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp {
	return v.value
}

func (v *NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) Set(val *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp(val *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) *NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp {
	return &NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp{value: val, isSet: true}
}

func (v NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
