/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the RedeemOnChainYieldsLockedProductResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RedeemOnChainYieldsLockedProductResponse{}

// RedeemOnChainYieldsLockedProductResponse struct for RedeemOnChainYieldsLockedProductResponse
type RedeemOnChainYieldsLockedProductResponse struct {
	RedeemId             *int64 `json:"redeemId,omitempty"`
	Success              *bool  `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RedeemOnChainYieldsLockedProductResponse RedeemOnChainYieldsLockedProductResponse

// NewRedeemOnChainYieldsLockedProductResponse instantiates a new RedeemOnChainYieldsLockedProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedeemOnChainYieldsLockedProductResponse() *RedeemOnChainYieldsLockedProductResponse {
	this := RedeemOnChainYieldsLockedProductResponse{}
	return &this
}

// NewRedeemOnChainYieldsLockedProductResponseWithDefaults instantiates a new RedeemOnChainYieldsLockedProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedeemOnChainYieldsLockedProductResponseWithDefaults() *RedeemOnChainYieldsLockedProductResponse {
	this := RedeemOnChainYieldsLockedProductResponse{}
	return &this
}

// GetRedeemId returns the RedeemId field value if set, zero value otherwise.
func (o *RedeemOnChainYieldsLockedProductResponse) GetRedeemId() int64 {
	if o == nil || common.IsNil(o.RedeemId) {
		var ret int64
		return ret
	}
	return *o.RedeemId
}

// GetRedeemIdOk returns a tuple with the RedeemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemOnChainYieldsLockedProductResponse) GetRedeemIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RedeemId) {
		return nil, false
	}
	return o.RedeemId, true
}

// HasRedeemId returns a boolean if a field has been set.
func (o *RedeemOnChainYieldsLockedProductResponse) HasRedeemId() bool {
	if o != nil && !common.IsNil(o.RedeemId) {
		return true
	}

	return false
}

// SetRedeemId gets a reference to the given int64 and assigns it to the RedeemId field.
func (o *RedeemOnChainYieldsLockedProductResponse) SetRedeemId(v int64) {
	o.RedeemId = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *RedeemOnChainYieldsLockedProductResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemOnChainYieldsLockedProductResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *RedeemOnChainYieldsLockedProductResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *RedeemOnChainYieldsLockedProductResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o RedeemOnChainYieldsLockedProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedeemOnChainYieldsLockedProductResponse) ToMap() (map[string]interface{}, error) {
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

func (o *RedeemOnChainYieldsLockedProductResponse) UnmarshalJSON(data []byte) (err error) {
	varRedeemOnChainYieldsLockedProductResponse := _RedeemOnChainYieldsLockedProductResponse{}

	err = json.Unmarshal(data, &varRedeemOnChainYieldsLockedProductResponse)

	if err != nil {
		return err
	}

	*o = RedeemOnChainYieldsLockedProductResponse(varRedeemOnChainYieldsLockedProductResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "redeemId")
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRedeemOnChainYieldsLockedProductResponse struct {
	value *RedeemOnChainYieldsLockedProductResponse
	isSet bool
}

func (v NullableRedeemOnChainYieldsLockedProductResponse) Get() *RedeemOnChainYieldsLockedProductResponse {
	return v.value
}

func (v *NullableRedeemOnChainYieldsLockedProductResponse) Set(val *RedeemOnChainYieldsLockedProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRedeemOnChainYieldsLockedProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRedeemOnChainYieldsLockedProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedeemOnChainYieldsLockedProductResponse(val *RedeemOnChainYieldsLockedProductResponse) *NullableRedeemOnChainYieldsLockedProductResponse {
	return &NullableRedeemOnChainYieldsLockedProductResponse{value: val, isSet: true}
}

func (v NullableRedeemOnChainYieldsLockedProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedeemOnChainYieldsLockedProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
