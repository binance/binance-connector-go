/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AvgPriceResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AvgPriceResponseResult{}

// AvgPriceResponseResult struct for AvgPriceResponseResult
type AvgPriceResponseResult struct {
	Mins                 *int64  `json:"mins,omitempty"`
	Price                *string `json:"price,omitempty"`
	CloseTime            *int64  `json:"closeTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AvgPriceResponseResult AvgPriceResponseResult

// NewAvgPriceResponseResult instantiates a new AvgPriceResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvgPriceResponseResult() *AvgPriceResponseResult {
	this := AvgPriceResponseResult{}
	return &this
}

// NewAvgPriceResponseResultWithDefaults instantiates a new AvgPriceResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvgPriceResponseResultWithDefaults() *AvgPriceResponseResult {
	this := AvgPriceResponseResult{}
	return &this
}

// GetMins returns the Mins field value if set, zero value otherwise.
func (o *AvgPriceResponseResult) GetMins() int64 {
	if o == nil || common.IsNil(o.Mins) {
		var ret int64
		return ret
	}
	return *o.Mins
}

// GetMinsOk returns a tuple with the Mins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponseResult) GetMinsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Mins) {
		return nil, false
	}
	return o.Mins, true
}

// HasMins returns a boolean if a field has been set.
func (o *AvgPriceResponseResult) HasMins() bool {
	if o != nil && !common.IsNil(o.Mins) {
		return true
	}

	return false
}

// SetMins gets a reference to the given int64 and assigns it to the Mins field.
func (o *AvgPriceResponseResult) SetMins(v int64) {
	o.Mins = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *AvgPriceResponseResult) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponseResult) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *AvgPriceResponseResult) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *AvgPriceResponseResult) SetPrice(v string) {
	o.Price = &v
}

// GetCloseTime returns the CloseTime field value if set, zero value otherwise.
func (o *AvgPriceResponseResult) GetCloseTime() int64 {
	if o == nil || common.IsNil(o.CloseTime) {
		var ret int64
		return ret
	}
	return *o.CloseTime
}

// GetCloseTimeOk returns a tuple with the CloseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponseResult) GetCloseTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CloseTime) {
		return nil, false
	}
	return o.CloseTime, true
}

// HasCloseTime returns a boolean if a field has been set.
func (o *AvgPriceResponseResult) HasCloseTime() bool {
	if o != nil && !common.IsNil(o.CloseTime) {
		return true
	}

	return false
}

// SetCloseTime gets a reference to the given int64 and assigns it to the CloseTime field.
func (o *AvgPriceResponseResult) SetCloseTime(v int64) {
	o.CloseTime = &v
}

func (o AvgPriceResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvgPriceResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Mins) {
		toSerialize["mins"] = o.Mins
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.CloseTime) {
		toSerialize["closeTime"] = o.CloseTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AvgPriceResponseResult) UnmarshalJSON(data []byte) (err error) {
	varAvgPriceResponseResult := _AvgPriceResponseResult{}

	err = json.Unmarshal(data, &varAvgPriceResponseResult)

	if err != nil {
		return err
	}

	*o = AvgPriceResponseResult(varAvgPriceResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mins")
		delete(additionalProperties, "price")
		delete(additionalProperties, "closeTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAvgPriceResponseResult struct {
	value *AvgPriceResponseResult
	isSet bool
}

func (v NullableAvgPriceResponseResult) Get() *AvgPriceResponseResult {
	return v.value
}

func (v *NullableAvgPriceResponseResult) Set(val *AvgPriceResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAvgPriceResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAvgPriceResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvgPriceResponseResult(val *AvgPriceResponseResult) *NullableAvgPriceResponseResult {
	return &NullableAvgPriceResponseResult{value: val, isSet: true}
}

func (v NullableAvgPriceResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvgPriceResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
