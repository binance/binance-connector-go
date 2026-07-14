/*
Portfolio Margin WebSocket Market Streams

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OpenOrderLoss type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OpenOrderLoss{}

// OpenOrderLoss struct for OpenOrderLoss
type OpenOrderLoss struct {
	// Event Time
	E *int64 `json:"E,omitempty"`
	// Update Data
	O                    []OpenOrderLossOInner `json:"O,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenOrderLoss OpenOrderLoss

// NewOpenOrderLoss instantiates a new OpenOrderLoss object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenOrderLoss() *OpenOrderLoss {
	this := OpenOrderLoss{}
	return &this
}

// NewOpenOrderLossWithDefaults instantiates a new OpenOrderLoss object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenOrderLossWithDefaults() *OpenOrderLoss {
	this := OpenOrderLoss{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OpenOrderLoss) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenOrderLoss) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *OpenOrderLoss) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *OpenOrderLoss) SetE(v int64) {
	o.E = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *OpenOrderLoss) GetO() []OpenOrderLossOInner {
	if o == nil || common.IsNil(o.O) {
		var ret []OpenOrderLossOInner
		return ret
	}
	return o.O
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenOrderLoss) GetOOk() ([]OpenOrderLossOInner, bool) {
	if o == nil || common.IsNil(o.O) {
		return nil, false
	}
	return o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *OpenOrderLoss) HasO() bool {
	if o != nil && !common.IsNil(o.O) {
		return true
	}

	return false
}

// SetO gets a reference to the given []OpenOrderLossOInner and assigns it to the O field.
func (o *OpenOrderLoss) SetO(v []OpenOrderLossOInner) {
	o.O = v
}

func (o OpenOrderLoss) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenOrderLoss) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.O) {
		toSerialize["O"] = o.O
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenOrderLoss) UnmarshalJSON(data []byte) (err error) {
	varOpenOrderLoss := _OpenOrderLoss{}

	err = json.Unmarshal(data, &varOpenOrderLoss)

	if err != nil {
		return err
	}

	*o = OpenOrderLoss(varOpenOrderLoss)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "O")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenOrderLoss struct {
	value *OpenOrderLoss
	isSet bool
}

func (v NullableOpenOrderLoss) Get() *OpenOrderLoss {
	return v.value
}

func (v *NullableOpenOrderLoss) Set(val *OpenOrderLoss) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenOrderLoss) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenOrderLoss) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenOrderLoss(val *OpenOrderLoss) *NullableOpenOrderLoss {
	return &NullableOpenOrderLoss{value: val, isSet: true}
}

func (v NullableOpenOrderLoss) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenOrderLoss) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
