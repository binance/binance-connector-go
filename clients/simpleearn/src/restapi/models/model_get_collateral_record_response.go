/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetCollateralRecordResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCollateralRecordResponse{}

// GetCollateralRecordResponse struct for GetCollateralRecordResponse
type GetCollateralRecordResponse struct {
	Rows                 []GetCollateralRecordResponseRowsInner `json:"rows,omitempty"`
	Total                *string                                `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCollateralRecordResponse GetCollateralRecordResponse

// NewGetCollateralRecordResponse instantiates a new GetCollateralRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCollateralRecordResponse() *GetCollateralRecordResponse {
	this := GetCollateralRecordResponse{}
	return &this
}

// NewGetCollateralRecordResponseWithDefaults instantiates a new GetCollateralRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCollateralRecordResponseWithDefaults() *GetCollateralRecordResponse {
	this := GetCollateralRecordResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetCollateralRecordResponse) GetRows() []GetCollateralRecordResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetCollateralRecordResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralRecordResponse) GetRowsOk() ([]GetCollateralRecordResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetCollateralRecordResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetCollateralRecordResponseRowsInner and assigns it to the Rows field.
func (o *GetCollateralRecordResponse) SetRows(v []GetCollateralRecordResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetCollateralRecordResponse) GetTotal() string {
	if o == nil || common.IsNil(o.Total) {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollateralRecordResponse) GetTotalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetCollateralRecordResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *GetCollateralRecordResponse) SetTotal(v string) {
	o.Total = &v
}

func (o GetCollateralRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCollateralRecordResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetCollateralRecordResponse) UnmarshalJSON(data []byte) (err error) {
	varGetCollateralRecordResponse := _GetCollateralRecordResponse{}

	err = json.Unmarshal(data, &varGetCollateralRecordResponse)

	if err != nil {
		return err
	}

	*o = GetCollateralRecordResponse(varGetCollateralRecordResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCollateralRecordResponse struct {
	value *GetCollateralRecordResponse
	isSet bool
}

func (v NullableGetCollateralRecordResponse) Get() *GetCollateralRecordResponse {
	return v.value
}

func (v *NullableGetCollateralRecordResponse) Set(val *GetCollateralRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCollateralRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCollateralRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCollateralRecordResponse(val *GetCollateralRecordResponse) *NullableGetCollateralRecordResponse {
	return &NullableGetCollateralRecordResponse{value: val, isSet: true}
}

func (v NullableGetCollateralRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCollateralRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
