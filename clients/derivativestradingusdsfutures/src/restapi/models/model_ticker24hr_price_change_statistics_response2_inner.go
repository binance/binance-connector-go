/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the Ticker24hrPriceChangeStatisticsResponse2Inner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Ticker24hrPriceChangeStatisticsResponse2Inner{}

// Ticker24hrPriceChangeStatisticsResponse2Inner struct for Ticker24hrPriceChangeStatisticsResponse2Inner
type Ticker24hrPriceChangeStatisticsResponse2Inner struct {
	Symbol               *string `json:"symbol,omitempty"`
	PriceChange          *string `json:"priceChange,omitempty"`
	PriceChangePercent   *string `json:"priceChangePercent,omitempty"`
	WeightedAvgPrice     *string `json:"weightedAvgPrice,omitempty"`
	LastPrice            *string `json:"lastPrice,omitempty"`
	LastQty              *string `json:"lastQty,omitempty"`
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

type _Ticker24hrPriceChangeStatisticsResponse2Inner Ticker24hrPriceChangeStatisticsResponse2Inner

// NewTicker24hrPriceChangeStatisticsResponse2Inner instantiates a new Ticker24hrPriceChangeStatisticsResponse2Inner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicker24hrPriceChangeStatisticsResponse2Inner() *Ticker24hrPriceChangeStatisticsResponse2Inner {
	this := Ticker24hrPriceChangeStatisticsResponse2Inner{}
	return &this
}

// NewTicker24hrPriceChangeStatisticsResponse2InnerWithDefaults instantiates a new Ticker24hrPriceChangeStatisticsResponse2Inner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicker24hrPriceChangeStatisticsResponse2InnerWithDefaults() *Ticker24hrPriceChangeStatisticsResponse2Inner {
	this := Ticker24hrPriceChangeStatisticsResponse2Inner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPriceChange returns the PriceChange field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetPriceChange() string {
	if o == nil || common.IsNil(o.PriceChange) {
		var ret string
		return ret
	}
	return *o.PriceChange
}

// GetPriceChangeOk returns a tuple with the PriceChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetPriceChangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceChange) {
		return nil, false
	}
	return o.PriceChange, true
}

// HasPriceChange returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) HasPriceChange() bool {
	if o != nil && !common.IsNil(o.PriceChange) {
		return true
	}

	return false
}

// SetPriceChange gets a reference to the given string and assigns it to the PriceChange field.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) SetPriceChange(v string) {
	o.PriceChange = &v
}

// GetPriceChangePercent returns the PriceChangePercent field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetPriceChangePercent() string {
	if o == nil || common.IsNil(o.PriceChangePercent) {
		var ret string
		return ret
	}
	return *o.PriceChangePercent
}

// GetPriceChangePercentOk returns a tuple with the PriceChangePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetPriceChangePercentOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceChangePercent) {
		return nil, false
	}
	return o.PriceChangePercent, true
}

// HasPriceChangePercent returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) HasPriceChangePercent() bool {
	if o != nil && !common.IsNil(o.PriceChangePercent) {
		return true
	}

	return false
}

// SetPriceChangePercent gets a reference to the given string and assigns it to the PriceChangePercent field.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) SetPriceChangePercent(v string) {
	o.PriceChangePercent = &v
}

// GetWeightedAvgPrice returns the WeightedAvgPrice field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetWeightedAvgPrice() string {
	if o == nil || common.IsNil(o.WeightedAvgPrice) {
		var ret string
		return ret
	}
	return *o.WeightedAvgPrice
}

// GetWeightedAvgPriceOk returns a tuple with the WeightedAvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetWeightedAvgPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.WeightedAvgPrice) {
		return nil, false
	}
	return o.WeightedAvgPrice, true
}

// HasWeightedAvgPrice returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) HasWeightedAvgPrice() bool {
	if o != nil && !common.IsNil(o.WeightedAvgPrice) {
		return true
	}

	return false
}

// SetWeightedAvgPrice gets a reference to the given string and assigns it to the WeightedAvgPrice field.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) SetWeightedAvgPrice(v string) {
	o.WeightedAvgPrice = &v
}

// GetLastPrice returns the LastPrice field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetLastPrice() string {
	if o == nil || common.IsNil(o.LastPrice) {
		var ret string
		return ret
	}
	return *o.LastPrice
}

// GetLastPriceOk returns a tuple with the LastPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetLastPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastPrice) {
		return nil, false
	}
	return o.LastPrice, true
}

