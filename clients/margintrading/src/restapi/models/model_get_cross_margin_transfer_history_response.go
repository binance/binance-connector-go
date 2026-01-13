/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetCrossMarginTransferHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCrossMarginTransferHistoryResponse{}

// GetCrossMarginTransferHistoryResponse struct for GetCrossMarginTransferHistoryResponse
type GetCrossMarginTransferHistoryResponse struct {
	Rows                 []GetCrossMarginTransferHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                           `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCrossMarginTransferHistoryResponse GetCrossMarginTransferHistoryResponse

// NewGetCrossMarginTransferHistoryResponse instantiates a new GetCrossMarginTransferHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCrossMarginTransferHistoryResponse() *GetCrossMarginTransferHistoryResponse {
	this := GetCrossMarginTransferHistoryResponse{}
	return &this
}

// NewGetCrossMarginTransferHistoryResponseWithDefaults instantiates a new GetCrossMarginTransferHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCrossMarginTransferHistoryResponseWithDefaults() *GetCrossMarginTransferHistoryResponse {
	this := GetCrossMarginTransferHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetCrossMarginTransferHistoryResponse) GetRows() []GetCrossMarginTransferHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetCrossMarginTransferHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCrossMarginTransferHistoryResponse) GetRowsOk() ([]GetCrossMarginTransferHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetCrossMarginTransferHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetCrossMarginTransferHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetCrossMarginTransferHistoryResponse) SetRows(v []GetCrossMarginTransferHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetCrossMarginTransferHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCrossMarginTransferHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetCrossMarginTransferHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetCrossMarginTransferHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetCrossMarginTransferHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCrossMarginTransferHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetCrossMarginTransferHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetCrossMarginTransferHistoryResponse := _GetCrossMarginTransferHistoryResponse{}

	err = json.Unmarshal(data, &varGetCrossMarginTransferHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetCrossMarginTransferHistoryResponse(varGetCrossMarginTransferHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCrossMarginTransferHistoryResponse struct {
	value *GetCrossMarginTransferHistoryResponse
	isSet bool
}

func (v NullableGetCrossMarginTransferHistoryResponse) Get() *GetCrossMarginTransferHistoryResponse {
	return v.value
}

func (v *NullableGetCrossMarginTransferHistoryResponse) Set(val *GetCrossMarginTransferHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCrossMarginTransferHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCrossMarginTransferHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCrossMarginTransferHistoryResponse(val *GetCrossMarginTransferHistoryResponse) *NullableGetCrossMarginTransferHistoryResponse {
	return &NullableGetCrossMarginTransferHistoryResponse{value: val, isSet: true}
}

func (v NullableGetCrossMarginTransferHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCrossMarginTransferHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
