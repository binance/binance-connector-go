/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the HashrateResaleDetailResponseDataProfitTransferDetailsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &HashrateResaleDetailResponseDataProfitTransferDetailsInner{}

// HashrateResaleDetailResponseDataProfitTransferDetailsInner struct for HashrateResaleDetailResponseDataProfitTransferDetailsInner
type HashrateResaleDetailResponseDataProfitTransferDetailsInner struct {
	PoolUsername         *string  `json:"poolUsername,omitempty"`
	ToPoolUsername       *string  `json:"toPoolUsername,omitempty"`
	AlgoName             *string  `json:"algoName,omitempty"`
	HashRate             *int64   `json:"hashRate,omitempty"`
	Day                  *int64   `json:"day,omitempty"`
	Amount               *float32 `json:"amount,omitempty"`
	CoinName             *string  `json:"coinName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HashrateResaleDetailResponseDataProfitTransferDetailsInner HashrateResaleDetailResponseDataProfitTransferDetailsInner

// NewHashrateResaleDetailResponseDataProfitTransferDetailsInner instantiates a new HashrateResaleDetailResponseDataProfitTransferDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHashrateResaleDetailResponseDataProfitTransferDetailsInner() *HashrateResaleDetailResponseDataProfitTransferDetailsInner {
	this := HashrateResaleDetailResponseDataProfitTransferDetailsInner{}
	return &this
}

// NewHashrateResaleDetailResponseDataProfitTransferDetailsInnerWithDefaults instantiates a new HashrateResaleDetailResponseDataProfitTransferDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHashrateResaleDetailResponseDataProfitTransferDetailsInnerWithDefaults() *HashrateResaleDetailResponseDataProfitTransferDetailsInner {
	this := HashrateResaleDetailResponseDataProfitTransferDetailsInner{}
	return &this
}

// GetPoolUsername returns the PoolUsername field value if set, zero value otherwise.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) GetPoolUsername() string {
	if o == nil || common.IsNil(o.PoolUsername) {
		var ret string
		return ret
	}
	return *o.PoolUsername
}

// GetPoolUsernameOk returns a tuple with the PoolUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) GetPoolUsernameOk() (*string, bool) {
	if o == nil || common.IsNil(o.PoolUsername) {
		return nil, false
	}
	return o.PoolUsername, true
}

// HasPoolUsername returns a boolean if a field has been set.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) HasPoolUsername() bool {
	if o != nil && !common.IsNil(o.PoolUsername) {
		return true
	}

	return false
}

// SetPoolUsername gets a reference to the given string and assigns it to the PoolUsername field.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) SetPoolUsername(v string) {
	o.PoolUsername = &v
}

// GetToPoolUsername returns the ToPoolUsername field value if set, zero value otherwise.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) GetToPoolUsername() string {
	if o == nil || common.IsNil(o.ToPoolUsername) {
		var ret string
		return ret
	}
	return *o.ToPoolUsername
}

// GetToPoolUsernameOk returns a tuple with the ToPoolUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) GetToPoolUsernameOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToPoolUsername) {
		return nil, false
	}
	return o.ToPoolUsername, true
}

// HasToPoolUsername returns a boolean if a field has been set.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) HasToPoolUsername() bool {
	if o != nil && !common.IsNil(o.ToPoolUsername) {
		return true
	}

	return false
}

// SetToPoolUsername gets a reference to the given string and assigns it to the ToPoolUsername field.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) SetToPoolUsername(v string) {
	o.ToPoolUsername = &v
}

// GetAlgoName returns the AlgoName field value if set, zero value otherwise.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) GetAlgoName() string {
	if o == nil || common.IsNil(o.AlgoName) {
		var ret string
		return ret
	}
	return *o.AlgoName
}

// GetAlgoNameOk returns a tuple with the AlgoName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) GetAlgoNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.AlgoName) {
		return nil, false
	}
	return o.AlgoName, true
}

// HasAlgoName returns a boolean if a field has been set.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) HasAlgoName() bool {
	if o != nil && !common.IsNil(o.AlgoName) {
		return true
	}

	return false
}

// SetAlgoName gets a reference to the given string and assigns it to the AlgoName field.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) SetAlgoName(v string) {
	o.AlgoName = &v
}

// GetHashRate returns the HashRate field value if set, zero value otherwise.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) GetHashRate() int64 {
	if o == nil || common.IsNil(o.HashRate) {
		var ret int64
		return ret
	}
	return *o.HashRate
}

// GetHashRateOk returns a tuple with the HashRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) GetHashRateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.HashRate) {
		return nil, false
	}
	return o.HashRate, true
}

// HasHashRate returns a boolean if a field has been set.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) HasHashRate() bool {
	if o != nil && !common.IsNil(o.HashRate) {
		return true
	}

	return false
}

// SetHashRate gets a reference to the given int64 and assigns it to the HashRate field.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) SetHashRate(v int64) {
	o.HashRate = &v
}

// GetDay returns the Day field value if set, zero value otherwise.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) GetDay() int64 {
	if o == nil || common.IsNil(o.Day) {
		var ret int64
		return ret
	}
	return *o.Day
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) GetDayOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Day) {
		return nil, false
	}
	return o.Day, true
}

// HasDay returns a boolean if a field has been set.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) HasDay() bool {
	if o != nil && !common.IsNil(o.Day) {
		return true
	}

	return false
}

// SetDay gets a reference to the given int64 and assigns it to the Day field.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) SetDay(v int64) {
	o.Day = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) GetAmount() float32 {
	if o == nil || common.IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) GetAmountOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) SetAmount(v float32) {
	o.Amount = &v
}

// GetCoinName returns the CoinName field value if set, zero value otherwise.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) GetCoinName() string {
	if o == nil || common.IsNil(o.CoinName) {
		var ret string
		return ret
	}
	return *o.CoinName
}

// GetCoinNameOk returns a tuple with the CoinName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) GetCoinNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.CoinName) {
		return nil, false
	}
	return o.CoinName, true
}

// HasCoinName returns a boolean if a field has been set.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) HasCoinName() bool {
	if o != nil && !common.IsNil(o.CoinName) {
		return true
	}

	return false
}

// SetCoinName gets a reference to the given string and assigns it to the CoinName field.
func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) SetCoinName(v string) {
	o.CoinName = &v
}

func (o HashrateResaleDetailResponseDataProfitTransferDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HashrateResaleDetailResponseDataProfitTransferDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PoolUsername) {
		toSerialize["poolUsername"] = o.PoolUsername
	}
	if !common.IsNil(o.ToPoolUsername) {
		toSerialize["toPoolUsername"] = o.ToPoolUsername
	}
	if !common.IsNil(o.AlgoName) {
		toSerialize["algoName"] = o.AlgoName
	}
	if !common.IsNil(o.HashRate) {
		toSerialize["hashRate"] = o.HashRate
	}
	if !common.IsNil(o.Day) {
		toSerialize["day"] = o.Day
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.CoinName) {
		toSerialize["coinName"] = o.CoinName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HashrateResaleDetailResponseDataProfitTransferDetailsInner) UnmarshalJSON(data []byte) (err error) {
	varHashrateResaleDetailResponseDataProfitTransferDetailsInner := _HashrateResaleDetailResponseDataProfitTransferDetailsInner{}

	err = json.Unmarshal(data, &varHashrateResaleDetailResponseDataProfitTransferDetailsInner)

	if err != nil {
		return err
	}

	*o = HashrateResaleDetailResponseDataProfitTransferDetailsInner(varHashrateResaleDetailResponseDataProfitTransferDetailsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "poolUsername")
		delete(additionalProperties, "toPoolUsername")
		delete(additionalProperties, "algoName")
		delete(additionalProperties, "hashRate")
		delete(additionalProperties, "day")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "coinName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHashrateResaleDetailResponseDataProfitTransferDetailsInner struct {
	value *HashrateResaleDetailResponseDataProfitTransferDetailsInner
	isSet bool
}

func (v NullableHashrateResaleDetailResponseDataProfitTransferDetailsInner) Get() *HashrateResaleDetailResponseDataProfitTransferDetailsInner {
	return v.value
}

func (v *NullableHashrateResaleDetailResponseDataProfitTransferDetailsInner) Set(val *HashrateResaleDetailResponseDataProfitTransferDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableHashrateResaleDetailResponseDataProfitTransferDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableHashrateResaleDetailResponseDataProfitTransferDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHashrateResaleDetailResponseDataProfitTransferDetailsInner(val *HashrateResaleDetailResponseDataProfitTransferDetailsInner) *NullableHashrateResaleDetailResponseDataProfitTransferDetailsInner {
	return &NullableHashrateResaleDetailResponseDataProfitTransferDetailsInner{value: val, isSet: true}
}

func (v NullableHashrateResaleDetailResponseDataProfitTransferDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHashrateResaleDetailResponseDataProfitTransferDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
