/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryLiquidationLoanRepayHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryLiquidationLoanRepayHistoryResponse{}

// QueryLiquidationLoanRepayHistoryResponse struct for QueryLiquidationLoanRepayHistoryResponse
type QueryLiquidationLoanRepayHistoryResponse struct {
	// Total number of repayment records
	Total                *int64                                              `json:"total,omitempty"`
	Rows                 []QueryLiquidationLoanRepayHistoryResponseRowsInner `json:"rows,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryLiquidationLoanRepayHistoryResponse QueryLiquidationLoanRepayHistoryResponse

// NewQueryLiquidationLoanRepayHistoryResponse instantiates a new QueryLiquidationLoanRepayHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryLiquidationLoanRepayHistoryResponse() *QueryLiquidationLoanRepayHistoryResponse {
	this := QueryLiquidationLoanRepayHistoryResponse{}
	return &this
}

// NewQueryLiquidationLoanRepayHistoryResponseWithDefaults instantiates a new QueryLiquidationLoanRepayHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryLiquidationLoanRepayHistoryResponseWithDefaults() *QueryLiquidationLoanRepayHistoryResponse {
	this := QueryLiquidationLoanRepayHistoryResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryLiquidationLoanRepayHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiquidationLoanRepayHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryLiquidationLoanRepayHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QueryLiquidationLoanRepayHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *QueryLiquidationLoanRepayHistoryResponse) GetRows() []QueryLiquidationLoanRepayHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []QueryLiquidationLoanRepayHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiquidationLoanRepayHistoryResponse) GetRowsOk() ([]QueryLiquidationLoanRepayHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *QueryLiquidationLoanRepayHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []QueryLiquidationLoanRepayHistoryResponseRowsInner and assigns it to the Rows field.
func (o *QueryLiquidationLoanRepayHistoryResponse) SetRows(v []QueryLiquidationLoanRepayHistoryResponseRowsInner) {
	o.Rows = v
}

func (o QueryLiquidationLoanRepayHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryLiquidationLoanRepayHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryLiquidationLoanRepayHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryLiquidationLoanRepayHistoryResponse := _QueryLiquidationLoanRepayHistoryResponse{}

	err = json.Unmarshal(data, &varQueryLiquidationLoanRepayHistoryResponse)

	if err != nil {
		return err
	}

	*o = QueryLiquidationLoanRepayHistoryResponse(varQueryLiquidationLoanRepayHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "rows")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryLiquidationLoanRepayHistoryResponse struct {
	value *QueryLiquidationLoanRepayHistoryResponse
	isSet bool
}

func (v NullableQueryLiquidationLoanRepayHistoryResponse) Get() *QueryLiquidationLoanRepayHistoryResponse {
	return v.value
}

func (v *NullableQueryLiquidationLoanRepayHistoryResponse) Set(val *QueryLiquidationLoanRepayHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryLiquidationLoanRepayHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryLiquidationLoanRepayHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryLiquidationLoanRepayHistoryResponse(val *QueryLiquidationLoanRepayHistoryResponse) *NullableQueryLiquidationLoanRepayHistoryResponse {
	return &NullableQueryLiquidationLoanRepayHistoryResponse{value: val, isSet: true}
}

func (v NullableQueryLiquidationLoanRepayHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryLiquidationLoanRepayHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
