/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSolStakingHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSolStakingHistoryResponse{}

// GetSolStakingHistoryResponse struct for GetSolStakingHistoryResponse
type GetSolStakingHistoryResponse struct {
	Rows                 []GetSolStakingHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                  `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSolStakingHistoryResponse GetSolStakingHistoryResponse

// NewGetSolStakingHistoryResponse instantiates a new GetSolStakingHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolStakingHistoryResponse() *GetSolStakingHistoryResponse {
	this := GetSolStakingHistoryResponse{}
	return &this
}

// NewGetSolStakingHistoryResponseWithDefaults instantiates a new GetSolStakingHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolStakingHistoryResponseWithDefaults() *GetSolStakingHistoryResponse {
	this := GetSolStakingHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetSolStakingHistoryResponse) GetRows() []GetSolStakingHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetSolStakingHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingHistoryResponse) GetRowsOk() ([]GetSolStakingHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetSolStakingHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetSolStakingHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetSolStakingHistoryResponse) SetRows(v []GetSolStakingHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetSolStakingHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetSolStakingHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetSolStakingHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetSolStakingHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSolStakingHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetSolStakingHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetSolStakingHistoryResponse := _GetSolStakingHistoryResponse{}

	err = json.Unmarshal(data, &varGetSolStakingHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetSolStakingHistoryResponse(varGetSolStakingHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSolStakingHistoryResponse struct {
	value *GetSolStakingHistoryResponse
	isSet bool
}

func (v NullableGetSolStakingHistoryResponse) Get() *GetSolStakingHistoryResponse {
	return v.value
}

func (v *NullableGetSolStakingHistoryResponse) Set(val *GetSolStakingHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolStakingHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolStakingHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolStakingHistoryResponse(val *GetSolStakingHistoryResponse) *NullableGetSolStakingHistoryResponse {
	return &NullableGetSolStakingHistoryResponse{value: val, isSet: true}
}

func (v NullableGetSolStakingHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSolStakingHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
