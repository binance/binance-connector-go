/*
Binance Convert REST API

OpenAPI Specification for the Binance Convert REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryOrderQuantityPrecisionPerAssetResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryOrderQuantityPrecisionPerAssetResponseInner{}

// QueryOrderQuantityPrecisionPerAssetResponseInner struct for QueryOrderQuantityPrecisionPerAssetResponseInner
type QueryOrderQuantityPrecisionPerAssetResponseInner struct {
	Asset                *string `json:"asset,omitempty"`
	Fraction             *int64  `json:"fraction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryOrderQuantityPrecisionPerAssetResponseInner QueryOrderQuantityPrecisionPerAssetResponseInner

// NewQueryOrderQuantityPrecisionPerAssetResponseInner instantiates a new QueryOrderQuantityPrecisionPerAssetResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryOrderQuantityPrecisionPerAssetResponseInner() *QueryOrderQuantityPrecisionPerAssetResponseInner {
	this := QueryOrderQuantityPrecisionPerAssetResponseInner{}
	return &this
}

// NewQueryOrderQuantityPrecisionPerAssetResponseInnerWithDefaults instantiates a new QueryOrderQuantityPrecisionPerAssetResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryOrderQuantityPrecisionPerAssetResponseInnerWithDefaults() *QueryOrderQuantityPrecisionPerAssetResponseInner {
	this := QueryOrderQuantityPrecisionPerAssetResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryOrderQuantityPrecisionPerAssetResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderQuantityPrecisionPerAssetResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryOrderQuantityPrecisionPerAssetResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryOrderQuantityPrecisionPerAssetResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetFraction returns the Fraction field value if set, zero value otherwise.
func (o *QueryOrderQuantityPrecisionPerAssetResponseInner) GetFraction() int64 {
	if o == nil || common.IsNil(o.Fraction) {
		var ret int64
		return ret
	}
	return *o.Fraction
}

// GetFractionOk returns a tuple with the Fraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderQuantityPrecisionPerAssetResponseInner) GetFractionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Fraction) {
		return nil, false
	}
	return o.Fraction, true
}

// HasFraction returns a boolean if a field has been set.
func (o *QueryOrderQuantityPrecisionPerAssetResponseInner) HasFraction() bool {
	if o != nil && !common.IsNil(o.Fraction) {
		return true
	}

	return false
}

// SetFraction gets a reference to the given int64 and assigns it to the Fraction field.
func (o *QueryOrderQuantityPrecisionPerAssetResponseInner) SetFraction(v int64) {
	o.Fraction = &v
}

func (o QueryOrderQuantityPrecisionPerAssetResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryOrderQuantityPrecisionPerAssetResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Fraction) {
		toSerialize["fraction"] = o.Fraction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryOrderQuantityPrecisionPerAssetResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryOrderQuantityPrecisionPerAssetResponseInner := _QueryOrderQuantityPrecisionPerAssetResponseInner{}

	err = json.Unmarshal(data, &varQueryOrderQuantityPrecisionPerAssetResponseInner)

	if err != nil {
		return err
	}

	*o = QueryOrderQuantityPrecisionPerAssetResponseInner(varQueryOrderQuantityPrecisionPerAssetResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "fraction")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryOrderQuantityPrecisionPerAssetResponseInner struct {
	value *QueryOrderQuantityPrecisionPerAssetResponseInner
	isSet bool
}

func (v NullableQueryOrderQuantityPrecisionPerAssetResponseInner) Get() *QueryOrderQuantityPrecisionPerAssetResponseInner {
	return v.value
}

func (v *NullableQueryOrderQuantityPrecisionPerAssetResponseInner) Set(val *QueryOrderQuantityPrecisionPerAssetResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryOrderQuantityPrecisionPerAssetResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryOrderQuantityPrecisionPerAssetResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryOrderQuantityPrecisionPerAssetResponseInner(val *QueryOrderQuantityPrecisionPerAssetResponseInner) *NullableQueryOrderQuantityPrecisionPerAssetResponseInner {
	return &NullableQueryOrderQuantityPrecisionPerAssetResponseInner{value: val, isSet: true}
}

func (v NullableQueryOrderQuantityPrecisionPerAssetResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryOrderQuantityPrecisionPerAssetResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
