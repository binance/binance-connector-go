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

// checks if the GetFlexibleLoanRepaymentHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleLoanRepaymentHistoryResponse{}

// GetFlexibleLoanRepaymentHistoryResponse struct for GetFlexibleLoanRepaymentHistoryResponse
type GetFlexibleLoanRepaymentHistoryResponse struct {
	Rows                 []GetFlexibleLoanRepaymentHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                             `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleLoanRepaymentHistoryResponse GetFlexibleLoanRepaymentHistoryResponse

// NewGetFlexibleLoanRepaymentHistoryResponse instantiates a new GetFlexibleLoanRepaymentHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleLoanRepaymentHistoryResponse() *GetFlexibleLoanRepaymentHistoryResponse {
	this := GetFlexibleLoanRepaymentHistoryResponse{}
	return &this
}

// NewGetFlexibleLoanRepaymentHistoryResponseWithDefaults instantiates a new GetFlexibleLoanRepaymentHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleLoanRepaymentHistoryResponseWithDefaults() *GetFlexibleLoanRepaymentHistoryResponse {
	this := GetFlexibleLoanRepaymentHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetFlexibleLoanRepaymentHistoryResponse) GetRows() []GetFlexibleLoanRepaymentHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetFlexibleLoanRepaymentHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanRepaymentHistoryResponse) GetRowsOk() ([]GetFlexibleLoanRepaymentHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetFlexibleLoanRepaymentHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetFlexibleLoanRepaymentHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetFlexibleLoanRepaymentHistoryResponse) SetRows(v []GetFlexibleLoanRepaymentHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetFlexibleLoanRepaymentHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanRepaymentHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetFlexibleLoanRepaymentHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetFlexibleLoanRepaymentHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetFlexibleLoanRepaymentHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleLoanRepaymentHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetFlexibleLoanRepaymentHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleLoanRepaymentHistoryResponse := _GetFlexibleLoanRepaymentHistoryResponse{}

	err = json.Unmarshal(data, &varGetFlexibleLoanRepaymentHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetFlexibleLoanRepaymentHistoryResponse(varGetFlexibleLoanRepaymentHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleLoanRepaymentHistoryResponse struct {
	value *GetFlexibleLoanRepaymentHistoryResponse
	isSet bool
}

func (v NullableGetFlexibleLoanRepaymentHistoryResponse) Get() *GetFlexibleLoanRepaymentHistoryResponse {
	return v.value
}

func (v *NullableGetFlexibleLoanRepaymentHistoryResponse) Set(val *GetFlexibleLoanRepaymentHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleLoanRepaymentHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleLoanRepaymentHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleLoanRepaymentHistoryResponse(val *GetFlexibleLoanRepaymentHistoryResponse) *NullableGetFlexibleLoanRepaymentHistoryResponse {
	return &NullableGetFlexibleLoanRepaymentHistoryResponse{value: val, isSet: true}
}

func (v NullableGetFlexibleLoanRepaymentHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleLoanRepaymentHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
