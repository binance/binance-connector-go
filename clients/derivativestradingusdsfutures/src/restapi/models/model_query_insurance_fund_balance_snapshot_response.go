/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// QueryInsuranceFundBalanceSnapshotResponse - struct for QueryInsuranceFundBalanceSnapshotResponse
type QueryInsuranceFundBalanceSnapshotResponse struct {
	QueryInsuranceFundBalanceSnapshotResponse1 *QueryInsuranceFundBalanceSnapshotResponse1
	QueryInsuranceFundBalanceSnapshotResponse2 *QueryInsuranceFundBalanceSnapshotResponse2
}

// QueryInsuranceFundBalanceSnapshotResponse1AsQueryInsuranceFundBalanceSnapshotResponse is a convenience function that returns QueryInsuranceFundBalanceSnapshotResponse1 wrapped in QueryInsuranceFundBalanceSnapshotResponse
func QueryInsuranceFundBalanceSnapshotResponse1AsQueryInsuranceFundBalanceSnapshotResponse(v *QueryInsuranceFundBalanceSnapshotResponse1) QueryInsuranceFundBalanceSnapshotResponse {
	return QueryInsuranceFundBalanceSnapshotResponse{
		QueryInsuranceFundBalanceSnapshotResponse1: v,
	}
}

// QueryInsuranceFundBalanceSnapshotResponse2AsQueryInsuranceFundBalanceSnapshotResponse is a convenience function that returns QueryInsuranceFundBalanceSnapshotResponse2 wrapped in QueryInsuranceFundBalanceSnapshotResponse
func QueryInsuranceFundBalanceSnapshotResponse2AsQueryInsuranceFundBalanceSnapshotResponse(v *QueryInsuranceFundBalanceSnapshotResponse2) QueryInsuranceFundBalanceSnapshotResponse {
	return QueryInsuranceFundBalanceSnapshotResponse{
		QueryInsuranceFundBalanceSnapshotResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *QueryInsuranceFundBalanceSnapshotResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into QueryInsuranceFundBalanceSnapshotResponse1
	err = json.Unmarshal(data, &dst.QueryInsuranceFundBalanceSnapshotResponse1)
	if err == nil {
		jsonQueryInsuranceFundBalanceSnapshotResponse1, _ := json.Marshal(dst.QueryInsuranceFundBalanceSnapshotResponse1)
		if string(jsonQueryInsuranceFundBalanceSnapshotResponse1) == "{}" { // empty struct
			dst.QueryInsuranceFundBalanceSnapshotResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.QueryInsuranceFundBalanceSnapshotResponse1 = nil
	}

	// try to unmarshal data into QueryInsuranceFundBalanceSnapshotResponse2
	err = json.Unmarshal(data, &dst.QueryInsuranceFundBalanceSnapshotResponse2)
	if err == nil {
		jsonQueryInsuranceFundBalanceSnapshotResponse2, _ := json.Marshal(dst.QueryInsuranceFundBalanceSnapshotResponse2)
		if string(jsonQueryInsuranceFundBalanceSnapshotResponse2) == "{}" { // empty struct
			dst.QueryInsuranceFundBalanceSnapshotResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.QueryInsuranceFundBalanceSnapshotResponse2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.QueryInsuranceFundBalanceSnapshotResponse1 = nil
		dst.QueryInsuranceFundBalanceSnapshotResponse2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(QueryInsuranceFundBalanceSnapshotResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(QueryInsuranceFundBalanceSnapshotResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src QueryInsuranceFundBalanceSnapshotResponse) MarshalJSON() ([]byte, error) {
	if src.QueryInsuranceFundBalanceSnapshotResponse1 != nil {
		return json.Marshal(&src.QueryInsuranceFundBalanceSnapshotResponse1)
	}

	if src.QueryInsuranceFundBalanceSnapshotResponse2 != nil {
		return json.Marshal(&src.QueryInsuranceFundBalanceSnapshotResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *QueryInsuranceFundBalanceSnapshotResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.QueryInsuranceFundBalanceSnapshotResponse1 != nil {
		return obj.QueryInsuranceFundBalanceSnapshotResponse1
	}

	if obj.QueryInsuranceFundBalanceSnapshotResponse2 != nil {
		return obj.QueryInsuranceFundBalanceSnapshotResponse2
	}

	// all schemas are nil
	return nil
}

type NullableQueryInsuranceFundBalanceSnapshotResponse struct {
	value *QueryInsuranceFundBalanceSnapshotResponse
	isSet bool
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse) Get() *QueryInsuranceFundBalanceSnapshotResponse {
	return v.value
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse) Set(val *QueryInsuranceFundBalanceSnapshotResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryInsuranceFundBalanceSnapshotResponse(val *QueryInsuranceFundBalanceSnapshotResponse) *NullableQueryInsuranceFundBalanceSnapshotResponse {
	return &NullableQueryInsuranceFundBalanceSnapshotResponse{value: val, isSet: true}
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
