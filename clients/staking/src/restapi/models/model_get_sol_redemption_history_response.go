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

// checks if the GetSolRedemptionHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSolRedemptionHistoryResponse{}

// GetSolRedemptionHistoryResponse struct for GetSolRedemptionHistoryResponse
type GetSolRedemptionHistoryResponse struct {
	Rows                 []GetSolRedemptionHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                     `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSolRedemptionHistoryResponse GetSolRedemptionHistoryResponse

// NewGetSolRedemptionHistoryResponse instantiates a new GetSolRedemptionHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolRedemptionHistoryResponse() *GetSolRedemptionHistoryResponse {
	this := GetSolRedemptionHistoryResponse{}
	return &this
}

// NewGetSolRedemptionHistoryResponseWithDefaults instantiates a new GetSolRedemptionHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolRedemptionHistoryResponseWithDefaults() *GetSolRedemptionHistoryResponse {
	this := GetSolRedemptionHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetSolRedemptionHistoryResponse) GetRows() []GetSolRedemptionHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetSolRedemptionHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolRedemptionHistoryResponse) GetRowsOk() ([]GetSolRedemptionHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetSolRedemptionHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetSolRedemptionHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetSolRedemptionHistoryResponse) SetRows(v []GetSolRedemptionHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetSolRedemptionHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolRedemptionHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetSolRedemptionHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetSolRedemptionHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetSolRedemptionHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSolRedemptionHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetSolRedemptionHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetSolRedemptionHistoryResponse := _GetSolRedemptionHistoryResponse{}

	err = json.Unmarshal(data, &varGetSolRedemptionHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetSolRedemptionHistoryResponse(varGetSolRedemptionHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSolRedemptionHistoryResponse struct {
	value *GetSolRedemptionHistoryResponse
	isSet bool
}

func (v NullableGetSolRedemptionHistoryResponse) Get() *GetSolRedemptionHistoryResponse {
	return v.value
}

func (v *NullableGetSolRedemptionHistoryResponse) Set(val *GetSolRedemptionHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolRedemptionHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolRedemptionHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolRedemptionHistoryResponse(val *GetSolRedemptionHistoryResponse) *NullableGetSolRedemptionHistoryResponse {
	return &NullableGetSolRedemptionHistoryResponse{value: val, isSet: true}
}

func (v NullableGetSolRedemptionHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSolRedemptionHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
