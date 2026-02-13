/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountCommissionResponseResultDiscount type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountCommissionResponseResultDiscount{}

// AccountCommissionResponseResultDiscount struct for AccountCommissionResponseResultDiscount
type AccountCommissionResponseResultDiscount struct {
	EnabledForAccount    *bool   `json:"enabledForAccount,omitempty"`
	EnabledForSymbol     *bool   `json:"enabledForSymbol,omitempty"`
	DiscountAsset        *string `json:"discountAsset,omitempty"`
	Discount             *string `json:"discount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountCommissionResponseResultDiscount AccountCommissionResponseResultDiscount

// NewAccountCommissionResponseResultDiscount instantiates a new AccountCommissionResponseResultDiscount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCommissionResponseResultDiscount() *AccountCommissionResponseResultDiscount {
	this := AccountCommissionResponseResultDiscount{}
	return &this
}

// NewAccountCommissionResponseResultDiscountWithDefaults instantiates a new AccountCommissionResponseResultDiscount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCommissionResponseResultDiscountWithDefaults() *AccountCommissionResponseResultDiscount {
	this := AccountCommissionResponseResultDiscount{}
	return &this
}

// GetEnabledForAccount returns the EnabledForAccount field value if set, zero value otherwise.
func (o *AccountCommissionResponseResultDiscount) GetEnabledForAccount() bool {
	if o == nil || common.IsNil(o.EnabledForAccount) {
		var ret bool
		return ret
	}
	return *o.EnabledForAccount
}

// GetEnabledForAccountOk returns a tuple with the EnabledForAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseResultDiscount) GetEnabledForAccountOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnabledForAccount) {
		return nil, false
	}
	return o.EnabledForAccount, true
}

// HasEnabledForAccount returns a boolean if a field has been set.
func (o *AccountCommissionResponseResultDiscount) HasEnabledForAccount() bool {
	if o != nil && !common.IsNil(o.EnabledForAccount) {
		return true
	}

	return false
}

// SetEnabledForAccount gets a reference to the given bool and assigns it to the EnabledForAccount field.
func (o *AccountCommissionResponseResultDiscount) SetEnabledForAccount(v bool) {
	o.EnabledForAccount = &v
}

// GetEnabledForSymbol returns the EnabledForSymbol field value if set, zero value otherwise.
func (o *AccountCommissionResponseResultDiscount) GetEnabledForSymbol() bool {
	if o == nil || common.IsNil(o.EnabledForSymbol) {
		var ret bool
		return ret
	}
	return *o.EnabledForSymbol
}

// GetEnabledForSymbolOk returns a tuple with the EnabledForSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseResultDiscount) GetEnabledForSymbolOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnabledForSymbol) {
		return nil, false
	}
	return o.EnabledForSymbol, true
}

// HasEnabledForSymbol returns a boolean if a field has been set.
func (o *AccountCommissionResponseResultDiscount) HasEnabledForSymbol() bool {
	if o != nil && !common.IsNil(o.EnabledForSymbol) {
		return true
	}

	return false
}

// SetEnabledForSymbol gets a reference to the given bool and assigns it to the EnabledForSymbol field.
func (o *AccountCommissionResponseResultDiscount) SetEnabledForSymbol(v bool) {
	o.EnabledForSymbol = &v
}

// GetDiscountAsset returns the DiscountAsset field value if set, zero value otherwise.
func (o *AccountCommissionResponseResultDiscount) GetDiscountAsset() string {
	if o == nil || common.IsNil(o.DiscountAsset) {
		var ret string
		return ret
	}
	return *o.DiscountAsset
}

// GetDiscountAssetOk returns a tuple with the DiscountAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseResultDiscount) GetDiscountAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.DiscountAsset) {
		return nil, false
	}
	return o.DiscountAsset, true
}

// HasDiscountAsset returns a boolean if a field has been set.
func (o *AccountCommissionResponseResultDiscount) HasDiscountAsset() bool {
	if o != nil && !common.IsNil(o.DiscountAsset) {
		return true
	}

	return false
}

// SetDiscountAsset gets a reference to the given string and assigns it to the DiscountAsset field.
func (o *AccountCommissionResponseResultDiscount) SetDiscountAsset(v string) {
	o.DiscountAsset = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *AccountCommissionResponseResultDiscount) GetDiscount() string {
	if o == nil || common.IsNil(o.Discount) {
		var ret string
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseResultDiscount) GetDiscountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *AccountCommissionResponseResultDiscount) HasDiscount() bool {
	if o != nil && !common.IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given string and assigns it to the Discount field.
func (o *AccountCommissionResponseResultDiscount) SetDiscount(v string) {
	o.Discount = &v
}

func (o AccountCommissionResponseResultDiscount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCommissionResponseResultDiscount) ToMap() (map[string]interface{}, error) {
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

func (o *AccountCommissionResponseResultDiscount) UnmarshalJSON(data []byte) (err error) {
	varAccountCommissionResponseResultDiscount := _AccountCommissionResponseResultDiscount{}

	err = json.Unmarshal(data, &varAccountCommissionResponseResultDiscount)

	if err != nil {
		return err
	}

	*o = AccountCommissionResponseResultDiscount(varAccountCommissionResponseResultDiscount)

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

type NullableAccountCommissionResponseResultDiscount struct {
	value *AccountCommissionResponseResultDiscount
	isSet bool
}

func (v NullableAccountCommissionResponseResultDiscount) Get() *AccountCommissionResponseResultDiscount {
	return v.value
}

func (v *NullableAccountCommissionResponseResultDiscount) Set(val *AccountCommissionResponseResultDiscount) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCommissionResponseResultDiscount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCommissionResponseResultDiscount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCommissionResponseResultDiscount(val *AccountCommissionResponseResultDiscount) *NullableAccountCommissionResponseResultDiscount {
	return &NullableAccountCommissionResponseResultDiscount{value: val, isSet: true}
}

func (v NullableAccountCommissionResponseResultDiscount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCommissionResponseResultDiscount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
