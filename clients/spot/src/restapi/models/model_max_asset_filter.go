/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MaxAssetFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MaxAssetFilter{}

// MaxAssetFilter struct for MaxAssetFilter
type MaxAssetFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	QtyExponent          *int32  `json:"qtyExponent,omitempty"`
	Limit                *string `json:"limit,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MaxAssetFilter MaxAssetFilter

// NewMaxAssetFilter instantiates a new MaxAssetFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaxAssetFilter() *MaxAssetFilter {
	this := MaxAssetFilter{}
	return &this
}

// NewMaxAssetFilterWithDefaults instantiates a new MaxAssetFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaxAssetFilterWithDefaults() *MaxAssetFilter {
	this := MaxAssetFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *MaxAssetFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxAssetFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *MaxAssetFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *MaxAssetFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetQtyExponent returns the QtyExponent field value if set, zero value otherwise.
func (o *MaxAssetFilter) GetQtyExponent() int32 {
	if o == nil || common.IsNil(o.QtyExponent) {
		var ret int32
		return ret
	}
	return *o.QtyExponent
}

// GetQtyExponentOk returns a tuple with the QtyExponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxAssetFilter) GetQtyExponentOk() (*int32, bool) {
	if o == nil || common.IsNil(o.QtyExponent) {
		return nil, false
	}
	return o.QtyExponent, true
}

// HasQtyExponent returns a boolean if a field has been set.
func (o *MaxAssetFilter) HasQtyExponent() bool {
	if o != nil && !common.IsNil(o.QtyExponent) {
		return true
	}

	return false
}

// SetQtyExponent gets a reference to the given int32 and assigns it to the QtyExponent field.
func (o *MaxAssetFilter) SetQtyExponent(v int32) {
	o.QtyExponent = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *MaxAssetFilter) GetLimit() string {
	if o == nil || common.IsNil(o.Limit) {
		var ret string
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxAssetFilter) GetLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *MaxAssetFilter) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given string and assigns it to the Limit field.
func (o *MaxAssetFilter) SetLimit(v string) {
	o.Limit = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *MaxAssetFilter) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxAssetFilter) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *MaxAssetFilter) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *MaxAssetFilter) SetAsset(v string) {
	o.Asset = &v
}

func (o MaxAssetFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaxAssetFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.QtyExponent) {
		toSerialize["qtyExponent"] = o.QtyExponent
	}
	if !common.IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MaxAssetFilter) UnmarshalJSON(data []byte) (err error) {
	varMaxAssetFilter := _MaxAssetFilter{}

	err = json.Unmarshal(data, &varMaxAssetFilter)

	if err != nil {
		return err
	}

	*o = MaxAssetFilter(varMaxAssetFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "qtyExponent")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "asset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMaxAssetFilter struct {
	value *MaxAssetFilter
	isSet bool
}

func (v NullableMaxAssetFilter) Get() *MaxAssetFilter {
	return v.value
}

func (v *NullableMaxAssetFilter) Set(val *MaxAssetFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMaxAssetFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMaxAssetFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaxAssetFilter(val *MaxAssetFilter) *NullableMaxAssetFilter {
	return &NullableMaxAssetFilter{value: val, isSet: true}
}

func (v NullableMaxAssetFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaxAssetFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
