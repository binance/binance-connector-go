/*
Sub Account REST API

Create and manage sub-accounts, control permissions, and transfer assets via the Sub Account API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MovePositionForSubAccountProductTypeParameter the model 'MovePositionForSubAccountProductTypeParameter'
type MovePositionForSubAccountProductTypeParameter string

// List of movePositionForSubAccount_productType_parameter
const (
	MovePositionForSubAccountProductTypeParameterUm MovePositionForSubAccountProductTypeParameter = "UM"
)

// All allowed values of MovePositionForSubAccountProductTypeParameter enum
var AllowedMovePositionForSubAccountProductTypeParameterEnumValues = []MovePositionForSubAccountProductTypeParameter{
	"UM",
}

func (v *MovePositionForSubAccountProductTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MovePositionForSubAccountProductTypeParameter(value)
	for _, existing := range AllowedMovePositionForSubAccountProductTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MovePositionForSubAccountProductTypeParameter", value)
}

// NewMovePositionForSubAccountProductTypeParameterFromValue returns a pointer to a valid MovePositionForSubAccountProductTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMovePositionForSubAccountProductTypeParameterFromValue(v string) (*MovePositionForSubAccountProductTypeParameter, error) {
	ev := MovePositionForSubAccountProductTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MovePositionForSubAccountProductTypeParameter: valid values are %v", v, AllowedMovePositionForSubAccountProductTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MovePositionForSubAccountProductTypeParameter) IsValid() bool {
	for _, existing := range AllowedMovePositionForSubAccountProductTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to movePositionForSubAccount_productType_parameter value
func (v MovePositionForSubAccountProductTypeParameter) Ptr() *MovePositionForSubAccountProductTypeParameter {
	return &v
}

type NullableMovePositionForSubAccountProductTypeParameter struct {
	value *MovePositionForSubAccountProductTypeParameter
	isSet bool
}

func (v NullableMovePositionForSubAccountProductTypeParameter) Get() *MovePositionForSubAccountProductTypeParameter {
	return v.value
}

func (v *NullableMovePositionForSubAccountProductTypeParameter) Set(val *MovePositionForSubAccountProductTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMovePositionForSubAccountProductTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMovePositionForSubAccountProductTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovePositionForSubAccountProductTypeParameter(val *MovePositionForSubAccountProductTypeParameter) *NullableMovePositionForSubAccountProductTypeParameter {
	return &NullableMovePositionForSubAccountProductTypeParameter{value: val, isSet: true}
}

func (v NullableMovePositionForSubAccountProductTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovePositionForSubAccountProductTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
