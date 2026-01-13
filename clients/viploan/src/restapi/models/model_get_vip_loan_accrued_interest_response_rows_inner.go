/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetVIPLoanAccruedInterestResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetVIPLoanAccruedInterestResponseRowsInner{}

// GetVIPLoanAccruedInterestResponseRowsInner struct for GetVIPLoanAccruedInterestResponseRowsInner
type GetVIPLoanAccruedInterestResponseRowsInner struct {
	LoanCoin             *string `json:"loanCoin,omitempty"`
	PrincipalAmount      *string `json:"principalAmount,omitempty"`
	InterestAmount       *string `json:"interestAmount,omitempty"`
	AnnualInterestRate   *string `json:"annualInterestRate,omitempty"`
	AccrualTime          *int64  `json:"accrualTime,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetVIPLoanAccruedInterestResponseRowsInner GetVIPLoanAccruedInterestResponseRowsInner

// NewGetVIPLoanAccruedInterestResponseRowsInner instantiates a new GetVIPLoanAccruedInterestResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVIPLoanAccruedInterestResponseRowsInner() *GetVIPLoanAccruedInterestResponseRowsInner {
	this := GetVIPLoanAccruedInterestResponseRowsInner{}
	return &this
}

// NewGetVIPLoanAccruedInterestResponseRowsInnerWithDefaults instantiates a new GetVIPLoanAccruedInterestResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVIPLoanAccruedInterestResponseRowsInnerWithDefaults() *GetVIPLoanAccruedInterestResponseRowsInner {
	this := GetVIPLoanAccruedInterestResponseRowsInner{}
	return &this
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetPrincipalAmount returns the PrincipalAmount field value if set, zero value otherwise.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) GetPrincipalAmount() string {
	if o == nil || common.IsNil(o.PrincipalAmount) {
		var ret string
		return ret
	}
	return *o.PrincipalAmount
}

// GetPrincipalAmountOk returns a tuple with the PrincipalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) GetPrincipalAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.PrincipalAmount) {
		return nil, false
	}
	return o.PrincipalAmount, true
}

// HasPrincipalAmount returns a boolean if a field has been set.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) HasPrincipalAmount() bool {
	if o != nil && !common.IsNil(o.PrincipalAmount) {
		return true
	}

	return false
}

// SetPrincipalAmount gets a reference to the given string and assigns it to the PrincipalAmount field.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) SetPrincipalAmount(v string) {
	o.PrincipalAmount = &v
}

// GetInterestAmount returns the InterestAmount field value if set, zero value otherwise.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) GetInterestAmount() string {
	if o == nil || common.IsNil(o.InterestAmount) {
		var ret string
		return ret
	}
	return *o.InterestAmount
}

// GetInterestAmountOk returns a tuple with the InterestAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) GetInterestAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.InterestAmount) {
		return nil, false
	}
	return o.InterestAmount, true
}

// HasInterestAmount returns a boolean if a field has been set.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) HasInterestAmount() bool {
	if o != nil && !common.IsNil(o.InterestAmount) {
		return true
	}

	return false
}

// SetInterestAmount gets a reference to the given string and assigns it to the InterestAmount field.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) SetInterestAmount(v string) {
	o.InterestAmount = &v
}

// GetAnnualInterestRate returns the AnnualInterestRate field value if set, zero value otherwise.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) GetAnnualInterestRate() string {
	if o == nil || common.IsNil(o.AnnualInterestRate) {
		var ret string
		return ret
	}
	return *o.AnnualInterestRate
}

// GetAnnualInterestRateOk returns a tuple with the AnnualInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) GetAnnualInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.AnnualInterestRate) {
		return nil, false
	}
	return o.AnnualInterestRate, true
}

// HasAnnualInterestRate returns a boolean if a field has been set.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) HasAnnualInterestRate() bool {
	if o != nil && !common.IsNil(o.AnnualInterestRate) {
		return true
	}

	return false
}

// SetAnnualInterestRate gets a reference to the given string and assigns it to the AnnualInterestRate field.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) SetAnnualInterestRate(v string) {
	o.AnnualInterestRate = &v
}

// GetAccrualTime returns the AccrualTime field value if set, zero value otherwise.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) GetAccrualTime() int64 {
	if o == nil || common.IsNil(o.AccrualTime) {
		var ret int64
		return ret
	}
	return *o.AccrualTime
}

// GetAccrualTimeOk returns a tuple with the AccrualTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) GetAccrualTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AccrualTime) {
		return nil, false
	}
	return o.AccrualTime, true
}

// HasAccrualTime returns a boolean if a field has been set.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) HasAccrualTime() bool {
	if o != nil && !common.IsNil(o.AccrualTime) {
		return true
	}

	return false
}

// SetAccrualTime gets a reference to the given int64 and assigns it to the AccrualTime field.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) SetAccrualTime(v int64) {
	o.AccrualTime = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetVIPLoanAccruedInterestResponseRowsInner) SetOrderId(v int64) {
	o.OrderId = &v
}

func (o GetVIPLoanAccruedInterestResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVIPLoanAccruedInterestResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.PrincipalAmount) {
		toSerialize["principalAmount"] = o.PrincipalAmount
	}
	if !common.IsNil(o.InterestAmount) {
		toSerialize["interestAmount"] = o.InterestAmount
	}
	if !common.IsNil(o.AnnualInterestRate) {
		toSerialize["annualInterestRate"] = o.AnnualInterestRate
	}
	if !common.IsNil(o.AccrualTime) {
		toSerialize["accrualTime"] = o.AccrualTime
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetVIPLoanAccruedInterestResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetVIPLoanAccruedInterestResponseRowsInner := _GetVIPLoanAccruedInterestResponseRowsInner{}

	err = json.Unmarshal(data, &varGetVIPLoanAccruedInterestResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetVIPLoanAccruedInterestResponseRowsInner(varGetVIPLoanAccruedInterestResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "principalAmount")
		delete(additionalProperties, "interestAmount")
		delete(additionalProperties, "annualInterestRate")
		delete(additionalProperties, "accrualTime")
		delete(additionalProperties, "orderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetVIPLoanAccruedInterestResponseRowsInner struct {
	value *GetVIPLoanAccruedInterestResponseRowsInner
	isSet bool
}

func (v NullableGetVIPLoanAccruedInterestResponseRowsInner) Get() *GetVIPLoanAccruedInterestResponseRowsInner {
	return v.value
}

func (v *NullableGetVIPLoanAccruedInterestResponseRowsInner) Set(val *GetVIPLoanAccruedInterestResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVIPLoanAccruedInterestResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVIPLoanAccruedInterestResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVIPLoanAccruedInterestResponseRowsInner(val *GetVIPLoanAccruedInterestResponseRowsInner) *NullableGetVIPLoanAccruedInterestResponseRowsInner {
	return &NullableGetVIPLoanAccruedInterestResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetVIPLoanAccruedInterestResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVIPLoanAccruedInterestResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
