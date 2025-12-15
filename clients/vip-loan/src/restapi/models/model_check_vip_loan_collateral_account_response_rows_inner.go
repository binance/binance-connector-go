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

// checks if the CheckVIPLoanCollateralAccountResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckVIPLoanCollateralAccountResponseRowsInner{}

// CheckVIPLoanCollateralAccountResponseRowsInner struct for CheckVIPLoanCollateralAccountResponseRowsInner
type CheckVIPLoanCollateralAccountResponseRowsInner struct {
	CollateralAccountId  *string `json:"collateralAccountId,omitempty"`
	CollateralCoin       *string `json:"collateralCoin,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CheckVIPLoanCollateralAccountResponseRowsInner CheckVIPLoanCollateralAccountResponseRowsInner

// NewCheckVIPLoanCollateralAccountResponseRowsInner instantiates a new CheckVIPLoanCollateralAccountResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckVIPLoanCollateralAccountResponseRowsInner() *CheckVIPLoanCollateralAccountResponseRowsInner {
	this := CheckVIPLoanCollateralAccountResponseRowsInner{}
	return &this
}

// NewCheckVIPLoanCollateralAccountResponseRowsInnerWithDefaults instantiates a new CheckVIPLoanCollateralAccountResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckVIPLoanCollateralAccountResponseRowsInnerWithDefaults() *CheckVIPLoanCollateralAccountResponseRowsInner {
	this := CheckVIPLoanCollateralAccountResponseRowsInner{}
	return &this
}

// GetCollateralAccountId returns the CollateralAccountId field value if set, zero value otherwise.
func (o *CheckVIPLoanCollateralAccountResponseRowsInner) GetCollateralAccountId() string {
	if o == nil || common.IsNil(o.CollateralAccountId) {
		var ret string
		return ret
	}
	return *o.CollateralAccountId
}

// GetCollateralAccountIdOk returns a tuple with the CollateralAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckVIPLoanCollateralAccountResponseRowsInner) GetCollateralAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralAccountId) {
		return nil, false
	}
	return o.CollateralAccountId, true
}

// HasCollateralAccountId returns a boolean if a field has been set.
func (o *CheckVIPLoanCollateralAccountResponseRowsInner) HasCollateralAccountId() bool {
	if o != nil && !common.IsNil(o.CollateralAccountId) {
		return true
	}

	return false
}

// SetCollateralAccountId gets a reference to the given string and assigns it to the CollateralAccountId field.
func (o *CheckVIPLoanCollateralAccountResponseRowsInner) SetCollateralAccountId(v string) {
	o.CollateralAccountId = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *CheckVIPLoanCollateralAccountResponseRowsInner) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckVIPLoanCollateralAccountResponseRowsInner) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *CheckVIPLoanCollateralAccountResponseRowsInner) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *CheckVIPLoanCollateralAccountResponseRowsInner) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

func (o CheckVIPLoanCollateralAccountResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckVIPLoanCollateralAccountResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CollateralAccountId) {
		toSerialize["collateralAccountId"] = o.CollateralAccountId
	}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CheckVIPLoanCollateralAccountResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varCheckVIPLoanCollateralAccountResponseRowsInner := _CheckVIPLoanCollateralAccountResponseRowsInner{}

	err = json.Unmarshal(data, &varCheckVIPLoanCollateralAccountResponseRowsInner)

	if err != nil {
		return err
	}

	*o = CheckVIPLoanCollateralAccountResponseRowsInner(varCheckVIPLoanCollateralAccountResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "collateralAccountId")
		delete(additionalProperties, "collateralCoin")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCheckVIPLoanCollateralAccountResponseRowsInner struct {
	value *CheckVIPLoanCollateralAccountResponseRowsInner
	isSet bool
}

func (v NullableCheckVIPLoanCollateralAccountResponseRowsInner) Get() *CheckVIPLoanCollateralAccountResponseRowsInner {
	return v.value
}

func (v *NullableCheckVIPLoanCollateralAccountResponseRowsInner) Set(val *CheckVIPLoanCollateralAccountResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckVIPLoanCollateralAccountResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckVIPLoanCollateralAccountResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckVIPLoanCollateralAccountResponseRowsInner(val *CheckVIPLoanCollateralAccountResponseRowsInner) *NullableCheckVIPLoanCollateralAccountResponseRowsInner {
	return &NullableCheckVIPLoanCollateralAccountResponseRowsInner{value: val, isSet: true}
}

func (v NullableCheckVIPLoanCollateralAccountResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckVIPLoanCollateralAccountResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
