/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// TickerPriceResponse - struct for TickerPriceResponse
type TickerPriceResponse struct {
	TickerPriceResponse1 *TickerPriceResponse1
	TickerPriceResponse2 *TickerPriceResponse2
}

// TickerPriceResponse1AsTickerPriceResponse is a convenience function that returns TickerPriceResponse1 wrapped in TickerPriceResponse
func TickerPriceResponse1AsTickerPriceResponse(v *TickerPriceResponse1) TickerPriceResponse {
	return TickerPriceResponse{
		TickerPriceResponse1: v,
	}
}

// TickerPriceResponse2AsTickerPriceResponse is a convenience function that returns TickerPriceResponse2 wrapped in TickerPriceResponse
func TickerPriceResponse2AsTickerPriceResponse(v *TickerPriceResponse2) TickerPriceResponse {
	return TickerPriceResponse{
		TickerPriceResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TickerPriceResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TickerPriceResponse1
	err = json.Unmarshal(data, &dst.TickerPriceResponse1)
	if err == nil {
		jsonTickerPriceResponse1, _ := json.Marshal(dst.TickerPriceResponse1)
		if string(jsonTickerPriceResponse1) == "{}" { // empty struct
			dst.TickerPriceResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.TickerPriceResponse1 = nil
	}

	// try to unmarshal data into TickerPriceResponse2
	err = json.Unmarshal(data, &dst.TickerPriceResponse2)
	if err == nil {
		jsonTickerPriceResponse2, _ := json.Marshal(dst.TickerPriceResponse2)
		if string(jsonTickerPriceResponse2) == "{}" { // empty struct
			dst.TickerPriceResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.TickerPriceResponse2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.TickerPriceResponse1 = nil
		dst.TickerPriceResponse2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TickerPriceResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TickerPriceResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TickerPriceResponse) MarshalJSON() ([]byte, error) {
	if src.TickerPriceResponse1 != nil {
		return json.Marshal(&src.TickerPriceResponse1)
	}

	if src.TickerPriceResponse2 != nil {
		return json.Marshal(&src.TickerPriceResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TickerPriceResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.TickerPriceResponse1 != nil {
		return obj.TickerPriceResponse1
	}

	if obj.TickerPriceResponse2 != nil {
		return obj.TickerPriceResponse2
	}

	// all schemas are nil
	return nil
}

type NullableTickerPriceResponse struct {
	value *TickerPriceResponse
	isSet bool
}

func (v NullableTickerPriceResponse) Get() *TickerPriceResponse {
	return v.value
}

func (v *NullableTickerPriceResponse) Set(val *TickerPriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerPriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerPriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerPriceResponse(val *TickerPriceResponse) *NullableTickerPriceResponse {
	return &NullableTickerPriceResponse{value: val, isSet: true}
}

func (v NullableTickerPriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerPriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
