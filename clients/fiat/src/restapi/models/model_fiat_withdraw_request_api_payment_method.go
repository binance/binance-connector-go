/*
Fiat REST API

Query Binance fiat deposit and withdrawal history.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// FiatWithdrawRequestApiPaymentMethod payment method; current supported: bank_transfer
type FiatWithdrawRequestApiPaymentMethod string

// List of fiatWithdrawRequest_apiPaymentMethod
const (
	FiatWithdrawRequestApiPaymentMethodBankTransfer FiatWithdrawRequestApiPaymentMethod = "bank_transfer"
)

// All allowed values of FiatWithdrawRequestApiPaymentMethod enum
var AllowedFiatWithdrawRequestApiPaymentMethodEnumValues = []FiatWithdrawRequestApiPaymentMethod{
	"bank_transfer",
}

func (v *FiatWithdrawRequestApiPaymentMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FiatWithdrawRequestApiPaymentMethod(value)
	for _, existing := range AllowedFiatWithdrawRequestApiPaymentMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FiatWithdrawRequestApiPaymentMethod", value)
}

// NewFiatWithdrawRequestApiPaymentMethodFromValue returns a pointer to a valid FiatWithdrawRequestApiPaymentMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFiatWithdrawRequestApiPaymentMethodFromValue(v string) (*FiatWithdrawRequestApiPaymentMethod, error) {
	ev := FiatWithdrawRequestApiPaymentMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FiatWithdrawRequestApiPaymentMethod: valid values are %v", v, AllowedFiatWithdrawRequestApiPaymentMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FiatWithdrawRequestApiPaymentMethod) IsValid() bool {
	for _, existing := range AllowedFiatWithdrawRequestApiPaymentMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to fiatWithdrawRequest_apiPaymentMethod value
func (v FiatWithdrawRequestApiPaymentMethod) Ptr() *FiatWithdrawRequestApiPaymentMethod {
	return &v
}

type NullableFiatWithdrawRequestApiPaymentMethod struct {
	value *FiatWithdrawRequestApiPaymentMethod
	isSet bool
}

func (v NullableFiatWithdrawRequestApiPaymentMethod) Get() *FiatWithdrawRequestApiPaymentMethod {
	return v.value
}

func (v *NullableFiatWithdrawRequestApiPaymentMethod) Set(val *FiatWithdrawRequestApiPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableFiatWithdrawRequestApiPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableFiatWithdrawRequestApiPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiatWithdrawRequestApiPaymentMethod(val *FiatWithdrawRequestApiPaymentMethod) *NullableFiatWithdrawRequestApiPaymentMethod {
	return &NullableFiatWithdrawRequestApiPaymentMethod{value: val, isSet: true}
}

func (v NullableFiatWithdrawRequestApiPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiatWithdrawRequestApiPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
