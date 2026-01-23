/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API

*/

package models

import (
	"encoding/json"
	"fmt"
)

// KlinesResponseDataItemInner - struct for KlinesResponseDataItemInner
type KlinesResponseDataItemInner struct {
	Int64  *int64
	String *string
}

// int64AsKlinesResponseDataItemInner is a convenience function that returns int64 wrapped in KlinesResponseDataItemInner
func Int64AsKlinesResponseDataItemInner(v *int64) KlinesResponseDataItemInner {
	return KlinesResponseDataItemInner{
		Int64: v,
	}
}

// stringAsKlinesResponseDataItemInner is a convenience function that returns string wrapped in KlinesResponseDataItemInner
func StringAsKlinesResponseDataItemInner(v *string) KlinesResponseDataItemInner {
	return KlinesResponseDataItemInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *KlinesResponseDataItemInner) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(KlinesResponseDataItemInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(KlinesResponseDataItemInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src KlinesResponseDataItemInner) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *KlinesResponseDataItemInner) GetActualInstance() interface{} {
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

type NullableKlinesResponseDataItemInner struct {
	value *KlinesResponseDataItemInner
	isSet bool
}

func (v NullableKlinesResponseDataItemInner) Get() *KlinesResponseDataItemInner {
	return v.value
}

func (v *NullableKlinesResponseDataItemInner) Set(val *KlinesResponseDataItemInner) {
	v.value = val
	v.isSet = true
}

func (v NullableKlinesResponseDataItemInner) IsSet() bool {
	return v.isSet
}

func (v *NullableKlinesResponseDataItemInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlinesResponseDataItemInner(val *KlinesResponseDataItemInner) *NullableKlinesResponseDataItemInner {
	return &NullableKlinesResponseDataItemInner{value: val, isSet: true}
}

func (v NullableKlinesResponseDataItemInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlinesResponseDataItemInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
