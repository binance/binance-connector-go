/*
Binance Dual Investment REST API

OpenAPI Specification for the Binance Dual Investment REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CheckDualInvestmentAccountsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckDualInvestmentAccountsResponse{}

// CheckDualInvestmentAccountsResponse struct for CheckDualInvestmentAccountsResponse
type CheckDualInvestmentAccountsResponse struct {
	TotalAmountInBTC     *string `json:"totalAmountInBTC,omitempty"`
	TotalAmountInUSDT    *string `json:"totalAmountInUSDT,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CheckDualInvestmentAccountsResponse CheckDualInvestmentAccountsResponse

// NewCheckDualInvestmentAccountsResponse instantiates a new CheckDualInvestmentAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckDualInvestmentAccountsResponse() *CheckDualInvestmentAccountsResponse {
	this := CheckDualInvestmentAccountsResponse{}
	return &this
}

// NewCheckDualInvestmentAccountsResponseWithDefaults instantiates a new CheckDualInvestmentAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckDualInvestmentAccountsResponseWithDefaults() *CheckDualInvestmentAccountsResponse {
	this := CheckDualInvestmentAccountsResponse{}
	return &this
}

// GetTotalAmountInBTC returns the TotalAmountInBTC field value if set, zero value otherwise.
func (o *CheckDualInvestmentAccountsResponse) GetTotalAmountInBTC() string {
	if o == nil || common.IsNil(o.TotalAmountInBTC) {
		var ret string
		return ret
	}
	return *o.TotalAmountInBTC
}

// GetTotalAmountInBTCOk returns a tuple with the TotalAmountInBTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckDualInvestmentAccountsResponse) GetTotalAmountInBTCOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAmountInBTC) {
		return nil, false
	}
	return o.TotalAmountInBTC, true
}

// HasTotalAmountInBTC returns a boolean if a field has been set.
func (o *CheckDualInvestmentAccountsResponse) HasTotalAmountInBTC() bool {
	if o != nil && !common.IsNil(o.TotalAmountInBTC) {
		return true
	}

	return false
}

// SetTotalAmountInBTC gets a reference to the given string and assigns it to the TotalAmountInBTC field.
func (o *CheckDualInvestmentAccountsResponse) SetTotalAmountInBTC(v string) {
	o.TotalAmountInBTC = &v
}

// GetTotalAmountInUSDT returns the TotalAmountInUSDT field value if set, zero value otherwise.
func (o *CheckDualInvestmentAccountsResponse) GetTotalAmountInUSDT() string {
	if o == nil || common.IsNil(o.TotalAmountInUSDT) {
		var ret string
		return ret
	}
	return *o.TotalAmountInUSDT
}

// GetTotalAmountInUSDTOk returns a tuple with the TotalAmountInUSDT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckDualInvestmentAccountsResponse) GetTotalAmountInUSDTOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAmountInUSDT) {
		return nil, false
	}
	return o.TotalAmountInUSDT, true
}

// HasTotalAmountInUSDT returns a boolean if a field has been set.
func (o *CheckDualInvestmentAccountsResponse) HasTotalAmountInUSDT() bool {
	if o != nil && !common.IsNil(o.TotalAmountInUSDT) {
		return true
	}

	return false
}

// SetTotalAmountInUSDT gets a reference to the given string and assigns it to the TotalAmountInUSDT field.
func (o *CheckDualInvestmentAccountsResponse) SetTotalAmountInUSDT(v string) {
	o.TotalAmountInUSDT = &v
}

func (o CheckDualInvestmentAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckDualInvestmentAccountsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TotalAmountInBTC) {
		toSerialize["totalAmountInBTC"] = o.TotalAmountInBTC
	}
	if !common.IsNil(o.TotalAmountInUSDT) {
		toSerialize["totalAmountInUSDT"] = o.TotalAmountInUSDT
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CheckDualInvestmentAccountsResponse) UnmarshalJSON(data []byte) (err error) {
	varCheckDualInvestmentAccountsResponse := _CheckDualInvestmentAccountsResponse{}

	err = json.Unmarshal(data, &varCheckDualInvestmentAccountsResponse)

	if err != nil {
		return err
	}

	*o = CheckDualInvestmentAccountsResponse(varCheckDualInvestmentAccountsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "totalAmountInBTC")
		delete(additionalProperties, "totalAmountInUSDT")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCheckDualInvestmentAccountsResponse struct {
	value *CheckDualInvestmentAccountsResponse
	isSet bool
}

func (v NullableCheckDualInvestmentAccountsResponse) Get() *CheckDualInvestmentAccountsResponse {
	return v.value
}

func (v *NullableCheckDualInvestmentAccountsResponse) Set(val *CheckDualInvestmentAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckDualInvestmentAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckDualInvestmentAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckDualInvestmentAccountsResponse(val *CheckDualInvestmentAccountsResponse) *NullableCheckDualInvestmentAccountsResponse {
	return &NullableCheckDualInvestmentAccountsResponse{value: val, isSet: true}
}

func (v NullableCheckDualInvestmentAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckDualInvestmentAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
