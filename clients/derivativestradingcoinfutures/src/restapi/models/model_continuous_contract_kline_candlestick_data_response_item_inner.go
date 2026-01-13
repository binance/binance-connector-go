/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ContinuousContractKlineCandlestickDataResponseItemInner - struct for ContinuousContractKlineCandlestickDataResponseItemInner
type ContinuousContractKlineCandlestickDataResponseItemInner struct {
	Int64  *int64
	String *string
}

// int64AsContinuousContractKlineCandlestickDataResponseItemInner is a convenience function that returns int64 wrapped in ContinuousContractKlineCandlestickDataResponseItemInner
func Int64AsContinuousContractKlineCandlestickDataResponseItemInner(v *int64) ContinuousContractKlineCandlestickDataResponseItemInner {
	return ContinuousContractKlineCandlestickDataResponseItemInner{
		Int64: v,
	}
}

// stringAsContinuousContractKlineCandlestickDataResponseItemInner is a convenience function that returns string wrapped in ContinuousContractKlineCandlestickDataResponseItemInner
func StringAsContinuousContractKlineCandlestickDataResponseItemInner(v *string) ContinuousContractKlineCandlestickDataResponseItemInner {
	return ContinuousContractKlineCandlestickDataResponseItemInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContinuousContractKlineCandlestickDataResponseItemInner) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(ContinuousContractKlineCandlestickDataResponseItemInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ContinuousContractKlineCandlestickDataResponseItemInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContinuousContractKlineCandlestickDataResponseItemInner) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContinuousContractKlineCandlestickDataResponseItemInner) GetActualInstance() interface{} {
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

type NullableContinuousContractKlineCandlestickDataResponseItemInner struct {
	value *ContinuousContractKlineCandlestickDataResponseItemInner
	isSet bool
}

func (v NullableContinuousContractKlineCandlestickDataResponseItemInner) Get() *ContinuousContractKlineCandlestickDataResponseItemInner {
	return v.value
}

func (v *NullableContinuousContractKlineCandlestickDataResponseItemInner) Set(val *ContinuousContractKlineCandlestickDataResponseItemInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousContractKlineCandlestickDataResponseItemInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousContractKlineCandlestickDataResponseItemInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousContractKlineCandlestickDataResponseItemInner(val *ContinuousContractKlineCandlestickDataResponseItemInner) *NullableContinuousContractKlineCandlestickDataResponseItemInner {
	return &NullableContinuousContractKlineCandlestickDataResponseItemInner{value: val, isSet: true}
}

func (v NullableContinuousContractKlineCandlestickDataResponseItemInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousContractKlineCandlestickDataResponseItemInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
