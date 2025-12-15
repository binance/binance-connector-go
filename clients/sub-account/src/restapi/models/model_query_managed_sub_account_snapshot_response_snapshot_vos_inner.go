/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryManagedSubAccountSnapshotResponseSnapshotVosInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountSnapshotResponseSnapshotVosInner{}

// QueryManagedSubAccountSnapshotResponseSnapshotVosInner struct for QueryManagedSubAccountSnapshotResponseSnapshotVosInner
type QueryManagedSubAccountSnapshotResponseSnapshotVosInner struct {
	Data                 *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData `json:"data,omitempty"`
	Type                 *string                                                     `json:"type,omitempty"`
	UpdateTime           *int64                                                      `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryManagedSubAccountSnapshotResponseSnapshotVosInner QueryManagedSubAccountSnapshotResponseSnapshotVosInner

// NewQueryManagedSubAccountSnapshotResponseSnapshotVosInner instantiates a new QueryManagedSubAccountSnapshotResponseSnapshotVosInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountSnapshotResponseSnapshotVosInner() *QueryManagedSubAccountSnapshotResponseSnapshotVosInner {
	this := QueryManagedSubAccountSnapshotResponseSnapshotVosInner{}
	return &this
}

// NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerWithDefaults instantiates a new QueryManagedSubAccountSnapshotResponseSnapshotVosInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerWithDefaults() *QueryManagedSubAccountSnapshotResponseSnapshotVosInner {
	this := QueryManagedSubAccountSnapshotResponseSnapshotVosInner{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInner) GetData() QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData {
	if o == nil || common.IsNil(o.Data) {
		var ret QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInner) GetDataOk() (*QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInner) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData and assigns it to the Data field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInner) SetData(v QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) {
	o.Data = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInner) SetType(v string) {
	o.Type = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o QueryManagedSubAccountSnapshotResponseSnapshotVosInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountSnapshotResponseSnapshotVosInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInner) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountSnapshotResponseSnapshotVosInner := _QueryManagedSubAccountSnapshotResponseSnapshotVosInner{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountSnapshotResponseSnapshotVosInner)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountSnapshotResponseSnapshotVosInner(varQueryManagedSubAccountSnapshotResponseSnapshotVosInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInner struct {
	value *QueryManagedSubAccountSnapshotResponseSnapshotVosInner
	isSet bool
}

func (v NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInner) Get() *QueryManagedSubAccountSnapshotResponseSnapshotVosInner {
	return v.value
}

func (v *NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInner) Set(val *QueryManagedSubAccountSnapshotResponseSnapshotVosInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountSnapshotResponseSnapshotVosInner(val *QueryManagedSubAccountSnapshotResponseSnapshotVosInner) *NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInner {
	return &NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInner{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
