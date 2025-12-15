/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetWbethUnwrapHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetWbethUnwrapHistoryResponse{}

// GetWbethUnwrapHistoryResponse struct for GetWbethUnwrapHistoryResponse
type GetWbethUnwrapHistoryResponse struct {
	Rows                 []GetWbethUnwrapHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetWbethUnwrapHistoryResponse GetWbethUnwrapHistoryResponse

// NewGetWbethUnwrapHistoryResponse instantiates a new GetWbethUnwrapHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWbethUnwrapHistoryResponse() *GetWbethUnwrapHistoryResponse {
	this := GetWbethUnwrapHistoryResponse{}
	return &this
}

// NewGetWbethUnwrapHistoryResponseWithDefaults instantiates a new GetWbethUnwrapHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWbethUnwrapHistoryResponseWithDefaults() *GetWbethUnwrapHistoryResponse {
	this := GetWbethUnwrapHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetWbethUnwrapHistoryResponse) GetRows() []GetWbethUnwrapHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetWbethUnwrapHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethUnwrapHistoryResponse) GetRowsOk() ([]GetWbethUnwrapHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetWbethUnwrapHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetWbethUnwrapHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetWbethUnwrapHistoryResponse) SetRows(v []GetWbethUnwrapHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetWbethUnwrapHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethUnwrapHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetWbethUnwrapHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetWbethUnwrapHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetWbethUnwrapHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWbethUnwrapHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetWbethUnwrapHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetWbethUnwrapHistoryResponse := _GetWbethUnwrapHistoryResponse{}

	err = json.Unmarshal(data, &varGetWbethUnwrapHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetWbethUnwrapHistoryResponse(varGetWbethUnwrapHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetWbethUnwrapHistoryResponse struct {
	value *GetWbethUnwrapHistoryResponse
	isSet bool
}

func (v NullableGetWbethUnwrapHistoryResponse) Get() *GetWbethUnwrapHistoryResponse {
	return v.value
}

func (v *NullableGetWbethUnwrapHistoryResponse) Set(val *GetWbethUnwrapHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWbethUnwrapHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWbethUnwrapHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWbethUnwrapHistoryResponse(val *GetWbethUnwrapHistoryResponse) *NullableGetWbethUnwrapHistoryResponse {
	return &NullableGetWbethUnwrapHistoryResponse{value: val, isSet: true}
}

func (v NullableGetWbethUnwrapHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWbethUnwrapHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
