/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FlexibleLoanBorrowResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FlexibleLoanBorrowResponse{}

// FlexibleLoanBorrowResponse struct for FlexibleLoanBorrowResponse
type FlexibleLoanBorrowResponse struct {
	LoanCoin             *string `json:"loanCoin,omitempty"`
	LoanAmount           *string `json:"loanAmount,omitempty"`
	CollateralCoin       *string `json:"collateralCoin,omitempty"`
	CollateralAmount     *string `json:"collateralAmount,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FlexibleLoanBorrowResponse FlexibleLoanBorrowResponse

// NewFlexibleLoanBorrowResponse instantiates a new FlexibleLoanBorrowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlexibleLoanBorrowResponse() *FlexibleLoanBorrowResponse {
	this := FlexibleLoanBorrowResponse{}
	return &this
}

// NewFlexibleLoanBorrowResponseWithDefaults instantiates a new FlexibleLoanBorrowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlexibleLoanBorrowResponseWithDefaults() *FlexibleLoanBorrowResponse {
	this := FlexibleLoanBorrowResponse{}
	return &this
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *FlexibleLoanBorrowResponse) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanBorrowResponse) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *FlexibleLoanBorrowResponse) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *FlexibleLoanBorrowResponse) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetLoanAmount returns the LoanAmount field value if set, zero value otherwise.
func (o *FlexibleLoanBorrowResponse) GetLoanAmount() string {
	if o == nil || common.IsNil(o.LoanAmount) {
		var ret string
		return ret
	}
	return *o.LoanAmount
}

// GetLoanAmountOk returns a tuple with the LoanAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanBorrowResponse) GetLoanAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanAmount) {
		return nil, false
	}
	return o.LoanAmount, true
}

// HasLoanAmount returns a boolean if a field has been set.
func (o *FlexibleLoanBorrowResponse) HasLoanAmount() bool {
	if o != nil && !common.IsNil(o.LoanAmount) {
		return true
	}

	return false
}

// SetLoanAmount gets a reference to the given string and assigns it to the LoanAmount field.
func (o *FlexibleLoanBorrowResponse) SetLoanAmount(v string) {
	o.LoanAmount = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *FlexibleLoanBorrowResponse) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanBorrowResponse) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *FlexibleLoanBorrowResponse) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *FlexibleLoanBorrowResponse) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetCollateralAmount returns the CollateralAmount field value if set, zero value otherwise.
func (o *FlexibleLoanBorrowResponse) GetCollateralAmount() string {
	if o == nil || common.IsNil(o.CollateralAmount) {
		var ret string
		return ret
	}
	return *o.CollateralAmount
}

// GetCollateralAmountOk returns a tuple with the CollateralAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanBorrowResponse) GetCollateralAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralAmount) {
		return nil, false
	}
	return o.CollateralAmount, true
}

// HasCollateralAmount returns a boolean if a field has been set.
func (o *FlexibleLoanBorrowResponse) HasCollateralAmount() bool {
	if o != nil && !common.IsNil(o.CollateralAmount) {
		return true
	}

	return false
}

// SetCollateralAmount gets a reference to the given string and assigns it to the CollateralAmount field.
func (o *FlexibleLoanBorrowResponse) SetCollateralAmount(v string) {
	o.CollateralAmount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FlexibleLoanBorrowResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleLoanBorrowResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FlexibleLoanBorrowResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FlexibleLoanBorrowResponse) SetStatus(v string) {
	o.Status = &v
}

func (o FlexibleLoanBorrowResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlexibleLoanBorrowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.LoanAmount) {
		toSerialize["loanAmount"] = o.LoanAmount
	}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !common.IsNil(o.CollateralAmount) {
		toSerialize["collateralAmount"] = o.CollateralAmount
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FlexibleLoanBorrowResponse) UnmarshalJSON(data []byte) (err error) {
	varFlexibleLoanBorrowResponse := _FlexibleLoanBorrowResponse{}

	err = json.Unmarshal(data, &varFlexibleLoanBorrowResponse)

	if err != nil {
		return err
	}

	*o = FlexibleLoanBorrowResponse(varFlexibleLoanBorrowResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "loanAmount")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "collateralAmount")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFlexibleLoanBorrowResponse struct {
	value *FlexibleLoanBorrowResponse
	isSet bool
}

func (v NullableFlexibleLoanBorrowResponse) Get() *FlexibleLoanBorrowResponse {
	return v.value
}

func (v *NullableFlexibleLoanBorrowResponse) Set(val *FlexibleLoanBorrowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFlexibleLoanBorrowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFlexibleLoanBorrowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlexibleLoanBorrowResponse(val *FlexibleLoanBorrowResponse) *NullableFlexibleLoanBorrowResponse {
	return &NullableFlexibleLoanBorrowResponse{value: val, isSet: true}
}

func (v NullableFlexibleLoanBorrowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlexibleLoanBorrowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
