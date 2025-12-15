/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SetAutoCancelAllOpenOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SetAutoCancelAllOpenOrdersResponse{}

// SetAutoCancelAllOpenOrdersResponse struct for SetAutoCancelAllOpenOrdersResponse
type SetAutoCancelAllOpenOrdersResponse struct {
	Underlying           *string `json:"underlying,omitempty"`
	CountdownTime        *int64  `json:"countdownTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SetAutoCancelAllOpenOrdersResponse SetAutoCancelAllOpenOrdersResponse

// NewSetAutoCancelAllOpenOrdersResponse instantiates a new SetAutoCancelAllOpenOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetAutoCancelAllOpenOrdersResponse() *SetAutoCancelAllOpenOrdersResponse {
	this := SetAutoCancelAllOpenOrdersResponse{}
	return &this
}

// NewSetAutoCancelAllOpenOrdersResponseWithDefaults instantiates a new SetAutoCancelAllOpenOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetAutoCancelAllOpenOrdersResponseWithDefaults() *SetAutoCancelAllOpenOrdersResponse {
	this := SetAutoCancelAllOpenOrdersResponse{}
	return &this
}

// GetUnderlying returns the Underlying field value if set, zero value otherwise.
func (o *SetAutoCancelAllOpenOrdersResponse) GetUnderlying() string {
	if o == nil || common.IsNil(o.Underlying) {
		var ret string
		return ret
	}
	return *o.Underlying
}

// GetUnderlyingOk returns a tuple with the Underlying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetAutoCancelAllOpenOrdersResponse) GetUnderlyingOk() (*string, bool) {
	if o == nil || common.IsNil(o.Underlying) {
		return nil, false
	}
	return o.Underlying, true
}

// HasUnderlying returns a boolean if a field has been set.
func (o *SetAutoCancelAllOpenOrdersResponse) HasUnderlying() bool {
	if o != nil && !common.IsNil(o.Underlying) {
		return true
	}

	return false
}

// SetUnderlying gets a reference to the given string and assigns it to the Underlying field.
func (o *SetAutoCancelAllOpenOrdersResponse) SetUnderlying(v string) {
	o.Underlying = &v
}

// GetCountdownTime returns the CountdownTime field value if set, zero value otherwise.
func (o *SetAutoCancelAllOpenOrdersResponse) GetCountdownTime() int64 {
	if o == nil || common.IsNil(o.CountdownTime) {
		var ret int64
		return ret
	}
	return *o.CountdownTime
}

// GetCountdownTimeOk returns a tuple with the CountdownTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetAutoCancelAllOpenOrdersResponse) GetCountdownTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CountdownTime) {
		return nil, false
	}
	return o.CountdownTime, true
}

// HasCountdownTime returns a boolean if a field has been set.
func (o *SetAutoCancelAllOpenOrdersResponse) HasCountdownTime() bool {
	if o != nil && !common.IsNil(o.CountdownTime) {
		return true
	}

	return false
}

// SetCountdownTime gets a reference to the given int64 and assigns it to the CountdownTime field.
func (o *SetAutoCancelAllOpenOrdersResponse) SetCountdownTime(v int64) {
	o.CountdownTime = &v
}

func (o SetAutoCancelAllOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetAutoCancelAllOpenOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Underlying) {
		toSerialize["underlying"] = o.Underlying
	}
	if !common.IsNil(o.CountdownTime) {
		toSerialize["countdownTime"] = o.CountdownTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SetAutoCancelAllOpenOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	varSetAutoCancelAllOpenOrdersResponse := _SetAutoCancelAllOpenOrdersResponse{}

	err = json.Unmarshal(data, &varSetAutoCancelAllOpenOrdersResponse)

	if err != nil {
		return err
	}

	*o = SetAutoCancelAllOpenOrdersResponse(varSetAutoCancelAllOpenOrdersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "underlying")
		delete(additionalProperties, "countdownTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetAutoCancelAllOpenOrdersResponse struct {
	value *SetAutoCancelAllOpenOrdersResponse
	isSet bool
}

func (v NullableSetAutoCancelAllOpenOrdersResponse) Get() *SetAutoCancelAllOpenOrdersResponse {
	return v.value
}

func (v *NullableSetAutoCancelAllOpenOrdersResponse) Set(val *SetAutoCancelAllOpenOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSetAutoCancelAllOpenOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSetAutoCancelAllOpenOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetAutoCancelAllOpenOrdersResponse(val *SetAutoCancelAllOpenOrdersResponse) *NullableSetAutoCancelAllOpenOrdersResponse {
	return &NullableSetAutoCancelAllOpenOrdersResponse{value: val, isSet: true}
}

func (v NullableSetAutoCancelAllOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetAutoCancelAllOpenOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
