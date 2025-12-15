/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetBnbBurnStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBnbBurnStatusResponse{}

// GetBnbBurnStatusResponse struct for GetBnbBurnStatusResponse
type GetBnbBurnStatusResponse struct {
	FeeBurn              *bool `json:"feeBurn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBnbBurnStatusResponse GetBnbBurnStatusResponse

// NewGetBnbBurnStatusResponse instantiates a new GetBnbBurnStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBnbBurnStatusResponse() *GetBnbBurnStatusResponse {
	this := GetBnbBurnStatusResponse{}
	return &this
}

// NewGetBnbBurnStatusResponseWithDefaults instantiates a new GetBnbBurnStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBnbBurnStatusResponseWithDefaults() *GetBnbBurnStatusResponse {
	this := GetBnbBurnStatusResponse{}
	return &this
}

// GetFeeBurn returns the FeeBurn field value if set, zero value otherwise.
func (o *GetBnbBurnStatusResponse) GetFeeBurn() bool {
	if o == nil || common.IsNil(o.FeeBurn) {
		var ret bool
		return ret
	}
	return *o.FeeBurn
}

// GetFeeBurnOk returns a tuple with the FeeBurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnbBurnStatusResponse) GetFeeBurnOk() (*bool, bool) {
	if o == nil || common.IsNil(o.FeeBurn) {
		return nil, false
	}
	return o.FeeBurn, true
}

// HasFeeBurn returns a boolean if a field has been set.
func (o *GetBnbBurnStatusResponse) HasFeeBurn() bool {
	if o != nil && !common.IsNil(o.FeeBurn) {
		return true
	}

	return false
}

// SetFeeBurn gets a reference to the given bool and assigns it to the FeeBurn field.
func (o *GetBnbBurnStatusResponse) SetFeeBurn(v bool) {
	o.FeeBurn = &v
}

func (o GetBnbBurnStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBnbBurnStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FeeBurn) {
		toSerialize["feeBurn"] = o.FeeBurn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetBnbBurnStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varGetBnbBurnStatusResponse := _GetBnbBurnStatusResponse{}

	err = json.Unmarshal(data, &varGetBnbBurnStatusResponse)

	if err != nil {
		return err
	}

	*o = GetBnbBurnStatusResponse(varGetBnbBurnStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "feeBurn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBnbBurnStatusResponse struct {
	value *GetBnbBurnStatusResponse
	isSet bool
}

func (v NullableGetBnbBurnStatusResponse) Get() *GetBnbBurnStatusResponse {
	return v.value
}

func (v *NullableGetBnbBurnStatusResponse) Set(val *GetBnbBurnStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBnbBurnStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBnbBurnStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBnbBurnStatusResponse(val *GetBnbBurnStatusResponse) *NullableGetBnbBurnStatusResponse {
	return &NullableGetBnbBurnStatusResponse{value: val, isSet: true}
}

func (v NullableGetBnbBurnStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBnbBurnStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
