/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PartialBookDepthLevelsParameter the model 'PartialBookDepthLevelsParameter'
type PartialBookDepthLevelsParameter string

// List of partialBookDepth_levels_parameter
const (
	PartialBookDepthLevelsParameterLevels5  PartialBookDepthLevelsParameter = "5"
	PartialBookDepthLevelsParameterLevels10 PartialBookDepthLevelsParameter = "10"
	PartialBookDepthLevelsParameterLevels20 PartialBookDepthLevelsParameter = "20"
)

// All allowed values of PartialBookDepthLevelsParameter enum
var AllowedPartialBookDepthLevelsParameterEnumValues = []PartialBookDepthLevelsParameter{
	"5",
	"10",
	"20",
}

func (v *PartialBookDepthLevelsParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PartialBookDepthLevelsParameter(value)
	for _, existing := range AllowedPartialBookDepthLevelsParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PartialBookDepthLevelsParameter", value)
}

// NewPartialBookDepthLevelsParameterFromValue returns a pointer to a valid PartialBookDepthLevelsParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPartialBookDepthLevelsParameterFromValue(v string) (*PartialBookDepthLevelsParameter, error) {
	ev := PartialBookDepthLevelsParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PartialBookDepthLevelsParameter: valid values are %v", v, AllowedPartialBookDepthLevelsParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PartialBookDepthLevelsParameter) IsValid() bool {
	for _, existing := range AllowedPartialBookDepthLevelsParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to partialBookDepth_levels_parameter value
func (v PartialBookDepthLevelsParameter) Ptr() *PartialBookDepthLevelsParameter {
	return &v
}

type NullablePartialBookDepthLevelsParameter struct {
	value *PartialBookDepthLevelsParameter
	isSet bool
}

func (v NullablePartialBookDepthLevelsParameter) Get() *PartialBookDepthLevelsParameter {
	return v.value
}

func (v *NullablePartialBookDepthLevelsParameter) Set(val *PartialBookDepthLevelsParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialBookDepthLevelsParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialBookDepthLevelsParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialBookDepthLevelsParameter(val *PartialBookDepthLevelsParameter) *NullablePartialBookDepthLevelsParameter {
	return &NullablePartialBookDepthLevelsParameter{value: val, isSet: true}
}

func (v NullablePartialBookDepthLevelsParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialBookDepthLevelsParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
