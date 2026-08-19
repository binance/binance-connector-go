/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CancelAllEquityOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelAllEquityOrdersResponse{}

// CancelAllEquityOrdersResponse struct for CancelAllEquityOrdersResponse
type CancelAllEquityOrdersResponse struct {
	// `true` when the cancel-all request was accepted by upstream. Per-order outcomes live in `/order/history`.
	Success              *bool `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelAllEquityOrdersResponse CancelAllEquityOrdersResponse

// NewCancelAllEquityOrdersResponse instantiates a new CancelAllEquityOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelAllEquityOrdersResponse() *CancelAllEquityOrdersResponse {
	this := CancelAllEquityOrdersResponse{}
	return &this
}

// NewCancelAllEquityOrdersResponseWithDefaults instantiates a new CancelAllEquityOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelAllEquityOrdersResponseWithDefaults() *CancelAllEquityOrdersResponse {
	this := CancelAllEquityOrdersResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CancelAllEquityOrdersResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAllEquityOrdersResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CancelAllEquityOrdersResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CancelAllEquityOrdersResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o CancelAllEquityOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelAllEquityOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CancelAllEquityOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	varCancelAllEquityOrdersResponse := _CancelAllEquityOrdersResponse{}

	err = json.Unmarshal(data, &varCancelAllEquityOrdersResponse)

	if err != nil {
		return err
	}

	*o = CancelAllEquityOrdersResponse(varCancelAllEquityOrdersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelAllEquityOrdersResponse struct {
	value *CancelAllEquityOrdersResponse
	isSet bool
}

func (v NullableCancelAllEquityOrdersResponse) Get() *CancelAllEquityOrdersResponse {
	return v.value
}

func (v *NullableCancelAllEquityOrdersResponse) Set(val *CancelAllEquityOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelAllEquityOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelAllEquityOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelAllEquityOrdersResponse(val *CancelAllEquityOrdersResponse) *NullableCancelAllEquityOrdersResponse {
	return &NullableCancelAllEquityOrdersResponse{value: val, isSet: true}
}

func (v NullableCancelAllEquityOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelAllEquityOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
