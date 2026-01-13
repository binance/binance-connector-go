/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TickerPriceResponse2Inner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TickerPriceResponse2Inner{}

// TickerPriceResponse2Inner struct for TickerPriceResponse2Inner
type TickerPriceResponse2Inner struct {
	Symbol               *string `json:"symbol,omitempty"`
	Price                *string `json:"price,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TickerPriceResponse2Inner TickerPriceResponse2Inner

// NewTickerPriceResponse2Inner instantiates a new TickerPriceResponse2Inner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTickerPriceResponse2Inner() *TickerPriceResponse2Inner {
	this := TickerPriceResponse2Inner{}
	return &this
}

// NewTickerPriceResponse2InnerWithDefaults instantiates a new TickerPriceResponse2Inner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTickerPriceResponse2InnerWithDefaults() *TickerPriceResponse2Inner {
	this := TickerPriceResponse2Inner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *TickerPriceResponse2Inner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerPriceResponse2Inner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *TickerPriceResponse2Inner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *TickerPriceResponse2Inner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *TickerPriceResponse2Inner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerPriceResponse2Inner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *TickerPriceResponse2Inner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *TickerPriceResponse2Inner) SetPrice(v string) {
	o.Price = &v
}

func (o TickerPriceResponse2Inner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TickerPriceResponse2Inner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TickerPriceResponse2Inner) UnmarshalJSON(data []byte) (err error) {
	varTickerPriceResponse2Inner := _TickerPriceResponse2Inner{}

	err = json.Unmarshal(data, &varTickerPriceResponse2Inner)

	if err != nil {
		return err
	}

	*o = TickerPriceResponse2Inner(varTickerPriceResponse2Inner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "price")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTickerPriceResponse2Inner struct {
	value *TickerPriceResponse2Inner
	isSet bool
}

func (v NullableTickerPriceResponse2Inner) Get() *TickerPriceResponse2Inner {
	return v.value
}

func (v *NullableTickerPriceResponse2Inner) Set(val *TickerPriceResponse2Inner) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerPriceResponse2Inner) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerPriceResponse2Inner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerPriceResponse2Inner(val *TickerPriceResponse2Inner) *NullableTickerPriceResponse2Inner {
	return &NullableTickerPriceResponse2Inner{value: val, isSet: true}
}

func (v NullableTickerPriceResponse2Inner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerPriceResponse2Inner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
