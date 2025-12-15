/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetLoanBorrowHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLoanBorrowHistoryResponse{}

// GetLoanBorrowHistoryResponse struct for GetLoanBorrowHistoryResponse
type GetLoanBorrowHistoryResponse struct {
	Rows                 []GetLoanBorrowHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                  `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLoanBorrowHistoryResponse GetLoanBorrowHistoryResponse

// NewGetLoanBorrowHistoryResponse instantiates a new GetLoanBorrowHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLoanBorrowHistoryResponse() *GetLoanBorrowHistoryResponse {
	this := GetLoanBorrowHistoryResponse{}
	return &this
}

// NewGetLoanBorrowHistoryResponseWithDefaults instantiates a new GetLoanBorrowHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLoanBorrowHistoryResponseWithDefaults() *GetLoanBorrowHistoryResponse {
	this := GetLoanBorrowHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetLoanBorrowHistoryResponse) GetRows() []GetLoanBorrowHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetLoanBorrowHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanBorrowHistoryResponse) GetRowsOk() ([]GetLoanBorrowHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetLoanBorrowHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetLoanBorrowHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetLoanBorrowHistoryResponse) SetRows(v []GetLoanBorrowHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetLoanBorrowHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanBorrowHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetLoanBorrowHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetLoanBorrowHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetLoanBorrowHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLoanBorrowHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetLoanBorrowHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetLoanBorrowHistoryResponse := _GetLoanBorrowHistoryResponse{}

	err = json.Unmarshal(data, &varGetLoanBorrowHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetLoanBorrowHistoryResponse(varGetLoanBorrowHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLoanBorrowHistoryResponse struct {
	value *GetLoanBorrowHistoryResponse
	isSet bool
}

func (v NullableGetLoanBorrowHistoryResponse) Get() *GetLoanBorrowHistoryResponse {
	return v.value
}

func (v *NullableGetLoanBorrowHistoryResponse) Set(val *GetLoanBorrowHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLoanBorrowHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLoanBorrowHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLoanBorrowHistoryResponse(val *GetLoanBorrowHistoryResponse) *NullableGetLoanBorrowHistoryResponse {
	return &NullableGetLoanBorrowHistoryResponse{value: val, isSet: true}
}

func (v NullableGetLoanBorrowHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLoanBorrowHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
