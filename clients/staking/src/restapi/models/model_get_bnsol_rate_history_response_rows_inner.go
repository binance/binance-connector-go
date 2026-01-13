/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetBnsolRateHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBnsolRateHistoryResponseRowsInner{}

// GetBnsolRateHistoryResponseRowsInner struct for GetBnsolRateHistoryResponseRowsInner
type GetBnsolRateHistoryResponseRowsInner struct {
	AnnualPercentageRate *string                                                 `json:"annualPercentageRate,omitempty"`
	ExchangeRate         *string                                                 `json:"exchangeRate,omitempty"`
	BoostRewards         []GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner `json:"boostRewards,omitempty"`
	Time                 *int64                                                  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBnsolRateHistoryResponseRowsInner GetBnsolRateHistoryResponseRowsInner

// NewGetBnsolRateHistoryResponseRowsInner instantiates a new GetBnsolRateHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBnsolRateHistoryResponseRowsInner() *GetBnsolRateHistoryResponseRowsInner {
	this := GetBnsolRateHistoryResponseRowsInner{}
	return &this
}

// NewGetBnsolRateHistoryResponseRowsInnerWithDefaults instantiates a new GetBnsolRateHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBnsolRateHistoryResponseRowsInnerWithDefaults() *GetBnsolRateHistoryResponseRowsInner {
	this := GetBnsolRateHistoryResponseRowsInner{}
	return &this
}

// GetAnnualPercentageRate returns the AnnualPercentageRate field value if set, zero value otherwise.
func (o *GetBnsolRateHistoryResponseRowsInner) GetAnnualPercentageRate() string {
	if o == nil || common.IsNil(o.AnnualPercentageRate) {
		var ret string
		return ret
	}
	return *o.AnnualPercentageRate
}

// GetAnnualPercentageRateOk returns a tuple with the AnnualPercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnsolRateHistoryResponseRowsInner) GetAnnualPercentageRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.AnnualPercentageRate) {
		return nil, false
	}
	return o.AnnualPercentageRate, true
}

// HasAnnualPercentageRate returns a boolean if a field has been set.
func (o *GetBnsolRateHistoryResponseRowsInner) HasAnnualPercentageRate() bool {
	if o != nil && !common.IsNil(o.AnnualPercentageRate) {
		return true
	}

	return false
}

// SetAnnualPercentageRate gets a reference to the given string and assigns it to the AnnualPercentageRate field.
func (o *GetBnsolRateHistoryResponseRowsInner) SetAnnualPercentageRate(v string) {
	o.AnnualPercentageRate = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *GetBnsolRateHistoryResponseRowsInner) GetExchangeRate() string {
	if o == nil || common.IsNil(o.ExchangeRate) {
		var ret string
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnsolRateHistoryResponseRowsInner) GetExchangeRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *GetBnsolRateHistoryResponseRowsInner) HasExchangeRate() bool {
	if o != nil && !common.IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given string and assigns it to the ExchangeRate field.
func (o *GetBnsolRateHistoryResponseRowsInner) SetExchangeRate(v string) {
	o.ExchangeRate = &v
}

// GetBoostRewards returns the BoostRewards field value if set, zero value otherwise.
func (o *GetBnsolRateHistoryResponseRowsInner) GetBoostRewards() []GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner {
	if o == nil || common.IsNil(o.BoostRewards) {
		var ret []GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner
		return ret
	}
	return o.BoostRewards
}

// GetBoostRewardsOk returns a tuple with the BoostRewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnsolRateHistoryResponseRowsInner) GetBoostRewardsOk() ([]GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner, bool) {
	if o == nil || common.IsNil(o.BoostRewards) {
		return nil, false
	}
	return o.BoostRewards, true
}

// HasBoostRewards returns a boolean if a field has been set.
func (o *GetBnsolRateHistoryResponseRowsInner) HasBoostRewards() bool {
	if o != nil && !common.IsNil(o.BoostRewards) {
		return true
	}

	return false
}

// SetBoostRewards gets a reference to the given []GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner and assigns it to the BoostRewards field.
func (o *GetBnsolRateHistoryResponseRowsInner) SetBoostRewards(v []GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) {
	o.BoostRewards = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetBnsolRateHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnsolRateHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetBnsolRateHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetBnsolRateHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

func (o GetBnsolRateHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBnsolRateHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AnnualPercentageRate) {
		toSerialize["annualPercentageRate"] = o.AnnualPercentageRate
	}
	if !common.IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	if !common.IsNil(o.BoostRewards) {
		toSerialize["boostRewards"] = o.BoostRewards
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetBnsolRateHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetBnsolRateHistoryResponseRowsInner := _GetBnsolRateHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetBnsolRateHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetBnsolRateHistoryResponseRowsInner(varGetBnsolRateHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "annualPercentageRate")
		delete(additionalProperties, "exchangeRate")
		delete(additionalProperties, "boostRewards")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBnsolRateHistoryResponseRowsInner struct {
	value *GetBnsolRateHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetBnsolRateHistoryResponseRowsInner) Get() *GetBnsolRateHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetBnsolRateHistoryResponseRowsInner) Set(val *GetBnsolRateHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBnsolRateHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBnsolRateHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBnsolRateHistoryResponseRowsInner(val *GetBnsolRateHistoryResponseRowsInner) *NullableGetBnsolRateHistoryResponseRowsInner {
	return &NullableGetBnsolRateHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetBnsolRateHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBnsolRateHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
