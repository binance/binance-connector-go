/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the HashrateResaleDetailResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &HashrateResaleDetailResponseData{}

// HashrateResaleDetailResponseData struct for HashrateResaleDetailResponseData
type HashrateResaleDetailResponseData struct {
	ProfitTransferDetails []HashrateResaleDetailResponseDataProfitTransferDetailsInner `json:"profitTransferDetails,omitempty"`
	TotalNum              *int64                                                       `json:"totalNum,omitempty"`
	PageSize              *int64                                                       `json:"pageSize,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _HashrateResaleDetailResponseData HashrateResaleDetailResponseData

// NewHashrateResaleDetailResponseData instantiates a new HashrateResaleDetailResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHashrateResaleDetailResponseData() *HashrateResaleDetailResponseData {
	this := HashrateResaleDetailResponseData{}
	return &this
}

// NewHashrateResaleDetailResponseDataWithDefaults instantiates a new HashrateResaleDetailResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHashrateResaleDetailResponseDataWithDefaults() *HashrateResaleDetailResponseData {
	this := HashrateResaleDetailResponseData{}
	return &this
}

// GetProfitTransferDetails returns the ProfitTransferDetails field value if set, zero value otherwise.
func (o *HashrateResaleDetailResponseData) GetProfitTransferDetails() []HashrateResaleDetailResponseDataProfitTransferDetailsInner {
	if o == nil || common.IsNil(o.ProfitTransferDetails) {
		var ret []HashrateResaleDetailResponseDataProfitTransferDetailsInner
		return ret
	}
	return o.ProfitTransferDetails
}

// GetProfitTransferDetailsOk returns a tuple with the ProfitTransferDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleDetailResponseData) GetProfitTransferDetailsOk() ([]HashrateResaleDetailResponseDataProfitTransferDetailsInner, bool) {
	if o == nil || common.IsNil(o.ProfitTransferDetails) {
		return nil, false
	}
	return o.ProfitTransferDetails, true
}

// HasProfitTransferDetails returns a boolean if a field has been set.
func (o *HashrateResaleDetailResponseData) HasProfitTransferDetails() bool {
	if o != nil && !common.IsNil(o.ProfitTransferDetails) {
		return true
	}

	return false
}

// SetProfitTransferDetails gets a reference to the given []HashrateResaleDetailResponseDataProfitTransferDetailsInner and assigns it to the ProfitTransferDetails field.
func (o *HashrateResaleDetailResponseData) SetProfitTransferDetails(v []HashrateResaleDetailResponseDataProfitTransferDetailsInner) {
	o.ProfitTransferDetails = v
}

// GetTotalNum returns the TotalNum field value if set, zero value otherwise.
func (o *HashrateResaleDetailResponseData) GetTotalNum() int64 {
	if o == nil || common.IsNil(o.TotalNum) {
		var ret int64
		return ret
	}
	return *o.TotalNum
}

// GetTotalNumOk returns a tuple with the TotalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleDetailResponseData) GetTotalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TotalNum) {
		return nil, false
	}
	return o.TotalNum, true
}

// HasTotalNum returns a boolean if a field has been set.
func (o *HashrateResaleDetailResponseData) HasTotalNum() bool {
	if o != nil && !common.IsNil(o.TotalNum) {
		return true
	}

	return false
}

// SetTotalNum gets a reference to the given int64 and assigns it to the TotalNum field.
func (o *HashrateResaleDetailResponseData) SetTotalNum(v int64) {
	o.TotalNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *HashrateResaleDetailResponseData) GetPageSize() int64 {
	if o == nil || common.IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleDetailResponseData) GetPageSizeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *HashrateResaleDetailResponseData) HasPageSize() bool {
	if o != nil && !common.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *HashrateResaleDetailResponseData) SetPageSize(v int64) {
	o.PageSize = &v
}

func (o HashrateResaleDetailResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HashrateResaleDetailResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ProfitTransferDetails) {
		toSerialize["profitTransferDetails"] = o.ProfitTransferDetails
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

func (o *HashrateResaleDetailResponseData) UnmarshalJSON(data []byte) (err error) {
	varHashrateResaleDetailResponseData := _HashrateResaleDetailResponseData{}

	err = json.Unmarshal(data, &varHashrateResaleDetailResponseData)

	if err != nil {
		return err
	}

	*o = HashrateResaleDetailResponseData(varHashrateResaleDetailResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "profitTransferDetails")
		delete(additionalProperties, "totalNum")
		delete(additionalProperties, "pageSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHashrateResaleDetailResponseData struct {
	value *HashrateResaleDetailResponseData
	isSet bool
}

func (v NullableHashrateResaleDetailResponseData) Get() *HashrateResaleDetailResponseData {
	return v.value
}

func (v *NullableHashrateResaleDetailResponseData) Set(val *HashrateResaleDetailResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableHashrateResaleDetailResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableHashrateResaleDetailResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHashrateResaleDetailResponseData(val *HashrateResaleDetailResponseData) *NullableHashrateResaleDetailResponseData {
	return &NullableHashrateResaleDetailResponseData{value: val, isSet: true}
}

func (v NullableHashrateResaleDetailResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHashrateResaleDetailResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
