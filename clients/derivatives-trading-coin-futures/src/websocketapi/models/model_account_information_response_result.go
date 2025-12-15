/*
Binance Derivatives Trading COIN Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountInformationResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountInformationResponseResult{}

// AccountInformationResponseResult struct for AccountInformationResponseResult
type AccountInformationResponseResult struct {
	FeeTier              *int64                                           `json:"feeTier,omitempty"`
	CanTrade             *bool                                            `json:"canTrade,omitempty"`
	CanDeposit           *bool                                            `json:"canDeposit,omitempty"`
	CanWithdraw          *bool                                            `json:"canWithdraw,omitempty"`
	UpdateTime           *int64                                           `json:"updateTime,omitempty"`
	Assets               []AccountInformationResponseResultAssetsInner    `json:"assets,omitempty"`
	Positions            []AccountInformationResponseResultPositionsInner `json:"positions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountInformationResponseResult AccountInformationResponseResult

// NewAccountInformationResponseResult instantiates a new AccountInformationResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInformationResponseResult() *AccountInformationResponseResult {
	this := AccountInformationResponseResult{}
	return &this
}

// NewAccountInformationResponseResultWithDefaults instantiates a new AccountInformationResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInformationResponseResultWithDefaults() *AccountInformationResponseResult {
	this := AccountInformationResponseResult{}
	return &this
}

// GetFeeTier returns the FeeTier field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetFeeTier() int64 {
	if o == nil || common.IsNil(o.FeeTier) {
		var ret int64
		return ret
	}
	return *o.FeeTier
}

// GetFeeTierOk returns a tuple with the FeeTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetFeeTierOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FeeTier) {
		return nil, false
	}
	return o.FeeTier, true
}

// HasFeeTier returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasFeeTier() bool {
	if o != nil && !common.IsNil(o.FeeTier) {
		return true
	}

	return false
}

// SetFeeTier gets a reference to the given int64 and assigns it to the FeeTier field.
func (o *AccountInformationResponseResult) SetFeeTier(v int64) {
	o.FeeTier = &v
}

// GetCanTrade returns the CanTrade field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetCanTrade() bool {
	if o == nil || common.IsNil(o.CanTrade) {
		var ret bool
		return ret
	}
	return *o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetCanTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanTrade) {
		return nil, false
	}
	return o.CanTrade, true
}

// HasCanTrade returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasCanTrade() bool {
	if o != nil && !common.IsNil(o.CanTrade) {
		return true
	}

	return false
}

// SetCanTrade gets a reference to the given bool and assigns it to the CanTrade field.
func (o *AccountInformationResponseResult) SetCanTrade(v bool) {
	o.CanTrade = &v
}

// GetCanDeposit returns the CanDeposit field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetCanDeposit() bool {
	if o == nil || common.IsNil(o.CanDeposit) {
		var ret bool
		return ret
	}
	return *o.CanDeposit
}

// GetCanDepositOk returns a tuple with the CanDeposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetCanDepositOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanDeposit) {
		return nil, false
	}
	return o.CanDeposit, true
}

// HasCanDeposit returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasCanDeposit() bool {
	if o != nil && !common.IsNil(o.CanDeposit) {
		return true
	}

	return false
}

// SetCanDeposit gets a reference to the given bool and assigns it to the CanDeposit field.
func (o *AccountInformationResponseResult) SetCanDeposit(v bool) {
	o.CanDeposit = &v
}

// GetCanWithdraw returns the CanWithdraw field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetCanWithdraw() bool {
	if o == nil || common.IsNil(o.CanWithdraw) {
		var ret bool
		return ret
	}
	return *o.CanWithdraw
}

// GetCanWithdrawOk returns a tuple with the CanWithdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetCanWithdrawOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanWithdraw) {
		return nil, false
	}
	return o.CanWithdraw, true
}

// HasCanWithdraw returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasCanWithdraw() bool {
	if o != nil && !common.IsNil(o.CanWithdraw) {
		return true
	}

	return false
}

// SetCanWithdraw gets a reference to the given bool and assigns it to the CanWithdraw field.
func (o *AccountInformationResponseResult) SetCanWithdraw(v bool) {
	o.CanWithdraw = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AccountInformationResponseResult) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetAssets() []AccountInformationResponseResultAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []AccountInformationResponseResultAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetAssetsOk() ([]AccountInformationResponseResultAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []AccountInformationResponseResultAssetsInner and assigns it to the Assets field.
func (o *AccountInformationResponseResult) SetAssets(v []AccountInformationResponseResultAssetsInner) {
	o.Assets = v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *AccountInformationResponseResult) GetPositions() []AccountInformationResponseResultPositionsInner {
	if o == nil || common.IsNil(o.Positions) {
		var ret []AccountInformationResponseResultPositionsInner
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResult) GetPositionsOk() ([]AccountInformationResponseResultPositionsInner, bool) {
	if o == nil || common.IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *AccountInformationResponseResult) HasPositions() bool {
	if o != nil && !common.IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given []AccountInformationResponseResultPositionsInner and assigns it to the Positions field.
func (o *AccountInformationResponseResult) SetPositions(v []AccountInformationResponseResultPositionsInner) {
	o.Positions = v
}

func (o AccountInformationResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInformationResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FeeTier) {
		toSerialize["feeTier"] = o.FeeTier
	}
	if !common.IsNil(o.CanTrade) {
		toSerialize["canTrade"] = o.CanTrade
	}
	if !common.IsNil(o.CanDeposit) {
		toSerialize["canDeposit"] = o.CanDeposit
	}
	if !common.IsNil(o.CanWithdraw) {
		toSerialize["canWithdraw"] = o.CanWithdraw
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !common.IsNil(o.Positions) {
		toSerialize["positions"] = o.Positions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountInformationResponseResult) UnmarshalJSON(data []byte) (err error) {
	varAccountInformationResponseResult := _AccountInformationResponseResult{}

	err = json.Unmarshal(data, &varAccountInformationResponseResult)

	if err != nil {
		return err
	}

	*o = AccountInformationResponseResult(varAccountInformationResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "feeTier")
		delete(additionalProperties, "canTrade")
		delete(additionalProperties, "canDeposit")
		delete(additionalProperties, "canWithdraw")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "assets")
		delete(additionalProperties, "positions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountInformationResponseResult struct {
	value *AccountInformationResponseResult
	isSet bool
}

func (v NullableAccountInformationResponseResult) Get() *AccountInformationResponseResult {
	return v.value
}

func (v *NullableAccountInformationResponseResult) Set(val *AccountInformationResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInformationResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInformationResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInformationResponseResult(val *AccountInformationResponseResult) *NullableAccountInformationResponseResult {
	return &NullableAccountInformationResponseResult{value: val, isSet: true}
}

func (v NullableAccountInformationResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInformationResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
