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

// checks if the RedeemSolResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RedeemSolResponse{}

// RedeemSolResponse struct for RedeemSolResponse
type RedeemSolResponse struct {
	Success              *bool   `json:"success,omitempty"`
	SolAmount            *string `json:"solAmount,omitempty"`
	ExchangeRate         *string `json:"exchangeRate,omitempty"`
	ArrivalTime          *int64  `json:"arrivalTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RedeemSolResponse RedeemSolResponse

// NewRedeemSolResponse instantiates a new RedeemSolResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedeemSolResponse() *RedeemSolResponse {
	this := RedeemSolResponse{}
	return &this
}

// NewRedeemSolResponseWithDefaults instantiates a new RedeemSolResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedeemSolResponseWithDefaults() *RedeemSolResponse {
	this := RedeemSolResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *RedeemSolResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemSolResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *RedeemSolResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *RedeemSolResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetSolAmount returns the SolAmount field value if set, zero value otherwise.
func (o *RedeemSolResponse) GetSolAmount() string {
	if o == nil || common.IsNil(o.SolAmount) {
		var ret string
		return ret
	}
	return *o.SolAmount
}

// GetSolAmountOk returns a tuple with the SolAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemSolResponse) GetSolAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.SolAmount) {
		return nil, false
	}
	return o.SolAmount, true
}

// HasSolAmount returns a boolean if a field has been set.
func (o *RedeemSolResponse) HasSolAmount() bool {
	if o != nil && !common.IsNil(o.SolAmount) {
		return true
	}

	return false
}

// SetSolAmount gets a reference to the given string and assigns it to the SolAmount field.
func (o *RedeemSolResponse) SetSolAmount(v string) {
	o.SolAmount = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *RedeemSolResponse) GetExchangeRate() string {
	if o == nil || common.IsNil(o.ExchangeRate) {
		var ret string
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemSolResponse) GetExchangeRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *RedeemSolResponse) HasExchangeRate() bool {
	if o != nil && !common.IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given string and assigns it to the ExchangeRate field.
func (o *RedeemSolResponse) SetExchangeRate(v string) {
	o.ExchangeRate = &v
}

// GetArrivalTime returns the ArrivalTime field value if set, zero value otherwise.
func (o *RedeemSolResponse) GetArrivalTime() int64 {
	if o == nil || common.IsNil(o.ArrivalTime) {
		var ret int64
		return ret
	}
	return *o.ArrivalTime
}

// GetArrivalTimeOk returns a tuple with the ArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemSolResponse) GetArrivalTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ArrivalTime) {
		return nil, false
	}
	return o.ArrivalTime, true
}

// HasArrivalTime returns a boolean if a field has been set.
func (o *RedeemSolResponse) HasArrivalTime() bool {
	if o != nil && !common.IsNil(o.ArrivalTime) {
		return true
	}

	return false
}

// SetArrivalTime gets a reference to the given int64 and assigns it to the ArrivalTime field.
func (o *RedeemSolResponse) SetArrivalTime(v int64) {
	o.ArrivalTime = &v
}

func (o RedeemSolResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedeemSolResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !common.IsNil(o.SolAmount) {
		toSerialize["solAmount"] = o.SolAmount
	}
	if !common.IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	if !common.IsNil(o.ArrivalTime) {
		toSerialize["arrivalTime"] = o.ArrivalTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RedeemSolResponse) UnmarshalJSON(data []byte) (err error) {
	varRedeemSolResponse := _RedeemSolResponse{}

	err = json.Unmarshal(data, &varRedeemSolResponse)

	if err != nil {
		return err
	}

	*o = RedeemSolResponse(varRedeemSolResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "solAmount")
		delete(additionalProperties, "exchangeRate")
		delete(additionalProperties, "arrivalTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRedeemSolResponse struct {
	value *RedeemSolResponse
	isSet bool
}

func (v NullableRedeemSolResponse) Get() *RedeemSolResponse {
	return v.value
}

func (v *NullableRedeemSolResponse) Set(val *RedeemSolResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRedeemSolResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRedeemSolResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedeemSolResponse(val *RedeemSolResponse) *NullableRedeemSolResponse {
	return &NullableRedeemSolResponse{value: val, isSet: true}
}

func (v NullableRedeemSolResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedeemSolResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
