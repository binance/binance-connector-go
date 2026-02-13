/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OrderTestResponseSpecialCommissionForOrder type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderTestResponseSpecialCommissionForOrder{}

// OrderTestResponseSpecialCommissionForOrder struct for OrderTestResponseSpecialCommissionForOrder
type OrderTestResponseSpecialCommissionForOrder struct {
	Maker                *string `json:"maker,omitempty"`
	Taker                *string `json:"taker,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderTestResponseSpecialCommissionForOrder OrderTestResponseSpecialCommissionForOrder

// NewOrderTestResponseSpecialCommissionForOrder instantiates a new OrderTestResponseSpecialCommissionForOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTestResponseSpecialCommissionForOrder() *OrderTestResponseSpecialCommissionForOrder {
	this := OrderTestResponseSpecialCommissionForOrder{}
	return &this
}

// NewOrderTestResponseSpecialCommissionForOrderWithDefaults instantiates a new OrderTestResponseSpecialCommissionForOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTestResponseSpecialCommissionForOrderWithDefaults() *OrderTestResponseSpecialCommissionForOrder {
	this := OrderTestResponseSpecialCommissionForOrder{}
	return &this
}

// GetMaker returns the Maker field value if set, zero value otherwise.
func (o *OrderTestResponseSpecialCommissionForOrder) GetMaker() string {
	if o == nil || common.IsNil(o.Maker) {
		var ret string
		return ret
	}
	return *o.Maker
}

// GetMakerOk returns a tuple with the Maker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTestResponseSpecialCommissionForOrder) GetMakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Maker) {
		return nil, false
	}
	return o.Maker, true
}

// HasMaker returns a boolean if a field has been set.
func (o *OrderTestResponseSpecialCommissionForOrder) HasMaker() bool {
	if o != nil && !common.IsNil(o.Maker) {
		return true
	}

	return false
}

// SetMaker gets a reference to the given string and assigns it to the Maker field.
func (o *OrderTestResponseSpecialCommissionForOrder) SetMaker(v string) {
	o.Maker = &v
}

// GetTaker returns the Taker field value if set, zero value otherwise.
func (o *OrderTestResponseSpecialCommissionForOrder) GetTaker() string {
	if o == nil || common.IsNil(o.Taker) {
		var ret string
		return ret
	}
	return *o.Taker
}

// GetTakerOk returns a tuple with the Taker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTestResponseSpecialCommissionForOrder) GetTakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Taker) {
		return nil, false
	}
	return o.Taker, true
}

// HasTaker returns a boolean if a field has been set.
func (o *OrderTestResponseSpecialCommissionForOrder) HasTaker() bool {
	if o != nil && !common.IsNil(o.Taker) {
		return true
	}

	return false
}

// SetTaker gets a reference to the given string and assigns it to the Taker field.
func (o *OrderTestResponseSpecialCommissionForOrder) SetTaker(v string) {
	o.Taker = &v
}

func (o OrderTestResponseSpecialCommissionForOrder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTestResponseSpecialCommissionForOrder) ToMap() (map[string]interface{}, error) {
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

func (o *OrderTestResponseSpecialCommissionForOrder) UnmarshalJSON(data []byte) (err error) {
	varOrderTestResponseSpecialCommissionForOrder := _OrderTestResponseSpecialCommissionForOrder{}

	err = json.Unmarshal(data, &varOrderTestResponseSpecialCommissionForOrder)

	if err != nil {
		return err
	}

	*o = OrderTestResponseSpecialCommissionForOrder(varOrderTestResponseSpecialCommissionForOrder)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "maker")
		delete(additionalProperties, "taker")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderTestResponseSpecialCommissionForOrder struct {
	value *OrderTestResponseSpecialCommissionForOrder
	isSet bool
}

func (v NullableOrderTestResponseSpecialCommissionForOrder) Get() *OrderTestResponseSpecialCommissionForOrder {
	return v.value
}

func (v *NullableOrderTestResponseSpecialCommissionForOrder) Set(val *OrderTestResponseSpecialCommissionForOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTestResponseSpecialCommissionForOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTestResponseSpecialCommissionForOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTestResponseSpecialCommissionForOrder(val *OrderTestResponseSpecialCommissionForOrder) *NullableOrderTestResponseSpecialCommissionForOrder {
	return &NullableOrderTestResponseSpecialCommissionForOrder{value: val, isSet: true}
}

func (v NullableOrderTestResponseSpecialCommissionForOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTestResponseSpecialCommissionForOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
