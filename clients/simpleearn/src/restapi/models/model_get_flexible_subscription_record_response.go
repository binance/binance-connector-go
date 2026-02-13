/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFlexibleSubscriptionRecordResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleSubscriptionRecordResponse{}

// GetFlexibleSubscriptionRecordResponse struct for GetFlexibleSubscriptionRecordResponse
type GetFlexibleSubscriptionRecordResponse struct {
	Rows                 []GetFlexibleSubscriptionRecordResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                           `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleSubscriptionRecordResponse GetFlexibleSubscriptionRecordResponse

// NewGetFlexibleSubscriptionRecordResponse instantiates a new GetFlexibleSubscriptionRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleSubscriptionRecordResponse() *GetFlexibleSubscriptionRecordResponse {
	this := GetFlexibleSubscriptionRecordResponse{}
	return &this
}

// NewGetFlexibleSubscriptionRecordResponseWithDefaults instantiates a new GetFlexibleSubscriptionRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleSubscriptionRecordResponseWithDefaults() *GetFlexibleSubscriptionRecordResponse {
	this := GetFlexibleSubscriptionRecordResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionRecordResponse) GetRows() []GetFlexibleSubscriptionRecordResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetFlexibleSubscriptionRecordResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionRecordResponse) GetRowsOk() ([]GetFlexibleSubscriptionRecordResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionRecordResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetFlexibleSubscriptionRecordResponseRowsInner and assigns it to the Rows field.
func (o *GetFlexibleSubscriptionRecordResponse) SetRows(v []GetFlexibleSubscriptionRecordResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionRecordResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionRecordResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionRecordResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetFlexibleSubscriptionRecordResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetFlexibleSubscriptionRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleSubscriptionRecordResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetFlexibleSubscriptionRecordResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleSubscriptionRecordResponse := _GetFlexibleSubscriptionRecordResponse{}

	err = json.Unmarshal(data, &varGetFlexibleSubscriptionRecordResponse)

	if err != nil {
		return err
	}

	*o = GetFlexibleSubscriptionRecordResponse(varGetFlexibleSubscriptionRecordResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleSubscriptionRecordResponse struct {
	value *GetFlexibleSubscriptionRecordResponse
	isSet bool
}

func (v NullableGetFlexibleSubscriptionRecordResponse) Get() *GetFlexibleSubscriptionRecordResponse {
	return v.value
}

func (v *NullableGetFlexibleSubscriptionRecordResponse) Set(val *GetFlexibleSubscriptionRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleSubscriptionRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleSubscriptionRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleSubscriptionRecordResponse(val *GetFlexibleSubscriptionRecordResponse) *NullableGetFlexibleSubscriptionRecordResponse {
	return &NullableGetFlexibleSubscriptionRecordResponse{value: val, isSet: true}
}

func (v NullableGetFlexibleSubscriptionRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleSubscriptionRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
