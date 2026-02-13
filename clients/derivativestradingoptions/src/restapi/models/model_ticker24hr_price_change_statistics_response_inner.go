/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the Ticker24hrPriceChangeStatisticsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Ticker24hrPriceChangeStatisticsResponseInner{}

// Ticker24hrPriceChangeStatisticsResponseInner struct for Ticker24hrPriceChangeStatisticsResponseInner
type Ticker24hrPriceChangeStatisticsResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	PriceChange          *string `json:"priceChange,omitempty"`
	PriceChangePercent   *string `json:"priceChangePercent,omitempty"`
	LastPrice            *string `json:"lastPrice,omitempty"`
	LastQty              *string `json:"lastQty,omitempty"`
	Open                 *string `json:"open,omitempty"`
	High                 *string `json:"high,omitempty"`
	Low                  *string `json:"low,omitempty"`
	Volume               *string `json:"volume,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	BidPrice             *string `json:"bidPrice,omitempty"`
	AskPrice             *string `json:"askPrice,omitempty"`
	OpenTime             *int64  `json:"openTime,omitempty"`
	CloseTime            *int64  `json:"closeTime,omitempty"`
	FirstTradeId         *int64  `json:"firstTradeId,omitempty"`
	TradeCount           *int64  `json:"tradeCount,omitempty"`
	StrikePrice          *string `json:"strikePrice,omitempty"`
	ExercisePrice        *string `json:"exercisePrice,omitempty"`
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

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetOpen() string {
	if o == nil || common.IsNil(o.Open) {
		var ret string
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetOpenOk() (*string, bool) {
	if o == nil || common.IsNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasOpen() bool {
	if o != nil && !common.IsNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given string and assigns it to the Open field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetOpen(v string) {
	o.Open = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetHigh() string {
	if o == nil || common.IsNil(o.High) {
		var ret string
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetHighOk() (*string, bool) {
	if o == nil || common.IsNil(o.High) {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasHigh() bool {
	if o != nil && !common.IsNil(o.High) {
		return true
	}

	return false
}

// SetHigh gets a reference to the given string and assigns it to the High field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetHigh(v string) {
	o.High = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetLow() string {
	if o == nil || common.IsNil(o.Low) {
		var ret string
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetLowOk() (*string, bool) {
	if o == nil || common.IsNil(o.Low) {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasLow() bool {
	if o != nil && !common.IsNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given string and assigns it to the Low field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetLow(v string) {
	o.Low = &v
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

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetAmount(v string) {
	o.Amount = &v
}

// GetBidPrice returns the BidPrice field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetBidPrice() string {
	if o == nil || common.IsNil(o.BidPrice) {
		var ret string
		return ret
	}
	return *o.BidPrice
}

// GetBidPriceOk returns a tuple with the BidPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetBidPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidPrice) {
		return nil, false
	}
	return o.BidPrice, true
}

// HasBidPrice returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasBidPrice() bool {
	if o != nil && !common.IsNil(o.BidPrice) {
		return true
	}

	return false
}

// SetBidPrice gets a reference to the given string and assigns it to the BidPrice field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetBidPrice(v string) {
	o.BidPrice = &v
}

// GetAskPrice returns the AskPrice field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetAskPrice() string {
	if o == nil || common.IsNil(o.AskPrice) {
		var ret string
		return ret
	}
	return *o.AskPrice
}

// GetAskPriceOk returns a tuple with the AskPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetAskPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskPrice) {
		return nil, false
	}
	return o.AskPrice, true
}

// HasAskPrice returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasAskPrice() bool {
	if o != nil && !common.IsNil(o.AskPrice) {
		return true
	}

	return false
}

// SetAskPrice gets a reference to the given string and assigns it to the AskPrice field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetAskPrice(v string) {
	o.AskPrice = &v
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

// GetFirstTradeId returns the FirstTradeId field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetFirstTradeId() int64 {
	if o == nil || common.IsNil(o.FirstTradeId) {
		var ret int64
		return ret
	}
	return *o.FirstTradeId
}

// GetFirstTradeIdOk returns a tuple with the FirstTradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetFirstTradeIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FirstTradeId) {
		return nil, false
	}
	return o.FirstTradeId, true
}

// HasFirstTradeId returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasFirstTradeId() bool {
	if o != nil && !common.IsNil(o.FirstTradeId) {
		return true
	}

	return false
}

// SetFirstTradeId gets a reference to the given int64 and assigns it to the FirstTradeId field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetFirstTradeId(v int64) {
	o.FirstTradeId = &v
}

// GetTradeCount returns the TradeCount field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetTradeCount() int64 {
	if o == nil || common.IsNil(o.TradeCount) {
		var ret int64
		return ret
	}
	return *o.TradeCount
}

// GetTradeCountOk returns a tuple with the TradeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetTradeCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeCount) {
		return nil, false
	}
	return o.TradeCount, true
}

// HasTradeCount returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasTradeCount() bool {
	if o != nil && !common.IsNil(o.TradeCount) {
		return true
	}

	return false
}

// SetTradeCount gets a reference to the given int64 and assigns it to the TradeCount field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetTradeCount(v int64) {
	o.TradeCount = &v
}

// GetStrikePrice returns the StrikePrice field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetStrikePrice() string {
	if o == nil || common.IsNil(o.StrikePrice) {
		var ret string
		return ret
	}
	return *o.StrikePrice
}

// GetStrikePriceOk returns a tuple with the StrikePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetStrikePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StrikePrice) {
		return nil, false
	}
	return o.StrikePrice, true
}

// HasStrikePrice returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasStrikePrice() bool {
	if o != nil && !common.IsNil(o.StrikePrice) {
		return true
	}

	return false
}

// SetStrikePrice gets a reference to the given string and assigns it to the StrikePrice field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetStrikePrice(v string) {
	o.StrikePrice = &v
}

// GetExercisePrice returns the ExercisePrice field value if set, zero value otherwise.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetExercisePrice() string {
	if o == nil || common.IsNil(o.ExercisePrice) {
		var ret string
		return ret
	}
	return *o.ExercisePrice
}

// GetExercisePriceOk returns a tuple with the ExercisePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) GetExercisePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExercisePrice) {
		return nil, false
	}
	return o.ExercisePrice, true
}

// HasExercisePrice returns a boolean if a field has been set.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) HasExercisePrice() bool {
	if o != nil && !common.IsNil(o.ExercisePrice) {
		return true
	}

	return false
}

// SetExercisePrice gets a reference to the given string and assigns it to the ExercisePrice field.
func (o *Ticker24hrPriceChangeStatisticsResponseInner) SetExercisePrice(v string) {
	o.ExercisePrice = &v
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
	if !common.IsNil(o.PriceChange) {
		toSerialize["priceChange"] = o.PriceChange
	}
	if !common.IsNil(o.PriceChangePercent) {
		toSerialize["priceChangePercent"] = o.PriceChangePercent
	}
	if !common.IsNil(o.LastPrice) {
		toSerialize["lastPrice"] = o.LastPrice
	}
	if !common.IsNil(o.LastQty) {
		toSerialize["lastQty"] = o.LastQty
	}
	if !common.IsNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	if !common.IsNil(o.High) {
		toSerialize["high"] = o.High
	}
	if !common.IsNil(o.Low) {
		toSerialize["low"] = o.Low
	}
	if !common.IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.BidPrice) {
		toSerialize["bidPrice"] = o.BidPrice
	}
	if !common.IsNil(o.AskPrice) {
		toSerialize["askPrice"] = o.AskPrice
	}
	if !common.IsNil(o.OpenTime) {
		toSerialize["openTime"] = o.OpenTime
	}
	if !common.IsNil(o.CloseTime) {
		toSerialize["closeTime"] = o.CloseTime
	}
	if !common.IsNil(o.FirstTradeId) {
		toSerialize["firstTradeId"] = o.FirstTradeId
	}
	if !common.IsNil(o.TradeCount) {
		toSerialize["tradeCount"] = o.TradeCount
	}
	if !common.IsNil(o.StrikePrice) {
		toSerialize["strikePrice"] = o.StrikePrice
	}
	if !common.IsNil(o.ExercisePrice) {
		toSerialize["exercisePrice"] = o.ExercisePrice
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
		delete(additionalProperties, "priceChange")
		delete(additionalProperties, "priceChangePercent")
		delete(additionalProperties, "lastPrice")
		delete(additionalProperties, "lastQty")
		delete(additionalProperties, "open")
		delete(additionalProperties, "high")
		delete(additionalProperties, "low")
		delete(additionalProperties, "volume")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "bidPrice")
		delete(additionalProperties, "askPrice")
		delete(additionalProperties, "openTime")
		delete(additionalProperties, "closeTime")
		delete(additionalProperties, "firstTradeId")
		delete(additionalProperties, "tradeCount")
		delete(additionalProperties, "strikePrice")
		delete(additionalProperties, "exercisePrice")
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
