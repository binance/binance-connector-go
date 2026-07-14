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

// MyFiltersResponseSymbolFiltersInner - struct for MyFiltersResponseSymbolFiltersInner
type MyFiltersResponseSymbolFiltersInner struct {
	IcebergPartsFilter        *IcebergPartsFilter
	LotSizeFilter             *LotSizeFilter
	MarketLotSizeFilter       *MarketLotSizeFilter
	MaxNumAlgoOrdersFilter    *MaxNumAlgoOrdersFilter
	MaxNumIcebergOrdersFilter *MaxNumIcebergOrdersFilter
	MaxNumOrderAmendsFilter   *MaxNumOrderAmendsFilter
	MaxNumOrderListsFilter    *MaxNumOrderListsFilter
	MaxNumOrdersFilter        *MaxNumOrdersFilter
	MaxPositionFilter         *MaxPositionFilter
	MinNotionalFilter         *MinNotionalFilter
	NotionalFilter            *NotionalFilter
	PercentPriceBySideFilter  *PercentPriceBySideFilter
	PercentPriceFilter        *PercentPriceFilter
	PriceFilter               *PriceFilter
	TPlusSellFilter           *TPlusSellFilter
	TrailingDeltaFilter       *TrailingDeltaFilter
}

// IcebergPartsFilterAsMyFiltersResponseSymbolFiltersInner is a convenience function that returns IcebergPartsFilter wrapped in MyFiltersResponseSymbolFiltersInner
func IcebergPartsFilterAsMyFiltersResponseSymbolFiltersInner(v *IcebergPartsFilter) MyFiltersResponseSymbolFiltersInner {
	return MyFiltersResponseSymbolFiltersInner{
		IcebergPartsFilter: v,
	}
}

// LotSizeFilterAsMyFiltersResponseSymbolFiltersInner is a convenience function that returns LotSizeFilter wrapped in MyFiltersResponseSymbolFiltersInner
func LotSizeFilterAsMyFiltersResponseSymbolFiltersInner(v *LotSizeFilter) MyFiltersResponseSymbolFiltersInner {
	return MyFiltersResponseSymbolFiltersInner{
		LotSizeFilter: v,
	}
}

// MarketLotSizeFilterAsMyFiltersResponseSymbolFiltersInner is a convenience function that returns MarketLotSizeFilter wrapped in MyFiltersResponseSymbolFiltersInner
func MarketLotSizeFilterAsMyFiltersResponseSymbolFiltersInner(v *MarketLotSizeFilter) MyFiltersResponseSymbolFiltersInner {
	return MyFiltersResponseSymbolFiltersInner{
		MarketLotSizeFilter: v,
	}
}

// MaxNumAlgoOrdersFilterAsMyFiltersResponseSymbolFiltersInner is a convenience function that returns MaxNumAlgoOrdersFilter wrapped in MyFiltersResponseSymbolFiltersInner
func MaxNumAlgoOrdersFilterAsMyFiltersResponseSymbolFiltersInner(v *MaxNumAlgoOrdersFilter) MyFiltersResponseSymbolFiltersInner {
	return MyFiltersResponseSymbolFiltersInner{
		MaxNumAlgoOrdersFilter: v,
	}
}

// MaxNumIcebergOrdersFilterAsMyFiltersResponseSymbolFiltersInner is a convenience function that returns MaxNumIcebergOrdersFilter wrapped in MyFiltersResponseSymbolFiltersInner
func MaxNumIcebergOrdersFilterAsMyFiltersResponseSymbolFiltersInner(v *MaxNumIcebergOrdersFilter) MyFiltersResponseSymbolFiltersInner {
	return MyFiltersResponseSymbolFiltersInner{
		MaxNumIcebergOrdersFilter: v,
	}
}

// MaxNumOrderAmendsFilterAsMyFiltersResponseSymbolFiltersInner is a convenience function that returns MaxNumOrderAmendsFilter wrapped in MyFiltersResponseSymbolFiltersInner
func MaxNumOrderAmendsFilterAsMyFiltersResponseSymbolFiltersInner(v *MaxNumOrderAmendsFilter) MyFiltersResponseSymbolFiltersInner {
	return MyFiltersResponseSymbolFiltersInner{
		MaxNumOrderAmendsFilter: v,
	}
}

