/*
VIP Loan REST API

Access over-collateralized loan services, manage positions, and monitor collateral via the VIP Loan API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetVIPLoanRepaymentHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetVIPLoanRepaymentHistoryResponseRowsInner{}

// GetVIPLoanRepaymentHistoryResponseRowsInner struct for GetVIPLoanRepaymentHistoryResponseRowsInner
type GetVIPLoanRepaymentHistoryResponseRowsInner struct {
	LoanCoin       *string `json:"loanCoin,omitempty"`
	RepayAmount    *string `json:"repayAmount,omitempty"`
	CollateralCoin *string `json:"collateralCoin,omitempty"`
	// Repayment status (`Repaid`, `Repaying`, `Failed`).
	RepayStatus          *string `json:"repayStatus,omitempty"`
	LoanDate             *string `json:"loanDate,omitempty"`
	RepayTime            *string `json:"repayTime,omitempty"`
	OrderId              *string `json:"orderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetVIPLoanRepaymentHistoryResponseRowsInner GetVIPLoanRepaymentHistoryResponseRowsInner

// NewGetVIPLoanRepaymentHistoryResponseRowsInner instantiates a new GetVIPLoanRepaymentHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVIPLoanRepaymentHistoryResponseRowsInner() *GetVIPLoanRepaymentHistoryResponseRowsInner {
	this := GetVIPLoanRepaymentHistoryResponseRowsInner{}
	return &this
}

// NewGetVIPLoanRepaymentHistoryResponseRowsInnerWithDefaults instantiates a new GetVIPLoanRepaymentHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVIPLoanRepaymentHistoryResponseRowsInnerWithDefaults() *GetVIPLoanRepaymentHistoryResponseRowsInner {
	this := GetVIPLoanRepaymentHistoryResponseRowsInner{}
	return &this
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetRepayAmount returns the RepayAmount field value if set, zero value otherwise.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetRepayAmount() string {
	if o == nil || common.IsNil(o.RepayAmount) {
		var ret string
		return ret
	}
	return *o.RepayAmount
}

// GetRepayAmountOk returns a tuple with the RepayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetRepayAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.RepayAmount) {
		return nil, false
	}
	return o.RepayAmount, true
}

// HasRepayAmount returns a boolean if a field has been set.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) HasRepayAmount() bool {
	if o != nil && !common.IsNil(o.RepayAmount) {
		return true
	}

	return false
}

// SetRepayAmount gets a reference to the given string and assigns it to the RepayAmount field.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) SetRepayAmount(v string) {
	o.RepayAmount = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetRepayStatus returns the RepayStatus field value if set, zero value otherwise.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetRepayStatus() string {
	if o == nil || common.IsNil(o.RepayStatus) {
		var ret string
		return ret
	}
	return *o.RepayStatus
}

// GetRepayStatusOk returns a tuple with the RepayStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetRepayStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.RepayStatus) {
		return nil, false
	}
	return o.RepayStatus, true
}

// HasRepayStatus returns a boolean if a field has been set.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) HasRepayStatus() bool {
	if o != nil && !common.IsNil(o.RepayStatus) {
		return true
	}

	return false
}

// SetRepayStatus gets a reference to the given string and assigns it to the RepayStatus field.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) SetRepayStatus(v string) {
	o.RepayStatus = &v
}

// GetLoanDate returns the LoanDate field value if set, zero value otherwise.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetLoanDate() string {
	if o == nil || common.IsNil(o.LoanDate) {
		var ret string
		return ret
	}
	return *o.LoanDate
}

// GetLoanDateOk returns a tuple with the LoanDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetLoanDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanDate) {
		return nil, false
	}
	return o.LoanDate, true
}

// HasLoanDate returns a boolean if a field has been set.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) HasLoanDate() bool {
	if o != nil && !common.IsNil(o.LoanDate) {
		return true
	}

	return false
}

// SetLoanDate gets a reference to the given string and assigns it to the LoanDate field.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) SetLoanDate(v string) {
	o.LoanDate = &v
}

// GetRepayTime returns the RepayTime field value if set, zero value otherwise.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetRepayTime() string {
	if o == nil || common.IsNil(o.RepayTime) {
		var ret string
		return ret
	}
	return *o.RepayTime
}

// GetRepayTimeOk returns a tuple with the RepayTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetRepayTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RepayTime) {
		return nil, false
	}
	return o.RepayTime, true
}

// HasRepayTime returns a boolean if a field has been set.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) HasRepayTime() bool {
	if o != nil && !common.IsNil(o.RepayTime) {
		return true
	}

	return false
}

// SetRepayTime gets a reference to the given string and assigns it to the RepayTime field.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) SetRepayTime(v string) {
	o.RepayTime = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) SetOrderId(v string) {
	o.OrderId = &v
}

func (o GetVIPLoanRepaymentHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVIPLoanRepaymentHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.RepayStatus) {
		toSerialize["repayStatus"] = o.RepayStatus
	}
	if !common.IsNil(o.LoanDate) {
		toSerialize["loanDate"] = o.LoanDate
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

func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetVIPLoanRepaymentHistoryResponseRowsInner := _GetVIPLoanRepaymentHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetVIPLoanRepaymentHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetVIPLoanRepaymentHistoryResponseRowsInner(varGetVIPLoanRepaymentHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "repayAmount")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "repayStatus")
		delete(additionalProperties, "loanDate")
		delete(additionalProperties, "repayTime")
		delete(additionalProperties, "orderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetVIPLoanRepaymentHistoryResponseRowsInner struct {
	value *GetVIPLoanRepaymentHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetVIPLoanRepaymentHistoryResponseRowsInner) Get() *GetVIPLoanRepaymentHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetVIPLoanRepaymentHistoryResponseRowsInner) Set(val *GetVIPLoanRepaymentHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVIPLoanRepaymentHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVIPLoanRepaymentHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVIPLoanRepaymentHistoryResponseRowsInner(val *GetVIPLoanRepaymentHistoryResponseRowsInner) *NullableGetVIPLoanRepaymentHistoryResponseRowsInner {
	return &NullableGetVIPLoanRepaymentHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetVIPLoanRepaymentHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVIPLoanRepaymentHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
