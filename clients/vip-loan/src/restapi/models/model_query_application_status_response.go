/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryApplicationStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryApplicationStatusResponse{}

// QueryApplicationStatusResponse struct for QueryApplicationStatusResponse
type QueryApplicationStatusResponse struct {
	Rows                 []QueryApplicationStatusResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                    `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryApplicationStatusResponse QueryApplicationStatusResponse

// NewQueryApplicationStatusResponse instantiates a new QueryApplicationStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryApplicationStatusResponse() *QueryApplicationStatusResponse {
	this := QueryApplicationStatusResponse{}
	return &this
}

// NewQueryApplicationStatusResponseWithDefaults instantiates a new QueryApplicationStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryApplicationStatusResponseWithDefaults() *QueryApplicationStatusResponse {
	this := QueryApplicationStatusResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *QueryApplicationStatusResponse) GetRows() []QueryApplicationStatusResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []QueryApplicationStatusResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryApplicationStatusResponse) GetRowsOk() ([]QueryApplicationStatusResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *QueryApplicationStatusResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []QueryApplicationStatusResponseRowsInner and assigns it to the Rows field.
func (o *QueryApplicationStatusResponse) SetRows(v []QueryApplicationStatusResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryApplicationStatusResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryApplicationStatusResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryApplicationStatusResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QueryApplicationStatusResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o QueryApplicationStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryApplicationStatusResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryApplicationStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryApplicationStatusResponse := _QueryApplicationStatusResponse{}

	err = json.Unmarshal(data, &varQueryApplicationStatusResponse)

	if err != nil {
		return err
	}

	*o = QueryApplicationStatusResponse(varQueryApplicationStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryApplicationStatusResponse struct {
	value *QueryApplicationStatusResponse
	isSet bool
}

func (v NullableQueryApplicationStatusResponse) Get() *QueryApplicationStatusResponse {
	return v.value
}

func (v *NullableQueryApplicationStatusResponse) Set(val *QueryApplicationStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryApplicationStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryApplicationStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryApplicationStatusResponse(val *QueryApplicationStatusResponse) *NullableQueryApplicationStatusResponse {
	return &NullableQueryApplicationStatusResponse{value: val, isSet: true}
}

func (v NullableQueryApplicationStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryApplicationStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
