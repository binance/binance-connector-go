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

// checks if the GetLoanRepaymentHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLoanRepaymentHistoryResponseRowsInner{}

// GetLoanRepaymentHistoryResponseRowsInner struct for GetLoanRepaymentHistoryResponseRowsInner
type GetLoanRepaymentHistoryResponseRowsInner struct {
	LoanCoin             *string `json:"loanCoin,omitempty"`
	RepayAmount          *string `json:"repayAmount,omitempty"`
	CollateralCoin       *string `json:"collateralCoin,omitempty"`
	CollateralUsed       *string `json:"collateralUsed,omitempty"`
	CollateralReturn     *string `json:"collateralReturn,omitempty"`
	RepayType            *string `json:"repayType,omitempty"`
	RepayStatus          *string `json:"repayStatus,omitempty"`
	RepayTime            *int64  `json:"repayTime,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLoanRepaymentHistoryResponseRowsInner GetLoanRepaymentHistoryResponseRowsInner

// NewGetLoanRepaymentHistoryResponseRowsInner instantiates a new GetLoanRepaymentHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLoanRepaymentHistoryResponseRowsInner() *GetLoanRepaymentHistoryResponseRowsInner {
	this := GetLoanRepaymentHistoryResponseRowsInner{}
	return &this
}

// NewGetLoanRepaymentHistoryResponseRowsInnerWithDefaults instantiates a new GetLoanRepaymentHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLoanRepaymentHistoryResponseRowsInnerWithDefaults() *GetLoanRepaymentHistoryResponseRowsInner {
	this := GetLoanRepaymentHistoryResponseRowsInner{}
	return &this
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *GetLoanRepaymentHistoryResponseRowsInner) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetRepayAmount returns the RepayAmount field value if set, zero value otherwise.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetRepayAmount() string {
	if o == nil || common.IsNil(o.RepayAmount) {
		var ret string
		return ret
	}
	return *o.RepayAmount
}

// GetRepayAmountOk returns a tuple with the RepayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetRepayAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.RepayAmount) {
		return nil, false
	}
	return o.RepayAmount, true
}

// HasRepayAmount returns a boolean if a field has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) HasRepayAmount() bool {
	if o != nil && !common.IsNil(o.RepayAmount) {
		return true
	}

	return false
}

// SetRepayAmount gets a reference to the given string and assigns it to the RepayAmount field.
func (o *GetLoanRepaymentHistoryResponseRowsInner) SetRepayAmount(v string) {
	o.RepayAmount = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *GetLoanRepaymentHistoryResponseRowsInner) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetCollateralUsed returns the CollateralUsed field value if set, zero value otherwise.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetCollateralUsed() string {
	if o == nil || common.IsNil(o.CollateralUsed) {
		var ret string
		return ret
	}
	return *o.CollateralUsed
}

// GetCollateralUsedOk returns a tuple with the CollateralUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetCollateralUsedOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralUsed) {
		return nil, false
	}
	return o.CollateralUsed, true
}

// HasCollateralUsed returns a boolean if a field has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) HasCollateralUsed() bool {
	if o != nil && !common.IsNil(o.CollateralUsed) {
		return true
	}

	return false
}

// SetCollateralUsed gets a reference to the given string and assigns it to the CollateralUsed field.
func (o *GetLoanRepaymentHistoryResponseRowsInner) SetCollateralUsed(v string) {
	o.CollateralUsed = &v
}

// GetCollateralReturn returns the CollateralReturn field value if set, zero value otherwise.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetCollateralReturn() string {
	if o == nil || common.IsNil(o.CollateralReturn) {
		var ret string
		return ret
	}
	return *o.CollateralReturn
}

// GetCollateralReturnOk returns a tuple with the CollateralReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetCollateralReturnOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralReturn) {
		return nil, false
	}
	return o.CollateralReturn, true
}

// HasCollateralReturn returns a boolean if a field has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) HasCollateralReturn() bool {
	if o != nil && !common.IsNil(o.CollateralReturn) {
		return true
	}

	return false
}

// SetCollateralReturn gets a reference to the given string and assigns it to the CollateralReturn field.
func (o *GetLoanRepaymentHistoryResponseRowsInner) SetCollateralReturn(v string) {
	o.CollateralReturn = &v
}

// GetRepayType returns the RepayType field value if set, zero value otherwise.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetRepayType() string {
	if o == nil || common.IsNil(o.RepayType) {
		var ret string
		return ret
	}
	return *o.RepayType
}

// GetRepayTypeOk returns a tuple with the RepayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetRepayTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RepayType) {
		return nil, false
	}
	return o.RepayType, true
}

// HasRepayType returns a boolean if a field has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) HasRepayType() bool {
	if o != nil && !common.IsNil(o.RepayType) {
		return true
	}

	return false
}

// SetRepayType gets a reference to the given string and assigns it to the RepayType field.
func (o *GetLoanRepaymentHistoryResponseRowsInner) SetRepayType(v string) {
	o.RepayType = &v
}

// GetRepayStatus returns the RepayStatus field value if set, zero value otherwise.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetRepayStatus() string {
	if o == nil || common.IsNil(o.RepayStatus) {
		var ret string
		return ret
	}
	return *o.RepayStatus
}

// GetRepayStatusOk returns a tuple with the RepayStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetRepayStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.RepayStatus) {
		return nil, false
	}
	return o.RepayStatus, true
}

// HasRepayStatus returns a boolean if a field has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) HasRepayStatus() bool {
	if o != nil && !common.IsNil(o.RepayStatus) {
		return true
	}

	return false
}

// SetRepayStatus gets a reference to the given string and assigns it to the RepayStatus field.
func (o *GetLoanRepaymentHistoryResponseRowsInner) SetRepayStatus(v string) {
	o.RepayStatus = &v
}

// GetRepayTime returns the RepayTime field value if set, zero value otherwise.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetRepayTime() int64 {
	if o == nil || common.IsNil(o.RepayTime) {
		var ret int64
		return ret
	}
	return *o.RepayTime
}

// GetRepayTimeOk returns a tuple with the RepayTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetRepayTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RepayTime) {
		return nil, false
	}
	return o.RepayTime, true
}

// HasRepayTime returns a boolean if a field has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) HasRepayTime() bool {
	if o != nil && !common.IsNil(o.RepayTime) {
		return true
	}

	return false
}

// SetRepayTime gets a reference to the given int64 and assigns it to the RepayTime field.
func (o *GetLoanRepaymentHistoryResponseRowsInner) SetRepayTime(v int64) {
	o.RepayTime = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetLoanRepaymentHistoryResponseRowsInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetLoanRepaymentHistoryResponseRowsInner) SetOrderId(v int64) {
	o.OrderId = &v
}

func (o GetLoanRepaymentHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLoanRepaymentHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.CollateralUsed) {
		toSerialize["collateralUsed"] = o.CollateralUsed
	}
	if !common.IsNil(o.CollateralReturn) {
		toSerialize["collateralReturn"] = o.CollateralReturn
	}
	if !common.IsNil(o.RepayType) {
		toSerialize["repayType"] = o.RepayType
	}
	if !common.IsNil(o.RepayStatus) {
		toSerialize["repayStatus"] = o.RepayStatus
	}
	if !common.IsNil(o.RepayTime) {
		toSerialize["repayTime"] = o.RepayTime
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetLoanRepaymentHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetLoanRepaymentHistoryResponseRowsInner := _GetLoanRepaymentHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetLoanRepaymentHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetLoanRepaymentHistoryResponseRowsInner(varGetLoanRepaymentHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "repayAmount")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "collateralUsed")
		delete(additionalProperties, "collateralReturn")
		delete(additionalProperties, "repayType")
		delete(additionalProperties, "repayStatus")
		delete(additionalProperties, "repayTime")
		delete(additionalProperties, "orderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLoanRepaymentHistoryResponseRowsInner struct {
	value *GetLoanRepaymentHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetLoanRepaymentHistoryResponseRowsInner) Get() *GetLoanRepaymentHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetLoanRepaymentHistoryResponseRowsInner) Set(val *GetLoanRepaymentHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLoanRepaymentHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLoanRepaymentHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLoanRepaymentHistoryResponseRowsInner(val *GetLoanRepaymentHistoryResponseRowsInner) *NullableGetLoanRepaymentHistoryResponseRowsInner {
	return &NullableGetLoanRepaymentHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetLoanRepaymentHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLoanRepaymentHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
