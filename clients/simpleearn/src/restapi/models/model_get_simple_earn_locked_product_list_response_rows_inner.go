/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetSimpleEarnLockedProductListResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSimpleEarnLockedProductListResponseRowsInner{}

// GetSimpleEarnLockedProductListResponseRowsInner struct for GetSimpleEarnLockedProductListResponseRowsInner
type GetSimpleEarnLockedProductListResponseRowsInner struct {
	ProjectId            *string                                                `json:"projectId,omitempty"`
	Detail               *GetSimpleEarnLockedProductListResponseRowsInnerDetail `json:"detail,omitempty"`
	Quota                *GetSimpleEarnLockedProductListResponseRowsInnerQuota  `json:"quota,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSimpleEarnLockedProductListResponseRowsInner GetSimpleEarnLockedProductListResponseRowsInner

// NewGetSimpleEarnLockedProductListResponseRowsInner instantiates a new GetSimpleEarnLockedProductListResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSimpleEarnLockedProductListResponseRowsInner() *GetSimpleEarnLockedProductListResponseRowsInner {
	this := GetSimpleEarnLockedProductListResponseRowsInner{}
	return &this
}

// NewGetSimpleEarnLockedProductListResponseRowsInnerWithDefaults instantiates a new GetSimpleEarnLockedProductListResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSimpleEarnLockedProductListResponseRowsInnerWithDefaults() *GetSimpleEarnLockedProductListResponseRowsInner {
	this := GetSimpleEarnLockedProductListResponseRowsInner{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInner) GetProjectId() string {
	if o == nil || common.IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInner) GetProjectIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInner) HasProjectId() bool {
	if o != nil && !common.IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GetSimpleEarnLockedProductListResponseRowsInner) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInner) GetDetail() GetSimpleEarnLockedProductListResponseRowsInnerDetail {
	if o == nil || common.IsNil(o.Detail) {
		var ret GetSimpleEarnLockedProductListResponseRowsInnerDetail
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInner) GetDetailOk() (*GetSimpleEarnLockedProductListResponseRowsInnerDetail, bool) {
	if o == nil || common.IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInner) HasDetail() bool {
	if o != nil && !common.IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given GetSimpleEarnLockedProductListResponseRowsInnerDetail and assigns it to the Detail field.
func (o *GetSimpleEarnLockedProductListResponseRowsInner) SetDetail(v GetSimpleEarnLockedProductListResponseRowsInnerDetail) {
	o.Detail = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInner) GetQuota() GetSimpleEarnLockedProductListResponseRowsInnerQuota {
	if o == nil || common.IsNil(o.Quota) {
		var ret GetSimpleEarnLockedProductListResponseRowsInnerQuota
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInner) GetQuotaOk() (*GetSimpleEarnLockedProductListResponseRowsInnerQuota, bool) {
	if o == nil || common.IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInner) HasQuota() bool {
	if o != nil && !common.IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given GetSimpleEarnLockedProductListResponseRowsInnerQuota and assigns it to the Quota field.
func (o *GetSimpleEarnLockedProductListResponseRowsInner) SetQuota(v GetSimpleEarnLockedProductListResponseRowsInnerQuota) {
	o.Quota = &v
}

func (o GetSimpleEarnLockedProductListResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSimpleEarnLockedProductListResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !common.IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !common.IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSimpleEarnLockedProductListResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetSimpleEarnLockedProductListResponseRowsInner := _GetSimpleEarnLockedProductListResponseRowsInner{}

	err = json.Unmarshal(data, &varGetSimpleEarnLockedProductListResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetSimpleEarnLockedProductListResponseRowsInner(varGetSimpleEarnLockedProductListResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "projectId")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "quota")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSimpleEarnLockedProductListResponseRowsInner struct {
	value *GetSimpleEarnLockedProductListResponseRowsInner
	isSet bool
}

func (v NullableGetSimpleEarnLockedProductListResponseRowsInner) Get() *GetSimpleEarnLockedProductListResponseRowsInner {
	return v.value
}

func (v *NullableGetSimpleEarnLockedProductListResponseRowsInner) Set(val *GetSimpleEarnLockedProductListResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSimpleEarnLockedProductListResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSimpleEarnLockedProductListResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSimpleEarnLockedProductListResponseRowsInner(val *GetSimpleEarnLockedProductListResponseRowsInner) *NullableGetSimpleEarnLockedProductListResponseRowsInner {
	return &NullableGetSimpleEarnLockedProductListResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetSimpleEarnLockedProductListResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSimpleEarnLockedProductListResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
