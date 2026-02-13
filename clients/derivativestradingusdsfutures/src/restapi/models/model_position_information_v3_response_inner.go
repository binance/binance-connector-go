/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PositionInformationV3ResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PositionInformationV3ResponseInner{}

// PositionInformationV3ResponseInner struct for PositionInformationV3ResponseInner
type PositionInformationV3ResponseInner struct {
	Symbol                 *string `json:"symbol,omitempty"`
	PositionSide           *string `json:"positionSide,omitempty"`
	PositionAmt            *string `json:"positionAmt,omitempty"`
	EntryPrice             *string `json:"entryPrice,omitempty"`
	BreakEvenPrice         *string `json:"breakEvenPrice,omitempty"`
	MarkPrice              *string `json:"markPrice,omitempty"`
	UnRealizedProfit       *string `json:"unRealizedProfit,omitempty"`
	LiquidationPrice       *string `json:"liquidationPrice,omitempty"`
	IsolatedMargin         *string `json:"isolatedMargin,omitempty"`
	Notional               *string `json:"notional,omitempty"`
	MarginAsset            *string `json:"marginAsset,omitempty"`
	IsolatedWallet         *string `json:"isolatedWallet,omitempty"`
	InitialMargin          *string `json:"initialMargin,omitempty"`
	MaintMargin            *string `json:"maintMargin,omitempty"`
	PositionInitialMargin  *string `json:"positionInitialMargin,omitempty"`
	OpenOrderInitialMargin *string `json:"openOrderInitialMargin,omitempty"`
	Adl                    *int64  `json:"adl,omitempty"`
	BidNotional            *string `json:"bidNotional,omitempty"`
	AskNotional            *string `json:"askNotional,omitempty"`
	UpdateTime             *int64  `json:"updateTime,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _PositionInformationV3ResponseInner PositionInformationV3ResponseInner

// NewPositionInformationV3ResponseInner instantiates a new PositionInformationV3ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPositionInformationV3ResponseInner() *PositionInformationV3ResponseInner {
	this := PositionInformationV3ResponseInner{}
	return &this
}

// NewPositionInformationV3ResponseInnerWithDefaults instantiates a new PositionInformationV3ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositionInformationV3ResponseInnerWithDefaults() *PositionInformationV3ResponseInner {
	this := PositionInformationV3ResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *PositionInformationV3ResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *PositionInformationV3ResponseInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetPositionAmt returns the PositionAmt field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetPositionAmt() string {
	if o == nil || common.IsNil(o.PositionAmt) {
		var ret string
		return ret
	}
	return *o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetPositionAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionAmt) {
		return nil, false
	}
	return o.PositionAmt, true
}

// HasPositionAmt returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasPositionAmt() bool {
	if o != nil && !common.IsNil(o.PositionAmt) {
		return true
	}

	return false
}

// SetPositionAmt gets a reference to the given string and assigns it to the PositionAmt field.
func (o *PositionInformationV3ResponseInner) SetPositionAmt(v string) {
	o.PositionAmt = &v
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetEntryPrice() string {
	if o == nil || common.IsNil(o.EntryPrice) {
		var ret string
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetEntryPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasEntryPrice() bool {
	if o != nil && !common.IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given string and assigns it to the EntryPrice field.
func (o *PositionInformationV3ResponseInner) SetEntryPrice(v string) {
	o.EntryPrice = &v
}

// GetBreakEvenPrice returns the BreakEvenPrice field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetBreakEvenPrice() string {
	if o == nil || common.IsNil(o.BreakEvenPrice) {
		var ret string
		return ret
	}
	return *o.BreakEvenPrice
}

// GetBreakEvenPriceOk returns a tuple with the BreakEvenPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetBreakEvenPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.BreakEvenPrice) {
		return nil, false
	}
	return o.BreakEvenPrice, true
}

// HasBreakEvenPrice returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasBreakEvenPrice() bool {
	if o != nil && !common.IsNil(o.BreakEvenPrice) {
		return true
	}

	return false
}

// SetBreakEvenPrice gets a reference to the given string and assigns it to the BreakEvenPrice field.
func (o *PositionInformationV3ResponseInner) SetBreakEvenPrice(v string) {
	o.BreakEvenPrice = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetMarkPrice() string {
	if o == nil || common.IsNil(o.MarkPrice) {
		var ret string
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetMarkPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasMarkPrice() bool {
	if o != nil && !common.IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given string and assigns it to the MarkPrice field.
func (o *PositionInformationV3ResponseInner) SetMarkPrice(v string) {
	o.MarkPrice = &v
}

// GetUnRealizedProfit returns the UnRealizedProfit field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetUnRealizedProfit() string {
	if o == nil || common.IsNil(o.UnRealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnRealizedProfit
}

// GetUnRealizedProfitOk returns a tuple with the UnRealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetUnRealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnRealizedProfit) {
		return nil, false
	}
	return o.UnRealizedProfit, true
}

// HasUnRealizedProfit returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasUnRealizedProfit() bool {
	if o != nil && !common.IsNil(o.UnRealizedProfit) {
		return true
	}

	return false
}

// SetUnRealizedProfit gets a reference to the given string and assigns it to the UnRealizedProfit field.
func (o *PositionInformationV3ResponseInner) SetUnRealizedProfit(v string) {
	o.UnRealizedProfit = &v
}

// GetLiquidationPrice returns the LiquidationPrice field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetLiquidationPrice() string {
	if o == nil || common.IsNil(o.LiquidationPrice) {
		var ret string
		return ret
	}
	return *o.LiquidationPrice
}

// GetLiquidationPriceOk returns a tuple with the LiquidationPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetLiquidationPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiquidationPrice) {
		return nil, false
	}
	return o.LiquidationPrice, true
}

// HasLiquidationPrice returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasLiquidationPrice() bool {
	if o != nil && !common.IsNil(o.LiquidationPrice) {
		return true
	}

	return false
}

// SetLiquidationPrice gets a reference to the given string and assigns it to the LiquidationPrice field.
func (o *PositionInformationV3ResponseInner) SetLiquidationPrice(v string) {
	o.LiquidationPrice = &v
}

// GetIsolatedMargin returns the IsolatedMargin field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetIsolatedMargin() string {
	if o == nil || common.IsNil(o.IsolatedMargin) {
		var ret string
		return ret
	}
	return *o.IsolatedMargin
}

// GetIsolatedMarginOk returns a tuple with the IsolatedMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetIsolatedMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsolatedMargin) {
		return nil, false
	}
	return o.IsolatedMargin, true
}

// HasIsolatedMargin returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasIsolatedMargin() bool {
	if o != nil && !common.IsNil(o.IsolatedMargin) {
		return true
	}

	return false
}

// SetIsolatedMargin gets a reference to the given string and assigns it to the IsolatedMargin field.
func (o *PositionInformationV3ResponseInner) SetIsolatedMargin(v string) {
	o.IsolatedMargin = &v
}

// GetNotional returns the Notional field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetNotional() string {
	if o == nil || common.IsNil(o.Notional) {
		var ret string
		return ret
	}
	return *o.Notional
}

// GetNotionalOk returns a tuple with the Notional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Notional) {
		return nil, false
	}
	return o.Notional, true
}

// HasNotional returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasNotional() bool {
	if o != nil && !common.IsNil(o.Notional) {
		return true
	}

	return false
}

// SetNotional gets a reference to the given string and assigns it to the Notional field.
func (o *PositionInformationV3ResponseInner) SetNotional(v string) {
	o.Notional = &v
}

// GetMarginAsset returns the MarginAsset field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetMarginAsset() string {
	if o == nil || common.IsNil(o.MarginAsset) {
		var ret string
		return ret
	}
	return *o.MarginAsset
}

// GetMarginAssetOk returns a tuple with the MarginAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetMarginAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginAsset) {
		return nil, false
	}
	return o.MarginAsset, true
}

// HasMarginAsset returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasMarginAsset() bool {
	if o != nil && !common.IsNil(o.MarginAsset) {
		return true
	}

	return false
}

// SetMarginAsset gets a reference to the given string and assigns it to the MarginAsset field.
func (o *PositionInformationV3ResponseInner) SetMarginAsset(v string) {
	o.MarginAsset = &v
}

// GetIsolatedWallet returns the IsolatedWallet field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetIsolatedWallet() string {
	if o == nil || common.IsNil(o.IsolatedWallet) {
		var ret string
		return ret
	}
	return *o.IsolatedWallet
}

// GetIsolatedWalletOk returns a tuple with the IsolatedWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetIsolatedWalletOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsolatedWallet) {
		return nil, false
	}
	return o.IsolatedWallet, true
}

// HasIsolatedWallet returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasIsolatedWallet() bool {
	if o != nil && !common.IsNil(o.IsolatedWallet) {
		return true
	}

	return false
}

// SetIsolatedWallet gets a reference to the given string and assigns it to the IsolatedWallet field.
func (o *PositionInformationV3ResponseInner) SetIsolatedWallet(v string) {
	o.IsolatedWallet = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetInitialMargin() string {
	if o == nil || common.IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasInitialMargin() bool {
	if o != nil && !common.IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *PositionInformationV3ResponseInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetMaintMargin returns the MaintMargin field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetMaintMargin() string {
	if o == nil || common.IsNil(o.MaintMargin) {
		var ret string
		return ret
	}
	return *o.MaintMargin
}

// GetMaintMarginOk returns a tuple with the MaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaintMargin) {
		return nil, false
	}
	return o.MaintMargin, true
}

// HasMaintMargin returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasMaintMargin() bool {
	if o != nil && !common.IsNil(o.MaintMargin) {
		return true
	}

	return false
}

// SetMaintMargin gets a reference to the given string and assigns it to the MaintMargin field.
func (o *PositionInformationV3ResponseInner) SetMaintMargin(v string) {
	o.MaintMargin = &v
}

// GetPositionInitialMargin returns the PositionInitialMargin field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetPositionInitialMargin() string {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.PositionInitialMargin
}

// GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		return nil, false
	}
	return o.PositionInitialMargin, true
}

// HasPositionInitialMargin returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.PositionInitialMargin) {
		return true
	}

	return false
}

// SetPositionInitialMargin gets a reference to the given string and assigns it to the PositionInitialMargin field.
func (o *PositionInformationV3ResponseInner) SetPositionInitialMargin(v string) {
	o.PositionInitialMargin = &v
}

// GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.OpenOrderInitialMargin
}

// GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		return nil, false
	}
	return o.OpenOrderInitialMargin, true
}

// HasOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.OpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetOpenOrderInitialMargin gets a reference to the given string and assigns it to the OpenOrderInitialMargin field.
func (o *PositionInformationV3ResponseInner) SetOpenOrderInitialMargin(v string) {
	o.OpenOrderInitialMargin = &v
}

// GetAdl returns the Adl field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetAdl() int64 {
	if o == nil || common.IsNil(o.Adl) {
		var ret int64
		return ret
	}
	return *o.Adl
}

// GetAdlOk returns a tuple with the Adl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetAdlOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Adl) {
		return nil, false
	}
	return o.Adl, true
}

// HasAdl returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasAdl() bool {
	if o != nil && !common.IsNil(o.Adl) {
		return true
	}

	return false
}

// SetAdl gets a reference to the given int64 and assigns it to the Adl field.
func (o *PositionInformationV3ResponseInner) SetAdl(v int64) {
	o.Adl = &v
}

// GetBidNotional returns the BidNotional field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetBidNotional() string {
	if o == nil || common.IsNil(o.BidNotional) {
		var ret string
		return ret
	}
	return *o.BidNotional
}

// GetBidNotionalOk returns a tuple with the BidNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetBidNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidNotional) {
		return nil, false
	}
	return o.BidNotional, true
}

// HasBidNotional returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasBidNotional() bool {
	if o != nil && !common.IsNil(o.BidNotional) {
		return true
	}

	return false
}

// SetBidNotional gets a reference to the given string and assigns it to the BidNotional field.
func (o *PositionInformationV3ResponseInner) SetBidNotional(v string) {
	o.BidNotional = &v
}

// GetAskNotional returns the AskNotional field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetAskNotional() string {
	if o == nil || common.IsNil(o.AskNotional) {
		var ret string
		return ret
	}
	return *o.AskNotional
}

// GetAskNotionalOk returns a tuple with the AskNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetAskNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskNotional) {
		return nil, false
	}
	return o.AskNotional, true
}

// HasAskNotional returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasAskNotional() bool {
	if o != nil && !common.IsNil(o.AskNotional) {
		return true
	}

	return false
}

// SetAskNotional gets a reference to the given string and assigns it to the AskNotional field.
func (o *PositionInformationV3ResponseInner) SetAskNotional(v string) {
	o.AskNotional = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *PositionInformationV3ResponseInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV3ResponseInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *PositionInformationV3ResponseInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *PositionInformationV3ResponseInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o PositionInformationV3ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PositionInformationV3ResponseInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.EntryPrice) {
		toSerialize["entryPrice"] = o.EntryPrice
	}
	if !common.IsNil(o.BreakEvenPrice) {
		toSerialize["breakEvenPrice"] = o.BreakEvenPrice
	}
	if !common.IsNil(o.MarkPrice) {
		toSerialize["markPrice"] = o.MarkPrice
	}
	if !common.IsNil(o.UnRealizedProfit) {
		toSerialize["unRealizedProfit"] = o.UnRealizedProfit
	}
	if !common.IsNil(o.LiquidationPrice) {
		toSerialize["liquidationPrice"] = o.LiquidationPrice
	}
	if !common.IsNil(o.IsolatedMargin) {
		toSerialize["isolatedMargin"] = o.IsolatedMargin
	}
	if !common.IsNil(o.Notional) {
		toSerialize["notional"] = o.Notional
	}
	if !common.IsNil(o.MarginAsset) {
		toSerialize["marginAsset"] = o.MarginAsset
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
	if !common.IsNil(o.PositionInitialMargin) {
		toSerialize["positionInitialMargin"] = o.PositionInitialMargin
	}
	if !common.IsNil(o.OpenOrderInitialMargin) {
		toSerialize["openOrderInitialMargin"] = o.OpenOrderInitialMargin
	}
	if !common.IsNil(o.Adl) {
		toSerialize["adl"] = o.Adl
	}
	if !common.IsNil(o.BidNotional) {
		toSerialize["bidNotional"] = o.BidNotional
	}
	if !common.IsNil(o.AskNotional) {
		toSerialize["askNotional"] = o.AskNotional
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PositionInformationV3ResponseInner) UnmarshalJSON(data []byte) (err error) {
	varPositionInformationV3ResponseInner := _PositionInformationV3ResponseInner{}

	err = json.Unmarshal(data, &varPositionInformationV3ResponseInner)

	if err != nil {
		return err
	}

	*o = PositionInformationV3ResponseInner(varPositionInformationV3ResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "positionAmt")
		delete(additionalProperties, "entryPrice")
		delete(additionalProperties, "breakEvenPrice")
		delete(additionalProperties, "markPrice")
		delete(additionalProperties, "unRealizedProfit")
		delete(additionalProperties, "liquidationPrice")
		delete(additionalProperties, "isolatedMargin")
		delete(additionalProperties, "notional")
		delete(additionalProperties, "marginAsset")
		delete(additionalProperties, "isolatedWallet")
		delete(additionalProperties, "initialMargin")
		delete(additionalProperties, "maintMargin")
		delete(additionalProperties, "positionInitialMargin")
		delete(additionalProperties, "openOrderInitialMargin")
		delete(additionalProperties, "adl")
		delete(additionalProperties, "bidNotional")
		delete(additionalProperties, "askNotional")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePositionInformationV3ResponseInner struct {
	value *PositionInformationV3ResponseInner
	isSet bool
}

func (v NullablePositionInformationV3ResponseInner) Get() *PositionInformationV3ResponseInner {
	return v.value
}

func (v *NullablePositionInformationV3ResponseInner) Set(val *PositionInformationV3ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePositionInformationV3ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePositionInformationV3ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositionInformationV3ResponseInner(val *PositionInformationV3ResponseInner) *NullablePositionInformationV3ResponseInner {
	return &NullablePositionInformationV3ResponseInner{value: val, isSet: true}
}

func (v NullablePositionInformationV3ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositionInformationV3ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
