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

// Ticker24hrResponse - struct for Ticker24hrResponse
type Ticker24hrResponse struct {
	Ticker24hrResponse1 *Ticker24hrResponse1
	Ticker24hrResponse2 *Ticker24hrResponse2
}

// Ticker24hrResponse1AsTicker24hrResponse is a convenience function that returns Ticker24hrResponse1 wrapped in Ticker24hrResponse
func Ticker24hrResponse1AsTicker24hrResponse(v *Ticker24hrResponse1) Ticker24hrResponse {
	return Ticker24hrResponse{
		Ticker24hrResponse1: v,
	}
}

// Ticker24hrResponse2AsTicker24hrResponse is a convenience function that returns Ticker24hrResponse2 wrapped in Ticker24hrResponse
func Ticker24hrResponse2AsTicker24hrResponse(v *Ticker24hrResponse2) Ticker24hrResponse {
	return Ticker24hrResponse{
		Ticker24hrResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Ticker24hrResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Ticker24hrResponse1
	err = json.Unmarshal(data, &dst.Ticker24hrResponse1)
	if err == nil {
		jsonTicker24hrResponse1, _ := json.Marshal(dst.Ticker24hrResponse1)
		if string(jsonTicker24hrResponse1) == "{}" { // empty struct
			dst.Ticker24hrResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.Ticker24hrResponse1 = nil
	}

	// try to unmarshal data into Ticker24hrResponse2
	err = json.Unmarshal(data, &dst.Ticker24hrResponse2)
	if err == nil {
		jsonTicker24hrResponse2, _ := json.Marshal(dst.Ticker24hrResponse2)
		if string(jsonTicker24hrResponse2) == "{}" { // empty struct
			dst.Ticker24hrResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.Ticker24hrResponse2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Ticker24hrResponse1 = nil
		dst.Ticker24hrResponse2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Ticker24hrResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Ticker24hrResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Ticker24hrResponse) MarshalJSON() ([]byte, error) {
	if src.Ticker24hrResponse1 != nil {
		return json.Marshal(&src.Ticker24hrResponse1)
	}

	if src.Ticker24hrResponse2 != nil {
		return json.Marshal(&src.Ticker24hrResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Ticker24hrResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Ticker24hrResponse1 != nil {
		return obj.Ticker24hrResponse1
	}

	if obj.Ticker24hrResponse2 != nil {
		return obj.Ticker24hrResponse2
	}

	// all schemas are nil
	return nil
}

type NullableTicker24hrResponse struct {
	value *Ticker24hrResponse
	isSet bool
}

func (v NullableTicker24hrResponse) Get() *Ticker24hrResponse {
	return v.value
}

func (v *NullableTicker24hrResponse) Set(val *Ticker24hrResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTicker24hrResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTicker24hrResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicker24hrResponse(val *Ticker24hrResponse) *NullableTicker24hrResponse {
	return &NullableTicker24hrResponse{value: val, isSet: true}
}

func (v NullableTicker24hrResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicker24hrResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
