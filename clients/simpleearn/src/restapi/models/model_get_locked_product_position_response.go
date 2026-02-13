/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetLockedProductPositionResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLockedProductPositionResponse{}

// GetLockedProductPositionResponse struct for GetLockedProductPositionResponse
type GetLockedProductPositionResponse struct {
	Rows                 []GetLockedProductPositionResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                      `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLockedProductPositionResponse GetLockedProductPositionResponse

// NewGetLockedProductPositionResponse instantiates a new GetLockedProductPositionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLockedProductPositionResponse() *GetLockedProductPositionResponse {
	this := GetLockedProductPositionResponse{}
	return &this
}

// NewGetLockedProductPositionResponseWithDefaults instantiates a new GetLockedProductPositionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLockedProductPositionResponseWithDefaults() *GetLockedProductPositionResponse {
	this := GetLockedProductPositionResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponse) GetRows() []GetLockedProductPositionResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetLockedProductPositionResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponse) GetRowsOk() ([]GetLockedProductPositionResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetLockedProductPositionResponseRowsInner and assigns it to the Rows field.
func (o *GetLockedProductPositionResponse) SetRows(v []GetLockedProductPositionResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetLockedProductPositionResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetLockedProductPositionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLockedProductPositionResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetLockedProductPositionResponse) UnmarshalJSON(data []byte) (err error) {
	varGetLockedProductPositionResponse := _GetLockedProductPositionResponse{}

	err = json.Unmarshal(data, &varGetLockedProductPositionResponse)

	if err != nil {
		return err
	}

	*o = GetLockedProductPositionResponse(varGetLockedProductPositionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLockedProductPositionResponse struct {
	value *GetLockedProductPositionResponse
	isSet bool
}

func (v NullableGetLockedProductPositionResponse) Get() *GetLockedProductPositionResponse {
	return v.value
}

func (v *NullableGetLockedProductPositionResponse) Set(val *GetLockedProductPositionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLockedProductPositionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLockedProductPositionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLockedProductPositionResponse(val *GetLockedProductPositionResponse) *NullableGetLockedProductPositionResponse {
	return &NullableGetLockedProductPositionResponse{value: val, isSet: true}
}

func (v NullableGetLockedProductPositionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLockedProductPositionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
