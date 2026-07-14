/*
Options REST API

Access market data, manage accounts, and trade Binance Options.
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the NewBlockTradeOrderLegsParameterInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NewBlockTradeOrderLegsParameterInner{}

// NewBlockTradeOrderLegsParameterInner struct for NewBlockTradeOrderLegsParameterInner
type NewBlockTradeOrderLegsParameterInner struct {
	// Option trading pair
	Symbol string                                   `json:"symbol"`
	Side   NewBlockTradeOrderLegsParameterInnerSide `json:"side"`
	Type   NewBlockTradeOrderLegsParameterInnerType `json:"type"`
	// Order quantity
	Quantity string `json:"quantity"`
	// Order price
	Price                *string `json:"price,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NewBlockTradeOrderLegsParameterInner NewBlockTradeOrderLegsParameterInner

// NewNewBlockTradeOrderLegsParameterInner instantiates a new NewBlockTradeOrderLegsParameterInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewBlockTradeOrderLegsParameterInner(symbol string, side NewBlockTradeOrderLegsParameterInnerSide, type_ NewBlockTradeOrderLegsParameterInnerType, quantity string) *NewBlockTradeOrderLegsParameterInner {
	this := NewBlockTradeOrderLegsParameterInner{}
	this.Symbol = symbol
	this.Side = side
	this.Type = type_
	this.Quantity = quantity
	return &this
}

// NewNewBlockTradeOrderLegsParameterInnerWithDefaults instantiates a new NewBlockTradeOrderLegsParameterInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewBlockTradeOrderLegsParameterInnerWithDefaults() *NewBlockTradeOrderLegsParameterInner {
	this := NewBlockTradeOrderLegsParameterInner{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *NewBlockTradeOrderLegsParameterInner) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *NewBlockTradeOrderLegsParameterInner) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *NewBlockTradeOrderLegsParameterInner) SetSymbol(v string) {
	o.Symbol = v
}

// GetSide returns the Side field value
func (o *NewBlockTradeOrderLegsParameterInner) GetSide() NewBlockTradeOrderLegsParameterInnerSide {
	if o == nil {
		var ret NewBlockTradeOrderLegsParameterInnerSide
		return ret
	}

	return o.Side
}

// GetSideOk returns a tuple with the Side field value
// and a boolean to check if the value has been set.
func (o *NewBlockTradeOrderLegsParameterInner) GetSideOk() (*NewBlockTradeOrderLegsParameterInnerSide, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Side, true
}

// SetSide sets field value
func (o *NewBlockTradeOrderLegsParameterInner) SetSide(v NewBlockTradeOrderLegsParameterInnerSide) {
	o.Side = v
}

// GetType returns the Type field value
func (o *NewBlockTradeOrderLegsParameterInner) GetType() NewBlockTradeOrderLegsParameterInnerType {
	if o == nil {
		var ret NewBlockTradeOrderLegsParameterInnerType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NewBlockTradeOrderLegsParameterInner) GetTypeOk() (*NewBlockTradeOrderLegsParameterInnerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NewBlockTradeOrderLegsParameterInner) SetType(v NewBlockTradeOrderLegsParameterInnerType) {
	o.Type = v
}

// GetQuantity returns the Quantity field value
func (o *NewBlockTradeOrderLegsParameterInner) GetQuantity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *NewBlockTradeOrderLegsParameterInner) GetQuantityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *NewBlockTradeOrderLegsParameterInner) SetQuantity(v string) {
	o.Quantity = v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *NewBlockTradeOrderLegsParameterInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewBlockTradeOrderLegsParameterInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *NewBlockTradeOrderLegsParameterInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *NewBlockTradeOrderLegsParameterInner) SetPrice(v string) {
	o.Price = &v
}

func (o NewBlockTradeOrderLegsParameterInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewBlockTradeOrderLegsParameterInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["side"] = o.Side
	toSerialize["type"] = o.Type
	toSerialize["quantity"] = o.Quantity
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NewBlockTradeOrderLegsParameterInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"side",
		"type",
		"quantity",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNewBlockTradeOrderLegsParameterInner := _NewBlockTradeOrderLegsParameterInner{}

	err = json.Unmarshal(data, &varNewBlockTradeOrderLegsParameterInner)

	if err != nil {
		return err
	}

	*o = NewBlockTradeOrderLegsParameterInner(varNewBlockTradeOrderLegsParameterInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "side")
		delete(additionalProperties, "type")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "price")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNewBlockTradeOrderLegsParameterInner struct {
	value *NewBlockTradeOrderLegsParameterInner
	isSet bool
}

func (v NullableNewBlockTradeOrderLegsParameterInner) Get() *NewBlockTradeOrderLegsParameterInner {
	return v.value
}

func (v *NullableNewBlockTradeOrderLegsParameterInner) Set(val *NewBlockTradeOrderLegsParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableNewBlockTradeOrderLegsParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableNewBlockTradeOrderLegsParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewBlockTradeOrderLegsParameterInner(val *NewBlockTradeOrderLegsParameterInner) *NullableNewBlockTradeOrderLegsParameterInner {
	return &NullableNewBlockTradeOrderLegsParameterInner{value: val, isSet: true}
}

func (v NullableNewBlockTradeOrderLegsParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewBlockTradeOrderLegsParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
