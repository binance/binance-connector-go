/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountStatusResponse{}

// AccountStatusResponse struct for AccountStatusResponse
type AccountStatusResponse struct {
	Data                 *string `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountStatusResponse AccountStatusResponse

// NewAccountStatusResponse instantiates a new AccountStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountStatusResponse() *AccountStatusResponse {
	this := AccountStatusResponse{}
	return &this
}

// NewAccountStatusResponseWithDefaults instantiates a new AccountStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountStatusResponseWithDefaults() *AccountStatusResponse {
	this := AccountStatusResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AccountStatusResponse) GetData() string {
	if o == nil || common.IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatusResponse) GetDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AccountStatusResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *AccountStatusResponse) SetData(v string) {
	o.Data = &v
}

func (o AccountStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varAccountStatusResponse := _AccountStatusResponse{}

	err = json.Unmarshal(data, &varAccountStatusResponse)

	if err != nil {
		return err
	}

	*o = AccountStatusResponse(varAccountStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountStatusResponse struct {
	value *AccountStatusResponse
	isSet bool
}

func (v NullableAccountStatusResponse) Get() *AccountStatusResponse {
	return v.value
}

func (v *NullableAccountStatusResponse) Set(val *AccountStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountStatusResponse(val *AccountStatusResponse) *NullableAccountStatusResponse {
	return &NullableAccountStatusResponse{value: val, isSet: true}
}

func (v NullableAccountStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
