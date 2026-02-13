/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the StatisticListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StatisticListResponse{}

// StatisticListResponse struct for StatisticListResponse
type StatisticListResponse struct {
	Code                 *int64                     `json:"code,omitempty"`
	Msg                  *string                    `json:"msg,omitempty"`
	Data                 *StatisticListResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StatisticListResponse StatisticListResponse

// NewStatisticListResponse instantiates a new StatisticListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticListResponse() *StatisticListResponse {
	this := StatisticListResponse{}
	return &this
}

// NewStatisticListResponseWithDefaults instantiates a new StatisticListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticListResponseWithDefaults() *StatisticListResponse {
	this := StatisticListResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *StatisticListResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticListResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *StatisticListResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *StatisticListResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *StatisticListResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticListResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *StatisticListResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *StatisticListResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StatisticListResponse) GetData() StatisticListResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret StatisticListResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticListResponse) GetDataOk() (*StatisticListResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StatisticListResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given StatisticListResponseData and assigns it to the Data field.
func (o *StatisticListResponse) SetData(v StatisticListResponseData) {
	o.Data = &v
}

func (o StatisticListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatisticListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StatisticListResponse) UnmarshalJSON(data []byte) (err error) {
	varStatisticListResponse := _StatisticListResponse{}

	err = json.Unmarshal(data, &varStatisticListResponse)

	if err != nil {
		return err
	}

	*o = StatisticListResponse(varStatisticListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStatisticListResponse struct {
	value *StatisticListResponse
	isSet bool
}

func (v NullableStatisticListResponse) Get() *StatisticListResponse {
	return v.value
}

func (v *NullableStatisticListResponse) Set(val *StatisticListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticListResponse(val *StatisticListResponse) *NullableStatisticListResponse {
	return &NullableStatisticListResponse{value: val, isSet: true}
}

func (v NullableStatisticListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
