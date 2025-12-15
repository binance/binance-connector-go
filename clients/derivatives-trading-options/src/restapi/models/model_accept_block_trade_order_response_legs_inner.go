/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AcceptBlockTradeOrderResponseLegsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AcceptBlockTradeOrderResponseLegsInner{}

// AcceptBlockTradeOrderResponseLegsInner struct for AcceptBlockTradeOrderResponseLegsInner
type AcceptBlockTradeOrderResponseLegsInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	Side                 *string `json:"side,omitempty"`
	Quantity             *string `json:"quantity,omitempty"`
	Price                *string `json:"price,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AcceptBlockTradeOrderResponseLegsInner AcceptBlockTradeOrderResponseLegsInner

// NewAcceptBlockTradeOrderResponseLegsInner instantiates a new AcceptBlockTradeOrderResponseLegsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptBlockTradeOrderResponseLegsInner() *AcceptBlockTradeOrderResponseLegsInner {
	this := AcceptBlockTradeOrderResponseLegsInner{}
	return &this
}

// NewAcceptBlockTradeOrderResponseLegsInnerWithDefaults instantiates a new AcceptBlockTradeOrderResponseLegsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptBlockTradeOrderResponseLegsInnerWithDefaults() *AcceptBlockTradeOrderResponseLegsInner {
	this := AcceptBlockTradeOrderResponseLegsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *AcceptBlockTradeOrderResponseLegsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptBlockTradeOrderResponseLegsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *AcceptBlockTradeOrderResponseLegsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *AcceptBlockTradeOrderResponseLegsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *AcceptBlockTradeOrderResponseLegsInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptBlockTradeOrderResponseLegsInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *AcceptBlockTradeOrderResponseLegsInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *AcceptBlockTradeOrderResponseLegsInner) SetSide(v string) {
	o.Side = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *AcceptBlockTradeOrderResponseLegsInner) GetQuantity() string {
	if o == nil || common.IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptBlockTradeOrderResponseLegsInner) GetQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *AcceptBlockTradeOrderResponseLegsInner) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *AcceptBlockTradeOrderResponseLegsInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *AcceptBlockTradeOrderResponseLegsInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptBlockTradeOrderResponseLegsInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *AcceptBlockTradeOrderResponseLegsInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *AcceptBlockTradeOrderResponseLegsInner) SetPrice(v string) {
	o.Price = &v
}

func (o AcceptBlockTradeOrderResponseLegsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceptBlockTradeOrderResponseLegsInner) ToMap() (map[string]interface{}, error) {
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

func (o *AcceptBlockTradeOrderResponseLegsInner) UnmarshalJSON(data []byte) (err error) {
	varAcceptBlockTradeOrderResponseLegsInner := _AcceptBlockTradeOrderResponseLegsInner{}

	err = json.Unmarshal(data, &varAcceptBlockTradeOrderResponseLegsInner)

	if err != nil {
		return err
	}

	*o = AcceptBlockTradeOrderResponseLegsInner(varAcceptBlockTradeOrderResponseLegsInner)

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

type NullableAcceptBlockTradeOrderResponseLegsInner struct {
	value *AcceptBlockTradeOrderResponseLegsInner
	isSet bool
}

func (v NullableAcceptBlockTradeOrderResponseLegsInner) Get() *AcceptBlockTradeOrderResponseLegsInner {
	return v.value
}

func (v *NullableAcceptBlockTradeOrderResponseLegsInner) Set(val *AcceptBlockTradeOrderResponseLegsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptBlockTradeOrderResponseLegsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptBlockTradeOrderResponseLegsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptBlockTradeOrderResponseLegsInner(val *AcceptBlockTradeOrderResponseLegsInner) *NullableAcceptBlockTradeOrderResponseLegsInner {
	return &NullableAcceptBlockTradeOrderResponseLegsInner{value: val, isSet: true}
}

func (v NullableAcceptBlockTradeOrderResponseLegsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptBlockTradeOrderResponseLegsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
