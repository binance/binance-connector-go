/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetInterestHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetInterestHistoryResponseRowsInner{}

// GetInterestHistoryResponseRowsInner struct for GetInterestHistoryResponseRowsInner
type GetInterestHistoryResponseRowsInner struct {
	TxId                 *int64  `json:"txId,omitempty"`
	InterestAccuredTime  *int64  `json:"interestAccuredTime,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	RawAsset             *string `json:"rawAsset,omitempty"`
	Principal            *string `json:"principal,omitempty"`
	Interest             *string `json:"interest,omitempty"`
	InterestRate         *string `json:"interestRate,omitempty"`
	Type                 *string `json:"type,omitempty"`
	IsolatedSymbol       *string `json:"isolatedSymbol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetInterestHistoryResponseRowsInner GetInterestHistoryResponseRowsInner

// NewGetInterestHistoryResponseRowsInner instantiates a new GetInterestHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInterestHistoryResponseRowsInner() *GetInterestHistoryResponseRowsInner {
	this := GetInterestHistoryResponseRowsInner{}
	return &this
}

// NewGetInterestHistoryResponseRowsInnerWithDefaults instantiates a new GetInterestHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInterestHistoryResponseRowsInnerWithDefaults() *GetInterestHistoryResponseRowsInner {
	this := GetInterestHistoryResponseRowsInner{}
	return &this
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *GetInterestHistoryResponseRowsInner) GetTxId() int64 {
	if o == nil || common.IsNil(o.TxId) {
		var ret int64
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterestHistoryResponseRowsInner) GetTxIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *GetInterestHistoryResponseRowsInner) HasTxId() bool {
	if o != nil && !common.IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given int64 and assigns it to the TxId field.
func (o *GetInterestHistoryResponseRowsInner) SetTxId(v int64) {
	o.TxId = &v
}

// GetInterestAccuredTime returns the InterestAccuredTime field value if set, zero value otherwise.
func (o *GetInterestHistoryResponseRowsInner) GetInterestAccuredTime() int64 {
	if o == nil || common.IsNil(o.InterestAccuredTime) {
		var ret int64
		return ret
	}
	return *o.InterestAccuredTime
}

// GetInterestAccuredTimeOk returns a tuple with the InterestAccuredTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterestHistoryResponseRowsInner) GetInterestAccuredTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.InterestAccuredTime) {
		return nil, false
	}
	return o.InterestAccuredTime, true
}

// HasInterestAccuredTime returns a boolean if a field has been set.
func (o *GetInterestHistoryResponseRowsInner) HasInterestAccuredTime() bool {
	if o != nil && !common.IsNil(o.InterestAccuredTime) {
		return true
	}

	return false
}

// SetInterestAccuredTime gets a reference to the given int64 and assigns it to the InterestAccuredTime field.
func (o *GetInterestHistoryResponseRowsInner) SetInterestAccuredTime(v int64) {
	o.InterestAccuredTime = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetInterestHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterestHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetInterestHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetInterestHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetRawAsset returns the RawAsset field value if set, zero value otherwise.
func (o *GetInterestHistoryResponseRowsInner) GetRawAsset() string {
	if o == nil || common.IsNil(o.RawAsset) {
		var ret string
		return ret
	}
	return *o.RawAsset
}

// GetRawAssetOk returns a tuple with the RawAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterestHistoryResponseRowsInner) GetRawAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.RawAsset) {
		return nil, false
	}
	return o.RawAsset, true
}

// HasRawAsset returns a boolean if a field has been set.
func (o *GetInterestHistoryResponseRowsInner) HasRawAsset() bool {
	if o != nil && !common.IsNil(o.RawAsset) {
		return true
	}

	return false
}

