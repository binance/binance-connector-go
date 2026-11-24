/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter the model 'PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter'
type PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter string

// List of placeEquityOrder_tokenizedAssetSettleWalletType_parameter
const (
	PlaceEquityOrderTokenizedAssetSettleWalletTypeParameterMain PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter = "MAIN"
	PlaceEquityOrderTokenizedAssetSettleWalletTypeParameterCard PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter = "CARD"
)

// All allowed values of PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter enum
var AllowedPlaceEquityOrderTokenizedAssetSettleWalletTypeParameterEnumValues = []PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter{
	"MAIN",
	"CARD",
}

func (v *PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter(value)
	for _, existing := range AllowedPlaceEquityOrderTokenizedAssetSettleWalletTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter", value)
}

// NewPlaceEquityOrderTokenizedAssetSettleWalletTypeParameterFromValue returns a pointer to a valid PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceEquityOrderTokenizedAssetSettleWalletTypeParameterFromValue(v string) (*PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter, error) {
	ev := PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter: valid values are %v", v, AllowedPlaceEquityOrderTokenizedAssetSettleWalletTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter) IsValid() bool {
	for _, existing := range AllowedPlaceEquityOrderTokenizedAssetSettleWalletTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeEquityOrder_tokenizedAssetSettleWalletType_parameter value
func (v PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter) Ptr() *PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter {
	return &v
}

type NullablePlaceEquityOrderTokenizedAssetSettleWalletTypeParameter struct {
	value *PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter
	isSet bool
}

func (v NullablePlaceEquityOrderTokenizedAssetSettleWalletTypeParameter) Get() *PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter {
	return v.value
}

func (v *NullablePlaceEquityOrderTokenizedAssetSettleWalletTypeParameter) Set(val *PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceEquityOrderTokenizedAssetSettleWalletTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceEquityOrderTokenizedAssetSettleWalletTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceEquityOrderTokenizedAssetSettleWalletTypeParameter(val *PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter) *NullablePlaceEquityOrderTokenizedAssetSettleWalletTypeParameter {
	return &NullablePlaceEquityOrderTokenizedAssetSettleWalletTypeParameter{value: val, isSet: true}
}

func (v NullablePlaceEquityOrderTokenizedAssetSettleWalletTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceEquityOrderTokenizedAssetSettleWalletTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
