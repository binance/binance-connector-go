/*
Binance Gift Card REST API

OpenAPI Specification for the Binance Gift Card REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the FetchTokenLimitResponseDataInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FetchTokenLimitResponseDataInner{}

// FetchTokenLimitResponseDataInner struct for FetchTokenLimitResponseDataInner
type FetchTokenLimitResponseDataInner struct {
	Coin                 *string `json:"coin,omitempty"`
	FromMin              *string `json:"fromMin,omitempty"`
	FromMax              *string `json:"fromMax,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FetchTokenLimitResponseDataInner FetchTokenLimitResponseDataInner

// NewFetchTokenLimitResponseDataInner instantiates a new FetchTokenLimitResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchTokenLimitResponseDataInner() *FetchTokenLimitResponseDataInner {
	this := FetchTokenLimitResponseDataInner{}
	return &this
}

// NewFetchTokenLimitResponseDataInnerWithDefaults instantiates a new FetchTokenLimitResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchTokenLimitResponseDataInnerWithDefaults() *FetchTokenLimitResponseDataInner {
	this := FetchTokenLimitResponseDataInner{}
	return &this
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *FetchTokenLimitResponseDataInner) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchTokenLimitResponseDataInner) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *FetchTokenLimitResponseDataInner) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *FetchTokenLimitResponseDataInner) SetCoin(v string) {
	o.Coin = &v
}

// GetFromMin returns the FromMin field value if set, zero value otherwise.
func (o *FetchTokenLimitResponseDataInner) GetFromMin() string {
	if o == nil || common.IsNil(o.FromMin) {
		var ret string
		return ret
	}
	return *o.FromMin
}

// GetFromMinOk returns a tuple with the FromMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchTokenLimitResponseDataInner) GetFromMinOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromMin) {
		return nil, false
	}
	return o.FromMin, true
}

// HasFromMin returns a boolean if a field has been set.
func (o *FetchTokenLimitResponseDataInner) HasFromMin() bool {
	if o != nil && !common.IsNil(o.FromMin) {
		return true
	}

	return false
}

// SetFromMin gets a reference to the given string and assigns it to the FromMin field.
func (o *FetchTokenLimitResponseDataInner) SetFromMin(v string) {
	o.FromMin = &v
}

// GetFromMax returns the FromMax field value if set, zero value otherwise.
func (o *FetchTokenLimitResponseDataInner) GetFromMax() string {
	if o == nil || common.IsNil(o.FromMax) {
		var ret string
		return ret
	}
	return *o.FromMax
}

// GetFromMaxOk returns a tuple with the FromMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchTokenLimitResponseDataInner) GetFromMaxOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromMax) {
		return nil, false
	}
	return o.FromMax, true
}

// HasFromMax returns a boolean if a field has been set.
func (o *FetchTokenLimitResponseDataInner) HasFromMax() bool {
	if o != nil && !common.IsNil(o.FromMax) {
		return true
	}

	return false
}

// SetFromMax gets a reference to the given string and assigns it to the FromMax field.
func (o *FetchTokenLimitResponseDataInner) SetFromMax(v string) {
	o.FromMax = &v
}

func (o FetchTokenLimitResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchTokenLimitResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !common.IsNil(o.FromMin) {
		toSerialize["fromMin"] = o.FromMin
	}
	if !common.IsNil(o.FromMax) {
		toSerialize["fromMax"] = o.FromMax
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FetchTokenLimitResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	varFetchTokenLimitResponseDataInner := _FetchTokenLimitResponseDataInner{}

	err = json.Unmarshal(data, &varFetchTokenLimitResponseDataInner)

	if err != nil {
		return err
	}

	*o = FetchTokenLimitResponseDataInner(varFetchTokenLimitResponseDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "coin")
		delete(additionalProperties, "fromMin")
		delete(additionalProperties, "fromMax")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFetchTokenLimitResponseDataInner struct {
	value *FetchTokenLimitResponseDataInner
	isSet bool
}

func (v NullableFetchTokenLimitResponseDataInner) Get() *FetchTokenLimitResponseDataInner {
	return v.value
}

func (v *NullableFetchTokenLimitResponseDataInner) Set(val *FetchTokenLimitResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchTokenLimitResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchTokenLimitResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchTokenLimitResponseDataInner(val *FetchTokenLimitResponseDataInner) *NullableFetchTokenLimitResponseDataInner {
	return &NullableFetchTokenLimitResponseDataInner{value: val, isSet: true}
}

func (v NullableFetchTokenLimitResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchTokenLimitResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
