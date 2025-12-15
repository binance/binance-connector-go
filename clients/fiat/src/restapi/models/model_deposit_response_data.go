/*
Binance Fiat REST API

OpenAPI Specification for the Binance Fiat REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DepositResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DepositResponseData{}

// DepositResponseData struct for DepositResponseData
type DepositResponseData struct {
	OrderId              *string `json:"orderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DepositResponseData DepositResponseData

// NewDepositResponseData instantiates a new DepositResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositResponseData() *DepositResponseData {
	this := DepositResponseData{}
	return &this
}

// NewDepositResponseDataWithDefaults instantiates a new DepositResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositResponseDataWithDefaults() *DepositResponseData {
	this := DepositResponseData{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *DepositResponseData) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositResponseData) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *DepositResponseData) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *DepositResponseData) SetOrderId(v string) {
	o.OrderId = &v
}

func (o DepositResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DepositResponseData) UnmarshalJSON(data []byte) (err error) {
	varDepositResponseData := _DepositResponseData{}

	err = json.Unmarshal(data, &varDepositResponseData)

	if err != nil {
		return err
	}

	*o = DepositResponseData(varDepositResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositResponseData struct {
	value *DepositResponseData
	isSet bool
}

func (v NullableDepositResponseData) Get() *DepositResponseData {
	return v.value
}

func (v *NullableDepositResponseData) Set(val *DepositResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositResponseData(val *DepositResponseData) *NullableDepositResponseData {
	return &NullableDepositResponseData{value: val, isSet: true}
}

func (v NullableDepositResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
