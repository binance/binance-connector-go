/*
Binance Gift Card REST API

OpenAPI Specification for the Binance Gift Card REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the RedeemABinanceGiftCardResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RedeemABinanceGiftCardResponseData{}

// RedeemABinanceGiftCardResponseData struct for RedeemABinanceGiftCardResponseData
type RedeemABinanceGiftCardResponseData struct {
	ReferenceNo          *string `json:"referenceNo,omitempty"`
	IdentityNo           *string `json:"identityNo,omitempty"`
	Token                *string `json:"token,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RedeemABinanceGiftCardResponseData RedeemABinanceGiftCardResponseData

// NewRedeemABinanceGiftCardResponseData instantiates a new RedeemABinanceGiftCardResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedeemABinanceGiftCardResponseData() *RedeemABinanceGiftCardResponseData {
	this := RedeemABinanceGiftCardResponseData{}
	return &this
}

// NewRedeemABinanceGiftCardResponseDataWithDefaults instantiates a new RedeemABinanceGiftCardResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedeemABinanceGiftCardResponseDataWithDefaults() *RedeemABinanceGiftCardResponseData {
	this := RedeemABinanceGiftCardResponseData{}
	return &this
}

// GetReferenceNo returns the ReferenceNo field value if set, zero value otherwise.
func (o *RedeemABinanceGiftCardResponseData) GetReferenceNo() string {
	if o == nil || common.IsNil(o.ReferenceNo) {
		var ret string
		return ret
	}
	return *o.ReferenceNo
}

// GetReferenceNoOk returns a tuple with the ReferenceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemABinanceGiftCardResponseData) GetReferenceNoOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReferenceNo) {
		return nil, false
	}
	return o.ReferenceNo, true
}

// HasReferenceNo returns a boolean if a field has been set.
func (o *RedeemABinanceGiftCardResponseData) HasReferenceNo() bool {
	if o != nil && !common.IsNil(o.ReferenceNo) {
		return true
	}

	return false
}

// SetReferenceNo gets a reference to the given string and assigns it to the ReferenceNo field.
func (o *RedeemABinanceGiftCardResponseData) SetReferenceNo(v string) {
	o.ReferenceNo = &v
}

// GetIdentityNo returns the IdentityNo field value if set, zero value otherwise.
func (o *RedeemABinanceGiftCardResponseData) GetIdentityNo() string {
	if o == nil || common.IsNil(o.IdentityNo) {
		var ret string
		return ret
	}
	return *o.IdentityNo
}

// GetIdentityNoOk returns a tuple with the IdentityNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemABinanceGiftCardResponseData) GetIdentityNoOk() (*string, bool) {
	if o == nil || common.IsNil(o.IdentityNo) {
		return nil, false
	}
	return o.IdentityNo, true
}

// HasIdentityNo returns a boolean if a field has been set.
func (o *RedeemABinanceGiftCardResponseData) HasIdentityNo() bool {
	if o != nil && !common.IsNil(o.IdentityNo) {
		return true
	}

	return false
}

// SetIdentityNo gets a reference to the given string and assigns it to the IdentityNo field.
func (o *RedeemABinanceGiftCardResponseData) SetIdentityNo(v string) {
	o.IdentityNo = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RedeemABinanceGiftCardResponseData) GetToken() string {
	if o == nil || common.IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemABinanceGiftCardResponseData) GetTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RedeemABinanceGiftCardResponseData) HasToken() bool {
	if o != nil && !common.IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RedeemABinanceGiftCardResponseData) SetToken(v string) {
	o.Token = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *RedeemABinanceGiftCardResponseData) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemABinanceGiftCardResponseData) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *RedeemABinanceGiftCardResponseData) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *RedeemABinanceGiftCardResponseData) SetAmount(v string) {
	o.Amount = &v
}

func (o RedeemABinanceGiftCardResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedeemABinanceGiftCardResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ReferenceNo) {
		toSerialize["referenceNo"] = o.ReferenceNo
	}
	if !common.IsNil(o.IdentityNo) {
		toSerialize["identityNo"] = o.IdentityNo
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

func (o *RedeemABinanceGiftCardResponseData) UnmarshalJSON(data []byte) (err error) {
	varRedeemABinanceGiftCardResponseData := _RedeemABinanceGiftCardResponseData{}

	err = json.Unmarshal(data, &varRedeemABinanceGiftCardResponseData)

	if err != nil {
		return err
	}

	*o = RedeemABinanceGiftCardResponseData(varRedeemABinanceGiftCardResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "referenceNo")
		delete(additionalProperties, "identityNo")
		delete(additionalProperties, "token")
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRedeemABinanceGiftCardResponseData struct {
	value *RedeemABinanceGiftCardResponseData
	isSet bool
}

func (v NullableRedeemABinanceGiftCardResponseData) Get() *RedeemABinanceGiftCardResponseData {
	return v.value
}

func (v *NullableRedeemABinanceGiftCardResponseData) Set(val *RedeemABinanceGiftCardResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableRedeemABinanceGiftCardResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableRedeemABinanceGiftCardResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedeemABinanceGiftCardResponseData(val *RedeemABinanceGiftCardResponseData) *NullableRedeemABinanceGiftCardResponseData {
	return &NullableRedeemABinanceGiftCardResponseData{value: val, isSet: true}
}

func (v NullableRedeemABinanceGiftCardResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedeemABinanceGiftCardResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
