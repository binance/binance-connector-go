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

// checks if the AutoCancelAllOpenOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AutoCancelAllOpenOrdersResponse{}

// AutoCancelAllOpenOrdersResponse struct for AutoCancelAllOpenOrdersResponse
type AutoCancelAllOpenOrdersResponse struct {
	Underlyings          []string `json:"underlyings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AutoCancelAllOpenOrdersResponse AutoCancelAllOpenOrdersResponse

// NewAutoCancelAllOpenOrdersResponse instantiates a new AutoCancelAllOpenOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoCancelAllOpenOrdersResponse() *AutoCancelAllOpenOrdersResponse {
	this := AutoCancelAllOpenOrdersResponse{}
	return &this
}

// NewAutoCancelAllOpenOrdersResponseWithDefaults instantiates a new AutoCancelAllOpenOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoCancelAllOpenOrdersResponseWithDefaults() *AutoCancelAllOpenOrdersResponse {
	this := AutoCancelAllOpenOrdersResponse{}
	return &this
}

// GetUnderlyings returns the Underlyings field value if set, zero value otherwise.
func (o *AutoCancelAllOpenOrdersResponse) GetUnderlyings() []string {
	if o == nil || common.IsNil(o.Underlyings) {
		var ret []string
		return ret
	}
	return o.Underlyings
}

// GetUnderlyingsOk returns a tuple with the Underlyings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoCancelAllOpenOrdersResponse) GetUnderlyingsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Underlyings) {
		return nil, false
	}
	return o.Underlyings, true
}

// HasUnderlyings returns a boolean if a field has been set.
func (o *AutoCancelAllOpenOrdersResponse) HasUnderlyings() bool {
	if o != nil && !common.IsNil(o.Underlyings) {
		return true
	}

	return false
}

// SetUnderlyings gets a reference to the given []string and assigns it to the Underlyings field.
func (o *AutoCancelAllOpenOrdersResponse) SetUnderlyings(v []string) {
	o.Underlyings = v
}

func (o AutoCancelAllOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoCancelAllOpenOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Underlyings) {
		toSerialize["underlyings"] = o.Underlyings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AutoCancelAllOpenOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	varAutoCancelAllOpenOrdersResponse := _AutoCancelAllOpenOrdersResponse{}

	err = json.Unmarshal(data, &varAutoCancelAllOpenOrdersResponse)

	if err != nil {
		return err
	}

	*o = AutoCancelAllOpenOrdersResponse(varAutoCancelAllOpenOrdersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "underlyings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAutoCancelAllOpenOrdersResponse struct {
	value *AutoCancelAllOpenOrdersResponse
	isSet bool
}

func (v NullableAutoCancelAllOpenOrdersResponse) Get() *AutoCancelAllOpenOrdersResponse {
	return v.value
}

func (v *NullableAutoCancelAllOpenOrdersResponse) Set(val *AutoCancelAllOpenOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoCancelAllOpenOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoCancelAllOpenOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoCancelAllOpenOrdersResponse(val *AutoCancelAllOpenOrdersResponse) *NullableAutoCancelAllOpenOrdersResponse {
	return &NullableAutoCancelAllOpenOrdersResponse{value: val, isSet: true}
}

func (v NullableAutoCancelAllOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoCancelAllOpenOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
