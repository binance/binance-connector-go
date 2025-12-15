/*
Binance Convert REST API

OpenAPI Specification for the Binance Convert REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CancelLimitOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelLimitOrderResponse{}

// CancelLimitOrderResponse struct for CancelLimitOrderResponse
type CancelLimitOrderResponse struct {
	OrderId              *int64  `json:"orderId,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelLimitOrderResponse CancelLimitOrderResponse

// NewCancelLimitOrderResponse instantiates a new CancelLimitOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelLimitOrderResponse() *CancelLimitOrderResponse {
	this := CancelLimitOrderResponse{}
	return &this
}

// NewCancelLimitOrderResponseWithDefaults instantiates a new CancelLimitOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelLimitOrderResponseWithDefaults() *CancelLimitOrderResponse {
	this := CancelLimitOrderResponse{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *CancelLimitOrderResponse) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelLimitOrderResponse) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *CancelLimitOrderResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *CancelLimitOrderResponse) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CancelLimitOrderResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelLimitOrderResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CancelLimitOrderResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CancelLimitOrderResponse) SetStatus(v string) {
	o.Status = &v
}

func (o CancelLimitOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelLimitOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CancelLimitOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varCancelLimitOrderResponse := _CancelLimitOrderResponse{}

	err = json.Unmarshal(data, &varCancelLimitOrderResponse)

	if err != nil {
		return err
	}

	*o = CancelLimitOrderResponse(varCancelLimitOrderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelLimitOrderResponse struct {
	value *CancelLimitOrderResponse
	isSet bool
}

func (v NullableCancelLimitOrderResponse) Get() *CancelLimitOrderResponse {
	return v.value
}

func (v *NullableCancelLimitOrderResponse) Set(val *CancelLimitOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelLimitOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelLimitOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelLimitOrderResponse(val *CancelLimitOrderResponse) *NullableCancelLimitOrderResponse {
	return &NullableCancelLimitOrderResponse{value: val, isSet: true}
}

func (v NullableCancelLimitOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelLimitOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
