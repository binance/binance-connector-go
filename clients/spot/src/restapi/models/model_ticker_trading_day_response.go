/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// TickerTradingDayResponse - struct for TickerTradingDayResponse
type TickerTradingDayResponse struct {
	TickerTradingDayResponse1 *TickerTradingDayResponse1
	TickerTradingDayResponse2 *TickerTradingDayResponse2
}

// TickerTradingDayResponse1AsTickerTradingDayResponse is a convenience function that returns TickerTradingDayResponse1 wrapped in TickerTradingDayResponse
func TickerTradingDayResponse1AsTickerTradingDayResponse(v *TickerTradingDayResponse1) TickerTradingDayResponse {
	return TickerTradingDayResponse{
		TickerTradingDayResponse1: v,
	}
}

// TickerTradingDayResponse2AsTickerTradingDayResponse is a convenience function that returns TickerTradingDayResponse2 wrapped in TickerTradingDayResponse
func TickerTradingDayResponse2AsTickerTradingDayResponse(v *TickerTradingDayResponse2) TickerTradingDayResponse {
	return TickerTradingDayResponse{
		TickerTradingDayResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TickerTradingDayResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TickerTradingDayResponse1
	err = json.Unmarshal(data, &dst.TickerTradingDayResponse1)
	if err == nil {
		jsonTickerTradingDayResponse1, _ := json.Marshal(dst.TickerTradingDayResponse1)
		if string(jsonTickerTradingDayResponse1) == "{}" { // empty struct
			dst.TickerTradingDayResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.TickerTradingDayResponse1 = nil
	}

	// try to unmarshal data into TickerTradingDayResponse2
	err = json.Unmarshal(data, &dst.TickerTradingDayResponse2)
	if err == nil {
		jsonTickerTradingDayResponse2, _ := json.Marshal(dst.TickerTradingDayResponse2)
		if string(jsonTickerTradingDayResponse2) == "{}" { // empty struct
			dst.TickerTradingDayResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.TickerTradingDayResponse2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.TickerTradingDayResponse1 = nil
		dst.TickerTradingDayResponse2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TickerTradingDayResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TickerTradingDayResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TickerTradingDayResponse) MarshalJSON() ([]byte, error) {
	if src.TickerTradingDayResponse1 != nil {
		return json.Marshal(&src.TickerTradingDayResponse1)
	}

	if src.TickerTradingDayResponse2 != nil {
		return json.Marshal(&src.TickerTradingDayResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TickerTradingDayResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.TickerTradingDayResponse1 != nil {
		return obj.TickerTradingDayResponse1
	}

	if obj.TickerTradingDayResponse2 != nil {
		return obj.TickerTradingDayResponse2
	}

	// all schemas are nil
	return nil
}

type NullableTickerTradingDayResponse struct {
	value *TickerTradingDayResponse
	isSet bool
}

func (v NullableTickerTradingDayResponse) Get() *TickerTradingDayResponse {
	return v.value
}

func (v *NullableTickerTradingDayResponse) Set(val *TickerTradingDayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerTradingDayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerTradingDayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerTradingDayResponse(val *TickerTradingDayResponse) *NullableTickerTradingDayResponse {
	return &NullableTickerTradingDayResponse{value: val, isSet: true}
}

func (v NullableTickerTradingDayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerTradingDayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
