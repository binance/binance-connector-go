/*
Fiat REST API

Query Binance fiat deposit and withdrawal history.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// DepositRequestApiPaymentMethod payment method; current supported: pix
type DepositRequestApiPaymentMethod string

// List of depositRequest_apiPaymentMethod
const (
	DepositRequestApiPaymentMethodPix DepositRequestApiPaymentMethod = "pix"
)

// All allowed values of DepositRequestApiPaymentMethod enum
var AllowedDepositRequestApiPaymentMethodEnumValues = []DepositRequestApiPaymentMethod{
	"pix",
}

func (v *DepositRequestApiPaymentMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DepositRequestApiPaymentMethod(value)
	for _, existing := range AllowedDepositRequestApiPaymentMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DepositRequestApiPaymentMethod", value)
}

// NewDepositRequestApiPaymentMethodFromValue returns a pointer to a valid DepositRequestApiPaymentMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDepositRequestApiPaymentMethodFromValue(v string) (*DepositRequestApiPaymentMethod, error) {
	ev := DepositRequestApiPaymentMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DepositRequestApiPaymentMethod: valid values are %v", v, AllowedDepositRequestApiPaymentMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DepositRequestApiPaymentMethod) IsValid() bool {
	for _, existing := range AllowedDepositRequestApiPaymentMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to depositRequest_apiPaymentMethod value
func (v DepositRequestApiPaymentMethod) Ptr() *DepositRequestApiPaymentMethod {
	return &v
}

type NullableDepositRequestApiPaymentMethod struct {
	value *DepositRequestApiPaymentMethod
	isSet bool
}

func (v NullableDepositRequestApiPaymentMethod) Get() *DepositRequestApiPaymentMethod {
	return v.value
}

func (v *NullableDepositRequestApiPaymentMethod) Set(val *DepositRequestApiPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositRequestApiPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositRequestApiPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositRequestApiPaymentMethod(val *DepositRequestApiPaymentMethod) *NullableDepositRequestApiPaymentMethod {
	return &NullableDepositRequestApiPaymentMethod{value: val, isSet: true}
}

func (v NullableDepositRequestApiPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositRequestApiPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
