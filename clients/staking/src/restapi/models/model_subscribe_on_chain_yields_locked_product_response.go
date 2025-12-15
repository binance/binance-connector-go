/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SubscribeOnChainYieldsLockedProductResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubscribeOnChainYieldsLockedProductResponse{}

// SubscribeOnChainYieldsLockedProductResponse struct for SubscribeOnChainYieldsLockedProductResponse
type SubscribeOnChainYieldsLockedProductResponse struct {
	PurchaseId           *int64  `json:"purchaseId,omitempty"`
	PositionId           *string `json:"positionId,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	Success              *bool   `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscribeOnChainYieldsLockedProductResponse SubscribeOnChainYieldsLockedProductResponse

// NewSubscribeOnChainYieldsLockedProductResponse instantiates a new SubscribeOnChainYieldsLockedProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribeOnChainYieldsLockedProductResponse() *SubscribeOnChainYieldsLockedProductResponse {
	this := SubscribeOnChainYieldsLockedProductResponse{}
	return &this
}

// NewSubscribeOnChainYieldsLockedProductResponseWithDefaults instantiates a new SubscribeOnChainYieldsLockedProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribeOnChainYieldsLockedProductResponseWithDefaults() *SubscribeOnChainYieldsLockedProductResponse {
	this := SubscribeOnChainYieldsLockedProductResponse{}
	return &this
}

// GetPurchaseId returns the PurchaseId field value if set, zero value otherwise.
func (o *SubscribeOnChainYieldsLockedProductResponse) GetPurchaseId() int64 {
	if o == nil || common.IsNil(o.PurchaseId) {
		var ret int64
		return ret
	}
	return *o.PurchaseId
}

// GetPurchaseIdOk returns a tuple with the PurchaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeOnChainYieldsLockedProductResponse) GetPurchaseIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PurchaseId) {
		return nil, false
	}
	return o.PurchaseId, true
}

// HasPurchaseId returns a boolean if a field has been set.
func (o *SubscribeOnChainYieldsLockedProductResponse) HasPurchaseId() bool {
	if o != nil && !common.IsNil(o.PurchaseId) {
		return true
	}

	return false
}

// SetPurchaseId gets a reference to the given int64 and assigns it to the PurchaseId field.
func (o *SubscribeOnChainYieldsLockedProductResponse) SetPurchaseId(v int64) {
	o.PurchaseId = &v
}

// GetPositionId returns the PositionId field value if set, zero value otherwise.
func (o *SubscribeOnChainYieldsLockedProductResponse) GetPositionId() string {
	if o == nil || common.IsNil(o.PositionId) {
		var ret string
		return ret
	}
	return *o.PositionId
}

// GetPositionIdOk returns a tuple with the PositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeOnChainYieldsLockedProductResponse) GetPositionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionId) {
		return nil, false
	}
	return o.PositionId, true
}

// HasPositionId returns a boolean if a field has been set.
func (o *SubscribeOnChainYieldsLockedProductResponse) HasPositionId() bool {
	if o != nil && !common.IsNil(o.PositionId) {
		return true
	}

	return false
}

// SetPositionId gets a reference to the given string and assigns it to the PositionId field.
func (o *SubscribeOnChainYieldsLockedProductResponse) SetPositionId(v string) {
	o.PositionId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SubscribeOnChainYieldsLockedProductResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeOnChainYieldsLockedProductResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SubscribeOnChainYieldsLockedProductResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *SubscribeOnChainYieldsLockedProductResponse) SetAmount(v string) {
	o.Amount = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SubscribeOnChainYieldsLockedProductResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeOnChainYieldsLockedProductResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SubscribeOnChainYieldsLockedProductResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SubscribeOnChainYieldsLockedProductResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o SubscribeOnChainYieldsLockedProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscribeOnChainYieldsLockedProductResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PurchaseId) {
		toSerialize["purchaseId"] = o.PurchaseId
	}
	if !common.IsNil(o.PositionId) {
		toSerialize["positionId"] = o.PositionId
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscribeOnChainYieldsLockedProductResponse) UnmarshalJSON(data []byte) (err error) {
	varSubscribeOnChainYieldsLockedProductResponse := _SubscribeOnChainYieldsLockedProductResponse{}

	err = json.Unmarshal(data, &varSubscribeOnChainYieldsLockedProductResponse)

	if err != nil {
		return err
	}

	*o = SubscribeOnChainYieldsLockedProductResponse(varSubscribeOnChainYieldsLockedProductResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "purchaseId")
		delete(additionalProperties, "positionId")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscribeOnChainYieldsLockedProductResponse struct {
	value *SubscribeOnChainYieldsLockedProductResponse
	isSet bool
}

func (v NullableSubscribeOnChainYieldsLockedProductResponse) Get() *SubscribeOnChainYieldsLockedProductResponse {
	return v.value
}

func (v *NullableSubscribeOnChainYieldsLockedProductResponse) Set(val *SubscribeOnChainYieldsLockedProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribeOnChainYieldsLockedProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribeOnChainYieldsLockedProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribeOnChainYieldsLockedProductResponse(val *SubscribeOnChainYieldsLockedProductResponse) *NullableSubscribeOnChainYieldsLockedProductResponse {
	return &NullableSubscribeOnChainYieldsLockedProductResponse{value: val, isSet: true}
}

func (v NullableSubscribeOnChainYieldsLockedProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribeOnChainYieldsLockedProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
