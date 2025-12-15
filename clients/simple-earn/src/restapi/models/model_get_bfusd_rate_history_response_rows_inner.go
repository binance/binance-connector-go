/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetBfusdRateHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBfusdRateHistoryResponseRowsInner{}

// GetBfusdRateHistoryResponseRowsInner struct for GetBfusdRateHistoryResponseRowsInner
type GetBfusdRateHistoryResponseRowsInner struct {
	AnnualPercentageRate *string `json:"annualPercentageRate,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBfusdRateHistoryResponseRowsInner GetBfusdRateHistoryResponseRowsInner

// NewGetBfusdRateHistoryResponseRowsInner instantiates a new GetBfusdRateHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBfusdRateHistoryResponseRowsInner() *GetBfusdRateHistoryResponseRowsInner {
	this := GetBfusdRateHistoryResponseRowsInner{}
	return &this
}

// NewGetBfusdRateHistoryResponseRowsInnerWithDefaults instantiates a new GetBfusdRateHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBfusdRateHistoryResponseRowsInnerWithDefaults() *GetBfusdRateHistoryResponseRowsInner {
	this := GetBfusdRateHistoryResponseRowsInner{}
	return &this
}

// GetAnnualPercentageRate returns the AnnualPercentageRate field value if set, zero value otherwise.
func (o *GetBfusdRateHistoryResponseRowsInner) GetAnnualPercentageRate() string {
	if o == nil || common.IsNil(o.AnnualPercentageRate) {
		var ret string
		return ret
	}
	return *o.AnnualPercentageRate
}

// GetAnnualPercentageRateOk returns a tuple with the AnnualPercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRateHistoryResponseRowsInner) GetAnnualPercentageRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.AnnualPercentageRate) {
		return nil, false
	}
	return o.AnnualPercentageRate, true
}

// HasAnnualPercentageRate returns a boolean if a field has been set.
func (o *GetBfusdRateHistoryResponseRowsInner) HasAnnualPercentageRate() bool {
	if o != nil && !common.IsNil(o.AnnualPercentageRate) {
		return true
	}

	return false
}

// SetAnnualPercentageRate gets a reference to the given string and assigns it to the AnnualPercentageRate field.
func (o *GetBfusdRateHistoryResponseRowsInner) SetAnnualPercentageRate(v string) {
	o.AnnualPercentageRate = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetBfusdRateHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRateHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetBfusdRateHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetBfusdRateHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

func (o GetBfusdRateHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBfusdRateHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AnnualPercentageRate) {
		toSerialize["annualPercentageRate"] = o.AnnualPercentageRate
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetBfusdRateHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetBfusdRateHistoryResponseRowsInner := _GetBfusdRateHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetBfusdRateHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetBfusdRateHistoryResponseRowsInner(varGetBfusdRateHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "annualPercentageRate")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBfusdRateHistoryResponseRowsInner struct {
	value *GetBfusdRateHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetBfusdRateHistoryResponseRowsInner) Get() *GetBfusdRateHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetBfusdRateHistoryResponseRowsInner) Set(val *GetBfusdRateHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBfusdRateHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBfusdRateHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBfusdRateHistoryResponseRowsInner(val *GetBfusdRateHistoryResponseRowsInner) *NullableGetBfusdRateHistoryResponseRowsInner {
	return &NullableGetBfusdRateHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetBfusdRateHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBfusdRateHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
