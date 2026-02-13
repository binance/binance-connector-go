/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetBoostRewardsHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBoostRewardsHistoryResponse{}

// GetBoostRewardsHistoryResponse struct for GetBoostRewardsHistoryResponse
type GetBoostRewardsHistoryResponse struct {
	Rows                 []GetBoostRewardsHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                    `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBoostRewardsHistoryResponse GetBoostRewardsHistoryResponse

// NewGetBoostRewardsHistoryResponse instantiates a new GetBoostRewardsHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBoostRewardsHistoryResponse() *GetBoostRewardsHistoryResponse {
	this := GetBoostRewardsHistoryResponse{}
	return &this
}

// NewGetBoostRewardsHistoryResponseWithDefaults instantiates a new GetBoostRewardsHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBoostRewardsHistoryResponseWithDefaults() *GetBoostRewardsHistoryResponse {
	this := GetBoostRewardsHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetBoostRewardsHistoryResponse) GetRows() []GetBoostRewardsHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetBoostRewardsHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBoostRewardsHistoryResponse) GetRowsOk() ([]GetBoostRewardsHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetBoostRewardsHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetBoostRewardsHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetBoostRewardsHistoryResponse) SetRows(v []GetBoostRewardsHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetBoostRewardsHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBoostRewardsHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetBoostRewardsHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetBoostRewardsHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetBoostRewardsHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBoostRewardsHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetBoostRewardsHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetBoostRewardsHistoryResponse := _GetBoostRewardsHistoryResponse{}

	err = json.Unmarshal(data, &varGetBoostRewardsHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetBoostRewardsHistoryResponse(varGetBoostRewardsHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBoostRewardsHistoryResponse struct {
	value *GetBoostRewardsHistoryResponse
	isSet bool
}

func (v NullableGetBoostRewardsHistoryResponse) Get() *GetBoostRewardsHistoryResponse {
	return v.value
}

func (v *NullableGetBoostRewardsHistoryResponse) Set(val *GetBoostRewardsHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBoostRewardsHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBoostRewardsHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBoostRewardsHistoryResponse(val *GetBoostRewardsHistoryResponse) *NullableGetBoostRewardsHistoryResponse {
	return &NullableGetBoostRewardsHistoryResponse{value: val, isSet: true}
}

func (v NullableGetBoostRewardsHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBoostRewardsHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