// MaxNumOrderListsFilterAsMyFiltersResponseSymbolFiltersInner is a convenience function that returns MaxNumOrderListsFilter wrapped in MyFiltersResponseSymbolFiltersInner
func MaxNumOrderListsFilterAsMyFiltersResponseSymbolFiltersInner(v *MaxNumOrderListsFilter) MyFiltersResponseSymbolFiltersInner {
	return MyFiltersResponseSymbolFiltersInner{
		MaxNumOrderListsFilter: v,
	}
}

// MaxNumOrdersFilterAsMyFiltersResponseSymbolFiltersInner is a convenience function that returns MaxNumOrdersFilter wrapped in MyFiltersResponseSymbolFiltersInner
func MaxNumOrdersFilterAsMyFiltersResponseSymbolFiltersInner(v *MaxNumOrdersFilter) MyFiltersResponseSymbolFiltersInner {
	return MyFiltersResponseSymbolFiltersInner{
		MaxNumOrdersFilter: v,
	}
}

// MaxPositionFilterAsMyFiltersResponseSymbolFiltersInner is a convenience function that returns MaxPositionFilter wrapped in MyFiltersResponseSymbolFiltersInner
func MaxPositionFilterAsMyFiltersResponseSymbolFiltersInner(v *MaxPositionFilter) MyFiltersResponseSymbolFiltersInner {
	return MyFiltersResponseSymbolFiltersInner{
		MaxPositionFilter: v,
	}
}

// MinNotionalFilterAsMyFiltersResponseSymbolFiltersInner is a convenience function that returns MinNotionalFilter wrapped in MyFiltersResponseSymbolFiltersInner
func MinNotionalFilterAsMyFiltersResponseSymbolFiltersInner(v *MinNotionalFilter) MyFiltersResponseSymbolFiltersInner {
	return MyFiltersResponseSymbolFiltersInner{
		MinNotionalFilter: v,
	}
}

// NotionalFilterAsMyFiltersResponseSymbolFiltersInner is a convenience function that returns NotionalFilter wrapped in MyFiltersResponseSymbolFiltersInner
func NotionalFilterAsMyFiltersResponseSymbolFiltersInner(v *NotionalFilter) MyFiltersResponseSymbolFiltersInner {
	return MyFiltersResponseSymbolFiltersInner{
		NotionalFilter: v,
	}
}

// PercentPriceBySideFilterAsMyFiltersResponseSymbolFiltersInner is a convenience function that returns PercentPriceBySideFilter wrapped in MyFiltersResponseSymbolFiltersInner
func PercentPriceBySideFilterAsMyFiltersResponseSymbolFiltersInner(v *PercentPriceBySideFilter) MyFiltersResponseSymbolFiltersInner {
	return MyFiltersResponseSymbolFiltersInner{
		PercentPriceBySideFilter: v,
	}
}

// PercentPriceFilterAsMyFiltersResponseSymbolFiltersInner is a convenience function that returns PercentPriceFilter wrapped in MyFiltersResponseSymbolFiltersInner
func PercentPriceFilterAsMyFiltersResponseSymbolFiltersInner(v *PercentPriceFilter) MyFiltersResponseSymbolFiltersInner {
	return MyFiltersResponseSymbolFiltersInner{
		PercentPriceFilter: v,
	}
}

// PriceFilterAsMyFiltersResponseSymbolFiltersInner is a convenience function that returns PriceFilter wrapped in MyFiltersResponseSymbolFiltersInner
func PriceFilterAsMyFiltersResponseSymbolFiltersInner(v *PriceFilter) MyFiltersResponseSymbolFiltersInner {
	return MyFiltersResponseSymbolFiltersInner{
		PriceFilter: v,
	}
}

// TPlusSellFilterAsMyFiltersResponseSymbolFiltersInner is a convenience function that returns TPlusSellFilter wrapped in MyFiltersResponseSymbolFiltersInner
func TPlusSellFilterAsMyFiltersResponseSymbolFiltersInner(v *TPlusSellFilter) MyFiltersResponseSymbolFiltersInner {
	return MyFiltersResponseSymbolFiltersInner{
		TPlusSellFilter: v,
	}
}

