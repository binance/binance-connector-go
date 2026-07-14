/*
Futures (COIN-M) REST API

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PremiumIndexKlineDataItemInner - struct for PremiumIndexKlineDataItemInner
type PremiumIndexKlineDataItemInner struct {
	Int64  *int64
	String *string
}

// int64AsPremiumIndexKlineDataItemInner is a convenience function that returns int64 wrapped in PremiumIndexKlineDataItemInner
func Int64AsPremiumIndexKlineDataItemInner(v *int64) PremiumIndexKlineDataItemInner {
	return PremiumIndexKlineDataItemInner{
		Int64: v,
	}
}

// stringAsPremiumIndexKlineDataItemInner is a convenience function that returns string wrapped in PremiumIndexKlineDataItemInner
func StringAsPremiumIndexKlineDataItemInner(v *string) PremiumIndexKlineDataItemInner {
	return PremiumIndexKlineDataItemInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PremiumIndexKlineDataItemInner) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(PremiumIndexKlineDataItemInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PremiumIndexKlineDataItemInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PremiumIndexKlineDataItemInner) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PremiumIndexKlineDataItemInner) GetActualInstance() interface{} {
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

type NullablePremiumIndexKlineDataItemInner struct {
	value *PremiumIndexKlineDataItemInner
	isSet bool
}

func (v NullablePremiumIndexKlineDataItemInner) Get() *PremiumIndexKlineDataItemInner {
	return v.value
}

func (v *NullablePremiumIndexKlineDataItemInner) Set(val *PremiumIndexKlineDataItemInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePremiumIndexKlineDataItemInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePremiumIndexKlineDataItemInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePremiumIndexKlineDataItemInner(val *PremiumIndexKlineDataItemInner) *NullablePremiumIndexKlineDataItemInner {
	return &NullablePremiumIndexKlineDataItemInner{value: val, isSet: true}
}

func (v NullablePremiumIndexKlineDataItemInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePremiumIndexKlineDataItemInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
