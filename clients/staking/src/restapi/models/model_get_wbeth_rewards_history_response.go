/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetWbethRewardsHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetWbethRewardsHistoryResponse{}

// GetWbethRewardsHistoryResponse struct for GetWbethRewardsHistoryResponse
type GetWbethRewardsHistoryResponse struct {
	EstRewardsInETH      *string                                   `json:"estRewardsInETH,omitempty"`
	Rows                 []GetWbethRewardsHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                    `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetWbethRewardsHistoryResponse GetWbethRewardsHistoryResponse

// NewGetWbethRewardsHistoryResponse instantiates a new GetWbethRewardsHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWbethRewardsHistoryResponse() *GetWbethRewardsHistoryResponse {
	this := GetWbethRewardsHistoryResponse{}
	return &this
}

// NewGetWbethRewardsHistoryResponseWithDefaults instantiates a new GetWbethRewardsHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWbethRewardsHistoryResponseWithDefaults() *GetWbethRewardsHistoryResponse {
	this := GetWbethRewardsHistoryResponse{}
	return &this
}

// GetEstRewardsInETH returns the EstRewardsInETH field value if set, zero value otherwise.
func (o *GetWbethRewardsHistoryResponse) GetEstRewardsInETH() string {
	if o == nil || common.IsNil(o.EstRewardsInETH) {
		var ret string
		return ret
	}
	return *o.EstRewardsInETH
}

// GetEstRewardsInETHOk returns a tuple with the EstRewardsInETH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethRewardsHistoryResponse) GetEstRewardsInETHOk() (*string, bool) {
	if o == nil || common.IsNil(o.EstRewardsInETH) {
		return nil, false
	}
	return o.EstRewardsInETH, true
}

// HasEstRewardsInETH returns a boolean if a field has been set.
func (o *GetWbethRewardsHistoryResponse) HasEstRewardsInETH() bool {
	if o != nil && !common.IsNil(o.EstRewardsInETH) {
		return true
	}

	return false
}

// SetEstRewardsInETH gets a reference to the given string and assigns it to the EstRewardsInETH field.
func (o *GetWbethRewardsHistoryResponse) SetEstRewardsInETH(v string) {
	o.EstRewardsInETH = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetWbethRewardsHistoryResponse) GetRows() []GetWbethRewardsHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetWbethRewardsHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethRewardsHistoryResponse) GetRowsOk() ([]GetWbethRewardsHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetWbethRewardsHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetWbethRewardsHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetWbethRewardsHistoryResponse) SetRows(v []GetWbethRewardsHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetWbethRewardsHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethRewardsHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetWbethRewardsHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetWbethRewardsHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetWbethRewardsHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWbethRewardsHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EstRewardsInETH) {
		toSerialize["estRewardsInETH"] = o.EstRewardsInETH
	}
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

func (o *GetWbethRewardsHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetWbethRewardsHistoryResponse := _GetWbethRewardsHistoryResponse{}

	err = json.Unmarshal(data, &varGetWbethRewardsHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetWbethRewardsHistoryResponse(varGetWbethRewardsHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "estRewardsInETH")
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetWbethRewardsHistoryResponse struct {
	value *GetWbethRewardsHistoryResponse
	isSet bool
}

func (v NullableGetWbethRewardsHistoryResponse) Get() *GetWbethRewardsHistoryResponse {
	return v.value
}

func (v *NullableGetWbethRewardsHistoryResponse) Set(val *GetWbethRewardsHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWbethRewardsHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWbethRewardsHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWbethRewardsHistoryResponse(val *GetWbethRewardsHistoryResponse) *NullableGetWbethRewardsHistoryResponse {
	return &NullableGetWbethRewardsHistoryResponse{value: val, isSet: true}
}

func (v NullableGetWbethRewardsHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWbethRewardsHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
