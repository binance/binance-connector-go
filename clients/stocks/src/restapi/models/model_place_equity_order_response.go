/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PlaceEquityOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PlaceEquityOrderResponse{}

// PlaceEquityOrderResponse struct for PlaceEquityOrderResponse
type PlaceEquityOrderResponse struct {
	// Acknowledgement code: `S` = accepted, `F` = failed. Not an order lifecycle status — to poll lifecycle, call `/order/detail` or `/order/history`.
	Status *string `json:"status,omitempty"`
	// Order id (UUID).
	OrderId *string `json:"orderId,omitempty"`
	// Echoes the supplied or server-generated client order id.
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlaceEquityOrderResponse PlaceEquityOrderResponse

// NewPlaceEquityOrderResponse instantiates a new PlaceEquityOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceEquityOrderResponse() *PlaceEquityOrderResponse {
	this := PlaceEquityOrderResponse{}
	return &this
}

// NewPlaceEquityOrderResponseWithDefaults instantiates a new PlaceEquityOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceEquityOrderResponseWithDefaults() *PlaceEquityOrderResponse {
	this := PlaceEquityOrderResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PlaceEquityOrderResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceEquityOrderResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PlaceEquityOrderResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PlaceEquityOrderResponse) SetStatus(v string) {
	o.Status = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *PlaceEquityOrderResponse) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceEquityOrderResponse) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *PlaceEquityOrderResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *PlaceEquityOrderResponse) SetOrderId(v string) {
	o.OrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *PlaceEquityOrderResponse) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceEquityOrderResponse) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *PlaceEquityOrderResponse) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *PlaceEquityOrderResponse) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

func (o PlaceEquityOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaceEquityOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlaceEquityOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varPlaceEquityOrderResponse := _PlaceEquityOrderResponse{}

	err = json.Unmarshal(data, &varPlaceEquityOrderResponse)

	if err != nil {
		return err
	}

	*o = PlaceEquityOrderResponse(varPlaceEquityOrderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "clientOrderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlaceEquityOrderResponse struct {
	value *PlaceEquityOrderResponse
	isSet bool
}

func (v NullablePlaceEquityOrderResponse) Get() *PlaceEquityOrderResponse {
	return v.value
}

func (v *NullablePlaceEquityOrderResponse) Set(val *PlaceEquityOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceEquityOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceEquityOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceEquityOrderResponse(val *PlaceEquityOrderResponse) *NullablePlaceEquityOrderResponse {
	return &NullablePlaceEquityOrderResponse{value: val, isSet: true}
}

func (v NullablePlaceEquityOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceEquityOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
