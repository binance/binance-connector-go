/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderTestResponseStandardCommissionForOrder type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderTestResponseStandardCommissionForOrder{}

// OrderTestResponseStandardCommissionForOrder struct for OrderTestResponseStandardCommissionForOrder
type OrderTestResponseStandardCommissionForOrder struct {
	Maker                *string `json:"maker,omitempty"`
	Taker                *string `json:"taker,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderTestResponseStandardCommissionForOrder OrderTestResponseStandardCommissionForOrder

// NewOrderTestResponseStandardCommissionForOrder instantiates a new OrderTestResponseStandardCommissionForOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTestResponseStandardCommissionForOrder() *OrderTestResponseStandardCommissionForOrder {
	this := OrderTestResponseStandardCommissionForOrder{}
	return &this
}

// NewOrderTestResponseStandardCommissionForOrderWithDefaults instantiates a new OrderTestResponseStandardCommissionForOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTestResponseStandardCommissionForOrderWithDefaults() *OrderTestResponseStandardCommissionForOrder {
	this := OrderTestResponseStandardCommissionForOrder{}
	return &this
}

// GetMaker returns the Maker field value if set, zero value otherwise.
func (o *OrderTestResponseStandardCommissionForOrder) GetMaker() string {
	if o == nil || common.IsNil(o.Maker) {
		var ret string
		return ret
	}
	return *o.Maker
}

// GetMakerOk returns a tuple with the Maker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTestResponseStandardCommissionForOrder) GetMakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Maker) {
		return nil, false
	}
	return o.Maker, true
}

// HasMaker returns a boolean if a field has been set.
func (o *OrderTestResponseStandardCommissionForOrder) HasMaker() bool {
	if o != nil && !common.IsNil(o.Maker) {
		return true
	}

	return false
}

// SetMaker gets a reference to the given string and assigns it to the Maker field.
func (o *OrderTestResponseStandardCommissionForOrder) SetMaker(v string) {
	o.Maker = &v
}

// GetTaker returns the Taker field value if set, zero value otherwise.
func (o *OrderTestResponseStandardCommissionForOrder) GetTaker() string {
	if o == nil || common.IsNil(o.Taker) {
		var ret string
		return ret
	}
	return *o.Taker
}

// GetTakerOk returns a tuple with the Taker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTestResponseStandardCommissionForOrder) GetTakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Taker) {
		return nil, false
	}
	return o.Taker, true
}

// HasTaker returns a boolean if a field has been set.
func (o *OrderTestResponseStandardCommissionForOrder) HasTaker() bool {
	if o != nil && !common.IsNil(o.Taker) {
		return true
	}

	return false
}

// SetTaker gets a reference to the given string and assigns it to the Taker field.
func (o *OrderTestResponseStandardCommissionForOrder) SetTaker(v string) {
	o.Taker = &v
}

func (o OrderTestResponseStandardCommissionForOrder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTestResponseStandardCommissionForOrder) ToMap() (map[string]interface{}, error) {
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

func (o *OrderTestResponseStandardCommissionForOrder) UnmarshalJSON(data []byte) (err error) {
	varOrderTestResponseStandardCommissionForOrder := _OrderTestResponseStandardCommissionForOrder{}

	err = json.Unmarshal(data, &varOrderTestResponseStandardCommissionForOrder)

	if err != nil {
		return err
	}

	*o = OrderTestResponseStandardCommissionForOrder(varOrderTestResponseStandardCommissionForOrder)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "maker")
		delete(additionalProperties, "taker")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderTestResponseStandardCommissionForOrder struct {
	value *OrderTestResponseStandardCommissionForOrder
	isSet bool
}

func (v NullableOrderTestResponseStandardCommissionForOrder) Get() *OrderTestResponseStandardCommissionForOrder {
	return v.value
}

func (v *NullableOrderTestResponseStandardCommissionForOrder) Set(val *OrderTestResponseStandardCommissionForOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTestResponseStandardCommissionForOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTestResponseStandardCommissionForOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTestResponseStandardCommissionForOrder(val *OrderTestResponseStandardCommissionForOrder) *NullableOrderTestResponseStandardCommissionForOrder {
	return &NullableOrderTestResponseStandardCommissionForOrder{value: val, isSet: true}
}

func (v NullableOrderTestResponseStandardCommissionForOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTestResponseStandardCommissionForOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
