/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountApiTradingStatusResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountApiTradingStatusResponseData{}

// AccountApiTradingStatusResponseData struct for AccountApiTradingStatusResponseData
type AccountApiTradingStatusResponseData struct {
	IsLocked             *bool                                                `json:"isLocked,omitempty"`
	PlannedRecoverTime   *int64                                               `json:"plannedRecoverTime,omitempty"`
	TriggerCondition     *AccountApiTradingStatusResponseDataTriggerCondition `json:"triggerCondition,omitempty"`
	UpdateTime           *int64                                               `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountApiTradingStatusResponseData AccountApiTradingStatusResponseData

// NewAccountApiTradingStatusResponseData instantiates a new AccountApiTradingStatusResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountApiTradingStatusResponseData() *AccountApiTradingStatusResponseData {
	this := AccountApiTradingStatusResponseData{}
	return &this
}

// NewAccountApiTradingStatusResponseDataWithDefaults instantiates a new AccountApiTradingStatusResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountApiTradingStatusResponseDataWithDefaults() *AccountApiTradingStatusResponseData {
	this := AccountApiTradingStatusResponseData{}
	return &this
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *AccountApiTradingStatusResponseData) GetIsLocked() bool {
	if o == nil || common.IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountApiTradingStatusResponseData) GetIsLockedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *AccountApiTradingStatusResponseData) HasIsLocked() bool {
	if o != nil && !common.IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *AccountApiTradingStatusResponseData) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetPlannedRecoverTime returns the PlannedRecoverTime field value if set, zero value otherwise.
func (o *AccountApiTradingStatusResponseData) GetPlannedRecoverTime() int64 {
	if o == nil || common.IsNil(o.PlannedRecoverTime) {
		var ret int64
		return ret
	}
	return *o.PlannedRecoverTime
}

// GetPlannedRecoverTimeOk returns a tuple with the PlannedRecoverTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountApiTradingStatusResponseData) GetPlannedRecoverTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PlannedRecoverTime) {
		return nil, false
	}
	return o.PlannedRecoverTime, true
}

// HasPlannedRecoverTime returns a boolean if a field has been set.
func (o *AccountApiTradingStatusResponseData) HasPlannedRecoverTime() bool {
	if o != nil && !common.IsNil(o.PlannedRecoverTime) {
		return true
	}

	return false
}

// SetPlannedRecoverTime gets a reference to the given int64 and assigns it to the PlannedRecoverTime field.
func (o *AccountApiTradingStatusResponseData) SetPlannedRecoverTime(v int64) {
	o.PlannedRecoverTime = &v
}

// GetTriggerCondition returns the TriggerCondition field value if set, zero value otherwise.
func (o *AccountApiTradingStatusResponseData) GetTriggerCondition() AccountApiTradingStatusResponseDataTriggerCondition {
	if o == nil || common.IsNil(o.TriggerCondition) {
		var ret AccountApiTradingStatusResponseDataTriggerCondition
		return ret
	}
	return *o.TriggerCondition
}

// GetTriggerConditionOk returns a tuple with the TriggerCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountApiTradingStatusResponseData) GetTriggerConditionOk() (*AccountApiTradingStatusResponseDataTriggerCondition, bool) {
	if o == nil || common.IsNil(o.TriggerCondition) {
		return nil, false
	}
	return o.TriggerCondition, true
}

// HasTriggerCondition returns a boolean if a field has been set.
func (o *AccountApiTradingStatusResponseData) HasTriggerCondition() bool {
	if o != nil && !common.IsNil(o.TriggerCondition) {
		return true
	}

	return false
}

// SetTriggerCondition gets a reference to the given AccountApiTradingStatusResponseDataTriggerCondition and assigns it to the TriggerCondition field.
func (o *AccountApiTradingStatusResponseData) SetTriggerCondition(v AccountApiTradingStatusResponseDataTriggerCondition) {
	o.TriggerCondition = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AccountApiTradingStatusResponseData) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountApiTradingStatusResponseData) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AccountApiTradingStatusResponseData) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AccountApiTradingStatusResponseData) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o AccountApiTradingStatusResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountApiTradingStatusResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if !common.IsNil(o.PlannedRecoverTime) {
		toSerialize["plannedRecoverTime"] = o.PlannedRecoverTime
	}
	if !common.IsNil(o.TriggerCondition) {
		toSerialize["triggerCondition"] = o.TriggerCondition
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountApiTradingStatusResponseData) UnmarshalJSON(data []byte) (err error) {
	varAccountApiTradingStatusResponseData := _AccountApiTradingStatusResponseData{}

	err = json.Unmarshal(data, &varAccountApiTradingStatusResponseData)

	if err != nil {
		return err
	}

	*o = AccountApiTradingStatusResponseData(varAccountApiTradingStatusResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "isLocked")
		delete(additionalProperties, "plannedRecoverTime")
		delete(additionalProperties, "triggerCondition")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountApiTradingStatusResponseData struct {
	value *AccountApiTradingStatusResponseData
	isSet bool
}

func (v NullableAccountApiTradingStatusResponseData) Get() *AccountApiTradingStatusResponseData {
	return v.value
}

func (v *NullableAccountApiTradingStatusResponseData) Set(val *AccountApiTradingStatusResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountApiTradingStatusResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountApiTradingStatusResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountApiTradingStatusResponseData(val *AccountApiTradingStatusResponseData) *NullableAccountApiTradingStatusResponseData {
	return &NullableAccountApiTradingStatusResponseData{value: val, isSet: true}
}

func (v NullableAccountApiTradingStatusResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountApiTradingStatusResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
