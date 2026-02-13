/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExtraBonusListResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExtraBonusListResponseData{}

// ExtraBonusListResponseData struct for ExtraBonusListResponseData
type ExtraBonusListResponseData struct {
	OtherProfits         []ExtraBonusListResponseDataOtherProfitsInner `json:"otherProfits,omitempty"`
	TotalNum             *int64                                        `json:"totalNum,omitempty"`
	PageSize             *int64                                        `json:"pageSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExtraBonusListResponseData ExtraBonusListResponseData

// NewExtraBonusListResponseData instantiates a new ExtraBonusListResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtraBonusListResponseData() *ExtraBonusListResponseData {
	this := ExtraBonusListResponseData{}
	return &this
}

// NewExtraBonusListResponseDataWithDefaults instantiates a new ExtraBonusListResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtraBonusListResponseDataWithDefaults() *ExtraBonusListResponseData {
	this := ExtraBonusListResponseData{}
	return &this
}

// GetOtherProfits returns the OtherProfits field value if set, zero value otherwise.
func (o *ExtraBonusListResponseData) GetOtherProfits() []ExtraBonusListResponseDataOtherProfitsInner {
	if o == nil || common.IsNil(o.OtherProfits) {
		var ret []ExtraBonusListResponseDataOtherProfitsInner
		return ret
	}
	return o.OtherProfits
}

// GetOtherProfitsOk returns a tuple with the OtherProfits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraBonusListResponseData) GetOtherProfitsOk() ([]ExtraBonusListResponseDataOtherProfitsInner, bool) {
	if o == nil || common.IsNil(o.OtherProfits) {
		return nil, false
	}
	return o.OtherProfits, true
}

// HasOtherProfits returns a boolean if a field has been set.
func (o *ExtraBonusListResponseData) HasOtherProfits() bool {
	if o != nil && !common.IsNil(o.OtherProfits) {
		return true
	}

	return false
}

// SetOtherProfits gets a reference to the given []ExtraBonusListResponseDataOtherProfitsInner and assigns it to the OtherProfits field.
func (o *ExtraBonusListResponseData) SetOtherProfits(v []ExtraBonusListResponseDataOtherProfitsInner) {
	o.OtherProfits = v
}

// GetTotalNum returns the TotalNum field value if set, zero value otherwise.
func (o *ExtraBonusListResponseData) GetTotalNum() int64 {
	if o == nil || common.IsNil(o.TotalNum) {
		var ret int64
		return ret
	}
	return *o.TotalNum
}

// GetTotalNumOk returns a tuple with the TotalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraBonusListResponseData) GetTotalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TotalNum) {
		return nil, false
	}
	return o.TotalNum, true
}

// HasTotalNum returns a boolean if a field has been set.
func (o *ExtraBonusListResponseData) HasTotalNum() bool {
	if o != nil && !common.IsNil(o.TotalNum) {
		return true
	}

	return false
}

// SetTotalNum gets a reference to the given int64 and assigns it to the TotalNum field.
func (o *ExtraBonusListResponseData) SetTotalNum(v int64) {
	o.TotalNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ExtraBonusListResponseData) GetPageSize() int64 {
	if o == nil || common.IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraBonusListResponseData) GetPageSizeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ExtraBonusListResponseData) HasPageSize() bool {
	if o != nil && !common.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *ExtraBonusListResponseData) SetPageSize(v int64) {
	o.PageSize = &v
}

func (o ExtraBonusListResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtraBonusListResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OtherProfits) {
		toSerialize["otherProfits"] = o.OtherProfits
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

func (o *ExtraBonusListResponseData) UnmarshalJSON(data []byte) (err error) {
	varExtraBonusListResponseData := _ExtraBonusListResponseData{}

	err = json.Unmarshal(data, &varExtraBonusListResponseData)

	if err != nil {
		return err
	}

	*o = ExtraBonusListResponseData(varExtraBonusListResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "otherProfits")
		delete(additionalProperties, "totalNum")
		delete(additionalProperties, "pageSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExtraBonusListResponseData struct {
	value *ExtraBonusListResponseData
	isSet bool
}

func (v NullableExtraBonusListResponseData) Get() *ExtraBonusListResponseData {
	return v.value
}

func (v *NullableExtraBonusListResponseData) Set(val *ExtraBonusListResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableExtraBonusListResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableExtraBonusListResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtraBonusListResponseData(val *ExtraBonusListResponseData) *NullableExtraBonusListResponseData {
	return &NullableExtraBonusListResponseData{value: val, isSet: true}
}

func (v NullableExtraBonusListResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtraBonusListResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
