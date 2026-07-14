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

// MyFiltersResponseExchangeFiltersInner - struct for MyFiltersResponseExchangeFiltersInner
type MyFiltersResponseExchangeFiltersInner struct {
	ExchangeMaxNumAlgoOrdersFilter    *ExchangeMaxNumAlgoOrdersFilter
	ExchangeMaxNumIcebergOrdersFilter *ExchangeMaxNumIcebergOrdersFilter
	ExchangeMaxNumOrderListsFilter    *ExchangeMaxNumOrderListsFilter
	ExchangeMaxNumOrdersFilter        *ExchangeMaxNumOrdersFilter
}

// ExchangeMaxNumAlgoOrdersFilterAsMyFiltersResponseExchangeFiltersInner is a convenience function that returns ExchangeMaxNumAlgoOrdersFilter wrapped in MyFiltersResponseExchangeFiltersInner
func ExchangeMaxNumAlgoOrdersFilterAsMyFiltersResponseExchangeFiltersInner(v *ExchangeMaxNumAlgoOrdersFilter) MyFiltersResponseExchangeFiltersInner {
	return MyFiltersResponseExchangeFiltersInner{
		ExchangeMaxNumAlgoOrdersFilter: v,
	}
}

// ExchangeMaxNumIcebergOrdersFilterAsMyFiltersResponseExchangeFiltersInner is a convenience function that returns ExchangeMaxNumIcebergOrdersFilter wrapped in MyFiltersResponseExchangeFiltersInner
func ExchangeMaxNumIcebergOrdersFilterAsMyFiltersResponseExchangeFiltersInner(v *ExchangeMaxNumIcebergOrdersFilter) MyFiltersResponseExchangeFiltersInner {
	return MyFiltersResponseExchangeFiltersInner{
		ExchangeMaxNumIcebergOrdersFilter: v,
	}
}

// ExchangeMaxNumOrderListsFilterAsMyFiltersResponseExchangeFiltersInner is a convenience function that returns ExchangeMaxNumOrderListsFilter wrapped in MyFiltersResponseExchangeFiltersInner
func ExchangeMaxNumOrderListsFilterAsMyFiltersResponseExchangeFiltersInner(v *ExchangeMaxNumOrderListsFilter) MyFiltersResponseExchangeFiltersInner {
	return MyFiltersResponseExchangeFiltersInner{
		ExchangeMaxNumOrderListsFilter: v,
	}
}

// ExchangeMaxNumOrdersFilterAsMyFiltersResponseExchangeFiltersInner is a convenience function that returns ExchangeMaxNumOrdersFilter wrapped in MyFiltersResponseExchangeFiltersInner
func ExchangeMaxNumOrdersFilterAsMyFiltersResponseExchangeFiltersInner(v *ExchangeMaxNumOrdersFilter) MyFiltersResponseExchangeFiltersInner {
	return MyFiltersResponseExchangeFiltersInner{
		ExchangeMaxNumOrdersFilter: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MyFiltersResponseExchangeFiltersInner) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal MyFiltersResponseExchangeFiltersInner as ExchangeMaxNumAlgoOrdersFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal MyFiltersResponseExchangeFiltersInner as ExchangeMaxNumIcebergOrdersFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal MyFiltersResponseExchangeFiltersInner as ExchangeMaxNumOrdersFilter: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal MyFiltersResponseExchangeFiltersInner as ExchangeMaxNumOrderListsFilter: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MyFiltersResponseExchangeFiltersInner) MarshalJSON() ([]byte, error) {
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
func (obj *MyFiltersResponseExchangeFiltersInner) GetActualInstance() interface{} {
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

type NullableMyFiltersResponseExchangeFiltersInner struct {
	value *MyFiltersResponseExchangeFiltersInner
	isSet bool
}

func (v NullableMyFiltersResponseExchangeFiltersInner) Get() *MyFiltersResponseExchangeFiltersInner {
	return v.value
}

func (v *NullableMyFiltersResponseExchangeFiltersInner) Set(val *MyFiltersResponseExchangeFiltersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMyFiltersResponseExchangeFiltersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMyFiltersResponseExchangeFiltersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyFiltersResponseExchangeFiltersInner(val *MyFiltersResponseExchangeFiltersInner) *NullableMyFiltersResponseExchangeFiltersInner {
	return &NullableMyFiltersResponseExchangeFiltersInner{value: val, isSet: true}
}

func (v NullableMyFiltersResponseExchangeFiltersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyFiltersResponseExchangeFiltersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
