/*
Convert REST API

Request quotes and execute cryptocurrency conversions via the Convert REST API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceLimitOrderWalletTypeParameter the model 'PlaceLimitOrderWalletTypeParameter'
type PlaceLimitOrderWalletTypeParameter string

// List of placeLimitOrder_walletType_parameter
const (
	PlaceLimitOrderWalletTypeParameterSpot            PlaceLimitOrderWalletTypeParameter = "SPOT"
	PlaceLimitOrderWalletTypeParameterFunding         PlaceLimitOrderWalletTypeParameter = "FUNDING"
	PlaceLimitOrderWalletTypeParameterEarn            PlaceLimitOrderWalletTypeParameter = "EARN"
	PlaceLimitOrderWalletTypeParameterSpotFunding     PlaceLimitOrderWalletTypeParameter = "SPOT_FUNDING"
	PlaceLimitOrderWalletTypeParameterFundingEarn     PlaceLimitOrderWalletTypeParameter = "FUNDING_EARN"
	PlaceLimitOrderWalletTypeParameterSpotFundingEarn PlaceLimitOrderWalletTypeParameter = "SPOT_FUNDING_EARN"
	PlaceLimitOrderWalletTypeParameterSpotEarn        PlaceLimitOrderWalletTypeParameter = "SPOT_EARN"
)

// All allowed values of PlaceLimitOrderWalletTypeParameter enum
var AllowedPlaceLimitOrderWalletTypeParameterEnumValues = []PlaceLimitOrderWalletTypeParameter{
	"SPOT",
	"FUNDING",
	"EARN",
	"SPOT_FUNDING",
	"FUNDING_EARN",
	"SPOT_FUNDING_EARN",
	"SPOT_EARN",
}

func (v *PlaceLimitOrderWalletTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceLimitOrderWalletTypeParameter(value)
	for _, existing := range AllowedPlaceLimitOrderWalletTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceLimitOrderWalletTypeParameter", value)
}

// NewPlaceLimitOrderWalletTypeParameterFromValue returns a pointer to a valid PlaceLimitOrderWalletTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceLimitOrderWalletTypeParameterFromValue(v string) (*PlaceLimitOrderWalletTypeParameter, error) {
	ev := PlaceLimitOrderWalletTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceLimitOrderWalletTypeParameter: valid values are %v", v, AllowedPlaceLimitOrderWalletTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceLimitOrderWalletTypeParameter) IsValid() bool {
	for _, existing := range AllowedPlaceLimitOrderWalletTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeLimitOrder_walletType_parameter value
func (v PlaceLimitOrderWalletTypeParameter) Ptr() *PlaceLimitOrderWalletTypeParameter {
	return &v
}

type NullablePlaceLimitOrderWalletTypeParameter struct {
	value *PlaceLimitOrderWalletTypeParameter
	isSet bool
}

func (v NullablePlaceLimitOrderWalletTypeParameter) Get() *PlaceLimitOrderWalletTypeParameter {
	return v.value
}

func (v *NullablePlaceLimitOrderWalletTypeParameter) Set(val *PlaceLimitOrderWalletTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceLimitOrderWalletTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceLimitOrderWalletTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceLimitOrderWalletTypeParameter(val *PlaceLimitOrderWalletTypeParameter) *NullablePlaceLimitOrderWalletTypeParameter {
	return &NullablePlaceLimitOrderWalletTypeParameter{value: val, isSet: true}
}

func (v NullablePlaceLimitOrderWalletTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceLimitOrderWalletTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
