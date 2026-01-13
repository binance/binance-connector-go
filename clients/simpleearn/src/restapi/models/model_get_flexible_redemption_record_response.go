/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFlexibleRedemptionRecordResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleRedemptionRecordResponse{}

// GetFlexibleRedemptionRecordResponse struct for GetFlexibleRedemptionRecordResponse
type GetFlexibleRedemptionRecordResponse struct {
	Rows                 []GetFlexibleRedemptionRecordResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                         `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleRedemptionRecordResponse GetFlexibleRedemptionRecordResponse

// NewGetFlexibleRedemptionRecordResponse instantiates a new GetFlexibleRedemptionRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleRedemptionRecordResponse() *GetFlexibleRedemptionRecordResponse {
	this := GetFlexibleRedemptionRecordResponse{}
	return &this
}

// NewGetFlexibleRedemptionRecordResponseWithDefaults instantiates a new GetFlexibleRedemptionRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleRedemptionRecordResponseWithDefaults() *GetFlexibleRedemptionRecordResponse {
	this := GetFlexibleRedemptionRecordResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetFlexibleRedemptionRecordResponse) GetRows() []GetFlexibleRedemptionRecordResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetFlexibleRedemptionRecordResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleRedemptionRecordResponse) GetRowsOk() ([]GetFlexibleRedemptionRecordResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetFlexibleRedemptionRecordResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetFlexibleRedemptionRecordResponseRowsInner and assigns it to the Rows field.
func (o *GetFlexibleRedemptionRecordResponse) SetRows(v []GetFlexibleRedemptionRecordResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetFlexibleRedemptionRecordResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleRedemptionRecordResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetFlexibleRedemptionRecordResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetFlexibleRedemptionRecordResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetFlexibleRedemptionRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleRedemptionRecordResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetFlexibleRedemptionRecordResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleRedemptionRecordResponse := _GetFlexibleRedemptionRecordResponse{}

	err = json.Unmarshal(data, &varGetFlexibleRedemptionRecordResponse)

	if err != nil {
		return err
	}

	*o = GetFlexibleRedemptionRecordResponse(varGetFlexibleRedemptionRecordResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleRedemptionRecordResponse struct {
	value *GetFlexibleRedemptionRecordResponse
	isSet bool
}

func (v NullableGetFlexibleRedemptionRecordResponse) Get() *GetFlexibleRedemptionRecordResponse {
	return v.value
}

func (v *NullableGetFlexibleRedemptionRecordResponse) Set(val *GetFlexibleRedemptionRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleRedemptionRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleRedemptionRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleRedemptionRecordResponse(val *GetFlexibleRedemptionRecordResponse) *NullableGetFlexibleRedemptionRecordResponse {
	return &NullableGetFlexibleRedemptionRecordResponse{value: val, isSet: true}
}

func (v NullableGetFlexibleRedemptionRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleRedemptionRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
