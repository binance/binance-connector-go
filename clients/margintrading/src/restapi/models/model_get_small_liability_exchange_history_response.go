/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetSmallLiabilityExchangeHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSmallLiabilityExchangeHistoryResponse{}

// GetSmallLiabilityExchangeHistoryResponse struct for GetSmallLiabilityExchangeHistoryResponse
type GetSmallLiabilityExchangeHistoryResponse struct {
	Total                *int64                                              `json:"total,omitempty"`
	Rows                 []GetSmallLiabilityExchangeHistoryResponseRowsInner `json:"rows,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSmallLiabilityExchangeHistoryResponse GetSmallLiabilityExchangeHistoryResponse

// NewGetSmallLiabilityExchangeHistoryResponse instantiates a new GetSmallLiabilityExchangeHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSmallLiabilityExchangeHistoryResponse() *GetSmallLiabilityExchangeHistoryResponse {
	this := GetSmallLiabilityExchangeHistoryResponse{}
	return &this
}

// NewGetSmallLiabilityExchangeHistoryResponseWithDefaults instantiates a new GetSmallLiabilityExchangeHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSmallLiabilityExchangeHistoryResponseWithDefaults() *GetSmallLiabilityExchangeHistoryResponse {
	this := GetSmallLiabilityExchangeHistoryResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetSmallLiabilityExchangeHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSmallLiabilityExchangeHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetSmallLiabilityExchangeHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetSmallLiabilityExchangeHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetSmallLiabilityExchangeHistoryResponse) GetRows() []GetSmallLiabilityExchangeHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetSmallLiabilityExchangeHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSmallLiabilityExchangeHistoryResponse) GetRowsOk() ([]GetSmallLiabilityExchangeHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetSmallLiabilityExchangeHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetSmallLiabilityExchangeHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetSmallLiabilityExchangeHistoryResponse) SetRows(v []GetSmallLiabilityExchangeHistoryResponseRowsInner) {
	o.Rows = v
}

func (o GetSmallLiabilityExchangeHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSmallLiabilityExchangeHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !common.IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSmallLiabilityExchangeHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetSmallLiabilityExchangeHistoryResponse := _GetSmallLiabilityExchangeHistoryResponse{}

	err = json.Unmarshal(data, &varGetSmallLiabilityExchangeHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetSmallLiabilityExchangeHistoryResponse(varGetSmallLiabilityExchangeHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "rows")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSmallLiabilityExchangeHistoryResponse struct {
	value *GetSmallLiabilityExchangeHistoryResponse
	isSet bool
}

func (v NullableGetSmallLiabilityExchangeHistoryResponse) Get() *GetSmallLiabilityExchangeHistoryResponse {
	return v.value
}

func (v *NullableGetSmallLiabilityExchangeHistoryResponse) Set(val *GetSmallLiabilityExchangeHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSmallLiabilityExchangeHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSmallLiabilityExchangeHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSmallLiabilityExchangeHistoryResponse(val *GetSmallLiabilityExchangeHistoryResponse) *NullableGetSmallLiabilityExchangeHistoryResponse {
	return &NullableGetSmallLiabilityExchangeHistoryResponse{value: val, isSet: true}
}

func (v NullableGetSmallLiabilityExchangeHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSmallLiabilityExchangeHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
