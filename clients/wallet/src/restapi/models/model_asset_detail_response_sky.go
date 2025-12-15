/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AssetDetailResponseSKY type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AssetDetailResponseSKY{}

// AssetDetailResponseSKY struct for AssetDetailResponseSKY
type AssetDetailResponseSKY struct {
	MinWithdrawAmount    *string  `json:"minWithdrawAmount,omitempty"`
	DepositStatus        *bool    `json:"depositStatus,omitempty"`
	WithdrawFee          *float32 `json:"withdrawFee,omitempty"`
	WithdrawStatus       *bool    `json:"withdrawStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetDetailResponseSKY AssetDetailResponseSKY

// NewAssetDetailResponseSKY instantiates a new AssetDetailResponseSKY object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDetailResponseSKY() *AssetDetailResponseSKY {
	this := AssetDetailResponseSKY{}
	return &this
}

// NewAssetDetailResponseSKYWithDefaults instantiates a new AssetDetailResponseSKY object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDetailResponseSKYWithDefaults() *AssetDetailResponseSKY {
	this := AssetDetailResponseSKY{}
	return &this
}

// GetMinWithdrawAmount returns the MinWithdrawAmount field value if set, zero value otherwise.
func (o *AssetDetailResponseSKY) GetMinWithdrawAmount() string {
	if o == nil || common.IsNil(o.MinWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.MinWithdrawAmount
}

// GetMinWithdrawAmountOk returns a tuple with the MinWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDetailResponseSKY) GetMinWithdrawAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinWithdrawAmount) {
		return nil, false
	}
	return o.MinWithdrawAmount, true
}

// HasMinWithdrawAmount returns a boolean if a field has been set.
func (o *AssetDetailResponseSKY) HasMinWithdrawAmount() bool {
	if o != nil && !common.IsNil(o.MinWithdrawAmount) {
		return true
	}

	return false
}

// SetMinWithdrawAmount gets a reference to the given string and assigns it to the MinWithdrawAmount field.
func (o *AssetDetailResponseSKY) SetMinWithdrawAmount(v string) {
	o.MinWithdrawAmount = &v
}

// GetDepositStatus returns the DepositStatus field value if set, zero value otherwise.
func (o *AssetDetailResponseSKY) GetDepositStatus() bool {
	if o == nil || common.IsNil(o.DepositStatus) {
		var ret bool
		return ret
	}
	return *o.DepositStatus
}

// GetDepositStatusOk returns a tuple with the DepositStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDetailResponseSKY) GetDepositStatusOk() (*bool, bool) {
	if o == nil || common.IsNil(o.DepositStatus) {
		return nil, false
	}
	return o.DepositStatus, true
}

// HasDepositStatus returns a boolean if a field has been set.
func (o *AssetDetailResponseSKY) HasDepositStatus() bool {
	if o != nil && !common.IsNil(o.DepositStatus) {
		return true
	}

	return false
}

// SetDepositStatus gets a reference to the given bool and assigns it to the DepositStatus field.
func (o *AssetDetailResponseSKY) SetDepositStatus(v bool) {
	o.DepositStatus = &v
}

// GetWithdrawFee returns the WithdrawFee field value if set, zero value otherwise.
func (o *AssetDetailResponseSKY) GetWithdrawFee() float32 {
	if o == nil || common.IsNil(o.WithdrawFee) {
		var ret float32
		return ret
	}
	return *o.WithdrawFee
}

// GetWithdrawFeeOk returns a tuple with the WithdrawFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDetailResponseSKY) GetWithdrawFeeOk() (*float32, bool) {
	if o == nil || common.IsNil(o.WithdrawFee) {
		return nil, false
	}
	return o.WithdrawFee, true
}

// HasWithdrawFee returns a boolean if a field has been set.
func (o *AssetDetailResponseSKY) HasWithdrawFee() bool {
	if o != nil && !common.IsNil(o.WithdrawFee) {
		return true
	}

	return false
}

// SetWithdrawFee gets a reference to the given float32 and assigns it to the WithdrawFee field.
func (o *AssetDetailResponseSKY) SetWithdrawFee(v float32) {
	o.WithdrawFee = &v
}

// GetWithdrawStatus returns the WithdrawStatus field value if set, zero value otherwise.
func (o *AssetDetailResponseSKY) GetWithdrawStatus() bool {
	if o == nil || common.IsNil(o.WithdrawStatus) {
		var ret bool
		return ret
	}
	return *o.WithdrawStatus
}

// GetWithdrawStatusOk returns a tuple with the WithdrawStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDetailResponseSKY) GetWithdrawStatusOk() (*bool, bool) {
	if o == nil || common.IsNil(o.WithdrawStatus) {
		return nil, false
	}
	return o.WithdrawStatus, true
}

// HasWithdrawStatus returns a boolean if a field has been set.
func (o *AssetDetailResponseSKY) HasWithdrawStatus() bool {
	if o != nil && !common.IsNil(o.WithdrawStatus) {
		return true
	}

	return false
}

// SetWithdrawStatus gets a reference to the given bool and assigns it to the WithdrawStatus field.
func (o *AssetDetailResponseSKY) SetWithdrawStatus(v bool) {
	o.WithdrawStatus = &v
}

func (o AssetDetailResponseSKY) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetDetailResponseSKY) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MinWithdrawAmount) {
		toSerialize["minWithdrawAmount"] = o.MinWithdrawAmount
	}
	if !common.IsNil(o.DepositStatus) {
		toSerialize["depositStatus"] = o.DepositStatus
	}
	if !common.IsNil(o.WithdrawFee) {
		toSerialize["withdrawFee"] = o.WithdrawFee
	}
	if !common.IsNil(o.WithdrawStatus) {
		toSerialize["withdrawStatus"] = o.WithdrawStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetDetailResponseSKY) UnmarshalJSON(data []byte) (err error) {
	varAssetDetailResponseSKY := _AssetDetailResponseSKY{}

	err = json.Unmarshal(data, &varAssetDetailResponseSKY)

	if err != nil {
		return err
	}

	*o = AssetDetailResponseSKY(varAssetDetailResponseSKY)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "minWithdrawAmount")
		delete(additionalProperties, "depositStatus")
		delete(additionalProperties, "withdrawFee")
		delete(additionalProperties, "withdrawStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetDetailResponseSKY struct {
	value *AssetDetailResponseSKY
	isSet bool
}

func (v NullableAssetDetailResponseSKY) Get() *AssetDetailResponseSKY {
	return v.value
}

func (v *NullableAssetDetailResponseSKY) Set(val *AssetDetailResponseSKY) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDetailResponseSKY) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDetailResponseSKY) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDetailResponseSKY(val *AssetDetailResponseSKY) *NullableAssetDetailResponseSKY {
	return &NullableAssetDetailResponseSKY{value: val, isSet: true}
}

func (v NullableAssetDetailResponseSKY) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDetailResponseSKY) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
