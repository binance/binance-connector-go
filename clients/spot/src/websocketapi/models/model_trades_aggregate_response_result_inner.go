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

// checks if the TradesAggregateResponseResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradesAggregateResponseResultInner{}

// TradesAggregateResponseResultInner struct for TradesAggregateResponseResultInner
type TradesAggregateResponseResultInner struct {
	A                    *int64  `json:"a,omitempty"`
	P                    *string `json:"p,omitempty"`
	Q                    *string `json:"q,omitempty"`
	F                    *int64  `json:"f,omitempty"`
	L                    *int64  `json:"l,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	Smallm               *bool   `json:"m,omitempty"`
	M                    *bool   `json:"M,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradesAggregateResponseResultInner TradesAggregateResponseResultInner

// NewTradesAggregateResponseResultInner instantiates a new TradesAggregateResponseResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradesAggregateResponseResultInner() *TradesAggregateResponseResultInner {
	this := TradesAggregateResponseResultInner{}
	return &this
}

// NewTradesAggregateResponseResultInnerWithDefaults instantiates a new TradesAggregateResponseResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradesAggregateResponseResultInnerWithDefaults() *TradesAggregateResponseResultInner {
	this := TradesAggregateResponseResultInner{}
	return &this
}

// GetA returns the A field value if set, zero value otherwise.
func (o *TradesAggregateResponseResultInner) GetA() int64 {
	if o == nil || common.IsNil(o.A) {
		var ret int64
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesAggregateResponseResultInner) GetAOk() (*int64, bool) {
	if o == nil || common.IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *TradesAggregateResponseResultInner) HasA() bool {
	if o != nil && !common.IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given int64 and assigns it to the A field.
func (o *TradesAggregateResponseResultInner) SetA(v int64) {
	o.A = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *TradesAggregateResponseResultInner) GetP() string {
	if o == nil || common.IsNil(o.P) {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesAggregateResponseResultInner) GetPOk() (*string, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *TradesAggregateResponseResultInner) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *TradesAggregateResponseResultInner) SetP(v string) {
	o.P = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *TradesAggregateResponseResultInner) GetQ() string {
	if o == nil || common.IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesAggregateResponseResultInner) GetQOk() (*string, bool) {
	if o == nil || common.IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *TradesAggregateResponseResultInner) HasQ() bool {
	if o != nil && !common.IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *TradesAggregateResponseResultInner) SetQ(v string) {
	o.Q = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *TradesAggregateResponseResultInner) GetF() int64 {
	if o == nil || common.IsNil(o.F) {
		var ret int64
		return ret
	}
	return *o.F
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesAggregateResponseResultInner) GetFOk() (*int64, bool) {
	if o == nil || common.IsNil(o.F) {
		return nil, false
	}
	return o.F, true
}

// HasF returns a boolean if a field has been set.
func (o *TradesAggregateResponseResultInner) HasF() bool {
	if o != nil && !common.IsNil(o.F) {
		return true
	}

	return false
}

// SetF gets a reference to the given int64 and assigns it to the F field.
func (o *TradesAggregateResponseResultInner) SetF(v int64) {
	o.F = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *TradesAggregateResponseResultInner) GetL() int64 {
	if o == nil || common.IsNil(o.L) {
		var ret int64
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesAggregateResponseResultInner) GetLOk() (*int64, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *TradesAggregateResponseResultInner) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given int64 and assigns it to the L field.
func (o *TradesAggregateResponseResultInner) SetL(v int64) {
	o.L = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *TradesAggregateResponseResultInner) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesAggregateResponseResultInner) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *TradesAggregateResponseResultInner) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *TradesAggregateResponseResultInner) SetT(v int64) {
	o.T = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *TradesAggregateResponseResultInner) GetSmallm() bool {
	if o == nil || common.IsNil(o.Smallm) {
		var ret bool
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesAggregateResponseResultInner) GetSmallmOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *TradesAggregateResponseResultInner) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *TradesAggregateResponseResultInner) SetSmallm(v bool) {
	o.Smallm = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *TradesAggregateResponseResultInner) GetM() bool {
	if o == nil || common.IsNil(o.M) {
		var ret bool
		return ret
	}
	return *o.M
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesAggregateResponseResultInner) GetMOk() (*bool, bool) {
	if o == nil || common.IsNil(o.M) {
		return nil, false
	}
	return o.M, true
}

// HasM returns a boolean if a field has been set.
func (o *TradesAggregateResponseResultInner) HasM() bool {
	if o != nil && !common.IsNil(o.M) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *TradesAggregateResponseResultInner) SetM(v bool) {
	o.M = &v
}

func (o TradesAggregateResponseResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradesAggregateResponseResultInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Smallm) {
		toSerialize["m"] = o.Smallm
	}
	if !common.IsNil(o.M) {
		toSerialize["M"] = o.M
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TradesAggregateResponseResultInner) UnmarshalJSON(data []byte) (err error) {
	varTradesAggregateResponseResultInner := _TradesAggregateResponseResultInner{}

	err = json.Unmarshal(data, &varTradesAggregateResponseResultInner)

	if err != nil {
		return err
	}

	*o = TradesAggregateResponseResultInner(varTradesAggregateResponseResultInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "a")
		delete(additionalProperties, "p")
		delete(additionalProperties, "q")
		delete(additionalProperties, "f")
		delete(additionalProperties, "l")
		delete(additionalProperties, "T")
		delete(additionalProperties, "m")
		delete(additionalProperties, "M")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradesAggregateResponseResultInner struct {
	value *TradesAggregateResponseResultInner
	isSet bool
}

func (v NullableTradesAggregateResponseResultInner) Get() *TradesAggregateResponseResultInner {
	return v.value
}

func (v *NullableTradesAggregateResponseResultInner) Set(val *TradesAggregateResponseResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTradesAggregateResponseResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTradesAggregateResponseResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradesAggregateResponseResultInner(val *TradesAggregateResponseResultInner) *NullableTradesAggregateResponseResultInner {
	return &NullableTradesAggregateResponseResultInner{value: val, isSet: true}
}

func (v NullableTradesAggregateResponseResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradesAggregateResponseResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
