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

// checks if the SolStakingAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SolStakingAccountResponse{}

// SolStakingAccountResponse struct for SolStakingAccountResponse
type SolStakingAccountResponse struct {
	BnsolAmount           *string `json:"bnsolAmount,omitempty"`
	HoldingInSOL          *string `json:"holdingInSOL,omitempty"`
	ThirtyDaysProfitInSOL *string `json:"thirtyDaysProfitInSOL,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _SolStakingAccountResponse SolStakingAccountResponse

// NewSolStakingAccountResponse instantiates a new SolStakingAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSolStakingAccountResponse() *SolStakingAccountResponse {
	this := SolStakingAccountResponse{}
	return &this
}

// NewSolStakingAccountResponseWithDefaults instantiates a new SolStakingAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSolStakingAccountResponseWithDefaults() *SolStakingAccountResponse {
	this := SolStakingAccountResponse{}
	return &this
}

// GetBnsolAmount returns the BnsolAmount field value if set, zero value otherwise.
func (o *SolStakingAccountResponse) GetBnsolAmount() string {
	if o == nil || common.IsNil(o.BnsolAmount) {
		var ret string
		return ret
	}
	return *o.BnsolAmount
}

// GetBnsolAmountOk returns a tuple with the BnsolAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolStakingAccountResponse) GetBnsolAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.BnsolAmount) {
		return nil, false
	}
	return o.BnsolAmount, true
}

// HasBnsolAmount returns a boolean if a field has been set.
func (o *SolStakingAccountResponse) HasBnsolAmount() bool {
	if o != nil && !common.IsNil(o.BnsolAmount) {
		return true
	}

	return false
}

// SetBnsolAmount gets a reference to the given string and assigns it to the BnsolAmount field.
func (o *SolStakingAccountResponse) SetBnsolAmount(v string) {
	o.BnsolAmount = &v
}

// GetHoldingInSOL returns the HoldingInSOL field value if set, zero value otherwise.
func (o *SolStakingAccountResponse) GetHoldingInSOL() string {
	if o == nil || common.IsNil(o.HoldingInSOL) {
		var ret string
		return ret
	}
	return *o.HoldingInSOL
}

// GetHoldingInSOLOk returns a tuple with the HoldingInSOL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolStakingAccountResponse) GetHoldingInSOLOk() (*string, bool) {
	if o == nil || common.IsNil(o.HoldingInSOL) {
		return nil, false
	}
	return o.HoldingInSOL, true
}

// HasHoldingInSOL returns a boolean if a field has been set.
func (o *SolStakingAccountResponse) HasHoldingInSOL() bool {
	if o != nil && !common.IsNil(o.HoldingInSOL) {
		return true
	}

	return false
}

// SetHoldingInSOL gets a reference to the given string and assigns it to the HoldingInSOL field.
func (o *SolStakingAccountResponse) SetHoldingInSOL(v string) {
	o.HoldingInSOL = &v
}

// GetThirtyDaysProfitInSOL returns the ThirtyDaysProfitInSOL field value if set, zero value otherwise.
func (o *SolStakingAccountResponse) GetThirtyDaysProfitInSOL() string {
	if o == nil || common.IsNil(o.ThirtyDaysProfitInSOL) {
		var ret string
		return ret
	}
	return *o.ThirtyDaysProfitInSOL
}

// GetThirtyDaysProfitInSOLOk returns a tuple with the ThirtyDaysProfitInSOL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolStakingAccountResponse) GetThirtyDaysProfitInSOLOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThirtyDaysProfitInSOL) {
		return nil, false
	}
	return o.ThirtyDaysProfitInSOL, true
}

// HasThirtyDaysProfitInSOL returns a boolean if a field has been set.
func (o *SolStakingAccountResponse) HasThirtyDaysProfitInSOL() bool {
	if o != nil && !common.IsNil(o.ThirtyDaysProfitInSOL) {
		return true
	}

	return false
}

// SetThirtyDaysProfitInSOL gets a reference to the given string and assigns it to the ThirtyDaysProfitInSOL field.
func (o *SolStakingAccountResponse) SetThirtyDaysProfitInSOL(v string) {
	o.ThirtyDaysProfitInSOL = &v
}

func (o SolStakingAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SolStakingAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BnsolAmount) {
		toSerialize["bnsolAmount"] = o.BnsolAmount
	}
	if !common.IsNil(o.HoldingInSOL) {
		toSerialize["holdingInSOL"] = o.HoldingInSOL
	}
	if !common.IsNil(o.ThirtyDaysProfitInSOL) {
		toSerialize["thirtyDaysProfitInSOL"] = o.ThirtyDaysProfitInSOL
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SolStakingAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varSolStakingAccountResponse := _SolStakingAccountResponse{}

	err = json.Unmarshal(data, &varSolStakingAccountResponse)

	if err != nil {
		return err
	}

	*o = SolStakingAccountResponse(varSolStakingAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bnsolAmount")
		delete(additionalProperties, "holdingInSOL")
		delete(additionalProperties, "thirtyDaysProfitInSOL")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSolStakingAccountResponse struct {
	value *SolStakingAccountResponse
	isSet bool
}

func (v NullableSolStakingAccountResponse) Get() *SolStakingAccountResponse {
	return v.value
}

func (v *NullableSolStakingAccountResponse) Set(val *SolStakingAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSolStakingAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSolStakingAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSolStakingAccountResponse(val *SolStakingAccountResponse) *NullableSolStakingAccountResponse {
	return &NullableSolStakingAccountResponse{value: val, isSet: true}
}

func (v NullableSolStakingAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSolStakingAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
