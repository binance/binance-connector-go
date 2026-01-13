/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetLoanRepaymentHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLoanRepaymentHistoryResponse{}

// GetLoanRepaymentHistoryResponse struct for GetLoanRepaymentHistoryResponse
type GetLoanRepaymentHistoryResponse struct {
	Rows                 []GetLoanRepaymentHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                     `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLoanRepaymentHistoryResponse GetLoanRepaymentHistoryResponse

// NewGetLoanRepaymentHistoryResponse instantiates a new GetLoanRepaymentHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLoanRepaymentHistoryResponse() *GetLoanRepaymentHistoryResponse {
	this := GetLoanRepaymentHistoryResponse{}
	return &this
}

// NewGetLoanRepaymentHistoryResponseWithDefaults instantiates a new GetLoanRepaymentHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLoanRepaymentHistoryResponseWithDefaults() *GetLoanRepaymentHistoryResponse {
	this := GetLoanRepaymentHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetLoanRepaymentHistoryResponse) GetRows() []GetLoanRepaymentHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetLoanRepaymentHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepaymentHistoryResponse) GetRowsOk() ([]GetLoanRepaymentHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetLoanRepaymentHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetLoanRepaymentHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetLoanRepaymentHistoryResponse) SetRows(v []GetLoanRepaymentHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetLoanRepaymentHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepaymentHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetLoanRepaymentHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetLoanRepaymentHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetLoanRepaymentHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLoanRepaymentHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetLoanRepaymentHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetLoanRepaymentHistoryResponse := _GetLoanRepaymentHistoryResponse{}

	err = json.Unmarshal(data, &varGetLoanRepaymentHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetLoanRepaymentHistoryResponse(varGetLoanRepaymentHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLoanRepaymentHistoryResponse struct {
	value *GetLoanRepaymentHistoryResponse
	isSet bool
}

func (v NullableGetLoanRepaymentHistoryResponse) Get() *GetLoanRepaymentHistoryResponse {
	return v.value
}

func (v *NullableGetLoanRepaymentHistoryResponse) Set(val *GetLoanRepaymentHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLoanRepaymentHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLoanRepaymentHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLoanRepaymentHistoryResponse(val *GetLoanRepaymentHistoryResponse) *NullableGetLoanRepaymentHistoryResponse {
	return &NullableGetLoanRepaymentHistoryResponse{value: val, isSet: true}
}

func (v NullableGetLoanRepaymentHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLoanRepaymentHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
