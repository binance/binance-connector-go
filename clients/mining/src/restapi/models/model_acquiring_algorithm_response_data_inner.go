/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AcquiringAlgorithmResponseDataInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AcquiringAlgorithmResponseDataInner{}

// AcquiringAlgorithmResponseDataInner struct for AcquiringAlgorithmResponseDataInner
type AcquiringAlgorithmResponseDataInner struct {
	AlgoName             *string `json:"algoName,omitempty"`
	AlgoId               *int64  `json:"algoId,omitempty"`
	PoolIndex            *int64  `json:"poolIndex,omitempty"`
	Unit                 *string `json:"unit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AcquiringAlgorithmResponseDataInner AcquiringAlgorithmResponseDataInner

// NewAcquiringAlgorithmResponseDataInner instantiates a new AcquiringAlgorithmResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcquiringAlgorithmResponseDataInner() *AcquiringAlgorithmResponseDataInner {
	this := AcquiringAlgorithmResponseDataInner{}
	return &this
}

// NewAcquiringAlgorithmResponseDataInnerWithDefaults instantiates a new AcquiringAlgorithmResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcquiringAlgorithmResponseDataInnerWithDefaults() *AcquiringAlgorithmResponseDataInner {
	this := AcquiringAlgorithmResponseDataInner{}
	return &this
}

// GetAlgoName returns the AlgoName field value if set, zero value otherwise.
func (o *AcquiringAlgorithmResponseDataInner) GetAlgoName() string {
	if o == nil || common.IsNil(o.AlgoName) {
		var ret string
		return ret
	}
	return *o.AlgoName
}

// GetAlgoNameOk returns a tuple with the AlgoName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquiringAlgorithmResponseDataInner) GetAlgoNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.AlgoName) {
		return nil, false
	}
	return o.AlgoName, true
}

// HasAlgoName returns a boolean if a field has been set.
func (o *AcquiringAlgorithmResponseDataInner) HasAlgoName() bool {
	if o != nil && !common.IsNil(o.AlgoName) {
		return true
	}

	return false
}

// SetAlgoName gets a reference to the given string and assigns it to the AlgoName field.
func (o *AcquiringAlgorithmResponseDataInner) SetAlgoName(v string) {
	o.AlgoName = &v
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *AcquiringAlgorithmResponseDataInner) GetAlgoId() int64 {
	if o == nil || common.IsNil(o.AlgoId) {
		var ret int64
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquiringAlgorithmResponseDataInner) GetAlgoIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *AcquiringAlgorithmResponseDataInner) HasAlgoId() bool {
	if o != nil && !common.IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given int64 and assigns it to the AlgoId field.
func (o *AcquiringAlgorithmResponseDataInner) SetAlgoId(v int64) {
	o.AlgoId = &v
}

// GetPoolIndex returns the PoolIndex field value if set, zero value otherwise.
func (o *AcquiringAlgorithmResponseDataInner) GetPoolIndex() int64 {
	if o == nil || common.IsNil(o.PoolIndex) {
		var ret int64
		return ret
	}
	return *o.PoolIndex
}

// GetPoolIndexOk returns a tuple with the PoolIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquiringAlgorithmResponseDataInner) GetPoolIndexOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PoolIndex) {
		return nil, false
	}
	return o.PoolIndex, true
}

// HasPoolIndex returns a boolean if a field has been set.
func (o *AcquiringAlgorithmResponseDataInner) HasPoolIndex() bool {
	if o != nil && !common.IsNil(o.PoolIndex) {
		return true
	}

	return false
}

// SetPoolIndex gets a reference to the given int64 and assigns it to the PoolIndex field.
func (o *AcquiringAlgorithmResponseDataInner) SetPoolIndex(v int64) {
	o.PoolIndex = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *AcquiringAlgorithmResponseDataInner) GetUnit() string {
	if o == nil || common.IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquiringAlgorithmResponseDataInner) GetUnitOk() (*string, bool) {
	if o == nil || common.IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *AcquiringAlgorithmResponseDataInner) HasUnit() bool {
	if o != nil && !common.IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *AcquiringAlgorithmResponseDataInner) SetUnit(v string) {
	o.Unit = &v
}

func (o AcquiringAlgorithmResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcquiringAlgorithmResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AlgoName) {
		toSerialize["algoName"] = o.AlgoName
	}
	if !common.IsNil(o.AlgoId) {
		toSerialize["algoId"] = o.AlgoId
	}
	if !common.IsNil(o.PoolIndex) {
		toSerialize["poolIndex"] = o.PoolIndex
	}
	if !common.IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AcquiringAlgorithmResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	varAcquiringAlgorithmResponseDataInner := _AcquiringAlgorithmResponseDataInner{}

	err = json.Unmarshal(data, &varAcquiringAlgorithmResponseDataInner)

	if err != nil {
		return err
	}

	*o = AcquiringAlgorithmResponseDataInner(varAcquiringAlgorithmResponseDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "algoName")
		delete(additionalProperties, "algoId")
		delete(additionalProperties, "poolIndex")
		delete(additionalProperties, "unit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcquiringAlgorithmResponseDataInner struct {
	value *AcquiringAlgorithmResponseDataInner
	isSet bool
}

func (v NullableAcquiringAlgorithmResponseDataInner) Get() *AcquiringAlgorithmResponseDataInner {
	return v.value
}

func (v *NullableAcquiringAlgorithmResponseDataInner) Set(val *AcquiringAlgorithmResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAcquiringAlgorithmResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAcquiringAlgorithmResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcquiringAlgorithmResponseDataInner(val *AcquiringAlgorithmResponseDataInner) *NullableAcquiringAlgorithmResponseDataInner {
	return &NullableAcquiringAlgorithmResponseDataInner{value: val, isSet: true}
}

func (v NullableAcquiringAlgorithmResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcquiringAlgorithmResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
