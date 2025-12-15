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

// checks if the SubscribeEthStakingResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubscribeEthStakingResponse{}

// SubscribeEthStakingResponse struct for SubscribeEthStakingResponse
type SubscribeEthStakingResponse struct {
	Success              *bool   `json:"success,omitempty"`
	WbethAmount          *string `json:"wbethAmount,omitempty"`
	ConversionRatio      *string `json:"conversionRatio,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscribeEthStakingResponse SubscribeEthStakingResponse

// NewSubscribeEthStakingResponse instantiates a new SubscribeEthStakingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribeEthStakingResponse() *SubscribeEthStakingResponse {
	this := SubscribeEthStakingResponse{}
	return &this
}

// NewSubscribeEthStakingResponseWithDefaults instantiates a new SubscribeEthStakingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribeEthStakingResponseWithDefaults() *SubscribeEthStakingResponse {
	this := SubscribeEthStakingResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SubscribeEthStakingResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeEthStakingResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SubscribeEthStakingResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SubscribeEthStakingResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetWbethAmount returns the WbethAmount field value if set, zero value otherwise.
func (o *SubscribeEthStakingResponse) GetWbethAmount() string {
	if o == nil || common.IsNil(o.WbethAmount) {
		var ret string
		return ret
	}
	return *o.WbethAmount
}

// GetWbethAmountOk returns a tuple with the WbethAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeEthStakingResponse) GetWbethAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.WbethAmount) {
		return nil, false
	}
	return o.WbethAmount, true
}

// HasWbethAmount returns a boolean if a field has been set.
func (o *SubscribeEthStakingResponse) HasWbethAmount() bool {
	if o != nil && !common.IsNil(o.WbethAmount) {
		return true
	}

	return false
}

// SetWbethAmount gets a reference to the given string and assigns it to the WbethAmount field.
func (o *SubscribeEthStakingResponse) SetWbethAmount(v string) {
	o.WbethAmount = &v
}

// GetConversionRatio returns the ConversionRatio field value if set, zero value otherwise.
func (o *SubscribeEthStakingResponse) GetConversionRatio() string {
	if o == nil || common.IsNil(o.ConversionRatio) {
		var ret string
		return ret
	}
	return *o.ConversionRatio
}

// GetConversionRatioOk returns a tuple with the ConversionRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeEthStakingResponse) GetConversionRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.ConversionRatio) {
		return nil, false
	}
	return o.ConversionRatio, true
}

// HasConversionRatio returns a boolean if a field has been set.
func (o *SubscribeEthStakingResponse) HasConversionRatio() bool {
	if o != nil && !common.IsNil(o.ConversionRatio) {
		return true
	}

	return false
}

// SetConversionRatio gets a reference to the given string and assigns it to the ConversionRatio field.
func (o *SubscribeEthStakingResponse) SetConversionRatio(v string) {
	o.ConversionRatio = &v
}

func (o SubscribeEthStakingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscribeEthStakingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !common.IsNil(o.WbethAmount) {
		toSerialize["wbethAmount"] = o.WbethAmount
	}
	if !common.IsNil(o.ConversionRatio) {
		toSerialize["conversionRatio"] = o.ConversionRatio
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscribeEthStakingResponse) UnmarshalJSON(data []byte) (err error) {
	varSubscribeEthStakingResponse := _SubscribeEthStakingResponse{}

	err = json.Unmarshal(data, &varSubscribeEthStakingResponse)

	if err != nil {
		return err
	}

	*o = SubscribeEthStakingResponse(varSubscribeEthStakingResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "wbethAmount")
		delete(additionalProperties, "conversionRatio")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscribeEthStakingResponse struct {
	value *SubscribeEthStakingResponse
	isSet bool
}

func (v NullableSubscribeEthStakingResponse) Get() *SubscribeEthStakingResponse {
	return v.value
}

func (v *NullableSubscribeEthStakingResponse) Set(val *SubscribeEthStakingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribeEthStakingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribeEthStakingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribeEthStakingResponse(val *SubscribeEthStakingResponse) *NullableSubscribeEthStakingResponse {
	return &NullableSubscribeEthStakingResponse{value: val, isSet: true}
}

func (v NullableSubscribeEthStakingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribeEthStakingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
