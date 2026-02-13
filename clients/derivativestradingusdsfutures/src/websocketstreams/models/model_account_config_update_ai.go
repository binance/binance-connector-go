/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountConfigUpdateAi type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountConfigUpdateAi{}

// AccountConfigUpdateAi struct for AccountConfigUpdateAi
type AccountConfigUpdateAi struct {
	Smallj               *bool `json:"j,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountConfigUpdateAi AccountConfigUpdateAi

// NewAccountConfigUpdateAi instantiates a new AccountConfigUpdateAi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountConfigUpdateAi() *AccountConfigUpdateAi {
	this := AccountConfigUpdateAi{}
	return &this
}

// NewAccountConfigUpdateAiWithDefaults instantiates a new AccountConfigUpdateAi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountConfigUpdateAiWithDefaults() *AccountConfigUpdateAi {
	this := AccountConfigUpdateAi{}
	return &this
}

// GetJ returns the J field value if set, zero value otherwise.
func (o *AccountConfigUpdateAi) GetSmallj() bool {
	if o == nil || common.IsNil(o.Smallj) {
		var ret bool
		return ret
	}
	return *o.Smallj
}

// GetJOk returns a tuple with the J field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountConfigUpdateAi) GetSmalljOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallj) {
		return nil, false
	}
	return o.Smallj, true
}

// HasJ returns a boolean if a field has been set.
func (o *AccountConfigUpdateAi) HasSmallj() bool {
	if o != nil && !common.IsNil(o.Smallj) {
		return true
	}

	return false
}

// SetJ gets a reference to the given bool and assigns it to the J field.
func (o *AccountConfigUpdateAi) SetSmallj(v bool) {
	o.Smallj = &v
}

func (o AccountConfigUpdateAi) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountConfigUpdateAi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smallj) {
		toSerialize["j"] = o.Smallj
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountConfigUpdateAi) UnmarshalJSON(data []byte) (err error) {
	varAccountConfigUpdateAi := _AccountConfigUpdateAi{}

	err = json.Unmarshal(data, &varAccountConfigUpdateAi)

	if err != nil {
		return err
	}

	*o = AccountConfigUpdateAi(varAccountConfigUpdateAi)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "j")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountConfigUpdateAi struct {
	value *AccountConfigUpdateAi
	isSet bool
}

func (v NullableAccountConfigUpdateAi) Get() *AccountConfigUpdateAi {
	return v.value
}

func (v *NullableAccountConfigUpdateAi) Set(val *AccountConfigUpdateAi) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountConfigUpdateAi) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountConfigUpdateAi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountConfigUpdateAi(val *AccountConfigUpdateAi) *NullableAccountConfigUpdateAi {
	return &NullableAccountConfigUpdateAi{value: val, isSet: true}
}

func (v NullableAccountConfigUpdateAi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountConfigUpdateAi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
