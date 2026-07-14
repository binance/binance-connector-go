/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarkPriceKlineCandlestickDataItemInner - struct for MarkPriceKlineCandlestickDataItemInner
type MarkPriceKlineCandlestickDataItemInner struct {
	Int64  *int64
	String *string
}

// int64AsMarkPriceKlineCandlestickDataItemInner is a convenience function that returns int64 wrapped in MarkPriceKlineCandlestickDataItemInner
func Int64AsMarkPriceKlineCandlestickDataItemInner(v *int64) MarkPriceKlineCandlestickDataItemInner {
	return MarkPriceKlineCandlestickDataItemInner{
		Int64: v,
	}
}

// stringAsMarkPriceKlineCandlestickDataItemInner is a convenience function that returns string wrapped in MarkPriceKlineCandlestickDataItemInner
func StringAsMarkPriceKlineCandlestickDataItemInner(v *string) MarkPriceKlineCandlestickDataItemInner {
	return MarkPriceKlineCandlestickDataItemInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MarkPriceKlineCandlestickDataItemInner) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(MarkPriceKlineCandlestickDataItemInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MarkPriceKlineCandlestickDataItemInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MarkPriceKlineCandlestickDataItemInner) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MarkPriceKlineCandlestickDataItemInner) GetActualInstance() interface{} {
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

type NullableMarkPriceKlineCandlestickDataItemInner struct {
	value *MarkPriceKlineCandlestickDataItemInner
	isSet bool
}

func (v NullableMarkPriceKlineCandlestickDataItemInner) Get() *MarkPriceKlineCandlestickDataItemInner {
	return v.value
}

func (v *NullableMarkPriceKlineCandlestickDataItemInner) Set(val *MarkPriceKlineCandlestickDataItemInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkPriceKlineCandlestickDataItemInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkPriceKlineCandlestickDataItemInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkPriceKlineCandlestickDataItemInner(val *MarkPriceKlineCandlestickDataItemInner) *NullableMarkPriceKlineCandlestickDataItemInner {
	return &NullableMarkPriceKlineCandlestickDataItemInner{value: val, isSet: true}
}

func (v NullableMarkPriceKlineCandlestickDataItemInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkPriceKlineCandlestickDataItemInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
