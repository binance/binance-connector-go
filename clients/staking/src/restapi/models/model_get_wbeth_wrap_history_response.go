/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetWbethWrapHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetWbethWrapHistoryResponse{}

// GetWbethWrapHistoryResponse struct for GetWbethWrapHistoryResponse
type GetWbethWrapHistoryResponse struct {
	Rows                 []GetWbethWrapHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetWbethWrapHistoryResponse GetWbethWrapHistoryResponse

// NewGetWbethWrapHistoryResponse instantiates a new GetWbethWrapHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWbethWrapHistoryResponse() *GetWbethWrapHistoryResponse {
	this := GetWbethWrapHistoryResponse{}
	return &this
}

// NewGetWbethWrapHistoryResponseWithDefaults instantiates a new GetWbethWrapHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWbethWrapHistoryResponseWithDefaults() *GetWbethWrapHistoryResponse {
	this := GetWbethWrapHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetWbethWrapHistoryResponse) GetRows() []GetWbethWrapHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetWbethWrapHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethWrapHistoryResponse) GetRowsOk() ([]GetWbethWrapHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetWbethWrapHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetWbethWrapHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetWbethWrapHistoryResponse) SetRows(v []GetWbethWrapHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetWbethWrapHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethWrapHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetWbethWrapHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetWbethWrapHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetWbethWrapHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWbethWrapHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetWbethWrapHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetWbethWrapHistoryResponse := _GetWbethWrapHistoryResponse{}

	err = json.Unmarshal(data, &varGetWbethWrapHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetWbethWrapHistoryResponse(varGetWbethWrapHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetWbethWrapHistoryResponse struct {
	value *GetWbethWrapHistoryResponse
	isSet bool
}

func (v NullableGetWbethWrapHistoryResponse) Get() *GetWbethWrapHistoryResponse {
	return v.value
}

func (v *NullableGetWbethWrapHistoryResponse) Set(val *GetWbethWrapHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWbethWrapHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWbethWrapHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWbethWrapHistoryResponse(val *GetWbethWrapHistoryResponse) *NullableGetWbethWrapHistoryResponse {
	return &NullableGetWbethWrapHistoryResponse{value: val, isSet: true}
}

func (v NullableGetWbethWrapHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWbethWrapHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
