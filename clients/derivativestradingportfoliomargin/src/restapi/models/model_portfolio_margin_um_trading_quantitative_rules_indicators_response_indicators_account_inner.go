/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner{}

// PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner struct for PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner
type PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner struct {
	Indicator            *string `json:"indicator,omitempty"`
	Value                *int64  `json:"value,omitempty"`
	TriggerValue         *int64  `json:"triggerValue,omitempty"`
	PlannedRecoverTime   *int64  `json:"plannedRecoverTime,omitempty"`
	IsLocked             *bool   `json:"isLocked,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner

// NewPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner instantiates a new PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner() *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner {
	this := PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner{}
	return &this
}

// NewPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInnerWithDefaults instantiates a new PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInnerWithDefaults() *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner {
	this := PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner{}
	return &this
}

// GetIndicator returns the Indicator field value if set, zero value otherwise.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetIndicator() string {
	if o == nil || common.IsNil(o.Indicator) {
		var ret string
		return ret
	}
	return *o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Indicator) {
		return nil, false
	}
	return o.Indicator, true
}

// HasIndicator returns a boolean if a field has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) HasIndicator() bool {
	if o != nil && !common.IsNil(o.Indicator) {
		return true
	}

	return false
}

// SetIndicator gets a reference to the given string and assigns it to the Indicator field.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) SetIndicator(v string) {
	o.Indicator = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetValue() int64 {
	if o == nil || common.IsNil(o.Value) {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetValueOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) SetValue(v int64) {
	o.Value = &v
}

// GetTriggerValue returns the TriggerValue field value if set, zero value otherwise.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetTriggerValue() int64 {
	if o == nil || common.IsNil(o.TriggerValue) {
		var ret int64
		return ret
	}
	return *o.TriggerValue
}

// GetTriggerValueOk returns a tuple with the TriggerValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetTriggerValueOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TriggerValue) {
		return nil, false
	}
	return o.TriggerValue, true
}

// HasTriggerValue returns a boolean if a field has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) HasTriggerValue() bool {
	if o != nil && !common.IsNil(o.TriggerValue) {
		return true
	}

	return false
}

// SetTriggerValue gets a reference to the given int64 and assigns it to the TriggerValue field.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) SetTriggerValue(v int64) {
	o.TriggerValue = &v
}

// GetPlannedRecoverTime returns the PlannedRecoverTime field value if set, zero value otherwise.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetPlannedRecoverTime() int64 {
	if o == nil || common.IsNil(o.PlannedRecoverTime) {
		var ret int64
		return ret
	}
	return *o.PlannedRecoverTime
}

// GetPlannedRecoverTimeOk returns a tuple with the PlannedRecoverTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetPlannedRecoverTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PlannedRecoverTime) {
		return nil, false
	}
	return o.PlannedRecoverTime, true
}

// HasPlannedRecoverTime returns a boolean if a field has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) HasPlannedRecoverTime() bool {
	if o != nil && !common.IsNil(o.PlannedRecoverTime) {
		return true
	}

	return false
}

// SetPlannedRecoverTime gets a reference to the given int64 and assigns it to the PlannedRecoverTime field.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) SetPlannedRecoverTime(v int64) {
	o.PlannedRecoverTime = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetIsLocked() bool {
	if o == nil || common.IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetIsLockedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) HasIsLocked() bool {
	if o != nil && !common.IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) SetIsLocked(v bool) {
	o.IsLocked = &v
}

func (o PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Indicator) {
		toSerialize["indicator"] = o.Indicator
	}
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !common.IsNil(o.TriggerValue) {
		toSerialize["triggerValue"] = o.TriggerValue
	}
	if !common.IsNil(o.PlannedRecoverTime) {
		toSerialize["plannedRecoverTime"] = o.PlannedRecoverTime
	}
	if !common.IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) UnmarshalJSON(data []byte) (err error) {
	varPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner := _PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner{}

	err = json.Unmarshal(data, &varPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner)

	if err != nil {
		return err
	}

	*o = PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner(varPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "indicator")
		delete(additionalProperties, "value")
		delete(additionalProperties, "triggerValue")
		delete(additionalProperties, "plannedRecoverTime")
		delete(additionalProperties, "isLocked")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner struct {
	value *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner
	isSet bool
}

func (v NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) Get() *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner {
	return v.value
}

func (v *NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) Set(val *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner(val *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) *NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner {
	return &NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner{value: val, isSet: true}
}

func (v NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
