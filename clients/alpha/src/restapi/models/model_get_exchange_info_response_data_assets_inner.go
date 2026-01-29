/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetExchangeInfoResponseDataAssetsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetExchangeInfoResponseDataAssetsInner{}

// GetExchangeInfoResponseDataAssetsInner struct for GetExchangeInfoResponseDataAssetsInner
type GetExchangeInfoResponseDataAssetsInner struct {
	Asset                *string `json:"asset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetExchangeInfoResponseDataAssetsInner GetExchangeInfoResponseDataAssetsInner

// NewGetExchangeInfoResponseDataAssetsInner instantiates a new GetExchangeInfoResponseDataAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExchangeInfoResponseDataAssetsInner() *GetExchangeInfoResponseDataAssetsInner {
	this := GetExchangeInfoResponseDataAssetsInner{}
	return &this
}

// NewGetExchangeInfoResponseDataAssetsInnerWithDefaults instantiates a new GetExchangeInfoResponseDataAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExchangeInfoResponseDataAssetsInnerWithDefaults() *GetExchangeInfoResponseDataAssetsInner {
	this := GetExchangeInfoResponseDataAssetsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataAssetsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataAssetsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataAssetsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetExchangeInfoResponseDataAssetsInner) SetAsset(v string) {
	o.Asset = &v
}

func (o GetExchangeInfoResponseDataAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExchangeInfoResponseDataAssetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetExchangeInfoResponseDataAssetsInner) UnmarshalJSON(data []byte) (err error) {
	varGetExchangeInfoResponseDataAssetsInner := _GetExchangeInfoResponseDataAssetsInner{}

	err = json.Unmarshal(data, &varGetExchangeInfoResponseDataAssetsInner)

	if err != nil {
		return err
	}

	*o = GetExchangeInfoResponseDataAssetsInner(varGetExchangeInfoResponseDataAssetsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetExchangeInfoResponseDataAssetsInner struct {
	value *GetExchangeInfoResponseDataAssetsInner
	isSet bool
}

func (v NullableGetExchangeInfoResponseDataAssetsInner) Get() *GetExchangeInfoResponseDataAssetsInner {
	return v.value
}

func (v *NullableGetExchangeInfoResponseDataAssetsInner) Set(val *GetExchangeInfoResponseDataAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExchangeInfoResponseDataAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExchangeInfoResponseDataAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExchangeInfoResponseDataAssetsInner(val *GetExchangeInfoResponseDataAssetsInner) *NullableGetExchangeInfoResponseDataAssetsInner {
	return &NullableGetExchangeInfoResponseDataAssetsInner{value: val, isSet: true}
}

func (v NullableGetExchangeInfoResponseDataAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExchangeInfoResponseDataAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
