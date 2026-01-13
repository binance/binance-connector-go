/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSimpleEarnFlexibleProductListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSimpleEarnFlexibleProductListResponse{}

// GetSimpleEarnFlexibleProductListResponse struct for GetSimpleEarnFlexibleProductListResponse
type GetSimpleEarnFlexibleProductListResponse struct {
	Rows                 []GetSimpleEarnFlexibleProductListResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                              `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSimpleEarnFlexibleProductListResponse GetSimpleEarnFlexibleProductListResponse

// NewGetSimpleEarnFlexibleProductListResponse instantiates a new GetSimpleEarnFlexibleProductListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSimpleEarnFlexibleProductListResponse() *GetSimpleEarnFlexibleProductListResponse {
	this := GetSimpleEarnFlexibleProductListResponse{}
	return &this
}

// NewGetSimpleEarnFlexibleProductListResponseWithDefaults instantiates a new GetSimpleEarnFlexibleProductListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSimpleEarnFlexibleProductListResponseWithDefaults() *GetSimpleEarnFlexibleProductListResponse {
	this := GetSimpleEarnFlexibleProductListResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleProductListResponse) GetRows() []GetSimpleEarnFlexibleProductListResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetSimpleEarnFlexibleProductListResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleProductListResponse) GetRowsOk() ([]GetSimpleEarnFlexibleProductListResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleProductListResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetSimpleEarnFlexibleProductListResponseRowsInner and assigns it to the Rows field.
func (o *GetSimpleEarnFlexibleProductListResponse) SetRows(v []GetSimpleEarnFlexibleProductListResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleProductListResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleProductListResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleProductListResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetSimpleEarnFlexibleProductListResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetSimpleEarnFlexibleProductListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSimpleEarnFlexibleProductListResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetSimpleEarnFlexibleProductListResponse) UnmarshalJSON(data []byte) (err error) {
	varGetSimpleEarnFlexibleProductListResponse := _GetSimpleEarnFlexibleProductListResponse{}

	err = json.Unmarshal(data, &varGetSimpleEarnFlexibleProductListResponse)

	if err != nil {
		return err
	}

	*o = GetSimpleEarnFlexibleProductListResponse(varGetSimpleEarnFlexibleProductListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSimpleEarnFlexibleProductListResponse struct {
	value *GetSimpleEarnFlexibleProductListResponse
	isSet bool
}

func (v NullableGetSimpleEarnFlexibleProductListResponse) Get() *GetSimpleEarnFlexibleProductListResponse {
	return v.value
}

func (v *NullableGetSimpleEarnFlexibleProductListResponse) Set(val *GetSimpleEarnFlexibleProductListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSimpleEarnFlexibleProductListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSimpleEarnFlexibleProductListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSimpleEarnFlexibleProductListResponse(val *GetSimpleEarnFlexibleProductListResponse) *NullableGetSimpleEarnFlexibleProductListResponse {
	return &NullableGetSimpleEarnFlexibleProductListResponse{value: val, isSet: true}
}

func (v NullableGetSimpleEarnFlexibleProductListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSimpleEarnFlexibleProductListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
