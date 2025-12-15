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

// checks if the FlexibleLoanRepayResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FlexibleLoanRepayResponse{}

// FlexibleLoanRepayResponse struct for FlexibleLoanRepayResponse
type FlexibleLoanRepayResponse struct {
	LoanCoin             *string `json:"loanCoin,omitempty"`
	CollateralCoin       *string `json:"collateralCoin,omitempty"`
	RemainingDebt        *string `json:"remainingDebt,omitempty"`
	RemainingCollateral  *string `json:"remainingCollateral,omitempty"`
	FullRepayment        *bool   `json:"fullRepayment,omitempty"`
	CurrentLTV           *string `json:"currentLTV,omitempty"`
	RepayStatus          *string `json:"repayStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FlexibleLoanRepayResponse FlexibleLoanRepayResponse

// NewFlexibleLoanRepayResponse instantiates a new FlexibleLoanRepayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlexibleLoanRepayResponse() *FlexibleLoanRepayResponse {
	this := FlexibleLoanRepayResponse{}
	return &this
}

// NewFlexibleLoanRepayResponseWithDefaults instantiates a new FlexibleLoanRepayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlexibleLoanRepayResponseWithDefaults() *FlexibleLoanRepayResponse {
	this := FlexibleLoanRepayResponse{}
	return &this
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *FlexibleLoanRepayResponse) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanRepayResponse) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *FlexibleLoanRepayResponse) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *FlexibleLoanRepayResponse) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *FlexibleLoanRepayResponse) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanRepayResponse) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *FlexibleLoanRepayResponse) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *FlexibleLoanRepayResponse) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetRemainingDebt returns the RemainingDebt field value if set, zero value otherwise.
func (o *FlexibleLoanRepayResponse) GetRemainingDebt() string {
	if o == nil || common.IsNil(o.RemainingDebt) {
		var ret string
		return ret
	}
	return *o.RemainingDebt
}

// GetRemainingDebtOk returns a tuple with the RemainingDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanRepayResponse) GetRemainingDebtOk() (*string, bool) {
	if o == nil || common.IsNil(o.RemainingDebt) {
		return nil, false
	}
	return o.RemainingDebt, true
}

// HasRemainingDebt returns a boolean if a field has been set.
func (o *FlexibleLoanRepayResponse) HasRemainingDebt() bool {
	if o != nil && !common.IsNil(o.RemainingDebt) {
		return true
	}

	return false
}

// SetRemainingDebt gets a reference to the given string and assigns it to the RemainingDebt field.
func (o *FlexibleLoanRepayResponse) SetRemainingDebt(v string) {
	o.RemainingDebt = &v
}

// GetRemainingCollateral returns the RemainingCollateral field value if set, zero value otherwise.
func (o *FlexibleLoanRepayResponse) GetRemainingCollateral() string {
	if o == nil || common.IsNil(o.RemainingCollateral) {
		var ret string
		return ret
	}
	return *o.RemainingCollateral
}

// GetRemainingCollateralOk returns a tuple with the RemainingCollateral field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanRepayResponse) GetRemainingCollateralOk() (*string, bool) {
	if o == nil || common.IsNil(o.RemainingCollateral) {
		return nil, false
	}
	return o.RemainingCollateral, true
}

// HasRemainingCollateral returns a boolean if a field has been set.
func (o *FlexibleLoanRepayResponse) HasRemainingCollateral() bool {
	if o != nil && !common.IsNil(o.RemainingCollateral) {
		return true
	}

	return false
}

// SetRemainingCollateral gets a reference to the given string and assigns it to the RemainingCollateral field.
func (o *FlexibleLoanRepayResponse) SetRemainingCollateral(v string) {
	o.RemainingCollateral = &v
}

// GetFullRepayment returns the FullRepayment field value if set, zero value otherwise.
func (o *FlexibleLoanRepayResponse) GetFullRepayment() bool {
	if o == nil || common.IsNil(o.FullRepayment) {
		var ret bool
		return ret
	}
	return *o.FullRepayment
}

// GetFullRepaymentOk returns a tuple with the FullRepayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanRepayResponse) GetFullRepaymentOk() (*bool, bool) {
	if o == nil || common.IsNil(o.FullRepayment) {
		return nil, false
	}
	return o.FullRepayment, true
}

// HasFullRepayment returns a boolean if a field has been set.
func (o *FlexibleLoanRepayResponse) HasFullRepayment() bool {
	if o != nil && !common.IsNil(o.FullRepayment) {
		return true
	}

	return false
}

// SetFullRepayment gets a reference to the given bool and assigns it to the FullRepayment field.
func (o *FlexibleLoanRepayResponse) SetFullRepayment(v bool) {
	o.FullRepayment = &v
}

// GetCurrentLTV returns the CurrentLTV field value if set, zero value otherwise.
func (o *FlexibleLoanRepayResponse) GetCurrentLTV() string {
	if o == nil || common.IsNil(o.CurrentLTV) {
		var ret string
		return ret
	}
	return *o.CurrentLTV
}

// GetCurrentLTVOk returns a tuple with the CurrentLTV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanRepayResponse) GetCurrentLTVOk() (*string, bool) {
	if o == nil || common.IsNil(o.CurrentLTV) {
		return nil, false
	}
	return o.CurrentLTV, true
}

// HasCurrentLTV returns a boolean if a field has been set.
func (o *FlexibleLoanRepayResponse) HasCurrentLTV() bool {
	if o != nil && !common.IsNil(o.CurrentLTV) {
		return true
	}

	return false
}

// SetCurrentLTV gets a reference to the given string and assigns it to the CurrentLTV field.
func (o *FlexibleLoanRepayResponse) SetCurrentLTV(v string) {
	o.CurrentLTV = &v
}

// GetRepayStatus returns the RepayStatus field value if set, zero value otherwise.
func (o *FlexibleLoanRepayResponse) GetRepayStatus() string {
	if o == nil || common.IsNil(o.RepayStatus) {
		var ret string
		return ret
	}
	return *o.RepayStatus
}

// GetRepayStatusOk returns a tuple with the RepayStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanRepayResponse) GetRepayStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.RepayStatus) {
		return nil, false
	}
	return o.RepayStatus, true
}

// HasRepayStatus returns a boolean if a field has been set.
func (o *FlexibleLoanRepayResponse) HasRepayStatus() bool {
	if o != nil && !common.IsNil(o.RepayStatus) {
		return true
	}

	return false
}

// SetRepayStatus gets a reference to the given string and assigns it to the RepayStatus field.
func (o *FlexibleLoanRepayResponse) SetRepayStatus(v string) {
	o.RepayStatus = &v
}

func (o FlexibleLoanRepayResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlexibleLoanRepayResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !common.IsNil(o.RemainingDebt) {
		toSerialize["remainingDebt"] = o.RemainingDebt
	}
	if !common.IsNil(o.RemainingCollateral) {
		toSerialize["remainingCollateral"] = o.RemainingCollateral
	}
	if !common.IsNil(o.FullRepayment) {
		toSerialize["fullRepayment"] = o.FullRepayment
	}
	if !common.IsNil(o.CurrentLTV) {
		toSerialize["currentLTV"] = o.CurrentLTV
	}
	if !common.IsNil(o.RepayStatus) {
		toSerialize["repayStatus"] = o.RepayStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FlexibleLoanRepayResponse) UnmarshalJSON(data []byte) (err error) {
	varFlexibleLoanRepayResponse := _FlexibleLoanRepayResponse{}

	err = json.Unmarshal(data, &varFlexibleLoanRepayResponse)

	if err != nil {
		return err
	}

	*o = FlexibleLoanRepayResponse(varFlexibleLoanRepayResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "remainingDebt")
		delete(additionalProperties, "remainingCollateral")
		delete(additionalProperties, "fullRepayment")
		delete(additionalProperties, "currentLTV")
		delete(additionalProperties, "repayStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFlexibleLoanRepayResponse struct {
	value *FlexibleLoanRepayResponse
	isSet bool
}

func (v NullableFlexibleLoanRepayResponse) Get() *FlexibleLoanRepayResponse {
	return v.value
}

func (v *NullableFlexibleLoanRepayResponse) Set(val *FlexibleLoanRepayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFlexibleLoanRepayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFlexibleLoanRepayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlexibleLoanRepayResponse(val *FlexibleLoanRepayResponse) *NullableFlexibleLoanRepayResponse {
	return &NullableFlexibleLoanRepayResponse{value: val, isSet: true}
}

func (v NullableFlexibleLoanRepayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlexibleLoanRepayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
