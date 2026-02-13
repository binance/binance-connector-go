/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PortfolioMarginCollateralRateResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PortfolioMarginCollateralRateResponseInner{}

// PortfolioMarginCollateralRateResponseInner struct for PortfolioMarginCollateralRateResponseInner
type PortfolioMarginCollateralRateResponseInner struct {
	Asset                *string `json:"asset,omitempty"`
	CollateralRate       *string `json:"collateralRate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortfolioMarginCollateralRateResponseInner PortfolioMarginCollateralRateResponseInner

// NewPortfolioMarginCollateralRateResponseInner instantiates a new PortfolioMarginCollateralRateResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioMarginCollateralRateResponseInner() *PortfolioMarginCollateralRateResponseInner {
	this := PortfolioMarginCollateralRateResponseInner{}
	return &this
}

// NewPortfolioMarginCollateralRateResponseInnerWithDefaults instantiates a new PortfolioMarginCollateralRateResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioMarginCollateralRateResponseInnerWithDefaults() *PortfolioMarginCollateralRateResponseInner {
	this := PortfolioMarginCollateralRateResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *PortfolioMarginCollateralRateResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginCollateralRateResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *PortfolioMarginCollateralRateResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *PortfolioMarginCollateralRateResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetCollateralRate returns the CollateralRate field value if set, zero value otherwise.
func (o *PortfolioMarginCollateralRateResponseInner) GetCollateralRate() string {
	if o == nil || common.IsNil(o.CollateralRate) {
		var ret string
		return ret
	}
	return *o.CollateralRate
}

// GetCollateralRateOk returns a tuple with the CollateralRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginCollateralRateResponseInner) GetCollateralRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralRate) {
		return nil, false
	}
	return o.CollateralRate, true
}

// HasCollateralRate returns a boolean if a field has been set.
func (o *PortfolioMarginCollateralRateResponseInner) HasCollateralRate() bool {
	if o != nil && !common.IsNil(o.CollateralRate) {
		return true
	}

	return false
}

// SetCollateralRate gets a reference to the given string and assigns it to the CollateralRate field.
func (o *PortfolioMarginCollateralRateResponseInner) SetCollateralRate(v string) {
	o.CollateralRate = &v
}

func (o PortfolioMarginCollateralRateResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioMarginCollateralRateResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.CollateralRate) {
		toSerialize["collateralRate"] = o.CollateralRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortfolioMarginCollateralRateResponseInner) UnmarshalJSON(data []byte) (err error) {
	varPortfolioMarginCollateralRateResponseInner := _PortfolioMarginCollateralRateResponseInner{}

	err = json.Unmarshal(data, &varPortfolioMarginCollateralRateResponseInner)

	if err != nil {
		return err
	}

	*o = PortfolioMarginCollateralRateResponseInner(varPortfolioMarginCollateralRateResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "collateralRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortfolioMarginCollateralRateResponseInner struct {
	value *PortfolioMarginCollateralRateResponseInner
	isSet bool
}

func (v NullablePortfolioMarginCollateralRateResponseInner) Get() *PortfolioMarginCollateralRateResponseInner {
	return v.value
}

func (v *NullablePortfolioMarginCollateralRateResponseInner) Set(val *PortfolioMarginCollateralRateResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMarginCollateralRateResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMarginCollateralRateResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioMarginCollateralRateResponseInner(val *PortfolioMarginCollateralRateResponseInner) *NullablePortfolioMarginCollateralRateResponseInner {
	return &NullablePortfolioMarginCollateralRateResponseInner{value: val, isSet: true}
}

func (v NullablePortfolioMarginCollateralRateResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioMarginCollateralRateResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
