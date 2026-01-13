/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the VipLoanBorrowResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &VipLoanBorrowResponse{}

// VipLoanBorrowResponse struct for VipLoanBorrowResponse
type VipLoanBorrowResponse struct {
	LoanAccountId        *string `json:"loanAccountId,omitempty"`
	RequestId            *string `json:"requestId,omitempty"`
	LoanCoin             *string `json:"loanCoin,omitempty"`
	IsFlexibleRate       *string `json:"isFlexibleRate,omitempty"`
	LoanAmount           *string `json:"loanAmount,omitempty"`
	CollateralAccountId  *string `json:"collateralAccountId,omitempty"`
	CollateralCoin       *string `json:"collateralCoin,omitempty"`
	LoanTerm             *string `json:"loanTerm,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VipLoanBorrowResponse VipLoanBorrowResponse

// NewVipLoanBorrowResponse instantiates a new VipLoanBorrowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVipLoanBorrowResponse() *VipLoanBorrowResponse {
	this := VipLoanBorrowResponse{}
	return &this
}

// NewVipLoanBorrowResponseWithDefaults instantiates a new VipLoanBorrowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVipLoanBorrowResponseWithDefaults() *VipLoanBorrowResponse {
	this := VipLoanBorrowResponse{}
	return &this
}

// GetLoanAccountId returns the LoanAccountId field value if set, zero value otherwise.
func (o *VipLoanBorrowResponse) GetLoanAccountId() string {
	if o == nil || common.IsNil(o.LoanAccountId) {
		var ret string
		return ret
	}
	return *o.LoanAccountId
}

// GetLoanAccountIdOk returns a tuple with the LoanAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanBorrowResponse) GetLoanAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanAccountId) {
		return nil, false
	}
	return o.LoanAccountId, true
}

// HasLoanAccountId returns a boolean if a field has been set.
func (o *VipLoanBorrowResponse) HasLoanAccountId() bool {
	if o != nil && !common.IsNil(o.LoanAccountId) {
		return true
	}

	return false
}

// SetLoanAccountId gets a reference to the given string and assigns it to the LoanAccountId field.
func (o *VipLoanBorrowResponse) SetLoanAccountId(v string) {
	o.LoanAccountId = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *VipLoanBorrowResponse) GetRequestId() string {
	if o == nil || common.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanBorrowResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *VipLoanBorrowResponse) HasRequestId() bool {
	if o != nil && !common.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *VipLoanBorrowResponse) SetRequestId(v string) {
	o.RequestId = &v
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *VipLoanBorrowResponse) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanBorrowResponse) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *VipLoanBorrowResponse) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *VipLoanBorrowResponse) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetIsFlexibleRate returns the IsFlexibleRate field value if set, zero value otherwise.
func (o *VipLoanBorrowResponse) GetIsFlexibleRate() string {
	if o == nil || common.IsNil(o.IsFlexibleRate) {
		var ret string
		return ret
	}
	return *o.IsFlexibleRate
}

// GetIsFlexibleRateOk returns a tuple with the IsFlexibleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanBorrowResponse) GetIsFlexibleRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsFlexibleRate) {
		return nil, false
	}
	return o.IsFlexibleRate, true
}

// HasIsFlexibleRate returns a boolean if a field has been set.
func (o *VipLoanBorrowResponse) HasIsFlexibleRate() bool {
	if o != nil && !common.IsNil(o.IsFlexibleRate) {
		return true
	}

	return false
}

// SetIsFlexibleRate gets a reference to the given string and assigns it to the IsFlexibleRate field.
func (o *VipLoanBorrowResponse) SetIsFlexibleRate(v string) {
	o.IsFlexibleRate = &v
}

// GetLoanAmount returns the LoanAmount field value if set, zero value otherwise.
func (o *VipLoanBorrowResponse) GetLoanAmount() string {
	if o == nil || common.IsNil(o.LoanAmount) {
		var ret string
		return ret
	}
	return *o.LoanAmount
}

// GetLoanAmountOk returns a tuple with the LoanAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanBorrowResponse) GetLoanAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanAmount) {
		return nil, false
	}
	return o.LoanAmount, true
}

// HasLoanAmount returns a boolean if a field has been set.
func (o *VipLoanBorrowResponse) HasLoanAmount() bool {
	if o != nil && !common.IsNil(o.LoanAmount) {
		return true
	}

	return false
}

// SetLoanAmount gets a reference to the given string and assigns it to the LoanAmount field.
func (o *VipLoanBorrowResponse) SetLoanAmount(v string) {
	o.LoanAmount = &v
}

// GetCollateralAccountId returns the CollateralAccountId field value if set, zero value otherwise.
func (o *VipLoanBorrowResponse) GetCollateralAccountId() string {
	if o == nil || common.IsNil(o.CollateralAccountId) {
		var ret string
		return ret
	}
	return *o.CollateralAccountId
}

// GetCollateralAccountIdOk returns a tuple with the CollateralAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanBorrowResponse) GetCollateralAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralAccountId) {
		return nil, false
	}
	return o.CollateralAccountId, true
}

// HasCollateralAccountId returns a boolean if a field has been set.
func (o *VipLoanBorrowResponse) HasCollateralAccountId() bool {
	if o != nil && !common.IsNil(o.CollateralAccountId) {
		return true
	}

	return false
}

// SetCollateralAccountId gets a reference to the given string and assigns it to the CollateralAccountId field.
func (o *VipLoanBorrowResponse) SetCollateralAccountId(v string) {
	o.CollateralAccountId = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *VipLoanBorrowResponse) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanBorrowResponse) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *VipLoanBorrowResponse) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *VipLoanBorrowResponse) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetLoanTerm returns the LoanTerm field value if set, zero value otherwise.
func (o *VipLoanBorrowResponse) GetLoanTerm() string {
	if o == nil || common.IsNil(o.LoanTerm) {
		var ret string
		return ret
	}
	return *o.LoanTerm
}

// GetLoanTermOk returns a tuple with the LoanTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipLoanBorrowResponse) GetLoanTermOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanTerm) {
		return nil, false
	}
	return o.LoanTerm, true
}

// HasLoanTerm returns a boolean if a field has been set.
func (o *VipLoanBorrowResponse) HasLoanTerm() bool {
	if o != nil && !common.IsNil(o.LoanTerm) {
		return true
	}

	return false
}

// SetLoanTerm gets a reference to the given string and assigns it to the LoanTerm field.
func (o *VipLoanBorrowResponse) SetLoanTerm(v string) {
	o.LoanTerm = &v
}

func (o VipLoanBorrowResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VipLoanBorrowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanAccountId) {
		toSerialize["loanAccountId"] = o.LoanAccountId
	}
	if !common.IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.IsFlexibleRate) {
		toSerialize["isFlexibleRate"] = o.IsFlexibleRate
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

func (o *VipLoanBorrowResponse) UnmarshalJSON(data []byte) (err error) {
	varVipLoanBorrowResponse := _VipLoanBorrowResponse{}

	err = json.Unmarshal(data, &varVipLoanBorrowResponse)

	if err != nil {
		return err
	}

	*o = VipLoanBorrowResponse(varVipLoanBorrowResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanAccountId")
		delete(additionalProperties, "requestId")
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "isFlexibleRate")
		delete(additionalProperties, "loanAmount")
		delete(additionalProperties, "collateralAccountId")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "loanTerm")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVipLoanBorrowResponse struct {
	value *VipLoanBorrowResponse
	isSet bool
}

func (v NullableVipLoanBorrowResponse) Get() *VipLoanBorrowResponse {
	return v.value
}

func (v *NullableVipLoanBorrowResponse) Set(val *VipLoanBorrowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVipLoanBorrowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVipLoanBorrowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVipLoanBorrowResponse(val *VipLoanBorrowResponse) *NullableVipLoanBorrowResponse {
	return &NullableVipLoanBorrowResponse{value: val, isSet: true}
}

func (v NullableVipLoanBorrowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVipLoanBorrowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
