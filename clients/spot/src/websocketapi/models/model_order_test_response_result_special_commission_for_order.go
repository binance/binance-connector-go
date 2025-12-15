/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderTestResponseResultSpecialCommissionForOrder type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderTestResponseResultSpecialCommissionForOrder{}

// OrderTestResponseResultSpecialCommissionForOrder struct for OrderTestResponseResultSpecialCommissionForOrder
type OrderTestResponseResultSpecialCommissionForOrder struct {
	Maker                *string `json:"maker,omitempty"`
	Taker                *string `json:"taker,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderTestResponseResultSpecialCommissionForOrder OrderTestResponseResultSpecialCommissionForOrder

// NewOrderTestResponseResultSpecialCommissionForOrder instantiates a new OrderTestResponseResultSpecialCommissionForOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTestResponseResultSpecialCommissionForOrder() *OrderTestResponseResultSpecialCommissionForOrder {
	this := OrderTestResponseResultSpecialCommissionForOrder{}
	return &this
}

// NewOrderTestResponseResultSpecialCommissionForOrderWithDefaults instantiates a new OrderTestResponseResultSpecialCommissionForOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTestResponseResultSpecialCommissionForOrderWithDefaults() *OrderTestResponseResultSpecialCommissionForOrder {
	this := OrderTestResponseResultSpecialCommissionForOrder{}
	return &this
}

// GetMaker returns the Maker field value if set, zero value otherwise.
func (o *OrderTestResponseResultSpecialCommissionForOrder) GetMaker() string {
	if o == nil || common.IsNil(o.Maker) {
		var ret string
		return ret
	}
	return *o.Maker
}

// GetMakerOk returns a tuple with the Maker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTestResponseResultSpecialCommissionForOrder) GetMakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Maker) {
		return nil, false
	}
	return o.Maker, true
}

// HasMaker returns a boolean if a field has been set.
func (o *OrderTestResponseResultSpecialCommissionForOrder) HasMaker() bool {
	if o != nil && !common.IsNil(o.Maker) {
		return true
	}

	return false
}

// SetMaker gets a reference to the given string and assigns it to the Maker field.
func (o *OrderTestResponseResultSpecialCommissionForOrder) SetMaker(v string) {
	o.Maker = &v
}

// GetTaker returns the Taker field value if set, zero value otherwise.
func (o *OrderTestResponseResultSpecialCommissionForOrder) GetTaker() string {
	if o == nil || common.IsNil(o.Taker) {
		var ret string
		return ret
	}
	return *o.Taker
}

// GetTakerOk returns a tuple with the Taker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTestResponseResultSpecialCommissionForOrder) GetTakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Taker) {
		return nil, false
	}
	return o.Taker, true
}

// HasTaker returns a boolean if a field has been set.
func (o *OrderTestResponseResultSpecialCommissionForOrder) HasTaker() bool {
	if o != nil && !common.IsNil(o.Taker) {
		return true
	}

	return false
}

// SetTaker gets a reference to the given string and assigns it to the Taker field.
func (o *OrderTestResponseResultSpecialCommissionForOrder) SetTaker(v string) {
	o.Taker = &v
}

func (o OrderTestResponseResultSpecialCommissionForOrder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTestResponseResultSpecialCommissionForOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Maker) {
		toSerialize["maker"] = o.Maker
	}
	if !common.IsNil(o.Taker) {
		toSerialize["taker"] = o.Taker
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderTestResponseResultSpecialCommissionForOrder) UnmarshalJSON(data []byte) (err error) {
	varOrderTestResponseResultSpecialCommissionForOrder := _OrderTestResponseResultSpecialCommissionForOrder{}

	err = json.Unmarshal(data, &varOrderTestResponseResultSpecialCommissionForOrder)

	if err != nil {
		return err
	}

	*o = OrderTestResponseResultSpecialCommissionForOrder(varOrderTestResponseResultSpecialCommissionForOrder)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "maker")
		delete(additionalProperties, "taker")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderTestResponseResultSpecialCommissionForOrder struct {
	value *OrderTestResponseResultSpecialCommissionForOrder
	isSet bool
}

func (v NullableOrderTestResponseResultSpecialCommissionForOrder) Get() *OrderTestResponseResultSpecialCommissionForOrder {
	return v.value
}

func (v *NullableOrderTestResponseResultSpecialCommissionForOrder) Set(val *OrderTestResponseResultSpecialCommissionForOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTestResponseResultSpecialCommissionForOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTestResponseResultSpecialCommissionForOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTestResponseResultSpecialCommissionForOrder(val *OrderTestResponseResultSpecialCommissionForOrder) *NullableOrderTestResponseResultSpecialCommissionForOrder {
	return &NullableOrderTestResponseResultSpecialCommissionForOrder{value: val, isSet: true}
}

func (v NullableOrderTestResponseResultSpecialCommissionForOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTestResponseResultSpecialCommissionForOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
