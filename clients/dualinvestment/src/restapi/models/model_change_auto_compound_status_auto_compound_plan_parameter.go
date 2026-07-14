/*
Dual Investment REST API

Query products, request quotes, and subscribe to Advanced Earn Dual Investment strategies.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ChangeAutoCompoundStatusAutoCompoundPlanParameter the model 'ChangeAutoCompoundStatusAutoCompoundPlanParameter'
type ChangeAutoCompoundStatusAutoCompoundPlanParameter string

// List of changeAutoCompoundStatus_autoCompoundPlan_parameter
const (
	ChangeAutoCompoundStatusAutoCompoundPlanParameterNone     ChangeAutoCompoundStatusAutoCompoundPlanParameter = "NONE"
	ChangeAutoCompoundStatusAutoCompoundPlanParameterStandard ChangeAutoCompoundStatusAutoCompoundPlanParameter = "STANDARD"
	ChangeAutoCompoundStatusAutoCompoundPlanParameterAdvanced ChangeAutoCompoundStatusAutoCompoundPlanParameter = "ADVANCED"
)

// All allowed values of ChangeAutoCompoundStatusAutoCompoundPlanParameter enum
var AllowedChangeAutoCompoundStatusAutoCompoundPlanParameterEnumValues = []ChangeAutoCompoundStatusAutoCompoundPlanParameter{
	"NONE",
	"STANDARD",
	"ADVANCED",
}

func (v *ChangeAutoCompoundStatusAutoCompoundPlanParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChangeAutoCompoundStatusAutoCompoundPlanParameter(value)
	for _, existing := range AllowedChangeAutoCompoundStatusAutoCompoundPlanParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChangeAutoCompoundStatusAutoCompoundPlanParameter", value)
}

// NewChangeAutoCompoundStatusAutoCompoundPlanParameterFromValue returns a pointer to a valid ChangeAutoCompoundStatusAutoCompoundPlanParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChangeAutoCompoundStatusAutoCompoundPlanParameterFromValue(v string) (*ChangeAutoCompoundStatusAutoCompoundPlanParameter, error) {
	ev := ChangeAutoCompoundStatusAutoCompoundPlanParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChangeAutoCompoundStatusAutoCompoundPlanParameter: valid values are %v", v, AllowedChangeAutoCompoundStatusAutoCompoundPlanParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChangeAutoCompoundStatusAutoCompoundPlanParameter) IsValid() bool {
	for _, existing := range AllowedChangeAutoCompoundStatusAutoCompoundPlanParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to changeAutoCompoundStatus_autoCompoundPlan_parameter value
func (v ChangeAutoCompoundStatusAutoCompoundPlanParameter) Ptr() *ChangeAutoCompoundStatusAutoCompoundPlanParameter {
	return &v
}

type NullableChangeAutoCompoundStatusAutoCompoundPlanParameter struct {
	value *ChangeAutoCompoundStatusAutoCompoundPlanParameter
	isSet bool
}

func (v NullableChangeAutoCompoundStatusAutoCompoundPlanParameter) Get() *ChangeAutoCompoundStatusAutoCompoundPlanParameter {
	return v.value
}

func (v *NullableChangeAutoCompoundStatusAutoCompoundPlanParameter) Set(val *ChangeAutoCompoundStatusAutoCompoundPlanParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeAutoCompoundStatusAutoCompoundPlanParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeAutoCompoundStatusAutoCompoundPlanParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeAutoCompoundStatusAutoCompoundPlanParameter(val *ChangeAutoCompoundStatusAutoCompoundPlanParameter) *NullableChangeAutoCompoundStatusAutoCompoundPlanParameter {
	return &NullableChangeAutoCompoundStatusAutoCompoundPlanParameter{value: val, isSet: true}
}

func (v NullableChangeAutoCompoundStatusAutoCompoundPlanParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeAutoCompoundStatusAutoCompoundPlanParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
