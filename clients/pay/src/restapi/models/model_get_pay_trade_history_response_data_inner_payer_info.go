/*
Binance Pay REST API

OpenAPI Specification for the Binance Pay REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetPayTradeHistoryResponseDataInnerPayerInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPayTradeHistoryResponseDataInnerPayerInfo{}

// GetPayTradeHistoryResponseDataInnerPayerInfo struct for GetPayTradeHistoryResponseDataInnerPayerInfo
type GetPayTradeHistoryResponseDataInnerPayerInfo struct {
	Name                 *string `json:"name,omitempty"`
	Type                 *string `json:"type,omitempty"`
	BinanceId            *string `json:"binanceId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPayTradeHistoryResponseDataInnerPayerInfo GetPayTradeHistoryResponseDataInnerPayerInfo

// NewGetPayTradeHistoryResponseDataInnerPayerInfo instantiates a new GetPayTradeHistoryResponseDataInnerPayerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPayTradeHistoryResponseDataInnerPayerInfo() *GetPayTradeHistoryResponseDataInnerPayerInfo {
	this := GetPayTradeHistoryResponseDataInnerPayerInfo{}
	return &this
}

// NewGetPayTradeHistoryResponseDataInnerPayerInfoWithDefaults instantiates a new GetPayTradeHistoryResponseDataInnerPayerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPayTradeHistoryResponseDataInnerPayerInfoWithDefaults() *GetPayTradeHistoryResponseDataInnerPayerInfo {
	this := GetPayTradeHistoryResponseDataInnerPayerInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerPayerInfo) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerPayerInfo) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerPayerInfo) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetPayTradeHistoryResponseDataInnerPayerInfo) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerPayerInfo) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerPayerInfo) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerPayerInfo) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetPayTradeHistoryResponseDataInnerPayerInfo) SetType(v string) {
	o.Type = &v
}

// GetBinanceId returns the BinanceId field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerPayerInfo) GetBinanceId() string {
	if o == nil || common.IsNil(o.BinanceId) {
		var ret string
		return ret
	}
	return *o.BinanceId
}

// GetBinanceIdOk returns a tuple with the BinanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerPayerInfo) GetBinanceIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BinanceId) {
		return nil, false
	}
	return o.BinanceId, true
}

// HasBinanceId returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerPayerInfo) HasBinanceId() bool {
	if o != nil && !common.IsNil(o.BinanceId) {
		return true
	}

	return false
}

// SetBinanceId gets a reference to the given string and assigns it to the BinanceId field.
func (o *GetPayTradeHistoryResponseDataInnerPayerInfo) SetBinanceId(v string) {
	o.BinanceId = &v
}

func (o GetPayTradeHistoryResponseDataInnerPayerInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPayTradeHistoryResponseDataInnerPayerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.BinanceId) {
		toSerialize["binanceId"] = o.BinanceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPayTradeHistoryResponseDataInnerPayerInfo) UnmarshalJSON(data []byte) (err error) {
	varGetPayTradeHistoryResponseDataInnerPayerInfo := _GetPayTradeHistoryResponseDataInnerPayerInfo{}

	err = json.Unmarshal(data, &varGetPayTradeHistoryResponseDataInnerPayerInfo)

	if err != nil {
		return err
	}

	*o = GetPayTradeHistoryResponseDataInnerPayerInfo(varGetPayTradeHistoryResponseDataInnerPayerInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "binanceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPayTradeHistoryResponseDataInnerPayerInfo struct {
	value *GetPayTradeHistoryResponseDataInnerPayerInfo
	isSet bool
}

func (v NullableGetPayTradeHistoryResponseDataInnerPayerInfo) Get() *GetPayTradeHistoryResponseDataInnerPayerInfo {
	return v.value
}

func (v *NullableGetPayTradeHistoryResponseDataInnerPayerInfo) Set(val *GetPayTradeHistoryResponseDataInnerPayerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPayTradeHistoryResponseDataInnerPayerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPayTradeHistoryResponseDataInnerPayerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPayTradeHistoryResponseDataInnerPayerInfo(val *GetPayTradeHistoryResponseDataInnerPayerInfo) *NullableGetPayTradeHistoryResponseDataInnerPayerInfo {
	return &NullableGetPayTradeHistoryResponseDataInnerPayerInfo{value: val, isSet: true}
}

func (v NullableGetPayTradeHistoryResponseDataInnerPayerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPayTradeHistoryResponseDataInnerPayerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
