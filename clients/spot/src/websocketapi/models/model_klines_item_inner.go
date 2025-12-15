/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// KlinesItemInner - struct for KlinesItemInner
type KlinesItemInner struct {
	Int64  *int64
	String *string
}

// int64AsKlinesItemInner is a convenience function that returns int64 wrapped in KlinesItemInner
func Int64AsKlinesItemInner(v *int64) KlinesItemInner {
	return KlinesItemInner{
		Int64: v,
	}
}

// stringAsKlinesItemInner is a convenience function that returns string wrapped in KlinesItemInner
func StringAsKlinesItemInner(v *string) KlinesItemInner {
	return KlinesItemInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *KlinesItemInner) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(KlinesItemInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(KlinesItemInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src KlinesItemInner) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *KlinesItemInner) GetActualInstance() interface{} {
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

type NullableKlinesItemInner struct {
	value *KlinesItemInner
	isSet bool
}

func (v NullableKlinesItemInner) Get() *KlinesItemInner {
	return v.value
}

func (v *NullableKlinesItemInner) Set(val *KlinesItemInner) {
	v.value = val
	v.isSet = true
}

func (v NullableKlinesItemInner) IsSet() bool {
	return v.isSet
}

func (v *NullableKlinesItemInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlinesItemInner(val *KlinesItemInner) *NullableKlinesItemInner {
	return &NullableKlinesItemInner{value: val, isSet: true}
}

func (v NullableKlinesItemInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlinesItemInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
