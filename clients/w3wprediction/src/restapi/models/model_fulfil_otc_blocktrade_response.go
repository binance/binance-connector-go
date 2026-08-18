/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the FulfilOtcBlocktradeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FulfilOtcBlocktradeResponse{}

// FulfilOtcBlocktradeResponse struct for FulfilOtcBlocktradeResponse
type FulfilOtcBlocktradeResponse struct {
	OrderId              *string `json:"orderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FulfilOtcBlocktradeResponse FulfilOtcBlocktradeResponse

// NewFulfilOtcBlocktradeResponse instantiates a new FulfilOtcBlocktradeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfilOtcBlocktradeResponse() *FulfilOtcBlocktradeResponse {
	this := FulfilOtcBlocktradeResponse{}
	return &this
}

// NewFulfilOtcBlocktradeResponseWithDefaults instantiates a new FulfilOtcBlocktradeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfilOtcBlocktradeResponseWithDefaults() *FulfilOtcBlocktradeResponse {
	this := FulfilOtcBlocktradeResponse{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *FulfilOtcBlocktradeResponse) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfilOtcBlocktradeResponse) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *FulfilOtcBlocktradeResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *FulfilOtcBlocktradeResponse) SetOrderId(v string) {
	o.OrderId = &v
}

func (o FulfilOtcBlocktradeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FulfilOtcBlocktradeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FulfilOtcBlocktradeResponse) UnmarshalJSON(data []byte) (err error) {
	varFulfilOtcBlocktradeResponse := _FulfilOtcBlocktradeResponse{}

	err = json.Unmarshal(data, &varFulfilOtcBlocktradeResponse)

	if err != nil {
		return err
	}

	*o = FulfilOtcBlocktradeResponse(varFulfilOtcBlocktradeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFulfilOtcBlocktradeResponse struct {
	value *FulfilOtcBlocktradeResponse
	isSet bool
}

func (v NullableFulfilOtcBlocktradeResponse) Get() *FulfilOtcBlocktradeResponse {
	return v.value
}

func (v *NullableFulfilOtcBlocktradeResponse) Set(val *FulfilOtcBlocktradeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfilOtcBlocktradeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfilOtcBlocktradeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfilOtcBlocktradeResponse(val *FulfilOtcBlocktradeResponse) *NullableFulfilOtcBlocktradeResponse {
	return &NullableFulfilOtcBlocktradeResponse{value: val, isSet: true}
}

func (v NullableFulfilOtcBlocktradeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfilOtcBlocktradeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
