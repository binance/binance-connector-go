/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the RequestForDetailMinerListResponseDataInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RequestForDetailMinerListResponseDataInner{}

// RequestForDetailMinerListResponseDataInner struct for RequestForDetailMinerListResponseDataInner
type RequestForDetailMinerListResponseDataInner struct {
	WorkerName           *string                                                        `json:"workerName,omitempty"`
	Type                 *string                                                        `json:"type,omitempty"`
	HashrateDatas        []RequestForDetailMinerListResponseDataInnerHashrateDatasInner `json:"hashrateDatas,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestForDetailMinerListResponseDataInner RequestForDetailMinerListResponseDataInner

// NewRequestForDetailMinerListResponseDataInner instantiates a new RequestForDetailMinerListResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestForDetailMinerListResponseDataInner() *RequestForDetailMinerListResponseDataInner {
	this := RequestForDetailMinerListResponseDataInner{}
	return &this
}

// NewRequestForDetailMinerListResponseDataInnerWithDefaults instantiates a new RequestForDetailMinerListResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestForDetailMinerListResponseDataInnerWithDefaults() *RequestForDetailMinerListResponseDataInner {
	this := RequestForDetailMinerListResponseDataInner{}
	return &this
}

// GetWorkerName returns the WorkerName field value if set, zero value otherwise.
func (o *RequestForDetailMinerListResponseDataInner) GetWorkerName() string {
	if o == nil || common.IsNil(o.WorkerName) {
		var ret string
		return ret
	}
	return *o.WorkerName
}

// GetWorkerNameOk returns a tuple with the WorkerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForDetailMinerListResponseDataInner) GetWorkerNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.WorkerName) {
		return nil, false
	}
	return o.WorkerName, true
}

// HasWorkerName returns a boolean if a field has been set.
func (o *RequestForDetailMinerListResponseDataInner) HasWorkerName() bool {
	if o != nil && !common.IsNil(o.WorkerName) {
		return true
	}

	return false
}

// SetWorkerName gets a reference to the given string and assigns it to the WorkerName field.
func (o *RequestForDetailMinerListResponseDataInner) SetWorkerName(v string) {
	o.WorkerName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RequestForDetailMinerListResponseDataInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForDetailMinerListResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RequestForDetailMinerListResponseDataInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RequestForDetailMinerListResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetHashrateDatas returns the HashrateDatas field value if set, zero value otherwise.
func (o *RequestForDetailMinerListResponseDataInner) GetHashrateDatas() []RequestForDetailMinerListResponseDataInnerHashrateDatasInner {
	if o == nil || common.IsNil(o.HashrateDatas) {
		var ret []RequestForDetailMinerListResponseDataInnerHashrateDatasInner
		return ret
	}
	return o.HashrateDatas
}

// GetHashrateDatasOk returns a tuple with the HashrateDatas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForDetailMinerListResponseDataInner) GetHashrateDatasOk() ([]RequestForDetailMinerListResponseDataInnerHashrateDatasInner, bool) {
	if o == nil || common.IsNil(o.HashrateDatas) {
		return nil, false
	}
	return o.HashrateDatas, true
}

// HasHashrateDatas returns a boolean if a field has been set.
func (o *RequestForDetailMinerListResponseDataInner) HasHashrateDatas() bool {
	if o != nil && !common.IsNil(o.HashrateDatas) {
		return true
	}

	return false
}

// SetHashrateDatas gets a reference to the given []RequestForDetailMinerListResponseDataInnerHashrateDatasInner and assigns it to the HashrateDatas field.
func (o *RequestForDetailMinerListResponseDataInner) SetHashrateDatas(v []RequestForDetailMinerListResponseDataInnerHashrateDatasInner) {
	o.HashrateDatas = v
}

func (o RequestForDetailMinerListResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestForDetailMinerListResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.WorkerName) {
		toSerialize["workerName"] = o.WorkerName
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.HashrateDatas) {
		toSerialize["hashrateDatas"] = o.HashrateDatas
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestForDetailMinerListResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	varRequestForDetailMinerListResponseDataInner := _RequestForDetailMinerListResponseDataInner{}

	err = json.Unmarshal(data, &varRequestForDetailMinerListResponseDataInner)

	if err != nil {
		return err
	}

	*o = RequestForDetailMinerListResponseDataInner(varRequestForDetailMinerListResponseDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "workerName")
		delete(additionalProperties, "type")
		delete(additionalProperties, "hashrateDatas")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestForDetailMinerListResponseDataInner struct {
	value *RequestForDetailMinerListResponseDataInner
	isSet bool
}

func (v NullableRequestForDetailMinerListResponseDataInner) Get() *RequestForDetailMinerListResponseDataInner {
	return v.value
}

func (v *NullableRequestForDetailMinerListResponseDataInner) Set(val *RequestForDetailMinerListResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestForDetailMinerListResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestForDetailMinerListResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestForDetailMinerListResponseDataInner(val *RequestForDetailMinerListResponseDataInner) *NullableRequestForDetailMinerListResponseDataInner {
	return &NullableRequestForDetailMinerListResponseDataInner{value: val, isSet: true}
}

func (v NullableRequestForDetailMinerListResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestForDetailMinerListResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
