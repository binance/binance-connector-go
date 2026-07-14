/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// KlinesResponseResultInnerInner - struct for KlinesResponseResultInnerInner
type KlinesResponseResultInnerInner struct {
	Int64  *int64
	String *string
}

// int64AsKlinesResponseResultInnerInner is a convenience function that returns int64 wrapped in KlinesResponseResultInnerInner
func Int64AsKlinesResponseResultInnerInner(v *int64) KlinesResponseResultInnerInner {
	return KlinesResponseResultInnerInner{
		Int64: v,
	}
}

// stringAsKlinesResponseResultInnerInner is a convenience function that returns string wrapped in KlinesResponseResultInnerInner
func StringAsKlinesResponseResultInnerInner(v *string) KlinesResponseResultInnerInner {
	return KlinesResponseResultInnerInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *KlinesResponseResultInnerInner) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(KlinesResponseResultInnerInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(KlinesResponseResultInnerInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src KlinesResponseResultInnerInner) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *KlinesResponseResultInnerInner) GetActualInstance() interface{} {
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

type NullableKlinesResponseResultInnerInner struct {
	value *KlinesResponseResultInnerInner
	isSet bool
}

func (v NullableKlinesResponseResultInnerInner) Get() *KlinesResponseResultInnerInner {
	return v.value
}

func (v *NullableKlinesResponseResultInnerInner) Set(val *KlinesResponseResultInnerInner) {
	v.value = val
	v.isSet = true
}

func (v NullableKlinesResponseResultInnerInner) IsSet() bool {
	return v.isSet
}

func (v *NullableKlinesResponseResultInnerInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlinesResponseResultInnerInner(val *KlinesResponseResultInnerInner) *NullableKlinesResponseResultInnerInner {
	return &NullableKlinesResponseResultInnerInner{value: val, isSet: true}
}

func (v NullableKlinesResponseResultInnerInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlinesResponseResultInnerInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
