/*
Binance Derivatives Trading Portfolio Margin WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the Openorderloss type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Openorderloss{}

// Openorderloss struct for Openorderloss
type Openorderloss struct {
	E                    *int64                `json:"E,omitempty"`
	O                    []OpenorderlossOInner `json:"O,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Openorderloss Openorderloss

// NewOpenorderloss instantiates a new Openorderloss object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenorderloss() *Openorderloss {
	this := Openorderloss{}
	return &this
}

// NewOpenorderlossWithDefaults instantiates a new Openorderloss object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenorderlossWithDefaults() *Openorderloss {
	this := Openorderloss{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *Openorderloss) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Openorderloss) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *Openorderloss) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *Openorderloss) SetE(v int64) {
	o.E = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *Openorderloss) GetO() []OpenorderlossOInner {
	if o == nil || common.IsNil(o.O) {
		var ret []OpenorderlossOInner
		return ret
	}
	return o.O
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Openorderloss) GetOOk() ([]OpenorderlossOInner, bool) {
	if o == nil || common.IsNil(o.O) {
		return nil, false
	}
	return o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *Openorderloss) HasO() bool {
	if o != nil && !common.IsNil(o.O) {
		return true
	}

	return false
}

// SetO gets a reference to the given []OpenorderlossOInner and assigns it to the O field.
func (o *Openorderloss) SetO(v []OpenorderlossOInner) {
	o.O = v
}

func (o Openorderloss) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Openorderloss) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.O) {
		toSerialize["O"] = o.O
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Openorderloss) UnmarshalJSON(data []byte) (err error) {
	varOpenorderloss := _Openorderloss{}

	err = json.Unmarshal(data, &varOpenorderloss)

	if err != nil {
		return err
	}

	*o = Openorderloss(varOpenorderloss)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "O")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenorderloss struct {
	value *Openorderloss
	isSet bool
}

func (v NullableOpenorderloss) Get() *Openorderloss {
	return v.value
}

func (v *NullableOpenorderloss) Set(val *Openorderloss) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenorderloss) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenorderloss) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenorderloss(val *Openorderloss) *NullableOpenorderloss {
	return &NullableOpenorderloss{value: val, isSet: true}
}

func (v NullableOpenorderloss) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenorderloss) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
