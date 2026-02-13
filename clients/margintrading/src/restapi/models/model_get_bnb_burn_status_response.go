/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetBnbBurnStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBnbBurnStatusResponse{}

// GetBnbBurnStatusResponse struct for GetBnbBurnStatusResponse
type GetBnbBurnStatusResponse struct {
	SpotBNBBurn          *bool `json:"spotBNBBurn,omitempty"`
	InterestBNBBurn      *bool `json:"interestBNBBurn,omitempty"`
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

// GetSpotBNBBurn returns the SpotBNBBurn field value if set, zero value otherwise.
func (o *GetBnbBurnStatusResponse) GetSpotBNBBurn() bool {
	if o == nil || common.IsNil(o.SpotBNBBurn) {
		var ret bool
		return ret
	}
	return *o.SpotBNBBurn
}

// GetSpotBNBBurnOk returns a tuple with the SpotBNBBurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnbBurnStatusResponse) GetSpotBNBBurnOk() (*bool, bool) {
	if o == nil || common.IsNil(o.SpotBNBBurn) {
		return nil, false
	}
	return o.SpotBNBBurn, true
}

// HasSpotBNBBurn returns a boolean if a field has been set.
func (o *GetBnbBurnStatusResponse) HasSpotBNBBurn() bool {
	if o != nil && !common.IsNil(o.SpotBNBBurn) {
		return true
	}

	return false
}

// SetSpotBNBBurn gets a reference to the given bool and assigns it to the SpotBNBBurn field.
func (o *GetBnbBurnStatusResponse) SetSpotBNBBurn(v bool) {
	o.SpotBNBBurn = &v
}

// GetInterestBNBBurn returns the InterestBNBBurn field value if set, zero value otherwise.
func (o *GetBnbBurnStatusResponse) GetInterestBNBBurn() bool {
	if o == nil || common.IsNil(o.InterestBNBBurn) {
		var ret bool
		return ret
	}
	return *o.InterestBNBBurn
}

// GetInterestBNBBurnOk returns a tuple with the InterestBNBBurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnbBurnStatusResponse) GetInterestBNBBurnOk() (*bool, bool) {
	if o == nil || common.IsNil(o.InterestBNBBurn) {
		return nil, false
	}
	return o.InterestBNBBurn, true
}

// HasInterestBNBBurn returns a boolean if a field has been set.
func (o *GetBnbBurnStatusResponse) HasInterestBNBBurn() bool {
	if o != nil && !common.IsNil(o.InterestBNBBurn) {
		return true
	}

	return false
}

// SetInterestBNBBurn gets a reference to the given bool and assigns it to the InterestBNBBurn field.
func (o *GetBnbBurnStatusResponse) SetInterestBNBBurn(v bool) {
	o.InterestBNBBurn = &v
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
	if !common.IsNil(o.SpotBNBBurn) {
		toSerialize["spotBNBBurn"] = o.SpotBNBBurn
	}
	if !common.IsNil(o.InterestBNBBurn) {
		toSerialize["interestBNBBurn"] = o.InterestBNBBurn
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
		delete(additionalProperties, "spotBNBBurn")
		delete(additionalProperties, "interestBNBBurn")
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
