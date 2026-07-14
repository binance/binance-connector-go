/*
VIP Loan REST API

Access over-collateralized loan services, manage positions, and monitor collateral via the VIP Loan API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the VipLoanFixedRateBorrowResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &VipLoanFixedRateBorrowResponse{}

// VipLoanFixedRateBorrowResponse struct for VipLoanFixedRateBorrowResponse
type VipLoanFixedRateBorrowResponse struct {
	// Echo of input parameter
	BorrowCoin *string `json:"borrowCoin,omitempty"`
	// Actual total borrow amount (aggregated when multiple supplyRequest)
	BorrowAmount *string `json:"borrowAmount,omitempty"`
	// Actual received amount
	ActualReceivedAmount *string `json:"actualReceivedAmount,omitempty"`
	// Echo of input parameter, comma-separated
	CollateralCoin *string `json:"collateralCoin,omitempty"`
	// Echo of input parameter, comma-separated
	CollateralAccountId *string `json:"collateralAccountId,omitempty"`
	// Actual borrow interest rate (weighted average when multiple)
	BorrowInterestRate *string `json:"borrowInterestRate,omitempty"`
	// `{loanTerm}Days`, e.g. \"30Days\"
	Duration *string `json:"duration,omitempty"`
	// Echo of input parameter
	AutoRepay *bool `json:"autoRepay,omitempty"`
	// Order ID
	OrderId *int64 `json:"orderId,omitempty"`
	// `Succeeds` / `Failed` / `Processing`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VipLoanFixedRateBorrowResponse VipLoanFixedRateBorrowResponse

// NewVipLoanFixedRateBorrowResponse instantiates a new VipLoanFixedRateBorrowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVipLoanFixedRateBorrowResponse() *VipLoanFixedRateBorrowResponse {
	this := VipLoanFixedRateBorrowResponse{}
	return &this
}

// NewVipLoanFixedRateBorrowResponseWithDefaults instantiates a new VipLoanFixedRateBorrowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVipLoanFixedRateBorrowResponseWithDefaults() *VipLoanFixedRateBorrowResponse {
	this := VipLoanFixedRateBorrowResponse{}
	return &this
}

// GetBorrowCoin returns the BorrowCoin field value if set, zero value otherwise.
func (o *VipLoanFixedRateBorrowResponse) GetBorrowCoin() string {
	if o == nil || common.IsNil(o.BorrowCoin) {
		var ret string
		return ret
	}
	return *o.BorrowCoin
}

// GetBorrowCoinOk returns a tuple with the BorrowCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanFixedRateBorrowResponse) GetBorrowCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.BorrowCoin) {
		return nil, false
	}
	return o.BorrowCoin, true
}

// HasBorrowCoin returns a boolean if a field has been set.
func (o *VipLoanFixedRateBorrowResponse) HasBorrowCoin() bool {
	if o != nil && !common.IsNil(o.BorrowCoin) {
		return true
	}

	return false
}

// SetBorrowCoin gets a reference to the given string and assigns it to the BorrowCoin field.
func (o *VipLoanFixedRateBorrowResponse) SetBorrowCoin(v string) {
	o.BorrowCoin = &v
}

// GetBorrowAmount returns the BorrowAmount field value if set, zero value otherwise.
func (o *VipLoanFixedRateBorrowResponse) GetBorrowAmount() string {
	if o == nil || common.IsNil(o.BorrowAmount) {
		var ret string
		return ret
	}
	return *o.BorrowAmount
}

// GetBorrowAmountOk returns a tuple with the BorrowAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanFixedRateBorrowResponse) GetBorrowAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.BorrowAmount) {
		return nil, false
	}
	return o.BorrowAmount, true
}

// HasBorrowAmount returns a boolean if a field has been set.
func (o *VipLoanFixedRateBorrowResponse) HasBorrowAmount() bool {
	if o != nil && !common.IsNil(o.BorrowAmount) {
		return true
	}

	return false
}

// SetBorrowAmount gets a reference to the given string and assigns it to the BorrowAmount field.
func (o *VipLoanFixedRateBorrowResponse) SetBorrowAmount(v string) {
	o.BorrowAmount = &v
}

// GetActualReceivedAmount returns the ActualReceivedAmount field value if set, zero value otherwise.
func (o *VipLoanFixedRateBorrowResponse) GetActualReceivedAmount() string {
	if o == nil || common.IsNil(o.ActualReceivedAmount) {
		var ret string
		return ret
	}
	return *o.ActualReceivedAmount
}

// GetActualReceivedAmountOk returns a tuple with the ActualReceivedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanFixedRateBorrowResponse) GetActualReceivedAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ActualReceivedAmount) {
		return nil, false
	}
	return o.ActualReceivedAmount, true
}

// HasActualReceivedAmount returns a boolean if a field has been set.
func (o *VipLoanFixedRateBorrowResponse) HasActualReceivedAmount() bool {
	if o != nil && !common.IsNil(o.ActualReceivedAmount) {
		return true
	}

	return false
}

// SetActualReceivedAmount gets a reference to the given string and assigns it to the ActualReceivedAmount field.
func (o *VipLoanFixedRateBorrowResponse) SetActualReceivedAmount(v string) {
	o.ActualReceivedAmount = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *VipLoanFixedRateBorrowResponse) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanFixedRateBorrowResponse) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *VipLoanFixedRateBorrowResponse) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *VipLoanFixedRateBorrowResponse) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetCollateralAccountId returns the CollateralAccountId field value if set, zero value otherwise.
func (o *VipLoanFixedRateBorrowResponse) GetCollateralAccountId() string {
	if o == nil || common.IsNil(o.CollateralAccountId) {
		var ret string
		return ret
	}
	return *o.CollateralAccountId
}

// GetCollateralAccountIdOk returns a tuple with the CollateralAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanFixedRateBorrowResponse) GetCollateralAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralAccountId) {
		return nil, false
	}
	return o.CollateralAccountId, true
}

// HasCollateralAccountId returns a boolean if a field has been set.
func (o *VipLoanFixedRateBorrowResponse) HasCollateralAccountId() bool {
	if o != nil && !common.IsNil(o.CollateralAccountId) {
		return true
	}

	return false
}

// SetCollateralAccountId gets a reference to the given string and assigns it to the CollateralAccountId field.
func (o *VipLoanFixedRateBorrowResponse) SetCollateralAccountId(v string) {
	o.CollateralAccountId = &v
}

// GetBorrowInterestRate returns the BorrowInterestRate field value if set, zero value otherwise.
func (o *VipLoanFixedRateBorrowResponse) GetBorrowInterestRate() string {
	if o == nil || common.IsNil(o.BorrowInterestRate) {
		var ret string
		return ret
	}
	return *o.BorrowInterestRate
}

// GetBorrowInterestRateOk returns a tuple with the BorrowInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanFixedRateBorrowResponse) GetBorrowInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.BorrowInterestRate) {
		return nil, false
	}
	return o.BorrowInterestRate, true
}

// HasBorrowInterestRate returns a boolean if a field has been set.
func (o *VipLoanFixedRateBorrowResponse) HasBorrowInterestRate() bool {
	if o != nil && !common.IsNil(o.BorrowInterestRate) {
		return true
	}

	return false
}

// SetBorrowInterestRate gets a reference to the given string and assigns it to the BorrowInterestRate field.
func (o *VipLoanFixedRateBorrowResponse) SetBorrowInterestRate(v string) {
	o.BorrowInterestRate = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *VipLoanFixedRateBorrowResponse) GetDuration() string {
	if o == nil || common.IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanFixedRateBorrowResponse) GetDurationOk() (*string, bool) {
	if o == nil || common.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *VipLoanFixedRateBorrowResponse) HasDuration() bool {
	if o != nil && !common.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *VipLoanFixedRateBorrowResponse) SetDuration(v string) {
	o.Duration = &v
}

// GetAutoRepay returns the AutoRepay field value if set, zero value otherwise.
func (o *VipLoanFixedRateBorrowResponse) GetAutoRepay() bool {
	if o == nil || common.IsNil(o.AutoRepay) {
		var ret bool
		return ret
	}
	return *o.AutoRepay
}

// GetAutoRepayOk returns a tuple with the AutoRepay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanFixedRateBorrowResponse) GetAutoRepayOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AutoRepay) {
		return nil, false
	}
	return o.AutoRepay, true
}

// HasAutoRepay returns a boolean if a field has been set.
func (o *VipLoanFixedRateBorrowResponse) HasAutoRepay() bool {
	if o != nil && !common.IsNil(o.AutoRepay) {
		return true
	}

	return false
}

// SetAutoRepay gets a reference to the given bool and assigns it to the AutoRepay field.
func (o *VipLoanFixedRateBorrowResponse) SetAutoRepay(v bool) {
	o.AutoRepay = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *VipLoanFixedRateBorrowResponse) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanFixedRateBorrowResponse) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *VipLoanFixedRateBorrowResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *VipLoanFixedRateBorrowResponse) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VipLoanFixedRateBorrowResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanFixedRateBorrowResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VipLoanFixedRateBorrowResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VipLoanFixedRateBorrowResponse) SetStatus(v string) {
	o.Status = &v
}

func (o VipLoanFixedRateBorrowResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VipLoanFixedRateBorrowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BorrowCoin) {
		toSerialize["borrowCoin"] = o.BorrowCoin
	}
	if !common.IsNil(o.BorrowAmount) {
		toSerialize["borrowAmount"] = o.BorrowAmount
	}
	if !common.IsNil(o.ActualReceivedAmount) {
		toSerialize["actualReceivedAmount"] = o.ActualReceivedAmount
	}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !common.IsNil(o.CollateralAccountId) {
		toSerialize["collateralAccountId"] = o.CollateralAccountId
	}
	if !common.IsNil(o.BorrowInterestRate) {
		toSerialize["borrowInterestRate"] = o.BorrowInterestRate
	}
	if !common.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !common.IsNil(o.AutoRepay) {
		toSerialize["autoRepay"] = o.AutoRepay
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VipLoanFixedRateBorrowResponse) UnmarshalJSON(data []byte) (err error) {
	varVipLoanFixedRateBorrowResponse := _VipLoanFixedRateBorrowResponse{}

	err = json.Unmarshal(data, &varVipLoanFixedRateBorrowResponse)

	if err != nil {
		return err
	}

	*o = VipLoanFixedRateBorrowResponse(varVipLoanFixedRateBorrowResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "borrowCoin")
		delete(additionalProperties, "borrowAmount")
		delete(additionalProperties, "actualReceivedAmount")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "collateralAccountId")
		delete(additionalProperties, "borrowInterestRate")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "autoRepay")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVipLoanFixedRateBorrowResponse struct {
	value *VipLoanFixedRateBorrowResponse
	isSet bool
}

func (v NullableVipLoanFixedRateBorrowResponse) Get() *VipLoanFixedRateBorrowResponse {
	return v.value
}

func (v *NullableVipLoanFixedRateBorrowResponse) Set(val *VipLoanFixedRateBorrowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVipLoanFixedRateBorrowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVipLoanFixedRateBorrowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVipLoanFixedRateBorrowResponse(val *VipLoanFixedRateBorrowResponse) *NullableVipLoanFixedRateBorrowResponse {
	return &NullableVipLoanFixedRateBorrowResponse{value: val, isSet: true}
}

func (v NullableVipLoanFixedRateBorrowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVipLoanFixedRateBorrowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
