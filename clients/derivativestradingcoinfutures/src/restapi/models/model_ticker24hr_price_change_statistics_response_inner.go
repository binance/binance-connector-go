/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the Ticker24hrPriceChangeStatisticsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Ticker24hrPriceChangeStatisticsResponseInner{}

// Ticker24hrPriceChangeStatisticsResponseInner struct for Ticker24hrPriceChangeStatisticsResponseInner
type Ticker24hrPriceChangeStatisticsResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	Pair                 *string `json:"pair,omitempty"`
	PriceChange          *string `json:"priceChange,omitempty"`
	PriceChangePercent   *string `json:"priceChangePercent,omitempty"`
	WeightedAvgPrice     *string `json:"weightedAvgPrice,omitempty"`
	LastPrice            *string `json:"lastPrice,omitempty"`
	LastQty              *string `json:"lastQty,omitempty"`
	OpenPrice            *string `json:"openPrice,omitempty"`
	HighPrice            *string `json:"highPrice,omitempty"`
	LowPrice             *string `json:"lowPrice,omitempty"`
	Volume               *string `json:"volume,omitempty"`
	BaseVolume           *string `json:"baseVolume,omitempty"`
	OpenTime             *int64  `json:"openTime,omitempty"`
	CloseTime            *int64  `json:"closeTime,omitempty"`
	FirstId              *int64  `json:"firstId,omitempty"`
	LastId               *int64  `json:"lastId,omitempty"`
	Count                *int64  `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Ticker24hrPriceChangeStatisticsResponseInner Ticker24hrPriceChangeStatisticsResponseInner

// NewTicker24hrPriceChangeStatisticsResponseInner instantiates a new Ticker24hrPriceChangeStatisticsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicker24hrPriceChangeStatisticsResponseInner() *Ticker24hrPriceChangeStatisticsResponseInner {
	this := Ticker24hrPriceChangeStatisticsResponseInner{}
	return &this
}

// NewTicker24hrPriceChangeStatisticsResponseInnerWithDefaults instantiates a new Ticker24hrPriceChangeStatisticsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicker24hrPriceChangeStatisticsResponseInnerWithDefaults() *Ticker24hrPriceChangeStatisticsResponseInner {
	this := Ticker24hrPriceChangeStatisticsResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetPair() string {
	if o == nil || common.IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetPairOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasPair() bool {
	if o != nil && !common.IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetPair(v string) {
	o.Pair = &v
}

// GetPriceChange returns the PriceChange field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetPriceChange() string {
	if o == nil || common.IsNil(o.PriceChange) {
		var ret string
		return ret
	}
	return *o.PriceChange
}

// GetPriceChangeOk returns a tuple with the PriceChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetPriceChangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceChange) {
		return nil, false
	}
	return o.PriceChange, true
}

// HasPriceChange returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasPriceChange() bool {
	if o != nil && !common.IsNil(o.PriceChange) {
		return true
	}

	return false
}

// SetPriceChange gets a reference to the given string and assigns it to the PriceChange field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetPriceChange(v string) {
	o.PriceChange = &v
}

// GetPriceChangePercent returns the PriceChangePercent field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetPriceChangePercent() string {
	if o == nil || common.IsNil(o.PriceChangePercent) {
		var ret string
		return ret
	}
	return *o.PriceChangePercent
}

// GetPriceChangePercentOk returns a tuple with the PriceChangePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetPriceChangePercentOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceChangePercent) {
		return nil, false
	}
	return o.PriceChangePercent, true
}

// HasPriceChangePercent returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasPriceChangePercent() bool {
	if o != nil && !common.IsNil(o.PriceChangePercent) {
		return true
	}

	return false
}

// SetPriceChangePercent gets a reference to the given string and assigns it to the PriceChangePercent field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetPriceChangePercent(v string) {
	o.PriceChangePercent = &v
}

// GetWeightedAvgPrice returns the WeightedAvgPrice field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetWeightedAvgPrice() string {
	if o == nil || common.IsNil(o.WeightedAvgPrice) {
		var ret string
		return ret
	}
	return *o.WeightedAvgPrice
}

// GetWeightedAvgPriceOk returns a tuple with the WeightedAvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetWeightedAvgPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.WeightedAvgPrice) {
		return nil, false
	}
	return o.WeightedAvgPrice, true
}

// HasWeightedAvgPrice returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasWeightedAvgPrice() bool {
	if o != nil && !common.IsNil(o.WeightedAvgPrice) {
		return true
	}

	return false
}

// SetWeightedAvgPrice gets a reference to the given string and assigns it to the WeightedAvgPrice field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetWeightedAvgPrice(v string) {
	o.WeightedAvgPrice = &v
}

// GetLastPrice returns the LastPrice field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetLastPrice() string {
	if o == nil || common.IsNil(o.LastPrice) {
		var ret string
		return ret
	}
	return *o.LastPrice
}

// GetLastPriceOk returns a tuple with the LastPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetLastPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastPrice) {
		return nil, false
	}
	return o.LastPrice, true
}

// HasLastPrice returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasLastPrice() bool {
	if o != nil && !common.IsNil(o.LastPrice) {
		return true
	}

	return false
}

// SetLastPrice gets a reference to the given string and assigns it to the LastPrice field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetLastPrice(v string) {
	o.LastPrice = &v
}

// GetLastQty returns the LastQty field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetLastQty() string {
	if o == nil || common.IsNil(o.LastQty) {
		var ret string
		return ret
	}
	return *o.LastQty
}

// GetLastQtyOk returns a tuple with the LastQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetLastQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastQty) {
		return nil, false
	}
	return o.LastQty, true
}

// HasLastQty returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasLastQty() bool {
	if o != nil && !common.IsNil(o.LastQty) {
		return true
	}

	return false
}

// SetLastQty gets a reference to the given string and assigns it to the LastQty field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetLastQty(v string) {
	o.LastQty = &v
}

// GetOpenPrice returns the OpenPrice field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetOpenPrice() string {
	if o == nil || common.IsNil(o.OpenPrice) {
		var ret string
		return ret
	}
	return *o.OpenPrice
}

// GetOpenPriceOk returns a tuple with the OpenPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetOpenPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpenPrice) {
		return nil, false
	}
	return o.OpenPrice, true
}

// HasOpenPrice returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasOpenPrice() bool {
	if o != nil && !common.IsNil(o.OpenPrice) {
		return true
	}

	return false
}

// SetOpenPrice gets a reference to the given string and assigns it to the OpenPrice field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetOpenPrice(v string) {
	o.OpenPrice = &v
}

// GetHighPrice returns the HighPrice field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetHighPrice() string {
	if o == nil || common.IsNil(o.HighPrice) {
		var ret string
		return ret
	}
	return *o.HighPrice
}

// GetHighPriceOk returns a tuple with the HighPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetHighPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.HighPrice) {
		return nil, false
	}
	return o.HighPrice, true
}

// HasHighPrice returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasHighPrice() bool {
	if o != nil && !common.IsNil(o.HighPrice) {
		return true
	}

	return false
}

// SetHighPrice gets a reference to the given string and assigns it to the HighPrice field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetHighPrice(v string) {
	o.HighPrice = &v
}

// GetLowPrice returns the LowPrice field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetLowPrice() string {
	if o == nil || common.IsNil(o.LowPrice) {
		var ret string
		return ret
	}
	return *o.LowPrice
}

// GetLowPriceOk returns a tuple with the LowPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetLowPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LowPrice) {
		return nil, false
	}
	return o.LowPrice, true
}

// HasLowPrice returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasLowPrice() bool {
	if o != nil && !common.IsNil(o.LowPrice) {
		return true
	}

	return false
}

// SetLowPrice gets a reference to the given string and assigns it to the LowPrice field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetLowPrice(v string) {
	o.LowPrice = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetVolume() string {
	if o == nil || common.IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetVolumeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasVolume() bool {
	if o != nil && !common.IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetVolume(v string) {
	o.Volume = &v
}

// GetBaseVolume returns the BaseVolume field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetBaseVolume() string {
	if o == nil || common.IsNil(o.BaseVolume) {
		var ret string
		return ret
	}
	return *o.BaseVolume
}

// GetBaseVolumeOk returns a tuple with the BaseVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetBaseVolumeOk() (*string, bool) {
	if o == nil || common.IsNil(o.BaseVolume) {
		return nil, false
	}
	return o.BaseVolume, true
}

// HasBaseVolume returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasBaseVolume() bool {
	if o != nil && !common.IsNil(o.BaseVolume) {
		return true
	}

	return false
}

// SetBaseVolume gets a reference to the given string and assigns it to the BaseVolume field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetBaseVolume(v string) {
	o.BaseVolume = &v
}

// GetOpenTime returns the OpenTime field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetOpenTime() int64 {
	if o == nil || common.IsNil(o.OpenTime) {
		var ret int64
		return ret
	}
	return *o.OpenTime
}

// GetOpenTimeOk returns a tuple with the OpenTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetOpenTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OpenTime) {
		return nil, false
	}
	return o.OpenTime, true
}

// HasOpenTime returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasOpenTime() bool {
	if o != nil && !common.IsNil(o.OpenTime) {
		return true
	}

	return false
}

// SetOpenTime gets a reference to the given int64 and assigns it to the OpenTime field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetOpenTime(v int64) {
	o.OpenTime = &v
}

// GetCloseTime returns the CloseTime field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetCloseTime() int64 {
	if o == nil || common.IsNil(o.CloseTime) {
		var ret int64
		return ret
	}
	return *o.CloseTime
}

// GetCloseTimeOk returns a tuple with the CloseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetCloseTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CloseTime) {
		return nil, false
	}
	return o.CloseTime, true
}

// HasCloseTime returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasCloseTime() bool {
	if o != nil && !common.IsNil(o.CloseTime) {
		return true
	}

	return false
}

// SetCloseTime gets a reference to the given int64 and assigns it to the CloseTime field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetCloseTime(v int64) {
	o.CloseTime = &v
}

// GetFirstId returns the FirstId field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetFirstId() int64 {
	if o == nil || common.IsNil(o.FirstId) {
		var ret int64
		return ret
	}
	return *o.FirstId
}

// GetFirstIdOk returns a tuple with the FirstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetFirstIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FirstId) {
		return nil, false
	}
	return o.FirstId, true
}

// HasFirstId returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasFirstId() bool {
	if o != nil && !common.IsNil(o.FirstId) {
		return true
	}

	return false
}

// SetFirstId gets a reference to the given int64 and assigns it to the FirstId field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetFirstId(v int64) {
	o.FirstId = &v
}

// GetLastId returns the LastId field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetLastId() int64 {
	if o == nil || common.IsNil(o.LastId) {
		var ret int64
		return ret
	}
	return *o.LastId
}

// GetLastIdOk returns a tuple with the LastId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetLastIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastId) {
		return nil, false
	}
	return o.LastId, true
}

// HasLastId returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasLastId() bool {
	if o != nil && !common.IsNil(o.LastId) {
		return true
	}

	return false
}

// SetLastId gets a reference to the given int64 and assigns it to the LastId field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetLastId(v int64) {
	o.LastId = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetCount(v int64) {
	o.Count = &v
}

func (o Ticker24hrPriceChangeStatisticsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ticker24hrPriceChangeStatisticsResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
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
	if !common.IsNil(o.BaseVolume) {
		toSerialize["baseVolume"] = o.BaseVolume
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

func (o *Ticker24hrPriceChangeStatisticsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varTicker24hrPriceChangeStatisticsResponseInner := _Ticker24hrPriceChangeStatisticsResponseInner{}

	err = json.Unmarshal(data, &varTicker24hrPriceChangeStatisticsResponseInner)

	if err != nil {
		return err
	}

	*o = Ticker24hrPriceChangeStatisticsResponseInner(varTicker24hrPriceChangeStatisticsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "pair")
		delete(additionalProperties, "priceChange")
		delete(additionalProperties, "priceChangePercent")
		delete(additionalProperties, "weightedAvgPrice")
		delete(additionalProperties, "lastPrice")
		delete(additionalProperties, "lastQty")
		delete(additionalProperties, "openPrice")
		delete(additionalProperties, "highPrice")
		delete(additionalProperties, "lowPrice")
		delete(additionalProperties, "volume")
		delete(additionalProperties, "baseVolume")
		delete(additionalProperties, "openTime")
		delete(additionalProperties, "closeTime")
		delete(additionalProperties, "firstId")
		delete(additionalProperties, "lastId")
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTicker24hrPriceChangeStatisticsResponseInner struct {
	value *Ticker24hrPriceChangeStatisticsResponseInner
	isSet bool
}

func (v NullableTicker24hrPriceChangeStatisticsResponseInner) Get() *Ticker24hrPriceChangeStatisticsResponseInner {
	return v.value
}

func (v *NullableTicker24hrPriceChangeStatisticsResponseInner) Set(val *Ticker24hrPriceChangeStatisticsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTicker24hrPriceChangeStatisticsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTicker24hrPriceChangeStatisticsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicker24hrPriceChangeStatisticsResponseInner(val *Ticker24hrPriceChangeStatisticsResponseInner) *NullableTicker24hrPriceChangeStatisticsResponseInner {
	return &NullableTicker24hrPriceChangeStatisticsResponseInner{value: val, isSet: true}
}

func (v NullableTicker24hrPriceChangeStatisticsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicker24hrPriceChangeStatisticsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
