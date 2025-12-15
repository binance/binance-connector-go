/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

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
	BuySellRatio         *string `json:"buySellRatio,omitempty"`
	BuyVol               *string `json:"buyVol,omitempty"`
	SellVol              *string `json:"sellVol,omitempty"`
	Timestamp            *string `json:"timestamp,omitempty"`
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

// GetBuySellRatio returns the BuySellRatio field value if set, zero value otherwise.
func (o *TakerBuySellVolumeResponseInner) GetBuySellRatio() string {
	if o == nil || common.IsNil(o.BuySellRatio) {
		var ret string
		return ret
	}
	return *o.BuySellRatio
}

// GetBuySellRatioOk returns a tuple with the BuySellRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TakerBuySellVolumeResponseInner) GetBuySellRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.BuySellRatio) {
		return nil, false
	}
	return o.BuySellRatio, true
}

// HasBuySellRatio returns a boolean if a field has been set.
func (o *TakerBuySellVolumeResponseInner) HasBuySellRatio() bool {
	if o != nil && !common.IsNil(o.BuySellRatio) {
		return true
	}

	return false
}

// SetBuySellRatio gets a reference to the given string and assigns it to the BuySellRatio field.
func (o *TakerBuySellVolumeResponseInner) SetBuySellRatio(v string) {
	o.BuySellRatio = &v
}

// GetBuyVol returns the BuyVol field value if set, zero value otherwise.
func (o *TakerBuySellVolumeResponseInner) GetBuyVol() string {
	if o == nil || common.IsNil(o.BuyVol) {
		var ret string
		return ret
	}
	return *o.BuyVol
}

// GetBuyVolOk returns a tuple with the BuyVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TakerBuySellVolumeResponseInner) GetBuyVolOk() (*string, bool) {
	if o == nil || common.IsNil(o.BuyVol) {
		return nil, false
	}
	return o.BuyVol, true
}

// HasBuyVol returns a boolean if a field has been set.
func (o *TakerBuySellVolumeResponseInner) HasBuyVol() bool {
	if o != nil && !common.IsNil(o.BuyVol) {
		return true
	}

	return false
}

// SetBuyVol gets a reference to the given string and assigns it to the BuyVol field.
func (o *TakerBuySellVolumeResponseInner) SetBuyVol(v string) {
	o.BuyVol = &v
}

// GetSellVol returns the SellVol field value if set, zero value otherwise.
func (o *TakerBuySellVolumeResponseInner) GetSellVol() string {
	if o == nil || common.IsNil(o.SellVol) {
		var ret string
		return ret
	}
	return *o.SellVol
}

// GetSellVolOk returns a tuple with the SellVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TakerBuySellVolumeResponseInner) GetSellVolOk() (*string, bool) {
	if o == nil || common.IsNil(o.SellVol) {
		return nil, false
	}
	return o.SellVol, true
}

// HasSellVol returns a boolean if a field has been set.
func (o *TakerBuySellVolumeResponseInner) HasSellVol() bool {
	if o != nil && !common.IsNil(o.SellVol) {
		return true
	}

	return false
}

// SetSellVol gets a reference to the given string and assigns it to the SellVol field.
func (o *TakerBuySellVolumeResponseInner) SetSellVol(v string) {
	o.SellVol = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TakerBuySellVolumeResponseInner) GetTimestamp() string {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TakerBuySellVolumeResponseInner) GetTimestampOk() (*string, bool) {
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

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *TakerBuySellVolumeResponseInner) SetTimestamp(v string) {
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
	if !common.IsNil(o.BuySellRatio) {
		toSerialize["buySellRatio"] = o.BuySellRatio
	}
	if !common.IsNil(o.BuyVol) {
		toSerialize["buyVol"] = o.BuyVol
	}
	if !common.IsNil(o.SellVol) {
		toSerialize["sellVol"] = o.SellVol
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
		delete(additionalProperties, "buySellRatio")
		delete(additionalProperties, "buyVol")
		delete(additionalProperties, "sellVol")
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
