/*
Futures (COIN-M) WebSocket Market Streams

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// IndexPriceStreamUpdateSpeedParameter the model 'IndexPriceStreamUpdateSpeedParameter'
type IndexPriceStreamUpdateSpeedParameter string

// List of indexPriceStream_updateSpeed_parameter
const (
	IndexPriceStreamUpdateSpeedParameterUpdateSpeed1s IndexPriceStreamUpdateSpeedParameter = "1s"
)

// All allowed values of IndexPriceStreamUpdateSpeedParameter enum
var AllowedIndexPriceStreamUpdateSpeedParameterEnumValues = []IndexPriceStreamUpdateSpeedParameter{
	"1s",
}

func (v *IndexPriceStreamUpdateSpeedParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IndexPriceStreamUpdateSpeedParameter(value)
	for _, existing := range AllowedIndexPriceStreamUpdateSpeedParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IndexPriceStreamUpdateSpeedParameter", value)
}

// NewIndexPriceStreamUpdateSpeedParameterFromValue returns a pointer to a valid IndexPriceStreamUpdateSpeedParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIndexPriceStreamUpdateSpeedParameterFromValue(v string) (*IndexPriceStreamUpdateSpeedParameter, error) {
	ev := IndexPriceStreamUpdateSpeedParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IndexPriceStreamUpdateSpeedParameter: valid values are %v", v, AllowedIndexPriceStreamUpdateSpeedParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IndexPriceStreamUpdateSpeedParameter) IsValid() bool {
	for _, existing := range AllowedIndexPriceStreamUpdateSpeedParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to indexPriceStream_updateSpeed_parameter value
func (v IndexPriceStreamUpdateSpeedParameter) Ptr() *IndexPriceStreamUpdateSpeedParameter {
	return &v
}

type NullableIndexPriceStreamUpdateSpeedParameter struct {
	value *IndexPriceStreamUpdateSpeedParameter
	isSet bool
}

func (v NullableIndexPriceStreamUpdateSpeedParameter) Get() *IndexPriceStreamUpdateSpeedParameter {
	return v.value
}

func (v *NullableIndexPriceStreamUpdateSpeedParameter) Set(val *IndexPriceStreamUpdateSpeedParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexPriceStreamUpdateSpeedParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexPriceStreamUpdateSpeedParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexPriceStreamUpdateSpeedParameter(val *IndexPriceStreamUpdateSpeedParameter) *NullableIndexPriceStreamUpdateSpeedParameter {
	return &NullableIndexPriceStreamUpdateSpeedParameter{value: val, isSet: true}
}

func (v NullableIndexPriceStreamUpdateSpeedParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexPriceStreamUpdateSpeedParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
