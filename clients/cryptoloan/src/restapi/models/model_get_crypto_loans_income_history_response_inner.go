/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetCryptoLoansIncomeHistoryResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCryptoLoansIncomeHistoryResponseInner{}

// GetCryptoLoansIncomeHistoryResponseInner struct for GetCryptoLoansIncomeHistoryResponseInner
type GetCryptoLoansIncomeHistoryResponseInner struct {
	Asset                *string `json:"asset,omitempty"`
	Type                 *string `json:"type,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	TranId               *string `json:"tranId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCryptoLoansIncomeHistoryResponseInner GetCryptoLoansIncomeHistoryResponseInner

// NewGetCryptoLoansIncomeHistoryResponseInner instantiates a new GetCryptoLoansIncomeHistoryResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCryptoLoansIncomeHistoryResponseInner() *GetCryptoLoansIncomeHistoryResponseInner {
	this := GetCryptoLoansIncomeHistoryResponseInner{}
	return &this
}

// NewGetCryptoLoansIncomeHistoryResponseInnerWithDefaults instantiates a new GetCryptoLoansIncomeHistoryResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCryptoLoansIncomeHistoryResponseInnerWithDefaults() *GetCryptoLoansIncomeHistoryResponseInner {
	this := GetCryptoLoansIncomeHistoryResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetCryptoLoansIncomeHistoryResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCryptoLoansIncomeHistoryResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetCryptoLoansIncomeHistoryResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetCryptoLoansIncomeHistoryResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetCryptoLoansIncomeHistoryResponseInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCryptoLoansIncomeHistoryResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetCryptoLoansIncomeHistoryResponseInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetCryptoLoansIncomeHistoryResponseInner) SetType(v string) {
	o.Type = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetCryptoLoansIncomeHistoryResponseInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCryptoLoansIncomeHistoryResponseInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetCryptoLoansIncomeHistoryResponseInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetCryptoLoansIncomeHistoryResponseInner) SetAmount(v string) {
	o.Amount = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetCryptoLoansIncomeHistoryResponseInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCryptoLoansIncomeHistoryResponseInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetCryptoLoansIncomeHistoryResponseInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *GetCryptoLoansIncomeHistoryResponseInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *GetCryptoLoansIncomeHistoryResponseInner) GetTranId() string {
	if o == nil || common.IsNil(o.TranId) {
		var ret string
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCryptoLoansIncomeHistoryResponseInner) GetTranIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *GetCryptoLoansIncomeHistoryResponseInner) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given string and assigns it to the TranId field.
func (o *GetCryptoLoansIncomeHistoryResponseInner) SetTranId(v string) {
	o.TranId = &v
}

func (o GetCryptoLoansIncomeHistoryResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCryptoLoansIncomeHistoryResponseInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCryptoLoansIncomeHistoryResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetCryptoLoansIncomeHistoryResponseInner := _GetCryptoLoansIncomeHistoryResponseInner{}

	err = json.Unmarshal(data, &varGetCryptoLoansIncomeHistoryResponseInner)

	if err != nil {
		return err
	}

	*o = GetCryptoLoansIncomeHistoryResponseInner(varGetCryptoLoansIncomeHistoryResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "type")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "tranId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCryptoLoansIncomeHistoryResponseInner struct {
	value *GetCryptoLoansIncomeHistoryResponseInner
	isSet bool
}

func (v NullableGetCryptoLoansIncomeHistoryResponseInner) Get() *GetCryptoLoansIncomeHistoryResponseInner {
	return v.value
}

func (v *NullableGetCryptoLoansIncomeHistoryResponseInner) Set(val *GetCryptoLoansIncomeHistoryResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCryptoLoansIncomeHistoryResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCryptoLoansIncomeHistoryResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCryptoLoansIncomeHistoryResponseInner(val *GetCryptoLoansIncomeHistoryResponseInner) *NullableGetCryptoLoansIncomeHistoryResponseInner {
	return &NullableGetCryptoLoansIncomeHistoryResponseInner{value: val, isSet: true}
}

func (v NullableGetCryptoLoansIncomeHistoryResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCryptoLoansIncomeHistoryResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
