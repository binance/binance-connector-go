/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryUserUniversalTransferHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUserUniversalTransferHistoryResponse{}

// QueryUserUniversalTransferHistoryResponse struct for QueryUserUniversalTransferHistoryResponse
type QueryUserUniversalTransferHistoryResponse struct {
	Total                *int64                                               `json:"total,omitempty"`
	Rows                 []QueryUserUniversalTransferHistoryResponseRowsInner `json:"rows,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryUserUniversalTransferHistoryResponse QueryUserUniversalTransferHistoryResponse

// NewQueryUserUniversalTransferHistoryResponse instantiates a new QueryUserUniversalTransferHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUserUniversalTransferHistoryResponse() *QueryUserUniversalTransferHistoryResponse {
	this := QueryUserUniversalTransferHistoryResponse{}
	return &this
}

// NewQueryUserUniversalTransferHistoryResponseWithDefaults instantiates a new QueryUserUniversalTransferHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUserUniversalTransferHistoryResponseWithDefaults() *QueryUserUniversalTransferHistoryResponse {
	this := QueryUserUniversalTransferHistoryResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryUserUniversalTransferHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserUniversalTransferHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryUserUniversalTransferHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QueryUserUniversalTransferHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *QueryUserUniversalTransferHistoryResponse) GetRows() []QueryUserUniversalTransferHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []QueryUserUniversalTransferHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserUniversalTransferHistoryResponse) GetRowsOk() ([]QueryUserUniversalTransferHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *QueryUserUniversalTransferHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []QueryUserUniversalTransferHistoryResponseRowsInner and assigns it to the Rows field.
func (o *QueryUserUniversalTransferHistoryResponse) SetRows(v []QueryUserUniversalTransferHistoryResponseRowsInner) {
	o.Rows = v
}

func (o QueryUserUniversalTransferHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUserUniversalTransferHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !common.IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryUserUniversalTransferHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryUserUniversalTransferHistoryResponse := _QueryUserUniversalTransferHistoryResponse{}

	err = json.Unmarshal(data, &varQueryUserUniversalTransferHistoryResponse)

	if err != nil {
		return err
	}

	*o = QueryUserUniversalTransferHistoryResponse(varQueryUserUniversalTransferHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "rows")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryUserUniversalTransferHistoryResponse struct {
	value *QueryUserUniversalTransferHistoryResponse
	isSet bool
}

func (v NullableQueryUserUniversalTransferHistoryResponse) Get() *QueryUserUniversalTransferHistoryResponse {
	return v.value
}

func (v *NullableQueryUserUniversalTransferHistoryResponse) Set(val *QueryUserUniversalTransferHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUserUniversalTransferHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUserUniversalTransferHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUserUniversalTransferHistoryResponse(val *QueryUserUniversalTransferHistoryResponse) *NullableQueryUserUniversalTransferHistoryResponse {
	return &NullableQueryUserUniversalTransferHistoryResponse{value: val, isSet: true}
}

func (v NullableQueryUserUniversalTransferHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUserUniversalTransferHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
