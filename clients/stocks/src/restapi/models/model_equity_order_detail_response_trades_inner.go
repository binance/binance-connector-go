/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the EquityOrderDetailResponseTradesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EquityOrderDetailResponseTradesInner{}

// EquityOrderDetailResponseTradesInner struct for EquityOrderDetailResponseTradesInner
type EquityOrderDetailResponseTradesInner struct {
	// Execution (fill) id.
	ExecutionId *string `json:"executionId,omitempty"`
	// Execution time (ms epoch).
	ExecutionAt *int64 `json:"executionAt,omitempty"`
	// Fill price for this execution (USD).
	Price *string `json:"price,omitempty"`
	// Fill quantity.
	Qty                  *string `json:"qty,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquityOrderDetailResponseTradesInner EquityOrderDetailResponseTradesInner

// NewEquityOrderDetailResponseTradesInner instantiates a new EquityOrderDetailResponseTradesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquityOrderDetailResponseTradesInner() *EquityOrderDetailResponseTradesInner {
	this := EquityOrderDetailResponseTradesInner{}
	return &this
}

// NewEquityOrderDetailResponseTradesInnerWithDefaults instantiates a new EquityOrderDetailResponseTradesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquityOrderDetailResponseTradesInnerWithDefaults() *EquityOrderDetailResponseTradesInner {
	this := EquityOrderDetailResponseTradesInner{}
	return &this
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *EquityOrderDetailResponseTradesInner) GetExecutionId() string {
	if o == nil || common.IsNil(o.ExecutionId) {
		var ret string
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderDetailResponseTradesInner) GetExecutionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutionId) {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *EquityOrderDetailResponseTradesInner) HasExecutionId() bool {
	if o != nil && !common.IsNil(o.ExecutionId) {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given string and assigns it to the ExecutionId field.
func (o *EquityOrderDetailResponseTradesInner) SetExecutionId(v string) {
	o.ExecutionId = &v
}

// GetExecutionAt returns the ExecutionAt field value if set, zero value otherwise.
func (o *EquityOrderDetailResponseTradesInner) GetExecutionAt() int64 {
	if o == nil || common.IsNil(o.ExecutionAt) {
		var ret int64
		return ret
	}
	return *o.ExecutionAt
}

// GetExecutionAtOk returns a tuple with the ExecutionAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderDetailResponseTradesInner) GetExecutionAtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExecutionAt) {
		return nil, false
	}
	return o.ExecutionAt, true
}

// HasExecutionAt returns a boolean if a field has been set.
func (o *EquityOrderDetailResponseTradesInner) HasExecutionAt() bool {
	if o != nil && !common.IsNil(o.ExecutionAt) {
		return true
	}

	return false
}

// SetExecutionAt gets a reference to the given int64 and assigns it to the ExecutionAt field.
func (o *EquityOrderDetailResponseTradesInner) SetExecutionAt(v int64) {
	o.ExecutionAt = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *EquityOrderDetailResponseTradesInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderDetailResponseTradesInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *EquityOrderDetailResponseTradesInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *EquityOrderDetailResponseTradesInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *EquityOrderDetailResponseTradesInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderDetailResponseTradesInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *EquityOrderDetailResponseTradesInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *EquityOrderDetailResponseTradesInner) SetQty(v string) {
	o.Qty = &v
}

func (o EquityOrderDetailResponseTradesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquityOrderDetailResponseTradesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ExecutionId) {
		toSerialize["executionId"] = o.ExecutionId
	}
	if !common.IsNil(o.ExecutionAt) {
		toSerialize["executionAt"] = o.ExecutionAt
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquityOrderDetailResponseTradesInner) UnmarshalJSON(data []byte) (err error) {
	varEquityOrderDetailResponseTradesInner := _EquityOrderDetailResponseTradesInner{}

	err = json.Unmarshal(data, &varEquityOrderDetailResponseTradesInner)

	if err != nil {
		return err
	}

	*o = EquityOrderDetailResponseTradesInner(varEquityOrderDetailResponseTradesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "executionId")
		delete(additionalProperties, "executionAt")
		delete(additionalProperties, "price")
		delete(additionalProperties, "qty")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquityOrderDetailResponseTradesInner struct {
	value *EquityOrderDetailResponseTradesInner
	isSet bool
}

func (v NullableEquityOrderDetailResponseTradesInner) Get() *EquityOrderDetailResponseTradesInner {
	return v.value
}

func (v *NullableEquityOrderDetailResponseTradesInner) Set(val *EquityOrderDetailResponseTradesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEquityOrderDetailResponseTradesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEquityOrderDetailResponseTradesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquityOrderDetailResponseTradesInner(val *EquityOrderDetailResponseTradesInner) *NullableEquityOrderDetailResponseTradesInner {
	return &NullableEquityOrderDetailResponseTradesInner{value: val, isSet: true}
}

func (v NullableEquityOrderDetailResponseTradesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquityOrderDetailResponseTradesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
