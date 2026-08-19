/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the LatestQuoteResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &LatestQuoteResponse{}

// LatestQuoteResponse struct for LatestQuoteResponse
type LatestQuoteResponse struct {
	// US-equity ticker.
	Symbol *string `json:"symbol,omitempty"`
	// Best bid price (USD).
	BidPrice *string `json:"bidPrice,omitempty"`
	// Best ask price (USD).
	AskPrice *string `json:"askPrice,omitempty"`
	// Best bid size (shares).
	BidSize *int32 `json:"bidSize,omitempty"`
	// Best ask size (shares).
	AskSize              *int32 `json:"askSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LatestQuoteResponse LatestQuoteResponse

// NewLatestQuoteResponse instantiates a new LatestQuoteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLatestQuoteResponse() *LatestQuoteResponse {
	this := LatestQuoteResponse{}
	return &this
}

// NewLatestQuoteResponseWithDefaults instantiates a new LatestQuoteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLatestQuoteResponseWithDefaults() *LatestQuoteResponse {
	this := LatestQuoteResponse{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *LatestQuoteResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatestQuoteResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *LatestQuoteResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *LatestQuoteResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetBidPrice returns the BidPrice field value if set, zero value otherwise.
func (o *LatestQuoteResponse) GetBidPrice() string {
	if o == nil || common.IsNil(o.BidPrice) {
		var ret string
		return ret
	}
	return *o.BidPrice
}

// GetBidPriceOk returns a tuple with the BidPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatestQuoteResponse) GetBidPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidPrice) {
		return nil, false
	}
	return o.BidPrice, true
}

// HasBidPrice returns a boolean if a field has been set.
func (o *LatestQuoteResponse) HasBidPrice() bool {
	if o != nil && !common.IsNil(o.BidPrice) {
		return true
	}

	return false
}

// SetBidPrice gets a reference to the given string and assigns it to the BidPrice field.
func (o *LatestQuoteResponse) SetBidPrice(v string) {
	o.BidPrice = &v
}

// GetAskPrice returns the AskPrice field value if set, zero value otherwise.
func (o *LatestQuoteResponse) GetAskPrice() string {
	if o == nil || common.IsNil(o.AskPrice) {
		var ret string
		return ret
	}
	return *o.AskPrice
}

// GetAskPriceOk returns a tuple with the AskPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatestQuoteResponse) GetAskPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskPrice) {
		return nil, false
	}
	return o.AskPrice, true
}

// HasAskPrice returns a boolean if a field has been set.
func (o *LatestQuoteResponse) HasAskPrice() bool {
	if o != nil && !common.IsNil(o.AskPrice) {
		return true
	}

	return false
}

// SetAskPrice gets a reference to the given string and assigns it to the AskPrice field.
func (o *LatestQuoteResponse) SetAskPrice(v string) {
	o.AskPrice = &v
}

// GetBidSize returns the BidSize field value if set, zero value otherwise.
func (o *LatestQuoteResponse) GetBidSize() int32 {
	if o == nil || common.IsNil(o.BidSize) {
		var ret int32
		return ret
	}
	return *o.BidSize
}

// GetBidSizeOk returns a tuple with the BidSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatestQuoteResponse) GetBidSizeOk() (*int32, bool) {
	if o == nil || common.IsNil(o.BidSize) {
		return nil, false
	}
	return o.BidSize, true
}

// HasBidSize returns a boolean if a field has been set.
func (o *LatestQuoteResponse) HasBidSize() bool {
	if o != nil && !common.IsNil(o.BidSize) {
		return true
	}

	return false
}

// SetBidSize gets a reference to the given int32 and assigns it to the BidSize field.
func (o *LatestQuoteResponse) SetBidSize(v int32) {
	o.BidSize = &v
}

// GetAskSize returns the AskSize field value if set, zero value otherwise.
func (o *LatestQuoteResponse) GetAskSize() int32 {
	if o == nil || common.IsNil(o.AskSize) {
		var ret int32
		return ret
	}
	return *o.AskSize
}

// GetAskSizeOk returns a tuple with the AskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatestQuoteResponse) GetAskSizeOk() (*int32, bool) {
	if o == nil || common.IsNil(o.AskSize) {
		return nil, false
	}
	return o.AskSize, true
}

// HasAskSize returns a boolean if a field has been set.
func (o *LatestQuoteResponse) HasAskSize() bool {
	if o != nil && !common.IsNil(o.AskSize) {
		return true
	}

	return false
}

// SetAskSize gets a reference to the given int32 and assigns it to the AskSize field.
func (o *LatestQuoteResponse) SetAskSize(v int32) {
	o.AskSize = &v
}

func (o LatestQuoteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LatestQuoteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.BidPrice) {
		toSerialize["bidPrice"] = o.BidPrice
	}
	if !common.IsNil(o.AskPrice) {
		toSerialize["askPrice"] = o.AskPrice
	}
	if !common.IsNil(o.BidSize) {
		toSerialize["bidSize"] = o.BidSize
	}
	if !common.IsNil(o.AskSize) {
		toSerialize["askSize"] = o.AskSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LatestQuoteResponse) UnmarshalJSON(data []byte) (err error) {
	varLatestQuoteResponse := _LatestQuoteResponse{}

	err = json.Unmarshal(data, &varLatestQuoteResponse)

	if err != nil {
		return err
	}

	*o = LatestQuoteResponse(varLatestQuoteResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "bidPrice")
		delete(additionalProperties, "askPrice")
		delete(additionalProperties, "bidSize")
		delete(additionalProperties, "askSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLatestQuoteResponse struct {
	value *LatestQuoteResponse
	isSet bool
}

func (v NullableLatestQuoteResponse) Get() *LatestQuoteResponse {
	return v.value
}

func (v *NullableLatestQuoteResponse) Set(val *LatestQuoteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLatestQuoteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLatestQuoteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLatestQuoteResponse(val *LatestQuoteResponse) *NullableLatestQuoteResponse {
	return &NullableLatestQuoteResponse{value: val, isSet: true}
}

func (v NullableLatestQuoteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLatestQuoteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
