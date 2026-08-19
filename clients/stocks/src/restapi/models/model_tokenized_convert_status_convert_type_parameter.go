/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// TokenizedConvertStatusConvertTypeParameter the model 'TokenizedConvertStatusConvertTypeParameter'
type TokenizedConvertStatusConvertTypeParameter string

// List of tokenizedConvertStatus_convertType_parameter
const (
	TokenizedConvertStatusConvertTypeParameterMint   TokenizedConvertStatusConvertTypeParameter = "MINT"
	TokenizedConvertStatusConvertTypeParameterRedeem TokenizedConvertStatusConvertTypeParameter = "REDEEM"
)

// All allowed values of TokenizedConvertStatusConvertTypeParameter enum
var AllowedTokenizedConvertStatusConvertTypeParameterEnumValues = []TokenizedConvertStatusConvertTypeParameter{
	"MINT",
	"REDEEM",
}

func (v *TokenizedConvertStatusConvertTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TokenizedConvertStatusConvertTypeParameter(value)
	for _, existing := range AllowedTokenizedConvertStatusConvertTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TokenizedConvertStatusConvertTypeParameter", value)
}

// NewTokenizedConvertStatusConvertTypeParameterFromValue returns a pointer to a valid TokenizedConvertStatusConvertTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTokenizedConvertStatusConvertTypeParameterFromValue(v string) (*TokenizedConvertStatusConvertTypeParameter, error) {
	ev := TokenizedConvertStatusConvertTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TokenizedConvertStatusConvertTypeParameter: valid values are %v", v, AllowedTokenizedConvertStatusConvertTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TokenizedConvertStatusConvertTypeParameter) IsValid() bool {
	for _, existing := range AllowedTokenizedConvertStatusConvertTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tokenizedConvertStatus_convertType_parameter value
func (v TokenizedConvertStatusConvertTypeParameter) Ptr() *TokenizedConvertStatusConvertTypeParameter {
	return &v
}

type NullableTokenizedConvertStatusConvertTypeParameter struct {
	value *TokenizedConvertStatusConvertTypeParameter
	isSet bool
}

func (v NullableTokenizedConvertStatusConvertTypeParameter) Get() *TokenizedConvertStatusConvertTypeParameter {
	return v.value
}

func (v *NullableTokenizedConvertStatusConvertTypeParameter) Set(val *TokenizedConvertStatusConvertTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizedConvertStatusConvertTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizedConvertStatusConvertTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizedConvertStatusConvertTypeParameter(val *TokenizedConvertStatusConvertTypeParameter) *NullableTokenizedConvertStatusConvertTypeParameter {
	return &NullableTokenizedConvertStatusConvertTypeParameter{value: val, isSet: true}
}

func (v NullableTokenizedConvertStatusConvertTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizedConvertStatusConvertTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
