/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarginAccountBorrowRepayIsIsolatedParameter the model 'MarginAccountBorrowRepayIsIsolatedParameter'
type MarginAccountBorrowRepayIsIsolatedParameter string

// List of marginAccountBorrowRepay_isIsolated_parameter
const (
	MarginAccountBorrowRepayIsIsolatedParameterTrue  MarginAccountBorrowRepayIsIsolatedParameter = "TRUE"
	MarginAccountBorrowRepayIsIsolatedParameterFalse MarginAccountBorrowRepayIsIsolatedParameter = "FALSE"
)

// All allowed values of MarginAccountBorrowRepayIsIsolatedParameter enum
var AllowedMarginAccountBorrowRepayIsIsolatedParameterEnumValues = []MarginAccountBorrowRepayIsIsolatedParameter{
	"TRUE",
	"FALSE",
}

func (v *MarginAccountBorrowRepayIsIsolatedParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarginAccountBorrowRepayIsIsolatedParameter(value)
	for _, existing := range AllowedMarginAccountBorrowRepayIsIsolatedParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarginAccountBorrowRepayIsIsolatedParameter", value)
}

// NewMarginAccountBorrowRepayIsIsolatedParameterFromValue returns a pointer to a valid MarginAccountBorrowRepayIsIsolatedParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarginAccountBorrowRepayIsIsolatedParameterFromValue(v string) (*MarginAccountBorrowRepayIsIsolatedParameter, error) {
	ev := MarginAccountBorrowRepayIsIsolatedParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarginAccountBorrowRepayIsIsolatedParameter: valid values are %v", v, AllowedMarginAccountBorrowRepayIsIsolatedParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarginAccountBorrowRepayIsIsolatedParameter) IsValid() bool {
	for _, existing := range AllowedMarginAccountBorrowRepayIsIsolatedParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to marginAccountBorrowRepay_isIsolated_parameter value
func (v MarginAccountBorrowRepayIsIsolatedParameter) Ptr() *MarginAccountBorrowRepayIsIsolatedParameter {
	return &v
}

type NullableMarginAccountBorrowRepayIsIsolatedParameter struct {
	value *MarginAccountBorrowRepayIsIsolatedParameter
	isSet bool
}

func (v NullableMarginAccountBorrowRepayIsIsolatedParameter) Get() *MarginAccountBorrowRepayIsIsolatedParameter {
	return v.value
}

func (v *NullableMarginAccountBorrowRepayIsIsolatedParameter) Set(val *MarginAccountBorrowRepayIsIsolatedParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountBorrowRepayIsIsolatedParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountBorrowRepayIsIsolatedParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountBorrowRepayIsIsolatedParameter(val *MarginAccountBorrowRepayIsIsolatedParameter) *NullableMarginAccountBorrowRepayIsIsolatedParameter {
	return &NullableMarginAccountBorrowRepayIsIsolatedParameter{value: val, isSet: true}
}

func (v NullableMarginAccountBorrowRepayIsIsolatedParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountBorrowRepayIsIsolatedParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
