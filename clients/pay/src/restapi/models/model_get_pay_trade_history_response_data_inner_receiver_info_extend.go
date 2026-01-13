/*
Binance Pay REST API

OpenAPI Specification for the Binance Pay REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetPayTradeHistoryResponseDataInnerReceiverInfoExtend type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPayTradeHistoryResponseDataInnerReceiverInfoExtend{}

// GetPayTradeHistoryResponseDataInnerReceiverInfoExtend struct for GetPayTradeHistoryResponseDataInnerReceiverInfoExtend
type GetPayTradeHistoryResponseDataInnerReceiverInfoExtend struct {
	InstitutionName      *string `json:"institutionName,omitempty"`
	CardNumber           *string `json:"cardNumber,omitempty"`
	DigitalWalletId      *string `json:"digitalWalletId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPayTradeHistoryResponseDataInnerReceiverInfoExtend GetPayTradeHistoryResponseDataInnerReceiverInfoExtend

// NewGetPayTradeHistoryResponseDataInnerReceiverInfoExtend instantiates a new GetPayTradeHistoryResponseDataInnerReceiverInfoExtend object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPayTradeHistoryResponseDataInnerReceiverInfoExtend() *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend {
	this := GetPayTradeHistoryResponseDataInnerReceiverInfoExtend{}
	return &this
}

// NewGetPayTradeHistoryResponseDataInnerReceiverInfoExtendWithDefaults instantiates a new GetPayTradeHistoryResponseDataInnerReceiverInfoExtend object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPayTradeHistoryResponseDataInnerReceiverInfoExtendWithDefaults() *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend {
	this := GetPayTradeHistoryResponseDataInnerReceiverInfoExtend{}
	return &this
}

// GetInstitutionName returns the InstitutionName field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) GetInstitutionName() string {
	if o == nil || common.IsNil(o.InstitutionName) {
		var ret string
		return ret
	}
	return *o.InstitutionName
}

// GetInstitutionNameOk returns a tuple with the InstitutionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) GetInstitutionNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.InstitutionName) {
		return nil, false
	}
	return o.InstitutionName, true
}

// HasInstitutionName returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) HasInstitutionName() bool {
	if o != nil && !common.IsNil(o.InstitutionName) {
		return true
	}

	return false
}

// SetInstitutionName gets a reference to the given string and assigns it to the InstitutionName field.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) SetInstitutionName(v string) {
	o.InstitutionName = &v
}

// GetCardNumber returns the CardNumber field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) GetCardNumber() string {
	if o == nil || common.IsNil(o.CardNumber) {
		var ret string
		return ret
	}
	return *o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) GetCardNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardNumber) {
		return nil, false
	}
	return o.CardNumber, true
}

// HasCardNumber returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) HasCardNumber() bool {
	if o != nil && !common.IsNil(o.CardNumber) {
		return true
	}

	return false
}

// SetCardNumber gets a reference to the given string and assigns it to the CardNumber field.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) SetCardNumber(v string) {
	o.CardNumber = &v
}

// GetDigitalWalletId returns the DigitalWalletId field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) GetDigitalWalletId() string {
	if o == nil || common.IsNil(o.DigitalWalletId) {
		var ret string
		return ret
	}
	return *o.DigitalWalletId
}

// GetDigitalWalletIdOk returns a tuple with the DigitalWalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) GetDigitalWalletIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.DigitalWalletId) {
		return nil, false
	}
	return o.DigitalWalletId, true
}

// HasDigitalWalletId returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) HasDigitalWalletId() bool {
	if o != nil && !common.IsNil(o.DigitalWalletId) {
		return true
	}

	return false
}

// SetDigitalWalletId gets a reference to the given string and assigns it to the DigitalWalletId field.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) SetDigitalWalletId(v string) {
	o.DigitalWalletId = &v
}

func (o GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.InstitutionName) {
		toSerialize["institutionName"] = o.InstitutionName
	}
	if !common.IsNil(o.CardNumber) {
		toSerialize["cardNumber"] = o.CardNumber
	}
	if !common.IsNil(o.DigitalWalletId) {
		toSerialize["digitalWalletId"] = o.DigitalWalletId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) UnmarshalJSON(data []byte) (err error) {
	varGetPayTradeHistoryResponseDataInnerReceiverInfoExtend := _GetPayTradeHistoryResponseDataInnerReceiverInfoExtend{}

	err = json.Unmarshal(data, &varGetPayTradeHistoryResponseDataInnerReceiverInfoExtend)

	if err != nil {
		return err
	}

	*o = GetPayTradeHistoryResponseDataInnerReceiverInfoExtend(varGetPayTradeHistoryResponseDataInnerReceiverInfoExtend)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "institutionName")
		delete(additionalProperties, "cardNumber")
		delete(additionalProperties, "digitalWalletId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPayTradeHistoryResponseDataInnerReceiverInfoExtend struct {
	value *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend
	isSet bool
}

func (v NullableGetPayTradeHistoryResponseDataInnerReceiverInfoExtend) Get() *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend {
	return v.value
}

func (v *NullableGetPayTradeHistoryResponseDataInnerReceiverInfoExtend) Set(val *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPayTradeHistoryResponseDataInnerReceiverInfoExtend) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPayTradeHistoryResponseDataInnerReceiverInfoExtend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPayTradeHistoryResponseDataInnerReceiverInfoExtend(val *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) *NullableGetPayTradeHistoryResponseDataInnerReceiverInfoExtend {
	return &NullableGetPayTradeHistoryResponseDataInnerReceiverInfoExtend{value: val, isSet: true}
}

func (v NullableGetPayTradeHistoryResponseDataInnerReceiverInfoExtend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPayTradeHistoryResponseDataInnerReceiverInfoExtend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
