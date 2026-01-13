/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/common"
)

// ExchangeFilters - struct for ExchangeFilters
type ExchangeFilters struct {
	ExchangeMaxNumAlgoOrdersFilter    *ExchangeMaxNumAlgoOrdersFilter
	ExchangeMaxNumIcebergOrdersFilter *ExchangeMaxNumIcebergOrdersFilter
	ExchangeMaxNumOrderListsFilter    *ExchangeMaxNumOrderListsFilter
	ExchangeMaxNumOrdersFilter        *ExchangeMaxNumOrdersFilter
}

// ExchangeMaxNumAlgoOrdersFilterAsExchangeFilters is a convenience function that returns ExchangeMaxNumAlgoOrdersFilter wrapped in ExchangeFilters
func ExchangeMaxNumAlgoOrdersFilterAsExchangeFilters(v *ExchangeMaxNumAlgoOrdersFilter) ExchangeFilters {
	return ExchangeFilters{
		ExchangeMaxNumAlgoOrdersFilter: v,
	}
}

// ExchangeMaxNumIcebergOrdersFilterAsExchangeFilters is a convenience function that returns ExchangeMaxNumIcebergOrdersFilter wrapped in ExchangeFilters
func ExchangeMaxNumIcebergOrdersFilterAsExchangeFilters(v *ExchangeMaxNumIcebergOrdersFilter) ExchangeFilters {
	return ExchangeFilters{
		ExchangeMaxNumIcebergOrdersFilter: v,
	}
}

// ExchangeMaxNumOrderListsFilterAsExchangeFilters is a convenience function that returns ExchangeMaxNumOrderListsFilter wrapped in ExchangeFilters
func ExchangeMaxNumOrderListsFilterAsExchangeFilters(v *ExchangeMaxNumOrderListsFilter) ExchangeFilters {
	return ExchangeFilters{
		ExchangeMaxNumOrderListsFilter: v,
	}
}

