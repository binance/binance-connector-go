/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetUmIncomeHistoryIncomeTypeParameter the model 'GetUmIncomeHistoryIncomeTypeParameter'
type GetUmIncomeHistoryIncomeTypeParameter string

// List of getUmIncomeHistory_incomeType_parameter
const (
	GetUmIncomeHistoryIncomeTypeParameterTransfer                 GetUmIncomeHistoryIncomeTypeParameter = "TRANSFER"
	GetUmIncomeHistoryIncomeTypeParameterWelcomeBonus             GetUmIncomeHistoryIncomeTypeParameter = "WELCOME_BONUS"
	GetUmIncomeHistoryIncomeTypeParameterRealizedPnl              GetUmIncomeHistoryIncomeTypeParameter = "REALIZED_PNL"
	GetUmIncomeHistoryIncomeTypeParameterFundingFee               GetUmIncomeHistoryIncomeTypeParameter = "FUNDING_FEE"
	GetUmIncomeHistoryIncomeTypeParameterCommission               GetUmIncomeHistoryIncomeTypeParameter = "COMMISSION"
	GetUmIncomeHistoryIncomeTypeParameterInsuranceClear           GetUmIncomeHistoryIncomeTypeParameter = "INSURANCE_CLEAR"
	GetUmIncomeHistoryIncomeTypeParameterReferralKickback         GetUmIncomeHistoryIncomeTypeParameter = "REFERRAL_KICKBACK"
	GetUmIncomeHistoryIncomeTypeParameterCommissionRebate         GetUmIncomeHistoryIncomeTypeParameter = "COMMISSION_REBATE"
	GetUmIncomeHistoryIncomeTypeParameterApiRebate                GetUmIncomeHistoryIncomeTypeParameter = "API_REBATE"
	GetUmIncomeHistoryIncomeTypeParameterContestReward            GetUmIncomeHistoryIncomeTypeParameter = "CONTEST_REWARD"
	GetUmIncomeHistoryIncomeTypeParameterCrossCollateralTransfer  GetUmIncomeHistoryIncomeTypeParameter = "CROSS_COLLATERAL_TRANSFER"
	GetUmIncomeHistoryIncomeTypeParameterOptionsPremiumFee        GetUmIncomeHistoryIncomeTypeParameter = "OPTIONS_PREMIUM_FEE"
	GetUmIncomeHistoryIncomeTypeParameterOptionsSettleProfit      GetUmIncomeHistoryIncomeTypeParameter = "OPTIONS_SETTLE_PROFIT"
	GetUmIncomeHistoryIncomeTypeParameterInternalTransfer         GetUmIncomeHistoryIncomeTypeParameter = "INTERNAL_TRANSFER"
	GetUmIncomeHistoryIncomeTypeParameterAutoExchange             GetUmIncomeHistoryIncomeTypeParameter = "AUTO_EXCHANGE"
	GetUmIncomeHistoryIncomeTypeParameterDeliveredSettelment      GetUmIncomeHistoryIncomeTypeParameter = "DELIVERED_SETTELMENT"
	GetUmIncomeHistoryIncomeTypeParameterCoinSwapDeposit          GetUmIncomeHistoryIncomeTypeParameter = "COIN_SWAP_DEPOSIT"
	GetUmIncomeHistoryIncomeTypeParameterCoinSwapWithdraw         GetUmIncomeHistoryIncomeTypeParameter = "COIN_SWAP_WITHDRAW"
	GetUmIncomeHistoryIncomeTypeParameterPositionLimitIncreaseFee GetUmIncomeHistoryIncomeTypeParameter = "POSITION_LIMIT_INCREASE_FEE"
)

// All allowed values of GetUmIncomeHistoryIncomeTypeParameter enum
var AllowedGetUmIncomeHistoryIncomeTypeParameterEnumValues = []GetUmIncomeHistoryIncomeTypeParameter{
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
}

func (v *GetUmIncomeHistoryIncomeTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetUmIncomeHistoryIncomeTypeParameter(value)
	for _, existing := range AllowedGetUmIncomeHistoryIncomeTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetUmIncomeHistoryIncomeTypeParameter", value)
}

// NewGetUmIncomeHistoryIncomeTypeParameterFromValue returns a pointer to a valid GetUmIncomeHistoryIncomeTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetUmIncomeHistoryIncomeTypeParameterFromValue(v string) (*GetUmIncomeHistoryIncomeTypeParameter, error) {
	ev := GetUmIncomeHistoryIncomeTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetUmIncomeHistoryIncomeTypeParameter: valid values are %v", v, AllowedGetUmIncomeHistoryIncomeTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetUmIncomeHistoryIncomeTypeParameter) IsValid() bool {
	for _, existing := range AllowedGetUmIncomeHistoryIncomeTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getUmIncomeHistory_incomeType_parameter value
func (v GetUmIncomeHistoryIncomeTypeParameter) Ptr() *GetUmIncomeHistoryIncomeTypeParameter {
	return &v
}

type NullableGetUmIncomeHistoryIncomeTypeParameter struct {
	value *GetUmIncomeHistoryIncomeTypeParameter
	isSet bool
}

func (v NullableGetUmIncomeHistoryIncomeTypeParameter) Get() *GetUmIncomeHistoryIncomeTypeParameter {
	return v.value
}

func (v *NullableGetUmIncomeHistoryIncomeTypeParameter) Set(val *GetUmIncomeHistoryIncomeTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUmIncomeHistoryIncomeTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUmIncomeHistoryIncomeTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUmIncomeHistoryIncomeTypeParameter(val *GetUmIncomeHistoryIncomeTypeParameter) *NullableGetUmIncomeHistoryIncomeTypeParameter {
	return &NullableGetUmIncomeHistoryIncomeTypeParameter{value: val, isSet: true}
}

func (v NullableGetUmIncomeHistoryIncomeTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUmIncomeHistoryIncomeTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
