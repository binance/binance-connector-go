/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner{}

// GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner struct for GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner
type GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner struct {
	Asset                *string `json:"asset,omitempty"`
	UniMaintainUsd       *string `json:"uniMaintainUsd,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner

// NewGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner instantiates a new GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner() *GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner {
	this := GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner{}
	return &this
}

// NewGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInnerWithDefaults instantiates a new GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInnerWithDefaults() *GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner {
	this := GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) SetAsset(v string) {
	o.Asset = &v
}

// GetUniMaintainUsd returns the UniMaintainUsd field value if set, zero value otherwise.
func (o *GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) GetUniMaintainUsd() string {
	if o == nil || common.IsNil(o.UniMaintainUsd) {
		var ret string
		return ret
	}
	return *o.UniMaintainUsd
}

// GetUniMaintainUsdOk returns a tuple with the UniMaintainUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) GetUniMaintainUsdOk() (*string, bool) {
	if o == nil || common.IsNil(o.UniMaintainUsd) {
		return nil, false
	}
	return o.UniMaintainUsd, true
}

// HasUniMaintainUsd returns a boolean if a field has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) HasUniMaintainUsd() bool {
	if o != nil && !common.IsNil(o.UniMaintainUsd) {
		return true
	}

	return false
}

// SetUniMaintainUsd gets a reference to the given string and assigns it to the UniMaintainUsd field.
func (o *GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) SetUniMaintainUsd(v string) {
	o.UniMaintainUsd = &v
}

func (o GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.UniMaintainUsd) {
		toSerialize["uniMaintainUsd"] = o.UniMaintainUsd
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) UnmarshalJSON(data []byte) (err error) {
	varGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner := _GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner{}

	err = json.Unmarshal(data, &varGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner)

	if err != nil {
		return err
	}

	*o = GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner(varGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "uniMaintainUsd")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner struct {
	value *GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner
	isSet bool
}

func (v NullableGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) Get() *GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner {
	return v.value
}

func (v *NullableGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) Set(val *GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner(val *GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) *NullableGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner {
	return &NullableGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner{value: val, isSet: true}
}

func (v NullableGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
