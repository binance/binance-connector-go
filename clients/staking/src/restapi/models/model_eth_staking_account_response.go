/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the EthStakingAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EthStakingAccountResponse{}

// EthStakingAccountResponse struct for EthStakingAccountResponse
type EthStakingAccountResponse struct {
	HoldingInETH          *string                            `json:"holdingInETH,omitempty"`
	Holdings              *EthStakingAccountResponseHoldings `json:"holdings,omitempty"`
	ThirtyDaysProfitInETH *string                            `json:"thirtyDaysProfitInETH,omitempty"`
	Profit                *EthStakingAccountResponseProfit   `json:"profit,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _EthStakingAccountResponse EthStakingAccountResponse

// NewEthStakingAccountResponse instantiates a new EthStakingAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthStakingAccountResponse() *EthStakingAccountResponse {
	this := EthStakingAccountResponse{}
	return &this
}

// NewEthStakingAccountResponseWithDefaults instantiates a new EthStakingAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthStakingAccountResponseWithDefaults() *EthStakingAccountResponse {
	this := EthStakingAccountResponse{}
	return &this
}

// GetHoldingInETH returns the HoldingInETH field value if set, zero value otherwise.
func (o *EthStakingAccountResponse) GetHoldingInETH() string {
	if o == nil || common.IsNil(o.HoldingInETH) {
		var ret string
		return ret
	}
	return *o.HoldingInETH
}

// GetHoldingInETHOk returns a tuple with the HoldingInETH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthStakingAccountResponse) GetHoldingInETHOk() (*string, bool) {
	if o == nil || common.IsNil(o.HoldingInETH) {
		return nil, false
	}
	return o.HoldingInETH, true
}

// HasHoldingInETH returns a boolean if a field has been set.
func (o *EthStakingAccountResponse) HasHoldingInETH() bool {
	if o != nil && !common.IsNil(o.HoldingInETH) {
		return true
	}

	return false
}

// SetHoldingInETH gets a reference to the given string and assigns it to the HoldingInETH field.
func (o *EthStakingAccountResponse) SetHoldingInETH(v string) {
	o.HoldingInETH = &v
}

// GetHoldings returns the Holdings field value if set, zero value otherwise.
func (o *EthStakingAccountResponse) GetHoldings() EthStakingAccountResponseHoldings {
	if o == nil || common.IsNil(o.Holdings) {
		var ret EthStakingAccountResponseHoldings
		return ret
	}
	return *o.Holdings
}

// GetHoldingsOk returns a tuple with the Holdings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthStakingAccountResponse) GetHoldingsOk() (*EthStakingAccountResponseHoldings, bool) {
	if o == nil || common.IsNil(o.Holdings) {
		return nil, false
	}
	return o.Holdings, true
}

// HasHoldings returns a boolean if a field has been set.
func (o *EthStakingAccountResponse) HasHoldings() bool {
	if o != nil && !common.IsNil(o.Holdings) {
		return true
	}

	return false
}

// SetHoldings gets a reference to the given EthStakingAccountResponseHoldings and assigns it to the Holdings field.
func (o *EthStakingAccountResponse) SetHoldings(v EthStakingAccountResponseHoldings) {
	o.Holdings = &v
}

// GetThirtyDaysProfitInETH returns the ThirtyDaysProfitInETH field value if set, zero value otherwise.
func (o *EthStakingAccountResponse) GetThirtyDaysProfitInETH() string {
	if o == nil || common.IsNil(o.ThirtyDaysProfitInETH) {
		var ret string
		return ret
	}
	return *o.ThirtyDaysProfitInETH
}

// GetThirtyDaysProfitInETHOk returns a tuple with the ThirtyDaysProfitInETH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthStakingAccountResponse) GetThirtyDaysProfitInETHOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThirtyDaysProfitInETH) {
		return nil, false
	}
	return o.ThirtyDaysProfitInETH, true
}

// HasThirtyDaysProfitInETH returns a boolean if a field has been set.
func (o *EthStakingAccountResponse) HasThirtyDaysProfitInETH() bool {
	if o != nil && !common.IsNil(o.ThirtyDaysProfitInETH) {
		return true
	}

	return false
}

// SetThirtyDaysProfitInETH gets a reference to the given string and assigns it to the ThirtyDaysProfitInETH field.
func (o *EthStakingAccountResponse) SetThirtyDaysProfitInETH(v string) {
	o.ThirtyDaysProfitInETH = &v
}

// GetProfit returns the Profit field value if set, zero value otherwise.
func (o *EthStakingAccountResponse) GetProfit() EthStakingAccountResponseProfit {
	if o == nil || common.IsNil(o.Profit) {
		var ret EthStakingAccountResponseProfit
		return ret
	}
	return *o.Profit
}

// GetProfitOk returns a tuple with the Profit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthStakingAccountResponse) GetProfitOk() (*EthStakingAccountResponseProfit, bool) {
	if o == nil || common.IsNil(o.Profit) {
		return nil, false
	}
	return o.Profit, true
}

// HasProfit returns a boolean if a field has been set.
func (o *EthStakingAccountResponse) HasProfit() bool {
	if o != nil && !common.IsNil(o.Profit) {
		return true
	}

	return false
}

// SetProfit gets a reference to the given EthStakingAccountResponseProfit and assigns it to the Profit field.
func (o *EthStakingAccountResponse) SetProfit(v EthStakingAccountResponseProfit) {
	o.Profit = &v
}

func (o EthStakingAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EthStakingAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.HoldingInETH) {
		toSerialize["holdingInETH"] = o.HoldingInETH
	}
	if !common.IsNil(o.Holdings) {
		toSerialize["holdings"] = o.Holdings
	}
	if !common.IsNil(o.ThirtyDaysProfitInETH) {
		toSerialize["thirtyDaysProfitInETH"] = o.ThirtyDaysProfitInETH
	}
	if !common.IsNil(o.Profit) {
		toSerialize["profit"] = o.Profit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EthStakingAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varEthStakingAccountResponse := _EthStakingAccountResponse{}

	err = json.Unmarshal(data, &varEthStakingAccountResponse)

	if err != nil {
		return err
	}

	*o = EthStakingAccountResponse(varEthStakingAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "holdingInETH")
		delete(additionalProperties, "holdings")
		delete(additionalProperties, "thirtyDaysProfitInETH")
		delete(additionalProperties, "profit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEthStakingAccountResponse struct {
	value *EthStakingAccountResponse
	isSet bool
}

func (v NullableEthStakingAccountResponse) Get() *EthStakingAccountResponse {
	return v.value
}

func (v *NullableEthStakingAccountResponse) Set(val *EthStakingAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEthStakingAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEthStakingAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthStakingAccountResponse(val *EthStakingAccountResponse) *NullableEthStakingAccountResponse {
	return &NullableEthStakingAccountResponse{value: val, isSet: true}
}

func (v NullableEthStakingAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthStakingAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
