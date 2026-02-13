/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner{}

// GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner struct for GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner
type GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner struct {
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

type _GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner

// NewGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner instantiates a new GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner() *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner {
	this := GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner{}
	return &this
}

// NewGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInnerWithDefaults instantiates a new GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInnerWithDefaults() *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner {
	this := GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner{}
	return &this
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) GetEntryPrice() string {
	if o == nil || common.IsNil(o.EntryPrice) {
		var ret string
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) GetEntryPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) HasEntryPrice() bool {
	if o != nil && !common.IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given string and assigns it to the EntryPrice field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) SetEntryPrice(v string) {
	o.EntryPrice = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) GetLeverage() string {
	if o == nil || common.IsNil(o.Leverage) {
		var ret string
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) GetLeverageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given string and assigns it to the Leverage field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) SetLeverage(v string) {
	o.Leverage = &v
}

// GetMaxNotional returns the MaxNotional field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) GetMaxNotional() string {
	if o == nil || common.IsNil(o.MaxNotional) {
		var ret string
		return ret
	}
	return *o.MaxNotional
}

// GetMaxNotionalOk returns a tuple with the MaxNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) GetMaxNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxNotional) {
		return nil, false
	}
	return o.MaxNotional, true
}

// HasMaxNotional returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) HasMaxNotional() bool {
	if o != nil && !common.IsNil(o.MaxNotional) {
		return true
	}

	return false
}

// SetMaxNotional gets a reference to the given string and assigns it to the MaxNotional field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) SetMaxNotional(v string) {
	o.MaxNotional = &v
}

// GetLiquidationPrice returns the LiquidationPrice field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) GetLiquidationPrice() string {
	if o == nil || common.IsNil(o.LiquidationPrice) {
		var ret string
		return ret
	}
	return *o.LiquidationPrice
}

// GetLiquidationPriceOk returns a tuple with the LiquidationPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) GetLiquidationPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiquidationPrice) {
		return nil, false
	}
	return o.LiquidationPrice, true
}

// HasLiquidationPrice returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) HasLiquidationPrice() bool {
	if o != nil && !common.IsNil(o.LiquidationPrice) {
		return true
	}

	return false
}

// SetLiquidationPrice gets a reference to the given string and assigns it to the LiquidationPrice field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) SetLiquidationPrice(v string) {
	o.LiquidationPrice = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) GetMarkPrice() string {
	if o == nil || common.IsNil(o.MarkPrice) {
		var ret string
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) GetMarkPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) HasMarkPrice() bool {
	if o != nil && !common.IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given string and assigns it to the MarkPrice field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) SetMarkPrice(v string) {
	o.MarkPrice = &v
}

// GetPositionAmount returns the PositionAmount field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) GetPositionAmount() string {
	if o == nil || common.IsNil(o.PositionAmount) {
		var ret string
		return ret
	}
	return *o.PositionAmount
}

// GetPositionAmountOk returns a tuple with the PositionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) GetPositionAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionAmount) {
		return nil, false
	}
	return o.PositionAmount, true
}

// HasPositionAmount returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) HasPositionAmount() bool {
	if o != nil && !common.IsNil(o.PositionAmount) {
		return true
	}

	return false
}

// SetPositionAmount gets a reference to the given string and assigns it to the PositionAmount field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) SetPositionAmount(v string) {
	o.PositionAmount = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetUnrealizedProfit returns the UnrealizedProfit field value if set, zero value otherwise.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) GetUnrealizedProfit() string {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfit
}

// GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) GetUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		return nil, false
	}
	return o.UnrealizedProfit, true
}

// HasUnrealizedProfit returns a boolean if a field has been set.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) HasUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.UnrealizedProfit) {
		return true
	}

	return false
}

// SetUnrealizedProfit gets a reference to the given string and assigns it to the UnrealizedProfit field.
func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) SetUnrealizedProfit(v string) {
	o.UnrealizedProfit = &v
}

func (o GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) ToMap() (map[string]interface{}, error) {
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

func (o *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) UnmarshalJSON(data []byte) (err error) {
	varGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner := _GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner{}

	err = json.Unmarshal(data, &varGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner)

	if err != nil {
		return err
	}

	*o = GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner(varGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner)

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

type NullableGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner struct {
	value *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner
	isSet bool
}

func (v NullableGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) Get() *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner {
	return v.value
}

func (v *NullableGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) Set(val *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner(val *GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) *NullableGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner {
	return &NullableGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner{value: val, isSet: true}
}

func (v NullableGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
