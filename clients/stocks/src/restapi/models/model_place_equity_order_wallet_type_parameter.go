/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceEquityOrderWalletTypeParameter the model 'PlaceEquityOrderWalletTypeParameter'
type PlaceEquityOrderWalletTypeParameter string

// List of placeEquityOrder_walletType_parameter
const (
	PlaceEquityOrderWalletTypeParameterCard PlaceEquityOrderWalletTypeParameter = "CARD"
	PlaceEquityOrderWalletTypeParameterMain PlaceEquityOrderWalletTypeParameter = "MAIN"
)

// All allowed values of PlaceEquityOrderWalletTypeParameter enum
var AllowedPlaceEquityOrderWalletTypeParameterEnumValues = []PlaceEquityOrderWalletTypeParameter{
	"CARD",
	"MAIN",
}

func (v *PlaceEquityOrderWalletTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceEquityOrderWalletTypeParameter(value)
	for _, existing := range AllowedPlaceEquityOrderWalletTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceEquityOrderWalletTypeParameter", value)
}

// NewPlaceEquityOrderWalletTypeParameterFromValue returns a pointer to a valid PlaceEquityOrderWalletTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceEquityOrderWalletTypeParameterFromValue(v string) (*PlaceEquityOrderWalletTypeParameter, error) {
	ev := PlaceEquityOrderWalletTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceEquityOrderWalletTypeParameter: valid values are %v", v, AllowedPlaceEquityOrderWalletTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceEquityOrderWalletTypeParameter) IsValid() bool {
	for _, existing := range AllowedPlaceEquityOrderWalletTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeEquityOrder_walletType_parameter value
func (v PlaceEquityOrderWalletTypeParameter) Ptr() *PlaceEquityOrderWalletTypeParameter {
	return &v
}

type NullablePlaceEquityOrderWalletTypeParameter struct {
	value *PlaceEquityOrderWalletTypeParameter
	isSet bool
}

func (v NullablePlaceEquityOrderWalletTypeParameter) Get() *PlaceEquityOrderWalletTypeParameter {
	return v.value
}

func (v *NullablePlaceEquityOrderWalletTypeParameter) Set(val *PlaceEquityOrderWalletTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceEquityOrderWalletTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceEquityOrderWalletTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceEquityOrderWalletTypeParameter(val *PlaceEquityOrderWalletTypeParameter) *NullablePlaceEquityOrderWalletTypeParameter {
	return &NullablePlaceEquityOrderWalletTypeParameter{value: val, isSet: true}
}

func (v NullablePlaceEquityOrderWalletTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceEquityOrderWalletTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
