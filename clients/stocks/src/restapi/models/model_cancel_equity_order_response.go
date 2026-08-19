/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CancelEquityOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelEquityOrderResponse{}

// CancelEquityOrderResponse struct for CancelEquityOrderResponse
type CancelEquityOrderResponse struct {
	// Echoes the requested order id.
	OrderId *string `json:"orderId,omitempty"`
	// Acknowledgement code: `S` = cancel accepted by upstream, `F` = cancel failed. Not an order lifecycle status — use `/order/detail` for lifecycle poll.
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelEquityOrderResponse CancelEquityOrderResponse

// NewCancelEquityOrderResponse instantiates a new CancelEquityOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelEquityOrderResponse() *CancelEquityOrderResponse {
	this := CancelEquityOrderResponse{}
	return &this
}

// NewCancelEquityOrderResponseWithDefaults instantiates a new CancelEquityOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelEquityOrderResponseWithDefaults() *CancelEquityOrderResponse {
	this := CancelEquityOrderResponse{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *CancelEquityOrderResponse) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelEquityOrderResponse) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *CancelEquityOrderResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *CancelEquityOrderResponse) SetOrderId(v string) {
	o.OrderId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CancelEquityOrderResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelEquityOrderResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CancelEquityOrderResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CancelEquityOrderResponse) SetStatus(v string) {
	o.Status = &v
}

func (o CancelEquityOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelEquityOrderResponse) ToMap() (map[string]interface{}, error) {
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

func (o *CancelEquityOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varCancelEquityOrderResponse := _CancelEquityOrderResponse{}

	err = json.Unmarshal(data, &varCancelEquityOrderResponse)

	if err != nil {
		return err
	}

	*o = CancelEquityOrderResponse(varCancelEquityOrderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelEquityOrderResponse struct {
	value *CancelEquityOrderResponse
	isSet bool
}

func (v NullableCancelEquityOrderResponse) Get() *CancelEquityOrderResponse {
	return v.value
}

func (v *NullableCancelEquityOrderResponse) Set(val *CancelEquityOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelEquityOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelEquityOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelEquityOrderResponse(val *CancelEquityOrderResponse) *NullableCancelEquityOrderResponse {
	return &NullableCancelEquityOrderResponse{value: val, isSet: true}
}

func (v NullableCancelEquityOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelEquityOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
