/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MultiAssetsModeAssetIndexResponse2Inner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MultiAssetsModeAssetIndexResponse2Inner{}

// MultiAssetsModeAssetIndexResponse2Inner struct for MultiAssetsModeAssetIndexResponse2Inner
type MultiAssetsModeAssetIndexResponse2Inner struct {
	Symbol                *string `json:"symbol,omitempty"`
	Time                  *int64  `json:"time,omitempty"`
	Index                 *string `json:"index,omitempty"`
	BidBuffer             *string `json:"bidBuffer,omitempty"`
	AskBuffer             *string `json:"askBuffer,omitempty"`
	BidRate               *string `json:"bidRate,omitempty"`
	AskRate               *string `json:"askRate,omitempty"`
	AutoExchangeBidBuffer *string `json:"autoExchangeBidBuffer,omitempty"`
	AutoExchangeAskBuffer *string `json:"autoExchangeAskBuffer,omitempty"`
	AutoExchangeBidRate   *string `json:"autoExchangeBidRate,omitempty"`
	AutoExchangeAskRate   *string `json:"autoExchangeAskRate,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _MultiAssetsModeAssetIndexResponse2Inner MultiAssetsModeAssetIndexResponse2Inner

// NewMultiAssetsModeAssetIndexResponse2Inner instantiates a new MultiAssetsModeAssetIndexResponse2Inner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiAssetsModeAssetIndexResponse2Inner() *MultiAssetsModeAssetIndexResponse2Inner {
	this := MultiAssetsModeAssetIndexResponse2Inner{}
	return &this
}

// NewMultiAssetsModeAssetIndexResponse2InnerWithDefaults instantiates a new MultiAssetsModeAssetIndexResponse2Inner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiAssetsModeAssetIndexResponse2InnerWithDefaults() *MultiAssetsModeAssetIndexResponse2Inner {
	this := MultiAssetsModeAssetIndexResponse2Inner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MultiAssetsModeAssetIndexResponse2Inner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *MultiAssetsModeAssetIndexResponse2Inner) SetTime(v int64) {
	o.Time = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetIndex() string {
	if o == nil || common.IsNil(o.Index) {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetIndexOk() (*string, bool) {
	if o == nil || common.IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) HasIndex() bool {
	if o != nil && !common.IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *MultiAssetsModeAssetIndexResponse2Inner) SetIndex(v string) {
	o.Index = &v
}

// GetBidBuffer returns the BidBuffer field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetBidBuffer() string {
	if o == nil || common.IsNil(o.BidBuffer) {
		var ret string
		return ret
	}
	return *o.BidBuffer
}

// GetBidBufferOk returns a tuple with the BidBuffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetBidBufferOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidBuffer) {
		return nil, false
	}
	return o.BidBuffer, true
}

// HasBidBuffer returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) HasBidBuffer() bool {
	if o != nil && !common.IsNil(o.BidBuffer) {
		return true
	}

	return false
}

// SetBidBuffer gets a reference to the given string and assigns it to the BidBuffer field.
func (o *MultiAssetsModeAssetIndexResponse2Inner) SetBidBuffer(v string) {
	o.BidBuffer = &v
}

// GetAskBuffer returns the AskBuffer field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetAskBuffer() string {
	if o == nil || common.IsNil(o.AskBuffer) {
		var ret string
		return ret
	}
	return *o.AskBuffer
}

// GetAskBufferOk returns a tuple with the AskBuffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetAskBufferOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskBuffer) {
		return nil, false
	}
	return o.AskBuffer, true
}

// HasAskBuffer returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) HasAskBuffer() bool {
	if o != nil && !common.IsNil(o.AskBuffer) {
		return true
	}

	return false
}

// SetAskBuffer gets a reference to the given string and assigns it to the AskBuffer field.
func (o *MultiAssetsModeAssetIndexResponse2Inner) SetAskBuffer(v string) {
	o.AskBuffer = &v
}

// GetBidRate returns the BidRate field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetBidRate() string {
	if o == nil || common.IsNil(o.BidRate) {
		var ret string
		return ret
	}
	return *o.BidRate
}

// GetBidRateOk returns a tuple with the BidRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetBidRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidRate) {
		return nil, false
	}
	return o.BidRate, true
}

// HasBidRate returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) HasBidRate() bool {
	if o != nil && !common.IsNil(o.BidRate) {
		return true
	}

	return false
}

// SetBidRate gets a reference to the given string and assigns it to the BidRate field.
func (o *MultiAssetsModeAssetIndexResponse2Inner) SetBidRate(v string) {
	o.BidRate = &v
}

// GetAskRate returns the AskRate field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetAskRate() string {
	if o == nil || common.IsNil(o.AskRate) {
		var ret string
		return ret
	}
	return *o.AskRate
}

// GetAskRateOk returns a tuple with the AskRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetAskRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskRate) {
		return nil, false
	}
	return o.AskRate, true
}

// HasAskRate returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) HasAskRate() bool {
	if o != nil && !common.IsNil(o.AskRate) {
		return true
	}

	return false
}

// SetAskRate gets a reference to the given string and assigns it to the AskRate field.
func (o *MultiAssetsModeAssetIndexResponse2Inner) SetAskRate(v string) {
	o.AskRate = &v
}

// GetAutoExchangeBidBuffer returns the AutoExchangeBidBuffer field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetAutoExchangeBidBuffer() string {
	if o == nil || common.IsNil(o.AutoExchangeBidBuffer) {
		var ret string
		return ret
	}
	return *o.AutoExchangeBidBuffer
}

// GetAutoExchangeBidBufferOk returns a tuple with the AutoExchangeBidBuffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetAutoExchangeBidBufferOk() (*string, bool) {
	if o == nil || common.IsNil(o.AutoExchangeBidBuffer) {
		return nil, false
	}
	return o.AutoExchangeBidBuffer, true
}

// HasAutoExchangeBidBuffer returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) HasAutoExchangeBidBuffer() bool {
	if o != nil && !common.IsNil(o.AutoExchangeBidBuffer) {
		return true
	}

	return false
}

// SetAutoExchangeBidBuffer gets a reference to the given string and assigns it to the AutoExchangeBidBuffer field.
func (o *MultiAssetsModeAssetIndexResponse2Inner) SetAutoExchangeBidBuffer(v string) {
	o.AutoExchangeBidBuffer = &v
}

// GetAutoExchangeAskBuffer returns the AutoExchangeAskBuffer field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetAutoExchangeAskBuffer() string {
	if o == nil || common.IsNil(o.AutoExchangeAskBuffer) {
		var ret string
		return ret
	}
	return *o.AutoExchangeAskBuffer
}

// GetAutoExchangeAskBufferOk returns a tuple with the AutoExchangeAskBuffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetAutoExchangeAskBufferOk() (*string, bool) {
	if o == nil || common.IsNil(o.AutoExchangeAskBuffer) {
		return nil, false
	}
	return o.AutoExchangeAskBuffer, true
}

// HasAutoExchangeAskBuffer returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) HasAutoExchangeAskBuffer() bool {
	if o != nil && !common.IsNil(o.AutoExchangeAskBuffer) {
		return true
	}

	return false
}

// SetAutoExchangeAskBuffer gets a reference to the given string and assigns it to the AutoExchangeAskBuffer field.
func (o *MultiAssetsModeAssetIndexResponse2Inner) SetAutoExchangeAskBuffer(v string) {
	o.AutoExchangeAskBuffer = &v
}

// GetAutoExchangeBidRate returns the AutoExchangeBidRate field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetAutoExchangeBidRate() string {
	if o == nil || common.IsNil(o.AutoExchangeBidRate) {
		var ret string
		return ret
	}
	return *o.AutoExchangeBidRate
}

// GetAutoExchangeBidRateOk returns a tuple with the AutoExchangeBidRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetAutoExchangeBidRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.AutoExchangeBidRate) {
		return nil, false
	}
	return o.AutoExchangeBidRate, true
}

// HasAutoExchangeBidRate returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) HasAutoExchangeBidRate() bool {
	if o != nil && !common.IsNil(o.AutoExchangeBidRate) {
		return true
	}

	return false
}

// SetAutoExchangeBidRate gets a reference to the given string and assigns it to the AutoExchangeBidRate field.
func (o *MultiAssetsModeAssetIndexResponse2Inner) SetAutoExchangeBidRate(v string) {
	o.AutoExchangeBidRate = &v
}

// GetAutoExchangeAskRate returns the AutoExchangeAskRate field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetAutoExchangeAskRate() string {
	if o == nil || common.IsNil(o.AutoExchangeAskRate) {
		var ret string
		return ret
	}
	return *o.AutoExchangeAskRate
}

// GetAutoExchangeAskRateOk returns a tuple with the AutoExchangeAskRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) GetAutoExchangeAskRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.AutoExchangeAskRate) {
		return nil, false
	}
	return o.AutoExchangeAskRate, true
}

// HasAutoExchangeAskRate returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponse2Inner) HasAutoExchangeAskRate() bool {
	if o != nil && !common.IsNil(o.AutoExchangeAskRate) {
		return true
	}

	return false
}

// SetAutoExchangeAskRate gets a reference to the given string and assigns it to the AutoExchangeAskRate field.
func (o *MultiAssetsModeAssetIndexResponse2Inner) SetAutoExchangeAskRate(v string) {
	o.AutoExchangeAskRate = &v
}

func (o MultiAssetsModeAssetIndexResponse2Inner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiAssetsModeAssetIndexResponse2Inner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !common.IsNil(o.BidBuffer) {
		toSerialize["bidBuffer"] = o.BidBuffer
	}
	if !common.IsNil(o.AskBuffer) {
		toSerialize["askBuffer"] = o.AskBuffer
	}
	if !common.IsNil(o.BidRate) {
		toSerialize["bidRate"] = o.BidRate
	}
	if !common.IsNil(o.AskRate) {
		toSerialize["askRate"] = o.AskRate
	}
	if !common.IsNil(o.AutoExchangeBidBuffer) {
		toSerialize["autoExchangeBidBuffer"] = o.AutoExchangeBidBuffer
	}
	if !common.IsNil(o.AutoExchangeAskBuffer) {
		toSerialize["autoExchangeAskBuffer"] = o.AutoExchangeAskBuffer
	}
	if !common.IsNil(o.AutoExchangeBidRate) {
		toSerialize["autoExchangeBidRate"] = o.AutoExchangeBidRate
	}
	if !common.IsNil(o.AutoExchangeAskRate) {
		toSerialize["autoExchangeAskRate"] = o.AutoExchangeAskRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MultiAssetsModeAssetIndexResponse2Inner) UnmarshalJSON(data []byte) (err error) {
	varMultiAssetsModeAssetIndexResponse2Inner := _MultiAssetsModeAssetIndexResponse2Inner{}

	err = json.Unmarshal(data, &varMultiAssetsModeAssetIndexResponse2Inner)

	if err != nil {
		return err
	}

	*o = MultiAssetsModeAssetIndexResponse2Inner(varMultiAssetsModeAssetIndexResponse2Inner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "time")
		delete(additionalProperties, "index")
		delete(additionalProperties, "bidBuffer")
		delete(additionalProperties, "askBuffer")
		delete(additionalProperties, "bidRate")
		delete(additionalProperties, "askRate")
		delete(additionalProperties, "autoExchangeBidBuffer")
		delete(additionalProperties, "autoExchangeAskBuffer")
		delete(additionalProperties, "autoExchangeBidRate")
		delete(additionalProperties, "autoExchangeAskRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMultiAssetsModeAssetIndexResponse2Inner struct {
	value *MultiAssetsModeAssetIndexResponse2Inner
	isSet bool
}

func (v NullableMultiAssetsModeAssetIndexResponse2Inner) Get() *MultiAssetsModeAssetIndexResponse2Inner {
	return v.value
}

func (v *NullableMultiAssetsModeAssetIndexResponse2Inner) Set(val *MultiAssetsModeAssetIndexResponse2Inner) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiAssetsModeAssetIndexResponse2Inner) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiAssetsModeAssetIndexResponse2Inner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiAssetsModeAssetIndexResponse2Inner(val *MultiAssetsModeAssetIndexResponse2Inner) *NullableMultiAssetsModeAssetIndexResponse2Inner {
	return &NullableMultiAssetsModeAssetIndexResponse2Inner{value: val, isSet: true}
}

func (v NullableMultiAssetsModeAssetIndexResponse2Inner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiAssetsModeAssetIndexResponse2Inner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
