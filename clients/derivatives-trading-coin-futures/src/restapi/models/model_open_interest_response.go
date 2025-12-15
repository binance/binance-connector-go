/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OpenInterestResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OpenInterestResponse{}

// OpenInterestResponse struct for OpenInterestResponse
type OpenInterestResponse struct {
	Symbol               *string `json:"symbol,omitempty"`
	Pair                 *string `json:"pair,omitempty"`
	OpenInterest         *string `json:"openInterest,omitempty"`
	ContractType         *string `json:"contractType,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenInterestResponse OpenInterestResponse

// NewOpenInterestResponse instantiates a new OpenInterestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenInterestResponse() *OpenInterestResponse {
	this := OpenInterestResponse{}
	return &this
}

// NewOpenInterestResponseWithDefaults instantiates a new OpenInterestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenInterestResponseWithDefaults() *OpenInterestResponse {
	this := OpenInterestResponse{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OpenInterestResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OpenInterestResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OpenInterestResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *OpenInterestResponse) GetPair() string {
	if o == nil || common.IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestResponse) GetPairOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *OpenInterestResponse) HasPair() bool {
	if o != nil && !common.IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *OpenInterestResponse) SetPair(v string) {
	o.Pair = &v
}

// GetOpenInterest returns the OpenInterest field value if set, zero value otherwise.
func (o *OpenInterestResponse) GetOpenInterest() string {
	if o == nil || common.IsNil(o.OpenInterest) {
		var ret string
		return ret
	}
	return *o.OpenInterest
}

// GetOpenInterestOk returns a tuple with the OpenInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestResponse) GetOpenInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpenInterest) {
		return nil, false
	}
	return o.OpenInterest, true
}

// HasOpenInterest returns a boolean if a field has been set.
func (o *OpenInterestResponse) HasOpenInterest() bool {
	if o != nil && !common.IsNil(o.OpenInterest) {
		return true
	}

	return false
}

// SetOpenInterest gets a reference to the given string and assigns it to the OpenInterest field.
func (o *OpenInterestResponse) SetOpenInterest(v string) {
	o.OpenInterest = &v
}

// GetContractType returns the ContractType field value if set, zero value otherwise.
func (o *OpenInterestResponse) GetContractType() string {
	if o == nil || common.IsNil(o.ContractType) {
		var ret string
		return ret
	}
	return *o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestResponse) GetContractTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContractType) {
		return nil, false
	}
	return o.ContractType, true
}

// HasContractType returns a boolean if a field has been set.
func (o *OpenInterestResponse) HasContractType() bool {
	if o != nil && !common.IsNil(o.ContractType) {
		return true
	}

	return false
}

// SetContractType gets a reference to the given string and assigns it to the ContractType field.
func (o *OpenInterestResponse) SetContractType(v string) {
	o.ContractType = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *OpenInterestResponse) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestResponse) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *OpenInterestResponse) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *OpenInterestResponse) SetTime(v int64) {
	o.Time = &v
}

func (o OpenInterestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenInterestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !common.IsNil(o.OpenInterest) {
		toSerialize["openInterest"] = o.OpenInterest
	}
	if !common.IsNil(o.ContractType) {
		toSerialize["contractType"] = o.ContractType
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenInterestResponse) UnmarshalJSON(data []byte) (err error) {
	varOpenInterestResponse := _OpenInterestResponse{}

	err = json.Unmarshal(data, &varOpenInterestResponse)

	if err != nil {
		return err
	}

	*o = OpenInterestResponse(varOpenInterestResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "pair")
		delete(additionalProperties, "openInterest")
		delete(additionalProperties, "contractType")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenInterestResponse struct {
	value *OpenInterestResponse
	isSet bool
}

func (v NullableOpenInterestResponse) Get() *OpenInterestResponse {
	return v.value
}

func (v *NullableOpenInterestResponse) Set(val *OpenInterestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenInterestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenInterestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenInterestResponse(val *OpenInterestResponse) *NullableOpenInterestResponse {
	return &NullableOpenInterestResponse{value: val, isSet: true}
}

func (v NullableOpenInterestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenInterestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
