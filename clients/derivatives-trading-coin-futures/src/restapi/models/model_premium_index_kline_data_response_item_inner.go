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

// PremiumIndexKlineDataResponseItemInner - struct for PremiumIndexKlineDataResponseItemInner
type PremiumIndexKlineDataResponseItemInner struct {
	Int64  *int64
	String *string
}

// int64AsPremiumIndexKlineDataResponseItemInner is a convenience function that returns int64 wrapped in PremiumIndexKlineDataResponseItemInner
func Int64AsPremiumIndexKlineDataResponseItemInner(v *int64) PremiumIndexKlineDataResponseItemInner {
	return PremiumIndexKlineDataResponseItemInner{
		Int64: v,
	}
}

// stringAsPremiumIndexKlineDataResponseItemInner is a convenience function that returns string wrapped in PremiumIndexKlineDataResponseItemInner
func StringAsPremiumIndexKlineDataResponseItemInner(v *string) PremiumIndexKlineDataResponseItemInner {
	return PremiumIndexKlineDataResponseItemInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PremiumIndexKlineDataResponseItemInner) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(PremiumIndexKlineDataResponseItemInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PremiumIndexKlineDataResponseItemInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PremiumIndexKlineDataResponseItemInner) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PremiumIndexKlineDataResponseItemInner) GetActualInstance() interface{} {
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

type NullablePremiumIndexKlineDataResponseItemInner struct {
	value *PremiumIndexKlineDataResponseItemInner
	isSet bool
}

func (v NullablePremiumIndexKlineDataResponseItemInner) Get() *PremiumIndexKlineDataResponseItemInner {
	return v.value
}

func (v *NullablePremiumIndexKlineDataResponseItemInner) Set(val *PremiumIndexKlineDataResponseItemInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePremiumIndexKlineDataResponseItemInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePremiumIndexKlineDataResponseItemInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePremiumIndexKlineDataResponseItemInner(val *PremiumIndexKlineDataResponseItemInner) *NullablePremiumIndexKlineDataResponseItemInner {
	return &NullablePremiumIndexKlineDataResponseItemInner{value: val, isSet: true}
}

func (v NullablePremiumIndexKlineDataResponseItemInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePremiumIndexKlineDataResponseItemInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
