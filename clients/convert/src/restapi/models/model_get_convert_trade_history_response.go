/*
Binance Convert REST API

OpenAPI Specification for the Binance Convert REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetConvertTradeHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetConvertTradeHistoryResponse{}

// GetConvertTradeHistoryResponse struct for GetConvertTradeHistoryResponse
type GetConvertTradeHistoryResponse struct {
	List                 []GetConvertTradeHistoryResponseListInner `json:"list,omitempty"`
	StartTime            *int64                                    `json:"startTime,omitempty"`
	EndTime              *int64                                    `json:"endTime,omitempty"`
	Limit                *int64                                    `json:"limit,omitempty"`
	MoreData             *bool                                     `json:"moreData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetConvertTradeHistoryResponse GetConvertTradeHistoryResponse

// NewGetConvertTradeHistoryResponse instantiates a new GetConvertTradeHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConvertTradeHistoryResponse() *GetConvertTradeHistoryResponse {
	this := GetConvertTradeHistoryResponse{}
	return &this
}

// NewGetConvertTradeHistoryResponseWithDefaults instantiates a new GetConvertTradeHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConvertTradeHistoryResponseWithDefaults() *GetConvertTradeHistoryResponse {
	this := GetConvertTradeHistoryResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *GetConvertTradeHistoryResponse) GetList() []GetConvertTradeHistoryResponseListInner {
	if o == nil || common.IsNil(o.List) {
		var ret []GetConvertTradeHistoryResponseListInner
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConvertTradeHistoryResponse) GetListOk() ([]GetConvertTradeHistoryResponseListInner, bool) {
	if o == nil || common.IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *GetConvertTradeHistoryResponse) HasList() bool {
	if o != nil && !common.IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []GetConvertTradeHistoryResponseListInner and assigns it to the List field.
func (o *GetConvertTradeHistoryResponse) SetList(v []GetConvertTradeHistoryResponseListInner) {
	o.List = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GetConvertTradeHistoryResponse) GetStartTime() int64 {
	if o == nil || common.IsNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConvertTradeHistoryResponse) GetStartTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GetConvertTradeHistoryResponse) HasStartTime() bool {
	if o != nil && !common.IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *GetConvertTradeHistoryResponse) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *GetConvertTradeHistoryResponse) GetEndTime() int64 {
	if o == nil || common.IsNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConvertTradeHistoryResponse) GetEndTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *GetConvertTradeHistoryResponse) HasEndTime() bool {
	if o != nil && !common.IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *GetConvertTradeHistoryResponse) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetConvertTradeHistoryResponse) GetLimit() int64 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConvertTradeHistoryResponse) GetLimitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetConvertTradeHistoryResponse) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *GetConvertTradeHistoryResponse) SetLimit(v int64) {
	o.Limit = &v
}

// GetMoreData returns the MoreData field value if set, zero value otherwise.
func (o *GetConvertTradeHistoryResponse) GetMoreData() bool {
	if o == nil || common.IsNil(o.MoreData) {
		var ret bool
		return ret
	}
	return *o.MoreData
}

// GetMoreDataOk returns a tuple with the MoreData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConvertTradeHistoryResponse) GetMoreDataOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MoreData) {
		return nil, false
	}
	return o.MoreData, true
}

// HasMoreData returns a boolean if a field has been set.
func (o *GetConvertTradeHistoryResponse) HasMoreData() bool {
	if o != nil && !common.IsNil(o.MoreData) {
		return true
	}

	return false
}

// SetMoreData gets a reference to the given bool and assigns it to the MoreData field.
func (o *GetConvertTradeHistoryResponse) SetMoreData(v bool) {
	o.MoreData = &v
}

func (o GetConvertTradeHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetConvertTradeHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	if !common.IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !common.IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !common.IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !common.IsNil(o.MoreData) {
		toSerialize["moreData"] = o.MoreData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetConvertTradeHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetConvertTradeHistoryResponse := _GetConvertTradeHistoryResponse{}

	err = json.Unmarshal(data, &varGetConvertTradeHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetConvertTradeHistoryResponse(varGetConvertTradeHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "list")
		delete(additionalProperties, "startTime")
		delete(additionalProperties, "endTime")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "moreData")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetConvertTradeHistoryResponse struct {
	value *GetConvertTradeHistoryResponse
	isSet bool
}

func (v NullableGetConvertTradeHistoryResponse) Get() *GetConvertTradeHistoryResponse {
	return v.value
}

func (v *NullableGetConvertTradeHistoryResponse) Set(val *GetConvertTradeHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConvertTradeHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConvertTradeHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConvertTradeHistoryResponse(val *GetConvertTradeHistoryResponse) *NullableGetConvertTradeHistoryResponse {
	return &NullableGetConvertTradeHistoryResponse{value: val, isSet: true}
}

func (v NullableGetConvertTradeHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConvertTradeHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
