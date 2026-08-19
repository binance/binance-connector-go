/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CreateRenewListenKeyResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateRenewListenKeyResponse{}

// CreateRenewListenKeyResponse struct for CreateRenewListenKeyResponse
type CreateRenewListenKeyResponse struct {
	// Session identifier for the stock user data stream.
	ListenKey            *string `json:"listenKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateRenewListenKeyResponse CreateRenewListenKeyResponse

// NewCreateRenewListenKeyResponse instantiates a new CreateRenewListenKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRenewListenKeyResponse() *CreateRenewListenKeyResponse {
	this := CreateRenewListenKeyResponse{}
	return &this
}

// NewCreateRenewListenKeyResponseWithDefaults instantiates a new CreateRenewListenKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRenewListenKeyResponseWithDefaults() *CreateRenewListenKeyResponse {
	this := CreateRenewListenKeyResponse{}
	return &this
}

// GetListenKey returns the ListenKey field value if set, zero value otherwise.
func (o *CreateRenewListenKeyResponse) GetListenKey() string {
	if o == nil || common.IsNil(o.ListenKey) {
		var ret string
		return ret
	}
	return *o.ListenKey
}

// GetListenKeyOk returns a tuple with the ListenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRenewListenKeyResponse) GetListenKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListenKey) {
		return nil, false
	}
	return o.ListenKey, true
}

// HasListenKey returns a boolean if a field has been set.
func (o *CreateRenewListenKeyResponse) HasListenKey() bool {
	if o != nil && !common.IsNil(o.ListenKey) {
		return true
	}

	return false
}

// SetListenKey gets a reference to the given string and assigns it to the ListenKey field.
func (o *CreateRenewListenKeyResponse) SetListenKey(v string) {
	o.ListenKey = &v
}

func (o CreateRenewListenKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRenewListenKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ListenKey) {
		toSerialize["listenKey"] = o.ListenKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateRenewListenKeyResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateRenewListenKeyResponse := _CreateRenewListenKeyResponse{}

	err = json.Unmarshal(data, &varCreateRenewListenKeyResponse)

	if err != nil {
		return err
	}

	*o = CreateRenewListenKeyResponse(varCreateRenewListenKeyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "listenKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateRenewListenKeyResponse struct {
	value *CreateRenewListenKeyResponse
	isSet bool
}

func (v NullableCreateRenewListenKeyResponse) Get() *CreateRenewListenKeyResponse {
	return v.value
}

func (v *NullableCreateRenewListenKeyResponse) Set(val *CreateRenewListenKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRenewListenKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRenewListenKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRenewListenKeyResponse(val *CreateRenewListenKeyResponse) *NullableCreateRenewListenKeyResponse {
	return &NullableCreateRenewListenKeyResponse{value: val, isSet: true}
}

func (v NullableCreateRenewListenKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRenewListenKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
