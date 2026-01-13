/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the KlineCandlestickDataResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KlineCandlestickDataResponseInner{}

// KlineCandlestickDataResponseInner struct for KlineCandlestickDataResponseInner
type KlineCandlestickDataResponseInner struct {
	Open                 *string `json:"open,omitempty"`
	High                 *string `json:"high,omitempty"`
	Low                  *string `json:"low,omitempty"`
	Close                *string `json:"close,omitempty"`
	Volume               *string `json:"volume,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	Interval             *string `json:"interval,omitempty"`
	TradeCount           *int64  `json:"tradeCount,omitempty"`
	TakerVolume          *string `json:"takerVolume,omitempty"`
	TakerAmount          *string `json:"takerAmount,omitempty"`
	OpenTime             *int64  `json:"openTime,omitempty"`
	CloseTime            *int64  `json:"closeTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KlineCandlestickDataResponseInner KlineCandlestickDataResponseInner

// NewKlineCandlestickDataResponseInner instantiates a new KlineCandlestickDataResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlineCandlestickDataResponseInner() *KlineCandlestickDataResponseInner {
	this := KlineCandlestickDataResponseInner{}
	return &this
}

// NewKlineCandlestickDataResponseInnerWithDefaults instantiates a new KlineCandlestickDataResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlineCandlestickDataResponseInnerWithDefaults() *KlineCandlestickDataResponseInner {
	this := KlineCandlestickDataResponseInner{}
	return &this
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *KlineCandlestickDataResponseInner) GetOpen() string {
	if o == nil || common.IsNil(o.Open) {
		var ret string
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineCandlestickDataResponseInner) GetOpenOk() (*string, bool) {
	if o == nil || common.IsNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *KlineCandlestickDataResponseInner) HasOpen() bool {
	if o != nil && !common.IsNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given string and assigns it to the Open field.
func (o *KlineCandlestickDataResponseInner) SetOpen(v string) {
	o.Open = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *KlineCandlestickDataResponseInner) GetHigh() string {
	if o == nil || common.IsNil(o.High) {
		var ret string
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineCandlestickDataResponseInner) GetHighOk() (*string, bool) {
	if o == nil || common.IsNil(o.High) {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *KlineCandlestickDataResponseInner) HasHigh() bool {
	if o != nil && !common.IsNil(o.High) {
		return true
	}

	return false
}

// SetHigh gets a reference to the given string and assigns it to the High field.
func (o *KlineCandlestickDataResponseInner) SetHigh(v string) {
	o.High = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *KlineCandlestickDataResponseInner) GetLow() string {
	if o == nil || common.IsNil(o.Low) {
		var ret string
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineCandlestickDataResponseInner) GetLowOk() (*string, bool) {
	if o == nil || common.IsNil(o.Low) {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *KlineCandlestickDataResponseInner) HasLow() bool {
	if o != nil && !common.IsNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given string and assigns it to the Low field.
func (o *KlineCandlestickDataResponseInner) SetLow(v string) {
	o.Low = &v
}

// GetClose returns the Close field value if set, zero value otherwise.
func (o *KlineCandlestickDataResponseInner) GetClose() string {
	if o == nil || common.IsNil(o.Close) {
		var ret string
		return ret
	}
	return *o.Close
}

// GetCloseOk returns a tuple with the Close field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineCandlestickDataResponseInner) GetCloseOk() (*string, bool) {
	if o == nil || common.IsNil(o.Close) {
		return nil, false
	}
	return o.Close, true
}

// HasClose returns a boolean if a field has been set.
func (o *KlineCandlestickDataResponseInner) HasClose() bool {
	if o != nil && !common.IsNil(o.Close) {
		return true
	}

	return false
}

// SetClose gets a reference to the given string and assigns it to the Close field.
func (o *KlineCandlestickDataResponseInner) SetClose(v string) {
	o.Close = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *KlineCandlestickDataResponseInner) GetVolume() string {
	if o == nil || common.IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineCandlestickDataResponseInner) GetVolumeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *KlineCandlestickDataResponseInner) HasVolume() bool {
	if o != nil && !common.IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *KlineCandlestickDataResponseInner) SetVolume(v string) {
	o.Volume = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *KlineCandlestickDataResponseInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineCandlestickDataResponseInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *KlineCandlestickDataResponseInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *KlineCandlestickDataResponseInner) SetAmount(v string) {
	o.Amount = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *KlineCandlestickDataResponseInner) GetInterval() string {
	if o == nil || common.IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineCandlestickDataResponseInner) GetIntervalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *KlineCandlestickDataResponseInner) HasInterval() bool {
	if o != nil && !common.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *KlineCandlestickDataResponseInner) SetInterval(v string) {
	o.Interval = &v
}

// GetTradeCount returns the TradeCount field value if set, zero value otherwise.
func (o *KlineCandlestickDataResponseInner) GetTradeCount() int64 {
	if o == nil || common.IsNil(o.TradeCount) {
		var ret int64
		return ret
	}
	return *o.TradeCount
}

// GetTradeCountOk returns a tuple with the TradeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineCandlestickDataResponseInner) GetTradeCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeCount) {
		return nil, false
	}
	return o.TradeCount, true
}

// HasTradeCount returns a boolean if a field has been set.
func (o *KlineCandlestickDataResponseInner) HasTradeCount() bool {
	if o != nil && !common.IsNil(o.TradeCount) {
		return true
	}

	return false
}

// SetTradeCount gets a reference to the given int64 and assigns it to the TradeCount field.
func (o *KlineCandlestickDataResponseInner) SetTradeCount(v int64) {
	o.TradeCount = &v
}

// GetTakerVolume returns the TakerVolume field value if set, zero value otherwise.
func (o *KlineCandlestickDataResponseInner) GetTakerVolume() string {
	if o == nil || common.IsNil(o.TakerVolume) {
		var ret string
		return ret
	}
	return *o.TakerVolume
}

// GetTakerVolumeOk returns a tuple with the TakerVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineCandlestickDataResponseInner) GetTakerVolumeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerVolume) {
		return nil, false
	}
	return o.TakerVolume, true
}

