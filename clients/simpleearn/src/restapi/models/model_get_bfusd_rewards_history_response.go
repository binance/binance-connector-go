/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetBfusdRewardsHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBfusdRewardsHistoryResponse{}

// GetBfusdRewardsHistoryResponse struct for GetBfusdRewardsHistoryResponse
type GetBfusdRewardsHistoryResponse struct {
	Rows                 []GetBfusdRewardsHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                    `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBfusdRewardsHistoryResponse GetBfusdRewardsHistoryResponse

// NewGetBfusdRewardsHistoryResponse instantiates a new GetBfusdRewardsHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBfusdRewardsHistoryResponse() *GetBfusdRewardsHistoryResponse {
	this := GetBfusdRewardsHistoryResponse{}
	return &this
}

// NewGetBfusdRewardsHistoryResponseWithDefaults instantiates a new GetBfusdRewardsHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBfusdRewardsHistoryResponseWithDefaults() *GetBfusdRewardsHistoryResponse {
	this := GetBfusdRewardsHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetBfusdRewardsHistoryResponse) GetRows() []GetBfusdRewardsHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetBfusdRewardsHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRewardsHistoryResponse) GetRowsOk() ([]GetBfusdRewardsHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetBfusdRewardsHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetBfusdRewardsHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetBfusdRewardsHistoryResponse) SetRows(v []GetBfusdRewardsHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetBfusdRewardsHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRewardsHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetBfusdRewardsHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetBfusdRewardsHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetBfusdRewardsHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBfusdRewardsHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetBfusdRewardsHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetBfusdRewardsHistoryResponse := _GetBfusdRewardsHistoryResponse{}

	err = json.Unmarshal(data, &varGetBfusdRewardsHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetBfusdRewardsHistoryResponse(varGetBfusdRewardsHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBfusdRewardsHistoryResponse struct {
	value *GetBfusdRewardsHistoryResponse
	isSet bool
}

func (v NullableGetBfusdRewardsHistoryResponse) Get() *GetBfusdRewardsHistoryResponse {
	return v.value
}

func (v *NullableGetBfusdRewardsHistoryResponse) Set(val *GetBfusdRewardsHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBfusdRewardsHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBfusdRewardsHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBfusdRewardsHistoryResponse(val *GetBfusdRewardsHistoryResponse) *NullableGetBfusdRewardsHistoryResponse {
	return &NullableGetBfusdRewardsHistoryResponse{value: val, isSet: true}
}

func (v NullableGetBfusdRewardsHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBfusdRewardsHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
