/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOnChainYieldsLockedRedemptionRecordResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOnChainYieldsLockedRedemptionRecordResponse{}

// GetOnChainYieldsLockedRedemptionRecordResponse struct for GetOnChainYieldsLockedRedemptionRecordResponse
type GetOnChainYieldsLockedRedemptionRecordResponse struct {
	Rows                 []GetOnChainYieldsLockedRedemptionRecordResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                                    `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOnChainYieldsLockedRedemptionRecordResponse GetOnChainYieldsLockedRedemptionRecordResponse

// NewGetOnChainYieldsLockedRedemptionRecordResponse instantiates a new GetOnChainYieldsLockedRedemptionRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOnChainYieldsLockedRedemptionRecordResponse() *GetOnChainYieldsLockedRedemptionRecordResponse {
	this := GetOnChainYieldsLockedRedemptionRecordResponse{}
	return &this
}

// NewGetOnChainYieldsLockedRedemptionRecordResponseWithDefaults instantiates a new GetOnChainYieldsLockedRedemptionRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOnChainYieldsLockedRedemptionRecordResponseWithDefaults() *GetOnChainYieldsLockedRedemptionRecordResponse {
	this := GetOnChainYieldsLockedRedemptionRecordResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRedemptionRecordResponse) GetRows() []GetOnChainYieldsLockedRedemptionRecordResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetOnChainYieldsLockedRedemptionRecordResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponse) GetRowsOk() ([]GetOnChainYieldsLockedRedemptionRecordResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetOnChainYieldsLockedRedemptionRecordResponseRowsInner and assigns it to the Rows field.
func (o *GetOnChainYieldsLockedRedemptionRecordResponse) SetRows(v []GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRedemptionRecordResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetOnChainYieldsLockedRedemptionRecordResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetOnChainYieldsLockedRedemptionRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOnChainYieldsLockedRedemptionRecordResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetOnChainYieldsLockedRedemptionRecordResponse) UnmarshalJSON(data []byte) (err error) {
	varGetOnChainYieldsLockedRedemptionRecordResponse := _GetOnChainYieldsLockedRedemptionRecordResponse{}

	err = json.Unmarshal(data, &varGetOnChainYieldsLockedRedemptionRecordResponse)

	if err != nil {
		return err
	}

	*o = GetOnChainYieldsLockedRedemptionRecordResponse(varGetOnChainYieldsLockedRedemptionRecordResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOnChainYieldsLockedRedemptionRecordResponse struct {
	value *GetOnChainYieldsLockedRedemptionRecordResponse
	isSet bool
}

func (v NullableGetOnChainYieldsLockedRedemptionRecordResponse) Get() *GetOnChainYieldsLockedRedemptionRecordResponse {
	return v.value
}

func (v *NullableGetOnChainYieldsLockedRedemptionRecordResponse) Set(val *GetOnChainYieldsLockedRedemptionRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOnChainYieldsLockedRedemptionRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOnChainYieldsLockedRedemptionRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOnChainYieldsLockedRedemptionRecordResponse(val *GetOnChainYieldsLockedRedemptionRecordResponse) *NullableGetOnChainYieldsLockedRedemptionRecordResponse {
	return &NullableGetOnChainYieldsLockedRedemptionRecordResponse{value: val, isSet: true}
}

func (v NullableGetOnChainYieldsLockedRedemptionRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOnChainYieldsLockedRedemptionRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
