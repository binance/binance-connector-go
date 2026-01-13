/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetRwusdRateHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetRwusdRateHistoryResponse{}

// GetRwusdRateHistoryResponse struct for GetRwusdRateHistoryResponse
type GetRwusdRateHistoryResponse struct {
	Rows                 []GetBfusdRateHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *string                                `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRwusdRateHistoryResponse GetRwusdRateHistoryResponse

// NewGetRwusdRateHistoryResponse instantiates a new GetRwusdRateHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRwusdRateHistoryResponse() *GetRwusdRateHistoryResponse {
	this := GetRwusdRateHistoryResponse{}
	return &this
}

// NewGetRwusdRateHistoryResponseWithDefaults instantiates a new GetRwusdRateHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRwusdRateHistoryResponseWithDefaults() *GetRwusdRateHistoryResponse {
	this := GetRwusdRateHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetRwusdRateHistoryResponse) GetRows() []GetBfusdRateHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetBfusdRateHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdRateHistoryResponse) GetRowsOk() ([]GetBfusdRateHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetRwusdRateHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetBfusdRateHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetRwusdRateHistoryResponse) SetRows(v []GetBfusdRateHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetRwusdRateHistoryResponse) GetTotal() string {
	if o == nil || common.IsNil(o.Total) {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdRateHistoryResponse) GetTotalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetRwusdRateHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *GetRwusdRateHistoryResponse) SetTotal(v string) {
	o.Total = &v
}

func (o GetRwusdRateHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRwusdRateHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetRwusdRateHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetRwusdRateHistoryResponse := _GetRwusdRateHistoryResponse{}

	err = json.Unmarshal(data, &varGetRwusdRateHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetRwusdRateHistoryResponse(varGetRwusdRateHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRwusdRateHistoryResponse struct {
	value *GetRwusdRateHistoryResponse
	isSet bool
}

func (v NullableGetRwusdRateHistoryResponse) Get() *GetRwusdRateHistoryResponse {
	return v.value
}

func (v *NullableGetRwusdRateHistoryResponse) Set(val *GetRwusdRateHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRwusdRateHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRwusdRateHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRwusdRateHistoryResponse(val *GetRwusdRateHistoryResponse) *NullableGetRwusdRateHistoryResponse {
	return &NullableGetRwusdRateHistoryResponse{value: val, isSet: true}
}

func (v NullableGetRwusdRateHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRwusdRateHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
