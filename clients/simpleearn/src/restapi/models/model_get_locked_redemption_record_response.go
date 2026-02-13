/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetLockedRedemptionRecordResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLockedRedemptionRecordResponse{}

// GetLockedRedemptionRecordResponse struct for GetLockedRedemptionRecordResponse
type GetLockedRedemptionRecordResponse struct {
	Rows                 []GetLockedRedemptionRecordResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                       `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLockedRedemptionRecordResponse GetLockedRedemptionRecordResponse

// NewGetLockedRedemptionRecordResponse instantiates a new GetLockedRedemptionRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLockedRedemptionRecordResponse() *GetLockedRedemptionRecordResponse {
	this := GetLockedRedemptionRecordResponse{}
	return &this
}

// NewGetLockedRedemptionRecordResponseWithDefaults instantiates a new GetLockedRedemptionRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLockedRedemptionRecordResponseWithDefaults() *GetLockedRedemptionRecordResponse {
	this := GetLockedRedemptionRecordResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponse) GetRows() []GetLockedRedemptionRecordResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetLockedRedemptionRecordResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponse) GetRowsOk() ([]GetLockedRedemptionRecordResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetLockedRedemptionRecordResponseRowsInner and assigns it to the Rows field.
func (o *GetLockedRedemptionRecordResponse) SetRows(v []GetLockedRedemptionRecordResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetLockedRedemptionRecordResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetLockedRedemptionRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLockedRedemptionRecordResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetLockedRedemptionRecordResponse) UnmarshalJSON(data []byte) (err error) {
	varGetLockedRedemptionRecordResponse := _GetLockedRedemptionRecordResponse{}

	err = json.Unmarshal(data, &varGetLockedRedemptionRecordResponse)

	if err != nil {
		return err
	}

	*o = GetLockedRedemptionRecordResponse(varGetLockedRedemptionRecordResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLockedRedemptionRecordResponse struct {
	value *GetLockedRedemptionRecordResponse
	isSet bool
}

func (v NullableGetLockedRedemptionRecordResponse) Get() *GetLockedRedemptionRecordResponse {
	return v.value
}

func (v *NullableGetLockedRedemptionRecordResponse) Set(val *GetLockedRedemptionRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLockedRedemptionRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLockedRedemptionRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLockedRedemptionRecordResponse(val *GetLockedRedemptionRecordResponse) *NullableGetLockedRedemptionRecordResponse {
	return &NullableGetLockedRedemptionRecordResponse{value: val, isSet: true}
}

func (v NullableGetLockedRedemptionRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLockedRedemptionRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
