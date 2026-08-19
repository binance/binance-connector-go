/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceEquityOrderOrderTypeParameter the model 'PlaceEquityOrderOrderTypeParameter'
type PlaceEquityOrderOrderTypeParameter string

// List of placeEquityOrder_orderType_parameter
const (
	PlaceEquityOrderOrderTypeParameterMarket PlaceEquityOrderOrderTypeParameter = "MARKET"
	PlaceEquityOrderOrderTypeParameterLimit  PlaceEquityOrderOrderTypeParameter = "LIMIT"
)

// All allowed values of PlaceEquityOrderOrderTypeParameter enum
var AllowedPlaceEquityOrderOrderTypeParameterEnumValues = []PlaceEquityOrderOrderTypeParameter{
	"MARKET",
	"LIMIT",
}

func (v *PlaceEquityOrderOrderTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceEquityOrderOrderTypeParameter(value)
	for _, existing := range AllowedPlaceEquityOrderOrderTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceEquityOrderOrderTypeParameter", value)
}

// NewPlaceEquityOrderOrderTypeParameterFromValue returns a pointer to a valid PlaceEquityOrderOrderTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceEquityOrderOrderTypeParameterFromValue(v string) (*PlaceEquityOrderOrderTypeParameter, error) {
	ev := PlaceEquityOrderOrderTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceEquityOrderOrderTypeParameter: valid values are %v", v, AllowedPlaceEquityOrderOrderTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceEquityOrderOrderTypeParameter) IsValid() bool {
	for _, existing := range AllowedPlaceEquityOrderOrderTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeEquityOrder_orderType_parameter value
func (v PlaceEquityOrderOrderTypeParameter) Ptr() *PlaceEquityOrderOrderTypeParameter {
	return &v
}

type NullablePlaceEquityOrderOrderTypeParameter struct {
	value *PlaceEquityOrderOrderTypeParameter
	isSet bool
}

func (v NullablePlaceEquityOrderOrderTypeParameter) Get() *PlaceEquityOrderOrderTypeParameter {
	return v.value
}

func (v *NullablePlaceEquityOrderOrderTypeParameter) Set(val *PlaceEquityOrderOrderTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceEquityOrderOrderTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceEquityOrderOrderTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceEquityOrderOrderTypeParameter(val *PlaceEquityOrderOrderTypeParameter) *NullablePlaceEquityOrderOrderTypeParameter {
	return &NullablePlaceEquityOrderOrderTypeParameter{value: val, isSet: true}
}

func (v NullablePlaceEquityOrderOrderTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceEquityOrderOrderTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
