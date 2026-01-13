/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetUnclaimedRewardsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetUnclaimedRewardsResponseInner{}

// GetUnclaimedRewardsResponseInner struct for GetUnclaimedRewardsResponseInner
type GetUnclaimedRewardsResponseInner struct {
	Amount               *string `json:"amount,omitempty"`
	RewardsAsset         *string `json:"rewardsAsset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUnclaimedRewardsResponseInner GetUnclaimedRewardsResponseInner

// NewGetUnclaimedRewardsResponseInner instantiates a new GetUnclaimedRewardsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUnclaimedRewardsResponseInner() *GetUnclaimedRewardsResponseInner {
	this := GetUnclaimedRewardsResponseInner{}
	return &this
}

// NewGetUnclaimedRewardsResponseInnerWithDefaults instantiates a new GetUnclaimedRewardsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUnclaimedRewardsResponseInnerWithDefaults() *GetUnclaimedRewardsResponseInner {
	this := GetUnclaimedRewardsResponseInner{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetUnclaimedRewardsResponseInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUnclaimedRewardsResponseInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetUnclaimedRewardsResponseInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetUnclaimedRewardsResponseInner) SetAmount(v string) {
	o.Amount = &v
}

// GetRewardsAsset returns the RewardsAsset field value if set, zero value otherwise.
func (o *GetUnclaimedRewardsResponseInner) GetRewardsAsset() string {
	if o == nil || common.IsNil(o.RewardsAsset) {
		var ret string
		return ret
	}
	return *o.RewardsAsset
}

// GetRewardsAssetOk returns a tuple with the RewardsAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUnclaimedRewardsResponseInner) GetRewardsAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardsAsset) {
		return nil, false
	}
	return o.RewardsAsset, true
}

// HasRewardsAsset returns a boolean if a field has been set.
func (o *GetUnclaimedRewardsResponseInner) HasRewardsAsset() bool {
	if o != nil && !common.IsNil(o.RewardsAsset) {
		return true
	}

	return false
}

// SetRewardsAsset gets a reference to the given string and assigns it to the RewardsAsset field.
func (o *GetUnclaimedRewardsResponseInner) SetRewardsAsset(v string) {
	o.RewardsAsset = &v
}

func (o GetUnclaimedRewardsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUnclaimedRewardsResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.RewardsAsset) {
		toSerialize["rewardsAsset"] = o.RewardsAsset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUnclaimedRewardsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetUnclaimedRewardsResponseInner := _GetUnclaimedRewardsResponseInner{}

	err = json.Unmarshal(data, &varGetUnclaimedRewardsResponseInner)

	if err != nil {
		return err
	}

	*o = GetUnclaimedRewardsResponseInner(varGetUnclaimedRewardsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "rewardsAsset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUnclaimedRewardsResponseInner struct {
	value *GetUnclaimedRewardsResponseInner
	isSet bool
}

func (v NullableGetUnclaimedRewardsResponseInner) Get() *GetUnclaimedRewardsResponseInner {
	return v.value
}

func (v *NullableGetUnclaimedRewardsResponseInner) Set(val *GetUnclaimedRewardsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUnclaimedRewardsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUnclaimedRewardsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUnclaimedRewardsResponseInner(val *GetUnclaimedRewardsResponseInner) *NullableGetUnclaimedRewardsResponseInner {
	return &NullableGetUnclaimedRewardsResponseInner{value: val, isSet: true}
}

func (v NullableGetUnclaimedRewardsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUnclaimedRewardsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
