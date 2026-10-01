/*
Spot REST API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ReferencePriceResponse - struct for ReferencePriceResponse
type ReferencePriceResponse struct {
	ReferencePriceResponse1 *ReferencePriceResponse1
	ReferencePriceResponse2 *ReferencePriceResponse2
}

// ReferencePriceResponse1AsReferencePriceResponse is a convenience function that returns ReferencePriceResponse1 wrapped in ReferencePriceResponse
func ReferencePriceResponse1AsReferencePriceResponse(v *ReferencePriceResponse1) ReferencePriceResponse {
	return ReferencePriceResponse{
		ReferencePriceResponse1: v,
	}
}

// ReferencePriceResponse2AsReferencePriceResponse is a convenience function that returns ReferencePriceResponse2 wrapped in ReferencePriceResponse
func ReferencePriceResponse2AsReferencePriceResponse(v *ReferencePriceResponse2) ReferencePriceResponse {
	return ReferencePriceResponse{
		ReferencePriceResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ReferencePriceResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ReferencePriceResponse1
	err = json.Unmarshal(data, &dst.ReferencePriceResponse1)
	if err == nil {
		jsonReferencePriceResponse1, _ := json.Marshal(dst.ReferencePriceResponse1)
		if string(jsonReferencePriceResponse1) == "{}" { // empty struct
			dst.ReferencePriceResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.ReferencePriceResponse1 = nil
	}

	// try to unmarshal data into ReferencePriceResponse2
	err = json.Unmarshal(data, &dst.ReferencePriceResponse2)
	if err == nil {
		jsonReferencePriceResponse2, _ := json.Marshal(dst.ReferencePriceResponse2)
		if string(jsonReferencePriceResponse2) == "{}" { // empty struct
			dst.ReferencePriceResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.ReferencePriceResponse2 = nil
	}

	if match > 1 { // more than 1 match
		kept := 0
		if dst.ReferencePriceResponse1 != nil {
			kept++
			if kept > 1 {
				dst.ReferencePriceResponse1 = nil
			}
		}
		if dst.ReferencePriceResponse2 != nil {
			kept++
			if kept > 1 {
				dst.ReferencePriceResponse2 = nil
			}
		}

		return nil
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ReferencePriceResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ReferencePriceResponse) MarshalJSON() ([]byte, error) {
	if src.ReferencePriceResponse1 != nil {
		return json.Marshal(&src.ReferencePriceResponse1)
	}

	if src.ReferencePriceResponse2 != nil {
		return json.Marshal(&src.ReferencePriceResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ReferencePriceResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ReferencePriceResponse1 != nil {
		return obj.ReferencePriceResponse1
	}

	if obj.ReferencePriceResponse2 != nil {
		return obj.ReferencePriceResponse2
	}

	// all schemas are nil
	return nil
}

type NullableReferencePriceResponse struct {
	value *ReferencePriceResponse
	isSet bool
}

func (v NullableReferencePriceResponse) Get() *ReferencePriceResponse {
	return v.value
}

func (v *NullableReferencePriceResponse) Set(val *ReferencePriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReferencePriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReferencePriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferencePriceResponse(val *ReferencePriceResponse) *NullableReferencePriceResponse {
	return &NullableReferencePriceResponse{value: val, isSet: true}
}

func (v NullableReferencePriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferencePriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
