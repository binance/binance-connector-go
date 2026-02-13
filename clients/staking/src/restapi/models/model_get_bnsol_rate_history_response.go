/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetBnsolRateHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBnsolRateHistoryResponse{}

// GetBnsolRateHistoryResponse struct for GetBnsolRateHistoryResponse
type GetBnsolRateHistoryResponse struct {
	Rows                 []GetBnsolRateHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *string                                `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBnsolRateHistoryResponse GetBnsolRateHistoryResponse

// NewGetBnsolRateHistoryResponse instantiates a new GetBnsolRateHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBnsolRateHistoryResponse() *GetBnsolRateHistoryResponse {
	this := GetBnsolRateHistoryResponse{}
	return &this
}

// NewGetBnsolRateHistoryResponseWithDefaults instantiates a new GetBnsolRateHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBnsolRateHistoryResponseWithDefaults() *GetBnsolRateHistoryResponse {
	this := GetBnsolRateHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetBnsolRateHistoryResponse) GetRows() []GetBnsolRateHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetBnsolRateHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnsolRateHistoryResponse) GetRowsOk() ([]GetBnsolRateHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetBnsolRateHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetBnsolRateHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetBnsolRateHistoryResponse) SetRows(v []GetBnsolRateHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetBnsolRateHistoryResponse) GetTotal() string {
	if o == nil || common.IsNil(o.Total) {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnsolRateHistoryResponse) GetTotalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetBnsolRateHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *GetBnsolRateHistoryResponse) SetTotal(v string) {
	o.Total = &v
}

func (o GetBnsolRateHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBnsolRateHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetBnsolRateHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetBnsolRateHistoryResponse := _GetBnsolRateHistoryResponse{}

	err = json.Unmarshal(data, &varGetBnsolRateHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetBnsolRateHistoryResponse(varGetBnsolRateHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBnsolRateHistoryResponse struct {
	value *GetBnsolRateHistoryResponse
	isSet bool
}

func (v NullableGetBnsolRateHistoryResponse) Get() *GetBnsolRateHistoryResponse {
	return v.value
}

func (v *NullableGetBnsolRateHistoryResponse) Set(val *GetBnsolRateHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBnsolRateHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBnsolRateHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBnsolRateHistoryResponse(val *GetBnsolRateHistoryResponse) *NullableGetBnsolRateHistoryResponse {
	return &NullableGetBnsolRateHistoryResponse{value: val, isSet: true}
}

func (v NullableGetBnsolRateHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBnsolRateHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
