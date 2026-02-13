/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// AssetFilters - struct for AssetFilters
type AssetFilters struct {
	MaxAssetFilter *MaxAssetFilter
}

// MaxAssetFilterAsAssetFilters is a convenience function that returns MaxAssetFilter wrapped in AssetFilters
func MaxAssetFilterAsAssetFilters(v *MaxAssetFilter) AssetFilters {
	return AssetFilters{
		MaxAssetFilter: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AssetFilters) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = common.NewStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'MAX_ASSET'
	if jsonDict["filterType"] == "MAX_ASSET" {
		// try to unmarshal JSON data into MaxAssetFilter
		err = json.Unmarshal(data, &dst.MaxAssetFilter)
		if err == nil {
			return nil // data stored in dst.MaxAssetFilter, return on the first match
		} else {
			dst.MaxAssetFilter = nil
			return fmt.Errorf("failed to unmarshal AssetFilters as MaxAssetFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MaxAssetFilter'
	if jsonDict["filterType"] == "MaxAssetFilter" {
		// try to unmarshal JSON data into MaxAssetFilter
		err = json.Unmarshal(data, &dst.MaxAssetFilter)
		if err == nil {
			return nil // data stored in dst.MaxAssetFilter, return on the first match
		} else {
			dst.MaxAssetFilter = nil
			return fmt.Errorf("failed to unmarshal AssetFilters as MaxAssetFilter: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AssetFilters) MarshalJSON() ([]byte, error) {
	if src.MaxAssetFilter != nil {
		return json.Marshal(&src.MaxAssetFilter)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AssetFilters) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MaxAssetFilter != nil {
		return obj.MaxAssetFilter
	}

	// all schemas are nil
	return nil
}

type NullableAssetFilters struct {
	value *AssetFilters
	isSet bool
}

func (v NullableAssetFilters) Get() *AssetFilters {
	return v.value
}

func (v *NullableAssetFilters) Set(val *AssetFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetFilters(val *AssetFilters) *NullableAssetFilters {
	return &NullableAssetFilters{value: val, isSet: true}
}

func (v NullableAssetFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
