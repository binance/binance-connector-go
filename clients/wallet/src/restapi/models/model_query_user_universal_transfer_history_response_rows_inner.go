/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryUserUniversalTransferHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUserUniversalTransferHistoryResponseRowsInner{}

// QueryUserUniversalTransferHistoryResponseRowsInner struct for QueryUserUniversalTransferHistoryResponseRowsInner
type QueryUserUniversalTransferHistoryResponseRowsInner struct {
	Asset                *string `json:"asset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	Type                 *string `json:"type,omitempty"`
	Status               *string `json:"status,omitempty"`
	TranId               *int64  `json:"tranId,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryUserUniversalTransferHistoryResponseRowsInner QueryUserUniversalTransferHistoryResponseRowsInner

// NewQueryUserUniversalTransferHistoryResponseRowsInner instantiates a new QueryUserUniversalTransferHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUserUniversalTransferHistoryResponseRowsInner() *QueryUserUniversalTransferHistoryResponseRowsInner {
	this := QueryUserUniversalTransferHistoryResponseRowsInner{}
	return &this
}

// NewQueryUserUniversalTransferHistoryResponseRowsInnerWithDefaults instantiates a new QueryUserUniversalTransferHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUserUniversalTransferHistoryResponseRowsInnerWithDefaults() *QueryUserUniversalTransferHistoryResponseRowsInner {
	this := QueryUserUniversalTransferHistoryResponseRowsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) SetTranId(v int64) {
	o.TranId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *QueryUserUniversalTransferHistoryResponseRowsInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o QueryUserUniversalTransferHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUserUniversalTransferHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryUserUniversalTransferHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryUserUniversalTransferHistoryResponseRowsInner := _QueryUserUniversalTransferHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varQueryUserUniversalTransferHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = QueryUserUniversalTransferHistoryResponseRowsInner(varQueryUserUniversalTransferHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tranId")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryUserUniversalTransferHistoryResponseRowsInner struct {
	value *QueryUserUniversalTransferHistoryResponseRowsInner
	isSet bool
}

func (v NullableQueryUserUniversalTransferHistoryResponseRowsInner) Get() *QueryUserUniversalTransferHistoryResponseRowsInner {
	return v.value
}

func (v *NullableQueryUserUniversalTransferHistoryResponseRowsInner) Set(val *QueryUserUniversalTransferHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUserUniversalTransferHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUserUniversalTransferHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUserUniversalTransferHistoryResponseRowsInner(val *QueryUserUniversalTransferHistoryResponseRowsInner) *NullableQueryUserUniversalTransferHistoryResponseRowsInner {
	return &NullableQueryUserUniversalTransferHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableQueryUserUniversalTransferHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUserUniversalTransferHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
