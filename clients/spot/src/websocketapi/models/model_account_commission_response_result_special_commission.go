/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountCommissionResponseResultSpecialCommission type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountCommissionResponseResultSpecialCommission{}

// AccountCommissionResponseResultSpecialCommission struct for AccountCommissionResponseResultSpecialCommission
type AccountCommissionResponseResultSpecialCommission struct {
	Maker                *string `json:"maker,omitempty"`
	Taker                *string `json:"taker,omitempty"`
	Buyer                *string `json:"buyer,omitempty"`
	Seller               *string `json:"seller,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountCommissionResponseResultSpecialCommission AccountCommissionResponseResultSpecialCommission

// NewAccountCommissionResponseResultSpecialCommission instantiates a new AccountCommissionResponseResultSpecialCommission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCommissionResponseResultSpecialCommission() *AccountCommissionResponseResultSpecialCommission {
	this := AccountCommissionResponseResultSpecialCommission{}
	return &this
}

// NewAccountCommissionResponseResultSpecialCommissionWithDefaults instantiates a new AccountCommissionResponseResultSpecialCommission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCommissionResponseResultSpecialCommissionWithDefaults() *AccountCommissionResponseResultSpecialCommission {
	this := AccountCommissionResponseResultSpecialCommission{}
	return &this
}

// GetMaker returns the Maker field value if set, zero value otherwise.
func (o *AccountCommissionResponseResultSpecialCommission) GetMaker() string {
	if o == nil || common.IsNil(o.Maker) {
		var ret string
		return ret
	}
	return *o.Maker
}

// GetMakerOk returns a tuple with the Maker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseResultSpecialCommission) GetMakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Maker) {
		return nil, false
	}
	return o.Maker, true
}

// HasMaker returns a boolean if a field has been set.
func (o *AccountCommissionResponseResultSpecialCommission) HasMaker() bool {
	if o != nil && !common.IsNil(o.Maker) {
		return true
	}

	return false
}

// SetMaker gets a reference to the given string and assigns it to the Maker field.
func (o *AccountCommissionResponseResultSpecialCommission) SetMaker(v string) {
	o.Maker = &v
}

// GetTaker returns the Taker field value if set, zero value otherwise.
func (o *AccountCommissionResponseResultSpecialCommission) GetTaker() string {
	if o == nil || common.IsNil(o.Taker) {
		var ret string
		return ret
	}
	return *o.Taker
}

// GetTakerOk returns a tuple with the Taker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseResultSpecialCommission) GetTakerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Taker) {
		return nil, false
	}
	return o.Taker, true
}

// HasTaker returns a boolean if a field has been set.
func (o *AccountCommissionResponseResultSpecialCommission) HasTaker() bool {
	if o != nil && !common.IsNil(o.Taker) {
		return true
	}

	return false
}

// SetTaker gets a reference to the given string and assigns it to the Taker field.
func (o *AccountCommissionResponseResultSpecialCommission) SetTaker(v string) {
	o.Taker = &v
}

// GetBuyer returns the Buyer field value if set, zero value otherwise.
func (o *AccountCommissionResponseResultSpecialCommission) GetBuyer() string {
	if o == nil || common.IsNil(o.Buyer) {
		var ret string
		return ret
	}
	return *o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseResultSpecialCommission) GetBuyerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Buyer) {
		return nil, false
	}
	return o.Buyer, true
}

// HasBuyer returns a boolean if a field has been set.
func (o *AccountCommissionResponseResultSpecialCommission) HasBuyer() bool {
	if o != nil && !common.IsNil(o.Buyer) {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given string and assigns it to the Buyer field.
func (o *AccountCommissionResponseResultSpecialCommission) SetBuyer(v string) {
	o.Buyer = &v
}

// GetSeller returns the Seller field value if set, zero value otherwise.
func (o *AccountCommissionResponseResultSpecialCommission) GetSeller() string {
	if o == nil || common.IsNil(o.Seller) {
		var ret string
		return ret
	}
	return *o.Seller
}

// GetSellerOk returns a tuple with the Seller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseResultSpecialCommission) GetSellerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Seller) {
		return nil, false
	}
	return o.Seller, true
}

// HasSeller returns a boolean if a field has been set.
func (o *AccountCommissionResponseResultSpecialCommission) HasSeller() bool {
	if o != nil && !common.IsNil(o.Seller) {
		return true
	}

	return false
}

// SetSeller gets a reference to the given string and assigns it to the Seller field.
func (o *AccountCommissionResponseResultSpecialCommission) SetSeller(v string) {
	o.Seller = &v
}

func (o AccountCommissionResponseResultSpecialCommission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCommissionResponseResultSpecialCommission) ToMap() (map[string]interface{}, error) {
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

func (o *AccountCommissionResponseResultSpecialCommission) UnmarshalJSON(data []byte) (err error) {
	varAccountCommissionResponseResultSpecialCommission := _AccountCommissionResponseResultSpecialCommission{}

	err = json.Unmarshal(data, &varAccountCommissionResponseResultSpecialCommission)

	if err != nil {
		return err
	}

	*o = AccountCommissionResponseResultSpecialCommission(varAccountCommissionResponseResultSpecialCommission)

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

type NullableAccountCommissionResponseResultSpecialCommission struct {
	value *AccountCommissionResponseResultSpecialCommission
	isSet bool
}

func (v NullableAccountCommissionResponseResultSpecialCommission) Get() *AccountCommissionResponseResultSpecialCommission {
	return v.value
}

func (v *NullableAccountCommissionResponseResultSpecialCommission) Set(val *AccountCommissionResponseResultSpecialCommission) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCommissionResponseResultSpecialCommission) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCommissionResponseResultSpecialCommission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCommissionResponseResultSpecialCommission(val *AccountCommissionResponseResultSpecialCommission) *NullableAccountCommissionResponseResultSpecialCommission {
	return &NullableAccountCommissionResponseResultSpecialCommission{value: val, isSet: true}
}

func (v NullableAccountCommissionResponseResultSpecialCommission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCommissionResponseResultSpecialCommission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
