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

// checks if the TakerBuySellVolumeResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TakerBuySellVolumeResponseInner{}

// TakerBuySellVolumeResponseInner struct for TakerBuySellVolumeResponseInner
type TakerBuySellVolumeResponseInner struct {
	Pair                 *string `json:"pair,omitempty"`
	ContractType         *string `json:"contractType,omitempty"`
	TakerBuyVol          *string `json:"takerBuyVol,omitempty"`
	TakerSellVol         *string `json:"takerSellVol,omitempty"`
	TakerBuyVolValue     *string `json:"takerBuyVolValue,omitempty"`
	TakerSellVolValue    *string `json:"takerSellVolValue,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TakerBuySellVolumeResponseInner TakerBuySellVolumeResponseInner

// NewTakerBuySellVolumeResponseInner instantiates a new TakerBuySellVolumeResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTakerBuySellVolumeResponseInner() *TakerBuySellVolumeResponseInner {
	this := TakerBuySellVolumeResponseInner{}
	return &this
}

// NewTakerBuySellVolumeResponseInnerWithDefaults instantiates a new TakerBuySellVolumeResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTakerBuySellVolumeResponseInnerWithDefaults() *TakerBuySellVolumeResponseInner {
	this := TakerBuySellVolumeResponseInner{}
	return &this
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *TakerBuySellVolumeResponseInner) GetPair() string {
	if o == nil || common.IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TakerBuySellVolumeResponseInner) GetPairOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *TakerBuySellVolumeResponseInner) HasPair() bool {
	if o != nil && !common.IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *TakerBuySellVolumeResponseInner) SetPair(v string) {
	o.Pair = &v
}

// GetContractType returns the ContractType field value if set, zero value otherwise.
func (o *TakerBuySellVolumeResponseInner) GetContractType() string {
	if o == nil || common.IsNil(o.ContractType) {
		var ret string
		return ret
	}
	return *o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TakerBuySellVolumeResponseInner) GetContractTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContractType) {
		return nil, false
	}
	return o.ContractType, true
}

// HasContractType returns a boolean if a field has been set.
func (o *TakerBuySellVolumeResponseInner) HasContractType() bool {
	if o != nil && !common.IsNil(o.ContractType) {
		return true
	}

	return false
}

// SetContractType gets a reference to the given string and assigns it to the ContractType field.
func (o *TakerBuySellVolumeResponseInner) SetContractType(v string) {
	o.ContractType = &v
}

// GetTakerBuyVol returns the TakerBuyVol field value if set, zero value otherwise.
func (o *TakerBuySellVolumeResponseInner) GetTakerBuyVol() string {
	if o == nil || common.IsNil(o.TakerBuyVol) {
		var ret string
		return ret
	}
	return *o.TakerBuyVol
}

// GetTakerBuyVolOk returns a tuple with the TakerBuyVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TakerBuySellVolumeResponseInner) GetTakerBuyVolOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerBuyVol) {
		return nil, false
	}
	return o.TakerBuyVol, true
}

// HasTakerBuyVol returns a boolean if a field has been set.
func (o *TakerBuySellVolumeResponseInner) HasTakerBuyVol() bool {
	if o != nil && !common.IsNil(o.TakerBuyVol) {
		return true
	}

	return false
}

// SetTakerBuyVol gets a reference to the given string and assigns it to the TakerBuyVol field.
func (o *TakerBuySellVolumeResponseInner) SetTakerBuyVol(v string) {
	o.TakerBuyVol = &v
}

// GetTakerSellVol returns the TakerSellVol field value if set, zero value otherwise.
func (o *TakerBuySellVolumeResponseInner) GetTakerSellVol() string {
	if o == nil || common.IsNil(o.TakerSellVol) {
		var ret string
		return ret
	}
	return *o.TakerSellVol
}

// GetTakerSellVolOk returns a tuple with the TakerSellVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TakerBuySellVolumeResponseInner) GetTakerSellVolOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerSellVol) {
		return nil, false
	}
	return o.TakerSellVol, true
}

// HasTakerSellVol returns a boolean if a field has been set.
func (o *TakerBuySellVolumeResponseInner) HasTakerSellVol() bool {
	if o != nil && !common.IsNil(o.TakerSellVol) {
		return true
	}

	return false
}

// SetTakerSellVol gets a reference to the given string and assigns it to the TakerSellVol field.
func (o *TakerBuySellVolumeResponseInner) SetTakerSellVol(v string) {
	o.TakerSellVol = &v
}

// GetTakerBuyVolValue returns the TakerBuyVolValue field value if set, zero value otherwise.
func (o *TakerBuySellVolumeResponseInner) GetTakerBuyVolValue() string {
	if o == nil || common.IsNil(o.TakerBuyVolValue) {
		var ret string
		return ret
	}
	return *o.TakerBuyVolValue
}

// GetTakerBuyVolValueOk returns a tuple with the TakerBuyVolValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TakerBuySellVolumeResponseInner) GetTakerBuyVolValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerBuyVolValue) {
		return nil, false
	}
	return o.TakerBuyVolValue, true
}

// HasTakerBuyVolValue returns a boolean if a field has been set.
func (o *TakerBuySellVolumeResponseInner) HasTakerBuyVolValue() bool {
	if o != nil && !common.IsNil(o.TakerBuyVolValue) {
		return true
	}

	return false
}

// SetTakerBuyVolValue gets a reference to the given string and assigns it to the TakerBuyVolValue field.
func (o *TakerBuySellVolumeResponseInner) SetTakerBuyVolValue(v string) {
	o.TakerBuyVolValue = &v
}

// GetTakerSellVolValue returns the TakerSellVolValue field value if set, zero value otherwise.
func (o *TakerBuySellVolumeResponseInner) GetTakerSellVolValue() string {
	if o == nil || common.IsNil(o.TakerSellVolValue) {
		var ret string
		return ret
	}
	return *o.TakerSellVolValue
}

// GetTakerSellVolValueOk returns a tuple with the TakerSellVolValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TakerBuySellVolumeResponseInner) GetTakerSellVolValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerSellVolValue) {
		return nil, false
	}
	return o.TakerSellVolValue, true
}

// HasTakerSellVolValue returns a boolean if a field has been set.
func (o *TakerBuySellVolumeResponseInner) HasTakerSellVolValue() bool {
	if o != nil && !common.IsNil(o.TakerSellVolValue) {
		return true
	}

	return false
}

// SetTakerSellVolValue gets a reference to the given string and assigns it to the TakerSellVolValue field.
func (o *TakerBuySellVolumeResponseInner) SetTakerSellVolValue(v string) {
	o.TakerSellVolValue = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TakerBuySellVolumeResponseInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TakerBuySellVolumeResponseInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TakerBuySellVolumeResponseInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *TakerBuySellVolumeResponseInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o TakerBuySellVolumeResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TakerBuySellVolumeResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !common.IsNil(o.ContractType) {
		toSerialize["contractType"] = o.ContractType
	}
	if !common.IsNil(o.TakerBuyVol) {
		toSerialize["takerBuyVol"] = o.TakerBuyVol
	}
	if !common.IsNil(o.TakerSellVol) {
		toSerialize["takerSellVol"] = o.TakerSellVol
	}
	if !common.IsNil(o.TakerBuyVolValue) {
		toSerialize["takerBuyVolValue"] = o.TakerBuyVolValue
	}
	if !common.IsNil(o.TakerSellVolValue) {
		toSerialize["takerSellVolValue"] = o.TakerSellVolValue
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TakerBuySellVolumeResponseInner) UnmarshalJSON(data []byte) (err error) {
	varTakerBuySellVolumeResponseInner := _TakerBuySellVolumeResponseInner{}

	err = json.Unmarshal(data, &varTakerBuySellVolumeResponseInner)

	if err != nil {
		return err
	}

	*o = TakerBuySellVolumeResponseInner(varTakerBuySellVolumeResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pair")
		delete(additionalProperties, "contractType")
		delete(additionalProperties, "takerBuyVol")
		delete(additionalProperties, "takerSellVol")
		delete(additionalProperties, "takerBuyVolValue")
		delete(additionalProperties, "takerSellVolValue")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTakerBuySellVolumeResponseInner struct {
	value *TakerBuySellVolumeResponseInner
	isSet bool
}

func (v NullableTakerBuySellVolumeResponseInner) Get() *TakerBuySellVolumeResponseInner {
	return v.value
}

func (v *NullableTakerBuySellVolumeResponseInner) Set(val *TakerBuySellVolumeResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTakerBuySellVolumeResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTakerBuySellVolumeResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTakerBuySellVolumeResponseInner(val *TakerBuySellVolumeResponseInner) *NullableTakerBuySellVolumeResponseInner {
	return &NullableTakerBuySellVolumeResponseInner{value: val, isSet: true}
}

func (v NullableTakerBuySellVolumeResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTakerBuySellVolumeResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
