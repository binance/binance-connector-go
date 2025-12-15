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

// checks if the GetUmCurrentPositionModeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetUmCurrentPositionModeResponse{}

// GetUmCurrentPositionModeResponse struct for GetUmCurrentPositionModeResponse
type GetUmCurrentPositionModeResponse struct {
	DualSidePosition     *bool `json:"dualSidePosition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUmCurrentPositionModeResponse GetUmCurrentPositionModeResponse

// NewGetUmCurrentPositionModeResponse instantiates a new GetUmCurrentPositionModeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUmCurrentPositionModeResponse() *GetUmCurrentPositionModeResponse {
	this := GetUmCurrentPositionModeResponse{}
	return &this
}

// NewGetUmCurrentPositionModeResponseWithDefaults instantiates a new GetUmCurrentPositionModeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUmCurrentPositionModeResponseWithDefaults() *GetUmCurrentPositionModeResponse {
	this := GetUmCurrentPositionModeResponse{}
	return &this
}

// GetDualSidePosition returns the DualSidePosition field value if set, zero value otherwise.
func (o *GetUmCurrentPositionModeResponse) GetDualSidePosition() bool {
	if o == nil || common.IsNil(o.DualSidePosition) {
		var ret bool
		return ret
	}
	return *o.DualSidePosition
}

// GetDualSidePositionOk returns a tuple with the DualSidePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmCurrentPositionModeResponse) GetDualSidePositionOk() (*bool, bool) {
	if o == nil || common.IsNil(o.DualSidePosition) {
		return nil, false
	}
	return o.DualSidePosition, true
}

// HasDualSidePosition returns a boolean if a field has been set.
func (o *GetUmCurrentPositionModeResponse) HasDualSidePosition() bool {
	if o != nil && !common.IsNil(o.DualSidePosition) {
		return true
	}

	return false
}

// SetDualSidePosition gets a reference to the given bool and assigns it to the DualSidePosition field.
func (o *GetUmCurrentPositionModeResponse) SetDualSidePosition(v bool) {
	o.DualSidePosition = &v
}

func (o GetUmCurrentPositionModeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUmCurrentPositionModeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DualSidePosition) {
		toSerialize["dualSidePosition"] = o.DualSidePosition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUmCurrentPositionModeResponse) UnmarshalJSON(data []byte) (err error) {
	varGetUmCurrentPositionModeResponse := _GetUmCurrentPositionModeResponse{}

	err = json.Unmarshal(data, &varGetUmCurrentPositionModeResponse)

	if err != nil {
		return err
	}

	*o = GetUmCurrentPositionModeResponse(varGetUmCurrentPositionModeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dualSidePosition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUmCurrentPositionModeResponse struct {
	value *GetUmCurrentPositionModeResponse
	isSet bool
}

func (v NullableGetUmCurrentPositionModeResponse) Get() *GetUmCurrentPositionModeResponse {
	return v.value
}

func (v *NullableGetUmCurrentPositionModeResponse) Set(val *GetUmCurrentPositionModeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUmCurrentPositionModeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUmCurrentPositionModeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUmCurrentPositionModeResponse(val *GetUmCurrentPositionModeResponse) *NullableGetUmCurrentPositionModeResponse {
	return &NullableGetUmCurrentPositionModeResponse{value: val, isSet: true}
}

func (v NullableGetUmCurrentPositionModeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUmCurrentPositionModeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
