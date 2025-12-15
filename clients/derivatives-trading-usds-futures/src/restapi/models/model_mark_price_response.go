/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarkPriceResponse - struct for MarkPriceResponse
type MarkPriceResponse struct {
	MarkPriceResponse1 *MarkPriceResponse1
	MarkPriceResponse2 *MarkPriceResponse2
}

// MarkPriceResponse1AsMarkPriceResponse is a convenience function that returns MarkPriceResponse1 wrapped in MarkPriceResponse
func MarkPriceResponse1AsMarkPriceResponse(v *MarkPriceResponse1) MarkPriceResponse {
	return MarkPriceResponse{
		MarkPriceResponse1: v,
	}
}

// MarkPriceResponse2AsMarkPriceResponse is a convenience function that returns MarkPriceResponse2 wrapped in MarkPriceResponse
func MarkPriceResponse2AsMarkPriceResponse(v *MarkPriceResponse2) MarkPriceResponse {
	return MarkPriceResponse{
		MarkPriceResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MarkPriceResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MarkPriceResponse1
	err = json.Unmarshal(data, &dst.MarkPriceResponse1)
	if err == nil {
		jsonMarkPriceResponse1, _ := json.Marshal(dst.MarkPriceResponse1)
		if string(jsonMarkPriceResponse1) == "{}" { // empty struct
			dst.MarkPriceResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.MarkPriceResponse1 = nil
	}

	// try to unmarshal data into MarkPriceResponse2
	err = json.Unmarshal(data, &dst.MarkPriceResponse2)
	if err == nil {
		jsonMarkPriceResponse2, _ := json.Marshal(dst.MarkPriceResponse2)
		if string(jsonMarkPriceResponse2) == "{}" { // empty struct
			dst.MarkPriceResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.MarkPriceResponse2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MarkPriceResponse1 = nil
		dst.MarkPriceResponse2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MarkPriceResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MarkPriceResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MarkPriceResponse) MarshalJSON() ([]byte, error) {
	if src.MarkPriceResponse1 != nil {
		return json.Marshal(&src.MarkPriceResponse1)
	}

	if src.MarkPriceResponse2 != nil {
		return json.Marshal(&src.MarkPriceResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MarkPriceResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MarkPriceResponse1 != nil {
		return obj.MarkPriceResponse1
	}

	if obj.MarkPriceResponse2 != nil {
		return obj.MarkPriceResponse2
	}

	// all schemas are nil
	return nil
}

type NullableMarkPriceResponse struct {
	value *MarkPriceResponse
	isSet bool
}

func (v NullableMarkPriceResponse) Get() *MarkPriceResponse {
	return v.value
}

func (v *NullableMarkPriceResponse) Set(val *MarkPriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkPriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkPriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkPriceResponse(val *MarkPriceResponse) *NullableMarkPriceResponse {
	return &NullableMarkPriceResponse{value: val, isSet: true}
}

func (v NullableMarkPriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkPriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
