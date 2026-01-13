/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryMarginRepayRecordResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginRepayRecordResponse{}

// QueryMarginRepayRecordResponse struct for QueryMarginRepayRecordResponse
type QueryMarginRepayRecordResponse struct {
	Rows                 []QueryMarginRepayRecordResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                    `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMarginRepayRecordResponse QueryMarginRepayRecordResponse

// NewQueryMarginRepayRecordResponse instantiates a new QueryMarginRepayRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginRepayRecordResponse() *QueryMarginRepayRecordResponse {
	this := QueryMarginRepayRecordResponse{}
	return &this
}

// NewQueryMarginRepayRecordResponseWithDefaults instantiates a new QueryMarginRepayRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginRepayRecordResponseWithDefaults() *QueryMarginRepayRecordResponse {
	this := QueryMarginRepayRecordResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *QueryMarginRepayRecordResponse) GetRows() []QueryMarginRepayRecordResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []QueryMarginRepayRecordResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginRepayRecordResponse) GetRowsOk() ([]QueryMarginRepayRecordResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *QueryMarginRepayRecordResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []QueryMarginRepayRecordResponseRowsInner and assigns it to the Rows field.
func (o *QueryMarginRepayRecordResponse) SetRows(v []QueryMarginRepayRecordResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryMarginRepayRecordResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginRepayRecordResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryMarginRepayRecordResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QueryMarginRepayRecordResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o QueryMarginRepayRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginRepayRecordResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryMarginRepayRecordResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryMarginRepayRecordResponse := _QueryMarginRepayRecordResponse{}

	err = json.Unmarshal(data, &varQueryMarginRepayRecordResponse)

	if err != nil {
		return err
	}

	*o = QueryMarginRepayRecordResponse(varQueryMarginRepayRecordResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryMarginRepayRecordResponse struct {
	value *QueryMarginRepayRecordResponse
	isSet bool
}

func (v NullableQueryMarginRepayRecordResponse) Get() *QueryMarginRepayRecordResponse {
	return v.value
}

func (v *NullableQueryMarginRepayRecordResponse) Set(val *QueryMarginRepayRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginRepayRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginRepayRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginRepayRecordResponse(val *QueryMarginRepayRecordResponse) *NullableQueryMarginRepayRecordResponse {
	return &NullableQueryMarginRepayRecordResponse{value: val, isSet: true}
}

func (v NullableQueryMarginRepayRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginRepayRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
