/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetCloudMiningPaymentAndRefundHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCloudMiningPaymentAndRefundHistoryResponse{}

// GetCloudMiningPaymentAndRefundHistoryResponse struct for GetCloudMiningPaymentAndRefundHistoryResponse
type GetCloudMiningPaymentAndRefundHistoryResponse struct {
	Total                *int64                                                   `json:"total,omitempty"`
	Rows                 []GetCloudMiningPaymentAndRefundHistoryResponseRowsInner `json:"rows,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCloudMiningPaymentAndRefundHistoryResponse GetCloudMiningPaymentAndRefundHistoryResponse

// NewGetCloudMiningPaymentAndRefundHistoryResponse instantiates a new GetCloudMiningPaymentAndRefundHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCloudMiningPaymentAndRefundHistoryResponse() *GetCloudMiningPaymentAndRefundHistoryResponse {
	this := GetCloudMiningPaymentAndRefundHistoryResponse{}
	return &this
}

// NewGetCloudMiningPaymentAndRefundHistoryResponseWithDefaults instantiates a new GetCloudMiningPaymentAndRefundHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCloudMiningPaymentAndRefundHistoryResponseWithDefaults() *GetCloudMiningPaymentAndRefundHistoryResponse {
	this := GetCloudMiningPaymentAndRefundHistoryResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetCloudMiningPaymentAndRefundHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudMiningPaymentAndRefundHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetCloudMiningPaymentAndRefundHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetCloudMiningPaymentAndRefundHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *GetCloudMiningPaymentAndRefundHistoryResponse) GetRows() []GetCloudMiningPaymentAndRefundHistoryResponseRowsInner {
	if o == nil || common.IsNil(o.Rows) {
		var ret []GetCloudMiningPaymentAndRefundHistoryResponseRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudMiningPaymentAndRefundHistoryResponse) GetRowsOk() ([]GetCloudMiningPaymentAndRefundHistoryResponseRowsInner, bool) {
	if o == nil || common.IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *GetCloudMiningPaymentAndRefundHistoryResponse) HasRows() bool {
	if o != nil && !common.IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []GetCloudMiningPaymentAndRefundHistoryResponseRowsInner and assigns it to the Rows field.
func (o *GetCloudMiningPaymentAndRefundHistoryResponse) SetRows(v []GetCloudMiningPaymentAndRefundHistoryResponseRowsInner) {
	o.Rows = v
}

func (o GetCloudMiningPaymentAndRefundHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCloudMiningPaymentAndRefundHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetCloudMiningPaymentAndRefundHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetCloudMiningPaymentAndRefundHistoryResponse := _GetCloudMiningPaymentAndRefundHistoryResponse{}

	err = json.Unmarshal(data, &varGetCloudMiningPaymentAndRefundHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetCloudMiningPaymentAndRefundHistoryResponse(varGetCloudMiningPaymentAndRefundHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "rows")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCloudMiningPaymentAndRefundHistoryResponse struct {
	value *GetCloudMiningPaymentAndRefundHistoryResponse
	isSet bool
}

func (v NullableGetCloudMiningPaymentAndRefundHistoryResponse) Get() *GetCloudMiningPaymentAndRefundHistoryResponse {
	return v.value
}

func (v *NullableGetCloudMiningPaymentAndRefundHistoryResponse) Set(val *GetCloudMiningPaymentAndRefundHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCloudMiningPaymentAndRefundHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCloudMiningPaymentAndRefundHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCloudMiningPaymentAndRefundHistoryResponse(val *GetCloudMiningPaymentAndRefundHistoryResponse) *NullableGetCloudMiningPaymentAndRefundHistoryResponse {
	return &NullableGetCloudMiningPaymentAndRefundHistoryResponse{value: val, isSet: true}
}

func (v NullableGetCloudMiningPaymentAndRefundHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCloudMiningPaymentAndRefundHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
