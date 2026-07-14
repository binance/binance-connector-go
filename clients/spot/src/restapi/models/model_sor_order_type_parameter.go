/*
Spot REST API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// SorOrderTypeParameter the model 'SorOrderTypeParameter'
type SorOrderTypeParameter string

// List of sorOrder_type_parameter
const (
	SorOrderTypeParameterMarket SorOrderTypeParameter = "MARKET"
	SorOrderTypeParameterLimit  SorOrderTypeParameter = "LIMIT"
)

// All allowed values of SorOrderTypeParameter enum
var AllowedSorOrderTypeParameterEnumValues = []SorOrderTypeParameter{
	"MARKET",
	"LIMIT",
}

func (v *SorOrderTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SorOrderTypeParameter(value)
	for _, existing := range AllowedSorOrderTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SorOrderTypeParameter", value)
}

// NewSorOrderTypeParameterFromValue returns a pointer to a valid SorOrderTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSorOrderTypeParameterFromValue(v string) (*SorOrderTypeParameter, error) {
	ev := SorOrderTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SorOrderTypeParameter: valid values are %v", v, AllowedSorOrderTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SorOrderTypeParameter) IsValid() bool {
	for _, existing := range AllowedSorOrderTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to sorOrder_type_parameter value
func (v SorOrderTypeParameter) Ptr() *SorOrderTypeParameter {
	return &v
}

type NullableSorOrderTypeParameter struct {
	value *SorOrderTypeParameter
	isSet bool
}

func (v NullableSorOrderTypeParameter) Get() *SorOrderTypeParameter {
	return v.value
}

func (v *NullableSorOrderTypeParameter) Set(val *SorOrderTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableSorOrderTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableSorOrderTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSorOrderTypeParameter(val *SorOrderTypeParameter) *NullableSorOrderTypeParameter {
	return &NullableSorOrderTypeParameter{value: val, isSet: true}
}

func (v NullableSorOrderTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSorOrderTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
