/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AutoCancelAllOpenOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AutoCancelAllOpenOrdersResponse{}

// AutoCancelAllOpenOrdersResponse struct for AutoCancelAllOpenOrdersResponse
type AutoCancelAllOpenOrdersResponse struct {
	Symbol               *string `json:"symbol,omitempty"`
	CountdownTime        *string `json:"countdownTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AutoCancelAllOpenOrdersResponse AutoCancelAllOpenOrdersResponse

// NewAutoCancelAllOpenOrdersResponse instantiates a new AutoCancelAllOpenOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoCancelAllOpenOrdersResponse() *AutoCancelAllOpenOrdersResponse {
	this := AutoCancelAllOpenOrdersResponse{}
	return &this
}

// NewAutoCancelAllOpenOrdersResponseWithDefaults instantiates a new AutoCancelAllOpenOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoCancelAllOpenOrdersResponseWithDefaults() *AutoCancelAllOpenOrdersResponse {
	this := AutoCancelAllOpenOrdersResponse{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *AutoCancelAllOpenOrdersResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoCancelAllOpenOrdersResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *AutoCancelAllOpenOrdersResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *AutoCancelAllOpenOrdersResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetCountdownTime returns the CountdownTime field value if set, zero value otherwise.
func (o *AutoCancelAllOpenOrdersResponse) GetCountdownTime() string {
	if o == nil || common.IsNil(o.CountdownTime) {
		var ret string
		return ret
	}
	return *o.CountdownTime
}

// GetCountdownTimeOk returns a tuple with the CountdownTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoCancelAllOpenOrdersResponse) GetCountdownTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CountdownTime) {
		return nil, false
	}
	return o.CountdownTime, true
}

// HasCountdownTime returns a boolean if a field has been set.
func (o *AutoCancelAllOpenOrdersResponse) HasCountdownTime() bool {
	if o != nil && !common.IsNil(o.CountdownTime) {
		return true
	}

	return false
}

// SetCountdownTime gets a reference to the given string and assigns it to the CountdownTime field.
func (o *AutoCancelAllOpenOrdersResponse) SetCountdownTime(v string) {
	o.CountdownTime = &v
}

func (o AutoCancelAllOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoCancelAllOpenOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.CountdownTime) {
		toSerialize["countdownTime"] = o.CountdownTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AutoCancelAllOpenOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	varAutoCancelAllOpenOrdersResponse := _AutoCancelAllOpenOrdersResponse{}

	err = json.Unmarshal(data, &varAutoCancelAllOpenOrdersResponse)

	if err != nil {
		return err
	}

	*o = AutoCancelAllOpenOrdersResponse(varAutoCancelAllOpenOrdersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "countdownTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAutoCancelAllOpenOrdersResponse struct {
	value *AutoCancelAllOpenOrdersResponse
	isSet bool
}

func (v NullableAutoCancelAllOpenOrdersResponse) Get() *AutoCancelAllOpenOrdersResponse {
	return v.value
}

func (v *NullableAutoCancelAllOpenOrdersResponse) Set(val *AutoCancelAllOpenOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoCancelAllOpenOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoCancelAllOpenOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoCancelAllOpenOrdersResponse(val *AutoCancelAllOpenOrdersResponse) *NullableAutoCancelAllOpenOrdersResponse {
	return &NullableAutoCancelAllOpenOrdersResponse{value: val, isSet: true}
}

func (v NullableAutoCancelAllOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoCancelAllOpenOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
