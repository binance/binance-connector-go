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

// checks if the QuerySubAccountListResponseSubAccountsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubAccountListResponseSubAccountsInner{}

// QuerySubAccountListResponseSubAccountsInner struct for QuerySubAccountListResponseSubAccountsInner
type QuerySubAccountListResponseSubAccountsInner struct {
	SubUserId                   *int64  `json:"subUserId,omitempty"`
	Email                       *string `json:"email,omitempty"`
	Remark                      *string `json:"remark,omitempty"`
	IsFreeze                    *bool   `json:"isFreeze,omitempty"`
	CreateTime                  *int64  `json:"createTime,omitempty"`
	IsManagedSubAccount         *bool   `json:"isManagedSubAccount,omitempty"`
	IsAssetManagementSubAccount *bool   `json:"isAssetManagementSubAccount,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _QuerySubAccountListResponseSubAccountsInner QuerySubAccountListResponseSubAccountsInner

// NewQuerySubAccountListResponseSubAccountsInner instantiates a new QuerySubAccountListResponseSubAccountsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubAccountListResponseSubAccountsInner() *QuerySubAccountListResponseSubAccountsInner {
	this := QuerySubAccountListResponseSubAccountsInner{}
	return &this
}

// NewQuerySubAccountListResponseSubAccountsInnerWithDefaults instantiates a new QuerySubAccountListResponseSubAccountsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubAccountListResponseSubAccountsInnerWithDefaults() *QuerySubAccountListResponseSubAccountsInner {
	this := QuerySubAccountListResponseSubAccountsInner{}
	return &this
}

// GetSubUserId returns the SubUserId field value if set, zero value otherwise.
func (o *QuerySubAccountListResponseSubAccountsInner) GetSubUserId() int64 {
	if o == nil || common.IsNil(o.SubUserId) {
		var ret int64
		return ret
	}
	return *o.SubUserId
}

// GetSubUserIdOk returns a tuple with the SubUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountListResponseSubAccountsInner) GetSubUserIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.SubUserId) {
		return nil, false
	}
	return o.SubUserId, true
}

// HasSubUserId returns a boolean if a field has been set.
func (o *QuerySubAccountListResponseSubAccountsInner) HasSubUserId() bool {
	if o != nil && !common.IsNil(o.SubUserId) {
		return true
	}

	return false
}

// SetSubUserId gets a reference to the given int64 and assigns it to the SubUserId field.
func (o *QuerySubAccountListResponseSubAccountsInner) SetSubUserId(v int64) {
	o.SubUserId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *QuerySubAccountListResponseSubAccountsInner) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountListResponseSubAccountsInner) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *QuerySubAccountListResponseSubAccountsInner) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *QuerySubAccountListResponseSubAccountsInner) SetEmail(v string) {
	o.Email = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *QuerySubAccountListResponseSubAccountsInner) GetRemark() string {
	if o == nil || common.IsNil(o.Remark) {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountListResponseSubAccountsInner) GetRemarkOk() (*string, bool) {
	if o == nil || common.IsNil(o.Remark) {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *QuerySubAccountListResponseSubAccountsInner) HasRemark() bool {
	if o != nil && !common.IsNil(o.Remark) {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *QuerySubAccountListResponseSubAccountsInner) SetRemark(v string) {
	o.Remark = &v
}

// GetIsFreeze returns the IsFreeze field value if set, zero value otherwise.
func (o *QuerySubAccountListResponseSubAccountsInner) GetIsFreeze() bool {
	if o == nil || common.IsNil(o.IsFreeze) {
		var ret bool
		return ret
	}
	return *o.IsFreeze
}

// GetIsFreezeOk returns a tuple with the IsFreeze field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountListResponseSubAccountsInner) GetIsFreezeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsFreeze) {
		return nil, false
	}
	return o.IsFreeze, true
}

// HasIsFreeze returns a boolean if a field has been set.
func (o *QuerySubAccountListResponseSubAccountsInner) HasIsFreeze() bool {
	if o != nil && !common.IsNil(o.IsFreeze) {
		return true
	}

	return false
}

// SetIsFreeze gets a reference to the given bool and assigns it to the IsFreeze field.
func (o *QuerySubAccountListResponseSubAccountsInner) SetIsFreeze(v bool) {
	o.IsFreeze = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *QuerySubAccountListResponseSubAccountsInner) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountListResponseSubAccountsInner) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *QuerySubAccountListResponseSubAccountsInner) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *QuerySubAccountListResponseSubAccountsInner) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetIsManagedSubAccount returns the IsManagedSubAccount field value if set, zero value otherwise.
func (o *QuerySubAccountListResponseSubAccountsInner) GetIsManagedSubAccount() bool {
	if o == nil || common.IsNil(o.IsManagedSubAccount) {
		var ret bool
		return ret
	}
	return *o.IsManagedSubAccount
}

// GetIsManagedSubAccountOk returns a tuple with the IsManagedSubAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountListResponseSubAccountsInner) GetIsManagedSubAccountOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsManagedSubAccount) {
		return nil, false
	}
	return o.IsManagedSubAccount, true
}

// HasIsManagedSubAccount returns a boolean if a field has been set.
func (o *QuerySubAccountListResponseSubAccountsInner) HasIsManagedSubAccount() bool {
	if o != nil && !common.IsNil(o.IsManagedSubAccount) {
		return true
	}

	return false
}

// SetIsManagedSubAccount gets a reference to the given bool and assigns it to the IsManagedSubAccount field.
func (o *QuerySubAccountListResponseSubAccountsInner) SetIsManagedSubAccount(v bool) {
	o.IsManagedSubAccount = &v
}

// GetIsAssetManagementSubAccount returns the IsAssetManagementSubAccount field value if set, zero value otherwise.
func (o *QuerySubAccountListResponseSubAccountsInner) GetIsAssetManagementSubAccount() bool {
	if o == nil || common.IsNil(o.IsAssetManagementSubAccount) {
		var ret bool
		return ret
	}
	return *o.IsAssetManagementSubAccount
}

// GetIsAssetManagementSubAccountOk returns a tuple with the IsAssetManagementSubAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountListResponseSubAccountsInner) GetIsAssetManagementSubAccountOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsAssetManagementSubAccount) {
		return nil, false
	}
	return o.IsAssetManagementSubAccount, true
}

// HasIsAssetManagementSubAccount returns a boolean if a field has been set.
func (o *QuerySubAccountListResponseSubAccountsInner) HasIsAssetManagementSubAccount() bool {
	if o != nil && !common.IsNil(o.IsAssetManagementSubAccount) {
		return true
	}

	return false
}

// SetIsAssetManagementSubAccount gets a reference to the given bool and assigns it to the IsAssetManagementSubAccount field.
func (o *QuerySubAccountListResponseSubAccountsInner) SetIsAssetManagementSubAccount(v bool) {
	o.IsAssetManagementSubAccount = &v
}

func (o QuerySubAccountListResponseSubAccountsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubAccountListResponseSubAccountsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.SubUserId) {
		toSerialize["subUserId"] = o.SubUserId
	}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.Remark) {
		toSerialize["remark"] = o.Remark
	}
	if !common.IsNil(o.IsFreeze) {
		toSerialize["isFreeze"] = o.IsFreeze
	}
	if !common.IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !common.IsNil(o.IsManagedSubAccount) {
		toSerialize["isManagedSubAccount"] = o.IsManagedSubAccount
	}
	if !common.IsNil(o.IsAssetManagementSubAccount) {
		toSerialize["isAssetManagementSubAccount"] = o.IsAssetManagementSubAccount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySubAccountListResponseSubAccountsInner) UnmarshalJSON(data []byte) (err error) {
	varQuerySubAccountListResponseSubAccountsInner := _QuerySubAccountListResponseSubAccountsInner{}

	err = json.Unmarshal(data, &varQuerySubAccountListResponseSubAccountsInner)

	if err != nil {
		return err
	}

	*o = QuerySubAccountListResponseSubAccountsInner(varQuerySubAccountListResponseSubAccountsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "subUserId")
		delete(additionalProperties, "email")
		delete(additionalProperties, "remark")
		delete(additionalProperties, "isFreeze")
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "isManagedSubAccount")
		delete(additionalProperties, "isAssetManagementSubAccount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySubAccountListResponseSubAccountsInner struct {
	value *QuerySubAccountListResponseSubAccountsInner
	isSet bool
}

func (v NullableQuerySubAccountListResponseSubAccountsInner) Get() *QuerySubAccountListResponseSubAccountsInner {
	return v.value
}

func (v *NullableQuerySubAccountListResponseSubAccountsInner) Set(val *QuerySubAccountListResponseSubAccountsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubAccountListResponseSubAccountsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubAccountListResponseSubAccountsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubAccountListResponseSubAccountsInner(val *QuerySubAccountListResponseSubAccountsInner) *NullableQuerySubAccountListResponseSubAccountsInner {
	return &NullableQuerySubAccountListResponseSubAccountsInner{value: val, isSet: true}
}

func (v NullableQuerySubAccountListResponseSubAccountsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubAccountListResponseSubAccountsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
