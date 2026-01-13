/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetVIPLoanInterestRateHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetVIPLoanInterestRateHistoryResponse{}

// GetVIPLoanInterestRateHistoryResponse struct for GetVIPLoanInterestRateHistoryResponse
type GetVIPLoanInterestRateHistoryResponse struct {
	Rows                 []GetVIPLoanInterestRateHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                           `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetVIPLoanInterestRateHistoryResponse GetVIPLoanInterestRateHistoryResponse

// NewGetVIPLoanInterestRateHistoryResponse instantiates a new GetVIPLoanInterestRateHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVIPLoanInterestRateHistoryResponse() *GetVIPLoanInterestRateHistoryResponse {
	this := GetVIPLoanInterestRateHistoryResponse{}
	return &this
}

// NewGetVIPLoanInterestRateHistoryResponseWithDefaults instantiates a new GetVIPLoanInterestRateHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVIPLoanInterestRateHistoryResponseWithDefaults() *GetVIPLoanInterestRateHistoryResponse {
	this := GetVIPLoanInterestRateHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetVIPLoanInterestRateHistoryResponse) GetRows() []GetVIPLoanInterestRateHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetVIPLoanInterestRateHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanInterestRateHistoryResponse) GetRowsOk() ([]GetVIPLoanInterestRateHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetVIPLoanInterestRateHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetVIPLoanInterestRateHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetVIPLoanInterestRateHistoryResponse) SetRows(v []GetVIPLoanInterestRateHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetVIPLoanInterestRateHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanInterestRateHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetVIPLoanInterestRateHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetVIPLoanInterestRateHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetVIPLoanInterestRateHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVIPLoanInterestRateHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetVIPLoanInterestRateHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetVIPLoanInterestRateHistoryResponse := _GetVIPLoanInterestRateHistoryResponse{}

	err = json.Unmarshal(data, &varGetVIPLoanInterestRateHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetVIPLoanInterestRateHistoryResponse(varGetVIPLoanInterestRateHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetVIPLoanInterestRateHistoryResponse struct {
	value *GetVIPLoanInterestRateHistoryResponse
	isSet bool
}

func (v NullableGetVIPLoanInterestRateHistoryResponse) Get() *GetVIPLoanInterestRateHistoryResponse {
	return v.value
}

func (v *NullableGetVIPLoanInterestRateHistoryResponse) Set(val *GetVIPLoanInterestRateHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVIPLoanInterestRateHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVIPLoanInterestRateHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVIPLoanInterestRateHistoryResponse(val *GetVIPLoanInterestRateHistoryResponse) *NullableGetVIPLoanInterestRateHistoryResponse {
	return &NullableGetVIPLoanInterestRateHistoryResponse{value: val, isSet: true}
}

func (v NullableGetVIPLoanInterestRateHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVIPLoanInterestRateHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
