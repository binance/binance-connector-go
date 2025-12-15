/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ExchangeInformationResponseAssetsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeInformationResponseAssetsInner{}

// ExchangeInformationResponseAssetsInner struct for ExchangeInformationResponseAssetsInner
type ExchangeInformationResponseAssetsInner struct {
	Asset                *string `json:"asset,omitempty"`
	MarginAvailable      *bool   `json:"marginAvailable,omitempty"`
	AutoAssetExchange    *string `json:"autoAssetExchange,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangeInformationResponseAssetsInner ExchangeInformationResponseAssetsInner

// NewExchangeInformationResponseAssetsInner instantiates a new ExchangeInformationResponseAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeInformationResponseAssetsInner() *ExchangeInformationResponseAssetsInner {
	this := ExchangeInformationResponseAssetsInner{}
	return &this
}

// NewExchangeInformationResponseAssetsInnerWithDefaults instantiates a new ExchangeInformationResponseAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeInformationResponseAssetsInnerWithDefaults() *ExchangeInformationResponseAssetsInner {
	this := ExchangeInformationResponseAssetsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *ExchangeInformationResponseAssetsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseAssetsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *ExchangeInformationResponseAssetsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *ExchangeInformationResponseAssetsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetMarginAvailable returns the MarginAvailable field value if set, zero value otherwise.
func (o *ExchangeInformationResponseAssetsInner) GetMarginAvailable() bool {
	if o == nil || common.IsNil(o.MarginAvailable) {
		var ret bool
		return ret
	}
	return *o.MarginAvailable
}

// GetMarginAvailableOk returns a tuple with the MarginAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseAssetsInner) GetMarginAvailableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MarginAvailable) {
		return nil, false
	}
	return o.MarginAvailable, true
}

// HasMarginAvailable returns a boolean if a field has been set.
func (o *ExchangeInformationResponseAssetsInner) HasMarginAvailable() bool {
	if o != nil && !common.IsNil(o.MarginAvailable) {
		return true
	}

	return false
}

// SetMarginAvailable gets a reference to the given bool and assigns it to the MarginAvailable field.
func (o *ExchangeInformationResponseAssetsInner) SetMarginAvailable(v bool) {
	o.MarginAvailable = &v
}

// GetAutoAssetExchange returns the AutoAssetExchange field value if set, zero value otherwise.
func (o *ExchangeInformationResponseAssetsInner) GetAutoAssetExchange() string {
	if o == nil || common.IsNil(o.AutoAssetExchange) {
		var ret string
		return ret
	}
	return *o.AutoAssetExchange
}

// GetAutoAssetExchangeOk returns a tuple with the AutoAssetExchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseAssetsInner) GetAutoAssetExchangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AutoAssetExchange) {
		return nil, false
	}
	return o.AutoAssetExchange, true
}

// HasAutoAssetExchange returns a boolean if a field has been set.
func (o *ExchangeInformationResponseAssetsInner) HasAutoAssetExchange() bool {
	if o != nil && !common.IsNil(o.AutoAssetExchange) {
		return true
	}

	return false
}

// SetAutoAssetExchange gets a reference to the given string and assigns it to the AutoAssetExchange field.
func (o *ExchangeInformationResponseAssetsInner) SetAutoAssetExchange(v string) {
	o.AutoAssetExchange = &v
}

func (o ExchangeInformationResponseAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeInformationResponseAssetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.MarginAvailable) {
		toSerialize["marginAvailable"] = o.MarginAvailable
	}
	if !common.IsNil(o.AutoAssetExchange) {
		toSerialize["autoAssetExchange"] = o.AutoAssetExchange
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExchangeInformationResponseAssetsInner) UnmarshalJSON(data []byte) (err error) {
	varExchangeInformationResponseAssetsInner := _ExchangeInformationResponseAssetsInner{}

	err = json.Unmarshal(data, &varExchangeInformationResponseAssetsInner)

	if err != nil {
		return err
	}

	*o = ExchangeInformationResponseAssetsInner(varExchangeInformationResponseAssetsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "marginAvailable")
		delete(additionalProperties, "autoAssetExchange")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeInformationResponseAssetsInner struct {
	value *ExchangeInformationResponseAssetsInner
	isSet bool
}

func (v NullableExchangeInformationResponseAssetsInner) Get() *ExchangeInformationResponseAssetsInner {
	return v.value
}

func (v *NullableExchangeInformationResponseAssetsInner) Set(val *ExchangeInformationResponseAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInformationResponseAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInformationResponseAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInformationResponseAssetsInner(val *ExchangeInformationResponseAssetsInner) *NullableExchangeInformationResponseAssetsInner {
	return &NullableExchangeInformationResponseAssetsInner{value: val, isSet: true}
}

func (v NullableExchangeInformationResponseAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInformationResponseAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
