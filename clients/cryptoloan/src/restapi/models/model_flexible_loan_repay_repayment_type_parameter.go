/*
Crypto Loan REST API

Access Binance Crypto Loans to query assets, subscribe to loans, and manage loan positions.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// FlexibleLoanRepayRepaymentTypeParameter the model 'FlexibleLoanRepayRepaymentTypeParameter'
type FlexibleLoanRepayRepaymentTypeParameter int64

// List of flexibleLoanRepay_repaymentType_parameter
const (
	FlexibleLoanRepayRepaymentTypeParameterRepaymentType1 FlexibleLoanRepayRepaymentTypeParameter = 1
	FlexibleLoanRepayRepaymentTypeParameterRepaymentType2 FlexibleLoanRepayRepaymentTypeParameter = 2
)

// All allowed values of FlexibleLoanRepayRepaymentTypeParameter enum
var AllowedFlexibleLoanRepayRepaymentTypeParameterEnumValues = []FlexibleLoanRepayRepaymentTypeParameter{
	1,
	2,
}

func (v *FlexibleLoanRepayRepaymentTypeParameter) UnmarshalJSON(src []byte) error {
	var value int64
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FlexibleLoanRepayRepaymentTypeParameter(value)
	for _, existing := range AllowedFlexibleLoanRepayRepaymentTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FlexibleLoanRepayRepaymentTypeParameter", value)
}

// NewFlexibleLoanRepayRepaymentTypeParameterFromValue returns a pointer to a valid FlexibleLoanRepayRepaymentTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFlexibleLoanRepayRepaymentTypeParameterFromValue(v int64) (*FlexibleLoanRepayRepaymentTypeParameter, error) {
	ev := FlexibleLoanRepayRepaymentTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FlexibleLoanRepayRepaymentTypeParameter: valid values are %v", v, AllowedFlexibleLoanRepayRepaymentTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FlexibleLoanRepayRepaymentTypeParameter) IsValid() bool {
	for _, existing := range AllowedFlexibleLoanRepayRepaymentTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to flexibleLoanRepay_repaymentType_parameter value
func (v FlexibleLoanRepayRepaymentTypeParameter) Ptr() *FlexibleLoanRepayRepaymentTypeParameter {
	return &v
}

type NullableFlexibleLoanRepayRepaymentTypeParameter struct {
	value *FlexibleLoanRepayRepaymentTypeParameter
	isSet bool
}

func (v NullableFlexibleLoanRepayRepaymentTypeParameter) Get() *FlexibleLoanRepayRepaymentTypeParameter {
	return v.value
}

func (v *NullableFlexibleLoanRepayRepaymentTypeParameter) Set(val *FlexibleLoanRepayRepaymentTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableFlexibleLoanRepayRepaymentTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableFlexibleLoanRepayRepaymentTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlexibleLoanRepayRepaymentTypeParameter(val *FlexibleLoanRepayRepaymentTypeParameter) *NullableFlexibleLoanRepayRepaymentTypeParameter {
	return &NullableFlexibleLoanRepayRepaymentTypeParameter{value: val, isSet: true}
}

func (v NullableFlexibleLoanRepayRepaymentTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlexibleLoanRepayRepaymentTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
