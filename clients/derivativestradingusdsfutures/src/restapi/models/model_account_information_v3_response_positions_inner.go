/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountInformationV3ResponsePositionsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountInformationV3ResponsePositionsInner{}

// AccountInformationV3ResponsePositionsInner struct for AccountInformationV3ResponsePositionsInner
type AccountInformationV3ResponsePositionsInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	PositionSide         *string `json:"positionSide,omitempty"`
	PositionAmt          *string `json:"positionAmt,omitempty"`
	UnrealizedProfit     *string `json:"unrealizedProfit,omitempty"`
	IsolatedMargin       *string `json:"isolatedMargin,omitempty"`
	Notional             *string `json:"notional,omitempty"`
	IsolatedWallet       *string `json:"isolatedWallet,omitempty"`
	InitialMargin        *string `json:"initialMargin,omitempty"`
	MaintMargin          *string `json:"maintMargin,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountInformationV3ResponsePositionsInner AccountInformationV3ResponsePositionsInner

// NewAccountInformationV3ResponsePositionsInner instantiates a new AccountInformationV3ResponsePositionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInformationV3ResponsePositionsInner() *AccountInformationV3ResponsePositionsInner {
	this := AccountInformationV3ResponsePositionsInner{}
	return &this
}

// NewAccountInformationV3ResponsePositionsInnerWithDefaults instantiates a new AccountInformationV3ResponsePositionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInformationV3ResponsePositionsInnerWithDefaults() *AccountInformationV3ResponsePositionsInner {
	this := AccountInformationV3ResponsePositionsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *AccountInformationV3ResponsePositionsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3ResponsePositionsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *AccountInformationV3ResponsePositionsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *AccountInformationV3ResponsePositionsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *AccountInformationV3ResponsePositionsInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3ResponsePositionsInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *AccountInformationV3ResponsePositionsInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *AccountInformationV3ResponsePositionsInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetPositionAmt returns the PositionAmt field value if set, zero value otherwise.
func (o *AccountInformationV3ResponsePositionsInner) GetPositionAmt() string {
	if o == nil || common.IsNil(o.PositionAmt) {
		var ret string
		return ret
	}
	return *o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3ResponsePositionsInner) GetPositionAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionAmt) {
		return nil, false
	}
	return o.PositionAmt, true
}

// HasPositionAmt returns a boolean if a field has been set.
func (o *AccountInformationV3ResponsePositionsInner) HasPositionAmt() bool {
	if o != nil && !common.IsNil(o.PositionAmt) {
		return true
	}

	return false
}

// SetPositionAmt gets a reference to the given string and assigns it to the PositionAmt field.
func (o *AccountInformationV3ResponsePositionsInner) SetPositionAmt(v string) {
	o.PositionAmt = &v
}

// GetUnrealizedProfit returns the UnrealizedProfit field value if set, zero value otherwise.
func (o *AccountInformationV3ResponsePositionsInner) GetUnrealizedProfit() string {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfit
}

// GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3ResponsePositionsInner) GetUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		return nil, false
	}
	return o.UnrealizedProfit, true
}

// HasUnrealizedProfit returns a boolean if a field has been set.
func (o *AccountInformationV3ResponsePositionsInner) HasUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.UnrealizedProfit) {
		return true
	}

	return false
}

// SetUnrealizedProfit gets a reference to the given string and assigns it to the UnrealizedProfit field.
func (o *AccountInformationV3ResponsePositionsInner) SetUnrealizedProfit(v string) {
	o.UnrealizedProfit = &v
}

// GetIsolatedMargin returns the IsolatedMargin field value if set, zero value otherwise.
func (o *AccountInformationV3ResponsePositionsInner) GetIsolatedMargin() string {
	if o == nil || common.IsNil(o.IsolatedMargin) {
		var ret string
		return ret
	}
	return *o.IsolatedMargin
}

// GetIsolatedMarginOk returns a tuple with the IsolatedMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3ResponsePositionsInner) GetIsolatedMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsolatedMargin) {
		return nil, false
	}
	return o.IsolatedMargin, true
}

// HasIsolatedMargin returns a boolean if a field has been set.
func (o *AccountInformationV3ResponsePositionsInner) HasIsolatedMargin() bool {
	if o != nil && !common.IsNil(o.IsolatedMargin) {
		return true
	}

	return false
}

// SetIsolatedMargin gets a reference to the given string and assigns it to the IsolatedMargin field.
func (o *AccountInformationV3ResponsePositionsInner) SetIsolatedMargin(v string) {
	o.IsolatedMargin = &v
}

// GetNotional returns the Notional field value if set, zero value otherwise.
func (o *AccountInformationV3ResponsePositionsInner) GetNotional() string {
	if o == nil || common.IsNil(o.Notional) {
		var ret string
		return ret
	}
	return *o.Notional
}

// GetNotionalOk returns a tuple with the Notional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3ResponsePositionsInner) GetNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Notional) {
		return nil, false
	}
	return o.Notional, true
}

// HasNotional returns a boolean if a field has been set.
func (o *AccountInformationV3ResponsePositionsInner) HasNotional() bool {
	if o != nil && !common.IsNil(o.Notional) {
		return true
	}

	return false
}

// SetNotional gets a reference to the given string and assigns it to the Notional field.
func (o *AccountInformationV3ResponsePositionsInner) SetNotional(v string) {
	o.Notional = &v
}

// GetIsolatedWallet returns the IsolatedWallet field value if set, zero value otherwise.
func (o *AccountInformationV3ResponsePositionsInner) GetIsolatedWallet() string {
	if o == nil || common.IsNil(o.IsolatedWallet) {
		var ret string
		return ret
	}
	return *o.IsolatedWallet
}

// GetIsolatedWalletOk returns a tuple with the IsolatedWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3ResponsePositionsInner) GetIsolatedWalletOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsolatedWallet) {
		return nil, false
	}
	return o.IsolatedWallet, true
}

// HasIsolatedWallet returns a boolean if a field has been set.
func (o *AccountInformationV3ResponsePositionsInner) HasIsolatedWallet() bool {
	if o != nil && !common.IsNil(o.IsolatedWallet) {
		return true
	}

	return false
}

// SetIsolatedWallet gets a reference to the given string and assigns it to the IsolatedWallet field.
func (o *AccountInformationV3ResponsePositionsInner) SetIsolatedWallet(v string) {
	o.IsolatedWallet = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *AccountInformationV3ResponsePositionsInner) GetInitialMargin() string {
	if o == nil || common.IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3ResponsePositionsInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationV3ResponsePositionsInner) HasInitialMargin() bool {
	if o != nil && !common.IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *AccountInformationV3ResponsePositionsInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetMaintMargin returns the MaintMargin field value if set, zero value otherwise.
func (o *AccountInformationV3ResponsePositionsInner) GetMaintMargin() string {
	if o == nil || common.IsNil(o.MaintMargin) {
		var ret string
		return ret
	}
	return *o.MaintMargin
}

// GetMaintMarginOk returns a tuple with the MaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3ResponsePositionsInner) GetMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaintMargin) {
		return nil, false
	}
	return o.MaintMargin, true
}

// HasMaintMargin returns a boolean if a field has been set.
func (o *AccountInformationV3ResponsePositionsInner) HasMaintMargin() bool {
	if o != nil && !common.IsNil(o.MaintMargin) {
		return true
	}

	return false
}

// SetMaintMargin gets a reference to the given string and assigns it to the MaintMargin field.
func (o *AccountInformationV3ResponsePositionsInner) SetMaintMargin(v string) {
	o.MaintMargin = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AccountInformationV3ResponsePositionsInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3ResponsePositionsInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AccountInformationV3ResponsePositionsInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AccountInformationV3ResponsePositionsInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o AccountInformationV3ResponsePositionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInformationV3ResponsePositionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !common.IsNil(o.PositionAmt) {
		toSerialize["positionAmt"] = o.PositionAmt
	}
	if !common.IsNil(o.UnrealizedProfit) {
		toSerialize["unrealizedProfit"] = o.UnrealizedProfit
	}
	if !common.IsNil(o.IsolatedMargin) {
		toSerialize["isolatedMargin"] = o.IsolatedMargin
	}
	if !common.IsNil(o.Notional) {
		toSerialize["notional"] = o.Notional
	}
	if !common.IsNil(o.IsolatedWallet) {
		toSerialize["isolatedWallet"] = o.IsolatedWallet
	}
	if !common.IsNil(o.InitialMargin) {
		toSerialize["initialMargin"] = o.InitialMargin
	}
	if !common.IsNil(o.MaintMargin) {
		toSerialize["maintMargin"] = o.MaintMargin
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountInformationV3ResponsePositionsInner) UnmarshalJSON(data []byte) (err error) {
	varAccountInformationV3ResponsePositionsInner := _AccountInformationV3ResponsePositionsInner{}

	err = json.Unmarshal(data, &varAccountInformationV3ResponsePositionsInner)

	if err != nil {
		return err
	}

	*o = AccountInformationV3ResponsePositionsInner(varAccountInformationV3ResponsePositionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "positionAmt")
		delete(additionalProperties, "unrealizedProfit")
		delete(additionalProperties, "isolatedMargin")
		delete(additionalProperties, "notional")
		delete(additionalProperties, "isolatedWallet")
		delete(additionalProperties, "initialMargin")
		delete(additionalProperties, "maintMargin")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountInformationV3ResponsePositionsInner struct {
	value *AccountInformationV3ResponsePositionsInner
	isSet bool
}

func (v NullableAccountInformationV3ResponsePositionsInner) Get() *AccountInformationV3ResponsePositionsInner {
	return v.value
}

func (v *NullableAccountInformationV3ResponsePositionsInner) Set(val *AccountInformationV3ResponsePositionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInformationV3ResponsePositionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInformationV3ResponsePositionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInformationV3ResponsePositionsInner(val *AccountInformationV3ResponsePositionsInner) *NullableAccountInformationV3ResponsePositionsInner {
	return &NullableAccountInformationV3ResponsePositionsInner{value: val, isSet: true}
}

func (v NullableAccountInformationV3ResponsePositionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInformationV3ResponsePositionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
