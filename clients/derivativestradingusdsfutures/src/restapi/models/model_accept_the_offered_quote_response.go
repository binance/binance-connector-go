/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AcceptTheOfferedQuoteResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AcceptTheOfferedQuoteResponse{}

// AcceptTheOfferedQuoteResponse struct for AcceptTheOfferedQuoteResponse
type AcceptTheOfferedQuoteResponse struct {
	OrderId              *string `json:"orderId,omitempty"`
	CreateTime           *int64  `json:"createTime,omitempty"`
	OrderStatus          *string `json:"orderStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AcceptTheOfferedQuoteResponse AcceptTheOfferedQuoteResponse

// NewAcceptTheOfferedQuoteResponse instantiates a new AcceptTheOfferedQuoteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptTheOfferedQuoteResponse() *AcceptTheOfferedQuoteResponse {
	this := AcceptTheOfferedQuoteResponse{}
	return &this
}

// NewAcceptTheOfferedQuoteResponseWithDefaults instantiates a new AcceptTheOfferedQuoteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptTheOfferedQuoteResponseWithDefaults() *AcceptTheOfferedQuoteResponse {
	this := AcceptTheOfferedQuoteResponse{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *AcceptTheOfferedQuoteResponse) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptTheOfferedQuoteResponse) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *AcceptTheOfferedQuoteResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *AcceptTheOfferedQuoteResponse) SetOrderId(v string) {
	o.OrderId = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *AcceptTheOfferedQuoteResponse) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptTheOfferedQuoteResponse) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *AcceptTheOfferedQuoteResponse) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *AcceptTheOfferedQuoteResponse) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *AcceptTheOfferedQuoteResponse) GetOrderStatus() string {
	if o == nil || common.IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptTheOfferedQuoteResponse) GetOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *AcceptTheOfferedQuoteResponse) HasOrderStatus() bool {
	if o != nil && !common.IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *AcceptTheOfferedQuoteResponse) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

func (o AcceptTheOfferedQuoteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceptTheOfferedQuoteResponse) ToMap() (map[string]interface{}, error) {
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

func (o *AcceptTheOfferedQuoteResponse) UnmarshalJSON(data []byte) (err error) {
	varAcceptTheOfferedQuoteResponse := _AcceptTheOfferedQuoteResponse{}

	err = json.Unmarshal(data, &varAcceptTheOfferedQuoteResponse)

	if err != nil {
		return err
	}

	*o = AcceptTheOfferedQuoteResponse(varAcceptTheOfferedQuoteResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "orderStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcceptTheOfferedQuoteResponse struct {
	value *AcceptTheOfferedQuoteResponse
	isSet bool
}

func (v NullableAcceptTheOfferedQuoteResponse) Get() *AcceptTheOfferedQuoteResponse {
	return v.value
}

func (v *NullableAcceptTheOfferedQuoteResponse) Set(val *AcceptTheOfferedQuoteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptTheOfferedQuoteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptTheOfferedQuoteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptTheOfferedQuoteResponse(val *AcceptTheOfferedQuoteResponse) *NullableAcceptTheOfferedQuoteResponse {
	return &NullableAcceptTheOfferedQuoteResponse{value: val, isSet: true}
}

func (v NullableAcceptTheOfferedQuoteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptTheOfferedQuoteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
