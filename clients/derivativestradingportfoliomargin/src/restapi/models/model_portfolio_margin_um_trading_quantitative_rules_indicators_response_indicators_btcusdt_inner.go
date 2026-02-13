/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner{}

// PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner struct for PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner
type PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner struct {
	IsLocked             *bool    `json:"isLocked,omitempty"`
	PlannedRecoverTime   *int64   `json:"plannedRecoverTime,omitempty"`
	Indicator            *string  `json:"indicator,omitempty"`
	Value                *float32 `json:"value,omitempty"`
	TriggerValue         *float32 `json:"triggerValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner

// NewPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner instantiates a new PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner() *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner {
	this := PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner{}
	return &this
}

// NewPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInnerWithDefaults instantiates a new PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInnerWithDefaults() *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner {
	this := PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner{}
	return &this
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetIsLocked() bool {
	if o == nil || common.IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetIsLockedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) HasIsLocked() bool {
	if o != nil && !common.IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetPlannedRecoverTime returns the PlannedRecoverTime field value if set, zero value otherwise.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetPlannedRecoverTime() int64 {
	if o == nil || common.IsNil(o.PlannedRecoverTime) {
		var ret int64
		return ret
	}
	return *o.PlannedRecoverTime
}

// GetPlannedRecoverTimeOk returns a tuple with the PlannedRecoverTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetPlannedRecoverTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PlannedRecoverTime) {
		return nil, false
	}
	return o.PlannedRecoverTime, true
}

// HasPlannedRecoverTime returns a boolean if a field has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) HasPlannedRecoverTime() bool {
	if o != nil && !common.IsNil(o.PlannedRecoverTime) {
		return true
	}

	return false
}

// SetPlannedRecoverTime gets a reference to the given int64 and assigns it to the PlannedRecoverTime field.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) SetPlannedRecoverTime(v int64) {
	o.PlannedRecoverTime = &v
}

// GetIndicator returns the Indicator field value if set, zero value otherwise.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetIndicator() string {
	if o == nil || common.IsNil(o.Indicator) {
		var ret string
		return ret
	}
	return *o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Indicator) {
		return nil, false
	}
	return o.Indicator, true
}

// HasIndicator returns a boolean if a field has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) HasIndicator() bool {
	if o != nil && !common.IsNil(o.Indicator) {
		return true
	}

	return false
}

// SetIndicator gets a reference to the given string and assigns it to the Indicator field.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) SetIndicator(v string) {
	o.Indicator = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetValue() float32 {
	if o == nil || common.IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetValueOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) SetValue(v float32) {
	o.Value = &v
}

// GetTriggerValue returns the TriggerValue field value if set, zero value otherwise.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetTriggerValue() float32 {
	if o == nil || common.IsNil(o.TriggerValue) {
		var ret float32
		return ret
	}
	return *o.TriggerValue
}

// GetTriggerValueOk returns a tuple with the TriggerValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetTriggerValueOk() (*float32, bool) {
	if o == nil || common.IsNil(o.TriggerValue) {
		return nil, false
	}
	return o.TriggerValue, true
}

// HasTriggerValue returns a boolean if a field has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) HasTriggerValue() bool {
	if o != nil && !common.IsNil(o.TriggerValue) {
		return true
	}

	return false
}

// SetTriggerValue gets a reference to the given float32 and assigns it to the TriggerValue field.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) SetTriggerValue(v float32) {
	o.TriggerValue = &v
}

func (o PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if !common.IsNil(o.PlannedRecoverTime) {
		toSerialize["plannedRecoverTime"] = o.PlannedRecoverTime
	}
	if !common.IsNil(o.Indicator) {
		toSerialize["indicator"] = o.Indicator
	}
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !common.IsNil(o.TriggerValue) {
		toSerialize["triggerValue"] = o.TriggerValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) UnmarshalJSON(data []byte) (err error) {
	varPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner := _PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner{}

	err = json.Unmarshal(data, &varPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner)

	if err != nil {
		return err
	}

	*o = PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner(varPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "isLocked")
		delete(additionalProperties, "plannedRecoverTime")
		delete(additionalProperties, "indicator")
		delete(additionalProperties, "value")
		delete(additionalProperties, "triggerValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner struct {
	value *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner
	isSet bool
}

func (v NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) Get() *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner {
	return v.value
}

func (v *NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) Set(val *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner(val *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) *NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner {
	return &NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner{value: val, isSet: true}
}

func (v NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
