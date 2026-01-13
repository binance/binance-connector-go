/*
Binance Gift Card REST API

OpenAPI Specification for the Binance Gift Card REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the VerifyBinanceGiftCardByGiftCardNumberResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &VerifyBinanceGiftCardByGiftCardNumberResponseData{}

// VerifyBinanceGiftCardByGiftCardNumberResponseData struct for VerifyBinanceGiftCardByGiftCardNumberResponseData
type VerifyBinanceGiftCardByGiftCardNumberResponseData struct {
	Valid                *bool   `json:"valid,omitempty"`
	Token                *string `json:"token,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VerifyBinanceGiftCardByGiftCardNumberResponseData VerifyBinanceGiftCardByGiftCardNumberResponseData

// NewVerifyBinanceGiftCardByGiftCardNumberResponseData instantiates a new VerifyBinanceGiftCardByGiftCardNumberResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyBinanceGiftCardByGiftCardNumberResponseData() *VerifyBinanceGiftCardByGiftCardNumberResponseData {
	this := VerifyBinanceGiftCardByGiftCardNumberResponseData{}
	return &this
}

// NewVerifyBinanceGiftCardByGiftCardNumberResponseDataWithDefaults instantiates a new VerifyBinanceGiftCardByGiftCardNumberResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyBinanceGiftCardByGiftCardNumberResponseDataWithDefaults() *VerifyBinanceGiftCardByGiftCardNumberResponseData {
	this := VerifyBinanceGiftCardByGiftCardNumberResponseData{}
	return &this
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *VerifyBinanceGiftCardByGiftCardNumberResponseData) GetValid() bool {
	if o == nil || common.IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyBinanceGiftCardByGiftCardNumberResponseData) GetValidOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *VerifyBinanceGiftCardByGiftCardNumberResponseData) HasValid() bool {
	if o != nil && !common.IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *VerifyBinanceGiftCardByGiftCardNumberResponseData) SetValid(v bool) {
	o.Valid = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *VerifyBinanceGiftCardByGiftCardNumberResponseData) GetToken() string {
	if o == nil || common.IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyBinanceGiftCardByGiftCardNumberResponseData) GetTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *VerifyBinanceGiftCardByGiftCardNumberResponseData) HasToken() bool {
	if o != nil && !common.IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *VerifyBinanceGiftCardByGiftCardNumberResponseData) SetToken(v string) {
	o.Token = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *VerifyBinanceGiftCardByGiftCardNumberResponseData) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyBinanceGiftCardByGiftCardNumberResponseData) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *VerifyBinanceGiftCardByGiftCardNumberResponseData) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *VerifyBinanceGiftCardByGiftCardNumberResponseData) SetAmount(v string) {
	o.Amount = &v
}

func (o VerifyBinanceGiftCardByGiftCardNumberResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyBinanceGiftCardByGiftCardNumberResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Valid) {
		toSerialize["valid"] = o.Valid
	}
	if !common.IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VerifyBinanceGiftCardByGiftCardNumberResponseData) UnmarshalJSON(data []byte) (err error) {
	varVerifyBinanceGiftCardByGiftCardNumberResponseData := _VerifyBinanceGiftCardByGiftCardNumberResponseData{}

	err = json.Unmarshal(data, &varVerifyBinanceGiftCardByGiftCardNumberResponseData)

	if err != nil {
		return err
	}

	*o = VerifyBinanceGiftCardByGiftCardNumberResponseData(varVerifyBinanceGiftCardByGiftCardNumberResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "valid")
		delete(additionalProperties, "token")
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVerifyBinanceGiftCardByGiftCardNumberResponseData struct {
	value *VerifyBinanceGiftCardByGiftCardNumberResponseData
	isSet bool
}

func (v NullableVerifyBinanceGiftCardByGiftCardNumberResponseData) Get() *VerifyBinanceGiftCardByGiftCardNumberResponseData {
	return v.value
}

func (v *NullableVerifyBinanceGiftCardByGiftCardNumberResponseData) Set(val *VerifyBinanceGiftCardByGiftCardNumberResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyBinanceGiftCardByGiftCardNumberResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyBinanceGiftCardByGiftCardNumberResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyBinanceGiftCardByGiftCardNumberResponseData(val *VerifyBinanceGiftCardByGiftCardNumberResponseData) *NullableVerifyBinanceGiftCardByGiftCardNumberResponseData {
	return &NullableVerifyBinanceGiftCardByGiftCardNumberResponseData{value: val, isSet: true}
}

func (v NullableVerifyBinanceGiftCardByGiftCardNumberResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyBinanceGiftCardByGiftCardNumberResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
