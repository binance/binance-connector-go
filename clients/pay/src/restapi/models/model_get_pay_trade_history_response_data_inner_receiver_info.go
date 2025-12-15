/*
Binance Pay REST API

OpenAPI Specification for the Binance Pay REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetPayTradeHistoryResponseDataInnerReceiverInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPayTradeHistoryResponseDataInnerReceiverInfo{}

// GetPayTradeHistoryResponseDataInnerReceiverInfo struct for GetPayTradeHistoryResponseDataInnerReceiverInfo
type GetPayTradeHistoryResponseDataInnerReceiverInfo struct {
	Name                 *string                                                `json:"name,omitempty"`
	Type                 *string                                                `json:"type,omitempty"`
	Email                *string                                                `json:"email,omitempty"`
	BinanceId            *string                                                `json:"binanceId,omitempty"`
	AccountId            *string                                                `json:"accountId,omitempty"`
	CountryCode          *string                                                `json:"countryCode,omitempty"`
	PhoneNumber          *string                                                `json:"phoneNumber,omitempty"`
	MobileCode           *string                                                `json:"mobileCode,omitempty"`
	Extend               *GetPayTradeHistoryResponseDataInnerReceiverInfoExtend `json:"extend,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPayTradeHistoryResponseDataInnerReceiverInfo GetPayTradeHistoryResponseDataInnerReceiverInfo

// NewGetPayTradeHistoryResponseDataInnerReceiverInfo instantiates a new GetPayTradeHistoryResponseDataInnerReceiverInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPayTradeHistoryResponseDataInnerReceiverInfo() *GetPayTradeHistoryResponseDataInnerReceiverInfo {
	this := GetPayTradeHistoryResponseDataInnerReceiverInfo{}
	return &this
}

// NewGetPayTradeHistoryResponseDataInnerReceiverInfoWithDefaults instantiates a new GetPayTradeHistoryResponseDataInnerReceiverInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPayTradeHistoryResponseDataInnerReceiverInfoWithDefaults() *GetPayTradeHistoryResponseDataInnerReceiverInfo {
	this := GetPayTradeHistoryResponseDataInnerReceiverInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) SetType(v string) {
	o.Type = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) SetEmail(v string) {
	o.Email = &v
}

// GetBinanceId returns the BinanceId field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetBinanceId() string {
	if o == nil || common.IsNil(o.BinanceId) {
		var ret string
		return ret
	}
	return *o.BinanceId
}

// GetBinanceIdOk returns a tuple with the BinanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetBinanceIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BinanceId) {
		return nil, false
	}
	return o.BinanceId, true
}

// HasBinanceId returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) HasBinanceId() bool {
	if o != nil && !common.IsNil(o.BinanceId) {
		return true
	}

	return false
}

// SetBinanceId gets a reference to the given string and assigns it to the BinanceId field.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) SetBinanceId(v string) {
	o.BinanceId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetAccountId() string {
	if o == nil || common.IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) HasAccountId() bool {
	if o != nil && !common.IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetCountryCode() string {
	if o == nil || common.IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetCountryCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) HasCountryCode() bool {
	if o != nil && !common.IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetPhoneNumber() string {
	if o == nil || common.IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetPhoneNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) HasPhoneNumber() bool {
	if o != nil && !common.IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetMobileCode returns the MobileCode field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetMobileCode() string {
	if o == nil || common.IsNil(o.MobileCode) {
		var ret string
		return ret
	}
	return *o.MobileCode
}

// GetMobileCodeOk returns a tuple with the MobileCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetMobileCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MobileCode) {
		return nil, false
	}
	return o.MobileCode, true
}

// HasMobileCode returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) HasMobileCode() bool {
	if o != nil && !common.IsNil(o.MobileCode) {
		return true
	}

	return false
}

// SetMobileCode gets a reference to the given string and assigns it to the MobileCode field.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) SetMobileCode(v string) {
	o.MobileCode = &v
}

// GetExtend returns the Extend field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetExtend() GetPayTradeHistoryResponseDataInnerReceiverInfoExtend {
	if o == nil || common.IsNil(o.Extend) {
		var ret GetPayTradeHistoryResponseDataInnerReceiverInfoExtend
		return ret
	}
	return *o.Extend
}

// GetExtendOk returns a tuple with the Extend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) GetExtendOk() (*GetPayTradeHistoryResponseDataInnerReceiverInfoExtend, bool) {
	if o == nil || common.IsNil(o.Extend) {
		return nil, false
	}
	return o.Extend, true
}

// HasExtend returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) HasExtend() bool {
	if o != nil && !common.IsNil(o.Extend) {
		return true
	}

	return false
}

// SetExtend gets a reference to the given GetPayTradeHistoryResponseDataInnerReceiverInfoExtend and assigns it to the Extend field.
func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) SetExtend(v GetPayTradeHistoryResponseDataInnerReceiverInfoExtend) {
	o.Extend = &v
}

func (o GetPayTradeHistoryResponseDataInnerReceiverInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPayTradeHistoryResponseDataInnerReceiverInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.BinanceId) {
		toSerialize["binanceId"] = o.BinanceId
	}
	if !common.IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !common.IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !common.IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !common.IsNil(o.MobileCode) {
		toSerialize["mobileCode"] = o.MobileCode
	}
	if !common.IsNil(o.Extend) {
		toSerialize["extend"] = o.Extend
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPayTradeHistoryResponseDataInnerReceiverInfo) UnmarshalJSON(data []byte) (err error) {
	varGetPayTradeHistoryResponseDataInnerReceiverInfo := _GetPayTradeHistoryResponseDataInnerReceiverInfo{}

	err = json.Unmarshal(data, &varGetPayTradeHistoryResponseDataInnerReceiverInfo)

	if err != nil {
		return err
	}

	*o = GetPayTradeHistoryResponseDataInnerReceiverInfo(varGetPayTradeHistoryResponseDataInnerReceiverInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "email")
		delete(additionalProperties, "binanceId")
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "countryCode")
		delete(additionalProperties, "phoneNumber")
		delete(additionalProperties, "mobileCode")
		delete(additionalProperties, "extend")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPayTradeHistoryResponseDataInnerReceiverInfo struct {
	value *GetPayTradeHistoryResponseDataInnerReceiverInfo
	isSet bool
}

func (v NullableGetPayTradeHistoryResponseDataInnerReceiverInfo) Get() *GetPayTradeHistoryResponseDataInnerReceiverInfo {
	return v.value
}

func (v *NullableGetPayTradeHistoryResponseDataInnerReceiverInfo) Set(val *GetPayTradeHistoryResponseDataInnerReceiverInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPayTradeHistoryResponseDataInnerReceiverInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPayTradeHistoryResponseDataInnerReceiverInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPayTradeHistoryResponseDataInnerReceiverInfo(val *GetPayTradeHistoryResponseDataInnerReceiverInfo) *NullableGetPayTradeHistoryResponseDataInnerReceiverInfo {
	return &NullableGetPayTradeHistoryResponseDataInnerReceiverInfo{value: val, isSet: true}
}

func (v NullableGetPayTradeHistoryResponseDataInnerReceiverInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPayTradeHistoryResponseDataInnerReceiverInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
