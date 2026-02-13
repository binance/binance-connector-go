/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryPortfolioMarginAssetIndexPriceResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPortfolioMarginAssetIndexPriceResponseInner{}

// QueryPortfolioMarginAssetIndexPriceResponseInner struct for QueryPortfolioMarginAssetIndexPriceResponseInner
type QueryPortfolioMarginAssetIndexPriceResponseInner struct {
	Asset                *string `json:"asset,omitempty"`
	AssetIndexPrice      *string `json:"assetIndexPrice,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryPortfolioMarginAssetIndexPriceResponseInner QueryPortfolioMarginAssetIndexPriceResponseInner

// NewQueryPortfolioMarginAssetIndexPriceResponseInner instantiates a new QueryPortfolioMarginAssetIndexPriceResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPortfolioMarginAssetIndexPriceResponseInner() *QueryPortfolioMarginAssetIndexPriceResponseInner {
	this := QueryPortfolioMarginAssetIndexPriceResponseInner{}
	return &this
}

// NewQueryPortfolioMarginAssetIndexPriceResponseInnerWithDefaults instantiates a new QueryPortfolioMarginAssetIndexPriceResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPortfolioMarginAssetIndexPriceResponseInnerWithDefaults() *QueryPortfolioMarginAssetIndexPriceResponseInner {
	this := QueryPortfolioMarginAssetIndexPriceResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryPortfolioMarginAssetIndexPriceResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginAssetIndexPriceResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryPortfolioMarginAssetIndexPriceResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryPortfolioMarginAssetIndexPriceResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAssetIndexPrice returns the AssetIndexPrice field value if set, zero value otherwise.
func (o *QueryPortfolioMarginAssetIndexPriceResponseInner) GetAssetIndexPrice() string {
	if o == nil || common.IsNil(o.AssetIndexPrice) {
		var ret string
		return ret
	}
	return *o.AssetIndexPrice
}

// GetAssetIndexPriceOk returns a tuple with the AssetIndexPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginAssetIndexPriceResponseInner) GetAssetIndexPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AssetIndexPrice) {
		return nil, false
	}
	return o.AssetIndexPrice, true
}

// HasAssetIndexPrice returns a boolean if a field has been set.
func (o *QueryPortfolioMarginAssetIndexPriceResponseInner) HasAssetIndexPrice() bool {
	if o != nil && !common.IsNil(o.AssetIndexPrice) {
		return true
	}

	return false
}

// SetAssetIndexPrice gets a reference to the given string and assigns it to the AssetIndexPrice field.
func (o *QueryPortfolioMarginAssetIndexPriceResponseInner) SetAssetIndexPrice(v string) {
	o.AssetIndexPrice = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *QueryPortfolioMarginAssetIndexPriceResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginAssetIndexPriceResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *QueryPortfolioMarginAssetIndexPriceResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *QueryPortfolioMarginAssetIndexPriceResponseInner) SetTime(v int64) {
	o.Time = &v
}

func (o QueryPortfolioMarginAssetIndexPriceResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPortfolioMarginAssetIndexPriceResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.AssetIndexPrice) {
		toSerialize["assetIndexPrice"] = o.AssetIndexPrice
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryPortfolioMarginAssetIndexPriceResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryPortfolioMarginAssetIndexPriceResponseInner := _QueryPortfolioMarginAssetIndexPriceResponseInner{}

	err = json.Unmarshal(data, &varQueryPortfolioMarginAssetIndexPriceResponseInner)

	if err != nil {
		return err
	}

	*o = QueryPortfolioMarginAssetIndexPriceResponseInner(varQueryPortfolioMarginAssetIndexPriceResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "assetIndexPrice")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryPortfolioMarginAssetIndexPriceResponseInner struct {
	value *QueryPortfolioMarginAssetIndexPriceResponseInner
	isSet bool
}

func (v NullableQueryPortfolioMarginAssetIndexPriceResponseInner) Get() *QueryPortfolioMarginAssetIndexPriceResponseInner {
	return v.value
}

func (v *NullableQueryPortfolioMarginAssetIndexPriceResponseInner) Set(val *QueryPortfolioMarginAssetIndexPriceResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPortfolioMarginAssetIndexPriceResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPortfolioMarginAssetIndexPriceResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPortfolioMarginAssetIndexPriceResponseInner(val *QueryPortfolioMarginAssetIndexPriceResponseInner) *NullableQueryPortfolioMarginAssetIndexPriceResponseInner {
	return &NullableQueryPortfolioMarginAssetIndexPriceResponseInner{value: val, isSet: true}
}

func (v NullableQueryPortfolioMarginAssetIndexPriceResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPortfolioMarginAssetIndexPriceResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
