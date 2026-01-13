/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFundingRateHistoryResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFundingRateHistoryResponseInner{}

// GetFundingRateHistoryResponseInner struct for GetFundingRateHistoryResponseInner
type GetFundingRateHistoryResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	FundingRate          *string `json:"fundingRate,omitempty"`
	FundingTime          *int64  `json:"fundingTime,omitempty"`
	MarkPrice            *string `json:"markPrice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFundingRateHistoryResponseInner GetFundingRateHistoryResponseInner

// NewGetFundingRateHistoryResponseInner instantiates a new GetFundingRateHistoryResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFundingRateHistoryResponseInner() *GetFundingRateHistoryResponseInner {
	this := GetFundingRateHistoryResponseInner{}
	return &this
}

// NewGetFundingRateHistoryResponseInnerWithDefaults instantiates a new GetFundingRateHistoryResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFundingRateHistoryResponseInnerWithDefaults() *GetFundingRateHistoryResponseInner {
	this := GetFundingRateHistoryResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetFundingRateHistoryResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundingRateHistoryResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetFundingRateHistoryResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetFundingRateHistoryResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetFundingRate returns the FundingRate field value if set, zero value otherwise.
func (o *GetFundingRateHistoryResponseInner) GetFundingRate() string {
	if o == nil || common.IsNil(o.FundingRate) {
		var ret string
		return ret
	}
	return *o.FundingRate
}

// GetFundingRateOk returns a tuple with the FundingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundingRateHistoryResponseInner) GetFundingRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.FundingRate) {
		return nil, false
	}
	return o.FundingRate, true
}

// HasFundingRate returns a boolean if a field has been set.
func (o *GetFundingRateHistoryResponseInner) HasFundingRate() bool {
	if o != nil && !common.IsNil(o.FundingRate) {
		return true
	}

	return false
}

// SetFundingRate gets a reference to the given string and assigns it to the FundingRate field.
func (o *GetFundingRateHistoryResponseInner) SetFundingRate(v string) {
	o.FundingRate = &v
}

// GetFundingTime returns the FundingTime field value if set, zero value otherwise.
func (o *GetFundingRateHistoryResponseInner) GetFundingTime() int64 {
	if o == nil || common.IsNil(o.FundingTime) {
		var ret int64
		return ret
	}
	return *o.FundingTime
}

// GetFundingTimeOk returns a tuple with the FundingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundingRateHistoryResponseInner) GetFundingTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FundingTime) {
		return nil, false
	}
	return o.FundingTime, true
}

// HasFundingTime returns a boolean if a field has been set.
func (o *GetFundingRateHistoryResponseInner) HasFundingTime() bool {
	if o != nil && !common.IsNil(o.FundingTime) {
		return true
	}

	return false
}

// SetFundingTime gets a reference to the given int64 and assigns it to the FundingTime field.
func (o *GetFundingRateHistoryResponseInner) SetFundingTime(v int64) {
	o.FundingTime = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *GetFundingRateHistoryResponseInner) GetMarkPrice() string {
	if o == nil || common.IsNil(o.MarkPrice) {
		var ret string
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundingRateHistoryResponseInner) GetMarkPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *GetFundingRateHistoryResponseInner) HasMarkPrice() bool {
	if o != nil && !common.IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given string and assigns it to the MarkPrice field.
func (o *GetFundingRateHistoryResponseInner) SetMarkPrice(v string) {
	o.MarkPrice = &v
}

func (o GetFundingRateHistoryResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFundingRateHistoryResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.FundingRate) {
		toSerialize["fundingRate"] = o.FundingRate
	}
	if !common.IsNil(o.FundingTime) {
		toSerialize["fundingTime"] = o.FundingTime
	}
	if !common.IsNil(o.MarkPrice) {
		toSerialize["markPrice"] = o.MarkPrice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFundingRateHistoryResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetFundingRateHistoryResponseInner := _GetFundingRateHistoryResponseInner{}

	err = json.Unmarshal(data, &varGetFundingRateHistoryResponseInner)

	if err != nil {
		return err
	}

	*o = GetFundingRateHistoryResponseInner(varGetFundingRateHistoryResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "fundingRate")
		delete(additionalProperties, "fundingTime")
		delete(additionalProperties, "markPrice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFundingRateHistoryResponseInner struct {
	value *GetFundingRateHistoryResponseInner
	isSet bool
}

func (v NullableGetFundingRateHistoryResponseInner) Get() *GetFundingRateHistoryResponseInner {
	return v.value
}

func (v *NullableGetFundingRateHistoryResponseInner) Set(val *GetFundingRateHistoryResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFundingRateHistoryResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFundingRateHistoryResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFundingRateHistoryResponseInner(val *GetFundingRateHistoryResponseInner) *NullableGetFundingRateHistoryResponseInner {
	return &NullableGetFundingRateHistoryResponseInner{value: val, isSet: true}
}

func (v NullableGetFundingRateHistoryResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFundingRateHistoryResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
