/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountCommissionResponseDiscount type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountCommissionResponseDiscount{}

// AccountCommissionResponseDiscount struct for AccountCommissionResponseDiscount
type AccountCommissionResponseDiscount struct {
	EnabledForAccount    *bool   `json:"enabledForAccount,omitempty"`
	EnabledForSymbol     *bool   `json:"enabledForSymbol,omitempty"`
	DiscountAsset        *string `json:"discountAsset,omitempty"`
	Discount             *string `json:"discount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountCommissionResponseDiscount AccountCommissionResponseDiscount

// NewAccountCommissionResponseDiscount instantiates a new AccountCommissionResponseDiscount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCommissionResponseDiscount() *AccountCommissionResponseDiscount {
	this := AccountCommissionResponseDiscount{}
	return &this
}

// NewAccountCommissionResponseDiscountWithDefaults instantiates a new AccountCommissionResponseDiscount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCommissionResponseDiscountWithDefaults() *AccountCommissionResponseDiscount {
	this := AccountCommissionResponseDiscount{}
	return &this
}

// GetEnabledForAccount returns the EnabledForAccount field value if set, zero value otherwise.
func (o *AccountCommissionResponseDiscount) GetEnabledForAccount() bool {
	if o == nil || common.IsNil(o.EnabledForAccount) {
		var ret bool
		return ret
	}
	return *o.EnabledForAccount
}

// GetEnabledForAccountOk returns a tuple with the EnabledForAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseDiscount) GetEnabledForAccountOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnabledForAccount) {
		return nil, false
	}
	return o.EnabledForAccount, true
}

// HasEnabledForAccount returns a boolean if a field has been set.
func (o *AccountCommissionResponseDiscount) HasEnabledForAccount() bool {
	if o != nil && !common.IsNil(o.EnabledForAccount) {
		return true
	}

	return false
}

// SetEnabledForAccount gets a reference to the given bool and assigns it to the EnabledForAccount field.
func (o *AccountCommissionResponseDiscount) SetEnabledForAccount(v bool) {
	o.EnabledForAccount = &v
}

// GetEnabledForSymbol returns the EnabledForSymbol field value if set, zero value otherwise.
func (o *AccountCommissionResponseDiscount) GetEnabledForSymbol() bool {
	if o == nil || common.IsNil(o.EnabledForSymbol) {
		var ret bool
		return ret
	}
	return *o.EnabledForSymbol
}

// GetEnabledForSymbolOk returns a tuple with the EnabledForSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseDiscount) GetEnabledForSymbolOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnabledForSymbol) {
		return nil, false
	}
	return o.EnabledForSymbol, true
}

// HasEnabledForSymbol returns a boolean if a field has been set.
func (o *AccountCommissionResponseDiscount) HasEnabledForSymbol() bool {
	if o != nil && !common.IsNil(o.EnabledForSymbol) {
		return true
	}

	return false
}

// SetEnabledForSymbol gets a reference to the given bool and assigns it to the EnabledForSymbol field.
func (o *AccountCommissionResponseDiscount) SetEnabledForSymbol(v bool) {
	o.EnabledForSymbol = &v
}

// GetDiscountAsset returns the DiscountAsset field value if set, zero value otherwise.
func (o *AccountCommissionResponseDiscount) GetDiscountAsset() string {
	if o == nil || common.IsNil(o.DiscountAsset) {
		var ret string
		return ret
	}
	return *o.DiscountAsset
}

// GetDiscountAssetOk returns a tuple with the DiscountAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseDiscount) GetDiscountAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.DiscountAsset) {
		return nil, false
	}
	return o.DiscountAsset, true
}

// HasDiscountAsset returns a boolean if a field has been set.
func (o *AccountCommissionResponseDiscount) HasDiscountAsset() bool {
	if o != nil && !common.IsNil(o.DiscountAsset) {
		return true
	}

	return false
}

// SetDiscountAsset gets a reference to the given string and assigns it to the DiscountAsset field.
func (o *AccountCommissionResponseDiscount) SetDiscountAsset(v string) {
	o.DiscountAsset = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *AccountCommissionResponseDiscount) GetDiscount() string {
	if o == nil || common.IsNil(o.Discount) {
		var ret string
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseDiscount) GetDiscountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *AccountCommissionResponseDiscount) HasDiscount() bool {
	if o != nil && !common.IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given string and assigns it to the Discount field.
func (o *AccountCommissionResponseDiscount) SetDiscount(v string) {
	o.Discount = &v
}

func (o AccountCommissionResponseDiscount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCommissionResponseDiscount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EnabledForAccount) {
		toSerialize["enabledForAccount"] = o.EnabledForAccount
	}
	if !common.IsNil(o.EnabledForSymbol) {
		toSerialize["enabledForSymbol"] = o.EnabledForSymbol
	}
	if !common.IsNil(o.DiscountAsset) {
		toSerialize["discountAsset"] = o.DiscountAsset
	}
	if !common.IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountCommissionResponseDiscount) UnmarshalJSON(data []byte) (err error) {
	varAccountCommissionResponseDiscount := _AccountCommissionResponseDiscount{}

	err = json.Unmarshal(data, &varAccountCommissionResponseDiscount)

	if err != nil {
		return err
	}

	*o = AccountCommissionResponseDiscount(varAccountCommissionResponseDiscount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabledForAccount")
		delete(additionalProperties, "enabledForSymbol")
		delete(additionalProperties, "discountAsset")
		delete(additionalProperties, "discount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountCommissionResponseDiscount struct {
	value *AccountCommissionResponseDiscount
	isSet bool
}

func (v NullableAccountCommissionResponseDiscount) Get() *AccountCommissionResponseDiscount {
	return v.value
}

func (v *NullableAccountCommissionResponseDiscount) Set(val *AccountCommissionResponseDiscount) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCommissionResponseDiscount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCommissionResponseDiscount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCommissionResponseDiscount(val *AccountCommissionResponseDiscount) *NullableAccountCommissionResponseDiscount {
	return &NullableAccountCommissionResponseDiscount{value: val, isSet: true}
}

func (v NullableAccountCommissionResponseDiscount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCommissionResponseDiscount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
