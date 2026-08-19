/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceEquityOrderSideParameter the model 'PlaceEquityOrderSideParameter'
type PlaceEquityOrderSideParameter string

// List of placeEquityOrder_side_parameter
const (
	PlaceEquityOrderSideParameterBuy  PlaceEquityOrderSideParameter = "BUY"
	PlaceEquityOrderSideParameterSell PlaceEquityOrderSideParameter = "SELL"
)

// All allowed values of PlaceEquityOrderSideParameter enum
var AllowedPlaceEquityOrderSideParameterEnumValues = []PlaceEquityOrderSideParameter{
	"BUY",
	"SELL",
}

func (v *PlaceEquityOrderSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceEquityOrderSideParameter(value)
	for _, existing := range AllowedPlaceEquityOrderSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceEquityOrderSideParameter", value)
}

// NewPlaceEquityOrderSideParameterFromValue returns a pointer to a valid PlaceEquityOrderSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceEquityOrderSideParameterFromValue(v string) (*PlaceEquityOrderSideParameter, error) {
	ev := PlaceEquityOrderSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceEquityOrderSideParameter: valid values are %v", v, AllowedPlaceEquityOrderSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceEquityOrderSideParameter) IsValid() bool {
	for _, existing := range AllowedPlaceEquityOrderSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeEquityOrder_side_parameter value
func (v PlaceEquityOrderSideParameter) Ptr() *PlaceEquityOrderSideParameter {
	return &v
}

type NullablePlaceEquityOrderSideParameter struct {
	value *PlaceEquityOrderSideParameter
	isSet bool
}

func (v NullablePlaceEquityOrderSideParameter) Get() *PlaceEquityOrderSideParameter {
	return v.value
}

func (v *NullablePlaceEquityOrderSideParameter) Set(val *PlaceEquityOrderSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceEquityOrderSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceEquityOrderSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceEquityOrderSideParameter(val *PlaceEquityOrderSideParameter) *NullablePlaceEquityOrderSideParameter {
	return &NullablePlaceEquityOrderSideParameter{value: val, isSet: true}
}

func (v NullablePlaceEquityOrderSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceEquityOrderSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
