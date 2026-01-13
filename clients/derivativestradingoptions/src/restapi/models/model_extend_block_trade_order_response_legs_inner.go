/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ExtendBlockTradeOrderResponseLegsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExtendBlockTradeOrderResponseLegsInner{}

// ExtendBlockTradeOrderResponseLegsInner struct for ExtendBlockTradeOrderResponseLegsInner
type ExtendBlockTradeOrderResponseLegsInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	Side                 *string `json:"side,omitempty"`
	Quantity             *string `json:"quantity,omitempty"`
	Price                *string `json:"price,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExtendBlockTradeOrderResponseLegsInner ExtendBlockTradeOrderResponseLegsInner

// NewExtendBlockTradeOrderResponseLegsInner instantiates a new ExtendBlockTradeOrderResponseLegsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendBlockTradeOrderResponseLegsInner() *ExtendBlockTradeOrderResponseLegsInner {
	this := ExtendBlockTradeOrderResponseLegsInner{}
	return &this
}

// NewExtendBlockTradeOrderResponseLegsInnerWithDefaults instantiates a new ExtendBlockTradeOrderResponseLegsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendBlockTradeOrderResponseLegsInnerWithDefaults() *ExtendBlockTradeOrderResponseLegsInner {
	this := ExtendBlockTradeOrderResponseLegsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ExtendBlockTradeOrderResponseLegsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendBlockTradeOrderResponseLegsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ExtendBlockTradeOrderResponseLegsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ExtendBlockTradeOrderResponseLegsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *ExtendBlockTradeOrderResponseLegsInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendBlockTradeOrderResponseLegsInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *ExtendBlockTradeOrderResponseLegsInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *ExtendBlockTradeOrderResponseLegsInner) SetSide(v string) {
	o.Side = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ExtendBlockTradeOrderResponseLegsInner) GetQuantity() string {
	if o == nil || common.IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendBlockTradeOrderResponseLegsInner) GetQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ExtendBlockTradeOrderResponseLegsInner) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *ExtendBlockTradeOrderResponseLegsInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ExtendBlockTradeOrderResponseLegsInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendBlockTradeOrderResponseLegsInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ExtendBlockTradeOrderResponseLegsInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *ExtendBlockTradeOrderResponseLegsInner) SetPrice(v string) {
	o.Price = &v
}

func (o ExtendBlockTradeOrderResponseLegsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtendBlockTradeOrderResponseLegsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExtendBlockTradeOrderResponseLegsInner) UnmarshalJSON(data []byte) (err error) {
	varExtendBlockTradeOrderResponseLegsInner := _ExtendBlockTradeOrderResponseLegsInner{}

	err = json.Unmarshal(data, &varExtendBlockTradeOrderResponseLegsInner)

	if err != nil {
		return err
	}

	*o = ExtendBlockTradeOrderResponseLegsInner(varExtendBlockTradeOrderResponseLegsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "side")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "price")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExtendBlockTradeOrderResponseLegsInner struct {
	value *ExtendBlockTradeOrderResponseLegsInner
	isSet bool
}

func (v NullableExtendBlockTradeOrderResponseLegsInner) Get() *ExtendBlockTradeOrderResponseLegsInner {
	return v.value
}

func (v *NullableExtendBlockTradeOrderResponseLegsInner) Set(val *ExtendBlockTradeOrderResponseLegsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendBlockTradeOrderResponseLegsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendBlockTradeOrderResponseLegsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendBlockTradeOrderResponseLegsInner(val *ExtendBlockTradeOrderResponseLegsInner) *NullableExtendBlockTradeOrderResponseLegsInner {
	return &NullableExtendBlockTradeOrderResponseLegsInner{value: val, isSet: true}
}

func (v NullableExtendBlockTradeOrderResponseLegsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendBlockTradeOrderResponseLegsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
