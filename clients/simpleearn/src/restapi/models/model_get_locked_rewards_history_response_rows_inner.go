/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetLockedRewardsHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLockedRewardsHistoryResponseRowsInner{}

// GetLockedRewardsHistoryResponseRowsInner struct for GetLockedRewardsHistoryResponseRowsInner
type GetLockedRewardsHistoryResponseRowsInner struct {
	PositionId           *int64  `json:"positionId,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	LockPeriod           *string `json:"lockPeriod,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLockedRewardsHistoryResponseRowsInner GetLockedRewardsHistoryResponseRowsInner

// NewGetLockedRewardsHistoryResponseRowsInner instantiates a new GetLockedRewardsHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLockedRewardsHistoryResponseRowsInner() *GetLockedRewardsHistoryResponseRowsInner {
	this := GetLockedRewardsHistoryResponseRowsInner{}
	return &this
}

// NewGetLockedRewardsHistoryResponseRowsInnerWithDefaults instantiates a new GetLockedRewardsHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLockedRewardsHistoryResponseRowsInnerWithDefaults() *GetLockedRewardsHistoryResponseRowsInner {
	this := GetLockedRewardsHistoryResponseRowsInner{}
	return &this
}

// GetPositionId returns the PositionId field value if set, zero value otherwise.
func (o *GetLockedRewardsHistoryResponseRowsInner) GetPositionId() int64 {
	if o == nil || common.IsNil(o.PositionId) {
		var ret int64
		return ret
	}
	return *o.PositionId
}

// GetPositionIdOk returns a tuple with the PositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRewardsHistoryResponseRowsInner) GetPositionIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PositionId) {
		return nil, false
	}
	return o.PositionId, true
}

// HasPositionId returns a boolean if a field has been set.
func (o *GetLockedRewardsHistoryResponseRowsInner) HasPositionId() bool {
	if o != nil && !common.IsNil(o.PositionId) {
		return true
	}

	return false
}

// SetPositionId gets a reference to the given int64 and assigns it to the PositionId field.
func (o *GetLockedRewardsHistoryResponseRowsInner) SetPositionId(v int64) {
	o.PositionId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetLockedRewardsHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRewardsHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetLockedRewardsHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetLockedRewardsHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetLockedRewardsHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRewardsHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetLockedRewardsHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetLockedRewardsHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetLockPeriod returns the LockPeriod field value if set, zero value otherwise.
func (o *GetLockedRewardsHistoryResponseRowsInner) GetLockPeriod() string {
	if o == nil || common.IsNil(o.LockPeriod) {
		var ret string
		return ret
	}
	return *o.LockPeriod
}

// GetLockPeriodOk returns a tuple with the LockPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRewardsHistoryResponseRowsInner) GetLockPeriodOk() (*string, bool) {
	if o == nil || common.IsNil(o.LockPeriod) {
		return nil, false
	}
	return o.LockPeriod, true
}

// HasLockPeriod returns a boolean if a field has been set.
func (o *GetLockedRewardsHistoryResponseRowsInner) HasLockPeriod() bool {
	if o != nil && !common.IsNil(o.LockPeriod) {
		return true
	}

	return false
}

// SetLockPeriod gets a reference to the given string and assigns it to the LockPeriod field.
func (o *GetLockedRewardsHistoryResponseRowsInner) SetLockPeriod(v string) {
	o.LockPeriod = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetLockedRewardsHistoryResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRewardsHistoryResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetLockedRewardsHistoryResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetLockedRewardsHistoryResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetLockedRewardsHistoryResponseRowsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRewardsHistoryResponseRowsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetLockedRewardsHistoryResponseRowsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetLockedRewardsHistoryResponseRowsInner) SetType(v string) {
	o.Type = &v
}

func (o GetLockedRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLockedRewardsHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PositionId) {
		toSerialize["positionId"] = o.PositionId
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.LockPeriod) {
		toSerialize["lockPeriod"] = o.LockPeriod
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetLockedRewardsHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetLockedRewardsHistoryResponseRowsInner := _GetLockedRewardsHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetLockedRewardsHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetLockedRewardsHistoryResponseRowsInner(varGetLockedRewardsHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "positionId")
		delete(additionalProperties, "time")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "lockPeriod")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLockedRewardsHistoryResponseRowsInner struct {
	value *GetLockedRewardsHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetLockedRewardsHistoryResponseRowsInner) Get() *GetLockedRewardsHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetLockedRewardsHistoryResponseRowsInner) Set(val *GetLockedRewardsHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLockedRewardsHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLockedRewardsHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLockedRewardsHistoryResponseRowsInner(val *GetLockedRewardsHistoryResponseRowsInner) *NullableGetLockedRewardsHistoryResponseRowsInner {
	return &NullableGetLockedRewardsHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetLockedRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLockedRewardsHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
