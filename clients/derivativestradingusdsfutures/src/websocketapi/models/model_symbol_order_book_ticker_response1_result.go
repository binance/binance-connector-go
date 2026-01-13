/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SymbolOrderBookTickerResponse1Result type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolOrderBookTickerResponse1Result{}

// SymbolOrderBookTickerResponse1Result struct for SymbolOrderBookTickerResponse1Result
type SymbolOrderBookTickerResponse1Result struct {
	LastUpdateId         *int64  `json:"lastUpdateId,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	BidPrice             *string `json:"bidPrice,omitempty"`
	BidQty               *string `json:"bidQty,omitempty"`
	AskPrice             *string `json:"askPrice,omitempty"`
	AskQty               *string `json:"askQty,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SymbolOrderBookTickerResponse1Result SymbolOrderBookTickerResponse1Result

// NewSymbolOrderBookTickerResponse1Result instantiates a new SymbolOrderBookTickerResponse1Result object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolOrderBookTickerResponse1Result() *SymbolOrderBookTickerResponse1Result {
	this := SymbolOrderBookTickerResponse1Result{}
	return &this
}

// NewSymbolOrderBookTickerResponse1ResultWithDefaults instantiates a new SymbolOrderBookTickerResponse1Result object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolOrderBookTickerResponse1ResultWithDefaults() *SymbolOrderBookTickerResponse1Result {
	this := SymbolOrderBookTickerResponse1Result{}
	return &this
}

// GetLastUpdateId returns the LastUpdateId field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse1Result) GetLastUpdateId() int64 {
	if o == nil || common.IsNil(o.LastUpdateId) {
		var ret int64
		return ret
	}
	return *o.LastUpdateId
}

// GetLastUpdateIdOk returns a tuple with the LastUpdateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse1Result) GetLastUpdateIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastUpdateId) {
		return nil, false
	}
	return o.LastUpdateId, true
}

// HasLastUpdateId returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse1Result) HasLastUpdateId() bool {
	if o != nil && !common.IsNil(o.LastUpdateId) {
		return true
	}

	return false
}

// SetLastUpdateId gets a reference to the given int64 and assigns it to the LastUpdateId field.
func (o *SymbolOrderBookTickerResponse1Result) SetLastUpdateId(v int64) {
	o.LastUpdateId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse1Result) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse1Result) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse1Result) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SymbolOrderBookTickerResponse1Result) SetSymbol(v string) {
	o.Symbol = &v
}

// GetBidPrice returns the BidPrice field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse1Result) GetBidPrice() string {
	if o == nil || common.IsNil(o.BidPrice) {
		var ret string
		return ret
	}
	return *o.BidPrice
}

// GetBidPriceOk returns a tuple with the BidPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse1Result) GetBidPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidPrice) {
		return nil, false
	}
	return o.BidPrice, true
}

// HasBidPrice returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse1Result) HasBidPrice() bool {
	if o != nil && !common.IsNil(o.BidPrice) {
		return true
	}

	return false
}

// SetBidPrice gets a reference to the given string and assigns it to the BidPrice field.
func (o *SymbolOrderBookTickerResponse1Result) SetBidPrice(v string) {
	o.BidPrice = &v
}

// GetBidQty returns the BidQty field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse1Result) GetBidQty() string {
	if o == nil || common.IsNil(o.BidQty) {
		var ret string
		return ret
	}
	return *o.BidQty
}

// GetBidQtyOk returns a tuple with the BidQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse1Result) GetBidQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidQty) {
		return nil, false
	}
	return o.BidQty, true
}

// HasBidQty returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse1Result) HasBidQty() bool {
	if o != nil && !common.IsNil(o.BidQty) {
		return true
	}

	return false
}

// SetBidQty gets a reference to the given string and assigns it to the BidQty field.
func (o *SymbolOrderBookTickerResponse1Result) SetBidQty(v string) {
	o.BidQty = &v
}

// GetAskPrice returns the AskPrice field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse1Result) GetAskPrice() string {
	if o == nil || common.IsNil(o.AskPrice) {
		var ret string
		return ret
	}
	return *o.AskPrice
}

// GetAskPriceOk returns a tuple with the AskPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse1Result) GetAskPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskPrice) {
		return nil, false
	}
	return o.AskPrice, true
}

// HasAskPrice returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse1Result) HasAskPrice() bool {
	if o != nil && !common.IsNil(o.AskPrice) {
		return true
	}

	return false
}

// SetAskPrice gets a reference to the given string and assigns it to the AskPrice field.
func (o *SymbolOrderBookTickerResponse1Result) SetAskPrice(v string) {
	o.AskPrice = &v
}

// GetAskQty returns the AskQty field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse1Result) GetAskQty() string {
	if o == nil || common.IsNil(o.AskQty) {
		var ret string
		return ret
	}
	return *o.AskQty
}

// GetAskQtyOk returns a tuple with the AskQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse1Result) GetAskQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskQty) {
		return nil, false
	}
	return o.AskQty, true
}

// HasAskQty returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse1Result) HasAskQty() bool {
	if o != nil && !common.IsNil(o.AskQty) {
		return true
	}

	return false
}

// SetAskQty gets a reference to the given string and assigns it to the AskQty field.
func (o *SymbolOrderBookTickerResponse1Result) SetAskQty(v string) {
	o.AskQty = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse1Result) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse1Result) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse1Result) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *SymbolOrderBookTickerResponse1Result) SetTime(v int64) {
	o.Time = &v
}

func (o SymbolOrderBookTickerResponse1Result) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolOrderBookTickerResponse1Result) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LastUpdateId) {
		toSerialize["lastUpdateId"] = o.LastUpdateId
	}
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
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SymbolOrderBookTickerResponse1Result) UnmarshalJSON(data []byte) (err error) {
	varSymbolOrderBookTickerResponse1Result := _SymbolOrderBookTickerResponse1Result{}

	err = json.Unmarshal(data, &varSymbolOrderBookTickerResponse1Result)

	if err != nil {
		return err
	}

	*o = SymbolOrderBookTickerResponse1Result(varSymbolOrderBookTickerResponse1Result)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lastUpdateId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "bidPrice")
		delete(additionalProperties, "bidQty")
		delete(additionalProperties, "askPrice")
		delete(additionalProperties, "askQty")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSymbolOrderBookTickerResponse1Result struct {
	value *SymbolOrderBookTickerResponse1Result
	isSet bool
}

func (v NullableSymbolOrderBookTickerResponse1Result) Get() *SymbolOrderBookTickerResponse1Result {
	return v.value
}

func (v *NullableSymbolOrderBookTickerResponse1Result) Set(val *SymbolOrderBookTickerResponse1Result) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolOrderBookTickerResponse1Result) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolOrderBookTickerResponse1Result) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolOrderBookTickerResponse1Result(val *SymbolOrderBookTickerResponse1Result) *NullableSymbolOrderBookTickerResponse1Result {
	return &NullableSymbolOrderBookTickerResponse1Result{value: val, isSet: true}
}

func (v NullableSymbolOrderBookTickerResponse1Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolOrderBookTickerResponse1Result) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
