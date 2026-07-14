/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetCrossMarginTransferHistoryTypeParameter the model 'GetCrossMarginTransferHistoryTypeParameter'
type GetCrossMarginTransferHistoryTypeParameter string

// List of getCrossMarginTransferHistory_type_parameter
const (
	GetCrossMarginTransferHistoryTypeParameterRollIn  GetCrossMarginTransferHistoryTypeParameter = "ROLL_IN"
	GetCrossMarginTransferHistoryTypeParameterRollOut GetCrossMarginTransferHistoryTypeParameter = "ROLL_OUT"
)

// All allowed values of GetCrossMarginTransferHistoryTypeParameter enum
var AllowedGetCrossMarginTransferHistoryTypeParameterEnumValues = []GetCrossMarginTransferHistoryTypeParameter{
	"ROLL_IN",
	"ROLL_OUT",
}

func (v *GetCrossMarginTransferHistoryTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetCrossMarginTransferHistoryTypeParameter(value)
	for _, existing := range AllowedGetCrossMarginTransferHistoryTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetCrossMarginTransferHistoryTypeParameter", value)
}

// NewGetCrossMarginTransferHistoryTypeParameterFromValue returns a pointer to a valid GetCrossMarginTransferHistoryTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetCrossMarginTransferHistoryTypeParameterFromValue(v string) (*GetCrossMarginTransferHistoryTypeParameter, error) {
	ev := GetCrossMarginTransferHistoryTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetCrossMarginTransferHistoryTypeParameter: valid values are %v", v, AllowedGetCrossMarginTransferHistoryTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetCrossMarginTransferHistoryTypeParameter) IsValid() bool {
	for _, existing := range AllowedGetCrossMarginTransferHistoryTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getCrossMarginTransferHistory_type_parameter value
func (v GetCrossMarginTransferHistoryTypeParameter) Ptr() *GetCrossMarginTransferHistoryTypeParameter {
	return &v
}

type NullableGetCrossMarginTransferHistoryTypeParameter struct {
	value *GetCrossMarginTransferHistoryTypeParameter
	isSet bool
}

func (v NullableGetCrossMarginTransferHistoryTypeParameter) Get() *GetCrossMarginTransferHistoryTypeParameter {
	return v.value
}

func (v *NullableGetCrossMarginTransferHistoryTypeParameter) Set(val *GetCrossMarginTransferHistoryTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCrossMarginTransferHistoryTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCrossMarginTransferHistoryTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCrossMarginTransferHistoryTypeParameter(val *GetCrossMarginTransferHistoryTypeParameter) *NullableGetCrossMarginTransferHistoryTypeParameter {
	return &NullableGetCrossMarginTransferHistoryTypeParameter{value: val, isSet: true}
}

func (v NullableGetCrossMarginTransferHistoryTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCrossMarginTransferHistoryTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
