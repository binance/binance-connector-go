/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFlexibleLoanRepaymentHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleLoanRepaymentHistoryResponseRowsInner{}

// GetFlexibleLoanRepaymentHistoryResponseRowsInner struct for GetFlexibleLoanRepaymentHistoryResponseRowsInner
type GetFlexibleLoanRepaymentHistoryResponseRowsInner struct {
	LoanCoin             *string `json:"loanCoin,omitempty"`
	RepayAmount          *string `json:"repayAmount,omitempty"`
	CollateralCoin       *string `json:"collateralCoin,omitempty"`
	CollateralReturn     *string `json:"collateralReturn,omitempty"`
	RepayStatus          *string `json:"repayStatus,omitempty"`
	RepayTime            *int64  `json:"repayTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleLoanRepaymentHistoryResponseRowsInner GetFlexibleLoanRepaymentHistoryResponseRowsInner

// NewGetFlexibleLoanRepaymentHistoryResponseRowsInner instantiates a new GetFlexibleLoanRepaymentHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleLoanRepaymentHistoryResponseRowsInner() *GetFlexibleLoanRepaymentHistoryResponseRowsInner {
	this := GetFlexibleLoanRepaymentHistoryResponseRowsInner{}
	return &this
}

// NewGetFlexibleLoanRepaymentHistoryResponseRowsInnerWithDefaults instantiates a new GetFlexibleLoanRepaymentHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleLoanRepaymentHistoryResponseRowsInnerWithDefaults() *GetFlexibleLoanRepaymentHistoryResponseRowsInner {
	this := GetFlexibleLoanRepaymentHistoryResponseRowsInner{}
	return &this
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetRepayAmount returns the RepayAmount field value if set, zero value otherwise.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) GetRepayAmount() string {
	if o == nil || common.IsNil(o.RepayAmount) {
		var ret string
		return ret
	}
	return *o.RepayAmount
}

// GetRepayAmountOk returns a tuple with the RepayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) GetRepayAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.RepayAmount) {
		return nil, false
	}
	return o.RepayAmount, true
}

// HasRepayAmount returns a boolean if a field has been set.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) HasRepayAmount() bool {
	if o != nil && !common.IsNil(o.RepayAmount) {
		return true
	}

	return false
}

// SetRepayAmount gets a reference to the given string and assigns it to the RepayAmount field.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) SetRepayAmount(v string) {
	o.RepayAmount = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetCollateralReturn returns the CollateralReturn field value if set, zero value otherwise.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) GetCollateralReturn() string {
	if o == nil || common.IsNil(o.CollateralReturn) {
		var ret string
		return ret
	}
	return *o.CollateralReturn
}

// GetCollateralReturnOk returns a tuple with the CollateralReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) GetCollateralReturnOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralReturn) {
		return nil, false
	}
	return o.CollateralReturn, true
}

// HasCollateralReturn returns a boolean if a field has been set.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) HasCollateralReturn() bool {
	if o != nil && !common.IsNil(o.CollateralReturn) {
		return true
	}

	return false
}

// SetCollateralReturn gets a reference to the given string and assigns it to the CollateralReturn field.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) SetCollateralReturn(v string) {
	o.CollateralReturn = &v
}

// GetRepayStatus returns the RepayStatus field value if set, zero value otherwise.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) GetRepayStatus() string {
	if o == nil || common.IsNil(o.RepayStatus) {
		var ret string
		return ret
	}
	return *o.RepayStatus
}

// GetRepayStatusOk returns a tuple with the RepayStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) GetRepayStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.RepayStatus) {
		return nil, false
	}
	return o.RepayStatus, true
}

// HasRepayStatus returns a boolean if a field has been set.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) HasRepayStatus() bool {
	if o != nil && !common.IsNil(o.RepayStatus) {
		return true
	}

	return false
}

// SetRepayStatus gets a reference to the given string and assigns it to the RepayStatus field.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) SetRepayStatus(v string) {
	o.RepayStatus = &v
}

// GetRepayTime returns the RepayTime field value if set, zero value otherwise.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) GetRepayTime() int64 {
	if o == nil || common.IsNil(o.RepayTime) {
		var ret int64
		return ret
	}
	return *o.RepayTime
}

// GetRepayTimeOk returns a tuple with the RepayTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) GetRepayTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RepayTime) {
		return nil, false
	}
	return o.RepayTime, true
}

// HasRepayTime returns a boolean if a field has been set.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) HasRepayTime() bool {
	if o != nil && !common.IsNil(o.RepayTime) {
		return true
	}

	return false
}

// SetRepayTime gets a reference to the given int64 and assigns it to the RepayTime field.
func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) SetRepayTime(v int64) {
	o.RepayTime = &v
}

func (o GetFlexibleLoanRepaymentHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleLoanRepaymentHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.RepayAmount) {
		toSerialize["repayAmount"] = o.RepayAmount
	}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !common.IsNil(o.CollateralReturn) {
		toSerialize["collateralReturn"] = o.CollateralReturn
	}
	if !common.IsNil(o.RepayStatus) {
		toSerialize["repayStatus"] = o.RepayStatus
	}
	if !common.IsNil(o.RepayTime) {
		toSerialize["repayTime"] = o.RepayTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFlexibleLoanRepaymentHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleLoanRepaymentHistoryResponseRowsInner := _GetFlexibleLoanRepaymentHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetFlexibleLoanRepaymentHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetFlexibleLoanRepaymentHistoryResponseRowsInner(varGetFlexibleLoanRepaymentHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "repayAmount")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "collateralReturn")
		delete(additionalProperties, "repayStatus")
		delete(additionalProperties, "repayTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleLoanRepaymentHistoryResponseRowsInner struct {
	value *GetFlexibleLoanRepaymentHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetFlexibleLoanRepaymentHistoryResponseRowsInner) Get() *GetFlexibleLoanRepaymentHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetFlexibleLoanRepaymentHistoryResponseRowsInner) Set(val *GetFlexibleLoanRepaymentHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleLoanRepaymentHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleLoanRepaymentHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleLoanRepaymentHistoryResponseRowsInner(val *GetFlexibleLoanRepaymentHistoryResponseRowsInner) *NullableGetFlexibleLoanRepaymentHistoryResponseRowsInner {
	return &NullableGetFlexibleLoanRepaymentHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetFlexibleLoanRepaymentHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleLoanRepaymentHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
