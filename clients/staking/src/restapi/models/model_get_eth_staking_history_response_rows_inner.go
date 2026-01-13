/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetEthStakingHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetEthStakingHistoryResponseRowsInner{}

// GetEthStakingHistoryResponseRowsInner struct for GetEthStakingHistoryResponseRowsInner
type GetEthStakingHistoryResponseRowsInner struct {
	Time                 *int64  `json:"time,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	DistributeAsset      *string `json:"distributeAsset,omitempty"`
	DistributeAmount     *string `json:"distributeAmount,omitempty"`
	ConversionRatio      *string `json:"conversionRatio,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetEthStakingHistoryResponseRowsInner GetEthStakingHistoryResponseRowsInner

// NewGetEthStakingHistoryResponseRowsInner instantiates a new GetEthStakingHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEthStakingHistoryResponseRowsInner() *GetEthStakingHistoryResponseRowsInner {
	this := GetEthStakingHistoryResponseRowsInner{}
	return &this
}

// NewGetEthStakingHistoryResponseRowsInnerWithDefaults instantiates a new GetEthStakingHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEthStakingHistoryResponseRowsInnerWithDefaults() *GetEthStakingHistoryResponseRowsInner {
	this := GetEthStakingHistoryResponseRowsInner{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetEthStakingHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthStakingHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetEthStakingHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetEthStakingHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetEthStakingHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthStakingHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetEthStakingHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetEthStakingHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetEthStakingHistoryResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthStakingHistoryResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetEthStakingHistoryResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetEthStakingHistoryResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetDistributeAsset returns the DistributeAsset field value if set, zero value otherwise.
func (o *GetEthStakingHistoryResponseRowsInner) GetDistributeAsset() string {
	if o == nil || common.IsNil(o.DistributeAsset) {
		var ret string
		return ret
	}
	return *o.DistributeAsset
}

// GetDistributeAssetOk returns a tuple with the DistributeAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthStakingHistoryResponseRowsInner) GetDistributeAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.DistributeAsset) {
		return nil, false
	}
	return o.DistributeAsset, true
}

// HasDistributeAsset returns a boolean if a field has been set.
func (o *GetEthStakingHistoryResponseRowsInner) HasDistributeAsset() bool {
	if o != nil && !common.IsNil(o.DistributeAsset) {
		return true
	}

	return false
}

// SetDistributeAsset gets a reference to the given string and assigns it to the DistributeAsset field.
func (o *GetEthStakingHistoryResponseRowsInner) SetDistributeAsset(v string) {
	o.DistributeAsset = &v
}

// GetDistributeAmount returns the DistributeAmount field value if set, zero value otherwise.
func (o *GetEthStakingHistoryResponseRowsInner) GetDistributeAmount() string {
	if o == nil || common.IsNil(o.DistributeAmount) {
		var ret string
		return ret
	}
	return *o.DistributeAmount
}

// GetDistributeAmountOk returns a tuple with the DistributeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthStakingHistoryResponseRowsInner) GetDistributeAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.DistributeAmount) {
		return nil, false
	}
	return o.DistributeAmount, true
}

// HasDistributeAmount returns a boolean if a field has been set.
func (o *GetEthStakingHistoryResponseRowsInner) HasDistributeAmount() bool {
	if o != nil && !common.IsNil(o.DistributeAmount) {
		return true
	}

	return false
}

// SetDistributeAmount gets a reference to the given string and assigns it to the DistributeAmount field.
func (o *GetEthStakingHistoryResponseRowsInner) SetDistributeAmount(v string) {
	o.DistributeAmount = &v
}

// GetConversionRatio returns the ConversionRatio field value if set, zero value otherwise.
func (o *GetEthStakingHistoryResponseRowsInner) GetConversionRatio() string {
	if o == nil || common.IsNil(o.ConversionRatio) {
		var ret string
		return ret
	}
	return *o.ConversionRatio
}

// GetConversionRatioOk returns a tuple with the ConversionRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthStakingHistoryResponseRowsInner) GetConversionRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.ConversionRatio) {
		return nil, false
	}
	return o.ConversionRatio, true
}

// HasConversionRatio returns a boolean if a field has been set.
func (o *GetEthStakingHistoryResponseRowsInner) HasConversionRatio() bool {
	if o != nil && !common.IsNil(o.ConversionRatio) {
		return true
	}

	return false
}

// SetConversionRatio gets a reference to the given string and assigns it to the ConversionRatio field.
func (o *GetEthStakingHistoryResponseRowsInner) SetConversionRatio(v string) {
	o.ConversionRatio = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetEthStakingHistoryResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEthStakingHistoryResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetEthStakingHistoryResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetEthStakingHistoryResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetEthStakingHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEthStakingHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.ConversionRatio) {
		toSerialize["conversionRatio"] = o.ConversionRatio
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetEthStakingHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetEthStakingHistoryResponseRowsInner := _GetEthStakingHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetEthStakingHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetEthStakingHistoryResponseRowsInner(varGetEthStakingHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "distributeAsset")
		delete(additionalProperties, "distributeAmount")
		delete(additionalProperties, "conversionRatio")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetEthStakingHistoryResponseRowsInner struct {
	value *GetEthStakingHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetEthStakingHistoryResponseRowsInner) Get() *GetEthStakingHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetEthStakingHistoryResponseRowsInner) Set(val *GetEthStakingHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEthStakingHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEthStakingHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEthStakingHistoryResponseRowsInner(val *GetEthStakingHistoryResponseRowsInner) *NullableGetEthStakingHistoryResponseRowsInner {
	return &NullableGetEthStakingHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetEthStakingHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEthStakingHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
