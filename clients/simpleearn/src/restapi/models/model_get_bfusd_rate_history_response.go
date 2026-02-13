/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetBfusdRateHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBfusdRateHistoryResponse{}

// GetBfusdRateHistoryResponse struct for GetBfusdRateHistoryResponse
type GetBfusdRateHistoryResponse struct {
	Rows                 []GetBfusdRateHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *string                                `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBfusdRateHistoryResponse GetBfusdRateHistoryResponse

// NewGetBfusdRateHistoryResponse instantiates a new GetBfusdRateHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBfusdRateHistoryResponse() *GetBfusdRateHistoryResponse {
	this := GetBfusdRateHistoryResponse{}
	return &this
}

// NewGetBfusdRateHistoryResponseWithDefaults instantiates a new GetBfusdRateHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBfusdRateHistoryResponseWithDefaults() *GetBfusdRateHistoryResponse {
	this := GetBfusdRateHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetBfusdRateHistoryResponse) GetRows() []GetBfusdRateHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetBfusdRateHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRateHistoryResponse) GetRowsOk() ([]GetBfusdRateHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetBfusdRateHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetBfusdRateHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetBfusdRateHistoryResponse) SetRows(v []GetBfusdRateHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetBfusdRateHistoryResponse) GetTotal() string {
	if o == nil || common.IsNil(o.Total) {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRateHistoryResponse) GetTotalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetBfusdRateHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *GetBfusdRateHistoryResponse) SetTotal(v string) {
	o.Total = &v
}

func (o GetBfusdRateHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBfusdRateHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetBfusdRateHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetBfusdRateHistoryResponse := _GetBfusdRateHistoryResponse{}

	err = json.Unmarshal(data, &varGetBfusdRateHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetBfusdRateHistoryResponse(varGetBfusdRateHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBfusdRateHistoryResponse struct {
	value *GetBfusdRateHistoryResponse
	isSet bool
}

func (v NullableGetBfusdRateHistoryResponse) Get() *GetBfusdRateHistoryResponse {
	return v.value
}

func (v *NullableGetBfusdRateHistoryResponse) Set(val *GetBfusdRateHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBfusdRateHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBfusdRateHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBfusdRateHistoryResponse(val *GetBfusdRateHistoryResponse) *NullableGetBfusdRateHistoryResponse {
	return &NullableGetBfusdRateHistoryResponse{value: val, isSet: true}
}

func (v NullableGetBfusdRateHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBfusdRateHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
