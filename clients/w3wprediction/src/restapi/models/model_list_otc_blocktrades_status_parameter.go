/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ListOtcBlocktradesStatusParameter the model 'ListOtcBlocktradesStatusParameter'
type ListOtcBlocktradesStatusParameter string

// List of listOtcBlocktrades_status_parameter
const (
	ListOtcBlocktradesStatusParameterOpen      ListOtcBlocktradesStatusParameter = "OPEN"
	ListOtcBlocktradesStatusParameterFulfilled ListOtcBlocktradesStatusParameter = "FULFILLED"
	ListOtcBlocktradesStatusParameterMatched   ListOtcBlocktradesStatusParameter = "MATCHED"
	ListOtcBlocktradesStatusParameterCancelled ListOtcBlocktradesStatusParameter = "CANCELLED"
	ListOtcBlocktradesStatusParameterExpired   ListOtcBlocktradesStatusParameter = "EXPIRED"
	ListOtcBlocktradesStatusParameterFailed    ListOtcBlocktradesStatusParameter = "FAILED"
)

// All allowed values of ListOtcBlocktradesStatusParameter enum
var AllowedListOtcBlocktradesStatusParameterEnumValues = []ListOtcBlocktradesStatusParameter{
	"OPEN",
	"FULFILLED",
	"MATCHED",
	"CANCELLED",
	"EXPIRED",
	"FAILED",
}

func (v *ListOtcBlocktradesStatusParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListOtcBlocktradesStatusParameter(value)
	for _, existing := range AllowedListOtcBlocktradesStatusParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListOtcBlocktradesStatusParameter", value)
}

// NewListOtcBlocktradesStatusParameterFromValue returns a pointer to a valid ListOtcBlocktradesStatusParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListOtcBlocktradesStatusParameterFromValue(v string) (*ListOtcBlocktradesStatusParameter, error) {
	ev := ListOtcBlocktradesStatusParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListOtcBlocktradesStatusParameter: valid values are %v", v, AllowedListOtcBlocktradesStatusParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListOtcBlocktradesStatusParameter) IsValid() bool {
	for _, existing := range AllowedListOtcBlocktradesStatusParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to listOtcBlocktrades_status_parameter value
func (v ListOtcBlocktradesStatusParameter) Ptr() *ListOtcBlocktradesStatusParameter {
	return &v
}

type NullableListOtcBlocktradesStatusParameter struct {
	value *ListOtcBlocktradesStatusParameter
	isSet bool
}

func (v NullableListOtcBlocktradesStatusParameter) Get() *ListOtcBlocktradesStatusParameter {
	return v.value
}

func (v *NullableListOtcBlocktradesStatusParameter) Set(val *ListOtcBlocktradesStatusParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableListOtcBlocktradesStatusParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableListOtcBlocktradesStatusParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOtcBlocktradesStatusParameter(val *ListOtcBlocktradesStatusParameter) *NullableListOtcBlocktradesStatusParameter {
	return &NullableListOtcBlocktradesStatusParameter{value: val, isSet: true}
}

func (v NullableListOtcBlocktradesStatusParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOtcBlocktradesStatusParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
