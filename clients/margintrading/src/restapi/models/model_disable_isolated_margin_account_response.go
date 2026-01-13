/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DisableIsolatedMarginAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DisableIsolatedMarginAccountResponse{}

// DisableIsolatedMarginAccountResponse struct for DisableIsolatedMarginAccountResponse
type DisableIsolatedMarginAccountResponse struct {
	Success              *bool   `json:"success,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DisableIsolatedMarginAccountResponse DisableIsolatedMarginAccountResponse

// NewDisableIsolatedMarginAccountResponse instantiates a new DisableIsolatedMarginAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisableIsolatedMarginAccountResponse() *DisableIsolatedMarginAccountResponse {
	this := DisableIsolatedMarginAccountResponse{}
	return &this
}

// NewDisableIsolatedMarginAccountResponseWithDefaults instantiates a new DisableIsolatedMarginAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisableIsolatedMarginAccountResponseWithDefaults() *DisableIsolatedMarginAccountResponse {
	this := DisableIsolatedMarginAccountResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DisableIsolatedMarginAccountResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisableIsolatedMarginAccountResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DisableIsolatedMarginAccountResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DisableIsolatedMarginAccountResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *DisableIsolatedMarginAccountResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisableIsolatedMarginAccountResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *DisableIsolatedMarginAccountResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *DisableIsolatedMarginAccountResponse) SetSymbol(v string) {
	o.Symbol = &v
}

func (o DisableIsolatedMarginAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisableIsolatedMarginAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DisableIsolatedMarginAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varDisableIsolatedMarginAccountResponse := _DisableIsolatedMarginAccountResponse{}

	err = json.Unmarshal(data, &varDisableIsolatedMarginAccountResponse)

	if err != nil {
		return err
	}

	*o = DisableIsolatedMarginAccountResponse(varDisableIsolatedMarginAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "symbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDisableIsolatedMarginAccountResponse struct {
	value *DisableIsolatedMarginAccountResponse
	isSet bool
}

func (v NullableDisableIsolatedMarginAccountResponse) Get() *DisableIsolatedMarginAccountResponse {
	return v.value
}

func (v *NullableDisableIsolatedMarginAccountResponse) Set(val *DisableIsolatedMarginAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDisableIsolatedMarginAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDisableIsolatedMarginAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisableIsolatedMarginAccountResponse(val *DisableIsolatedMarginAccountResponse) *NullableDisableIsolatedMarginAccountResponse {
	return &NullableDisableIsolatedMarginAccountResponse{value: val, isSet: true}
}

func (v NullableDisableIsolatedMarginAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisableIsolatedMarginAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
