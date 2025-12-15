/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SymbolOrderBookTickerResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolOrderBookTickerResponseInner{}

// SymbolOrderBookTickerResponseInner struct for SymbolOrderBookTickerResponseInner
type SymbolOrderBookTickerResponseInner struct {
	LastUpdateId         *int64  `json:"lastUpdateId,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	Pair                 *string `json:"pair,omitempty"`
	BidPrice             *string `json:"bidPrice,omitempty"`
	BidQty               *string `json:"bidQty,omitempty"`
	AskPrice             *string `json:"askPrice,omitempty"`
	AskQty               *string `json:"askQty,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SymbolOrderBookTickerResponseInner SymbolOrderBookTickerResponseInner

// NewSymbolOrderBookTickerResponseInner instantiates a new SymbolOrderBookTickerResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolOrderBookTickerResponseInner() *SymbolOrderBookTickerResponseInner {
	this := SymbolOrderBookTickerResponseInner{}
	return &this
}

// NewSymbolOrderBookTickerResponseInnerWithDefaults instantiates a new SymbolOrderBookTickerResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolOrderBookTickerResponseInnerWithDefaults() *SymbolOrderBookTickerResponseInner {
	this := SymbolOrderBookTickerResponseInner{}
	return &this
}

// GetLastUpdateId returns the LastUpdateId field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponseInner) GetLastUpdateId() int64 {
	if o == nil || common.IsNil(o.LastUpdateId) {
		var ret int64
		return ret
	}
	return *o.LastUpdateId
}

// GetLastUpdateIdOk returns a tuple with the LastUpdateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponseInner) GetLastUpdateIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastUpdateId) {
		return nil, false
	}
	return o.LastUpdateId, true
}

// HasLastUpdateId returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponseInner) HasLastUpdateId() bool {
	if o != nil && !common.IsNil(o.LastUpdateId) {
		return true
	}

	return false
}

// SetLastUpdateId gets a reference to the given int64 and assigns it to the LastUpdateId field.
func (o *SymbolOrderBookTickerResponseInner) SetLastUpdateId(v int64) {
	o.LastUpdateId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SymbolOrderBookTickerResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponseInner) GetPair() string {
	if o == nil || common.IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponseInner) GetPairOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponseInner) HasPair() bool {
	if o != nil && !common.IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *SymbolOrderBookTickerResponseInner) SetPair(v string) {
	o.Pair = &v
}

// GetBidPrice returns the BidPrice field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponseInner) GetBidPrice() string {
	if o == nil || common.IsNil(o.BidPrice) {
		var ret string
		return ret
	}
	return *o.BidPrice
}

// GetBidPriceOk returns a tuple with the BidPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponseInner) GetBidPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidPrice) {
		return nil, false
	}
	return o.BidPrice, true
}

// HasBidPrice returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponseInner) HasBidPrice() bool {
	if o != nil && !common.IsNil(o.BidPrice) {
		return true
	}

	return false
}

// SetBidPrice gets a reference to the given string and assigns it to the BidPrice field.
func (o *SymbolOrderBookTickerResponseInner) SetBidPrice(v string) {
	o.BidPrice = &v
}

// GetBidQty returns the BidQty field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponseInner) GetBidQty() string {
	if o == nil || common.IsNil(o.BidQty) {
		var ret string
		return ret
	}
	return *o.BidQty
}

// GetBidQtyOk returns a tuple with the BidQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponseInner) GetBidQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidQty) {
		return nil, false
	}
	return o.BidQty, true
}

// HasBidQty returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponseInner) HasBidQty() bool {
	if o != nil && !common.IsNil(o.BidQty) {
		return true
	}

	return false
}

// SetBidQty gets a reference to the given string and assigns it to the BidQty field.
func (o *SymbolOrderBookTickerResponseInner) SetBidQty(v string) {
	o.BidQty = &v
}

// GetAskPrice returns the AskPrice field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponseInner) GetAskPrice() string {
	if o == nil || common.IsNil(o.AskPrice) {
		var ret string
		return ret
	}
	return *o.AskPrice
}

// GetAskPriceOk returns a tuple with the AskPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponseInner) GetAskPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskPrice) {
		return nil, false
	}
	return o.AskPrice, true
}

// HasAskPrice returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponseInner) HasAskPrice() bool {
	if o != nil && !common.IsNil(o.AskPrice) {
		return true
	}

	return false
}

// SetAskPrice gets a reference to the given string and assigns it to the AskPrice field.
func (o *SymbolOrderBookTickerResponseInner) SetAskPrice(v string) {
	o.AskPrice = &v
}

// GetAskQty returns the AskQty field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponseInner) GetAskQty() string {
	if o == nil || common.IsNil(o.AskQty) {
		var ret string
		return ret
	}
	return *o.AskQty
}

// GetAskQtyOk returns a tuple with the AskQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponseInner) GetAskQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskQty) {
		return nil, false
	}
	return o.AskQty, true
}

// HasAskQty returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponseInner) HasAskQty() bool {
	if o != nil && !common.IsNil(o.AskQty) {
		return true
	}

	return false
}

// SetAskQty gets a reference to the given string and assigns it to the AskQty field.
func (o *SymbolOrderBookTickerResponseInner) SetAskQty(v string) {
	o.AskQty = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *SymbolOrderBookTickerResponseInner) SetTime(v int64) {
	o.Time = &v
}

func (o SymbolOrderBookTickerResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolOrderBookTickerResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LastUpdateId) {
		toSerialize["lastUpdateId"] = o.LastUpdateId
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
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
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SymbolOrderBookTickerResponseInner) UnmarshalJSON(data []byte) (err error) {
	varSymbolOrderBookTickerResponseInner := _SymbolOrderBookTickerResponseInner{}

	err = json.Unmarshal(data, &varSymbolOrderBookTickerResponseInner)

	if err != nil {
		return err
	}

	*o = SymbolOrderBookTickerResponseInner(varSymbolOrderBookTickerResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lastUpdateId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "pair")
		delete(additionalProperties, "bidPrice")
		delete(additionalProperties, "bidQty")
		delete(additionalProperties, "askPrice")
		delete(additionalProperties, "askQty")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSymbolOrderBookTickerResponseInner struct {
	value *SymbolOrderBookTickerResponseInner
	isSet bool
}

func (v NullableSymbolOrderBookTickerResponseInner) Get() *SymbolOrderBookTickerResponseInner {
	return v.value
}

func (v *NullableSymbolOrderBookTickerResponseInner) Set(val *SymbolOrderBookTickerResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolOrderBookTickerResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolOrderBookTickerResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolOrderBookTickerResponseInner(val *SymbolOrderBookTickerResponseInner) *NullableSymbolOrderBookTickerResponseInner {
	return &NullableSymbolOrderBookTickerResponseInner{value: val, isSet: true}
}

func (v NullableSymbolOrderBookTickerResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolOrderBookTickerResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
