/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the LotSizeFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &LotSizeFilter{}

// LotSizeFilter struct for LotSizeFilter
type LotSizeFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	QtyExponent          *int32  `json:"qtyExponent,omitempty"`
	MinQty               *string `json:"minQty,omitempty"`
	MaxQty               *string `json:"maxQty,omitempty"`
	StepSize             *string `json:"stepSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LotSizeFilter LotSizeFilter

// NewLotSizeFilter instantiates a new LotSizeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLotSizeFilter() *LotSizeFilter {
	this := LotSizeFilter{}
	return &this
}

// NewLotSizeFilterWithDefaults instantiates a new LotSizeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLotSizeFilterWithDefaults() *LotSizeFilter {
	this := LotSizeFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *LotSizeFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LotSizeFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *LotSizeFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *LotSizeFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetQtyExponent returns the QtyExponent field value if set, zero value otherwise.
func (o *LotSizeFilter) GetQtyExponent() int32 {
	if o == nil || common.IsNil(o.QtyExponent) {
		var ret int32
		return ret
	}
	return *o.QtyExponent
}

// GetQtyExponentOk returns a tuple with the QtyExponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LotSizeFilter) GetQtyExponentOk() (*int32, bool) {
	if o == nil || common.IsNil(o.QtyExponent) {
		return nil, false
	}
	return o.QtyExponent, true
}

// HasQtyExponent returns a boolean if a field has been set.
func (o *LotSizeFilter) HasQtyExponent() bool {
	if o != nil && !common.IsNil(o.QtyExponent) {
		return true
	}

	return false
}

// SetQtyExponent gets a reference to the given int32 and assigns it to the QtyExponent field.
func (o *LotSizeFilter) SetQtyExponent(v int32) {
	o.QtyExponent = &v
}

// GetMinQty returns the MinQty field value if set, zero value otherwise.
func (o *LotSizeFilter) GetMinQty() string {
	if o == nil || common.IsNil(o.MinQty) {
		var ret string
		return ret
	}
	return *o.MinQty
}

// GetMinQtyOk returns a tuple with the MinQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LotSizeFilter) GetMinQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinQty) {
		return nil, false
	}
	return o.MinQty, true
}

// HasMinQty returns a boolean if a field has been set.
func (o *LotSizeFilter) HasMinQty() bool {
	if o != nil && !common.IsNil(o.MinQty) {
		return true
	}

	return false
}

// SetMinQty gets a reference to the given string and assigns it to the MinQty field.
func (o *LotSizeFilter) SetMinQty(v string) {
	o.MinQty = &v
}

// GetMaxQty returns the MaxQty field value if set, zero value otherwise.
func (o *LotSizeFilter) GetMaxQty() string {
	if o == nil || common.IsNil(o.MaxQty) {
		var ret string
		return ret
	}
	return *o.MaxQty
}

// GetMaxQtyOk returns a tuple with the MaxQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LotSizeFilter) GetMaxQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxQty) {
		return nil, false
	}
	return o.MaxQty, true
}

// HasMaxQty returns a boolean if a field has been set.
func (o *LotSizeFilter) HasMaxQty() bool {
	if o != nil && !common.IsNil(o.MaxQty) {
		return true
	}

	return false
}

// SetMaxQty gets a reference to the given string and assigns it to the MaxQty field.
func (o *LotSizeFilter) SetMaxQty(v string) {
	o.MaxQty = &v
}

// GetStepSize returns the StepSize field value if set, zero value otherwise.
func (o *LotSizeFilter) GetStepSize() string {
	if o == nil || common.IsNil(o.StepSize) {
		var ret string
		return ret
	}
	return *o.StepSize
}

// GetStepSizeOk returns a tuple with the StepSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LotSizeFilter) GetStepSizeOk() (*string, bool) {
	if o == nil || common.IsNil(o.StepSize) {
		return nil, false
	}
	return o.StepSize, true
}

// HasStepSize returns a boolean if a field has been set.
func (o *LotSizeFilter) HasStepSize() bool {
	if o != nil && !common.IsNil(o.StepSize) {
		return true
	}

	return false
}

// SetStepSize gets a reference to the given string and assigns it to the StepSize field.
func (o *LotSizeFilter) SetStepSize(v string) {
	o.StepSize = &v
}

func (o LotSizeFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LotSizeFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.QtyExponent) {
		toSerialize["qtyExponent"] = o.QtyExponent
	}
	if !common.IsNil(o.MinQty) {
		toSerialize["minQty"] = o.MinQty
	}
	if !common.IsNil(o.MaxQty) {
		toSerialize["maxQty"] = o.MaxQty
	}
	if !common.IsNil(o.StepSize) {
		toSerialize["stepSize"] = o.StepSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LotSizeFilter) UnmarshalJSON(data []byte) (err error) {
	varLotSizeFilter := _LotSizeFilter{}

	err = json.Unmarshal(data, &varLotSizeFilter)

	if err != nil {
		return err
	}

	*o = LotSizeFilter(varLotSizeFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "qtyExponent")
		delete(additionalProperties, "minQty")
		delete(additionalProperties, "maxQty")
		delete(additionalProperties, "stepSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLotSizeFilter struct {
	value *LotSizeFilter
	isSet bool
}

func (v NullableLotSizeFilter) Get() *LotSizeFilter {
	return v.value
}

func (v *NullableLotSizeFilter) Set(val *LotSizeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableLotSizeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableLotSizeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLotSizeFilter(val *LotSizeFilter) *NullableLotSizeFilter {
	return &NullableLotSizeFilter{value: val, isSet: true}
}

func (v NullableLotSizeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLotSizeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
