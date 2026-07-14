/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// QueryCrossIsolatedMarginCapitalFlowTypeParameter the model 'QueryCrossIsolatedMarginCapitalFlowTypeParameter'
type QueryCrossIsolatedMarginCapitalFlowTypeParameter string

// List of queryCrossIsolatedMarginCapitalFlow_type_parameter
const (
	QueryCrossIsolatedMarginCapitalFlowTypeParameterTransfer            QueryCrossIsolatedMarginCapitalFlowTypeParameter = "TRANSFER"
	QueryCrossIsolatedMarginCapitalFlowTypeParameterBorrow              QueryCrossIsolatedMarginCapitalFlowTypeParameter = "BORROW"
	QueryCrossIsolatedMarginCapitalFlowTypeParameterRepay               QueryCrossIsolatedMarginCapitalFlowTypeParameter = "REPAY"
	QueryCrossIsolatedMarginCapitalFlowTypeParameterBuyIncome           QueryCrossIsolatedMarginCapitalFlowTypeParameter = "BUY_INCOME"
	QueryCrossIsolatedMarginCapitalFlowTypeParameterBuyExpense          QueryCrossIsolatedMarginCapitalFlowTypeParameter = "BUY_EXPENSE"
	QueryCrossIsolatedMarginCapitalFlowTypeParameterSellIncome          QueryCrossIsolatedMarginCapitalFlowTypeParameter = "SELL_INCOME"
	QueryCrossIsolatedMarginCapitalFlowTypeParameterSellExpense         QueryCrossIsolatedMarginCapitalFlowTypeParameter = "SELL_EXPENSE"
	QueryCrossIsolatedMarginCapitalFlowTypeParameterTradingCommission   QueryCrossIsolatedMarginCapitalFlowTypeParameter = "TRADING_COMMISSION"
	QueryCrossIsolatedMarginCapitalFlowTypeParameterBuyLiquidation      QueryCrossIsolatedMarginCapitalFlowTypeParameter = "BUY_LIQUIDATION"
	QueryCrossIsolatedMarginCapitalFlowTypeParameterSellLiquidation     QueryCrossIsolatedMarginCapitalFlowTypeParameter = "SELL_LIQUIDATION"
	QueryCrossIsolatedMarginCapitalFlowTypeParameterRepayLiquidation    QueryCrossIsolatedMarginCapitalFlowTypeParameter = "REPAY_LIQUIDATION"
	QueryCrossIsolatedMarginCapitalFlowTypeParameterOtherLiquidation    QueryCrossIsolatedMarginCapitalFlowTypeParameter = "OTHER_LIQUIDATION"
	QueryCrossIsolatedMarginCapitalFlowTypeParameterLiquidationFee      QueryCrossIsolatedMarginCapitalFlowTypeParameter = "LIQUIDATION_FEE"
	QueryCrossIsolatedMarginCapitalFlowTypeParameterSmallBalanceConvert QueryCrossIsolatedMarginCapitalFlowTypeParameter = "SMALL_BALANCE_CONVERT"
	QueryCrossIsolatedMarginCapitalFlowTypeParameterCommissionReturn    QueryCrossIsolatedMarginCapitalFlowTypeParameter = "COMMISSION_RETURN"
	QueryCrossIsolatedMarginCapitalFlowTypeParameterSmallConvert        QueryCrossIsolatedMarginCapitalFlowTypeParameter = "SMALL_CONVERT"
)

// All allowed values of QueryCrossIsolatedMarginCapitalFlowTypeParameter enum
var AllowedQueryCrossIsolatedMarginCapitalFlowTypeParameterEnumValues = []QueryCrossIsolatedMarginCapitalFlowTypeParameter{
	"TRANSFER",
	"BORROW",
	"REPAY",
	"BUY_INCOME",
	"BUY_EXPENSE",
	"SELL_INCOME",
	"SELL_EXPENSE",
	"TRADING_COMMISSION",
	"BUY_LIQUIDATION",
	"SELL_LIQUIDATION",
	"REPAY_LIQUIDATION",
	"OTHER_LIQUIDATION",
	"LIQUIDATION_FEE",
	"SMALL_BALANCE_CONVERT",
	"COMMISSION_RETURN",
	"SMALL_CONVERT",
}

func (v *QueryCrossIsolatedMarginCapitalFlowTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryCrossIsolatedMarginCapitalFlowTypeParameter(value)
	for _, existing := range AllowedQueryCrossIsolatedMarginCapitalFlowTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryCrossIsolatedMarginCapitalFlowTypeParameter", value)
}

// NewQueryCrossIsolatedMarginCapitalFlowTypeParameterFromValue returns a pointer to a valid QueryCrossIsolatedMarginCapitalFlowTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryCrossIsolatedMarginCapitalFlowTypeParameterFromValue(v string) (*QueryCrossIsolatedMarginCapitalFlowTypeParameter, error) {
	ev := QueryCrossIsolatedMarginCapitalFlowTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryCrossIsolatedMarginCapitalFlowTypeParameter: valid values are %v", v, AllowedQueryCrossIsolatedMarginCapitalFlowTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryCrossIsolatedMarginCapitalFlowTypeParameter) IsValid() bool {
	for _, existing := range AllowedQueryCrossIsolatedMarginCapitalFlowTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to queryCrossIsolatedMarginCapitalFlow_type_parameter value
func (v QueryCrossIsolatedMarginCapitalFlowTypeParameter) Ptr() *QueryCrossIsolatedMarginCapitalFlowTypeParameter {
	return &v
}

type NullableQueryCrossIsolatedMarginCapitalFlowTypeParameter struct {
	value *QueryCrossIsolatedMarginCapitalFlowTypeParameter
	isSet bool
}

func (v NullableQueryCrossIsolatedMarginCapitalFlowTypeParameter) Get() *QueryCrossIsolatedMarginCapitalFlowTypeParameter {
	return v.value
}

func (v *NullableQueryCrossIsolatedMarginCapitalFlowTypeParameter) Set(val *QueryCrossIsolatedMarginCapitalFlowTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCrossIsolatedMarginCapitalFlowTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCrossIsolatedMarginCapitalFlowTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCrossIsolatedMarginCapitalFlowTypeParameter(val *QueryCrossIsolatedMarginCapitalFlowTypeParameter) *NullableQueryCrossIsolatedMarginCapitalFlowTypeParameter {
	return &NullableQueryCrossIsolatedMarginCapitalFlowTypeParameter{value: val, isSet: true}
}

func (v NullableQueryCrossIsolatedMarginCapitalFlowTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCrossIsolatedMarginCapitalFlowTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
