/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the VipLoanRepayResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &VipLoanRepayResponse{}

// VipLoanRepayResponse struct for VipLoanRepayResponse
type VipLoanRepayResponse struct {
	LoanCoin             *string `json:"loanCoin,omitempty"`
	RepayAmount          *string `json:"repayAmount,omitempty"`
	RemainingPrincipal   *string `json:"remainingPrincipal,omitempty"`
	RemainingInterest    *string `json:"remainingInterest,omitempty"`
	CollateralCoin       *string `json:"collateralCoin,omitempty"`
	CurrentLTV           *string `json:"currentLTV,omitempty"`
	RepayStatus          *string `json:"repayStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VipLoanRepayResponse VipLoanRepayResponse

// NewVipLoanRepayResponse instantiates a new VipLoanRepayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVipLoanRepayResponse() *VipLoanRepayResponse {
	this := VipLoanRepayResponse{}
	return &this
}

// NewVipLoanRepayResponseWithDefaults instantiates a new VipLoanRepayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVipLoanRepayResponseWithDefaults() *VipLoanRepayResponse {
	this := VipLoanRepayResponse{}
	return &this
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *VipLoanRepayResponse) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanRepayResponse) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *VipLoanRepayResponse) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *VipLoanRepayResponse) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetRepayAmount returns the RepayAmount field value if set, zero value otherwise.
func (o *VipLoanRepayResponse) GetRepayAmount() string {
	if o == nil || common.IsNil(o.RepayAmount) {
		var ret string
		return ret
	}
	return *o.RepayAmount
}

// GetRepayAmountOk returns a tuple with the RepayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanRepayResponse) GetRepayAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.RepayAmount) {
		return nil, false
	}
	return o.RepayAmount, true
}

// HasRepayAmount returns a boolean if a field has been set.
func (o *VipLoanRepayResponse) HasRepayAmount() bool {
	if o != nil && !common.IsNil(o.RepayAmount) {
		return true
	}

	return false
}

// SetRepayAmount gets a reference to the given string and assigns it to the RepayAmount field.
func (o *VipLoanRepayResponse) SetRepayAmount(v string) {
	o.RepayAmount = &v
}

// GetRemainingPrincipal returns the RemainingPrincipal field value if set, zero value otherwise.
func (o *VipLoanRepayResponse) GetRemainingPrincipal() string {
	if o == nil || common.IsNil(o.RemainingPrincipal) {
		var ret string
		return ret
	}
	return *o.RemainingPrincipal
}

// GetRemainingPrincipalOk returns a tuple with the RemainingPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanRepayResponse) GetRemainingPrincipalOk() (*string, bool) {
	if o == nil || common.IsNil(o.RemainingPrincipal) {
		return nil, false
	}
	return o.RemainingPrincipal, true
}

// HasRemainingPrincipal returns a boolean if a field has been set.
func (o *VipLoanRepayResponse) HasRemainingPrincipal() bool {
	if o != nil && !common.IsNil(o.RemainingPrincipal) {
		return true
	}

	return false
}

// SetRemainingPrincipal gets a reference to the given string and assigns it to the RemainingPrincipal field.
func (o *VipLoanRepayResponse) SetRemainingPrincipal(v string) {
	o.RemainingPrincipal = &v
}

// GetRemainingInterest returns the RemainingInterest field value if set, zero value otherwise.
func (o *VipLoanRepayResponse) GetRemainingInterest() string {
	if o == nil || common.IsNil(o.RemainingInterest) {
		var ret string
		return ret
	}
	return *o.RemainingInterest
}

// GetRemainingInterestOk returns a tuple with the RemainingInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanRepayResponse) GetRemainingInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.RemainingInterest) {
		return nil, false
	}
	return o.RemainingInterest, true
}

// HasRemainingInterest returns a boolean if a field has been set.
func (o *VipLoanRepayResponse) HasRemainingInterest() bool {
	if o != nil && !common.IsNil(o.RemainingInterest) {
		return true
	}

	return false
}

// SetRemainingInterest gets a reference to the given string and assigns it to the RemainingInterest field.
func (o *VipLoanRepayResponse) SetRemainingInterest(v string) {
	o.RemainingInterest = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *VipLoanRepayResponse) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanRepayResponse) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *VipLoanRepayResponse) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *VipLoanRepayResponse) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetCurrentLTV returns the CurrentLTV field value if set, zero value otherwise.
func (o *VipLoanRepayResponse) GetCurrentLTV() string {
	if o == nil || common.IsNil(o.CurrentLTV) {
		var ret string
		return ret
	}
	return *o.CurrentLTV
}

// GetCurrentLTVOk returns a tuple with the CurrentLTV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanRepayResponse) GetCurrentLTVOk() (*string, bool) {
	if o == nil || common.IsNil(o.CurrentLTV) {
		return nil, false
	}
	return o.CurrentLTV, true
}

// HasCurrentLTV returns a boolean if a field has been set.
func (o *VipLoanRepayResponse) HasCurrentLTV() bool {
	if o != nil && !common.IsNil(o.CurrentLTV) {
		return true
	}

	return false
}

// SetCurrentLTV gets a reference to the given string and assigns it to the CurrentLTV field.
func (o *VipLoanRepayResponse) SetCurrentLTV(v string) {
	o.CurrentLTV = &v
}

// GetRepayStatus returns the RepayStatus field value if set, zero value otherwise.
func (o *VipLoanRepayResponse) GetRepayStatus() string {
	if o == nil || common.IsNil(o.RepayStatus) {
		var ret string
		return ret
	}
	return *o.RepayStatus
}

// GetRepayStatusOk returns a tuple with the RepayStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanRepayResponse) GetRepayStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.RepayStatus) {
		return nil, false
	}
	return o.RepayStatus, true
}

// HasRepayStatus returns a boolean if a field has been set.
func (o *VipLoanRepayResponse) HasRepayStatus() bool {
	if o != nil && !common.IsNil(o.RepayStatus) {
		return true
	}

	return false
}

// SetRepayStatus gets a reference to the given string and assigns it to the RepayStatus field.
func (o *VipLoanRepayResponse) SetRepayStatus(v string) {
	o.RepayStatus = &v
}

func (o VipLoanRepayResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VipLoanRepayResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.RepayAmount) {
		toSerialize["repayAmount"] = o.RepayAmount
	}
	if !common.IsNil(o.RemainingPrincipal) {
		toSerialize["remainingPrincipal"] = o.RemainingPrincipal
	}
	if !common.IsNil(o.RemainingInterest) {
		toSerialize["remainingInterest"] = o.RemainingInterest
	}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
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

func (o *VipLoanRepayResponse) UnmarshalJSON(data []byte) (err error) {
	varVipLoanRepayResponse := _VipLoanRepayResponse{}

	err = json.Unmarshal(data, &varVipLoanRepayResponse)

	if err != nil {
		return err
	}

	*o = VipLoanRepayResponse(varVipLoanRepayResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "repayAmount")
		delete(additionalProperties, "remainingPrincipal")
		delete(additionalProperties, "remainingInterest")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "currentLTV")
		delete(additionalProperties, "repayStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVipLoanRepayResponse struct {
	value *VipLoanRepayResponse
	isSet bool
}

func (v NullableVipLoanRepayResponse) Get() *VipLoanRepayResponse {
	return v.value
}

func (v *NullableVipLoanRepayResponse) Set(val *VipLoanRepayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVipLoanRepayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVipLoanRepayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVipLoanRepayResponse(val *VipLoanRepayResponse) *NullableVipLoanRepayResponse {
	return &NullableVipLoanRepayResponse{value: val, isSet: true}
}

func (v NullableVipLoanRepayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVipLoanRepayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
