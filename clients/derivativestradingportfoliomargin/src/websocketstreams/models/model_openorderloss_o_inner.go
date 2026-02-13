/*
Binance Derivatives Trading Portfolio Margin WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OpenorderlossOInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OpenorderlossOInner{}

// OpenorderlossOInner struct for OpenorderlossOInner
type OpenorderlossOInner struct {
	Smalla               *string `json:"a,omitempty"`
	Smallo               *string `json:"o,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenorderlossOInner OpenorderlossOInner

// NewOpenorderlossOInner instantiates a new OpenorderlossOInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenorderlossOInner() *OpenorderlossOInner {
	this := OpenorderlossOInner{}
	return &this
}

// NewOpenorderlossOInnerWithDefaults instantiates a new OpenorderlossOInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenorderlossOInnerWithDefaults() *OpenorderlossOInner {
	this := OpenorderlossOInner{}
	return &this
}

// GetA returns the A field value if set, zero value otherwise.
func (o *OpenorderlossOInner) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenorderlossOInner) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *OpenorderlossOInner) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *OpenorderlossOInner) SetSmalla(v string) {
	o.Smalla = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *OpenorderlossOInner) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenorderlossOInner) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *OpenorderlossOInner) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *OpenorderlossOInner) SetSmallo(v string) {
	o.Smallo = &v
}

func (o OpenorderlossOInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenorderlossOInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenorderlossOInner) UnmarshalJSON(data []byte) (err error) {
	varOpenorderlossOInner := _OpenorderlossOInner{}

	err = json.Unmarshal(data, &varOpenorderlossOInner)

	if err != nil {
		return err
	}

	*o = OpenorderlossOInner(varOpenorderlossOInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "a")
		delete(additionalProperties, "o")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenorderlossOInner struct {
	value *OpenorderlossOInner
	isSet bool
}

func (v NullableOpenorderlossOInner) Get() *OpenorderlossOInner {
	return v.value
}

func (v *NullableOpenorderlossOInner) Set(val *OpenorderlossOInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenorderlossOInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenorderlossOInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenorderlossOInner(val *OpenorderlossOInner) *NullableOpenorderlossOInner {
	return &NullableOpenorderlossOInner{value: val, isSet: true}
}

func (v NullableOpenorderlossOInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenorderlossOInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
