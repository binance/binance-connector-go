/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetWbethRateHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetWbethRateHistoryResponseRowsInner{}

// GetWbethRateHistoryResponseRowsInner struct for GetWbethRateHistoryResponseRowsInner
type GetWbethRateHistoryResponseRowsInner struct {
	AnnualPercentageRate *string `json:"annualPercentageRate,omitempty"`
	ExchangeRate         *string `json:"exchangeRate,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetWbethRateHistoryResponseRowsInner GetWbethRateHistoryResponseRowsInner

// NewGetWbethRateHistoryResponseRowsInner instantiates a new GetWbethRateHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWbethRateHistoryResponseRowsInner() *GetWbethRateHistoryResponseRowsInner {
	this := GetWbethRateHistoryResponseRowsInner{}
	return &this
}

// NewGetWbethRateHistoryResponseRowsInnerWithDefaults instantiates a new GetWbethRateHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWbethRateHistoryResponseRowsInnerWithDefaults() *GetWbethRateHistoryResponseRowsInner {
	this := GetWbethRateHistoryResponseRowsInner{}
	return &this
}

// GetAnnualPercentageRate returns the AnnualPercentageRate field value if set, zero value otherwise.
func (o *GetWbethRateHistoryResponseRowsInner) GetAnnualPercentageRate() string {
	if o == nil || common.IsNil(o.AnnualPercentageRate) {
		var ret string
		return ret
	}
	return *o.AnnualPercentageRate
}

// GetAnnualPercentageRateOk returns a tuple with the AnnualPercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethRateHistoryResponseRowsInner) GetAnnualPercentageRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.AnnualPercentageRate) {
		return nil, false
	}
	return o.AnnualPercentageRate, true
}

// HasAnnualPercentageRate returns a boolean if a field has been set.
func (o *GetWbethRateHistoryResponseRowsInner) HasAnnualPercentageRate() bool {
	if o != nil && !common.IsNil(o.AnnualPercentageRate) {
		return true
	}

	return false
}

// SetAnnualPercentageRate gets a reference to the given string and assigns it to the AnnualPercentageRate field.
func (o *GetWbethRateHistoryResponseRowsInner) SetAnnualPercentageRate(v string) {
	o.AnnualPercentageRate = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *GetWbethRateHistoryResponseRowsInner) GetExchangeRate() string {
	if o == nil || common.IsNil(o.ExchangeRate) {
		var ret string
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethRateHistoryResponseRowsInner) GetExchangeRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *GetWbethRateHistoryResponseRowsInner) HasExchangeRate() bool {
	if o != nil && !common.IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given string and assigns it to the ExchangeRate field.
func (o *GetWbethRateHistoryResponseRowsInner) SetExchangeRate(v string) {
	o.ExchangeRate = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetWbethRateHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethRateHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetWbethRateHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetWbethRateHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

func (o GetWbethRateHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWbethRateHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AnnualPercentageRate) {
		toSerialize["annualPercentageRate"] = o.AnnualPercentageRate
	}
	if !common.IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetWbethRateHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetWbethRateHistoryResponseRowsInner := _GetWbethRateHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetWbethRateHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetWbethRateHistoryResponseRowsInner(varGetWbethRateHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "annualPercentageRate")
		delete(additionalProperties, "exchangeRate")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetWbethRateHistoryResponseRowsInner struct {
	value *GetWbethRateHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetWbethRateHistoryResponseRowsInner) Get() *GetWbethRateHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetWbethRateHistoryResponseRowsInner) Set(val *GetWbethRateHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWbethRateHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWbethRateHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWbethRateHistoryResponseRowsInner(val *GetWbethRateHistoryResponseRowsInner) *NullableGetWbethRateHistoryResponseRowsInner {
	return &NullableGetWbethRateHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetWbethRateHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWbethRateHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
