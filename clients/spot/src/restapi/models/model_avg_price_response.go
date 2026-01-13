/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AvgPriceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AvgPriceResponse{}

// AvgPriceResponse struct for AvgPriceResponse
type AvgPriceResponse struct {
	Mins                 *int64  `json:"mins,omitempty"`
	Price                *string `json:"price,omitempty"`
	CloseTime            *int64  `json:"closeTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AvgPriceResponse AvgPriceResponse

// NewAvgPriceResponse instantiates a new AvgPriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvgPriceResponse() *AvgPriceResponse {
	this := AvgPriceResponse{}
	return &this
}

// NewAvgPriceResponseWithDefaults instantiates a new AvgPriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvgPriceResponseWithDefaults() *AvgPriceResponse {
	this := AvgPriceResponse{}
	return &this
}

// GetMins returns the Mins field value if set, zero value otherwise.
func (o *AvgPriceResponse) GetMins() int64 {
	if o == nil || common.IsNil(o.Mins) {
		var ret int64
		return ret
	}
	return *o.Mins
}

// GetMinsOk returns a tuple with the Mins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponse) GetMinsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Mins) {
		return nil, false
	}
	return o.Mins, true
}

// HasMins returns a boolean if a field has been set.
func (o *AvgPriceResponse) HasMins() bool {
	if o != nil && !common.IsNil(o.Mins) {
		return true
	}

	return false
}

// SetMins gets a reference to the given int64 and assigns it to the Mins field.
func (o *AvgPriceResponse) SetMins(v int64) {
	o.Mins = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *AvgPriceResponse) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponse) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *AvgPriceResponse) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *AvgPriceResponse) SetPrice(v string) {
	o.Price = &v
}

// GetCloseTime returns the CloseTime field value if set, zero value otherwise.
func (o *AvgPriceResponse) GetCloseTime() int64 {
	if o == nil || common.IsNil(o.CloseTime) {
		var ret int64
		return ret
	}
	return *o.CloseTime
}

// GetCloseTimeOk returns a tuple with the CloseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvgPriceResponse) GetCloseTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CloseTime) {
		return nil, false
	}
	return o.CloseTime, true
}

// HasCloseTime returns a boolean if a field has been set.
func (o *AvgPriceResponse) HasCloseTime() bool {
	if o != nil && !common.IsNil(o.CloseTime) {
		return true
	}

	return false
}

// SetCloseTime gets a reference to the given int64 and assigns it to the CloseTime field.
func (o *AvgPriceResponse) SetCloseTime(v int64) {
	o.CloseTime = &v
}

func (o AvgPriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvgPriceResponse) ToMap() (map[string]interface{}, error) {
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

func (o *AvgPriceResponse) UnmarshalJSON(data []byte) (err error) {
	varAvgPriceResponse := _AvgPriceResponse{}

	err = json.Unmarshal(data, &varAvgPriceResponse)

	if err != nil {
		return err
	}

	*o = AvgPriceResponse(varAvgPriceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mins")
		delete(additionalProperties, "price")
		delete(additionalProperties, "closeTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAvgPriceResponse struct {
	value *AvgPriceResponse
	isSet bool
}

func (v NullableAvgPriceResponse) Get() *AvgPriceResponse {
	return v.value
}

func (v *NullableAvgPriceResponse) Set(val *AvgPriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAvgPriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAvgPriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvgPriceResponse(val *AvgPriceResponse) *NullableAvgPriceResponse {
	return &NullableAvgPriceResponse{value: val, isSet: true}
}

func (v NullableAvgPriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvgPriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
