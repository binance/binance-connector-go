/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CreateOtcBlocktradeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateOtcBlocktradeResponse{}

// CreateOtcBlocktradeResponse struct for CreateOtcBlocktradeResponse
type CreateOtcBlocktradeResponse struct {
	OrderId              *string `json:"orderId,omitempty"`
	SecretToken          *string `json:"secretToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateOtcBlocktradeResponse CreateOtcBlocktradeResponse

// NewCreateOtcBlocktradeResponse instantiates a new CreateOtcBlocktradeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOtcBlocktradeResponse() *CreateOtcBlocktradeResponse {
	this := CreateOtcBlocktradeResponse{}
	return &this
}

// NewCreateOtcBlocktradeResponseWithDefaults instantiates a new CreateOtcBlocktradeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOtcBlocktradeResponseWithDefaults() *CreateOtcBlocktradeResponse {
	this := CreateOtcBlocktradeResponse{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *CreateOtcBlocktradeResponse) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOtcBlocktradeResponse) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *CreateOtcBlocktradeResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *CreateOtcBlocktradeResponse) SetOrderId(v string) {
	o.OrderId = &v
}

// GetSecretToken returns the SecretToken field value if set, zero value otherwise.
func (o *CreateOtcBlocktradeResponse) GetSecretToken() string {
	if o == nil || common.IsNil(o.SecretToken) {
		var ret string
		return ret
	}
	return *o.SecretToken
}

// GetSecretTokenOk returns a tuple with the SecretToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOtcBlocktradeResponse) GetSecretTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.SecretToken) {
		return nil, false
	}
	return o.SecretToken, true
}

// HasSecretToken returns a boolean if a field has been set.
func (o *CreateOtcBlocktradeResponse) HasSecretToken() bool {
	if o != nil && !common.IsNil(o.SecretToken) {
		return true
	}

	return false
}

// SetSecretToken gets a reference to the given string and assigns it to the SecretToken field.
func (o *CreateOtcBlocktradeResponse) SetSecretToken(v string) {
	o.SecretToken = &v
}

func (o CreateOtcBlocktradeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOtcBlocktradeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.SecretToken) {
		toSerialize["secretToken"] = o.SecretToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateOtcBlocktradeResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateOtcBlocktradeResponse := _CreateOtcBlocktradeResponse{}

	err = json.Unmarshal(data, &varCreateOtcBlocktradeResponse)

	if err != nil {
		return err
	}

	*o = CreateOtcBlocktradeResponse(varCreateOtcBlocktradeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "secretToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateOtcBlocktradeResponse struct {
	value *CreateOtcBlocktradeResponse
	isSet bool
}

func (v NullableCreateOtcBlocktradeResponse) Get() *CreateOtcBlocktradeResponse {
	return v.value
}

func (v *NullableCreateOtcBlocktradeResponse) Set(val *CreateOtcBlocktradeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOtcBlocktradeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOtcBlocktradeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOtcBlocktradeResponse(val *CreateOtcBlocktradeResponse) *NullableCreateOtcBlocktradeResponse {
	return &NullableCreateOtcBlocktradeResponse{value: val, isSet: true}
}

func (v NullableCreateOtcBlocktradeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOtcBlocktradeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
