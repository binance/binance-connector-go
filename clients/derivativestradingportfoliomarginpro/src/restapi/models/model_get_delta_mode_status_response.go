/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetDeltaModeStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDeltaModeStatusResponse{}

// GetDeltaModeStatusResponse struct for GetDeltaModeStatusResponse
type GetDeltaModeStatusResponse struct {
	DeltaEnabled         *bool `json:"deltaEnabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDeltaModeStatusResponse GetDeltaModeStatusResponse

// NewGetDeltaModeStatusResponse instantiates a new GetDeltaModeStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeltaModeStatusResponse() *GetDeltaModeStatusResponse {
	this := GetDeltaModeStatusResponse{}
	return &this
}

// NewGetDeltaModeStatusResponseWithDefaults instantiates a new GetDeltaModeStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeltaModeStatusResponseWithDefaults() *GetDeltaModeStatusResponse {
	this := GetDeltaModeStatusResponse{}
	return &this
}

// GetDeltaEnabled returns the DeltaEnabled field value if set, zero value otherwise.
func (o *GetDeltaModeStatusResponse) GetDeltaEnabled() bool {
	if o == nil || common.IsNil(o.DeltaEnabled) {
		var ret bool
		return ret
	}
	return *o.DeltaEnabled
}

// GetDeltaEnabledOk returns a tuple with the DeltaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeltaModeStatusResponse) GetDeltaEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.DeltaEnabled) {
		return nil, false
	}
	return o.DeltaEnabled, true
}

// HasDeltaEnabled returns a boolean if a field has been set.
func (o *GetDeltaModeStatusResponse) HasDeltaEnabled() bool {
	if o != nil && !common.IsNil(o.DeltaEnabled) {
		return true
	}

	return false
}

// SetDeltaEnabled gets a reference to the given bool and assigns it to the DeltaEnabled field.
func (o *GetDeltaModeStatusResponse) SetDeltaEnabled(v bool) {
	o.DeltaEnabled = &v
}

func (o GetDeltaModeStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeltaModeStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DeltaEnabled) {
		toSerialize["deltaEnabled"] = o.DeltaEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDeltaModeStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varGetDeltaModeStatusResponse := _GetDeltaModeStatusResponse{}

	err = json.Unmarshal(data, &varGetDeltaModeStatusResponse)

	if err != nil {
		return err
	}

	*o = GetDeltaModeStatusResponse(varGetDeltaModeStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deltaEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDeltaModeStatusResponse struct {
	value *GetDeltaModeStatusResponse
	isSet bool
}

func (v NullableGetDeltaModeStatusResponse) Get() *GetDeltaModeStatusResponse {
	return v.value
}

func (v *NullableGetDeltaModeStatusResponse) Set(val *GetDeltaModeStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeltaModeStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeltaModeStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeltaModeStatusResponse(val *GetDeltaModeStatusResponse) *NullableGetDeltaModeStatusResponse {
	return &NullableGetDeltaModeStatusResponse{value: val, isSet: true}
}

func (v NullableGetDeltaModeStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeltaModeStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
