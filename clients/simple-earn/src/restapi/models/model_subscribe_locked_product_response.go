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

// checks if the SubscribeLockedProductResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubscribeLockedProductResponse{}

// SubscribeLockedProductResponse struct for SubscribeLockedProductResponse
type SubscribeLockedProductResponse struct {
	PurchaseId           *int64  `json:"purchaseId,omitempty"`
	PositionId           *string `json:"positionId,omitempty"`
	Success              *bool   `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscribeLockedProductResponse SubscribeLockedProductResponse

// NewSubscribeLockedProductResponse instantiates a new SubscribeLockedProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribeLockedProductResponse() *SubscribeLockedProductResponse {
	this := SubscribeLockedProductResponse{}
	return &this
}

// NewSubscribeLockedProductResponseWithDefaults instantiates a new SubscribeLockedProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribeLockedProductResponseWithDefaults() *SubscribeLockedProductResponse {
	this := SubscribeLockedProductResponse{}
	return &this
}

// GetPurchaseId returns the PurchaseId field value if set, zero value otherwise.
func (o *SubscribeLockedProductResponse) GetPurchaseId() int64 {
	if o == nil || common.IsNil(o.PurchaseId) {
		var ret int64
		return ret
	}
	return *o.PurchaseId
}

// GetPurchaseIdOk returns a tuple with the PurchaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeLockedProductResponse) GetPurchaseIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PurchaseId) {
		return nil, false
	}
	return o.PurchaseId, true
}

// HasPurchaseId returns a boolean if a field has been set.
func (o *SubscribeLockedProductResponse) HasPurchaseId() bool {
	if o != nil && !common.IsNil(o.PurchaseId) {
		return true
	}

	return false
}

// SetPurchaseId gets a reference to the given int64 and assigns it to the PurchaseId field.
func (o *SubscribeLockedProductResponse) SetPurchaseId(v int64) {
	o.PurchaseId = &v
}

// GetPositionId returns the PositionId field value if set, zero value otherwise.
func (o *SubscribeLockedProductResponse) GetPositionId() string {
	if o == nil || common.IsNil(o.PositionId) {
		var ret string
		return ret
	}
	return *o.PositionId
}

// GetPositionIdOk returns a tuple with the PositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeLockedProductResponse) GetPositionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionId) {
		return nil, false
	}
	return o.PositionId, true
}

// HasPositionId returns a boolean if a field has been set.
func (o *SubscribeLockedProductResponse) HasPositionId() bool {
	if o != nil && !common.IsNil(o.PositionId) {
		return true
	}

	return false
}

// SetPositionId gets a reference to the given string and assigns it to the PositionId field.
func (o *SubscribeLockedProductResponse) SetPositionId(v string) {
	o.PositionId = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SubscribeLockedProductResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeLockedProductResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SubscribeLockedProductResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SubscribeLockedProductResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o SubscribeLockedProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscribeLockedProductResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PurchaseId) {
		toSerialize["purchaseId"] = o.PurchaseId
	}
	if !common.IsNil(o.PositionId) {
		toSerialize["positionId"] = o.PositionId
	}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscribeLockedProductResponse) UnmarshalJSON(data []byte) (err error) {
	varSubscribeLockedProductResponse := _SubscribeLockedProductResponse{}

	err = json.Unmarshal(data, &varSubscribeLockedProductResponse)

	if err != nil {
		return err
	}

	*o = SubscribeLockedProductResponse(varSubscribeLockedProductResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "purchaseId")
		delete(additionalProperties, "positionId")
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscribeLockedProductResponse struct {
	value *SubscribeLockedProductResponse
	isSet bool
}

func (v NullableSubscribeLockedProductResponse) Get() *SubscribeLockedProductResponse {
	return v.value
}

func (v *NullableSubscribeLockedProductResponse) Set(val *SubscribeLockedProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribeLockedProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribeLockedProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribeLockedProductResponse(val *SubscribeLockedProductResponse) *NullableSubscribeLockedProductResponse {
	return &NullableSubscribeLockedProductResponse{value: val, isSet: true}
}

func (v NullableSubscribeLockedProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribeLockedProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
