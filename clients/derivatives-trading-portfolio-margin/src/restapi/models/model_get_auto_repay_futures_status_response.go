/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetAutoRepayFuturesStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetAutoRepayFuturesStatusResponse{}

// GetAutoRepayFuturesStatusResponse struct for GetAutoRepayFuturesStatusResponse
type GetAutoRepayFuturesStatusResponse struct {
	AutoRepay            *bool `json:"autoRepay,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetAutoRepayFuturesStatusResponse GetAutoRepayFuturesStatusResponse

// NewGetAutoRepayFuturesStatusResponse instantiates a new GetAutoRepayFuturesStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAutoRepayFuturesStatusResponse() *GetAutoRepayFuturesStatusResponse {
	this := GetAutoRepayFuturesStatusResponse{}
	return &this
}

// NewGetAutoRepayFuturesStatusResponseWithDefaults instantiates a new GetAutoRepayFuturesStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAutoRepayFuturesStatusResponseWithDefaults() *GetAutoRepayFuturesStatusResponse {
	this := GetAutoRepayFuturesStatusResponse{}
	return &this
}

// GetAutoRepay returns the AutoRepay field value if set, zero value otherwise.
func (o *GetAutoRepayFuturesStatusResponse) GetAutoRepay() bool {
	if o == nil || common.IsNil(o.AutoRepay) {
		var ret bool
		return ret
	}
	return *o.AutoRepay
}

// GetAutoRepayOk returns a tuple with the AutoRepay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAutoRepayFuturesStatusResponse) GetAutoRepayOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AutoRepay) {
		return nil, false
	}
	return o.AutoRepay, true
}

// HasAutoRepay returns a boolean if a field has been set.
func (o *GetAutoRepayFuturesStatusResponse) HasAutoRepay() bool {
	if o != nil && !common.IsNil(o.AutoRepay) {
		return true
	}

	return false
}

// SetAutoRepay gets a reference to the given bool and assigns it to the AutoRepay field.
func (o *GetAutoRepayFuturesStatusResponse) SetAutoRepay(v bool) {
	o.AutoRepay = &v
}

func (o GetAutoRepayFuturesStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAutoRepayFuturesStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AutoRepay) {
		toSerialize["autoRepay"] = o.AutoRepay
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAutoRepayFuturesStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varGetAutoRepayFuturesStatusResponse := _GetAutoRepayFuturesStatusResponse{}

	err = json.Unmarshal(data, &varGetAutoRepayFuturesStatusResponse)

	if err != nil {
		return err
	}

	*o = GetAutoRepayFuturesStatusResponse(varGetAutoRepayFuturesStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "autoRepay")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAutoRepayFuturesStatusResponse struct {
	value *GetAutoRepayFuturesStatusResponse
	isSet bool
}

func (v NullableGetAutoRepayFuturesStatusResponse) Get() *GetAutoRepayFuturesStatusResponse {
	return v.value
}

func (v *NullableGetAutoRepayFuturesStatusResponse) Set(val *GetAutoRepayFuturesStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAutoRepayFuturesStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAutoRepayFuturesStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAutoRepayFuturesStatusResponse(val *GetAutoRepayFuturesStatusResponse) *NullableGetAutoRepayFuturesStatusResponse {
	return &NullableGetAutoRepayFuturesStatusResponse{value: val, isSet: true}
}

func (v NullableGetAutoRepayFuturesStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAutoRepayFuturesStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
