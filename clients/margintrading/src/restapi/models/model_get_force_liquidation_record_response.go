/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetForceLiquidationRecordResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetForceLiquidationRecordResponse{}

// GetForceLiquidationRecordResponse struct for GetForceLiquidationRecordResponse
type GetForceLiquidationRecordResponse struct {
	Rows                 []GetForceLiquidationRecordResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                       `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetForceLiquidationRecordResponse GetForceLiquidationRecordResponse

// NewGetForceLiquidationRecordResponse instantiates a new GetForceLiquidationRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetForceLiquidationRecordResponse() *GetForceLiquidationRecordResponse {
	this := GetForceLiquidationRecordResponse{}
	return &this
}

// NewGetForceLiquidationRecordResponseWithDefaults instantiates a new GetForceLiquidationRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetForceLiquidationRecordResponseWithDefaults() *GetForceLiquidationRecordResponse {
	this := GetForceLiquidationRecordResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetForceLiquidationRecordResponse) GetRows() []GetForceLiquidationRecordResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetForceLiquidationRecordResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetForceLiquidationRecordResponse) GetRowsOk() ([]GetForceLiquidationRecordResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetForceLiquidationRecordResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetForceLiquidationRecordResponseRowsInner and assigns it to the Rows field.
func (o *GetForceLiquidationRecordResponse) SetRows(v []GetForceLiquidationRecordResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetForceLiquidationRecordResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetForceLiquidationRecordResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetForceLiquidationRecordResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetForceLiquidationRecordResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetForceLiquidationRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetForceLiquidationRecordResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetForceLiquidationRecordResponse) UnmarshalJSON(data []byte) (err error) {
	varGetForceLiquidationRecordResponse := _GetForceLiquidationRecordResponse{}

	err = json.Unmarshal(data, &varGetForceLiquidationRecordResponse)

	if err != nil {
		return err
	}

	*o = GetForceLiquidationRecordResponse(varGetForceLiquidationRecordResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetForceLiquidationRecordResponse struct {
	value *GetForceLiquidationRecordResponse
	isSet bool
}

func (v NullableGetForceLiquidationRecordResponse) Get() *GetForceLiquidationRecordResponse {
	return v.value
}

func (v *NullableGetForceLiquidationRecordResponse) Set(val *GetForceLiquidationRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetForceLiquidationRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetForceLiquidationRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetForceLiquidationRecordResponse(val *GetForceLiquidationRecordResponse) *NullableGetForceLiquidationRecordResponse {
	return &NullableGetForceLiquidationRecordResponse{value: val, isSet: true}
}

func (v NullableGetForceLiquidationRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetForceLiquidationRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
