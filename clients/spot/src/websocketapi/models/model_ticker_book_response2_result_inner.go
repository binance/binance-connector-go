/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TickerBookResponse2ResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TickerBookResponse2ResultInner{}

// TickerBookResponse2ResultInner struct for TickerBookResponse2ResultInner
type TickerBookResponse2ResultInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	BidPrice             *string `json:"bidPrice,omitempty"`
	BidQty               *string `json:"bidQty,omitempty"`
	AskPrice             *string `json:"askPrice,omitempty"`
	AskQty               *string `json:"askQty,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TickerBookResponse2ResultInner TickerBookResponse2ResultInner

// NewTickerBookResponse2ResultInner instantiates a new TickerBookResponse2ResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTickerBookResponse2ResultInner() *TickerBookResponse2ResultInner {
	this := TickerBookResponse2ResultInner{}
	return &this
}

// NewTickerBookResponse2ResultInnerWithDefaults instantiates a new TickerBookResponse2ResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTickerBookResponse2ResultInnerWithDefaults() *TickerBookResponse2ResultInner {
	this := TickerBookResponse2ResultInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *TickerBookResponse2ResultInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerBookResponse2ResultInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *TickerBookResponse2ResultInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *TickerBookResponse2ResultInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetBidPrice returns the BidPrice field value if set, zero value otherwise.
func (o *TickerBookResponse2ResultInner) GetBidPrice() string {
	if o == nil || common.IsNil(o.BidPrice) {
		var ret string
		return ret
	}
	return *o.BidPrice
}

// GetBidPriceOk returns a tuple with the BidPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerBookResponse2ResultInner) GetBidPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidPrice) {
		return nil, false
	}
	return o.BidPrice, true
}

// HasBidPrice returns a boolean if a field has been set.
func (o *TickerBookResponse2ResultInner) HasBidPrice() bool {
	if o != nil && !common.IsNil(o.BidPrice) {
		return true
	}

	return false
}

// SetBidPrice gets a reference to the given string and assigns it to the BidPrice field.
func (o *TickerBookResponse2ResultInner) SetBidPrice(v string) {
	o.BidPrice = &v
}

// GetBidQty returns the BidQty field value if set, zero value otherwise.
func (o *TickerBookResponse2ResultInner) GetBidQty() string {
	if o == nil || common.IsNil(o.BidQty) {
		var ret string
		return ret
	}
	return *o.BidQty
}

// GetBidQtyOk returns a tuple with the BidQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerBookResponse2ResultInner) GetBidQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidQty) {
		return nil, false
	}
	return o.BidQty, true
}

// HasBidQty returns a boolean if a field has been set.
func (o *TickerBookResponse2ResultInner) HasBidQty() bool {
	if o != nil && !common.IsNil(o.BidQty) {
		return true
	}

	return false
}

// SetBidQty gets a reference to the given string and assigns it to the BidQty field.
func (o *TickerBookResponse2ResultInner) SetBidQty(v string) {
	o.BidQty = &v
}

// GetAskPrice returns the AskPrice field value if set, zero value otherwise.
func (o *TickerBookResponse2ResultInner) GetAskPrice() string {
	if o == nil || common.IsNil(o.AskPrice) {
		var ret string
		return ret
	}
	return *o.AskPrice
}

// GetAskPriceOk returns a tuple with the AskPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerBookResponse2ResultInner) GetAskPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskPrice) {
		return nil, false
	}
	return o.AskPrice, true
}

// HasAskPrice returns a boolean if a field has been set.
func (o *TickerBookResponse2ResultInner) HasAskPrice() bool {
	if o != nil && !common.IsNil(o.AskPrice) {
		return true
	}

	return false
}

// SetAskPrice gets a reference to the given string and assigns it to the AskPrice field.
func (o *TickerBookResponse2ResultInner) SetAskPrice(v string) {
	o.AskPrice = &v
}

// GetAskQty returns the AskQty field value if set, zero value otherwise.
func (o *TickerBookResponse2ResultInner) GetAskQty() string {
	if o == nil || common.IsNil(o.AskQty) {
		var ret string
		return ret
	}
	return *o.AskQty
}

// GetAskQtyOk returns a tuple with the AskQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerBookResponse2ResultInner) GetAskQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskQty) {
		return nil, false
	}
	return o.AskQty, true
}

// HasAskQty returns a boolean if a field has been set.
func (o *TickerBookResponse2ResultInner) HasAskQty() bool {
	if o != nil && !common.IsNil(o.AskQty) {
		return true
	}

	return false
}

// SetAskQty gets a reference to the given string and assigns it to the AskQty field.
func (o *TickerBookResponse2ResultInner) SetAskQty(v string) {
	o.AskQty = &v
}

func (o TickerBookResponse2ResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TickerBookResponse2ResultInner) ToMap() (map[string]interface{}, error) {
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

func (o *TickerBookResponse2ResultInner) UnmarshalJSON(data []byte) (err error) {
	varTickerBookResponse2ResultInner := _TickerBookResponse2ResultInner{}

	err = json.Unmarshal(data, &varTickerBookResponse2ResultInner)

	if err != nil {
		return err
	}

	*o = TickerBookResponse2ResultInner(varTickerBookResponse2ResultInner)

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

type NullableTickerBookResponse2ResultInner struct {
	value *TickerBookResponse2ResultInner
	isSet bool
}

func (v NullableTickerBookResponse2ResultInner) Get() *TickerBookResponse2ResultInner {
	return v.value
}

func (v *NullableTickerBookResponse2ResultInner) Set(val *TickerBookResponse2ResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerBookResponse2ResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerBookResponse2ResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerBookResponse2ResultInner(val *TickerBookResponse2ResultInner) *NullableTickerBookResponse2ResultInner {
	return &NullableTickerBookResponse2ResultInner{value: val, isSet: true}
}

func (v NullableTickerBookResponse2ResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerBookResponse2ResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
