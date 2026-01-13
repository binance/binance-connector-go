/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetDetailOnSubAccountsFuturesAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDetailOnSubAccountsFuturesAccountResponse{}

// GetDetailOnSubAccountsFuturesAccountResponse struct for GetDetailOnSubAccountsFuturesAccountResponse
type GetDetailOnSubAccountsFuturesAccountResponse struct {
	Email                       *string                                                                      `json:"email,omitempty"`
	Asset                       *string                                                                      `json:"asset,omitempty"`
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

type _GetDetailOnSubAccountsFuturesAccountResponse GetDetailOnSubAccountsFuturesAccountResponse

// NewGetDetailOnSubAccountsFuturesAccountResponse instantiates a new GetDetailOnSubAccountsFuturesAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDetailOnSubAccountsFuturesAccountResponse() *GetDetailOnSubAccountsFuturesAccountResponse {
	this := GetDetailOnSubAccountsFuturesAccountResponse{}
	return &this
}

// NewGetDetailOnSubAccountsFuturesAccountResponseWithDefaults instantiates a new GetDetailOnSubAccountsFuturesAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDetailOnSubAccountsFuturesAccountResponseWithDefaults() *GetDetailOnSubAccountsFuturesAccountResponse {
	this := GetDetailOnSubAccountsFuturesAccountResponse{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetEmail(v string) {
	o.Email = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetAsset(v string) {
	o.Asset = &v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetAssets() []GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetAssetsOk() ([]GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner and assigns it to the Assets field.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetAssets(v []GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) {
	o.Assets = v
}

// GetCanDeposit returns the CanDeposit field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetCanDeposit() bool {
	if o == nil || common.IsNil(o.CanDeposit) {
		var ret bool
		return ret
	}
	return *o.CanDeposit
}

// GetCanDepositOk returns a tuple with the CanDeposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetCanDepositOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanDeposit) {
		return nil, false
	}
	return o.CanDeposit, true
}

// HasCanDeposit returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasCanDeposit() bool {
	if o != nil && !common.IsNil(o.CanDeposit) {
		return true
	}

	return false
}

// SetCanDeposit gets a reference to the given bool and assigns it to the CanDeposit field.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetCanDeposit(v bool) {
	o.CanDeposit = &v
}

// GetCanTrade returns the CanTrade field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetCanTrade() bool {
	if o == nil || common.IsNil(o.CanTrade) {
		var ret bool
		return ret
	}
	return *o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetCanTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanTrade) {
		return nil, false
	}
	return o.CanTrade, true
}

// HasCanTrade returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasCanTrade() bool {
	if o != nil && !common.IsNil(o.CanTrade) {
		return true
	}

	return false
}

// SetCanTrade gets a reference to the given bool and assigns it to the CanTrade field.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetCanTrade(v bool) {
	o.CanTrade = &v
}

// GetCanWithdraw returns the CanWithdraw field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetCanWithdraw() bool {
	if o == nil || common.IsNil(o.CanWithdraw) {
		var ret bool
		return ret
	}
	return *o.CanWithdraw
}

// GetCanWithdrawOk returns a tuple with the CanWithdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetCanWithdrawOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanWithdraw) {
		return nil, false
	}
	return o.CanWithdraw, true
}

// HasCanWithdraw returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasCanWithdraw() bool {
	if o != nil && !common.IsNil(o.CanWithdraw) {
		return true
	}

	return false
}

// SetCanWithdraw gets a reference to the given bool and assigns it to the CanWithdraw field.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetCanWithdraw(v bool) {
	o.CanWithdraw = &v
}

// GetFeeTier returns the FeeTier field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetFeeTier() int64 {
	if o == nil || common.IsNil(o.FeeTier) {
		var ret int64
		return ret
	}
	return *o.FeeTier
}

// GetFeeTierOk returns a tuple with the FeeTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetFeeTierOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FeeTier) {
		return nil, false
	}
	return o.FeeTier, true
}

// HasFeeTier returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasFeeTier() bool {
	if o != nil && !common.IsNil(o.FeeTier) {
		return true
	}

	return false
}

// SetFeeTier gets a reference to the given int64 and assigns it to the FeeTier field.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetFeeTier(v int64) {
	o.FeeTier = &v
}

// GetMaxWithdrawAmount returns the MaxWithdrawAmount field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetMaxWithdrawAmount() string {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.MaxWithdrawAmount
}

// GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetMaxWithdrawAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		return nil, false
	}
	return o.MaxWithdrawAmount, true
}

// HasMaxWithdrawAmount returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasMaxWithdrawAmount() bool {
	if o != nil && !common.IsNil(o.MaxWithdrawAmount) {
		return true
	}

	return false
}

// SetMaxWithdrawAmount gets a reference to the given string and assigns it to the MaxWithdrawAmount field.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetMaxWithdrawAmount(v string) {
	o.MaxWithdrawAmount = &v
}

// GetTotalInitialMargin returns the TotalInitialMargin field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalInitialMargin() string {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalInitialMargin
}

// GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		return nil, false
	}
	return o.TotalInitialMargin, true
}

