/*
Convert REST API

Request quotes and execute cryptocurrency conversions via the Convert REST API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceLimitOrderExpiredTypeParameter the model 'PlaceLimitOrderExpiredTypeParameter'
type PlaceLimitOrderExpiredTypeParameter string

// List of placeLimitOrder_expiredType_parameter
const (
	PlaceLimitOrderExpiredTypeParameterExpiredType1D  PlaceLimitOrderExpiredTypeParameter = "1_D"
	PlaceLimitOrderExpiredTypeParameterExpiredType3D  PlaceLimitOrderExpiredTypeParameter = "3_D"
	PlaceLimitOrderExpiredTypeParameterExpiredType7D  PlaceLimitOrderExpiredTypeParameter = "7_D"
	PlaceLimitOrderExpiredTypeParameterExpiredType30D PlaceLimitOrderExpiredTypeParameter = "30_D"
)

// All allowed values of PlaceLimitOrderExpiredTypeParameter enum
var AllowedPlaceLimitOrderExpiredTypeParameterEnumValues = []PlaceLimitOrderExpiredTypeParameter{
	"1_D",
	"3_D",
	"7_D",
	"30_D",
}

func (v *PlaceLimitOrderExpiredTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceLimitOrderExpiredTypeParameter(value)
	for _, existing := range AllowedPlaceLimitOrderExpiredTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceLimitOrderExpiredTypeParameter", value)
}

// NewPlaceLimitOrderExpiredTypeParameterFromValue returns a pointer to a valid PlaceLimitOrderExpiredTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceLimitOrderExpiredTypeParameterFromValue(v string) (*PlaceLimitOrderExpiredTypeParameter, error) {
	ev := PlaceLimitOrderExpiredTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceLimitOrderExpiredTypeParameter: valid values are %v", v, AllowedPlaceLimitOrderExpiredTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceLimitOrderExpiredTypeParameter) IsValid() bool {
	for _, existing := range AllowedPlaceLimitOrderExpiredTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeLimitOrder_expiredType_parameter value
func (v PlaceLimitOrderExpiredTypeParameter) Ptr() *PlaceLimitOrderExpiredTypeParameter {
	return &v
}

type NullablePlaceLimitOrderExpiredTypeParameter struct {
	value *PlaceLimitOrderExpiredTypeParameter
	isSet bool
}

func (v NullablePlaceLimitOrderExpiredTypeParameter) Get() *PlaceLimitOrderExpiredTypeParameter {
	return v.value
}

func (v *NullablePlaceLimitOrderExpiredTypeParameter) Set(val *PlaceLimitOrderExpiredTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceLimitOrderExpiredTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceLimitOrderExpiredTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceLimitOrderExpiredTypeParameter(val *PlaceLimitOrderExpiredTypeParameter) *NullablePlaceLimitOrderExpiredTypeParameter {
	return &NullablePlaceLimitOrderExpiredTypeParameter{value: val, isSet: true}
}

func (v NullablePlaceLimitOrderExpiredTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceLimitOrderExpiredTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
