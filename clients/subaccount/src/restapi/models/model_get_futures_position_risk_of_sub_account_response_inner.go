/*
Sub Account REST API

Create and manage sub-accounts, control permissions, and transfer assets via the Sub Account API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFuturesPositionRiskOfSubAccountResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFuturesPositionRiskOfSubAccountResponseInner{}

// GetFuturesPositionRiskOfSubAccountResponseInner struct for GetFuturesPositionRiskOfSubAccountResponseInner
type GetFuturesPositionRiskOfSubAccountResponseInner struct {
	EntryPrice           *string `json:"entryPrice,omitempty"`
	Leverage             *string `json:"leverage,omitempty"`
	MaxNotional          *string `json:"maxNotional,omitempty"`
	LiquidationPrice     *string `json:"liquidationPrice,omitempty"`
	MarkPrice            *string `json:"markPrice,omitempty"`
	PositionAmount       *string `json:"positionAmount,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	UnrealizedProfit     *string `json:"unrealizedProfit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFuturesPositionRiskOfSubAccountResponseInner GetFuturesPositionRiskOfSubAccountResponseInner

// NewGetFuturesPositionRiskOfSubAccountResponseInner instantiates a new GetFuturesPositionRiskOfSubAccountResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFuturesPositionRiskOfSubAccountResponseInner() *GetFuturesPositionRiskOfSubAccountResponseInner {
	this := GetFuturesPositionRiskOfSubAccountResponseInner{}
	return &this
}

// NewGetFuturesPositionRiskOfSubAccountResponseInnerWithDefaults instantiates a new GetFuturesPositionRiskOfSubAccountResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFuturesPositionRiskOfSubAccountResponseInnerWithDefaults() *GetFuturesPositionRiskOfSubAccountResponseInner {
	this := GetFuturesPositionRiskOfSubAccountResponseInner{}
	return &this
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) GetEntryPrice() string {
	if o == nil || common.IsNil(o.EntryPrice) {
		var ret string
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) GetEntryPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) HasEntryPrice() bool {
	if o != nil && !common.IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given string and assigns it to the EntryPrice field.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) SetEntryPrice(v string) {
	o.EntryPrice = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) GetLeverage() string {
	if o == nil || common.IsNil(o.Leverage) {
		var ret string
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) GetLeverageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given string and assigns it to the Leverage field.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) SetLeverage(v string) {
	o.Leverage = &v
}

// GetMaxNotional returns the MaxNotional field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) GetMaxNotional() string {
	if o == nil || common.IsNil(o.MaxNotional) {
		var ret string
		return ret
	}
	return *o.MaxNotional
}

// GetMaxNotionalOk returns a tuple with the MaxNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) GetMaxNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxNotional) {
		return nil, false
	}
	return o.MaxNotional, true
}

// HasMaxNotional returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) HasMaxNotional() bool {
	if o != nil && !common.IsNil(o.MaxNotional) {
		return true
	}

	return false
}

// SetMaxNotional gets a reference to the given string and assigns it to the MaxNotional field.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) SetMaxNotional(v string) {
	o.MaxNotional = &v
}

// GetLiquidationPrice returns the LiquidationPrice field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) GetLiquidationPrice() string {
	if o == nil || common.IsNil(o.LiquidationPrice) {
		var ret string
		return ret
	}
	return *o.LiquidationPrice
}

// GetLiquidationPriceOk returns a tuple with the LiquidationPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) GetLiquidationPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiquidationPrice) {
		return nil, false
	}
	return o.LiquidationPrice, true
}

// HasLiquidationPrice returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) HasLiquidationPrice() bool {
	if o != nil && !common.IsNil(o.LiquidationPrice) {
		return true
	}

	return false
}

// SetLiquidationPrice gets a reference to the given string and assigns it to the LiquidationPrice field.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) SetLiquidationPrice(v string) {
	o.LiquidationPrice = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) GetMarkPrice() string {
	if o == nil || common.IsNil(o.MarkPrice) {
		var ret string
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) GetMarkPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) HasMarkPrice() bool {
	if o != nil && !common.IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given string and assigns it to the MarkPrice field.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) SetMarkPrice(v string) {
	o.MarkPrice = &v
}

// GetPositionAmount returns the PositionAmount field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) GetPositionAmount() string {
	if o == nil || common.IsNil(o.PositionAmount) {
		var ret string
		return ret
	}
	return *o.PositionAmount
}

// GetPositionAmountOk returns a tuple with the PositionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) GetPositionAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionAmount) {
		return nil, false
	}
	return o.PositionAmount, true
}

// HasPositionAmount returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) HasPositionAmount() bool {
	if o != nil && !common.IsNil(o.PositionAmount) {
		return true
	}

	return false
}

// SetPositionAmount gets a reference to the given string and assigns it to the PositionAmount field.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) SetPositionAmount(v string) {
	o.PositionAmount = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetUnrealizedProfit returns the UnrealizedProfit field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) GetUnrealizedProfit() string {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfit
}

// GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) GetUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		return nil, false
	}
	return o.UnrealizedProfit, true
}

// HasUnrealizedProfit returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) HasUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.UnrealizedProfit) {
		return true
	}

	return false
}

// SetUnrealizedProfit gets a reference to the given string and assigns it to the UnrealizedProfit field.
func (o *GetFuturesPositionRiskOfSubAccountResponseInner) SetUnrealizedProfit(v string) {
	o.UnrealizedProfit = &v
}

func (o GetFuturesPositionRiskOfSubAccountResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFuturesPositionRiskOfSubAccountResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EntryPrice) {
		toSerialize["entryPrice"] = o.EntryPrice
	}
	if !common.IsNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}
	if !common.IsNil(o.MaxNotional) {
		toSerialize["maxNotional"] = o.MaxNotional
	}
	if !common.IsNil(o.LiquidationPrice) {
		toSerialize["liquidationPrice"] = o.LiquidationPrice
	}
	if !common.IsNil(o.MarkPrice) {
		toSerialize["markPrice"] = o.MarkPrice
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

func (o *GetFuturesPositionRiskOfSubAccountResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetFuturesPositionRiskOfSubAccountResponseInner := _GetFuturesPositionRiskOfSubAccountResponseInner{}

	err = json.Unmarshal(data, &varGetFuturesPositionRiskOfSubAccountResponseInner)

	if err != nil {
		return err
	}

	*o = GetFuturesPositionRiskOfSubAccountResponseInner(varGetFuturesPositionRiskOfSubAccountResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entryPrice")
		delete(additionalProperties, "leverage")
		delete(additionalProperties, "maxNotional")
		delete(additionalProperties, "liquidationPrice")
		delete(additionalProperties, "markPrice")
		delete(additionalProperties, "positionAmount")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "unrealizedProfit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFuturesPositionRiskOfSubAccountResponseInner struct {
	value *GetFuturesPositionRiskOfSubAccountResponseInner
	isSet bool
}

func (v NullableGetFuturesPositionRiskOfSubAccountResponseInner) Get() *GetFuturesPositionRiskOfSubAccountResponseInner {
	return v.value
}

func (v *NullableGetFuturesPositionRiskOfSubAccountResponseInner) Set(val *GetFuturesPositionRiskOfSubAccountResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFuturesPositionRiskOfSubAccountResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFuturesPositionRiskOfSubAccountResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFuturesPositionRiskOfSubAccountResponseInner(val *GetFuturesPositionRiskOfSubAccountResponseInner) *NullableGetFuturesPositionRiskOfSubAccountResponseInner {
	return &NullableGetFuturesPositionRiskOfSubAccountResponseInner{value: val, isSet: true}
}

func (v NullableGetFuturesPositionRiskOfSubAccountResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFuturesPositionRiskOfSubAccountResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
