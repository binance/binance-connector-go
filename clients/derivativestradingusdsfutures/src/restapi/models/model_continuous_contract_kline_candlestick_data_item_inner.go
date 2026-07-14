/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ContinuousContractKlineCandlestickDataItemInner - struct for ContinuousContractKlineCandlestickDataItemInner
type ContinuousContractKlineCandlestickDataItemInner struct {
	Int64  *int64
	String *string
}

// int64AsContinuousContractKlineCandlestickDataItemInner is a convenience function that returns int64 wrapped in ContinuousContractKlineCandlestickDataItemInner
func Int64AsContinuousContractKlineCandlestickDataItemInner(v *int64) ContinuousContractKlineCandlestickDataItemInner {
	return ContinuousContractKlineCandlestickDataItemInner{
		Int64: v,
	}
}

// stringAsContinuousContractKlineCandlestickDataItemInner is a convenience function that returns string wrapped in ContinuousContractKlineCandlestickDataItemInner
func StringAsContinuousContractKlineCandlestickDataItemInner(v *string) ContinuousContractKlineCandlestickDataItemInner {
	return ContinuousContractKlineCandlestickDataItemInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContinuousContractKlineCandlestickDataItemInner) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(ContinuousContractKlineCandlestickDataItemInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ContinuousContractKlineCandlestickDataItemInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContinuousContractKlineCandlestickDataItemInner) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContinuousContractKlineCandlestickDataItemInner) GetActualInstance() interface{} {
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

type NullableContinuousContractKlineCandlestickDataItemInner struct {
	value *ContinuousContractKlineCandlestickDataItemInner
	isSet bool
}

func (v NullableContinuousContractKlineCandlestickDataItemInner) Get() *ContinuousContractKlineCandlestickDataItemInner {
	return v.value
}

func (v *NullableContinuousContractKlineCandlestickDataItemInner) Set(val *ContinuousContractKlineCandlestickDataItemInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousContractKlineCandlestickDataItemInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousContractKlineCandlestickDataItemInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousContractKlineCandlestickDataItemInner(val *ContinuousContractKlineCandlestickDataItemInner) *NullableContinuousContractKlineCandlestickDataItemInner {
	return &NullableContinuousContractKlineCandlestickDataItemInner{value: val, isSet: true}
}

func (v NullableContinuousContractKlineCandlestickDataItemInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousContractKlineCandlestickDataItemInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
