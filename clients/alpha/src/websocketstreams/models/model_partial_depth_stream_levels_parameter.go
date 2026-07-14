/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PartialDepthStreamLevelsParameter the model 'PartialDepthStreamLevelsParameter'
type PartialDepthStreamLevelsParameter string

// List of partialDepthStream_levels_parameter
const (
	PartialDepthStreamLevelsParameterLevels5  PartialDepthStreamLevelsParameter = "5"
	PartialDepthStreamLevelsParameterLevels10 PartialDepthStreamLevelsParameter = "10"
	PartialDepthStreamLevelsParameterLevels20 PartialDepthStreamLevelsParameter = "20"
)

// All allowed values of PartialDepthStreamLevelsParameter enum
var AllowedPartialDepthStreamLevelsParameterEnumValues = []PartialDepthStreamLevelsParameter{
	"5",
	"10",
	"20",
}

func (v *PartialDepthStreamLevelsParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PartialDepthStreamLevelsParameter(value)
	for _, existing := range AllowedPartialDepthStreamLevelsParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PartialDepthStreamLevelsParameter", value)
}

// NewPartialDepthStreamLevelsParameterFromValue returns a pointer to a valid PartialDepthStreamLevelsParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPartialDepthStreamLevelsParameterFromValue(v string) (*PartialDepthStreamLevelsParameter, error) {
	ev := PartialDepthStreamLevelsParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PartialDepthStreamLevelsParameter: valid values are %v", v, AllowedPartialDepthStreamLevelsParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PartialDepthStreamLevelsParameter) IsValid() bool {
	for _, existing := range AllowedPartialDepthStreamLevelsParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to partialDepthStream_levels_parameter value
func (v PartialDepthStreamLevelsParameter) Ptr() *PartialDepthStreamLevelsParameter {
	return &v
}

type NullablePartialDepthStreamLevelsParameter struct {
	value *PartialDepthStreamLevelsParameter
	isSet bool
}

func (v NullablePartialDepthStreamLevelsParameter) Get() *PartialDepthStreamLevelsParameter {
	return v.value
}

func (v *NullablePartialDepthStreamLevelsParameter) Set(val *PartialDepthStreamLevelsParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialDepthStreamLevelsParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialDepthStreamLevelsParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialDepthStreamLevelsParameter(val *PartialDepthStreamLevelsParameter) *NullablePartialDepthStreamLevelsParameter {
	return &NullablePartialDepthStreamLevelsParameter{value: val, isSet: true}
}

func (v NullablePartialDepthStreamLevelsParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialDepthStreamLevelsParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
