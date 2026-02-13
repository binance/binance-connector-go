/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TickerPriceResponse2ResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TickerPriceResponse2ResultInner{}

// TickerPriceResponse2ResultInner struct for TickerPriceResponse2ResultInner
type TickerPriceResponse2ResultInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	Price                *string `json:"price,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TickerPriceResponse2ResultInner TickerPriceResponse2ResultInner

// NewTickerPriceResponse2ResultInner instantiates a new TickerPriceResponse2ResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTickerPriceResponse2ResultInner() *TickerPriceResponse2ResultInner {
	this := TickerPriceResponse2ResultInner{}
	return &this
}

// NewTickerPriceResponse2ResultInnerWithDefaults instantiates a new TickerPriceResponse2ResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTickerPriceResponse2ResultInnerWithDefaults() *TickerPriceResponse2ResultInner {
	this := TickerPriceResponse2ResultInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *TickerPriceResponse2ResultInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerPriceResponse2ResultInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *TickerPriceResponse2ResultInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *TickerPriceResponse2ResultInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *TickerPriceResponse2ResultInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerPriceResponse2ResultInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *TickerPriceResponse2ResultInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *TickerPriceResponse2ResultInner) SetPrice(v string) {
	o.Price = &v
}

func (o TickerPriceResponse2ResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TickerPriceResponse2ResultInner) ToMap() (map[string]interface{}, error) {
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

func (o *TickerPriceResponse2ResultInner) UnmarshalJSON(data []byte) (err error) {
	varTickerPriceResponse2ResultInner := _TickerPriceResponse2ResultInner{}

	err = json.Unmarshal(data, &varTickerPriceResponse2ResultInner)

	if err != nil {
		return err
	}

	*o = TickerPriceResponse2ResultInner(varTickerPriceResponse2ResultInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "price")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTickerPriceResponse2ResultInner struct {
	value *TickerPriceResponse2ResultInner
	isSet bool
}

func (v NullableTickerPriceResponse2ResultInner) Get() *TickerPriceResponse2ResultInner {
	return v.value
}

func (v *NullableTickerPriceResponse2ResultInner) Set(val *TickerPriceResponse2ResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerPriceResponse2ResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerPriceResponse2ResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerPriceResponse2ResultInner(val *TickerPriceResponse2ResultInner) *NullableTickerPriceResponse2ResultInner {
	return &NullableTickerPriceResponse2ResultInner{value: val, isSet: true}
}

func (v NullableTickerPriceResponse2ResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerPriceResponse2ResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
