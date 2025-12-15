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

// checks if the QueryUserDelegationHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUserDelegationHistoryResponseRowsInner{}

// QueryUserDelegationHistoryResponseRowsInner struct for QueryUserDelegationHistoryResponseRowsInner
type QueryUserDelegationHistoryResponseRowsInner struct {
	ClientTranId         *string `json:"clientTranId,omitempty"`
	TransferType         *string `json:"transferType,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryUserDelegationHistoryResponseRowsInner QueryUserDelegationHistoryResponseRowsInner

// NewQueryUserDelegationHistoryResponseRowsInner instantiates a new QueryUserDelegationHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUserDelegationHistoryResponseRowsInner() *QueryUserDelegationHistoryResponseRowsInner {
	this := QueryUserDelegationHistoryResponseRowsInner{}
	return &this
}

// NewQueryUserDelegationHistoryResponseRowsInnerWithDefaults instantiates a new QueryUserDelegationHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUserDelegationHistoryResponseRowsInnerWithDefaults() *QueryUserDelegationHistoryResponseRowsInner {
	this := QueryUserDelegationHistoryResponseRowsInner{}
	return &this
}

// GetClientTranId returns the ClientTranId field value if set, zero value otherwise.
func (o *QueryUserDelegationHistoryResponseRowsInner) GetClientTranId() string {
	if o == nil || common.IsNil(o.ClientTranId) {
		var ret string
		return ret
	}
	return *o.ClientTranId
}

// GetClientTranIdOk returns a tuple with the ClientTranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserDelegationHistoryResponseRowsInner) GetClientTranIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientTranId) {
		return nil, false
	}
	return o.ClientTranId, true
}

// HasClientTranId returns a boolean if a field has been set.
func (o *QueryUserDelegationHistoryResponseRowsInner) HasClientTranId() bool {
	if o != nil && !common.IsNil(o.ClientTranId) {
		return true
	}

	return false
}

// SetClientTranId gets a reference to the given string and assigns it to the ClientTranId field.
func (o *QueryUserDelegationHistoryResponseRowsInner) SetClientTranId(v string) {
	o.ClientTranId = &v
}

// GetTransferType returns the TransferType field value if set, zero value otherwise.
func (o *QueryUserDelegationHistoryResponseRowsInner) GetTransferType() string {
	if o == nil || common.IsNil(o.TransferType) {
		var ret string
		return ret
	}
	return *o.TransferType
}

// GetTransferTypeOk returns a tuple with the TransferType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserDelegationHistoryResponseRowsInner) GetTransferTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferType) {
		return nil, false
	}
	return o.TransferType, true
}

// HasTransferType returns a boolean if a field has been set.
func (o *QueryUserDelegationHistoryResponseRowsInner) HasTransferType() bool {
	if o != nil && !common.IsNil(o.TransferType) {
		return true
	}

	return false
}

// SetTransferType gets a reference to the given string and assigns it to the TransferType field.
func (o *QueryUserDelegationHistoryResponseRowsInner) SetTransferType(v string) {
	o.TransferType = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryUserDelegationHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserDelegationHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryUserDelegationHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryUserDelegationHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QueryUserDelegationHistoryResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserDelegationHistoryResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QueryUserDelegationHistoryResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QueryUserDelegationHistoryResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *QueryUserDelegationHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserDelegationHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *QueryUserDelegationHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *QueryUserDelegationHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

func (o QueryUserDelegationHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUserDelegationHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ClientTranId) {
		toSerialize["clientTranId"] = o.ClientTranId
	}
	if !common.IsNil(o.TransferType) {
		toSerialize["transferType"] = o.TransferType
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryUserDelegationHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryUserDelegationHistoryResponseRowsInner := _QueryUserDelegationHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varQueryUserDelegationHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = QueryUserDelegationHistoryResponseRowsInner(varQueryUserDelegationHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "clientTranId")
		delete(additionalProperties, "transferType")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryUserDelegationHistoryResponseRowsInner struct {
	value *QueryUserDelegationHistoryResponseRowsInner
	isSet bool
}

func (v NullableQueryUserDelegationHistoryResponseRowsInner) Get() *QueryUserDelegationHistoryResponseRowsInner {
	return v.value
}

func (v *NullableQueryUserDelegationHistoryResponseRowsInner) Set(val *QueryUserDelegationHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUserDelegationHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUserDelegationHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUserDelegationHistoryResponseRowsInner(val *QueryUserDelegationHistoryResponseRowsInner) *NullableQueryUserDelegationHistoryResponseRowsInner {
	return &NullableQueryUserDelegationHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableQueryUserDelegationHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUserDelegationHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
