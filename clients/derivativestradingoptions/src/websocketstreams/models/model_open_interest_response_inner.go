/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OpenInterestResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OpenInterestResponseInner{}

// OpenInterestResponseInner struct for OpenInterestResponseInner
type OpenInterestResponseInner struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smallo               *string `json:"o,omitempty"`
	Smallh               *string `json:"h,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenInterestResponseInner OpenInterestResponseInner

// NewOpenInterestResponseInner instantiates a new OpenInterestResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenInterestResponseInner() *OpenInterestResponseInner {
	this := OpenInterestResponseInner{}
	return &this
}

// NewOpenInterestResponseInnerWithDefaults instantiates a new OpenInterestResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenInterestResponseInnerWithDefaults() *OpenInterestResponseInner {
	this := OpenInterestResponseInner{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OpenInterestResponseInner) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestResponseInner) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *OpenInterestResponseInner) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *OpenInterestResponseInner) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OpenInterestResponseInner) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestResponseInner) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *OpenInterestResponseInner) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *OpenInterestResponseInner) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *OpenInterestResponseInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestResponseInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *OpenInterestResponseInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *OpenInterestResponseInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *OpenInterestResponseInner) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestResponseInner) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *OpenInterestResponseInner) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *OpenInterestResponseInner) SetSmallo(v string) {
	o.Smallo = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *OpenInterestResponseInner) GetSmallh() string {
	if o == nil || common.IsNil(o.Smallh) {
		var ret string
		return ret
	}
	return *o.Smallh
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestResponseInner) GetSmallhOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallh) {
		return nil, false
	}
	return o.Smallh, true
}

// HasH returns a boolean if a field has been set.
func (o *OpenInterestResponseInner) HasSmallh() bool {
	if o != nil && !common.IsNil(o.Smallh) {
		return true
	}

	return false
}

// SetH gets a reference to the given string and assigns it to the H field.
func (o *OpenInterestResponseInner) SetSmallh(v string) {
	o.Smallh = &v
}

func (o OpenInterestResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenInterestResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}
	if !common.IsNil(o.Smallh) {
		toSerialize["h"] = o.Smallh
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenInterestResponseInner) UnmarshalJSON(data []byte) (err error) {
	varOpenInterestResponseInner := _OpenInterestResponseInner{}

	err = json.Unmarshal(data, &varOpenInterestResponseInner)

	if err != nil {
		return err
	}

	*o = OpenInterestResponseInner(varOpenInterestResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "o")
		delete(additionalProperties, "h")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenInterestResponseInner struct {
	value *OpenInterestResponseInner
	isSet bool
}

func (v NullableOpenInterestResponseInner) Get() *OpenInterestResponseInner {
	return v.value
}

func (v *NullableOpenInterestResponseInner) Set(val *OpenInterestResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenInterestResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenInterestResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenInterestResponseInner(val *OpenInterestResponseInner) *NullableOpenInterestResponseInner {
	return &NullableOpenInterestResponseInner{value: val, isSet: true}
}

func (v NullableOpenInterestResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenInterestResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
