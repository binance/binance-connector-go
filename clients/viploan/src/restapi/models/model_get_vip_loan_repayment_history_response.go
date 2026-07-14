/*
VIP Loan REST API

Access over-collateralized loan services, manage positions, and monitor collateral via the VIP Loan API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetVIPLoanRepaymentHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetVIPLoanRepaymentHistoryResponse{}

// GetVIPLoanRepaymentHistoryResponse struct for GetVIPLoanRepaymentHistoryResponse
type GetVIPLoanRepaymentHistoryResponse struct {
	Rows                 []GetVIPLoanRepaymentHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                        `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetVIPLoanRepaymentHistoryResponse GetVIPLoanRepaymentHistoryResponse

// NewGetVIPLoanRepaymentHistoryResponse instantiates a new GetVIPLoanRepaymentHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVIPLoanRepaymentHistoryResponse() *GetVIPLoanRepaymentHistoryResponse {
	this := GetVIPLoanRepaymentHistoryResponse{}
	return &this
}

// NewGetVIPLoanRepaymentHistoryResponseWithDefaults instantiates a new GetVIPLoanRepaymentHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVIPLoanRepaymentHistoryResponseWithDefaults() *GetVIPLoanRepaymentHistoryResponse {
	this := GetVIPLoanRepaymentHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetVIPLoanRepaymentHistoryResponse) GetRows() []GetVIPLoanRepaymentHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetVIPLoanRepaymentHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanRepaymentHistoryResponse) GetRowsOk() ([]GetVIPLoanRepaymentHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetVIPLoanRepaymentHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetVIPLoanRepaymentHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetVIPLoanRepaymentHistoryResponse) SetRows(v []GetVIPLoanRepaymentHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetVIPLoanRepaymentHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanRepaymentHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetVIPLoanRepaymentHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetVIPLoanRepaymentHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetVIPLoanRepaymentHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVIPLoanRepaymentHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetVIPLoanRepaymentHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetVIPLoanRepaymentHistoryResponse := _GetVIPLoanRepaymentHistoryResponse{}

	err = json.Unmarshal(data, &varGetVIPLoanRepaymentHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetVIPLoanRepaymentHistoryResponse(varGetVIPLoanRepaymentHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetVIPLoanRepaymentHistoryResponse struct {
	value *GetVIPLoanRepaymentHistoryResponse
	isSet bool
}

func (v NullableGetVIPLoanRepaymentHistoryResponse) Get() *GetVIPLoanRepaymentHistoryResponse {
	return v.value
}

func (v *NullableGetVIPLoanRepaymentHistoryResponse) Set(val *GetVIPLoanRepaymentHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVIPLoanRepaymentHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVIPLoanRepaymentHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVIPLoanRepaymentHistoryResponse(val *GetVIPLoanRepaymentHistoryResponse) *NullableGetVIPLoanRepaymentHistoryResponse {
	return &NullableGetVIPLoanRepaymentHistoryResponse{value: val, isSet: true}
}

func (v NullableGetVIPLoanRepaymentHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVIPLoanRepaymentHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
