/*
Binance Convert REST API

OpenAPI Specification for the Binance Convert REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AcceptQuoteResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AcceptQuoteResponse{}

// AcceptQuoteResponse struct for AcceptQuoteResponse
type AcceptQuoteResponse struct {
	OrderId              *string `json:"orderId,omitempty"`
	CreateTime           *int64  `json:"createTime,omitempty"`
	OrderStatus          *string `json:"orderStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AcceptQuoteResponse AcceptQuoteResponse

// NewAcceptQuoteResponse instantiates a new AcceptQuoteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptQuoteResponse() *AcceptQuoteResponse {
	this := AcceptQuoteResponse{}
	return &this
}

// NewAcceptQuoteResponseWithDefaults instantiates a new AcceptQuoteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptQuoteResponseWithDefaults() *AcceptQuoteResponse {
	this := AcceptQuoteResponse{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *AcceptQuoteResponse) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptQuoteResponse) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *AcceptQuoteResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *AcceptQuoteResponse) SetOrderId(v string) {
	o.OrderId = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *AcceptQuoteResponse) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptQuoteResponse) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *AcceptQuoteResponse) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *AcceptQuoteResponse) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *AcceptQuoteResponse) GetOrderStatus() string {
	if o == nil || common.IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptQuoteResponse) GetOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *AcceptQuoteResponse) HasOrderStatus() bool {
	if o != nil && !common.IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *AcceptQuoteResponse) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

func (o AcceptQuoteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceptQuoteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !common.IsNil(o.OrderStatus) {
		toSerialize["orderStatus"] = o.OrderStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AcceptQuoteResponse) UnmarshalJSON(data []byte) (err error) {
	varAcceptQuoteResponse := _AcceptQuoteResponse{}

	err = json.Unmarshal(data, &varAcceptQuoteResponse)

	if err != nil {
		return err
	}

	*o = AcceptQuoteResponse(varAcceptQuoteResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "orderStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcceptQuoteResponse struct {
	value *AcceptQuoteResponse
	isSet bool
}

func (v NullableAcceptQuoteResponse) Get() *AcceptQuoteResponse {
	return v.value
}

func (v *NullableAcceptQuoteResponse) Set(val *AcceptQuoteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptQuoteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptQuoteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptQuoteResponse(val *AcceptQuoteResponse) *NullableAcceptQuoteResponse {
	return &NullableAcceptQuoteResponse{value: val, isSet: true}
}

func (v NullableAcceptQuoteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptQuoteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
