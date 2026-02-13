/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetCmCurrentPositionModeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCmCurrentPositionModeResponse{}

// GetCmCurrentPositionModeResponse struct for GetCmCurrentPositionModeResponse
type GetCmCurrentPositionModeResponse struct {
	DualSidePosition     *bool `json:"dualSidePosition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCmCurrentPositionModeResponse GetCmCurrentPositionModeResponse

// NewGetCmCurrentPositionModeResponse instantiates a new GetCmCurrentPositionModeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCmCurrentPositionModeResponse() *GetCmCurrentPositionModeResponse {
	this := GetCmCurrentPositionModeResponse{}
	return &this
}

// NewGetCmCurrentPositionModeResponseWithDefaults instantiates a new GetCmCurrentPositionModeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCmCurrentPositionModeResponseWithDefaults() *GetCmCurrentPositionModeResponse {
	this := GetCmCurrentPositionModeResponse{}
	return &this
}

// GetDualSidePosition returns the DualSidePosition field value if set, zero value otherwise.
func (o *GetCmCurrentPositionModeResponse) GetDualSidePosition() bool {
	if o == nil || common.IsNil(o.DualSidePosition) {
		var ret bool
		return ret
	}
	return *o.DualSidePosition
}

// GetDualSidePositionOk returns a tuple with the DualSidePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCmCurrentPositionModeResponse) GetDualSidePositionOk() (*bool, bool) {
	if o == nil || common.IsNil(o.DualSidePosition) {
		return nil, false
	}
	return o.DualSidePosition, true
}

// HasDualSidePosition returns a boolean if a field has been set.
func (o *GetCmCurrentPositionModeResponse) HasDualSidePosition() bool {
	if o != nil && !common.IsNil(o.DualSidePosition) {
		return true
	}

	return false
}

// SetDualSidePosition gets a reference to the given bool and assigns it to the DualSidePosition field.
func (o *GetCmCurrentPositionModeResponse) SetDualSidePosition(v bool) {
	o.DualSidePosition = &v
}

func (o GetCmCurrentPositionModeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCmCurrentPositionModeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DualSidePosition) {
		toSerialize["dualSidePosition"] = o.DualSidePosition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCmCurrentPositionModeResponse) UnmarshalJSON(data []byte) (err error) {
	varGetCmCurrentPositionModeResponse := _GetCmCurrentPositionModeResponse{}

	err = json.Unmarshal(data, &varGetCmCurrentPositionModeResponse)

	if err != nil {
		return err
	}

	*o = GetCmCurrentPositionModeResponse(varGetCmCurrentPositionModeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dualSidePosition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCmCurrentPositionModeResponse struct {
	value *GetCmCurrentPositionModeResponse
	isSet bool
}

func (v NullableGetCmCurrentPositionModeResponse) Get() *GetCmCurrentPositionModeResponse {
	return v.value
}

func (v *NullableGetCmCurrentPositionModeResponse) Set(val *GetCmCurrentPositionModeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCmCurrentPositionModeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCmCurrentPositionModeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCmCurrentPositionModeResponse(val *GetCmCurrentPositionModeResponse) *NullableGetCmCurrentPositionModeResponse {
	return &NullableGetCmCurrentPositionModeResponse{value: val, isSet: true}
}

func (v NullableGetCmCurrentPositionModeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCmCurrentPositionModeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
