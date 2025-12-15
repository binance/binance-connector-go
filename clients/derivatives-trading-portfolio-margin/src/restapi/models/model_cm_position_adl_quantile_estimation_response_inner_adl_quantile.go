/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CmPositionAdlQuantileEstimationResponseInnerAdlQuantile type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CmPositionAdlQuantileEstimationResponseInnerAdlQuantile{}

// CmPositionAdlQuantileEstimationResponseInnerAdlQuantile struct for CmPositionAdlQuantileEstimationResponseInnerAdlQuantile
type CmPositionAdlQuantileEstimationResponseInnerAdlQuantile struct {
	LONG                 *int64 `json:"LONG,omitempty"`
	SHORT                *int64 `json:"SHORT,omitempty"`
	HEDGE                *int64 `json:"HEDGE,omitempty"`
	BOTH                 *int64 `json:"BOTH,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CmPositionAdlQuantileEstimationResponseInnerAdlQuantile CmPositionAdlQuantileEstimationResponseInnerAdlQuantile

// NewCmPositionAdlQuantileEstimationResponseInnerAdlQuantile instantiates a new CmPositionAdlQuantileEstimationResponseInnerAdlQuantile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmPositionAdlQuantileEstimationResponseInnerAdlQuantile() *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile {
	this := CmPositionAdlQuantileEstimationResponseInnerAdlQuantile{}
	return &this
}

// NewCmPositionAdlQuantileEstimationResponseInnerAdlQuantileWithDefaults instantiates a new CmPositionAdlQuantileEstimationResponseInnerAdlQuantile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmPositionAdlQuantileEstimationResponseInnerAdlQuantileWithDefaults() *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile {
	this := CmPositionAdlQuantileEstimationResponseInnerAdlQuantile{}
	return &this
}

// GetLONG returns the LONG field value if set, zero value otherwise.
func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) GetLONG() int64 {
	if o == nil || common.IsNil(o.LONG) {
		var ret int64
		return ret
	}
	return *o.LONG
}

// GetLONGOk returns a tuple with the LONG field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) GetLONGOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LONG) {
		return nil, false
	}
	return o.LONG, true
}

// HasLONG returns a boolean if a field has been set.
func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) HasLONG() bool {
	if o != nil && !common.IsNil(o.LONG) {
		return true
	}

	return false
}

// SetLONG gets a reference to the given int64 and assigns it to the LONG field.
func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) SetLONG(v int64) {
	o.LONG = &v
}

// GetSHORT returns the SHORT field value if set, zero value otherwise.
func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) GetSHORT() int64 {
	if o == nil || common.IsNil(o.SHORT) {
		var ret int64
		return ret
	}
	return *o.SHORT
}

// GetSHORTOk returns a tuple with the SHORT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) GetSHORTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.SHORT) {
		return nil, false
	}
	return o.SHORT, true
}

// HasSHORT returns a boolean if a field has been set.
func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) HasSHORT() bool {
	if o != nil && !common.IsNil(o.SHORT) {
		return true
	}

	return false
}

// SetSHORT gets a reference to the given int64 and assigns it to the SHORT field.
func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) SetSHORT(v int64) {
	o.SHORT = &v
}

// GetHEDGE returns the HEDGE field value if set, zero value otherwise.
func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) GetHEDGE() int64 {
	if o == nil || common.IsNil(o.HEDGE) {
		var ret int64
		return ret
	}
	return *o.HEDGE
}

// GetHEDGEOk returns a tuple with the HEDGE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) GetHEDGEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.HEDGE) {
		return nil, false
	}
	return o.HEDGE, true
}

// HasHEDGE returns a boolean if a field has been set.
func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) HasHEDGE() bool {
	if o != nil && !common.IsNil(o.HEDGE) {
		return true
	}

	return false
}

// SetHEDGE gets a reference to the given int64 and assigns it to the HEDGE field.
func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) SetHEDGE(v int64) {
	o.HEDGE = &v
}

// GetBOTH returns the BOTH field value if set, zero value otherwise.
func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) GetBOTH() int64 {
	if o == nil || common.IsNil(o.BOTH) {
		var ret int64
		return ret
	}
	return *o.BOTH
}

// GetBOTHOk returns a tuple with the BOTH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) GetBOTHOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BOTH) {
		return nil, false
	}
	return o.BOTH, true
}

// HasBOTH returns a boolean if a field has been set.
func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) HasBOTH() bool {
	if o != nil && !common.IsNil(o.BOTH) {
		return true
	}

	return false
}

// SetBOTH gets a reference to the given int64 and assigns it to the BOTH field.
func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) SetBOTH(v int64) {
	o.BOTH = &v
}

func (o CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LONG) {
		toSerialize["LONG"] = o.LONG
	}
	if !common.IsNil(o.SHORT) {
		toSerialize["SHORT"] = o.SHORT
	}
	if !common.IsNil(o.HEDGE) {
		toSerialize["HEDGE"] = o.HEDGE
	}
	if !common.IsNil(o.BOTH) {
		toSerialize["BOTH"] = o.BOTH
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) UnmarshalJSON(data []byte) (err error) {
	varCmPositionAdlQuantileEstimationResponseInnerAdlQuantile := _CmPositionAdlQuantileEstimationResponseInnerAdlQuantile{}

	err = json.Unmarshal(data, &varCmPositionAdlQuantileEstimationResponseInnerAdlQuantile)

	if err != nil {
		return err
	}

	*o = CmPositionAdlQuantileEstimationResponseInnerAdlQuantile(varCmPositionAdlQuantileEstimationResponseInnerAdlQuantile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "LONG")
		delete(additionalProperties, "SHORT")
		delete(additionalProperties, "HEDGE")
		delete(additionalProperties, "BOTH")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCmPositionAdlQuantileEstimationResponseInnerAdlQuantile struct {
	value *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile
	isSet bool
}

func (v NullableCmPositionAdlQuantileEstimationResponseInnerAdlQuantile) Get() *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile {
	return v.value
}

func (v *NullableCmPositionAdlQuantileEstimationResponseInnerAdlQuantile) Set(val *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) {
	v.value = val
	v.isSet = true
}

func (v NullableCmPositionAdlQuantileEstimationResponseInnerAdlQuantile) IsSet() bool {
	return v.isSet
}

func (v *NullableCmPositionAdlQuantileEstimationResponseInnerAdlQuantile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmPositionAdlQuantileEstimationResponseInnerAdlQuantile(val *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) *NullableCmPositionAdlQuantileEstimationResponseInnerAdlQuantile {
	return &NullableCmPositionAdlQuantileEstimationResponseInnerAdlQuantile{value: val, isSet: true}
}

func (v NullableCmPositionAdlQuantileEstimationResponseInnerAdlQuantile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmPositionAdlQuantileEstimationResponseInnerAdlQuantile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
