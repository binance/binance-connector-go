/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFlexibleLoanLtvAdjustmentHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleLoanLtvAdjustmentHistoryResponse{}

// GetFlexibleLoanLtvAdjustmentHistoryResponse struct for GetFlexibleLoanLtvAdjustmentHistoryResponse
type GetFlexibleLoanLtvAdjustmentHistoryResponse struct {
	Rows                 []GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                                 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleLoanLtvAdjustmentHistoryResponse GetFlexibleLoanLtvAdjustmentHistoryResponse

// NewGetFlexibleLoanLtvAdjustmentHistoryResponse instantiates a new GetFlexibleLoanLtvAdjustmentHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleLoanLtvAdjustmentHistoryResponse() *GetFlexibleLoanLtvAdjustmentHistoryResponse {
	this := GetFlexibleLoanLtvAdjustmentHistoryResponse{}
	return &this
}

// NewGetFlexibleLoanLtvAdjustmentHistoryResponseWithDefaults instantiates a new GetFlexibleLoanLtvAdjustmentHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleLoanLtvAdjustmentHistoryResponseWithDefaults() *GetFlexibleLoanLtvAdjustmentHistoryResponse {
	this := GetFlexibleLoanLtvAdjustmentHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) GetRows() []GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) GetRowsOk() ([]GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) SetRows(v []GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetFlexibleLoanLtvAdjustmentHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleLoanLtvAdjustmentHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleLoanLtvAdjustmentHistoryResponse := _GetFlexibleLoanLtvAdjustmentHistoryResponse{}

	err = json.Unmarshal(data, &varGetFlexibleLoanLtvAdjustmentHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetFlexibleLoanLtvAdjustmentHistoryResponse(varGetFlexibleLoanLtvAdjustmentHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleLoanLtvAdjustmentHistoryResponse struct {
	value *GetFlexibleLoanLtvAdjustmentHistoryResponse
	isSet bool
}

func (v NullableGetFlexibleLoanLtvAdjustmentHistoryResponse) Get() *GetFlexibleLoanLtvAdjustmentHistoryResponse {
	return v.value
}

func (v *NullableGetFlexibleLoanLtvAdjustmentHistoryResponse) Set(val *GetFlexibleLoanLtvAdjustmentHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleLoanLtvAdjustmentHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleLoanLtvAdjustmentHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleLoanLtvAdjustmentHistoryResponse(val *GetFlexibleLoanLtvAdjustmentHistoryResponse) *NullableGetFlexibleLoanLtvAdjustmentHistoryResponse {
	return &NullableGetFlexibleLoanLtvAdjustmentHistoryResponse{value: val, isSet: true}
}

func (v NullableGetFlexibleLoanLtvAdjustmentHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleLoanLtvAdjustmentHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
