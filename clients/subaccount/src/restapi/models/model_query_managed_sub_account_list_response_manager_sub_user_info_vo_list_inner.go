/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner{}

// QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner struct for QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner
type QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner struct {
	RootUserId               *int64  `json:"rootUserId,omitempty"`
	ManagersubUserId         *int64  `json:"managersubUserId,omitempty"`
	BindParentUserId         *int64  `json:"bindParentUserId,omitempty"`
	Email                    *string `json:"email,omitempty"`
	InsertTimeStamp          *int64  `json:"insertTimeStamp,omitempty"`
	BindParentEmail          *string `json:"bindParentEmail,omitempty"`
	IsSubUserEnabled         *bool   `json:"isSubUserEnabled,omitempty"`
	IsUserActive             *bool   `json:"isUserActive,omitempty"`
	IsMarginEnabled          *bool   `json:"isMarginEnabled,omitempty"`
	IsFutureEnabled          *bool   `json:"isFutureEnabled,omitempty"`
	IsSignedLVTRiskAgreement *bool   `json:"isSignedLVTRiskAgreement,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner

// NewQueryManagedSubAccountListResponseManagerSubUserInfoVoListInner instantiates a new QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountListResponseManagerSubUserInfoVoListInner() *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner {
	this := QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner{}
	return &this
}

// NewQueryManagedSubAccountListResponseManagerSubUserInfoVoListInnerWithDefaults instantiates a new QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountListResponseManagerSubUserInfoVoListInnerWithDefaults() *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner {
	this := QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner{}
	return &this
}

// GetRootUserId returns the RootUserId field value if set, zero value otherwise.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetRootUserId() int64 {
	if o == nil || common.IsNil(o.RootUserId) {
		var ret int64
		return ret
	}
	return *o.RootUserId
}

// GetRootUserIdOk returns a tuple with the RootUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetRootUserIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RootUserId) {
		return nil, false
	}
	return o.RootUserId, true
}

// HasRootUserId returns a boolean if a field has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) HasRootUserId() bool {
	if o != nil && !common.IsNil(o.RootUserId) {
		return true
	}

	return false
}

// SetRootUserId gets a reference to the given int64 and assigns it to the RootUserId field.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) SetRootUserId(v int64) {
	o.RootUserId = &v
}

// GetManagersubUserId returns the ManagersubUserId field value if set, zero value otherwise.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetManagersubUserId() int64 {
	if o == nil || common.IsNil(o.ManagersubUserId) {
		var ret int64
		return ret
	}
	return *o.ManagersubUserId
}

// GetManagersubUserIdOk returns a tuple with the ManagersubUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetManagersubUserIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ManagersubUserId) {
		return nil, false
	}
	return o.ManagersubUserId, true
}

// HasManagersubUserId returns a boolean if a field has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) HasManagersubUserId() bool {
	if o != nil && !common.IsNil(o.ManagersubUserId) {
		return true
	}

	return false
}

// SetManagersubUserId gets a reference to the given int64 and assigns it to the ManagersubUserId field.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) SetManagersubUserId(v int64) {
	o.ManagersubUserId = &v
}

// GetBindParentUserId returns the BindParentUserId field value if set, zero value otherwise.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetBindParentUserId() int64 {
	if o == nil || common.IsNil(o.BindParentUserId) {
		var ret int64
		return ret
	}
	return *o.BindParentUserId
}

// GetBindParentUserIdOk returns a tuple with the BindParentUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetBindParentUserIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BindParentUserId) {
		return nil, false
	}
	return o.BindParentUserId, true
}

// HasBindParentUserId returns a boolean if a field has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) HasBindParentUserId() bool {
	if o != nil && !common.IsNil(o.BindParentUserId) {
		return true
	}

	return false
}

// SetBindParentUserId gets a reference to the given int64 and assigns it to the BindParentUserId field.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) SetBindParentUserId(v int64) {
	o.BindParentUserId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) SetEmail(v string) {
	o.Email = &v
}

// GetInsertTimeStamp returns the InsertTimeStamp field value if set, zero value otherwise.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetInsertTimeStamp() int64 {
	if o == nil || common.IsNil(o.InsertTimeStamp) {
		var ret int64
		return ret
	}
	return *o.InsertTimeStamp
}

// GetInsertTimeStampOk returns a tuple with the InsertTimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetInsertTimeStampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.InsertTimeStamp) {
		return nil, false
	}
	return o.InsertTimeStamp, true
}

// HasInsertTimeStamp returns a boolean if a field has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) HasInsertTimeStamp() bool {
	if o != nil && !common.IsNil(o.InsertTimeStamp) {
		return true
	}

	return false
}

// SetInsertTimeStamp gets a reference to the given int64 and assigns it to the InsertTimeStamp field.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) SetInsertTimeStamp(v int64) {
	o.InsertTimeStamp = &v
}

// GetBindParentEmail returns the BindParentEmail field value if set, zero value otherwise.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetBindParentEmail() string {
	if o == nil || common.IsNil(o.BindParentEmail) {
		var ret string
		return ret
	}
	return *o.BindParentEmail
}

// GetBindParentEmailOk returns a tuple with the BindParentEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetBindParentEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.BindParentEmail) {
		return nil, false
	}
	return o.BindParentEmail, true
}

// HasBindParentEmail returns a boolean if a field has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) HasBindParentEmail() bool {
	if o != nil && !common.IsNil(o.BindParentEmail) {
		return true
	}

	return false
}

// SetBindParentEmail gets a reference to the given string and assigns it to the BindParentEmail field.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) SetBindParentEmail(v string) {
	o.BindParentEmail = &v
}

// GetIsSubUserEnabled returns the IsSubUserEnabled field value if set, zero value otherwise.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetIsSubUserEnabled() bool {
	if o == nil || common.IsNil(o.IsSubUserEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSubUserEnabled
}

// GetIsSubUserEnabledOk returns a tuple with the IsSubUserEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetIsSubUserEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsSubUserEnabled) {
		return nil, false
	}
	return o.IsSubUserEnabled, true
}

// HasIsSubUserEnabled returns a boolean if a field has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) HasIsSubUserEnabled() bool {
	if o != nil && !common.IsNil(o.IsSubUserEnabled) {
		return true
	}

	return false
}

// SetIsSubUserEnabled gets a reference to the given bool and assigns it to the IsSubUserEnabled field.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) SetIsSubUserEnabled(v bool) {
	o.IsSubUserEnabled = &v
}

// GetIsUserActive returns the IsUserActive field value if set, zero value otherwise.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetIsUserActive() bool {
	if o == nil || common.IsNil(o.IsUserActive) {
		var ret bool
		return ret
	}
	return *o.IsUserActive
}

// GetIsUserActiveOk returns a tuple with the IsUserActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetIsUserActiveOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsUserActive) {
		return nil, false
	}
	return o.IsUserActive, true
}

// HasIsUserActive returns a boolean if a field has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) HasIsUserActive() bool {
	if o != nil && !common.IsNil(o.IsUserActive) {
		return true
	}

	return false
}

// SetIsUserActive gets a reference to the given bool and assigns it to the IsUserActive field.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) SetIsUserActive(v bool) {
	o.IsUserActive = &v
}

// GetIsMarginEnabled returns the IsMarginEnabled field value if set, zero value otherwise.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetIsMarginEnabled() bool {
	if o == nil || common.IsNil(o.IsMarginEnabled) {
		var ret bool
		return ret
	}
	return *o.IsMarginEnabled
}

// GetIsMarginEnabledOk returns a tuple with the IsMarginEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetIsMarginEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsMarginEnabled) {
		return nil, false
	}
	return o.IsMarginEnabled, true
}

// HasIsMarginEnabled returns a boolean if a field has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) HasIsMarginEnabled() bool {
	if o != nil && !common.IsNil(o.IsMarginEnabled) {
		return true
	}

	return false
}

// SetIsMarginEnabled gets a reference to the given bool and assigns it to the IsMarginEnabled field.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) SetIsMarginEnabled(v bool) {
	o.IsMarginEnabled = &v
}

// GetIsFutureEnabled returns the IsFutureEnabled field value if set, zero value otherwise.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetIsFutureEnabled() bool {
	if o == nil || common.IsNil(o.IsFutureEnabled) {
		var ret bool
		return ret
	}
	return *o.IsFutureEnabled
}

// GetIsFutureEnabledOk returns a tuple with the IsFutureEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetIsFutureEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsFutureEnabled) {
		return nil, false
	}
	return o.IsFutureEnabled, true
}

// HasIsFutureEnabled returns a boolean if a field has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) HasIsFutureEnabled() bool {
	if o != nil && !common.IsNil(o.IsFutureEnabled) {
		return true
	}

	return false
}

// SetIsFutureEnabled gets a reference to the given bool and assigns it to the IsFutureEnabled field.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) SetIsFutureEnabled(v bool) {
	o.IsFutureEnabled = &v
}

// GetIsSignedLVTRiskAgreement returns the IsSignedLVTRiskAgreement field value if set, zero value otherwise.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetIsSignedLVTRiskAgreement() bool {
	if o == nil || common.IsNil(o.IsSignedLVTRiskAgreement) {
		var ret bool
		return ret
	}
	return *o.IsSignedLVTRiskAgreement
}

// GetIsSignedLVTRiskAgreementOk returns a tuple with the IsSignedLVTRiskAgreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) GetIsSignedLVTRiskAgreementOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsSignedLVTRiskAgreement) {
		return nil, false
	}
	return o.IsSignedLVTRiskAgreement, true
}

// HasIsSignedLVTRiskAgreement returns a boolean if a field has been set.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) HasIsSignedLVTRiskAgreement() bool {
	if o != nil && !common.IsNil(o.IsSignedLVTRiskAgreement) {
		return true
	}

	return false
}

// SetIsSignedLVTRiskAgreement gets a reference to the given bool and assigns it to the IsSignedLVTRiskAgreement field.
func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) SetIsSignedLVTRiskAgreement(v bool) {
	o.IsSignedLVTRiskAgreement = &v
}

func (o QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RootUserId) {
		toSerialize["rootUserId"] = o.RootUserId
	}
	if !common.IsNil(o.ManagersubUserId) {
		toSerialize["managersubUserId"] = o.ManagersubUserId
	}
	if !common.IsNil(o.BindParentUserId) {
		toSerialize["bindParentUserId"] = o.BindParentUserId
	}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.InsertTimeStamp) {
		toSerialize["insertTimeStamp"] = o.InsertTimeStamp
	}
	if !common.IsNil(o.BindParentEmail) {
		toSerialize["bindParentEmail"] = o.BindParentEmail
	}
	if !common.IsNil(o.IsSubUserEnabled) {
		toSerialize["isSubUserEnabled"] = o.IsSubUserEnabled
	}
	if !common.IsNil(o.IsUserActive) {
		toSerialize["isUserActive"] = o.IsUserActive
	}
	if !common.IsNil(o.IsMarginEnabled) {
		toSerialize["isMarginEnabled"] = o.IsMarginEnabled
	}
	if !common.IsNil(o.IsFutureEnabled) {
		toSerialize["isFutureEnabled"] = o.IsFutureEnabled
	}
	if !common.IsNil(o.IsSignedLVTRiskAgreement) {
		toSerialize["isSignedLVTRiskAgreement"] = o.IsSignedLVTRiskAgreement
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountListResponseManagerSubUserInfoVoListInner := _QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountListResponseManagerSubUserInfoVoListInner)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner(varQueryManagedSubAccountListResponseManagerSubUserInfoVoListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rootUserId")
		delete(additionalProperties, "managersubUserId")
		delete(additionalProperties, "bindParentUserId")
		delete(additionalProperties, "email")
		delete(additionalProperties, "insertTimeStamp")
		delete(additionalProperties, "bindParentEmail")
		delete(additionalProperties, "isSubUserEnabled")
		delete(additionalProperties, "isUserActive")
		delete(additionalProperties, "isMarginEnabled")
		delete(additionalProperties, "isFutureEnabled")
		delete(additionalProperties, "isSignedLVTRiskAgreement")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountListResponseManagerSubUserInfoVoListInner struct {
	value *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner
	isSet bool
}

func (v NullableQueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) Get() *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner {
	return v.value
}

func (v *NullableQueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) Set(val *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountListResponseManagerSubUserInfoVoListInner(val *QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) *NullableQueryManagedSubAccountListResponseManagerSubUserInfoVoListInner {
	return &NullableQueryManagedSubAccountListResponseManagerSubUserInfoVoListInner{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
