/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFundingRateHistoryOfPerpetualFuturesResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFundingRateHistoryOfPerpetualFuturesResponseInner{}

// GetFundingRateHistoryOfPerpetualFuturesResponseInner struct for GetFundingRateHistoryOfPerpetualFuturesResponseInner
type GetFundingRateHistoryOfPerpetualFuturesResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	FundingTime          *int64  `json:"fundingTime,omitempty"`
	FundingRate          *string `json:"fundingRate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFundingRateHistoryOfPerpetualFuturesResponseInner GetFundingRateHistoryOfPerpetualFuturesResponseInner

// NewGetFundingRateHistoryOfPerpetualFuturesResponseInner instantiates a new GetFundingRateHistoryOfPerpetualFuturesResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFundingRateHistoryOfPerpetualFuturesResponseInner() *GetFundingRateHistoryOfPerpetualFuturesResponseInner {
	this := GetFundingRateHistoryOfPerpetualFuturesResponseInner{}
	return &this
}

// NewGetFundingRateHistoryOfPerpetualFuturesResponseInnerWithDefaults instantiates a new GetFundingRateHistoryOfPerpetualFuturesResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFundingRateHistoryOfPerpetualFuturesResponseInnerWithDefaults() *GetFundingRateHistoryOfPerpetualFuturesResponseInner {
	this := GetFundingRateHistoryOfPerpetualFuturesResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetFundingRateHistoryOfPerpetualFuturesResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundingRateHistoryOfPerpetualFuturesResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetFundingRateHistoryOfPerpetualFuturesResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetFundingRateHistoryOfPerpetualFuturesResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetFundingTime returns the FundingTime field value if set, zero value otherwise.
func (o *GetFundingRateHistoryOfPerpetualFuturesResponseInner) GetFundingTime() int64 {
	if o == nil || common.IsNil(o.FundingTime) {
		var ret int64
		return ret
	}
	return *o.FundingTime
}

// GetFundingTimeOk returns a tuple with the FundingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundingRateHistoryOfPerpetualFuturesResponseInner) GetFundingTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FundingTime) {
		return nil, false
	}
	return o.FundingTime, true
}

// HasFundingTime returns a boolean if a field has been set.
func (o *GetFundingRateHistoryOfPerpetualFuturesResponseInner) HasFundingTime() bool {
	if o != nil && !common.IsNil(o.FundingTime) {
		return true
	}

	return false
}

// SetFundingTime gets a reference to the given int64 and assigns it to the FundingTime field.
func (o *GetFundingRateHistoryOfPerpetualFuturesResponseInner) SetFundingTime(v int64) {
	o.FundingTime = &v
}

// GetFundingRate returns the FundingRate field value if set, zero value otherwise.
func (o *GetFundingRateHistoryOfPerpetualFuturesResponseInner) GetFundingRate() string {
	if o == nil || common.IsNil(o.FundingRate) {
		var ret string
		return ret
	}
	return *o.FundingRate
}

// GetFundingRateOk returns a tuple with the FundingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundingRateHistoryOfPerpetualFuturesResponseInner) GetFundingRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.FundingRate) {
		return nil, false
	}
	return o.FundingRate, true
}

// HasFundingRate returns a boolean if a field has been set.
func (o *GetFundingRateHistoryOfPerpetualFuturesResponseInner) HasFundingRate() bool {
	if o != nil && !common.IsNil(o.FundingRate) {
		return true
	}

	return false
}

// SetFundingRate gets a reference to the given string and assigns it to the FundingRate field.
func (o *GetFundingRateHistoryOfPerpetualFuturesResponseInner) SetFundingRate(v string) {
	o.FundingRate = &v
}

func (o GetFundingRateHistoryOfPerpetualFuturesResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFundingRateHistoryOfPerpetualFuturesResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.FundingTime) {
		toSerialize["fundingTime"] = o.FundingTime
	}
	if !common.IsNil(o.FundingRate) {
		toSerialize["fundingRate"] = o.FundingRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFundingRateHistoryOfPerpetualFuturesResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetFundingRateHistoryOfPerpetualFuturesResponseInner := _GetFundingRateHistoryOfPerpetualFuturesResponseInner{}

	err = json.Unmarshal(data, &varGetFundingRateHistoryOfPerpetualFuturesResponseInner)

	if err != nil {
		return err
	}

	*o = GetFundingRateHistoryOfPerpetualFuturesResponseInner(varGetFundingRateHistoryOfPerpetualFuturesResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "fundingTime")
		delete(additionalProperties, "fundingRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFundingRateHistoryOfPerpetualFuturesResponseInner struct {
	value *GetFundingRateHistoryOfPerpetualFuturesResponseInner
	isSet bool
}

func (v NullableGetFundingRateHistoryOfPerpetualFuturesResponseInner) Get() *GetFundingRateHistoryOfPerpetualFuturesResponseInner {
	return v.value
}

func (v *NullableGetFundingRateHistoryOfPerpetualFuturesResponseInner) Set(val *GetFundingRateHistoryOfPerpetualFuturesResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFundingRateHistoryOfPerpetualFuturesResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFundingRateHistoryOfPerpetualFuturesResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFundingRateHistoryOfPerpetualFuturesResponseInner(val *GetFundingRateHistoryOfPerpetualFuturesResponseInner) *NullableGetFundingRateHistoryOfPerpetualFuturesResponseInner {
	return &NullableGetFundingRateHistoryOfPerpetualFuturesResponseInner{value: val, isSet: true}
}

func (v NullableGetFundingRateHistoryOfPerpetualFuturesResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFundingRateHistoryOfPerpetualFuturesResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
