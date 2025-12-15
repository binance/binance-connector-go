/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SubscribeSolStakingResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubscribeSolStakingResponse{}

// SubscribeSolStakingResponse struct for SubscribeSolStakingResponse
type SubscribeSolStakingResponse struct {
	Success              *bool   `json:"success,omitempty"`
	BnsolAmount          *string `json:"bnsolAmount,omitempty"`
	ExchangeRate         *string `json:"exchangeRate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscribeSolStakingResponse SubscribeSolStakingResponse

// NewSubscribeSolStakingResponse instantiates a new SubscribeSolStakingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribeSolStakingResponse() *SubscribeSolStakingResponse {
	this := SubscribeSolStakingResponse{}
	return &this
}

// NewSubscribeSolStakingResponseWithDefaults instantiates a new SubscribeSolStakingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribeSolStakingResponseWithDefaults() *SubscribeSolStakingResponse {
	this := SubscribeSolStakingResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SubscribeSolStakingResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeSolStakingResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SubscribeSolStakingResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SubscribeSolStakingResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetBnsolAmount returns the BnsolAmount field value if set, zero value otherwise.
func (o *SubscribeSolStakingResponse) GetBnsolAmount() string {
	if o == nil || common.IsNil(o.BnsolAmount) {
		var ret string
		return ret
	}
	return *o.BnsolAmount
}

// GetBnsolAmountOk returns a tuple with the BnsolAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeSolStakingResponse) GetBnsolAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.BnsolAmount) {
		return nil, false
	}
	return o.BnsolAmount, true
}

// HasBnsolAmount returns a boolean if a field has been set.
func (o *SubscribeSolStakingResponse) HasBnsolAmount() bool {
	if o != nil && !common.IsNil(o.BnsolAmount) {
		return true
	}

	return false
}

// SetBnsolAmount gets a reference to the given string and assigns it to the BnsolAmount field.
func (o *SubscribeSolStakingResponse) SetBnsolAmount(v string) {
	o.BnsolAmount = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *SubscribeSolStakingResponse) GetExchangeRate() string {
	if o == nil || common.IsNil(o.ExchangeRate) {
		var ret string
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeSolStakingResponse) GetExchangeRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *SubscribeSolStakingResponse) HasExchangeRate() bool {
	if o != nil && !common.IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given string and assigns it to the ExchangeRate field.
func (o *SubscribeSolStakingResponse) SetExchangeRate(v string) {
	o.ExchangeRate = &v
}

func (o SubscribeSolStakingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscribeSolStakingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !common.IsNil(o.BnsolAmount) {
		toSerialize["bnsolAmount"] = o.BnsolAmount
	}
	if !common.IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscribeSolStakingResponse) UnmarshalJSON(data []byte) (err error) {
	varSubscribeSolStakingResponse := _SubscribeSolStakingResponse{}

	err = json.Unmarshal(data, &varSubscribeSolStakingResponse)

	if err != nil {
		return err
	}

	*o = SubscribeSolStakingResponse(varSubscribeSolStakingResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "bnsolAmount")
		delete(additionalProperties, "exchangeRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscribeSolStakingResponse struct {
	value *SubscribeSolStakingResponse
	isSet bool
}

func (v NullableSubscribeSolStakingResponse) Get() *SubscribeSolStakingResponse {
	return v.value
}

func (v *NullableSubscribeSolStakingResponse) Set(val *SubscribeSolStakingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribeSolStakingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribeSolStakingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribeSolStakingResponse(val *SubscribeSolStakingResponse) *NullableSubscribeSolStakingResponse {
	return &NullableSubscribeSolStakingResponse{value: val, isSet: true}
}

func (v NullableSubscribeSolStakingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribeSolStakingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
