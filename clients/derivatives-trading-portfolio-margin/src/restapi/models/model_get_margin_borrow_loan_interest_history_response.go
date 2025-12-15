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

// checks if the GetMarginBorrowLoanInterestHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetMarginBorrowLoanInterestHistoryResponse{}

// GetMarginBorrowLoanInterestHistoryResponse struct for GetMarginBorrowLoanInterestHistoryResponse
type GetMarginBorrowLoanInterestHistoryResponse struct {
	Rows                 []GetMarginBorrowLoanInterestHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                                `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetMarginBorrowLoanInterestHistoryResponse GetMarginBorrowLoanInterestHistoryResponse

// NewGetMarginBorrowLoanInterestHistoryResponse instantiates a new GetMarginBorrowLoanInterestHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarginBorrowLoanInterestHistoryResponse() *GetMarginBorrowLoanInterestHistoryResponse {
	this := GetMarginBorrowLoanInterestHistoryResponse{}
	return &this
}

// NewGetMarginBorrowLoanInterestHistoryResponseWithDefaults instantiates a new GetMarginBorrowLoanInterestHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarginBorrowLoanInterestHistoryResponseWithDefaults() *GetMarginBorrowLoanInterestHistoryResponse {
	this := GetMarginBorrowLoanInterestHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetMarginBorrowLoanInterestHistoryResponse) GetRows() []GetMarginBorrowLoanInterestHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetMarginBorrowLoanInterestHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponse) GetRowsOk() ([]GetMarginBorrowLoanInterestHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetMarginBorrowLoanInterestHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetMarginBorrowLoanInterestHistoryResponse) SetRows(v []GetMarginBorrowLoanInterestHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetMarginBorrowLoanInterestHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetMarginBorrowLoanInterestHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetMarginBorrowLoanInterestHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetMarginBorrowLoanInterestHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarginBorrowLoanInterestHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetMarginBorrowLoanInterestHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetMarginBorrowLoanInterestHistoryResponse := _GetMarginBorrowLoanInterestHistoryResponse{}

	err = json.Unmarshal(data, &varGetMarginBorrowLoanInterestHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetMarginBorrowLoanInterestHistoryResponse(varGetMarginBorrowLoanInterestHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMarginBorrowLoanInterestHistoryResponse struct {
	value *GetMarginBorrowLoanInterestHistoryResponse
	isSet bool
}

func (v NullableGetMarginBorrowLoanInterestHistoryResponse) Get() *GetMarginBorrowLoanInterestHistoryResponse {
	return v.value
}

func (v *NullableGetMarginBorrowLoanInterestHistoryResponse) Set(val *GetMarginBorrowLoanInterestHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarginBorrowLoanInterestHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarginBorrowLoanInterestHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarginBorrowLoanInterestHistoryResponse(val *GetMarginBorrowLoanInterestHistoryResponse) *NullableGetMarginBorrowLoanInterestHistoryResponse {
	return &NullableGetMarginBorrowLoanInterestHistoryResponse{value: val, isSet: true}
}

func (v NullableGetMarginBorrowLoanInterestHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarginBorrowLoanInterestHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
