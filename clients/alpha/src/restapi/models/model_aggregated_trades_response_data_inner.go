/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AggregatedTradesResponseDataInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AggregatedTradesResponseDataInner{}

// AggregatedTradesResponseDataInner struct for AggregatedTradesResponseDataInner
type AggregatedTradesResponseDataInner struct {
	A                    *int64  `json:"a,omitempty"`
	P                    *string `json:"p,omitempty"`
	Q                    *string `json:"q,omitempty"`
	F                    *int64  `json:"f,omitempty"`
	L                    *int64  `json:"l,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	M                    *bool   `json:"m,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AggregatedTradesResponseDataInner AggregatedTradesResponseDataInner

// NewAggregatedTradesResponseDataInner instantiates a new AggregatedTradesResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregatedTradesResponseDataInner() *AggregatedTradesResponseDataInner {
	this := AggregatedTradesResponseDataInner{}
	return &this
}

// NewAggregatedTradesResponseDataInnerWithDefaults instantiates a new AggregatedTradesResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregatedTradesResponseDataInnerWithDefaults() *AggregatedTradesResponseDataInner {
	this := AggregatedTradesResponseDataInner{}
	return &this
}

// GetA returns the A field value if set, zero value otherwise.
func (o *AggregatedTradesResponseDataInner) GetA() int64 {
	if o == nil || common.IsNil(o.A) {
		var ret int64
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedTradesResponseDataInner) GetAOk() (*int64, bool) {
	if o == nil || common.IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *AggregatedTradesResponseDataInner) HasA() bool {
	if o != nil && !common.IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given int64 and assigns it to the A field.
func (o *AggregatedTradesResponseDataInner) SetA(v int64) {
	o.A = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *AggregatedTradesResponseDataInner) GetP() string {
	if o == nil || common.IsNil(o.P) {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedTradesResponseDataInner) GetPOk() (*string, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *AggregatedTradesResponseDataInner) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *AggregatedTradesResponseDataInner) SetP(v string) {
	o.P = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *AggregatedTradesResponseDataInner) GetQ() string {
	if o == nil || common.IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedTradesResponseDataInner) GetQOk() (*string, bool) {
	if o == nil || common.IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *AggregatedTradesResponseDataInner) HasQ() bool {
	if o != nil && !common.IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *AggregatedTradesResponseDataInner) SetQ(v string) {
	o.Q = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *AggregatedTradesResponseDataInner) GetF() int64 {
	if o == nil || common.IsNil(o.F) {
		var ret int64
		return ret
	}
	return *o.F
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedTradesResponseDataInner) GetFOk() (*int64, bool) {
	if o == nil || common.IsNil(o.F) {
		return nil, false
	}
	return o.F, true
}

// HasF returns a boolean if a field has been set.
func (o *AggregatedTradesResponseDataInner) HasF() bool {
	if o != nil && !common.IsNil(o.F) {
		return true
	}

	return false
}

// SetF gets a reference to the given int64 and assigns it to the F field.
func (o *AggregatedTradesResponseDataInner) SetF(v int64) {
	o.F = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *AggregatedTradesResponseDataInner) GetL() int64 {
	if o == nil || common.IsNil(o.L) {
		var ret int64
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedTradesResponseDataInner) GetLOk() (*int64, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *AggregatedTradesResponseDataInner) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given int64 and assigns it to the L field.
func (o *AggregatedTradesResponseDataInner) SetL(v int64) {
	o.L = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *AggregatedTradesResponseDataInner) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedTradesResponseDataInner) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *AggregatedTradesResponseDataInner) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *AggregatedTradesResponseDataInner) SetT(v int64) {
	o.T = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *AggregatedTradesResponseDataInner) GetM() bool {
	if o == nil || common.IsNil(o.M) {
		var ret bool
		return ret
	}
	return *o.M
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedTradesResponseDataInner) GetMOk() (*bool, bool) {
	if o == nil || common.IsNil(o.M) {
		return nil, false
	}
	return o.M, true
}

// HasM returns a boolean if a field has been set.
func (o *AggregatedTradesResponseDataInner) HasM() bool {
	if o != nil && !common.IsNil(o.M) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *AggregatedTradesResponseDataInner) SetM(v bool) {
	o.M = &v
}

func (o AggregatedTradesResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregatedTradesResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.A) {
		toSerialize["a"] = o.A
	}
	if !common.IsNil(o.P) {
		toSerialize["p"] = o.P
	}
	if !common.IsNil(o.Q) {
		toSerialize["q"] = o.Q
	}
	if !common.IsNil(o.F) {
		toSerialize["f"] = o.F
	}
	if !common.IsNil(o.L) {
		toSerialize["l"] = o.L
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.M) {
		toSerialize["m"] = o.M
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AggregatedTradesResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	varAggregatedTradesResponseDataInner := _AggregatedTradesResponseDataInner{}

	err = json.Unmarshal(data, &varAggregatedTradesResponseDataInner)

	if err != nil {
		return err
	}

	*o = AggregatedTradesResponseDataInner(varAggregatedTradesResponseDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "a")
		delete(additionalProperties, "p")
		delete(additionalProperties, "q")
		delete(additionalProperties, "f")
		delete(additionalProperties, "l")
		delete(additionalProperties, "T")
		delete(additionalProperties, "m")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAggregatedTradesResponseDataInner struct {
	value *AggregatedTradesResponseDataInner
	isSet bool
}

func (v NullableAggregatedTradesResponseDataInner) Get() *AggregatedTradesResponseDataInner {
	return v.value
}

func (v *NullableAggregatedTradesResponseDataInner) Set(val *AggregatedTradesResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregatedTradesResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregatedTradesResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregatedTradesResponseDataInner(val *AggregatedTradesResponseDataInner) *NullableAggregatedTradesResponseDataInner {
	return &NullableAggregatedTradesResponseDataInner{value: val, isSet: true}
}

func (v NullableAggregatedTradesResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregatedTradesResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
