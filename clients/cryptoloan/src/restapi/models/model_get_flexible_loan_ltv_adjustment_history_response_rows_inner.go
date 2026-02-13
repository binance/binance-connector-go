/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner{}

// GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner struct for GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner
type GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner struct {
	LoanCoin             *string `json:"loanCoin,omitempty"`
	CollateralCoin       *string `json:"collateralCoin,omitempty"`
	Direction            *string `json:"direction,omitempty"`
	CollateralAmount     *string `json:"collateralAmount,omitempty"`
	PreLTV               *string `json:"preLTV,omitempty"`
	AfterLTV             *string `json:"afterLTV,omitempty"`
	AdjustTime           *int64  `json:"adjustTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner

// NewGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner instantiates a new GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner() *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner {
	this := GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner{}
	return &this
}

// NewGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInnerWithDefaults instantiates a new GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInnerWithDefaults() *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner {
	this := GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner{}
	return &this
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) GetDirection() string {
	if o == nil || common.IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) GetDirectionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) HasDirection() bool {
	if o != nil && !common.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) SetDirection(v string) {
	o.Direction = &v
}

// GetCollateralAmount returns the CollateralAmount field value if set, zero value otherwise.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) GetCollateralAmount() string {
	if o == nil || common.IsNil(o.CollateralAmount) {
		var ret string
		return ret
	}
	return *o.CollateralAmount
}

// GetCollateralAmountOk returns a tuple with the CollateralAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) GetCollateralAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralAmount) {
		return nil, false
	}
	return o.CollateralAmount, true
}

// HasCollateralAmount returns a boolean if a field has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) HasCollateralAmount() bool {
	if o != nil && !common.IsNil(o.CollateralAmount) {
		return true
	}

	return false
}

// SetCollateralAmount gets a reference to the given string and assigns it to the CollateralAmount field.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) SetCollateralAmount(v string) {
	o.CollateralAmount = &v
}

// GetPreLTV returns the PreLTV field value if set, zero value otherwise.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) GetPreLTV() string {
	if o == nil || common.IsNil(o.PreLTV) {
		var ret string
		return ret
	}
	return *o.PreLTV
}

// GetPreLTVOk returns a tuple with the PreLTV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) GetPreLTVOk() (*string, bool) {
	if o == nil || common.IsNil(o.PreLTV) {
		return nil, false
	}
	return o.PreLTV, true
}

// HasPreLTV returns a boolean if a field has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) HasPreLTV() bool {
	if o != nil && !common.IsNil(o.PreLTV) {
		return true
	}

	return false
}

// SetPreLTV gets a reference to the given string and assigns it to the PreLTV field.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) SetPreLTV(v string) {
	o.PreLTV = &v
}

// GetAfterLTV returns the AfterLTV field value if set, zero value otherwise.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) GetAfterLTV() string {
	if o == nil || common.IsNil(o.AfterLTV) {
		var ret string
		return ret
	}
	return *o.AfterLTV
}

// GetAfterLTVOk returns a tuple with the AfterLTV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) GetAfterLTVOk() (*string, bool) {
	if o == nil || common.IsNil(o.AfterLTV) {
		return nil, false
	}
	return o.AfterLTV, true
}

// HasAfterLTV returns a boolean if a field has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) HasAfterLTV() bool {
	if o != nil && !common.IsNil(o.AfterLTV) {
		return true
	}

	return false
}

// SetAfterLTV gets a reference to the given string and assigns it to the AfterLTV field.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) SetAfterLTV(v string) {
	o.AfterLTV = &v
}

// GetAdjustTime returns the AdjustTime field value if set, zero value otherwise.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) GetAdjustTime() int64 {
	if o == nil || common.IsNil(o.AdjustTime) {
		var ret int64
		return ret
	}
	return *o.AdjustTime
}

// GetAdjustTimeOk returns a tuple with the AdjustTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) GetAdjustTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AdjustTime) {
		return nil, false
	}
	return o.AdjustTime, true
}

// HasAdjustTime returns a boolean if a field has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) HasAdjustTime() bool {
	if o != nil && !common.IsNil(o.AdjustTime) {
		return true
	}

	return false
}

// SetAdjustTime gets a reference to the given int64 and assigns it to the AdjustTime field.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) SetAdjustTime(v int64) {
	o.AdjustTime = &v
}

func (o GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.CollateralAmount) {
		toSerialize["collateralAmount"] = o.CollateralAmount
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner := _GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner(varGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "direction")
		delete(additionalProperties, "collateralAmount")
		delete(additionalProperties, "preLTV")
		delete(additionalProperties, "afterLTV")
		delete(additionalProperties, "adjustTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner struct {
	value *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) Get() *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) Set(val *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner(val *GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) *NullableGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner {
	return &NullableGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