// TrailingDeltaFilterAsMyFiltersResponseSymbolFiltersInner is a convenience function that returns TrailingDeltaFilter wrapped in MyFiltersResponseSymbolFiltersInner
func TrailingDeltaFilterAsMyFiltersResponseSymbolFiltersInner(v *TrailingDeltaFilter) MyFiltersResponseSymbolFiltersInner {
	return MyFiltersResponseSymbolFiltersInner{
		TrailingDeltaFilter: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MyFiltersResponseSymbolFiltersInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = common.NewStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ICEBERG_PARTS'
	if jsonDict["filterType"] == "ICEBERG_PARTS" {
		// try to unmarshal JSON data into IcebergPartsFilter
		err = json.Unmarshal(data, &dst.IcebergPartsFilter)
		if err == nil {
			return nil // data stored in dst.IcebergPartsFilter, return on the first match
		} else {
			dst.IcebergPartsFilter = nil
			return fmt.Errorf("failed to unmarshal MyFiltersResponseSymbolFiltersInner as IcebergPartsFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LOT_SIZE'
	if jsonDict["filterType"] == "LOT_SIZE" {
		// try to unmarshal JSON data into LotSizeFilter
		err = json.Unmarshal(data, &dst.LotSizeFilter)
		if err == nil {
			return nil // data stored in dst.LotSizeFilter, return on the first match
		} else {
			dst.LotSizeFilter = nil
			return fmt.Errorf("failed to unmarshal MyFiltersResponseSymbolFiltersInner as LotSizeFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MARKET_LOT_SIZE'
	if jsonDict["filterType"] == "MARKET_LOT_SIZE" {
		// try to unmarshal JSON data into MarketLotSizeFilter
		err = json.Unmarshal(data, &dst.MarketLotSizeFilter)
		if err == nil {
			return nil // data stored in dst.MarketLotSizeFilter, return on the first match
		} else {
			dst.MarketLotSizeFilter = nil
			return fmt.Errorf("failed to unmarshal MyFiltersResponseSymbolFiltersInner as MarketLotSizeFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MAX_NUM_ALGO_ORDERS'
	if jsonDict["filterType"] == "MAX_NUM_ALGO_ORDERS" {
		// try to unmarshal JSON data into MaxNumAlgoOrdersFilter
		err = json.Unmarshal(data, &dst.MaxNumAlgoOrdersFilter)
		if err == nil {
			return nil // data stored in dst.MaxNumAlgoOrdersFilter, return on the first match
		} else {
			dst.MaxNumAlgoOrdersFilter = nil
			return fmt.Errorf("failed to unmarshal MyFiltersResponseSymbolFiltersInner as MaxNumAlgoOrdersFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MAX_NUM_ICEBERG_ORDERS'
	if jsonDict["filterType"] == "MAX_NUM_ICEBERG_ORDERS" {
		// try to unmarshal JSON data into MaxNumIcebergOrdersFilter
		err = json.Unmarshal(data, &dst.MaxNumIcebergOrdersFilter)
		if err == nil {
			return nil // data stored in dst.MaxNumIcebergOrdersFilter, return on the first match
		} else {
			dst.MaxNumIcebergOrdersFilter = nil
			return fmt.Errorf("failed to unmarshal MyFiltersResponseSymbolFiltersInner as MaxNumIcebergOrdersFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MAX_NUM_ORDERS'
	if jsonDict["filterType"] == "MAX_NUM_ORDERS" {
		// try to unmarshal JSON data into MaxNumOrdersFilter
		err = json.Unmarshal(data, &dst.MaxNumOrdersFilter)
		if err == nil {
			return nil // data stored in dst.MaxNumOrdersFilter, return on the first match
		} else {
			dst.MaxNumOrdersFilter = nil
			return fmt.Errorf("failed to unmarshal MyFiltersResponseSymbolFiltersInner as MaxNumOrdersFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MAX_NUM_ORDER_AMENDS'
	if jsonDict["filterType"] == "MAX_NUM_ORDER_AMENDS" {
		// try to unmarshal JSON data into MaxNumOrderAmendsFilter
		err = json.Unmarshal(data, &dst.MaxNumOrderAmendsFilter)
		if err == nil {
			return nil // data stored in dst.MaxNumOrderAmendsFilter, return on the first match
		} else {
			dst.MaxNumOrderAmendsFilter = nil
			return fmt.Errorf("failed to unmarshal MyFiltersResponseSymbolFiltersInner as MaxNumOrderAmendsFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MAX_NUM_ORDER_LISTS'
	if jsonDict["filterType"] == "MAX_NUM_ORDER_LISTS" {
		// try to unmarshal JSON data into MaxNumOrderListsFilter
		err = json.Unmarshal(data, &dst.MaxNumOrderListsFilter)
		if err == nil {
			return nil // data stored in dst.MaxNumOrderListsFilter, return on the first match
		} else {
			dst.MaxNumOrderListsFilter = nil
			return fmt.Errorf("failed to unmarshal MyFiltersResponseSymbolFiltersInner as MaxNumOrderListsFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MAX_POSITION'
	if jsonDict["filterType"] == "MAX_POSITION" {
		// try to unmarshal JSON data into MaxPositionFilter
		err = json.Unmarshal(data, &dst.MaxPositionFilter)
		if err == nil {
			return nil // data stored in dst.MaxPositionFilter, return on the first match
		} else {
			dst.MaxPositionFilter = nil
			return fmt.Errorf("failed to unmarshal MyFiltersResponseSymbolFiltersInner as MaxPositionFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MIN_NOTIONAL'
	if jsonDict["filterType"] == "MIN_NOTIONAL" {
		// try to unmarshal JSON data into MinNotionalFilter
		err = json.Unmarshal(data, &dst.MinNotionalFilter)
		if err == nil {
			return nil // data stored in dst.MinNotionalFilter, return on the first match
		} else {
			dst.MinNotionalFilter = nil
			return fmt.Errorf("failed to unmarshal MyFiltersResponseSymbolFiltersInner as MinNotionalFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'NOTIONAL'
	if jsonDict["filterType"] == "NOTIONAL" {
		// try to unmarshal JSON data into NotionalFilter
		err = json.Unmarshal(data, &dst.NotionalFilter)
		if err == nil {
			return nil // data stored in dst.NotionalFilter, return on the first match
		} else {
			dst.NotionalFilter = nil
			return fmt.Errorf("failed to unmarshal MyFiltersResponseSymbolFiltersInner as NotionalFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PERCENT_PRICE'
	if jsonDict["filterType"] == "PERCENT_PRICE" {
		// try to unmarshal JSON data into PercentPriceFilter
		err = json.Unmarshal(data, &dst.PercentPriceFilter)
		if err == nil {
			return nil // data stored in dst.PercentPriceFilter, return on the first match
		} else {
			dst.PercentPriceFilter = nil
			return fmt.Errorf("failed to unmarshal MyFiltersResponseSymbolFiltersInner as PercentPriceFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PERCENT_PRICE_BY_SIDE'
	if jsonDict["filterType"] == "PERCENT_PRICE_BY_SIDE" {
		// try to unmarshal JSON data into PercentPriceBySideFilter
		err = json.Unmarshal(data, &dst.PercentPriceBySideFilter)
		if err == nil {
			return nil // data stored in dst.PercentPriceBySideFilter, return on the first match
		} else {
			dst.PercentPriceBySideFilter = nil
			return fmt.Errorf("failed to unmarshal MyFiltersResponseSymbolFiltersInner as PercentPriceBySideFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PRICE_FILTER'
	if jsonDict["filterType"] == "PRICE_FILTER" {
		// try to unmarshal JSON data into PriceFilter
		err = json.Unmarshal(data, &dst.PriceFilter)
		if err == nil {
			return nil // data stored in dst.PriceFilter, return on the first match
		} else {
			dst.PriceFilter = nil
			return fmt.Errorf("failed to unmarshal MyFiltersResponseSymbolFiltersInner as PriceFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TRAILING_DELTA'
	if jsonDict["filterType"] == "TRAILING_DELTA" {
		// try to unmarshal JSON data into TrailingDeltaFilter
		err = json.Unmarshal(data, &dst.TrailingDeltaFilter)
		if err == nil {
			return nil // data stored in dst.TrailingDeltaFilter, return on the first match
		} else {
			dst.TrailingDeltaFilter = nil
			return fmt.Errorf("failed to unmarshal MyFiltersResponseSymbolFiltersInner as TrailingDeltaFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'T_PLUS_SELL'
	if jsonDict["filterType"] == "T_PLUS_SELL" {
		// try to unmarshal JSON data into TPlusSellFilter
		err = json.Unmarshal(data, &dst.TPlusSellFilter)
		if err == nil {
			return nil // data stored in dst.TPlusSellFilter, return on the first match
		} else {
			dst.TPlusSellFilter = nil
			return fmt.Errorf("failed to unmarshal MyFiltersResponseSymbolFiltersInner as TPlusSellFilter: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MyFiltersResponseSymbolFiltersInner) MarshalJSON() ([]byte, error) {
	if src.IcebergPartsFilter != nil {
		return json.Marshal(&src.IcebergPartsFilter)
	}

	if src.LotSizeFilter != nil {
		return json.Marshal(&src.LotSizeFilter)
	}

	if src.MarketLotSizeFilter != nil {
		return json.Marshal(&src.MarketLotSizeFilter)
	}

	if src.MaxNumAlgoOrdersFilter != nil {
		return json.Marshal(&src.MaxNumAlgoOrdersFilter)
	}

	if src.MaxNumIcebergOrdersFilter != nil {
		return json.Marshal(&src.MaxNumIcebergOrdersFilter)
	}

	if src.MaxNumOrderAmendsFilter != nil {
		return json.Marshal(&src.MaxNumOrderAmendsFilter)
	}

	if src.MaxNumOrderListsFilter != nil {
		return json.Marshal(&src.MaxNumOrderListsFilter)
	}

	if src.MaxNumOrdersFilter != nil {
		return json.Marshal(&src.MaxNumOrdersFilter)
	}

	if src.MaxPositionFilter != nil {
		return json.Marshal(&src.MaxPositionFilter)
	}

	if src.MinNotionalFilter != nil {
		return json.Marshal(&src.MinNotionalFilter)
	}

	if src.NotionalFilter != nil {
		return json.Marshal(&src.NotionalFilter)
	}

	if src.PercentPriceBySideFilter != nil {
		return json.Marshal(&src.PercentPriceBySideFilter)
	}

	if src.PercentPriceFilter != nil {
		return json.Marshal(&src.PercentPriceFilter)
	}

	if src.PriceFilter != nil {
		return json.Marshal(&src.PriceFilter)
	}

	if src.TPlusSellFilter != nil {
		return json.Marshal(&src.TPlusSellFilter)
	}

	if src.TrailingDeltaFilter != nil {
		return json.Marshal(&src.TrailingDeltaFilter)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MyFiltersResponseSymbolFiltersInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IcebergPartsFilter != nil {
		return obj.IcebergPartsFilter
	}

	if obj.LotSizeFilter != nil {
		return obj.LotSizeFilter
	}

	if obj.MarketLotSizeFilter != nil {
		return obj.MarketLotSizeFilter
	}

	if obj.MaxNumAlgoOrdersFilter != nil {
		return obj.MaxNumAlgoOrdersFilter
	}

	if obj.MaxNumIcebergOrdersFilter != nil {
		return obj.MaxNumIcebergOrdersFilter
	}

	if obj.MaxNumOrderAmendsFilter != nil {
		return obj.MaxNumOrderAmendsFilter
	}

	if obj.MaxNumOrderListsFilter != nil {
		return obj.MaxNumOrderListsFilter
	}

	if obj.MaxNumOrdersFilter != nil {
		return obj.MaxNumOrdersFilter
	}

	if obj.MaxPositionFilter != nil {
		return obj.MaxPositionFilter
	}

	if obj.MinNotionalFilter != nil {
		return obj.MinNotionalFilter
	}

	if obj.NotionalFilter != nil {
		return obj.NotionalFilter
	}

	if obj.PercentPriceBySideFilter != nil {
		return obj.PercentPriceBySideFilter
	}

	if obj.PercentPriceFilter != nil {
		return obj.PercentPriceFilter
	}

	if obj.PriceFilter != nil {
		return obj.PriceFilter
	}

	if obj.TPlusSellFilter != nil {
		return obj.TPlusSellFilter
	}

	if obj.TrailingDeltaFilter != nil {
		return obj.TrailingDeltaFilter
	}

	// all schemas are nil
	return nil
}

type NullableMyFiltersResponseSymbolFiltersInner struct {
	value *MyFiltersResponseSymbolFiltersInner
	isSet bool
}

func (v NullableMyFiltersResponseSymbolFiltersInner) Get() *MyFiltersResponseSymbolFiltersInner {
	return v.value
}

func (v *NullableMyFiltersResponseSymbolFiltersInner) Set(val *MyFiltersResponseSymbolFiltersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMyFiltersResponseSymbolFiltersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMyFiltersResponseSymbolFiltersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyFiltersResponseSymbolFiltersInner(val *MyFiltersResponseSymbolFiltersInner) *NullableMyFiltersResponseSymbolFiltersInner {
	return &NullableMyFiltersResponseSymbolFiltersInner{value: val, isSet: true}
}

func (v NullableMyFiltersResponseSymbolFiltersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyFiltersResponseSymbolFiltersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
