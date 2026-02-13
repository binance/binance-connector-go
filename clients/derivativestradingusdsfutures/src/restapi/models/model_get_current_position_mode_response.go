/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetCurrentPositionModeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCurrentPositionModeResponse{}

// GetCurrentPositionModeResponse struct for GetCurrentPositionModeResponse
type GetCurrentPositionModeResponse struct {
	DualSidePosition     *bool `json:"dualSidePosition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCurrentPositionModeResponse GetCurrentPositionModeResponse

// NewGetCurrentPositionModeResponse instantiates a new GetCurrentPositionModeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCurrentPositionModeResponse() *GetCurrentPositionModeResponse {
	this := GetCurrentPositionModeResponse{}
	return &this
}

// NewGetCurrentPositionModeResponseWithDefaults instantiates a new GetCurrentPositionModeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCurrentPositionModeResponseWithDefaults() *GetCurrentPositionModeResponse {
	this := GetCurrentPositionModeResponse{}
	return &this
}

// GetDualSidePosition returns the DualSidePosition field value if set, zero value otherwise.
func (o *GetCurrentPositionModeResponse) GetDualSidePosition() bool {
	if o == nil || common.IsNil(o.DualSidePosition) {
		var ret bool
		return ret
	}
	return *o.DualSidePosition
}

// GetDualSidePositionOk returns a tuple with the DualSidePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentPositionModeResponse) GetDualSidePositionOk() (*bool, bool) {
	if o == nil || common.IsNil(o.DualSidePosition) {
		return nil, false
	}
	return o.DualSidePosition, true
}

// HasDualSidePosition returns a boolean if a field has been set.
func (o *GetCurrentPositionModeResponse) HasDualSidePosition() bool {
	if o != nil && !common.IsNil(o.DualSidePosition) {
		return true
	}

	return false
}

// SetDualSidePosition gets a reference to the given bool and assigns it to the DualSidePosition field.
func (o *GetCurrentPositionModeResponse) SetDualSidePosition(v bool) {
	o.DualSidePosition = &v
}

func (o GetCurrentPositionModeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCurrentPositionModeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DualSidePosition) {
		toSerialize["dualSidePosition"] = o.DualSidePosition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCurrentPositionModeResponse) UnmarshalJSON(data []byte) (err error) {
	varGetCurrentPositionModeResponse := _GetCurrentPositionModeResponse{}

	err = json.Unmarshal(data, &varGetCurrentPositionModeResponse)

	if err != nil {
		return err
	}

	*o = GetCurrentPositionModeResponse(varGetCurrentPositionModeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dualSidePosition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCurrentPositionModeResponse struct {
	value *GetCurrentPositionModeResponse
	isSet bool
}

func (v NullableGetCurrentPositionModeResponse) Get() *GetCurrentPositionModeResponse {
	return v.value
}

func (v *NullableGetCurrentPositionModeResponse) Set(val *GetCurrentPositionModeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCurrentPositionModeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCurrentPositionModeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCurrentPositionModeResponse(val *GetCurrentPositionModeResponse) *NullableGetCurrentPositionModeResponse {
	return &NullableGetCurrentPositionModeResponse{value: val, isSet: true}
}

func (v NullableGetCurrentPositionModeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCurrentPositionModeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
