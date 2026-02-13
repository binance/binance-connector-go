/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFlexibleLoanBorrowHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleLoanBorrowHistoryResponse{}

// GetFlexibleLoanBorrowHistoryResponse struct for GetFlexibleLoanBorrowHistoryResponse
type GetFlexibleLoanBorrowHistoryResponse struct {
	Rows                 []GetFlexibleLoanBorrowHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                          `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleLoanBorrowHistoryResponse GetFlexibleLoanBorrowHistoryResponse

// NewGetFlexibleLoanBorrowHistoryResponse instantiates a new GetFlexibleLoanBorrowHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleLoanBorrowHistoryResponse() *GetFlexibleLoanBorrowHistoryResponse {
	this := GetFlexibleLoanBorrowHistoryResponse{}
	return &this
}

// NewGetFlexibleLoanBorrowHistoryResponseWithDefaults instantiates a new GetFlexibleLoanBorrowHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleLoanBorrowHistoryResponseWithDefaults() *GetFlexibleLoanBorrowHistoryResponse {
	this := GetFlexibleLoanBorrowHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetFlexibleLoanBorrowHistoryResponse) GetRows() []GetFlexibleLoanBorrowHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetFlexibleLoanBorrowHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanBorrowHistoryResponse) GetRowsOk() ([]GetFlexibleLoanBorrowHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetFlexibleLoanBorrowHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetFlexibleLoanBorrowHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetFlexibleLoanBorrowHistoryResponse) SetRows(v []GetFlexibleLoanBorrowHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetFlexibleLoanBorrowHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanBorrowHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetFlexibleLoanBorrowHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetFlexibleLoanBorrowHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetFlexibleLoanBorrowHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleLoanBorrowHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetFlexibleLoanBorrowHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleLoanBorrowHistoryResponse := _GetFlexibleLoanBorrowHistoryResponse{}

	err = json.Unmarshal(data, &varGetFlexibleLoanBorrowHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetFlexibleLoanBorrowHistoryResponse(varGetFlexibleLoanBorrowHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleLoanBorrowHistoryResponse struct {
	value *GetFlexibleLoanBorrowHistoryResponse
	isSet bool
}

func (v NullableGetFlexibleLoanBorrowHistoryResponse) Get() *GetFlexibleLoanBorrowHistoryResponse {
	return v.value
}

func (v *NullableGetFlexibleLoanBorrowHistoryResponse) Set(val *GetFlexibleLoanBorrowHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleLoanBorrowHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleLoanBorrowHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleLoanBorrowHistoryResponse(val *GetFlexibleLoanBorrowHistoryResponse) *NullableGetFlexibleLoanBorrowHistoryResponse {
	return &NullableGetFlexibleLoanBorrowHistoryResponse{value: val, isSet: true}
}

func (v NullableGetFlexibleLoanBorrowHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleLoanBorrowHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
