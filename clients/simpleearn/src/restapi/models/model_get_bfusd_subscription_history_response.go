/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetBfusdSubscriptionHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBfusdSubscriptionHistoryResponse{}

// GetBfusdSubscriptionHistoryResponse struct for GetBfusdSubscriptionHistoryResponse
type GetBfusdSubscriptionHistoryResponse struct {
	Rows                 []GetBfusdSubscriptionHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                         `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBfusdSubscriptionHistoryResponse GetBfusdSubscriptionHistoryResponse

// NewGetBfusdSubscriptionHistoryResponse instantiates a new GetBfusdSubscriptionHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBfusdSubscriptionHistoryResponse() *GetBfusdSubscriptionHistoryResponse {
	this := GetBfusdSubscriptionHistoryResponse{}
	return &this
}

// NewGetBfusdSubscriptionHistoryResponseWithDefaults instantiates a new GetBfusdSubscriptionHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBfusdSubscriptionHistoryResponseWithDefaults() *GetBfusdSubscriptionHistoryResponse {
	this := GetBfusdSubscriptionHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetBfusdSubscriptionHistoryResponse) GetRows() []GetBfusdSubscriptionHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetBfusdSubscriptionHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdSubscriptionHistoryResponse) GetRowsOk() ([]GetBfusdSubscriptionHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetBfusdSubscriptionHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetBfusdSubscriptionHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetBfusdSubscriptionHistoryResponse) SetRows(v []GetBfusdSubscriptionHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetBfusdSubscriptionHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdSubscriptionHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetBfusdSubscriptionHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetBfusdSubscriptionHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetBfusdSubscriptionHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBfusdSubscriptionHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetBfusdSubscriptionHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetBfusdSubscriptionHistoryResponse := _GetBfusdSubscriptionHistoryResponse{}

	err = json.Unmarshal(data, &varGetBfusdSubscriptionHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetBfusdSubscriptionHistoryResponse(varGetBfusdSubscriptionHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBfusdSubscriptionHistoryResponse struct {
	value *GetBfusdSubscriptionHistoryResponse
	isSet bool
}

func (v NullableGetBfusdSubscriptionHistoryResponse) Get() *GetBfusdSubscriptionHistoryResponse {
	return v.value
}

func (v *NullableGetBfusdSubscriptionHistoryResponse) Set(val *GetBfusdSubscriptionHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBfusdSubscriptionHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBfusdSubscriptionHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBfusdSubscriptionHistoryResponse(val *GetBfusdSubscriptionHistoryResponse) *NullableGetBfusdSubscriptionHistoryResponse {
	return &NullableGetBfusdSubscriptionHistoryResponse{value: val, isSet: true}
}

func (v NullableGetBfusdSubscriptionHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBfusdSubscriptionHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
