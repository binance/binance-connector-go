/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner{}

// QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner struct for QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner
type QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner struct {
	AssetNames           []string                                                                          `json:"assetNames,omitempty"`
	Rank                 *int64                                                                            `json:"rank,omitempty"`
	Brackets             []QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner `json:"brackets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner

// NewQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner instantiates a new QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner() *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner {
	this := QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner{}
	return &this
}

// NewQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerWithDefaults instantiates a new QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerWithDefaults() *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner {
	this := QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner{}
	return &this
}

// GetAssetNames returns the AssetNames field value if set, zero value otherwise.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) GetAssetNames() []string {
	if o == nil || common.IsNil(o.AssetNames) {
		var ret []string
		return ret
	}
	return o.AssetNames
}

// GetAssetNamesOk returns a tuple with the AssetNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) GetAssetNamesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AssetNames) {
		return nil, false
	}
	return o.AssetNames, true
}

// HasAssetNames returns a boolean if a field has been set.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) HasAssetNames() bool {
	if o != nil && !common.IsNil(o.AssetNames) {
		return true
	}

	return false
}

// SetAssetNames gets a reference to the given []string and assigns it to the AssetNames field.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) SetAssetNames(v []string) {
	o.AssetNames = v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) GetRank() int64 {
	if o == nil || common.IsNil(o.Rank) {
		var ret int64
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) GetRankOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) HasRank() bool {
	if o != nil && !common.IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given int64 and assigns it to the Rank field.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) SetRank(v int64) {
	o.Rank = &v
}

// GetBrackets returns the Brackets field value if set, zero value otherwise.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) GetBrackets() []QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner {
	if o == nil || common.IsNil(o.Brackets) {
		var ret []QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner
		return ret
	}
	return o.Brackets
}

// GetBracketsOk returns a tuple with the Brackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) GetBracketsOk() ([]QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner, bool) {
	if o == nil || common.IsNil(o.Brackets) {
		return nil, false
	}
	return o.Brackets, true
}

// HasBrackets returns a boolean if a field has been set.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) HasBrackets() bool {
	if o != nil && !common.IsNil(o.Brackets) {
		return true
	}

	return false
}

// SetBrackets gets a reference to the given []QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner and assigns it to the Brackets field.
func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) SetBrackets(v []QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInnerBracketsInner) {
	o.Brackets = v
}

func (o QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AssetNames) {
		toSerialize["assetNames"] = o.AssetNames
	}
	if !common.IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	if !common.IsNil(o.Brackets) {
		toSerialize["brackets"] = o.Brackets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner := _QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner{}

	err = json.Unmarshal(data, &varQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner)

	if err != nil {
		return err
	}

	*o = QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner(varQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assetNames")
		delete(additionalProperties, "rank")
		delete(additionalProperties, "brackets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner struct {
	value *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner
	isSet bool
}

func (v NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) Get() *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner {
	return v.value
}

func (v *NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) Set(val *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner(val *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) *NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner {
	return &NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner{value: val, isSet: true}
}

func (v NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
