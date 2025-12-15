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

// checks if the GetUmFuturesBnbBurnStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetUmFuturesBnbBurnStatusResponse{}

// GetUmFuturesBnbBurnStatusResponse struct for GetUmFuturesBnbBurnStatusResponse
type GetUmFuturesBnbBurnStatusResponse struct {
	FeeBurn              *bool `json:"feeBurn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUmFuturesBnbBurnStatusResponse GetUmFuturesBnbBurnStatusResponse

// NewGetUmFuturesBnbBurnStatusResponse instantiates a new GetUmFuturesBnbBurnStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUmFuturesBnbBurnStatusResponse() *GetUmFuturesBnbBurnStatusResponse {
	this := GetUmFuturesBnbBurnStatusResponse{}
	return &this
}

// NewGetUmFuturesBnbBurnStatusResponseWithDefaults instantiates a new GetUmFuturesBnbBurnStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUmFuturesBnbBurnStatusResponseWithDefaults() *GetUmFuturesBnbBurnStatusResponse {
	this := GetUmFuturesBnbBurnStatusResponse{}
	return &this
}

// GetFeeBurn returns the FeeBurn field value if set, zero value otherwise.
func (o *GetUmFuturesBnbBurnStatusResponse) GetFeeBurn() bool {
	if o == nil || common.IsNil(o.FeeBurn) {
		var ret bool
		return ret
	}
	return *o.FeeBurn
}

// GetFeeBurnOk returns a tuple with the FeeBurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmFuturesBnbBurnStatusResponse) GetFeeBurnOk() (*bool, bool) {
	if o == nil || common.IsNil(o.FeeBurn) {
		return nil, false
	}
	return o.FeeBurn, true
}

// HasFeeBurn returns a boolean if a field has been set.
func (o *GetUmFuturesBnbBurnStatusResponse) HasFeeBurn() bool {
	if o != nil && !common.IsNil(o.FeeBurn) {
		return true
	}

	return false
}

// SetFeeBurn gets a reference to the given bool and assigns it to the FeeBurn field.
func (o *GetUmFuturesBnbBurnStatusResponse) SetFeeBurn(v bool) {
	o.FeeBurn = &v
}

func (o GetUmFuturesBnbBurnStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUmFuturesBnbBurnStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FeeBurn) {
		toSerialize["feeBurn"] = o.FeeBurn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUmFuturesBnbBurnStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varGetUmFuturesBnbBurnStatusResponse := _GetUmFuturesBnbBurnStatusResponse{}

	err = json.Unmarshal(data, &varGetUmFuturesBnbBurnStatusResponse)

	if err != nil {
		return err
	}

	*o = GetUmFuturesBnbBurnStatusResponse(varGetUmFuturesBnbBurnStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "feeBurn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUmFuturesBnbBurnStatusResponse struct {
	value *GetUmFuturesBnbBurnStatusResponse
	isSet bool
}

func (v NullableGetUmFuturesBnbBurnStatusResponse) Get() *GetUmFuturesBnbBurnStatusResponse {
	return v.value
}

func (v *NullableGetUmFuturesBnbBurnStatusResponse) Set(val *GetUmFuturesBnbBurnStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUmFuturesBnbBurnStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUmFuturesBnbBurnStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUmFuturesBnbBurnStatusResponse(val *GetUmFuturesBnbBurnStatusResponse) *NullableGetUmFuturesBnbBurnStatusResponse {
	return &NullableGetUmFuturesBnbBurnStatusResponse{value: val, isSet: true}
}

func (v NullableGetUmFuturesBnbBurnStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUmFuturesBnbBurnStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
