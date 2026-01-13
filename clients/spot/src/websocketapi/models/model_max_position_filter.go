/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MaxPositionFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MaxPositionFilter{}

// MaxPositionFilter struct for MaxPositionFilter
type MaxPositionFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	QtyExponent          *int32  `json:"qtyExponent,omitempty"`
	MaxPosition          *string `json:"maxPosition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MaxPositionFilter MaxPositionFilter

// NewMaxPositionFilter instantiates a new MaxPositionFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaxPositionFilter() *MaxPositionFilter {
	this := MaxPositionFilter{}
	return &this
}

// NewMaxPositionFilterWithDefaults instantiates a new MaxPositionFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaxPositionFilterWithDefaults() *MaxPositionFilter {
	this := MaxPositionFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *MaxPositionFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxPositionFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *MaxPositionFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *MaxPositionFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetQtyExponent returns the QtyExponent field value if set, zero value otherwise.
func (o *MaxPositionFilter) GetQtyExponent() int32 {
	if o == nil || common.IsNil(o.QtyExponent) {
		var ret int32
		return ret
	}
	return *o.QtyExponent
}

// GetQtyExponentOk returns a tuple with the QtyExponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxPositionFilter) GetQtyExponentOk() (*int32, bool) {
	if o == nil || common.IsNil(o.QtyExponent) {
		return nil, false
	}
	return o.QtyExponent, true
}

// HasQtyExponent returns a boolean if a field has been set.
func (o *MaxPositionFilter) HasQtyExponent() bool {
	if o != nil && !common.IsNil(o.QtyExponent) {
		return true
	}

	return false
}

// SetQtyExponent gets a reference to the given int32 and assigns it to the QtyExponent field.
func (o *MaxPositionFilter) SetQtyExponent(v int32) {
	o.QtyExponent = &v
}

// GetMaxPosition returns the MaxPosition field value if set, zero value otherwise.
func (o *MaxPositionFilter) GetMaxPosition() string {
	if o == nil || common.IsNil(o.MaxPosition) {
		var ret string
		return ret
	}
	return *o.MaxPosition
}

// GetMaxPositionOk returns a tuple with the MaxPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxPositionFilter) GetMaxPositionOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxPosition) {
		return nil, false
	}
	return o.MaxPosition, true
}

// HasMaxPosition returns a boolean if a field has been set.
func (o *MaxPositionFilter) HasMaxPosition() bool {
	if o != nil && !common.IsNil(o.MaxPosition) {
		return true
	}

	return false
}

// SetMaxPosition gets a reference to the given string and assigns it to the MaxPosition field.
func (o *MaxPositionFilter) SetMaxPosition(v string) {
	o.MaxPosition = &v
}

func (o MaxPositionFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaxPositionFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.QtyExponent) {
		toSerialize["qtyExponent"] = o.QtyExponent
	}
	if !common.IsNil(o.MaxPosition) {
		toSerialize["maxPosition"] = o.MaxPosition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MaxPositionFilter) UnmarshalJSON(data []byte) (err error) {
	varMaxPositionFilter := _MaxPositionFilter{}

	err = json.Unmarshal(data, &varMaxPositionFilter)

	if err != nil {
		return err
	}

	*o = MaxPositionFilter(varMaxPositionFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "qtyExponent")
		delete(additionalProperties, "maxPosition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMaxPositionFilter struct {
	value *MaxPositionFilter
	isSet bool
}

func (v NullableMaxPositionFilter) Get() *MaxPositionFilter {
	return v.value
}

func (v *NullableMaxPositionFilter) Set(val *MaxPositionFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMaxPositionFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMaxPositionFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaxPositionFilter(val *MaxPositionFilter) *NullableMaxPositionFilter {
	return &NullableMaxPositionFilter{value: val, isSet: true}
}

func (v NullableMaxPositionFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaxPositionFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
