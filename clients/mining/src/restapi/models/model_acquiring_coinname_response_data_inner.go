/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AcquiringCoinnameResponseDataInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AcquiringCoinnameResponseDataInner{}

// AcquiringCoinnameResponseDataInner struct for AcquiringCoinnameResponseDataInner
type AcquiringCoinnameResponseDataInner struct {
	CoinName             *string `json:"coinName,omitempty"`
	CoinId               *int64  `json:"coinId,omitempty"`
	PoolIndex            *int64  `json:"poolIndex,omitempty"`
	AlgoId               *int64  `json:"algoId,omitempty"`
	AlgoName             *string `json:"algoName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AcquiringCoinnameResponseDataInner AcquiringCoinnameResponseDataInner

// NewAcquiringCoinnameResponseDataInner instantiates a new AcquiringCoinnameResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcquiringCoinnameResponseDataInner() *AcquiringCoinnameResponseDataInner {
	this := AcquiringCoinnameResponseDataInner{}
	return &this
}

// NewAcquiringCoinnameResponseDataInnerWithDefaults instantiates a new AcquiringCoinnameResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcquiringCoinnameResponseDataInnerWithDefaults() *AcquiringCoinnameResponseDataInner {
	this := AcquiringCoinnameResponseDataInner{}
	return &this
}

// GetCoinName returns the CoinName field value if set, zero value otherwise.
func (o *AcquiringCoinnameResponseDataInner) GetCoinName() string {
	if o == nil || common.IsNil(o.CoinName) {
		var ret string
		return ret
	}
	return *o.CoinName
}

// GetCoinNameOk returns a tuple with the CoinName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquiringCoinnameResponseDataInner) GetCoinNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.CoinName) {
		return nil, false
	}
	return o.CoinName, true
}

// HasCoinName returns a boolean if a field has been set.
func (o *AcquiringCoinnameResponseDataInner) HasCoinName() bool {
	if o != nil && !common.IsNil(o.CoinName) {
		return true
	}

	return false
}

// SetCoinName gets a reference to the given string and assigns it to the CoinName field.
func (o *AcquiringCoinnameResponseDataInner) SetCoinName(v string) {
	o.CoinName = &v
}

// GetCoinId returns the CoinId field value if set, zero value otherwise.
func (o *AcquiringCoinnameResponseDataInner) GetCoinId() int64 {
	if o == nil || common.IsNil(o.CoinId) {
		var ret int64
		return ret
	}
	return *o.CoinId
}

// GetCoinIdOk returns a tuple with the CoinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquiringCoinnameResponseDataInner) GetCoinIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CoinId) {
		return nil, false
	}
	return o.CoinId, true
}

// HasCoinId returns a boolean if a field has been set.
func (o *AcquiringCoinnameResponseDataInner) HasCoinId() bool {
	if o != nil && !common.IsNil(o.CoinId) {
		return true
	}

	return false
}

// SetCoinId gets a reference to the given int64 and assigns it to the CoinId field.
func (o *AcquiringCoinnameResponseDataInner) SetCoinId(v int64) {
	o.CoinId = &v
}

// GetPoolIndex returns the PoolIndex field value if set, zero value otherwise.
func (o *AcquiringCoinnameResponseDataInner) GetPoolIndex() int64 {
	if o == nil || common.IsNil(o.PoolIndex) {
		var ret int64
		return ret
	}
	return *o.PoolIndex
}

// GetPoolIndexOk returns a tuple with the PoolIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquiringCoinnameResponseDataInner) GetPoolIndexOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PoolIndex) {
		return nil, false
	}
	return o.PoolIndex, true
}

// HasPoolIndex returns a boolean if a field has been set.
func (o *AcquiringCoinnameResponseDataInner) HasPoolIndex() bool {
	if o != nil && !common.IsNil(o.PoolIndex) {
		return true
	}

	return false
}

// SetPoolIndex gets a reference to the given int64 and assigns it to the PoolIndex field.
func (o *AcquiringCoinnameResponseDataInner) SetPoolIndex(v int64) {
	o.PoolIndex = &v
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *AcquiringCoinnameResponseDataInner) GetAlgoId() int64 {
	if o == nil || common.IsNil(o.AlgoId) {
		var ret int64
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquiringCoinnameResponseDataInner) GetAlgoIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *AcquiringCoinnameResponseDataInner) HasAlgoId() bool {
	if o != nil && !common.IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given int64 and assigns it to the AlgoId field.
func (o *AcquiringCoinnameResponseDataInner) SetAlgoId(v int64) {
	o.AlgoId = &v
}

// GetAlgoName returns the AlgoName field value if set, zero value otherwise.
func (o *AcquiringCoinnameResponseDataInner) GetAlgoName() string {
	if o == nil || common.IsNil(o.AlgoName) {
		var ret string
		return ret
	}
	return *o.AlgoName
}

// GetAlgoNameOk returns a tuple with the AlgoName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquiringCoinnameResponseDataInner) GetAlgoNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.AlgoName) {
		return nil, false
	}
	return o.AlgoName, true
}

// HasAlgoName returns a boolean if a field has been set.
func (o *AcquiringCoinnameResponseDataInner) HasAlgoName() bool {
	if o != nil && !common.IsNil(o.AlgoName) {
		return true
	}

	return false
}

// SetAlgoName gets a reference to the given string and assigns it to the AlgoName field.
func (o *AcquiringCoinnameResponseDataInner) SetAlgoName(v string) {
	o.AlgoName = &v
}

func (o AcquiringCoinnameResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcquiringCoinnameResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CoinName) {
		toSerialize["coinName"] = o.CoinName
	}
	if !common.IsNil(o.CoinId) {
		toSerialize["coinId"] = o.CoinId
	}
	if !common.IsNil(o.PoolIndex) {
		toSerialize["poolIndex"] = o.PoolIndex
	}
	if !common.IsNil(o.AlgoId) {
		toSerialize["algoId"] = o.AlgoId
	}
	if !common.IsNil(o.AlgoName) {
		toSerialize["algoName"] = o.AlgoName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AcquiringCoinnameResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	varAcquiringCoinnameResponseDataInner := _AcquiringCoinnameResponseDataInner{}

	err = json.Unmarshal(data, &varAcquiringCoinnameResponseDataInner)

	if err != nil {
		return err
	}

	*o = AcquiringCoinnameResponseDataInner(varAcquiringCoinnameResponseDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "coinName")
		delete(additionalProperties, "coinId")
		delete(additionalProperties, "poolIndex")
		delete(additionalProperties, "algoId")
		delete(additionalProperties, "algoName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcquiringCoinnameResponseDataInner struct {
	value *AcquiringCoinnameResponseDataInner
	isSet bool
}

func (v NullableAcquiringCoinnameResponseDataInner) Get() *AcquiringCoinnameResponseDataInner {
	return v.value
}

func (v *NullableAcquiringCoinnameResponseDataInner) Set(val *AcquiringCoinnameResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAcquiringCoinnameResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAcquiringCoinnameResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcquiringCoinnameResponseDataInner(val *AcquiringCoinnameResponseDataInner) *NullableAcquiringCoinnameResponseDataInner {
	return &NullableAcquiringCoinnameResponseDataInner{value: val, isSet: true}
}

func (v NullableAcquiringCoinnameResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcquiringCoinnameResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
