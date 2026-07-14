/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetIncomeHistoryIncomeTypeParameter the model 'GetIncomeHistoryIncomeTypeParameter'
type GetIncomeHistoryIncomeTypeParameter string

// List of getIncomeHistory_incomeType_parameter
const (
	GetIncomeHistoryIncomeTypeParameterTransfer                  GetIncomeHistoryIncomeTypeParameter = "TRANSFER"
	GetIncomeHistoryIncomeTypeParameterWelcomeBonus              GetIncomeHistoryIncomeTypeParameter = "WELCOME_BONUS"
	GetIncomeHistoryIncomeTypeParameterRealizedPnl               GetIncomeHistoryIncomeTypeParameter = "REALIZED_PNL"
	GetIncomeHistoryIncomeTypeParameterFundingFee                GetIncomeHistoryIncomeTypeParameter = "FUNDING_FEE"
	GetIncomeHistoryIncomeTypeParameterCommission                GetIncomeHistoryIncomeTypeParameter = "COMMISSION"
	GetIncomeHistoryIncomeTypeParameterInsuranceClear            GetIncomeHistoryIncomeTypeParameter = "INSURANCE_CLEAR"
	GetIncomeHistoryIncomeTypeParameterReferralKickback          GetIncomeHistoryIncomeTypeParameter = "REFERRAL_KICKBACK"
	GetIncomeHistoryIncomeTypeParameterCommissionRebate          GetIncomeHistoryIncomeTypeParameter = "COMMISSION_REBATE"
	GetIncomeHistoryIncomeTypeParameterApiRebate                 GetIncomeHistoryIncomeTypeParameter = "API_REBATE"
	GetIncomeHistoryIncomeTypeParameterContestReward             GetIncomeHistoryIncomeTypeParameter = "CONTEST_REWARD"
	GetIncomeHistoryIncomeTypeParameterCrossCollateralTransfer   GetIncomeHistoryIncomeTypeParameter = "CROSS_COLLATERAL_TRANSFER"
	GetIncomeHistoryIncomeTypeParameterOptionsPremiumFee         GetIncomeHistoryIncomeTypeParameter = "OPTIONS_PREMIUM_FEE"
	GetIncomeHistoryIncomeTypeParameterOptionsSettleProfit       GetIncomeHistoryIncomeTypeParameter = "OPTIONS_SETTLE_PROFIT"
	GetIncomeHistoryIncomeTypeParameterInternalTransfer          GetIncomeHistoryIncomeTypeParameter = "INTERNAL_TRANSFER"
	GetIncomeHistoryIncomeTypeParameterAutoExchange              GetIncomeHistoryIncomeTypeParameter = "AUTO_EXCHANGE"
	GetIncomeHistoryIncomeTypeParameterDeliveredSettelment       GetIncomeHistoryIncomeTypeParameter = "DELIVERED_SETTELMENT"
	GetIncomeHistoryIncomeTypeParameterCoinSwapDeposit           GetIncomeHistoryIncomeTypeParameter = "COIN_SWAP_DEPOSIT"
	GetIncomeHistoryIncomeTypeParameterCoinSwapWithdraw          GetIncomeHistoryIncomeTypeParameter = "COIN_SWAP_WITHDRAW"
	GetIncomeHistoryIncomeTypeParameterPositionLimitIncreaseFee  GetIncomeHistoryIncomeTypeParameter = "POSITION_LIMIT_INCREASE_FEE"
	GetIncomeHistoryIncomeTypeParameterStrategyUmfuturesTransfer GetIncomeHistoryIncomeTypeParameter = "STRATEGY_UMFUTURES_TRANSFER"
	GetIncomeHistoryIncomeTypeParameterFeeReturn                 GetIncomeHistoryIncomeTypeParameter = "FEE_RETURN"
	GetIncomeHistoryIncomeTypeParameterBfusdReward               GetIncomeHistoryIncomeTypeParameter = "BFUSD_REWARD"
)

// All allowed values of GetIncomeHistoryIncomeTypeParameter enum
var AllowedGetIncomeHistoryIncomeTypeParameterEnumValues = []GetIncomeHistoryIncomeTypeParameter{
	"TRANSFER",
	"WELCOME_BONUS",
	"REALIZED_PNL",
	"FUNDING_FEE",
	"COMMISSION",
	"INSURANCE_CLEAR",
	"REFERRAL_KICKBACK",
	"COMMISSION_REBATE",
	"API_REBATE",
	"CONTEST_REWARD",
	"CROSS_COLLATERAL_TRANSFER",
	"OPTIONS_PREMIUM_FEE",
	"OPTIONS_SETTLE_PROFIT",
	"INTERNAL_TRANSFER",
	"AUTO_EXCHANGE",
	"DELIVERED_SETTELMENT",
	"COIN_SWAP_DEPOSIT",
	"COIN_SWAP_WITHDRAW",
	"POSITION_LIMIT_INCREASE_FEE",
	"STRATEGY_UMFUTURES_TRANSFER",
	"FEE_RETURN",
	"BFUSD_REWARD",
}

func (v *GetIncomeHistoryIncomeTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetIncomeHistoryIncomeTypeParameter(value)
	for _, existing := range AllowedGetIncomeHistoryIncomeTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetIncomeHistoryIncomeTypeParameter", value)
}

// NewGetIncomeHistoryIncomeTypeParameterFromValue returns a pointer to a valid GetIncomeHistoryIncomeTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetIncomeHistoryIncomeTypeParameterFromValue(v string) (*GetIncomeHistoryIncomeTypeParameter, error) {
	ev := GetIncomeHistoryIncomeTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetIncomeHistoryIncomeTypeParameter: valid values are %v", v, AllowedGetIncomeHistoryIncomeTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetIncomeHistoryIncomeTypeParameter) IsValid() bool {
	for _, existing := range AllowedGetIncomeHistoryIncomeTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getIncomeHistory_incomeType_parameter value
func (v GetIncomeHistoryIncomeTypeParameter) Ptr() *GetIncomeHistoryIncomeTypeParameter {
	return &v
}

type NullableGetIncomeHistoryIncomeTypeParameter struct {
	value *GetIncomeHistoryIncomeTypeParameter
	isSet bool
}

func (v NullableGetIncomeHistoryIncomeTypeParameter) Get() *GetIncomeHistoryIncomeTypeParameter {
	return v.value
}

func (v *NullableGetIncomeHistoryIncomeTypeParameter) Set(val *GetIncomeHistoryIncomeTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIncomeHistoryIncomeTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIncomeHistoryIncomeTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIncomeHistoryIncomeTypeParameter(val *GetIncomeHistoryIncomeTypeParameter) *NullableGetIncomeHistoryIncomeTypeParameter {
	return &NullableGetIncomeHistoryIncomeTypeParameter{value: val, isSet: true}
}

func (v NullableGetIncomeHistoryIncomeTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIncomeHistoryIncomeTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
