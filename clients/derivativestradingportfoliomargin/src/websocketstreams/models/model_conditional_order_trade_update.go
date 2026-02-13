/*
Binance Derivatives Trading Portfolio Margin WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ConditionalOrderTradeUpdate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ConditionalOrderTradeUpdate{}

// ConditionalOrderTradeUpdate struct for ConditionalOrderTradeUpdate
type ConditionalOrderTradeUpdate struct {
	T                    *int64                         `json:"T,omitempty"`
	E                    *int64                         `json:"E,omitempty"`
	Smallfs              *string                        `json:"fs,omitempty"`
	Smallso              *ConditionalOrderTradeUpdateSo `json:"so,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConditionalOrderTradeUpdate ConditionalOrderTradeUpdate

// NewConditionalOrderTradeUpdate instantiates a new ConditionalOrderTradeUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionalOrderTradeUpdate() *ConditionalOrderTradeUpdate {
	this := ConditionalOrderTradeUpdate{}
	return &this
}

// NewConditionalOrderTradeUpdateWithDefaults instantiates a new ConditionalOrderTradeUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionalOrderTradeUpdateWithDefaults() *ConditionalOrderTradeUpdate {
	this := ConditionalOrderTradeUpdate{}
	return &this
}

// GetT returns the T field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdate) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdate) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdate) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *ConditionalOrderTradeUpdate) SetT(v int64) {
	o.T = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdate) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdate) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdate) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *ConditionalOrderTradeUpdate) SetE(v int64) {
	o.E = &v
}

// GetFs returns the Fs field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdate) GetSmallfs() string {
	if o == nil || common.IsNil(o.Smallfs) {
		var ret string
		return ret
	}
	return *o.Smallfs
}

// GetFsOk returns a tuple with the Fs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdate) GetSmallfsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallfs) {
		return nil, false
	}
	return o.Smallfs, true
}

// HasFs returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdate) HasSmallfs() bool {
	if o != nil && !common.IsNil(o.Smallfs) {
		return true
	}

	return false
}

// SetFs gets a reference to the given string and assigns it to the Fs field.
func (o *ConditionalOrderTradeUpdate) SetSmallfs(v string) {
	o.Smallfs = &v
}

// GetSo returns the So field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdate) GetSmallso() ConditionalOrderTradeUpdateSo {
	if o == nil || common.IsNil(o.Smallso) {
		var ret ConditionalOrderTradeUpdateSo
		return ret
	}
	return *o.Smallso
}

// GetSoOk returns a tuple with the So field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdate) GetSmallsoOk() (*ConditionalOrderTradeUpdateSo, bool) {
	if o == nil || common.IsNil(o.Smallso) {
		return nil, false
	}
	return o.Smallso, true
}

// HasSo returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdate) HasSmallso() bool {
	if o != nil && !common.IsNil(o.Smallso) {
		return true
	}

	return false
}

// SetSo gets a reference to the given ConditionalOrderTradeUpdateSo and assigns it to the So field.
func (o *ConditionalOrderTradeUpdate) SetSmallso(v ConditionalOrderTradeUpdateSo) {
	o.Smallso = &v
}

func (o ConditionalOrderTradeUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionalOrderTradeUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smallfs) {
		toSerialize["fs"] = o.Smallfs
	}
	if !common.IsNil(o.Smallso) {
		toSerialize["so"] = o.Smallso
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConditionalOrderTradeUpdate) UnmarshalJSON(data []byte) (err error) {
	varConditionalOrderTradeUpdate := _ConditionalOrderTradeUpdate{}

	err = json.Unmarshal(data, &varConditionalOrderTradeUpdate)

	if err != nil {
		return err
	}

	*o = ConditionalOrderTradeUpdate(varConditionalOrderTradeUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "T")
		delete(additionalProperties, "E")
		delete(additionalProperties, "fs")
		delete(additionalProperties, "so")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConditionalOrderTradeUpdate struct {
	value *ConditionalOrderTradeUpdate
	isSet bool
}

func (v NullableConditionalOrderTradeUpdate) Get() *ConditionalOrderTradeUpdate {
	return v.value
}

func (v *NullableConditionalOrderTradeUpdate) Set(val *ConditionalOrderTradeUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionalOrderTradeUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionalOrderTradeUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionalOrderTradeUpdate(val *ConditionalOrderTradeUpdate) *NullableConditionalOrderTradeUpdate {
	return &NullableConditionalOrderTradeUpdate{value: val, isSet: true}
}

func (v NullableConditionalOrderTradeUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionalOrderTradeUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
