/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CheckServerTimeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckServerTimeResponse{}

// CheckServerTimeResponse struct for CheckServerTimeResponse
type CheckServerTimeResponse struct {
	ServerTime           *int64 `json:"serverTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CheckServerTimeResponse CheckServerTimeResponse

// NewCheckServerTimeResponse instantiates a new CheckServerTimeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckServerTimeResponse() *CheckServerTimeResponse {
	this := CheckServerTimeResponse{}
	return &this
}

// NewCheckServerTimeResponseWithDefaults instantiates a new CheckServerTimeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckServerTimeResponseWithDefaults() *CheckServerTimeResponse {
	this := CheckServerTimeResponse{}
	return &this
}

// GetServerTime returns the ServerTime field value if set, zero value otherwise.
func (o *CheckServerTimeResponse) GetServerTime() int64 {
	if o == nil || common.IsNil(o.ServerTime) {
		var ret int64
		return ret
	}
	return *o.ServerTime
}

// GetServerTimeOk returns a tuple with the ServerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckServerTimeResponse) GetServerTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ServerTime) {
		return nil, false
	}
	return o.ServerTime, true
}

// HasServerTime returns a boolean if a field has been set.
func (o *CheckServerTimeResponse) HasServerTime() bool {
	if o != nil && !common.IsNil(o.ServerTime) {
		return true
	}

	return false
}

// SetServerTime gets a reference to the given int64 and assigns it to the ServerTime field.
func (o *CheckServerTimeResponse) SetServerTime(v int64) {
	o.ServerTime = &v
}

func (o CheckServerTimeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckServerTimeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ServerTime) {
		toSerialize["serverTime"] = o.ServerTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CheckServerTimeResponse) UnmarshalJSON(data []byte) (err error) {
	varCheckServerTimeResponse := _CheckServerTimeResponse{}

	err = json.Unmarshal(data, &varCheckServerTimeResponse)

	if err != nil {
		return err
	}

	*o = CheckServerTimeResponse(varCheckServerTimeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "serverTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCheckServerTimeResponse struct {
	value *CheckServerTimeResponse
	isSet bool
}

func (v NullableCheckServerTimeResponse) Get() *CheckServerTimeResponse {
	return v.value
}

func (v *NullableCheckServerTimeResponse) Set(val *CheckServerTimeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckServerTimeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckServerTimeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckServerTimeResponse(val *CheckServerTimeResponse) *NullableCheckServerTimeResponse {
	return &NullableCheckServerTimeResponse{value: val, isSet: true}
}

func (v NullableCheckServerTimeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckServerTimeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
