/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOtcBlocktradeDetailResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOtcBlocktradeDetailResponse{}

// GetOtcBlocktradeDetailResponse struct for GetOtcBlocktradeDetailResponse
type GetOtcBlocktradeDetailResponse struct {
	OrderId              *string                                  `json:"orderId,omitempty"`
	Status               *string                                  `json:"status,omitempty"`
	OrderData            *GetOtcBlocktradeDetailResponseOrderData `json:"orderData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOtcBlocktradeDetailResponse GetOtcBlocktradeDetailResponse

// NewGetOtcBlocktradeDetailResponse instantiates a new GetOtcBlocktradeDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOtcBlocktradeDetailResponse() *GetOtcBlocktradeDetailResponse {
	this := GetOtcBlocktradeDetailResponse{}
	return &this
}

// NewGetOtcBlocktradeDetailResponseWithDefaults instantiates a new GetOtcBlocktradeDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOtcBlocktradeDetailResponseWithDefaults() *GetOtcBlocktradeDetailResponse {
	this := GetOtcBlocktradeDetailResponse{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponse) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponse) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *GetOtcBlocktradeDetailResponse) SetOrderId(v string) {
	o.OrderId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetOtcBlocktradeDetailResponse) SetStatus(v string) {
	o.Status = &v
}

// GetOrderData returns the OrderData field value if set, zero value otherwise.
func (o *GetOtcBlocktradeDetailResponse) GetOrderData() GetOtcBlocktradeDetailResponseOrderData {
	if o == nil || common.IsNil(o.OrderData) {
		var ret GetOtcBlocktradeDetailResponseOrderData
		return ret
	}
	return *o.OrderData
}

// GetOrderDataOk returns a tuple with the OrderData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeDetailResponse) GetOrderDataOk() (*GetOtcBlocktradeDetailResponseOrderData, bool) {
	if o == nil || common.IsNil(o.OrderData) {
		return nil, false
	}
	return o.OrderData, true
}

// HasOrderData returns a boolean if a field has been set.
func (o *GetOtcBlocktradeDetailResponse) HasOrderData() bool {
	if o != nil && !common.IsNil(o.OrderData) {
		return true
	}

	return false
}

// SetOrderData gets a reference to the given GetOtcBlocktradeDetailResponseOrderData and assigns it to the OrderData field.
func (o *GetOtcBlocktradeDetailResponse) SetOrderData(v GetOtcBlocktradeDetailResponseOrderData) {
	o.OrderData = &v
}

func (o GetOtcBlocktradeDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOtcBlocktradeDetailResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetOtcBlocktradeDetailResponse) UnmarshalJSON(data []byte) (err error) {
	varGetOtcBlocktradeDetailResponse := _GetOtcBlocktradeDetailResponse{}

	err = json.Unmarshal(data, &varGetOtcBlocktradeDetailResponse)

	if err != nil {
		return err
	}

	*o = GetOtcBlocktradeDetailResponse(varGetOtcBlocktradeDetailResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "orderData")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOtcBlocktradeDetailResponse struct {
	value *GetOtcBlocktradeDetailResponse
	isSet bool
}

func (v NullableGetOtcBlocktradeDetailResponse) Get() *GetOtcBlocktradeDetailResponse {
	return v.value
}

func (v *NullableGetOtcBlocktradeDetailResponse) Set(val *GetOtcBlocktradeDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOtcBlocktradeDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOtcBlocktradeDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOtcBlocktradeDetailResponse(val *GetOtcBlocktradeDetailResponse) *NullableGetOtcBlocktradeDetailResponse {
	return &NullableGetOtcBlocktradeDetailResponse{value: val, isSet: true}
}

func (v NullableGetOtcBlocktradeDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOtcBlocktradeDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
