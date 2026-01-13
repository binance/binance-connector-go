/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner{}

// GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner struct for GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner
type GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner struct {
	EntryPrice           *string `json:"entryPrice,omitempty"`
	MarkPrice            *string `json:"markPrice,omitempty"`
	Leverage             *string `json:"leverage,omitempty"`
	Isolated             *string `json:"isolated,omitempty"`
	IsolatedWallet       *string `json:"isolatedWallet,omitempty"`
	IsolatedMargin       *string `json:"isolatedMargin,omitempty"`
	IsAutoAddMargin      *string `json:"isAutoAddMargin,omitempty"`
	PositionSide         *string `json:"positionSide,omitempty"`
	PositionAmount       *string `json:"positionAmount,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	UnrealizedProfit     *string `json:"unrealizedProfit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner

// NewGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner instantiates a new GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner() *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner {
	this := GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner{}
	return &this
}

// NewGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInnerWithDefaults instantiates a new GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInnerWithDefaults() *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner {
	this := GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner{}
	return &this
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetEntryPrice() string {
	if o == nil || common.IsNil(o.EntryPrice) {
		var ret string
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetEntryPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) HasEntryPrice() bool {
	if o != nil && !common.IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given string and assigns it to the EntryPrice field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) SetEntryPrice(v string) {
	o.EntryPrice = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetMarkPrice() string {
	if o == nil || common.IsNil(o.MarkPrice) {
		var ret string
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetMarkPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) HasMarkPrice() bool {
	if o != nil && !common.IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given string and assigns it to the MarkPrice field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) SetMarkPrice(v string) {
	o.MarkPrice = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetLeverage() string {
	if o == nil || common.IsNil(o.Leverage) {
		var ret string
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetLeverageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given string and assigns it to the Leverage field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) SetLeverage(v string) {
	o.Leverage = &v
}

// GetIsolated returns the Isolated field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetIsolated() string {
	if o == nil || common.IsNil(o.Isolated) {
		var ret string
		return ret
	}
	return *o.Isolated
}

// GetIsolatedOk returns a tuple with the Isolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetIsolatedOk() (*string, bool) {
	if o == nil || common.IsNil(o.Isolated) {
		return nil, false
	}
	return o.Isolated, true
}

// HasIsolated returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) HasIsolated() bool {
	if o != nil && !common.IsNil(o.Isolated) {
		return true
	}

	return false
}

// SetIsolated gets a reference to the given string and assigns it to the Isolated field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) SetIsolated(v string) {
	o.Isolated = &v
}

// GetIsolatedWallet returns the IsolatedWallet field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetIsolatedWallet() string {
	if o == nil || common.IsNil(o.IsolatedWallet) {
		var ret string
		return ret
	}
	return *o.IsolatedWallet
}

// GetIsolatedWalletOk returns a tuple with the IsolatedWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetIsolatedWalletOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsolatedWallet) {
		return nil, false
	}
	return o.IsolatedWallet, true
}

// HasIsolatedWallet returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) HasIsolatedWallet() bool {
	if o != nil && !common.IsNil(o.IsolatedWallet) {
		return true
	}

	return false
}

// SetIsolatedWallet gets a reference to the given string and assigns it to the IsolatedWallet field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) SetIsolatedWallet(v string) {
	o.IsolatedWallet = &v
}

// GetIsolatedMargin returns the IsolatedMargin field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetIsolatedMargin() string {
	if o == nil || common.IsNil(o.IsolatedMargin) {
		var ret string
		return ret
	}
	return *o.IsolatedMargin
}

// GetIsolatedMarginOk returns a tuple with the IsolatedMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetIsolatedMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsolatedMargin) {
		return nil, false
	}
	return o.IsolatedMargin, true
}

// HasIsolatedMargin returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) HasIsolatedMargin() bool {
	if o != nil && !common.IsNil(o.IsolatedMargin) {
		return true
	}

	return false
}

// SetIsolatedMargin gets a reference to the given string and assigns it to the IsolatedMargin field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) SetIsolatedMargin(v string) {
	o.IsolatedMargin = &v
}

