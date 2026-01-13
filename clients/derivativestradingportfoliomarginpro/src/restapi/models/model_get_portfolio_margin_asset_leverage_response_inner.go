/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetPortfolioMarginAssetLeverageResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPortfolioMarginAssetLeverageResponseInner{}

// GetPortfolioMarginAssetLeverageResponseInner struct for GetPortfolioMarginAssetLeverageResponseInner
type GetPortfolioMarginAssetLeverageResponseInner struct {
	Asset                *string `json:"asset,omitempty"`
	Leverage             *int64  `json:"leverage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPortfolioMarginAssetLeverageResponseInner GetPortfolioMarginAssetLeverageResponseInner

// NewGetPortfolioMarginAssetLeverageResponseInner instantiates a new GetPortfolioMarginAssetLeverageResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPortfolioMarginAssetLeverageResponseInner() *GetPortfolioMarginAssetLeverageResponseInner {
	this := GetPortfolioMarginAssetLeverageResponseInner{}
	return &this
}

// NewGetPortfolioMarginAssetLeverageResponseInnerWithDefaults instantiates a new GetPortfolioMarginAssetLeverageResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPortfolioMarginAssetLeverageResponseInnerWithDefaults() *GetPortfolioMarginAssetLeverageResponseInner {
	this := GetPortfolioMarginAssetLeverageResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetPortfolioMarginAssetLeverageResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginAssetLeverageResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetPortfolioMarginAssetLeverageResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetPortfolioMarginAssetLeverageResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *GetPortfolioMarginAssetLeverageResponseInner) GetLeverage() int64 {
	if o == nil || common.IsNil(o.Leverage) {
		var ret int64
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginAssetLeverageResponseInner) GetLeverageOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *GetPortfolioMarginAssetLeverageResponseInner) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given int64 and assigns it to the Leverage field.
func (o *GetPortfolioMarginAssetLeverageResponseInner) SetLeverage(v int64) {
	o.Leverage = &v
}

func (o GetPortfolioMarginAssetLeverageResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPortfolioMarginAssetLeverageResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPortfolioMarginAssetLeverageResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetPortfolioMarginAssetLeverageResponseInner := _GetPortfolioMarginAssetLeverageResponseInner{}

	err = json.Unmarshal(data, &varGetPortfolioMarginAssetLeverageResponseInner)

	if err != nil {
		return err
	}

	*o = GetPortfolioMarginAssetLeverageResponseInner(varGetPortfolioMarginAssetLeverageResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "leverage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPortfolioMarginAssetLeverageResponseInner struct {
	value *GetPortfolioMarginAssetLeverageResponseInner
	isSet bool
}

func (v NullableGetPortfolioMarginAssetLeverageResponseInner) Get() *GetPortfolioMarginAssetLeverageResponseInner {
	return v.value
}

func (v *NullableGetPortfolioMarginAssetLeverageResponseInner) Set(val *GetPortfolioMarginAssetLeverageResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPortfolioMarginAssetLeverageResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPortfolioMarginAssetLeverageResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPortfolioMarginAssetLeverageResponseInner(val *GetPortfolioMarginAssetLeverageResponseInner) *NullableGetPortfolioMarginAssetLeverageResponseInner {
	return &NullableGetPortfolioMarginAssetLeverageResponseInner{value: val, isSet: true}
}

func (v NullableGetPortfolioMarginAssetLeverageResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPortfolioMarginAssetLeverageResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
