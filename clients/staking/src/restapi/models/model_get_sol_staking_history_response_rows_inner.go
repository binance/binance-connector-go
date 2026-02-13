/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetSolStakingHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSolStakingHistoryResponseRowsInner{}

// GetSolStakingHistoryResponseRowsInner struct for GetSolStakingHistoryResponseRowsInner
type GetSolStakingHistoryResponseRowsInner struct {
	Time                 *int64  `json:"time,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	DistributeAsset      *string `json:"distributeAsset,omitempty"`
	DistributeAmount     *string `json:"distributeAmount,omitempty"`
	ExchangeRate         *string `json:"exchangeRate,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSolStakingHistoryResponseRowsInner GetSolStakingHistoryResponseRowsInner

// NewGetSolStakingHistoryResponseRowsInner instantiates a new GetSolStakingHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolStakingHistoryResponseRowsInner() *GetSolStakingHistoryResponseRowsInner {
	this := GetSolStakingHistoryResponseRowsInner{}
	return &this
}

// NewGetSolStakingHistoryResponseRowsInnerWithDefaults instantiates a new GetSolStakingHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolStakingHistoryResponseRowsInnerWithDefaults() *GetSolStakingHistoryResponseRowsInner {
	this := GetSolStakingHistoryResponseRowsInner{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetSolStakingHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetSolStakingHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetSolStakingHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSolStakingHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSolStakingHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSolStakingHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetSolStakingHistoryResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingHistoryResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetSolStakingHistoryResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetSolStakingHistoryResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetDistributeAsset returns the DistributeAsset field value if set, zero value otherwise.
func (o *GetSolStakingHistoryResponseRowsInner) GetDistributeAsset() string {
	if o == nil || common.IsNil(o.DistributeAsset) {
		var ret string
		return ret
	}
	return *o.DistributeAsset
}

// GetDistributeAssetOk returns a tuple with the DistributeAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingHistoryResponseRowsInner) GetDistributeAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.DistributeAsset) {
		return nil, false
	}
	return o.DistributeAsset, true
}

// HasDistributeAsset returns a boolean if a field has been set.
func (o *GetSolStakingHistoryResponseRowsInner) HasDistributeAsset() bool {
	if o != nil && !common.IsNil(o.DistributeAsset) {
		return true
	}

	return false
}

// SetDistributeAsset gets a reference to the given string and assigns it to the DistributeAsset field.
func (o *GetSolStakingHistoryResponseRowsInner) SetDistributeAsset(v string) {
	o.DistributeAsset = &v
}

// GetDistributeAmount returns the DistributeAmount field value if set, zero value otherwise.
func (o *GetSolStakingHistoryResponseRowsInner) GetDistributeAmount() string {
	if o == nil || common.IsNil(o.DistributeAmount) {
		var ret string
		return ret
	}
	return *o.DistributeAmount
}

// GetDistributeAmountOk returns a tuple with the DistributeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingHistoryResponseRowsInner) GetDistributeAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.DistributeAmount) {
		return nil, false
	}
	return o.DistributeAmount, true
}

// HasDistributeAmount returns a boolean if a field has been set.
func (o *GetSolStakingHistoryResponseRowsInner) HasDistributeAmount() bool {
	if o != nil && !common.IsNil(o.DistributeAmount) {
		return true
	}

	return false
}

// SetDistributeAmount gets a reference to the given string and assigns it to the DistributeAmount field.
func (o *GetSolStakingHistoryResponseRowsInner) SetDistributeAmount(v string) {
	o.DistributeAmount = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *GetSolStakingHistoryResponseRowsInner) GetExchangeRate() string {
	if o == nil || common.IsNil(o.ExchangeRate) {
		var ret string
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingHistoryResponseRowsInner) GetExchangeRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *GetSolStakingHistoryResponseRowsInner) HasExchangeRate() bool {
	if o != nil && !common.IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given string and assigns it to the ExchangeRate field.
func (o *GetSolStakingHistoryResponseRowsInner) SetExchangeRate(v string) {
	o.ExchangeRate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetSolStakingHistoryResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingHistoryResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetSolStakingHistoryResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetSolStakingHistoryResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetSolStakingHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSolStakingHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.DistributeAsset) {
		toSerialize["distributeAsset"] = o.DistributeAsset
	}
	if !common.IsNil(o.DistributeAmount) {
		toSerialize["distributeAmount"] = o.DistributeAmount
	}
	if !common.IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSolStakingHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetSolStakingHistoryResponseRowsInner := _GetSolStakingHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetSolStakingHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetSolStakingHistoryResponseRowsInner(varGetSolStakingHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "distributeAsset")
		delete(additionalProperties, "distributeAmount")
		delete(additionalProperties, "exchangeRate")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSolStakingHistoryResponseRowsInner struct {
	value *GetSolStakingHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetSolStakingHistoryResponseRowsInner) Get() *GetSolStakingHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetSolStakingHistoryResponseRowsInner) Set(val *GetSolStakingHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolStakingHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolStakingHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolStakingHistoryResponseRowsInner(val *GetSolStakingHistoryResponseRowsInner) *NullableGetSolStakingHistoryResponseRowsInner {
	return &NullableGetSolStakingHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetSolStakingHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSolStakingHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
