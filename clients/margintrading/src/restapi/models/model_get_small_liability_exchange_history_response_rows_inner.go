/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSmallLiabilityExchangeHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSmallLiabilityExchangeHistoryResponseRowsInner{}

// GetSmallLiabilityExchangeHistoryResponseRowsInner struct for GetSmallLiabilityExchangeHistoryResponseRowsInner
type GetSmallLiabilityExchangeHistoryResponseRowsInner struct {
	Asset                *string `json:"asset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	TargetAsset          *string `json:"targetAsset,omitempty"`
	TargetAmount         *string `json:"targetAmount,omitempty"`
	BizType              *string `json:"bizType,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSmallLiabilityExchangeHistoryResponseRowsInner GetSmallLiabilityExchangeHistoryResponseRowsInner

// NewGetSmallLiabilityExchangeHistoryResponseRowsInner instantiates a new GetSmallLiabilityExchangeHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSmallLiabilityExchangeHistoryResponseRowsInner() *GetSmallLiabilityExchangeHistoryResponseRowsInner {
	this := GetSmallLiabilityExchangeHistoryResponseRowsInner{}
	return &this
}

// NewGetSmallLiabilityExchangeHistoryResponseRowsInnerWithDefaults instantiates a new GetSmallLiabilityExchangeHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSmallLiabilityExchangeHistoryResponseRowsInnerWithDefaults() *GetSmallLiabilityExchangeHistoryResponseRowsInner {
	this := GetSmallLiabilityExchangeHistoryResponseRowsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetTargetAsset returns the TargetAsset field value if set, zero value otherwise.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) GetTargetAsset() string {
	if o == nil || common.IsNil(o.TargetAsset) {
		var ret string
		return ret
	}
	return *o.TargetAsset
}

// GetTargetAssetOk returns a tuple with the TargetAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) GetTargetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.TargetAsset) {
		return nil, false
	}
	return o.TargetAsset, true
}

// HasTargetAsset returns a boolean if a field has been set.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) HasTargetAsset() bool {
	if o != nil && !common.IsNil(o.TargetAsset) {
		return true
	}

	return false
}

// SetTargetAsset gets a reference to the given string and assigns it to the TargetAsset field.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) SetTargetAsset(v string) {
	o.TargetAsset = &v
}

// GetTargetAmount returns the TargetAmount field value if set, zero value otherwise.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) GetTargetAmount() string {
	if o == nil || common.IsNil(o.TargetAmount) {
		var ret string
		return ret
	}
	return *o.TargetAmount
}

// GetTargetAmountOk returns a tuple with the TargetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) GetTargetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TargetAmount) {
		return nil, false
	}
	return o.TargetAmount, true
}

// HasTargetAmount returns a boolean if a field has been set.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) HasTargetAmount() bool {
	if o != nil && !common.IsNil(o.TargetAmount) {
		return true
	}

	return false
}

// SetTargetAmount gets a reference to the given string and assigns it to the TargetAmount field.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) SetTargetAmount(v string) {
	o.TargetAmount = &v
}

// GetBizType returns the BizType field value if set, zero value otherwise.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) GetBizType() string {
	if o == nil || common.IsNil(o.BizType) {
		var ret string
		return ret
	}
	return *o.BizType
}

// GetBizTypeOk returns a tuple with the BizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) GetBizTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.BizType) {
		return nil, false
	}
	return o.BizType, true
}

// HasBizType returns a boolean if a field has been set.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) HasBizType() bool {
	if o != nil && !common.IsNil(o.BizType) {
		return true
	}

	return false
}

// SetBizType gets a reference to the given string and assigns it to the BizType field.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) SetBizType(v string) {
	o.BizType = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o GetSmallLiabilityExchangeHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSmallLiabilityExchangeHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.TargetAsset) {
		toSerialize["targetAsset"] = o.TargetAsset
	}
	if !common.IsNil(o.TargetAmount) {
		toSerialize["targetAmount"] = o.TargetAmount
	}
	if !common.IsNil(o.BizType) {
		toSerialize["bizType"] = o.BizType
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSmallLiabilityExchangeHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetSmallLiabilityExchangeHistoryResponseRowsInner := _GetSmallLiabilityExchangeHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetSmallLiabilityExchangeHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetSmallLiabilityExchangeHistoryResponseRowsInner(varGetSmallLiabilityExchangeHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "targetAsset")
		delete(additionalProperties, "targetAmount")
		delete(additionalProperties, "bizType")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSmallLiabilityExchangeHistoryResponseRowsInner struct {
	value *GetSmallLiabilityExchangeHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetSmallLiabilityExchangeHistoryResponseRowsInner) Get() *GetSmallLiabilityExchangeHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetSmallLiabilityExchangeHistoryResponseRowsInner) Set(val *GetSmallLiabilityExchangeHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSmallLiabilityExchangeHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSmallLiabilityExchangeHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSmallLiabilityExchangeHistoryResponseRowsInner(val *GetSmallLiabilityExchangeHistoryResponseRowsInner) *NullableGetSmallLiabilityExchangeHistoryResponseRowsInner {
	return &NullableGetSmallLiabilityExchangeHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetSmallLiabilityExchangeHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSmallLiabilityExchangeHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
