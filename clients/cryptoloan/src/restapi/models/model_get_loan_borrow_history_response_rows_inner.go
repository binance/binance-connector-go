/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetLoanBorrowHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLoanBorrowHistoryResponseRowsInner{}

// GetLoanBorrowHistoryResponseRowsInner struct for GetLoanBorrowHistoryResponseRowsInner
type GetLoanBorrowHistoryResponseRowsInner struct {
	OrderId                 *int64  `json:"orderId,omitempty"`
	LoanCoin                *string `json:"loanCoin,omitempty"`
	InitialLoanAmount       *string `json:"initialLoanAmount,omitempty"`
	HourlyInterestRate      *string `json:"hourlyInterestRate,omitempty"`
	LoanTerm                *string `json:"loanTerm,omitempty"`
	CollateralCoin          *string `json:"collateralCoin,omitempty"`
	InitialCollateralAmount *string `json:"initialCollateralAmount,omitempty"`
	BorrowTime              *int64  `json:"borrowTime,omitempty"`
	Status                  *string `json:"status,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _GetLoanBorrowHistoryResponseRowsInner GetLoanBorrowHistoryResponseRowsInner

// NewGetLoanBorrowHistoryResponseRowsInner instantiates a new GetLoanBorrowHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLoanBorrowHistoryResponseRowsInner() *GetLoanBorrowHistoryResponseRowsInner {
	this := GetLoanBorrowHistoryResponseRowsInner{}
	return &this
}

// NewGetLoanBorrowHistoryResponseRowsInnerWithDefaults instantiates a new GetLoanBorrowHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLoanBorrowHistoryResponseRowsInnerWithDefaults() *GetLoanBorrowHistoryResponseRowsInner {
	this := GetLoanBorrowHistoryResponseRowsInner{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetLoanBorrowHistoryResponseRowsInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *GetLoanBorrowHistoryResponseRowsInner) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetInitialLoanAmount returns the InitialLoanAmount field value if set, zero value otherwise.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetInitialLoanAmount() string {
	if o == nil || common.IsNil(o.InitialLoanAmount) {
		var ret string
		return ret
	}
	return *o.InitialLoanAmount
}

// GetInitialLoanAmountOk returns a tuple with the InitialLoanAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetInitialLoanAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialLoanAmount) {
		return nil, false
	}
	return o.InitialLoanAmount, true
}

// HasInitialLoanAmount returns a boolean if a field has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) HasInitialLoanAmount() bool {
	if o != nil && !common.IsNil(o.InitialLoanAmount) {
		return true
	}

	return false
}

// SetInitialLoanAmount gets a reference to the given string and assigns it to the InitialLoanAmount field.
func (o *GetLoanBorrowHistoryResponseRowsInner) SetInitialLoanAmount(v string) {
	o.InitialLoanAmount = &v
}

// GetHourlyInterestRate returns the HourlyInterestRate field value if set, zero value otherwise.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetHourlyInterestRate() string {
	if o == nil || common.IsNil(o.HourlyInterestRate) {
		var ret string
		return ret
	}
	return *o.HourlyInterestRate
}

// GetHourlyInterestRateOk returns a tuple with the HourlyInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetHourlyInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.HourlyInterestRate) {
		return nil, false
	}
	return o.HourlyInterestRate, true
}

// HasHourlyInterestRate returns a boolean if a field has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) HasHourlyInterestRate() bool {
	if o != nil && !common.IsNil(o.HourlyInterestRate) {
		return true
	}

	return false
}

// SetHourlyInterestRate gets a reference to the given string and assigns it to the HourlyInterestRate field.
func (o *GetLoanBorrowHistoryResponseRowsInner) SetHourlyInterestRate(v string) {
	o.HourlyInterestRate = &v
}

// GetLoanTerm returns the LoanTerm field value if set, zero value otherwise.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetLoanTerm() string {
	if o == nil || common.IsNil(o.LoanTerm) {
		var ret string
		return ret
	}
	return *o.LoanTerm
}

// GetLoanTermOk returns a tuple with the LoanTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetLoanTermOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanTerm) {
		return nil, false
	}
	return o.LoanTerm, true
}

// HasLoanTerm returns a boolean if a field has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) HasLoanTerm() bool {
	if o != nil && !common.IsNil(o.LoanTerm) {
		return true
	}

	return false
}

// SetLoanTerm gets a reference to the given string and assigns it to the LoanTerm field.
func (o *GetLoanBorrowHistoryResponseRowsInner) SetLoanTerm(v string) {
	o.LoanTerm = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *GetLoanBorrowHistoryResponseRowsInner) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetInitialCollateralAmount returns the InitialCollateralAmount field value if set, zero value otherwise.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetInitialCollateralAmount() string {
	if o == nil || common.IsNil(o.InitialCollateralAmount) {
		var ret string
		return ret
	}
	return *o.InitialCollateralAmount
}

// GetInitialCollateralAmountOk returns a tuple with the InitialCollateralAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetInitialCollateralAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialCollateralAmount) {
		return nil, false
	}
	return o.InitialCollateralAmount, true
}

// HasInitialCollateralAmount returns a boolean if a field has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) HasInitialCollateralAmount() bool {
	if o != nil && !common.IsNil(o.InitialCollateralAmount) {
		return true
	}

	return false
}

// SetInitialCollateralAmount gets a reference to the given string and assigns it to the InitialCollateralAmount field.
func (o *GetLoanBorrowHistoryResponseRowsInner) SetInitialCollateralAmount(v string) {
	o.InitialCollateralAmount = &v
}

// GetBorrowTime returns the BorrowTime field value if set, zero value otherwise.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetBorrowTime() int64 {
	if o == nil || common.IsNil(o.BorrowTime) {
		var ret int64
		return ret
	}
	return *o.BorrowTime
}

// GetBorrowTimeOk returns a tuple with the BorrowTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetBorrowTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BorrowTime) {
		return nil, false
	}
	return o.BorrowTime, true
}

// HasBorrowTime returns a boolean if a field has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) HasBorrowTime() bool {
	if o != nil && !common.IsNil(o.BorrowTime) {
		return true
	}

	return false
}

// SetBorrowTime gets a reference to the given int64 and assigns it to the BorrowTime field.
func (o *GetLoanBorrowHistoryResponseRowsInner) SetBorrowTime(v int64) {
	o.BorrowTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetLoanBorrowHistoryResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetLoanBorrowHistoryResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetLoanBorrowHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLoanBorrowHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.InitialLoanAmount) {
		toSerialize["initialLoanAmount"] = o.InitialLoanAmount
	}
	if !common.IsNil(o.HourlyInterestRate) {
		toSerialize["hourlyInterestRate"] = o.HourlyInterestRate
	}
	if !common.IsNil(o.LoanTerm) {
		toSerialize["loanTerm"] = o.LoanTerm
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

func (o *GetLoanBorrowHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetLoanBorrowHistoryResponseRowsInner := _GetLoanBorrowHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetLoanBorrowHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetLoanBorrowHistoryResponseRowsInner(varGetLoanBorrowHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "initialLoanAmount")
		delete(additionalProperties, "hourlyInterestRate")
		delete(additionalProperties, "loanTerm")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "initialCollateralAmount")
		delete(additionalProperties, "borrowTime")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLoanBorrowHistoryResponseRowsInner struct {
	value *GetLoanBorrowHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetLoanBorrowHistoryResponseRowsInner) Get() *GetLoanBorrowHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetLoanBorrowHistoryResponseRowsInner) Set(val *GetLoanBorrowHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLoanBorrowHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLoanBorrowHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLoanBorrowHistoryResponseRowsInner(val *GetLoanBorrowHistoryResponseRowsInner) *NullableGetLoanBorrowHistoryResponseRowsInner {
	return &NullableGetLoanBorrowHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetLoanBorrowHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLoanBorrowHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
