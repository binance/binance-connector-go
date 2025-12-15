/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the EthStakingAccountResponseProfit type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EthStakingAccountResponseProfit{}

// EthStakingAccountResponseProfit struct for EthStakingAccountResponseProfit
type EthStakingAccountResponseProfit struct {
	AmountFromWBETH      *string `json:"amountFromWBETH,omitempty"`
	AmountFromBETH       *string `json:"amountFromBETH,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EthStakingAccountResponseProfit EthStakingAccountResponseProfit

// NewEthStakingAccountResponseProfit instantiates a new EthStakingAccountResponseProfit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthStakingAccountResponseProfit() *EthStakingAccountResponseProfit {
	this := EthStakingAccountResponseProfit{}
	return &this
}

// NewEthStakingAccountResponseProfitWithDefaults instantiates a new EthStakingAccountResponseProfit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthStakingAccountResponseProfitWithDefaults() *EthStakingAccountResponseProfit {
	this := EthStakingAccountResponseProfit{}
	return &this
}

// GetAmountFromWBETH returns the AmountFromWBETH field value if set, zero value otherwise.
func (o *EthStakingAccountResponseProfit) GetAmountFromWBETH() string {
	if o == nil || common.IsNil(o.AmountFromWBETH) {
		var ret string
		return ret
	}
	return *o.AmountFromWBETH
}

// GetAmountFromWBETHOk returns a tuple with the AmountFromWBETH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthStakingAccountResponseProfit) GetAmountFromWBETHOk() (*string, bool) {
	if o == nil || common.IsNil(o.AmountFromWBETH) {
		return nil, false
	}
	return o.AmountFromWBETH, true
}

// HasAmountFromWBETH returns a boolean if a field has been set.
func (o *EthStakingAccountResponseProfit) HasAmountFromWBETH() bool {
	if o != nil && !common.IsNil(o.AmountFromWBETH) {
		return true
	}

	return false
}

// SetAmountFromWBETH gets a reference to the given string and assigns it to the AmountFromWBETH field.
func (o *EthStakingAccountResponseProfit) SetAmountFromWBETH(v string) {
	o.AmountFromWBETH = &v
}

// GetAmountFromBETH returns the AmountFromBETH field value if set, zero value otherwise.
func (o *EthStakingAccountResponseProfit) GetAmountFromBETH() string {
	if o == nil || common.IsNil(o.AmountFromBETH) {
		var ret string
		return ret
	}
	return *o.AmountFromBETH
}

// GetAmountFromBETHOk returns a tuple with the AmountFromBETH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthStakingAccountResponseProfit) GetAmountFromBETHOk() (*string, bool) {
	if o == nil || common.IsNil(o.AmountFromBETH) {
		return nil, false
	}
	return o.AmountFromBETH, true
}

// HasAmountFromBETH returns a boolean if a field has been set.
func (o *EthStakingAccountResponseProfit) HasAmountFromBETH() bool {
	if o != nil && !common.IsNil(o.AmountFromBETH) {
		return true
	}

	return false
}

// SetAmountFromBETH gets a reference to the given string and assigns it to the AmountFromBETH field.
func (o *EthStakingAccountResponseProfit) SetAmountFromBETH(v string) {
	o.AmountFromBETH = &v
}

func (o EthStakingAccountResponseProfit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EthStakingAccountResponseProfit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AmountFromWBETH) {
		toSerialize["amountFromWBETH"] = o.AmountFromWBETH
	}
	if !common.IsNil(o.AmountFromBETH) {
		toSerialize["amountFromBETH"] = o.AmountFromBETH
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EthStakingAccountResponseProfit) UnmarshalJSON(data []byte) (err error) {
	varEthStakingAccountResponseProfit := _EthStakingAccountResponseProfit{}

	err = json.Unmarshal(data, &varEthStakingAccountResponseProfit)

	if err != nil {
		return err
	}

	*o = EthStakingAccountResponseProfit(varEthStakingAccountResponseProfit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amountFromWBETH")
		delete(additionalProperties, "amountFromBETH")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEthStakingAccountResponseProfit struct {
	value *EthStakingAccountResponseProfit
	isSet bool
}

func (v NullableEthStakingAccountResponseProfit) Get() *EthStakingAccountResponseProfit {
	return v.value
}

func (v *NullableEthStakingAccountResponseProfit) Set(val *EthStakingAccountResponseProfit) {
	v.value = val
	v.isSet = true
}

func (v NullableEthStakingAccountResponseProfit) IsSet() bool {
	return v.isSet
}

func (v *NullableEthStakingAccountResponseProfit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthStakingAccountResponseProfit(val *EthStakingAccountResponseProfit) *NullableEthStakingAccountResponseProfit {
	return &NullableEthStakingAccountResponseProfit{value: val, isSet: true}
}

func (v NullableEthStakingAccountResponseProfit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthStakingAccountResponseProfit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
