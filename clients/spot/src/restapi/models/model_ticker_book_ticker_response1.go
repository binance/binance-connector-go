/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TickerBookTickerResponse1 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TickerBookTickerResponse1{}

// TickerBookTickerResponse1 struct for TickerBookTickerResponse1
type TickerBookTickerResponse1 struct {
	Symbol               *string `json:"symbol,omitempty"`
	BidPrice             *string `json:"bidPrice,omitempty"`
	BidQty               *string `json:"bidQty,omitempty"`
	AskPrice             *string `json:"askPrice,omitempty"`
	AskQty               *string `json:"askQty,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TickerBookTickerResponse1 TickerBookTickerResponse1

// NewTickerBookTickerResponse1 instantiates a new TickerBookTickerResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTickerBookTickerResponse1() *TickerBookTickerResponse1 {
	this := TickerBookTickerResponse1{}
	return &this
}

// NewTickerBookTickerResponse1WithDefaults instantiates a new TickerBookTickerResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTickerBookTickerResponse1WithDefaults() *TickerBookTickerResponse1 {
	this := TickerBookTickerResponse1{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *TickerBookTickerResponse1) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerBookTickerResponse1) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *TickerBookTickerResponse1) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *TickerBookTickerResponse1) SetSymbol(v string) {
	o.Symbol = &v
}

// GetBidPrice returns the BidPrice field value if set, zero value otherwise.
func (o *TickerBookTickerResponse1) GetBidPrice() string {
	if o == nil || common.IsNil(o.BidPrice) {
		var ret string
		return ret
	}
	return *o.BidPrice
}

// GetBidPriceOk returns a tuple with the BidPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerBookTickerResponse1) GetBidPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidPrice) {
		return nil, false
	}
	return o.BidPrice, true
}

// HasBidPrice returns a boolean if a field has been set.
func (o *TickerBookTickerResponse1) HasBidPrice() bool {
	if o != nil && !common.IsNil(o.BidPrice) {
		return true
	}

	return false
}

// SetBidPrice gets a reference to the given string and assigns it to the BidPrice field.
func (o *TickerBookTickerResponse1) SetBidPrice(v string) {
	o.BidPrice = &v
}

// GetBidQty returns the BidQty field value if set, zero value otherwise.
func (o *TickerBookTickerResponse1) GetBidQty() string {
	if o == nil || common.IsNil(o.BidQty) {
		var ret string
		return ret
	}
	return *o.BidQty
}

// GetBidQtyOk returns a tuple with the BidQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerBookTickerResponse1) GetBidQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidQty) {
		return nil, false
	}
	return o.BidQty, true
}

// HasBidQty returns a boolean if a field has been set.
func (o *TickerBookTickerResponse1) HasBidQty() bool {
	if o != nil && !common.IsNil(o.BidQty) {
		return true
	}

	return false
}

// SetBidQty gets a reference to the given string and assigns it to the BidQty field.
func (o *TickerBookTickerResponse1) SetBidQty(v string) {
	o.BidQty = &v
}

// GetAskPrice returns the AskPrice field value if set, zero value otherwise.
func (o *TickerBookTickerResponse1) GetAskPrice() string {
	if o == nil || common.IsNil(o.AskPrice) {
		var ret string
		return ret
	}
	return *o.AskPrice
}

// GetAskPriceOk returns a tuple with the AskPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerBookTickerResponse1) GetAskPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskPrice) {
		return nil, false
	}
	return o.AskPrice, true
}

// HasAskPrice returns a boolean if a field has been set.
func (o *TickerBookTickerResponse1) HasAskPrice() bool {
	if o != nil && !common.IsNil(o.AskPrice) {
		return true
	}

	return false
}

// SetAskPrice gets a reference to the given string and assigns it to the AskPrice field.
func (o *TickerBookTickerResponse1) SetAskPrice(v string) {
	o.AskPrice = &v
}

// GetAskQty returns the AskQty field value if set, zero value otherwise.
func (o *TickerBookTickerResponse1) GetAskQty() string {
	if o == nil || common.IsNil(o.AskQty) {
		var ret string
		return ret
	}
	return *o.AskQty
}

// GetAskQtyOk returns a tuple with the AskQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerBookTickerResponse1) GetAskQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskQty) {
		return nil, false
	}
	return o.AskQty, true
}

// HasAskQty returns a boolean if a field has been set.
func (o *TickerBookTickerResponse1) HasAskQty() bool {
	if o != nil && !common.IsNil(o.AskQty) {
		return true
	}

	return false
}

// SetAskQty gets a reference to the given string and assigns it to the AskQty field.
func (o *TickerBookTickerResponse1) SetAskQty(v string) {
	o.AskQty = &v
}

func (o TickerBookTickerResponse1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TickerBookTickerResponse1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.BidPrice) {
		toSerialize["bidPrice"] = o.BidPrice
	}
	if !common.IsNil(o.BidQty) {
		toSerialize["bidQty"] = o.BidQty
	}
	if !common.IsNil(o.AskPrice) {
		toSerialize["askPrice"] = o.AskPrice
	}
	if !common.IsNil(o.AskQty) {
		toSerialize["askQty"] = o.AskQty
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TickerBookTickerResponse1) UnmarshalJSON(data []byte) (err error) {
	varTickerBookTickerResponse1 := _TickerBookTickerResponse1{}

	err = json.Unmarshal(data, &varTickerBookTickerResponse1)

	if err != nil {
		return err
	}

	*o = TickerBookTickerResponse1(varTickerBookTickerResponse1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "bidPrice")
		delete(additionalProperties, "bidQty")
		delete(additionalProperties, "askPrice")
		delete(additionalProperties, "askQty")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTickerBookTickerResponse1 struct {
	value *TickerBookTickerResponse1
	isSet bool
}

func (v NullableTickerBookTickerResponse1) Get() *TickerBookTickerResponse1 {
	return v.value
}

func (v *NullableTickerBookTickerResponse1) Set(val *TickerBookTickerResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerBookTickerResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerBookTickerResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerBookTickerResponse1(val *TickerBookTickerResponse1) *NullableTickerBookTickerResponse1 {
	return &NullableTickerBookTickerResponse1{value: val, isSet: true}
}

func (v NullableTickerBookTickerResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerBookTickerResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
