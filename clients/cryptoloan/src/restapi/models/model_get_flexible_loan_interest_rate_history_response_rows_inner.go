/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFlexibleLoanInterestRateHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleLoanInterestRateHistoryResponseRowsInner{}

// GetFlexibleLoanInterestRateHistoryResponseRowsInner struct for GetFlexibleLoanInterestRateHistoryResponseRowsInner
type GetFlexibleLoanInterestRateHistoryResponseRowsInner struct {
	Coin                   *string `json:"coin,omitempty"`
	AnnualizedInterestRate *string `json:"annualizedInterestRate,omitempty"`
	Time                   *int64  `json:"time,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _GetFlexibleLoanInterestRateHistoryResponseRowsInner GetFlexibleLoanInterestRateHistoryResponseRowsInner

// NewGetFlexibleLoanInterestRateHistoryResponseRowsInner instantiates a new GetFlexibleLoanInterestRateHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleLoanInterestRateHistoryResponseRowsInner() *GetFlexibleLoanInterestRateHistoryResponseRowsInner {
	this := GetFlexibleLoanInterestRateHistoryResponseRowsInner{}
	return &this
}

// NewGetFlexibleLoanInterestRateHistoryResponseRowsInnerWithDefaults instantiates a new GetFlexibleLoanInterestRateHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleLoanInterestRateHistoryResponseRowsInnerWithDefaults() *GetFlexibleLoanInterestRateHistoryResponseRowsInner {
	this := GetFlexibleLoanInterestRateHistoryResponseRowsInner{}
	return &this
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *GetFlexibleLoanInterestRateHistoryResponseRowsInner) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanInterestRateHistoryResponseRowsInner) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *GetFlexibleLoanInterestRateHistoryResponseRowsInner) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *GetFlexibleLoanInterestRateHistoryResponseRowsInner) SetCoin(v string) {
	o.Coin = &v
}

// GetAnnualizedInterestRate returns the AnnualizedInterestRate field value if set, zero value otherwise.
func (o *GetFlexibleLoanInterestRateHistoryResponseRowsInner) GetAnnualizedInterestRate() string {
	if o == nil || common.IsNil(o.AnnualizedInterestRate) {
		var ret string
		return ret
	}
	return *o.AnnualizedInterestRate
}

// GetAnnualizedInterestRateOk returns a tuple with the AnnualizedInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanInterestRateHistoryResponseRowsInner) GetAnnualizedInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.AnnualizedInterestRate) {
		return nil, false
	}
	return o.AnnualizedInterestRate, true
}

// HasAnnualizedInterestRate returns a boolean if a field has been set.
func (o *GetFlexibleLoanInterestRateHistoryResponseRowsInner) HasAnnualizedInterestRate() bool {
	if o != nil && !common.IsNil(o.AnnualizedInterestRate) {
		return true
	}

	return false
}

// SetAnnualizedInterestRate gets a reference to the given string and assigns it to the AnnualizedInterestRate field.
func (o *GetFlexibleLoanInterestRateHistoryResponseRowsInner) SetAnnualizedInterestRate(v string) {
	o.AnnualizedInterestRate = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetFlexibleLoanInterestRateHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanInterestRateHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetFlexibleLoanInterestRateHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetFlexibleLoanInterestRateHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

func (o GetFlexibleLoanInterestRateHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleLoanInterestRateHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !common.IsNil(o.AnnualizedInterestRate) {
		toSerialize["annualizedInterestRate"] = o.AnnualizedInterestRate
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFlexibleLoanInterestRateHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleLoanInterestRateHistoryResponseRowsInner := _GetFlexibleLoanInterestRateHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetFlexibleLoanInterestRateHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetFlexibleLoanInterestRateHistoryResponseRowsInner(varGetFlexibleLoanInterestRateHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "coin")
		delete(additionalProperties, "annualizedInterestRate")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleLoanInterestRateHistoryResponseRowsInner struct {
	value *GetFlexibleLoanInterestRateHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetFlexibleLoanInterestRateHistoryResponseRowsInner) Get() *GetFlexibleLoanInterestRateHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetFlexibleLoanInterestRateHistoryResponseRowsInner) Set(val *GetFlexibleLoanInterestRateHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleLoanInterestRateHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleLoanInterestRateHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleLoanInterestRateHistoryResponseRowsInner(val *GetFlexibleLoanInterestRateHistoryResponseRowsInner) *NullableGetFlexibleLoanInterestRateHistoryResponseRowsInner {
	return &NullableGetFlexibleLoanInterestRateHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetFlexibleLoanInterestRateHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleLoanInterestRateHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
