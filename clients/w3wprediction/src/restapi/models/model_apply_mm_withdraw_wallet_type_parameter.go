/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ApplyMmWithdrawWalletTypeParameter the model 'ApplyMmWithdrawWalletTypeParameter'
type ApplyMmWithdrawWalletTypeParameter int32

// List of applyMmWithdraw_walletType_parameter
const (
	ApplyMmWithdrawWalletTypeParameterWalletType0 ApplyMmWithdrawWalletTypeParameter = 0
	ApplyMmWithdrawWalletTypeParameterWalletType1 ApplyMmWithdrawWalletTypeParameter = 1
)

// All allowed values of ApplyMmWithdrawWalletTypeParameter enum
var AllowedApplyMmWithdrawWalletTypeParameterEnumValues = []ApplyMmWithdrawWalletTypeParameter{
	0,
	1,
}

func (v *ApplyMmWithdrawWalletTypeParameter) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApplyMmWithdrawWalletTypeParameter(value)
	for _, existing := range AllowedApplyMmWithdrawWalletTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApplyMmWithdrawWalletTypeParameter", value)
}

// NewApplyMmWithdrawWalletTypeParameterFromValue returns a pointer to a valid ApplyMmWithdrawWalletTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApplyMmWithdrawWalletTypeParameterFromValue(v int32) (*ApplyMmWithdrawWalletTypeParameter, error) {
	ev := ApplyMmWithdrawWalletTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApplyMmWithdrawWalletTypeParameter: valid values are %v", v, AllowedApplyMmWithdrawWalletTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApplyMmWithdrawWalletTypeParameter) IsValid() bool {
	for _, existing := range AllowedApplyMmWithdrawWalletTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to applyMmWithdraw_walletType_parameter value
func (v ApplyMmWithdrawWalletTypeParameter) Ptr() *ApplyMmWithdrawWalletTypeParameter {
	return &v
}

type NullableApplyMmWithdrawWalletTypeParameter struct {
	value *ApplyMmWithdrawWalletTypeParameter
	isSet bool
}

func (v NullableApplyMmWithdrawWalletTypeParameter) Get() *ApplyMmWithdrawWalletTypeParameter {
	return v.value
}

func (v *NullableApplyMmWithdrawWalletTypeParameter) Set(val *ApplyMmWithdrawWalletTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyMmWithdrawWalletTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyMmWithdrawWalletTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyMmWithdrawWalletTypeParameter(val *ApplyMmWithdrawWalletTypeParameter) *NullableApplyMmWithdrawWalletTypeParameter {
	return &NullableApplyMmWithdrawWalletTypeParameter{value: val, isSet: true}
}

func (v NullableApplyMmWithdrawWalletTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyMmWithdrawWalletTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
