/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DailyAccountSnapshotResponseSnapshotVosInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DailyAccountSnapshotResponseSnapshotVosInner{}

// DailyAccountSnapshotResponseSnapshotVosInner struct for DailyAccountSnapshotResponseSnapshotVosInner
type DailyAccountSnapshotResponseSnapshotVosInner struct {
	Data                 *DailyAccountSnapshotResponseSnapshotVosInnerData `json:"data,omitempty"`
	Type                 *string                                           `json:"type,omitempty"`
	UpdateTime           *int64                                            `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DailyAccountSnapshotResponseSnapshotVosInner DailyAccountSnapshotResponseSnapshotVosInner

// NewDailyAccountSnapshotResponseSnapshotVosInner instantiates a new DailyAccountSnapshotResponseSnapshotVosInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyAccountSnapshotResponseSnapshotVosInner() *DailyAccountSnapshotResponseSnapshotVosInner {
	this := DailyAccountSnapshotResponseSnapshotVosInner{}
	return &this
}

// NewDailyAccountSnapshotResponseSnapshotVosInnerWithDefaults instantiates a new DailyAccountSnapshotResponseSnapshotVosInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyAccountSnapshotResponseSnapshotVosInnerWithDefaults() *DailyAccountSnapshotResponseSnapshotVosInner {
	this := DailyAccountSnapshotResponseSnapshotVosInner{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInner) GetData() DailyAccountSnapshotResponseSnapshotVosInnerData {
	if o == nil || common.IsNil(o.Data) {
		var ret DailyAccountSnapshotResponseSnapshotVosInnerData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInner) GetDataOk() (*DailyAccountSnapshotResponseSnapshotVosInnerData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInner) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DailyAccountSnapshotResponseSnapshotVosInnerData and assigns it to the Data field.
func (o *DailyAccountSnapshotResponseSnapshotVosInner) SetData(v DailyAccountSnapshotResponseSnapshotVosInnerData) {
	o.Data = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DailyAccountSnapshotResponseSnapshotVosInner) SetType(v string) {
	o.Type = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *DailyAccountSnapshotResponseSnapshotVosInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o DailyAccountSnapshotResponseSnapshotVosInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DailyAccountSnapshotResponseSnapshotVosInner) ToMap() (map[string]interface{}, error) {
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

func (o *DailyAccountSnapshotResponseSnapshotVosInner) UnmarshalJSON(data []byte) (err error) {
	varDailyAccountSnapshotResponseSnapshotVosInner := _DailyAccountSnapshotResponseSnapshotVosInner{}

	err = json.Unmarshal(data, &varDailyAccountSnapshotResponseSnapshotVosInner)

	if err != nil {
		return err
	}

	*o = DailyAccountSnapshotResponseSnapshotVosInner(varDailyAccountSnapshotResponseSnapshotVosInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDailyAccountSnapshotResponseSnapshotVosInner struct {
	value *DailyAccountSnapshotResponseSnapshotVosInner
	isSet bool
}

func (v NullableDailyAccountSnapshotResponseSnapshotVosInner) Get() *DailyAccountSnapshotResponseSnapshotVosInner {
	return v.value
}

func (v *NullableDailyAccountSnapshotResponseSnapshotVosInner) Set(val *DailyAccountSnapshotResponseSnapshotVosInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyAccountSnapshotResponseSnapshotVosInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyAccountSnapshotResponseSnapshotVosInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyAccountSnapshotResponseSnapshotVosInner(val *DailyAccountSnapshotResponseSnapshotVosInner) *NullableDailyAccountSnapshotResponseSnapshotVosInner {
	return &NullableDailyAccountSnapshotResponseSnapshotVosInner{value: val, isSet: true}
}

func (v NullableDailyAccountSnapshotResponseSnapshotVosInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyAccountSnapshotResponseSnapshotVosInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
