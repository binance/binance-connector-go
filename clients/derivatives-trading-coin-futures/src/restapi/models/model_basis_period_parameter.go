/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// BasisPeriodParameter the model 'BasisPeriodParameter'
type BasisPeriodParameter string

// List of basis_period_parameter
const (
	BasisPeriodParameterPeriod5m  BasisPeriodParameter = "5m"
	BasisPeriodParameterPeriod15m BasisPeriodParameter = "15m"
	BasisPeriodParameterPeriod30m BasisPeriodParameter = "30m"
	BasisPeriodParameterPeriod1h  BasisPeriodParameter = "1h"
	BasisPeriodParameterPeriod2h  BasisPeriodParameter = "2h"
	BasisPeriodParameterPeriod4h  BasisPeriodParameter = "4h"
	BasisPeriodParameterPeriod6h  BasisPeriodParameter = "6h"
	BasisPeriodParameterPeriod12h BasisPeriodParameter = "12h"
	BasisPeriodParameterPeriod1d  BasisPeriodParameter = "1d"
)

// All allowed values of BasisPeriodParameter enum
var AllowedBasisPeriodParameterEnumValues = []BasisPeriodParameter{
	"5m",
	"15m",
	"30m",
	"1h",
	"2h",
	"4h",
	"6h",
	"12h",
	"1d",
}

func (v *BasisPeriodParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BasisPeriodParameter(value)
	for _, existing := range AllowedBasisPeriodParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BasisPeriodParameter", value)
}

// NewBasisPeriodParameterFromValue returns a pointer to a valid BasisPeriodParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBasisPeriodParameterFromValue(v string) (*BasisPeriodParameter, error) {
	ev := BasisPeriodParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BasisPeriodParameter: valid values are %v", v, AllowedBasisPeriodParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BasisPeriodParameter) IsValid() bool {
	for _, existing := range AllowedBasisPeriodParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to basis_period_parameter value
func (v BasisPeriodParameter) Ptr() *BasisPeriodParameter {
	return &v
}

type NullableBasisPeriodParameter struct {
	value *BasisPeriodParameter
	isSet bool
}

func (v NullableBasisPeriodParameter) Get() *BasisPeriodParameter {
	return v.value
}

func (v *NullableBasisPeriodParameter) Set(val *BasisPeriodParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableBasisPeriodParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableBasisPeriodParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasisPeriodParameter(val *BasisPeriodParameter) *NullableBasisPeriodParameter {
	return &NullableBasisPeriodParameter{value: val, isSet: true}
}

func (v NullableBasisPeriodParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasisPeriodParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
