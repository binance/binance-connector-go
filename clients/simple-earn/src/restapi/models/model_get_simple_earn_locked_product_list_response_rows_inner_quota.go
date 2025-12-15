/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSimpleEarnLockedProductListResponseRowsInnerQuota type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSimpleEarnLockedProductListResponseRowsInnerQuota{}

// GetSimpleEarnLockedProductListResponseRowsInnerQuota struct for GetSimpleEarnLockedProductListResponseRowsInnerQuota
type GetSimpleEarnLockedProductListResponseRowsInnerQuota struct {
	TotalPersonalQuota   *string `json:"totalPersonalQuota,omitempty"`
	Minimum              *string `json:"minimum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSimpleEarnLockedProductListResponseRowsInnerQuota GetSimpleEarnLockedProductListResponseRowsInnerQuota

// NewGetSimpleEarnLockedProductListResponseRowsInnerQuota instantiates a new GetSimpleEarnLockedProductListResponseRowsInnerQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSimpleEarnLockedProductListResponseRowsInnerQuota() *GetSimpleEarnLockedProductListResponseRowsInnerQuota {
	this := GetSimpleEarnLockedProductListResponseRowsInnerQuota{}
	return &this
}

// NewGetSimpleEarnLockedProductListResponseRowsInnerQuotaWithDefaults instantiates a new GetSimpleEarnLockedProductListResponseRowsInnerQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSimpleEarnLockedProductListResponseRowsInnerQuotaWithDefaults() *GetSimpleEarnLockedProductListResponseRowsInnerQuota {
	this := GetSimpleEarnLockedProductListResponseRowsInnerQuota{}
	return &this
}

// GetTotalPersonalQuota returns the TotalPersonalQuota field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerQuota) GetTotalPersonalQuota() string {
	if o == nil || common.IsNil(o.TotalPersonalQuota) {
		var ret string
		return ret
	}
	return *o.TotalPersonalQuota
}

// GetTotalPersonalQuotaOk returns a tuple with the TotalPersonalQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerQuota) GetTotalPersonalQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalPersonalQuota) {
		return nil, false
	}
	return o.TotalPersonalQuota, true
}

// HasTotalPersonalQuota returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerQuota) HasTotalPersonalQuota() bool {
	if o != nil && !common.IsNil(o.TotalPersonalQuota) {
		return true
	}

	return false
}

// SetTotalPersonalQuota gets a reference to the given string and assigns it to the TotalPersonalQuota field.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerQuota) SetTotalPersonalQuota(v string) {
	o.TotalPersonalQuota = &v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerQuota) GetMinimum() string {
	if o == nil || common.IsNil(o.Minimum) {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerQuota) GetMinimumOk() (*string, bool) {
	if o == nil || common.IsNil(o.Minimum) {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerQuota) HasMinimum() bool {
	if o != nil && !common.IsNil(o.Minimum) {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerQuota) SetMinimum(v string) {
	o.Minimum = &v
}

func (o GetSimpleEarnLockedProductListResponseRowsInnerQuota) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSimpleEarnLockedProductListResponseRowsInnerQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TotalPersonalQuota) {
		toSerialize["totalPersonalQuota"] = o.TotalPersonalQuota
	}
	if !common.IsNil(o.Minimum) {
		toSerialize["minimum"] = o.Minimum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSimpleEarnLockedProductListResponseRowsInnerQuota) UnmarshalJSON(data []byte) (err error) {
	varGetSimpleEarnLockedProductListResponseRowsInnerQuota := _GetSimpleEarnLockedProductListResponseRowsInnerQuota{}

	err = json.Unmarshal(data, &varGetSimpleEarnLockedProductListResponseRowsInnerQuota)

	if err != nil {
		return err
	}

	*o = GetSimpleEarnLockedProductListResponseRowsInnerQuota(varGetSimpleEarnLockedProductListResponseRowsInnerQuota)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "totalPersonalQuota")
		delete(additionalProperties, "minimum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSimpleEarnLockedProductListResponseRowsInnerQuota struct {
	value *GetSimpleEarnLockedProductListResponseRowsInnerQuota
	isSet bool
}

func (v NullableGetSimpleEarnLockedProductListResponseRowsInnerQuota) Get() *GetSimpleEarnLockedProductListResponseRowsInnerQuota {
	return v.value
}

func (v *NullableGetSimpleEarnLockedProductListResponseRowsInnerQuota) Set(val *GetSimpleEarnLockedProductListResponseRowsInnerQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSimpleEarnLockedProductListResponseRowsInnerQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSimpleEarnLockedProductListResponseRowsInnerQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSimpleEarnLockedProductListResponseRowsInnerQuota(val *GetSimpleEarnLockedProductListResponseRowsInnerQuota) *NullableGetSimpleEarnLockedProductListResponseRowsInnerQuota {
	return &NullableGetSimpleEarnLockedProductListResponseRowsInnerQuota{value: val, isSet: true}
}

func (v NullableGetSimpleEarnLockedProductListResponseRowsInnerQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSimpleEarnLockedProductListResponseRowsInnerQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
