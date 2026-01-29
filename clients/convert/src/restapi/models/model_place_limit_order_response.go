/*
Binance Convert REST API

OpenAPI Specification for the Binance Convert REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PlaceLimitOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PlaceLimitOrderResponse{}

// PlaceLimitOrderResponse struct for PlaceLimitOrderResponse
type PlaceLimitOrderResponse struct {
	OrderId              *int64  `json:"orderId,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlaceLimitOrderResponse PlaceLimitOrderResponse

// NewPlaceLimitOrderResponse instantiates a new PlaceLimitOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceLimitOrderResponse() *PlaceLimitOrderResponse {
	this := PlaceLimitOrderResponse{}
	return &this
}

// NewPlaceLimitOrderResponseWithDefaults instantiates a new PlaceLimitOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceLimitOrderResponseWithDefaults() *PlaceLimitOrderResponse {
	this := PlaceLimitOrderResponse{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *PlaceLimitOrderResponse) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceLimitOrderResponse) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *PlaceLimitOrderResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *PlaceLimitOrderResponse) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PlaceLimitOrderResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceLimitOrderResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PlaceLimitOrderResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PlaceLimitOrderResponse) SetStatus(v string) {
	o.Status = &v
}

func (o PlaceLimitOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaceLimitOrderResponse) ToMap() (map[string]interface{}, error) {
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

func (o *PlaceLimitOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varPlaceLimitOrderResponse := _PlaceLimitOrderResponse{}

	err = json.Unmarshal(data, &varPlaceLimitOrderResponse)

	if err != nil {
		return err
	}

	*o = PlaceLimitOrderResponse(varPlaceLimitOrderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlaceLimitOrderResponse struct {
	value *PlaceLimitOrderResponse
	isSet bool
}

func (v NullablePlaceLimitOrderResponse) Get() *PlaceLimitOrderResponse {
	return v.value
}

func (v *NullablePlaceLimitOrderResponse) Set(val *PlaceLimitOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceLimitOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceLimitOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceLimitOrderResponse(val *PlaceLimitOrderResponse) *NullablePlaceLimitOrderResponse {
	return &NullablePlaceLimitOrderResponse{value: val, isSet: true}
}

func (v NullablePlaceLimitOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceLimitOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
