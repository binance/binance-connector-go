/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UserCommissionRateResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UserCommissionRateResponse{}

// UserCommissionRateResponse struct for UserCommissionRateResponse
type UserCommissionRateResponse struct {
	Symbol               *string `json:"symbol,omitempty"`
	MakerCommissionRate  *string `json:"makerCommissionRate,omitempty"`
	TakerCommissionRate  *string `json:"takerCommissionRate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserCommissionRateResponse UserCommissionRateResponse

// NewUserCommissionRateResponse instantiates a new UserCommissionRateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCommissionRateResponse() *UserCommissionRateResponse {
	this := UserCommissionRateResponse{}
	return &this
}

// NewUserCommissionRateResponseWithDefaults instantiates a new UserCommissionRateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCommissionRateResponseWithDefaults() *UserCommissionRateResponse {
	this := UserCommissionRateResponse{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UserCommissionRateResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCommissionRateResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UserCommissionRateResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UserCommissionRateResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetMakerCommissionRate returns the MakerCommissionRate field value if set, zero value otherwise.
func (o *UserCommissionRateResponse) GetMakerCommissionRate() string {
	if o == nil || common.IsNil(o.MakerCommissionRate) {
		var ret string
		return ret
	}
	return *o.MakerCommissionRate
}

// GetMakerCommissionRateOk returns a tuple with the MakerCommissionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCommissionRateResponse) GetMakerCommissionRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.MakerCommissionRate) {
		return nil, false
	}
	return o.MakerCommissionRate, true
}

// HasMakerCommissionRate returns a boolean if a field has been set.
func (o *UserCommissionRateResponse) HasMakerCommissionRate() bool {
	if o != nil && !common.IsNil(o.MakerCommissionRate) {
		return true
	}

	return false
}

// SetMakerCommissionRate gets a reference to the given string and assigns it to the MakerCommissionRate field.
func (o *UserCommissionRateResponse) SetMakerCommissionRate(v string) {
	o.MakerCommissionRate = &v
}

// GetTakerCommissionRate returns the TakerCommissionRate field value if set, zero value otherwise.
func (o *UserCommissionRateResponse) GetTakerCommissionRate() string {
	if o == nil || common.IsNil(o.TakerCommissionRate) {
		var ret string
		return ret
	}
	return *o.TakerCommissionRate
}

// GetTakerCommissionRateOk returns a tuple with the TakerCommissionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCommissionRateResponse) GetTakerCommissionRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerCommissionRate) {
		return nil, false
	}
	return o.TakerCommissionRate, true
}

// HasTakerCommissionRate returns a boolean if a field has been set.
func (o *UserCommissionRateResponse) HasTakerCommissionRate() bool {
	if o != nil && !common.IsNil(o.TakerCommissionRate) {
		return true
	}

	return false
}

// SetTakerCommissionRate gets a reference to the given string and assigns it to the TakerCommissionRate field.
func (o *UserCommissionRateResponse) SetTakerCommissionRate(v string) {
	o.TakerCommissionRate = &v
}

func (o UserCommissionRateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCommissionRateResponse) ToMap() (map[string]interface{}, error) {
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

func (o *UserCommissionRateResponse) UnmarshalJSON(data []byte) (err error) {
	varUserCommissionRateResponse := _UserCommissionRateResponse{}

	err = json.Unmarshal(data, &varUserCommissionRateResponse)

	if err != nil {
		return err
	}

	*o = UserCommissionRateResponse(varUserCommissionRateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "makerCommissionRate")
		delete(additionalProperties, "takerCommissionRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserCommissionRateResponse struct {
	value *UserCommissionRateResponse
	isSet bool
}

func (v NullableUserCommissionRateResponse) Get() *UserCommissionRateResponse {
	return v.value
}

func (v *NullableUserCommissionRateResponse) Set(val *UserCommissionRateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCommissionRateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCommissionRateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCommissionRateResponse(val *UserCommissionRateResponse) *NullableUserCommissionRateResponse {
	return &NullableUserCommissionRateResponse{value: val, isSet: true}
}

func (v NullableUserCommissionRateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCommissionRateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
