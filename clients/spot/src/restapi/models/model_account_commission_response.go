/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountCommissionResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountCommissionResponse{}

// AccountCommissionResponse struct for AccountCommissionResponse
type AccountCommissionResponse struct {
	Symbol               *string                                      `json:"symbol,omitempty"`
	StandardCommission   *AccountCommissionResponseStandardCommission `json:"standardCommission,omitempty"`
	SpecialCommission    *AccountCommissionResponseSpecialCommission  `json:"specialCommission,omitempty"`
	TaxCommission        *AccountCommissionResponseTaxCommission      `json:"taxCommission,omitempty"`
	Discount             *AccountCommissionResponseDiscount           `json:"discount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountCommissionResponse AccountCommissionResponse

// NewAccountCommissionResponse instantiates a new AccountCommissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCommissionResponse() *AccountCommissionResponse {
	this := AccountCommissionResponse{}
	return &this
}

// NewAccountCommissionResponseWithDefaults instantiates a new AccountCommissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCommissionResponseWithDefaults() *AccountCommissionResponse {
	this := AccountCommissionResponse{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *AccountCommissionResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *AccountCommissionResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *AccountCommissionResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetStandardCommission returns the StandardCommission field value if set, zero value otherwise.
func (o *AccountCommissionResponse) GetStandardCommission() AccountCommissionResponseStandardCommission {
	if o == nil || common.IsNil(o.StandardCommission) {
		var ret AccountCommissionResponseStandardCommission
		return ret
	}
	return *o.StandardCommission
}

// GetStandardCommissionOk returns a tuple with the StandardCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponse) GetStandardCommissionOk() (*AccountCommissionResponseStandardCommission, bool) {
	if o == nil || common.IsNil(o.StandardCommission) {
		return nil, false
	}
	return o.StandardCommission, true
}

// HasStandardCommission returns a boolean if a field has been set.
func (o *AccountCommissionResponse) HasStandardCommission() bool {
	if o != nil && !common.IsNil(o.StandardCommission) {
		return true
	}

	return false
}

// SetStandardCommission gets a reference to the given AccountCommissionResponseStandardCommission and assigns it to the StandardCommission field.
func (o *AccountCommissionResponse) SetStandardCommission(v AccountCommissionResponseStandardCommission) {
	o.StandardCommission = &v
}

// GetSpecialCommission returns the SpecialCommission field value if set, zero value otherwise.
func (o *AccountCommissionResponse) GetSpecialCommission() AccountCommissionResponseSpecialCommission {
	if o == nil || common.IsNil(o.SpecialCommission) {
		var ret AccountCommissionResponseSpecialCommission
		return ret
	}
	return *o.SpecialCommission
}

// GetSpecialCommissionOk returns a tuple with the SpecialCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponse) GetSpecialCommissionOk() (*AccountCommissionResponseSpecialCommission, bool) {
	if o == nil || common.IsNil(o.SpecialCommission) {
		return nil, false
	}
	return o.SpecialCommission, true
}

// HasSpecialCommission returns a boolean if a field has been set.
func (o *AccountCommissionResponse) HasSpecialCommission() bool {
	if o != nil && !common.IsNil(o.SpecialCommission) {
		return true
	}

	return false
}

// SetSpecialCommission gets a reference to the given AccountCommissionResponseSpecialCommission and assigns it to the SpecialCommission field.
func (o *AccountCommissionResponse) SetSpecialCommission(v AccountCommissionResponseSpecialCommission) {
	o.SpecialCommission = &v
}

// GetTaxCommission returns the TaxCommission field value if set, zero value otherwise.
func (o *AccountCommissionResponse) GetTaxCommission() AccountCommissionResponseTaxCommission {
	if o == nil || common.IsNil(o.TaxCommission) {
		var ret AccountCommissionResponseTaxCommission
		return ret
	}
	return *o.TaxCommission
}

// GetTaxCommissionOk returns a tuple with the TaxCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponse) GetTaxCommissionOk() (*AccountCommissionResponseTaxCommission, bool) {
	if o == nil || common.IsNil(o.TaxCommission) {
		return nil, false
	}
	return o.TaxCommission, true
}

// HasTaxCommission returns a boolean if a field has been set.
func (o *AccountCommissionResponse) HasTaxCommission() bool {
	if o != nil && !common.IsNil(o.TaxCommission) {
		return true
	}

	return false
}

// SetTaxCommission gets a reference to the given AccountCommissionResponseTaxCommission and assigns it to the TaxCommission field.
func (o *AccountCommissionResponse) SetTaxCommission(v AccountCommissionResponseTaxCommission) {
	o.TaxCommission = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *AccountCommissionResponse) GetDiscount() AccountCommissionResponseDiscount {
	if o == nil || common.IsNil(o.Discount) {
		var ret AccountCommissionResponseDiscount
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponse) GetDiscountOk() (*AccountCommissionResponseDiscount, bool) {
	if o == nil || common.IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *AccountCommissionResponse) HasDiscount() bool {
	if o != nil && !common.IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given AccountCommissionResponseDiscount and assigns it to the Discount field.
func (o *AccountCommissionResponse) SetDiscount(v AccountCommissionResponseDiscount) {
	o.Discount = &v
}

func (o AccountCommissionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCommissionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.StandardCommission) {
		toSerialize["standardCommission"] = o.StandardCommission
	}
	if !common.IsNil(o.SpecialCommission) {
		toSerialize["specialCommission"] = o.SpecialCommission
	}
	if !common.IsNil(o.TaxCommission) {
		toSerialize["taxCommission"] = o.TaxCommission
	}
	if !common.IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountCommissionResponse) UnmarshalJSON(data []byte) (err error) {
	varAccountCommissionResponse := _AccountCommissionResponse{}

	err = json.Unmarshal(data, &varAccountCommissionResponse)

	if err != nil {
		return err
	}

	*o = AccountCommissionResponse(varAccountCommissionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "standardCommission")
		delete(additionalProperties, "specialCommission")
		delete(additionalProperties, "taxCommission")
		delete(additionalProperties, "discount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountCommissionResponse struct {
	value *AccountCommissionResponse
	isSet bool
}

func (v NullableAccountCommissionResponse) Get() *AccountCommissionResponse {
	return v.value
}

func (v *NullableAccountCommissionResponse) Set(val *AccountCommissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCommissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCommissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCommissionResponse(val *AccountCommissionResponse) *NullableAccountCommissionResponse {
	return &NullableAccountCommissionResponse{value: val, isSet: true}
}

func (v NullableAccountCommissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCommissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