// HasLastPrice returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) HasLastPrice() bool {
	if o != nil && !common.IsNil(o.LastPrice) {
		return true
	}

	return false
}

// SetLastPrice gets a reference to the given string and assigns it to the LastPrice field.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) SetLastPrice(v string) {
	o.LastPrice = &v
}

// GetLastQty returns the LastQty field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetLastQty() string {
	if o == nil || common.IsNil(o.LastQty) {
		var ret string
		return ret
	}
	return *o.LastQty
}

// GetLastQtyOk returns a tuple with the LastQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetLastQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastQty) {
		return nil, false
	}
	return o.LastQty, true
}

// HasLastQty returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) HasLastQty() bool {
	if o != nil && !common.IsNil(o.LastQty) {
		return true
	}

	return false
}

// SetLastQty gets a reference to the given string and assigns it to the LastQty field.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) SetLastQty(v string) {
	o.LastQty = &v
}

// GetOpenPrice returns the OpenPrice field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetOpenPrice() string {
	if o == nil || common.IsNil(o.OpenPrice) {
		var ret string
		return ret
	}
	return *o.OpenPrice
}

// GetOpenPriceOk returns a tuple with the OpenPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetOpenPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpenPrice) {
		return nil, false
	}
	return o.OpenPrice, true
}

// HasOpenPrice returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) HasOpenPrice() bool {
	if o != nil && !common.IsNil(o.OpenPrice) {
		return true
	}

	return false
}

// SetOpenPrice gets a reference to the given string and assigns it to the OpenPrice field.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) SetOpenPrice(v string) {
	o.OpenPrice = &v
}

// GetHighPrice returns the HighPrice field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetHighPrice() string {
	if o == nil || common.IsNil(o.HighPrice) {
		var ret string
		return ret
	}
	return *o.HighPrice
}

// GetHighPriceOk returns a tuple with the HighPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetHighPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.HighPrice) {
		return nil, false
	}
	return o.HighPrice, true
}

// HasHighPrice returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) HasHighPrice() bool {
	if o != nil && !common.IsNil(o.HighPrice) {
		return true
	}

	return false
}

// SetHighPrice gets a reference to the given string and assigns it to the HighPrice field.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) SetHighPrice(v string) {
	o.HighPrice = &v
}

// GetLowPrice returns the LowPrice field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetLowPrice() string {
	if o == nil || common.IsNil(o.LowPrice) {
		var ret string
		return ret
	}
	return *o.LowPrice
}

// GetLowPriceOk returns a tuple with the LowPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetLowPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LowPrice) {
		return nil, false
	}
	return o.LowPrice, true
}

// HasLowPrice returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) HasLowPrice() bool {
	if o != nil && !common.IsNil(o.LowPrice) {
		return true
	}

	return false
}

// SetLowPrice gets a reference to the given string and assigns it to the LowPrice field.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) SetLowPrice(v string) {
	o.LowPrice = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetVolume() string {
	if o == nil || common.IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetVolumeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) HasVolume() bool {
	if o != nil && !common.IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) SetVolume(v string) {
	o.Volume = &v
}

// GetQuoteVolume returns the QuoteVolume field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetQuoteVolume() string {
	if o == nil || common.IsNil(o.QuoteVolume) {
		var ret string
		return ret
	}
	return *o.QuoteVolume
}

// GetQuoteVolumeOk returns a tuple with the QuoteVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetQuoteVolumeOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteVolume) {
		return nil, false
	}
	return o.QuoteVolume, true
}

// HasQuoteVolume returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) HasQuoteVolume() bool {
	if o != nil && !common.IsNil(o.QuoteVolume) {
		return true
	}

	return false
}

// SetQuoteVolume gets a reference to the given string and assigns it to the QuoteVolume field.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) SetQuoteVolume(v string) {
	o.QuoteVolume = &v
}

// GetOpenTime returns the OpenTime field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetOpenTime() int64 {
	if o == nil || common.IsNil(o.OpenTime) {
		var ret int64
		return ret
	}
	return *o.OpenTime
}

// GetOpenTimeOk returns a tuple with the OpenTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetOpenTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OpenTime) {
		return nil, false
	}
	return o.OpenTime, true
}

// HasOpenTime returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) HasOpenTime() bool {
	if o != nil && !common.IsNil(o.OpenTime) {
		return true
	}

	return false
}

