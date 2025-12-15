/*
Binance Rebate REST API

OpenAPI Specification for the Binance Rebate REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSpotRebateHistoryRecordsResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSpotRebateHistoryRecordsResponseData{}

// GetSpotRebateHistoryRecordsResponseData struct for GetSpotRebateHistoryRecordsResponseData
type GetSpotRebateHistoryRecordsResponseData struct {
	Page                 *int64                                             `json:"page,omitempty"`
	TotalRecords         *int64                                             `json:"totalRecords,omitempty"`
	TotalPageNum         *int64                                             `json:"totalPageNum,omitempty"`
	Data                 []GetSpotRebateHistoryRecordsResponseDataDataInner `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSpotRebateHistoryRecordsResponseData GetSpotRebateHistoryRecordsResponseData

// NewGetSpotRebateHistoryRecordsResponseData instantiates a new GetSpotRebateHistoryRecordsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSpotRebateHistoryRecordsResponseData() *GetSpotRebateHistoryRecordsResponseData {
	this := GetSpotRebateHistoryRecordsResponseData{}
	return &this
}

// NewGetSpotRebateHistoryRecordsResponseDataWithDefaults instantiates a new GetSpotRebateHistoryRecordsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSpotRebateHistoryRecordsResponseDataWithDefaults() *GetSpotRebateHistoryRecordsResponseData {
	this := GetSpotRebateHistoryRecordsResponseData{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetSpotRebateHistoryRecordsResponseData) GetPage() int64 {
	if o == nil || common.IsNil(o.Page) {
		var ret int64
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotRebateHistoryRecordsResponseData) GetPageOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetSpotRebateHistoryRecordsResponseData) HasPage() bool {
	if o != nil && !common.IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int64 and assigns it to the Page field.
func (o *GetSpotRebateHistoryRecordsResponseData) SetPage(v int64) {
	o.Page = &v
}

// GetTotalRecords returns the TotalRecords field value if set, zero value otherwise.
func (o *GetSpotRebateHistoryRecordsResponseData) GetTotalRecords() int64 {
	if o == nil || common.IsNil(o.TotalRecords) {
		var ret int64
		return ret
	}
	return *o.TotalRecords
}

// GetTotalRecordsOk returns a tuple with the TotalRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotRebateHistoryRecordsResponseData) GetTotalRecordsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TotalRecords) {
		return nil, false
	}
	return o.TotalRecords, true
}

// HasTotalRecords returns a boolean if a field has been set.
func (o *GetSpotRebateHistoryRecordsResponseData) HasTotalRecords() bool {
	if o != nil && !common.IsNil(o.TotalRecords) {
		return true
	}

	return false
}

// SetTotalRecords gets a reference to the given int64 and assigns it to the TotalRecords field.
func (o *GetSpotRebateHistoryRecordsResponseData) SetTotalRecords(v int64) {
	o.TotalRecords = &v
}

// GetTotalPageNum returns the TotalPageNum field value if set, zero value otherwise.
func (o *GetSpotRebateHistoryRecordsResponseData) GetTotalPageNum() int64 {
	if o == nil || common.IsNil(o.TotalPageNum) {
		var ret int64
		return ret
	}
	return *o.TotalPageNum
}

// GetTotalPageNumOk returns a tuple with the TotalPageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotRebateHistoryRecordsResponseData) GetTotalPageNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TotalPageNum) {
		return nil, false
	}
	return o.TotalPageNum, true
}

// HasTotalPageNum returns a boolean if a field has been set.
func (o *GetSpotRebateHistoryRecordsResponseData) HasTotalPageNum() bool {
	if o != nil && !common.IsNil(o.TotalPageNum) {
		return true
	}

	return false
}

// SetTotalPageNum gets a reference to the given int64 and assigns it to the TotalPageNum field.
func (o *GetSpotRebateHistoryRecordsResponseData) SetTotalPageNum(v int64) {
	o.TotalPageNum = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetSpotRebateHistoryRecordsResponseData) GetData() []GetSpotRebateHistoryRecordsResponseDataDataInner {
	if o == nil || common.IsNil(o.Data) {
		var ret []GetSpotRebateHistoryRecordsResponseDataDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotRebateHistoryRecordsResponseData) GetDataOk() ([]GetSpotRebateHistoryRecordsResponseDataDataInner, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetSpotRebateHistoryRecordsResponseData) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetSpotRebateHistoryRecordsResponseDataDataInner and assigns it to the Data field.
func (o *GetSpotRebateHistoryRecordsResponseData) SetData(v []GetSpotRebateHistoryRecordsResponseDataDataInner) {
	o.Data = v
}

func (o GetSpotRebateHistoryRecordsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSpotRebateHistoryRecordsResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !common.IsNil(o.TotalRecords) {
		toSerialize["totalRecords"] = o.TotalRecords
	}
	if !common.IsNil(o.TotalPageNum) {
		toSerialize["totalPageNum"] = o.TotalPageNum
	}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSpotRebateHistoryRecordsResponseData) UnmarshalJSON(data []byte) (err error) {
	varGetSpotRebateHistoryRecordsResponseData := _GetSpotRebateHistoryRecordsResponseData{}

	err = json.Unmarshal(data, &varGetSpotRebateHistoryRecordsResponseData)

	if err != nil {
		return err
	}

	*o = GetSpotRebateHistoryRecordsResponseData(varGetSpotRebateHistoryRecordsResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "page")
		delete(additionalProperties, "totalRecords")
		delete(additionalProperties, "totalPageNum")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSpotRebateHistoryRecordsResponseData struct {
	value *GetSpotRebateHistoryRecordsResponseData
	isSet bool
}

func (v NullableGetSpotRebateHistoryRecordsResponseData) Get() *GetSpotRebateHistoryRecordsResponseData {
	return v.value
}

func (v *NullableGetSpotRebateHistoryRecordsResponseData) Set(val *GetSpotRebateHistoryRecordsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSpotRebateHistoryRecordsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSpotRebateHistoryRecordsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSpotRebateHistoryRecordsResponseData(val *GetSpotRebateHistoryRecordsResponseData) *NullableGetSpotRebateHistoryRecordsResponseData {
	return &NullableGetSpotRebateHistoryRecordsResponseData{value: val, isSet: true}
}

func (v NullableGetSpotRebateHistoryRecordsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSpotRebateHistoryRecordsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
