/*
Binance Pay REST API

OpenAPI Specification for the Binance Pay REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner{}

// GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner struct for GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner
type GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner struct {
	Var1                 *string `json:"1,omitempty"`
	Var2                 *string `json:"2,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner

// NewGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner instantiates a new GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner() *GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner {
	this := GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner{}
	return &this
}

// NewGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInnerWithDefaults instantiates a new GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInnerWithDefaults() *GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner {
	this := GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner{}
	return &this
}

// GetVar1 returns the Var1 field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) GetVar1() string {
	if o == nil || common.IsNil(o.Var1) {
		var ret string
		return ret
	}
	return *o.Var1
}

// GetVar1Ok returns a tuple with the Var1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) GetVar1Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Var1) {
		return nil, false
	}
	return o.Var1, true
}

// HasVar1 returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) HasVar1() bool {
	if o != nil && !common.IsNil(o.Var1) {
		return true
	}

	return false
}

// SetVar1 gets a reference to the given string and assigns it to the Var1 field.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) SetVar1(v string) {
	o.Var1 = &v
}

// GetVar2 returns the Var2 field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) GetVar2() string {
	if o == nil || common.IsNil(o.Var2) {
		var ret string
		return ret
	}
	return *o.Var2
}

// GetVar2Ok returns a tuple with the Var2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) GetVar2Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Var2) {
		return nil, false
	}
	return o.Var2, true
}

// HasVar2 returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) HasVar2() bool {
	if o != nil && !common.IsNil(o.Var2) {
		return true
	}

	return false
}

// SetVar2 gets a reference to the given string and assigns it to the Var2 field.
func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) SetVar2(v string) {
	o.Var2 = &v
}

func (o GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Var1) {
		toSerialize["1"] = o.Var1
	}
	if !common.IsNil(o.Var2) {
		toSerialize["2"] = o.Var2
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) UnmarshalJSON(data []byte) (err error) {
	varGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner := _GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner{}

	err = json.Unmarshal(data, &varGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner)

	if err != nil {
		return err
	}

	*o = GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner(varGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "1")
		delete(additionalProperties, "2")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner struct {
	value *GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner
	isSet bool
}

func (v NullableGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) Get() *GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner {
	return v.value
}

func (v *NullableGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) Set(val *GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner(val *GetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) *NullableGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner {
	return &NullableGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner{value: val, isSet: true}
}

func (v NullableGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPayTradeHistoryResponseDataInnerFundsDetailInnerWalletAssetCostInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
