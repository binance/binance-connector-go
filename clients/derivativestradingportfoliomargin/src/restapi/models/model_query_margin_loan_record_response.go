/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryMarginLoanRecordResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginLoanRecordResponse{}

// QueryMarginLoanRecordResponse struct for QueryMarginLoanRecordResponse
type QueryMarginLoanRecordResponse struct {
	Rows                 []QueryMarginLoanRecordResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMarginLoanRecordResponse QueryMarginLoanRecordResponse

// NewQueryMarginLoanRecordResponse instantiates a new QueryMarginLoanRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginLoanRecordResponse() *QueryMarginLoanRecordResponse {
	this := QueryMarginLoanRecordResponse{}
	return &this
}

// NewQueryMarginLoanRecordResponseWithDefaults instantiates a new QueryMarginLoanRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginLoanRecordResponseWithDefaults() *QueryMarginLoanRecordResponse {
	this := QueryMarginLoanRecordResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *QueryMarginLoanRecordResponse) GetRows() []QueryMarginLoanRecordResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []QueryMarginLoanRecordResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginLoanRecordResponse) GetRowsOk() ([]QueryMarginLoanRecordResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *QueryMarginLoanRecordResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []QueryMarginLoanRecordResponseRowsInner and assigns it to the Rows field.
func (o *QueryMarginLoanRecordResponse) SetRows(v []QueryMarginLoanRecordResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryMarginLoanRecordResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginLoanRecordResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryMarginLoanRecordResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QueryMarginLoanRecordResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o QueryMarginLoanRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginLoanRecordResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}
	if !common.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryMarginLoanRecordResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryMarginLoanRecordResponse := _QueryMarginLoanRecordResponse{}

	err = json.Unmarshal(data, &varQueryMarginLoanRecordResponse)

	if err != nil {
		return err
	}

	*o = QueryMarginLoanRecordResponse(varQueryMarginLoanRecordResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryMarginLoanRecordResponse struct {
	value *QueryMarginLoanRecordResponse
	isSet bool
}

func (v NullableQueryMarginLoanRecordResponse) Get() *QueryMarginLoanRecordResponse {
	return v.value
}

func (v *NullableQueryMarginLoanRecordResponse) Set(val *QueryMarginLoanRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginLoanRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginLoanRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginLoanRecordResponse(val *QueryMarginLoanRecordResponse) *NullableQueryMarginLoanRecordResponse {
	return &NullableQueryMarginLoanRecordResponse{value: val, isSet: true}
}

func (v NullableQueryMarginLoanRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginLoanRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
