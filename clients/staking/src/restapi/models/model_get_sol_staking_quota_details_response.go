/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSolStakingQuotaDetailsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSolStakingQuotaDetailsResponse{}

// GetSolStakingQuotaDetailsResponse struct for GetSolStakingQuotaDetailsResponse
type GetSolStakingQuotaDetailsResponse struct {
	LeftStakingPersonalQuota    *string `json:"leftStakingPersonalQuota,omitempty"`
	LeftRedemptionPersonalQuota *string `json:"leftRedemptionPersonalQuota,omitempty"`
	MinStakeAmount              *string `json:"minStakeAmount,omitempty"`
	MinRedeemAmount             *string `json:"minRedeemAmount,omitempty"`
	RedeemPeriod                *int64  `json:"redeemPeriod,omitempty"`
	Stakeable                   *bool   `json:"stakeable,omitempty"`
	Redeemable                  *bool   `json:"redeemable,omitempty"`
	SoldOut                     *bool   `json:"soldOut,omitempty"`
	CommissionFee               *string `json:"commissionFee,omitempty"`
	NextEpochTime               *int64  `json:"nextEpochTime,omitempty"`
	Calculating                 *bool   `json:"calculating,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _GetSolStakingQuotaDetailsResponse GetSolStakingQuotaDetailsResponse

// NewGetSolStakingQuotaDetailsResponse instantiates a new GetSolStakingQuotaDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolStakingQuotaDetailsResponse() *GetSolStakingQuotaDetailsResponse {
	this := GetSolStakingQuotaDetailsResponse{}
	return &this
}

// NewGetSolStakingQuotaDetailsResponseWithDefaults instantiates a new GetSolStakingQuotaDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolStakingQuotaDetailsResponseWithDefaults() *GetSolStakingQuotaDetailsResponse {
	this := GetSolStakingQuotaDetailsResponse{}
	return &this
}

// GetLeftStakingPersonalQuota returns the LeftStakingPersonalQuota field value if set, zero value otherwise.
func (o *GetSolStakingQuotaDetailsResponse) GetLeftStakingPersonalQuota() string {
	if o == nil || common.IsNil(o.LeftStakingPersonalQuota) {
		var ret string
		return ret
	}
	return *o.LeftStakingPersonalQuota
}

// GetLeftStakingPersonalQuotaOk returns a tuple with the LeftStakingPersonalQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingQuotaDetailsResponse) GetLeftStakingPersonalQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.LeftStakingPersonalQuota) {
		return nil, false
	}
	return o.LeftStakingPersonalQuota, true
}

// HasLeftStakingPersonalQuota returns a boolean if a field has been set.
func (o *GetSolStakingQuotaDetailsResponse) HasLeftStakingPersonalQuota() bool {
	if o != nil && !common.IsNil(o.LeftStakingPersonalQuota) {
		return true
	}

	return false
}

// SetLeftStakingPersonalQuota gets a reference to the given string and assigns it to the LeftStakingPersonalQuota field.
func (o *GetSolStakingQuotaDetailsResponse) SetLeftStakingPersonalQuota(v string) {
	o.LeftStakingPersonalQuota = &v
}

// GetLeftRedemptionPersonalQuota returns the LeftRedemptionPersonalQuota field value if set, zero value otherwise.
func (o *GetSolStakingQuotaDetailsResponse) GetLeftRedemptionPersonalQuota() string {
	if o == nil || common.IsNil(o.LeftRedemptionPersonalQuota) {
		var ret string
		return ret
	}
	return *o.LeftRedemptionPersonalQuota
}

// GetLeftRedemptionPersonalQuotaOk returns a tuple with the LeftRedemptionPersonalQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingQuotaDetailsResponse) GetLeftRedemptionPersonalQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.LeftRedemptionPersonalQuota) {
		return nil, false
	}
	return o.LeftRedemptionPersonalQuota, true
}

// HasLeftRedemptionPersonalQuota returns a boolean if a field has been set.
func (o *GetSolStakingQuotaDetailsResponse) HasLeftRedemptionPersonalQuota() bool {
	if o != nil && !common.IsNil(o.LeftRedemptionPersonalQuota) {
		return true
	}

	return false
}

// SetLeftRedemptionPersonalQuota gets a reference to the given string and assigns it to the LeftRedemptionPersonalQuota field.
func (o *GetSolStakingQuotaDetailsResponse) SetLeftRedemptionPersonalQuota(v string) {
	o.LeftRedemptionPersonalQuota = &v
}

// GetMinStakeAmount returns the MinStakeAmount field value if set, zero value otherwise.
func (o *GetSolStakingQuotaDetailsResponse) GetMinStakeAmount() string {
	if o == nil || common.IsNil(o.MinStakeAmount) {
		var ret string
		return ret
	}
	return *o.MinStakeAmount
}

// GetMinStakeAmountOk returns a tuple with the MinStakeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingQuotaDetailsResponse) GetMinStakeAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinStakeAmount) {
		return nil, false
	}
	return o.MinStakeAmount, true
}

// HasMinStakeAmount returns a boolean if a field has been set.
func (o *GetSolStakingQuotaDetailsResponse) HasMinStakeAmount() bool {
	if o != nil && !common.IsNil(o.MinStakeAmount) {
		return true
	}

	return false
}

// SetMinStakeAmount gets a reference to the given string and assigns it to the MinStakeAmount field.
func (o *GetSolStakingQuotaDetailsResponse) SetMinStakeAmount(v string) {
	o.MinStakeAmount = &v
}

// GetMinRedeemAmount returns the MinRedeemAmount field value if set, zero value otherwise.
func (o *GetSolStakingQuotaDetailsResponse) GetMinRedeemAmount() string {
	if o == nil || common.IsNil(o.MinRedeemAmount) {
		var ret string
		return ret
	}
	return *o.MinRedeemAmount
}

// GetMinRedeemAmountOk returns a tuple with the MinRedeemAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingQuotaDetailsResponse) GetMinRedeemAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinRedeemAmount) {
		return nil, false
	}
	return o.MinRedeemAmount, true
}

// HasMinRedeemAmount returns a boolean if a field has been set.
func (o *GetSolStakingQuotaDetailsResponse) HasMinRedeemAmount() bool {
	if o != nil && !common.IsNil(o.MinRedeemAmount) {
		return true
	}

	return false
}

// SetMinRedeemAmount gets a reference to the given string and assigns it to the MinRedeemAmount field.
func (o *GetSolStakingQuotaDetailsResponse) SetMinRedeemAmount(v string) {
	o.MinRedeemAmount = &v
}

// GetRedeemPeriod returns the RedeemPeriod field value if set, zero value otherwise.
func (o *GetSolStakingQuotaDetailsResponse) GetRedeemPeriod() int64 {
	if o == nil || common.IsNil(o.RedeemPeriod) {
		var ret int64
		return ret
	}
	return *o.RedeemPeriod
}

// GetRedeemPeriodOk returns a tuple with the RedeemPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingQuotaDetailsResponse) GetRedeemPeriodOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RedeemPeriod) {
		return nil, false
	}
	return o.RedeemPeriod, true
}

// HasRedeemPeriod returns a boolean if a field has been set.
func (o *GetSolStakingQuotaDetailsResponse) HasRedeemPeriod() bool {
	if o != nil && !common.IsNil(o.RedeemPeriod) {
		return true
	}

	return false
}

// SetRedeemPeriod gets a reference to the given int64 and assigns it to the RedeemPeriod field.
func (o *GetSolStakingQuotaDetailsResponse) SetRedeemPeriod(v int64) {
	o.RedeemPeriod = &v
}

// GetStakeable returns the Stakeable field value if set, zero value otherwise.
func (o *GetSolStakingQuotaDetailsResponse) GetStakeable() bool {
	if o == nil || common.IsNil(o.Stakeable) {
		var ret bool
		return ret
	}
	return *o.Stakeable
}

// GetStakeableOk returns a tuple with the Stakeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingQuotaDetailsResponse) GetStakeableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Stakeable) {
		return nil, false
	}
	return o.Stakeable, true
}

// HasStakeable returns a boolean if a field has been set.
func (o *GetSolStakingQuotaDetailsResponse) HasStakeable() bool {
	if o != nil && !common.IsNil(o.Stakeable) {
		return true
	}

	return false
}

// SetStakeable gets a reference to the given bool and assigns it to the Stakeable field.
func (o *GetSolStakingQuotaDetailsResponse) SetStakeable(v bool) {
	o.Stakeable = &v
}

// GetRedeemable returns the Redeemable field value if set, zero value otherwise.
func (o *GetSolStakingQuotaDetailsResponse) GetRedeemable() bool {
	if o == nil || common.IsNil(o.Redeemable) {
		var ret bool
		return ret
	}
	return *o.Redeemable
}

// GetRedeemableOk returns a tuple with the Redeemable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingQuotaDetailsResponse) GetRedeemableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Redeemable) {
		return nil, false
	}
	return o.Redeemable, true
}

// HasRedeemable returns a boolean if a field has been set.
func (o *GetSolStakingQuotaDetailsResponse) HasRedeemable() bool {
	if o != nil && !common.IsNil(o.Redeemable) {
		return true
	}

	return false
}

// SetRedeemable gets a reference to the given bool and assigns it to the Redeemable field.
func (o *GetSolStakingQuotaDetailsResponse) SetRedeemable(v bool) {
	o.Redeemable = &v
}

// GetSoldOut returns the SoldOut field value if set, zero value otherwise.
func (o *GetSolStakingQuotaDetailsResponse) GetSoldOut() bool {
	if o == nil || common.IsNil(o.SoldOut) {
		var ret bool
		return ret
	}
	return *o.SoldOut
}

// GetSoldOutOk returns a tuple with the SoldOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingQuotaDetailsResponse) GetSoldOutOk() (*bool, bool) {
	if o == nil || common.IsNil(o.SoldOut) {
		return nil, false
	}
	return o.SoldOut, true
}

// HasSoldOut returns a boolean if a field has been set.
func (o *GetSolStakingQuotaDetailsResponse) HasSoldOut() bool {
	if o != nil && !common.IsNil(o.SoldOut) {
		return true
	}

	return false
}

// SetSoldOut gets a reference to the given bool and assigns it to the SoldOut field.
func (o *GetSolStakingQuotaDetailsResponse) SetSoldOut(v bool) {
	o.SoldOut = &v
}

// GetCommissionFee returns the CommissionFee field value if set, zero value otherwise.
func (o *GetSolStakingQuotaDetailsResponse) GetCommissionFee() string {
	if o == nil || common.IsNil(o.CommissionFee) {
		var ret string
		return ret
	}
	return *o.CommissionFee
}

// GetCommissionFeeOk returns a tuple with the CommissionFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingQuotaDetailsResponse) GetCommissionFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CommissionFee) {
		return nil, false
	}
	return o.CommissionFee, true
}

// HasCommissionFee returns a boolean if a field has been set.
func (o *GetSolStakingQuotaDetailsResponse) HasCommissionFee() bool {
	if o != nil && !common.IsNil(o.CommissionFee) {
		return true
	}

	return false
}

// SetCommissionFee gets a reference to the given string and assigns it to the CommissionFee field.
func (o *GetSolStakingQuotaDetailsResponse) SetCommissionFee(v string) {
	o.CommissionFee = &v
}

// GetNextEpochTime returns the NextEpochTime field value if set, zero value otherwise.
func (o *GetSolStakingQuotaDetailsResponse) GetNextEpochTime() int64 {
	if o == nil || common.IsNil(o.NextEpochTime) {
		var ret int64
		return ret
	}
	return *o.NextEpochTime
}

// GetNextEpochTimeOk returns a tuple with the NextEpochTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingQuotaDetailsResponse) GetNextEpochTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.NextEpochTime) {
		return nil, false
	}
	return o.NextEpochTime, true
}

// HasNextEpochTime returns a boolean if a field has been set.
func (o *GetSolStakingQuotaDetailsResponse) HasNextEpochTime() bool {
	if o != nil && !common.IsNil(o.NextEpochTime) {
		return true
	}

	return false
}

// SetNextEpochTime gets a reference to the given int64 and assigns it to the NextEpochTime field.
func (o *GetSolStakingQuotaDetailsResponse) SetNextEpochTime(v int64) {
	o.NextEpochTime = &v
}

// GetCalculating returns the Calculating field value if set, zero value otherwise.
func (o *GetSolStakingQuotaDetailsResponse) GetCalculating() bool {
	if o == nil || common.IsNil(o.Calculating) {
		var ret bool
		return ret
	}
	return *o.Calculating
}

// GetCalculatingOk returns a tuple with the Calculating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingQuotaDetailsResponse) GetCalculatingOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Calculating) {
		return nil, false
	}
	return o.Calculating, true
}

// HasCalculating returns a boolean if a field has been set.
func (o *GetSolStakingQuotaDetailsResponse) HasCalculating() bool {
	if o != nil && !common.IsNil(o.Calculating) {
		return true
	}

	return false
}

// SetCalculating gets a reference to the given bool and assigns it to the Calculating field.
func (o *GetSolStakingQuotaDetailsResponse) SetCalculating(v bool) {
	o.Calculating = &v
}

func (o GetSolStakingQuotaDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSolStakingQuotaDetailsResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.SoldOut) {
		toSerialize["soldOut"] = o.SoldOut
	}
	if !common.IsNil(o.CommissionFee) {
		toSerialize["commissionFee"] = o.CommissionFee
	}
	if !common.IsNil(o.NextEpochTime) {
		toSerialize["nextEpochTime"] = o.NextEpochTime
	}
	if !common.IsNil(o.Calculating) {
		toSerialize["calculating"] = o.Calculating
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSolStakingQuotaDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	varGetSolStakingQuotaDetailsResponse := _GetSolStakingQuotaDetailsResponse{}

	err = json.Unmarshal(data, &varGetSolStakingQuotaDetailsResponse)

	if err != nil {
		return err
	}

	*o = GetSolStakingQuotaDetailsResponse(varGetSolStakingQuotaDetailsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "leftStakingPersonalQuota")
		delete(additionalProperties, "leftRedemptionPersonalQuota")
		delete(additionalProperties, "minStakeAmount")
		delete(additionalProperties, "minRedeemAmount")
		delete(additionalProperties, "redeemPeriod")
		delete(additionalProperties, "stakeable")
		delete(additionalProperties, "redeemable")
		delete(additionalProperties, "soldOut")
		delete(additionalProperties, "commissionFee")
		delete(additionalProperties, "nextEpochTime")
		delete(additionalProperties, "calculating")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSolStakingQuotaDetailsResponse struct {
	value *GetSolStakingQuotaDetailsResponse
	isSet bool
}

func (v NullableGetSolStakingQuotaDetailsResponse) Get() *GetSolStakingQuotaDetailsResponse {
	return v.value
}

func (v *NullableGetSolStakingQuotaDetailsResponse) Set(val *GetSolStakingQuotaDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolStakingQuotaDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolStakingQuotaDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolStakingQuotaDetailsResponse(val *GetSolStakingQuotaDetailsResponse) *NullableGetSolStakingQuotaDetailsResponse {
	return &NullableGetSolStakingQuotaDetailsResponse{value: val, isSet: true}
}

func (v NullableGetSolStakingQuotaDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSolStakingQuotaDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
