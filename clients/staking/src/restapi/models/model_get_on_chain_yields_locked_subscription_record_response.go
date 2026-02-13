/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOnChainYieldsLockedSubscriptionRecordResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOnChainYieldsLockedSubscriptionRecordResponse{}

// GetOnChainYieldsLockedSubscriptionRecordResponse struct for GetOnChainYieldsLockedSubscriptionRecordResponse
type GetOnChainYieldsLockedSubscriptionRecordResponse struct {
	Rows                 []GetOnChainYieldsLockedSubscriptionRecordResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                                      `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOnChainYieldsLockedSubscriptionRecordResponse GetOnChainYieldsLockedSubscriptionRecordResponse

// NewGetOnChainYieldsLockedSubscriptionRecordResponse instantiates a new GetOnChainYieldsLockedSubscriptionRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOnChainYieldsLockedSubscriptionRecordResponse() *GetOnChainYieldsLockedSubscriptionRecordResponse {
	this := GetOnChainYieldsLockedSubscriptionRecordResponse{}
	return &this
}

// NewGetOnChainYieldsLockedSubscriptionRecordResponseWithDefaults instantiates a new GetOnChainYieldsLockedSubscriptionRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOnChainYieldsLockedSubscriptionRecordResponseWithDefaults() *GetOnChainYieldsLockedSubscriptionRecordResponse {
	this := GetOnChainYieldsLockedSubscriptionRecordResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedSubscriptionRecordResponse) GetRows() []GetOnChainYieldsLockedSubscriptionRecordResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetOnChainYieldsLockedSubscriptionRecordResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedSubscriptionRecordResponse) GetRowsOk() ([]GetOnChainYieldsLockedSubscriptionRecordResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedSubscriptionRecordResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetOnChainYieldsLockedSubscriptionRecordResponseRowsInner and assigns it to the Rows field.
func (o *GetOnChainYieldsLockedSubscriptionRecordResponse) SetRows(v []GetOnChainYieldsLockedSubscriptionRecordResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedSubscriptionRecordResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedSubscriptionRecordResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedSubscriptionRecordResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetOnChainYieldsLockedSubscriptionRecordResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetOnChainYieldsLockedSubscriptionRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOnChainYieldsLockedSubscriptionRecordResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetOnChainYieldsLockedSubscriptionRecordResponse) UnmarshalJSON(data []byte) (err error) {
	varGetOnChainYieldsLockedSubscriptionRecordResponse := _GetOnChainYieldsLockedSubscriptionRecordResponse{}

	err = json.Unmarshal(data, &varGetOnChainYieldsLockedSubscriptionRecordResponse)

	if err != nil {
		return err
	}

	*o = GetOnChainYieldsLockedSubscriptionRecordResponse(varGetOnChainYieldsLockedSubscriptionRecordResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOnChainYieldsLockedSubscriptionRecordResponse struct {
	value *GetOnChainYieldsLockedSubscriptionRecordResponse
	isSet bool
}

func (v NullableGetOnChainYieldsLockedSubscriptionRecordResponse) Get() *GetOnChainYieldsLockedSubscriptionRecordResponse {
	return v.value
}

func (v *NullableGetOnChainYieldsLockedSubscriptionRecordResponse) Set(val *GetOnChainYieldsLockedSubscriptionRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOnChainYieldsLockedSubscriptionRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOnChainYieldsLockedSubscriptionRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOnChainYieldsLockedSubscriptionRecordResponse(val *GetOnChainYieldsLockedSubscriptionRecordResponse) *NullableGetOnChainYieldsLockedSubscriptionRecordResponse {
	return &NullableGetOnChainYieldsLockedSubscriptionRecordResponse{value: val, isSet: true}
}

func (v NullableGetOnChainYieldsLockedSubscriptionRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOnChainYieldsLockedSubscriptionRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
