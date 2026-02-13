/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the WrapBethResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &WrapBethResponse{}

// WrapBethResponse struct for WrapBethResponse
type WrapBethResponse struct {
	Success              *bool   `json:"success,omitempty"`
	WbethAmount          *string `json:"wbethAmount,omitempty"`
	ExchangeRate         *string `json:"exchangeRate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WrapBethResponse WrapBethResponse

// NewWrapBethResponse instantiates a new WrapBethResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWrapBethResponse() *WrapBethResponse {
	this := WrapBethResponse{}
	return &this
}

// NewWrapBethResponseWithDefaults instantiates a new WrapBethResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWrapBethResponseWithDefaults() *WrapBethResponse {
	this := WrapBethResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *WrapBethResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WrapBethResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *WrapBethResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *WrapBethResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetWbethAmount returns the WbethAmount field value if set, zero value otherwise.
func (o *WrapBethResponse) GetWbethAmount() string {
	if o == nil || common.IsNil(o.WbethAmount) {
		var ret string
		return ret
	}
	return *o.WbethAmount
}

// GetWbethAmountOk returns a tuple with the WbethAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WrapBethResponse) GetWbethAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.WbethAmount) {
		return nil, false
	}
	return o.WbethAmount, true
}

// HasWbethAmount returns a boolean if a field has been set.
func (o *WrapBethResponse) HasWbethAmount() bool {
	if o != nil && !common.IsNil(o.WbethAmount) {
		return true
	}

	return false
}

// SetWbethAmount gets a reference to the given string and assigns it to the WbethAmount field.
func (o *WrapBethResponse) SetWbethAmount(v string) {
	o.WbethAmount = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *WrapBethResponse) GetExchangeRate() string {
	if o == nil || common.IsNil(o.ExchangeRate) {
		var ret string
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WrapBethResponse) GetExchangeRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *WrapBethResponse) HasExchangeRate() bool {
	if o != nil && !common.IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given string and assigns it to the ExchangeRate field.
func (o *WrapBethResponse) SetExchangeRate(v string) {
	o.ExchangeRate = &v
}

func (o WrapBethResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WrapBethResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !common.IsNil(o.WbethAmount) {
		toSerialize["wbethAmount"] = o.WbethAmount
	}
	if !common.IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WrapBethResponse) UnmarshalJSON(data []byte) (err error) {
	varWrapBethResponse := _WrapBethResponse{}

	err = json.Unmarshal(data, &varWrapBethResponse)

	if err != nil {
		return err
	}

	*o = WrapBethResponse(varWrapBethResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "wbethAmount")
		delete(additionalProperties, "exchangeRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWrapBethResponse struct {
	value *WrapBethResponse
	isSet bool
}

func (v NullableWrapBethResponse) Get() *WrapBethResponse {
	return v.value
}

func (v *NullableWrapBethResponse) Set(val *WrapBethResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWrapBethResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWrapBethResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWrapBethResponse(val *WrapBethResponse) *NullableWrapBethResponse {
	return &NullableWrapBethResponse{value: val, isSet: true}
}

func (v NullableWrapBethResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWrapBethResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
