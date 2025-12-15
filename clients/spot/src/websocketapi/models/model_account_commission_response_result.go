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

// checks if the AccountCommissionResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountCommissionResponseResult{}

// AccountCommissionResponseResult struct for AccountCommissionResponseResult
type AccountCommissionResponseResult struct {
	Symbol               *string                                            `json:"symbol,omitempty"`
	StandardCommission   *AccountCommissionResponseResultStandardCommission `json:"standardCommission,omitempty"`
	SpecialCommission    *AccountCommissionResponseResultSpecialCommission  `json:"specialCommission,omitempty"`
	TaxCommission        *AccountCommissionResponseResultTaxCommission      `json:"taxCommission,omitempty"`
	Discount             *AccountCommissionResponseResultDiscount           `json:"discount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountCommissionResponseResult AccountCommissionResponseResult

// NewAccountCommissionResponseResult instantiates a new AccountCommissionResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCommissionResponseResult() *AccountCommissionResponseResult {
	this := AccountCommissionResponseResult{}
	return &this
}

// NewAccountCommissionResponseResultWithDefaults instantiates a new AccountCommissionResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCommissionResponseResultWithDefaults() *AccountCommissionResponseResult {
	this := AccountCommissionResponseResult{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *AccountCommissionResponseResult) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseResult) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *AccountCommissionResponseResult) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *AccountCommissionResponseResult) SetSymbol(v string) {
	o.Symbol = &v
}

// GetStandardCommission returns the StandardCommission field value if set, zero value otherwise.
func (o *AccountCommissionResponseResult) GetStandardCommission() AccountCommissionResponseResultStandardCommission {
	if o == nil || common.IsNil(o.StandardCommission) {
		var ret AccountCommissionResponseResultStandardCommission
		return ret
	}
	return *o.StandardCommission
}

// GetStandardCommissionOk returns a tuple with the StandardCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseResult) GetStandardCommissionOk() (*AccountCommissionResponseResultStandardCommission, bool) {
	if o == nil || common.IsNil(o.StandardCommission) {
		return nil, false
	}
	return o.StandardCommission, true
}

// HasStandardCommission returns a boolean if a field has been set.
func (o *AccountCommissionResponseResult) HasStandardCommission() bool {
	if o != nil && !common.IsNil(o.StandardCommission) {
		return true
	}

	return false
}

// SetStandardCommission gets a reference to the given AccountCommissionResponseResultStandardCommission and assigns it to the StandardCommission field.
func (o *AccountCommissionResponseResult) SetStandardCommission(v AccountCommissionResponseResultStandardCommission) {
	o.StandardCommission = &v
}

// GetSpecialCommission returns the SpecialCommission field value if set, zero value otherwise.
func (o *AccountCommissionResponseResult) GetSpecialCommission() AccountCommissionResponseResultSpecialCommission {
	if o == nil || common.IsNil(o.SpecialCommission) {
		var ret AccountCommissionResponseResultSpecialCommission
		return ret
	}
	return *o.SpecialCommission
}

// GetSpecialCommissionOk returns a tuple with the SpecialCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseResult) GetSpecialCommissionOk() (*AccountCommissionResponseResultSpecialCommission, bool) {
	if o == nil || common.IsNil(o.SpecialCommission) {
		return nil, false
	}
	return o.SpecialCommission, true
}

// HasSpecialCommission returns a boolean if a field has been set.
func (o *AccountCommissionResponseResult) HasSpecialCommission() bool {
	if o != nil && !common.IsNil(o.SpecialCommission) {
		return true
	}

	return false
}

// SetSpecialCommission gets a reference to the given AccountCommissionResponseResultSpecialCommission and assigns it to the SpecialCommission field.
func (o *AccountCommissionResponseResult) SetSpecialCommission(v AccountCommissionResponseResultSpecialCommission) {
	o.SpecialCommission = &v
}

// GetTaxCommission returns the TaxCommission field value if set, zero value otherwise.
func (o *AccountCommissionResponseResult) GetTaxCommission() AccountCommissionResponseResultTaxCommission {
	if o == nil || common.IsNil(o.TaxCommission) {
		var ret AccountCommissionResponseResultTaxCommission
		return ret
	}
	return *o.TaxCommission
}

// GetTaxCommissionOk returns a tuple with the TaxCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseResult) GetTaxCommissionOk() (*AccountCommissionResponseResultTaxCommission, bool) {
	if o == nil || common.IsNil(o.TaxCommission) {
		return nil, false
	}
	return o.TaxCommission, true
}

// HasTaxCommission returns a boolean if a field has been set.
func (o *AccountCommissionResponseResult) HasTaxCommission() bool {
	if o != nil && !common.IsNil(o.TaxCommission) {
		return true
	}

	return false
}

// SetTaxCommission gets a reference to the given AccountCommissionResponseResultTaxCommission and assigns it to the TaxCommission field.
func (o *AccountCommissionResponseResult) SetTaxCommission(v AccountCommissionResponseResultTaxCommission) {
	o.TaxCommission = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *AccountCommissionResponseResult) GetDiscount() AccountCommissionResponseResultDiscount {
	if o == nil || common.IsNil(o.Discount) {
		var ret AccountCommissionResponseResultDiscount
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseResult) GetDiscountOk() (*AccountCommissionResponseResultDiscount, bool) {
	if o == nil || common.IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *AccountCommissionResponseResult) HasDiscount() bool {
	if o != nil && !common.IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given AccountCommissionResponseResultDiscount and assigns it to the Discount field.
func (o *AccountCommissionResponseResult) SetDiscount(v AccountCommissionResponseResultDiscount) {
	o.Discount = &v
}

func (o AccountCommissionResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCommissionResponseResult) ToMap() (map[string]interface{}, error) {
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

func (o *AccountCommissionResponseResult) UnmarshalJSON(data []byte) (err error) {
	varAccountCommissionResponseResult := _AccountCommissionResponseResult{}

	err = json.Unmarshal(data, &varAccountCommissionResponseResult)

	if err != nil {
		return err
	}

	*o = AccountCommissionResponseResult(varAccountCommissionResponseResult)

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

type NullableAccountCommissionResponseResult struct {
	value *AccountCommissionResponseResult
	isSet bool
}

func (v NullableAccountCommissionResponseResult) Get() *AccountCommissionResponseResult {
	return v.value
}

func (v *NullableAccountCommissionResponseResult) Set(val *AccountCommissionResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCommissionResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCommissionResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCommissionResponseResult(val *AccountCommissionResponseResult) *NullableAccountCommissionResponseResult {
	return &NullableAccountCommissionResponseResult{value: val, isSet: true}
}

func (v NullableAccountCommissionResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCommissionResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
