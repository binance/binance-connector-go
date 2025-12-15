/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderTradeUpdate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderTradeUpdate{}

// OrderTradeUpdate struct for OrderTradeUpdate
type OrderTradeUpdate struct {
	E                    *int64                   `json:"E,omitempty"`
	Smallo               []OrderTradeUpdateOInner `json:"o,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderTradeUpdate OrderTradeUpdate

// NewOrderTradeUpdate instantiates a new OrderTradeUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTradeUpdate() *OrderTradeUpdate {
	this := OrderTradeUpdate{}
	return &this
}

// NewOrderTradeUpdateWithDefaults instantiates a new OrderTradeUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTradeUpdateWithDefaults() *OrderTradeUpdate {
	this := OrderTradeUpdate{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OrderTradeUpdate) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdate) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *OrderTradeUpdate) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *OrderTradeUpdate) SetE(v int64) {
	o.E = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *OrderTradeUpdate) GetSmallo() []OrderTradeUpdateOInner {
	if o == nil || common.IsNil(o.Smallo) {
		var ret []OrderTradeUpdateOInner
		return ret
	}
	return o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdate) GetSmalloOk() ([]OrderTradeUpdateOInner, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *OrderTradeUpdate) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given []OrderTradeUpdateOInner and assigns it to the O field.
func (o *OrderTradeUpdate) SetSmallo(v []OrderTradeUpdateOInner) {
	o.Smallo = v
}

func (o OrderTradeUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTradeUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderTradeUpdate) UnmarshalJSON(data []byte) (err error) {
	varOrderTradeUpdate := _OrderTradeUpdate{}

	err = json.Unmarshal(data, &varOrderTradeUpdate)

	if err != nil {
		return err
	}

	*o = OrderTradeUpdate(varOrderTradeUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "o")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderTradeUpdate struct {
	value *OrderTradeUpdate
	isSet bool
}

func (v NullableOrderTradeUpdate) Get() *OrderTradeUpdate {
	return v.value
}

func (v *NullableOrderTradeUpdate) Set(val *OrderTradeUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTradeUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTradeUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTradeUpdate(val *OrderTradeUpdate) *NullableOrderTradeUpdate {
	return &NullableOrderTradeUpdate{value: val, isSet: true}
}

func (v NullableOrderTradeUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTradeUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
