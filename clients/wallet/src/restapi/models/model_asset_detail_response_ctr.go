/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AssetDetailResponseCTR type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AssetDetailResponseCTR{}

// AssetDetailResponseCTR struct for AssetDetailResponseCTR
type AssetDetailResponseCTR struct {
	MinWithdrawAmount    *string `json:"minWithdrawAmount,omitempty"`
	DepositStatus        *bool   `json:"depositStatus,omitempty"`
	WithdrawFee          *int64  `json:"withdrawFee,omitempty"`
	WithdrawStatus       *bool   `json:"withdrawStatus,omitempty"`
	DepositTip           *string `json:"depositTip,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetDetailResponseCTR AssetDetailResponseCTR

// NewAssetDetailResponseCTR instantiates a new AssetDetailResponseCTR object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDetailResponseCTR() *AssetDetailResponseCTR {
	this := AssetDetailResponseCTR{}
	return &this
}

// NewAssetDetailResponseCTRWithDefaults instantiates a new AssetDetailResponseCTR object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDetailResponseCTRWithDefaults() *AssetDetailResponseCTR {
	this := AssetDetailResponseCTR{}
	return &this
}

// GetMinWithdrawAmount returns the MinWithdrawAmount field value if set, zero value otherwise.
func (o *AssetDetailResponseCTR) GetMinWithdrawAmount() string {
	if o == nil || common.IsNil(o.MinWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.MinWithdrawAmount
}

// GetMinWithdrawAmountOk returns a tuple with the MinWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDetailResponseCTR) GetMinWithdrawAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinWithdrawAmount) {
		return nil, false
	}
	return o.MinWithdrawAmount, true
}

// HasMinWithdrawAmount returns a boolean if a field has been set.
func (o *AssetDetailResponseCTR) HasMinWithdrawAmount() bool {
	if o != nil && !common.IsNil(o.MinWithdrawAmount) {
		return true
	}

	return false
}

// SetMinWithdrawAmount gets a reference to the given string and assigns it to the MinWithdrawAmount field.
func (o *AssetDetailResponseCTR) SetMinWithdrawAmount(v string) {
	o.MinWithdrawAmount = &v
}

// GetDepositStatus returns the DepositStatus field value if set, zero value otherwise.
func (o *AssetDetailResponseCTR) GetDepositStatus() bool {
	if o == nil || common.IsNil(o.DepositStatus) {
		var ret bool
		return ret
	}
	return *o.DepositStatus
}

// GetDepositStatusOk returns a tuple with the DepositStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDetailResponseCTR) GetDepositStatusOk() (*bool, bool) {
	if o == nil || common.IsNil(o.DepositStatus) {
		return nil, false
	}
	return o.DepositStatus, true
}

// HasDepositStatus returns a boolean if a field has been set.
func (o *AssetDetailResponseCTR) HasDepositStatus() bool {
	if o != nil && !common.IsNil(o.DepositStatus) {
		return true
	}

	return false
}

// SetDepositStatus gets a reference to the given bool and assigns it to the DepositStatus field.
func (o *AssetDetailResponseCTR) SetDepositStatus(v bool) {
	o.DepositStatus = &v
}

// GetWithdrawFee returns the WithdrawFee field value if set, zero value otherwise.
func (o *AssetDetailResponseCTR) GetWithdrawFee() int64 {
	if o == nil || common.IsNil(o.WithdrawFee) {
		var ret int64
		return ret
	}
	return *o.WithdrawFee
}

// GetWithdrawFeeOk returns a tuple with the WithdrawFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDetailResponseCTR) GetWithdrawFeeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.WithdrawFee) {
		return nil, false
	}
	return o.WithdrawFee, true
}

// HasWithdrawFee returns a boolean if a field has been set.
func (o *AssetDetailResponseCTR) HasWithdrawFee() bool {
	if o != nil && !common.IsNil(o.WithdrawFee) {
		return true
	}

	return false
}

// SetWithdrawFee gets a reference to the given int64 and assigns it to the WithdrawFee field.
func (o *AssetDetailResponseCTR) SetWithdrawFee(v int64) {
	o.WithdrawFee = &v
}

// GetWithdrawStatus returns the WithdrawStatus field value if set, zero value otherwise.
func (o *AssetDetailResponseCTR) GetWithdrawStatus() bool {
	if o == nil || common.IsNil(o.WithdrawStatus) {
		var ret bool
		return ret
	}
	return *o.WithdrawStatus
}

// GetWithdrawStatusOk returns a tuple with the WithdrawStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDetailResponseCTR) GetWithdrawStatusOk() (*bool, bool) {
	if o == nil || common.IsNil(o.WithdrawStatus) {
		return nil, false
	}
	return o.WithdrawStatus, true
}

// HasWithdrawStatus returns a boolean if a field has been set.
func (o *AssetDetailResponseCTR) HasWithdrawStatus() bool {
	if o != nil && !common.IsNil(o.WithdrawStatus) {
		return true
	}

	return false
}

// SetWithdrawStatus gets a reference to the given bool and assigns it to the WithdrawStatus field.
func (o *AssetDetailResponseCTR) SetWithdrawStatus(v bool) {
	o.WithdrawStatus = &v
}

// GetDepositTip returns the DepositTip field value if set, zero value otherwise.
func (o *AssetDetailResponseCTR) GetDepositTip() string {
	if o == nil || common.IsNil(o.DepositTip) {
		var ret string
		return ret
	}
	return *o.DepositTip
}

// GetDepositTipOk returns a tuple with the DepositTip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDetailResponseCTR) GetDepositTipOk() (*string, bool) {
	if o == nil || common.IsNil(o.DepositTip) {
		return nil, false
	}
	return o.DepositTip, true
}

// HasDepositTip returns a boolean if a field has been set.
func (o *AssetDetailResponseCTR) HasDepositTip() bool {
	if o != nil && !common.IsNil(o.DepositTip) {
		return true
	}

	return false
}

// SetDepositTip gets a reference to the given string and assigns it to the DepositTip field.
func (o *AssetDetailResponseCTR) SetDepositTip(v string) {
	o.DepositTip = &v
}

func (o AssetDetailResponseCTR) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetDetailResponseCTR) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.DepositTip) {
		toSerialize["depositTip"] = o.DepositTip
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetDetailResponseCTR) UnmarshalJSON(data []byte) (err error) {
	varAssetDetailResponseCTR := _AssetDetailResponseCTR{}

	err = json.Unmarshal(data, &varAssetDetailResponseCTR)

	if err != nil {
		return err
	}

	*o = AssetDetailResponseCTR(varAssetDetailResponseCTR)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "minWithdrawAmount")
		delete(additionalProperties, "depositStatus")
		delete(additionalProperties, "withdrawFee")
		delete(additionalProperties, "withdrawStatus")
		delete(additionalProperties, "depositTip")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetDetailResponseCTR struct {
	value *AssetDetailResponseCTR
	isSet bool
}

func (v NullableAssetDetailResponseCTR) Get() *AssetDetailResponseCTR {
	return v.value
}

func (v *NullableAssetDetailResponseCTR) Set(val *AssetDetailResponseCTR) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDetailResponseCTR) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDetailResponseCTR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDetailResponseCTR(val *AssetDetailResponseCTR) *NullableAssetDetailResponseCTR {
	return &NullableAssetDetailResponseCTR{value: val, isSet: true}
}

func (v NullableAssetDetailResponseCTR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDetailResponseCTR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
