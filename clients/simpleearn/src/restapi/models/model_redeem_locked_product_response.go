/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the RedeemLockedProductResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RedeemLockedProductResponse{}

// RedeemLockedProductResponse struct for RedeemLockedProductResponse
type RedeemLockedProductResponse struct {
	RedeemId             *int64 `json:"redeemId,omitempty"`
	Success              *bool  `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RedeemLockedProductResponse RedeemLockedProductResponse

// NewRedeemLockedProductResponse instantiates a new RedeemLockedProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedeemLockedProductResponse() *RedeemLockedProductResponse {
	this := RedeemLockedProductResponse{}
	return &this
}

// NewRedeemLockedProductResponseWithDefaults instantiates a new RedeemLockedProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedeemLockedProductResponseWithDefaults() *RedeemLockedProductResponse {
	this := RedeemLockedProductResponse{}
	return &this
}

// GetRedeemId returns the RedeemId field value if set, zero value otherwise.
func (o *RedeemLockedProductResponse) GetRedeemId() int64 {
	if o == nil || common.IsNil(o.RedeemId) {
		var ret int64
		return ret
	}
	return *o.RedeemId
}

// GetRedeemIdOk returns a tuple with the RedeemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemLockedProductResponse) GetRedeemIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RedeemId) {
		return nil, false
	}
	return o.RedeemId, true
}

// HasRedeemId returns a boolean if a field has been set.
func (o *RedeemLockedProductResponse) HasRedeemId() bool {
	if o != nil && !common.IsNil(o.RedeemId) {
		return true
	}

	return false
}

// SetRedeemId gets a reference to the given int64 and assigns it to the RedeemId field.
func (o *RedeemLockedProductResponse) SetRedeemId(v int64) {
	o.RedeemId = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *RedeemLockedProductResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemLockedProductResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *RedeemLockedProductResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *RedeemLockedProductResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o RedeemLockedProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedeemLockedProductResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RedeemId) {
		toSerialize["redeemId"] = o.RedeemId
	}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RedeemLockedProductResponse) UnmarshalJSON(data []byte) (err error) {
	varRedeemLockedProductResponse := _RedeemLockedProductResponse{}

	err = json.Unmarshal(data, &varRedeemLockedProductResponse)

	if err != nil {
		return err
	}

	*o = RedeemLockedProductResponse(varRedeemLockedProductResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "redeemId")
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRedeemLockedProductResponse struct {
	value *RedeemLockedProductResponse
	isSet bool
}

func (v NullableRedeemLockedProductResponse) Get() *RedeemLockedProductResponse {
	return v.value
}

func (v *NullableRedeemLockedProductResponse) Set(val *RedeemLockedProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRedeemLockedProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRedeemLockedProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedeemLockedProductResponse(val *RedeemLockedProductResponse) *NullableRedeemLockedProductResponse {
	return &NullableRedeemLockedProductResponse{value: val, isSet: true}
}

func (v NullableRedeemLockedProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedeemLockedProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
