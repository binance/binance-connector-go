/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PartialDepthStreamIntervalParameter the model 'PartialDepthStreamIntervalParameter'
type PartialDepthStreamIntervalParameter string

// List of partialDepthStream_interval_parameter
const (
	PartialDepthStreamIntervalParameterInterval0ms   PartialDepthStreamIntervalParameter = "0ms"
	PartialDepthStreamIntervalParameterInterval100ms PartialDepthStreamIntervalParameter = "100ms"
	PartialDepthStreamIntervalParameterInterval500ms PartialDepthStreamIntervalParameter = "500ms"
)

// All allowed values of PartialDepthStreamIntervalParameter enum
var AllowedPartialDepthStreamIntervalParameterEnumValues = []PartialDepthStreamIntervalParameter{
	"0ms",
	"100ms",
	"500ms",
}

func (v *PartialDepthStreamIntervalParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PartialDepthStreamIntervalParameter(value)
	for _, existing := range AllowedPartialDepthStreamIntervalParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PartialDepthStreamIntervalParameter", value)
}

// NewPartialDepthStreamIntervalParameterFromValue returns a pointer to a valid PartialDepthStreamIntervalParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPartialDepthStreamIntervalParameterFromValue(v string) (*PartialDepthStreamIntervalParameter, error) {
	ev := PartialDepthStreamIntervalParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PartialDepthStreamIntervalParameter: valid values are %v", v, AllowedPartialDepthStreamIntervalParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PartialDepthStreamIntervalParameter) IsValid() bool {
	for _, existing := range AllowedPartialDepthStreamIntervalParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to partialDepthStream_interval_parameter value
func (v PartialDepthStreamIntervalParameter) Ptr() *PartialDepthStreamIntervalParameter {
	return &v
}

type NullablePartialDepthStreamIntervalParameter struct {
	value *PartialDepthStreamIntervalParameter
	isSet bool
}

func (v NullablePartialDepthStreamIntervalParameter) Get() *PartialDepthStreamIntervalParameter {
	return v.value
}

func (v *NullablePartialDepthStreamIntervalParameter) Set(val *PartialDepthStreamIntervalParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialDepthStreamIntervalParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialDepthStreamIntervalParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialDepthStreamIntervalParameter(val *PartialDepthStreamIntervalParameter) *NullablePartialDepthStreamIntervalParameter {
	return &NullablePartialDepthStreamIntervalParameter{value: val, isSet: true}
}

func (v NullablePartialDepthStreamIntervalParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialDepthStreamIntervalParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
