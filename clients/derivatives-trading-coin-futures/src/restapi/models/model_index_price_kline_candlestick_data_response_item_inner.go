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

// IndexPriceKlineCandlestickDataResponseItemInner - struct for IndexPriceKlineCandlestickDataResponseItemInner
type IndexPriceKlineCandlestickDataResponseItemInner struct {
	Int64  *int64
	String *string
}

// int64AsIndexPriceKlineCandlestickDataResponseItemInner is a convenience function that returns int64 wrapped in IndexPriceKlineCandlestickDataResponseItemInner
func Int64AsIndexPriceKlineCandlestickDataResponseItemInner(v *int64) IndexPriceKlineCandlestickDataResponseItemInner {
	return IndexPriceKlineCandlestickDataResponseItemInner{
		Int64: v,
	}
}

// stringAsIndexPriceKlineCandlestickDataResponseItemInner is a convenience function that returns string wrapped in IndexPriceKlineCandlestickDataResponseItemInner
func StringAsIndexPriceKlineCandlestickDataResponseItemInner(v *string) IndexPriceKlineCandlestickDataResponseItemInner {
	return IndexPriceKlineCandlestickDataResponseItemInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IndexPriceKlineCandlestickDataResponseItemInner) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(IndexPriceKlineCandlestickDataResponseItemInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IndexPriceKlineCandlestickDataResponseItemInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IndexPriceKlineCandlestickDataResponseItemInner) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IndexPriceKlineCandlestickDataResponseItemInner) GetActualInstance() interface{} {
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

type NullableIndexPriceKlineCandlestickDataResponseItemInner struct {
	value *IndexPriceKlineCandlestickDataResponseItemInner
	isSet bool
}

func (v NullableIndexPriceKlineCandlestickDataResponseItemInner) Get() *IndexPriceKlineCandlestickDataResponseItemInner {
	return v.value
}

func (v *NullableIndexPriceKlineCandlestickDataResponseItemInner) Set(val *IndexPriceKlineCandlestickDataResponseItemInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexPriceKlineCandlestickDataResponseItemInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexPriceKlineCandlestickDataResponseItemInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexPriceKlineCandlestickDataResponseItemInner(val *IndexPriceKlineCandlestickDataResponseItemInner) *NullableIndexPriceKlineCandlestickDataResponseItemInner {
	return &NullableIndexPriceKlineCandlestickDataResponseItemInner{value: val, isSet: true}
}

func (v NullableIndexPriceKlineCandlestickDataResponseItemInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexPriceKlineCandlestickDataResponseItemInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
