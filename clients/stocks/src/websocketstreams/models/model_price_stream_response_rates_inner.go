/*
Stocks Trading WebSocket Streams

WebSocket stream definitions for Binance Stocks Trading. Base URL: wss://nbstream.binance.com/equity
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PriceStreamResponseRatesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PriceStreamResponseRatesInner{}

// PriceStreamResponseRatesInner struct for PriceStreamResponseRatesInner
type PriceStreamResponseRatesInner struct {
	// Symbol (UPPERCASE ticker), e.g. `\"AAPL\"`.
	S *string `json:"s,omitempty"`
	// Internal asset code, `EQ_{symbol}`. Reference only; do not use as a trading identifier.
	Ac *string `json:"ac,omitempty"`
	// Latest price, trailing-zero stripped.
	P *string `json:"p,omitempty"`
	// Price time (epoch milliseconds UTC); may be null.
	T *int64 `json:"t,omitempty"`
	// Previous day's RTH close price (reference); absent when unknown.
	Pc *string `json:"pc,omitempty"`
	// Today's RTH close price; absent before After-Hours.
	Tc *string `json:"tc,omitempty"`
	// Per-symbol market phase: `C` Closed, `ON` Overnight, `PRE` Pre-Market, `O` Market Open, `POST` Post-Market.
	Mp                   *string `json:"mp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PriceStreamResponseRatesInner PriceStreamResponseRatesInner

// NewPriceStreamResponseRatesInner instantiates a new PriceStreamResponseRatesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceStreamResponseRatesInner() *PriceStreamResponseRatesInner {
	this := PriceStreamResponseRatesInner{}
	return &this
}

// NewPriceStreamResponseRatesInnerWithDefaults instantiates a new PriceStreamResponseRatesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceStreamResponseRatesInnerWithDefaults() *PriceStreamResponseRatesInner {
	this := PriceStreamResponseRatesInner{}
	return &this
}

// GetS returns the S field value if set, zero value otherwise.
func (o *PriceStreamResponseRatesInner) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceStreamResponseRatesInner) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *PriceStreamResponseRatesInner) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *PriceStreamResponseRatesInner) SetS(v string) {
	o.S = &v
}

// GetAc returns the Ac field value if set, zero value otherwise.
func (o *PriceStreamResponseRatesInner) GetAc() string {
	if o == nil || common.IsNil(o.Ac) {
		var ret string
		return ret
	}
	return *o.Ac
}

// GetAcOk returns a tuple with the Ac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceStreamResponseRatesInner) GetAcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Ac) {
		return nil, false
	}
	return o.Ac, true
}

// HasAc returns a boolean if a field has been set.
func (o *PriceStreamResponseRatesInner) HasAc() bool {
	if o != nil && !common.IsNil(o.Ac) {
		return true
	}

	return false
}

// SetAc gets a reference to the given string and assigns it to the Ac field.
func (o *PriceStreamResponseRatesInner) SetAc(v string) {
	o.Ac = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *PriceStreamResponseRatesInner) GetP() string {
	if o == nil || common.IsNil(o.P) {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceStreamResponseRatesInner) GetPOk() (*string, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *PriceStreamResponseRatesInner) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *PriceStreamResponseRatesInner) SetP(v string) {
	o.P = &v
}

// GetT returns the T field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceStreamResponseRatesInner) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceStreamResponseRatesInner) GetTOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *PriceStreamResponseRatesInner) HasT() bool {
	if o != nil && common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given NullableInt64 and assigns it to the T field.
func (o *PriceStreamResponseRatesInner) SetT(v int64) {
	o.T = &v
}

// GetPc returns the Pc field value if set, zero value otherwise.
func (o *PriceStreamResponseRatesInner) GetPc() string {
	if o == nil || common.IsNil(o.Pc) {
		var ret string
		return ret
	}
	return *o.Pc
}

// GetPcOk returns a tuple with the Pc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceStreamResponseRatesInner) GetPcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pc) {
		return nil, false
	}
	return o.Pc, true
}

// HasPc returns a boolean if a field has been set.
func (o *PriceStreamResponseRatesInner) HasPc() bool {
	if o != nil && !common.IsNil(o.Pc) {
		return true
	}

	return false
}

// SetPc gets a reference to the given string and assigns it to the Pc field.
func (o *PriceStreamResponseRatesInner) SetPc(v string) {
	o.Pc = &v
}

// GetTc returns the Tc field value if set, zero value otherwise.
func (o *PriceStreamResponseRatesInner) GetTc() string {
	if o == nil || common.IsNil(o.Tc) {
		var ret string
		return ret
	}
	return *o.Tc
}

// GetTcOk returns a tuple with the Tc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceStreamResponseRatesInner) GetTcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Tc) {
		return nil, false
	}
	return o.Tc, true
}

// HasTc returns a boolean if a field has been set.
func (o *PriceStreamResponseRatesInner) HasTc() bool {
	if o != nil && !common.IsNil(o.Tc) {
		return true
	}

	return false
}

// SetTc gets a reference to the given string and assigns it to the Tc field.
func (o *PriceStreamResponseRatesInner) SetTc(v string) {
	o.Tc = &v
}

// GetMp returns the Mp field value if set, zero value otherwise.
func (o *PriceStreamResponseRatesInner) GetMp() string {
	if o == nil || common.IsNil(o.Mp) {
		var ret string
		return ret
	}
	return *o.Mp
}

// GetMpOk returns a tuple with the Mp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceStreamResponseRatesInner) GetMpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Mp) {
		return nil, false
	}
	return o.Mp, true
}

// HasMp returns a boolean if a field has been set.
func (o *PriceStreamResponseRatesInner) HasMp() bool {
	if o != nil && !common.IsNil(o.Mp) {
		return true
	}

	return false
}

// SetMp gets a reference to the given string and assigns it to the Mp field.
func (o *PriceStreamResponseRatesInner) SetMp(v string) {
	o.Mp = &v
}

func (o PriceStreamResponseRatesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceStreamResponseRatesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.S) {
		toSerialize["s"] = o.S
	}
	if !common.IsNil(o.Ac) {
		toSerialize["ac"] = o.Ac
	}
	if !common.IsNil(o.P) {
		toSerialize["p"] = o.P
	}
	if !common.IsNil(o.T) {
		toSerialize["t"] = o.T
	}
	if !common.IsNil(o.Pc) {
		toSerialize["pc"] = o.Pc
	}
	if !common.IsNil(o.Tc) {
		toSerialize["tc"] = o.Tc
	}
	if !common.IsNil(o.Mp) {
		toSerialize["mp"] = o.Mp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PriceStreamResponseRatesInner) UnmarshalJSON(data []byte) (err error) {
	varPriceStreamResponseRatesInner := _PriceStreamResponseRatesInner{}

	err = json.Unmarshal(data, &varPriceStreamResponseRatesInner)

	if err != nil {
		return err
	}

	*o = PriceStreamResponseRatesInner(varPriceStreamResponseRatesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "s")
		delete(additionalProperties, "ac")
		delete(additionalProperties, "p")
		delete(additionalProperties, "t")
		delete(additionalProperties, "pc")
		delete(additionalProperties, "tc")
		delete(additionalProperties, "mp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePriceStreamResponseRatesInner struct {
	value *PriceStreamResponseRatesInner
	isSet bool
}

func (v NullablePriceStreamResponseRatesInner) Get() *PriceStreamResponseRatesInner {
	return v.value
}

func (v *NullablePriceStreamResponseRatesInner) Set(val *PriceStreamResponseRatesInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceStreamResponseRatesInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceStreamResponseRatesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceStreamResponseRatesInner(val *PriceStreamResponseRatesInner) *NullablePriceStreamResponseRatesInner {
	return &NullablePriceStreamResponseRatesInner{value: val, isSet: true}
}

func (v NullablePriceStreamResponseRatesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceStreamResponseRatesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
