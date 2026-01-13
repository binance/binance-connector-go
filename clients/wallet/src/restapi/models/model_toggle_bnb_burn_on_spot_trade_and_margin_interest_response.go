/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ToggleBnbBurnOnSpotTradeAndMarginInterestResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ToggleBnbBurnOnSpotTradeAndMarginInterestResponse{}

// ToggleBnbBurnOnSpotTradeAndMarginInterestResponse struct for ToggleBnbBurnOnSpotTradeAndMarginInterestResponse
type ToggleBnbBurnOnSpotTradeAndMarginInterestResponse struct {
	SpotBNBBurn          *bool `json:"spotBNBBurn,omitempty"`
	InterestBNBBurn      *bool `json:"interestBNBBurn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ToggleBnbBurnOnSpotTradeAndMarginInterestResponse ToggleBnbBurnOnSpotTradeAndMarginInterestResponse

// NewToggleBnbBurnOnSpotTradeAndMarginInterestResponse instantiates a new ToggleBnbBurnOnSpotTradeAndMarginInterestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToggleBnbBurnOnSpotTradeAndMarginInterestResponse() *ToggleBnbBurnOnSpotTradeAndMarginInterestResponse {
	this := ToggleBnbBurnOnSpotTradeAndMarginInterestResponse{}
	return &this
}

// NewToggleBnbBurnOnSpotTradeAndMarginInterestResponseWithDefaults instantiates a new ToggleBnbBurnOnSpotTradeAndMarginInterestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToggleBnbBurnOnSpotTradeAndMarginInterestResponseWithDefaults() *ToggleBnbBurnOnSpotTradeAndMarginInterestResponse {
	this := ToggleBnbBurnOnSpotTradeAndMarginInterestResponse{}
	return &this
}

// GetSpotBNBBurn returns the SpotBNBBurn field value if set, zero value otherwise.
func (o *ToggleBnbBurnOnSpotTradeAndMarginInterestResponse) GetSpotBNBBurn() bool {
	if o == nil || common.IsNil(o.SpotBNBBurn) {
		var ret bool
		return ret
	}
	return *o.SpotBNBBurn
}

// GetSpotBNBBurnOk returns a tuple with the SpotBNBBurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToggleBnbBurnOnSpotTradeAndMarginInterestResponse) GetSpotBNBBurnOk() (*bool, bool) {
	if o == nil || common.IsNil(o.SpotBNBBurn) {
		return nil, false
	}
	return o.SpotBNBBurn, true
}

// HasSpotBNBBurn returns a boolean if a field has been set.
func (o *ToggleBnbBurnOnSpotTradeAndMarginInterestResponse) HasSpotBNBBurn() bool {
	if o != nil && !common.IsNil(o.SpotBNBBurn) {
		return true
	}

	return false
}

// SetSpotBNBBurn gets a reference to the given bool and assigns it to the SpotBNBBurn field.
func (o *ToggleBnbBurnOnSpotTradeAndMarginInterestResponse) SetSpotBNBBurn(v bool) {
	o.SpotBNBBurn = &v
}

// GetInterestBNBBurn returns the InterestBNBBurn field value if set, zero value otherwise.
func (o *ToggleBnbBurnOnSpotTradeAndMarginInterestResponse) GetInterestBNBBurn() bool {
	if o == nil || common.IsNil(o.InterestBNBBurn) {
		var ret bool
		return ret
	}
	return *o.InterestBNBBurn
}

// GetInterestBNBBurnOk returns a tuple with the InterestBNBBurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToggleBnbBurnOnSpotTradeAndMarginInterestResponse) GetInterestBNBBurnOk() (*bool, bool) {
	if o == nil || common.IsNil(o.InterestBNBBurn) {
		return nil, false
	}
	return o.InterestBNBBurn, true
}

// HasInterestBNBBurn returns a boolean if a field has been set.
func (o *ToggleBnbBurnOnSpotTradeAndMarginInterestResponse) HasInterestBNBBurn() bool {
	if o != nil && !common.IsNil(o.InterestBNBBurn) {
		return true
	}

	return false
}

// SetInterestBNBBurn gets a reference to the given bool and assigns it to the InterestBNBBurn field.
func (o *ToggleBnbBurnOnSpotTradeAndMarginInterestResponse) SetInterestBNBBurn(v bool) {
	o.InterestBNBBurn = &v
}

func (o ToggleBnbBurnOnSpotTradeAndMarginInterestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToggleBnbBurnOnSpotTradeAndMarginInterestResponse) ToMap() (map[string]interface{}, error) {
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

func (o *ToggleBnbBurnOnSpotTradeAndMarginInterestResponse) UnmarshalJSON(data []byte) (err error) {
	varToggleBnbBurnOnSpotTradeAndMarginInterestResponse := _ToggleBnbBurnOnSpotTradeAndMarginInterestResponse{}

	err = json.Unmarshal(data, &varToggleBnbBurnOnSpotTradeAndMarginInterestResponse)

	if err != nil {
		return err
	}

	*o = ToggleBnbBurnOnSpotTradeAndMarginInterestResponse(varToggleBnbBurnOnSpotTradeAndMarginInterestResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "spotBNBBurn")
		delete(additionalProperties, "interestBNBBurn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableToggleBnbBurnOnSpotTradeAndMarginInterestResponse struct {
	value *ToggleBnbBurnOnSpotTradeAndMarginInterestResponse
	isSet bool
}

func (v NullableToggleBnbBurnOnSpotTradeAndMarginInterestResponse) Get() *ToggleBnbBurnOnSpotTradeAndMarginInterestResponse {
	return v.value
}

func (v *NullableToggleBnbBurnOnSpotTradeAndMarginInterestResponse) Set(val *ToggleBnbBurnOnSpotTradeAndMarginInterestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableToggleBnbBurnOnSpotTradeAndMarginInterestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableToggleBnbBurnOnSpotTradeAndMarginInterestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToggleBnbBurnOnSpotTradeAndMarginInterestResponse(val *ToggleBnbBurnOnSpotTradeAndMarginInterestResponse) *NullableToggleBnbBurnOnSpotTradeAndMarginInterestResponse {
	return &NullableToggleBnbBurnOnSpotTradeAndMarginInterestResponse{value: val, isSet: true}
}

func (v NullableToggleBnbBurnOnSpotTradeAndMarginInterestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToggleBnbBurnOnSpotTradeAndMarginInterestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
