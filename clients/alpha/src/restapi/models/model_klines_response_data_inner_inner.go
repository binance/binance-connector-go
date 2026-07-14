/*
Alpha Trading REST API

APIs for Binance Alpha Trading.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// KlinesResponseDataInnerInner - struct for KlinesResponseDataInnerInner
type KlinesResponseDataInnerInner struct {
	Int64  *int64
	String *string
}

// int64AsKlinesResponseDataInnerInner is a convenience function that returns int64 wrapped in KlinesResponseDataInnerInner
func Int64AsKlinesResponseDataInnerInner(v *int64) KlinesResponseDataInnerInner {
	return KlinesResponseDataInnerInner{
		Int64: v,
	}
}

// stringAsKlinesResponseDataInnerInner is a convenience function that returns string wrapped in KlinesResponseDataInnerInner
func StringAsKlinesResponseDataInnerInner(v *string) KlinesResponseDataInnerInner {
	return KlinesResponseDataInnerInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *KlinesResponseDataInnerInner) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(KlinesResponseDataInnerInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(KlinesResponseDataInnerInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src KlinesResponseDataInnerInner) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *KlinesResponseDataInnerInner) GetActualInstance() interface{} {
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

type NullableKlinesResponseDataInnerInner struct {
	value *KlinesResponseDataInnerInner
	isSet bool
}

func (v NullableKlinesResponseDataInnerInner) Get() *KlinesResponseDataInnerInner {
	return v.value
}

func (v *NullableKlinesResponseDataInnerInner) Set(val *KlinesResponseDataInnerInner) {
	v.value = val
	v.isSet = true
}

func (v NullableKlinesResponseDataInnerInner) IsSet() bool {
	return v.isSet
}

func (v *NullableKlinesResponseDataInnerInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlinesResponseDataInnerInner(val *KlinesResponseDataInnerInner) *NullableKlinesResponseDataInnerInner {
	return &NullableKlinesResponseDataInnerInner{value: val, isSet: true}
}

func (v NullableKlinesResponseDataInnerInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlinesResponseDataInnerInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
