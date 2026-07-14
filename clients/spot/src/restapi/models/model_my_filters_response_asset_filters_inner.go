/*
Spot REST API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// MyFiltersResponseAssetFiltersInner - struct for MyFiltersResponseAssetFiltersInner
type MyFiltersResponseAssetFiltersInner struct {
	MaxAssetFilter *MaxAssetFilter
}

// MaxAssetFilterAsMyFiltersResponseAssetFiltersInner is a convenience function that returns MaxAssetFilter wrapped in MyFiltersResponseAssetFiltersInner
func MaxAssetFilterAsMyFiltersResponseAssetFiltersInner(v *MaxAssetFilter) MyFiltersResponseAssetFiltersInner {
	return MyFiltersResponseAssetFiltersInner{
		MaxAssetFilter: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MyFiltersResponseAssetFiltersInner) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal MyFiltersResponseAssetFiltersInner as MaxAssetFilter: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MyFiltersResponseAssetFiltersInner) MarshalJSON() ([]byte, error) {
	if src.MaxAssetFilter != nil {
		return json.Marshal(&src.MaxAssetFilter)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MyFiltersResponseAssetFiltersInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MaxAssetFilter != nil {
		return obj.MaxAssetFilter
	}

	// all schemas are nil
	return nil
}

type NullableMyFiltersResponseAssetFiltersInner struct {
	value *MyFiltersResponseAssetFiltersInner
	isSet bool
}

func (v NullableMyFiltersResponseAssetFiltersInner) Get() *MyFiltersResponseAssetFiltersInner {
	return v.value
}

func (v *NullableMyFiltersResponseAssetFiltersInner) Set(val *MyFiltersResponseAssetFiltersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMyFiltersResponseAssetFiltersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMyFiltersResponseAssetFiltersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyFiltersResponseAssetFiltersInner(val *MyFiltersResponseAssetFiltersInner) *NullableMyFiltersResponseAssetFiltersInner {
	return &NullableMyFiltersResponseAssetFiltersInner{value: val, isSet: true}
}

func (v NullableMyFiltersResponseAssetFiltersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyFiltersResponseAssetFiltersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
