/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetRwusdRedemptionHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetRwusdRedemptionHistoryResponse{}

// GetRwusdRedemptionHistoryResponse struct for GetRwusdRedemptionHistoryResponse
type GetRwusdRedemptionHistoryResponse struct {
	Rows                 []GetRwusdRedemptionHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                       `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRwusdRedemptionHistoryResponse GetRwusdRedemptionHistoryResponse

// NewGetRwusdRedemptionHistoryResponse instantiates a new GetRwusdRedemptionHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRwusdRedemptionHistoryResponse() *GetRwusdRedemptionHistoryResponse {
	this := GetRwusdRedemptionHistoryResponse{}
	return &this
}

// NewGetRwusdRedemptionHistoryResponseWithDefaults instantiates a new GetRwusdRedemptionHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRwusdRedemptionHistoryResponseWithDefaults() *GetRwusdRedemptionHistoryResponse {
	this := GetRwusdRedemptionHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetRwusdRedemptionHistoryResponse) GetRows() []GetRwusdRedemptionHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetRwusdRedemptionHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdRedemptionHistoryResponse) GetRowsOk() ([]GetRwusdRedemptionHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetRwusdRedemptionHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetRwusdRedemptionHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetRwusdRedemptionHistoryResponse) SetRows(v []GetRwusdRedemptionHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetRwusdRedemptionHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdRedemptionHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetRwusdRedemptionHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetRwusdRedemptionHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetRwusdRedemptionHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRwusdRedemptionHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetRwusdRedemptionHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetRwusdRedemptionHistoryResponse := _GetRwusdRedemptionHistoryResponse{}

	err = json.Unmarshal(data, &varGetRwusdRedemptionHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetRwusdRedemptionHistoryResponse(varGetRwusdRedemptionHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRwusdRedemptionHistoryResponse struct {
	value *GetRwusdRedemptionHistoryResponse
	isSet bool
}

func (v NullableGetRwusdRedemptionHistoryResponse) Get() *GetRwusdRedemptionHistoryResponse {
	return v.value
}

func (v *NullableGetRwusdRedemptionHistoryResponse) Set(val *GetRwusdRedemptionHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRwusdRedemptionHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRwusdRedemptionHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRwusdRedemptionHistoryResponse(val *GetRwusdRedemptionHistoryResponse) *NullableGetRwusdRedemptionHistoryResponse {
	return &NullableGetRwusdRedemptionHistoryResponse{value: val, isSet: true}
}

func (v NullableGetRwusdRedemptionHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRwusdRedemptionHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
