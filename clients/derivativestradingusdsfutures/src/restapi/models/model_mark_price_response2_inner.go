/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MarkPriceResponse2Inner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarkPriceResponse2Inner{}

// MarkPriceResponse2Inner struct for MarkPriceResponse2Inner
type MarkPriceResponse2Inner struct {
	Symbol               *string `json:"symbol,omitempty"`
	MarkPrice            *string `json:"markPrice,omitempty"`
	IndexPrice           *string `json:"indexPrice,omitempty"`
	EstimatedSettlePrice *string `json:"estimatedSettlePrice,omitempty"`
	LastFundingRate      *string `json:"lastFundingRate,omitempty"`
	InterestRate         *string `json:"interestRate,omitempty"`
	NextFundingTime      *int64  `json:"nextFundingTime,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarkPriceResponse2Inner MarkPriceResponse2Inner

// NewMarkPriceResponse2Inner instantiates a new MarkPriceResponse2Inner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkPriceResponse2Inner() *MarkPriceResponse2Inner {
	this := MarkPriceResponse2Inner{}
	return &this
}

// NewMarkPriceResponse2InnerWithDefaults instantiates a new MarkPriceResponse2Inner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkPriceResponse2InnerWithDefaults() *MarkPriceResponse2Inner {
	this := MarkPriceResponse2Inner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarkPriceResponse2Inner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponse2Inner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarkPriceResponse2Inner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarkPriceResponse2Inner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *MarkPriceResponse2Inner) GetMarkPrice() string {
	if o == nil || common.IsNil(o.MarkPrice) {
		var ret string
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponse2Inner) GetMarkPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *MarkPriceResponse2Inner) HasMarkPrice() bool {
	if o != nil && !common.IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given string and assigns it to the MarkPrice field.
func (o *MarkPriceResponse2Inner) SetMarkPrice(v string) {
	o.MarkPrice = &v
}

// GetIndexPrice returns the IndexPrice field value if set, zero value otherwise.
func (o *MarkPriceResponse2Inner) GetIndexPrice() string {
	if o == nil || common.IsNil(o.IndexPrice) {
		var ret string
		return ret
	}
	return *o.IndexPrice
}

// GetIndexPriceOk returns a tuple with the IndexPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponse2Inner) GetIndexPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.IndexPrice) {
		return nil, false
	}
	return o.IndexPrice, true
}

// HasIndexPrice returns a boolean if a field has been set.
func (o *MarkPriceResponse2Inner) HasIndexPrice() bool {
	if o != nil && !common.IsNil(o.IndexPrice) {
		return true
	}

	return false
}

// SetIndexPrice gets a reference to the given string and assigns it to the IndexPrice field.
func (o *MarkPriceResponse2Inner) SetIndexPrice(v string) {
	o.IndexPrice = &v
}

// GetEstimatedSettlePrice returns the EstimatedSettlePrice field value if set, zero value otherwise.
func (o *MarkPriceResponse2Inner) GetEstimatedSettlePrice() string {
	if o == nil || common.IsNil(o.EstimatedSettlePrice) {
		var ret string
		return ret
	}
	return *o.EstimatedSettlePrice
}

// GetEstimatedSettlePriceOk returns a tuple with the EstimatedSettlePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponse2Inner) GetEstimatedSettlePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EstimatedSettlePrice) {
		return nil, false
	}
	return o.EstimatedSettlePrice, true
}

// HasEstimatedSettlePrice returns a boolean if a field has been set.
func (o *MarkPriceResponse2Inner) HasEstimatedSettlePrice() bool {
	if o != nil && !common.IsNil(o.EstimatedSettlePrice) {
		return true
	}

	return false
}

// SetEstimatedSettlePrice gets a reference to the given string and assigns it to the EstimatedSettlePrice field.
func (o *MarkPriceResponse2Inner) SetEstimatedSettlePrice(v string) {
	o.EstimatedSettlePrice = &v
}

// GetLastFundingRate returns the LastFundingRate field value if set, zero value otherwise.
func (o *MarkPriceResponse2Inner) GetLastFundingRate() string {
	if o == nil || common.IsNil(o.LastFundingRate) {
		var ret string
		return ret
	}
	return *o.LastFundingRate
}

// GetLastFundingRateOk returns a tuple with the LastFundingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponse2Inner) GetLastFundingRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastFundingRate) {
		return nil, false
	}
	return o.LastFundingRate, true
}

// HasLastFundingRate returns a boolean if a field has been set.
func (o *MarkPriceResponse2Inner) HasLastFundingRate() bool {
	if o != nil && !common.IsNil(o.LastFundingRate) {
		return true
	}

	return false
}

// SetLastFundingRate gets a reference to the given string and assigns it to the LastFundingRate field.
func (o *MarkPriceResponse2Inner) SetLastFundingRate(v string) {
	o.LastFundingRate = &v
}

// GetInterestRate returns the InterestRate field value if set, zero value otherwise.
func (o *MarkPriceResponse2Inner) GetInterestRate() string {
	if o == nil || common.IsNil(o.InterestRate) {
		var ret string
		return ret
	}
	return *o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponse2Inner) GetInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.InterestRate) {
		return nil, false
	}
	return o.InterestRate, true
}

// HasInterestRate returns a boolean if a field has been set.
func (o *MarkPriceResponse2Inner) HasInterestRate() bool {
	if o != nil && !common.IsNil(o.InterestRate) {
		return true
	}

	return false
}

// SetInterestRate gets a reference to the given string and assigns it to the InterestRate field.
func (o *MarkPriceResponse2Inner) SetInterestRate(v string) {
	o.InterestRate = &v
}

// GetNextFundingTime returns the NextFundingTime field value if set, zero value otherwise.
func (o *MarkPriceResponse2Inner) GetNextFundingTime() int64 {
	if o == nil || common.IsNil(o.NextFundingTime) {
		var ret int64
		return ret
	}
	return *o.NextFundingTime
}

// GetNextFundingTimeOk returns a tuple with the NextFundingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponse2Inner) GetNextFundingTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.NextFundingTime) {
		return nil, false
	}
	return o.NextFundingTime, true
}

// HasNextFundingTime returns a boolean if a field has been set.
func (o *MarkPriceResponse2Inner) HasNextFundingTime() bool {
	if o != nil && !common.IsNil(o.NextFundingTime) {
		return true
	}

	return false
}

// SetNextFundingTime gets a reference to the given int64 and assigns it to the NextFundingTime field.
func (o *MarkPriceResponse2Inner) SetNextFundingTime(v int64) {
	o.NextFundingTime = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *MarkPriceResponse2Inner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponse2Inner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *MarkPriceResponse2Inner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *MarkPriceResponse2Inner) SetTime(v int64) {
	o.Time = &v
}

func (o MarkPriceResponse2Inner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkPriceResponse2Inner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
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

func (o *MarkPriceResponse2Inner) UnmarshalJSON(data []byte) (err error) {
	varMarkPriceResponse2Inner := _MarkPriceResponse2Inner{}

	err = json.Unmarshal(data, &varMarkPriceResponse2Inner)

	if err != nil {
		return err
	}

	*o = MarkPriceResponse2Inner(varMarkPriceResponse2Inner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
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

type NullableMarkPriceResponse2Inner struct {
	value *MarkPriceResponse2Inner
	isSet bool
}

func (v NullableMarkPriceResponse2Inner) Get() *MarkPriceResponse2Inner {
	return v.value
}

func (v *NullableMarkPriceResponse2Inner) Set(val *MarkPriceResponse2Inner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkPriceResponse2Inner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkPriceResponse2Inner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkPriceResponse2Inner(val *MarkPriceResponse2Inner) *NullableMarkPriceResponse2Inner {
	return &NullableMarkPriceResponse2Inner{value: val, isSet: true}
}

func (v NullableMarkPriceResponse2Inner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkPriceResponse2Inner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
