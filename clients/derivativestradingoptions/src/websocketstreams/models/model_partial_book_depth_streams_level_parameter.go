/*
Options WebSocket Market Streams

Access market data, manage accounts, and trade Binance Options.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PartialBookDepthStreamsLevelParameter the model 'PartialBookDepthStreamsLevelParameter'
type PartialBookDepthStreamsLevelParameter string

// List of partialBookDepthStreams_level_parameter
const (
	PartialBookDepthStreamsLevelParameterLevel5  PartialBookDepthStreamsLevelParameter = "5"
	PartialBookDepthStreamsLevelParameterLevel10 PartialBookDepthStreamsLevelParameter = "10"
	PartialBookDepthStreamsLevelParameterLevel20 PartialBookDepthStreamsLevelParameter = "20"
)

// All allowed values of PartialBookDepthStreamsLevelParameter enum
var AllowedPartialBookDepthStreamsLevelParameterEnumValues = []PartialBookDepthStreamsLevelParameter{
	"5",
	"10",
	"20",
}

func (v *PartialBookDepthStreamsLevelParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PartialBookDepthStreamsLevelParameter(value)
	for _, existing := range AllowedPartialBookDepthStreamsLevelParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PartialBookDepthStreamsLevelParameter", value)
}

// NewPartialBookDepthStreamsLevelParameterFromValue returns a pointer to a valid PartialBookDepthStreamsLevelParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPartialBookDepthStreamsLevelParameterFromValue(v string) (*PartialBookDepthStreamsLevelParameter, error) {
	ev := PartialBookDepthStreamsLevelParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PartialBookDepthStreamsLevelParameter: valid values are %v", v, AllowedPartialBookDepthStreamsLevelParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PartialBookDepthStreamsLevelParameter) IsValid() bool {
	for _, existing := range AllowedPartialBookDepthStreamsLevelParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to partialBookDepthStreams_level_parameter value
func (v PartialBookDepthStreamsLevelParameter) Ptr() *PartialBookDepthStreamsLevelParameter {
	return &v
}

type NullablePartialBookDepthStreamsLevelParameter struct {
	value *PartialBookDepthStreamsLevelParameter
	isSet bool
}

func (v NullablePartialBookDepthStreamsLevelParameter) Get() *PartialBookDepthStreamsLevelParameter {
	return v.value
}

func (v *NullablePartialBookDepthStreamsLevelParameter) Set(val *PartialBookDepthStreamsLevelParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialBookDepthStreamsLevelParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialBookDepthStreamsLevelParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialBookDepthStreamsLevelParameter(val *PartialBookDepthStreamsLevelParameter) *NullablePartialBookDepthStreamsLevelParameter {
	return &NullablePartialBookDepthStreamsLevelParameter{value: val, isSet: true}
}

func (v NullablePartialBookDepthStreamsLevelParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialBookDepthStreamsLevelParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
