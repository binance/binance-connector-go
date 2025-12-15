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

// MultiAssetsModeAssetIndexResponse - struct for MultiAssetsModeAssetIndexResponse
type MultiAssetsModeAssetIndexResponse struct {
	MultiAssetsModeAssetIndexResponse1 *MultiAssetsModeAssetIndexResponse1
	MultiAssetsModeAssetIndexResponse2 *MultiAssetsModeAssetIndexResponse2
}

// MultiAssetsModeAssetIndexResponse1AsMultiAssetsModeAssetIndexResponse is a convenience function that returns MultiAssetsModeAssetIndexResponse1 wrapped in MultiAssetsModeAssetIndexResponse
func MultiAssetsModeAssetIndexResponse1AsMultiAssetsModeAssetIndexResponse(v *MultiAssetsModeAssetIndexResponse1) MultiAssetsModeAssetIndexResponse {
	return MultiAssetsModeAssetIndexResponse{
		MultiAssetsModeAssetIndexResponse1: v,
	}
}

// MultiAssetsModeAssetIndexResponse2AsMultiAssetsModeAssetIndexResponse is a convenience function that returns MultiAssetsModeAssetIndexResponse2 wrapped in MultiAssetsModeAssetIndexResponse
func MultiAssetsModeAssetIndexResponse2AsMultiAssetsModeAssetIndexResponse(v *MultiAssetsModeAssetIndexResponse2) MultiAssetsModeAssetIndexResponse {
	return MultiAssetsModeAssetIndexResponse{
		MultiAssetsModeAssetIndexResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MultiAssetsModeAssetIndexResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MultiAssetsModeAssetIndexResponse1
	err = json.Unmarshal(data, &dst.MultiAssetsModeAssetIndexResponse1)
	if err == nil {
		jsonMultiAssetsModeAssetIndexResponse1, _ := json.Marshal(dst.MultiAssetsModeAssetIndexResponse1)
		if string(jsonMultiAssetsModeAssetIndexResponse1) == "{}" { // empty struct
			dst.MultiAssetsModeAssetIndexResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.MultiAssetsModeAssetIndexResponse1 = nil
	}

	// try to unmarshal data into MultiAssetsModeAssetIndexResponse2
	err = json.Unmarshal(data, &dst.MultiAssetsModeAssetIndexResponse2)
	if err == nil {
		jsonMultiAssetsModeAssetIndexResponse2, _ := json.Marshal(dst.MultiAssetsModeAssetIndexResponse2)
		if string(jsonMultiAssetsModeAssetIndexResponse2) == "{}" { // empty struct
			dst.MultiAssetsModeAssetIndexResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.MultiAssetsModeAssetIndexResponse2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MultiAssetsModeAssetIndexResponse1 = nil
		dst.MultiAssetsModeAssetIndexResponse2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MultiAssetsModeAssetIndexResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MultiAssetsModeAssetIndexResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MultiAssetsModeAssetIndexResponse) MarshalJSON() ([]byte, error) {
	if src.MultiAssetsModeAssetIndexResponse1 != nil {
		return json.Marshal(&src.MultiAssetsModeAssetIndexResponse1)
	}

	if src.MultiAssetsModeAssetIndexResponse2 != nil {
		return json.Marshal(&src.MultiAssetsModeAssetIndexResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MultiAssetsModeAssetIndexResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MultiAssetsModeAssetIndexResponse1 != nil {
		return obj.MultiAssetsModeAssetIndexResponse1
	}

	if obj.MultiAssetsModeAssetIndexResponse2 != nil {
		return obj.MultiAssetsModeAssetIndexResponse2
	}

	// all schemas are nil
	return nil
}

type NullableMultiAssetsModeAssetIndexResponse struct {
	value *MultiAssetsModeAssetIndexResponse
	isSet bool
}

func (v NullableMultiAssetsModeAssetIndexResponse) Get() *MultiAssetsModeAssetIndexResponse {
	return v.value
}

func (v *NullableMultiAssetsModeAssetIndexResponse) Set(val *MultiAssetsModeAssetIndexResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiAssetsModeAssetIndexResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiAssetsModeAssetIndexResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiAssetsModeAssetIndexResponse(val *MultiAssetsModeAssetIndexResponse) *NullableMultiAssetsModeAssetIndexResponse {
	return &NullableMultiAssetsModeAssetIndexResponse{value: val, isSet: true}
}

func (v NullableMultiAssetsModeAssetIndexResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiAssetsModeAssetIndexResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
