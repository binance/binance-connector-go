/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse{}

// QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse struct for QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse
type QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse struct {
	Total                *int64                                                               `json:"total,omitempty"`
	Rows                 []QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner `json:"rows,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse

// NewQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse instantiates a new QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse() *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse {
	this := QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse{}
	return &this
}

// NewQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseWithDefaults instantiates a new QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseWithDefaults() *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse {
	this := QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) GetRows() []QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) GetRowsOk() ([]QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner and assigns it to the Rows field.
func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) SetRows(v []QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponseRowsInner) {
	o.Rows = v
}

func (o QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse := _QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse{}

	err = json.Unmarshal(data, &varQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse)

	if err != nil {
		return err
	}

	*o = QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse(varQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "rows")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse struct {
	value *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse
	isSet bool
}

func (v NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) Get() *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse {
	return v.value
}

func (v *NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) Set(val *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse(val *QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) *NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse {
	return &NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse{value: val, isSet: true}
}

func (v NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
