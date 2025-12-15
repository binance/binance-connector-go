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

// checks if the GetSoftStakingRewardsHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSoftStakingRewardsHistoryResponse{}

// GetSoftStakingRewardsHistoryResponse struct for GetSoftStakingRewardsHistoryResponse
type GetSoftStakingRewardsHistoryResponse struct {
	Rows                 []GetSoftStakingRewardsHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                          `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSoftStakingRewardsHistoryResponse GetSoftStakingRewardsHistoryResponse

// NewGetSoftStakingRewardsHistoryResponse instantiates a new GetSoftStakingRewardsHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSoftStakingRewardsHistoryResponse() *GetSoftStakingRewardsHistoryResponse {
	this := GetSoftStakingRewardsHistoryResponse{}
	return &this
}

// NewGetSoftStakingRewardsHistoryResponseWithDefaults instantiates a new GetSoftStakingRewardsHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSoftStakingRewardsHistoryResponseWithDefaults() *GetSoftStakingRewardsHistoryResponse {
	this := GetSoftStakingRewardsHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetSoftStakingRewardsHistoryResponse) GetRows() []GetSoftStakingRewardsHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetSoftStakingRewardsHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingRewardsHistoryResponse) GetRowsOk() ([]GetSoftStakingRewardsHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetSoftStakingRewardsHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetSoftStakingRewardsHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetSoftStakingRewardsHistoryResponse) SetRows(v []GetSoftStakingRewardsHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetSoftStakingRewardsHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingRewardsHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetSoftStakingRewardsHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetSoftStakingRewardsHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetSoftStakingRewardsHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSoftStakingRewardsHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetSoftStakingRewardsHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetSoftStakingRewardsHistoryResponse := _GetSoftStakingRewardsHistoryResponse{}

	err = json.Unmarshal(data, &varGetSoftStakingRewardsHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetSoftStakingRewardsHistoryResponse(varGetSoftStakingRewardsHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSoftStakingRewardsHistoryResponse struct {
	value *GetSoftStakingRewardsHistoryResponse
	isSet bool
}

func (v NullableGetSoftStakingRewardsHistoryResponse) Get() *GetSoftStakingRewardsHistoryResponse {
	return v.value
}

func (v *NullableGetSoftStakingRewardsHistoryResponse) Set(val *GetSoftStakingRewardsHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSoftStakingRewardsHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSoftStakingRewardsHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSoftStakingRewardsHistoryResponse(val *GetSoftStakingRewardsHistoryResponse) *NullableGetSoftStakingRewardsHistoryResponse {
	return &NullableGetSoftStakingRewardsHistoryResponse{value: val, isSet: true}
}

func (v NullableGetSoftStakingRewardsHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSoftStakingRewardsHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
