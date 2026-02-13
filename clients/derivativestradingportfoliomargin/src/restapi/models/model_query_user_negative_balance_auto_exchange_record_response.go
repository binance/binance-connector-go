/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryUserNegativeBalanceAutoExchangeRecordResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUserNegativeBalanceAutoExchangeRecordResponse{}

// QueryUserNegativeBalanceAutoExchangeRecordResponse struct for QueryUserNegativeBalanceAutoExchangeRecordResponse
type QueryUserNegativeBalanceAutoExchangeRecordResponse struct {
	Total                *int64                                                        `json:"total,omitempty"`
	Rows                 []QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner `json:"rows,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryUserNegativeBalanceAutoExchangeRecordResponse QueryUserNegativeBalanceAutoExchangeRecordResponse

// NewQueryUserNegativeBalanceAutoExchangeRecordResponse instantiates a new QueryUserNegativeBalanceAutoExchangeRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUserNegativeBalanceAutoExchangeRecordResponse() *QueryUserNegativeBalanceAutoExchangeRecordResponse {
	this := QueryUserNegativeBalanceAutoExchangeRecordResponse{}
	return &this
}

// NewQueryUserNegativeBalanceAutoExchangeRecordResponseWithDefaults instantiates a new QueryUserNegativeBalanceAutoExchangeRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUserNegativeBalanceAutoExchangeRecordResponseWithDefaults() *QueryUserNegativeBalanceAutoExchangeRecordResponse {
	this := QueryUserNegativeBalanceAutoExchangeRecordResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) GetRows() []QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) GetRowsOk() ([]QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner and assigns it to the Rows field.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) SetRows(v []QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) {
	o.Rows = v
}

func (o QueryUserNegativeBalanceAutoExchangeRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUserNegativeBalanceAutoExchangeRecordResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryUserNegativeBalanceAutoExchangeRecordResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryUserNegativeBalanceAutoExchangeRecordResponse := _QueryUserNegativeBalanceAutoExchangeRecordResponse{}

	err = json.Unmarshal(data, &varQueryUserNegativeBalanceAutoExchangeRecordResponse)

	if err != nil {
		return err
	}

	*o = QueryUserNegativeBalanceAutoExchangeRecordResponse(varQueryUserNegativeBalanceAutoExchangeRecordResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "rows")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryUserNegativeBalanceAutoExchangeRecordResponse struct {
	value *QueryUserNegativeBalanceAutoExchangeRecordResponse
	isSet bool
}

func (v NullableQueryUserNegativeBalanceAutoExchangeRecordResponse) Get() *QueryUserNegativeBalanceAutoExchangeRecordResponse {
	return v.value
}

func (v *NullableQueryUserNegativeBalanceAutoExchangeRecordResponse) Set(val *QueryUserNegativeBalanceAutoExchangeRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUserNegativeBalanceAutoExchangeRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUserNegativeBalanceAutoExchangeRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUserNegativeBalanceAutoExchangeRecordResponse(val *QueryUserNegativeBalanceAutoExchangeRecordResponse) *NullableQueryUserNegativeBalanceAutoExchangeRecordResponse {
	return &NullableQueryUserNegativeBalanceAutoExchangeRecordResponse{value: val, isSet: true}
}

func (v NullableQueryUserNegativeBalanceAutoExchangeRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUserNegativeBalanceAutoExchangeRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
