/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// BasisContractTypeParameter the model 'BasisContractTypeParameter'
type BasisContractTypeParameter string

// List of basis_contractType_parameter
const (
	BasisContractTypeParameterPerpetual           BasisContractTypeParameter = "PERPETUAL"
	BasisContractTypeParameterCurrentMonth        BasisContractTypeParameter = "CURRENT_MONTH"
	BasisContractTypeParameterNextMonth           BasisContractTypeParameter = "NEXT_MONTH"
	BasisContractTypeParameterCurrentQuarter      BasisContractTypeParameter = "CURRENT_QUARTER"
	BasisContractTypeParameterNextQuarter         BasisContractTypeParameter = "NEXT_QUARTER"
	BasisContractTypeParameterPerpetualDelivering BasisContractTypeParameter = "PERPETUAL_DELIVERING"
)

// All allowed values of BasisContractTypeParameter enum
var AllowedBasisContractTypeParameterEnumValues = []BasisContractTypeParameter{
	"PERPETUAL",
	"CURRENT_MONTH",
	"NEXT_MONTH",
	"CURRENT_QUARTER",
	"NEXT_QUARTER",
	"PERPETUAL_DELIVERING",
}

func (v *BasisContractTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BasisContractTypeParameter(value)
	for _, existing := range AllowedBasisContractTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BasisContractTypeParameter", value)
}

// NewBasisContractTypeParameterFromValue returns a pointer to a valid BasisContractTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBasisContractTypeParameterFromValue(v string) (*BasisContractTypeParameter, error) {
	ev := BasisContractTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BasisContractTypeParameter: valid values are %v", v, AllowedBasisContractTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BasisContractTypeParameter) IsValid() bool {
	for _, existing := range AllowedBasisContractTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to basis_contractType_parameter value
func (v BasisContractTypeParameter) Ptr() *BasisContractTypeParameter {
	return &v
}

type NullableBasisContractTypeParameter struct {
	value *BasisContractTypeParameter
	isSet bool
}

func (v NullableBasisContractTypeParameter) Get() *BasisContractTypeParameter {
	return v.value
}

func (v *NullableBasisContractTypeParameter) Set(val *BasisContractTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableBasisContractTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableBasisContractTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasisContractTypeParameter(val *BasisContractTypeParameter) *NullableBasisContractTypeParameter {
	return &NullableBasisContractTypeParameter{value: val, isSet: true}
}

func (v NullableBasisContractTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasisContractTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
