/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFutureHourlyInterestRateResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFutureHourlyInterestRateResponseInner{}

// GetFutureHourlyInterestRateResponseInner struct for GetFutureHourlyInterestRateResponseInner
type GetFutureHourlyInterestRateResponseInner struct {
	Asset                  *string `json:"asset,omitempty"`
	NextHourlyInterestRate *string `json:"nextHourlyInterestRate,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _GetFutureHourlyInterestRateResponseInner GetFutureHourlyInterestRateResponseInner

// NewGetFutureHourlyInterestRateResponseInner instantiates a new GetFutureHourlyInterestRateResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFutureHourlyInterestRateResponseInner() *GetFutureHourlyInterestRateResponseInner {
	this := GetFutureHourlyInterestRateResponseInner{}
	return &this
}

// NewGetFutureHourlyInterestRateResponseInnerWithDefaults instantiates a new GetFutureHourlyInterestRateResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFutureHourlyInterestRateResponseInnerWithDefaults() *GetFutureHourlyInterestRateResponseInner {
	this := GetFutureHourlyInterestRateResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetFutureHourlyInterestRateResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFutureHourlyInterestRateResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetFutureHourlyInterestRateResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetFutureHourlyInterestRateResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetNextHourlyInterestRate returns the NextHourlyInterestRate field value if set, zero value otherwise.
func (o *GetFutureHourlyInterestRateResponseInner) GetNextHourlyInterestRate() string {
	if o == nil || common.IsNil(o.NextHourlyInterestRate) {
		var ret string
		return ret
	}
	return *o.NextHourlyInterestRate
}

// GetNextHourlyInterestRateOk returns a tuple with the NextHourlyInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFutureHourlyInterestRateResponseInner) GetNextHourlyInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.NextHourlyInterestRate) {
		return nil, false
	}
	return o.NextHourlyInterestRate, true
}

// HasNextHourlyInterestRate returns a boolean if a field has been set.
func (o *GetFutureHourlyInterestRateResponseInner) HasNextHourlyInterestRate() bool {
	if o != nil && !common.IsNil(o.NextHourlyInterestRate) {
		return true
	}

	return false
}

// SetNextHourlyInterestRate gets a reference to the given string and assigns it to the NextHourlyInterestRate field.
func (o *GetFutureHourlyInterestRateResponseInner) SetNextHourlyInterestRate(v string) {
	o.NextHourlyInterestRate = &v
}

func (o GetFutureHourlyInterestRateResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFutureHourlyInterestRateResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.NextHourlyInterestRate) {
		toSerialize["nextHourlyInterestRate"] = o.NextHourlyInterestRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFutureHourlyInterestRateResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetFutureHourlyInterestRateResponseInner := _GetFutureHourlyInterestRateResponseInner{}

	err = json.Unmarshal(data, &varGetFutureHourlyInterestRateResponseInner)

	if err != nil {
		return err
	}

	*o = GetFutureHourlyInterestRateResponseInner(varGetFutureHourlyInterestRateResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "nextHourlyInterestRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFutureHourlyInterestRateResponseInner struct {
	value *GetFutureHourlyInterestRateResponseInner
	isSet bool
}

func (v NullableGetFutureHourlyInterestRateResponseInner) Get() *GetFutureHourlyInterestRateResponseInner {
	return v.value
}

func (v *NullableGetFutureHourlyInterestRateResponseInner) Set(val *GetFutureHourlyInterestRateResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFutureHourlyInterestRateResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFutureHourlyInterestRateResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFutureHourlyInterestRateResponseInner(val *GetFutureHourlyInterestRateResponseInner) *NullableGetFutureHourlyInterestRateResponseInner {
	return &NullableGetFutureHourlyInterestRateResponseInner{value: val, isSet: true}
}

func (v NullableGetFutureHourlyInterestRateResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFutureHourlyInterestRateResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
