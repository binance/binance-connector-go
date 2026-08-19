/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceEquityOrderTimeInForceParameter the model 'PlaceEquityOrderTimeInForceParameter'
type PlaceEquityOrderTimeInForceParameter string

// List of placeEquityOrder_timeInForce_parameter
const (
	PlaceEquityOrderTimeInForceParameterDay PlaceEquityOrderTimeInForceParameter = "DAY"
	PlaceEquityOrderTimeInForceParameterGtc PlaceEquityOrderTimeInForceParameter = "GTC"
)

// All allowed values of PlaceEquityOrderTimeInForceParameter enum
var AllowedPlaceEquityOrderTimeInForceParameterEnumValues = []PlaceEquityOrderTimeInForceParameter{
	"DAY",
	"GTC",
}

func (v *PlaceEquityOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceEquityOrderTimeInForceParameter(value)
	for _, existing := range AllowedPlaceEquityOrderTimeInForceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceEquityOrderTimeInForceParameter", value)
}

// NewPlaceEquityOrderTimeInForceParameterFromValue returns a pointer to a valid PlaceEquityOrderTimeInForceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceEquityOrderTimeInForceParameterFromValue(v string) (*PlaceEquityOrderTimeInForceParameter, error) {
	ev := PlaceEquityOrderTimeInForceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceEquityOrderTimeInForceParameter: valid values are %v", v, AllowedPlaceEquityOrderTimeInForceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceEquityOrderTimeInForceParameter) IsValid() bool {
	for _, existing := range AllowedPlaceEquityOrderTimeInForceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeEquityOrder_timeInForce_parameter value
func (v PlaceEquityOrderTimeInForceParameter) Ptr() *PlaceEquityOrderTimeInForceParameter {
	return &v
}

type NullablePlaceEquityOrderTimeInForceParameter struct {
	value *PlaceEquityOrderTimeInForceParameter
	isSet bool
}

func (v NullablePlaceEquityOrderTimeInForceParameter) Get() *PlaceEquityOrderTimeInForceParameter {
	return v.value
}

func (v *NullablePlaceEquityOrderTimeInForceParameter) Set(val *PlaceEquityOrderTimeInForceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceEquityOrderTimeInForceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceEquityOrderTimeInForceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceEquityOrderTimeInForceParameter(val *PlaceEquityOrderTimeInForceParameter) *NullablePlaceEquityOrderTimeInForceParameter {
	return &NullablePlaceEquityOrderTimeInForceParameter{value: val, isSet: true}
}

func (v NullablePlaceEquityOrderTimeInForceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceEquityOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
