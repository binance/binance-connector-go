/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountCommissionResponseTaxCommission type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountCommissionResponseTaxCommission{}

// AccountCommissionResponseTaxCommission struct for AccountCommissionResponseTaxCommission
type AccountCommissionResponseTaxCommission struct {
	Maker                *string `json:"maker,omitempty"`
	Taker                *string `json:"taker,omitempty"`
	Buyer                *string `json:"buyer,omitempty"`
	Seller               *string `json:"seller,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountCommissionResponseTaxCommission AccountCommissionResponseTaxCommission

// NewAccountCommissionResponseTaxCommission instantiates a new AccountCommissionResponseTaxCommission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCommissionResponseTaxCommission() *AccountCommissionResponseTaxCommission {
	this := AccountCommissionResponseTaxCommission{}
	return &this
}

// NewAccountCommissionResponseTaxCommissionWithDefaults instantiates a new AccountCommissionResponseTaxCommission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCommissionResponseTaxCommissionWithDefaults() *AccountCommissionResponseTaxCommission {
	this := AccountCommissionResponseTaxCommission{}
	return &this
}

// GetMaker returns the Maker field value if set, zero value otherwise.
func (o *AccountCommissionResponseTaxCommission) GetMaker() string {
	if o == nil || common.IsNil(o.Maker) {
		var ret string
		return ret
	}
	return *o.Maker
}

// GetMakerOk returns a tuple with the Maker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseTaxCommission) GetMakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Maker) {
		return nil, false
	}
	return o.Maker, true
}

// HasMaker returns a boolean if a field has been set.
func (o *AccountCommissionResponseTaxCommission) HasMaker() bool {
	if o != nil && !common.IsNil(o.Maker) {
		return true
	}

	return false
}

// SetMaker gets a reference to the given string and assigns it to the Maker field.
func (o *AccountCommissionResponseTaxCommission) SetMaker(v string) {
	o.Maker = &v
}

// GetTaker returns the Taker field value if set, zero value otherwise.
func (o *AccountCommissionResponseTaxCommission) GetTaker() string {
	if o == nil || common.IsNil(o.Taker) {
		var ret string
		return ret
	}
	return *o.Taker
}

// GetTakerOk returns a tuple with the Taker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseTaxCommission) GetTakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Taker) {
		return nil, false
	}
	return o.Taker, true
}

// HasTaker returns a boolean if a field has been set.
func (o *AccountCommissionResponseTaxCommission) HasTaker() bool {
	if o != nil && !common.IsNil(o.Taker) {
		return true
	}

	return false
}

// SetTaker gets a reference to the given string and assigns it to the Taker field.
func (o *AccountCommissionResponseTaxCommission) SetTaker(v string) {
	o.Taker = &v
}

// GetBuyer returns the Buyer field value if set, zero value otherwise.
func (o *AccountCommissionResponseTaxCommission) GetBuyer() string {
	if o == nil || common.IsNil(o.Buyer) {
		var ret string
		return ret
	}
	return *o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseTaxCommission) GetBuyerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Buyer) {
		return nil, false
	}
	return o.Buyer, true
}

// HasBuyer returns a boolean if a field has been set.
func (o *AccountCommissionResponseTaxCommission) HasBuyer() bool {
	if o != nil && !common.IsNil(o.Buyer) {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given string and assigns it to the Buyer field.
func (o *AccountCommissionResponseTaxCommission) SetBuyer(v string) {
	o.Buyer = &v
}

// GetSeller returns the Seller field value if set, zero value otherwise.
func (o *AccountCommissionResponseTaxCommission) GetSeller() string {
	if o == nil || common.IsNil(o.Seller) {
		var ret string
		return ret
	}
	return *o.Seller
}

// GetSellerOk returns a tuple with the Seller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseTaxCommission) GetSellerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Seller) {
		return nil, false
	}
	return o.Seller, true
}

// HasSeller returns a boolean if a field has been set.
func (o *AccountCommissionResponseTaxCommission) HasSeller() bool {
	if o != nil && !common.IsNil(o.Seller) {
		return true
	}

	return false
}

// SetSeller gets a reference to the given string and assigns it to the Seller field.
func (o *AccountCommissionResponseTaxCommission) SetSeller(v string) {
	o.Seller = &v
}

func (o AccountCommissionResponseTaxCommission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCommissionResponseTaxCommission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Maker) {
		toSerialize["maker"] = o.Maker
	}
	if !common.IsNil(o.Taker) {
		toSerialize["taker"] = o.Taker
	}
	if !common.IsNil(o.Buyer) {
		toSerialize["buyer"] = o.Buyer
	}
	if !common.IsNil(o.Seller) {
		toSerialize["seller"] = o.Seller
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountCommissionResponseTaxCommission) UnmarshalJSON(data []byte) (err error) {
	varAccountCommissionResponseTaxCommission := _AccountCommissionResponseTaxCommission{}

	err = json.Unmarshal(data, &varAccountCommissionResponseTaxCommission)

	if err != nil {
		return err
	}

	*o = AccountCommissionResponseTaxCommission(varAccountCommissionResponseTaxCommission)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "maker")
		delete(additionalProperties, "taker")
		delete(additionalProperties, "buyer")
		delete(additionalProperties, "seller")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountCommissionResponseTaxCommission struct {
	value *AccountCommissionResponseTaxCommission
	isSet bool
}

func (v NullableAccountCommissionResponseTaxCommission) Get() *AccountCommissionResponseTaxCommission {
	return v.value
}

func (v *NullableAccountCommissionResponseTaxCommission) Set(val *AccountCommissionResponseTaxCommission) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCommissionResponseTaxCommission) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCommissionResponseTaxCommission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCommissionResponseTaxCommission(val *AccountCommissionResponseTaxCommission) *NullableAccountCommissionResponseTaxCommission {
	return &NullableAccountCommissionResponseTaxCommission{value: val, isSet: true}
}

func (v NullableAccountCommissionResponseTaxCommission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCommissionResponseTaxCommission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
