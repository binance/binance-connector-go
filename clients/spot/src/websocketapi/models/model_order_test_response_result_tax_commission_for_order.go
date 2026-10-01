/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OrderTestResponseResultTaxCommissionForOrder type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderTestResponseResultTaxCommissionForOrder{}

// OrderTestResponseResultTaxCommissionForOrder Tax commission rates for trades from the order.
type OrderTestResponseResultTaxCommissionForOrder struct {
	Maker                *string `json:"maker,omitempty"`
	Taker                *string `json:"taker,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderTestResponseResultTaxCommissionForOrder OrderTestResponseResultTaxCommissionForOrder

// NewOrderTestResponseResultTaxCommissionForOrder instantiates a new OrderTestResponseResultTaxCommissionForOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTestResponseResultTaxCommissionForOrder() *OrderTestResponseResultTaxCommissionForOrder {
	this := OrderTestResponseResultTaxCommissionForOrder{}
	return &this
}

// NewOrderTestResponseResultTaxCommissionForOrderWithDefaults instantiates a new OrderTestResponseResultTaxCommissionForOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTestResponseResultTaxCommissionForOrderWithDefaults() *OrderTestResponseResultTaxCommissionForOrder {
	this := OrderTestResponseResultTaxCommissionForOrder{}
	return &this
}

// GetMaker returns the Maker field value if set, zero value otherwise.
func (o *OrderTestResponseResultTaxCommissionForOrder) GetMaker() string {
	if o == nil || common.IsNil(o.Maker) {
		var ret string
		return ret
	}
	return *o.Maker
}

// GetMakerOk returns a tuple with the Maker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTestResponseResultTaxCommissionForOrder) GetMakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Maker) {
		return nil, false
	}
	return o.Maker, true
}

// HasMaker returns a boolean if a field has been set.
func (o *OrderTestResponseResultTaxCommissionForOrder) HasMaker() bool {
	if o != nil && !common.IsNil(o.Maker) {
		return true
	}

	return false
}

// SetMaker gets a reference to the given string and assigns it to the Maker field.
func (o *OrderTestResponseResultTaxCommissionForOrder) SetMaker(v string) {
	o.Maker = &v
}

// GetTaker returns the Taker field value if set, zero value otherwise.
func (o *OrderTestResponseResultTaxCommissionForOrder) GetTaker() string {
	if o == nil || common.IsNil(o.Taker) {
		var ret string
		return ret
	}
	return *o.Taker
}

// GetTakerOk returns a tuple with the Taker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTestResponseResultTaxCommissionForOrder) GetTakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Taker) {
		return nil, false
	}
	return o.Taker, true
}

// HasTaker returns a boolean if a field has been set.
func (o *OrderTestResponseResultTaxCommissionForOrder) HasTaker() bool {
	if o != nil && !common.IsNil(o.Taker) {
		return true
	}

	return false
}

// SetTaker gets a reference to the given string and assigns it to the Taker field.
func (o *OrderTestResponseResultTaxCommissionForOrder) SetTaker(v string) {
	o.Taker = &v
}

func (o OrderTestResponseResultTaxCommissionForOrder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTestResponseResultTaxCommissionForOrder) ToMap() (map[string]interface{}, error) {
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

func (o *OrderTestResponseResultTaxCommissionForOrder) UnmarshalJSON(data []byte) (err error) {
	varOrderTestResponseResultTaxCommissionForOrder := _OrderTestResponseResultTaxCommissionForOrder{}

	err = json.Unmarshal(data, &varOrderTestResponseResultTaxCommissionForOrder)

	if err != nil {
		return err
	}

	*o = OrderTestResponseResultTaxCommissionForOrder(varOrderTestResponseResultTaxCommissionForOrder)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "maker")
		delete(additionalProperties, "taker")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderTestResponseResultTaxCommissionForOrder struct {
	value *OrderTestResponseResultTaxCommissionForOrder
	isSet bool
}

func (v NullableOrderTestResponseResultTaxCommissionForOrder) Get() *OrderTestResponseResultTaxCommissionForOrder {
	return v.value
}

func (v *NullableOrderTestResponseResultTaxCommissionForOrder) Set(val *OrderTestResponseResultTaxCommissionForOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTestResponseResultTaxCommissionForOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTestResponseResultTaxCommissionForOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTestResponseResultTaxCommissionForOrder(val *OrderTestResponseResultTaxCommissionForOrder) *NullableOrderTestResponseResultTaxCommissionForOrder {
	return &NullableOrderTestResponseResultTaxCommissionForOrder{value: val, isSet: true}
}

func (v NullableOrderTestResponseResultTaxCommissionForOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTestResponseResultTaxCommissionForOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
