/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CheckCollateralRepayRateResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckCollateralRepayRateResponse{}

// CheckCollateralRepayRateResponse struct for CheckCollateralRepayRateResponse
type CheckCollateralRepayRateResponse struct {
	LoanCoin             *string `json:"loanCoin,omitempty"`
	CollateralCoin       *string `json:"collateralCoin,omitempty"`
	Rate                 *string `json:"rate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CheckCollateralRepayRateResponse CheckCollateralRepayRateResponse

// NewCheckCollateralRepayRateResponse instantiates a new CheckCollateralRepayRateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckCollateralRepayRateResponse() *CheckCollateralRepayRateResponse {
	this := CheckCollateralRepayRateResponse{}
	return &this
}

// NewCheckCollateralRepayRateResponseWithDefaults instantiates a new CheckCollateralRepayRateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckCollateralRepayRateResponseWithDefaults() *CheckCollateralRepayRateResponse {
	this := CheckCollateralRepayRateResponse{}
	return &this
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *CheckCollateralRepayRateResponse) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckCollateralRepayRateResponse) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *CheckCollateralRepayRateResponse) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *CheckCollateralRepayRateResponse) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *CheckCollateralRepayRateResponse) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckCollateralRepayRateResponse) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *CheckCollateralRepayRateResponse) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *CheckCollateralRepayRateResponse) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *CheckCollateralRepayRateResponse) GetRate() string {
	if o == nil || common.IsNil(o.Rate) {
		var ret string
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckCollateralRepayRateResponse) GetRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *CheckCollateralRepayRateResponse) HasRate() bool {
	if o != nil && !common.IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given string and assigns it to the Rate field.
func (o *CheckCollateralRepayRateResponse) SetRate(v string) {
	o.Rate = &v
}

func (o CheckCollateralRepayRateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckCollateralRepayRateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !common.IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CheckCollateralRepayRateResponse) UnmarshalJSON(data []byte) (err error) {
	varCheckCollateralRepayRateResponse := _CheckCollateralRepayRateResponse{}

	err = json.Unmarshal(data, &varCheckCollateralRepayRateResponse)

	if err != nil {
		return err
	}

	*o = CheckCollateralRepayRateResponse(varCheckCollateralRepayRateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "rate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCheckCollateralRepayRateResponse struct {
	value *CheckCollateralRepayRateResponse
	isSet bool
}

func (v NullableCheckCollateralRepayRateResponse) Get() *CheckCollateralRepayRateResponse {
	return v.value
}

func (v *NullableCheckCollateralRepayRateResponse) Set(val *CheckCollateralRepayRateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckCollateralRepayRateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckCollateralRepayRateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckCollateralRepayRateResponse(val *CheckCollateralRepayRateResponse) *NullableCheckCollateralRepayRateResponse {
	return &NullableCheckCollateralRepayRateResponse{value: val, isSet: true}
}

func (v NullableCheckCollateralRepayRateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckCollateralRepayRateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
