/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOnChainYieldsLockedProductListResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOnChainYieldsLockedProductListResponseRowsInner{}

// GetOnChainYieldsLockedProductListResponseRowsInner struct for GetOnChainYieldsLockedProductListResponseRowsInner
type GetOnChainYieldsLockedProductListResponseRowsInner struct {
	ProjectId            *string                                                   `json:"projectId,omitempty"`
	Detail               *GetOnChainYieldsLockedProductListResponseRowsInnerDetail `json:"detail,omitempty"`
	Quota                *GetOnChainYieldsLockedProductListResponseRowsInnerQuota  `json:"quota,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOnChainYieldsLockedProductListResponseRowsInner GetOnChainYieldsLockedProductListResponseRowsInner

// NewGetOnChainYieldsLockedProductListResponseRowsInner instantiates a new GetOnChainYieldsLockedProductListResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOnChainYieldsLockedProductListResponseRowsInner() *GetOnChainYieldsLockedProductListResponseRowsInner {
	this := GetOnChainYieldsLockedProductListResponseRowsInner{}
	return &this
}

// NewGetOnChainYieldsLockedProductListResponseRowsInnerWithDefaults instantiates a new GetOnChainYieldsLockedProductListResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOnChainYieldsLockedProductListResponseRowsInnerWithDefaults() *GetOnChainYieldsLockedProductListResponseRowsInner {
	this := GetOnChainYieldsLockedProductListResponseRowsInner{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductListResponseRowsInner) GetProjectId() string {
	if o == nil || common.IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInner) GetProjectIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInner) HasProjectId() bool {
	if o != nil && !common.IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GetOnChainYieldsLockedProductListResponseRowsInner) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductListResponseRowsInner) GetDetail() GetOnChainYieldsLockedProductListResponseRowsInnerDetail {
	if o == nil || common.IsNil(o.Detail) {
		var ret GetOnChainYieldsLockedProductListResponseRowsInnerDetail
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInner) GetDetailOk() (*GetOnChainYieldsLockedProductListResponseRowsInnerDetail, bool) {
	if o == nil || common.IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInner) HasDetail() bool {
	if o != nil && !common.IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given GetOnChainYieldsLockedProductListResponseRowsInnerDetail and assigns it to the Detail field.
func (o *GetOnChainYieldsLockedProductListResponseRowsInner) SetDetail(v GetOnChainYieldsLockedProductListResponseRowsInnerDetail) {
	o.Detail = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductListResponseRowsInner) GetQuota() GetOnChainYieldsLockedProductListResponseRowsInnerQuota {
	if o == nil || common.IsNil(o.Quota) {
		var ret GetOnChainYieldsLockedProductListResponseRowsInnerQuota
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInner) GetQuotaOk() (*GetOnChainYieldsLockedProductListResponseRowsInnerQuota, bool) {
	if o == nil || common.IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInner) HasQuota() bool {
	if o != nil && !common.IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given GetOnChainYieldsLockedProductListResponseRowsInnerQuota and assigns it to the Quota field.
func (o *GetOnChainYieldsLockedProductListResponseRowsInner) SetQuota(v GetOnChainYieldsLockedProductListResponseRowsInnerQuota) {
	o.Quota = &v
}

func (o GetOnChainYieldsLockedProductListResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOnChainYieldsLockedProductListResponseRowsInner) ToMap() (map[string]interface{}, error) {
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

func (o *GetOnChainYieldsLockedProductListResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetOnChainYieldsLockedProductListResponseRowsInner := _GetOnChainYieldsLockedProductListResponseRowsInner{}

	err = json.Unmarshal(data, &varGetOnChainYieldsLockedProductListResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetOnChainYieldsLockedProductListResponseRowsInner(varGetOnChainYieldsLockedProductListResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "projectId")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "quota")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOnChainYieldsLockedProductListResponseRowsInner struct {
	value *GetOnChainYieldsLockedProductListResponseRowsInner
	isSet bool
}

func (v NullableGetOnChainYieldsLockedProductListResponseRowsInner) Get() *GetOnChainYieldsLockedProductListResponseRowsInner {
	return v.value
}

func (v *NullableGetOnChainYieldsLockedProductListResponseRowsInner) Set(val *GetOnChainYieldsLockedProductListResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOnChainYieldsLockedProductListResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOnChainYieldsLockedProductListResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOnChainYieldsLockedProductListResponseRowsInner(val *GetOnChainYieldsLockedProductListResponseRowsInner) *NullableGetOnChainYieldsLockedProductListResponseRowsInner {
	return &NullableGetOnChainYieldsLockedProductListResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetOnChainYieldsLockedProductListResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOnChainYieldsLockedProductListResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
