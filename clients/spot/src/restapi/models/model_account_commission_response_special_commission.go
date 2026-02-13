/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountCommissionResponseSpecialCommission type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountCommissionResponseSpecialCommission{}

// AccountCommissionResponseSpecialCommission struct for AccountCommissionResponseSpecialCommission
type AccountCommissionResponseSpecialCommission struct {
	Maker                *string `json:"maker,omitempty"`
	Taker                *string `json:"taker,omitempty"`
	Buyer                *string `json:"buyer,omitempty"`
	Seller               *string `json:"seller,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountCommissionResponseSpecialCommission AccountCommissionResponseSpecialCommission

// NewAccountCommissionResponseSpecialCommission instantiates a new AccountCommissionResponseSpecialCommission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCommissionResponseSpecialCommission() *AccountCommissionResponseSpecialCommission {
	this := AccountCommissionResponseSpecialCommission{}
	return &this
}

// NewAccountCommissionResponseSpecialCommissionWithDefaults instantiates a new AccountCommissionResponseSpecialCommission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCommissionResponseSpecialCommissionWithDefaults() *AccountCommissionResponseSpecialCommission {
	this := AccountCommissionResponseSpecialCommission{}
	return &this
}

// GetMaker returns the Maker field value if set, zero value otherwise.
func (o *AccountCommissionResponseSpecialCommission) GetMaker() string {
	if o == nil || common.IsNil(o.Maker) {
		var ret string
		return ret
	}
	return *o.Maker
}

// GetMakerOk returns a tuple with the Maker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseSpecialCommission) GetMakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Maker) {
		return nil, false
	}
	return o.Maker, true
}

// HasMaker returns a boolean if a field has been set.
func (o *AccountCommissionResponseSpecialCommission) HasMaker() bool {
	if o != nil && !common.IsNil(o.Maker) {
		return true
	}

	return false
}

// SetMaker gets a reference to the given string and assigns it to the Maker field.
func (o *AccountCommissionResponseSpecialCommission) SetMaker(v string) {
	o.Maker = &v
}

// GetTaker returns the Taker field value if set, zero value otherwise.
func (o *AccountCommissionResponseSpecialCommission) GetTaker() string {
	if o == nil || common.IsNil(o.Taker) {
		var ret string
		return ret
	}
	return *o.Taker
}

// GetTakerOk returns a tuple with the Taker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseSpecialCommission) GetTakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Taker) {
		return nil, false
	}
	return o.Taker, true
}

// HasTaker returns a boolean if a field has been set.
func (o *AccountCommissionResponseSpecialCommission) HasTaker() bool {
	if o != nil && !common.IsNil(o.Taker) {
		return true
	}

	return false
}

// SetTaker gets a reference to the given string and assigns it to the Taker field.
func (o *AccountCommissionResponseSpecialCommission) SetTaker(v string) {
	o.Taker = &v
}

// GetBuyer returns the Buyer field value if set, zero value otherwise.
func (o *AccountCommissionResponseSpecialCommission) GetBuyer() string {
	if o == nil || common.IsNil(o.Buyer) {
		var ret string
		return ret
	}
	return *o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseSpecialCommission) GetBuyerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Buyer) {
		return nil, false
	}
	return o.Buyer, true
}

// HasBuyer returns a boolean if a field has been set.
func (o *AccountCommissionResponseSpecialCommission) HasBuyer() bool {
	if o != nil && !common.IsNil(o.Buyer) {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given string and assigns it to the Buyer field.
func (o *AccountCommissionResponseSpecialCommission) SetBuyer(v string) {
	o.Buyer = &v
}

// GetSeller returns the Seller field value if set, zero value otherwise.
func (o *AccountCommissionResponseSpecialCommission) GetSeller() string {
	if o == nil || common.IsNil(o.Seller) {
		var ret string
		return ret
	}
	return *o.Seller
}

// GetSellerOk returns a tuple with the Seller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseSpecialCommission) GetSellerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Seller) {
		return nil, false
	}
	return o.Seller, true
}

// HasSeller returns a boolean if a field has been set.
func (o *AccountCommissionResponseSpecialCommission) HasSeller() bool {
	if o != nil && !common.IsNil(o.Seller) {
		return true
	}

	return false
}

// SetSeller gets a reference to the given string and assigns it to the Seller field.
func (o *AccountCommissionResponseSpecialCommission) SetSeller(v string) {
	o.Seller = &v
}

func (o AccountCommissionResponseSpecialCommission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCommissionResponseSpecialCommission) ToMap() (map[string]interface{}, error) {
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

func (o *AccountCommissionResponseSpecialCommission) UnmarshalJSON(data []byte) (err error) {
	varAccountCommissionResponseSpecialCommission := _AccountCommissionResponseSpecialCommission{}

	err = json.Unmarshal(data, &varAccountCommissionResponseSpecialCommission)

	if err != nil {
		return err
	}

	*o = AccountCommissionResponseSpecialCommission(varAccountCommissionResponseSpecialCommission)

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

type NullableAccountCommissionResponseSpecialCommission struct {
	value *AccountCommissionResponseSpecialCommission
	isSet bool
}

func (v NullableAccountCommissionResponseSpecialCommission) Get() *AccountCommissionResponseSpecialCommission {
	return v.value
}

func (v *NullableAccountCommissionResponseSpecialCommission) Set(val *AccountCommissionResponseSpecialCommission) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCommissionResponseSpecialCommission) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCommissionResponseSpecialCommission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCommissionResponseSpecialCommission(val *AccountCommissionResponseSpecialCommission) *NullableAccountCommissionResponseSpecialCommission {
	return &NullableAccountCommissionResponseSpecialCommission{value: val, isSet: true}
}

func (v NullableAccountCommissionResponseSpecialCommission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCommissionResponseSpecialCommission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
