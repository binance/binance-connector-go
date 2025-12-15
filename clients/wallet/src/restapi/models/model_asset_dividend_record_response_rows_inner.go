/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AssetDividendRecordResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AssetDividendRecordResponseRowsInner{}

// AssetDividendRecordResponseRowsInner struct for AssetDividendRecordResponseRowsInner
type AssetDividendRecordResponseRowsInner struct {
	Id                   *int64  `json:"id,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	DivTime              *int64  `json:"divTime,omitempty"`
	EnInfo               *string `json:"enInfo,omitempty"`
	TranId               *int64  `json:"tranId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetDividendRecordResponseRowsInner AssetDividendRecordResponseRowsInner

// NewAssetDividendRecordResponseRowsInner instantiates a new AssetDividendRecordResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDividendRecordResponseRowsInner() *AssetDividendRecordResponseRowsInner {
	this := AssetDividendRecordResponseRowsInner{}
	return &this
}

// NewAssetDividendRecordResponseRowsInnerWithDefaults instantiates a new AssetDividendRecordResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDividendRecordResponseRowsInnerWithDefaults() *AssetDividendRecordResponseRowsInner {
	this := AssetDividendRecordResponseRowsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssetDividendRecordResponseRowsInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDividendRecordResponseRowsInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssetDividendRecordResponseRowsInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AssetDividendRecordResponseRowsInner) SetId(v int64) {
	o.Id = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AssetDividendRecordResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDividendRecordResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AssetDividendRecordResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AssetDividendRecordResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *AssetDividendRecordResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDividendRecordResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *AssetDividendRecordResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *AssetDividendRecordResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetDivTime returns the DivTime field value if set, zero value otherwise.
func (o *AssetDividendRecordResponseRowsInner) GetDivTime() int64 {
	if o == nil || common.IsNil(o.DivTime) {
		var ret int64
		return ret
	}
	return *o.DivTime
}

// GetDivTimeOk returns a tuple with the DivTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDividendRecordResponseRowsInner) GetDivTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.DivTime) {
		return nil, false
	}
	return o.DivTime, true
}

// HasDivTime returns a boolean if a field has been set.
func (o *AssetDividendRecordResponseRowsInner) HasDivTime() bool {
	if o != nil && !common.IsNil(o.DivTime) {
		return true
	}

	return false
}

// SetDivTime gets a reference to the given int64 and assigns it to the DivTime field.
func (o *AssetDividendRecordResponseRowsInner) SetDivTime(v int64) {
	o.DivTime = &v
}

// GetEnInfo returns the EnInfo field value if set, zero value otherwise.
func (o *AssetDividendRecordResponseRowsInner) GetEnInfo() string {
	if o == nil || common.IsNil(o.EnInfo) {
		var ret string
		return ret
	}
	return *o.EnInfo
}

// GetEnInfoOk returns a tuple with the EnInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDividendRecordResponseRowsInner) GetEnInfoOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnInfo) {
		return nil, false
	}
	return o.EnInfo, true
}

// HasEnInfo returns a boolean if a field has been set.
func (o *AssetDividendRecordResponseRowsInner) HasEnInfo() bool {
	if o != nil && !common.IsNil(o.EnInfo) {
		return true
	}

	return false
}

// SetEnInfo gets a reference to the given string and assigns it to the EnInfo field.
func (o *AssetDividendRecordResponseRowsInner) SetEnInfo(v string) {
	o.EnInfo = &v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *AssetDividendRecordResponseRowsInner) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDividendRecordResponseRowsInner) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *AssetDividendRecordResponseRowsInner) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *AssetDividendRecordResponseRowsInner) SetTranId(v int64) {
	o.TranId = &v
}

func (o AssetDividendRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetDividendRecordResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.DivTime) {
		toSerialize["divTime"] = o.DivTime
	}
	if !common.IsNil(o.EnInfo) {
		toSerialize["enInfo"] = o.EnInfo
	}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetDividendRecordResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varAssetDividendRecordResponseRowsInner := _AssetDividendRecordResponseRowsInner{}

	err = json.Unmarshal(data, &varAssetDividendRecordResponseRowsInner)

	if err != nil {
		return err
	}

	*o = AssetDividendRecordResponseRowsInner(varAssetDividendRecordResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "divTime")
		delete(additionalProperties, "enInfo")
		delete(additionalProperties, "tranId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetDividendRecordResponseRowsInner struct {
	value *AssetDividendRecordResponseRowsInner
	isSet bool
}

func (v NullableAssetDividendRecordResponseRowsInner) Get() *AssetDividendRecordResponseRowsInner {
	return v.value
}

func (v *NullableAssetDividendRecordResponseRowsInner) Set(val *AssetDividendRecordResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDividendRecordResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDividendRecordResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDividendRecordResponseRowsInner(val *AssetDividendRecordResponseRowsInner) *NullableAssetDividendRecordResponseRowsInner {
	return &NullableAssetDividendRecordResponseRowsInner{value: val, isSet: true}
}

func (v NullableAssetDividendRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDividendRecordResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
