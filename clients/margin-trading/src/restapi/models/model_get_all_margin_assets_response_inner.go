/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetAllMarginAssetsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetAllMarginAssetsResponseInner{}

// GetAllMarginAssetsResponseInner struct for GetAllMarginAssetsResponseInner
type GetAllMarginAssetsResponseInner struct {
	AssetFullName        *string `json:"assetFullName,omitempty"`
	AssetName            *string `json:"assetName,omitempty"`
	IsBorrowable         *bool   `json:"isBorrowable,omitempty"`
	IsMortgageable       *bool   `json:"isMortgageable,omitempty"`
	UserMinBorrow        *string `json:"userMinBorrow,omitempty"`
	UserMinRepay         *string `json:"userMinRepay,omitempty"`
	DelistTime           *int64  `json:"delistTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetAllMarginAssetsResponseInner GetAllMarginAssetsResponseInner

// NewGetAllMarginAssetsResponseInner instantiates a new GetAllMarginAssetsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllMarginAssetsResponseInner() *GetAllMarginAssetsResponseInner {
	this := GetAllMarginAssetsResponseInner{}
	return &this
}

// NewGetAllMarginAssetsResponseInnerWithDefaults instantiates a new GetAllMarginAssetsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllMarginAssetsResponseInnerWithDefaults() *GetAllMarginAssetsResponseInner {
	this := GetAllMarginAssetsResponseInner{}
	return &this
}

// GetAssetFullName returns the AssetFullName field value if set, zero value otherwise.
func (o *GetAllMarginAssetsResponseInner) GetAssetFullName() string {
	if o == nil || common.IsNil(o.AssetFullName) {
		var ret string
		return ret
	}
	return *o.AssetFullName
}

// GetAssetFullNameOk returns a tuple with the AssetFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllMarginAssetsResponseInner) GetAssetFullNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.AssetFullName) {
		return nil, false
	}
	return o.AssetFullName, true
}

// HasAssetFullName returns a boolean if a field has been set.
func (o *GetAllMarginAssetsResponseInner) HasAssetFullName() bool {
	if o != nil && !common.IsNil(o.AssetFullName) {
		return true
	}

	return false
}

// SetAssetFullName gets a reference to the given string and assigns it to the AssetFullName field.
func (o *GetAllMarginAssetsResponseInner) SetAssetFullName(v string) {
	o.AssetFullName = &v
}

// GetAssetName returns the AssetName field value if set, zero value otherwise.
func (o *GetAllMarginAssetsResponseInner) GetAssetName() string {
	if o == nil || common.IsNil(o.AssetName) {
		var ret string
		return ret
	}
	return *o.AssetName
}

// GetAssetNameOk returns a tuple with the AssetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllMarginAssetsResponseInner) GetAssetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.AssetName) {
		return nil, false
	}
	return o.AssetName, true
}

// HasAssetName returns a boolean if a field has been set.
func (o *GetAllMarginAssetsResponseInner) HasAssetName() bool {
	if o != nil && !common.IsNil(o.AssetName) {
		return true
	}

	return false
}

// SetAssetName gets a reference to the given string and assigns it to the AssetName field.
func (o *GetAllMarginAssetsResponseInner) SetAssetName(v string) {
	o.AssetName = &v
}

// GetIsBorrowable returns the IsBorrowable field value if set, zero value otherwise.
func (o *GetAllMarginAssetsResponseInner) GetIsBorrowable() bool {
	if o == nil || common.IsNil(o.IsBorrowable) {
		var ret bool
		return ret
	}
	return *o.IsBorrowable
}

// GetIsBorrowableOk returns a tuple with the IsBorrowable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllMarginAssetsResponseInner) GetIsBorrowableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBorrowable) {
		return nil, false
	}
	return o.IsBorrowable, true
}

// HasIsBorrowable returns a boolean if a field has been set.
func (o *GetAllMarginAssetsResponseInner) HasIsBorrowable() bool {
	if o != nil && !common.IsNil(o.IsBorrowable) {
		return true
	}

	return false
}

// SetIsBorrowable gets a reference to the given bool and assigns it to the IsBorrowable field.
func (o *GetAllMarginAssetsResponseInner) SetIsBorrowable(v bool) {
	o.IsBorrowable = &v
}

// GetIsMortgageable returns the IsMortgageable field value if set, zero value otherwise.
func (o *GetAllMarginAssetsResponseInner) GetIsMortgageable() bool {
	if o == nil || common.IsNil(o.IsMortgageable) {
		var ret bool
		return ret
	}
	return *o.IsMortgageable
}

// GetIsMortgageableOk returns a tuple with the IsMortgageable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllMarginAssetsResponseInner) GetIsMortgageableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsMortgageable) {
		return nil, false
	}
	return o.IsMortgageable, true
}

// HasIsMortgageable returns a boolean if a field has been set.
func (o *GetAllMarginAssetsResponseInner) HasIsMortgageable() bool {
	if o != nil && !common.IsNil(o.IsMortgageable) {
		return true
	}

	return false
}

// SetIsMortgageable gets a reference to the given bool and assigns it to the IsMortgageable field.
func (o *GetAllMarginAssetsResponseInner) SetIsMortgageable(v bool) {
	o.IsMortgageable = &v
}

// GetUserMinBorrow returns the UserMinBorrow field value if set, zero value otherwise.
func (o *GetAllMarginAssetsResponseInner) GetUserMinBorrow() string {
	if o == nil || common.IsNil(o.UserMinBorrow) {
		var ret string
		return ret
	}
	return *o.UserMinBorrow
}

// GetUserMinBorrowOk returns a tuple with the UserMinBorrow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllMarginAssetsResponseInner) GetUserMinBorrowOk() (*string, bool) {
	if o == nil || common.IsNil(o.UserMinBorrow) {
		return nil, false
	}
	return o.UserMinBorrow, true
}

// HasUserMinBorrow returns a boolean if a field has been set.
func (o *GetAllMarginAssetsResponseInner) HasUserMinBorrow() bool {
	if o != nil && !common.IsNil(o.UserMinBorrow) {
		return true
	}

	return false
}

// SetUserMinBorrow gets a reference to the given string and assigns it to the UserMinBorrow field.
func (o *GetAllMarginAssetsResponseInner) SetUserMinBorrow(v string) {
	o.UserMinBorrow = &v
}

// GetUserMinRepay returns the UserMinRepay field value if set, zero value otherwise.
func (o *GetAllMarginAssetsResponseInner) GetUserMinRepay() string {
	if o == nil || common.IsNil(o.UserMinRepay) {
		var ret string
		return ret
	}
	return *o.UserMinRepay
}

// GetUserMinRepayOk returns a tuple with the UserMinRepay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllMarginAssetsResponseInner) GetUserMinRepayOk() (*string, bool) {
	if o == nil || common.IsNil(o.UserMinRepay) {
		return nil, false
	}
	return o.UserMinRepay, true
}

// HasUserMinRepay returns a boolean if a field has been set.
func (o *GetAllMarginAssetsResponseInner) HasUserMinRepay() bool {
	if o != nil && !common.IsNil(o.UserMinRepay) {
		return true
	}

	return false
}

// SetUserMinRepay gets a reference to the given string and assigns it to the UserMinRepay field.
func (o *GetAllMarginAssetsResponseInner) SetUserMinRepay(v string) {
	o.UserMinRepay = &v
}

// GetDelistTime returns the DelistTime field value if set, zero value otherwise.
func (o *GetAllMarginAssetsResponseInner) GetDelistTime() int64 {
	if o == nil || common.IsNil(o.DelistTime) {
		var ret int64
		return ret
	}
	return *o.DelistTime
}

// GetDelistTimeOk returns a tuple with the DelistTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllMarginAssetsResponseInner) GetDelistTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.DelistTime) {
		return nil, false
	}
	return o.DelistTime, true
}

// HasDelistTime returns a boolean if a field has been set.
func (o *GetAllMarginAssetsResponseInner) HasDelistTime() bool {
	if o != nil && !common.IsNil(o.DelistTime) {
		return true
	}

	return false
}

// SetDelistTime gets a reference to the given int64 and assigns it to the DelistTime field.
func (o *GetAllMarginAssetsResponseInner) SetDelistTime(v int64) {
	o.DelistTime = &v
}

func (o GetAllMarginAssetsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllMarginAssetsResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AssetFullName) {
		toSerialize["assetFullName"] = o.AssetFullName
	}
	if !common.IsNil(o.AssetName) {
		toSerialize["assetName"] = o.AssetName
	}
	if !common.IsNil(o.IsBorrowable) {
		toSerialize["isBorrowable"] = o.IsBorrowable
	}
	if !common.IsNil(o.IsMortgageable) {
		toSerialize["isMortgageable"] = o.IsMortgageable
	}
	if !common.IsNil(o.UserMinBorrow) {
		toSerialize["userMinBorrow"] = o.UserMinBorrow
	}
	if !common.IsNil(o.UserMinRepay) {
		toSerialize["userMinRepay"] = o.UserMinRepay
	}
	if !common.IsNil(o.DelistTime) {
		toSerialize["delistTime"] = o.DelistTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAllMarginAssetsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetAllMarginAssetsResponseInner := _GetAllMarginAssetsResponseInner{}

	err = json.Unmarshal(data, &varGetAllMarginAssetsResponseInner)

	if err != nil {
		return err
	}

	*o = GetAllMarginAssetsResponseInner(varGetAllMarginAssetsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assetFullName")
		delete(additionalProperties, "assetName")
		delete(additionalProperties, "isBorrowable")
		delete(additionalProperties, "isMortgageable")
		delete(additionalProperties, "userMinBorrow")
		delete(additionalProperties, "userMinRepay")
		delete(additionalProperties, "delistTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAllMarginAssetsResponseInner struct {
	value *GetAllMarginAssetsResponseInner
	isSet bool
}

func (v NullableGetAllMarginAssetsResponseInner) Get() *GetAllMarginAssetsResponseInner {
	return v.value
}

func (v *NullableGetAllMarginAssetsResponseInner) Set(val *GetAllMarginAssetsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllMarginAssetsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllMarginAssetsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllMarginAssetsResponseInner(val *GetAllMarginAssetsResponseInner) *NullableGetAllMarginAssetsResponseInner {
	return &NullableGetAllMarginAssetsResponseInner{value: val, isSet: true}
}

func (v NullableGetAllMarginAssetsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllMarginAssetsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
