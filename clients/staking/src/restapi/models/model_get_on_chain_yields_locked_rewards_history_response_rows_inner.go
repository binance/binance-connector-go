/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOnChainYieldsLockedRewardsHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOnChainYieldsLockedRewardsHistoryResponseRowsInner{}

// GetOnChainYieldsLockedRewardsHistoryResponseRowsInner struct for GetOnChainYieldsLockedRewardsHistoryResponseRowsInner
type GetOnChainYieldsLockedRewardsHistoryResponseRowsInner struct {
	PositionId           *string `json:"positionId,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	LockPeriod           *string `json:"lockPeriod,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOnChainYieldsLockedRewardsHistoryResponseRowsInner GetOnChainYieldsLockedRewardsHistoryResponseRowsInner

// NewGetOnChainYieldsLockedRewardsHistoryResponseRowsInner instantiates a new GetOnChainYieldsLockedRewardsHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOnChainYieldsLockedRewardsHistoryResponseRowsInner() *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner {
	this := GetOnChainYieldsLockedRewardsHistoryResponseRowsInner{}
	return &this
}

// NewGetOnChainYieldsLockedRewardsHistoryResponseRowsInnerWithDefaults instantiates a new GetOnChainYieldsLockedRewardsHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOnChainYieldsLockedRewardsHistoryResponseRowsInnerWithDefaults() *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner {
	this := GetOnChainYieldsLockedRewardsHistoryResponseRowsInner{}
	return &this
}

// GetPositionId returns the PositionId field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) GetPositionId() string {
	if o == nil || common.IsNil(o.PositionId) {
		var ret string
		return ret
	}
	return *o.PositionId
}

// GetPositionIdOk returns a tuple with the PositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) GetPositionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionId) {
		return nil, false
	}
	return o.PositionId, true
}

// HasPositionId returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) HasPositionId() bool {
	if o != nil && !common.IsNil(o.PositionId) {
		return true
	}

	return false
}

// SetPositionId gets a reference to the given string and assigns it to the PositionId field.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) SetPositionId(v string) {
	o.PositionId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetLockPeriod returns the LockPeriod field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) GetLockPeriod() string {
	if o == nil || common.IsNil(o.LockPeriod) {
		var ret string
		return ret
	}
	return *o.LockPeriod
}

// GetLockPeriodOk returns a tuple with the LockPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) GetLockPeriodOk() (*string, bool) {
	if o == nil || common.IsNil(o.LockPeriod) {
		return nil, false
	}
	return o.LockPeriod, true
}

// HasLockPeriod returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) HasLockPeriod() bool {
	if o != nil && !common.IsNil(o.LockPeriod) {
		return true
	}

	return false
}

// SetLockPeriod gets a reference to the given string and assigns it to the LockPeriod field.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) SetLockPeriod(v string) {
	o.LockPeriod = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

func (o GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetOnChainYieldsLockedRewardsHistoryResponseRowsInner := _GetOnChainYieldsLockedRewardsHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetOnChainYieldsLockedRewardsHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetOnChainYieldsLockedRewardsHistoryResponseRowsInner(varGetOnChainYieldsLockedRewardsHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "positionId")
		delete(additionalProperties, "time")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "lockPeriod")
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOnChainYieldsLockedRewardsHistoryResponseRowsInner struct {
	value *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetOnChainYieldsLockedRewardsHistoryResponseRowsInner) Get() *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetOnChainYieldsLockedRewardsHistoryResponseRowsInner) Set(val *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOnChainYieldsLockedRewardsHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOnChainYieldsLockedRewardsHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOnChainYieldsLockedRewardsHistoryResponseRowsInner(val *GetOnChainYieldsLockedRewardsHistoryResponseRowsInner) *NullableGetOnChainYieldsLockedRewardsHistoryResponseRowsInner {
	return &NullableGetOnChainYieldsLockedRewardsHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetOnChainYieldsLockedRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOnChainYieldsLockedRewardsHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
