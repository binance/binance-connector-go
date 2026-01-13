/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// TickerBookTickerResponse - struct for TickerBookTickerResponse
type TickerBookTickerResponse struct {
	TickerBookTickerResponse1 *TickerBookTickerResponse1
	TickerBookTickerResponse2 *TickerBookTickerResponse2
}

// TickerBookTickerResponse1AsTickerBookTickerResponse is a convenience function that returns TickerBookTickerResponse1 wrapped in TickerBookTickerResponse
func TickerBookTickerResponse1AsTickerBookTickerResponse(v *TickerBookTickerResponse1) TickerBookTickerResponse {
	return TickerBookTickerResponse{
		TickerBookTickerResponse1: v,
	}
}

// TickerBookTickerResponse2AsTickerBookTickerResponse is a convenience function that returns TickerBookTickerResponse2 wrapped in TickerBookTickerResponse
func TickerBookTickerResponse2AsTickerBookTickerResponse(v *TickerBookTickerResponse2) TickerBookTickerResponse {
	return TickerBookTickerResponse{
		TickerBookTickerResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TickerBookTickerResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TickerBookTickerResponse1
	err = json.Unmarshal(data, &dst.TickerBookTickerResponse1)
	if err == nil {
		jsonTickerBookTickerResponse1, _ := json.Marshal(dst.TickerBookTickerResponse1)
		if string(jsonTickerBookTickerResponse1) == "{}" { // empty struct
			dst.TickerBookTickerResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.TickerBookTickerResponse1 = nil
	}

	// try to unmarshal data into TickerBookTickerResponse2
	err = json.Unmarshal(data, &dst.TickerBookTickerResponse2)
	if err == nil {
		jsonTickerBookTickerResponse2, _ := json.Marshal(dst.TickerBookTickerResponse2)
		if string(jsonTickerBookTickerResponse2) == "{}" { // empty struct
			dst.TickerBookTickerResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.TickerBookTickerResponse2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.TickerBookTickerResponse1 = nil
		dst.TickerBookTickerResponse2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TickerBookTickerResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TickerBookTickerResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TickerBookTickerResponse) MarshalJSON() ([]byte, error) {
	if src.TickerBookTickerResponse1 != nil {
		return json.Marshal(&src.TickerBookTickerResponse1)
	}

	if src.TickerBookTickerResponse2 != nil {
		return json.Marshal(&src.TickerBookTickerResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TickerBookTickerResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.TickerBookTickerResponse1 != nil {
		return obj.TickerBookTickerResponse1
	}

	if obj.TickerBookTickerResponse2 != nil {
		return obj.TickerBookTickerResponse2
	}

	// all schemas are nil
	return nil
}

type NullableTickerBookTickerResponse struct {
	value *TickerBookTickerResponse
	isSet bool
}

func (v NullableTickerBookTickerResponse) Get() *TickerBookTickerResponse {
	return v.value
}

func (v *NullableTickerBookTickerResponse) Set(val *TickerBookTickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerBookTickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerBookTickerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerBookTickerResponse(val *TickerBookTickerResponse) *NullableTickerBookTickerResponse {
	return &NullableTickerBookTickerResponse{value: val, isSet: true}
}

func (v NullableTickerBookTickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerBookTickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
