/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetSimpleEarnLockedProductListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSimpleEarnLockedProductListResponse{}

// GetSimpleEarnLockedProductListResponse struct for GetSimpleEarnLockedProductListResponse
type GetSimpleEarnLockedProductListResponse struct {
	Rows                 []GetSimpleEarnLockedProductListResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                            `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSimpleEarnLockedProductListResponse GetSimpleEarnLockedProductListResponse

// NewGetSimpleEarnLockedProductListResponse instantiates a new GetSimpleEarnLockedProductListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSimpleEarnLockedProductListResponse() *GetSimpleEarnLockedProductListResponse {
	this := GetSimpleEarnLockedProductListResponse{}
	return &this
}

// NewGetSimpleEarnLockedProductListResponseWithDefaults instantiates a new GetSimpleEarnLockedProductListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSimpleEarnLockedProductListResponseWithDefaults() *GetSimpleEarnLockedProductListResponse {
	this := GetSimpleEarnLockedProductListResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponse) GetRows() []GetSimpleEarnLockedProductListResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetSimpleEarnLockedProductListResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponse) GetRowsOk() ([]GetSimpleEarnLockedProductListResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetSimpleEarnLockedProductListResponseRowsInner and assigns it to the Rows field.
func (o *GetSimpleEarnLockedProductListResponse) SetRows(v []GetSimpleEarnLockedProductListResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetSimpleEarnLockedProductListResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetSimpleEarnLockedProductListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSimpleEarnLockedProductListResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetSimpleEarnLockedProductListResponse) UnmarshalJSON(data []byte) (err error) {
	varGetSimpleEarnLockedProductListResponse := _GetSimpleEarnLockedProductListResponse{}

	err = json.Unmarshal(data, &varGetSimpleEarnLockedProductListResponse)

	if err != nil {
		return err
	}

	*o = GetSimpleEarnLockedProductListResponse(varGetSimpleEarnLockedProductListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSimpleEarnLockedProductListResponse struct {
	value *GetSimpleEarnLockedProductListResponse
	isSet bool
}

func (v NullableGetSimpleEarnLockedProductListResponse) Get() *GetSimpleEarnLockedProductListResponse {
	return v.value
}

func (v *NullableGetSimpleEarnLockedProductListResponse) Set(val *GetSimpleEarnLockedProductListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSimpleEarnLockedProductListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSimpleEarnLockedProductListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSimpleEarnLockedProductListResponse(val *GetSimpleEarnLockedProductListResponse) *NullableGetSimpleEarnLockedProductListResponse {
	return &NullableGetSimpleEarnLockedProductListResponse{value: val, isSet: true}
}

func (v NullableGetSimpleEarnLockedProductListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSimpleEarnLockedProductListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
