/*
Binance Rebate REST API

OpenAPI Specification for the Binance Rebate REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSpotRebateHistoryRecordsResponseDataDataInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSpotRebateHistoryRecordsResponseDataDataInner{}

// GetSpotRebateHistoryRecordsResponseDataDataInner struct for GetSpotRebateHistoryRecordsResponseDataDataInner
type GetSpotRebateHistoryRecordsResponseDataDataInner struct {
	Asset                *string `json:"asset,omitempty"`
	Type                 *int64  `json:"type,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSpotRebateHistoryRecordsResponseDataDataInner GetSpotRebateHistoryRecordsResponseDataDataInner

// NewGetSpotRebateHistoryRecordsResponseDataDataInner instantiates a new GetSpotRebateHistoryRecordsResponseDataDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSpotRebateHistoryRecordsResponseDataDataInner() *GetSpotRebateHistoryRecordsResponseDataDataInner {
	this := GetSpotRebateHistoryRecordsResponseDataDataInner{}
	return &this
}

// NewGetSpotRebateHistoryRecordsResponseDataDataInnerWithDefaults instantiates a new GetSpotRebateHistoryRecordsResponseDataDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSpotRebateHistoryRecordsResponseDataDataInnerWithDefaults() *GetSpotRebateHistoryRecordsResponseDataDataInner {
	this := GetSpotRebateHistoryRecordsResponseDataDataInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) SetAsset(v string) {
	o.Asset = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) GetType() int64 {
	if o == nil || common.IsNil(o.Type) {
		var ret int64
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) GetTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int64 and assigns it to the Type field.
func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) SetType(v int64) {
	o.Type = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) SetAmount(v string) {
	o.Amount = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o GetSpotRebateHistoryRecordsResponseDataDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSpotRebateHistoryRecordsResponseDataDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSpotRebateHistoryRecordsResponseDataDataInner) UnmarshalJSON(data []byte) (err error) {
	varGetSpotRebateHistoryRecordsResponseDataDataInner := _GetSpotRebateHistoryRecordsResponseDataDataInner{}

	err = json.Unmarshal(data, &varGetSpotRebateHistoryRecordsResponseDataDataInner)

	if err != nil {
		return err
	}

	*o = GetSpotRebateHistoryRecordsResponseDataDataInner(varGetSpotRebateHistoryRecordsResponseDataDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "type")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSpotRebateHistoryRecordsResponseDataDataInner struct {
	value *GetSpotRebateHistoryRecordsResponseDataDataInner
	isSet bool
}

func (v NullableGetSpotRebateHistoryRecordsResponseDataDataInner) Get() *GetSpotRebateHistoryRecordsResponseDataDataInner {
	return v.value
}

func (v *NullableGetSpotRebateHistoryRecordsResponseDataDataInner) Set(val *GetSpotRebateHistoryRecordsResponseDataDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSpotRebateHistoryRecordsResponseDataDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSpotRebateHistoryRecordsResponseDataDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSpotRebateHistoryRecordsResponseDataDataInner(val *GetSpotRebateHistoryRecordsResponseDataDataInner) *NullableGetSpotRebateHistoryRecordsResponseDataDataInner {
	return &NullableGetSpotRebateHistoryRecordsResponseDataDataInner{value: val, isSet: true}
}

func (v NullableGetSpotRebateHistoryRecordsResponseDataDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSpotRebateHistoryRecordsResponseDataDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
