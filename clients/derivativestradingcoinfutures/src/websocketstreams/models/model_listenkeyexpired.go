/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the Listenkeyexpired type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Listenkeyexpired{}

// Listenkeyexpired struct for Listenkeyexpired
type Listenkeyexpired struct {
	E                    *int64  `json:"E,omitempty"`
	ListenKey            *string `json:"listenKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Listenkeyexpired Listenkeyexpired

// NewListenkeyexpired instantiates a new Listenkeyexpired object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListenkeyexpired() *Listenkeyexpired {
	this := Listenkeyexpired{}
	return &this
}

// NewListenkeyexpiredWithDefaults instantiates a new Listenkeyexpired object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListenkeyexpiredWithDefaults() *Listenkeyexpired {
	this := Listenkeyexpired{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *Listenkeyexpired) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listenkeyexpired) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *Listenkeyexpired) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *Listenkeyexpired) SetE(v int64) {
	o.E = &v
}

// GetListenKey returns the ListenKey field value if set, zero value otherwise.
func (o *Listenkeyexpired) GetListenKey() string {
	if o == nil || common.IsNil(o.ListenKey) {
		var ret string
		return ret
	}
	return *o.ListenKey
}

// GetListenKeyOk returns a tuple with the ListenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listenkeyexpired) GetListenKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListenKey) {
		return nil, false
	}
	return o.ListenKey, true
}

// HasListenKey returns a boolean if a field has been set.
func (o *Listenkeyexpired) HasListenKey() bool {
	if o != nil && !common.IsNil(o.ListenKey) {
		return true
	}

	return false
}

// SetListenKey gets a reference to the given string and assigns it to the ListenKey field.
func (o *Listenkeyexpired) SetListenKey(v string) {
	o.ListenKey = &v
}

func (o Listenkeyexpired) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Listenkeyexpired) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.ListenKey) {
		toSerialize["listenKey"] = o.ListenKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Listenkeyexpired) UnmarshalJSON(data []byte) (err error) {
	varListenkeyexpired := _Listenkeyexpired{}

	err = json.Unmarshal(data, &varListenkeyexpired)

	if err != nil {
		return err
	}

	*o = Listenkeyexpired(varListenkeyexpired)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "listenKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListenkeyexpired struct {
	value *Listenkeyexpired
	isSet bool
}

func (v NullableListenkeyexpired) Get() *Listenkeyexpired {
	return v.value
}

func (v *NullableListenkeyexpired) Set(val *Listenkeyexpired) {
	v.value = val
	v.isSet = true
}

func (v NullableListenkeyexpired) IsSet() bool {
	return v.isSet
}

func (v *NullableListenkeyexpired) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListenkeyexpired(val *Listenkeyexpired) *NullableListenkeyexpired {
	return &NullableListenkeyexpired{value: val, isSet: true}
}

func (v NullableListenkeyexpired) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListenkeyexpired) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
