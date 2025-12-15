/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountApiTradingStatusResponseDataTriggerCondition type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountApiTradingStatusResponseDataTriggerCondition{}

// AccountApiTradingStatusResponseDataTriggerCondition struct for AccountApiTradingStatusResponseDataTriggerCondition
type AccountApiTradingStatusResponseDataTriggerCondition struct {
	GCR                  *int64 `json:"GCR,omitempty"`
	IFER                 *int64 `json:"IFER,omitempty"`
	UFR                  *int64 `json:"UFR,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountApiTradingStatusResponseDataTriggerCondition AccountApiTradingStatusResponseDataTriggerCondition

// NewAccountApiTradingStatusResponseDataTriggerCondition instantiates a new AccountApiTradingStatusResponseDataTriggerCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountApiTradingStatusResponseDataTriggerCondition() *AccountApiTradingStatusResponseDataTriggerCondition {
	this := AccountApiTradingStatusResponseDataTriggerCondition{}
	return &this
}

// NewAccountApiTradingStatusResponseDataTriggerConditionWithDefaults instantiates a new AccountApiTradingStatusResponseDataTriggerCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountApiTradingStatusResponseDataTriggerConditionWithDefaults() *AccountApiTradingStatusResponseDataTriggerCondition {
	this := AccountApiTradingStatusResponseDataTriggerCondition{}
	return &this
}

// GetGCR returns the GCR field value if set, zero value otherwise.
func (o *AccountApiTradingStatusResponseDataTriggerCondition) GetGCR() int64 {
	if o == nil || common.IsNil(o.GCR) {
		var ret int64
		return ret
	}
	return *o.GCR
}

// GetGCROk returns a tuple with the GCR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountApiTradingStatusResponseDataTriggerCondition) GetGCROk() (*int64, bool) {
	if o == nil || common.IsNil(o.GCR) {
		return nil, false
	}
	return o.GCR, true
}

// HasGCR returns a boolean if a field has been set.
func (o *AccountApiTradingStatusResponseDataTriggerCondition) HasGCR() bool {
	if o != nil && !common.IsNil(o.GCR) {
		return true
	}

	return false
}

// SetGCR gets a reference to the given int64 and assigns it to the GCR field.
func (o *AccountApiTradingStatusResponseDataTriggerCondition) SetGCR(v int64) {
	o.GCR = &v
}

// GetIFER returns the IFER field value if set, zero value otherwise.
func (o *AccountApiTradingStatusResponseDataTriggerCondition) GetIFER() int64 {
	if o == nil || common.IsNil(o.IFER) {
		var ret int64
		return ret
	}
	return *o.IFER
}

// GetIFEROk returns a tuple with the IFER field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountApiTradingStatusResponseDataTriggerCondition) GetIFEROk() (*int64, bool) {
	if o == nil || common.IsNil(o.IFER) {
		return nil, false
	}
	return o.IFER, true
}

// HasIFER returns a boolean if a field has been set.
func (o *AccountApiTradingStatusResponseDataTriggerCondition) HasIFER() bool {
	if o != nil && !common.IsNil(o.IFER) {
		return true
	}

	return false
}

// SetIFER gets a reference to the given int64 and assigns it to the IFER field.
func (o *AccountApiTradingStatusResponseDataTriggerCondition) SetIFER(v int64) {
	o.IFER = &v
}

// GetUFR returns the UFR field value if set, zero value otherwise.
func (o *AccountApiTradingStatusResponseDataTriggerCondition) GetUFR() int64 {
	if o == nil || common.IsNil(o.UFR) {
		var ret int64
		return ret
	}
	return *o.UFR
}

// GetUFROk returns a tuple with the UFR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountApiTradingStatusResponseDataTriggerCondition) GetUFROk() (*int64, bool) {
	if o == nil || common.IsNil(o.UFR) {
		return nil, false
	}
	return o.UFR, true
}

// HasUFR returns a boolean if a field has been set.
func (o *AccountApiTradingStatusResponseDataTriggerCondition) HasUFR() bool {
	if o != nil && !common.IsNil(o.UFR) {
		return true
	}

	return false
}

// SetUFR gets a reference to the given int64 and assigns it to the UFR field.
func (o *AccountApiTradingStatusResponseDataTriggerCondition) SetUFR(v int64) {
	o.UFR = &v
}

func (o AccountApiTradingStatusResponseDataTriggerCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountApiTradingStatusResponseDataTriggerCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.GCR) {
		toSerialize["GCR"] = o.GCR
	}
	if !common.IsNil(o.IFER) {
		toSerialize["IFER"] = o.IFER
	}
	if !common.IsNil(o.UFR) {
		toSerialize["UFR"] = o.UFR
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountApiTradingStatusResponseDataTriggerCondition) UnmarshalJSON(data []byte) (err error) {
	varAccountApiTradingStatusResponseDataTriggerCondition := _AccountApiTradingStatusResponseDataTriggerCondition{}

	err = json.Unmarshal(data, &varAccountApiTradingStatusResponseDataTriggerCondition)

	if err != nil {
		return err
	}

	*o = AccountApiTradingStatusResponseDataTriggerCondition(varAccountApiTradingStatusResponseDataTriggerCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "GCR")
		delete(additionalProperties, "IFER")
		delete(additionalProperties, "UFR")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountApiTradingStatusResponseDataTriggerCondition struct {
	value *AccountApiTradingStatusResponseDataTriggerCondition
	isSet bool
}

func (v NullableAccountApiTradingStatusResponseDataTriggerCondition) Get() *AccountApiTradingStatusResponseDataTriggerCondition {
	return v.value
}

func (v *NullableAccountApiTradingStatusResponseDataTriggerCondition) Set(val *AccountApiTradingStatusResponseDataTriggerCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountApiTradingStatusResponseDataTriggerCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountApiTradingStatusResponseDataTriggerCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountApiTradingStatusResponseDataTriggerCondition(val *AccountApiTradingStatusResponseDataTriggerCondition) *NullableAccountApiTradingStatusResponseDataTriggerCondition {
	return &NullableAccountApiTradingStatusResponseDataTriggerCondition{value: val, isSet: true}
}

func (v NullableAccountApiTradingStatusResponseDataTriggerCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountApiTradingStatusResponseDataTriggerCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
