/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ContractKlineStreamResponseK type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ContractKlineStreamResponseK{}

// ContractKlineStreamResponseK Kline payload
type ContractKlineStreamResponseK struct {
	// Open price
	O *string `json:"o,omitempty"`
	// Close price
	C *string `json:"c,omitempty"`
	// High price
	H *string `json:"h,omitempty"`
	// Low price
	L *string `json:"l,omitempty"`
	// Volume
	V *string `json:"v,omitempty"`
	// Kline open time
	Ot *int64 `json:"ot,omitempty"`
	// Kline close time
	Ct *int64 `json:"ct,omitempty"`
	// Interval
	I                    *string `json:"i,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContractKlineStreamResponseK ContractKlineStreamResponseK

// NewContractKlineStreamResponseK instantiates a new ContractKlineStreamResponseK object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractKlineStreamResponseK() *ContractKlineStreamResponseK {
	this := ContractKlineStreamResponseK{}
	return &this
}

// NewContractKlineStreamResponseKWithDefaults instantiates a new ContractKlineStreamResponseK object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractKlineStreamResponseKWithDefaults() *ContractKlineStreamResponseK {
	this := ContractKlineStreamResponseK{}
	return &this
}

// GetO returns the O field value if set, zero value otherwise.
func (o *ContractKlineStreamResponseK) GetO() string {
	if o == nil || common.IsNil(o.O) {
		var ret string
		return ret
	}
	return *o.O
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractKlineStreamResponseK) GetOOk() (*string, bool) {
	if o == nil || common.IsNil(o.O) {
		return nil, false
	}
	return o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *ContractKlineStreamResponseK) HasO() bool {
	if o != nil && !common.IsNil(o.O) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *ContractKlineStreamResponseK) SetO(v string) {
	o.O = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *ContractKlineStreamResponseK) GetC() string {
	if o == nil || common.IsNil(o.C) {
		var ret string
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractKlineStreamResponseK) GetCOk() (*string, bool) {
	if o == nil || common.IsNil(o.C) {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *ContractKlineStreamResponseK) HasC() bool {
	if o != nil && !common.IsNil(o.C) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *ContractKlineStreamResponseK) SetC(v string) {
	o.C = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *ContractKlineStreamResponseK) GetH() string {
	if o == nil || common.IsNil(o.H) {
		var ret string
		return ret
	}
	return *o.H
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractKlineStreamResponseK) GetHOk() (*string, bool) {
	if o == nil || common.IsNil(o.H) {
		return nil, false
	}
	return o.H, true
}

// HasH returns a boolean if a field has been set.
func (o *ContractKlineStreamResponseK) HasH() bool {
	if o != nil && !common.IsNil(o.H) {
		return true
	}

	return false
}

// SetH gets a reference to the given string and assigns it to the H field.
func (o *ContractKlineStreamResponseK) SetH(v string) {
	o.H = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *ContractKlineStreamResponseK) GetL() string {
	if o == nil || common.IsNil(o.L) {
		var ret string
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractKlineStreamResponseK) GetLOk() (*string, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *ContractKlineStreamResponseK) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *ContractKlineStreamResponseK) SetL(v string) {
	o.L = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *ContractKlineStreamResponseK) GetV() string {
	if o == nil || common.IsNil(o.V) {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractKlineStreamResponseK) GetVOk() (*string, bool) {
	if o == nil || common.IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *ContractKlineStreamResponseK) HasV() bool {
	if o != nil && !common.IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *ContractKlineStreamResponseK) SetV(v string) {
	o.V = &v
}

// GetOt returns the Ot field value if set, zero value otherwise.
func (o *ContractKlineStreamResponseK) GetOt() int64 {
	if o == nil || common.IsNil(o.Ot) {
		var ret int64
		return ret
	}
	return *o.Ot
}

// GetOtOk returns a tuple with the Ot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractKlineStreamResponseK) GetOtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Ot) {
		return nil, false
	}
	return o.Ot, true
}

// HasOt returns a boolean if a field has been set.
func (o *ContractKlineStreamResponseK) HasOt() bool {
	if o != nil && !common.IsNil(o.Ot) {
		return true
	}

	return false
}

// SetOt gets a reference to the given int64 and assigns it to the Ot field.
func (o *ContractKlineStreamResponseK) SetOt(v int64) {
	o.Ot = &v
}

// GetCt returns the Ct field value if set, zero value otherwise.
func (o *ContractKlineStreamResponseK) GetCt() int64 {
	if o == nil || common.IsNil(o.Ct) {
		var ret int64
		return ret
	}
	return *o.Ct
}

// GetCtOk returns a tuple with the Ct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractKlineStreamResponseK) GetCtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Ct) {
		return nil, false
	}
	return o.Ct, true
}

// HasCt returns a boolean if a field has been set.
func (o *ContractKlineStreamResponseK) HasCt() bool {
	if o != nil && !common.IsNil(o.Ct) {
		return true
	}

	return false
}

// SetCt gets a reference to the given int64 and assigns it to the Ct field.
func (o *ContractKlineStreamResponseK) SetCt(v int64) {
	o.Ct = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *ContractKlineStreamResponseK) GetI() string {
	if o == nil || common.IsNil(o.I) {
		var ret string
		return ret
	}
	return *o.I
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractKlineStreamResponseK) GetIOk() (*string, bool) {
	if o == nil || common.IsNil(o.I) {
		return nil, false
	}
	return o.I, true
}

// HasI returns a boolean if a field has been set.
func (o *ContractKlineStreamResponseK) HasI() bool {
	if o != nil && !common.IsNil(o.I) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *ContractKlineStreamResponseK) SetI(v string) {
	o.I = &v
}

func (o ContractKlineStreamResponseK) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractKlineStreamResponseK) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.O) {
		toSerialize["o"] = o.O
	}
	if !common.IsNil(o.C) {
		toSerialize["c"] = o.C
	}
	if !common.IsNil(o.H) {
		toSerialize["h"] = o.H
	}
	if !common.IsNil(o.L) {
		toSerialize["l"] = o.L
	}
	if !common.IsNil(o.V) {
		toSerialize["v"] = o.V
	}
	if !common.IsNil(o.Ot) {
		toSerialize["ot"] = o.Ot
	}
	if !common.IsNil(o.Ct) {
		toSerialize["ct"] = o.Ct
	}
	if !common.IsNil(o.I) {
		toSerialize["i"] = o.I
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContractKlineStreamResponseK) UnmarshalJSON(data []byte) (err error) {
	varContractKlineStreamResponseK := _ContractKlineStreamResponseK{}

	err = json.Unmarshal(data, &varContractKlineStreamResponseK)

	if err != nil {
		return err
	}

	*o = ContractKlineStreamResponseK(varContractKlineStreamResponseK)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "o")
		delete(additionalProperties, "c")
		delete(additionalProperties, "h")
		delete(additionalProperties, "l")
		delete(additionalProperties, "v")
		delete(additionalProperties, "ot")
		delete(additionalProperties, "ct")
		delete(additionalProperties, "i")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContractKlineStreamResponseK struct {
	value *ContractKlineStreamResponseK
	isSet bool
}

func (v NullableContractKlineStreamResponseK) Get() *ContractKlineStreamResponseK {
	return v.value
}

func (v *NullableContractKlineStreamResponseK) Set(val *ContractKlineStreamResponseK) {
	v.value = val
	v.isSet = true
}

func (v NullableContractKlineStreamResponseK) IsSet() bool {
	return v.isSet
}

func (v *NullableContractKlineStreamResponseK) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractKlineStreamResponseK(val *ContractKlineStreamResponseK) *NullableContractKlineStreamResponseK {
	return &NullableContractKlineStreamResponseK{value: val, isSet: true}
}

func (v NullableContractKlineStreamResponseK) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractKlineStreamResponseK) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
