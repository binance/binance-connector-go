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

// SymbolPriceTickerResponse - struct for SymbolPriceTickerResponse
type SymbolPriceTickerResponse struct {
	SymbolPriceTickerResponse1 *SymbolPriceTickerResponse1
	SymbolPriceTickerResponse2 *SymbolPriceTickerResponse2
}

// SymbolPriceTickerResponse1AsSymbolPriceTickerResponse is a convenience function that returns SymbolPriceTickerResponse1 wrapped in SymbolPriceTickerResponse
func SymbolPriceTickerResponse1AsSymbolPriceTickerResponse(v *SymbolPriceTickerResponse1) SymbolPriceTickerResponse {
	return SymbolPriceTickerResponse{
		SymbolPriceTickerResponse1: v,
	}
}

// SymbolPriceTickerResponse2AsSymbolPriceTickerResponse is a convenience function that returns SymbolPriceTickerResponse2 wrapped in SymbolPriceTickerResponse
func SymbolPriceTickerResponse2AsSymbolPriceTickerResponse(v *SymbolPriceTickerResponse2) SymbolPriceTickerResponse {
	return SymbolPriceTickerResponse{
		SymbolPriceTickerResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SymbolPriceTickerResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SymbolPriceTickerResponse1
	err = json.Unmarshal(data, &dst.SymbolPriceTickerResponse1)
	if err == nil {
		jsonSymbolPriceTickerResponse1, _ := json.Marshal(dst.SymbolPriceTickerResponse1)
		if string(jsonSymbolPriceTickerResponse1) == "{}" { // empty struct
			dst.SymbolPriceTickerResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.SymbolPriceTickerResponse1 = nil
	}

	// try to unmarshal data into SymbolPriceTickerResponse2
	err = json.Unmarshal(data, &dst.SymbolPriceTickerResponse2)
	if err == nil {
		jsonSymbolPriceTickerResponse2, _ := json.Marshal(dst.SymbolPriceTickerResponse2)
		if string(jsonSymbolPriceTickerResponse2) == "{}" { // empty struct
			dst.SymbolPriceTickerResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.SymbolPriceTickerResponse2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SymbolPriceTickerResponse1 = nil
		dst.SymbolPriceTickerResponse2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SymbolPriceTickerResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SymbolPriceTickerResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SymbolPriceTickerResponse) MarshalJSON() ([]byte, error) {
	if src.SymbolPriceTickerResponse1 != nil {
		return json.Marshal(&src.SymbolPriceTickerResponse1)
	}

	if src.SymbolPriceTickerResponse2 != nil {
		return json.Marshal(&src.SymbolPriceTickerResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SymbolPriceTickerResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SymbolPriceTickerResponse1 != nil {
		return obj.SymbolPriceTickerResponse1
	}

	if obj.SymbolPriceTickerResponse2 != nil {
		return obj.SymbolPriceTickerResponse2
	}

	// all schemas are nil
	return nil
}

type NullableSymbolPriceTickerResponse struct {
	value *SymbolPriceTickerResponse
	isSet bool
}

func (v NullableSymbolPriceTickerResponse) Get() *SymbolPriceTickerResponse {
	return v.value
}

func (v *NullableSymbolPriceTickerResponse) Set(val *SymbolPriceTickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolPriceTickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolPriceTickerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolPriceTickerResponse(val *SymbolPriceTickerResponse) *NullableSymbolPriceTickerResponse {
	return &NullableSymbolPriceTickerResponse{value: val, isSet: true}
}

func (v NullableSymbolPriceTickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolPriceTickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
