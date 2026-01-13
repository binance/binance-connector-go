/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MinNotionalFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MinNotionalFilter{}

// MinNotionalFilter struct for MinNotionalFilter
type MinNotionalFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	PriceExponent        *int32  `json:"priceExponent,omitempty"`
	MinNotional          *string `json:"minNotional,omitempty"`
	ApplyToMarket        *bool   `json:"applyToMarket,omitempty"`
	AvgPriceMins         *int32  `json:"avgPriceMins,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MinNotionalFilter MinNotionalFilter

// NewMinNotionalFilter instantiates a new MinNotionalFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinNotionalFilter() *MinNotionalFilter {
	this := MinNotionalFilter{}
	return &this
}

// NewMinNotionalFilterWithDefaults instantiates a new MinNotionalFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinNotionalFilterWithDefaults() *MinNotionalFilter {
	this := MinNotionalFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *MinNotionalFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinNotionalFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *MinNotionalFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *MinNotionalFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetPriceExponent returns the PriceExponent field value if set, zero value otherwise.
func (o *MinNotionalFilter) GetPriceExponent() int32 {
	if o == nil || common.IsNil(o.PriceExponent) {
		var ret int32
		return ret
	}
	return *o.PriceExponent
}

// GetPriceExponentOk returns a tuple with the PriceExponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinNotionalFilter) GetPriceExponentOk() (*int32, bool) {
	if o == nil || common.IsNil(o.PriceExponent) {
		return nil, false
	}
	return o.PriceExponent, true
}

// HasPriceExponent returns a boolean if a field has been set.
func (o *MinNotionalFilter) HasPriceExponent() bool {
	if o != nil && !common.IsNil(o.PriceExponent) {
		return true
	}

	return false
}

// SetPriceExponent gets a reference to the given int32 and assigns it to the PriceExponent field.
func (o *MinNotionalFilter) SetPriceExponent(v int32) {
	o.PriceExponent = &v
}

// GetMinNotional returns the MinNotional field value if set, zero value otherwise.
func (o *MinNotionalFilter) GetMinNotional() string {
	if o == nil || common.IsNil(o.MinNotional) {
		var ret string
		return ret
	}
	return *o.MinNotional
}

// GetMinNotionalOk returns a tuple with the MinNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinNotionalFilter) GetMinNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinNotional) {
		return nil, false
	}
	return o.MinNotional, true
}

// HasMinNotional returns a boolean if a field has been set.
func (o *MinNotionalFilter) HasMinNotional() bool {
	if o != nil && !common.IsNil(o.MinNotional) {
		return true
	}

	return false
}

// SetMinNotional gets a reference to the given string and assigns it to the MinNotional field.
func (o *MinNotionalFilter) SetMinNotional(v string) {
	o.MinNotional = &v
}

// GetApplyToMarket returns the ApplyToMarket field value if set, zero value otherwise.
func (o *MinNotionalFilter) GetApplyToMarket() bool {
	if o == nil || common.IsNil(o.ApplyToMarket) {
		var ret bool
		return ret
	}
	return *o.ApplyToMarket
}

// GetApplyToMarketOk returns a tuple with the ApplyToMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinNotionalFilter) GetApplyToMarketOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ApplyToMarket) {
		return nil, false
	}
	return o.ApplyToMarket, true
}

// HasApplyToMarket returns a boolean if a field has been set.
func (o *MinNotionalFilter) HasApplyToMarket() bool {
	if o != nil && !common.IsNil(o.ApplyToMarket) {
		return true
	}

	return false
}

// SetApplyToMarket gets a reference to the given bool and assigns it to the ApplyToMarket field.
func (o *MinNotionalFilter) SetApplyToMarket(v bool) {
	o.ApplyToMarket = &v
}

// GetAvgPriceMins returns the AvgPriceMins field value if set, zero value otherwise.
func (o *MinNotionalFilter) GetAvgPriceMins() int32 {
	if o == nil || common.IsNil(o.AvgPriceMins) {
		var ret int32
		return ret
	}
	return *o.AvgPriceMins
}

// GetAvgPriceMinsOk returns a tuple with the AvgPriceMins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinNotionalFilter) GetAvgPriceMinsOk() (*int32, bool) {
	if o == nil || common.IsNil(o.AvgPriceMins) {
		return nil, false
	}
	return o.AvgPriceMins, true
}

// HasAvgPriceMins returns a boolean if a field has been set.
func (o *MinNotionalFilter) HasAvgPriceMins() bool {
	if o != nil && !common.IsNil(o.AvgPriceMins) {
		return true
	}

	return false
}

// SetAvgPriceMins gets a reference to the given int32 and assigns it to the AvgPriceMins field.
func (o *MinNotionalFilter) SetAvgPriceMins(v int32) {
	o.AvgPriceMins = &v
}

func (o MinNotionalFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MinNotionalFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.PriceExponent) {
		toSerialize["priceExponent"] = o.PriceExponent
	}
	if !common.IsNil(o.MinNotional) {
		toSerialize["minNotional"] = o.MinNotional
	}
	if !common.IsNil(o.ApplyToMarket) {
		toSerialize["applyToMarket"] = o.ApplyToMarket
	}
	if !common.IsNil(o.AvgPriceMins) {
		toSerialize["avgPriceMins"] = o.AvgPriceMins
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MinNotionalFilter) UnmarshalJSON(data []byte) (err error) {
	varMinNotionalFilter := _MinNotionalFilter{}

	err = json.Unmarshal(data, &varMinNotionalFilter)

	if err != nil {
		return err
	}

	*o = MinNotionalFilter(varMinNotionalFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "priceExponent")
		delete(additionalProperties, "minNotional")
		delete(additionalProperties, "applyToMarket")
		delete(additionalProperties, "avgPriceMins")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMinNotionalFilter struct {
	value *MinNotionalFilter
	isSet bool
}

func (v NullableMinNotionalFilter) Get() *MinNotionalFilter {
	return v.value
}

func (v *NullableMinNotionalFilter) Set(val *MinNotionalFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMinNotionalFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMinNotionalFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinNotionalFilter(val *MinNotionalFilter) *NullableMinNotionalFilter {
	return &NullableMinNotionalFilter{value: val, isSet: true}
}

func (v NullableMinNotionalFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinNotionalFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
