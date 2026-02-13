/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFlexibleProductPositionResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleProductPositionResponse{}

// GetFlexibleProductPositionResponse struct for GetFlexibleProductPositionResponse
type GetFlexibleProductPositionResponse struct {
	Rows                 []GetFlexibleProductPositionResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                        `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleProductPositionResponse GetFlexibleProductPositionResponse

// NewGetFlexibleProductPositionResponse instantiates a new GetFlexibleProductPositionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleProductPositionResponse() *GetFlexibleProductPositionResponse {
	this := GetFlexibleProductPositionResponse{}
	return &this
}

// NewGetFlexibleProductPositionResponseWithDefaults instantiates a new GetFlexibleProductPositionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleProductPositionResponseWithDefaults() *GetFlexibleProductPositionResponse {
	this := GetFlexibleProductPositionResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponse) GetRows() []GetFlexibleProductPositionResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetFlexibleProductPositionResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponse) GetRowsOk() ([]GetFlexibleProductPositionResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetFlexibleProductPositionResponseRowsInner and assigns it to the Rows field.
func (o *GetFlexibleProductPositionResponse) SetRows(v []GetFlexibleProductPositionResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetFlexibleProductPositionResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetFlexibleProductPositionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleProductPositionResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetFlexibleProductPositionResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleProductPositionResponse := _GetFlexibleProductPositionResponse{}

	err = json.Unmarshal(data, &varGetFlexibleProductPositionResponse)

	if err != nil {
		return err
	}

	*o = GetFlexibleProductPositionResponse(varGetFlexibleProductPositionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleProductPositionResponse struct {
	value *GetFlexibleProductPositionResponse
	isSet bool
}

func (v NullableGetFlexibleProductPositionResponse) Get() *GetFlexibleProductPositionResponse {
	return v.value
}

func (v *NullableGetFlexibleProductPositionResponse) Set(val *GetFlexibleProductPositionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleProductPositionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleProductPositionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleProductPositionResponse(val *GetFlexibleProductPositionResponse) *NullableGetFlexibleProductPositionResponse {
	return &NullableGetFlexibleProductPositionResponse{value: val, isSet: true}
}

func (v NullableGetFlexibleProductPositionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleProductPositionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
