/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetInterestHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetInterestHistoryResponse{}

// GetInterestHistoryResponse struct for GetInterestHistoryResponse
type GetInterestHistoryResponse struct {
	Rows                 []GetInterestHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetInterestHistoryResponse GetInterestHistoryResponse

// NewGetInterestHistoryResponse instantiates a new GetInterestHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInterestHistoryResponse() *GetInterestHistoryResponse {
	this := GetInterestHistoryResponse{}
	return &this
}

// NewGetInterestHistoryResponseWithDefaults instantiates a new GetInterestHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInterestHistoryResponseWithDefaults() *GetInterestHistoryResponse {
	this := GetInterestHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetInterestHistoryResponse) GetRows() []GetInterestHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetInterestHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterestHistoryResponse) GetRowsOk() ([]GetInterestHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetInterestHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetInterestHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetInterestHistoryResponse) SetRows(v []GetInterestHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetInterestHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterestHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetInterestHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetInterestHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetInterestHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInterestHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetInterestHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetInterestHistoryResponse := _GetInterestHistoryResponse{}

	err = json.Unmarshal(data, &varGetInterestHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetInterestHistoryResponse(varGetInterestHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetInterestHistoryResponse struct {
	value *GetInterestHistoryResponse
	isSet bool
}

func (v NullableGetInterestHistoryResponse) Get() *GetInterestHistoryResponse {
	return v.value
}

func (v *NullableGetInterestHistoryResponse) Set(val *GetInterestHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInterestHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInterestHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInterestHistoryResponse(val *GetInterestHistoryResponse) *NullableGetInterestHistoryResponse {
	return &NullableGetInterestHistoryResponse{value: val, isSet: true}
}

func (v NullableGetInterestHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInterestHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
