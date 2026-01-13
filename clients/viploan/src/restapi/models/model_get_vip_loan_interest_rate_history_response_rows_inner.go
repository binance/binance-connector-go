/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetVIPLoanInterestRateHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetVIPLoanInterestRateHistoryResponseRowsInner{}

// GetVIPLoanInterestRateHistoryResponseRowsInner struct for GetVIPLoanInterestRateHistoryResponseRowsInner
type GetVIPLoanInterestRateHistoryResponseRowsInner struct {
	Coin                   *string `json:"coin,omitempty"`
	AnnualizedInterestRate *string `json:"annualizedInterestRate,omitempty"`
	Time                   *int64  `json:"time,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _GetVIPLoanInterestRateHistoryResponseRowsInner GetVIPLoanInterestRateHistoryResponseRowsInner

// NewGetVIPLoanInterestRateHistoryResponseRowsInner instantiates a new GetVIPLoanInterestRateHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVIPLoanInterestRateHistoryResponseRowsInner() *GetVIPLoanInterestRateHistoryResponseRowsInner {
	this := GetVIPLoanInterestRateHistoryResponseRowsInner{}
	return &this
}

// NewGetVIPLoanInterestRateHistoryResponseRowsInnerWithDefaults instantiates a new GetVIPLoanInterestRateHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVIPLoanInterestRateHistoryResponseRowsInnerWithDefaults() *GetVIPLoanInterestRateHistoryResponseRowsInner {
	this := GetVIPLoanInterestRateHistoryResponseRowsInner{}
	return &this
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *GetVIPLoanInterestRateHistoryResponseRowsInner) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanInterestRateHistoryResponseRowsInner) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *GetVIPLoanInterestRateHistoryResponseRowsInner) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *GetVIPLoanInterestRateHistoryResponseRowsInner) SetCoin(v string) {
	o.Coin = &v
}

// GetAnnualizedInterestRate returns the AnnualizedInterestRate field value if set, zero value otherwise.
func (o *GetVIPLoanInterestRateHistoryResponseRowsInner) GetAnnualizedInterestRate() string {
	if o == nil || common.IsNil(o.AnnualizedInterestRate) {
		var ret string
		return ret
	}
	return *o.AnnualizedInterestRate
}

// GetAnnualizedInterestRateOk returns a tuple with the AnnualizedInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanInterestRateHistoryResponseRowsInner) GetAnnualizedInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.AnnualizedInterestRate) {
		return nil, false
	}
	return o.AnnualizedInterestRate, true
}

// HasAnnualizedInterestRate returns a boolean if a field has been set.
func (o *GetVIPLoanInterestRateHistoryResponseRowsInner) HasAnnualizedInterestRate() bool {
	if o != nil && !common.IsNil(o.AnnualizedInterestRate) {
		return true
	}

	return false
}

// SetAnnualizedInterestRate gets a reference to the given string and assigns it to the AnnualizedInterestRate field.
func (o *GetVIPLoanInterestRateHistoryResponseRowsInner) SetAnnualizedInterestRate(v string) {
	o.AnnualizedInterestRate = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetVIPLoanInterestRateHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanInterestRateHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetVIPLoanInterestRateHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetVIPLoanInterestRateHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

func (o GetVIPLoanInterestRateHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVIPLoanInterestRateHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
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

func (o *GetVIPLoanInterestRateHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetVIPLoanInterestRateHistoryResponseRowsInner := _GetVIPLoanInterestRateHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetVIPLoanInterestRateHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetVIPLoanInterestRateHistoryResponseRowsInner(varGetVIPLoanInterestRateHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "coin")
		delete(additionalProperties, "annualizedInterestRate")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetVIPLoanInterestRateHistoryResponseRowsInner struct {
	value *GetVIPLoanInterestRateHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetVIPLoanInterestRateHistoryResponseRowsInner) Get() *GetVIPLoanInterestRateHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetVIPLoanInterestRateHistoryResponseRowsInner) Set(val *GetVIPLoanInterestRateHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVIPLoanInterestRateHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVIPLoanInterestRateHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVIPLoanInterestRateHistoryResponseRowsInner(val *GetVIPLoanInterestRateHistoryResponseRowsInner) *NullableGetVIPLoanInterestRateHistoryResponseRowsInner {
	return &NullableGetVIPLoanInterestRateHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetVIPLoanInterestRateHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVIPLoanInterestRateHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
