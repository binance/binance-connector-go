/*
Portfolio Margin Pro REST API

Access advanced account management and high-frequency trading with Binance Portfolio Margin Pro.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// FundCollectionByAssetFutureTypeParameter the model 'FundCollectionByAssetFutureTypeParameter'
type FundCollectionByAssetFutureTypeParameter string

// List of fundCollectionByAsset_futureType_parameter
const (
	FundCollectionByAssetFutureTypeParameterUm     FundCollectionByAssetFutureTypeParameter = "UM"
	FundCollectionByAssetFutureTypeParameterCm     FundCollectionByAssetFutureTypeParameter = "CM"
	FundCollectionByAssetFutureTypeParameterOption FundCollectionByAssetFutureTypeParameter = "OPTION"
)

// All allowed values of FundCollectionByAssetFutureTypeParameter enum
var AllowedFundCollectionByAssetFutureTypeParameterEnumValues = []FundCollectionByAssetFutureTypeParameter{
	"UM",
	"CM",
	"OPTION",
}

func (v *FundCollectionByAssetFutureTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FundCollectionByAssetFutureTypeParameter(value)
	for _, existing := range AllowedFundCollectionByAssetFutureTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FundCollectionByAssetFutureTypeParameter", value)
}

// NewFundCollectionByAssetFutureTypeParameterFromValue returns a pointer to a valid FundCollectionByAssetFutureTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFundCollectionByAssetFutureTypeParameterFromValue(v string) (*FundCollectionByAssetFutureTypeParameter, error) {
	ev := FundCollectionByAssetFutureTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FundCollectionByAssetFutureTypeParameter: valid values are %v", v, AllowedFundCollectionByAssetFutureTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FundCollectionByAssetFutureTypeParameter) IsValid() bool {
	for _, existing := range AllowedFundCollectionByAssetFutureTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to fundCollectionByAsset_futureType_parameter value
func (v FundCollectionByAssetFutureTypeParameter) Ptr() *FundCollectionByAssetFutureTypeParameter {
	return &v
}

type NullableFundCollectionByAssetFutureTypeParameter struct {
	value *FundCollectionByAssetFutureTypeParameter
	isSet bool
}

func (v NullableFundCollectionByAssetFutureTypeParameter) Get() *FundCollectionByAssetFutureTypeParameter {
	return v.value
}

func (v *NullableFundCollectionByAssetFutureTypeParameter) Set(val *FundCollectionByAssetFutureTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableFundCollectionByAssetFutureTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableFundCollectionByAssetFutureTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundCollectionByAssetFutureTypeParameter(val *FundCollectionByAssetFutureTypeParameter) *NullableFundCollectionByAssetFutureTypeParameter {
	return &NullableFundCollectionByAssetFutureTypeParameter{value: val, isSet: true}
}

func (v NullableFundCollectionByAssetFutureTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundCollectionByAssetFutureTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
