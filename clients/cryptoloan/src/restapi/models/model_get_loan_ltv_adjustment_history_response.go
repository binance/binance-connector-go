/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetLoanLtvAdjustmentHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLoanLtvAdjustmentHistoryResponse{}

// GetLoanLtvAdjustmentHistoryResponse struct for GetLoanLtvAdjustmentHistoryResponse
type GetLoanLtvAdjustmentHistoryResponse struct {
	Rows                 []GetLoanLtvAdjustmentHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                         `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLoanLtvAdjustmentHistoryResponse GetLoanLtvAdjustmentHistoryResponse

// NewGetLoanLtvAdjustmentHistoryResponse instantiates a new GetLoanLtvAdjustmentHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLoanLtvAdjustmentHistoryResponse() *GetLoanLtvAdjustmentHistoryResponse {
	this := GetLoanLtvAdjustmentHistoryResponse{}
	return &this
}

// NewGetLoanLtvAdjustmentHistoryResponseWithDefaults instantiates a new GetLoanLtvAdjustmentHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLoanLtvAdjustmentHistoryResponseWithDefaults() *GetLoanLtvAdjustmentHistoryResponse {
	this := GetLoanLtvAdjustmentHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetLoanLtvAdjustmentHistoryResponse) GetRows() []GetLoanLtvAdjustmentHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetLoanLtvAdjustmentHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanLtvAdjustmentHistoryResponse) GetRowsOk() ([]GetLoanLtvAdjustmentHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetLoanLtvAdjustmentHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetLoanLtvAdjustmentHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetLoanLtvAdjustmentHistoryResponse) SetRows(v []GetLoanLtvAdjustmentHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetLoanLtvAdjustmentHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanLtvAdjustmentHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetLoanLtvAdjustmentHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetLoanLtvAdjustmentHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetLoanLtvAdjustmentHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLoanLtvAdjustmentHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetLoanLtvAdjustmentHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetLoanLtvAdjustmentHistoryResponse := _GetLoanLtvAdjustmentHistoryResponse{}

	err = json.Unmarshal(data, &varGetLoanLtvAdjustmentHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetLoanLtvAdjustmentHistoryResponse(varGetLoanLtvAdjustmentHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLoanLtvAdjustmentHistoryResponse struct {
	value *GetLoanLtvAdjustmentHistoryResponse
	isSet bool
}

func (v NullableGetLoanLtvAdjustmentHistoryResponse) Get() *GetLoanLtvAdjustmentHistoryResponse {
	return v.value
}

func (v *NullableGetLoanLtvAdjustmentHistoryResponse) Set(val *GetLoanLtvAdjustmentHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLoanLtvAdjustmentHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLoanLtvAdjustmentHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLoanLtvAdjustmentHistoryResponse(val *GetLoanLtvAdjustmentHistoryResponse) *NullableGetLoanLtvAdjustmentHistoryResponse {
	return &NullableGetLoanLtvAdjustmentHistoryResponse{value: val, isSet: true}
}

func (v NullableGetLoanLtvAdjustmentHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLoanLtvAdjustmentHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
