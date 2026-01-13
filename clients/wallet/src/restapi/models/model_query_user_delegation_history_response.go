/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryUserDelegationHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUserDelegationHistoryResponse{}

// QueryUserDelegationHistoryResponse struct for QueryUserDelegationHistoryResponse
type QueryUserDelegationHistoryResponse struct {
	Total                *int64                                        `json:"total,omitempty"`
	Rows                 []QueryUserDelegationHistoryResponseRowsInner `json:"rows,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryUserDelegationHistoryResponse QueryUserDelegationHistoryResponse

// NewQueryUserDelegationHistoryResponse instantiates a new QueryUserDelegationHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUserDelegationHistoryResponse() *QueryUserDelegationHistoryResponse {
	this := QueryUserDelegationHistoryResponse{}
	return &this
}

// NewQueryUserDelegationHistoryResponseWithDefaults instantiates a new QueryUserDelegationHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUserDelegationHistoryResponseWithDefaults() *QueryUserDelegationHistoryResponse {
	this := QueryUserDelegationHistoryResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryUserDelegationHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserDelegationHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryUserDelegationHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QueryUserDelegationHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *QueryUserDelegationHistoryResponse) GetRows() []QueryUserDelegationHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []QueryUserDelegationHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserDelegationHistoryResponse) GetRowsOk() ([]QueryUserDelegationHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *QueryUserDelegationHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []QueryUserDelegationHistoryResponseRowsInner and assigns it to the Rows field.
func (o *QueryUserDelegationHistoryResponse) SetRows(v []QueryUserDelegationHistoryResponseRowsInner) {
	o.Rows = v
}

func (o QueryUserDelegationHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUserDelegationHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryUserDelegationHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryUserDelegationHistoryResponse := _QueryUserDelegationHistoryResponse{}

	err = json.Unmarshal(data, &varQueryUserDelegationHistoryResponse)

	if err != nil {
		return err
	}

	*o = QueryUserDelegationHistoryResponse(varQueryUserDelegationHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "rows")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryUserDelegationHistoryResponse struct {
	value *QueryUserDelegationHistoryResponse
	isSet bool
}

func (v NullableQueryUserDelegationHistoryResponse) Get() *QueryUserDelegationHistoryResponse {
	return v.value
}

func (v *NullableQueryUserDelegationHistoryResponse) Set(val *QueryUserDelegationHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUserDelegationHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUserDelegationHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUserDelegationHistoryResponse(val *QueryUserDelegationHistoryResponse) *NullableQueryUserDelegationHistoryResponse {
	return &NullableQueryUserDelegationHistoryResponse{value: val, isSet: true}
}

func (v NullableQueryUserDelegationHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUserDelegationHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
