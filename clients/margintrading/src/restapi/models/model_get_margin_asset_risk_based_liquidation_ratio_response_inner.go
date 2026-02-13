/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetMarginAssetRiskBasedLiquidationRatioResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetMarginAssetRiskBasedLiquidationRatioResponseInner{}

// GetMarginAssetRiskBasedLiquidationRatioResponseInner struct for GetMarginAssetRiskBasedLiquidationRatioResponseInner
type GetMarginAssetRiskBasedLiquidationRatioResponseInner struct {
	Asset                     *string `json:"asset,omitempty"`
	RiskBasedLiquidationRatio *string `json:"riskBasedLiquidationRatio,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _GetMarginAssetRiskBasedLiquidationRatioResponseInner GetMarginAssetRiskBasedLiquidationRatioResponseInner

// NewGetMarginAssetRiskBasedLiquidationRatioResponseInner instantiates a new GetMarginAssetRiskBasedLiquidationRatioResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarginAssetRiskBasedLiquidationRatioResponseInner() *GetMarginAssetRiskBasedLiquidationRatioResponseInner {
	this := GetMarginAssetRiskBasedLiquidationRatioResponseInner{}
	return &this
}

// NewGetMarginAssetRiskBasedLiquidationRatioResponseInnerWithDefaults instantiates a new GetMarginAssetRiskBasedLiquidationRatioResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarginAssetRiskBasedLiquidationRatioResponseInnerWithDefaults() *GetMarginAssetRiskBasedLiquidationRatioResponseInner {
	this := GetMarginAssetRiskBasedLiquidationRatioResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetMarginAssetRiskBasedLiquidationRatioResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginAssetRiskBasedLiquidationRatioResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetMarginAssetRiskBasedLiquidationRatioResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetMarginAssetRiskBasedLiquidationRatioResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetRiskBasedLiquidationRatio returns the RiskBasedLiquidationRatio field value if set, zero value otherwise.
func (o *GetMarginAssetRiskBasedLiquidationRatioResponseInner) GetRiskBasedLiquidationRatio() string {
	if o == nil || common.IsNil(o.RiskBasedLiquidationRatio) {
		var ret string
		return ret
	}
	return *o.RiskBasedLiquidationRatio
}

// GetRiskBasedLiquidationRatioOk returns a tuple with the RiskBasedLiquidationRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginAssetRiskBasedLiquidationRatioResponseInner) GetRiskBasedLiquidationRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskBasedLiquidationRatio) {
		return nil, false
	}
	return o.RiskBasedLiquidationRatio, true
}

// HasRiskBasedLiquidationRatio returns a boolean if a field has been set.
func (o *GetMarginAssetRiskBasedLiquidationRatioResponseInner) HasRiskBasedLiquidationRatio() bool {
	if o != nil && !common.IsNil(o.RiskBasedLiquidationRatio) {
		return true
	}

	return false
}

// SetRiskBasedLiquidationRatio gets a reference to the given string and assigns it to the RiskBasedLiquidationRatio field.
func (o *GetMarginAssetRiskBasedLiquidationRatioResponseInner) SetRiskBasedLiquidationRatio(v string) {
	o.RiskBasedLiquidationRatio = &v
}

func (o GetMarginAssetRiskBasedLiquidationRatioResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarginAssetRiskBasedLiquidationRatioResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.RiskBasedLiquidationRatio) {
		toSerialize["riskBasedLiquidationRatio"] = o.RiskBasedLiquidationRatio
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetMarginAssetRiskBasedLiquidationRatioResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetMarginAssetRiskBasedLiquidationRatioResponseInner := _GetMarginAssetRiskBasedLiquidationRatioResponseInner{}

	err = json.Unmarshal(data, &varGetMarginAssetRiskBasedLiquidationRatioResponseInner)

	if err != nil {
		return err
	}

	*o = GetMarginAssetRiskBasedLiquidationRatioResponseInner(varGetMarginAssetRiskBasedLiquidationRatioResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "riskBasedLiquidationRatio")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMarginAssetRiskBasedLiquidationRatioResponseInner struct {
	value *GetMarginAssetRiskBasedLiquidationRatioResponseInner
	isSet bool
}

func (v NullableGetMarginAssetRiskBasedLiquidationRatioResponseInner) Get() *GetMarginAssetRiskBasedLiquidationRatioResponseInner {
	return v.value
}

func (v *NullableGetMarginAssetRiskBasedLiquidationRatioResponseInner) Set(val *GetMarginAssetRiskBasedLiquidationRatioResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarginAssetRiskBasedLiquidationRatioResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarginAssetRiskBasedLiquidationRatioResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarginAssetRiskBasedLiquidationRatioResponseInner(val *GetMarginAssetRiskBasedLiquidationRatioResponseInner) *NullableGetMarginAssetRiskBasedLiquidationRatioResponseInner {
	return &NullableGetMarginAssetRiskBasedLiquidationRatioResponseInner{value: val, isSet: true}
}

func (v NullableGetMarginAssetRiskBasedLiquidationRatioResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarginAssetRiskBasedLiquidationRatioResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
