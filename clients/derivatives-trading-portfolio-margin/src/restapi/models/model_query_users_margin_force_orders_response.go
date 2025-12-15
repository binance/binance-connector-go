/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryUsersMarginForceOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUsersMarginForceOrdersResponse{}

// QueryUsersMarginForceOrdersResponse struct for QueryUsersMarginForceOrdersResponse
type QueryUsersMarginForceOrdersResponse struct {
	Rows                 []QueryUsersMarginForceOrdersResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                         `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryUsersMarginForceOrdersResponse QueryUsersMarginForceOrdersResponse

// NewQueryUsersMarginForceOrdersResponse instantiates a new QueryUsersMarginForceOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUsersMarginForceOrdersResponse() *QueryUsersMarginForceOrdersResponse {
	this := QueryUsersMarginForceOrdersResponse{}
	return &this
}

// NewQueryUsersMarginForceOrdersResponseWithDefaults instantiates a new QueryUsersMarginForceOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUsersMarginForceOrdersResponseWithDefaults() *QueryUsersMarginForceOrdersResponse {
	this := QueryUsersMarginForceOrdersResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *QueryUsersMarginForceOrdersResponse) GetRows() []QueryUsersMarginForceOrdersResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []QueryUsersMarginForceOrdersResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUsersMarginForceOrdersResponse) GetRowsOk() ([]QueryUsersMarginForceOrdersResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *QueryUsersMarginForceOrdersResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []QueryUsersMarginForceOrdersResponseRowsInner and assigns it to the Rows field.
func (o *QueryUsersMarginForceOrdersResponse) SetRows(v []QueryUsersMarginForceOrdersResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryUsersMarginForceOrdersResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUsersMarginForceOrdersResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryUsersMarginForceOrdersResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QueryUsersMarginForceOrdersResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o QueryUsersMarginForceOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUsersMarginForceOrdersResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryUsersMarginForceOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryUsersMarginForceOrdersResponse := _QueryUsersMarginForceOrdersResponse{}

	err = json.Unmarshal(data, &varQueryUsersMarginForceOrdersResponse)

	if err != nil {
		return err
	}

	*o = QueryUsersMarginForceOrdersResponse(varQueryUsersMarginForceOrdersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryUsersMarginForceOrdersResponse struct {
	value *QueryUsersMarginForceOrdersResponse
	isSet bool
}

func (v NullableQueryUsersMarginForceOrdersResponse) Get() *QueryUsersMarginForceOrdersResponse {
	return v.value
}

func (v *NullableQueryUsersMarginForceOrdersResponse) Set(val *QueryUsersMarginForceOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUsersMarginForceOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUsersMarginForceOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUsersMarginForceOrdersResponse(val *QueryUsersMarginForceOrdersResponse) *NullableQueryUsersMarginForceOrdersResponse {
	return &NullableQueryUsersMarginForceOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryUsersMarginForceOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUsersMarginForceOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
