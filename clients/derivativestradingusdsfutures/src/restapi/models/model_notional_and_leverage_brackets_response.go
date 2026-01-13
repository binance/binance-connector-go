/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NotionalAndLeverageBracketsResponse - struct for NotionalAndLeverageBracketsResponse
type NotionalAndLeverageBracketsResponse struct {
	NotionalAndLeverageBracketsResponse1 *NotionalAndLeverageBracketsResponse1
	NotionalAndLeverageBracketsResponse2 *NotionalAndLeverageBracketsResponse2
}

// NotionalAndLeverageBracketsResponse1AsNotionalAndLeverageBracketsResponse is a convenience function that returns NotionalAndLeverageBracketsResponse1 wrapped in NotionalAndLeverageBracketsResponse
func NotionalAndLeverageBracketsResponse1AsNotionalAndLeverageBracketsResponse(v *NotionalAndLeverageBracketsResponse1) NotionalAndLeverageBracketsResponse {
	return NotionalAndLeverageBracketsResponse{
		NotionalAndLeverageBracketsResponse1: v,
	}
}

// NotionalAndLeverageBracketsResponse2AsNotionalAndLeverageBracketsResponse is a convenience function that returns NotionalAndLeverageBracketsResponse2 wrapped in NotionalAndLeverageBracketsResponse
func NotionalAndLeverageBracketsResponse2AsNotionalAndLeverageBracketsResponse(v *NotionalAndLeverageBracketsResponse2) NotionalAndLeverageBracketsResponse {
	return NotionalAndLeverageBracketsResponse{
		NotionalAndLeverageBracketsResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NotionalAndLeverageBracketsResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NotionalAndLeverageBracketsResponse1
	err = json.Unmarshal(data, &dst.NotionalAndLeverageBracketsResponse1)
	if err == nil {
		jsonNotionalAndLeverageBracketsResponse1, _ := json.Marshal(dst.NotionalAndLeverageBracketsResponse1)
		if string(jsonNotionalAndLeverageBracketsResponse1) == "{}" { // empty struct
			dst.NotionalAndLeverageBracketsResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.NotionalAndLeverageBracketsResponse1 = nil
	}

	// try to unmarshal data into NotionalAndLeverageBracketsResponse2
	err = json.Unmarshal(data, &dst.NotionalAndLeverageBracketsResponse2)
	if err == nil {
		jsonNotionalAndLeverageBracketsResponse2, _ := json.Marshal(dst.NotionalAndLeverageBracketsResponse2)
		if string(jsonNotionalAndLeverageBracketsResponse2) == "{}" { // empty struct
			dst.NotionalAndLeverageBracketsResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.NotionalAndLeverageBracketsResponse2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.NotionalAndLeverageBracketsResponse1 = nil
		dst.NotionalAndLeverageBracketsResponse2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(NotionalAndLeverageBracketsResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(NotionalAndLeverageBracketsResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NotionalAndLeverageBracketsResponse) MarshalJSON() ([]byte, error) {
	if src.NotionalAndLeverageBracketsResponse1 != nil {
		return json.Marshal(&src.NotionalAndLeverageBracketsResponse1)
	}

	if src.NotionalAndLeverageBracketsResponse2 != nil {
		return json.Marshal(&src.NotionalAndLeverageBracketsResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NotionalAndLeverageBracketsResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.NotionalAndLeverageBracketsResponse1 != nil {
		return obj.NotionalAndLeverageBracketsResponse1
	}

	if obj.NotionalAndLeverageBracketsResponse2 != nil {
		return obj.NotionalAndLeverageBracketsResponse2
	}

	// all schemas are nil
	return nil
}

type NullableNotionalAndLeverageBracketsResponse struct {
	value *NotionalAndLeverageBracketsResponse
	isSet bool
}

func (v NullableNotionalAndLeverageBracketsResponse) Get() *NotionalAndLeverageBracketsResponse {
	return v.value
}

func (v *NullableNotionalAndLeverageBracketsResponse) Set(val *NotionalAndLeverageBracketsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotionalAndLeverageBracketsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotionalAndLeverageBracketsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotionalAndLeverageBracketsResponse(val *NotionalAndLeverageBracketsResponse) *NullableNotionalAndLeverageBracketsResponse {
	return &NullableNotionalAndLeverageBracketsResponse{value: val, isSet: true}
}

func (v NullableNotionalAndLeverageBracketsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotionalAndLeverageBracketsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
