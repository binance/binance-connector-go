/*
Spot WebSocket Market Streams

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// DiffBookDepthUpdateSpeedParameter the model 'DiffBookDepthUpdateSpeedParameter'
type DiffBookDepthUpdateSpeedParameter string

// List of diffBookDepth_updateSpeed_parameter
const (
	DiffBookDepthUpdateSpeedParameterUpdateSpeed100ms DiffBookDepthUpdateSpeedParameter = "100ms"
)

// All allowed values of DiffBookDepthUpdateSpeedParameter enum
var AllowedDiffBookDepthUpdateSpeedParameterEnumValues = []DiffBookDepthUpdateSpeedParameter{
	"100ms",
}

func (v *DiffBookDepthUpdateSpeedParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DiffBookDepthUpdateSpeedParameter(value)
	for _, existing := range AllowedDiffBookDepthUpdateSpeedParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DiffBookDepthUpdateSpeedParameter", value)
}

// NewDiffBookDepthUpdateSpeedParameterFromValue returns a pointer to a valid DiffBookDepthUpdateSpeedParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDiffBookDepthUpdateSpeedParameterFromValue(v string) (*DiffBookDepthUpdateSpeedParameter, error) {
	ev := DiffBookDepthUpdateSpeedParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DiffBookDepthUpdateSpeedParameter: valid values are %v", v, AllowedDiffBookDepthUpdateSpeedParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DiffBookDepthUpdateSpeedParameter) IsValid() bool {
	for _, existing := range AllowedDiffBookDepthUpdateSpeedParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to diffBookDepth_updateSpeed_parameter value
func (v DiffBookDepthUpdateSpeedParameter) Ptr() *DiffBookDepthUpdateSpeedParameter {
	return &v
}

type NullableDiffBookDepthUpdateSpeedParameter struct {
	value *DiffBookDepthUpdateSpeedParameter
	isSet bool
}

func (v NullableDiffBookDepthUpdateSpeedParameter) Get() *DiffBookDepthUpdateSpeedParameter {
	return v.value
}

func (v *NullableDiffBookDepthUpdateSpeedParameter) Set(val *DiffBookDepthUpdateSpeedParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDiffBookDepthUpdateSpeedParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDiffBookDepthUpdateSpeedParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiffBookDepthUpdateSpeedParameter(val *DiffBookDepthUpdateSpeedParameter) *NullableDiffBookDepthUpdateSpeedParameter {
	return &NullableDiffBookDepthUpdateSpeedParameter{value: val, isSet: true}
}

func (v NullableDiffBookDepthUpdateSpeedParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiffBookDepthUpdateSpeedParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
