/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner{}

// GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner struct for GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner
type GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner struct {
	Asset                  *string `json:"asset,omitempty"`
	InitialMargin          *string `json:"initialMargin,omitempty"`
	MaintenanceMargin      *string `json:"maintenanceMargin,omitempty"`
	MarginBalance          *string `json:"marginBalance,omitempty"`
	MaxWithdrawAmount      *string `json:"maxWithdrawAmount,omitempty"`
	OpenOrderInitialMargin *string `json:"openOrderInitialMargin,omitempty"`
	PositionInitialMargin  *string `json:"positionInitialMargin,omitempty"`
	UnrealizedProfit       *string `json:"unrealizedProfit,omitempty"`
	WalletBalance          *string `json:"walletBalance,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner

// NewGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner instantiates a new GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner() *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner {
	this := GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner{}
	return &this
}

// NewGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInnerWithDefaults instantiates a new GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInnerWithDefaults() *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner {
	this := GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetInitialMargin() string {
	if o == nil || common.IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) HasInitialMargin() bool {
	if o != nil && !common.IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetMaintenanceMargin returns the MaintenanceMargin field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetMaintenanceMargin() string {
	if o == nil || common.IsNil(o.MaintenanceMargin) {
		var ret string
		return ret
	}
	return *o.MaintenanceMargin
}

// GetMaintenanceMarginOk returns a tuple with the MaintenanceMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetMaintenanceMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaintenanceMargin) {
		return nil, false
	}
	return o.MaintenanceMargin, true
}

// HasMaintenanceMargin returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) HasMaintenanceMargin() bool {
	if o != nil && !common.IsNil(o.MaintenanceMargin) {
		return true
	}

	return false
}

// SetMaintenanceMargin gets a reference to the given string and assigns it to the MaintenanceMargin field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) SetMaintenanceMargin(v string) {
	o.MaintenanceMargin = &v
}

// GetMarginBalance returns the MarginBalance field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetMarginBalance() string {
	if o == nil || common.IsNil(o.MarginBalance) {
		var ret string
		return ret
	}
	return *o.MarginBalance
}

// GetMarginBalanceOk returns a tuple with the MarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginBalance) {
		return nil, false
	}
	return o.MarginBalance, true
}

// HasMarginBalance returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) HasMarginBalance() bool {
	if o != nil && !common.IsNil(o.MarginBalance) {
		return true
	}

	return false
}

// SetMarginBalance gets a reference to the given string and assigns it to the MarginBalance field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) SetMarginBalance(v string) {
	o.MarginBalance = &v
}

// GetMaxWithdrawAmount returns the MaxWithdrawAmount field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetMaxWithdrawAmount() string {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.MaxWithdrawAmount
}

// GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetMaxWithdrawAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		return nil, false
	}
	return o.MaxWithdrawAmount, true
}

// HasMaxWithdrawAmount returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) HasMaxWithdrawAmount() bool {
	if o != nil && !common.IsNil(o.MaxWithdrawAmount) {
		return true
	}

	return false
}

// SetMaxWithdrawAmount gets a reference to the given string and assigns it to the MaxWithdrawAmount field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) SetMaxWithdrawAmount(v string) {
	o.MaxWithdrawAmount = &v
}

// GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.OpenOrderInitialMargin
}

// GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		return nil, false
	}
	return o.OpenOrderInitialMargin, true
}

// HasOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) HasOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.OpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetOpenOrderInitialMargin gets a reference to the given string and assigns it to the OpenOrderInitialMargin field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) SetOpenOrderInitialMargin(v string) {
	o.OpenOrderInitialMargin = &v
}

// GetPositionInitialMargin returns the PositionInitialMargin field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetPositionInitialMargin() string {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.PositionInitialMargin
}

// GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		return nil, false
	}
	return o.PositionInitialMargin, true
}

// HasPositionInitialMargin returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) HasPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.PositionInitialMargin) {
		return true
	}

	return false
}

// SetPositionInitialMargin gets a reference to the given string and assigns it to the PositionInitialMargin field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) SetPositionInitialMargin(v string) {
	o.PositionInitialMargin = &v
}

// GetUnrealizedProfit returns the UnrealizedProfit field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetUnrealizedProfit() string {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfit
}

// GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		return nil, false
	}
	return o.UnrealizedProfit, true
}

// HasUnrealizedProfit returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) HasUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.UnrealizedProfit) {
		return true
	}

	return false
}

// SetUnrealizedProfit gets a reference to the given string and assigns it to the UnrealizedProfit field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) SetUnrealizedProfit(v string) {
	o.UnrealizedProfit = &v
}

// GetWalletBalance returns the WalletBalance field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetWalletBalance() string {
	if o == nil || common.IsNil(o.WalletBalance) {
		var ret string
		return ret
	}
	return *o.WalletBalance
}

// GetWalletBalanceOk returns a tuple with the WalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) GetWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletBalance) {
		return nil, false
	}
	return o.WalletBalance, true
}

// HasWalletBalance returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) HasWalletBalance() bool {
	if o != nil && !common.IsNil(o.WalletBalance) {
		return true
	}

	return false
}

// SetWalletBalance gets a reference to the given string and assigns it to the WalletBalance field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) SetWalletBalance(v string) {
	o.WalletBalance = &v
}

func (o GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.InitialMargin) {
		toSerialize["initialMargin"] = o.InitialMargin
	}
	if !common.IsNil(o.MaintenanceMargin) {
		toSerialize["maintenanceMargin"] = o.MaintenanceMargin
	}
	if !common.IsNil(o.MarginBalance) {
		toSerialize["marginBalance"] = o.MarginBalance
	}
	if !common.IsNil(o.MaxWithdrawAmount) {
		toSerialize["maxWithdrawAmount"] = o.MaxWithdrawAmount
	}
	if !common.IsNil(o.OpenOrderInitialMargin) {
		toSerialize["openOrderInitialMargin"] = o.OpenOrderInitialMargin
	}
	if !common.IsNil(o.PositionInitialMargin) {
		toSerialize["positionInitialMargin"] = o.PositionInitialMargin
	}
	if !common.IsNil(o.UnrealizedProfit) {
		toSerialize["unrealizedProfit"] = o.UnrealizedProfit
	}
	if !common.IsNil(o.WalletBalance) {
		toSerialize["walletBalance"] = o.WalletBalance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) UnmarshalJSON(data []byte) (err error) {
	varGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner := _GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner{}

	err = json.Unmarshal(data, &varGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner)

	if err != nil {
		return err
	}

	*o = GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner(varGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "initialMargin")
		delete(additionalProperties, "maintenanceMargin")
		delete(additionalProperties, "marginBalance")
		delete(additionalProperties, "maxWithdrawAmount")
		delete(additionalProperties, "openOrderInitialMargin")
		delete(additionalProperties, "positionInitialMargin")
		delete(additionalProperties, "unrealizedProfit")
		delete(additionalProperties, "walletBalance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner struct {
	value *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner
	isSet bool
}

func (v NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) Get() *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner {
	return v.value
}

func (v *NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) Set(val *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner(val *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) *NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner {
	return &NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner{value: val, isSet: true}
}

func (v NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
