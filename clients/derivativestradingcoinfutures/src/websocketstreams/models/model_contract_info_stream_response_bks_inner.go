/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ContractInfoStreamResponseBksInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ContractInfoStreamResponseBksInner{}

// ContractInfoStreamResponseBksInner struct for ContractInfoStreamResponseBksInner
type ContractInfoStreamResponseBksInner struct {
	Smallbs              *int64   `json:"bs,omitempty"`
	Bnf                  *int64   `json:"bnf,omitempty"`
	Bnc                  *int64   `json:"bnc,omitempty"`
	Mmr                  *float32 `json:"mmr,omitempty"`
	Smallcf              *int64   `json:"cf,omitempty"`
	Smallmi              *int64   `json:"mi,omitempty"`
	Smallma              *int64   `json:"ma,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContractInfoStreamResponseBksInner ContractInfoStreamResponseBksInner

// NewContractInfoStreamResponseBksInner instantiates a new ContractInfoStreamResponseBksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractInfoStreamResponseBksInner() *ContractInfoStreamResponseBksInner {
	this := ContractInfoStreamResponseBksInner{}
	return &this
}

// NewContractInfoStreamResponseBksInnerWithDefaults instantiates a new ContractInfoStreamResponseBksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractInfoStreamResponseBksInnerWithDefaults() *ContractInfoStreamResponseBksInner {
	this := ContractInfoStreamResponseBksInner{}
	return &this
}

// GetBs returns the Bs field value if set, zero value otherwise.
func (o *ContractInfoStreamResponseBksInner) GetSmallbs() int64 {
	if o == nil || common.IsNil(o.Smallbs) {
		var ret int64
		return ret
	}
	return *o.Smallbs
}

// GetBsOk returns a tuple with the Bs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInfoStreamResponseBksInner) GetSmallbsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallbs) {
		return nil, false
	}
	return o.Smallbs, true
}

// HasBs returns a boolean if a field has been set.
func (o *ContractInfoStreamResponseBksInner) HasSmallbs() bool {
	if o != nil && !common.IsNil(o.Smallbs) {
		return true
	}

	return false
}

// SetBs gets a reference to the given int64 and assigns it to the Bs field.
func (o *ContractInfoStreamResponseBksInner) SetSmallbs(v int64) {
	o.Smallbs = &v
}

// GetBnf returns the Bnf field value if set, zero value otherwise.
func (o *ContractInfoStreamResponseBksInner) GetBnf() int64 {
	if o == nil || common.IsNil(o.Bnf) {
		var ret int64
		return ret
	}
	return *o.Bnf
}

// GetBnfOk returns a tuple with the Bnf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInfoStreamResponseBksInner) GetBnfOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Bnf) {
		return nil, false
	}
	return o.Bnf, true
}

// HasBnf returns a boolean if a field has been set.
func (o *ContractInfoStreamResponseBksInner) HasBnf() bool {
	if o != nil && !common.IsNil(o.Bnf) {
		return true
	}

	return false
}

// SetBnf gets a reference to the given int64 and assigns it to the Bnf field.
func (o *ContractInfoStreamResponseBksInner) SetBnf(v int64) {
	o.Bnf = &v
}

// GetBnc returns the Bnc field value if set, zero value otherwise.
func (o *ContractInfoStreamResponseBksInner) GetBnc() int64 {
	if o == nil || common.IsNil(o.Bnc) {
		var ret int64
		return ret
	}
	return *o.Bnc
}

// GetBncOk returns a tuple with the Bnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInfoStreamResponseBksInner) GetBncOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Bnc) {
		return nil, false
	}
	return o.Bnc, true
}

// HasBnc returns a boolean if a field has been set.
func (o *ContractInfoStreamResponseBksInner) HasBnc() bool {
	if o != nil && !common.IsNil(o.Bnc) {
		return true
	}

	return false
}

// SetBnc gets a reference to the given int64 and assigns it to the Bnc field.
func (o *ContractInfoStreamResponseBksInner) SetBnc(v int64) {
	o.Bnc = &v
}

// GetMmr returns the Mmr field value if set, zero value otherwise.
func (o *ContractInfoStreamResponseBksInner) GetMmr() float32 {
	if o == nil || common.IsNil(o.Mmr) {
		var ret float32
		return ret
	}
	return *o.Mmr
}

// GetMmrOk returns a tuple with the Mmr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInfoStreamResponseBksInner) GetMmrOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Mmr) {
		return nil, false
	}
	return o.Mmr, true
}

// HasMmr returns a boolean if a field has been set.
func (o *ContractInfoStreamResponseBksInner) HasMmr() bool {
	if o != nil && !common.IsNil(o.Mmr) {
		return true
	}

	return false
}

// SetMmr gets a reference to the given float32 and assigns it to the Mmr field.
func (o *ContractInfoStreamResponseBksInner) SetMmr(v float32) {
	o.Mmr = &v
}

// GetCf returns the Cf field value if set, zero value otherwise.
func (o *ContractInfoStreamResponseBksInner) GetSmallcf() int64 {
	if o == nil || common.IsNil(o.Smallcf) {
		var ret int64
		return ret
	}
	return *o.Smallcf
}

// GetCfOk returns a tuple with the Cf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInfoStreamResponseBksInner) GetSmallcfOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallcf) {
		return nil, false
	}
	return o.Smallcf, true
}

// HasCf returns a boolean if a field has been set.
func (o *ContractInfoStreamResponseBksInner) HasSmallcf() bool {
	if o != nil && !common.IsNil(o.Smallcf) {
		return true
	}

	return false
}

// SetCf gets a reference to the given int64 and assigns it to the Cf field.
func (o *ContractInfoStreamResponseBksInner) SetSmallcf(v int64) {
	o.Smallcf = &v
}

// GetMi returns the Mi field value if set, zero value otherwise.
func (o *ContractInfoStreamResponseBksInner) GetSmallmi() int64 {
	if o == nil || common.IsNil(o.Smallmi) {
		var ret int64
		return ret
	}
	return *o.Smallmi
}

// GetMiOk returns a tuple with the Mi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInfoStreamResponseBksInner) GetSmallmiOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallmi) {
		return nil, false
	}
	return o.Smallmi, true
}

// HasMi returns a boolean if a field has been set.
func (o *ContractInfoStreamResponseBksInner) HasSmallmi() bool {
	if o != nil && !common.IsNil(o.Smallmi) {
		return true
	}

	return false
}

// SetMi gets a reference to the given int64 and assigns it to the Mi field.
func (o *ContractInfoStreamResponseBksInner) SetSmallmi(v int64) {
	o.Smallmi = &v
}

// GetMa returns the Ma field value if set, zero value otherwise.
func (o *ContractInfoStreamResponseBksInner) GetSmallma() int64 {
	if o == nil || common.IsNil(o.Smallma) {
		var ret int64
		return ret
	}
	return *o.Smallma
}

// GetMaOk returns a tuple with the Ma field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInfoStreamResponseBksInner) GetSmallmaOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallma) {
		return nil, false
	}
	return o.Smallma, true
}

// HasMa returns a boolean if a field has been set.
func (o *ContractInfoStreamResponseBksInner) HasSmallma() bool {
	if o != nil && !common.IsNil(o.Smallma) {
		return true
	}

	return false
}

// SetMa gets a reference to the given int64 and assigns it to the Ma field.
func (o *ContractInfoStreamResponseBksInner) SetSmallma(v int64) {
	o.Smallma = &v
}

func (o ContractInfoStreamResponseBksInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractInfoStreamResponseBksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smallbs) {
		toSerialize["bs"] = o.Smallbs
	}
	if !common.IsNil(o.Bnf) {
		toSerialize["bnf"] = o.Bnf
	}
	if !common.IsNil(o.Bnc) {
		toSerialize["bnc"] = o.Bnc
	}
	if !common.IsNil(o.Mmr) {
		toSerialize["mmr"] = o.Mmr
	}
	if !common.IsNil(o.Smallcf) {
		toSerialize["cf"] = o.Smallcf
	}
	if !common.IsNil(o.Smallmi) {
		toSerialize["mi"] = o.Smallmi
	}
	if !common.IsNil(o.Smallma) {
		toSerialize["ma"] = o.Smallma
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContractInfoStreamResponseBksInner) UnmarshalJSON(data []byte) (err error) {
	varContractInfoStreamResponseBksInner := _ContractInfoStreamResponseBksInner{}

	err = json.Unmarshal(data, &varContractInfoStreamResponseBksInner)

	if err != nil {
		return err
	}

	*o = ContractInfoStreamResponseBksInner(varContractInfoStreamResponseBksInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bs")
		delete(additionalProperties, "bnf")
		delete(additionalProperties, "bnc")
		delete(additionalProperties, "mmr")
		delete(additionalProperties, "cf")
		delete(additionalProperties, "mi")
		delete(additionalProperties, "ma")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContractInfoStreamResponseBksInner struct {
	value *ContractInfoStreamResponseBksInner
	isSet bool
}

func (v NullableContractInfoStreamResponseBksInner) Get() *ContractInfoStreamResponseBksInner {
	return v.value
}

func (v *NullableContractInfoStreamResponseBksInner) Set(val *ContractInfoStreamResponseBksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContractInfoStreamResponseBksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContractInfoStreamResponseBksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractInfoStreamResponseBksInner(val *ContractInfoStreamResponseBksInner) *NullableContractInfoStreamResponseBksInner {
	return &NullableContractInfoStreamResponseBksInner{value: val, isSet: true}
}

func (v NullableContractInfoStreamResponseBksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractInfoStreamResponseBksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
