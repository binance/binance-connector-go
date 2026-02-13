/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner{}

// QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner struct for QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner
type QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner struct {
	Leverage              *int64   `json:"leverage,omitempty"`
	MaxDebt               *float32 `json:"maxDebt,omitempty"`
	MaintenanceMarginRate *float32 `json:"maintenanceMarginRate,omitempty"`
	InitialMarginRate     *float32 `json:"initialMarginRate,omitempty"`
	FastNum               *float32 `json:"fastNum,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner

// NewQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner instantiates a new QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner() *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner {
	this := QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner{}
	return &this
}

// NewQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInnerWithDefaults instantiates a new QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInnerWithDefaults() *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner {
	this := QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner{}
	return &this
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) GetLeverage() int64 {
	if o == nil || common.IsNil(o.Leverage) {
		var ret int64
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) GetLeverageOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given int64 and assigns it to the Leverage field.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) SetLeverage(v int64) {
	o.Leverage = &v
}

// GetMaxDebt returns the MaxDebt field value if set, zero value otherwise.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) GetMaxDebt() float32 {
	if o == nil || common.IsNil(o.MaxDebt) {
		var ret float32
		return ret
	}
	return *o.MaxDebt
}

// GetMaxDebtOk returns a tuple with the MaxDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) GetMaxDebtOk() (*float32, bool) {
	if o == nil || common.IsNil(o.MaxDebt) {
		return nil, false
	}
	return o.MaxDebt, true
}

// HasMaxDebt returns a boolean if a field has been set.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) HasMaxDebt() bool {
	if o != nil && !common.IsNil(o.MaxDebt) {
		return true
	}

	return false
}

// SetMaxDebt gets a reference to the given float32 and assigns it to the MaxDebt field.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) SetMaxDebt(v float32) {
	o.MaxDebt = &v
}

// GetMaintenanceMarginRate returns the MaintenanceMarginRate field value if set, zero value otherwise.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) GetMaintenanceMarginRate() float32 {
	if o == nil || common.IsNil(o.MaintenanceMarginRate) {
		var ret float32
		return ret
	}
	return *o.MaintenanceMarginRate
}

// GetMaintenanceMarginRateOk returns a tuple with the MaintenanceMarginRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) GetMaintenanceMarginRateOk() (*float32, bool) {
	if o == nil || common.IsNil(o.MaintenanceMarginRate) {
		return nil, false
	}
	return o.MaintenanceMarginRate, true
}

// HasMaintenanceMarginRate returns a boolean if a field has been set.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) HasMaintenanceMarginRate() bool {
	if o != nil && !common.IsNil(o.MaintenanceMarginRate) {
		return true
	}

	return false
}

// SetMaintenanceMarginRate gets a reference to the given float32 and assigns it to the MaintenanceMarginRate field.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) SetMaintenanceMarginRate(v float32) {
	o.MaintenanceMarginRate = &v
}

// GetInitialMarginRate returns the InitialMarginRate field value if set, zero value otherwise.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) GetInitialMarginRate() float32 {
	if o == nil || common.IsNil(o.InitialMarginRate) {
		var ret float32
		return ret
	}
	return *o.InitialMarginRate
}

// GetInitialMarginRateOk returns a tuple with the InitialMarginRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) GetInitialMarginRateOk() (*float32, bool) {
	if o == nil || common.IsNil(o.InitialMarginRate) {
		return nil, false
	}
	return o.InitialMarginRate, true
}

// HasInitialMarginRate returns a boolean if a field has been set.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) HasInitialMarginRate() bool {
	if o != nil && !common.IsNil(o.InitialMarginRate) {
		return true
	}

	return false
}

// SetInitialMarginRate gets a reference to the given float32 and assigns it to the InitialMarginRate field.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) SetInitialMarginRate(v float32) {
	o.InitialMarginRate = &v
}

// GetFastNum returns the FastNum field value if set, zero value otherwise.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) GetFastNum() float32 {
	if o == nil || common.IsNil(o.FastNum) {
		var ret float32
		return ret
	}
	return *o.FastNum
}

// GetFastNumOk returns a tuple with the FastNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) GetFastNumOk() (*float32, bool) {
	if o == nil || common.IsNil(o.FastNum) {
		return nil, false
	}
	return o.FastNum, true
}

// HasFastNum returns a boolean if a field has been set.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) HasFastNum() bool {
	if o != nil && !common.IsNil(o.FastNum) {
		return true
	}

	return false
}

// SetFastNum gets a reference to the given float32 and assigns it to the FastNum field.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) SetFastNum(v float32) {
	o.FastNum = &v
}

func (o QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}
	if !common.IsNil(o.MaxDebt) {
		toSerialize["maxDebt"] = o.MaxDebt
	}
	if !common.IsNil(o.MaintenanceMarginRate) {
		toSerialize["maintenanceMarginRate"] = o.MaintenanceMarginRate
	}
	if !common.IsNil(o.InitialMarginRate) {
		toSerialize["initialMarginRate"] = o.InitialMarginRate
	}
	if !common.IsNil(o.FastNum) {
		toSerialize["fastNum"] = o.FastNum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner := _QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner{}

	err = json.Unmarshal(data, &varQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner)

	if err != nil {
		return err
	}

	*o = QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner(varQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "leverage")
		delete(additionalProperties, "maxDebt")
		delete(additionalProperties, "maintenanceMarginRate")
		delete(additionalProperties, "initialMarginRate")
		delete(additionalProperties, "fastNum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner struct {
	value *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner
	isSet bool
}

func (v NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) Get() *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner {
	return v.value
}

func (v *NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) Set(val *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner(val *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) *NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner {
	return &NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner{value: val, isSet: true}
}

func (v NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
