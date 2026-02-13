/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the UserUniversalTransferResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UserUniversalTransferResponse{}

// UserUniversalTransferResponse struct for UserUniversalTransferResponse
type UserUniversalTransferResponse struct {
	TranId               *int64 `json:"tranId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserUniversalTransferResponse UserUniversalTransferResponse

// NewUserUniversalTransferResponse instantiates a new UserUniversalTransferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUniversalTransferResponse() *UserUniversalTransferResponse {
	this := UserUniversalTransferResponse{}
	return &this
}

// NewUserUniversalTransferResponseWithDefaults instantiates a new UserUniversalTransferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUniversalTransferResponseWithDefaults() *UserUniversalTransferResponse {
	this := UserUniversalTransferResponse{}
	return &this
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *UserUniversalTransferResponse) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUniversalTransferResponse) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *UserUniversalTransferResponse) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *UserUniversalTransferResponse) SetTranId(v int64) {
	o.TranId = &v
}

func (o UserUniversalTransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUniversalTransferResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserUniversalTransferResponse) UnmarshalJSON(data []byte) (err error) {
	varUserUniversalTransferResponse := _UserUniversalTransferResponse{}

	err = json.Unmarshal(data, &varUserUniversalTransferResponse)

	if err != nil {
		return err
	}

	*o = UserUniversalTransferResponse(varUserUniversalTransferResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tranId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserUniversalTransferResponse struct {
	value *UserUniversalTransferResponse
	isSet bool
}

func (v NullableUserUniversalTransferResponse) Get() *UserUniversalTransferResponse {
	return v.value
}

func (v *NullableUserUniversalTransferResponse) Set(val *UserUniversalTransferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUniversalTransferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUniversalTransferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUniversalTransferResponse(val *UserUniversalTransferResponse) *NullableUserUniversalTransferResponse {
	return &NullableUserUniversalTransferResponse{value: val, isSet: true}
}

func (v NullableUserUniversalTransferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUniversalTransferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
