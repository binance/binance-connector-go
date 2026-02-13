/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner{}

// QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner struct for QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner
type QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner struct {
	Asset                *string `json:"asset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	RepayTime            *int64  `json:"repayTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner

// NewQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner instantiates a new QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner() *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner {
	this := QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner{}
	return &this
}

// NewQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInnerWithDefaults instantiates a new QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInnerWithDefaults() *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner {
	this := QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetRepayTime returns the RepayTime field value if set, zero value otherwise.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) GetRepayTime() int64 {
	if o == nil || common.IsNil(o.RepayTime) {
		var ret int64
		return ret
	}
	return *o.RepayTime
}

// GetRepayTimeOk returns a tuple with the RepayTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) GetRepayTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RepayTime) {
		return nil, false
	}
	return o.RepayTime, true
}

// HasRepayTime returns a boolean if a field has been set.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) HasRepayTime() bool {
	if o != nil && !common.IsNil(o.RepayTime) {
		return true
	}

	return false
}

// SetRepayTime gets a reference to the given int64 and assigns it to the RepayTime field.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) SetRepayTime(v int64) {
	o.RepayTime = &v
}

func (o QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.RepayTime) {
		toSerialize["repayTime"] = o.RepayTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner := _QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner(varQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "repayTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner struct {
	value *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner
	isSet bool
}

func (v NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) Get() *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner {
	return v.value
}

func (v *NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) Set(val *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner(val *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) *NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner {
	return &NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
