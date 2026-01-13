/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// KlineCandlestickDataResponseItemInner - struct for KlineCandlestickDataResponseItemInner
type KlineCandlestickDataResponseItemInner struct {
	Int64  *int64
	String *string
}

// int64AsKlineCandlestickDataResponseItemInner is a convenience function that returns int64 wrapped in KlineCandlestickDataResponseItemInner
func Int64AsKlineCandlestickDataResponseItemInner(v *int64) KlineCandlestickDataResponseItemInner {
	return KlineCandlestickDataResponseItemInner{
		Int64: v,
	}
}

// stringAsKlineCandlestickDataResponseItemInner is a convenience function that returns string wrapped in KlineCandlestickDataResponseItemInner
func StringAsKlineCandlestickDataResponseItemInner(v *string) KlineCandlestickDataResponseItemInner {
	return KlineCandlestickDataResponseItemInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *KlineCandlestickDataResponseItemInner) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(KlineCandlestickDataResponseItemInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(KlineCandlestickDataResponseItemInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src KlineCandlestickDataResponseItemInner) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *KlineCandlestickDataResponseItemInner) GetActualInstance() interface{} {
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

type NullableKlineCandlestickDataResponseItemInner struct {
	value *KlineCandlestickDataResponseItemInner
	isSet bool
}

func (v NullableKlineCandlestickDataResponseItemInner) Get() *KlineCandlestickDataResponseItemInner {
	return v.value
}

func (v *NullableKlineCandlestickDataResponseItemInner) Set(val *KlineCandlestickDataResponseItemInner) {
	v.value = val
	v.isSet = true
}

func (v NullableKlineCandlestickDataResponseItemInner) IsSet() bool {
	return v.isSet
}

func (v *NullableKlineCandlestickDataResponseItemInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlineCandlestickDataResponseItemInner(val *KlineCandlestickDataResponseItemInner) *NullableKlineCandlestickDataResponseItemInner {
	return &NullableKlineCandlestickDataResponseItemInner{value: val, isSet: true}
}

func (v NullableKlineCandlestickDataResponseItemInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlineCandlestickDataResponseItemInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
