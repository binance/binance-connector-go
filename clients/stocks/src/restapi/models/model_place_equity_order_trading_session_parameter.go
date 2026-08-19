/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceEquityOrderTradingSessionParameter the model 'PlaceEquityOrderTradingSessionParameter'
type PlaceEquityOrderTradingSessionParameter string

// List of placeEquityOrder_tradingSession_parameter
const (
	PlaceEquityOrderTradingSessionParameterRth               PlaceEquityOrderTradingSessionParameter = "RTH"
	PlaceEquityOrderTradingSessionParameterExtended          PlaceEquityOrderTradingSessionParameter = "EXTENDED"
	PlaceEquityOrderTradingSessionParameterTradingSession24H PlaceEquityOrderTradingSessionParameter = "24H"
)

// All allowed values of PlaceEquityOrderTradingSessionParameter enum
var AllowedPlaceEquityOrderTradingSessionParameterEnumValues = []PlaceEquityOrderTradingSessionParameter{
	"RTH",
	"EXTENDED",
	"24H",
}

func (v *PlaceEquityOrderTradingSessionParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceEquityOrderTradingSessionParameter(value)
	for _, existing := range AllowedPlaceEquityOrderTradingSessionParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceEquityOrderTradingSessionParameter", value)
}

// NewPlaceEquityOrderTradingSessionParameterFromValue returns a pointer to a valid PlaceEquityOrderTradingSessionParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceEquityOrderTradingSessionParameterFromValue(v string) (*PlaceEquityOrderTradingSessionParameter, error) {
	ev := PlaceEquityOrderTradingSessionParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceEquityOrderTradingSessionParameter: valid values are %v", v, AllowedPlaceEquityOrderTradingSessionParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceEquityOrderTradingSessionParameter) IsValid() bool {
	for _, existing := range AllowedPlaceEquityOrderTradingSessionParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeEquityOrder_tradingSession_parameter value
func (v PlaceEquityOrderTradingSessionParameter) Ptr() *PlaceEquityOrderTradingSessionParameter {
	return &v
}

type NullablePlaceEquityOrderTradingSessionParameter struct {
	value *PlaceEquityOrderTradingSessionParameter
	isSet bool
}

func (v NullablePlaceEquityOrderTradingSessionParameter) Get() *PlaceEquityOrderTradingSessionParameter {
	return v.value
}

func (v *NullablePlaceEquityOrderTradingSessionParameter) Set(val *PlaceEquityOrderTradingSessionParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceEquityOrderTradingSessionParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceEquityOrderTradingSessionParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceEquityOrderTradingSessionParameter(val *PlaceEquityOrderTradingSessionParameter) *NullablePlaceEquityOrderTradingSessionParameter {
	return &NullablePlaceEquityOrderTradingSessionParameter{value: val, isSet: true}
}

func (v NullablePlaceEquityOrderTradingSessionParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceEquityOrderTradingSessionParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
