/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// TestOrderPositionSideParameter the model 'TestOrderPositionSideParameter'
type TestOrderPositionSideParameter string

// List of testOrder_positionSide_parameter
const (
	TestOrderPositionSideParameterBoth  TestOrderPositionSideParameter = "BOTH"
	TestOrderPositionSideParameterLong  TestOrderPositionSideParameter = "LONG"
	TestOrderPositionSideParameterShort TestOrderPositionSideParameter = "SHORT"
)

// All allowed values of TestOrderPositionSideParameter enum
var AllowedTestOrderPositionSideParameterEnumValues = []TestOrderPositionSideParameter{
	"BOTH",
	"LONG",
	"SHORT",
}

func (v *TestOrderPositionSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TestOrderPositionSideParameter(value)
	for _, existing := range AllowedTestOrderPositionSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TestOrderPositionSideParameter", value)
}

// NewTestOrderPositionSideParameterFromValue returns a pointer to a valid TestOrderPositionSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTestOrderPositionSideParameterFromValue(v string) (*TestOrderPositionSideParameter, error) {
	ev := TestOrderPositionSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TestOrderPositionSideParameter: valid values are %v", v, AllowedTestOrderPositionSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TestOrderPositionSideParameter) IsValid() bool {
	for _, existing := range AllowedTestOrderPositionSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to testOrder_positionSide_parameter value
func (v TestOrderPositionSideParameter) Ptr() *TestOrderPositionSideParameter {
	return &v
}

type NullableTestOrderPositionSideParameter struct {
	value *TestOrderPositionSideParameter
	isSet bool
}

func (v NullableTestOrderPositionSideParameter) Get() *TestOrderPositionSideParameter {
	return v.value
}

func (v *NullableTestOrderPositionSideParameter) Set(val *TestOrderPositionSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableTestOrderPositionSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableTestOrderPositionSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestOrderPositionSideParameter(val *TestOrderPositionSideParameter) *NullableTestOrderPositionSideParameter {
	return &NullableTestOrderPositionSideParameter{value: val, isSet: true}
}

func (v NullableTestOrderPositionSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestOrderPositionSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
