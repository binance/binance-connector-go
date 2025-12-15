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

// checks if the GetLoanLtvAdjustmentHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLoanLtvAdjustmentHistoryResponseRowsInner{}

// GetLoanLtvAdjustmentHistoryResponseRowsInner struct for GetLoanLtvAdjustmentHistoryResponseRowsInner
type GetLoanLtvAdjustmentHistoryResponseRowsInner struct {
	LoanCoin             *string `json:"loanCoin,omitempty"`
	CollateralCoin       *string `json:"collateralCoin,omitempty"`
	Direction            *string `json:"direction,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	PreLTV               *string `json:"preLTV,omitempty"`
	AfterLTV             *string `json:"afterLTV,omitempty"`
	AdjustTime           *int64  `json:"adjustTime,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLoanLtvAdjustmentHistoryResponseRowsInner GetLoanLtvAdjustmentHistoryResponseRowsInner

// NewGetLoanLtvAdjustmentHistoryResponseRowsInner instantiates a new GetLoanLtvAdjustmentHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLoanLtvAdjustmentHistoryResponseRowsInner() *GetLoanLtvAdjustmentHistoryResponseRowsInner {
	this := GetLoanLtvAdjustmentHistoryResponseRowsInner{}
	return &this
}

// NewGetLoanLtvAdjustmentHistoryResponseRowsInnerWithDefaults instantiates a new GetLoanLtvAdjustmentHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLoanLtvAdjustmentHistoryResponseRowsInnerWithDefaults() *GetLoanLtvAdjustmentHistoryResponseRowsInner {
	this := GetLoanLtvAdjustmentHistoryResponseRowsInner{}
	return &this
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) GetDirection() string {
	if o == nil || common.IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) GetDirectionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) HasDirection() bool {
	if o != nil && !common.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) SetDirection(v string) {
	o.Direction = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetPreLTV returns the PreLTV field value if set, zero value otherwise.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) GetPreLTV() string {
	if o == nil || common.IsNil(o.PreLTV) {
		var ret string
		return ret
	}
	return *o.PreLTV
}

// GetPreLTVOk returns a tuple with the PreLTV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) GetPreLTVOk() (*string, bool) {
	if o == nil || common.IsNil(o.PreLTV) {
		return nil, false
	}
	return o.PreLTV, true
}

// HasPreLTV returns a boolean if a field has been set.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) HasPreLTV() bool {
	if o != nil && !common.IsNil(o.PreLTV) {
		return true
	}

	return false
}

// SetPreLTV gets a reference to the given string and assigns it to the PreLTV field.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) SetPreLTV(v string) {
	o.PreLTV = &v
}

// GetAfterLTV returns the AfterLTV field value if set, zero value otherwise.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) GetAfterLTV() string {
	if o == nil || common.IsNil(o.AfterLTV) {
		var ret string
		return ret
	}
	return *o.AfterLTV
}

// GetAfterLTVOk returns a tuple with the AfterLTV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) GetAfterLTVOk() (*string, bool) {
	if o == nil || common.IsNil(o.AfterLTV) {
		return nil, false
	}
	return o.AfterLTV, true
}

// HasAfterLTV returns a boolean if a field has been set.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) HasAfterLTV() bool {
	if o != nil && !common.IsNil(o.AfterLTV) {
		return true
	}

	return false
}

// SetAfterLTV gets a reference to the given string and assigns it to the AfterLTV field.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) SetAfterLTV(v string) {
	o.AfterLTV = &v
}

// GetAdjustTime returns the AdjustTime field value if set, zero value otherwise.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) GetAdjustTime() int64 {
	if o == nil || common.IsNil(o.AdjustTime) {
		var ret int64
		return ret
	}
	return *o.AdjustTime
}

// GetAdjustTimeOk returns a tuple with the AdjustTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) GetAdjustTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AdjustTime) {
		return nil, false
	}
	return o.AdjustTime, true
}

// HasAdjustTime returns a boolean if a field has been set.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) HasAdjustTime() bool {
	if o != nil && !common.IsNil(o.AdjustTime) {
		return true
	}

	return false
}

// SetAdjustTime gets a reference to the given int64 and assigns it to the AdjustTime field.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) SetAdjustTime(v int64) {
	o.AdjustTime = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) SetOrderId(v int64) {
	o.OrderId = &v
}

func (o GetLoanLtvAdjustmentHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLoanLtvAdjustmentHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !common.IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.PreLTV) {
		toSerialize["preLTV"] = o.PreLTV
	}
	if !common.IsNil(o.AfterLTV) {
		toSerialize["afterLTV"] = o.AfterLTV
	}
	if !common.IsNil(o.AdjustTime) {
		toSerialize["adjustTime"] = o.AdjustTime
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetLoanLtvAdjustmentHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetLoanLtvAdjustmentHistoryResponseRowsInner := _GetLoanLtvAdjustmentHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetLoanLtvAdjustmentHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetLoanLtvAdjustmentHistoryResponseRowsInner(varGetLoanLtvAdjustmentHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "direction")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "preLTV")
		delete(additionalProperties, "afterLTV")
		delete(additionalProperties, "adjustTime")
		delete(additionalProperties, "orderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLoanLtvAdjustmentHistoryResponseRowsInner struct {
	value *GetLoanLtvAdjustmentHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetLoanLtvAdjustmentHistoryResponseRowsInner) Get() *GetLoanLtvAdjustmentHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetLoanLtvAdjustmentHistoryResponseRowsInner) Set(val *GetLoanLtvAdjustmentHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLoanLtvAdjustmentHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLoanLtvAdjustmentHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLoanLtvAdjustmentHistoryResponseRowsInner(val *GetLoanLtvAdjustmentHistoryResponseRowsInner) *NullableGetLoanLtvAdjustmentHistoryResponseRowsInner {
	return &NullableGetLoanLtvAdjustmentHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetLoanLtvAdjustmentHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLoanLtvAdjustmentHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
