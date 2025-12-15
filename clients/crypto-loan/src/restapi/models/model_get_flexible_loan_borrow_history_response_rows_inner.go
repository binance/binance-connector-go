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

// checks if the GetFlexibleLoanBorrowHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleLoanBorrowHistoryResponseRowsInner{}

// GetFlexibleLoanBorrowHistoryResponseRowsInner struct for GetFlexibleLoanBorrowHistoryResponseRowsInner
type GetFlexibleLoanBorrowHistoryResponseRowsInner struct {
	LoanCoin                *string `json:"loanCoin,omitempty"`
	InitialLoanAmount       *string `json:"initialLoanAmount,omitempty"`
	CollateralCoin          *string `json:"collateralCoin,omitempty"`
	InitialCollateralAmount *string `json:"initialCollateralAmount,omitempty"`
	BorrowTime              *int64  `json:"borrowTime,omitempty"`
	Status                  *string `json:"status,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _GetFlexibleLoanBorrowHistoryResponseRowsInner GetFlexibleLoanBorrowHistoryResponseRowsInner

// NewGetFlexibleLoanBorrowHistoryResponseRowsInner instantiates a new GetFlexibleLoanBorrowHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleLoanBorrowHistoryResponseRowsInner() *GetFlexibleLoanBorrowHistoryResponseRowsInner {
	this := GetFlexibleLoanBorrowHistoryResponseRowsInner{}
	return &this
}

// NewGetFlexibleLoanBorrowHistoryResponseRowsInnerWithDefaults instantiates a new GetFlexibleLoanBorrowHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleLoanBorrowHistoryResponseRowsInnerWithDefaults() *GetFlexibleLoanBorrowHistoryResponseRowsInner {
	this := GetFlexibleLoanBorrowHistoryResponseRowsInner{}
	return &this
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetInitialLoanAmount returns the InitialLoanAmount field value if set, zero value otherwise.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) GetInitialLoanAmount() string {
	if o == nil || common.IsNil(o.InitialLoanAmount) {
		var ret string
		return ret
	}
	return *o.InitialLoanAmount
}

// GetInitialLoanAmountOk returns a tuple with the InitialLoanAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) GetInitialLoanAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialLoanAmount) {
		return nil, false
	}
	return o.InitialLoanAmount, true
}

// HasInitialLoanAmount returns a boolean if a field has been set.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) HasInitialLoanAmount() bool {
	if o != nil && !common.IsNil(o.InitialLoanAmount) {
		return true
	}

	return false
}

// SetInitialLoanAmount gets a reference to the given string and assigns it to the InitialLoanAmount field.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) SetInitialLoanAmount(v string) {
	o.InitialLoanAmount = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetInitialCollateralAmount returns the InitialCollateralAmount field value if set, zero value otherwise.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) GetInitialCollateralAmount() string {
	if o == nil || common.IsNil(o.InitialCollateralAmount) {
		var ret string
		return ret
	}
	return *o.InitialCollateralAmount
}

// GetInitialCollateralAmountOk returns a tuple with the InitialCollateralAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) GetInitialCollateralAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialCollateralAmount) {
		return nil, false
	}
	return o.InitialCollateralAmount, true
}

// HasInitialCollateralAmount returns a boolean if a field has been set.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) HasInitialCollateralAmount() bool {
	if o != nil && !common.IsNil(o.InitialCollateralAmount) {
		return true
	}

	return false
}

// SetInitialCollateralAmount gets a reference to the given string and assigns it to the InitialCollateralAmount field.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) SetInitialCollateralAmount(v string) {
	o.InitialCollateralAmount = &v
}

// GetBorrowTime returns the BorrowTime field value if set, zero value otherwise.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) GetBorrowTime() int64 {
	if o == nil || common.IsNil(o.BorrowTime) {
		var ret int64
		return ret
	}
	return *o.BorrowTime
}

// GetBorrowTimeOk returns a tuple with the BorrowTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) GetBorrowTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BorrowTime) {
		return nil, false
	}
	return o.BorrowTime, true
}

// HasBorrowTime returns a boolean if a field has been set.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) HasBorrowTime() bool {
	if o != nil && !common.IsNil(o.BorrowTime) {
		return true
	}

	return false
}

// SetBorrowTime gets a reference to the given int64 and assigns it to the BorrowTime field.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) SetBorrowTime(v int64) {
	o.BorrowTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetFlexibleLoanBorrowHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleLoanBorrowHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.InitialLoanAmount) {
		toSerialize["initialLoanAmount"] = o.InitialLoanAmount
	}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !common.IsNil(o.InitialCollateralAmount) {
		toSerialize["initialCollateralAmount"] = o.InitialCollateralAmount
	}
	if !common.IsNil(o.BorrowTime) {
		toSerialize["borrowTime"] = o.BorrowTime
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFlexibleLoanBorrowHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleLoanBorrowHistoryResponseRowsInner := _GetFlexibleLoanBorrowHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetFlexibleLoanBorrowHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetFlexibleLoanBorrowHistoryResponseRowsInner(varGetFlexibleLoanBorrowHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "initialLoanAmount")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "initialCollateralAmount")
		delete(additionalProperties, "borrowTime")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleLoanBorrowHistoryResponseRowsInner struct {
	value *GetFlexibleLoanBorrowHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetFlexibleLoanBorrowHistoryResponseRowsInner) Get() *GetFlexibleLoanBorrowHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetFlexibleLoanBorrowHistoryResponseRowsInner) Set(val *GetFlexibleLoanBorrowHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleLoanBorrowHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleLoanBorrowHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleLoanBorrowHistoryResponseRowsInner(val *GetFlexibleLoanBorrowHistoryResponseRowsInner) *NullableGetFlexibleLoanBorrowHistoryResponseRowsInner {
	return &NullableGetFlexibleLoanBorrowHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetFlexibleLoanBorrowHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleLoanBorrowHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
