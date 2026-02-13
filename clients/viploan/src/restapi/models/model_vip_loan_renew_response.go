/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the VipLoanRenewResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &VipLoanRenewResponse{}

// VipLoanRenewResponse struct for VipLoanRenewResponse
type VipLoanRenewResponse struct {
	LoanAccountId        *string `json:"loanAccountId,omitempty"`
	LoanCoin             *string `json:"loanCoin,omitempty"`
	LoanAmount           *string `json:"loanAmount,omitempty"`
	CollateralAccountId  *string `json:"collateralAccountId,omitempty"`
	CollateralCoin       *string `json:"collateralCoin,omitempty"`
	LoanTerm             *string `json:"loanTerm,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VipLoanRenewResponse VipLoanRenewResponse

// NewVipLoanRenewResponse instantiates a new VipLoanRenewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVipLoanRenewResponse() *VipLoanRenewResponse {
	this := VipLoanRenewResponse{}
	return &this
}

// NewVipLoanRenewResponseWithDefaults instantiates a new VipLoanRenewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVipLoanRenewResponseWithDefaults() *VipLoanRenewResponse {
	this := VipLoanRenewResponse{}
	return &this
}

// GetLoanAccountId returns the LoanAccountId field value if set, zero value otherwise.
func (o *VipLoanRenewResponse) GetLoanAccountId() string {
	if o == nil || common.IsNil(o.LoanAccountId) {
		var ret string
		return ret
	}
	return *o.LoanAccountId
}

// GetLoanAccountIdOk returns a tuple with the LoanAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanRenewResponse) GetLoanAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanAccountId) {
		return nil, false
	}
	return o.LoanAccountId, true
}

// HasLoanAccountId returns a boolean if a field has been set.
func (o *VipLoanRenewResponse) HasLoanAccountId() bool {
	if o != nil && !common.IsNil(o.LoanAccountId) {
		return true
	}

	return false
}

// SetLoanAccountId gets a reference to the given string and assigns it to the LoanAccountId field.
func (o *VipLoanRenewResponse) SetLoanAccountId(v string) {
	o.LoanAccountId = &v
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *VipLoanRenewResponse) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanRenewResponse) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *VipLoanRenewResponse) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *VipLoanRenewResponse) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetLoanAmount returns the LoanAmount field value if set, zero value otherwise.
func (o *VipLoanRenewResponse) GetLoanAmount() string {
	if o == nil || common.IsNil(o.LoanAmount) {
		var ret string
		return ret
	}
	return *o.LoanAmount
}

// GetLoanAmountOk returns a tuple with the LoanAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanRenewResponse) GetLoanAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanAmount) {
		return nil, false
	}
	return o.LoanAmount, true
}

// HasLoanAmount returns a boolean if a field has been set.
func (o *VipLoanRenewResponse) HasLoanAmount() bool {
	if o != nil && !common.IsNil(o.LoanAmount) {
		return true
	}

	return false
}

// SetLoanAmount gets a reference to the given string and assigns it to the LoanAmount field.
func (o *VipLoanRenewResponse) SetLoanAmount(v string) {
	o.LoanAmount = &v
}

// GetCollateralAccountId returns the CollateralAccountId field value if set, zero value otherwise.
func (o *VipLoanRenewResponse) GetCollateralAccountId() string {
	if o == nil || common.IsNil(o.CollateralAccountId) {
		var ret string
		return ret
	}
	return *o.CollateralAccountId
}

// GetCollateralAccountIdOk returns a tuple with the CollateralAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanRenewResponse) GetCollateralAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralAccountId) {
		return nil, false
	}
	return o.CollateralAccountId, true
}

// HasCollateralAccountId returns a boolean if a field has been set.
func (o *VipLoanRenewResponse) HasCollateralAccountId() bool {
	if o != nil && !common.IsNil(o.CollateralAccountId) {
		return true
	}

	return false
}

// SetCollateralAccountId gets a reference to the given string and assigns it to the CollateralAccountId field.
func (o *VipLoanRenewResponse) SetCollateralAccountId(v string) {
	o.CollateralAccountId = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *VipLoanRenewResponse) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanRenewResponse) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *VipLoanRenewResponse) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *VipLoanRenewResponse) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetLoanTerm returns the LoanTerm field value if set, zero value otherwise.
func (o *VipLoanRenewResponse) GetLoanTerm() string {
	if o == nil || common.IsNil(o.LoanTerm) {
		var ret string
		return ret
	}
	return *o.LoanTerm
}

// GetLoanTermOk returns a tuple with the LoanTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanRenewResponse) GetLoanTermOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanTerm) {
		return nil, false
	}
	return o.LoanTerm, true
}

// HasLoanTerm returns a boolean if a field has been set.
func (o *VipLoanRenewResponse) HasLoanTerm() bool {
	if o != nil && !common.IsNil(o.LoanTerm) {
		return true
	}

	return false
}

// SetLoanTerm gets a reference to the given string and assigns it to the LoanTerm field.
func (o *VipLoanRenewResponse) SetLoanTerm(v string) {
	o.LoanTerm = &v
}

func (o VipLoanRenewResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VipLoanRenewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanAccountId) {
		toSerialize["loanAccountId"] = o.LoanAccountId
	}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.LoanAmount) {
		toSerialize["loanAmount"] = o.LoanAmount
	}
	if !common.IsNil(o.CollateralAccountId) {
		toSerialize["collateralAccountId"] = o.CollateralAccountId
	}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !common.IsNil(o.LoanTerm) {
		toSerialize["loanTerm"] = o.LoanTerm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VipLoanRenewResponse) UnmarshalJSON(data []byte) (err error) {
	varVipLoanRenewResponse := _VipLoanRenewResponse{}

	err = json.Unmarshal(data, &varVipLoanRenewResponse)

	if err != nil {
		return err
	}

	*o = VipLoanRenewResponse(varVipLoanRenewResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanAccountId")
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "loanAmount")
		delete(additionalProperties, "collateralAccountId")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "loanTerm")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVipLoanRenewResponse struct {
	value *VipLoanRenewResponse
	isSet bool
}

func (v NullableVipLoanRenewResponse) Get() *VipLoanRenewResponse {
	return v.value
}

func (v *NullableVipLoanRenewResponse) Set(val *VipLoanRenewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVipLoanRenewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVipLoanRenewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVipLoanRenewResponse(val *VipLoanRenewResponse) *NullableVipLoanRenewResponse {
	return &NullableVipLoanRenewResponse{value: val, isSet: true}
}

func (v NullableVipLoanRenewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVipLoanRenewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
