/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CrossMarginCollateralRatioResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CrossMarginCollateralRatioResponseInner{}

// CrossMarginCollateralRatioResponseInner struct for CrossMarginCollateralRatioResponseInner
type CrossMarginCollateralRatioResponseInner struct {
	Collaterals          []CrossMarginCollateralRatioResponseInnerCollateralsInner `json:"collaterals,omitempty"`
	AssetNames           []string                                                  `json:"assetNames,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CrossMarginCollateralRatioResponseInner CrossMarginCollateralRatioResponseInner

// NewCrossMarginCollateralRatioResponseInner instantiates a new CrossMarginCollateralRatioResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrossMarginCollateralRatioResponseInner() *CrossMarginCollateralRatioResponseInner {
	this := CrossMarginCollateralRatioResponseInner{}
	return &this
}

// NewCrossMarginCollateralRatioResponseInnerWithDefaults instantiates a new CrossMarginCollateralRatioResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrossMarginCollateralRatioResponseInnerWithDefaults() *CrossMarginCollateralRatioResponseInner {
	this := CrossMarginCollateralRatioResponseInner{}
	return &this
}

// GetCollaterals returns the Collaterals field value if set, zero value otherwise.
func (o *CrossMarginCollateralRatioResponseInner) GetCollaterals() []CrossMarginCollateralRatioResponseInnerCollateralsInner {
	if o == nil || common.IsNil(o.Collaterals) {
		var ret []CrossMarginCollateralRatioResponseInnerCollateralsInner
		return ret
	}
	return o.Collaterals
}

// GetCollateralsOk returns a tuple with the Collaterals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossMarginCollateralRatioResponseInner) GetCollateralsOk() ([]CrossMarginCollateralRatioResponseInnerCollateralsInner, bool) {
	if o == nil || common.IsNil(o.Collaterals) {
		return nil, false
	}
	return o.Collaterals, true
}

// HasCollaterals returns a boolean if a field has been set.
func (o *CrossMarginCollateralRatioResponseInner) HasCollaterals() bool {
	if o != nil && !common.IsNil(o.Collaterals) {
		return true
	}

	return false
}

// SetCollaterals gets a reference to the given []CrossMarginCollateralRatioResponseInnerCollateralsInner and assigns it to the Collaterals field.
func (o *CrossMarginCollateralRatioResponseInner) SetCollaterals(v []CrossMarginCollateralRatioResponseInnerCollateralsInner) {
	o.Collaterals = v
}

// GetAssetNames returns the AssetNames field value if set, zero value otherwise.
func (o *CrossMarginCollateralRatioResponseInner) GetAssetNames() []string {
	if o == nil || common.IsNil(o.AssetNames) {
		var ret []string
		return ret
	}
	return o.AssetNames
}

// GetAssetNamesOk returns a tuple with the AssetNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossMarginCollateralRatioResponseInner) GetAssetNamesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AssetNames) {
		return nil, false
	}
	return o.AssetNames, true
}

// HasAssetNames returns a boolean if a field has been set.
func (o *CrossMarginCollateralRatioResponseInner) HasAssetNames() bool {
	if o != nil && !common.IsNil(o.AssetNames) {
		return true
	}

	return false
}

// SetAssetNames gets a reference to the given []string and assigns it to the AssetNames field.
func (o *CrossMarginCollateralRatioResponseInner) SetAssetNames(v []string) {
	o.AssetNames = v
}

func (o CrossMarginCollateralRatioResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CrossMarginCollateralRatioResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Collaterals) {
		toSerialize["collaterals"] = o.Collaterals
	}
	if !common.IsNil(o.AssetNames) {
		toSerialize["assetNames"] = o.AssetNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CrossMarginCollateralRatioResponseInner) UnmarshalJSON(data []byte) (err error) {
	varCrossMarginCollateralRatioResponseInner := _CrossMarginCollateralRatioResponseInner{}

	err = json.Unmarshal(data, &varCrossMarginCollateralRatioResponseInner)

	if err != nil {
		return err
	}

	*o = CrossMarginCollateralRatioResponseInner(varCrossMarginCollateralRatioResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "collaterals")
		delete(additionalProperties, "assetNames")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCrossMarginCollateralRatioResponseInner struct {
	value *CrossMarginCollateralRatioResponseInner
	isSet bool
}

func (v NullableCrossMarginCollateralRatioResponseInner) Get() *CrossMarginCollateralRatioResponseInner {
	return v.value
}

func (v *NullableCrossMarginCollateralRatioResponseInner) Set(val *CrossMarginCollateralRatioResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCrossMarginCollateralRatioResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCrossMarginCollateralRatioResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrossMarginCollateralRatioResponseInner(val *CrossMarginCollateralRatioResponseInner) *NullableCrossMarginCollateralRatioResponseInner {
	return &NullableCrossMarginCollateralRatioResponseInner{value: val, isSet: true}
}

func (v NullableCrossMarginCollateralRatioResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrossMarginCollateralRatioResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
