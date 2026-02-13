/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetEthStakingHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetEthStakingHistoryResponse{}

// GetEthStakingHistoryResponse struct for GetEthStakingHistoryResponse
type GetEthStakingHistoryResponse struct {
	Rows                 []GetEthStakingHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                  `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetEthStakingHistoryResponse GetEthStakingHistoryResponse

// NewGetEthStakingHistoryResponse instantiates a new GetEthStakingHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEthStakingHistoryResponse() *GetEthStakingHistoryResponse {
	this := GetEthStakingHistoryResponse{}
	return &this
}

// NewGetEthStakingHistoryResponseWithDefaults instantiates a new GetEthStakingHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEthStakingHistoryResponseWithDefaults() *GetEthStakingHistoryResponse {
	this := GetEthStakingHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetEthStakingHistoryResponse) GetRows() []GetEthStakingHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetEthStakingHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthStakingHistoryResponse) GetRowsOk() ([]GetEthStakingHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetEthStakingHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetEthStakingHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetEthStakingHistoryResponse) SetRows(v []GetEthStakingHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetEthStakingHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthStakingHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetEthStakingHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetEthStakingHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetEthStakingHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEthStakingHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetEthStakingHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetEthStakingHistoryResponse := _GetEthStakingHistoryResponse{}

	err = json.Unmarshal(data, &varGetEthStakingHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetEthStakingHistoryResponse(varGetEthStakingHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetEthStakingHistoryResponse struct {
	value *GetEthStakingHistoryResponse
	isSet bool
}

func (v NullableGetEthStakingHistoryResponse) Get() *GetEthStakingHistoryResponse {
	return v.value
}

func (v *NullableGetEthStakingHistoryResponse) Set(val *GetEthStakingHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEthStakingHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEthStakingHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEthStakingHistoryResponse(val *GetEthStakingHistoryResponse) *NullableGetEthStakingHistoryResponse {
	return &NullableGetEthStakingHistoryResponse{value: val, isSet: true}
}

func (v NullableGetEthStakingHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEthStakingHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
