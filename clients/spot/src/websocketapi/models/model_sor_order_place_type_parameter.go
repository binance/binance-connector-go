/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// SorOrderPlaceTypeParameter the model 'SorOrderPlaceTypeParameter'
type SorOrderPlaceTypeParameter string

// List of sorOrderPlace_type_parameter
const (
	SorOrderPlaceTypeParameterMarket SorOrderPlaceTypeParameter = "MARKET"
	SorOrderPlaceTypeParameterLimit  SorOrderPlaceTypeParameter = "LIMIT"
)

// All allowed values of SorOrderPlaceTypeParameter enum
var AllowedSorOrderPlaceTypeParameterEnumValues = []SorOrderPlaceTypeParameter{
	"MARKET",
	"LIMIT",
}

func (v *SorOrderPlaceTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SorOrderPlaceTypeParameter(value)
	for _, existing := range AllowedSorOrderPlaceTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SorOrderPlaceTypeParameter", value)
}

// NewSorOrderPlaceTypeParameterFromValue returns a pointer to a valid SorOrderPlaceTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSorOrderPlaceTypeParameterFromValue(v string) (*SorOrderPlaceTypeParameter, error) {
	ev := SorOrderPlaceTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SorOrderPlaceTypeParameter: valid values are %v", v, AllowedSorOrderPlaceTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SorOrderPlaceTypeParameter) IsValid() bool {
	for _, existing := range AllowedSorOrderPlaceTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to sorOrderPlace_type_parameter value
func (v SorOrderPlaceTypeParameter) Ptr() *SorOrderPlaceTypeParameter {
	return &v
}

type NullableSorOrderPlaceTypeParameter struct {
	value *SorOrderPlaceTypeParameter
	isSet bool
}

func (v NullableSorOrderPlaceTypeParameter) Get() *SorOrderPlaceTypeParameter {
	return v.value
}

func (v *NullableSorOrderPlaceTypeParameter) Set(val *SorOrderPlaceTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableSorOrderPlaceTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableSorOrderPlaceTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSorOrderPlaceTypeParameter(val *SorOrderPlaceTypeParameter) *NullableSorOrderPlaceTypeParameter {
	return &NullableSorOrderPlaceTypeParameter{value: val, isSet: true}
}

func (v NullableSorOrderPlaceTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSorOrderPlaceTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
