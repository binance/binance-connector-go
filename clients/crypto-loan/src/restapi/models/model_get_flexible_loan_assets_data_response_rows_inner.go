/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFlexibleLoanAssetsDataResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleLoanAssetsDataResponseRowsInner{}

// GetFlexibleLoanAssetsDataResponseRowsInner struct for GetFlexibleLoanAssetsDataResponseRowsInner
type GetFlexibleLoanAssetsDataResponseRowsInner struct {
	LoanCoin             *string `json:"loanCoin,omitempty"`
	FlexibleInterestRate *string `json:"flexibleInterestRate,omitempty"`
	FlexibleMinLimit     *string `json:"flexibleMinLimit,omitempty"`
	FlexibleMaxLimit     *string `json:"flexibleMaxLimit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleLoanAssetsDataResponseRowsInner GetFlexibleLoanAssetsDataResponseRowsInner

// NewGetFlexibleLoanAssetsDataResponseRowsInner instantiates a new GetFlexibleLoanAssetsDataResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleLoanAssetsDataResponseRowsInner() *GetFlexibleLoanAssetsDataResponseRowsInner {
	this := GetFlexibleLoanAssetsDataResponseRowsInner{}
	return &this
}

// NewGetFlexibleLoanAssetsDataResponseRowsInnerWithDefaults instantiates a new GetFlexibleLoanAssetsDataResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleLoanAssetsDataResponseRowsInnerWithDefaults() *GetFlexibleLoanAssetsDataResponseRowsInner {
	this := GetFlexibleLoanAssetsDataResponseRowsInner{}
	return &this
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *GetFlexibleLoanAssetsDataResponseRowsInner) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanAssetsDataResponseRowsInner) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *GetFlexibleLoanAssetsDataResponseRowsInner) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *GetFlexibleLoanAssetsDataResponseRowsInner) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetFlexibleInterestRate returns the FlexibleInterestRate field value if set, zero value otherwise.
func (o *GetFlexibleLoanAssetsDataResponseRowsInner) GetFlexibleInterestRate() string {
	if o == nil || common.IsNil(o.FlexibleInterestRate) {
		var ret string
		return ret
	}
	return *o.FlexibleInterestRate
}

// GetFlexibleInterestRateOk returns a tuple with the FlexibleInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanAssetsDataResponseRowsInner) GetFlexibleInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.FlexibleInterestRate) {
		return nil, false
	}
	return o.FlexibleInterestRate, true
}

// HasFlexibleInterestRate returns a boolean if a field has been set.
func (o *GetFlexibleLoanAssetsDataResponseRowsInner) HasFlexibleInterestRate() bool {
	if o != nil && !common.IsNil(o.FlexibleInterestRate) {
		return true
	}

	return false
}

// SetFlexibleInterestRate gets a reference to the given string and assigns it to the FlexibleInterestRate field.
func (o *GetFlexibleLoanAssetsDataResponseRowsInner) SetFlexibleInterestRate(v string) {
	o.FlexibleInterestRate = &v
}

// GetFlexibleMinLimit returns the FlexibleMinLimit field value if set, zero value otherwise.
func (o *GetFlexibleLoanAssetsDataResponseRowsInner) GetFlexibleMinLimit() string {
	if o == nil || common.IsNil(o.FlexibleMinLimit) {
		var ret string
		return ret
	}
	return *o.FlexibleMinLimit
}

// GetFlexibleMinLimitOk returns a tuple with the FlexibleMinLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanAssetsDataResponseRowsInner) GetFlexibleMinLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.FlexibleMinLimit) {
		return nil, false
	}
	return o.FlexibleMinLimit, true
}

// HasFlexibleMinLimit returns a boolean if a field has been set.
func (o *GetFlexibleLoanAssetsDataResponseRowsInner) HasFlexibleMinLimit() bool {
	if o != nil && !common.IsNil(o.FlexibleMinLimit) {
		return true
	}

	return false
}

// SetFlexibleMinLimit gets a reference to the given string and assigns it to the FlexibleMinLimit field.
func (o *GetFlexibleLoanAssetsDataResponseRowsInner) SetFlexibleMinLimit(v string) {
	o.FlexibleMinLimit = &v
}

// GetFlexibleMaxLimit returns the FlexibleMaxLimit field value if set, zero value otherwise.
func (o *GetFlexibleLoanAssetsDataResponseRowsInner) GetFlexibleMaxLimit() string {
	if o == nil || common.IsNil(o.FlexibleMaxLimit) {
		var ret string
		return ret
	}
	return *o.FlexibleMaxLimit
}

// GetFlexibleMaxLimitOk returns a tuple with the FlexibleMaxLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanAssetsDataResponseRowsInner) GetFlexibleMaxLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.FlexibleMaxLimit) {
		return nil, false
	}
	return o.FlexibleMaxLimit, true
}

// HasFlexibleMaxLimit returns a boolean if a field has been set.
func (o *GetFlexibleLoanAssetsDataResponseRowsInner) HasFlexibleMaxLimit() bool {
	if o != nil && !common.IsNil(o.FlexibleMaxLimit) {
		return true
	}

	return false
}

// SetFlexibleMaxLimit gets a reference to the given string and assigns it to the FlexibleMaxLimit field.
func (o *GetFlexibleLoanAssetsDataResponseRowsInner) SetFlexibleMaxLimit(v string) {
	o.FlexibleMaxLimit = &v
}

func (o GetFlexibleLoanAssetsDataResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleLoanAssetsDataResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.FlexibleInterestRate) {
		toSerialize["flexibleInterestRate"] = o.FlexibleInterestRate
	}
	if !common.IsNil(o.FlexibleMinLimit) {
		toSerialize["flexibleMinLimit"] = o.FlexibleMinLimit
	}
	if !common.IsNil(o.FlexibleMaxLimit) {
		toSerialize["flexibleMaxLimit"] = o.FlexibleMaxLimit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFlexibleLoanAssetsDataResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleLoanAssetsDataResponseRowsInner := _GetFlexibleLoanAssetsDataResponseRowsInner{}

	err = json.Unmarshal(data, &varGetFlexibleLoanAssetsDataResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetFlexibleLoanAssetsDataResponseRowsInner(varGetFlexibleLoanAssetsDataResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "flexibleInterestRate")
		delete(additionalProperties, "flexibleMinLimit")
		delete(additionalProperties, "flexibleMaxLimit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleLoanAssetsDataResponseRowsInner struct {
	value *GetFlexibleLoanAssetsDataResponseRowsInner
	isSet bool
}

func (v NullableGetFlexibleLoanAssetsDataResponseRowsInner) Get() *GetFlexibleLoanAssetsDataResponseRowsInner {
	return v.value
}

func (v *NullableGetFlexibleLoanAssetsDataResponseRowsInner) Set(val *GetFlexibleLoanAssetsDataResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleLoanAssetsDataResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleLoanAssetsDataResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleLoanAssetsDataResponseRowsInner(val *GetFlexibleLoanAssetsDataResponseRowsInner) *NullableGetFlexibleLoanAssetsDataResponseRowsInner {
	return &NullableGetFlexibleLoanAssetsDataResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetFlexibleLoanAssetsDataResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleLoanAssetsDataResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
