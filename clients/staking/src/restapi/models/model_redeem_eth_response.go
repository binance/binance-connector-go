/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the RedeemEthResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RedeemEthResponse{}

// RedeemEthResponse struct for RedeemEthResponse
type RedeemEthResponse struct {
	Success              *bool   `json:"success,omitempty"`
	EthAmount            *string `json:"ethAmount,omitempty"`
	ConversionRatio      *string `json:"conversionRatio,omitempty"`
	ArrivalTime          *int64  `json:"arrivalTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RedeemEthResponse RedeemEthResponse

// NewRedeemEthResponse instantiates a new RedeemEthResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedeemEthResponse() *RedeemEthResponse {
	this := RedeemEthResponse{}
	return &this
}

// NewRedeemEthResponseWithDefaults instantiates a new RedeemEthResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedeemEthResponseWithDefaults() *RedeemEthResponse {
	this := RedeemEthResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *RedeemEthResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemEthResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *RedeemEthResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *RedeemEthResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetEthAmount returns the EthAmount field value if set, zero value otherwise.
func (o *RedeemEthResponse) GetEthAmount() string {
	if o == nil || common.IsNil(o.EthAmount) {
		var ret string
		return ret
	}
	return *o.EthAmount
}

// GetEthAmountOk returns a tuple with the EthAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemEthResponse) GetEthAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.EthAmount) {
		return nil, false
	}
	return o.EthAmount, true
}

// HasEthAmount returns a boolean if a field has been set.
func (o *RedeemEthResponse) HasEthAmount() bool {
	if o != nil && !common.IsNil(o.EthAmount) {
		return true
	}

	return false
}

// SetEthAmount gets a reference to the given string and assigns it to the EthAmount field.
func (o *RedeemEthResponse) SetEthAmount(v string) {
	o.EthAmount = &v
}

// GetConversionRatio returns the ConversionRatio field value if set, zero value otherwise.
func (o *RedeemEthResponse) GetConversionRatio() string {
	if o == nil || common.IsNil(o.ConversionRatio) {
		var ret string
		return ret
	}
	return *o.ConversionRatio
}

// GetConversionRatioOk returns a tuple with the ConversionRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemEthResponse) GetConversionRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.ConversionRatio) {
		return nil, false
	}
	return o.ConversionRatio, true
}

// HasConversionRatio returns a boolean if a field has been set.
func (o *RedeemEthResponse) HasConversionRatio() bool {
	if o != nil && !common.IsNil(o.ConversionRatio) {
		return true
	}

	return false
}

// SetConversionRatio gets a reference to the given string and assigns it to the ConversionRatio field.
func (o *RedeemEthResponse) SetConversionRatio(v string) {
	o.ConversionRatio = &v
}

// GetArrivalTime returns the ArrivalTime field value if set, zero value otherwise.
func (o *RedeemEthResponse) GetArrivalTime() int64 {
	if o == nil || common.IsNil(o.ArrivalTime) {
		var ret int64
		return ret
	}
	return *o.ArrivalTime
}

// GetArrivalTimeOk returns a tuple with the ArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemEthResponse) GetArrivalTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ArrivalTime) {
		return nil, false
	}
	return o.ArrivalTime, true
}

// HasArrivalTime returns a boolean if a field has been set.
func (o *RedeemEthResponse) HasArrivalTime() bool {
	if o != nil && !common.IsNil(o.ArrivalTime) {
		return true
	}

	return false
}

// SetArrivalTime gets a reference to the given int64 and assigns it to the ArrivalTime field.
func (o *RedeemEthResponse) SetArrivalTime(v int64) {
	o.ArrivalTime = &v
}

func (o RedeemEthResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedeemEthResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !common.IsNil(o.EthAmount) {
		toSerialize["ethAmount"] = o.EthAmount
	}
	if !common.IsNil(o.ConversionRatio) {
		toSerialize["conversionRatio"] = o.ConversionRatio
	}
	if !common.IsNil(o.ArrivalTime) {
		toSerialize["arrivalTime"] = o.ArrivalTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RedeemEthResponse) UnmarshalJSON(data []byte) (err error) {
	varRedeemEthResponse := _RedeemEthResponse{}

	err = json.Unmarshal(data, &varRedeemEthResponse)

	if err != nil {
		return err
	}

	*o = RedeemEthResponse(varRedeemEthResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "ethAmount")
		delete(additionalProperties, "conversionRatio")
		delete(additionalProperties, "arrivalTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRedeemEthResponse struct {
	value *RedeemEthResponse
	isSet bool
}

func (v NullableRedeemEthResponse) Get() *RedeemEthResponse {
	return v.value
}

func (v *NullableRedeemEthResponse) Set(val *RedeemEthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRedeemEthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRedeemEthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedeemEthResponse(val *RedeemEthResponse) *NullableRedeemEthResponse {
	return &NullableRedeemEthResponse{value: val, isSet: true}
}

func (v NullableRedeemEthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedeemEthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
