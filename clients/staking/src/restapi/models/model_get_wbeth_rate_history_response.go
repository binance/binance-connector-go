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

// checks if the GetWbethRateHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetWbethRateHistoryResponse{}

// GetWbethRateHistoryResponse struct for GetWbethRateHistoryResponse
type GetWbethRateHistoryResponse struct {
	Rows                 []GetWbethRateHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *string                                `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetWbethRateHistoryResponse GetWbethRateHistoryResponse

// NewGetWbethRateHistoryResponse instantiates a new GetWbethRateHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWbethRateHistoryResponse() *GetWbethRateHistoryResponse {
	this := GetWbethRateHistoryResponse{}
	return &this
}

// NewGetWbethRateHistoryResponseWithDefaults instantiates a new GetWbethRateHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWbethRateHistoryResponseWithDefaults() *GetWbethRateHistoryResponse {
	this := GetWbethRateHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetWbethRateHistoryResponse) GetRows() []GetWbethRateHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetWbethRateHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethRateHistoryResponse) GetRowsOk() ([]GetWbethRateHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetWbethRateHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetWbethRateHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetWbethRateHistoryResponse) SetRows(v []GetWbethRateHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetWbethRateHistoryResponse) GetTotal() string {
	if o == nil || common.IsNil(o.Total) {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethRateHistoryResponse) GetTotalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetWbethRateHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *GetWbethRateHistoryResponse) SetTotal(v string) {
	o.Total = &v
}

func (o GetWbethRateHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWbethRateHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetWbethRateHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetWbethRateHistoryResponse := _GetWbethRateHistoryResponse{}

	err = json.Unmarshal(data, &varGetWbethRateHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetWbethRateHistoryResponse(varGetWbethRateHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetWbethRateHistoryResponse struct {
	value *GetWbethRateHistoryResponse
	isSet bool
}

func (v NullableGetWbethRateHistoryResponse) Get() *GetWbethRateHistoryResponse {
	return v.value
}

func (v *NullableGetWbethRateHistoryResponse) Set(val *GetWbethRateHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWbethRateHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWbethRateHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWbethRateHistoryResponse(val *GetWbethRateHistoryResponse) *NullableGetWbethRateHistoryResponse {
	return &NullableGetWbethRateHistoryResponse{value: val, isSet: true}
}

func (v NullableGetWbethRateHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWbethRateHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
