/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TimeResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TimeResponseResult{}

// TimeResponseResult struct for TimeResponseResult
type TimeResponseResult struct {
	ServerTime           *int64 `json:"serverTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TimeResponseResult TimeResponseResult

// NewTimeResponseResult instantiates a new TimeResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeResponseResult() *TimeResponseResult {
	this := TimeResponseResult{}
	return &this
}

// NewTimeResponseResultWithDefaults instantiates a new TimeResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeResponseResultWithDefaults() *TimeResponseResult {
	this := TimeResponseResult{}
	return &this
}

// GetServerTime returns the ServerTime field value if set, zero value otherwise.
func (o *TimeResponseResult) GetServerTime() int64 {
	if o == nil || common.IsNil(o.ServerTime) {
		var ret int64
		return ret
	}
	return *o.ServerTime
}

// GetServerTimeOk returns a tuple with the ServerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeResponseResult) GetServerTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ServerTime) {
		return nil, false
	}
	return o.ServerTime, true
}

// HasServerTime returns a boolean if a field has been set.
func (o *TimeResponseResult) HasServerTime() bool {
	if o != nil && !common.IsNil(o.ServerTime) {
		return true
	}

	return false
}

// SetServerTime gets a reference to the given int64 and assigns it to the ServerTime field.
func (o *TimeResponseResult) SetServerTime(v int64) {
	o.ServerTime = &v
}

func (o TimeResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ServerTime) {
		toSerialize["serverTime"] = o.ServerTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TimeResponseResult) UnmarshalJSON(data []byte) (err error) {
	varTimeResponseResult := _TimeResponseResult{}

	err = json.Unmarshal(data, &varTimeResponseResult)

	if err != nil {
		return err
	}

	*o = TimeResponseResult(varTimeResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "serverTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTimeResponseResult struct {
	value *TimeResponseResult
	isSet bool
}

func (v NullableTimeResponseResult) Get() *TimeResponseResult {
	return v.value
}

func (v *NullableTimeResponseResult) Set(val *TimeResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeResponseResult(val *TimeResponseResult) *NullableTimeResponseResult {
	return &NullableTimeResponseResult{value: val, isSet: true}
}

func (v NullableTimeResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
