/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFlexibleLoanLiquidationHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleLoanLiquidationHistoryResponseRowsInner{}

// GetFlexibleLoanLiquidationHistoryResponseRowsInner struct for GetFlexibleLoanLiquidationHistoryResponseRowsInner
type GetFlexibleLoanLiquidationHistoryResponseRowsInner struct {
	LoanCoin                    *string `json:"loanCoin,omitempty"`
	LiquidationDebt             *string `json:"liquidationDebt,omitempty"`
	CollateralCoin              *string `json:"collateralCoin,omitempty"`
	LiquidationCollateralAmount *string `json:"liquidationCollateralAmount,omitempty"`
	ReturnCollateralAmount      *string `json:"returnCollateralAmount,omitempty"`
	LiquidationFee              *string `json:"liquidationFee,omitempty"`
	LiquidationStartingPrice    *string `json:"liquidationStartingPrice,omitempty"`
	LiquidationStartingTime     *int64  `json:"liquidationStartingTime,omitempty"`
	Status                      *string `json:"status,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _GetFlexibleLoanLiquidationHistoryResponseRowsInner GetFlexibleLoanLiquidationHistoryResponseRowsInner

// NewGetFlexibleLoanLiquidationHistoryResponseRowsInner instantiates a new GetFlexibleLoanLiquidationHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleLoanLiquidationHistoryResponseRowsInner() *GetFlexibleLoanLiquidationHistoryResponseRowsInner {
	this := GetFlexibleLoanLiquidationHistoryResponseRowsInner{}
	return &this
}

// NewGetFlexibleLoanLiquidationHistoryResponseRowsInnerWithDefaults instantiates a new GetFlexibleLoanLiquidationHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleLoanLiquidationHistoryResponseRowsInnerWithDefaults() *GetFlexibleLoanLiquidationHistoryResponseRowsInner {
	this := GetFlexibleLoanLiquidationHistoryResponseRowsInner{}
	return &this
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetLiquidationDebt returns the LiquidationDebt field value if set, zero value otherwise.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetLiquidationDebt() string {
	if o == nil || common.IsNil(o.LiquidationDebt) {
		var ret string
		return ret
	}
	return *o.LiquidationDebt
}

// GetLiquidationDebtOk returns a tuple with the LiquidationDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetLiquidationDebtOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiquidationDebt) {
		return nil, false
	}
	return o.LiquidationDebt, true
}

// HasLiquidationDebt returns a boolean if a field has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) HasLiquidationDebt() bool {
	if o != nil && !common.IsNil(o.LiquidationDebt) {
		return true
	}

	return false
}

// SetLiquidationDebt gets a reference to the given string and assigns it to the LiquidationDebt field.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) SetLiquidationDebt(v string) {
	o.LiquidationDebt = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetLiquidationCollateralAmount returns the LiquidationCollateralAmount field value if set, zero value otherwise.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetLiquidationCollateralAmount() string {
	if o == nil || common.IsNil(o.LiquidationCollateralAmount) {
		var ret string
		return ret
	}
	return *o.LiquidationCollateralAmount
}

// GetLiquidationCollateralAmountOk returns a tuple with the LiquidationCollateralAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetLiquidationCollateralAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiquidationCollateralAmount) {
		return nil, false
	}
	return o.LiquidationCollateralAmount, true
}

// HasLiquidationCollateralAmount returns a boolean if a field has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) HasLiquidationCollateralAmount() bool {
	if o != nil && !common.IsNil(o.LiquidationCollateralAmount) {
		return true
	}

	return false
}

// SetLiquidationCollateralAmount gets a reference to the given string and assigns it to the LiquidationCollateralAmount field.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) SetLiquidationCollateralAmount(v string) {
	o.LiquidationCollateralAmount = &v
}

// GetReturnCollateralAmount returns the ReturnCollateralAmount field value if set, zero value otherwise.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetReturnCollateralAmount() string {
	if o == nil || common.IsNil(o.ReturnCollateralAmount) {
		var ret string
		return ret
	}
	return *o.ReturnCollateralAmount
}

// GetReturnCollateralAmountOk returns a tuple with the ReturnCollateralAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetReturnCollateralAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReturnCollateralAmount) {
		return nil, false
	}
	return o.ReturnCollateralAmount, true
}

// HasReturnCollateralAmount returns a boolean if a field has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) HasReturnCollateralAmount() bool {
	if o != nil && !common.IsNil(o.ReturnCollateralAmount) {
		return true
	}

	return false
}

// SetReturnCollateralAmount gets a reference to the given string and assigns it to the ReturnCollateralAmount field.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) SetReturnCollateralAmount(v string) {
	o.ReturnCollateralAmount = &v
}

// GetLiquidationFee returns the LiquidationFee field value if set, zero value otherwise.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetLiquidationFee() string {
	if o == nil || common.IsNil(o.LiquidationFee) {
		var ret string
		return ret
	}
	return *o.LiquidationFee
}

// GetLiquidationFeeOk returns a tuple with the LiquidationFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetLiquidationFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiquidationFee) {
		return nil, false
	}
	return o.LiquidationFee, true
}

// HasLiquidationFee returns a boolean if a field has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) HasLiquidationFee() bool {
	if o != nil && !common.IsNil(o.LiquidationFee) {
		return true
	}

	return false
}

// SetLiquidationFee gets a reference to the given string and assigns it to the LiquidationFee field.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) SetLiquidationFee(v string) {
	o.LiquidationFee = &v
}

// GetLiquidationStartingPrice returns the LiquidationStartingPrice field value if set, zero value otherwise.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetLiquidationStartingPrice() string {
	if o == nil || common.IsNil(o.LiquidationStartingPrice) {
		var ret string
		return ret
	}
	return *o.LiquidationStartingPrice
}

// GetLiquidationStartingPriceOk returns a tuple with the LiquidationStartingPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetLiquidationStartingPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiquidationStartingPrice) {
		return nil, false
	}
	return o.LiquidationStartingPrice, true
}

// HasLiquidationStartingPrice returns a boolean if a field has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) HasLiquidationStartingPrice() bool {
	if o != nil && !common.IsNil(o.LiquidationStartingPrice) {
		return true
	}

	return false
}

// SetLiquidationStartingPrice gets a reference to the given string and assigns it to the LiquidationStartingPrice field.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) SetLiquidationStartingPrice(v string) {
	o.LiquidationStartingPrice = &v
}

// GetLiquidationStartingTime returns the LiquidationStartingTime field value if set, zero value otherwise.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetLiquidationStartingTime() int64 {
	if o == nil || common.IsNil(o.LiquidationStartingTime) {
		var ret int64
		return ret
	}
	return *o.LiquidationStartingTime
}

// GetLiquidationStartingTimeOk returns a tuple with the LiquidationStartingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetLiquidationStartingTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LiquidationStartingTime) {
		return nil, false
	}
	return o.LiquidationStartingTime, true
}

// HasLiquidationStartingTime returns a boolean if a field has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) HasLiquidationStartingTime() bool {
	if o != nil && !common.IsNil(o.LiquidationStartingTime) {
		return true
	}

	return false
}

// SetLiquidationStartingTime gets a reference to the given int64 and assigns it to the LiquidationStartingTime field.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) SetLiquidationStartingTime(v int64) {
	o.LiquidationStartingTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetFlexibleLoanLiquidationHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleLoanLiquidationHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.LiquidationDebt) {
		toSerialize["liquidationDebt"] = o.LiquidationDebt
	}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !common.IsNil(o.LiquidationCollateralAmount) {
		toSerialize["liquidationCollateralAmount"] = o.LiquidationCollateralAmount
	}
	if !common.IsNil(o.ReturnCollateralAmount) {
		toSerialize["returnCollateralAmount"] = o.ReturnCollateralAmount
	}
	if !common.IsNil(o.LiquidationFee) {
		toSerialize["liquidationFee"] = o.LiquidationFee
	}
	if !common.IsNil(o.LiquidationStartingPrice) {
		toSerialize["liquidationStartingPrice"] = o.LiquidationStartingPrice
	}
	if !common.IsNil(o.LiquidationStartingTime) {
		toSerialize["liquidationStartingTime"] = o.LiquidationStartingTime
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFlexibleLoanLiquidationHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleLoanLiquidationHistoryResponseRowsInner := _GetFlexibleLoanLiquidationHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetFlexibleLoanLiquidationHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetFlexibleLoanLiquidationHistoryResponseRowsInner(varGetFlexibleLoanLiquidationHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "liquidationDebt")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "liquidationCollateralAmount")
		delete(additionalProperties, "returnCollateralAmount")
		delete(additionalProperties, "liquidationFee")
		delete(additionalProperties, "liquidationStartingPrice")
		delete(additionalProperties, "liquidationStartingTime")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleLoanLiquidationHistoryResponseRowsInner struct {
	value *GetFlexibleLoanLiquidationHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetFlexibleLoanLiquidationHistoryResponseRowsInner) Get() *GetFlexibleLoanLiquidationHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetFlexibleLoanLiquidationHistoryResponseRowsInner) Set(val *GetFlexibleLoanLiquidationHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleLoanLiquidationHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleLoanLiquidationHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleLoanLiquidationHistoryResponseRowsInner(val *GetFlexibleLoanLiquidationHistoryResponseRowsInner) *NullableGetFlexibleLoanLiquidationHistoryResponseRowsInner {
	return &NullableGetFlexibleLoanLiquidationHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetFlexibleLoanLiquidationHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleLoanLiquidationHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
