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

// checks if the GetOnChainYieldsLockedProductListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOnChainYieldsLockedProductListResponse{}

// GetOnChainYieldsLockedProductListResponse struct for GetOnChainYieldsLockedProductListResponse
type GetOnChainYieldsLockedProductListResponse struct {
	Rows                 []GetOnChainYieldsLockedProductListResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                               `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOnChainYieldsLockedProductListResponse GetOnChainYieldsLockedProductListResponse

// NewGetOnChainYieldsLockedProductListResponse instantiates a new GetOnChainYieldsLockedProductListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOnChainYieldsLockedProductListResponse() *GetOnChainYieldsLockedProductListResponse {
	this := GetOnChainYieldsLockedProductListResponse{}
	return &this
}

// NewGetOnChainYieldsLockedProductListResponseWithDefaults instantiates a new GetOnChainYieldsLockedProductListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOnChainYieldsLockedProductListResponseWithDefaults() *GetOnChainYieldsLockedProductListResponse {
	this := GetOnChainYieldsLockedProductListResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductListResponse) GetRows() []GetOnChainYieldsLockedProductListResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetOnChainYieldsLockedProductListResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductListResponse) GetRowsOk() ([]GetOnChainYieldsLockedProductListResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductListResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetOnChainYieldsLockedProductListResponseRowsInner and assigns it to the Rows field.
func (o *GetOnChainYieldsLockedProductListResponse) SetRows(v []GetOnChainYieldsLockedProductListResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductListResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductListResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductListResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetOnChainYieldsLockedProductListResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetOnChainYieldsLockedProductListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOnChainYieldsLockedProductListResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetOnChainYieldsLockedProductListResponse) UnmarshalJSON(data []byte) (err error) {
	varGetOnChainYieldsLockedProductListResponse := _GetOnChainYieldsLockedProductListResponse{}

	err = json.Unmarshal(data, &varGetOnChainYieldsLockedProductListResponse)

	if err != nil {
		return err
	}

	*o = GetOnChainYieldsLockedProductListResponse(varGetOnChainYieldsLockedProductListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOnChainYieldsLockedProductListResponse struct {
	value *GetOnChainYieldsLockedProductListResponse
	isSet bool
}

func (v NullableGetOnChainYieldsLockedProductListResponse) Get() *GetOnChainYieldsLockedProductListResponse {
	return v.value
}

func (v *NullableGetOnChainYieldsLockedProductListResponse) Set(val *GetOnChainYieldsLockedProductListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOnChainYieldsLockedProductListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOnChainYieldsLockedProductListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOnChainYieldsLockedProductListResponse(val *GetOnChainYieldsLockedProductListResponse) *NullableGetOnChainYieldsLockedProductListResponse {
	return &NullableGetOnChainYieldsLockedProductListResponse{value: val, isSet: true}
}

func (v NullableGetOnChainYieldsLockedProductListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOnChainYieldsLockedProductListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
