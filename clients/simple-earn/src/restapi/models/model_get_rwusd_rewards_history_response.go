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

// checks if the GetRwusdRewardsHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetRwusdRewardsHistoryResponse{}

// GetRwusdRewardsHistoryResponse struct for GetRwusdRewardsHistoryResponse
type GetRwusdRewardsHistoryResponse struct {
	Rows                 []GetRwusdRewardsHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                    `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRwusdRewardsHistoryResponse GetRwusdRewardsHistoryResponse

// NewGetRwusdRewardsHistoryResponse instantiates a new GetRwusdRewardsHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRwusdRewardsHistoryResponse() *GetRwusdRewardsHistoryResponse {
	this := GetRwusdRewardsHistoryResponse{}
	return &this
}

// NewGetRwusdRewardsHistoryResponseWithDefaults instantiates a new GetRwusdRewardsHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRwusdRewardsHistoryResponseWithDefaults() *GetRwusdRewardsHistoryResponse {
	this := GetRwusdRewardsHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetRwusdRewardsHistoryResponse) GetRows() []GetRwusdRewardsHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetRwusdRewardsHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdRewardsHistoryResponse) GetRowsOk() ([]GetRwusdRewardsHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetRwusdRewardsHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetRwusdRewardsHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetRwusdRewardsHistoryResponse) SetRows(v []GetRwusdRewardsHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetRwusdRewardsHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdRewardsHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetRwusdRewardsHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetRwusdRewardsHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetRwusdRewardsHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRwusdRewardsHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetRwusdRewardsHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetRwusdRewardsHistoryResponse := _GetRwusdRewardsHistoryResponse{}

	err = json.Unmarshal(data, &varGetRwusdRewardsHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetRwusdRewardsHistoryResponse(varGetRwusdRewardsHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRwusdRewardsHistoryResponse struct {
	value *GetRwusdRewardsHistoryResponse
	isSet bool
}

func (v NullableGetRwusdRewardsHistoryResponse) Get() *GetRwusdRewardsHistoryResponse {
	return v.value
}

func (v *NullableGetRwusdRewardsHistoryResponse) Set(val *GetRwusdRewardsHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRwusdRewardsHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRwusdRewardsHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRwusdRewardsHistoryResponse(val *GetRwusdRewardsHistoryResponse) *NullableGetRwusdRewardsHistoryResponse {
	return &NullableGetRwusdRewardsHistoryResponse{value: val, isSet: true}
}

func (v NullableGetRwusdRewardsHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRwusdRewardsHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
