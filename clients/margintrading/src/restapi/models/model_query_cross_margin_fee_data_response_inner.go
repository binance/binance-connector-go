/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryCrossMarginFeeDataResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCrossMarginFeeDataResponseInner{}

// QueryCrossMarginFeeDataResponseInner struct for QueryCrossMarginFeeDataResponseInner
type QueryCrossMarginFeeDataResponseInner struct {
	VipLevel             *int64   `json:"vipLevel,omitempty"`
	Coin                 *string  `json:"coin,omitempty"`
	TransferIn           *bool    `json:"transferIn,omitempty"`
	Borrowable           *bool    `json:"borrowable,omitempty"`
	DailyInterest        *string  `json:"dailyInterest,omitempty"`
	YearlyInterest       *string  `json:"yearlyInterest,omitempty"`
	BorrowLimit          *string  `json:"borrowLimit,omitempty"`
	MarginablePairs      []string `json:"marginablePairs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryCrossMarginFeeDataResponseInner QueryCrossMarginFeeDataResponseInner

// NewQueryCrossMarginFeeDataResponseInner instantiates a new QueryCrossMarginFeeDataResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCrossMarginFeeDataResponseInner() *QueryCrossMarginFeeDataResponseInner {
	this := QueryCrossMarginFeeDataResponseInner{}
	return &this
}

// NewQueryCrossMarginFeeDataResponseInnerWithDefaults instantiates a new QueryCrossMarginFeeDataResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCrossMarginFeeDataResponseInnerWithDefaults() *QueryCrossMarginFeeDataResponseInner {
	this := QueryCrossMarginFeeDataResponseInner{}
	return &this
}

// GetVipLevel returns the VipLevel field value if set, zero value otherwise.
func (o *QueryCrossMarginFeeDataResponseInner) GetVipLevel() int64 {
	if o == nil || common.IsNil(o.VipLevel) {
		var ret int64
		return ret
	}
	return *o.VipLevel
}

// GetVipLevelOk returns a tuple with the VipLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginFeeDataResponseInner) GetVipLevelOk() (*int64, bool) {
	if o == nil || common.IsNil(o.VipLevel) {
		return nil, false
	}
	return o.VipLevel, true
}

// HasVipLevel returns a boolean if a field has been set.
func (o *QueryCrossMarginFeeDataResponseInner) HasVipLevel() bool {
	if o != nil && !common.IsNil(o.VipLevel) {
		return true
	}

	return false
}

// SetVipLevel gets a reference to the given int64 and assigns it to the VipLevel field.
func (o *QueryCrossMarginFeeDataResponseInner) SetVipLevel(v int64) {
	o.VipLevel = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *QueryCrossMarginFeeDataResponseInner) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginFeeDataResponseInner) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *QueryCrossMarginFeeDataResponseInner) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *QueryCrossMarginFeeDataResponseInner) SetCoin(v string) {
	o.Coin = &v
}

// GetTransferIn returns the TransferIn field value if set, zero value otherwise.
func (o *QueryCrossMarginFeeDataResponseInner) GetTransferIn() bool {
	if o == nil || common.IsNil(o.TransferIn) {
		var ret bool
		return ret
	}
	return *o.TransferIn
}

// GetTransferInOk returns a tuple with the TransferIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginFeeDataResponseInner) GetTransferInOk() (*bool, bool) {
	if o == nil || common.IsNil(o.TransferIn) {
		return nil, false
	}
	return o.TransferIn, true
}

// HasTransferIn returns a boolean if a field has been set.
func (o *QueryCrossMarginFeeDataResponseInner) HasTransferIn() bool {
	if o != nil && !common.IsNil(o.TransferIn) {
		return true
	}

	return false
}

// SetTransferIn gets a reference to the given bool and assigns it to the TransferIn field.
func (o *QueryCrossMarginFeeDataResponseInner) SetTransferIn(v bool) {
	o.TransferIn = &v
}

// GetBorrowable returns the Borrowable field value if set, zero value otherwise.
func (o *QueryCrossMarginFeeDataResponseInner) GetBorrowable() bool {
	if o == nil || common.IsNil(o.Borrowable) {
		var ret bool
		return ret
	}
	return *o.Borrowable
}

// GetBorrowableOk returns a tuple with the Borrowable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginFeeDataResponseInner) GetBorrowableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Borrowable) {
		return nil, false
	}
	return o.Borrowable, true
}

// HasBorrowable returns a boolean if a field has been set.
func (o *QueryCrossMarginFeeDataResponseInner) HasBorrowable() bool {
	if o != nil && !common.IsNil(o.Borrowable) {
		return true
	}

	return false
}

// SetBorrowable gets a reference to the given bool and assigns it to the Borrowable field.
func (o *QueryCrossMarginFeeDataResponseInner) SetBorrowable(v bool) {
	o.Borrowable = &v
}

// GetDailyInterest returns the DailyInterest field value if set, zero value otherwise.
func (o *QueryCrossMarginFeeDataResponseInner) GetDailyInterest() string {
	if o == nil || common.IsNil(o.DailyInterest) {
		var ret string
		return ret
	}
	return *o.DailyInterest
}

// GetDailyInterestOk returns a tuple with the DailyInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginFeeDataResponseInner) GetDailyInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.DailyInterest) {
		return nil, false
	}
	return o.DailyInterest, true
}

// HasDailyInterest returns a boolean if a field has been set.
func (o *QueryCrossMarginFeeDataResponseInner) HasDailyInterest() bool {
	if o != nil && !common.IsNil(o.DailyInterest) {
		return true
	}

	return false
}

// SetDailyInterest gets a reference to the given string and assigns it to the DailyInterest field.
func (o *QueryCrossMarginFeeDataResponseInner) SetDailyInterest(v string) {
	o.DailyInterest = &v
}

// GetYearlyInterest returns the YearlyInterest field value if set, zero value otherwise.
func (o *QueryCrossMarginFeeDataResponseInner) GetYearlyInterest() string {
	if o == nil || common.IsNil(o.YearlyInterest) {
		var ret string
		return ret
	}
	return *o.YearlyInterest
}

// GetYearlyInterestOk returns a tuple with the YearlyInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginFeeDataResponseInner) GetYearlyInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.YearlyInterest) {
		return nil, false
	}
	return o.YearlyInterest, true
}

// HasYearlyInterest returns a boolean if a field has been set.
func (o *QueryCrossMarginFeeDataResponseInner) HasYearlyInterest() bool {
	if o != nil && !common.IsNil(o.YearlyInterest) {
		return true
	}

	return false
}

// SetYearlyInterest gets a reference to the given string and assigns it to the YearlyInterest field.
func (o *QueryCrossMarginFeeDataResponseInner) SetYearlyInterest(v string) {
	o.YearlyInterest = &v
}

// GetBorrowLimit returns the BorrowLimit field value if set, zero value otherwise.
func (o *QueryCrossMarginFeeDataResponseInner) GetBorrowLimit() string {
	if o == nil || common.IsNil(o.BorrowLimit) {
		var ret string
		return ret
	}
	return *o.BorrowLimit
}

// GetBorrowLimitOk returns a tuple with the BorrowLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginFeeDataResponseInner) GetBorrowLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.BorrowLimit) {
		return nil, false
	}
	return o.BorrowLimit, true
}

// HasBorrowLimit returns a boolean if a field has been set.
func (o *QueryCrossMarginFeeDataResponseInner) HasBorrowLimit() bool {
	if o != nil && !common.IsNil(o.BorrowLimit) {
		return true
	}

	return false
}

// SetBorrowLimit gets a reference to the given string and assigns it to the BorrowLimit field.
func (o *QueryCrossMarginFeeDataResponseInner) SetBorrowLimit(v string) {
	o.BorrowLimit = &v
}

// GetMarginablePairs returns the MarginablePairs field value if set, zero value otherwise.
func (o *QueryCrossMarginFeeDataResponseInner) GetMarginablePairs() []string {
	if o == nil || common.IsNil(o.MarginablePairs) {
		var ret []string
		return ret
	}
	return o.MarginablePairs
}

// GetMarginablePairsOk returns a tuple with the MarginablePairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginFeeDataResponseInner) GetMarginablePairsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.MarginablePairs) {
		return nil, false
	}
	return o.MarginablePairs, true
}

// HasMarginablePairs returns a boolean if a field has been set.
func (o *QueryCrossMarginFeeDataResponseInner) HasMarginablePairs() bool {
	if o != nil && !common.IsNil(o.MarginablePairs) {
		return true
	}

	return false
}

// SetMarginablePairs gets a reference to the given []string and assigns it to the MarginablePairs field.
func (o *QueryCrossMarginFeeDataResponseInner) SetMarginablePairs(v []string) {
	o.MarginablePairs = v
}

func (o QueryCrossMarginFeeDataResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCrossMarginFeeDataResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.VipLevel) {
		toSerialize["vipLevel"] = o.VipLevel
	}
	if !common.IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !common.IsNil(o.TransferIn) {
		toSerialize["transferIn"] = o.TransferIn
	}
	if !common.IsNil(o.Borrowable) {
		toSerialize["borrowable"] = o.Borrowable
	}
	if !common.IsNil(o.DailyInterest) {
		toSerialize["dailyInterest"] = o.DailyInterest
	}
	if !common.IsNil(o.YearlyInterest) {
		toSerialize["yearlyInterest"] = o.YearlyInterest
	}
	if !common.IsNil(o.BorrowLimit) {
		toSerialize["borrowLimit"] = o.BorrowLimit
	}
	if !common.IsNil(o.MarginablePairs) {
		toSerialize["marginablePairs"] = o.MarginablePairs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryCrossMarginFeeDataResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryCrossMarginFeeDataResponseInner := _QueryCrossMarginFeeDataResponseInner{}

	err = json.Unmarshal(data, &varQueryCrossMarginFeeDataResponseInner)

	if err != nil {
		return err
	}

	*o = QueryCrossMarginFeeDataResponseInner(varQueryCrossMarginFeeDataResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vipLevel")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "transferIn")
		delete(additionalProperties, "borrowable")
		delete(additionalProperties, "dailyInterest")
		delete(additionalProperties, "yearlyInterest")
		delete(additionalProperties, "borrowLimit")
		delete(additionalProperties, "marginablePairs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryCrossMarginFeeDataResponseInner struct {
	value *QueryCrossMarginFeeDataResponseInner
	isSet bool
}

func (v NullableQueryCrossMarginFeeDataResponseInner) Get() *QueryCrossMarginFeeDataResponseInner {
	return v.value
}

func (v *NullableQueryCrossMarginFeeDataResponseInner) Set(val *QueryCrossMarginFeeDataResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCrossMarginFeeDataResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCrossMarginFeeDataResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCrossMarginFeeDataResponseInner(val *QueryCrossMarginFeeDataResponseInner) *NullableQueryCrossMarginFeeDataResponseInner {
	return &NullableQueryCrossMarginFeeDataResponseInner{value: val, isSet: true}
}

func (v NullableQueryCrossMarginFeeDataResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCrossMarginFeeDataResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
