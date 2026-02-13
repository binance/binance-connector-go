/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFlexibleRewardsHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleRewardsHistoryResponse{}

// GetFlexibleRewardsHistoryResponse struct for GetFlexibleRewardsHistoryResponse
type GetFlexibleRewardsHistoryResponse struct {
	Rows                 []GetFlexibleRewardsHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                       `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleRewardsHistoryResponse GetFlexibleRewardsHistoryResponse

// NewGetFlexibleRewardsHistoryResponse instantiates a new GetFlexibleRewardsHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleRewardsHistoryResponse() *GetFlexibleRewardsHistoryResponse {
	this := GetFlexibleRewardsHistoryResponse{}
	return &this
}

// NewGetFlexibleRewardsHistoryResponseWithDefaults instantiates a new GetFlexibleRewardsHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleRewardsHistoryResponseWithDefaults() *GetFlexibleRewardsHistoryResponse {
	this := GetFlexibleRewardsHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetFlexibleRewardsHistoryResponse) GetRows() []GetFlexibleRewardsHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetFlexibleRewardsHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleRewardsHistoryResponse) GetRowsOk() ([]GetFlexibleRewardsHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetFlexibleRewardsHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetFlexibleRewardsHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetFlexibleRewardsHistoryResponse) SetRows(v []GetFlexibleRewardsHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetFlexibleRewardsHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleRewardsHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetFlexibleRewardsHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetFlexibleRewardsHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetFlexibleRewardsHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleRewardsHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetFlexibleRewardsHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleRewardsHistoryResponse := _GetFlexibleRewardsHistoryResponse{}

	err = json.Unmarshal(data, &varGetFlexibleRewardsHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetFlexibleRewardsHistoryResponse(varGetFlexibleRewardsHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleRewardsHistoryResponse struct {
	value *GetFlexibleRewardsHistoryResponse
	isSet bool
}

func (v NullableGetFlexibleRewardsHistoryResponse) Get() *GetFlexibleRewardsHistoryResponse {
	return v.value
}

func (v *NullableGetFlexibleRewardsHistoryResponse) Set(val *GetFlexibleRewardsHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleRewardsHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleRewardsHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleRewardsHistoryResponse(val *GetFlexibleRewardsHistoryResponse) *NullableGetFlexibleRewardsHistoryResponse {
	return &NullableGetFlexibleRewardsHistoryResponse{value: val, isSet: true}
}

func (v NullableGetFlexibleRewardsHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleRewardsHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
