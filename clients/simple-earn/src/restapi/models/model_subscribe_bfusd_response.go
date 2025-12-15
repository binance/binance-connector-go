/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SubscribeBfusdResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubscribeBfusdResponse{}

// SubscribeBfusdResponse struct for SubscribeBfusdResponse
type SubscribeBfusdResponse struct {
	Success              *bool   `json:"success,omitempty"`
	BfusdAmount          *string `json:"bfusdAmount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscribeBfusdResponse SubscribeBfusdResponse

// NewSubscribeBfusdResponse instantiates a new SubscribeBfusdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribeBfusdResponse() *SubscribeBfusdResponse {
	this := SubscribeBfusdResponse{}
	return &this
}

// NewSubscribeBfusdResponseWithDefaults instantiates a new SubscribeBfusdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribeBfusdResponseWithDefaults() *SubscribeBfusdResponse {
	this := SubscribeBfusdResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SubscribeBfusdResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeBfusdResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SubscribeBfusdResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SubscribeBfusdResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetBfusdAmount returns the BfusdAmount field value if set, zero value otherwise.
func (o *SubscribeBfusdResponse) GetBfusdAmount() string {
	if o == nil || common.IsNil(o.BfusdAmount) {
		var ret string
		return ret
	}
	return *o.BfusdAmount
}

// GetBfusdAmountOk returns a tuple with the BfusdAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeBfusdResponse) GetBfusdAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.BfusdAmount) {
		return nil, false
	}
	return o.BfusdAmount, true
}

// HasBfusdAmount returns a boolean if a field has been set.
func (o *SubscribeBfusdResponse) HasBfusdAmount() bool {
	if o != nil && !common.IsNil(o.BfusdAmount) {
		return true
	}

	return false
}

// SetBfusdAmount gets a reference to the given string and assigns it to the BfusdAmount field.
func (o *SubscribeBfusdResponse) SetBfusdAmount(v string) {
	o.BfusdAmount = &v
}

func (o SubscribeBfusdResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscribeBfusdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !common.IsNil(o.BfusdAmount) {
		toSerialize["bfusdAmount"] = o.BfusdAmount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscribeBfusdResponse) UnmarshalJSON(data []byte) (err error) {
	varSubscribeBfusdResponse := _SubscribeBfusdResponse{}

	err = json.Unmarshal(data, &varSubscribeBfusdResponse)

	if err != nil {
		return err
	}

	*o = SubscribeBfusdResponse(varSubscribeBfusdResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "bfusdAmount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscribeBfusdResponse struct {
	value *SubscribeBfusdResponse
	isSet bool
}

func (v NullableSubscribeBfusdResponse) Get() *SubscribeBfusdResponse {
	return v.value
}

func (v *NullableSubscribeBfusdResponse) Set(val *SubscribeBfusdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribeBfusdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribeBfusdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribeBfusdResponse(val *SubscribeBfusdResponse) *NullableSubscribeBfusdResponse {
	return &NullableSubscribeBfusdResponse{value: val, isSet: true}
}

func (v NullableSubscribeBfusdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribeBfusdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
