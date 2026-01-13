/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountCommissionResponseStandardCommission type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountCommissionResponseStandardCommission{}

// AccountCommissionResponseStandardCommission struct for AccountCommissionResponseStandardCommission
type AccountCommissionResponseStandardCommission struct {
	Maker                *string `json:"maker,omitempty"`
	Taker                *string `json:"taker,omitempty"`
	Buyer                *string `json:"buyer,omitempty"`
	Seller               *string `json:"seller,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountCommissionResponseStandardCommission AccountCommissionResponseStandardCommission

// NewAccountCommissionResponseStandardCommission instantiates a new AccountCommissionResponseStandardCommission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCommissionResponseStandardCommission() *AccountCommissionResponseStandardCommission {
	this := AccountCommissionResponseStandardCommission{}
	return &this
}

// NewAccountCommissionResponseStandardCommissionWithDefaults instantiates a new AccountCommissionResponseStandardCommission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCommissionResponseStandardCommissionWithDefaults() *AccountCommissionResponseStandardCommission {
	this := AccountCommissionResponseStandardCommission{}
	return &this
}

// GetMaker returns the Maker field value if set, zero value otherwise.
func (o *AccountCommissionResponseStandardCommission) GetMaker() string {
	if o == nil || common.IsNil(o.Maker) {
		var ret string
		return ret
	}
	return *o.Maker
}

// GetMakerOk returns a tuple with the Maker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseStandardCommission) GetMakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Maker) {
		return nil, false
	}
	return o.Maker, true
}

// HasMaker returns a boolean if a field has been set.
func (o *AccountCommissionResponseStandardCommission) HasMaker() bool {
	if o != nil && !common.IsNil(o.Maker) {
		return true
	}

	return false
}

// SetMaker gets a reference to the given string and assigns it to the Maker field.
func (o *AccountCommissionResponseStandardCommission) SetMaker(v string) {
	o.Maker = &v
}

// GetTaker returns the Taker field value if set, zero value otherwise.
func (o *AccountCommissionResponseStandardCommission) GetTaker() string {
	if o == nil || common.IsNil(o.Taker) {
		var ret string
		return ret
	}
	return *o.Taker
}

// GetTakerOk returns a tuple with the Taker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseStandardCommission) GetTakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Taker) {
		return nil, false
	}
	return o.Taker, true
}

// HasTaker returns a boolean if a field has been set.
func (o *AccountCommissionResponseStandardCommission) HasTaker() bool {
	if o != nil && !common.IsNil(o.Taker) {
		return true
	}

	return false
}

// SetTaker gets a reference to the given string and assigns it to the Taker field.
func (o *AccountCommissionResponseStandardCommission) SetTaker(v string) {
	o.Taker = &v
}

// GetBuyer returns the Buyer field value if set, zero value otherwise.
func (o *AccountCommissionResponseStandardCommission) GetBuyer() string {
	if o == nil || common.IsNil(o.Buyer) {
		var ret string
		return ret
	}
	return *o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseStandardCommission) GetBuyerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Buyer) {
		return nil, false
	}
	return o.Buyer, true
}

// HasBuyer returns a boolean if a field has been set.
func (o *AccountCommissionResponseStandardCommission) HasBuyer() bool {
	if o != nil && !common.IsNil(o.Buyer) {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given string and assigns it to the Buyer field.
func (o *AccountCommissionResponseStandardCommission) SetBuyer(v string) {
	o.Buyer = &v
}

// GetSeller returns the Seller field value if set, zero value otherwise.
func (o *AccountCommissionResponseStandardCommission) GetSeller() string {
	if o == nil || common.IsNil(o.Seller) {
		var ret string
		return ret
	}
	return *o.Seller
}

// GetSellerOk returns a tuple with the Seller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseStandardCommission) GetSellerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Seller) {
		return nil, false
	}
	return o.Seller, true
}

// HasSeller returns a boolean if a field has been set.
func (o *AccountCommissionResponseStandardCommission) HasSeller() bool {
	if o != nil && !common.IsNil(o.Seller) {
		return true
	}

	return false
}

// SetSeller gets a reference to the given string and assigns it to the Seller field.
func (o *AccountCommissionResponseStandardCommission) SetSeller(v string) {
	o.Seller = &v
}

func (o AccountCommissionResponseStandardCommission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCommissionResponseStandardCommission) ToMap() (map[string]interface{}, error) {
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

func (o *AccountCommissionResponseStandardCommission) UnmarshalJSON(data []byte) (err error) {
	varAccountCommissionResponseStandardCommission := _AccountCommissionResponseStandardCommission{}

	err = json.Unmarshal(data, &varAccountCommissionResponseStandardCommission)

	if err != nil {
		return err
	}

	*o = AccountCommissionResponseStandardCommission(varAccountCommissionResponseStandardCommission)

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

type NullableAccountCommissionResponseStandardCommission struct {
	value *AccountCommissionResponseStandardCommission
	isSet bool
}

func (v NullableAccountCommissionResponseStandardCommission) Get() *AccountCommissionResponseStandardCommission {
	return v.value
}

func (v *NullableAccountCommissionResponseStandardCommission) Set(val *AccountCommissionResponseStandardCommission) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCommissionResponseStandardCommission) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCommissionResponseStandardCommission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCommissionResponseStandardCommission(val *AccountCommissionResponseStandardCommission) *NullableAccountCommissionResponseStandardCommission {
	return &NullableAccountCommissionResponseStandardCommission{value: val, isSet: true}
}

func (v NullableAccountCommissionResponseStandardCommission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCommissionResponseStandardCommission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
