/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFlexibleLoanLiquidationHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleLoanLiquidationHistoryResponse{}

// GetFlexibleLoanLiquidationHistoryResponse struct for GetFlexibleLoanLiquidationHistoryResponse
type GetFlexibleLoanLiquidationHistoryResponse struct {
	Rows                 []GetFlexibleLoanLiquidationHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                               `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleLoanLiquidationHistoryResponse GetFlexibleLoanLiquidationHistoryResponse

// NewGetFlexibleLoanLiquidationHistoryResponse instantiates a new GetFlexibleLoanLiquidationHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleLoanLiquidationHistoryResponse() *GetFlexibleLoanLiquidationHistoryResponse {
	this := GetFlexibleLoanLiquidationHistoryResponse{}
	return &this
}

// NewGetFlexibleLoanLiquidationHistoryResponseWithDefaults instantiates a new GetFlexibleLoanLiquidationHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleLoanLiquidationHistoryResponseWithDefaults() *GetFlexibleLoanLiquidationHistoryResponse {
	this := GetFlexibleLoanLiquidationHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetFlexibleLoanLiquidationHistoryResponse) GetRows() []GetFlexibleLoanLiquidationHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetFlexibleLoanLiquidationHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponse) GetRowsOk() ([]GetFlexibleLoanLiquidationHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetFlexibleLoanLiquidationHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetFlexibleLoanLiquidationHistoryResponse) SetRows(v []GetFlexibleLoanLiquidationHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetFlexibleLoanLiquidationHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetFlexibleLoanLiquidationHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetFlexibleLoanLiquidationHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetFlexibleLoanLiquidationHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleLoanLiquidationHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetFlexibleLoanLiquidationHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleLoanLiquidationHistoryResponse := _GetFlexibleLoanLiquidationHistoryResponse{}

	err = json.Unmarshal(data, &varGetFlexibleLoanLiquidationHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetFlexibleLoanLiquidationHistoryResponse(varGetFlexibleLoanLiquidationHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleLoanLiquidationHistoryResponse struct {
	value *GetFlexibleLoanLiquidationHistoryResponse
	isSet bool
}

func (v NullableGetFlexibleLoanLiquidationHistoryResponse) Get() *GetFlexibleLoanLiquidationHistoryResponse {
	return v.value
}

func (v *NullableGetFlexibleLoanLiquidationHistoryResponse) Set(val *GetFlexibleLoanLiquidationHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleLoanLiquidationHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleLoanLiquidationHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleLoanLiquidationHistoryResponse(val *GetFlexibleLoanLiquidationHistoryResponse) *NullableGetFlexibleLoanLiquidationHistoryResponse {
	return &NullableGetFlexibleLoanLiquidationHistoryResponse{value: val, isSet: true}
}

func (v NullableGetFlexibleLoanLiquidationHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleLoanLiquidationHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
