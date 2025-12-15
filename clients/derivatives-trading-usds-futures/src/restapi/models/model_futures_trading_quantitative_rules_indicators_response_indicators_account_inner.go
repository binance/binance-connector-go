/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner{}

// FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner struct for FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner
type FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner struct {
	Indicator            *string `json:"indicator,omitempty"`
	Value                *int64  `json:"value,omitempty"`
	TriggerValue         *int64  `json:"triggerValue,omitempty"`
	PlannedRecoverTime   *int64  `json:"plannedRecoverTime,omitempty"`
	IsLocked             *bool   `json:"isLocked,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner

// NewFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner instantiates a new FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner() *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner {
	this := FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner{}
	return &this
}

// NewFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInnerWithDefaults instantiates a new FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInnerWithDefaults() *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner {
	this := FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner{}
	return &this
}

// GetIndicator returns the Indicator field value if set, zero value otherwise.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetIndicator() string {
	if o == nil || common.IsNil(o.Indicator) {
		var ret string
		return ret
	}
	return *o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Indicator) {
		return nil, false
	}
	return o.Indicator, true
}

// HasIndicator returns a boolean if a field has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) HasIndicator() bool {
	if o != nil && !common.IsNil(o.Indicator) {
		return true
	}

	return false
}

// SetIndicator gets a reference to the given string and assigns it to the Indicator field.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) SetIndicator(v string) {
	o.Indicator = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetValue() int64 {
	if o == nil || common.IsNil(o.Value) {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetValueOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) SetValue(v int64) {
	o.Value = &v
}

// GetTriggerValue returns the TriggerValue field value if set, zero value otherwise.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetTriggerValue() int64 {
	if o == nil || common.IsNil(o.TriggerValue) {
		var ret int64
		return ret
	}
	return *o.TriggerValue
}

// GetTriggerValueOk returns a tuple with the TriggerValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetTriggerValueOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TriggerValue) {
		return nil, false
	}
	return o.TriggerValue, true
}

// HasTriggerValue returns a boolean if a field has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) HasTriggerValue() bool {
	if o != nil && !common.IsNil(o.TriggerValue) {
		return true
	}

	return false
}

// SetTriggerValue gets a reference to the given int64 and assigns it to the TriggerValue field.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) SetTriggerValue(v int64) {
	o.TriggerValue = &v
}

// GetPlannedRecoverTime returns the PlannedRecoverTime field value if set, zero value otherwise.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetPlannedRecoverTime() int64 {
	if o == nil || common.IsNil(o.PlannedRecoverTime) {
		var ret int64
		return ret
	}
	return *o.PlannedRecoverTime
}

// GetPlannedRecoverTimeOk returns a tuple with the PlannedRecoverTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetPlannedRecoverTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PlannedRecoverTime) {
		return nil, false
	}
	return o.PlannedRecoverTime, true
}

// HasPlannedRecoverTime returns a boolean if a field has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) HasPlannedRecoverTime() bool {
	if o != nil && !common.IsNil(o.PlannedRecoverTime) {
		return true
	}

	return false
}

// SetPlannedRecoverTime gets a reference to the given int64 and assigns it to the PlannedRecoverTime field.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) SetPlannedRecoverTime(v int64) {
	o.PlannedRecoverTime = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetIsLocked() bool {
	if o == nil || common.IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) GetIsLockedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) HasIsLocked() bool {
	if o != nil && !common.IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) SetIsLocked(v bool) {
	o.IsLocked = &v
}

func (o FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) ToMap() (map[string]interface{}, error) {
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

func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) UnmarshalJSON(data []byte) (err error) {
	varFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner := _FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner{}

	err = json.Unmarshal(data, &varFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner)

	if err != nil {
		return err
	}

	*o = FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner(varFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner)

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

type NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner struct {
	value *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner
	isSet bool
}

func (v NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) Get() *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner {
	return v.value
}

func (v *NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) Set(val *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner(val *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) *NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner {
	return &NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner{value: val, isSet: true}
}

func (v NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
