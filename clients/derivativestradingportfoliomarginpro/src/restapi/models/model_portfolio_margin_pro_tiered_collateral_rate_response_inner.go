/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PortfolioMarginProTieredCollateralRateResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PortfolioMarginProTieredCollateralRateResponseInner{}

// PortfolioMarginProTieredCollateralRateResponseInner struct for PortfolioMarginProTieredCollateralRateResponseInner
type PortfolioMarginProTieredCollateralRateResponseInner struct {
	Asset                *string                                                                  `json:"asset,omitempty"`
	CollateralInfo       []PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner `json:"collateralInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortfolioMarginProTieredCollateralRateResponseInner PortfolioMarginProTieredCollateralRateResponseInner

// NewPortfolioMarginProTieredCollateralRateResponseInner instantiates a new PortfolioMarginProTieredCollateralRateResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioMarginProTieredCollateralRateResponseInner() *PortfolioMarginProTieredCollateralRateResponseInner {
	this := PortfolioMarginProTieredCollateralRateResponseInner{}
	return &this
}

// NewPortfolioMarginProTieredCollateralRateResponseInnerWithDefaults instantiates a new PortfolioMarginProTieredCollateralRateResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioMarginProTieredCollateralRateResponseInnerWithDefaults() *PortfolioMarginProTieredCollateralRateResponseInner {
	this := PortfolioMarginProTieredCollateralRateResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *PortfolioMarginProTieredCollateralRateResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginProTieredCollateralRateResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *PortfolioMarginProTieredCollateralRateResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *PortfolioMarginProTieredCollateralRateResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetCollateralInfo returns the CollateralInfo field value if set, zero value otherwise.
func (o *PortfolioMarginProTieredCollateralRateResponseInner) GetCollateralInfo() []PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner {
	if o == nil || common.IsNil(o.CollateralInfo) {
		var ret []PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner
		return ret
	}
	return o.CollateralInfo
}

// GetCollateralInfoOk returns a tuple with the CollateralInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginProTieredCollateralRateResponseInner) GetCollateralInfoOk() ([]PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner, bool) {
	if o == nil || common.IsNil(o.CollateralInfo) {
		return nil, false
	}
	return o.CollateralInfo, true
}

// HasCollateralInfo returns a boolean if a field has been set.
func (o *PortfolioMarginProTieredCollateralRateResponseInner) HasCollateralInfo() bool {
	if o != nil && !common.IsNil(o.CollateralInfo) {
		return true
	}

	return false
}

// SetCollateralInfo gets a reference to the given []PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner and assigns it to the CollateralInfo field.
func (o *PortfolioMarginProTieredCollateralRateResponseInner) SetCollateralInfo(v []PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) {
	o.CollateralInfo = v
}

func (o PortfolioMarginProTieredCollateralRateResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioMarginProTieredCollateralRateResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.CollateralInfo) {
		toSerialize["collateralInfo"] = o.CollateralInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortfolioMarginProTieredCollateralRateResponseInner) UnmarshalJSON(data []byte) (err error) {
	varPortfolioMarginProTieredCollateralRateResponseInner := _PortfolioMarginProTieredCollateralRateResponseInner{}

	err = json.Unmarshal(data, &varPortfolioMarginProTieredCollateralRateResponseInner)

	if err != nil {
		return err
	}

	*o = PortfolioMarginProTieredCollateralRateResponseInner(varPortfolioMarginProTieredCollateralRateResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "collateralInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortfolioMarginProTieredCollateralRateResponseInner struct {
	value *PortfolioMarginProTieredCollateralRateResponseInner
	isSet bool
}

func (v NullablePortfolioMarginProTieredCollateralRateResponseInner) Get() *PortfolioMarginProTieredCollateralRateResponseInner {
	return v.value
}

func (v *NullablePortfolioMarginProTieredCollateralRateResponseInner) Set(val *PortfolioMarginProTieredCollateralRateResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMarginProTieredCollateralRateResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMarginProTieredCollateralRateResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioMarginProTieredCollateralRateResponseInner(val *PortfolioMarginProTieredCollateralRateResponseInner) *NullablePortfolioMarginProTieredCollateralRateResponseInner {
	return &NullablePortfolioMarginProTieredCollateralRateResponseInner{value: val, isSet: true}
}

func (v NullablePortfolioMarginProTieredCollateralRateResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioMarginProTieredCollateralRateResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
