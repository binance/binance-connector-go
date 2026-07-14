/*
Futures (USDⓈ-M) WebSocket Market Streams

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarkPriceStreamUpdateSpeedParameter the model 'MarkPriceStreamUpdateSpeedParameter'
type MarkPriceStreamUpdateSpeedParameter string

// List of markPriceStream_updateSpeed_parameter
const (
	MarkPriceStreamUpdateSpeedParameterUpdateSpeed1s MarkPriceStreamUpdateSpeedParameter = "1s"
)

// All allowed values of MarkPriceStreamUpdateSpeedParameter enum
var AllowedMarkPriceStreamUpdateSpeedParameterEnumValues = []MarkPriceStreamUpdateSpeedParameter{
	"1s",
}

func (v *MarkPriceStreamUpdateSpeedParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarkPriceStreamUpdateSpeedParameter(value)
	for _, existing := range AllowedMarkPriceStreamUpdateSpeedParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarkPriceStreamUpdateSpeedParameter", value)
}

// NewMarkPriceStreamUpdateSpeedParameterFromValue returns a pointer to a valid MarkPriceStreamUpdateSpeedParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarkPriceStreamUpdateSpeedParameterFromValue(v string) (*MarkPriceStreamUpdateSpeedParameter, error) {
	ev := MarkPriceStreamUpdateSpeedParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarkPriceStreamUpdateSpeedParameter: valid values are %v", v, AllowedMarkPriceStreamUpdateSpeedParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarkPriceStreamUpdateSpeedParameter) IsValid() bool {
	for _, existing := range AllowedMarkPriceStreamUpdateSpeedParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to markPriceStream_updateSpeed_parameter value
func (v MarkPriceStreamUpdateSpeedParameter) Ptr() *MarkPriceStreamUpdateSpeedParameter {
	return &v
}

type NullableMarkPriceStreamUpdateSpeedParameter struct {
	value *MarkPriceStreamUpdateSpeedParameter
	isSet bool
}

func (v NullableMarkPriceStreamUpdateSpeedParameter) Get() *MarkPriceStreamUpdateSpeedParameter {
	return v.value
}

func (v *NullableMarkPriceStreamUpdateSpeedParameter) Set(val *MarkPriceStreamUpdateSpeedParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkPriceStreamUpdateSpeedParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkPriceStreamUpdateSpeedParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkPriceStreamUpdateSpeedParameter(val *MarkPriceStreamUpdateSpeedParameter) *NullableMarkPriceStreamUpdateSpeedParameter {
	return &NullableMarkPriceStreamUpdateSpeedParameter{value: val, isSet: true}
}

func (v NullableMarkPriceStreamUpdateSpeedParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkPriceStreamUpdateSpeedParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
