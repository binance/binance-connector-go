/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarkPriceKlineCandlestickDataResponseItemInner - struct for MarkPriceKlineCandlestickDataResponseItemInner
type MarkPriceKlineCandlestickDataResponseItemInner struct {
	Int64  *int64
	String *string
}

// int64AsMarkPriceKlineCandlestickDataResponseItemInner is a convenience function that returns int64 wrapped in MarkPriceKlineCandlestickDataResponseItemInner
func Int64AsMarkPriceKlineCandlestickDataResponseItemInner(v *int64) MarkPriceKlineCandlestickDataResponseItemInner {
	return MarkPriceKlineCandlestickDataResponseItemInner{
		Int64: v,
	}
}

// stringAsMarkPriceKlineCandlestickDataResponseItemInner is a convenience function that returns string wrapped in MarkPriceKlineCandlestickDataResponseItemInner
func StringAsMarkPriceKlineCandlestickDataResponseItemInner(v *string) MarkPriceKlineCandlestickDataResponseItemInner {
	return MarkPriceKlineCandlestickDataResponseItemInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MarkPriceKlineCandlestickDataResponseItemInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int64
	err = json.Unmarshal(data, &dst.Int64)
	if err == nil {
		jsonInt64, _ := json.Marshal(dst.Int64)
		if string(jsonInt64) == "{}" { // empty struct
			dst.Int64 = nil
		} else {
			match++
		}
	} else {
		dst.Int64 = nil
	}

	// try to unmarshal data into String
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int64 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MarkPriceKlineCandlestickDataResponseItemInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MarkPriceKlineCandlestickDataResponseItemInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MarkPriceKlineCandlestickDataResponseItemInner) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MarkPriceKlineCandlestickDataResponseItemInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Int64 != nil {
		return obj.Int64
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableMarkPriceKlineCandlestickDataResponseItemInner struct {
	value *MarkPriceKlineCandlestickDataResponseItemInner
	isSet bool
}

func (v NullableMarkPriceKlineCandlestickDataResponseItemInner) Get() *MarkPriceKlineCandlestickDataResponseItemInner {
	return v.value
}

func (v *NullableMarkPriceKlineCandlestickDataResponseItemInner) Set(val *MarkPriceKlineCandlestickDataResponseItemInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkPriceKlineCandlestickDataResponseItemInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkPriceKlineCandlestickDataResponseItemInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkPriceKlineCandlestickDataResponseItemInner(val *MarkPriceKlineCandlestickDataResponseItemInner) *NullableMarkPriceKlineCandlestickDataResponseItemInner {
	return &NullableMarkPriceKlineCandlestickDataResponseItemInner{value: val, isSet: true}
}

func (v NullableMarkPriceKlineCandlestickDataResponseItemInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkPriceKlineCandlestickDataResponseItemInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
