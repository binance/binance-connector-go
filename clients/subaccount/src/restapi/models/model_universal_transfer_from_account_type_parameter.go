/*
Sub Account REST API

Create and manage sub-accounts, control permissions, and transfer assets via the Sub Account API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// UniversalTransferFromAccountTypeParameter the model 'UniversalTransferFromAccountTypeParameter'
type UniversalTransferFromAccountTypeParameter string

// List of universalTransfer_fromAccountType_parameter
const (
	UniversalTransferFromAccountTypeParameterSpot           UniversalTransferFromAccountTypeParameter = "SPOT"
	UniversalTransferFromAccountTypeParameterUsdtFuture     UniversalTransferFromAccountTypeParameter = "USDT_FUTURE"
	UniversalTransferFromAccountTypeParameterCoinFuture     UniversalTransferFromAccountTypeParameter = "COIN_FUTURE"
	UniversalTransferFromAccountTypeParameterMargin         UniversalTransferFromAccountTypeParameter = "MARGIN"
	UniversalTransferFromAccountTypeParameterIsolatedMargin UniversalTransferFromAccountTypeParameter = "ISOLATED_MARGIN"
)

// All allowed values of UniversalTransferFromAccountTypeParameter enum
var AllowedUniversalTransferFromAccountTypeParameterEnumValues = []UniversalTransferFromAccountTypeParameter{
	"SPOT",
	"USDT_FUTURE",
	"COIN_FUTURE",
	"MARGIN",
	"ISOLATED_MARGIN",
}

func (v *UniversalTransferFromAccountTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UniversalTransferFromAccountTypeParameter(value)
	for _, existing := range AllowedUniversalTransferFromAccountTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UniversalTransferFromAccountTypeParameter", value)
}

// NewUniversalTransferFromAccountTypeParameterFromValue returns a pointer to a valid UniversalTransferFromAccountTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUniversalTransferFromAccountTypeParameterFromValue(v string) (*UniversalTransferFromAccountTypeParameter, error) {
	ev := UniversalTransferFromAccountTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UniversalTransferFromAccountTypeParameter: valid values are %v", v, AllowedUniversalTransferFromAccountTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UniversalTransferFromAccountTypeParameter) IsValid() bool {
	for _, existing := range AllowedUniversalTransferFromAccountTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to universalTransfer_fromAccountType_parameter value
func (v UniversalTransferFromAccountTypeParameter) Ptr() *UniversalTransferFromAccountTypeParameter {
	return &v
}

type NullableUniversalTransferFromAccountTypeParameter struct {
	value *UniversalTransferFromAccountTypeParameter
	isSet bool
}

func (v NullableUniversalTransferFromAccountTypeParameter) Get() *UniversalTransferFromAccountTypeParameter {
	return v.value
}

func (v *NullableUniversalTransferFromAccountTypeParameter) Set(val *UniversalTransferFromAccountTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableUniversalTransferFromAccountTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableUniversalTransferFromAccountTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniversalTransferFromAccountTypeParameter(val *UniversalTransferFromAccountTypeParameter) *NullableUniversalTransferFromAccountTypeParameter {
	return &NullableUniversalTransferFromAccountTypeParameter{value: val, isSet: true}
}

func (v NullableUniversalTransferFromAccountTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniversalTransferFromAccountTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
