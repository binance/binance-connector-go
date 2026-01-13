/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the RequestForMinerListResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RequestForMinerListResponseData{}

// RequestForMinerListResponseData struct for RequestForMinerListResponseData
type RequestForMinerListResponseData struct {
	WorkerDatas          []RequestForMinerListResponseDataWorkerDatasInner `json:"workerDatas,omitempty"`
	TotalNum             *int64                                            `json:"totalNum,omitempty"`
	PageSize             *int64                                            `json:"pageSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestForMinerListResponseData RequestForMinerListResponseData

// NewRequestForMinerListResponseData instantiates a new RequestForMinerListResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestForMinerListResponseData() *RequestForMinerListResponseData {
	this := RequestForMinerListResponseData{}
	return &this
}

// NewRequestForMinerListResponseDataWithDefaults instantiates a new RequestForMinerListResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestForMinerListResponseDataWithDefaults() *RequestForMinerListResponseData {
	this := RequestForMinerListResponseData{}
	return &this
}

// GetWorkerDatas returns the WorkerDatas field value if set, zero value otherwise.
func (o *RequestForMinerListResponseData) GetWorkerDatas() []RequestForMinerListResponseDataWorkerDatasInner {
	if o == nil || common.IsNil(o.WorkerDatas) {
		var ret []RequestForMinerListResponseDataWorkerDatasInner
		return ret
	}
	return o.WorkerDatas
}

// GetWorkerDatasOk returns a tuple with the WorkerDatas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForMinerListResponseData) GetWorkerDatasOk() ([]RequestForMinerListResponseDataWorkerDatasInner, bool) {
	if o == nil || common.IsNil(o.WorkerDatas) {
		return nil, false
	}
	return o.WorkerDatas, true
}

// HasWorkerDatas returns a boolean if a field has been set.
func (o *RequestForMinerListResponseData) HasWorkerDatas() bool {
	if o != nil && !common.IsNil(o.WorkerDatas) {
		return true
	}

	return false
}

// SetWorkerDatas gets a reference to the given []RequestForMinerListResponseDataWorkerDatasInner and assigns it to the WorkerDatas field.
func (o *RequestForMinerListResponseData) SetWorkerDatas(v []RequestForMinerListResponseDataWorkerDatasInner) {
	o.WorkerDatas = v
}

// GetTotalNum returns the TotalNum field value if set, zero value otherwise.
func (o *RequestForMinerListResponseData) GetTotalNum() int64 {
	if o == nil || common.IsNil(o.TotalNum) {
		var ret int64
		return ret
	}
	return *o.TotalNum
}

// GetTotalNumOk returns a tuple with the TotalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForMinerListResponseData) GetTotalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TotalNum) {
		return nil, false
	}
	return o.TotalNum, true
}

// HasTotalNum returns a boolean if a field has been set.
func (o *RequestForMinerListResponseData) HasTotalNum() bool {
	if o != nil && !common.IsNil(o.TotalNum) {
		return true
	}

	return false
}

// SetTotalNum gets a reference to the given int64 and assigns it to the TotalNum field.
func (o *RequestForMinerListResponseData) SetTotalNum(v int64) {
	o.TotalNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *RequestForMinerListResponseData) GetPageSize() int64 {
	if o == nil || common.IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForMinerListResponseData) GetPageSizeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *RequestForMinerListResponseData) HasPageSize() bool {
	if o != nil && !common.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *RequestForMinerListResponseData) SetPageSize(v int64) {
	o.PageSize = &v
}

func (o RequestForMinerListResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestForMinerListResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.WorkerDatas) {
		toSerialize["workerDatas"] = o.WorkerDatas
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

func (o *RequestForMinerListResponseData) UnmarshalJSON(data []byte) (err error) {
	varRequestForMinerListResponseData := _RequestForMinerListResponseData{}

	err = json.Unmarshal(data, &varRequestForMinerListResponseData)

	if err != nil {
		return err
	}

	*o = RequestForMinerListResponseData(varRequestForMinerListResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "workerDatas")
		delete(additionalProperties, "totalNum")
		delete(additionalProperties, "pageSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestForMinerListResponseData struct {
	value *RequestForMinerListResponseData
	isSet bool
}

func (v NullableRequestForMinerListResponseData) Get() *RequestForMinerListResponseData {
	return v.value
}

func (v *NullableRequestForMinerListResponseData) Set(val *RequestForMinerListResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestForMinerListResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestForMinerListResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestForMinerListResponseData(val *RequestForMinerListResponseData) *NullableRequestForMinerListResponseData {
	return &NullableRequestForMinerListResponseData{value: val, isSet: true}
}

func (v NullableRequestForMinerListResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestForMinerListResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
