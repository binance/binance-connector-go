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

// checks if the TickerPriceResponse1Result type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TickerPriceResponse1Result{}

// TickerPriceResponse1Result struct for TickerPriceResponse1Result
type TickerPriceResponse1Result struct {
	Symbol               *string `json:"symbol,omitempty"`
	Price                *string `json:"price,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TickerPriceResponse1Result TickerPriceResponse1Result

// NewTickerPriceResponse1Result instantiates a new TickerPriceResponse1Result object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTickerPriceResponse1Result() *TickerPriceResponse1Result {
	this := TickerPriceResponse1Result{}
	return &this
}

// NewTickerPriceResponse1ResultWithDefaults instantiates a new TickerPriceResponse1Result object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTickerPriceResponse1ResultWithDefaults() *TickerPriceResponse1Result {
	this := TickerPriceResponse1Result{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *TickerPriceResponse1Result) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerPriceResponse1Result) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *TickerPriceResponse1Result) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *TickerPriceResponse1Result) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *TickerPriceResponse1Result) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerPriceResponse1Result) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *TickerPriceResponse1Result) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *TickerPriceResponse1Result) SetPrice(v string) {
	o.Price = &v
}

func (o TickerPriceResponse1Result) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TickerPriceResponse1Result) ToMap() (map[string]interface{}, error) {
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

func (o *TickerPriceResponse1Result) UnmarshalJSON(data []byte) (err error) {
	varTickerPriceResponse1Result := _TickerPriceResponse1Result{}

	err = json.Unmarshal(data, &varTickerPriceResponse1Result)

	if err != nil {
		return err
	}

	*o = TickerPriceResponse1Result(varTickerPriceResponse1Result)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "price")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTickerPriceResponse1Result struct {
	value *TickerPriceResponse1Result
	isSet bool
}

func (v NullableTickerPriceResponse1Result) Get() *TickerPriceResponse1Result {
	return v.value
}

func (v *NullableTickerPriceResponse1Result) Set(val *TickerPriceResponse1Result) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerPriceResponse1Result) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerPriceResponse1Result) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerPriceResponse1Result(val *TickerPriceResponse1Result) *NullableTickerPriceResponse1Result {
	return &NullableTickerPriceResponse1Result{value: val, isSet: true}
}

func (v NullableTickerPriceResponse1Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerPriceResponse1Result) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
