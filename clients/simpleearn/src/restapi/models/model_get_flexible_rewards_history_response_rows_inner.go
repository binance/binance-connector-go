/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFlexibleRewardsHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleRewardsHistoryResponseRowsInner{}

// GetFlexibleRewardsHistoryResponseRowsInner struct for GetFlexibleRewardsHistoryResponseRowsInner
type GetFlexibleRewardsHistoryResponseRowsInner struct {
	Asset                *string `json:"asset,omitempty"`
	Rewards              *string `json:"rewards,omitempty"`
	ProjectId            *string `json:"projectId,omitempty"`
	Type                 *string `json:"type,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleRewardsHistoryResponseRowsInner GetFlexibleRewardsHistoryResponseRowsInner

// NewGetFlexibleRewardsHistoryResponseRowsInner instantiates a new GetFlexibleRewardsHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleRewardsHistoryResponseRowsInner() *GetFlexibleRewardsHistoryResponseRowsInner {
	this := GetFlexibleRewardsHistoryResponseRowsInner{}
	return &this
}

// NewGetFlexibleRewardsHistoryResponseRowsInnerWithDefaults instantiates a new GetFlexibleRewardsHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleRewardsHistoryResponseRowsInnerWithDefaults() *GetFlexibleRewardsHistoryResponseRowsInner {
	this := GetFlexibleRewardsHistoryResponseRowsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetRewards returns the Rewards field value if set, zero value otherwise.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) GetRewards() string {
	if o == nil || common.IsNil(o.Rewards) {
		var ret string
		return ret
	}
	return *o.Rewards
}

// GetRewardsOk returns a tuple with the Rewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) GetRewardsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Rewards) {
		return nil, false
	}
	return o.Rewards, true
}

// HasRewards returns a boolean if a field has been set.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) HasRewards() bool {
	if o != nil && !common.IsNil(o.Rewards) {
		return true
	}

	return false
}

// SetRewards gets a reference to the given string and assigns it to the Rewards field.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) SetRewards(v string) {
	o.Rewards = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) GetProjectId() string {
	if o == nil || common.IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) GetProjectIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) HasProjectId() bool {
	if o != nil && !common.IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) SetType(v string) {
	o.Type = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetFlexibleRewardsHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

func (o GetFlexibleRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleRewardsHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Rewards) {
		toSerialize["rewards"] = o.Rewards
	}
	if !common.IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFlexibleRewardsHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleRewardsHistoryResponseRowsInner := _GetFlexibleRewardsHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetFlexibleRewardsHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetFlexibleRewardsHistoryResponseRowsInner(varGetFlexibleRewardsHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "rewards")
		delete(additionalProperties, "projectId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleRewardsHistoryResponseRowsInner struct {
	value *GetFlexibleRewardsHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetFlexibleRewardsHistoryResponseRowsInner) Get() *GetFlexibleRewardsHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetFlexibleRewardsHistoryResponseRowsInner) Set(val *GetFlexibleRewardsHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleRewardsHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleRewardsHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleRewardsHistoryResponseRowsInner(val *GetFlexibleRewardsHistoryResponseRowsInner) *NullableGetFlexibleRewardsHistoryResponseRowsInner {
	return &NullableGetFlexibleRewardsHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetFlexibleRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleRewardsHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