// SetRawAsset gets a reference to the given string and assigns it to the RawAsset field.
func (o *GetInterestHistoryResponseRowsInner) SetRawAsset(v string) {
	o.RawAsset = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *GetInterestHistoryResponseRowsInner) GetPrincipal() string {
	if o == nil || common.IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterestHistoryResponseRowsInner) GetPrincipalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *GetInterestHistoryResponseRowsInner) HasPrincipal() bool {
	if o != nil && !common.IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *GetInterestHistoryResponseRowsInner) SetPrincipal(v string) {
	o.Principal = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *GetInterestHistoryResponseRowsInner) GetInterest() string {
	if o == nil || common.IsNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterestHistoryResponseRowsInner) GetInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *GetInterestHistoryResponseRowsInner) HasInterest() bool {
	if o != nil && !common.IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *GetInterestHistoryResponseRowsInner) SetInterest(v string) {
	o.Interest = &v
}

// GetInterestRate returns the InterestRate field value if set, zero value otherwise.
func (o *GetInterestHistoryResponseRowsInner) GetInterestRate() string {
	if o == nil || common.IsNil(o.InterestRate) {
		var ret string
		return ret
	}
	return *o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterestHistoryResponseRowsInner) GetInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.InterestRate) {
		return nil, false
	}
	return o.InterestRate, true
}

// HasInterestRate returns a boolean if a field has been set.
func (o *GetInterestHistoryResponseRowsInner) HasInterestRate() bool {
	if o != nil && !common.IsNil(o.InterestRate) {
		return true
	}

	return false
}

// SetInterestRate gets a reference to the given string and assigns it to the InterestRate field.
func (o *GetInterestHistoryResponseRowsInner) SetInterestRate(v string) {
	o.InterestRate = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetInterestHistoryResponseRowsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterestHistoryResponseRowsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetInterestHistoryResponseRowsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetInterestHistoryResponseRowsInner) SetType(v string) {
	o.Type = &v
}

// GetIsolatedSymbol returns the IsolatedSymbol field value if set, zero value otherwise.
func (o *GetInterestHistoryResponseRowsInner) GetIsolatedSymbol() string {
	if o == nil || common.IsNil(o.IsolatedSymbol) {
		var ret string
		return ret
	}
	return *o.IsolatedSymbol
}

// GetIsolatedSymbolOk returns a tuple with the IsolatedSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterestHistoryResponseRowsInner) GetIsolatedSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsolatedSymbol) {
		return nil, false
	}
	return o.IsolatedSymbol, true
}

// HasIsolatedSymbol returns a boolean if a field has been set.
func (o *GetInterestHistoryResponseRowsInner) HasIsolatedSymbol() bool {
	if o != nil && !common.IsNil(o.IsolatedSymbol) {
		return true
	}

	return false
}

// SetIsolatedSymbol gets a reference to the given string and assigns it to the IsolatedSymbol field.
func (o *GetInterestHistoryResponseRowsInner) SetIsolatedSymbol(v string) {
	o.IsolatedSymbol = &v
}

func (o GetInterestHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInterestHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TxId) {
		toSerialize["txId"] = o.TxId
	}
	if !common.IsNil(o.InterestAccuredTime) {
		toSerialize["interestAccuredTime"] = o.InterestAccuredTime
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.RawAsset) {
		toSerialize["rawAsset"] = o.RawAsset
	}
	if !common.IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}
	if !common.IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !common.IsNil(o.InterestRate) {
		toSerialize["interestRate"] = o.InterestRate
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.IsolatedSymbol) {
		toSerialize["isolatedSymbol"] = o.IsolatedSymbol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetInterestHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetInterestHistoryResponseRowsInner := _GetInterestHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetInterestHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetInterestHistoryResponseRowsInner(varGetInterestHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "txId")
		delete(additionalProperties, "interestAccuredTime")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "rawAsset")
		delete(additionalProperties, "principal")
		delete(additionalProperties, "interest")
		delete(additionalProperties, "interestRate")
		delete(additionalProperties, "type")
		delete(additionalProperties, "isolatedSymbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetInterestHistoryResponseRowsInner struct {
	value *GetInterestHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetInterestHistoryResponseRowsInner) Get() *GetInterestHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetInterestHistoryResponseRowsInner) Set(val *GetInterestHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInterestHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInterestHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInterestHistoryResponseRowsInner(val *GetInterestHistoryResponseRowsInner) *NullableGetInterestHistoryResponseRowsInner {
	return &NullableGetInterestHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetInterestHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInterestHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