// HasTakerVolume returns a boolean if a field has been set.
func (o *KlineCandlestickDataResponseInner) HasTakerVolume() bool {
	if o != nil && !common.IsNil(o.TakerVolume) {
		return true
	}

	return false
}

// SetTakerVolume gets a reference to the given string and assigns it to the TakerVolume field.
func (o *KlineCandlestickDataResponseInner) SetTakerVolume(v string) {
	o.TakerVolume = &v
}

// GetTakerAmount returns the TakerAmount field value if set, zero value otherwise.
func (o *KlineCandlestickDataResponseInner) GetTakerAmount() string {
	if o == nil || common.IsNil(o.TakerAmount) {
		var ret string
		return ret
	}
	return *o.TakerAmount
}

// GetTakerAmountOk returns a tuple with the TakerAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineCandlestickDataResponseInner) GetTakerAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerAmount) {
		return nil, false
	}
	return o.TakerAmount, true
}

// HasTakerAmount returns a boolean if a field has been set.
func (o *KlineCandlestickDataResponseInner) HasTakerAmount() bool {
	if o != nil && !common.IsNil(o.TakerAmount) {
		return true
	}

	return false
}

// SetTakerAmount gets a reference to the given string and assigns it to the TakerAmount field.
func (o *KlineCandlestickDataResponseInner) SetTakerAmount(v string) {
	o.TakerAmount = &v
}

// GetOpenTime returns the OpenTime field value if set, zero value otherwise.
func (o *KlineCandlestickDataResponseInner) GetOpenTime() int64 {
	if o == nil || common.IsNil(o.OpenTime) {
		var ret int64
		return ret
	}
	return *o.OpenTime
}

