/*
Sub Account REST API

Create and manage sub-accounts, control permissions, and transfer assets via the Sub Account API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetMovePositionHistoryForSubAccountProductTypeParameter the model 'GetMovePositionHistoryForSubAccountProductTypeParameter'
type GetMovePositionHistoryForSubAccountProductTypeParameter string

// List of getMovePositionHistoryForSubAccount_productType_parameter
const (
	GetMovePositionHistoryForSubAccountProductTypeParameterUm     GetMovePositionHistoryForSubAccountProductTypeParameter = "UM"
	GetMovePositionHistoryForSubAccountProductTypeParameterOption GetMovePositionHistoryForSubAccountProductTypeParameter = "OPTION"
)

// All allowed values of GetMovePositionHistoryForSubAccountProductTypeParameter enum
var AllowedGetMovePositionHistoryForSubAccountProductTypeParameterEnumValues = []GetMovePositionHistoryForSubAccountProductTypeParameter{
	"UM",
	"OPTION",
}

func (v *GetMovePositionHistoryForSubAccountProductTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetMovePositionHistoryForSubAccountProductTypeParameter(value)
	for _, existing := range AllowedGetMovePositionHistoryForSubAccountProductTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetMovePositionHistoryForSubAccountProductTypeParameter", value)
}

// NewGetMovePositionHistoryForSubAccountProductTypeParameterFromValue returns a pointer to a valid GetMovePositionHistoryForSubAccountProductTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetMovePositionHistoryForSubAccountProductTypeParameterFromValue(v string) (*GetMovePositionHistoryForSubAccountProductTypeParameter, error) {
	ev := GetMovePositionHistoryForSubAccountProductTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetMovePositionHistoryForSubAccountProductTypeParameter: valid values are %v", v, AllowedGetMovePositionHistoryForSubAccountProductTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetMovePositionHistoryForSubAccountProductTypeParameter) IsValid() bool {
	for _, existing := range AllowedGetMovePositionHistoryForSubAccountProductTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getMovePositionHistoryForSubAccount_productType_parameter value
func (v GetMovePositionHistoryForSubAccountProductTypeParameter) Ptr() *GetMovePositionHistoryForSubAccountProductTypeParameter {
	return &v
}

type NullableGetMovePositionHistoryForSubAccountProductTypeParameter struct {
	value *GetMovePositionHistoryForSubAccountProductTypeParameter
	isSet bool
}

func (v NullableGetMovePositionHistoryForSubAccountProductTypeParameter) Get() *GetMovePositionHistoryForSubAccountProductTypeParameter {
	return v.value
}

func (v *NullableGetMovePositionHistoryForSubAccountProductTypeParameter) Set(val *GetMovePositionHistoryForSubAccountProductTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMovePositionHistoryForSubAccountProductTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMovePositionHistoryForSubAccountProductTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMovePositionHistoryForSubAccountProductTypeParameter(val *GetMovePositionHistoryForSubAccountProductTypeParameter) *NullableGetMovePositionHistoryForSubAccountProductTypeParameter {
	return &NullableGetMovePositionHistoryForSubAccountProductTypeParameter{value: val, isSet: true}
}

func (v NullableGetMovePositionHistoryForSubAccountProductTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMovePositionHistoryForSubAccountProductTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
