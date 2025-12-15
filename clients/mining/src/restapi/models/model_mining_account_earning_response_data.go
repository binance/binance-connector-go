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

// checks if the MiningAccountEarningResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MiningAccountEarningResponseData{}

// MiningAccountEarningResponseData struct for MiningAccountEarningResponseData
type MiningAccountEarningResponseData struct {
	AccountProfits       []MiningAccountEarningResponseDataAccountProfitsInner `json:"accountProfits,omitempty"`
	TotalNum             *int64                                                `json:"totalNum,omitempty"`
	PageSize             *int64                                                `json:"pageSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MiningAccountEarningResponseData MiningAccountEarningResponseData

// NewMiningAccountEarningResponseData instantiates a new MiningAccountEarningResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiningAccountEarningResponseData() *MiningAccountEarningResponseData {
	this := MiningAccountEarningResponseData{}
	return &this
}

// NewMiningAccountEarningResponseDataWithDefaults instantiates a new MiningAccountEarningResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiningAccountEarningResponseDataWithDefaults() *MiningAccountEarningResponseData {
	this := MiningAccountEarningResponseData{}
	return &this
}

// GetAccountProfits returns the AccountProfits field value if set, zero value otherwise.
func (o *MiningAccountEarningResponseData) GetAccountProfits() []MiningAccountEarningResponseDataAccountProfitsInner {
	if o == nil || common.IsNil(o.AccountProfits) {
		var ret []MiningAccountEarningResponseDataAccountProfitsInner
		return ret
	}
	return o.AccountProfits
}

// GetAccountProfitsOk returns a tuple with the AccountProfits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiningAccountEarningResponseData) GetAccountProfitsOk() ([]MiningAccountEarningResponseDataAccountProfitsInner, bool) {
	if o == nil || common.IsNil(o.AccountProfits) {
		return nil, false
	}
	return o.AccountProfits, true
}

// HasAccountProfits returns a boolean if a field has been set.
func (o *MiningAccountEarningResponseData) HasAccountProfits() bool {
	if o != nil && !common.IsNil(o.AccountProfits) {
		return true
	}

	return false
}

// SetAccountProfits gets a reference to the given []MiningAccountEarningResponseDataAccountProfitsInner and assigns it to the AccountProfits field.
func (o *MiningAccountEarningResponseData) SetAccountProfits(v []MiningAccountEarningResponseDataAccountProfitsInner) {
	o.AccountProfits = v
}

// GetTotalNum returns the TotalNum field value if set, zero value otherwise.
func (o *MiningAccountEarningResponseData) GetTotalNum() int64 {
	if o == nil || common.IsNil(o.TotalNum) {
		var ret int64
		return ret
	}
	return *o.TotalNum
}

// GetTotalNumOk returns a tuple with the TotalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiningAccountEarningResponseData) GetTotalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TotalNum) {
		return nil, false
	}
	return o.TotalNum, true
}

// HasTotalNum returns a boolean if a field has been set.
func (o *MiningAccountEarningResponseData) HasTotalNum() bool {
	if o != nil && !common.IsNil(o.TotalNum) {
		return true
	}

	return false
}

// SetTotalNum gets a reference to the given int64 and assigns it to the TotalNum field.
func (o *MiningAccountEarningResponseData) SetTotalNum(v int64) {
	o.TotalNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *MiningAccountEarningResponseData) GetPageSize() int64 {
	if o == nil || common.IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiningAccountEarningResponseData) GetPageSizeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *MiningAccountEarningResponseData) HasPageSize() bool {
	if o != nil && !common.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *MiningAccountEarningResponseData) SetPageSize(v int64) {
	o.PageSize = &v
}

func (o MiningAccountEarningResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiningAccountEarningResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountProfits) {
		toSerialize["accountProfits"] = o.AccountProfits
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

func (o *MiningAccountEarningResponseData) UnmarshalJSON(data []byte) (err error) {
	varMiningAccountEarningResponseData := _MiningAccountEarningResponseData{}

	err = json.Unmarshal(data, &varMiningAccountEarningResponseData)

	if err != nil {
		return err
	}

	*o = MiningAccountEarningResponseData(varMiningAccountEarningResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountProfits")
		delete(additionalProperties, "totalNum")
		delete(additionalProperties, "pageSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMiningAccountEarningResponseData struct {
	value *MiningAccountEarningResponseData
	isSet bool
}

func (v NullableMiningAccountEarningResponseData) Get() *MiningAccountEarningResponseData {
	return v.value
}

func (v *NullableMiningAccountEarningResponseData) Set(val *MiningAccountEarningResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMiningAccountEarningResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMiningAccountEarningResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiningAccountEarningResponseData(val *MiningAccountEarningResponseData) *NullableMiningAccountEarningResponseData {
	return &NullableMiningAccountEarningResponseData{value: val, isSet: true}
}

func (v NullableMiningAccountEarningResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiningAccountEarningResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
