/*
Convert REST API

Request quotes and execute cryptocurrency conversions via the Convert REST API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceLimitOrderSideParameter the model 'PlaceLimitOrderSideParameter'
type PlaceLimitOrderSideParameter string

// List of placeLimitOrder_side_parameter
const (
	PlaceLimitOrderSideParameterBuy  PlaceLimitOrderSideParameter = "BUY"
	PlaceLimitOrderSideParameterSell PlaceLimitOrderSideParameter = "SELL"
)

// All allowed values of PlaceLimitOrderSideParameter enum
var AllowedPlaceLimitOrderSideParameterEnumValues = []PlaceLimitOrderSideParameter{
	"BUY",
	"SELL",
}

func (v *PlaceLimitOrderSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceLimitOrderSideParameter(value)
	for _, existing := range AllowedPlaceLimitOrderSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceLimitOrderSideParameter", value)
}

// NewPlaceLimitOrderSideParameterFromValue returns a pointer to a valid PlaceLimitOrderSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceLimitOrderSideParameterFromValue(v string) (*PlaceLimitOrderSideParameter, error) {
	ev := PlaceLimitOrderSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceLimitOrderSideParameter: valid values are %v", v, AllowedPlaceLimitOrderSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceLimitOrderSideParameter) IsValid() bool {
	for _, existing := range AllowedPlaceLimitOrderSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeLimitOrder_side_parameter value
func (v PlaceLimitOrderSideParameter) Ptr() *PlaceLimitOrderSideParameter {
	return &v
}

type NullablePlaceLimitOrderSideParameter struct {
	value *PlaceLimitOrderSideParameter
	isSet bool
}

func (v NullablePlaceLimitOrderSideParameter) Get() *PlaceLimitOrderSideParameter {
	return v.value
}

func (v *NullablePlaceLimitOrderSideParameter) Set(val *PlaceLimitOrderSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceLimitOrderSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceLimitOrderSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceLimitOrderSideParameter(val *PlaceLimitOrderSideParameter) *NullablePlaceLimitOrderSideParameter {
	return &NullablePlaceLimitOrderSideParameter{value: val, isSet: true}
}

func (v NullablePlaceLimitOrderSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceLimitOrderSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
