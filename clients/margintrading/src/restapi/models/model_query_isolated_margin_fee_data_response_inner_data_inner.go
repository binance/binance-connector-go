/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryIsolatedMarginFeeDataResponseInnerDataInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryIsolatedMarginFeeDataResponseInnerDataInner{}

// QueryIsolatedMarginFeeDataResponseInnerDataInner struct for QueryIsolatedMarginFeeDataResponseInnerDataInner
type QueryIsolatedMarginFeeDataResponseInnerDataInner struct {
	Coin                 *string `json:"coin,omitempty"`
	DailyInterest        *string `json:"dailyInterest,omitempty"`
	BorrowLimit          *string `json:"borrowLimit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryIsolatedMarginFeeDataResponseInnerDataInner QueryIsolatedMarginFeeDataResponseInnerDataInner

// NewQueryIsolatedMarginFeeDataResponseInnerDataInner instantiates a new QueryIsolatedMarginFeeDataResponseInnerDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryIsolatedMarginFeeDataResponseInnerDataInner() *QueryIsolatedMarginFeeDataResponseInnerDataInner {
	this := QueryIsolatedMarginFeeDataResponseInnerDataInner{}
	return &this
}

// NewQueryIsolatedMarginFeeDataResponseInnerDataInnerWithDefaults instantiates a new QueryIsolatedMarginFeeDataResponseInnerDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryIsolatedMarginFeeDataResponseInnerDataInnerWithDefaults() *QueryIsolatedMarginFeeDataResponseInnerDataInner {
	this := QueryIsolatedMarginFeeDataResponseInnerDataInner{}
	return &this
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *QueryIsolatedMarginFeeDataResponseInnerDataInner) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginFeeDataResponseInnerDataInner) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *QueryIsolatedMarginFeeDataResponseInnerDataInner) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *QueryIsolatedMarginFeeDataResponseInnerDataInner) SetCoin(v string) {
	o.Coin = &v
}

// GetDailyInterest returns the DailyInterest field value if set, zero value otherwise.
func (o *QueryIsolatedMarginFeeDataResponseInnerDataInner) GetDailyInterest() string {
	if o == nil || common.IsNil(o.DailyInterest) {
		var ret string
		return ret
	}
	return *o.DailyInterest
}

// GetDailyInterestOk returns a tuple with the DailyInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginFeeDataResponseInnerDataInner) GetDailyInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.DailyInterest) {
		return nil, false
	}
	return o.DailyInterest, true
}

// HasDailyInterest returns a boolean if a field has been set.
func (o *QueryIsolatedMarginFeeDataResponseInnerDataInner) HasDailyInterest() bool {
	if o != nil && !common.IsNil(o.DailyInterest) {
		return true
	}

	return false
}

// SetDailyInterest gets a reference to the given string and assigns it to the DailyInterest field.
func (o *QueryIsolatedMarginFeeDataResponseInnerDataInner) SetDailyInterest(v string) {
	o.DailyInterest = &v
}

// GetBorrowLimit returns the BorrowLimit field value if set, zero value otherwise.
func (o *QueryIsolatedMarginFeeDataResponseInnerDataInner) GetBorrowLimit() string {
	if o == nil || common.IsNil(o.BorrowLimit) {
		var ret string
		return ret
	}
	return *o.BorrowLimit
}

// GetBorrowLimitOk returns a tuple with the BorrowLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginFeeDataResponseInnerDataInner) GetBorrowLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.BorrowLimit) {
		return nil, false
	}
	return o.BorrowLimit, true
}

// HasBorrowLimit returns a boolean if a field has been set.
func (o *QueryIsolatedMarginFeeDataResponseInnerDataInner) HasBorrowLimit() bool {
	if o != nil && !common.IsNil(o.BorrowLimit) {
		return true
	}

	return false
}

// SetBorrowLimit gets a reference to the given string and assigns it to the BorrowLimit field.
func (o *QueryIsolatedMarginFeeDataResponseInnerDataInner) SetBorrowLimit(v string) {
	o.BorrowLimit = &v
}

func (o QueryIsolatedMarginFeeDataResponseInnerDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryIsolatedMarginFeeDataResponseInnerDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !common.IsNil(o.DailyInterest) {
		toSerialize["dailyInterest"] = o.DailyInterest
	}
	if !common.IsNil(o.BorrowLimit) {
		toSerialize["borrowLimit"] = o.BorrowLimit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryIsolatedMarginFeeDataResponseInnerDataInner) UnmarshalJSON(data []byte) (err error) {
	varQueryIsolatedMarginFeeDataResponseInnerDataInner := _QueryIsolatedMarginFeeDataResponseInnerDataInner{}

	err = json.Unmarshal(data, &varQueryIsolatedMarginFeeDataResponseInnerDataInner)

	if err != nil {
		return err
	}

	*o = QueryIsolatedMarginFeeDataResponseInnerDataInner(varQueryIsolatedMarginFeeDataResponseInnerDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "coin")
		delete(additionalProperties, "dailyInterest")
		delete(additionalProperties, "borrowLimit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryIsolatedMarginFeeDataResponseInnerDataInner struct {
	value *QueryIsolatedMarginFeeDataResponseInnerDataInner
	isSet bool
}

func (v NullableQueryIsolatedMarginFeeDataResponseInnerDataInner) Get() *QueryIsolatedMarginFeeDataResponseInnerDataInner {
	return v.value
}

func (v *NullableQueryIsolatedMarginFeeDataResponseInnerDataInner) Set(val *QueryIsolatedMarginFeeDataResponseInnerDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryIsolatedMarginFeeDataResponseInnerDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryIsolatedMarginFeeDataResponseInnerDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryIsolatedMarginFeeDataResponseInnerDataInner(val *QueryIsolatedMarginFeeDataResponseInnerDataInner) *NullableQueryIsolatedMarginFeeDataResponseInnerDataInner {
	return &NullableQueryIsolatedMarginFeeDataResponseInnerDataInner{value: val, isSet: true}
}

func (v NullableQueryIsolatedMarginFeeDataResponseInnerDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryIsolatedMarginFeeDataResponseInnerDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
