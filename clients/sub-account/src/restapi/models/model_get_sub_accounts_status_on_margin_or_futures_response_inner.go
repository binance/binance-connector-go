/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSubAccountsStatusOnMarginOrFuturesResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSubAccountsStatusOnMarginOrFuturesResponseInner{}

// GetSubAccountsStatusOnMarginOrFuturesResponseInner struct for GetSubAccountsStatusOnMarginOrFuturesResponseInner
type GetSubAccountsStatusOnMarginOrFuturesResponseInner struct {
	Email                *string `json:"email,omitempty"`
	IsSubUserEnabled     *bool   `json:"isSubUserEnabled,omitempty"`
	IsUserActive         *bool   `json:"isUserActive,omitempty"`
	InsertTime           *int64  `json:"insertTime,omitempty"`
	IsMarginEnabled      *bool   `json:"isMarginEnabled,omitempty"`
	IsFutureEnabled      *bool   `json:"isFutureEnabled,omitempty"`
	Mobile               *int64  `json:"mobile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSubAccountsStatusOnMarginOrFuturesResponseInner GetSubAccountsStatusOnMarginOrFuturesResponseInner

// NewGetSubAccountsStatusOnMarginOrFuturesResponseInner instantiates a new GetSubAccountsStatusOnMarginOrFuturesResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubAccountsStatusOnMarginOrFuturesResponseInner() *GetSubAccountsStatusOnMarginOrFuturesResponseInner {
	this := GetSubAccountsStatusOnMarginOrFuturesResponseInner{}
	return &this
}

// NewGetSubAccountsStatusOnMarginOrFuturesResponseInnerWithDefaults instantiates a new GetSubAccountsStatusOnMarginOrFuturesResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubAccountsStatusOnMarginOrFuturesResponseInnerWithDefaults() *GetSubAccountsStatusOnMarginOrFuturesResponseInner {
	this := GetSubAccountsStatusOnMarginOrFuturesResponseInner{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) SetEmail(v string) {
	o.Email = &v
}

// GetIsSubUserEnabled returns the IsSubUserEnabled field value if set, zero value otherwise.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) GetIsSubUserEnabled() bool {
	if o == nil || common.IsNil(o.IsSubUserEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSubUserEnabled
}

// GetIsSubUserEnabledOk returns a tuple with the IsSubUserEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) GetIsSubUserEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsSubUserEnabled) {
		return nil, false
	}
	return o.IsSubUserEnabled, true
}

// HasIsSubUserEnabled returns a boolean if a field has been set.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) HasIsSubUserEnabled() bool {
	if o != nil && !common.IsNil(o.IsSubUserEnabled) {
		return true
	}

	return false
}

// SetIsSubUserEnabled gets a reference to the given bool and assigns it to the IsSubUserEnabled field.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) SetIsSubUserEnabled(v bool) {
	o.IsSubUserEnabled = &v
}

// GetIsUserActive returns the IsUserActive field value if set, zero value otherwise.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) GetIsUserActive() bool {
	if o == nil || common.IsNil(o.IsUserActive) {
		var ret bool
		return ret
	}
	return *o.IsUserActive
}

// GetIsUserActiveOk returns a tuple with the IsUserActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) GetIsUserActiveOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsUserActive) {
		return nil, false
	}
	return o.IsUserActive, true
}

// HasIsUserActive returns a boolean if a field has been set.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) HasIsUserActive() bool {
	if o != nil && !common.IsNil(o.IsUserActive) {
		return true
	}

	return false
}

// SetIsUserActive gets a reference to the given bool and assigns it to the IsUserActive field.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) SetIsUserActive(v bool) {
	o.IsUserActive = &v
}

// GetInsertTime returns the InsertTime field value if set, zero value otherwise.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) GetInsertTime() int64 {
	if o == nil || common.IsNil(o.InsertTime) {
		var ret int64
		return ret
	}
	return *o.InsertTime
}

// GetInsertTimeOk returns a tuple with the InsertTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) GetInsertTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.InsertTime) {
		return nil, false
	}
	return o.InsertTime, true
}

// HasInsertTime returns a boolean if a field has been set.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) HasInsertTime() bool {
	if o != nil && !common.IsNil(o.InsertTime) {
		return true
	}

	return false
}

// SetInsertTime gets a reference to the given int64 and assigns it to the InsertTime field.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) SetInsertTime(v int64) {
	o.InsertTime = &v
}

// GetIsMarginEnabled returns the IsMarginEnabled field value if set, zero value otherwise.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) GetIsMarginEnabled() bool {
	if o == nil || common.IsNil(o.IsMarginEnabled) {
		var ret bool
		return ret
	}
	return *o.IsMarginEnabled
}

// GetIsMarginEnabledOk returns a tuple with the IsMarginEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) GetIsMarginEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsMarginEnabled) {
		return nil, false
	}
	return o.IsMarginEnabled, true
}

// HasIsMarginEnabled returns a boolean if a field has been set.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) HasIsMarginEnabled() bool {
	if o != nil && !common.IsNil(o.IsMarginEnabled) {
		return true
	}

	return false
}

// SetIsMarginEnabled gets a reference to the given bool and assigns it to the IsMarginEnabled field.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) SetIsMarginEnabled(v bool) {
	o.IsMarginEnabled = &v
}

// GetIsFutureEnabled returns the IsFutureEnabled field value if set, zero value otherwise.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) GetIsFutureEnabled() bool {
	if o == nil || common.IsNil(o.IsFutureEnabled) {
		var ret bool
		return ret
	}
	return *o.IsFutureEnabled
}

// GetIsFutureEnabledOk returns a tuple with the IsFutureEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) GetIsFutureEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsFutureEnabled) {
		return nil, false
	}
	return o.IsFutureEnabled, true
}

// HasIsFutureEnabled returns a boolean if a field has been set.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) HasIsFutureEnabled() bool {
	if o != nil && !common.IsNil(o.IsFutureEnabled) {
		return true
	}

	return false
}

// SetIsFutureEnabled gets a reference to the given bool and assigns it to the IsFutureEnabled field.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) SetIsFutureEnabled(v bool) {
	o.IsFutureEnabled = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) GetMobile() int64 {
	if o == nil || common.IsNil(o.Mobile) {
		var ret int64
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) GetMobileOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) HasMobile() bool {
	if o != nil && !common.IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given int64 and assigns it to the Mobile field.
func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) SetMobile(v int64) {
	o.Mobile = &v
}

func (o GetSubAccountsStatusOnMarginOrFuturesResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSubAccountsStatusOnMarginOrFuturesResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.IsSubUserEnabled) {
		toSerialize["isSubUserEnabled"] = o.IsSubUserEnabled
	}
	if !common.IsNil(o.IsUserActive) {
		toSerialize["isUserActive"] = o.IsUserActive
	}
	if !common.IsNil(o.InsertTime) {
		toSerialize["insertTime"] = o.InsertTime
	}
	if !common.IsNil(o.IsMarginEnabled) {
		toSerialize["isMarginEnabled"] = o.IsMarginEnabled
	}
	if !common.IsNil(o.IsFutureEnabled) {
		toSerialize["isFutureEnabled"] = o.IsFutureEnabled
	}
	if !common.IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSubAccountsStatusOnMarginOrFuturesResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetSubAccountsStatusOnMarginOrFuturesResponseInner := _GetSubAccountsStatusOnMarginOrFuturesResponseInner{}

	err = json.Unmarshal(data, &varGetSubAccountsStatusOnMarginOrFuturesResponseInner)

	if err != nil {
		return err
	}

	*o = GetSubAccountsStatusOnMarginOrFuturesResponseInner(varGetSubAccountsStatusOnMarginOrFuturesResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "isSubUserEnabled")
		delete(additionalProperties, "isUserActive")
		delete(additionalProperties, "insertTime")
		delete(additionalProperties, "isMarginEnabled")
		delete(additionalProperties, "isFutureEnabled")
		delete(additionalProperties, "mobile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSubAccountsStatusOnMarginOrFuturesResponseInner struct {
	value *GetSubAccountsStatusOnMarginOrFuturesResponseInner
	isSet bool
}

func (v NullableGetSubAccountsStatusOnMarginOrFuturesResponseInner) Get() *GetSubAccountsStatusOnMarginOrFuturesResponseInner {
	return v.value
}

func (v *NullableGetSubAccountsStatusOnMarginOrFuturesResponseInner) Set(val *GetSubAccountsStatusOnMarginOrFuturesResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubAccountsStatusOnMarginOrFuturesResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubAccountsStatusOnMarginOrFuturesResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubAccountsStatusOnMarginOrFuturesResponseInner(val *GetSubAccountsStatusOnMarginOrFuturesResponseInner) *NullableGetSubAccountsStatusOnMarginOrFuturesResponseInner {
	return &NullableGetSubAccountsStatusOnMarginOrFuturesResponseInner{value: val, isSet: true}
}

func (v NullableGetSubAccountsStatusOnMarginOrFuturesResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubAccountsStatusOnMarginOrFuturesResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
