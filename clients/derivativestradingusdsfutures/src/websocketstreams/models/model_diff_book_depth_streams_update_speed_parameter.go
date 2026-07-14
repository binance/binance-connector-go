/*
Futures (USDⓈ-M) WebSocket Market Streams

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// DiffBookDepthStreamsUpdateSpeedParameter the model 'DiffBookDepthStreamsUpdateSpeedParameter'
type DiffBookDepthStreamsUpdateSpeedParameter string

// List of diffBookDepthStreams_updateSpeed_parameter
const (
	DiffBookDepthStreamsUpdateSpeedParameterUpdateSpeed100ms DiffBookDepthStreamsUpdateSpeedParameter = "100ms"
	DiffBookDepthStreamsUpdateSpeedParameterUpdateSpeed500ms DiffBookDepthStreamsUpdateSpeedParameter = "500ms"
)

// All allowed values of DiffBookDepthStreamsUpdateSpeedParameter enum
var AllowedDiffBookDepthStreamsUpdateSpeedParameterEnumValues = []DiffBookDepthStreamsUpdateSpeedParameter{
	"100ms",
	"500ms",
}

func (v *DiffBookDepthStreamsUpdateSpeedParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DiffBookDepthStreamsUpdateSpeedParameter(value)
	for _, existing := range AllowedDiffBookDepthStreamsUpdateSpeedParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DiffBookDepthStreamsUpdateSpeedParameter", value)
}

// NewDiffBookDepthStreamsUpdateSpeedParameterFromValue returns a pointer to a valid DiffBookDepthStreamsUpdateSpeedParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDiffBookDepthStreamsUpdateSpeedParameterFromValue(v string) (*DiffBookDepthStreamsUpdateSpeedParameter, error) {
	ev := DiffBookDepthStreamsUpdateSpeedParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DiffBookDepthStreamsUpdateSpeedParameter: valid values are %v", v, AllowedDiffBookDepthStreamsUpdateSpeedParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DiffBookDepthStreamsUpdateSpeedParameter) IsValid() bool {
	for _, existing := range AllowedDiffBookDepthStreamsUpdateSpeedParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to diffBookDepthStreams_updateSpeed_parameter value
func (v DiffBookDepthStreamsUpdateSpeedParameter) Ptr() *DiffBookDepthStreamsUpdateSpeedParameter {
	return &v
}

type NullableDiffBookDepthStreamsUpdateSpeedParameter struct {
	value *DiffBookDepthStreamsUpdateSpeedParameter
	isSet bool
}

func (v NullableDiffBookDepthStreamsUpdateSpeedParameter) Get() *DiffBookDepthStreamsUpdateSpeedParameter {
	return v.value
}

func (v *NullableDiffBookDepthStreamsUpdateSpeedParameter) Set(val *DiffBookDepthStreamsUpdateSpeedParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDiffBookDepthStreamsUpdateSpeedParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDiffBookDepthStreamsUpdateSpeedParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiffBookDepthStreamsUpdateSpeedParameter(val *DiffBookDepthStreamsUpdateSpeedParameter) *NullableDiffBookDepthStreamsUpdateSpeedParameter {
	return &NullableDiffBookDepthStreamsUpdateSpeedParameter{value: val, isSet: true}
}

func (v NullableDiffBookDepthStreamsUpdateSpeedParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiffBookDepthStreamsUpdateSpeedParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
