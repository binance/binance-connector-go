/*
Wallet REST API

Query balances, manage assets, and perform wallet operations via the Binance Wallet API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// DepositHistoryStatusParameter the model 'DepositHistoryStatusParameter'
type DepositHistoryStatusParameter int64

// List of depositHistory_status_parameter
const (
	DepositHistoryStatusParameterStatus0 DepositHistoryStatusParameter = 0
	DepositHistoryStatusParameterStatus1 DepositHistoryStatusParameter = 1
	DepositHistoryStatusParameterStatus2 DepositHistoryStatusParameter = 2
	DepositHistoryStatusParameterStatus6 DepositHistoryStatusParameter = 6
	DepositHistoryStatusParameterStatus7 DepositHistoryStatusParameter = 7
	DepositHistoryStatusParameterStatus8 DepositHistoryStatusParameter = 8
)

// All allowed values of DepositHistoryStatusParameter enum
var AllowedDepositHistoryStatusParameterEnumValues = []DepositHistoryStatusParameter{
	0,
	1,
	2,
	6,
	7,
	8,
}

func (v *DepositHistoryStatusParameter) UnmarshalJSON(src []byte) error {
	var value int64
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DepositHistoryStatusParameter(value)
	for _, existing := range AllowedDepositHistoryStatusParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DepositHistoryStatusParameter", value)
}

// NewDepositHistoryStatusParameterFromValue returns a pointer to a valid DepositHistoryStatusParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDepositHistoryStatusParameterFromValue(v int64) (*DepositHistoryStatusParameter, error) {
	ev := DepositHistoryStatusParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DepositHistoryStatusParameter: valid values are %v", v, AllowedDepositHistoryStatusParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DepositHistoryStatusParameter) IsValid() bool {
	for _, existing := range AllowedDepositHistoryStatusParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to depositHistory_status_parameter value
func (v DepositHistoryStatusParameter) Ptr() *DepositHistoryStatusParameter {
	return &v
}

type NullableDepositHistoryStatusParameter struct {
	value *DepositHistoryStatusParameter
	isSet bool
}

func (v NullableDepositHistoryStatusParameter) Get() *DepositHistoryStatusParameter {
	return v.value
}

func (v *NullableDepositHistoryStatusParameter) Set(val *DepositHistoryStatusParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositHistoryStatusParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositHistoryStatusParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositHistoryStatusParameter(val *DepositHistoryStatusParameter) *NullableDepositHistoryStatusParameter {
	return &NullableDepositHistoryStatusParameter{value: val, isSet: true}
}

func (v NullableDepositHistoryStatusParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositHistoryStatusParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
