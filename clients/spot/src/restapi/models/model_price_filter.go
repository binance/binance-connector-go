/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PriceFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PriceFilter{}

// PriceFilter struct for PriceFilter
type PriceFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	PriceExponent        *int32  `json:"priceExponent,omitempty"`
	MinPrice             *string `json:"minPrice,omitempty"`
	MaxPrice             *string `json:"maxPrice,omitempty"`
	TickSize             *string `json:"tickSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PriceFilter PriceFilter

// NewPriceFilter instantiates a new PriceFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceFilter() *PriceFilter {
	this := PriceFilter{}
	return &this
}

// NewPriceFilterWithDefaults instantiates a new PriceFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceFilterWithDefaults() *PriceFilter {
	this := PriceFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *PriceFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *PriceFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *PriceFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetPriceExponent returns the PriceExponent field value if set, zero value otherwise.
func (o *PriceFilter) GetPriceExponent() int32 {
	if o == nil || common.IsNil(o.PriceExponent) {
		var ret int32
		return ret
	}
	return *o.PriceExponent
}

// GetPriceExponentOk returns a tuple with the PriceExponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceFilter) GetPriceExponentOk() (*int32, bool) {
	if o == nil || common.IsNil(o.PriceExponent) {
		return nil, false
	}
	return o.PriceExponent, true
}

// HasPriceExponent returns a boolean if a field has been set.
func (o *PriceFilter) HasPriceExponent() bool {
	if o != nil && !common.IsNil(o.PriceExponent) {
		return true
	}

	return false
}

// SetPriceExponent gets a reference to the given int32 and assigns it to the PriceExponent field.
func (o *PriceFilter) SetPriceExponent(v int32) {
	o.PriceExponent = &v
}

// GetMinPrice returns the MinPrice field value if set, zero value otherwise.
func (o *PriceFilter) GetMinPrice() string {
	if o == nil || common.IsNil(o.MinPrice) {
		var ret string
		return ret
	}
	return *o.MinPrice
}

// GetMinPriceOk returns a tuple with the MinPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceFilter) GetMinPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinPrice) {
		return nil, false
	}
	return o.MinPrice, true
}

// HasMinPrice returns a boolean if a field has been set.
func (o *PriceFilter) HasMinPrice() bool {
	if o != nil && !common.IsNil(o.MinPrice) {
		return true
	}

	return false
}

// SetMinPrice gets a reference to the given string and assigns it to the MinPrice field.
func (o *PriceFilter) SetMinPrice(v string) {
	o.MinPrice = &v
}

// GetMaxPrice returns the MaxPrice field value if set, zero value otherwise.
func (o *PriceFilter) GetMaxPrice() string {
	if o == nil || common.IsNil(o.MaxPrice) {
		var ret string
		return ret
	}
	return *o.MaxPrice
}

// GetMaxPriceOk returns a tuple with the MaxPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceFilter) GetMaxPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxPrice) {
		return nil, false
	}
	return o.MaxPrice, true
}

// HasMaxPrice returns a boolean if a field has been set.
func (o *PriceFilter) HasMaxPrice() bool {
	if o != nil && !common.IsNil(o.MaxPrice) {
		return true
	}

	return false
}

// SetMaxPrice gets a reference to the given string and assigns it to the MaxPrice field.
func (o *PriceFilter) SetMaxPrice(v string) {
	o.MaxPrice = &v
}

// GetTickSize returns the TickSize field value if set, zero value otherwise.
func (o *PriceFilter) GetTickSize() string {
	if o == nil || common.IsNil(o.TickSize) {
		var ret string
		return ret
	}
	return *o.TickSize
}

// GetTickSizeOk returns a tuple with the TickSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceFilter) GetTickSizeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TickSize) {
		return nil, false
	}
	return o.TickSize, true
}

// HasTickSize returns a boolean if a field has been set.
func (o *PriceFilter) HasTickSize() bool {
	if o != nil && !common.IsNil(o.TickSize) {
		return true
	}

	return false
}

// SetTickSize gets a reference to the given string and assigns it to the TickSize field.
func (o *PriceFilter) SetTickSize(v string) {
	o.TickSize = &v
}

func (o PriceFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.PriceExponent) {
		toSerialize["priceExponent"] = o.PriceExponent
	}
	if !common.IsNil(o.MinPrice) {
		toSerialize["minPrice"] = o.MinPrice
	}
	if !common.IsNil(o.MaxPrice) {
		toSerialize["maxPrice"] = o.MaxPrice
	}
	if !common.IsNil(o.TickSize) {
		toSerialize["tickSize"] = o.TickSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PriceFilter) UnmarshalJSON(data []byte) (err error) {
	varPriceFilter := _PriceFilter{}

	err = json.Unmarshal(data, &varPriceFilter)

	if err != nil {
		return err
	}

	*o = PriceFilter(varPriceFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "priceExponent")
		delete(additionalProperties, "minPrice")
		delete(additionalProperties, "maxPrice")
		delete(additionalProperties, "tickSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePriceFilter struct {
	value *PriceFilter
	isSet bool
}

func (v NullablePriceFilter) Get() *PriceFilter {
	return v.value
}

func (v *NullablePriceFilter) Set(val *PriceFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceFilter(val *PriceFilter) *NullablePriceFilter {
	return &NullablePriceFilter{value: val, isSet: true}
}

func (v NullablePriceFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
