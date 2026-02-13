/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner{}

// FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner struct for FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner
type FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner struct {
	IsLocked             *bool    `json:"isLocked,omitempty"`
	PlannedRecoverTime   *int64   `json:"plannedRecoverTime,omitempty"`
	Indicator            *string  `json:"indicator,omitempty"`
	Value                *float32 `json:"value,omitempty"`
	TriggerValue         *float32 `json:"triggerValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner

// NewFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner instantiates a new FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner() *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner {
	this := FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner{}
	return &this
}

// NewFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInnerWithDefaults instantiates a new FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInnerWithDefaults() *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner {
	this := FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner{}
	return &this
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetIsLocked() bool {
	if o == nil || common.IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetIsLockedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) HasIsLocked() bool {
	if o != nil && !common.IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetPlannedRecoverTime returns the PlannedRecoverTime field value if set, zero value otherwise.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetPlannedRecoverTime() int64 {
	if o == nil || common.IsNil(o.PlannedRecoverTime) {
		var ret int64
		return ret
	}
	return *o.PlannedRecoverTime
}

// GetPlannedRecoverTimeOk returns a tuple with the PlannedRecoverTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetPlannedRecoverTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PlannedRecoverTime) {
		return nil, false
	}
	return o.PlannedRecoverTime, true
}

// HasPlannedRecoverTime returns a boolean if a field has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) HasPlannedRecoverTime() bool {
	if o != nil && !common.IsNil(o.PlannedRecoverTime) {
		return true
	}

	return false
}

// SetPlannedRecoverTime gets a reference to the given int64 and assigns it to the PlannedRecoverTime field.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) SetPlannedRecoverTime(v int64) {
	o.PlannedRecoverTime = &v
}

// GetIndicator returns the Indicator field value if set, zero value otherwise.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetIndicator() string {
	if o == nil || common.IsNil(o.Indicator) {
		var ret string
		return ret
	}
	return *o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Indicator) {
		return nil, false
	}
	return o.Indicator, true
}

// HasIndicator returns a boolean if a field has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) HasIndicator() bool {
	if o != nil && !common.IsNil(o.Indicator) {
		return true
	}

	return false
}

// SetIndicator gets a reference to the given string and assigns it to the Indicator field.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) SetIndicator(v string) {
	o.Indicator = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetValue() float32 {
	if o == nil || common.IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetValueOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) SetValue(v float32) {
	o.Value = &v
}

// GetTriggerValue returns the TriggerValue field value if set, zero value otherwise.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetTriggerValue() float32 {
	if o == nil || common.IsNil(o.TriggerValue) {
		var ret float32
		return ret
	}
	return *o.TriggerValue
}

// GetTriggerValueOk returns a tuple with the TriggerValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) GetTriggerValueOk() (*float32, bool) {
	if o == nil || common.IsNil(o.TriggerValue) {
		return nil, false
	}
	return o.TriggerValue, true
}

// HasTriggerValue returns a boolean if a field has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) HasTriggerValue() bool {
	if o != nil && !common.IsNil(o.TriggerValue) {
		return true
	}

	return false
}

// SetTriggerValue gets a reference to the given float32 and assigns it to the TriggerValue field.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) SetTriggerValue(v float32) {
	o.TriggerValue = &v
}

func (o FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) ToMap() (map[string]interface{}, error) {
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

func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) UnmarshalJSON(data []byte) (err error) {
	varFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner := _FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner{}

	err = json.Unmarshal(data, &varFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner)

	if err != nil {
		return err
	}

	*o = FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner(varFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner)

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

type NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner struct {
	value *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner
	isSet bool
}

func (v NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) Get() *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner {
	return v.value
}

func (v *NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) Set(val *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner(val *FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) *NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner {
	return &NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner{value: val, isSet: true}
}

func (v NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
