/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PreviewOtcBlocktradeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PreviewOtcBlocktradeResponse{}

// PreviewOtcBlocktradeResponse struct for PreviewOtcBlocktradeResponse
type PreviewOtcBlocktradeResponse struct {
	OrderId              *string                                `json:"orderId,omitempty"`
	Status               *string                                `json:"status,omitempty"`
	OrderData            *PreviewOtcBlocktradeResponseOrderData `json:"orderData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PreviewOtcBlocktradeResponse PreviewOtcBlocktradeResponse

// NewPreviewOtcBlocktradeResponse instantiates a new PreviewOtcBlocktradeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviewOtcBlocktradeResponse() *PreviewOtcBlocktradeResponse {
	this := PreviewOtcBlocktradeResponse{}
	return &this
}

// NewPreviewOtcBlocktradeResponseWithDefaults instantiates a new PreviewOtcBlocktradeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviewOtcBlocktradeResponseWithDefaults() *PreviewOtcBlocktradeResponse {
	this := PreviewOtcBlocktradeResponse{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PreviewOtcBlocktradeResponse) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreviewOtcBlocktradeResponse) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *PreviewOtcBlocktradeResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given NullableString and assigns it to the OrderId field.
func (o *PreviewOtcBlocktradeResponse) SetOrderId(v string) {
	o.OrderId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PreviewOtcBlocktradeResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviewOtcBlocktradeResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PreviewOtcBlocktradeResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PreviewOtcBlocktradeResponse) SetStatus(v string) {
	o.Status = &v
}

// GetOrderData returns the OrderData field value if set, zero value otherwise.
func (o *PreviewOtcBlocktradeResponse) GetOrderData() PreviewOtcBlocktradeResponseOrderData {
	if o == nil || common.IsNil(o.OrderData) {
		var ret PreviewOtcBlocktradeResponseOrderData
		return ret
	}
	return *o.OrderData
}

// GetOrderDataOk returns a tuple with the OrderData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviewOtcBlocktradeResponse) GetOrderDataOk() (*PreviewOtcBlocktradeResponseOrderData, bool) {
	if o == nil || common.IsNil(o.OrderData) {
		return nil, false
	}
	return o.OrderData, true
}

// HasOrderData returns a boolean if a field has been set.
func (o *PreviewOtcBlocktradeResponse) HasOrderData() bool {
	if o != nil && !common.IsNil(o.OrderData) {
		return true
	}

	return false
}

// SetOrderData gets a reference to the given PreviewOtcBlocktradeResponseOrderData and assigns it to the OrderData field.
func (o *PreviewOtcBlocktradeResponse) SetOrderData(v PreviewOtcBlocktradeResponseOrderData) {
	o.OrderData = &v
}

func (o PreviewOtcBlocktradeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreviewOtcBlocktradeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.OrderData) {
		toSerialize["orderData"] = o.OrderData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PreviewOtcBlocktradeResponse) UnmarshalJSON(data []byte) (err error) {
	varPreviewOtcBlocktradeResponse := _PreviewOtcBlocktradeResponse{}

	err = json.Unmarshal(data, &varPreviewOtcBlocktradeResponse)

	if err != nil {
		return err
	}

	*o = PreviewOtcBlocktradeResponse(varPreviewOtcBlocktradeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "orderData")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePreviewOtcBlocktradeResponse struct {
	value *PreviewOtcBlocktradeResponse
	isSet bool
}

func (v NullablePreviewOtcBlocktradeResponse) Get() *PreviewOtcBlocktradeResponse {
	return v.value
}

func (v *NullablePreviewOtcBlocktradeResponse) Set(val *PreviewOtcBlocktradeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviewOtcBlocktradeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviewOtcBlocktradeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviewOtcBlocktradeResponse(val *PreviewOtcBlocktradeResponse) *NullablePreviewOtcBlocktradeResponse {
	return &NullablePreviewOtcBlocktradeResponse{value: val, isSet: true}
}

func (v NullablePreviewOtcBlocktradeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreviewOtcBlocktradeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