// GetIsAutoAddMargin returns the IsAutoAddMargin field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetIsAutoAddMargin() string {
	if o == nil || common.IsNil(o.IsAutoAddMargin) {
		var ret string
		return ret
	}
	return *o.IsAutoAddMargin
}

// GetIsAutoAddMarginOk returns a tuple with the IsAutoAddMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetIsAutoAddMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsAutoAddMargin) {
		return nil, false
	}
	return o.IsAutoAddMargin, true
}

// HasIsAutoAddMargin returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) HasIsAutoAddMargin() bool {
	if o != nil && !common.IsNil(o.IsAutoAddMargin) {
		return true
	}

	return false
}

// SetIsAutoAddMargin gets a reference to the given string and assigns it to the IsAutoAddMargin field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) SetIsAutoAddMargin(v string) {
	o.IsAutoAddMargin = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetPositionAmount returns the PositionAmount field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetPositionAmount() string {
	if o == nil || common.IsNil(o.PositionAmount) {
		var ret string
		return ret
	}
	return *o.PositionAmount
}

// GetPositionAmountOk returns a tuple with the PositionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetPositionAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionAmount) {
		return nil, false
	}
	return o.PositionAmount, true
}

// HasPositionAmount returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) HasPositionAmount() bool {
	if o != nil && !common.IsNil(o.PositionAmount) {
		return true
	}

	return false
}

// SetPositionAmount gets a reference to the given string and assigns it to the PositionAmount field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) SetPositionAmount(v string) {
	o.PositionAmount = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetUnrealizedProfit returns the UnrealizedProfit field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetUnrealizedProfit() string {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfit
}

// GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) GetUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		return nil, false
	}
	return o.UnrealizedProfit, true
}

// HasUnrealizedProfit returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) HasUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.UnrealizedProfit) {
		return true
	}

	return false
}

// SetUnrealizedProfit gets a reference to the given string and assigns it to the UnrealizedProfit field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) SetUnrealizedProfit(v string) {
	o.UnrealizedProfit = &v
}

func (o GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EntryPrice) {
		toSerialize["entryPrice"] = o.EntryPrice
	}
	if !common.IsNil(o.MarkPrice) {
		toSerialize["markPrice"] = o.MarkPrice
	}
	if !common.IsNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}
	if !common.IsNil(o.Isolated) {
		toSerialize["isolated"] = o.Isolated
	}
	if !common.IsNil(o.IsolatedWallet) {
		toSerialize["isolatedWallet"] = o.IsolatedWallet
	}
	if !common.IsNil(o.IsolatedMargin) {
		toSerialize["isolatedMargin"] = o.IsolatedMargin
	}
	if !common.IsNil(o.IsAutoAddMargin) {
		toSerialize["isAutoAddMargin"] = o.IsAutoAddMargin
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !common.IsNil(o.PositionAmount) {
		toSerialize["positionAmount"] = o.PositionAmount
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.UnrealizedProfit) {
		toSerialize["unrealizedProfit"] = o.UnrealizedProfit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) UnmarshalJSON(data []byte) (err error) {
	varGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner := _GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner{}

	err = json.Unmarshal(data, &varGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner)

	if err != nil {
		return err
	}

	*o = GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner(varGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entryPrice")
		delete(additionalProperties, "markPrice")
		delete(additionalProperties, "leverage")
		delete(additionalProperties, "isolated")
		delete(additionalProperties, "isolatedWallet")
		delete(additionalProperties, "isolatedMargin")
		delete(additionalProperties, "isAutoAddMargin")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "positionAmount")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "unrealizedProfit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner struct {
	value *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner
	isSet bool
}

func (v NullableGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) Get() *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner {
	return v.value
}

func (v *NullableGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) Set(val *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner(val *GetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) *NullableGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner {
	return &NullableGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner{value: val, isSet: true}
}

func (v NullableGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFuturesPositionRiskOfSubAccountV2ResponseDeliveryPositionRiskVosInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
