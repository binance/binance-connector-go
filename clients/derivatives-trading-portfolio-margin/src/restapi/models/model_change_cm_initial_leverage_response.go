/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ChangeCmInitialLeverageResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ChangeCmInitialLeverageResponse{}

// ChangeCmInitialLeverageResponse struct for ChangeCmInitialLeverageResponse
type ChangeCmInitialLeverageResponse struct {
	Leverage             *int64  `json:"leverage,omitempty"`
	MaxQty               *string `json:"maxQty,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChangeCmInitialLeverageResponse ChangeCmInitialLeverageResponse

// NewChangeCmInitialLeverageResponse instantiates a new ChangeCmInitialLeverageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeCmInitialLeverageResponse() *ChangeCmInitialLeverageResponse {
	this := ChangeCmInitialLeverageResponse{}
	return &this
}

// NewChangeCmInitialLeverageResponseWithDefaults instantiates a new ChangeCmInitialLeverageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeCmInitialLeverageResponseWithDefaults() *ChangeCmInitialLeverageResponse {
	this := ChangeCmInitialLeverageResponse{}
	return &this
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *ChangeCmInitialLeverageResponse) GetLeverage() int64 {
	if o == nil || common.IsNil(o.Leverage) {
		var ret int64
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeCmInitialLeverageResponse) GetLeverageOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *ChangeCmInitialLeverageResponse) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given int64 and assigns it to the Leverage field.
func (o *ChangeCmInitialLeverageResponse) SetLeverage(v int64) {
	o.Leverage = &v
}

// GetMaxQty returns the MaxQty field value if set, zero value otherwise.
func (o *ChangeCmInitialLeverageResponse) GetMaxQty() string {
	if o == nil || common.IsNil(o.MaxQty) {
		var ret string
		return ret
	}
	return *o.MaxQty
}

// GetMaxQtyOk returns a tuple with the MaxQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeCmInitialLeverageResponse) GetMaxQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxQty) {
		return nil, false
	}
	return o.MaxQty, true
}

// HasMaxQty returns a boolean if a field has been set.
func (o *ChangeCmInitialLeverageResponse) HasMaxQty() bool {
	if o != nil && !common.IsNil(o.MaxQty) {
		return true
	}

	return false
}

// SetMaxQty gets a reference to the given string and assigns it to the MaxQty field.
func (o *ChangeCmInitialLeverageResponse) SetMaxQty(v string) {
	o.MaxQty = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ChangeCmInitialLeverageResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeCmInitialLeverageResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ChangeCmInitialLeverageResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ChangeCmInitialLeverageResponse) SetSymbol(v string) {
	o.Symbol = &v
}

func (o ChangeCmInitialLeverageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeCmInitialLeverageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}
	if !common.IsNil(o.MaxQty) {
		toSerialize["maxQty"] = o.MaxQty
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChangeCmInitialLeverageResponse) UnmarshalJSON(data []byte) (err error) {
	varChangeCmInitialLeverageResponse := _ChangeCmInitialLeverageResponse{}

	err = json.Unmarshal(data, &varChangeCmInitialLeverageResponse)

	if err != nil {
		return err
	}

	*o = ChangeCmInitialLeverageResponse(varChangeCmInitialLeverageResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "leverage")
		delete(additionalProperties, "maxQty")
		delete(additionalProperties, "symbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChangeCmInitialLeverageResponse struct {
	value *ChangeCmInitialLeverageResponse
	isSet bool
}

func (v NullableChangeCmInitialLeverageResponse) Get() *ChangeCmInitialLeverageResponse {
	return v.value
}

func (v *NullableChangeCmInitialLeverageResponse) Set(val *ChangeCmInitialLeverageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeCmInitialLeverageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeCmInitialLeverageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeCmInitialLeverageResponse(val *ChangeCmInitialLeverageResponse) *NullableChangeCmInitialLeverageResponse {
	return &NullableChangeCmInitialLeverageResponse{value: val, isSet: true}
}

func (v NullableChangeCmInitialLeverageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeCmInitialLeverageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
