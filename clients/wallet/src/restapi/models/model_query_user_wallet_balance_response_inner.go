/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryUserWalletBalanceResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUserWalletBalanceResponseInner{}

// QueryUserWalletBalanceResponseInner struct for QueryUserWalletBalanceResponseInner
type QueryUserWalletBalanceResponseInner struct {
	Activate             *bool   `json:"activate,omitempty"`
	Balance              *string `json:"balance,omitempty"`
	WalletName           *string `json:"walletName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryUserWalletBalanceResponseInner QueryUserWalletBalanceResponseInner

// NewQueryUserWalletBalanceResponseInner instantiates a new QueryUserWalletBalanceResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUserWalletBalanceResponseInner() *QueryUserWalletBalanceResponseInner {
	this := QueryUserWalletBalanceResponseInner{}
	return &this
}

// NewQueryUserWalletBalanceResponseInnerWithDefaults instantiates a new QueryUserWalletBalanceResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUserWalletBalanceResponseInnerWithDefaults() *QueryUserWalletBalanceResponseInner {
	this := QueryUserWalletBalanceResponseInner{}
	return &this
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *QueryUserWalletBalanceResponseInner) GetActivate() bool {
	if o == nil || common.IsNil(o.Activate) {
		var ret bool
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserWalletBalanceResponseInner) GetActivateOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Activate) {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *QueryUserWalletBalanceResponseInner) HasActivate() bool {
	if o != nil && !common.IsNil(o.Activate) {
		return true
	}

	return false
}

// SetActivate gets a reference to the given bool and assigns it to the Activate field.
func (o *QueryUserWalletBalanceResponseInner) SetActivate(v bool) {
	o.Activate = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *QueryUserWalletBalanceResponseInner) GetBalance() string {
	if o == nil || common.IsNil(o.Balance) {
		var ret string
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserWalletBalanceResponseInner) GetBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *QueryUserWalletBalanceResponseInner) HasBalance() bool {
	if o != nil && !common.IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given string and assigns it to the Balance field.
func (o *QueryUserWalletBalanceResponseInner) SetBalance(v string) {
	o.Balance = &v
}

// GetWalletName returns the WalletName field value if set, zero value otherwise.
func (o *QueryUserWalletBalanceResponseInner) GetWalletName() string {
	if o == nil || common.IsNil(o.WalletName) {
		var ret string
		return ret
	}
	return *o.WalletName
}

// GetWalletNameOk returns a tuple with the WalletName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserWalletBalanceResponseInner) GetWalletNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletName) {
		return nil, false
	}
	return o.WalletName, true
}

// HasWalletName returns a boolean if a field has been set.
func (o *QueryUserWalletBalanceResponseInner) HasWalletName() bool {
	if o != nil && !common.IsNil(o.WalletName) {
		return true
	}

	return false
}

// SetWalletName gets a reference to the given string and assigns it to the WalletName field.
func (o *QueryUserWalletBalanceResponseInner) SetWalletName(v string) {
	o.WalletName = &v
}

func (o QueryUserWalletBalanceResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUserWalletBalanceResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Activate) {
		toSerialize["activate"] = o.Activate
	}
	if !common.IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !common.IsNil(o.WalletName) {
		toSerialize["walletName"] = o.WalletName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryUserWalletBalanceResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryUserWalletBalanceResponseInner := _QueryUserWalletBalanceResponseInner{}

	err = json.Unmarshal(data, &varQueryUserWalletBalanceResponseInner)

	if err != nil {
		return err
	}

	*o = QueryUserWalletBalanceResponseInner(varQueryUserWalletBalanceResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activate")
		delete(additionalProperties, "balance")
		delete(additionalProperties, "walletName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryUserWalletBalanceResponseInner struct {
	value *QueryUserWalletBalanceResponseInner
	isSet bool
}

func (v NullableQueryUserWalletBalanceResponseInner) Get() *QueryUserWalletBalanceResponseInner {
	return v.value
}

func (v *NullableQueryUserWalletBalanceResponseInner) Set(val *QueryUserWalletBalanceResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUserWalletBalanceResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUserWalletBalanceResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUserWalletBalanceResponseInner(val *QueryUserWalletBalanceResponseInner) *NullableQueryUserWalletBalanceResponseInner {
	return &NullableQueryUserWalletBalanceResponseInner{value: val, isSet: true}
}

func (v NullableQueryUserWalletBalanceResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUserWalletBalanceResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
