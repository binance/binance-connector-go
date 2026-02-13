/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner{}

// QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner struct for QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner
type QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner struct {
	Type                 *string                                                                `json:"type,omitempty"`
	UpdateTime           *int64                                                                 `json:"updateTime,omitempty"`
	Data                 *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner

// NewQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner instantiates a new QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner() *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner {
	this := QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner{}
	return &this
}

// NewQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerWithDefaults instantiates a new QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerWithDefaults() *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner {
	this := QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) SetType(v string) {
	o.Type = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) GetData() QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData {
	if o == nil || common.IsNil(o.Data) {
		var ret QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) GetDataOk() (*QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData and assigns it to the Data field.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) SetData(v QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) {
	o.Data = &v
}

func (o QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner := _QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner(varQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner struct {
	value *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner
	isSet bool
}

func (v NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) Get() *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner {
	return v.value
}

func (v *NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) Set(val *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner(val *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) *NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner {
	return &NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
