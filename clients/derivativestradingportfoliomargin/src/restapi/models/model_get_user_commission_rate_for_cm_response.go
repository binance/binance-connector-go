/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetUserCommissionRateForCmResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetUserCommissionRateForCmResponse{}

// GetUserCommissionRateForCmResponse struct for GetUserCommissionRateForCmResponse
type GetUserCommissionRateForCmResponse struct {
	Symbol               *string `json:"symbol,omitempty"`
	MakerCommissionRate  *string `json:"makerCommissionRate,omitempty"`
	TakerCommissionRate  *string `json:"takerCommissionRate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUserCommissionRateForCmResponse GetUserCommissionRateForCmResponse

// NewGetUserCommissionRateForCmResponse instantiates a new GetUserCommissionRateForCmResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserCommissionRateForCmResponse() *GetUserCommissionRateForCmResponse {
	this := GetUserCommissionRateForCmResponse{}
	return &this
}

// NewGetUserCommissionRateForCmResponseWithDefaults instantiates a new GetUserCommissionRateForCmResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserCommissionRateForCmResponseWithDefaults() *GetUserCommissionRateForCmResponse {
	this := GetUserCommissionRateForCmResponse{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetUserCommissionRateForCmResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserCommissionRateForCmResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetUserCommissionRateForCmResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetUserCommissionRateForCmResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetMakerCommissionRate returns the MakerCommissionRate field value if set, zero value otherwise.
func (o *GetUserCommissionRateForCmResponse) GetMakerCommissionRate() string {
	if o == nil || common.IsNil(o.MakerCommissionRate) {
		var ret string
		return ret
	}
	return *o.MakerCommissionRate
}

// GetMakerCommissionRateOk returns a tuple with the MakerCommissionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserCommissionRateForCmResponse) GetMakerCommissionRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.MakerCommissionRate) {
		return nil, false
	}
	return o.MakerCommissionRate, true
}

// HasMakerCommissionRate returns a boolean if a field has been set.
func (o *GetUserCommissionRateForCmResponse) HasMakerCommissionRate() bool {
	if o != nil && !common.IsNil(o.MakerCommissionRate) {
		return true
	}

	return false
}

// SetMakerCommissionRate gets a reference to the given string and assigns it to the MakerCommissionRate field.
func (o *GetUserCommissionRateForCmResponse) SetMakerCommissionRate(v string) {
	o.MakerCommissionRate = &v
}

// GetTakerCommissionRate returns the TakerCommissionRate field value if set, zero value otherwise.
func (o *GetUserCommissionRateForCmResponse) GetTakerCommissionRate() string {
	if o == nil || common.IsNil(o.TakerCommissionRate) {
		var ret string
		return ret
	}
	return *o.TakerCommissionRate
}

// GetTakerCommissionRateOk returns a tuple with the TakerCommissionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserCommissionRateForCmResponse) GetTakerCommissionRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerCommissionRate) {
		return nil, false
	}
	return o.TakerCommissionRate, true
}

// HasTakerCommissionRate returns a boolean if a field has been set.
func (o *GetUserCommissionRateForCmResponse) HasTakerCommissionRate() bool {
	if o != nil && !common.IsNil(o.TakerCommissionRate) {
		return true
	}

	return false
}

// SetTakerCommissionRate gets a reference to the given string and assigns it to the TakerCommissionRate field.
func (o *GetUserCommissionRateForCmResponse) SetTakerCommissionRate(v string) {
	o.TakerCommissionRate = &v
}

func (o GetUserCommissionRateForCmResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserCommissionRateForCmResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.MakerCommissionRate) {
		toSerialize["makerCommissionRate"] = o.MakerCommissionRate
	}
	if !common.IsNil(o.TakerCommissionRate) {
		toSerialize["takerCommissionRate"] = o.TakerCommissionRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUserCommissionRateForCmResponse) UnmarshalJSON(data []byte) (err error) {
	varGetUserCommissionRateForCmResponse := _GetUserCommissionRateForCmResponse{}

	err = json.Unmarshal(data, &varGetUserCommissionRateForCmResponse)

	if err != nil {
		return err
	}

	*o = GetUserCommissionRateForCmResponse(varGetUserCommissionRateForCmResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "makerCommissionRate")
		delete(additionalProperties, "takerCommissionRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUserCommissionRateForCmResponse struct {
	value *GetUserCommissionRateForCmResponse
	isSet bool
}

func (v NullableGetUserCommissionRateForCmResponse) Get() *GetUserCommissionRateForCmResponse {
	return v.value
}

func (v *NullableGetUserCommissionRateForCmResponse) Set(val *GetUserCommissionRateForCmResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserCommissionRateForCmResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserCommissionRateForCmResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserCommissionRateForCmResponse(val *GetUserCommissionRateForCmResponse) *NullableGetUserCommissionRateForCmResponse {
	return &NullableGetUserCommissionRateForCmResponse{value: val, isSet: true}
}

func (v NullableGetUserCommissionRateForCmResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserCommissionRateForCmResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
