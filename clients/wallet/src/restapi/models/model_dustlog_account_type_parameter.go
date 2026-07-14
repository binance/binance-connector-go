/*
Wallet REST API

Query balances, manage assets, and perform wallet operations via the Binance Wallet API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// DustlogAccountTypeParameter the model 'DustlogAccountTypeParameter'
type DustlogAccountTypeParameter string

// List of dustlog_accountType_parameter
const (
	DustlogAccountTypeParameterSpot   DustlogAccountTypeParameter = "SPOT"
	DustlogAccountTypeParameterMargin DustlogAccountTypeParameter = "MARGIN"
)

// All allowed values of DustlogAccountTypeParameter enum
var AllowedDustlogAccountTypeParameterEnumValues = []DustlogAccountTypeParameter{
	"SPOT",
	"MARGIN",
}

func (v *DustlogAccountTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DustlogAccountTypeParameter(value)
	for _, existing := range AllowedDustlogAccountTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DustlogAccountTypeParameter", value)
}

// NewDustlogAccountTypeParameterFromValue returns a pointer to a valid DustlogAccountTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDustlogAccountTypeParameterFromValue(v string) (*DustlogAccountTypeParameter, error) {
	ev := DustlogAccountTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DustlogAccountTypeParameter: valid values are %v", v, AllowedDustlogAccountTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DustlogAccountTypeParameter) IsValid() bool {
	for _, existing := range AllowedDustlogAccountTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dustlog_accountType_parameter value
func (v DustlogAccountTypeParameter) Ptr() *DustlogAccountTypeParameter {
	return &v
}

type NullableDustlogAccountTypeParameter struct {
	value *DustlogAccountTypeParameter
	isSet bool
}

func (v NullableDustlogAccountTypeParameter) Get() *DustlogAccountTypeParameter {
	return v.value
}

func (v *NullableDustlogAccountTypeParameter) Set(val *DustlogAccountTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDustlogAccountTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDustlogAccountTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDustlogAccountTypeParameter(val *DustlogAccountTypeParameter) *NullableDustlogAccountTypeParameter {
	return &NullableDustlogAccountTypeParameter{value: val, isSet: true}
}

func (v NullableDustlogAccountTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDustlogAccountTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
