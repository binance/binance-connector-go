/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// TickerBookResponse - struct for TickerBookResponse
type TickerBookResponse struct {
	TickerBookResponse1 *TickerBookResponse1
	TickerBookResponse2 *TickerBookResponse2
}

// TickerBookResponse1AsTickerBookResponse is a convenience function that returns TickerBookResponse1 wrapped in TickerBookResponse
func TickerBookResponse1AsTickerBookResponse(v *TickerBookResponse1) TickerBookResponse {
	return TickerBookResponse{
		TickerBookResponse1: v,
	}
}

// TickerBookResponse2AsTickerBookResponse is a convenience function that returns TickerBookResponse2 wrapped in TickerBookResponse
func TickerBookResponse2AsTickerBookResponse(v *TickerBookResponse2) TickerBookResponse {
	return TickerBookResponse{
		TickerBookResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TickerBookResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TickerBookResponse1
	err = json.Unmarshal(data, &dst.TickerBookResponse1)
	if err == nil {
		jsonTickerBookResponse1, _ := json.Marshal(dst.TickerBookResponse1)
		if string(jsonTickerBookResponse1) == "{}" { // empty struct
			dst.TickerBookResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.TickerBookResponse1 = nil
	}

	// try to unmarshal data into TickerBookResponse2
	err = json.Unmarshal(data, &dst.TickerBookResponse2)
	if err == nil {
		jsonTickerBookResponse2, _ := json.Marshal(dst.TickerBookResponse2)
		if string(jsonTickerBookResponse2) == "{}" { // empty struct
			dst.TickerBookResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.TickerBookResponse2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.TickerBookResponse1 = nil
		dst.TickerBookResponse2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TickerBookResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TickerBookResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TickerBookResponse) MarshalJSON() ([]byte, error) {
	if src.TickerBookResponse1 != nil {
		return json.Marshal(&src.TickerBookResponse1)
	}

	if src.TickerBookResponse2 != nil {
		return json.Marshal(&src.TickerBookResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TickerBookResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.TickerBookResponse1 != nil {
		return obj.TickerBookResponse1
	}

	if obj.TickerBookResponse2 != nil {
		return obj.TickerBookResponse2
	}

	// all schemas are nil
	return nil
}

type NullableTickerBookResponse struct {
	value *TickerBookResponse
	isSet bool
}

func (v NullableTickerBookResponse) Get() *TickerBookResponse {
	return v.value
}

func (v *NullableTickerBookResponse) Set(val *TickerBookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerBookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerBookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerBookResponse(val *TickerBookResponse) *NullableTickerBookResponse {
	return &NullableTickerBookResponse{value: val, isSet: true}
}

func (v NullableTickerBookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerBookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
