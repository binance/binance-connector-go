/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// TickerResponse - struct for TickerResponse
type TickerResponse struct {
	TickerResponse1 *TickerResponse1
	TickerResponse2 *TickerResponse2
}

// TickerResponse1AsTickerResponse is a convenience function that returns TickerResponse1 wrapped in TickerResponse
func TickerResponse1AsTickerResponse(v *TickerResponse1) TickerResponse {
	return TickerResponse{
		TickerResponse1: v,
	}
}

// TickerResponse2AsTickerResponse is a convenience function that returns TickerResponse2 wrapped in TickerResponse
func TickerResponse2AsTickerResponse(v *TickerResponse2) TickerResponse {
	return TickerResponse{
		TickerResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TickerResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TickerResponse1
	err = json.Unmarshal(data, &dst.TickerResponse1)
	if err == nil {
		jsonTickerResponse1, _ := json.Marshal(dst.TickerResponse1)
		if string(jsonTickerResponse1) == "{}" { // empty struct
			dst.TickerResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.TickerResponse1 = nil
	}

	// try to unmarshal data into TickerResponse2
	err = json.Unmarshal(data, &dst.TickerResponse2)
	if err == nil {
		jsonTickerResponse2, _ := json.Marshal(dst.TickerResponse2)
		if string(jsonTickerResponse2) == "{}" { // empty struct
			dst.TickerResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.TickerResponse2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.TickerResponse1 = nil
		dst.TickerResponse2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TickerResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TickerResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TickerResponse) MarshalJSON() ([]byte, error) {
	if src.TickerResponse1 != nil {
		return json.Marshal(&src.TickerResponse1)
	}

	if src.TickerResponse2 != nil {
		return json.Marshal(&src.TickerResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TickerResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.TickerResponse1 != nil {
		return obj.TickerResponse1
	}

	if obj.TickerResponse2 != nil {
		return obj.TickerResponse2
	}

	// all schemas are nil
	return nil
}

type NullableTickerResponse struct {
	value *TickerResponse
	isSet bool
}

func (v NullableTickerResponse) Get() *TickerResponse {
	return v.value
}

func (v *NullableTickerResponse) Set(val *TickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerResponse(val *TickerResponse) *NullableTickerResponse {
	return &NullableTickerResponse{value: val, isSet: true}
}

func (v NullableTickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