// SetOpenTime gets a reference to the given int64 and assigns it to the OpenTime field.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) SetOpenTime(v int64) {
	o.OpenTime = &v
}

// GetCloseTime returns the CloseTime field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetCloseTime() int64 {
	if o == nil || common.IsNil(o.CloseTime) {
		var ret int64
		return ret
	}
	return *o.CloseTime
}

// GetCloseTimeOk returns a tuple with the CloseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetCloseTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CloseTime) {
		return nil, false
	}
	return o.CloseTime, true
}

// HasCloseTime returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) HasCloseTime() bool {
	if o != nil && !common.IsNil(o.CloseTime) {
		return true
	}

	return false
}

// SetCloseTime gets a reference to the given int64 and assigns it to the CloseTime field.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) SetCloseTime(v int64) {
	o.CloseTime = &v
}

// GetFirstId returns the FirstId field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetFirstId() int64 {
	if o == nil || common.IsNil(o.FirstId) {
		var ret int64
		return ret
	}
	return *o.FirstId
}

// GetFirstIdOk returns a tuple with the FirstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetFirstIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FirstId) {
		return nil, false
	}
	return o.FirstId, true
}

// HasFirstId returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) HasFirstId() bool {
	if o != nil && !common.IsNil(o.FirstId) {
		return true
	}

	return false
}

// SetFirstId gets a reference to the given int64 and assigns it to the FirstId field.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) SetFirstId(v int64) {
	o.FirstId = &v
}

// GetLastId returns the LastId field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetLastId() int64 {
	if o == nil || common.IsNil(o.LastId) {
		var ret int64
		return ret
	}
	return *o.LastId
}

// GetLastIdOk returns a tuple with the LastId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetLastIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastId) {
		return nil, false
	}
	return o.LastId, true
}

// HasLastId returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) HasLastId() bool {
	if o != nil && !common.IsNil(o.LastId) {
		return true
	}

	return false
}

// SetLastId gets a reference to the given int64 and assigns it to the LastId field.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) SetLastId(v int64) {
	o.LastId = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) SetCount(v int64) {
	o.Count = &v
}

func (o Ticker24hrPriceChangeStatisticsResponse2Inner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ticker24hrPriceChangeStatisticsResponse2Inner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.LastPrice) {
		toSerialize["lastPrice"] = o.LastPrice
	}
	if !common.IsNil(o.LastQty) {
		toSerialize["lastQty"] = o.LastQty
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

func (o *Ticker24hrPriceChangeStatisticsResponse2Inner) UnmarshalJSON(data []byte) (err error) {
	varTicker24hrPriceChangeStatisticsResponse2Inner := _Ticker24hrPriceChangeStatisticsResponse2Inner{}

	err = json.Unmarshal(data, &varTicker24hrPriceChangeStatisticsResponse2Inner)

	if err != nil {
		return err
	}

	*o = Ticker24hrPriceChangeStatisticsResponse2Inner(varTicker24hrPriceChangeStatisticsResponse2Inner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "priceChange")
		delete(additionalProperties, "priceChangePercent")
		delete(additionalProperties, "weightedAvgPrice")
		delete(additionalProperties, "lastPrice")
		delete(additionalProperties, "lastQty")
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

type NullableTicker24hrPriceChangeStatisticsResponse2Inner struct {
	value *Ticker24hrPriceChangeStatisticsResponse2Inner
	isSet bool
}

func (v NullableTicker24hrPriceChangeStatisticsResponse2Inner) Get() *Ticker24hrPriceChangeStatisticsResponse2Inner {
	return v.value
}

func (v *NullableTicker24hrPriceChangeStatisticsResponse2Inner) Set(val *Ticker24hrPriceChangeStatisticsResponse2Inner) {
	v.value = val
	v.isSet = true
}

func (v NullableTicker24hrPriceChangeStatisticsResponse2Inner) IsSet() bool {
	return v.isSet
}

func (v *NullableTicker24hrPriceChangeStatisticsResponse2Inner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicker24hrPriceChangeStatisticsResponse2Inner(val *Ticker24hrPriceChangeStatisticsResponse2Inner) *NullableTicker24hrPriceChangeStatisticsResponse2Inner {
	return &NullableTicker24hrPriceChangeStatisticsResponse2Inner{value: val, isSet: true}
}

func (v NullableTicker24hrPriceChangeStatisticsResponse2Inner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicker24hrPriceChangeStatisticsResponse2Inner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
