/*
Binance Pay REST API

OpenAPI Specification for the Binance Pay REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetPayTradeHistoryResponseDataInnerFundsDetailInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPayTradeHistoryResponseDataInnerFundsDetailInner{}

// GetPayTradeHistoryResponseDataInnerFundsDetailInner struct for GetPayTradeHistoryResponseDataInnerFundsDetailInner
type GetPayTradeHistoryResponseDataInnerFundsDetailInner struct {
	Currency             *string                                                                   `json:"currency,omitempty"`
	Amount               *string                                                                   `json:"amount,omitempty"`
	WalletAssetCost      []GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner `json:"walletAssetCost,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPayTradeHistoryResponseDataInnerFundsDetailInner GetPayTradeHistoryResponseDataInnerFundsDetailInner

// NewGetPayTradeHistoryResponseDataInnerFundsDetailInner instantiates a new GetPayTradeHistoryResponseDataInnerFundsDetailInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPayTradeHistoryResponseDataInnerFundsDetailInner() *GetPayTradeHistoryResponseDataInnerFundsDetailInner {
	this := GetPayTradeHistoryResponseDataInnerFundsDetailInner{}
	return &this
}

// NewGetPayTradeHistoryResponseDataInnerFundsDetailInnerWithDefaults instantiates a new GetPayTradeHistoryResponseDataInnerFundsDetailInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPayTradeHistoryResponseDataInnerFundsDetailInnerWithDefaults() *GetPayTradeHistoryResponseDataInnerFundsDetailInner {
	this := GetPayTradeHistoryResponseDataInnerFundsDetailInner{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) GetCurrency() string {
	if o == nil || common.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) GetCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) HasCurrency() bool {
	if o != nil && !common.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) SetCurrency(v string) {
	o.Currency = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) SetAmount(v string) {
	o.Amount = &v
}

// GetWalletAssetCost returns the WalletAssetCost field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) GetWalletAssetCost() []GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner {
	if o == nil || common.IsNil(o.WalletAssetCost) {
		var ret []GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner
		return ret
	}
	return o.WalletAssetCost
}

// GetWalletAssetCostOk returns a tuple with the WalletAssetCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) GetWalletAssetCostOk() ([]GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner, bool) {
	if o == nil || common.IsNil(o.WalletAssetCost) {
		return nil, false
	}
	return o.WalletAssetCost, true
}

// HasWalletAssetCost returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) HasWalletAssetCost() bool {
	if o != nil && !common.IsNil(o.WalletAssetCost) {
		return true
	}

	return false
}

// SetWalletAssetCost gets a reference to the given []GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner and assigns it to the WalletAssetCost field.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) SetWalletAssetCost(v []GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) {
	o.WalletAssetCost = v
}

func (o GetPayTradeHistoryResponseDataInnerFundsDetailInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPayTradeHistoryResponseDataInnerFundsDetailInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.WalletAssetCost) {
		toSerialize["walletAssetCost"] = o.WalletAssetCost
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) UnmarshalJSON(data []byte) (err error) {
	varGetPayTradeHistoryResponseDataInnerFundsDetailInner := _GetPayTradeHistoryResponseDataInnerFundsDetailInner{}

	err = json.Unmarshal(data, &varGetPayTradeHistoryResponseDataInnerFundsDetailInner)

	if err != nil {
		return err
	}

	*o = GetPayTradeHistoryResponseDataInnerFundsDetailInner(varGetPayTradeHistoryResponseDataInnerFundsDetailInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "currency")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "walletAssetCost")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPayTradeHistoryResponseDataInnerFundsDetailInner struct {
	value *GetPayTradeHistoryResponseDataInnerFundsDetailInner
	isSet bool
}

func (v NullableGetPayTradeHistoryResponseDataInnerFundsDetailInner) Get() *GetPayTradeHistoryResponseDataInnerFundsDetailInner {
	return v.value
}

func (v *NullableGetPayTradeHistoryResponseDataInnerFundsDetailInner) Set(val *GetPayTradeHistoryResponseDataInnerFundsDetailInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPayTradeHistoryResponseDataInnerFundsDetailInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPayTradeHistoryResponseDataInnerFundsDetailInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPayTradeHistoryResponseDataInnerFundsDetailInner(val *GetPayTradeHistoryResponseDataInnerFundsDetailInner) *NullableGetPayTradeHistoryResponseDataInnerFundsDetailInner {
	return &NullableGetPayTradeHistoryResponseDataInnerFundsDetailInner{value: val, isSet: true}
}

func (v NullableGetPayTradeHistoryResponseDataInnerFundsDetailInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPayTradeHistoryResponseDataInnerFundsDetailInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