// GetOpenTimeOk returns a tuple with the OpenTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineCandlestickDataResponseInner) GetOpenTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OpenTime) {
		return nil, false
	}
	return o.OpenTime, true
}

// HasOpenTime returns a boolean if a field has been set.
func (o *KlineCandlestickDataResponseInner) HasOpenTime() bool {
	if o != nil && !common.IsNil(o.OpenTime) {
		return true
	}

	return false
}

// SetOpenTime gets a reference to the given int64 and assigns it to the OpenTime field.
func (o *KlineCandlestickDataResponseInner) SetOpenTime(v int64) {
	o.OpenTime = &v
}

// GetCloseTime returns the CloseTime field value if set, zero value otherwise.
func (o *KlineCandlestickDataResponseInner) GetCloseTime() int64 {
	if o == nil || common.IsNil(o.CloseTime) {
		var ret int64
		return ret
	}
	return *o.CloseTime
}

// GetCloseTimeOk returns a tuple with the CloseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineCandlestickDataResponseInner) GetCloseTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CloseTime) {
		return nil, false
	}
	return o.CloseTime, true
}

// HasCloseTime returns a boolean if a field has been set.
func (o *KlineCandlestickDataResponseInner) HasCloseTime() bool {
	if o != nil && !common.IsNil(o.CloseTime) {
		return true
	}

	return false
}

// SetCloseTime gets a reference to the given int64 and assigns it to the CloseTime field.
func (o *KlineCandlestickDataResponseInner) SetCloseTime(v int64) {
	o.CloseTime = &v
}

func (o KlineCandlestickDataResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlineCandlestickDataResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	if !common.IsNil(o.High) {
		toSerialize["high"] = o.High
	}
	if !common.IsNil(o.Low) {
		toSerialize["low"] = o.Low
	}
	if !common.IsNil(o.Close) {
		toSerialize["close"] = o.Close
	}
	if !common.IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !common.IsNil(o.TradeCount) {
		toSerialize["tradeCount"] = o.TradeCount
	}
	if !common.IsNil(o.TakerVolume) {
		toSerialize["takerVolume"] = o.TakerVolume
	}
	if !common.IsNil(o.TakerAmount) {
		toSerialize["takerAmount"] = o.TakerAmount
	}
	if !common.IsNil(o.OpenTime) {
		toSerialize["openTime"] = o.OpenTime
	}
	if !common.IsNil(o.CloseTime) {
		toSerialize["closeTime"] = o.CloseTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KlineCandlestickDataResponseInner) UnmarshalJSON(data []byte) (err error) {
	varKlineCandlestickDataResponseInner := _KlineCandlestickDataResponseInner{}

	err = json.Unmarshal(data, &varKlineCandlestickDataResponseInner)

	if err != nil {
		return err
	}

	*o = KlineCandlestickDataResponseInner(varKlineCandlestickDataResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "open")
		delete(additionalProperties, "high")
		delete(additionalProperties, "low")
		delete(additionalProperties, "close")
		delete(additionalProperties, "volume")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "tradeCount")
		delete(additionalProperties, "takerVolume")
		delete(additionalProperties, "takerAmount")
		delete(additionalProperties, "openTime")
		delete(additionalProperties, "closeTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKlineCandlestickDataResponseInner struct {
	value *KlineCandlestickDataResponseInner
	isSet bool
}

func (v NullableKlineCandlestickDataResponseInner) Get() *KlineCandlestickDataResponseInner {
	return v.value
}

func (v *NullableKlineCandlestickDataResponseInner) Set(val *KlineCandlestickDataResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableKlineCandlestickDataResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableKlineCandlestickDataResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlineCandlestickDataResponseInner(val *KlineCandlestickDataResponseInner) *NullableKlineCandlestickDataResponseInner {
	return &NullableKlineCandlestickDataResponseInner{value: val, isSet: true}
}

func (v NullableKlineCandlestickDataResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlineCandlestickDataResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
