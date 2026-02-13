/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOnChainYieldsLockedProductListResponseRowsInnerQuota type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOnChainYieldsLockedProductListResponseRowsInnerQuota{}

// GetOnChainYieldsLockedProductListResponseRowsInnerQuota struct for GetOnChainYieldsLockedProductListResponseRowsInnerQuota
type GetOnChainYieldsLockedProductListResponseRowsInnerQuota struct {
	TotalPersonalQuota   *string `json:"totalPersonalQuota,omitempty"`
	Minimum              *string `json:"minimum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOnChainYieldsLockedProductListResponseRowsInnerQuota GetOnChainYieldsLockedProductListResponseRowsInnerQuota

// NewGetOnChainYieldsLockedProductListResponseRowsInnerQuota instantiates a new GetOnChainYieldsLockedProductListResponseRowsInnerQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOnChainYieldsLockedProductListResponseRowsInnerQuota() *GetOnChainYieldsLockedProductListResponseRowsInnerQuota {
	this := GetOnChainYieldsLockedProductListResponseRowsInnerQuota{}
	return &this
}

// NewGetOnChainYieldsLockedProductListResponseRowsInnerQuotaWithDefaults instantiates a new GetOnChainYieldsLockedProductListResponseRowsInnerQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOnChainYieldsLockedProductListResponseRowsInnerQuotaWithDefaults() *GetOnChainYieldsLockedProductListResponseRowsInnerQuota {
	this := GetOnChainYieldsLockedProductListResponseRowsInnerQuota{}
	return &this
}

// GetTotalPersonalQuota returns the TotalPersonalQuota field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerQuota) GetTotalPersonalQuota() string {
	if o == nil || common.IsNil(o.TotalPersonalQuota) {
		var ret string
		return ret
	}
	return *o.TotalPersonalQuota
}

// GetTotalPersonalQuotaOk returns a tuple with the TotalPersonalQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerQuota) GetTotalPersonalQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalPersonalQuota) {
		return nil, false
	}
	return o.TotalPersonalQuota, true
}

// HasTotalPersonalQuota returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerQuota) HasTotalPersonalQuota() bool {
	if o != nil && !common.IsNil(o.TotalPersonalQuota) {
		return true
	}

	return false
}

// SetTotalPersonalQuota gets a reference to the given string and assigns it to the TotalPersonalQuota field.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerQuota) SetTotalPersonalQuota(v string) {
	o.TotalPersonalQuota = &v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerQuota) GetMinimum() string {
	if o == nil || common.IsNil(o.Minimum) {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerQuota) GetMinimumOk() (*string, bool) {
	if o == nil || common.IsNil(o.Minimum) {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerQuota) HasMinimum() bool {
	if o != nil && !common.IsNil(o.Minimum) {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerQuota) SetMinimum(v string) {
	o.Minimum = &v
}

func (o GetOnChainYieldsLockedProductListResponseRowsInnerQuota) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOnChainYieldsLockedProductListResponseRowsInnerQuota) ToMap() (map[string]interface{}, error) {
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

func (o *GetOnChainYieldsLockedProductListResponseRowsInnerQuota) UnmarshalJSON(data []byte) (err error) {
	varGetOnChainYieldsLockedProductListResponseRowsInnerQuota := _GetOnChainYieldsLockedProductListResponseRowsInnerQuota{}

	err = json.Unmarshal(data, &varGetOnChainYieldsLockedProductListResponseRowsInnerQuota)

	if err != nil {
		return err
	}

	*o = GetOnChainYieldsLockedProductListResponseRowsInnerQuota(varGetOnChainYieldsLockedProductListResponseRowsInnerQuota)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "totalPersonalQuota")
		delete(additionalProperties, "minimum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOnChainYieldsLockedProductListResponseRowsInnerQuota struct {
	value *GetOnChainYieldsLockedProductListResponseRowsInnerQuota
	isSet bool
}

func (v NullableGetOnChainYieldsLockedProductListResponseRowsInnerQuota) Get() *GetOnChainYieldsLockedProductListResponseRowsInnerQuota {
	return v.value
}

func (v *NullableGetOnChainYieldsLockedProductListResponseRowsInnerQuota) Set(val *GetOnChainYieldsLockedProductListResponseRowsInnerQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOnChainYieldsLockedProductListResponseRowsInnerQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOnChainYieldsLockedProductListResponseRowsInnerQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOnChainYieldsLockedProductListResponseRowsInnerQuota(val *GetOnChainYieldsLockedProductListResponseRowsInnerQuota) *NullableGetOnChainYieldsLockedProductListResponseRowsInnerQuota {
	return &NullableGetOnChainYieldsLockedProductListResponseRowsInnerQuota{value: val, isSet: true}
}

func (v NullableGetOnChainYieldsLockedProductListResponseRowsInnerQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOnChainYieldsLockedProductListResponseRowsInnerQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
