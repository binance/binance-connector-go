/*
Alpha Trading REST API

APIs for Binance Alpha Trading.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// FullDepthLimitParameter the model 'FullDepthLimitParameter'
type FullDepthLimitParameter int64

// List of fullDepth_limit_parameter
const (
	FullDepthLimitParameterLimit5    FullDepthLimitParameter = 5
	FullDepthLimitParameterLimit10   FullDepthLimitParameter = 10
	FullDepthLimitParameterLimit20   FullDepthLimitParameter = 20
	FullDepthLimitParameterLimit50   FullDepthLimitParameter = 50
	FullDepthLimitParameterLimit100  FullDepthLimitParameter = 100
	FullDepthLimitParameterLimit500  FullDepthLimitParameter = 500
	FullDepthLimitParameterLimit1000 FullDepthLimitParameter = 1000
)

// All allowed values of FullDepthLimitParameter enum
var AllowedFullDepthLimitParameterEnumValues = []FullDepthLimitParameter{
	5,
	10,
	20,
	50,
	100,
	500,
	1000,
}

func (v *FullDepthLimitParameter) UnmarshalJSON(src []byte) error {
	var value int64
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FullDepthLimitParameter(value)
	for _, existing := range AllowedFullDepthLimitParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FullDepthLimitParameter", value)
}

// NewFullDepthLimitParameterFromValue returns a pointer to a valid FullDepthLimitParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFullDepthLimitParameterFromValue(v int64) (*FullDepthLimitParameter, error) {
	ev := FullDepthLimitParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FullDepthLimitParameter: valid values are %v", v, AllowedFullDepthLimitParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FullDepthLimitParameter) IsValid() bool {
	for _, existing := range AllowedFullDepthLimitParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to fullDepth_limit_parameter value
func (v FullDepthLimitParameter) Ptr() *FullDepthLimitParameter {
	return &v
}

type NullableFullDepthLimitParameter struct {
	value *FullDepthLimitParameter
	isSet bool
}

func (v NullableFullDepthLimitParameter) Get() *FullDepthLimitParameter {
	return v.value
}

func (v *NullableFullDepthLimitParameter) Set(val *FullDepthLimitParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableFullDepthLimitParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableFullDepthLimitParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullDepthLimitParameter(val *FullDepthLimitParameter) *NullableFullDepthLimitParameter {
	return &NullableFullDepthLimitParameter{value: val, isSet: true}
}

func (v NullableFullDepthLimitParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullDepthLimitParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
