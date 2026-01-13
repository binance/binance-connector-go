/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ChangeUmInitialLeverageResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ChangeUmInitialLeverageResponse{}

// ChangeUmInitialLeverageResponse struct for ChangeUmInitialLeverageResponse
type ChangeUmInitialLeverageResponse struct {
	Leverage             *int64  `json:"leverage,omitempty"`
	MaxNotionalValue     *string `json:"maxNotionalValue,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChangeUmInitialLeverageResponse ChangeUmInitialLeverageResponse

// NewChangeUmInitialLeverageResponse instantiates a new ChangeUmInitialLeverageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeUmInitialLeverageResponse() *ChangeUmInitialLeverageResponse {
	this := ChangeUmInitialLeverageResponse{}
	return &this
}

// NewChangeUmInitialLeverageResponseWithDefaults instantiates a new ChangeUmInitialLeverageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeUmInitialLeverageResponseWithDefaults() *ChangeUmInitialLeverageResponse {
	this := ChangeUmInitialLeverageResponse{}
	return &this
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *ChangeUmInitialLeverageResponse) GetLeverage() int64 {
	if o == nil || common.IsNil(o.Leverage) {
		var ret int64
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeUmInitialLeverageResponse) GetLeverageOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *ChangeUmInitialLeverageResponse) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given int64 and assigns it to the Leverage field.
func (o *ChangeUmInitialLeverageResponse) SetLeverage(v int64) {
	o.Leverage = &v
}

// GetMaxNotionalValue returns the MaxNotionalValue field value if set, zero value otherwise.
func (o *ChangeUmInitialLeverageResponse) GetMaxNotionalValue() string {
	if o == nil || common.IsNil(o.MaxNotionalValue) {
		var ret string
		return ret
	}
	return *o.MaxNotionalValue
}

// GetMaxNotionalValueOk returns a tuple with the MaxNotionalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeUmInitialLeverageResponse) GetMaxNotionalValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxNotionalValue) {
		return nil, false
	}
	return o.MaxNotionalValue, true
}

// HasMaxNotionalValue returns a boolean if a field has been set.
func (o *ChangeUmInitialLeverageResponse) HasMaxNotionalValue() bool {
	if o != nil && !common.IsNil(o.MaxNotionalValue) {
		return true
	}

	return false
}

// SetMaxNotionalValue gets a reference to the given string and assigns it to the MaxNotionalValue field.
func (o *ChangeUmInitialLeverageResponse) SetMaxNotionalValue(v string) {
	o.MaxNotionalValue = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ChangeUmInitialLeverageResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeUmInitialLeverageResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ChangeUmInitialLeverageResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ChangeUmInitialLeverageResponse) SetSymbol(v string) {
	o.Symbol = &v
}

func (o ChangeUmInitialLeverageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeUmInitialLeverageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}
	if !common.IsNil(o.MaxNotionalValue) {
		toSerialize["maxNotionalValue"] = o.MaxNotionalValue
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChangeUmInitialLeverageResponse) UnmarshalJSON(data []byte) (err error) {
	varChangeUmInitialLeverageResponse := _ChangeUmInitialLeverageResponse{}

	err = json.Unmarshal(data, &varChangeUmInitialLeverageResponse)

	if err != nil {
		return err
	}

	*o = ChangeUmInitialLeverageResponse(varChangeUmInitialLeverageResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "leverage")
		delete(additionalProperties, "maxNotionalValue")
		delete(additionalProperties, "symbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChangeUmInitialLeverageResponse struct {
	value *ChangeUmInitialLeverageResponse
	isSet bool
}

func (v NullableChangeUmInitialLeverageResponse) Get() *ChangeUmInitialLeverageResponse {
	return v.value
}

func (v *NullableChangeUmInitialLeverageResponse) Set(val *ChangeUmInitialLeverageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeUmInitialLeverageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeUmInitialLeverageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeUmInitialLeverageResponse(val *ChangeUmInitialLeverageResponse) *NullableChangeUmInitialLeverageResponse {
	return &NullableChangeUmInitialLeverageResponse{value: val, isSet: true}
}

func (v NullableChangeUmInitialLeverageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeUmInitialLeverageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
