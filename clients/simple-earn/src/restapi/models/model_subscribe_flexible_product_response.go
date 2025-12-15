/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SubscribeFlexibleProductResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubscribeFlexibleProductResponse{}

// SubscribeFlexibleProductResponse struct for SubscribeFlexibleProductResponse
type SubscribeFlexibleProductResponse struct {
	PurchaseId           *int64 `json:"purchaseId,omitempty"`
	Success              *bool  `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscribeFlexibleProductResponse SubscribeFlexibleProductResponse

// NewSubscribeFlexibleProductResponse instantiates a new SubscribeFlexibleProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribeFlexibleProductResponse() *SubscribeFlexibleProductResponse {
	this := SubscribeFlexibleProductResponse{}
	return &this
}

// NewSubscribeFlexibleProductResponseWithDefaults instantiates a new SubscribeFlexibleProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribeFlexibleProductResponseWithDefaults() *SubscribeFlexibleProductResponse {
	this := SubscribeFlexibleProductResponse{}
	return &this
}

// GetPurchaseId returns the PurchaseId field value if set, zero value otherwise.
func (o *SubscribeFlexibleProductResponse) GetPurchaseId() int64 {
	if o == nil || common.IsNil(o.PurchaseId) {
		var ret int64
		return ret
	}
	return *o.PurchaseId
}

// GetPurchaseIdOk returns a tuple with the PurchaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeFlexibleProductResponse) GetPurchaseIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PurchaseId) {
		return nil, false
	}
	return o.PurchaseId, true
}

// HasPurchaseId returns a boolean if a field has been set.
func (o *SubscribeFlexibleProductResponse) HasPurchaseId() bool {
	if o != nil && !common.IsNil(o.PurchaseId) {
		return true
	}

	return false
}

// SetPurchaseId gets a reference to the given int64 and assigns it to the PurchaseId field.
func (o *SubscribeFlexibleProductResponse) SetPurchaseId(v int64) {
	o.PurchaseId = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SubscribeFlexibleProductResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeFlexibleProductResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SubscribeFlexibleProductResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SubscribeFlexibleProductResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o SubscribeFlexibleProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscribeFlexibleProductResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PurchaseId) {
		toSerialize["purchaseId"] = o.PurchaseId
	}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscribeFlexibleProductResponse) UnmarshalJSON(data []byte) (err error) {
	varSubscribeFlexibleProductResponse := _SubscribeFlexibleProductResponse{}

	err = json.Unmarshal(data, &varSubscribeFlexibleProductResponse)

	if err != nil {
		return err
	}

	*o = SubscribeFlexibleProductResponse(varSubscribeFlexibleProductResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "purchaseId")
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscribeFlexibleProductResponse struct {
	value *SubscribeFlexibleProductResponse
	isSet bool
}

func (v NullableSubscribeFlexibleProductResponse) Get() *SubscribeFlexibleProductResponse {
	return v.value
}

func (v *NullableSubscribeFlexibleProductResponse) Set(val *SubscribeFlexibleProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribeFlexibleProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribeFlexibleProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribeFlexibleProductResponse(val *SubscribeFlexibleProductResponse) *NullableSubscribeFlexibleProductResponse {
	return &NullableSubscribeFlexibleProductResponse{value: val, isSet: true}
}

func (v NullableSubscribeFlexibleProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribeFlexibleProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
