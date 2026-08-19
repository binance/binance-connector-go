/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the SignUsEquityDisclaimerResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SignUsEquityDisclaimerResponse{}

// SignUsEquityDisclaimerResponse struct for SignUsEquityDisclaimerResponse
type SignUsEquityDisclaimerResponse struct {
	// `true` = acceptance recorded successfully.
	Success              *bool `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SignUsEquityDisclaimerResponse SignUsEquityDisclaimerResponse

// NewSignUsEquityDisclaimerResponse instantiates a new SignUsEquityDisclaimerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignUsEquityDisclaimerResponse() *SignUsEquityDisclaimerResponse {
	this := SignUsEquityDisclaimerResponse{}
	return &this
}

// NewSignUsEquityDisclaimerResponseWithDefaults instantiates a new SignUsEquityDisclaimerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignUsEquityDisclaimerResponseWithDefaults() *SignUsEquityDisclaimerResponse {
	this := SignUsEquityDisclaimerResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SignUsEquityDisclaimerResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignUsEquityDisclaimerResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SignUsEquityDisclaimerResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SignUsEquityDisclaimerResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o SignUsEquityDisclaimerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignUsEquityDisclaimerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SignUsEquityDisclaimerResponse) UnmarshalJSON(data []byte) (err error) {
	varSignUsEquityDisclaimerResponse := _SignUsEquityDisclaimerResponse{}

	err = json.Unmarshal(data, &varSignUsEquityDisclaimerResponse)

	if err != nil {
		return err
	}

	*o = SignUsEquityDisclaimerResponse(varSignUsEquityDisclaimerResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSignUsEquityDisclaimerResponse struct {
	value *SignUsEquityDisclaimerResponse
	isSet bool
}

func (v NullableSignUsEquityDisclaimerResponse) Get() *SignUsEquityDisclaimerResponse {
	return v.value
}

func (v *NullableSignUsEquityDisclaimerResponse) Set(val *SignUsEquityDisclaimerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSignUsEquityDisclaimerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSignUsEquityDisclaimerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignUsEquityDisclaimerResponse(val *SignUsEquityDisclaimerResponse) *NullableSignUsEquityDisclaimerResponse {
	return &NullableSignUsEquityDisclaimerResponse{value: val, isSet: true}
}

func (v NullableSignUsEquityDisclaimerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignUsEquityDisclaimerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
