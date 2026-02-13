/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetBfusdRedemptionHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBfusdRedemptionHistoryResponse{}

// GetBfusdRedemptionHistoryResponse struct for GetBfusdRedemptionHistoryResponse
type GetBfusdRedemptionHistoryResponse struct {
	Rows                 []GetBfusdRedemptionHistoryResponseRowsInner `json:"rows,omitempty"`
	Total                *int64                                       `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBfusdRedemptionHistoryResponse GetBfusdRedemptionHistoryResponse

// NewGetBfusdRedemptionHistoryResponse instantiates a new GetBfusdRedemptionHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBfusdRedemptionHistoryResponse() *GetBfusdRedemptionHistoryResponse {
	this := GetBfusdRedemptionHistoryResponse{}
	return &this
}

// NewGetBfusdRedemptionHistoryResponseWithDefaults instantiates a new GetBfusdRedemptionHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBfusdRedemptionHistoryResponseWithDefaults() *GetBfusdRedemptionHistoryResponse {
	this := GetBfusdRedemptionHistoryResponse{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetBfusdRedemptionHistoryResponse) GetRows() []GetBfusdRedemptionHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetBfusdRedemptionHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRedemptionHistoryResponse) GetRowsOk() ([]GetBfusdRedemptionHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetBfusdRedemptionHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetBfusdRedemptionHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetBfusdRedemptionHistoryResponse) SetRows(v []GetBfusdRedemptionHistoryResponseRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetBfusdRedemptionHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRedemptionHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetBfusdRedemptionHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetBfusdRedemptionHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o GetBfusdRedemptionHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBfusdRedemptionHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetBfusdRedemptionHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetBfusdRedemptionHistoryResponse := _GetBfusdRedemptionHistoryResponse{}

	err = json.Unmarshal(data, &varGetBfusdRedemptionHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetBfusdRedemptionHistoryResponse(varGetBfusdRedemptionHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBfusdRedemptionHistoryResponse struct {
	value *GetBfusdRedemptionHistoryResponse
	isSet bool
}

func (v NullableGetBfusdRedemptionHistoryResponse) Get() *GetBfusdRedemptionHistoryResponse {
	return v.value
}

func (v *NullableGetBfusdRedemptionHistoryResponse) Set(val *GetBfusdRedemptionHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBfusdRedemptionHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBfusdRedemptionHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBfusdRedemptionHistoryResponse(val *GetBfusdRedemptionHistoryResponse) *NullableGetBfusdRedemptionHistoryResponse {
	return &NullableGetBfusdRedemptionHistoryResponse{value: val, isSet: true}
}

func (v NullableGetBfusdRedemptionHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBfusdRedemptionHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
