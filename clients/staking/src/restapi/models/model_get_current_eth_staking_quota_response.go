/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetCurrentEthStakingQuotaResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCurrentEthStakingQuotaResponse{}

// GetCurrentEthStakingQuotaResponse struct for GetCurrentEthStakingQuotaResponse
type GetCurrentEthStakingQuotaResponse struct {
	LeftStakingPersonalQuota    *string `json:"leftStakingPersonalQuota,omitempty"`
	LeftRedemptionPersonalQuota *string `json:"leftRedemptionPersonalQuota,omitempty"`
	MinStakeAmount              *string `json:"minStakeAmount,omitempty"`
	MinRedeemAmount             *string `json:"minRedeemAmount,omitempty"`
	RedeemPeriod                *int64  `json:"redeemPeriod,omitempty"`
	Stakeable                   *bool   `json:"stakeable,omitempty"`
	Redeemable                  *bool   `json:"redeemable,omitempty"`
	CommissionFee               *string `json:"commissionFee,omitempty"`
	Calculating                 *bool   `json:"calculating,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _GetCurrentEthStakingQuotaResponse GetCurrentEthStakingQuotaResponse

// NewGetCurrentEthStakingQuotaResponse instantiates a new GetCurrentEthStakingQuotaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCurrentEthStakingQuotaResponse() *GetCurrentEthStakingQuotaResponse {
	this := GetCurrentEthStakingQuotaResponse{}
	return &this
}

// NewGetCurrentEthStakingQuotaResponseWithDefaults instantiates a new GetCurrentEthStakingQuotaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCurrentEthStakingQuotaResponseWithDefaults() *GetCurrentEthStakingQuotaResponse {
	this := GetCurrentEthStakingQuotaResponse{}
	return &this
}

// GetLeftStakingPersonalQuota returns the LeftStakingPersonalQuota field value if set, zero value otherwise.
func (o *GetCurrentEthStakingQuotaResponse) GetLeftStakingPersonalQuota() string {
	if o == nil || common.IsNil(o.LeftStakingPersonalQuota) {
		var ret string
		return ret
	}
	return *o.LeftStakingPersonalQuota
}

// GetLeftStakingPersonalQuotaOk returns a tuple with the LeftStakingPersonalQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentEthStakingQuotaResponse) GetLeftStakingPersonalQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.LeftStakingPersonalQuota) {
		return nil, false
	}
	return o.LeftStakingPersonalQuota, true
}

// HasLeftStakingPersonalQuota returns a boolean if a field has been set.
func (o *GetCurrentEthStakingQuotaResponse) HasLeftStakingPersonalQuota() bool {
	if o != nil && !common.IsNil(o.LeftStakingPersonalQuota) {
		return true
	}

	return false
}

// SetLeftStakingPersonalQuota gets a reference to the given string and assigns it to the LeftStakingPersonalQuota field.
func (o *GetCurrentEthStakingQuotaResponse) SetLeftStakingPersonalQuota(v string) {
	o.LeftStakingPersonalQuota = &v
}

// GetLeftRedemptionPersonalQuota returns the LeftRedemptionPersonalQuota field value if set, zero value otherwise.
func (o *GetCurrentEthStakingQuotaResponse) GetLeftRedemptionPersonalQuota() string {
	if o == nil || common.IsNil(o.LeftRedemptionPersonalQuota) {
		var ret string
		return ret
	}
	return *o.LeftRedemptionPersonalQuota
}

// GetLeftRedemptionPersonalQuotaOk returns a tuple with the LeftRedemptionPersonalQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentEthStakingQuotaResponse) GetLeftRedemptionPersonalQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.LeftRedemptionPersonalQuota) {
		return nil, false
	}
	return o.LeftRedemptionPersonalQuota, true
}

// HasLeftRedemptionPersonalQuota returns a boolean if a field has been set.
func (o *GetCurrentEthStakingQuotaResponse) HasLeftRedemptionPersonalQuota() bool {
	if o != nil && !common.IsNil(o.LeftRedemptionPersonalQuota) {
		return true
	}

	return false
}

// SetLeftRedemptionPersonalQuota gets a reference to the given string and assigns it to the LeftRedemptionPersonalQuota field.
func (o *GetCurrentEthStakingQuotaResponse) SetLeftRedemptionPersonalQuota(v string) {
	o.LeftRedemptionPersonalQuota = &v
}

// GetMinStakeAmount returns the MinStakeAmount field value if set, zero value otherwise.
func (o *GetCurrentEthStakingQuotaResponse) GetMinStakeAmount() string {
	if o == nil || common.IsNil(o.MinStakeAmount) {
		var ret string
		return ret
	}
	return *o.MinStakeAmount
}

// GetMinStakeAmountOk returns a tuple with the MinStakeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentEthStakingQuotaResponse) GetMinStakeAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinStakeAmount) {
		return nil, false
	}
	return o.MinStakeAmount, true
}

// HasMinStakeAmount returns a boolean if a field has been set.
func (o *GetCurrentEthStakingQuotaResponse) HasMinStakeAmount() bool {
	if o != nil && !common.IsNil(o.MinStakeAmount) {
		return true
	}

	return false
}

// SetMinStakeAmount gets a reference to the given string and assigns it to the MinStakeAmount field.
func (o *GetCurrentEthStakingQuotaResponse) SetMinStakeAmount(v string) {
	o.MinStakeAmount = &v
}

// GetMinRedeemAmount returns the MinRedeemAmount field value if set, zero value otherwise.
func (o *GetCurrentEthStakingQuotaResponse) GetMinRedeemAmount() string {
	if o == nil || common.IsNil(o.MinRedeemAmount) {
		var ret string
		return ret
	}
	return *o.MinRedeemAmount
}

// GetMinRedeemAmountOk returns a tuple with the MinRedeemAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentEthStakingQuotaResponse) GetMinRedeemAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinRedeemAmount) {
		return nil, false
	}
	return o.MinRedeemAmount, true
}

// HasMinRedeemAmount returns a boolean if a field has been set.
func (o *GetCurrentEthStakingQuotaResponse) HasMinRedeemAmount() bool {
	if o != nil && !common.IsNil(o.MinRedeemAmount) {
		return true
	}

	return false
}

// SetMinRedeemAmount gets a reference to the given string and assigns it to the MinRedeemAmount field.
func (o *GetCurrentEthStakingQuotaResponse) SetMinRedeemAmount(v string) {
	o.MinRedeemAmount = &v
}

// GetRedeemPeriod returns the RedeemPeriod field value if set, zero value otherwise.
func (o *GetCurrentEthStakingQuotaResponse) GetRedeemPeriod() int64 {
	if o == nil || common.IsNil(o.RedeemPeriod) {
		var ret int64
		return ret
	}
	return *o.RedeemPeriod
}

// GetRedeemPeriodOk returns a tuple with the RedeemPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentEthStakingQuotaResponse) GetRedeemPeriodOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RedeemPeriod) {
		return nil, false
	}
	return o.RedeemPeriod, true
}

// HasRedeemPeriod returns a boolean if a field has been set.
func (o *GetCurrentEthStakingQuotaResponse) HasRedeemPeriod() bool {
	if o != nil && !common.IsNil(o.RedeemPeriod) {
		return true
	}

	return false
}

// SetRedeemPeriod gets a reference to the given int64 and assigns it to the RedeemPeriod field.
func (o *GetCurrentEthStakingQuotaResponse) SetRedeemPeriod(v int64) {
	o.RedeemPeriod = &v
}

// GetStakeable returns the Stakeable field value if set, zero value otherwise.
func (o *GetCurrentEthStakingQuotaResponse) GetStakeable() bool {
	if o == nil || common.IsNil(o.Stakeable) {
		var ret bool
		return ret
	}
	return *o.Stakeable
}

// GetStakeableOk returns a tuple with the Stakeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentEthStakingQuotaResponse) GetStakeableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Stakeable) {
		return nil, false
	}
	return o.Stakeable, true
}

// HasStakeable returns a boolean if a field has been set.
func (o *GetCurrentEthStakingQuotaResponse) HasStakeable() bool {
	if o != nil && !common.IsNil(o.Stakeable) {
		return true
	}

	return false
}

// SetStakeable gets a reference to the given bool and assigns it to the Stakeable field.
func (o *GetCurrentEthStakingQuotaResponse) SetStakeable(v bool) {
	o.Stakeable = &v
}

// GetRedeemable returns the Redeemable field value if set, zero value otherwise.
func (o *GetCurrentEthStakingQuotaResponse) GetRedeemable() bool {
	if o == nil || common.IsNil(o.Redeemable) {
		var ret bool
		return ret
	}
	return *o.Redeemable
}

// GetRedeemableOk returns a tuple with the Redeemable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentEthStakingQuotaResponse) GetRedeemableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Redeemable) {
		return nil, false
	}
	return o.Redeemable, true
}

// HasRedeemable returns a boolean if a field has been set.
func (o *GetCurrentEthStakingQuotaResponse) HasRedeemable() bool {
	if o != nil && !common.IsNil(o.Redeemable) {
		return true
	}

	return false
}

// SetRedeemable gets a reference to the given bool and assigns it to the Redeemable field.
func (o *GetCurrentEthStakingQuotaResponse) SetRedeemable(v bool) {
	o.Redeemable = &v
}

// GetCommissionFee returns the CommissionFee field value if set, zero value otherwise.
func (o *GetCurrentEthStakingQuotaResponse) GetCommissionFee() string {
	if o == nil || common.IsNil(o.CommissionFee) {
		var ret string
		return ret
	}
	return *o.CommissionFee
}

// GetCommissionFeeOk returns a tuple with the CommissionFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentEthStakingQuotaResponse) GetCommissionFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CommissionFee) {
		return nil, false
	}
	return o.CommissionFee, true
}

// HasCommissionFee returns a boolean if a field has been set.
func (o *GetCurrentEthStakingQuotaResponse) HasCommissionFee() bool {
	if o != nil && !common.IsNil(o.CommissionFee) {
		return true
	}

	return false
}

// SetCommissionFee gets a reference to the given string and assigns it to the CommissionFee field.
func (o *GetCurrentEthStakingQuotaResponse) SetCommissionFee(v string) {
	o.CommissionFee = &v
}

// GetCalculating returns the Calculating field value if set, zero value otherwise.
func (o *GetCurrentEthStakingQuotaResponse) GetCalculating() bool {
	if o == nil || common.IsNil(o.Calculating) {
		var ret bool
		return ret
	}
	return *o.Calculating
}

// GetCalculatingOk returns a tuple with the Calculating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentEthStakingQuotaResponse) GetCalculatingOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Calculating) {
		return nil, false
	}
	return o.Calculating, true
}

// HasCalculating returns a boolean if a field has been set.
func (o *GetCurrentEthStakingQuotaResponse) HasCalculating() bool {
	if o != nil && !common.IsNil(o.Calculating) {
		return true
	}

	return false
}

// SetCalculating gets a reference to the given bool and assigns it to the Calculating field.
func (o *GetCurrentEthStakingQuotaResponse) SetCalculating(v bool) {
	o.Calculating = &v
}

func (o GetCurrentEthStakingQuotaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCurrentEthStakingQuotaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LeftStakingPersonalQuota) {
		toSerialize["leftStakingPersonalQuota"] = o.LeftStakingPersonalQuota
	}
	if !common.IsNil(o.LeftRedemptionPersonalQuota) {
		toSerialize["leftRedemptionPersonalQuota"] = o.LeftRedemptionPersonalQuota
	}
	if !common.IsNil(o.MinStakeAmount) {
		toSerialize["minStakeAmount"] = o.MinStakeAmount
	}
	if !common.IsNil(o.MinRedeemAmount) {
		toSerialize["minRedeemAmount"] = o.MinRedeemAmount
	}
	if !common.IsNil(o.RedeemPeriod) {
		toSerialize["redeemPeriod"] = o.RedeemPeriod
	}
	if !common.IsNil(o.Stakeable) {
		toSerialize["stakeable"] = o.Stakeable
	}
	if !common.IsNil(o.Redeemable) {
		toSerialize["redeemable"] = o.Redeemable
	}
	if !common.IsNil(o.CommissionFee) {
		toSerialize["commissionFee"] = o.CommissionFee
	}
	if !common.IsNil(o.Calculating) {
		toSerialize["calculating"] = o.Calculating
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCurrentEthStakingQuotaResponse) UnmarshalJSON(data []byte) (err error) {
	varGetCurrentEthStakingQuotaResponse := _GetCurrentEthStakingQuotaResponse{}

	err = json.Unmarshal(data, &varGetCurrentEthStakingQuotaResponse)

	if err != nil {
		return err
	}

	*o = GetCurrentEthStakingQuotaResponse(varGetCurrentEthStakingQuotaResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "leftStakingPersonalQuota")
		delete(additionalProperties, "leftRedemptionPersonalQuota")
		delete(additionalProperties, "minStakeAmount")
		delete(additionalProperties, "minRedeemAmount")
		delete(additionalProperties, "redeemPeriod")
		delete(additionalProperties, "stakeable")
		delete(additionalProperties, "redeemable")
		delete(additionalProperties, "commissionFee")
		delete(additionalProperties, "calculating")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCurrentEthStakingQuotaResponse struct {
	value *GetCurrentEthStakingQuotaResponse
	isSet bool
}

func (v NullableGetCurrentEthStakingQuotaResponse) Get() *GetCurrentEthStakingQuotaResponse {
	return v.value
}

func (v *NullableGetCurrentEthStakingQuotaResponse) Set(val *GetCurrentEthStakingQuotaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCurrentEthStakingQuotaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCurrentEthStakingQuotaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCurrentEthStakingQuotaResponse(val *GetCurrentEthStakingQuotaResponse) *NullableGetCurrentEthStakingQuotaResponse {
	return &NullableGetCurrentEthStakingQuotaResponse{value: val, isSet: true}
}

func (v NullableGetCurrentEthStakingQuotaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCurrentEthStakingQuotaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
