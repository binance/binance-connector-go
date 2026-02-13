/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the WithdrawResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &WithdrawResponse{}

// WithdrawResponse struct for WithdrawResponse
type WithdrawResponse struct {
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WithdrawResponse WithdrawResponse

// NewWithdrawResponse instantiates a new WithdrawResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithdrawResponse() *WithdrawResponse {
	this := WithdrawResponse{}
	return &this
}

// NewWithdrawResponseWithDefaults instantiates a new WithdrawResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithdrawResponseWithDefaults() *WithdrawResponse {
	this := WithdrawResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WithdrawResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WithdrawResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WithdrawResponse) SetId(v string) {
	o.Id = &v
}

func (o WithdrawResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WithdrawResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WithdrawResponse) UnmarshalJSON(data []byte) (err error) {
	varWithdrawResponse := _WithdrawResponse{}

	err = json.Unmarshal(data, &varWithdrawResponse)

	if err != nil {
		return err
	}

	*o = WithdrawResponse(varWithdrawResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWithdrawResponse struct {
	value *WithdrawResponse
	isSet bool
}

func (v NullableWithdrawResponse) Get() *WithdrawResponse {
	return v.value
}

func (v *NullableWithdrawResponse) Set(val *WithdrawResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWithdrawResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWithdrawResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWithdrawResponse(val *WithdrawResponse) *NullableWithdrawResponse {
	return &NullableWithdrawResponse{value: val, isSet: true}
}

func (v NullableWithdrawResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithdrawResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
