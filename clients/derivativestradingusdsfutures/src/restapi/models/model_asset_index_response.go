/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// AssetIndexResponse - struct for AssetIndexResponse
type AssetIndexResponse struct {
	AssetIndexResponse1 *AssetIndexResponse1
	AssetIndexResponse2 *AssetIndexResponse2
}

// AssetIndexResponse1AsAssetIndexResponse is a convenience function that returns AssetIndexResponse1 wrapped in AssetIndexResponse
func AssetIndexResponse1AsAssetIndexResponse(v *AssetIndexResponse1) AssetIndexResponse {
	return AssetIndexResponse{
		AssetIndexResponse1: v,
	}
}

// AssetIndexResponse2AsAssetIndexResponse is a convenience function that returns AssetIndexResponse2 wrapped in AssetIndexResponse
func AssetIndexResponse2AsAssetIndexResponse(v *AssetIndexResponse2) AssetIndexResponse {
	return AssetIndexResponse{
		AssetIndexResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AssetIndexResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AssetIndexResponse1
	err = json.Unmarshal(data, &dst.AssetIndexResponse1)
	if err == nil {
		jsonAssetIndexResponse1, _ := json.Marshal(dst.AssetIndexResponse1)
		if string(jsonAssetIndexResponse1) == "{}" { // empty struct
			dst.AssetIndexResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.AssetIndexResponse1 = nil
	}

	// try to unmarshal data into AssetIndexResponse2
	err = json.Unmarshal(data, &dst.AssetIndexResponse2)
	if err == nil {
		jsonAssetIndexResponse2, _ := json.Marshal(dst.AssetIndexResponse2)
		if string(jsonAssetIndexResponse2) == "{}" { // empty struct
			dst.AssetIndexResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.AssetIndexResponse2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AssetIndexResponse1 = nil
		dst.AssetIndexResponse2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AssetIndexResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AssetIndexResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AssetIndexResponse) MarshalJSON() ([]byte, error) {
	if src.AssetIndexResponse1 != nil {
		return json.Marshal(&src.AssetIndexResponse1)
	}

	if src.AssetIndexResponse2 != nil {
		return json.Marshal(&src.AssetIndexResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AssetIndexResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AssetIndexResponse1 != nil {
		return obj.AssetIndexResponse1
	}

	if obj.AssetIndexResponse2 != nil {
		return obj.AssetIndexResponse2
	}

	// all schemas are nil
	return nil
}

type NullableAssetIndexResponse struct {
	value *AssetIndexResponse
	isSet bool
}

func (v NullableAssetIndexResponse) Get() *AssetIndexResponse {
	return v.value
}

func (v *NullableAssetIndexResponse) Set(val *AssetIndexResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetIndexResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetIndexResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetIndexResponse(val *AssetIndexResponse) *NullableAssetIndexResponse {
	return &NullableAssetIndexResponse{value: val, isSet: true}
}

func (v NullableAssetIndexResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetIndexResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
