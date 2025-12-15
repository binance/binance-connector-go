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

// checks if the EarningsListResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EarningsListResponseData{}

// EarningsListResponseData struct for EarningsListResponseData
type EarningsListResponseData struct {
	AccountProfits       []EarningsListResponseDataAccountProfitsInner `json:"accountProfits,omitempty"`
	TotalNum             *int64                                        `json:"totalNum,omitempty"`
	PageSize             *int64                                        `json:"pageSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EarningsListResponseData EarningsListResponseData

// NewEarningsListResponseData instantiates a new EarningsListResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEarningsListResponseData() *EarningsListResponseData {
	this := EarningsListResponseData{}
	return &this
}

// NewEarningsListResponseDataWithDefaults instantiates a new EarningsListResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEarningsListResponseDataWithDefaults() *EarningsListResponseData {
	this := EarningsListResponseData{}
	return &this
}

// GetAccountProfits returns the AccountProfits field value if set, zero value otherwise.
func (o *EarningsListResponseData) GetAccountProfits() []EarningsListResponseDataAccountProfitsInner {
	if o == nil || common.IsNil(o.AccountProfits) {
		var ret []EarningsListResponseDataAccountProfitsInner
		return ret
	}
	return o.AccountProfits
}

// GetAccountProfitsOk returns a tuple with the AccountProfits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsListResponseData) GetAccountProfitsOk() ([]EarningsListResponseDataAccountProfitsInner, bool) {
	if o == nil || common.IsNil(o.AccountProfits) {
		return nil, false
	}
	return o.AccountProfits, true
}

// HasAccountProfits returns a boolean if a field has been set.
func (o *EarningsListResponseData) HasAccountProfits() bool {
	if o != nil && !common.IsNil(o.AccountProfits) {
		return true
	}

	return false
}

// SetAccountProfits gets a reference to the given []EarningsListResponseDataAccountProfitsInner and assigns it to the AccountProfits field.
func (o *EarningsListResponseData) SetAccountProfits(v []EarningsListResponseDataAccountProfitsInner) {
	o.AccountProfits = v
}

// GetTotalNum returns the TotalNum field value if set, zero value otherwise.
func (o *EarningsListResponseData) GetTotalNum() int64 {
	if o == nil || common.IsNil(o.TotalNum) {
		var ret int64
		return ret
	}
	return *o.TotalNum
}

// GetTotalNumOk returns a tuple with the TotalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsListResponseData) GetTotalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TotalNum) {
		return nil, false
	}
	return o.TotalNum, true
}

// HasTotalNum returns a boolean if a field has been set.
func (o *EarningsListResponseData) HasTotalNum() bool {
	if o != nil && !common.IsNil(o.TotalNum) {
		return true
	}

	return false
}

// SetTotalNum gets a reference to the given int64 and assigns it to the TotalNum field.
func (o *EarningsListResponseData) SetTotalNum(v int64) {
	o.TotalNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *EarningsListResponseData) GetPageSize() int64 {
	if o == nil || common.IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsListResponseData) GetPageSizeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *EarningsListResponseData) HasPageSize() bool {
	if o != nil && !common.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *EarningsListResponseData) SetPageSize(v int64) {
	o.PageSize = &v
}

func (o EarningsListResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EarningsListResponseData) ToMap() (map[string]interface{}, error) {
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

func (o *EarningsListResponseData) UnmarshalJSON(data []byte) (err error) {
	varEarningsListResponseData := _EarningsListResponseData{}

	err = json.Unmarshal(data, &varEarningsListResponseData)

	if err != nil {
		return err
	}

	*o = EarningsListResponseData(varEarningsListResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountProfits")
		delete(additionalProperties, "totalNum")
		delete(additionalProperties, "pageSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEarningsListResponseData struct {
	value *EarningsListResponseData
	isSet bool
}

func (v NullableEarningsListResponseData) Get() *EarningsListResponseData {
	return v.value
}

func (v *NullableEarningsListResponseData) Set(val *EarningsListResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableEarningsListResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableEarningsListResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEarningsListResponseData(val *EarningsListResponseData) *NullableEarningsListResponseData {
	return &NullableEarningsListResponseData{value: val, isSet: true}
}

func (v NullableEarningsListResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEarningsListResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
