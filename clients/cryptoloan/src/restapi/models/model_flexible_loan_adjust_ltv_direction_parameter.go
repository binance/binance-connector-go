/*
Crypto Loan REST API

Access Binance Crypto Loans to query assets, subscribe to loans, and manage loan positions.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// FlexibleLoanAdjustLtvDirectionParameter the model 'FlexibleLoanAdjustLtvDirectionParameter'
type FlexibleLoanAdjustLtvDirectionParameter string

// List of flexibleLoanAdjustLtv_direction_parameter
const (
	FlexibleLoanAdjustLtvDirectionParameterAdditional FlexibleLoanAdjustLtvDirectionParameter = "ADDITIONAL"
	FlexibleLoanAdjustLtvDirectionParameterReduced    FlexibleLoanAdjustLtvDirectionParameter = "REDUCED"
)

// All allowed values of FlexibleLoanAdjustLtvDirectionParameter enum
var AllowedFlexibleLoanAdjustLtvDirectionParameterEnumValues = []FlexibleLoanAdjustLtvDirectionParameter{
	"ADDITIONAL",
	"REDUCED",
}

func (v *FlexibleLoanAdjustLtvDirectionParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FlexibleLoanAdjustLtvDirectionParameter(value)
	for _, existing := range AllowedFlexibleLoanAdjustLtvDirectionParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FlexibleLoanAdjustLtvDirectionParameter", value)
}

// NewFlexibleLoanAdjustLtvDirectionParameterFromValue returns a pointer to a valid FlexibleLoanAdjustLtvDirectionParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFlexibleLoanAdjustLtvDirectionParameterFromValue(v string) (*FlexibleLoanAdjustLtvDirectionParameter, error) {
	ev := FlexibleLoanAdjustLtvDirectionParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FlexibleLoanAdjustLtvDirectionParameter: valid values are %v", v, AllowedFlexibleLoanAdjustLtvDirectionParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FlexibleLoanAdjustLtvDirectionParameter) IsValid() bool {
	for _, existing := range AllowedFlexibleLoanAdjustLtvDirectionParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to flexibleLoanAdjustLtv_direction_parameter value
func (v FlexibleLoanAdjustLtvDirectionParameter) Ptr() *FlexibleLoanAdjustLtvDirectionParameter {
	return &v
}

type NullableFlexibleLoanAdjustLtvDirectionParameter struct {
	value *FlexibleLoanAdjustLtvDirectionParameter
	isSet bool
}

func (v NullableFlexibleLoanAdjustLtvDirectionParameter) Get() *FlexibleLoanAdjustLtvDirectionParameter {
	return v.value
}

func (v *NullableFlexibleLoanAdjustLtvDirectionParameter) Set(val *FlexibleLoanAdjustLtvDirectionParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableFlexibleLoanAdjustLtvDirectionParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableFlexibleLoanAdjustLtvDirectionParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlexibleLoanAdjustLtvDirectionParameter(val *FlexibleLoanAdjustLtvDirectionParameter) *NullableFlexibleLoanAdjustLtvDirectionParameter {
	return &NullableFlexibleLoanAdjustLtvDirectionParameter{value: val, isSet: true}
}

func (v NullableFlexibleLoanAdjustLtvDirectionParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlexibleLoanAdjustLtvDirectionParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
