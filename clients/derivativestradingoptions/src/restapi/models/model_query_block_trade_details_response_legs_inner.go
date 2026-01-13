/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryBlockTradeDetailsResponseLegsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryBlockTradeDetailsResponseLegsInner{}

// QueryBlockTradeDetailsResponseLegsInner struct for QueryBlockTradeDetailsResponseLegsInner
type QueryBlockTradeDetailsResponseLegsInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	Side                 *string `json:"side,omitempty"`
	Quantity             *string `json:"quantity,omitempty"`
	Price                *string `json:"price,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryBlockTradeDetailsResponseLegsInner QueryBlockTradeDetailsResponseLegsInner

// NewQueryBlockTradeDetailsResponseLegsInner instantiates a new QueryBlockTradeDetailsResponseLegsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryBlockTradeDetailsResponseLegsInner() *QueryBlockTradeDetailsResponseLegsInner {
	this := QueryBlockTradeDetailsResponseLegsInner{}
	return &this
}

// NewQueryBlockTradeDetailsResponseLegsInnerWithDefaults instantiates a new QueryBlockTradeDetailsResponseLegsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryBlockTradeDetailsResponseLegsInnerWithDefaults() *QueryBlockTradeDetailsResponseLegsInner {
	this := QueryBlockTradeDetailsResponseLegsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryBlockTradeDetailsResponseLegsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBlockTradeDetailsResponseLegsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryBlockTradeDetailsResponseLegsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryBlockTradeDetailsResponseLegsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *QueryBlockTradeDetailsResponseLegsInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBlockTradeDetailsResponseLegsInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *QueryBlockTradeDetailsResponseLegsInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *QueryBlockTradeDetailsResponseLegsInner) SetSide(v string) {
	o.Side = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *QueryBlockTradeDetailsResponseLegsInner) GetQuantity() string {
	if o == nil || common.IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBlockTradeDetailsResponseLegsInner) GetQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *QueryBlockTradeDetailsResponseLegsInner) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *QueryBlockTradeDetailsResponseLegsInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QueryBlockTradeDetailsResponseLegsInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryBlockTradeDetailsResponseLegsInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QueryBlockTradeDetailsResponseLegsInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *QueryBlockTradeDetailsResponseLegsInner) SetPrice(v string) {
	o.Price = &v
}

func (o QueryBlockTradeDetailsResponseLegsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryBlockTradeDetailsResponseLegsInner) ToMap() (map[string]interface{}, error) {
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

func (o *QueryBlockTradeDetailsResponseLegsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryBlockTradeDetailsResponseLegsInner := _QueryBlockTradeDetailsResponseLegsInner{}

	err = json.Unmarshal(data, &varQueryBlockTradeDetailsResponseLegsInner)

	if err != nil {
		return err
	}

	*o = QueryBlockTradeDetailsResponseLegsInner(varQueryBlockTradeDetailsResponseLegsInner)

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

type NullableQueryBlockTradeDetailsResponseLegsInner struct {
	value *QueryBlockTradeDetailsResponseLegsInner
	isSet bool
}

func (v NullableQueryBlockTradeDetailsResponseLegsInner) Get() *QueryBlockTradeDetailsResponseLegsInner {
	return v.value
}

func (v *NullableQueryBlockTradeDetailsResponseLegsInner) Set(val *QueryBlockTradeDetailsResponseLegsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryBlockTradeDetailsResponseLegsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryBlockTradeDetailsResponseLegsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryBlockTradeDetailsResponseLegsInner(val *QueryBlockTradeDetailsResponseLegsInner) *NullableQueryBlockTradeDetailsResponseLegsInner {
	return &NullableQueryBlockTradeDetailsResponseLegsInner{value: val, isSet: true}
}

func (v NullableQueryBlockTradeDetailsResponseLegsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryBlockTradeDetailsResponseLegsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
