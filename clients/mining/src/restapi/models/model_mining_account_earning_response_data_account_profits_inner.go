/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MiningAccountEarningResponseDataAccountProfitsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MiningAccountEarningResponseDataAccountProfitsInner{}

// MiningAccountEarningResponseDataAccountProfitsInner struct for MiningAccountEarningResponseDataAccountProfitsInner
type MiningAccountEarningResponseDataAccountProfitsInner struct {
	Time                 *int64   `json:"time,omitempty"`
	CoinName             *string  `json:"coinName,omitempty"`
	Type                 *int64   `json:"type,omitempty"`
	Puid                 *int64   `json:"puid,omitempty"`
	SubName              *string  `json:"subName,omitempty"`
	Amount               *float32 `json:"amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MiningAccountEarningResponseDataAccountProfitsInner MiningAccountEarningResponseDataAccountProfitsInner

// NewMiningAccountEarningResponseDataAccountProfitsInner instantiates a new MiningAccountEarningResponseDataAccountProfitsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiningAccountEarningResponseDataAccountProfitsInner() *MiningAccountEarningResponseDataAccountProfitsInner {
	this := MiningAccountEarningResponseDataAccountProfitsInner{}
	return &this
}

// NewMiningAccountEarningResponseDataAccountProfitsInnerWithDefaults instantiates a new MiningAccountEarningResponseDataAccountProfitsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiningAccountEarningResponseDataAccountProfitsInnerWithDefaults() *MiningAccountEarningResponseDataAccountProfitsInner {
	this := MiningAccountEarningResponseDataAccountProfitsInner{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) SetTime(v int64) {
	o.Time = &v
}

// GetCoinName returns the CoinName field value if set, zero value otherwise.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) GetCoinName() string {
	if o == nil || common.IsNil(o.CoinName) {
		var ret string
		return ret
	}
	return *o.CoinName
}

// GetCoinNameOk returns a tuple with the CoinName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) GetCoinNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.CoinName) {
		return nil, false
	}
	return o.CoinName, true
}

// HasCoinName returns a boolean if a field has been set.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) HasCoinName() bool {
	if o != nil && !common.IsNil(o.CoinName) {
		return true
	}

	return false
}

// SetCoinName gets a reference to the given string and assigns it to the CoinName field.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) SetCoinName(v string) {
	o.CoinName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) GetType() int64 {
	if o == nil || common.IsNil(o.Type) {
		var ret int64
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) GetTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int64 and assigns it to the Type field.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) SetType(v int64) {
	o.Type = &v
}

// GetPuid returns the Puid field value if set, zero value otherwise.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) GetPuid() int64 {
	if o == nil || common.IsNil(o.Puid) {
		var ret int64
		return ret
	}
	return *o.Puid
}

// GetPuidOk returns a tuple with the Puid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) GetPuidOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Puid) {
		return nil, false
	}
	return o.Puid, true
}

// HasPuid returns a boolean if a field has been set.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) HasPuid() bool {
	if o != nil && !common.IsNil(o.Puid) {
		return true
	}

	return false
}

// SetPuid gets a reference to the given int64 and assigns it to the Puid field.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) SetPuid(v int64) {
	o.Puid = &v
}

// GetSubName returns the SubName field value if set, zero value otherwise.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) GetSubName() string {
	if o == nil || common.IsNil(o.SubName) {
		var ret string
		return ret
	}
	return *o.SubName
}

// GetSubNameOk returns a tuple with the SubName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) GetSubNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.SubName) {
		return nil, false
	}
	return o.SubName, true
}

// HasSubName returns a boolean if a field has been set.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) HasSubName() bool {
	if o != nil && !common.IsNil(o.SubName) {
		return true
	}

	return false
}

// SetSubName gets a reference to the given string and assigns it to the SubName field.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) SetSubName(v string) {
	o.SubName = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) GetAmount() float32 {
	if o == nil || common.IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) GetAmountOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *MiningAccountEarningResponseDataAccountProfitsInner) SetAmount(v float32) {
	o.Amount = &v
}

func (o MiningAccountEarningResponseDataAccountProfitsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiningAccountEarningResponseDataAccountProfitsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.CoinName) {
		toSerialize["coinName"] = o.CoinName
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.Puid) {
		toSerialize["puid"] = o.Puid
	}
	if !common.IsNil(o.SubName) {
		toSerialize["subName"] = o.SubName
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MiningAccountEarningResponseDataAccountProfitsInner) UnmarshalJSON(data []byte) (err error) {
	varMiningAccountEarningResponseDataAccountProfitsInner := _MiningAccountEarningResponseDataAccountProfitsInner{}

	err = json.Unmarshal(data, &varMiningAccountEarningResponseDataAccountProfitsInner)

	if err != nil {
		return err
	}

	*o = MiningAccountEarningResponseDataAccountProfitsInner(varMiningAccountEarningResponseDataAccountProfitsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "coinName")
		delete(additionalProperties, "type")
		delete(additionalProperties, "puid")
		delete(additionalProperties, "subName")
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMiningAccountEarningResponseDataAccountProfitsInner struct {
	value *MiningAccountEarningResponseDataAccountProfitsInner
	isSet bool
}

func (v NullableMiningAccountEarningResponseDataAccountProfitsInner) Get() *MiningAccountEarningResponseDataAccountProfitsInner {
	return v.value
}

func (v *NullableMiningAccountEarningResponseDataAccountProfitsInner) Set(val *MiningAccountEarningResponseDataAccountProfitsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMiningAccountEarningResponseDataAccountProfitsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMiningAccountEarningResponseDataAccountProfitsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiningAccountEarningResponseDataAccountProfitsInner(val *MiningAccountEarningResponseDataAccountProfitsInner) *NullableMiningAccountEarningResponseDataAccountProfitsInner {
	return &NullableMiningAccountEarningResponseDataAccountProfitsInner{value: val, isSet: true}
}

func (v NullableMiningAccountEarningResponseDataAccountProfitsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiningAccountEarningResponseDataAccountProfitsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
