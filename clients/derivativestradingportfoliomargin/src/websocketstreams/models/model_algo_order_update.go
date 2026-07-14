/*
Portfolio Margin WebSocket Market Streams

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AlgoOrderUpdate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AlgoOrderUpdate{}

// AlgoOrderUpdate struct for AlgoOrderUpdate
type AlgoOrderUpdate struct {
	// Transaction Time
	T *int64 `json:"T,omitempty"`
	// Event Time
	E *int64 `json:"E,omitempty"`
	// Futures segment, `UM` for USDS-M Futures, `CM` for Coin-M Futures
	Smallfs              *string            `json:"fs,omitempty"`
	Smallao              *AlgoOrderUpdateAo `json:"ao,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AlgoOrderUpdate AlgoOrderUpdate

// NewAlgoOrderUpdate instantiates a new AlgoOrderUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlgoOrderUpdate() *AlgoOrderUpdate {
	this := AlgoOrderUpdate{}
	return &this
}

// NewAlgoOrderUpdateWithDefaults instantiates a new AlgoOrderUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlgoOrderUpdateWithDefaults() *AlgoOrderUpdate {
	this := AlgoOrderUpdate{}
	return &this
}

// GetT returns the T field value if set, zero value otherwise.
func (o *AlgoOrderUpdate) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdate) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *AlgoOrderUpdate) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *AlgoOrderUpdate) SetT(v int64) {
	o.T = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AlgoOrderUpdate) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdate) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AlgoOrderUpdate) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AlgoOrderUpdate) SetE(v int64) {
	o.E = &v
}

// GetFs returns the Fs field value if set, zero value otherwise.
func (o *AlgoOrderUpdate) GetSmallfs() string {
	if o == nil || common.IsNil(o.Smallfs) {
		var ret string
		return ret
	}
	return *o.Smallfs
}

// GetFsOk returns a tuple with the Fs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdate) GetSmallfsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallfs) {
		return nil, false
	}
	return o.Smallfs, true
}

// HasFs returns a boolean if a field has been set.
func (o *AlgoOrderUpdate) HasSmallfs() bool {
	if o != nil && !common.IsNil(o.Smallfs) {
		return true
	}

	return false
}

// SetFs gets a reference to the given string and assigns it to the Fs field.
func (o *AlgoOrderUpdate) SetSmallfs(v string) {
	o.Smallfs = &v
}

// GetAo returns the Ao field value if set, zero value otherwise.
func (o *AlgoOrderUpdate) GetSmallao() AlgoOrderUpdateAo {
	if o == nil || common.IsNil(o.Smallao) {
		var ret AlgoOrderUpdateAo
		return ret
	}
	return *o.Smallao
}

// GetAoOk returns a tuple with the Ao field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdate) GetSmallaoOk() (*AlgoOrderUpdateAo, bool) {
	if o == nil || common.IsNil(o.Smallao) {
		return nil, false
	}
	return o.Smallao, true
}

// HasAo returns a boolean if a field has been set.
func (o *AlgoOrderUpdate) HasSmallao() bool {
	if o != nil && !common.IsNil(o.Smallao) {
		return true
	}

	return false
}

// SetAo gets a reference to the given AlgoOrderUpdateAo and assigns it to the Ao field.
func (o *AlgoOrderUpdate) SetSmallao(v AlgoOrderUpdateAo) {
	o.Smallao = &v
}

func (o AlgoOrderUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlgoOrderUpdate) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Smallao) {
		toSerialize["ao"] = o.Smallao
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AlgoOrderUpdate) UnmarshalJSON(data []byte) (err error) {
	varAlgoOrderUpdate := _AlgoOrderUpdate{}

	err = json.Unmarshal(data, &varAlgoOrderUpdate)

	if err != nil {
		return err
	}

	*o = AlgoOrderUpdate(varAlgoOrderUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "T")
		delete(additionalProperties, "E")
		delete(additionalProperties, "fs")
		delete(additionalProperties, "ao")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAlgoOrderUpdate struct {
	value *AlgoOrderUpdate
	isSet bool
}

func (v NullableAlgoOrderUpdate) Get() *AlgoOrderUpdate {
	return v.value
}

func (v *NullableAlgoOrderUpdate) Set(val *AlgoOrderUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAlgoOrderUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAlgoOrderUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlgoOrderUpdate(val *AlgoOrderUpdate) *NullableAlgoOrderUpdate {
	return &NullableAlgoOrderUpdate{value: val, isSet: true}
}

func (v NullableAlgoOrderUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlgoOrderUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
