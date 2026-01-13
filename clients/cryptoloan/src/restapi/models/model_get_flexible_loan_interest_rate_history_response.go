/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFlexibleLoanInterestRateHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleLoanInterestRateHistoryResponse{}

// GetFlexibleLoanInterestRateHistoryResponse struct for GetFlexibleLoanInterestRateHistoryResponse
type GetFlexibleLoanInterestRateHistoryResponse struct {
	Rows                 []GetFlexibleLoanInterestRateHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                                `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleLoanInterestRateHistoryResponse GetFlexibleLoanInterestRateHistoryResponse

// NewGetFlexibleLoanInterestRateHistoryResponse instantiates a new GetFlexibleLoanInterestRateHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleLoanInterestRateHistoryResponse() *GetFlexibleLoanInterestRateHistoryResponse {
	this := GetFlexibleLoanInterestRateHistoryResponse{}
	return &this
}

// NewGetFlexibleLoanInterestRateHistoryResponseWithDefaults instantiates a new GetFlexibleLoanInterestRateHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleLoanInterestRateHistoryResponseWithDefaults() *GetFlexibleLoanInterestRateHistoryResponse {
	this := GetFlexibleLoanInterestRateHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetFlexibleLoanInterestRateHistoryResponse) GetRows() []GetFlexibleLoanInterestRateHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetFlexibleLoanInterestRateHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanInterestRateHistoryResponse) GetRowsOk() ([]GetFlexibleLoanInterestRateHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetFlexibleLoanInterestRateHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetFlexibleLoanInterestRateHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetFlexibleLoanInterestRateHistoryResponse) SetRows(v []GetFlexibleLoanInterestRateHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetFlexibleLoanInterestRateHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanInterestRateHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetFlexibleLoanInterestRateHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetFlexibleLoanInterestRateHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetFlexibleLoanInterestRateHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleLoanInterestRateHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetFlexibleLoanInterestRateHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleLoanInterestRateHistoryResponse := _GetFlexibleLoanInterestRateHistoryResponse{}

	err = json.Unmarshal(data, &varGetFlexibleLoanInterestRateHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetFlexibleLoanInterestRateHistoryResponse(varGetFlexibleLoanInterestRateHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleLoanInterestRateHistoryResponse struct {
	value *GetFlexibleLoanInterestRateHistoryResponse
	isSet bool
}

func (v NullableGetFlexibleLoanInterestRateHistoryResponse) Get() *GetFlexibleLoanInterestRateHistoryResponse {
	return v.value
}

func (v *NullableGetFlexibleLoanInterestRateHistoryResponse) Set(val *GetFlexibleLoanInterestRateHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleLoanInterestRateHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleLoanInterestRateHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleLoanInterestRateHistoryResponse(val *GetFlexibleLoanInterestRateHistoryResponse) *NullableGetFlexibleLoanInterestRateHistoryResponse {
	return &NullableGetFlexibleLoanInterestRateHistoryResponse{value: val, isSet: true}
}

func (v NullableGetFlexibleLoanInterestRateHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleLoanInterestRateHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
