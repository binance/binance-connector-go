/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TimeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TimeResponse{}

// TimeResponse struct for TimeResponse
type TimeResponse struct {
	ServerTime           *int64 `json:"serverTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TimeResponse TimeResponse

// NewTimeResponse instantiates a new TimeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeResponse() *TimeResponse {
	this := TimeResponse{}
	return &this
}

// NewTimeResponseWithDefaults instantiates a new TimeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeResponseWithDefaults() *TimeResponse {
	this := TimeResponse{}
	return &this
}

// GetServerTime returns the ServerTime field value if set, zero value otherwise.
func (o *TimeResponse) GetServerTime() int64 {
	if o == nil || common.IsNil(o.ServerTime) {
		var ret int64
		return ret
	}
	return *o.ServerTime
}

// GetServerTimeOk returns a tuple with the ServerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeResponse) GetServerTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ServerTime) {
		return nil, false
	}
	return o.ServerTime, true
}

// HasServerTime returns a boolean if a field has been set.
func (o *TimeResponse) HasServerTime() bool {
	if o != nil && !common.IsNil(o.ServerTime) {
		return true
	}

	return false
}

// SetServerTime gets a reference to the given int64 and assigns it to the ServerTime field.
func (o *TimeResponse) SetServerTime(v int64) {
	o.ServerTime = &v
}

func (o TimeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ServerTime) {
		toSerialize["serverTime"] = o.ServerTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TimeResponse) UnmarshalJSON(data []byte) (err error) {
	varTimeResponse := _TimeResponse{}

	err = json.Unmarshal(data, &varTimeResponse)

	if err != nil {
		return err
	}

	*o = TimeResponse(varTimeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "serverTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTimeResponse struct {
	value *TimeResponse
	isSet bool
}

func (v NullableTimeResponse) Get() *TimeResponse {
	return v.value
}

func (v *NullableTimeResponse) Set(val *TimeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeResponse(val *TimeResponse) *NullableTimeResponse {
	return &NullableTimeResponse{value: val, isSet: true}
}

func (v NullableTimeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
