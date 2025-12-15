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

// AdlRiskResponse - struct for AdlRiskResponse
type AdlRiskResponse struct {
	AdlRiskResponse1 *AdlRiskResponse1
	AdlRiskResponse2 *AdlRiskResponse2
}

// AdlRiskResponse1AsAdlRiskResponse is a convenience function that returns AdlRiskResponse1 wrapped in AdlRiskResponse
func AdlRiskResponse1AsAdlRiskResponse(v *AdlRiskResponse1) AdlRiskResponse {
	return AdlRiskResponse{
		AdlRiskResponse1: v,
	}
}

// AdlRiskResponse2AsAdlRiskResponse is a convenience function that returns AdlRiskResponse2 wrapped in AdlRiskResponse
func AdlRiskResponse2AsAdlRiskResponse(v *AdlRiskResponse2) AdlRiskResponse {
	return AdlRiskResponse{
		AdlRiskResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AdlRiskResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AdlRiskResponse1
	err = json.Unmarshal(data, &dst.AdlRiskResponse1)
	if err == nil {
		jsonAdlRiskResponse1, _ := json.Marshal(dst.AdlRiskResponse1)
		if string(jsonAdlRiskResponse1) == "{}" { // empty struct
			dst.AdlRiskResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.AdlRiskResponse1 = nil
	}

	// try to unmarshal data into AdlRiskResponse2
	err = json.Unmarshal(data, &dst.AdlRiskResponse2)
	if err == nil {
		jsonAdlRiskResponse2, _ := json.Marshal(dst.AdlRiskResponse2)
		if string(jsonAdlRiskResponse2) == "{}" { // empty struct
			dst.AdlRiskResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.AdlRiskResponse2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AdlRiskResponse1 = nil
		dst.AdlRiskResponse2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AdlRiskResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AdlRiskResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AdlRiskResponse) MarshalJSON() ([]byte, error) {
	if src.AdlRiskResponse1 != nil {
		return json.Marshal(&src.AdlRiskResponse1)
	}

	if src.AdlRiskResponse2 != nil {
		return json.Marshal(&src.AdlRiskResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AdlRiskResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AdlRiskResponse1 != nil {
		return obj.AdlRiskResponse1
	}

	if obj.AdlRiskResponse2 != nil {
		return obj.AdlRiskResponse2
	}

	// all schemas are nil
	return nil
}

type NullableAdlRiskResponse struct {
	value *AdlRiskResponse
	isSet bool
}

func (v NullableAdlRiskResponse) Get() *AdlRiskResponse {
	return v.value
}

func (v *NullableAdlRiskResponse) Set(val *AdlRiskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAdlRiskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAdlRiskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdlRiskResponse(val *AdlRiskResponse) *NullableAdlRiskResponse {
	return &NullableAdlRiskResponse{value: val, isSet: true}
}

func (v NullableAdlRiskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdlRiskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
