/*
Fiat REST API

Query Binance fiat deposit and withdrawal history.
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the FiatWithdrawRequestAccountInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FiatWithdrawRequestAccountInfo{}

// FiatWithdrawRequestAccountInfo withdraw account info
type FiatWithdrawRequestAccountInfo struct {
	// Your destination bank account number is required to receive the withdrawal. In Argentina, this will be your CBU/CVU; in Mexico, it will be your CLABE.
	AccountNumber string `json:"accountNumber"`
	// Bank agency code. If contains a hyphen (e.g. `123-4`), enter `123` only.
	Agency *string `json:"agency,omitempty"`
	// Bank code used for PIX routing.
	BankCodeForPix *string `json:"bankCodeForPix,omitempty"`
	// Account type, e.g. `current` (Checking Account), `saving` (Savings Account), etc.
	AccountType          *string `json:"accountType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FiatWithdrawRequestAccountInfo FiatWithdrawRequestAccountInfo

// NewFiatWithdrawRequestAccountInfo instantiates a new FiatWithdrawRequestAccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiatWithdrawRequestAccountInfo(accountNumber string) *FiatWithdrawRequestAccountInfo {
	this := FiatWithdrawRequestAccountInfo{}
	this.AccountNumber = accountNumber
	return &this
}

// NewFiatWithdrawRequestAccountInfoWithDefaults instantiates a new FiatWithdrawRequestAccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiatWithdrawRequestAccountInfoWithDefaults() *FiatWithdrawRequestAccountInfo {
	this := FiatWithdrawRequestAccountInfo{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *FiatWithdrawRequestAccountInfo) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *FiatWithdrawRequestAccountInfo) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *FiatWithdrawRequestAccountInfo) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetAgency returns the Agency field value if set, zero value otherwise.
func (o *FiatWithdrawRequestAccountInfo) GetAgency() string {
	if o == nil || common.IsNil(o.Agency) {
		var ret string
		return ret
	}
	return *o.Agency
}

// GetAgencyOk returns a tuple with the Agency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatWithdrawRequestAccountInfo) GetAgencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Agency) {
		return nil, false
	}
	return o.Agency, true
}

// HasAgency returns a boolean if a field has been set.
func (o *FiatWithdrawRequestAccountInfo) HasAgency() bool {
	if o != nil && !common.IsNil(o.Agency) {
		return true
	}

	return false
}

// SetAgency gets a reference to the given string and assigns it to the Agency field.
func (o *FiatWithdrawRequestAccountInfo) SetAgency(v string) {
	o.Agency = &v
}

// GetBankCodeForPix returns the BankCodeForPix field value if set, zero value otherwise.
func (o *FiatWithdrawRequestAccountInfo) GetBankCodeForPix() string {
	if o == nil || common.IsNil(o.BankCodeForPix) {
		var ret string
		return ret
	}
	return *o.BankCodeForPix
}

// GetBankCodeForPixOk returns a tuple with the BankCodeForPix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatWithdrawRequestAccountInfo) GetBankCodeForPixOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankCodeForPix) {
		return nil, false
	}
	return o.BankCodeForPix, true
}

// HasBankCodeForPix returns a boolean if a field has been set.
func (o *FiatWithdrawRequestAccountInfo) HasBankCodeForPix() bool {
	if o != nil && !common.IsNil(o.BankCodeForPix) {
		return true
	}

	return false
}

// SetBankCodeForPix gets a reference to the given string and assigns it to the BankCodeForPix field.
func (o *FiatWithdrawRequestAccountInfo) SetBankCodeForPix(v string) {
	o.BankCodeForPix = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *FiatWithdrawRequestAccountInfo) GetAccountType() string {
	if o == nil || common.IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatWithdrawRequestAccountInfo) GetAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *FiatWithdrawRequestAccountInfo) HasAccountType() bool {
	if o != nil && !common.IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *FiatWithdrawRequestAccountInfo) SetAccountType(v string) {
	o.AccountType = &v
}

func (o FiatWithdrawRequestAccountInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FiatWithdrawRequestAccountInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	if !common.IsNil(o.Agency) {
		toSerialize["agency"] = o.Agency
	}
	if !common.IsNil(o.BankCodeForPix) {
		toSerialize["bankCodeForPix"] = o.BankCodeForPix
	}
	if !common.IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FiatWithdrawRequestAccountInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountNumber",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFiatWithdrawRequestAccountInfo := _FiatWithdrawRequestAccountInfo{}

	err = json.Unmarshal(data, &varFiatWithdrawRequestAccountInfo)

	if err != nil {
		return err
	}

	*o = FiatWithdrawRequestAccountInfo(varFiatWithdrawRequestAccountInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountNumber")
		delete(additionalProperties, "agency")
		delete(additionalProperties, "bankCodeForPix")
		delete(additionalProperties, "accountType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFiatWithdrawRequestAccountInfo struct {
	value *FiatWithdrawRequestAccountInfo
	isSet bool
}

func (v NullableFiatWithdrawRequestAccountInfo) Get() *FiatWithdrawRequestAccountInfo {
	return v.value
}

func (v *NullableFiatWithdrawRequestAccountInfo) Set(val *FiatWithdrawRequestAccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFiatWithdrawRequestAccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFiatWithdrawRequestAccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiatWithdrawRequestAccountInfo(val *FiatWithdrawRequestAccountInfo) *NullableFiatWithdrawRequestAccountInfo {
	return &NullableFiatWithdrawRequestAccountInfo{value: val, isSet: true}
}

func (v NullableFiatWithdrawRequestAccountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiatWithdrawRequestAccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
