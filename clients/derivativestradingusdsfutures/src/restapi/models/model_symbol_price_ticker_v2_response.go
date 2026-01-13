/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// SymbolPriceTickerV2Response - struct for SymbolPriceTickerV2Response
type SymbolPriceTickerV2Response struct {
	SymbolPriceTickerV2Response1 *SymbolPriceTickerV2Response1
	SymbolPriceTickerV2Response2 *SymbolPriceTickerV2Response2
}

// SymbolPriceTickerV2Response1AsSymbolPriceTickerV2Response is a convenience function that returns SymbolPriceTickerV2Response1 wrapped in SymbolPriceTickerV2Response
func SymbolPriceTickerV2Response1AsSymbolPriceTickerV2Response(v *SymbolPriceTickerV2Response1) SymbolPriceTickerV2Response {
	return SymbolPriceTickerV2Response{
		SymbolPriceTickerV2Response1: v,
	}
}

// SymbolPriceTickerV2Response2AsSymbolPriceTickerV2Response is a convenience function that returns SymbolPriceTickerV2Response2 wrapped in SymbolPriceTickerV2Response
func SymbolPriceTickerV2Response2AsSymbolPriceTickerV2Response(v *SymbolPriceTickerV2Response2) SymbolPriceTickerV2Response {
	return SymbolPriceTickerV2Response{
		SymbolPriceTickerV2Response2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SymbolPriceTickerV2Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SymbolPriceTickerV2Response1
	err = json.Unmarshal(data, &dst.SymbolPriceTickerV2Response1)
	if err == nil {
		jsonSymbolPriceTickerV2Response1, _ := json.Marshal(dst.SymbolPriceTickerV2Response1)
		if string(jsonSymbolPriceTickerV2Response1) == "{}" { // empty struct
			dst.SymbolPriceTickerV2Response1 = nil
		} else {
			match++
		}
	} else {
		dst.SymbolPriceTickerV2Response1 = nil
	}

	// try to unmarshal data into SymbolPriceTickerV2Response2
	err = json.Unmarshal(data, &dst.SymbolPriceTickerV2Response2)
	if err == nil {
		jsonSymbolPriceTickerV2Response2, _ := json.Marshal(dst.SymbolPriceTickerV2Response2)
		if string(jsonSymbolPriceTickerV2Response2) == "{}" { // empty struct
			dst.SymbolPriceTickerV2Response2 = nil
		} else {
			match++
		}
	} else {
		dst.SymbolPriceTickerV2Response2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SymbolPriceTickerV2Response1 = nil
		dst.SymbolPriceTickerV2Response2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SymbolPriceTickerV2Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SymbolPriceTickerV2Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SymbolPriceTickerV2Response) MarshalJSON() ([]byte, error) {
	if src.SymbolPriceTickerV2Response1 != nil {
		return json.Marshal(&src.SymbolPriceTickerV2Response1)
	}

	if src.SymbolPriceTickerV2Response2 != nil {
		return json.Marshal(&src.SymbolPriceTickerV2Response2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SymbolPriceTickerV2Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SymbolPriceTickerV2Response1 != nil {
		return obj.SymbolPriceTickerV2Response1
	}

	if obj.SymbolPriceTickerV2Response2 != nil {
		return obj.SymbolPriceTickerV2Response2
	}

	// all schemas are nil
	return nil
}

type NullableSymbolPriceTickerV2Response struct {
	value *SymbolPriceTickerV2Response
	isSet bool
}

func (v NullableSymbolPriceTickerV2Response) Get() *SymbolPriceTickerV2Response {
	return v.value
}

func (v *NullableSymbolPriceTickerV2Response) Set(val *SymbolPriceTickerV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolPriceTickerV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolPriceTickerV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolPriceTickerV2Response(val *SymbolPriceTickerV2Response) *NullableSymbolPriceTickerV2Response {
	return &NullableSymbolPriceTickerV2Response{value: val, isSet: true}
}

func (v NullableSymbolPriceTickerV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolPriceTickerV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
