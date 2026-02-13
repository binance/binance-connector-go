/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the SubscribeRwusdResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubscribeRwusdResponse{}

// SubscribeRwusdResponse struct for SubscribeRwusdResponse
type SubscribeRwusdResponse struct {
	Success              *bool   `json:"success,omitempty"`
	RwusdAmount          *string `json:"rwusdAmount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscribeRwusdResponse SubscribeRwusdResponse

// NewSubscribeRwusdResponse instantiates a new SubscribeRwusdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribeRwusdResponse() *SubscribeRwusdResponse {
	this := SubscribeRwusdResponse{}
	return &this
}

// NewSubscribeRwusdResponseWithDefaults instantiates a new SubscribeRwusdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribeRwusdResponseWithDefaults() *SubscribeRwusdResponse {
	this := SubscribeRwusdResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SubscribeRwusdResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeRwusdResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SubscribeRwusdResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SubscribeRwusdResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetRwusdAmount returns the RwusdAmount field value if set, zero value otherwise.
func (o *SubscribeRwusdResponse) GetRwusdAmount() string {
	if o == nil || common.IsNil(o.RwusdAmount) {
		var ret string
		return ret
	}
	return *o.RwusdAmount
}

// GetRwusdAmountOk returns a tuple with the RwusdAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeRwusdResponse) GetRwusdAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.RwusdAmount) {
		return nil, false
	}
	return o.RwusdAmount, true
}

// HasRwusdAmount returns a boolean if a field has been set.
func (o *SubscribeRwusdResponse) HasRwusdAmount() bool {
	if o != nil && !common.IsNil(o.RwusdAmount) {
		return true
	}

	return false
}

// SetRwusdAmount gets a reference to the given string and assigns it to the RwusdAmount field.
func (o *SubscribeRwusdResponse) SetRwusdAmount(v string) {
	o.RwusdAmount = &v
}

func (o SubscribeRwusdResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscribeRwusdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !common.IsNil(o.RwusdAmount) {
		toSerialize["rwusdAmount"] = o.RwusdAmount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscribeRwusdResponse) UnmarshalJSON(data []byte) (err error) {
	varSubscribeRwusdResponse := _SubscribeRwusdResponse{}

	err = json.Unmarshal(data, &varSubscribeRwusdResponse)

	if err != nil {
		return err
	}

	*o = SubscribeRwusdResponse(varSubscribeRwusdResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "rwusdAmount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscribeRwusdResponse struct {
	value *SubscribeRwusdResponse
	isSet bool
}

func (v NullableSubscribeRwusdResponse) Get() *SubscribeRwusdResponse {
	return v.value
}

func (v *NullableSubscribeRwusdResponse) Set(val *SubscribeRwusdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribeRwusdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribeRwusdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribeRwusdResponse(val *SubscribeRwusdResponse) *NullableSubscribeRwusdResponse {
	return &NullableSubscribeRwusdResponse{value: val, isSet: true}
}

func (v NullableSubscribeRwusdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribeRwusdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
