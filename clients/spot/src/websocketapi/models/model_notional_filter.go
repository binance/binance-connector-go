/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the NotionalFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NotionalFilter{}

// NotionalFilter struct for NotionalFilter
type NotionalFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	PriceExponent        *int32  `json:"priceExponent,omitempty"`
	MinNotional          *string `json:"minNotional,omitempty"`
	ApplyMinToMarket     *bool   `json:"applyMinToMarket,omitempty"`
	MaxNotional          *string `json:"maxNotional,omitempty"`
	ApplyMaxToMarket     *bool   `json:"applyMaxToMarket,omitempty"`
	AvgPriceMins         *int32  `json:"avgPriceMins,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotionalFilter NotionalFilter

// NewNotionalFilter instantiates a new NotionalFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotionalFilter() *NotionalFilter {
	this := NotionalFilter{}
	return &this
}

// NewNotionalFilterWithDefaults instantiates a new NotionalFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotionalFilterWithDefaults() *NotionalFilter {
	this := NotionalFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *NotionalFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *NotionalFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *NotionalFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetPriceExponent returns the PriceExponent field value if set, zero value otherwise.
func (o *NotionalFilter) GetPriceExponent() int32 {
	if o == nil || common.IsNil(o.PriceExponent) {
		var ret int32
		return ret
	}
	return *o.PriceExponent
}

// GetPriceExponentOk returns a tuple with the PriceExponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalFilter) GetPriceExponentOk() (*int32, bool) {
	if o == nil || common.IsNil(o.PriceExponent) {
		return nil, false
	}
	return o.PriceExponent, true
}

// HasPriceExponent returns a boolean if a field has been set.
func (o *NotionalFilter) HasPriceExponent() bool {
	if o != nil && !common.IsNil(o.PriceExponent) {
		return true
	}

	return false
}

// SetPriceExponent gets a reference to the given int32 and assigns it to the PriceExponent field.
func (o *NotionalFilter) SetPriceExponent(v int32) {
	o.PriceExponent = &v
}

// GetMinNotional returns the MinNotional field value if set, zero value otherwise.
func (o *NotionalFilter) GetMinNotional() string {
	if o == nil || common.IsNil(o.MinNotional) {
		var ret string
		return ret
	}
	return *o.MinNotional
}

// GetMinNotionalOk returns a tuple with the MinNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalFilter) GetMinNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinNotional) {
		return nil, false
	}
	return o.MinNotional, true
}

// HasMinNotional returns a boolean if a field has been set.
func (o *NotionalFilter) HasMinNotional() bool {
	if o != nil && !common.IsNil(o.MinNotional) {
		return true
	}

	return false
}

// SetMinNotional gets a reference to the given string and assigns it to the MinNotional field.
func (o *NotionalFilter) SetMinNotional(v string) {
	o.MinNotional = &v
}

// GetApplyMinToMarket returns the ApplyMinToMarket field value if set, zero value otherwise.
func (o *NotionalFilter) GetApplyMinToMarket() bool {
	if o == nil || common.IsNil(o.ApplyMinToMarket) {
		var ret bool
		return ret
	}
	return *o.ApplyMinToMarket
}

// GetApplyMinToMarketOk returns a tuple with the ApplyMinToMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalFilter) GetApplyMinToMarketOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ApplyMinToMarket) {
		return nil, false
	}
	return o.ApplyMinToMarket, true
}

// HasApplyMinToMarket returns a boolean if a field has been set.
func (o *NotionalFilter) HasApplyMinToMarket() bool {
	if o != nil && !common.IsNil(o.ApplyMinToMarket) {
		return true
	}

	return false
}

// SetApplyMinToMarket gets a reference to the given bool and assigns it to the ApplyMinToMarket field.
func (o *NotionalFilter) SetApplyMinToMarket(v bool) {
	o.ApplyMinToMarket = &v
}

// GetMaxNotional returns the MaxNotional field value if set, zero value otherwise.
func (o *NotionalFilter) GetMaxNotional() string {
	if o == nil || common.IsNil(o.MaxNotional) {
		var ret string
		return ret
	}
	return *o.MaxNotional
}

// GetMaxNotionalOk returns a tuple with the MaxNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalFilter) GetMaxNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxNotional) {
		return nil, false
	}
	return o.MaxNotional, true
}

// HasMaxNotional returns a boolean if a field has been set.
func (o *NotionalFilter) HasMaxNotional() bool {
	if o != nil && !common.IsNil(o.MaxNotional) {
		return true
	}

	return false
}

// SetMaxNotional gets a reference to the given string and assigns it to the MaxNotional field.
func (o *NotionalFilter) SetMaxNotional(v string) {
	o.MaxNotional = &v
}

// GetApplyMaxToMarket returns the ApplyMaxToMarket field value if set, zero value otherwise.
func (o *NotionalFilter) GetApplyMaxToMarket() bool {
	if o == nil || common.IsNil(o.ApplyMaxToMarket) {
		var ret bool
		return ret
	}
	return *o.ApplyMaxToMarket
}

// GetApplyMaxToMarketOk returns a tuple with the ApplyMaxToMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalFilter) GetApplyMaxToMarketOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ApplyMaxToMarket) {
		return nil, false
	}
	return o.ApplyMaxToMarket, true
}

// HasApplyMaxToMarket returns a boolean if a field has been set.
func (o *NotionalFilter) HasApplyMaxToMarket() bool {
	if o != nil && !common.IsNil(o.ApplyMaxToMarket) {
		return true
	}

	return false
}

// SetApplyMaxToMarket gets a reference to the given bool and assigns it to the ApplyMaxToMarket field.
func (o *NotionalFilter) SetApplyMaxToMarket(v bool) {
	o.ApplyMaxToMarket = &v
}

// GetAvgPriceMins returns the AvgPriceMins field value if set, zero value otherwise.
func (o *NotionalFilter) GetAvgPriceMins() int32 {
	if o == nil || common.IsNil(o.AvgPriceMins) {
		var ret int32
		return ret
	}
	return *o.AvgPriceMins
}

// GetAvgPriceMinsOk returns a tuple with the AvgPriceMins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalFilter) GetAvgPriceMinsOk() (*int32, bool) {
	if o == nil || common.IsNil(o.AvgPriceMins) {
		return nil, false
	}
	return o.AvgPriceMins, true
}

// HasAvgPriceMins returns a boolean if a field has been set.
func (o *NotionalFilter) HasAvgPriceMins() bool {
	if o != nil && !common.IsNil(o.AvgPriceMins) {
		return true
	}

	return false
}

// SetAvgPriceMins gets a reference to the given int32 and assigns it to the AvgPriceMins field.
func (o *NotionalFilter) SetAvgPriceMins(v int32) {
	o.AvgPriceMins = &v
}

func (o NotionalFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotionalFilter) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.ApplyMinToMarket) {
		toSerialize["applyMinToMarket"] = o.ApplyMinToMarket
	}
	if !common.IsNil(o.MaxNotional) {
		toSerialize["maxNotional"] = o.MaxNotional
	}
	if !common.IsNil(o.ApplyMaxToMarket) {
		toSerialize["applyMaxToMarket"] = o.ApplyMaxToMarket
	}
	if !common.IsNil(o.AvgPriceMins) {
		toSerialize["avgPriceMins"] = o.AvgPriceMins
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NotionalFilter) UnmarshalJSON(data []byte) (err error) {
	varNotionalFilter := _NotionalFilter{}

	err = json.Unmarshal(data, &varNotionalFilter)

	if err != nil {
		return err
	}

	*o = NotionalFilter(varNotionalFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "priceExponent")
		delete(additionalProperties, "minNotional")
		delete(additionalProperties, "applyMinToMarket")
		delete(additionalProperties, "maxNotional")
		delete(additionalProperties, "applyMaxToMarket")
		delete(additionalProperties, "avgPriceMins")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotionalFilter struct {
	value *NotionalFilter
	isSet bool
}

func (v NullableNotionalFilter) Get() *NotionalFilter {
	return v.value
}

func (v *NullableNotionalFilter) Set(val *NotionalFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableNotionalFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableNotionalFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotionalFilter(val *NotionalFilter) *NullableNotionalFilter {
	return &NullableNotionalFilter{value: val, isSet: true}
}

func (v NullableNotionalFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotionalFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
