/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the IndexPriceAndMarkPriceResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IndexPriceAndMarkPriceResponseInner{}

// IndexPriceAndMarkPriceResponseInner struct for IndexPriceAndMarkPriceResponseInner
type IndexPriceAndMarkPriceResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	Pair                 *string `json:"pair,omitempty"`
	MarkPrice            *string `json:"markPrice,omitempty"`
	IndexPrice           *string `json:"indexPrice,omitempty"`
	EstimatedSettlePrice *string `json:"estimatedSettlePrice,omitempty"`
	LastFundingRate      *string `json:"lastFundingRate,omitempty"`
	InterestRate         *string `json:"interestRate,omitempty"`
	NextFundingTime      *int64  `json:"nextFundingTime,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IndexPriceAndMarkPriceResponseInner IndexPriceAndMarkPriceResponseInner

// NewIndexPriceAndMarkPriceResponseInner instantiates a new IndexPriceAndMarkPriceResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexPriceAndMarkPriceResponseInner() *IndexPriceAndMarkPriceResponseInner {
	this := IndexPriceAndMarkPriceResponseInner{}
	return &this
}

// NewIndexPriceAndMarkPriceResponseInnerWithDefaults instantiates a new IndexPriceAndMarkPriceResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexPriceAndMarkPriceResponseInnerWithDefaults() *IndexPriceAndMarkPriceResponseInner {
	this := IndexPriceAndMarkPriceResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *IndexPriceAndMarkPriceResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceAndMarkPriceResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *IndexPriceAndMarkPriceResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *IndexPriceAndMarkPriceResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *IndexPriceAndMarkPriceResponseInner) GetPair() string {
	if o == nil || common.IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceAndMarkPriceResponseInner) GetPairOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *IndexPriceAndMarkPriceResponseInner) HasPair() bool {
	if o != nil && !common.IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *IndexPriceAndMarkPriceResponseInner) SetPair(v string) {
	o.Pair = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *IndexPriceAndMarkPriceResponseInner) GetMarkPrice() string {
	if o == nil || common.IsNil(o.MarkPrice) {
		var ret string
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceAndMarkPriceResponseInner) GetMarkPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *IndexPriceAndMarkPriceResponseInner) HasMarkPrice() bool {
	if o != nil && !common.IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given string and assigns it to the MarkPrice field.
func (o *IndexPriceAndMarkPriceResponseInner) SetMarkPrice(v string) {
	o.MarkPrice = &v
}

// GetIndexPrice returns the IndexPrice field value if set, zero value otherwise.
func (o *IndexPriceAndMarkPriceResponseInner) GetIndexPrice() string {
	if o == nil || common.IsNil(o.IndexPrice) {
		var ret string
		return ret
	}
	return *o.IndexPrice
}

// GetIndexPriceOk returns a tuple with the IndexPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceAndMarkPriceResponseInner) GetIndexPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.IndexPrice) {
		return nil, false
	}
	return o.IndexPrice, true
}

// HasIndexPrice returns a boolean if a field has been set.
func (o *IndexPriceAndMarkPriceResponseInner) HasIndexPrice() bool {
	if o != nil && !common.IsNil(o.IndexPrice) {
		return true
	}

	return false
}

// SetIndexPrice gets a reference to the given string and assigns it to the IndexPrice field.
func (o *IndexPriceAndMarkPriceResponseInner) SetIndexPrice(v string) {
	o.IndexPrice = &v
}

// GetEstimatedSettlePrice returns the EstimatedSettlePrice field value if set, zero value otherwise.
func (o *IndexPriceAndMarkPriceResponseInner) GetEstimatedSettlePrice() string {
	if o == nil || common.IsNil(o.EstimatedSettlePrice) {
		var ret string
		return ret
	}
	return *o.EstimatedSettlePrice
}

// GetEstimatedSettlePriceOk returns a tuple with the EstimatedSettlePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceAndMarkPriceResponseInner) GetEstimatedSettlePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EstimatedSettlePrice) {
		return nil, false
	}
	return o.EstimatedSettlePrice, true
}

// HasEstimatedSettlePrice returns a boolean if a field has been set.
func (o *IndexPriceAndMarkPriceResponseInner) HasEstimatedSettlePrice() bool {
	if o != nil && !common.IsNil(o.EstimatedSettlePrice) {
		return true
	}

	return false
}

// SetEstimatedSettlePrice gets a reference to the given string and assigns it to the EstimatedSettlePrice field.
func (o *IndexPriceAndMarkPriceResponseInner) SetEstimatedSettlePrice(v string) {
	o.EstimatedSettlePrice = &v
}

// GetLastFundingRate returns the LastFundingRate field value if set, zero value otherwise.
func (o *IndexPriceAndMarkPriceResponseInner) GetLastFundingRate() string {
	if o == nil || common.IsNil(o.LastFundingRate) {
		var ret string
		return ret
	}
	return *o.LastFundingRate
}

// GetLastFundingRateOk returns a tuple with the LastFundingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceAndMarkPriceResponseInner) GetLastFundingRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastFundingRate) {
		return nil, false
	}
	return o.LastFundingRate, true
}

// HasLastFundingRate returns a boolean if a field has been set.
func (o *IndexPriceAndMarkPriceResponseInner) HasLastFundingRate() bool {
	if o != nil && !common.IsNil(o.LastFundingRate) {
		return true
	}

	return false
}

// SetLastFundingRate gets a reference to the given string and assigns it to the LastFundingRate field.
func (o *IndexPriceAndMarkPriceResponseInner) SetLastFundingRate(v string) {
	o.LastFundingRate = &v
}

// GetInterestRate returns the InterestRate field value if set, zero value otherwise.
func (o *IndexPriceAndMarkPriceResponseInner) GetInterestRate() string {
	if o == nil || common.IsNil(o.InterestRate) {
		var ret string
		return ret
	}
	return *o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceAndMarkPriceResponseInner) GetInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.InterestRate) {
		return nil, false
	}
	return o.InterestRate, true
}

// HasInterestRate returns a boolean if a field has been set.
func (o *IndexPriceAndMarkPriceResponseInner) HasInterestRate() bool {
	if o != nil && !common.IsNil(o.InterestRate) {
		return true
	}

	return false
}

// SetInterestRate gets a reference to the given string and assigns it to the InterestRate field.
func (o *IndexPriceAndMarkPriceResponseInner) SetInterestRate(v string) {
	o.InterestRate = &v
}

// GetNextFundingTime returns the NextFundingTime field value if set, zero value otherwise.
func (o *IndexPriceAndMarkPriceResponseInner) GetNextFundingTime() int64 {
	if o == nil || common.IsNil(o.NextFundingTime) {
		var ret int64
		return ret
	}
	return *o.NextFundingTime
}

// GetNextFundingTimeOk returns a tuple with the NextFundingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceAndMarkPriceResponseInner) GetNextFundingTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.NextFundingTime) {
		return nil, false
	}
	return o.NextFundingTime, true
}

// HasNextFundingTime returns a boolean if a field has been set.
func (o *IndexPriceAndMarkPriceResponseInner) HasNextFundingTime() bool {
	if o != nil && !common.IsNil(o.NextFundingTime) {
		return true
	}

	return false
}

// SetNextFundingTime gets a reference to the given int64 and assigns it to the NextFundingTime field.
func (o *IndexPriceAndMarkPriceResponseInner) SetNextFundingTime(v int64) {
	o.NextFundingTime = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *IndexPriceAndMarkPriceResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceAndMarkPriceResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *IndexPriceAndMarkPriceResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *IndexPriceAndMarkPriceResponseInner) SetTime(v int64) {
	o.Time = &v
}

func (o IndexPriceAndMarkPriceResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexPriceAndMarkPriceResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !common.IsNil(o.MarkPrice) {
		toSerialize["markPrice"] = o.MarkPrice
	}
	if !common.IsNil(o.IndexPrice) {
		toSerialize["indexPrice"] = o.IndexPrice
	}
	if !common.IsNil(o.EstimatedSettlePrice) {
		toSerialize["estimatedSettlePrice"] = o.EstimatedSettlePrice
	}
	if !common.IsNil(o.LastFundingRate) {
		toSerialize["lastFundingRate"] = o.LastFundingRate
	}
	if !common.IsNil(o.InterestRate) {
		toSerialize["interestRate"] = o.InterestRate
	}
	if !common.IsNil(o.NextFundingTime) {
		toSerialize["nextFundingTime"] = o.NextFundingTime
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IndexPriceAndMarkPriceResponseInner) UnmarshalJSON(data []byte) (err error) {
	varIndexPriceAndMarkPriceResponseInner := _IndexPriceAndMarkPriceResponseInner{}

	err = json.Unmarshal(data, &varIndexPriceAndMarkPriceResponseInner)

	if err != nil {
		return err
	}

	*o = IndexPriceAndMarkPriceResponseInner(varIndexPriceAndMarkPriceResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "pair")
		delete(additionalProperties, "markPrice")
		delete(additionalProperties, "indexPrice")
		delete(additionalProperties, "estimatedSettlePrice")
		delete(additionalProperties, "lastFundingRate")
		delete(additionalProperties, "interestRate")
		delete(additionalProperties, "nextFundingTime")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIndexPriceAndMarkPriceResponseInner struct {
	value *IndexPriceAndMarkPriceResponseInner
	isSet bool
}

func (v NullableIndexPriceAndMarkPriceResponseInner) Get() *IndexPriceAndMarkPriceResponseInner {
	return v.value
}

func (v *NullableIndexPriceAndMarkPriceResponseInner) Set(val *IndexPriceAndMarkPriceResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexPriceAndMarkPriceResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexPriceAndMarkPriceResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexPriceAndMarkPriceResponseInner(val *IndexPriceAndMarkPriceResponseInner) *NullableIndexPriceAndMarkPriceResponseInner {
	return &NullableIndexPriceAndMarkPriceResponseInner{value: val, isSet: true}
}

func (v NullableIndexPriceAndMarkPriceResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexPriceAndMarkPriceResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
