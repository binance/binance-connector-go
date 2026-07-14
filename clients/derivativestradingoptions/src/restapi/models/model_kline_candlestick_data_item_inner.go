/*
Options REST API

Access market data, manage accounts, and trade Binance Options.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// KlineCandlestickDataItemInner - struct for KlineCandlestickDataItemInner
type KlineCandlestickDataItemInner struct {
	Int64  *int64
	String *string
}

// int64AsKlineCandlestickDataItemInner is a convenience function that returns int64 wrapped in KlineCandlestickDataItemInner
func Int64AsKlineCandlestickDataItemInner(v *int64) KlineCandlestickDataItemInner {
	return KlineCandlestickDataItemInner{
		Int64: v,
	}
}

// stringAsKlineCandlestickDataItemInner is a convenience function that returns string wrapped in KlineCandlestickDataItemInner
func StringAsKlineCandlestickDataItemInner(v *string) KlineCandlestickDataItemInner {
	return KlineCandlestickDataItemInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *KlineCandlestickDataItemInner) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(KlineCandlestickDataItemInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(KlineCandlestickDataItemInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src KlineCandlestickDataItemInner) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *KlineCandlestickDataItemInner) GetActualInstance() interface{} {
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

type NullableKlineCandlestickDataItemInner struct {
	value *KlineCandlestickDataItemInner
	isSet bool
}

func (v NullableKlineCandlestickDataItemInner) Get() *KlineCandlestickDataItemInner {
	return v.value
}

func (v *NullableKlineCandlestickDataItemInner) Set(val *KlineCandlestickDataItemInner) {
	v.value = val
	v.isSet = true
}

func (v NullableKlineCandlestickDataItemInner) IsSet() bool {
	return v.isSet
}

func (v *NullableKlineCandlestickDataItemInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlineCandlestickDataItemInner(val *KlineCandlestickDataItemInner) *NullableKlineCandlestickDataItemInner {
	return &NullableKlineCandlestickDataItemInner{value: val, isSet: true}
}

func (v NullableKlineCandlestickDataItemInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlineCandlestickDataItemInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
