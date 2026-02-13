/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the UmPositionAdlQuantileEstimationResponseInnerAdlQuantile type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UmPositionAdlQuantileEstimationResponseInnerAdlQuantile{}

// UmPositionAdlQuantileEstimationResponseInnerAdlQuantile struct for UmPositionAdlQuantileEstimationResponseInnerAdlQuantile
type UmPositionAdlQuantileEstimationResponseInnerAdlQuantile struct {
	LONG                 *int64 `json:"LONG,omitempty"`
	SHORT                *int64 `json:"SHORT,omitempty"`
	BOTH                 *int64 `json:"BOTH,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UmPositionAdlQuantileEstimationResponseInnerAdlQuantile UmPositionAdlQuantileEstimationResponseInnerAdlQuantile

// NewUmPositionAdlQuantileEstimationResponseInnerAdlQuantile instantiates a new UmPositionAdlQuantileEstimationResponseInnerAdlQuantile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmPositionAdlQuantileEstimationResponseInnerAdlQuantile() *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile {
	this := UmPositionAdlQuantileEstimationResponseInnerAdlQuantile{}
	return &this
}

// NewUmPositionAdlQuantileEstimationResponseInnerAdlQuantileWithDefaults instantiates a new UmPositionAdlQuantileEstimationResponseInnerAdlQuantile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmPositionAdlQuantileEstimationResponseInnerAdlQuantileWithDefaults() *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile {
	this := UmPositionAdlQuantileEstimationResponseInnerAdlQuantile{}
	return &this
}

// GetLONG returns the LONG field value if set, zero value otherwise.
func (o *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) GetLONG() int64 {
	if o == nil || common.IsNil(o.LONG) {
		var ret int64
		return ret
	}
	return *o.LONG
}

// GetLONGOk returns a tuple with the LONG field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) GetLONGOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LONG) {
		return nil, false
	}
	return o.LONG, true
}

// HasLONG returns a boolean if a field has been set.
func (o *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) HasLONG() bool {
	if o != nil && !common.IsNil(o.LONG) {
		return true
	}

	return false
}

// SetLONG gets a reference to the given int64 and assigns it to the LONG field.
func (o *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) SetLONG(v int64) {
	o.LONG = &v
}

// GetSHORT returns the SHORT field value if set, zero value otherwise.
func (o *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) GetSHORT() int64 {
	if o == nil || common.IsNil(o.SHORT) {
		var ret int64
		return ret
	}
	return *o.SHORT
}

// GetSHORTOk returns a tuple with the SHORT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) GetSHORTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.SHORT) {
		return nil, false
	}
	return o.SHORT, true
}

// HasSHORT returns a boolean if a field has been set.
func (o *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) HasSHORT() bool {
	if o != nil && !common.IsNil(o.SHORT) {
		return true
	}

	return false
}

// SetSHORT gets a reference to the given int64 and assigns it to the SHORT field.
func (o *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) SetSHORT(v int64) {
	o.SHORT = &v
}

// GetBOTH returns the BOTH field value if set, zero value otherwise.
func (o *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) GetBOTH() int64 {
	if o == nil || common.IsNil(o.BOTH) {
		var ret int64
		return ret
	}
	return *o.BOTH
}

// GetBOTHOk returns a tuple with the BOTH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) GetBOTHOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BOTH) {
		return nil, false
	}
	return o.BOTH, true
}

// HasBOTH returns a boolean if a field has been set.
func (o *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) HasBOTH() bool {
	if o != nil && !common.IsNil(o.BOTH) {
		return true
	}

	return false
}

// SetBOTH gets a reference to the given int64 and assigns it to the BOTH field.
func (o *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) SetBOTH(v int64) {
	o.BOTH = &v
}

func (o UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LONG) {
		toSerialize["LONG"] = o.LONG
	}
	if !common.IsNil(o.SHORT) {
		toSerialize["SHORT"] = o.SHORT
	}
	if !common.IsNil(o.BOTH) {
		toSerialize["BOTH"] = o.BOTH
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) UnmarshalJSON(data []byte) (err error) {
	varUmPositionAdlQuantileEstimationResponseInnerAdlQuantile := _UmPositionAdlQuantileEstimationResponseInnerAdlQuantile{}

	err = json.Unmarshal(data, &varUmPositionAdlQuantileEstimationResponseInnerAdlQuantile)

	if err != nil {
		return err
	}

	*o = UmPositionAdlQuantileEstimationResponseInnerAdlQuantile(varUmPositionAdlQuantileEstimationResponseInnerAdlQuantile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "LONG")
		delete(additionalProperties, "SHORT")
		delete(additionalProperties, "BOTH")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUmPositionAdlQuantileEstimationResponseInnerAdlQuantile struct {
	value *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile
	isSet bool
}

func (v NullableUmPositionAdlQuantileEstimationResponseInnerAdlQuantile) Get() *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile {
	return v.value
}

func (v *NullableUmPositionAdlQuantileEstimationResponseInnerAdlQuantile) Set(val *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) {
	v.value = val
	v.isSet = true
}

func (v NullableUmPositionAdlQuantileEstimationResponseInnerAdlQuantile) IsSet() bool {
	return v.isSet
}

func (v *NullableUmPositionAdlQuantileEstimationResponseInnerAdlQuantile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmPositionAdlQuantileEstimationResponseInnerAdlQuantile(val *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) *NullableUmPositionAdlQuantileEstimationResponseInnerAdlQuantile {
	return &NullableUmPositionAdlQuantileEstimationResponseInnerAdlQuantile{value: val, isSet: true}
}

func (v NullableUmPositionAdlQuantileEstimationResponseInnerAdlQuantile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmPositionAdlQuantileEstimationResponseInnerAdlQuantile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