// ExchangeMaxNumOrdersFilterAsExchangeFilters is a convenience function that returns ExchangeMaxNumOrdersFilter wrapped in ExchangeFilters
func ExchangeMaxNumOrdersFilterAsExchangeFilters(v *ExchangeMaxNumOrdersFilter) ExchangeFilters {
	return ExchangeFilters{
		ExchangeMaxNumOrdersFilter: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ExchangeFilters) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = common.NewStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'EXCHANGE_MAX_NUM_ALGO_ORDERS'
	if jsonDict["filterType"] == "EXCHANGE_MAX_NUM_ALGO_ORDERS" {
		// try to unmarshal JSON data into ExchangeMaxNumAlgoOrdersFilter
		err = json.Unmarshal(data, &dst.ExchangeMaxNumAlgoOrdersFilter)
		if err == nil {
			return nil // data stored in dst.ExchangeMaxNumAlgoOrdersFilter, return on the first match
		} else {
			dst.ExchangeMaxNumAlgoOrdersFilter = nil
			return fmt.Errorf("failed to unmarshal ExchangeFilters as ExchangeMaxNumAlgoOrdersFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EXCHANGE_MAX_NUM_ICEBERG_ORDERS'
	if jsonDict["filterType"] == "EXCHANGE_MAX_NUM_ICEBERG_ORDERS" {
		// try to unmarshal JSON data into ExchangeMaxNumIcebergOrdersFilter
		err = json.Unmarshal(data, &dst.ExchangeMaxNumIcebergOrdersFilter)
		if err == nil {
			return nil // data stored in dst.ExchangeMaxNumIcebergOrdersFilter, return on the first match
		} else {
			dst.ExchangeMaxNumIcebergOrdersFilter = nil
			return fmt.Errorf("failed to unmarshal ExchangeFilters as ExchangeMaxNumIcebergOrdersFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EXCHANGE_MAX_NUM_ORDERS'
	if jsonDict["filterType"] == "EXCHANGE_MAX_NUM_ORDERS" {
		// try to unmarshal JSON data into ExchangeMaxNumOrdersFilter
		err = json.Unmarshal(data, &dst.ExchangeMaxNumOrdersFilter)
		if err == nil {
			return nil // data stored in dst.ExchangeMaxNumOrdersFilter, return on the first match
		} else {
			dst.ExchangeMaxNumOrdersFilter = nil
			return fmt.Errorf("failed to unmarshal ExchangeFilters as ExchangeMaxNumOrdersFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EXCHANGE_MAX_NUM_ORDER_LISTS'
	if jsonDict["filterType"] == "EXCHANGE_MAX_NUM_ORDER_LISTS" {
		// try to unmarshal JSON data into ExchangeMaxNumOrderListsFilter
		err = json.Unmarshal(data, &dst.ExchangeMaxNumOrderListsFilter)
		if err == nil {
			return nil // data stored in dst.ExchangeMaxNumOrderListsFilter, return on the first match
		} else {
			dst.ExchangeMaxNumOrderListsFilter = nil
			return fmt.Errorf("failed to unmarshal ExchangeFilters as ExchangeMaxNumOrderListsFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ExchangeMaxNumAlgoOrdersFilter'
	if jsonDict["filterType"] == "ExchangeMaxNumAlgoOrdersFilter" {
		// try to unmarshal JSON data into ExchangeMaxNumAlgoOrdersFilter
		err = json.Unmarshal(data, &dst.ExchangeMaxNumAlgoOrdersFilter)
		if err == nil {
			return nil // data stored in dst.ExchangeMaxNumAlgoOrdersFilter, return on the first match
		} else {
			dst.ExchangeMaxNumAlgoOrdersFilter = nil
			return fmt.Errorf("failed to unmarshal ExchangeFilters as ExchangeMaxNumAlgoOrdersFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ExchangeMaxNumIcebergOrdersFilter'
	if jsonDict["filterType"] == "ExchangeMaxNumIcebergOrdersFilter" {
		// try to unmarshal JSON data into ExchangeMaxNumIcebergOrdersFilter
		err = json.Unmarshal(data, &dst.ExchangeMaxNumIcebergOrdersFilter)
		if err == nil {
			return nil // data stored in dst.ExchangeMaxNumIcebergOrdersFilter, return on the first match
		} else {
			dst.ExchangeMaxNumIcebergOrdersFilter = nil
			return fmt.Errorf("failed to unmarshal ExchangeFilters as ExchangeMaxNumIcebergOrdersFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ExchangeMaxNumOrderListsFilter'
	if jsonDict["filterType"] == "ExchangeMaxNumOrderListsFilter" {
		// try to unmarshal JSON data into ExchangeMaxNumOrderListsFilter
		err = json.Unmarshal(data, &dst.ExchangeMaxNumOrderListsFilter)
		if err == nil {
			return nil // data stored in dst.ExchangeMaxNumOrderListsFilter, return on the first match
		} else {
			dst.ExchangeMaxNumOrderListsFilter = nil
			return fmt.Errorf("failed to unmarshal ExchangeFilters as ExchangeMaxNumOrderListsFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ExchangeMaxNumOrdersFilter'
	if jsonDict["filterType"] == "ExchangeMaxNumOrdersFilter" {
		// try to unmarshal JSON data into ExchangeMaxNumOrdersFilter
		err = json.Unmarshal(data, &dst.ExchangeMaxNumOrdersFilter)
		if err == nil {
			return nil // data stored in dst.ExchangeMaxNumOrdersFilter, return on the first match
		} else {
			dst.ExchangeMaxNumOrdersFilter = nil
			return fmt.Errorf("failed to unmarshal ExchangeFilters as ExchangeMaxNumOrdersFilter: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ExchangeFilters) MarshalJSON() ([]byte, error) {
	if src.ExchangeMaxNumAlgoOrdersFilter != nil {
		return json.Marshal(&src.ExchangeMaxNumAlgoOrdersFilter)
	}

	if src.ExchangeMaxNumIcebergOrdersFilter != nil {
		return json.Marshal(&src.ExchangeMaxNumIcebergOrdersFilter)
	}

	if src.ExchangeMaxNumOrderListsFilter != nil {
		return json.Marshal(&src.ExchangeMaxNumOrderListsFilter)
	}

	if src.ExchangeMaxNumOrdersFilter != nil {
		return json.Marshal(&src.ExchangeMaxNumOrdersFilter)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ExchangeFilters) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ExchangeMaxNumAlgoOrdersFilter != nil {
		return obj.ExchangeMaxNumAlgoOrdersFilter
	}

	if obj.ExchangeMaxNumIcebergOrdersFilter != nil {
		return obj.ExchangeMaxNumIcebergOrdersFilter
	}

	if obj.ExchangeMaxNumOrderListsFilter != nil {
		return obj.ExchangeMaxNumOrderListsFilter
	}

	if obj.ExchangeMaxNumOrdersFilter != nil {
		return obj.ExchangeMaxNumOrdersFilter
	}

	// all schemas are nil
	return nil
}

type NullableExchangeFilters struct {
	value *ExchangeFilters
	isSet bool
}

func (v NullableExchangeFilters) Get() *ExchangeFilters {
	return v.value
}

func (v *NullableExchangeFilters) Set(val *ExchangeFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeFilters(val *ExchangeFilters) *NullableExchangeFilters {
	return &NullableExchangeFilters{value: val, isSet: true}
}

func (v NullableExchangeFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
