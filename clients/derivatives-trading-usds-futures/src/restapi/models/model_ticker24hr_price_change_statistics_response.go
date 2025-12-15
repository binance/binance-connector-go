/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// Ticker24hrPriceChangeStatisticsResponse - struct for Ticker24hrPriceChangeStatisticsResponse
type Ticker24hrPriceChangeStatisticsResponse struct {
	Ticker24hrPriceChangeStatisticsResponse1 *Ticker24hrPriceChangeStatisticsResponse1
	Ticker24hrPriceChangeStatisticsResponse2 *Ticker24hrPriceChangeStatisticsResponse2
}

// Ticker24hrPriceChangeStatisticsResponse1AsTicker24hrPriceChangeStatisticsResponse is a convenience function that returns Ticker24hrPriceChangeStatisticsResponse1 wrapped in Ticker24hrPriceChangeStatisticsResponse
func Ticker24hrPriceChangeStatisticsResponse1AsTicker24hrPriceChangeStatisticsResponse(v *Ticker24hrPriceChangeStatisticsResponse1) Ticker24hrPriceChangeStatisticsResponse {
	return Ticker24hrPriceChangeStatisticsResponse{
		Ticker24hrPriceChangeStatisticsResponse1: v,
	}
}

// Ticker24hrPriceChangeStatisticsResponse2AsTicker24hrPriceChangeStatisticsResponse is a convenience function that returns Ticker24hrPriceChangeStatisticsResponse2 wrapped in Ticker24hrPriceChangeStatisticsResponse
func Ticker24hrPriceChangeStatisticsResponse2AsTicker24hrPriceChangeStatisticsResponse(v *Ticker24hrPriceChangeStatisticsResponse2) Ticker24hrPriceChangeStatisticsResponse {
	return Ticker24hrPriceChangeStatisticsResponse{
		Ticker24hrPriceChangeStatisticsResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Ticker24hrPriceChangeStatisticsResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Ticker24hrPriceChangeStatisticsResponse1
	err = json.Unmarshal(data, &dst.Ticker24hrPriceChangeStatisticsResponse1)
	if err == nil {
		jsonTicker24hrPriceChangeStatisticsResponse1, _ := json.Marshal(dst.Ticker24hrPriceChangeStatisticsResponse1)
		if string(jsonTicker24hrPriceChangeStatisticsResponse1) == "{}" { // empty struct
			dst.Ticker24hrPriceChangeStatisticsResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.Ticker24hrPriceChangeStatisticsResponse1 = nil
	}

	// try to unmarshal data into Ticker24hrPriceChangeStatisticsResponse2
	err = json.Unmarshal(data, &dst.Ticker24hrPriceChangeStatisticsResponse2)
	if err == nil {
		jsonTicker24hrPriceChangeStatisticsResponse2, _ := json.Marshal(dst.Ticker24hrPriceChangeStatisticsResponse2)
		if string(jsonTicker24hrPriceChangeStatisticsResponse2) == "{}" { // empty struct
			dst.Ticker24hrPriceChangeStatisticsResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.Ticker24hrPriceChangeStatisticsResponse2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Ticker24hrPriceChangeStatisticsResponse1 = nil
		dst.Ticker24hrPriceChangeStatisticsResponse2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Ticker24hrPriceChangeStatisticsResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Ticker24hrPriceChangeStatisticsResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Ticker24hrPriceChangeStatisticsResponse) MarshalJSON() ([]byte, error) {
	if src.Ticker24hrPriceChangeStatisticsResponse1 != nil {
		return json.Marshal(&src.Ticker24hrPriceChangeStatisticsResponse1)
	}

	if src.Ticker24hrPriceChangeStatisticsResponse2 != nil {
		return json.Marshal(&src.Ticker24hrPriceChangeStatisticsResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Ticker24hrPriceChangeStatisticsResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Ticker24hrPriceChangeStatisticsResponse1 != nil {
		return obj.Ticker24hrPriceChangeStatisticsResponse1
	}

	if obj.Ticker24hrPriceChangeStatisticsResponse2 != nil {
		return obj.Ticker24hrPriceChangeStatisticsResponse2
	}

	// all schemas are nil
	return nil
}

type NullableTicker24hrPriceChangeStatisticsResponse struct {
	value *Ticker24hrPriceChangeStatisticsResponse
	isSet bool
}

func (v NullableTicker24hrPriceChangeStatisticsResponse) Get() *Ticker24hrPriceChangeStatisticsResponse {
	return v.value
}

func (v *NullableTicker24hrPriceChangeStatisticsResponse) Set(val *Ticker24hrPriceChangeStatisticsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTicker24hrPriceChangeStatisticsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTicker24hrPriceChangeStatisticsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicker24hrPriceChangeStatisticsResponse(val *Ticker24hrPriceChangeStatisticsResponse) *NullableTicker24hrPriceChangeStatisticsResponse {
	return &NullableTicker24hrPriceChangeStatisticsResponse{value: val, isSet: true}
}

func (v NullableTicker24hrPriceChangeStatisticsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicker24hrPriceChangeStatisticsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
