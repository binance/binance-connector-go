/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner{}

// QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner struct for QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner
type QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner struct {
	Asset                *string  `json:"asset,omitempty"`
	NegativeBalance      *float32 `json:"negativeBalance,omitempty"`
	NegativeMaxThreshold *int64   `json:"negativeMaxThreshold,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner

// NewQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner instantiates a new QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner() *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner {
	this := QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner{}
	return &this
}

// NewQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInnerWithDefaults instantiates a new QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInnerWithDefaults() *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner {
	this := QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetNegativeBalance returns the NegativeBalance field value if set, zero value otherwise.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) GetNegativeBalance() float32 {
	if o == nil || common.IsNil(o.NegativeBalance) {
		var ret float32
		return ret
	}
	return *o.NegativeBalance
}

// GetNegativeBalanceOk returns a tuple with the NegativeBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) GetNegativeBalanceOk() (*float32, bool) {
	if o == nil || common.IsNil(o.NegativeBalance) {
		return nil, false
	}
	return o.NegativeBalance, true
}

// HasNegativeBalance returns a boolean if a field has been set.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) HasNegativeBalance() bool {
	if o != nil && !common.IsNil(o.NegativeBalance) {
		return true
	}

	return false
}

// SetNegativeBalance gets a reference to the given float32 and assigns it to the NegativeBalance field.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) SetNegativeBalance(v float32) {
	o.NegativeBalance = &v
}

// GetNegativeMaxThreshold returns the NegativeMaxThreshold field value if set, zero value otherwise.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) GetNegativeMaxThreshold() int64 {
	if o == nil || common.IsNil(o.NegativeMaxThreshold) {
		var ret int64
		return ret
	}
	return *o.NegativeMaxThreshold
}

// GetNegativeMaxThresholdOk returns a tuple with the NegativeMaxThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) GetNegativeMaxThresholdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.NegativeMaxThreshold) {
		return nil, false
	}
	return o.NegativeMaxThreshold, true
}

// HasNegativeMaxThreshold returns a boolean if a field has been set.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) HasNegativeMaxThreshold() bool {
	if o != nil && !common.IsNil(o.NegativeMaxThreshold) {
		return true
	}

	return false
}

// SetNegativeMaxThreshold gets a reference to the given int64 and assigns it to the NegativeMaxThreshold field.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) SetNegativeMaxThreshold(v int64) {
	o.NegativeMaxThreshold = &v
}

func (o QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.NegativeBalance) {
		toSerialize["negativeBalance"] = o.NegativeBalance
	}
	if !common.IsNil(o.NegativeMaxThreshold) {
		toSerialize["negativeMaxThreshold"] = o.NegativeMaxThreshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner := _QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner{}

	err = json.Unmarshal(data, &varQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner)

	if err != nil {
		return err
	}

	*o = QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner(varQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "negativeBalance")
		delete(additionalProperties, "negativeMaxThreshold")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner struct {
	value *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner
	isSet bool
}

func (v NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) Get() *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner {
	return v.value
}

func (v *NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) Set(val *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner(val *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) *NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner {
	return &NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner{value: val, isSet: true}
}

func (v NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
