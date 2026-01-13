/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetRateHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetRateHistoryResponse{}

// GetRateHistoryResponse struct for GetRateHistoryResponse
type GetRateHistoryResponse struct {
	Rows                 []GetRateHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *string                           `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRateHistoryResponse GetRateHistoryResponse

// NewGetRateHistoryResponse instantiates a new GetRateHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRateHistoryResponse() *GetRateHistoryResponse {
	this := GetRateHistoryResponse{}
	return &this
}

// NewGetRateHistoryResponseWithDefaults instantiates a new GetRateHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRateHistoryResponseWithDefaults() *GetRateHistoryResponse {
	this := GetRateHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetRateHistoryResponse) GetRows() []GetRateHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetRateHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRateHistoryResponse) GetRowsOk() ([]GetRateHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetRateHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetRateHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetRateHistoryResponse) SetRows(v []GetRateHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetRateHistoryResponse) GetTotal() string {
	if o == nil || common.IsNil(o.Total) {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRateHistoryResponse) GetTotalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetRateHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *GetRateHistoryResponse) SetTotal(v string) {
	o.Total = &v
}

func (o GetRateHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRateHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetRateHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetRateHistoryResponse := _GetRateHistoryResponse{}

	err = json.Unmarshal(data, &varGetRateHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetRateHistoryResponse(varGetRateHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRateHistoryResponse struct {
	value *GetRateHistoryResponse
	isSet bool
}

func (v NullableGetRateHistoryResponse) Get() *GetRateHistoryResponse {
	return v.value
}

func (v *NullableGetRateHistoryResponse) Set(val *GetRateHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRateHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRateHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRateHistoryResponse(val *GetRateHistoryResponse) *NullableGetRateHistoryResponse {
	return &NullableGetRateHistoryResponse{value: val, isSet: true}
}

func (v NullableGetRateHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRateHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
