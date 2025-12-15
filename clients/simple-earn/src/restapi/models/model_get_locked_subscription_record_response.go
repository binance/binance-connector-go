/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetLockedSubscriptionRecordResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLockedSubscriptionRecordResponse{}

// GetLockedSubscriptionRecordResponse struct for GetLockedSubscriptionRecordResponse
type GetLockedSubscriptionRecordResponse struct {
	Rows                 []GetLockedSubscriptionRecordResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                         `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLockedSubscriptionRecordResponse GetLockedSubscriptionRecordResponse

// NewGetLockedSubscriptionRecordResponse instantiates a new GetLockedSubscriptionRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLockedSubscriptionRecordResponse() *GetLockedSubscriptionRecordResponse {
	this := GetLockedSubscriptionRecordResponse{}
	return &this
}

// NewGetLockedSubscriptionRecordResponseWithDefaults instantiates a new GetLockedSubscriptionRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLockedSubscriptionRecordResponseWithDefaults() *GetLockedSubscriptionRecordResponse {
	this := GetLockedSubscriptionRecordResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetLockedSubscriptionRecordResponse) GetRows() []GetLockedSubscriptionRecordResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetLockedSubscriptionRecordResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionRecordResponse) GetRowsOk() ([]GetLockedSubscriptionRecordResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetLockedSubscriptionRecordResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetLockedSubscriptionRecordResponseRowsInner and assigns it to the Rows field.
func (o *GetLockedSubscriptionRecordResponse) SetRows(v []GetLockedSubscriptionRecordResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetLockedSubscriptionRecordResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionRecordResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetLockedSubscriptionRecordResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetLockedSubscriptionRecordResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetLockedSubscriptionRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLockedSubscriptionRecordResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetLockedSubscriptionRecordResponse) UnmarshalJSON(data []byte) (err error) {
	varGetLockedSubscriptionRecordResponse := _GetLockedSubscriptionRecordResponse{}

	err = json.Unmarshal(data, &varGetLockedSubscriptionRecordResponse)

	if err != nil {
		return err
	}

	*o = GetLockedSubscriptionRecordResponse(varGetLockedSubscriptionRecordResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLockedSubscriptionRecordResponse struct {
	value *GetLockedSubscriptionRecordResponse
	isSet bool
}

func (v NullableGetLockedSubscriptionRecordResponse) Get() *GetLockedSubscriptionRecordResponse {
	return v.value
}

func (v *NullableGetLockedSubscriptionRecordResponse) Set(val *GetLockedSubscriptionRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLockedSubscriptionRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLockedSubscriptionRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLockedSubscriptionRecordResponse(val *GetLockedSubscriptionRecordResponse) *NullableGetLockedSubscriptionRecordResponse {
	return &NullableGetLockedSubscriptionRecordResponse{value: val, isSet: true}
}

func (v NullableGetLockedSubscriptionRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLockedSubscriptionRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
