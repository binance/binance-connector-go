/*
Futures (USDⓈ-M) WebSocket API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the SymbolOrderBookTickerResponse2ResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolOrderBookTickerResponse2ResultInner{}

// SymbolOrderBookTickerResponse2ResultInner struct for SymbolOrderBookTickerResponse2ResultInner
type SymbolOrderBookTickerResponse2ResultInner struct {
	LastUpdateId *int64  `json:"lastUpdateId,omitempty"`
	Symbol       *string `json:"symbol,omitempty"`
	BidPrice     *string `json:"bidPrice,omitempty"`
	BidQty       *string `json:"bidQty,omitempty"`
	AskPrice     *string `json:"askPrice,omitempty"`
	AskQty       *string `json:"askQty,omitempty"`
	// Transaction time
	Time                 *int64 `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SymbolOrderBookTickerResponse2ResultInner SymbolOrderBookTickerResponse2ResultInner

// NewSymbolOrderBookTickerResponse2ResultInner instantiates a new SymbolOrderBookTickerResponse2ResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolOrderBookTickerResponse2ResultInner() *SymbolOrderBookTickerResponse2ResultInner {
	this := SymbolOrderBookTickerResponse2ResultInner{}
	return &this
}

// NewSymbolOrderBookTickerResponse2ResultInnerWithDefaults instantiates a new SymbolOrderBookTickerResponse2ResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolOrderBookTickerResponse2ResultInnerWithDefaults() *SymbolOrderBookTickerResponse2ResultInner {
	this := SymbolOrderBookTickerResponse2ResultInner{}
	return &this
}

// GetLastUpdateId returns the LastUpdateId field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse2ResultInner) GetLastUpdateId() int64 {
	if o == nil || common.IsNil(o.LastUpdateId) {
		var ret int64
		return ret
	}
	return *o.LastUpdateId
}

// GetLastUpdateIdOk returns a tuple with the LastUpdateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse2ResultInner) GetLastUpdateIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastUpdateId) {
		return nil, false
	}
	return o.LastUpdateId, true
}

// HasLastUpdateId returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse2ResultInner) HasLastUpdateId() bool {
	if o != nil && !common.IsNil(o.LastUpdateId) {
		return true
	}

	return false
}

// SetLastUpdateId gets a reference to the given int64 and assigns it to the LastUpdateId field.
func (o *SymbolOrderBookTickerResponse2ResultInner) SetLastUpdateId(v int64) {
	o.LastUpdateId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse2ResultInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse2ResultInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse2ResultInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SymbolOrderBookTickerResponse2ResultInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetBidPrice returns the BidPrice field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse2ResultInner) GetBidPrice() string {
	if o == nil || common.IsNil(o.BidPrice) {
		var ret string
		return ret
	}
	return *o.BidPrice
}

// GetBidPriceOk returns a tuple with the BidPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse2ResultInner) GetBidPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidPrice) {
		return nil, false
	}
	return o.BidPrice, true
}

// HasBidPrice returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse2ResultInner) HasBidPrice() bool {
	if o != nil && !common.IsNil(o.BidPrice) {
		return true
	}

	return false
}

// SetBidPrice gets a reference to the given string and assigns it to the BidPrice field.
func (o *SymbolOrderBookTickerResponse2ResultInner) SetBidPrice(v string) {
	o.BidPrice = &v
}

// GetBidQty returns the BidQty field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse2ResultInner) GetBidQty() string {
	if o == nil || common.IsNil(o.BidQty) {
		var ret string
		return ret
	}
	return *o.BidQty
}

// GetBidQtyOk returns a tuple with the BidQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse2ResultInner) GetBidQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidQty) {
		return nil, false
	}
	return o.BidQty, true
}

// HasBidQty returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse2ResultInner) HasBidQty() bool {
	if o != nil && !common.IsNil(o.BidQty) {
		return true
	}

	return false
}

// SetBidQty gets a reference to the given string and assigns it to the BidQty field.
func (o *SymbolOrderBookTickerResponse2ResultInner) SetBidQty(v string) {
	o.BidQty = &v
}

// GetAskPrice returns the AskPrice field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse2ResultInner) GetAskPrice() string {
	if o == nil || common.IsNil(o.AskPrice) {
		var ret string
		return ret
	}
	return *o.AskPrice
}

// GetAskPriceOk returns a tuple with the AskPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse2ResultInner) GetAskPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskPrice) {
		return nil, false
	}
	return o.AskPrice, true
}

// HasAskPrice returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse2ResultInner) HasAskPrice() bool {
	if o != nil && !common.IsNil(o.AskPrice) {
		return true
	}

	return false
}

// SetAskPrice gets a reference to the given string and assigns it to the AskPrice field.
func (o *SymbolOrderBookTickerResponse2ResultInner) SetAskPrice(v string) {
	o.AskPrice = &v
}

// GetAskQty returns the AskQty field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse2ResultInner) GetAskQty() string {
	if o == nil || common.IsNil(o.AskQty) {
		var ret string
		return ret
	}
	return *o.AskQty
}

// GetAskQtyOk returns a tuple with the AskQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse2ResultInner) GetAskQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskQty) {
		return nil, false
	}
	return o.AskQty, true
}

// HasAskQty returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse2ResultInner) HasAskQty() bool {
	if o != nil && !common.IsNil(o.AskQty) {
		return true
	}

	return false
}

// SetAskQty gets a reference to the given string and assigns it to the AskQty field.
func (o *SymbolOrderBookTickerResponse2ResultInner) SetAskQty(v string) {
	o.AskQty = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse2ResultInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse2ResultInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse2ResultInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *SymbolOrderBookTickerResponse2ResultInner) SetTime(v int64) {
	o.Time = &v
}

func (o SymbolOrderBookTickerResponse2ResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolOrderBookTickerResponse2ResultInner) ToMap() (map[string]interface{}, error) {
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

func (o *SymbolOrderBookTickerResponse2ResultInner) UnmarshalJSON(data []byte) (err error) {
	varSymbolOrderBookTickerResponse2ResultInner := _SymbolOrderBookTickerResponse2ResultInner{}

	err = json.Unmarshal(data, &varSymbolOrderBookTickerResponse2ResultInner)

	if err != nil {
		return err
	}

	*o = SymbolOrderBookTickerResponse2ResultInner(varSymbolOrderBookTickerResponse2ResultInner)

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

type NullableSymbolOrderBookTickerResponse2ResultInner struct {
	value *SymbolOrderBookTickerResponse2ResultInner
	isSet bool
}

func (v NullableSymbolOrderBookTickerResponse2ResultInner) Get() *SymbolOrderBookTickerResponse2ResultInner {
	return v.value
}

func (v *NullableSymbolOrderBookTickerResponse2ResultInner) Set(val *SymbolOrderBookTickerResponse2ResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolOrderBookTickerResponse2ResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolOrderBookTickerResponse2ResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolOrderBookTickerResponse2ResultInner(val *SymbolOrderBookTickerResponse2ResultInner) *NullableSymbolOrderBookTickerResponse2ResultInner {
	return &NullableSymbolOrderBookTickerResponse2ResultInner{value: val, isSet: true}
}

func (v NullableSymbolOrderBookTickerResponse2ResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolOrderBookTickerResponse2ResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
