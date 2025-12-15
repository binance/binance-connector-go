/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderTradeUpdateOInnerFiInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderTradeUpdateOInnerFiInner{}

// OrderTradeUpdateOInnerFiInner struct for OrderTradeUpdateOInnerFiInner
type OrderTradeUpdateOInnerFiInner struct {
	Smallt               *string `json:"t,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	Smallq               *string `json:"q,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	Smallm               *string `json:"m,omitempty"`
	Smallf               *string `json:"f,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderTradeUpdateOInnerFiInner OrderTradeUpdateOInnerFiInner

// NewOrderTradeUpdateOInnerFiInner instantiates a new OrderTradeUpdateOInnerFiInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTradeUpdateOInnerFiInner() *OrderTradeUpdateOInnerFiInner {
	this := OrderTradeUpdateOInnerFiInner{}
	return &this
}

// NewOrderTradeUpdateOInnerFiInnerWithDefaults instantiates a new OrderTradeUpdateOInnerFiInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTradeUpdateOInnerFiInnerWithDefaults() *OrderTradeUpdateOInnerFiInner {
	this := OrderTradeUpdateOInnerFiInner{}
	return &this
}

// GetT returns the T field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInnerFiInner) GetSmallt() string {
	if o == nil || common.IsNil(o.Smallt) {
		var ret string
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInnerFiInner) GetSmalltOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInnerFiInner) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given string and assigns it to the T field.
func (o *OrderTradeUpdateOInnerFiInner) SetSmallt(v string) {
	o.Smallt = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInnerFiInner) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInnerFiInner) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInnerFiInner) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *OrderTradeUpdateOInnerFiInner) SetSmallp(v string) {
	o.Smallp = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInnerFiInner) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInnerFiInner) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInnerFiInner) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *OrderTradeUpdateOInnerFiInner) SetSmallq(v string) {
	o.Smallq = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInnerFiInner) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInnerFiInner) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInnerFiInner) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *OrderTradeUpdateOInnerFiInner) SetT(v int64) {
	o.T = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInnerFiInner) GetSmallm() string {
	if o == nil || common.IsNil(o.Smallm) {
		var ret string
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInnerFiInner) GetSmallmOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInnerFiInner) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given string and assigns it to the M field.
func (o *OrderTradeUpdateOInnerFiInner) SetSmallm(v string) {
	o.Smallm = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *OrderTradeUpdateOInnerFiInner) GetSmallf() string {
	if o == nil || common.IsNil(o.Smallf) {
		var ret string
		return ret
	}
	return *o.Smallf
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateOInnerFiInner) GetSmallfOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallf) {
		return nil, false
	}
	return o.Smallf, true
}

// HasF returns a boolean if a field has been set.
func (o *OrderTradeUpdateOInnerFiInner) HasSmallf() bool {
	if o != nil && !common.IsNil(o.Smallf) {
		return true
	}

	return false
}

// SetF gets a reference to the given string and assigns it to the F field.
func (o *OrderTradeUpdateOInnerFiInner) SetSmallf(v string) {
	o.Smallf = &v
}

func (o OrderTradeUpdateOInnerFiInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTradeUpdateOInnerFiInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smallt) {
		toSerialize["t"] = o.Smallt
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smallm) {
		toSerialize["m"] = o.Smallm
	}
	if !common.IsNil(o.Smallf) {
		toSerialize["f"] = o.Smallf
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderTradeUpdateOInnerFiInner) UnmarshalJSON(data []byte) (err error) {
	varOrderTradeUpdateOInnerFiInner := _OrderTradeUpdateOInnerFiInner{}

	err = json.Unmarshal(data, &varOrderTradeUpdateOInnerFiInner)

	if err != nil {
		return err
	}

	*o = OrderTradeUpdateOInnerFiInner(varOrderTradeUpdateOInnerFiInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "t")
		delete(additionalProperties, "p")
		delete(additionalProperties, "q")
		delete(additionalProperties, "T")
		delete(additionalProperties, "m")
		delete(additionalProperties, "f")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderTradeUpdateOInnerFiInner struct {
	value *OrderTradeUpdateOInnerFiInner
	isSet bool
}

func (v NullableOrderTradeUpdateOInnerFiInner) Get() *OrderTradeUpdateOInnerFiInner {
	return v.value
}

func (v *NullableOrderTradeUpdateOInnerFiInner) Set(val *OrderTradeUpdateOInnerFiInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTradeUpdateOInnerFiInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTradeUpdateOInnerFiInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTradeUpdateOInnerFiInner(val *OrderTradeUpdateOInnerFiInner) *NullableOrderTradeUpdateOInnerFiInner {
	return &NullableOrderTradeUpdateOInnerFiInner{value: val, isSet: true}
}

func (v NullableOrderTradeUpdateOInnerFiInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTradeUpdateOInnerFiInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
