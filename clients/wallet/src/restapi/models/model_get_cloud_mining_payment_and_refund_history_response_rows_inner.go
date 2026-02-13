/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetCloudMiningPaymentAndRefundHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCloudMiningPaymentAndRefundHistoryResponseRowsInner{}

// GetCloudMiningPaymentAndRefundHistoryResponseRowsInner struct for GetCloudMiningPaymentAndRefundHistoryResponseRowsInner
type GetCloudMiningPaymentAndRefundHistoryResponseRowsInner struct {
	CreateTime           *int64  `json:"createTime,omitempty"`
	TranId               *int64  `json:"tranId,omitempty"`
	Type                 *int64  `json:"type,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCloudMiningPaymentAndRefundHistoryResponseRowsInner GetCloudMiningPaymentAndRefundHistoryResponseRowsInner

// NewGetCloudMiningPaymentAndRefundHistoryResponseRowsInner instantiates a new GetCloudMiningPaymentAndRefundHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCloudMiningPaymentAndRefundHistoryResponseRowsInner() *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner {
	this := GetCloudMiningPaymentAndRefundHistoryResponseRowsInner{}
	return &this
}

// NewGetCloudMiningPaymentAndRefundHistoryResponseRowsInnerWithDefaults instantiates a new GetCloudMiningPaymentAndRefundHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCloudMiningPaymentAndRefundHistoryResponseRowsInnerWithDefaults() *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner {
	this := GetCloudMiningPaymentAndRefundHistoryResponseRowsInner{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) SetTranId(v int64) {
	o.TranId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) GetType() int64 {
	if o == nil || common.IsNil(o.Type) {
		var ret int64
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) GetTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int64 and assigns it to the Type field.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) SetType(v int64) {
	o.Type = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetCloudMiningPaymentAndRefundHistoryResponseRowsInner := _GetCloudMiningPaymentAndRefundHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetCloudMiningPaymentAndRefundHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetCloudMiningPaymentAndRefundHistoryResponseRowsInner(varGetCloudMiningPaymentAndRefundHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "tranId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCloudMiningPaymentAndRefundHistoryResponseRowsInner struct {
	value *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetCloudMiningPaymentAndRefundHistoryResponseRowsInner) Get() *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetCloudMiningPaymentAndRefundHistoryResponseRowsInner) Set(val *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCloudMiningPaymentAndRefundHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCloudMiningPaymentAndRefundHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCloudMiningPaymentAndRefundHistoryResponseRowsInner(val *GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) *NullableGetCloudMiningPaymentAndRefundHistoryResponseRowsInner {
	return &NullableGetCloudMiningPaymentAndRefundHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetCloudMiningPaymentAndRefundHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCloudMiningPaymentAndRefundHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
