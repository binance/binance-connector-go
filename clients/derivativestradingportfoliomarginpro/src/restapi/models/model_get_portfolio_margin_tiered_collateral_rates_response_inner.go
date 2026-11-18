/*
Portfolio Margin Pro REST API

Access advanced account management and high-frequency trading with Binance Portfolio Margin Pro.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetPortfolioMarginTieredCollateralRatesResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPortfolioMarginTieredCollateralRatesResponseInner{}

// GetPortfolioMarginTieredCollateralRatesResponseInner struct for GetPortfolioMarginTieredCollateralRatesResponseInner
type GetPortfolioMarginTieredCollateralRatesResponseInner struct {
	// asset.
	Asset                *string                                                                   `json:"asset,omitempty"`
	CollateralInfo       []GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner `json:"collateralInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPortfolioMarginTieredCollateralRatesResponseInner GetPortfolioMarginTieredCollateralRatesResponseInner

// NewGetPortfolioMarginTieredCollateralRatesResponseInner instantiates a new GetPortfolioMarginTieredCollateralRatesResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPortfolioMarginTieredCollateralRatesResponseInner() *GetPortfolioMarginTieredCollateralRatesResponseInner {
	this := GetPortfolioMarginTieredCollateralRatesResponseInner{}
	return &this
}

// NewGetPortfolioMarginTieredCollateralRatesResponseInnerWithDefaults instantiates a new GetPortfolioMarginTieredCollateralRatesResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPortfolioMarginTieredCollateralRatesResponseInnerWithDefaults() *GetPortfolioMarginTieredCollateralRatesResponseInner {
	this := GetPortfolioMarginTieredCollateralRatesResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetCollateralInfo returns the CollateralInfo field value if set, zero value otherwise.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInner) GetCollateralInfo() []GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner {
	if o == nil || common.IsNil(o.CollateralInfo) {
		var ret []GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner
		return ret
	}
	return o.CollateralInfo
}

// GetCollateralInfoOk returns a tuple with the CollateralInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInner) GetCollateralInfoOk() ([]GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner, bool) {
	if o == nil || common.IsNil(o.CollateralInfo) {
		return nil, false
	}
	return o.CollateralInfo, true
}

// HasCollateralInfo returns a boolean if a field has been set.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInner) HasCollateralInfo() bool {
	if o != nil && !common.IsNil(o.CollateralInfo) {
		return true
	}

	return false
}

// SetCollateralInfo gets a reference to the given []GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner and assigns it to the CollateralInfo field.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInner) SetCollateralInfo(v []GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) {
	o.CollateralInfo = v
}

func (o GetPortfolioMarginTieredCollateralRatesResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPortfolioMarginTieredCollateralRatesResponseInner) ToMap() (map[string]interface{}, error) {
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

func (o *GetPortfolioMarginTieredCollateralRatesResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetPortfolioMarginTieredCollateralRatesResponseInner := _GetPortfolioMarginTieredCollateralRatesResponseInner{}

	err = json.Unmarshal(data, &varGetPortfolioMarginTieredCollateralRatesResponseInner)

	if err != nil {
		return err
	}

	*o = GetPortfolioMarginTieredCollateralRatesResponseInner(varGetPortfolioMarginTieredCollateralRatesResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "collateralInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPortfolioMarginTieredCollateralRatesResponseInner struct {
	value *GetPortfolioMarginTieredCollateralRatesResponseInner
	isSet bool
}

func (v NullableGetPortfolioMarginTieredCollateralRatesResponseInner) Get() *GetPortfolioMarginTieredCollateralRatesResponseInner {
	return v.value
}

func (v *NullableGetPortfolioMarginTieredCollateralRatesResponseInner) Set(val *GetPortfolioMarginTieredCollateralRatesResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPortfolioMarginTieredCollateralRatesResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPortfolioMarginTieredCollateralRatesResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPortfolioMarginTieredCollateralRatesResponseInner(val *GetPortfolioMarginTieredCollateralRatesResponseInner) *NullableGetPortfolioMarginTieredCollateralRatesResponseInner {
	return &NullableGetPortfolioMarginTieredCollateralRatesResponseInner{value: val, isSet: true}
}

func (v NullableGetPortfolioMarginTieredCollateralRatesResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPortfolioMarginTieredCollateralRatesResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
