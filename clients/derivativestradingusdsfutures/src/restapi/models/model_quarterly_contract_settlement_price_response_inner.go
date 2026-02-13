/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QuarterlyContractSettlementPriceResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuarterlyContractSettlementPriceResponseInner{}

// QuarterlyContractSettlementPriceResponseInner struct for QuarterlyContractSettlementPriceResponseInner
type QuarterlyContractSettlementPriceResponseInner struct {
	DeliveryTime         *int64   `json:"deliveryTime,omitempty"`
	DeliveryPrice        *float32 `json:"deliveryPrice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuarterlyContractSettlementPriceResponseInner QuarterlyContractSettlementPriceResponseInner

// NewQuarterlyContractSettlementPriceResponseInner instantiates a new QuarterlyContractSettlementPriceResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuarterlyContractSettlementPriceResponseInner() *QuarterlyContractSettlementPriceResponseInner {
	this := QuarterlyContractSettlementPriceResponseInner{}
	return &this
}

// NewQuarterlyContractSettlementPriceResponseInnerWithDefaults instantiates a new QuarterlyContractSettlementPriceResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuarterlyContractSettlementPriceResponseInnerWithDefaults() *QuarterlyContractSettlementPriceResponseInner {
	this := QuarterlyContractSettlementPriceResponseInner{}
	return &this
}

// GetDeliveryTime returns the DeliveryTime field value if set, zero value otherwise.
func (o *QuarterlyContractSettlementPriceResponseInner) GetDeliveryTime() int64 {
	if o == nil || common.IsNil(o.DeliveryTime) {
		var ret int64
		return ret
	}
	return *o.DeliveryTime
}

// GetDeliveryTimeOk returns a tuple with the DeliveryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuarterlyContractSettlementPriceResponseInner) GetDeliveryTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.DeliveryTime) {
		return nil, false
	}
	return o.DeliveryTime, true
}

// HasDeliveryTime returns a boolean if a field has been set.
func (o *QuarterlyContractSettlementPriceResponseInner) HasDeliveryTime() bool {
	if o != nil && !common.IsNil(o.DeliveryTime) {
		return true
	}

	return false
}

// SetDeliveryTime gets a reference to the given int64 and assigns it to the DeliveryTime field.
func (o *QuarterlyContractSettlementPriceResponseInner) SetDeliveryTime(v int64) {
	o.DeliveryTime = &v
}

// GetDeliveryPrice returns the DeliveryPrice field value if set, zero value otherwise.
func (o *QuarterlyContractSettlementPriceResponseInner) GetDeliveryPrice() float32 {
	if o == nil || common.IsNil(o.DeliveryPrice) {
		var ret float32
		return ret
	}
	return *o.DeliveryPrice
}

// GetDeliveryPriceOk returns a tuple with the DeliveryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuarterlyContractSettlementPriceResponseInner) GetDeliveryPriceOk() (*float32, bool) {
	if o == nil || common.IsNil(o.DeliveryPrice) {
		return nil, false
	}
	return o.DeliveryPrice, true
}

// HasDeliveryPrice returns a boolean if a field has been set.
func (o *QuarterlyContractSettlementPriceResponseInner) HasDeliveryPrice() bool {
	if o != nil && !common.IsNil(o.DeliveryPrice) {
		return true
	}

	return false
}

// SetDeliveryPrice gets a reference to the given float32 and assigns it to the DeliveryPrice field.
func (o *QuarterlyContractSettlementPriceResponseInner) SetDeliveryPrice(v float32) {
	o.DeliveryPrice = &v
}

func (o QuarterlyContractSettlementPriceResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuarterlyContractSettlementPriceResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DeliveryTime) {
		toSerialize["deliveryTime"] = o.DeliveryTime
	}
	if !common.IsNil(o.DeliveryPrice) {
		toSerialize["deliveryPrice"] = o.DeliveryPrice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuarterlyContractSettlementPriceResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQuarterlyContractSettlementPriceResponseInner := _QuarterlyContractSettlementPriceResponseInner{}

	err = json.Unmarshal(data, &varQuarterlyContractSettlementPriceResponseInner)

	if err != nil {
		return err
	}

	*o = QuarterlyContractSettlementPriceResponseInner(varQuarterlyContractSettlementPriceResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deliveryTime")
		delete(additionalProperties, "deliveryPrice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuarterlyContractSettlementPriceResponseInner struct {
	value *QuarterlyContractSettlementPriceResponseInner
	isSet bool
}

func (v NullableQuarterlyContractSettlementPriceResponseInner) Get() *QuarterlyContractSettlementPriceResponseInner {
	return v.value
}

func (v *NullableQuarterlyContractSettlementPriceResponseInner) Set(val *QuarterlyContractSettlementPriceResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuarterlyContractSettlementPriceResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuarterlyContractSettlementPriceResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuarterlyContractSettlementPriceResponseInner(val *QuarterlyContractSettlementPriceResponseInner) *NullableQuarterlyContractSettlementPriceResponseInner {
	return &NullableQuarterlyContractSettlementPriceResponseInner{value: val, isSet: true}
}

func (v NullableQuarterlyContractSettlementPriceResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuarterlyContractSettlementPriceResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
