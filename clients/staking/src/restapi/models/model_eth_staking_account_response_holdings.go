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

// checks if the EthStakingAccountResponseHoldings type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EthStakingAccountResponseHoldings{}

// EthStakingAccountResponseHoldings struct for EthStakingAccountResponseHoldings
type EthStakingAccountResponseHoldings struct {
	WbethAmount          *string `json:"wbethAmount,omitempty"`
	BethAmount           *string `json:"bethAmount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EthStakingAccountResponseHoldings EthStakingAccountResponseHoldings

// NewEthStakingAccountResponseHoldings instantiates a new EthStakingAccountResponseHoldings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthStakingAccountResponseHoldings() *EthStakingAccountResponseHoldings {
	this := EthStakingAccountResponseHoldings{}
	return &this
}

// NewEthStakingAccountResponseHoldingsWithDefaults instantiates a new EthStakingAccountResponseHoldings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthStakingAccountResponseHoldingsWithDefaults() *EthStakingAccountResponseHoldings {
	this := EthStakingAccountResponseHoldings{}
	return &this
}

// GetWbethAmount returns the WbethAmount field value if set, zero value otherwise.
func (o *EthStakingAccountResponseHoldings) GetWbethAmount() string {
	if o == nil || common.IsNil(o.WbethAmount) {
		var ret string
		return ret
	}
	return *o.WbethAmount
}

// GetWbethAmountOk returns a tuple with the WbethAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthStakingAccountResponseHoldings) GetWbethAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.WbethAmount) {
		return nil, false
	}
	return o.WbethAmount, true
}

// HasWbethAmount returns a boolean if a field has been set.
func (o *EthStakingAccountResponseHoldings) HasWbethAmount() bool {
	if o != nil && !common.IsNil(o.WbethAmount) {
		return true
	}

	return false
}

// SetWbethAmount gets a reference to the given string and assigns it to the WbethAmount field.
func (o *EthStakingAccountResponseHoldings) SetWbethAmount(v string) {
	o.WbethAmount = &v
}

// GetBethAmount returns the BethAmount field value if set, zero value otherwise.
func (o *EthStakingAccountResponseHoldings) GetBethAmount() string {
	if o == nil || common.IsNil(o.BethAmount) {
		var ret string
		return ret
	}
	return *o.BethAmount
}

// GetBethAmountOk returns a tuple with the BethAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthStakingAccountResponseHoldings) GetBethAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.BethAmount) {
		return nil, false
	}
	return o.BethAmount, true
}

// HasBethAmount returns a boolean if a field has been set.
func (o *EthStakingAccountResponseHoldings) HasBethAmount() bool {
	if o != nil && !common.IsNil(o.BethAmount) {
		return true
	}

	return false
}

// SetBethAmount gets a reference to the given string and assigns it to the BethAmount field.
func (o *EthStakingAccountResponseHoldings) SetBethAmount(v string) {
	o.BethAmount = &v
}

func (o EthStakingAccountResponseHoldings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EthStakingAccountResponseHoldings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.WbethAmount) {
		toSerialize["wbethAmount"] = o.WbethAmount
	}
	if !common.IsNil(o.BethAmount) {
		toSerialize["bethAmount"] = o.BethAmount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EthStakingAccountResponseHoldings) UnmarshalJSON(data []byte) (err error) {
	varEthStakingAccountResponseHoldings := _EthStakingAccountResponseHoldings{}

	err = json.Unmarshal(data, &varEthStakingAccountResponseHoldings)

	if err != nil {
		return err
	}

	*o = EthStakingAccountResponseHoldings(varEthStakingAccountResponseHoldings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "wbethAmount")
		delete(additionalProperties, "bethAmount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEthStakingAccountResponseHoldings struct {
	value *EthStakingAccountResponseHoldings
	isSet bool
}

func (v NullableEthStakingAccountResponseHoldings) Get() *EthStakingAccountResponseHoldings {
	return v.value
}

func (v *NullableEthStakingAccountResponseHoldings) Set(val *EthStakingAccountResponseHoldings) {
	v.value = val
	v.isSet = true
}

func (v NullableEthStakingAccountResponseHoldings) IsSet() bool {
	return v.isSet
}

func (v *NullableEthStakingAccountResponseHoldings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthStakingAccountResponseHoldings(val *EthStakingAccountResponseHoldings) *NullableEthStakingAccountResponseHoldings {
	return &NullableEthStakingAccountResponseHoldings{value: val, isSet: true}
}

func (v NullableEthStakingAccountResponseHoldings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthStakingAccountResponseHoldings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
