/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// SymbolOrderBookTickerResponse - struct for SymbolOrderBookTickerResponse
type SymbolOrderBookTickerResponse struct {
	SymbolOrderBookTickerResponse1 *SymbolOrderBookTickerResponse1
	SymbolOrderBookTickerResponse2 *SymbolOrderBookTickerResponse2
}

// SymbolOrderBookTickerResponse1AsSymbolOrderBookTickerResponse is a convenience function that returns SymbolOrderBookTickerResponse1 wrapped in SymbolOrderBookTickerResponse
func SymbolOrderBookTickerResponse1AsSymbolOrderBookTickerResponse(v *SymbolOrderBookTickerResponse1) SymbolOrderBookTickerResponse {
	return SymbolOrderBookTickerResponse{
		SymbolOrderBookTickerResponse1: v,
	}
}

// SymbolOrderBookTickerResponse2AsSymbolOrderBookTickerResponse is a convenience function that returns SymbolOrderBookTickerResponse2 wrapped in SymbolOrderBookTickerResponse
func SymbolOrderBookTickerResponse2AsSymbolOrderBookTickerResponse(v *SymbolOrderBookTickerResponse2) SymbolOrderBookTickerResponse {
	return SymbolOrderBookTickerResponse{
		SymbolOrderBookTickerResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SymbolOrderBookTickerResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SymbolOrderBookTickerResponse1
	err = json.Unmarshal(data, &dst.SymbolOrderBookTickerResponse1)
	if err == nil {
		jsonSymbolOrderBookTickerResponse1, _ := json.Marshal(dst.SymbolOrderBookTickerResponse1)
		if string(jsonSymbolOrderBookTickerResponse1) == "{}" { // empty struct
			dst.SymbolOrderBookTickerResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.SymbolOrderBookTickerResponse1 = nil
	}

	// try to unmarshal data into SymbolOrderBookTickerResponse2
	err = json.Unmarshal(data, &dst.SymbolOrderBookTickerResponse2)
	if err == nil {
		jsonSymbolOrderBookTickerResponse2, _ := json.Marshal(dst.SymbolOrderBookTickerResponse2)
		if string(jsonSymbolOrderBookTickerResponse2) == "{}" { // empty struct
			dst.SymbolOrderBookTickerResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.SymbolOrderBookTickerResponse2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SymbolOrderBookTickerResponse1 = nil
		dst.SymbolOrderBookTickerResponse2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SymbolOrderBookTickerResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SymbolOrderBookTickerResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SymbolOrderBookTickerResponse) MarshalJSON() ([]byte, error) {
	if src.SymbolOrderBookTickerResponse1 != nil {
		return json.Marshal(&src.SymbolOrderBookTickerResponse1)
	}

	if src.SymbolOrderBookTickerResponse2 != nil {
		return json.Marshal(&src.SymbolOrderBookTickerResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SymbolOrderBookTickerResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SymbolOrderBookTickerResponse1 != nil {
		return obj.SymbolOrderBookTickerResponse1
	}

	if obj.SymbolOrderBookTickerResponse2 != nil {
		return obj.SymbolOrderBookTickerResponse2
	}

	// all schemas are nil
	return nil
}

type NullableSymbolOrderBookTickerResponse struct {
	value *SymbolOrderBookTickerResponse
	isSet bool
}

func (v NullableSymbolOrderBookTickerResponse) Get() *SymbolOrderBookTickerResponse {
	return v.value
}

func (v *NullableSymbolOrderBookTickerResponse) Set(val *SymbolOrderBookTickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolOrderBookTickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolOrderBookTickerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolOrderBookTickerResponse(val *SymbolOrderBookTickerResponse) *NullableSymbolOrderBookTickerResponse {
	return &NullableSymbolOrderBookTickerResponse{value: val, isSet: true}
}

func (v NullableSymbolOrderBookTickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolOrderBookTickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
