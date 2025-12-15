/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetLockedRewardsHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLockedRewardsHistoryResponse{}

// GetLockedRewardsHistoryResponse struct for GetLockedRewardsHistoryResponse
type GetLockedRewardsHistoryResponse struct {
	Rows                 []GetLockedRewardsHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                     `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLockedRewardsHistoryResponse GetLockedRewardsHistoryResponse

// NewGetLockedRewardsHistoryResponse instantiates a new GetLockedRewardsHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLockedRewardsHistoryResponse() *GetLockedRewardsHistoryResponse {
	this := GetLockedRewardsHistoryResponse{}
	return &this
}

// NewGetLockedRewardsHistoryResponseWithDefaults instantiates a new GetLockedRewardsHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLockedRewardsHistoryResponseWithDefaults() *GetLockedRewardsHistoryResponse {
	this := GetLockedRewardsHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetLockedRewardsHistoryResponse) GetRows() []GetLockedRewardsHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetLockedRewardsHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRewardsHistoryResponse) GetRowsOk() ([]GetLockedRewardsHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetLockedRewardsHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetLockedRewardsHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetLockedRewardsHistoryResponse) SetRows(v []GetLockedRewardsHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetLockedRewardsHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRewardsHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetLockedRewardsHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetLockedRewardsHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetLockedRewardsHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLockedRewardsHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetLockedRewardsHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetLockedRewardsHistoryResponse := _GetLockedRewardsHistoryResponse{}

	err = json.Unmarshal(data, &varGetLockedRewardsHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetLockedRewardsHistoryResponse(varGetLockedRewardsHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLockedRewardsHistoryResponse struct {
	value *GetLockedRewardsHistoryResponse
	isSet bool
}

func (v NullableGetLockedRewardsHistoryResponse) Get() *GetLockedRewardsHistoryResponse {
	return v.value
}

func (v *NullableGetLockedRewardsHistoryResponse) Set(val *GetLockedRewardsHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLockedRewardsHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLockedRewardsHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLockedRewardsHistoryResponse(val *GetLockedRewardsHistoryResponse) *NullableGetLockedRewardsHistoryResponse {
	return &NullableGetLockedRewardsHistoryResponse{value: val, isSet: true}
}

func (v NullableGetLockedRewardsHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLockedRewardsHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
