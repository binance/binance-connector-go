/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FetchDepositAddressListWithNetworkResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FetchDepositAddressListWithNetworkResponseInner{}

// FetchDepositAddressListWithNetworkResponseInner struct for FetchDepositAddressListWithNetworkResponseInner
type FetchDepositAddressListWithNetworkResponseInner struct {
	Coin                 *string `json:"coin,omitempty"`
	Address              *string `json:"address,omitempty"`
	Tag                  *string `json:"tag,omitempty"`
	IsDefault            *int64  `json:"isDefault,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FetchDepositAddressListWithNetworkResponseInner FetchDepositAddressListWithNetworkResponseInner

// NewFetchDepositAddressListWithNetworkResponseInner instantiates a new FetchDepositAddressListWithNetworkResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchDepositAddressListWithNetworkResponseInner() *FetchDepositAddressListWithNetworkResponseInner {
	this := FetchDepositAddressListWithNetworkResponseInner{}
	return &this
}

// NewFetchDepositAddressListWithNetworkResponseInnerWithDefaults instantiates a new FetchDepositAddressListWithNetworkResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchDepositAddressListWithNetworkResponseInnerWithDefaults() *FetchDepositAddressListWithNetworkResponseInner {
	this := FetchDepositAddressListWithNetworkResponseInner{}
	return &this
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *FetchDepositAddressListWithNetworkResponseInner) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchDepositAddressListWithNetworkResponseInner) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *FetchDepositAddressListWithNetworkResponseInner) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *FetchDepositAddressListWithNetworkResponseInner) SetCoin(v string) {
	o.Coin = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *FetchDepositAddressListWithNetworkResponseInner) GetAddress() string {
	if o == nil || common.IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchDepositAddressListWithNetworkResponseInner) GetAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *FetchDepositAddressListWithNetworkResponseInner) HasAddress() bool {
	if o != nil && !common.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *FetchDepositAddressListWithNetworkResponseInner) SetAddress(v string) {
	o.Address = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *FetchDepositAddressListWithNetworkResponseInner) GetTag() string {
	if o == nil || common.IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchDepositAddressListWithNetworkResponseInner) GetTagOk() (*string, bool) {
	if o == nil || common.IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *FetchDepositAddressListWithNetworkResponseInner) HasTag() bool {
	if o != nil && !common.IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *FetchDepositAddressListWithNetworkResponseInner) SetTag(v string) {
	o.Tag = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *FetchDepositAddressListWithNetworkResponseInner) GetIsDefault() int64 {
	if o == nil || common.IsNil(o.IsDefault) {
		var ret int64
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchDepositAddressListWithNetworkResponseInner) GetIsDefaultOk() (*int64, bool) {
	if o == nil || common.IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *FetchDepositAddressListWithNetworkResponseInner) HasIsDefault() bool {
	if o != nil && !common.IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given int64 and assigns it to the IsDefault field.
func (o *FetchDepositAddressListWithNetworkResponseInner) SetIsDefault(v int64) {
	o.IsDefault = &v
}

func (o FetchDepositAddressListWithNetworkResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchDepositAddressListWithNetworkResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !common.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !common.IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !common.IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FetchDepositAddressListWithNetworkResponseInner) UnmarshalJSON(data []byte) (err error) {
	varFetchDepositAddressListWithNetworkResponseInner := _FetchDepositAddressListWithNetworkResponseInner{}

	err = json.Unmarshal(data, &varFetchDepositAddressListWithNetworkResponseInner)

	if err != nil {
		return err
	}

	*o = FetchDepositAddressListWithNetworkResponseInner(varFetchDepositAddressListWithNetworkResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "coin")
		delete(additionalProperties, "address")
		delete(additionalProperties, "tag")
		delete(additionalProperties, "isDefault")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFetchDepositAddressListWithNetworkResponseInner struct {
	value *FetchDepositAddressListWithNetworkResponseInner
	isSet bool
}

func (v NullableFetchDepositAddressListWithNetworkResponseInner) Get() *FetchDepositAddressListWithNetworkResponseInner {
	return v.value
}

func (v *NullableFetchDepositAddressListWithNetworkResponseInner) Set(val *FetchDepositAddressListWithNetworkResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchDepositAddressListWithNetworkResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchDepositAddressListWithNetworkResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchDepositAddressListWithNetworkResponseInner(val *FetchDepositAddressListWithNetworkResponseInner) *NullableFetchDepositAddressListWithNetworkResponseInner {
	return &NullableFetchDepositAddressListWithNetworkResponseInner{value: val, isSet: true}
}

func (v NullableFetchDepositAddressListWithNetworkResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchDepositAddressListWithNetworkResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