// HasTotalInitialMargin returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasTotalInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalInitialMargin) {
		return true
	}

	return false
}

// SetTotalInitialMargin gets a reference to the given string and assigns it to the TotalInitialMargin field.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetTotalInitialMargin(v string) {
	o.TotalInitialMargin = &v
}

// GetTotalMaintenanceMargin returns the TotalMaintenanceMargin field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalMaintenanceMargin() string {
	if o == nil || common.IsNil(o.TotalMaintenanceMargin) {
		var ret string
		return ret
	}
	return *o.TotalMaintenanceMargin
}

// GetTotalMaintenanceMarginOk returns a tuple with the TotalMaintenanceMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalMaintenanceMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMaintenanceMargin) {
		return nil, false
	}
	return o.TotalMaintenanceMargin, true
}

// HasTotalMaintenanceMargin returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasTotalMaintenanceMargin() bool {
	if o != nil && !common.IsNil(o.TotalMaintenanceMargin) {
		return true
	}

	return false
}

// SetTotalMaintenanceMargin gets a reference to the given string and assigns it to the TotalMaintenanceMargin field.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetTotalMaintenanceMargin(v string) {
	o.TotalMaintenanceMargin = &v
}

// GetTotalMarginBalance returns the TotalMarginBalance field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalMarginBalance() string {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		var ret string
		return ret
	}
	return *o.TotalMarginBalance
}

// GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		return nil, false
	}
	return o.TotalMarginBalance, true
}

// HasTotalMarginBalance returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasTotalMarginBalance() bool {
	if o != nil && !common.IsNil(o.TotalMarginBalance) {
		return true
	}

	return false
}

// SetTotalMarginBalance gets a reference to the given string and assigns it to the TotalMarginBalance field.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetTotalMarginBalance(v string) {
	o.TotalMarginBalance = &v
}

// GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalOpenOrderInitialMargin
}

// GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		return nil, false
	}
	return o.TotalOpenOrderInitialMargin, true
}

// HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasTotalOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalOpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetTotalOpenOrderInitialMargin gets a reference to the given string and assigns it to the TotalOpenOrderInitialMargin field.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetTotalOpenOrderInitialMargin(v string) {
	o.TotalOpenOrderInitialMargin = &v
}

// GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalPositionInitialMargin() string {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalPositionInitialMargin
}

// GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		return nil, false
	}
	return o.TotalPositionInitialMargin, true
}

// HasTotalPositionInitialMargin returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasTotalPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalPositionInitialMargin) {
		return true
	}

	return false
}

// SetTotalPositionInitialMargin gets a reference to the given string and assigns it to the TotalPositionInitialMargin field.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetTotalPositionInitialMargin(v string) {
	o.TotalPositionInitialMargin = &v
}

// GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalUnrealizedProfit() string {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.TotalUnrealizedProfit
}

// GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		return nil, false
	}
	return o.TotalUnrealizedProfit, true
}

// HasTotalUnrealizedProfit returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasTotalUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.TotalUnrealizedProfit) {
		return true
	}

	return false
}

// SetTotalUnrealizedProfit gets a reference to the given string and assigns it to the TotalUnrealizedProfit field.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetTotalUnrealizedProfit(v string) {
	o.TotalUnrealizedProfit = &v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalWalletBalance() string {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		return nil, false
	}
	return o.TotalWalletBalance, true
}

// HasTotalWalletBalance returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasTotalWalletBalance() bool {
	if o != nil && !common.IsNil(o.TotalWalletBalance) {
		return true
	}

	return false
}

// SetTotalWalletBalance gets a reference to the given string and assigns it to the TotalWalletBalance field.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o GetDetailOnSubAccountsFuturesAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDetailOnSubAccountsFuturesAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
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

func (o *GetDetailOnSubAccountsFuturesAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varGetDetailOnSubAccountsFuturesAccountResponse := _GetDetailOnSubAccountsFuturesAccountResponse{}

	err = json.Unmarshal(data, &varGetDetailOnSubAccountsFuturesAccountResponse)

	if err != nil {
		return err
	}

	*o = GetDetailOnSubAccountsFuturesAccountResponse(varGetDetailOnSubAccountsFuturesAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "asset")
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

type NullableGetDetailOnSubAccountsFuturesAccountResponse struct {
	value *GetDetailOnSubAccountsFuturesAccountResponse
	isSet bool
}

func (v NullableGetDetailOnSubAccountsFuturesAccountResponse) Get() *GetDetailOnSubAccountsFuturesAccountResponse {
	return v.value
}

func (v *NullableGetDetailOnSubAccountsFuturesAccountResponse) Set(val *GetDetailOnSubAccountsFuturesAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDetailOnSubAccountsFuturesAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDetailOnSubAccountsFuturesAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDetailOnSubAccountsFuturesAccountResponse(val *GetDetailOnSubAccountsFuturesAccountResponse) *NullableGetDetailOnSubAccountsFuturesAccountResponse {
	return &NullableGetDetailOnSubAccountsFuturesAccountResponse{value: val, isSet: true}
}

func (v NullableGetDetailOnSubAccountsFuturesAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDetailOnSubAccountsFuturesAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
