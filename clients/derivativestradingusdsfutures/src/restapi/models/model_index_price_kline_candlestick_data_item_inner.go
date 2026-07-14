/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// IndexPriceKlineCandlestickDataItemInner - struct for IndexPriceKlineCandlestickDataItemInner
type IndexPriceKlineCandlestickDataItemInner struct {
	Int64  *int64
	String *string
}

// int64AsIndexPriceKlineCandlestickDataItemInner is a convenience function that returns int64 wrapped in IndexPriceKlineCandlestickDataItemInner
func Int64AsIndexPriceKlineCandlestickDataItemInner(v *int64) IndexPriceKlineCandlestickDataItemInner {
	return IndexPriceKlineCandlestickDataItemInner{
		Int64: v,
	}
}

// stringAsIndexPriceKlineCandlestickDataItemInner is a convenience function that returns string wrapped in IndexPriceKlineCandlestickDataItemInner
func StringAsIndexPriceKlineCandlestickDataItemInner(v *string) IndexPriceKlineCandlestickDataItemInner {
	return IndexPriceKlineCandlestickDataItemInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IndexPriceKlineCandlestickDataItemInner) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(IndexPriceKlineCandlestickDataItemInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IndexPriceKlineCandlestickDataItemInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IndexPriceKlineCandlestickDataItemInner) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IndexPriceKlineCandlestickDataItemInner) GetActualInstance() interface{} {
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

type NullableIndexPriceKlineCandlestickDataItemInner struct {
	value *IndexPriceKlineCandlestickDataItemInner
	isSet bool
}

func (v NullableIndexPriceKlineCandlestickDataItemInner) Get() *IndexPriceKlineCandlestickDataItemInner {
	return v.value
}

func (v *NullableIndexPriceKlineCandlestickDataItemInner) Set(val *IndexPriceKlineCandlestickDataItemInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexPriceKlineCandlestickDataItemInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexPriceKlineCandlestickDataItemInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexPriceKlineCandlestickDataItemInner(val *IndexPriceKlineCandlestickDataItemInner) *NullableIndexPriceKlineCandlestickDataItemInner {
	return &NullableIndexPriceKlineCandlestickDataItemInner{value: val, isSet: true}
}

func (v NullableIndexPriceKlineCandlestickDataItemInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexPriceKlineCandlestickDataItemInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
