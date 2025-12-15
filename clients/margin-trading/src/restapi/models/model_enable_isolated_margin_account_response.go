/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the EnableIsolatedMarginAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EnableIsolatedMarginAccountResponse{}

// EnableIsolatedMarginAccountResponse struct for EnableIsolatedMarginAccountResponse
type EnableIsolatedMarginAccountResponse struct {
	Success              *bool   `json:"success,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnableIsolatedMarginAccountResponse EnableIsolatedMarginAccountResponse

// NewEnableIsolatedMarginAccountResponse instantiates a new EnableIsolatedMarginAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableIsolatedMarginAccountResponse() *EnableIsolatedMarginAccountResponse {
	this := EnableIsolatedMarginAccountResponse{}
	return &this
}

// NewEnableIsolatedMarginAccountResponseWithDefaults instantiates a new EnableIsolatedMarginAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableIsolatedMarginAccountResponseWithDefaults() *EnableIsolatedMarginAccountResponse {
	this := EnableIsolatedMarginAccountResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *EnableIsolatedMarginAccountResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableIsolatedMarginAccountResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *EnableIsolatedMarginAccountResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *EnableIsolatedMarginAccountResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *EnableIsolatedMarginAccountResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableIsolatedMarginAccountResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *EnableIsolatedMarginAccountResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *EnableIsolatedMarginAccountResponse) SetSymbol(v string) {
	o.Symbol = &v
}

func (o EnableIsolatedMarginAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableIsolatedMarginAccountResponse) ToMap() (map[string]interface{}, error) {
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

func (o *EnableIsolatedMarginAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varEnableIsolatedMarginAccountResponse := _EnableIsolatedMarginAccountResponse{}

	err = json.Unmarshal(data, &varEnableIsolatedMarginAccountResponse)

	if err != nil {
		return err
	}

	*o = EnableIsolatedMarginAccountResponse(varEnableIsolatedMarginAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "symbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnableIsolatedMarginAccountResponse struct {
	value *EnableIsolatedMarginAccountResponse
	isSet bool
}

func (v NullableEnableIsolatedMarginAccountResponse) Get() *EnableIsolatedMarginAccountResponse {
	return v.value
}

func (v *NullableEnableIsolatedMarginAccountResponse) Set(val *EnableIsolatedMarginAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableIsolatedMarginAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableIsolatedMarginAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableIsolatedMarginAccountResponse(val *EnableIsolatedMarginAccountResponse) *NullableEnableIsolatedMarginAccountResponse {
	return &NullableEnableIsolatedMarginAccountResponse{value: val, isSet: true}
}

func (v NullableEnableIsolatedMarginAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableIsolatedMarginAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
