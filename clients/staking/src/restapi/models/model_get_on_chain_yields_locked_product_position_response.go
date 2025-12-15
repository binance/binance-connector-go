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

// checks if the GetOnChainYieldsLockedProductPositionResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOnChainYieldsLockedProductPositionResponse{}

// GetOnChainYieldsLockedProductPositionResponse struct for GetOnChainYieldsLockedProductPositionResponse
type GetOnChainYieldsLockedProductPositionResponse struct {
	Rows                 []GetOnChainYieldsLockedProductPositionResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOnChainYieldsLockedProductPositionResponse GetOnChainYieldsLockedProductPositionResponse

// NewGetOnChainYieldsLockedProductPositionResponse instantiates a new GetOnChainYieldsLockedProductPositionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOnChainYieldsLockedProductPositionResponse() *GetOnChainYieldsLockedProductPositionResponse {
	this := GetOnChainYieldsLockedProductPositionResponse{}
	return &this
}

// NewGetOnChainYieldsLockedProductPositionResponseWithDefaults instantiates a new GetOnChainYieldsLockedProductPositionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOnChainYieldsLockedProductPositionResponseWithDefaults() *GetOnChainYieldsLockedProductPositionResponse {
	this := GetOnChainYieldsLockedProductPositionResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponse) GetRows() []GetOnChainYieldsLockedProductPositionResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetOnChainYieldsLockedProductPositionResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponse) GetRowsOk() ([]GetOnChainYieldsLockedProductPositionResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetOnChainYieldsLockedProductPositionResponseRowsInner and assigns it to the Rows field.
func (o *GetOnChainYieldsLockedProductPositionResponse) SetRows(v []GetOnChainYieldsLockedProductPositionResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetOnChainYieldsLockedProductPositionResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetOnChainYieldsLockedProductPositionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOnChainYieldsLockedProductPositionResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetOnChainYieldsLockedProductPositionResponse) UnmarshalJSON(data []byte) (err error) {
	varGetOnChainYieldsLockedProductPositionResponse := _GetOnChainYieldsLockedProductPositionResponse{}

	err = json.Unmarshal(data, &varGetOnChainYieldsLockedProductPositionResponse)

	if err != nil {
		return err
	}

	*o = GetOnChainYieldsLockedProductPositionResponse(varGetOnChainYieldsLockedProductPositionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOnChainYieldsLockedProductPositionResponse struct {
	value *GetOnChainYieldsLockedProductPositionResponse
	isSet bool
}

func (v NullableGetOnChainYieldsLockedProductPositionResponse) Get() *GetOnChainYieldsLockedProductPositionResponse {
	return v.value
}

func (v *NullableGetOnChainYieldsLockedProductPositionResponse) Set(val *GetOnChainYieldsLockedProductPositionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOnChainYieldsLockedProductPositionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOnChainYieldsLockedProductPositionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOnChainYieldsLockedProductPositionResponse(val *GetOnChainYieldsLockedProductPositionResponse) *NullableGetOnChainYieldsLockedProductPositionResponse {
	return &NullableGetOnChainYieldsLockedProductPositionResponse{value: val, isSet: true}
}

func (v NullableGetOnChainYieldsLockedProductPositionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOnChainYieldsLockedProductPositionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
