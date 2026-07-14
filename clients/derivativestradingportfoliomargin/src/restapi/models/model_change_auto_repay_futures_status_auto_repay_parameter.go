/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ChangeAutoRepayFuturesStatusAutoRepayParameter the model 'ChangeAutoRepayFuturesStatusAutoRepayParameter'
type ChangeAutoRepayFuturesStatusAutoRepayParameter string

// List of changeAutoRepayFuturesStatus_autoRepay_parameter
const (
	ChangeAutoRepayFuturesStatusAutoRepayParameterTrue  ChangeAutoRepayFuturesStatusAutoRepayParameter = "true"
	ChangeAutoRepayFuturesStatusAutoRepayParameterFalse ChangeAutoRepayFuturesStatusAutoRepayParameter = "false"
)

// All allowed values of ChangeAutoRepayFuturesStatusAutoRepayParameter enum
var AllowedChangeAutoRepayFuturesStatusAutoRepayParameterEnumValues = []ChangeAutoRepayFuturesStatusAutoRepayParameter{
	"true",
	"false",
}

func (v *ChangeAutoRepayFuturesStatusAutoRepayParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChangeAutoRepayFuturesStatusAutoRepayParameter(value)
	for _, existing := range AllowedChangeAutoRepayFuturesStatusAutoRepayParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChangeAutoRepayFuturesStatusAutoRepayParameter", value)
}

// NewChangeAutoRepayFuturesStatusAutoRepayParameterFromValue returns a pointer to a valid ChangeAutoRepayFuturesStatusAutoRepayParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChangeAutoRepayFuturesStatusAutoRepayParameterFromValue(v string) (*ChangeAutoRepayFuturesStatusAutoRepayParameter, error) {
	ev := ChangeAutoRepayFuturesStatusAutoRepayParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChangeAutoRepayFuturesStatusAutoRepayParameter: valid values are %v", v, AllowedChangeAutoRepayFuturesStatusAutoRepayParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChangeAutoRepayFuturesStatusAutoRepayParameter) IsValid() bool {
	for _, existing := range AllowedChangeAutoRepayFuturesStatusAutoRepayParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to changeAutoRepayFuturesStatus_autoRepay_parameter value
func (v ChangeAutoRepayFuturesStatusAutoRepayParameter) Ptr() *ChangeAutoRepayFuturesStatusAutoRepayParameter {
	return &v
}

type NullableChangeAutoRepayFuturesStatusAutoRepayParameter struct {
	value *ChangeAutoRepayFuturesStatusAutoRepayParameter
	isSet bool
}

func (v NullableChangeAutoRepayFuturesStatusAutoRepayParameter) Get() *ChangeAutoRepayFuturesStatusAutoRepayParameter {
	return v.value
}

func (v *NullableChangeAutoRepayFuturesStatusAutoRepayParameter) Set(val *ChangeAutoRepayFuturesStatusAutoRepayParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeAutoRepayFuturesStatusAutoRepayParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeAutoRepayFuturesStatusAutoRepayParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeAutoRepayFuturesStatusAutoRepayParameter(val *ChangeAutoRepayFuturesStatusAutoRepayParameter) *NullableChangeAutoRepayFuturesStatusAutoRepayParameter {
	return &NullableChangeAutoRepayFuturesStatusAutoRepayParameter{value: val, isSet: true}
}

func (v NullableChangeAutoRepayFuturesStatusAutoRepayParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeAutoRepayFuturesStatusAutoRepayParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
