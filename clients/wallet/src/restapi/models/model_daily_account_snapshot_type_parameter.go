/*
Wallet REST API

Query balances, manage assets, and perform wallet operations via the Binance Wallet API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// DailyAccountSnapshotTypeParameter the model 'DailyAccountSnapshotTypeParameter'
type DailyAccountSnapshotTypeParameter string

// List of dailyAccountSnapshot_type_parameter
const (
	DailyAccountSnapshotTypeParameterSpot    DailyAccountSnapshotTypeParameter = "SPOT"
	DailyAccountSnapshotTypeParameterMargin  DailyAccountSnapshotTypeParameter = "MARGIN"
	DailyAccountSnapshotTypeParameterFutures DailyAccountSnapshotTypeParameter = "FUTURES"
)

// All allowed values of DailyAccountSnapshotTypeParameter enum
var AllowedDailyAccountSnapshotTypeParameterEnumValues = []DailyAccountSnapshotTypeParameter{
	"SPOT",
	"MARGIN",
	"FUTURES",
}

func (v *DailyAccountSnapshotTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DailyAccountSnapshotTypeParameter(value)
	for _, existing := range AllowedDailyAccountSnapshotTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DailyAccountSnapshotTypeParameter", value)
}

// NewDailyAccountSnapshotTypeParameterFromValue returns a pointer to a valid DailyAccountSnapshotTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDailyAccountSnapshotTypeParameterFromValue(v string) (*DailyAccountSnapshotTypeParameter, error) {
	ev := DailyAccountSnapshotTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DailyAccountSnapshotTypeParameter: valid values are %v", v, AllowedDailyAccountSnapshotTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DailyAccountSnapshotTypeParameter) IsValid() bool {
	for _, existing := range AllowedDailyAccountSnapshotTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dailyAccountSnapshot_type_parameter value
func (v DailyAccountSnapshotTypeParameter) Ptr() *DailyAccountSnapshotTypeParameter {
	return &v
}

type NullableDailyAccountSnapshotTypeParameter struct {
	value *DailyAccountSnapshotTypeParameter
	isSet bool
}

func (v NullableDailyAccountSnapshotTypeParameter) Get() *DailyAccountSnapshotTypeParameter {
	return v.value
}

func (v *NullableDailyAccountSnapshotTypeParameter) Set(val *DailyAccountSnapshotTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyAccountSnapshotTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyAccountSnapshotTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyAccountSnapshotTypeParameter(val *DailyAccountSnapshotTypeParameter) *NullableDailyAccountSnapshotTypeParameter {
	return &NullableDailyAccountSnapshotTypeParameter{value: val, isSet: true}
}

func (v NullableDailyAccountSnapshotTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyAccountSnapshotTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
