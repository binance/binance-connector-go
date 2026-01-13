/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the HashrateResaleListResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &HashrateResaleListResponseData{}

// HashrateResaleListResponseData struct for HashrateResaleListResponseData
type HashrateResaleListResponseData struct {
	ConfigDetails        []HashrateResaleListResponseDataConfigDetailsInner `json:"configDetails,omitempty"`
	TotalNum             *int64                                             `json:"totalNum,omitempty"`
	PageSize             *int64                                             `json:"pageSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HashrateResaleListResponseData HashrateResaleListResponseData

// NewHashrateResaleListResponseData instantiates a new HashrateResaleListResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHashrateResaleListResponseData() *HashrateResaleListResponseData {
	this := HashrateResaleListResponseData{}
	return &this
}

// NewHashrateResaleListResponseDataWithDefaults instantiates a new HashrateResaleListResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHashrateResaleListResponseDataWithDefaults() *HashrateResaleListResponseData {
	this := HashrateResaleListResponseData{}
	return &this
}

// GetConfigDetails returns the ConfigDetails field value if set, zero value otherwise.
func (o *HashrateResaleListResponseData) GetConfigDetails() []HashrateResaleListResponseDataConfigDetailsInner {
	if o == nil || common.IsNil(o.ConfigDetails) {
		var ret []HashrateResaleListResponseDataConfigDetailsInner
		return ret
	}
	return o.ConfigDetails
}

// GetConfigDetailsOk returns a tuple with the ConfigDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleListResponseData) GetConfigDetailsOk() ([]HashrateResaleListResponseDataConfigDetailsInner, bool) {
	if o == nil || common.IsNil(o.ConfigDetails) {
		return nil, false
	}
	return o.ConfigDetails, true
}

// HasConfigDetails returns a boolean if a field has been set.
func (o *HashrateResaleListResponseData) HasConfigDetails() bool {
	if o != nil && !common.IsNil(o.ConfigDetails) {
		return true
	}

	return false
}

// SetConfigDetails gets a reference to the given []HashrateResaleListResponseDataConfigDetailsInner and assigns it to the ConfigDetails field.
func (o *HashrateResaleListResponseData) SetConfigDetails(v []HashrateResaleListResponseDataConfigDetailsInner) {
	o.ConfigDetails = v
}

// GetTotalNum returns the TotalNum field value if set, zero value otherwise.
func (o *HashrateResaleListResponseData) GetTotalNum() int64 {
	if o == nil || common.IsNil(o.TotalNum) {
		var ret int64
		return ret
	}
	return *o.TotalNum
}

// GetTotalNumOk returns a tuple with the TotalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleListResponseData) GetTotalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TotalNum) {
		return nil, false
	}
	return o.TotalNum, true
}

// HasTotalNum returns a boolean if a field has been set.
func (o *HashrateResaleListResponseData) HasTotalNum() bool {
	if o != nil && !common.IsNil(o.TotalNum) {
		return true
	}

	return false
}

// SetTotalNum gets a reference to the given int64 and assigns it to the TotalNum field.
func (o *HashrateResaleListResponseData) SetTotalNum(v int64) {
	o.TotalNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *HashrateResaleListResponseData) GetPageSize() int64 {
	if o == nil || common.IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleListResponseData) GetPageSizeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *HashrateResaleListResponseData) HasPageSize() bool {
	if o != nil && !common.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *HashrateResaleListResponseData) SetPageSize(v int64) {
	o.PageSize = &v
}

func (o HashrateResaleListResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HashrateResaleListResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ConfigDetails) {
		toSerialize["configDetails"] = o.ConfigDetails
	}
	if !common.IsNil(o.TotalNum) {
		toSerialize["totalNum"] = o.TotalNum
	}
	if !common.IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HashrateResaleListResponseData) UnmarshalJSON(data []byte) (err error) {
	varHashrateResaleListResponseData := _HashrateResaleListResponseData{}

	err = json.Unmarshal(data, &varHashrateResaleListResponseData)

	if err != nil {
		return err
	}

	*o = HashrateResaleListResponseData(varHashrateResaleListResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "configDetails")
		delete(additionalProperties, "totalNum")
		delete(additionalProperties, "pageSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHashrateResaleListResponseData struct {
	value *HashrateResaleListResponseData
	isSet bool
}

func (v NullableHashrateResaleListResponseData) Get() *HashrateResaleListResponseData {
	return v.value
}

func (v *NullableHashrateResaleListResponseData) Set(val *HashrateResaleListResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableHashrateResaleListResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableHashrateResaleListResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHashrateResaleListResponseData(val *HashrateResaleListResponseData) *NullableHashrateResaleListResponseData {
	return &NullableHashrateResaleListResponseData{value: val, isSet: true}
}

func (v NullableHashrateResaleListResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHashrateResaleListResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
