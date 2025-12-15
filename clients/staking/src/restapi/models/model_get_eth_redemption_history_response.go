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

// checks if the GetEthRedemptionHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetEthRedemptionHistoryResponse{}

// GetEthRedemptionHistoryResponse struct for GetEthRedemptionHistoryResponse
type GetEthRedemptionHistoryResponse struct {
	Rows                 []GetEthRedemptionHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                     `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetEthRedemptionHistoryResponse GetEthRedemptionHistoryResponse

// NewGetEthRedemptionHistoryResponse instantiates a new GetEthRedemptionHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEthRedemptionHistoryResponse() *GetEthRedemptionHistoryResponse {
	this := GetEthRedemptionHistoryResponse{}
	return &this
}

// NewGetEthRedemptionHistoryResponseWithDefaults instantiates a new GetEthRedemptionHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEthRedemptionHistoryResponseWithDefaults() *GetEthRedemptionHistoryResponse {
	this := GetEthRedemptionHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetEthRedemptionHistoryResponse) GetRows() []GetEthRedemptionHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetEthRedemptionHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthRedemptionHistoryResponse) GetRowsOk() ([]GetEthRedemptionHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetEthRedemptionHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetEthRedemptionHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetEthRedemptionHistoryResponse) SetRows(v []GetEthRedemptionHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetEthRedemptionHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthRedemptionHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetEthRedemptionHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetEthRedemptionHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetEthRedemptionHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEthRedemptionHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetEthRedemptionHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetEthRedemptionHistoryResponse := _GetEthRedemptionHistoryResponse{}

	err = json.Unmarshal(data, &varGetEthRedemptionHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetEthRedemptionHistoryResponse(varGetEthRedemptionHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetEthRedemptionHistoryResponse struct {
	value *GetEthRedemptionHistoryResponse
	isSet bool
}

func (v NullableGetEthRedemptionHistoryResponse) Get() *GetEthRedemptionHistoryResponse {
	return v.value
}

func (v *NullableGetEthRedemptionHistoryResponse) Set(val *GetEthRedemptionHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEthRedemptionHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEthRedemptionHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEthRedemptionHistoryResponse(val *GetEthRedemptionHistoryResponse) *NullableGetEthRedemptionHistoryResponse {
	return &NullableGetEthRedemptionHistoryResponse{value: val, isSet: true}
}

func (v NullableGetEthRedemptionHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEthRedemptionHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
