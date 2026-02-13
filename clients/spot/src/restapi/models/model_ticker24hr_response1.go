/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the Ticker24hrResponse1 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Ticker24hrResponse1{}

// Ticker24hrResponse1 struct for Ticker24hrResponse1
type Ticker24hrResponse1 struct {
	Symbol               *string `json:"symbol,omitempty"`
	PriceChange          *string `json:"priceChange,omitempty"`
	PriceChangePercent   *string `json:"priceChangePercent,omitempty"`
	WeightedAvgPrice     *string `json:"weightedAvgPrice,omitempty"`
	PrevClosePrice       *string `json:"prevClosePrice,omitempty"`
	LastPrice            *string `json:"lastPrice,omitempty"`
	LastQty              *string `json:"lastQty,omitempty"`
	BidPrice             *string `json:"bidPrice,omitempty"`
	BidQty               *string `json:"bidQty,omitempty"`
	AskPrice             *string `json:"askPrice,omitempty"`
	AskQty               *string `json:"askQty,omitempty"`
	OpenPrice            *string `json:"openPrice,omitempty"`
	HighPrice            *string `json:"highPrice,omitempty"`
	LowPrice             *string `json:"lowPrice,omitempty"`
	Volume               *string `json:"volume,omitempty"`
	QuoteVolume          *string `json:"quoteVolume,omitempty"`
	OpenTime             *int64  `json:"openTime,omitempty"`
	CloseTime            *int64  `json:"closeTime,omitempty"`
	FirstId              *int64  `json:"firstId,omitempty"`
	LastId               *int64  `json:"lastId,omitempty"`
	Count                *int64  `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Ticker24hrResponse1 Ticker24hrResponse1

// NewTicker24hrResponse1 instantiates a new Ticker24hrResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicker24hrResponse1() *Ticker24hrResponse1 {
	this := Ticker24hrResponse1{}
	return &this
}

// NewTicker24hrResponse1WithDefaults instantiates a new Ticker24hrResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicker24hrResponse1WithDefaults() *Ticker24hrResponse1 {
	this := Ticker24hrResponse1{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *Ticker24hrResponse1) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPriceChange returns the PriceChange field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetPriceChange() string {
	if o == nil || common.IsNil(o.PriceChange) {
		var ret string
		return ret
	}
	return *o.PriceChange
}

// GetPriceChangeOk returns a tuple with the PriceChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetPriceChangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceChange) {
		return nil, false
	}
	return o.PriceChange, true
}

// HasPriceChange returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasPriceChange() bool {
	if o != nil && !common.IsNil(o.PriceChange) {
		return true
	}

	return false
}

// SetPriceChange gets a reference to the given string and assigns it to the PriceChange field.
func (o *Ticker24hrResponse1) SetPriceChange(v string) {
	o.PriceChange = &v
}

// GetPriceChangePercent returns the PriceChangePercent field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetPriceChangePercent() string {
	if o == nil || common.IsNil(o.PriceChangePercent) {
		var ret string
		return ret
	}
	return *o.PriceChangePercent
}

// GetPriceChangePercentOk returns a tuple with the PriceChangePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetPriceChangePercentOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceChangePercent) {
		return nil, false
	}
	return o.PriceChangePercent, true
}

// HasPriceChangePercent returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasPriceChangePercent() bool {
	if o != nil && !common.IsNil(o.PriceChangePercent) {
		return true
	}

	return false
}

// SetPriceChangePercent gets a reference to the given string and assigns it to the PriceChangePercent field.
func (o *Ticker24hrResponse1) SetPriceChangePercent(v string) {
	o.PriceChangePercent = &v
}

// GetWeightedAvgPrice returns the WeightedAvgPrice field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetWeightedAvgPrice() string {
	if o == nil || common.IsNil(o.WeightedAvgPrice) {
		var ret string
		return ret
	}
	return *o.WeightedAvgPrice
}

// GetWeightedAvgPriceOk returns a tuple with the WeightedAvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetWeightedAvgPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.WeightedAvgPrice) {
		return nil, false
	}
	return o.WeightedAvgPrice, true
}

// HasWeightedAvgPrice returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasWeightedAvgPrice() bool {
	if o != nil && !common.IsNil(o.WeightedAvgPrice) {
		return true
	}

	return false
}

// SetWeightedAvgPrice gets a reference to the given string and assigns it to the WeightedAvgPrice field.
func (o *Ticker24hrResponse1) SetWeightedAvgPrice(v string) {
	o.WeightedAvgPrice = &v
}

// GetPrevClosePrice returns the PrevClosePrice field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetPrevClosePrice() string {
	if o == nil || common.IsNil(o.PrevClosePrice) {
		var ret string
		return ret
	}
	return *o.PrevClosePrice
}

// GetPrevClosePriceOk returns a tuple with the PrevClosePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetPrevClosePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.PrevClosePrice) {
		return nil, false
	}
	return o.PrevClosePrice, true
}

// HasPrevClosePrice returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasPrevClosePrice() bool {
	if o != nil && !common.IsNil(o.PrevClosePrice) {
		return true
	}

	return false
}

// SetPrevClosePrice gets a reference to the given string and assigns it to the PrevClosePrice field.
func (o *Ticker24hrResponse1) SetPrevClosePrice(v string) {
	o.PrevClosePrice = &v
}

// GetLastPrice returns the LastPrice field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetLastPrice() string {
	if o == nil || common.IsNil(o.LastPrice) {
		var ret string
		return ret
	}
	return *o.LastPrice
}

// GetLastPriceOk returns a tuple with the LastPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetLastPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastPrice) {
		return nil, false
	}
	return o.LastPrice, true
}

// HasLastPrice returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasLastPrice() bool {
	if o != nil && !common.IsNil(o.LastPrice) {
		return true
	}

	return false
}

// SetLastPrice gets a reference to the given string and assigns it to the LastPrice field.
func (o *Ticker24hrResponse1) SetLastPrice(v string) {
	o.LastPrice = &v
}

// GetLastQty returns the LastQty field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetLastQty() string {
	if o == nil || common.IsNil(o.LastQty) {
		var ret string
		return ret
	}
	return *o.LastQty
}

// GetLastQtyOk returns a tuple with the LastQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetLastQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastQty) {
		return nil, false
	}
	return o.LastQty, true
}

// HasLastQty returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasLastQty() bool {
	if o != nil && !common.IsNil(o.LastQty) {
		return true
	}

	return false
}

// SetLastQty gets a reference to the given string and assigns it to the LastQty field.
func (o *Ticker24hrResponse1) SetLastQty(v string) {
	o.LastQty = &v
}

// GetBidPrice returns the BidPrice field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetBidPrice() string {
	if o == nil || common.IsNil(o.BidPrice) {
		var ret string
		return ret
	}
	return *o.BidPrice
}

// GetBidPriceOk returns a tuple with the BidPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetBidPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidPrice) {
		return nil, false
	}
	return o.BidPrice, true
}

// HasBidPrice returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasBidPrice() bool {
	if o != nil && !common.IsNil(o.BidPrice) {
		return true
	}

	return false
}

// SetBidPrice gets a reference to the given string and assigns it to the BidPrice field.
func (o *Ticker24hrResponse1) SetBidPrice(v string) {
	o.BidPrice = &v
}

// GetBidQty returns the BidQty field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetBidQty() string {
	if o == nil || common.IsNil(o.BidQty) {
		var ret string
		return ret
	}
	return *o.BidQty
}

// GetBidQtyOk returns a tuple with the BidQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetBidQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidQty) {
		return nil, false
	}
	return o.BidQty, true
}

// HasBidQty returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasBidQty() bool {
	if o != nil && !common.IsNil(o.BidQty) {
		return true
	}

	return false
}

// SetBidQty gets a reference to the given string and assigns it to the BidQty field.
func (o *Ticker24hrResponse1) SetBidQty(v string) {
	o.BidQty = &v
}

// GetAskPrice returns the AskPrice field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetAskPrice() string {
	if o == nil || common.IsNil(o.AskPrice) {
		var ret string
		return ret
	}
	return *o.AskPrice
}

// GetAskPriceOk returns a tuple with the AskPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetAskPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskPrice) {
		return nil, false
	}
	return o.AskPrice, true
}

// HasAskPrice returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasAskPrice() bool {
	if o != nil && !common.IsNil(o.AskPrice) {
		return true
	}

	return false
}

// SetAskPrice gets a reference to the given string and assigns it to the AskPrice field.
func (o *Ticker24hrResponse1) SetAskPrice(v string) {
	o.AskPrice = &v
}

// GetAskQty returns the AskQty field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetAskQty() string {
	if o == nil || common.IsNil(o.AskQty) {
		var ret string
		return ret
	}
	return *o.AskQty
}

// GetAskQtyOk returns a tuple with the AskQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetAskQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskQty) {
		return nil, false
	}
	return o.AskQty, true
}

// HasAskQty returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasAskQty() bool {
	if o != nil && !common.IsNil(o.AskQty) {
		return true
	}

	return false
}

// SetAskQty gets a reference to the given string and assigns it to the AskQty field.
func (o *Ticker24hrResponse1) SetAskQty(v string) {
	o.AskQty = &v
}

// GetOpenPrice returns the OpenPrice field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetOpenPrice() string {
	if o == nil || common.IsNil(o.OpenPrice) {
		var ret string
		return ret
	}
	return *o.OpenPrice
}

// GetOpenPriceOk returns a tuple with the OpenPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetOpenPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpenPrice) {
		return nil, false
	}
	return o.OpenPrice, true
}

// HasOpenPrice returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasOpenPrice() bool {
	if o != nil && !common.IsNil(o.OpenPrice) {
		return true
	}

	return false
}

// SetOpenPrice gets a reference to the given string and assigns it to the OpenPrice field.
func (o *Ticker24hrResponse1) SetOpenPrice(v string) {
	o.OpenPrice = &v
}

// GetHighPrice returns the HighPrice field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetHighPrice() string {
	if o == nil || common.IsNil(o.HighPrice) {
		var ret string
		return ret
	}
	return *o.HighPrice
}

// GetHighPriceOk returns a tuple with the HighPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetHighPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.HighPrice) {
		return nil, false
	}
	return o.HighPrice, true
}

// HasHighPrice returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasHighPrice() bool {
	if o != nil && !common.IsNil(o.HighPrice) {
		return true
	}

	return false
}

// SetHighPrice gets a reference to the given string and assigns it to the HighPrice field.
func (o *Ticker24hrResponse1) SetHighPrice(v string) {
	o.HighPrice = &v
}

// GetLowPrice returns the LowPrice field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetLowPrice() string {
	if o == nil || common.IsNil(o.LowPrice) {
		var ret string
		return ret
	}
	return *o.LowPrice
}

// GetLowPriceOk returns a tuple with the LowPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetLowPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LowPrice) {
		return nil, false
	}
	return o.LowPrice, true
}

// HasLowPrice returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasLowPrice() bool {
	if o != nil && !common.IsNil(o.LowPrice) {
		return true
	}

	return false
}

// SetLowPrice gets a reference to the given string and assigns it to the LowPrice field.
func (o *Ticker24hrResponse1) SetLowPrice(v string) {
	o.LowPrice = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetVolume() string {
	if o == nil || common.IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetVolumeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasVolume() bool {
	if o != nil && !common.IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *Ticker24hrResponse1) SetVolume(v string) {
	o.Volume = &v
}

// GetQuoteVolume returns the QuoteVolume field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetQuoteVolume() string {
	if o == nil || common.IsNil(o.QuoteVolume) {
		var ret string
		return ret
	}
	return *o.QuoteVolume
}

// GetQuoteVolumeOk returns a tuple with the QuoteVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetQuoteVolumeOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteVolume) {
		return nil, false
	}
	return o.QuoteVolume, true
}

// HasQuoteVolume returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasQuoteVolume() bool {
	if o != nil && !common.IsNil(o.QuoteVolume) {
		return true
	}

	return false
}

// SetQuoteVolume gets a reference to the given string and assigns it to the QuoteVolume field.
func (o *Ticker24hrResponse1) SetQuoteVolume(v string) {
	o.QuoteVolume = &v
}

// GetOpenTime returns the OpenTime field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetOpenTime() int64 {
	if o == nil || common.IsNil(o.OpenTime) {
		var ret int64
		return ret
	}
	return *o.OpenTime
}

// GetOpenTimeOk returns a tuple with the OpenTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetOpenTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OpenTime) {
		return nil, false
	}
	return o.OpenTime, true
}

// HasOpenTime returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasOpenTime() bool {
	if o != nil && !common.IsNil(o.OpenTime) {
		return true
	}

	return false
}

// SetOpenTime gets a reference to the given int64 and assigns it to the OpenTime field.
func (o *Ticker24hrResponse1) SetOpenTime(v int64) {
	o.OpenTime = &v
}

// GetCloseTime returns the CloseTime field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetCloseTime() int64 {
	if o == nil || common.IsNil(o.CloseTime) {
		var ret int64
		return ret
	}
	return *o.CloseTime
}

// GetCloseTimeOk returns a tuple with the CloseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetCloseTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CloseTime) {
		return nil, false
	}
	return o.CloseTime, true
}

// HasCloseTime returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasCloseTime() bool {
	if o != nil && !common.IsNil(o.CloseTime) {
		return true
	}

	return false
}

// SetCloseTime gets a reference to the given int64 and assigns it to the CloseTime field.
func (o *Ticker24hrResponse1) SetCloseTime(v int64) {
	o.CloseTime = &v
}

// GetFirstId returns the FirstId field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetFirstId() int64 {
	if o == nil || common.IsNil(o.FirstId) {
		var ret int64
		return ret
	}
	return *o.FirstId
}

// GetFirstIdOk returns a tuple with the FirstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetFirstIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FirstId) {
		return nil, false
	}
	return o.FirstId, true
}

// HasFirstId returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasFirstId() bool {
	if o != nil && !common.IsNil(o.FirstId) {
		return true
	}

	return false
}

// SetFirstId gets a reference to the given int64 and assigns it to the FirstId field.
func (o *Ticker24hrResponse1) SetFirstId(v int64) {
	o.FirstId = &v
}

// GetLastId returns the LastId field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetLastId() int64 {
	if o == nil || common.IsNil(o.LastId) {
		var ret int64
		return ret
	}
	return *o.LastId
}

// GetLastIdOk returns a tuple with the LastId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetLastIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastId) {
		return nil, false
	}
	return o.LastId, true
}

// HasLastId returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasLastId() bool {
	if o != nil && !common.IsNil(o.LastId) {
		return true
	}

	return false
}

// SetLastId gets a reference to the given int64 and assigns it to the LastId field.
func (o *Ticker24hrResponse1) SetLastId(v int64) {
	o.LastId = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Ticker24hrResponse1) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse1) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Ticker24hrResponse1) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *Ticker24hrResponse1) SetCount(v int64) {
	o.Count = &v
}

func (o Ticker24hrResponse1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ticker24hrResponse1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.PriceChange) {
		toSerialize["priceChange"] = o.PriceChange
	}
	if !common.IsNil(o.PriceChangePercent) {
		toSerialize["priceChangePercent"] = o.PriceChangePercent
	}
	if !common.IsNil(o.WeightedAvgPrice) {
		toSerialize["weightedAvgPrice"] = o.WeightedAvgPrice
	}
	if !common.IsNil(o.PrevClosePrice) {
		toSerialize["prevClosePrice"] = o.PrevClosePrice
	}
	if !common.IsNil(o.LastPrice) {
		toSerialize["lastPrice"] = o.LastPrice
	}
	if !common.IsNil(o.LastQty) {
		toSerialize["lastQty"] = o.LastQty
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
	if !common.IsNil(o.OpenPrice) {
		toSerialize["openPrice"] = o.OpenPrice
	}
	if !common.IsNil(o.HighPrice) {
		toSerialize["highPrice"] = o.HighPrice
	}
	if !common.IsNil(o.LowPrice) {
		toSerialize["lowPrice"] = o.LowPrice
	}
	if !common.IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	if !common.IsNil(o.QuoteVolume) {
		toSerialize["quoteVolume"] = o.QuoteVolume
	}
	if !common.IsNil(o.OpenTime) {
		toSerialize["openTime"] = o.OpenTime
	}
	if !common.IsNil(o.CloseTime) {
		toSerialize["closeTime"] = o.CloseTime
	}
	if !common.IsNil(o.FirstId) {
		toSerialize["firstId"] = o.FirstId
	}
	if !common.IsNil(o.LastId) {
		toSerialize["lastId"] = o.LastId
	}
	if !common.IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Ticker24hrResponse1) UnmarshalJSON(data []byte) (err error) {
	varTicker24hrResponse1 := _Ticker24hrResponse1{}

	err = json.Unmarshal(data, &varTicker24hrResponse1)

	if err != nil {
		return err
	}

	*o = Ticker24hrResponse1(varTicker24hrResponse1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "priceChange")
		delete(additionalProperties, "priceChangePercent")
		delete(additionalProperties, "weightedAvgPrice")
		delete(additionalProperties, "prevClosePrice")
		delete(additionalProperties, "lastPrice")
		delete(additionalProperties, "lastQty")
		delete(additionalProperties, "bidPrice")
		delete(additionalProperties, "bidQty")
		delete(additionalProperties, "askPrice")
		delete(additionalProperties, "askQty")
		delete(additionalProperties, "openPrice")
		delete(additionalProperties, "highPrice")
		delete(additionalProperties, "lowPrice")
		delete(additionalProperties, "volume")
		delete(additionalProperties, "quoteVolume")
		delete(additionalProperties, "openTime")
		delete(additionalProperties, "closeTime")
		delete(additionalProperties, "firstId")
		delete(additionalProperties, "lastId")
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTicker24hrResponse1 struct {
	value *Ticker24hrResponse1
	isSet bool
}

func (v NullableTicker24hrResponse1) Get() *Ticker24hrResponse1 {
	return v.value
}

func (v *NullableTicker24hrResponse1) Set(val *Ticker24hrResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableTicker24hrResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableTicker24hrResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicker24hrResponse1(val *Ticker24hrResponse1) *NullableTicker24hrResponse1 {
	return &NullableTicker24hrResponse1{value: val, isSet: true}
}

func (v NullableTicker24hrResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicker24hrResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
