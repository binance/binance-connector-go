/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/common"
)

// SymbolFilters - struct for SymbolFilters
type SymbolFilters struct {
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

// IcebergPartsFilterAsSymbolFilters is a convenience function that returns IcebergPartsFilter wrapped in SymbolFilters
func IcebergPartsFilterAsSymbolFilters(v *IcebergPartsFilter) SymbolFilters {
	return SymbolFilters{
		IcebergPartsFilter: v,
	}
}

// LotSizeFilterAsSymbolFilters is a convenience function that returns LotSizeFilter wrapped in SymbolFilters
func LotSizeFilterAsSymbolFilters(v *LotSizeFilter) SymbolFilters {
	return SymbolFilters{
		LotSizeFilter: v,
	}
}

// MarketLotSizeFilterAsSymbolFilters is a convenience function that returns MarketLotSizeFilter wrapped in SymbolFilters
func MarketLotSizeFilterAsSymbolFilters(v *MarketLotSizeFilter) SymbolFilters {
	return SymbolFilters{
		MarketLotSizeFilter: v,
	}
}

// MaxNumAlgoOrdersFilterAsSymbolFilters is a convenience function that returns MaxNumAlgoOrdersFilter wrapped in SymbolFilters
func MaxNumAlgoOrdersFilterAsSymbolFilters(v *MaxNumAlgoOrdersFilter) SymbolFilters {
	return SymbolFilters{
		MaxNumAlgoOrdersFilter: v,
	}
}

// MaxNumIcebergOrdersFilterAsSymbolFilters is a convenience function that returns MaxNumIcebergOrdersFilter wrapped in SymbolFilters
func MaxNumIcebergOrdersFilterAsSymbolFilters(v *MaxNumIcebergOrdersFilter) SymbolFilters {
	return SymbolFilters{
		MaxNumIcebergOrdersFilter: v,
	}
}

// MaxNumOrderAmendsFilterAsSymbolFilters is a convenience function that returns MaxNumOrderAmendsFilter wrapped in SymbolFilters
func MaxNumOrderAmendsFilterAsSymbolFilters(v *MaxNumOrderAmendsFilter) SymbolFilters {
	return SymbolFilters{
		MaxNumOrderAmendsFilter: v,
	}
}

// MaxNumOrderListsFilterAsSymbolFilters is a convenience function that returns MaxNumOrderListsFilter wrapped in SymbolFilters
func MaxNumOrderListsFilterAsSymbolFilters(v *MaxNumOrderListsFilter) SymbolFilters {
	return SymbolFilters{
		MaxNumOrderListsFilter: v,
	}
}

// MaxNumOrdersFilterAsSymbolFilters is a convenience function that returns MaxNumOrdersFilter wrapped in SymbolFilters
func MaxNumOrdersFilterAsSymbolFilters(v *MaxNumOrdersFilter) SymbolFilters {
	return SymbolFilters{
		MaxNumOrdersFilter: v,
	}
}

// MaxPositionFilterAsSymbolFilters is a convenience function that returns MaxPositionFilter wrapped in SymbolFilters
func MaxPositionFilterAsSymbolFilters(v *MaxPositionFilter) SymbolFilters {
	return SymbolFilters{
		MaxPositionFilter: v,
	}
}

// MinNotionalFilterAsSymbolFilters is a convenience function that returns MinNotionalFilter wrapped in SymbolFilters
func MinNotionalFilterAsSymbolFilters(v *MinNotionalFilter) SymbolFilters {
	return SymbolFilters{
		MinNotionalFilter: v,
	}
}

// NotionalFilterAsSymbolFilters is a convenience function that returns NotionalFilter wrapped in SymbolFilters
func NotionalFilterAsSymbolFilters(v *NotionalFilter) SymbolFilters {
	return SymbolFilters{
		NotionalFilter: v,
	}
}

// PercentPriceBySideFilterAsSymbolFilters is a convenience function that returns PercentPriceBySideFilter wrapped in SymbolFilters
func PercentPriceBySideFilterAsSymbolFilters(v *PercentPriceBySideFilter) SymbolFilters {
	return SymbolFilters{
		PercentPriceBySideFilter: v,
	}
}

// PercentPriceFilterAsSymbolFilters is a convenience function that returns PercentPriceFilter wrapped in SymbolFilters
func PercentPriceFilterAsSymbolFilters(v *PercentPriceFilter) SymbolFilters {
	return SymbolFilters{
		PercentPriceFilter: v,
	}
}

// PriceFilterAsSymbolFilters is a convenience function that returns PriceFilter wrapped in SymbolFilters
func PriceFilterAsSymbolFilters(v *PriceFilter) SymbolFilters {
	return SymbolFilters{
		PriceFilter: v,
	}
}

// TPlusSellFilterAsSymbolFilters is a convenience function that returns TPlusSellFilter wrapped in SymbolFilters
func TPlusSellFilterAsSymbolFilters(v *TPlusSellFilter) SymbolFilters {
	return SymbolFilters{
		TPlusSellFilter: v,
	}
}

// TrailingDeltaFilterAsSymbolFilters is a convenience function that returns TrailingDeltaFilter wrapped in SymbolFilters
func TrailingDeltaFilterAsSymbolFilters(v *TrailingDeltaFilter) SymbolFilters {
	return SymbolFilters{
		TrailingDeltaFilter: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SymbolFilters) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal SymbolFilters as IcebergPartsFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SymbolFilters as LotSizeFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SymbolFilters as MarketLotSizeFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SymbolFilters as MaxNumAlgoOrdersFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SymbolFilters as MaxNumIcebergOrdersFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SymbolFilters as MaxNumOrdersFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SymbolFilters as MaxNumOrderAmendsFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SymbolFilters as MaxNumOrderListsFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SymbolFilters as MaxPositionFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SymbolFilters as MinNotionalFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SymbolFilters as NotionalFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SymbolFilters as PercentPriceFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SymbolFilters as PercentPriceBySideFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SymbolFilters as PriceFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SymbolFilters as TrailingDeltaFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal SymbolFilters as TPlusSellFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IcebergPartsFilter'
	if jsonDict["filterType"] == "IcebergPartsFilter" {
		// try to unmarshal JSON data into IcebergPartsFilter
		err = json.Unmarshal(data, &dst.IcebergPartsFilter)
		if err == nil {
			return nil // data stored in dst.IcebergPartsFilter, return on the first match
		} else {
			dst.IcebergPartsFilter = nil
			return fmt.Errorf("failed to unmarshal SymbolFilters as IcebergPartsFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LotSizeFilter'
	if jsonDict["filterType"] == "LotSizeFilter" {
		// try to unmarshal JSON data into LotSizeFilter
		err = json.Unmarshal(data, &dst.LotSizeFilter)
		if err == nil {
			return nil // data stored in dst.LotSizeFilter, return on the first match
		} else {
			dst.LotSizeFilter = nil
			return fmt.Errorf("failed to unmarshal SymbolFilters as LotSizeFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MarketLotSizeFilter'
	if jsonDict["filterType"] == "MarketLotSizeFilter" {
		// try to unmarshal JSON data into MarketLotSizeFilter
		err = json.Unmarshal(data, &dst.MarketLotSizeFilter)
		if err == nil {
			return nil // data stored in dst.MarketLotSizeFilter, return on the first match
		} else {
			dst.MarketLotSizeFilter = nil
			return fmt.Errorf("failed to unmarshal SymbolFilters as MarketLotSizeFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MaxNumAlgoOrdersFilter'
	if jsonDict["filterType"] == "MaxNumAlgoOrdersFilter" {
		// try to unmarshal JSON data into MaxNumAlgoOrdersFilter
		err = json.Unmarshal(data, &dst.MaxNumAlgoOrdersFilter)
		if err == nil {
			return nil // data stored in dst.MaxNumAlgoOrdersFilter, return on the first match
		} else {
			dst.MaxNumAlgoOrdersFilter = nil
			return fmt.Errorf("failed to unmarshal SymbolFilters as MaxNumAlgoOrdersFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MaxNumIcebergOrdersFilter'
	if jsonDict["filterType"] == "MaxNumIcebergOrdersFilter" {
		// try to unmarshal JSON data into MaxNumIcebergOrdersFilter
		err = json.Unmarshal(data, &dst.MaxNumIcebergOrdersFilter)
		if err == nil {
			return nil // data stored in dst.MaxNumIcebergOrdersFilter, return on the first match
		} else {
			dst.MaxNumIcebergOrdersFilter = nil
			return fmt.Errorf("failed to unmarshal SymbolFilters as MaxNumIcebergOrdersFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MaxNumOrderAmendsFilter'
	if jsonDict["filterType"] == "MaxNumOrderAmendsFilter" {
		// try to unmarshal JSON data into MaxNumOrderAmendsFilter
		err = json.Unmarshal(data, &dst.MaxNumOrderAmendsFilter)
		if err == nil {
			return nil // data stored in dst.MaxNumOrderAmendsFilter, return on the first match
		} else {
			dst.MaxNumOrderAmendsFilter = nil
			return fmt.Errorf("failed to unmarshal SymbolFilters as MaxNumOrderAmendsFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MaxNumOrderListsFilter'
	if jsonDict["filterType"] == "MaxNumOrderListsFilter" {
		// try to unmarshal JSON data into MaxNumOrderListsFilter
		err = json.Unmarshal(data, &dst.MaxNumOrderListsFilter)
		if err == nil {
			return nil // data stored in dst.MaxNumOrderListsFilter, return on the first match
		} else {
			dst.MaxNumOrderListsFilter = nil
			return fmt.Errorf("failed to unmarshal SymbolFilters as MaxNumOrderListsFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MaxNumOrdersFilter'
	if jsonDict["filterType"] == "MaxNumOrdersFilter" {
		// try to unmarshal JSON data into MaxNumOrdersFilter
		err = json.Unmarshal(data, &dst.MaxNumOrdersFilter)
		if err == nil {
			return nil // data stored in dst.MaxNumOrdersFilter, return on the first match
		} else {
			dst.MaxNumOrdersFilter = nil
			return fmt.Errorf("failed to unmarshal SymbolFilters as MaxNumOrdersFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MaxPositionFilter'
	if jsonDict["filterType"] == "MaxPositionFilter" {
		// try to unmarshal JSON data into MaxPositionFilter
		err = json.Unmarshal(data, &dst.MaxPositionFilter)
		if err == nil {
			return nil // data stored in dst.MaxPositionFilter, return on the first match
		} else {
			dst.MaxPositionFilter = nil
			return fmt.Errorf("failed to unmarshal SymbolFilters as MaxPositionFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MinNotionalFilter'
	if jsonDict["filterType"] == "MinNotionalFilter" {
		// try to unmarshal JSON data into MinNotionalFilter
		err = json.Unmarshal(data, &dst.MinNotionalFilter)
		if err == nil {
			return nil // data stored in dst.MinNotionalFilter, return on the first match
		} else {
			dst.MinNotionalFilter = nil
			return fmt.Errorf("failed to unmarshal SymbolFilters as MinNotionalFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'NotionalFilter'
	if jsonDict["filterType"] == "NotionalFilter" {
		// try to unmarshal JSON data into NotionalFilter
		err = json.Unmarshal(data, &dst.NotionalFilter)
		if err == nil {
			return nil // data stored in dst.NotionalFilter, return on the first match
		} else {
			dst.NotionalFilter = nil
			return fmt.Errorf("failed to unmarshal SymbolFilters as NotionalFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PercentPriceBySideFilter'
	if jsonDict["filterType"] == "PercentPriceBySideFilter" {
		// try to unmarshal JSON data into PercentPriceBySideFilter
		err = json.Unmarshal(data, &dst.PercentPriceBySideFilter)
		if err == nil {
			return nil // data stored in dst.PercentPriceBySideFilter, return on the first match
		} else {
			dst.PercentPriceBySideFilter = nil
			return fmt.Errorf("failed to unmarshal SymbolFilters as PercentPriceBySideFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PercentPriceFilter'
	if jsonDict["filterType"] == "PercentPriceFilter" {
		// try to unmarshal JSON data into PercentPriceFilter
		err = json.Unmarshal(data, &dst.PercentPriceFilter)
		if err == nil {
			return nil // data stored in dst.PercentPriceFilter, return on the first match
		} else {
			dst.PercentPriceFilter = nil
			return fmt.Errorf("failed to unmarshal SymbolFilters as PercentPriceFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PriceFilter'
	if jsonDict["filterType"] == "PriceFilter" {
		// try to unmarshal JSON data into PriceFilter
		err = json.Unmarshal(data, &dst.PriceFilter)
		if err == nil {
			return nil // data stored in dst.PriceFilter, return on the first match
		} else {
			dst.PriceFilter = nil
			return fmt.Errorf("failed to unmarshal SymbolFilters as PriceFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TPlusSellFilter'
	if jsonDict["filterType"] == "TPlusSellFilter" {
		// try to unmarshal JSON data into TPlusSellFilter
		err = json.Unmarshal(data, &dst.TPlusSellFilter)
		if err == nil {
			return nil // data stored in dst.TPlusSellFilter, return on the first match
		} else {
			dst.TPlusSellFilter = nil
			return fmt.Errorf("failed to unmarshal SymbolFilters as TPlusSellFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TrailingDeltaFilter'
	if jsonDict["filterType"] == "TrailingDeltaFilter" {
		// try to unmarshal JSON data into TrailingDeltaFilter
		err = json.Unmarshal(data, &dst.TrailingDeltaFilter)
		if err == nil {
			return nil // data stored in dst.TrailingDeltaFilter, return on the first match
		} else {
			dst.TrailingDeltaFilter = nil
			return fmt.Errorf("failed to unmarshal SymbolFilters as TrailingDeltaFilter: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SymbolFilters) MarshalJSON() ([]byte, error) {
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
func (obj *SymbolFilters) GetActualInstance() interface{} {
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

type NullableSymbolFilters struct {
	value *SymbolFilters
	isSet bool
}

func (v NullableSymbolFilters) Get() *SymbolFilters {
	return v.value
}

func (v *NullableSymbolFilters) Set(val *SymbolFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolFilters(val *SymbolFilters) *NullableSymbolFilters {
	return &NullableSymbolFilters{value: val, isSet: true}
}

func (v NullableSymbolFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
