/*
Futures (COIN-M) WebSocket Market Streams

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PartialBookDepthStreamsLevelsParameter the model 'PartialBookDepthStreamsLevelsParameter'
type PartialBookDepthStreamsLevelsParameter string

// List of partialBookDepthStreams_levels_parameter
const (
	PartialBookDepthStreamsLevelsParameterLevels5  PartialBookDepthStreamsLevelsParameter = "5"
	PartialBookDepthStreamsLevelsParameterLevels10 PartialBookDepthStreamsLevelsParameter = "10"
	PartialBookDepthStreamsLevelsParameterLevels20 PartialBookDepthStreamsLevelsParameter = "20"
)

// All allowed values of PartialBookDepthStreamsLevelsParameter enum
var AllowedPartialBookDepthStreamsLevelsParameterEnumValues = []PartialBookDepthStreamsLevelsParameter{
	"5",
	"10",
	"20",
}

func (v *PartialBookDepthStreamsLevelsParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PartialBookDepthStreamsLevelsParameter(value)
	for _, existing := range AllowedPartialBookDepthStreamsLevelsParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PartialBookDepthStreamsLevelsParameter", value)
}

// NewPartialBookDepthStreamsLevelsParameterFromValue returns a pointer to a valid PartialBookDepthStreamsLevelsParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPartialBookDepthStreamsLevelsParameterFromValue(v string) (*PartialBookDepthStreamsLevelsParameter, error) {
	ev := PartialBookDepthStreamsLevelsParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PartialBookDepthStreamsLevelsParameter: valid values are %v", v, AllowedPartialBookDepthStreamsLevelsParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PartialBookDepthStreamsLevelsParameter) IsValid() bool {
	for _, existing := range AllowedPartialBookDepthStreamsLevelsParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to partialBookDepthStreams_levels_parameter value
func (v PartialBookDepthStreamsLevelsParameter) Ptr() *PartialBookDepthStreamsLevelsParameter {
	return &v
}

type NullablePartialBookDepthStreamsLevelsParameter struct {
	value *PartialBookDepthStreamsLevelsParameter
	isSet bool
}

func (v NullablePartialBookDepthStreamsLevelsParameter) Get() *PartialBookDepthStreamsLevelsParameter {
	return v.value
}

func (v *NullablePartialBookDepthStreamsLevelsParameter) Set(val *PartialBookDepthStreamsLevelsParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialBookDepthStreamsLevelsParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialBookDepthStreamsLevelsParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialBookDepthStreamsLevelsParameter(val *PartialBookDepthStreamsLevelsParameter) *NullablePartialBookDepthStreamsLevelsParameter {
	return &NullablePartialBookDepthStreamsLevelsParameter{value: val, isSet: true}
}

func (v NullablePartialBookDepthStreamsLevelsParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialBookDepthStreamsLevelsParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
