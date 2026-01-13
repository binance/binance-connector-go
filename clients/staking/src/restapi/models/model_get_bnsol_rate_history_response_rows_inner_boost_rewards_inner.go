/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner{}

// GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner struct for GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner
type GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner struct {
	BoostAPR             *string `json:"boostAPR,omitempty"`
	RewardsAsset         *string `json:"rewardsAsset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner

// NewGetBnsolRateHistoryResponseRowsInnerBoostRewardsInner instantiates a new GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBnsolRateHistoryResponseRowsInnerBoostRewardsInner() *GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner {
	this := GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner{}
	return &this
}

// NewGetBnsolRateHistoryResponseRowsInnerBoostRewardsInnerWithDefaults instantiates a new GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBnsolRateHistoryResponseRowsInnerBoostRewardsInnerWithDefaults() *GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner {
	this := GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner{}
	return &this
}

// GetBoostAPR returns the BoostAPR field value if set, zero value otherwise.
func (o *GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) GetBoostAPR() string {
	if o == nil || common.IsNil(o.BoostAPR) {
		var ret string
		return ret
	}
	return *o.BoostAPR
}

// GetBoostAPROk returns a tuple with the BoostAPR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) GetBoostAPROk() (*string, bool) {
	if o == nil || common.IsNil(o.BoostAPR) {
		return nil, false
	}
	return o.BoostAPR, true
}

// HasBoostAPR returns a boolean if a field has been set.
func (o *GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) HasBoostAPR() bool {
	if o != nil && !common.IsNil(o.BoostAPR) {
		return true
	}

	return false
}

// SetBoostAPR gets a reference to the given string and assigns it to the BoostAPR field.
func (o *GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) SetBoostAPR(v string) {
	o.BoostAPR = &v
}

// GetRewardsAsset returns the RewardsAsset field value if set, zero value otherwise.
func (o *GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) GetRewardsAsset() string {
	if o == nil || common.IsNil(o.RewardsAsset) {
		var ret string
		return ret
	}
	return *o.RewardsAsset
}

// GetRewardsAssetOk returns a tuple with the RewardsAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) GetRewardsAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardsAsset) {
		return nil, false
	}
	return o.RewardsAsset, true
}

// HasRewardsAsset returns a boolean if a field has been set.
func (o *GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) HasRewardsAsset() bool {
	if o != nil && !common.IsNil(o.RewardsAsset) {
		return true
	}

	return false
}

// SetRewardsAsset gets a reference to the given string and assigns it to the RewardsAsset field.
func (o *GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) SetRewardsAsset(v string) {
	o.RewardsAsset = &v
}

func (o GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BoostAPR) {
		toSerialize["boostAPR"] = o.BoostAPR
	}
	if !common.IsNil(o.RewardsAsset) {
		toSerialize["rewardsAsset"] = o.RewardsAsset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) UnmarshalJSON(data []byte) (err error) {
	varGetBnsolRateHistoryResponseRowsInnerBoostRewardsInner := _GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner{}

	err = json.Unmarshal(data, &varGetBnsolRateHistoryResponseRowsInnerBoostRewardsInner)

	if err != nil {
		return err
	}

	*o = GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner(varGetBnsolRateHistoryResponseRowsInnerBoostRewardsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "boostAPR")
		delete(additionalProperties, "rewardsAsset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBnsolRateHistoryResponseRowsInnerBoostRewardsInner struct {
	value *GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner
	isSet bool
}

func (v NullableGetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) Get() *GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner {
	return v.value
}

func (v *NullableGetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) Set(val *GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBnsolRateHistoryResponseRowsInnerBoostRewardsInner(val *GetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) *NullableGetBnsolRateHistoryResponseRowsInnerBoostRewardsInner {
	return &NullableGetBnsolRateHistoryResponseRowsInnerBoostRewardsInner{value: val, isSet: true}
}

func (v NullableGetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBnsolRateHistoryResponseRowsInnerBoostRewardsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
