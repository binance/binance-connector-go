/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ExchangeInformationResponseOptionSymbolsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeInformationResponseOptionSymbolsInner{}

// ExchangeInformationResponseOptionSymbolsInner struct for ExchangeInformationResponseOptionSymbolsInner
type ExchangeInformationResponseOptionSymbolsInner struct {
	ExpiryDate           *int64                                                      `json:"expiryDate,omitempty"`
	Filters              []ExchangeInformationResponseOptionSymbolsInnerFiltersInner `json:"filters,omitempty"`
	Symbol               *string                                                     `json:"symbol,omitempty"`
	Side                 *string                                                     `json:"side,omitempty"`
	StrikePrice          *string                                                     `json:"strikePrice,omitempty"`
	Underlying           *string                                                     `json:"underlying,omitempty"`
	Unit                 *int64                                                      `json:"unit,omitempty"`
	MakerFeeRate         *string                                                     `json:"makerFeeRate,omitempty"`
	TakerFeeRate         *string                                                     `json:"takerFeeRate,omitempty"`
	LiquidationFeeRate   *string                                                     `json:"liquidationFeeRate,omitempty"`
	MinQty               *string                                                     `json:"minQty,omitempty"`
	MaxQty               *string                                                     `json:"maxQty,omitempty"`
	InitialMargin        *string                                                     `json:"initialMargin,omitempty"`
	MaintenanceMargin    *string                                                     `json:"maintenanceMargin,omitempty"`
	MinInitialMargin     *string                                                     `json:"minInitialMargin,omitempty"`
	MinMaintenanceMargin *string                                                     `json:"minMaintenanceMargin,omitempty"`
	PriceScale           *int64                                                      `json:"priceScale,omitempty"`
	QuantityScale        *int64                                                      `json:"quantityScale,omitempty"`
	QuoteAsset           *string                                                     `json:"quoteAsset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangeInformationResponseOptionSymbolsInner ExchangeInformationResponseOptionSymbolsInner

// NewExchangeInformationResponseOptionSymbolsInner instantiates a new ExchangeInformationResponseOptionSymbolsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeInformationResponseOptionSymbolsInner() *ExchangeInformationResponseOptionSymbolsInner {
	this := ExchangeInformationResponseOptionSymbolsInner{}
	return &this
}

// NewExchangeInformationResponseOptionSymbolsInnerWithDefaults instantiates a new ExchangeInformationResponseOptionSymbolsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeInformationResponseOptionSymbolsInnerWithDefaults() *ExchangeInformationResponseOptionSymbolsInner {
	this := ExchangeInformationResponseOptionSymbolsInner{}
	return &this
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetExpiryDate() int64 {
	if o == nil || common.IsNil(o.ExpiryDate) {
		var ret int64
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetExpiryDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExpiryDate) {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasExpiryDate() bool {
	if o != nil && !common.IsNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given int64 and assigns it to the ExpiryDate field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetExpiryDate(v int64) {
	o.ExpiryDate = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetFilters() []ExchangeInformationResponseOptionSymbolsInnerFiltersInner {
	if o == nil || common.IsNil(o.Filters) {
		var ret []ExchangeInformationResponseOptionSymbolsInnerFiltersInner
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetFiltersOk() ([]ExchangeInformationResponseOptionSymbolsInnerFiltersInner, bool) {
	if o == nil || common.IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasFilters() bool {
	if o != nil && !common.IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []ExchangeInformationResponseOptionSymbolsInnerFiltersInner and assigns it to the Filters field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetFilters(v []ExchangeInformationResponseOptionSymbolsInnerFiltersInner) {
	o.Filters = v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetSide(v string) {
	o.Side = &v
}

// GetStrikePrice returns the StrikePrice field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetStrikePrice() string {
	if o == nil || common.IsNil(o.StrikePrice) {
		var ret string
		return ret
	}
	return *o.StrikePrice
}

// GetStrikePriceOk returns a tuple with the StrikePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetStrikePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StrikePrice) {
		return nil, false
	}
	return o.StrikePrice, true
}

// HasStrikePrice returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasStrikePrice() bool {
	if o != nil && !common.IsNil(o.StrikePrice) {
		return true
	}

	return false
}

// SetStrikePrice gets a reference to the given string and assigns it to the StrikePrice field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetStrikePrice(v string) {
	o.StrikePrice = &v
}

// GetUnderlying returns the Underlying field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetUnderlying() string {
	if o == nil || common.IsNil(o.Underlying) {
		var ret string
		return ret
	}
	return *o.Underlying
}

// GetUnderlyingOk returns a tuple with the Underlying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetUnderlyingOk() (*string, bool) {
	if o == nil || common.IsNil(o.Underlying) {
		return nil, false
	}
	return o.Underlying, true
}

// HasUnderlying returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasUnderlying() bool {
	if o != nil && !common.IsNil(o.Underlying) {
		return true
	}

	return false
}

// SetUnderlying gets a reference to the given string and assigns it to the Underlying field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetUnderlying(v string) {
	o.Underlying = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetUnit() int64 {
	if o == nil || common.IsNil(o.Unit) {
		var ret int64
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetUnitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasUnit() bool {
	if o != nil && !common.IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given int64 and assigns it to the Unit field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetUnit(v int64) {
	o.Unit = &v
}

// GetMakerFeeRate returns the MakerFeeRate field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetMakerFeeRate() string {
	if o == nil || common.IsNil(o.MakerFeeRate) {
		var ret string
		return ret
	}
	return *o.MakerFeeRate
}

// GetMakerFeeRateOk returns a tuple with the MakerFeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetMakerFeeRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.MakerFeeRate) {
		return nil, false
	}
	return o.MakerFeeRate, true
}

// HasMakerFeeRate returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasMakerFeeRate() bool {
	if o != nil && !common.IsNil(o.MakerFeeRate) {
		return true
	}

	return false
}

// SetMakerFeeRate gets a reference to the given string and assigns it to the MakerFeeRate field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetMakerFeeRate(v string) {
	o.MakerFeeRate = &v
}

// GetTakerFeeRate returns the TakerFeeRate field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetTakerFeeRate() string {
	if o == nil || common.IsNil(o.TakerFeeRate) {
		var ret string
		return ret
	}
	return *o.TakerFeeRate
}

// GetTakerFeeRateOk returns a tuple with the TakerFeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetTakerFeeRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerFeeRate) {
		return nil, false
	}
	return o.TakerFeeRate, true
}

// HasTakerFeeRate returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasTakerFeeRate() bool {
	if o != nil && !common.IsNil(o.TakerFeeRate) {
		return true
	}

	return false
}

// SetTakerFeeRate gets a reference to the given string and assigns it to the TakerFeeRate field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetTakerFeeRate(v string) {
	o.TakerFeeRate = &v
}

// GetLiquidationFeeRate returns the LiquidationFeeRate field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetLiquidationFeeRate() string {
	if o == nil || common.IsNil(o.LiquidationFeeRate) {
		var ret string
		return ret
	}
	return *o.LiquidationFeeRate
}

// GetLiquidationFeeRateOk returns a tuple with the LiquidationFeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetLiquidationFeeRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiquidationFeeRate) {
		return nil, false
	}
	return o.LiquidationFeeRate, true
}

// HasLiquidationFeeRate returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasLiquidationFeeRate() bool {
	if o != nil && !common.IsNil(o.LiquidationFeeRate) {
		return true
	}

	return false
}

// SetLiquidationFeeRate gets a reference to the given string and assigns it to the LiquidationFeeRate field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetLiquidationFeeRate(v string) {
	o.LiquidationFeeRate = &v
}

// GetMinQty returns the MinQty field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetMinQty() string {
	if o == nil || common.IsNil(o.MinQty) {
		var ret string
		return ret
	}
	return *o.MinQty
}

// GetMinQtyOk returns a tuple with the MinQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetMinQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinQty) {
		return nil, false
	}
	return o.MinQty, true
}

// HasMinQty returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasMinQty() bool {
	if o != nil && !common.IsNil(o.MinQty) {
		return true
	}

	return false
}

// SetMinQty gets a reference to the given string and assigns it to the MinQty field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetMinQty(v string) {
	o.MinQty = &v
}

// GetMaxQty returns the MaxQty field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetMaxQty() string {
	if o == nil || common.IsNil(o.MaxQty) {
		var ret string
		return ret
	}
	return *o.MaxQty
}

// GetMaxQtyOk returns a tuple with the MaxQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetMaxQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxQty) {
		return nil, false
	}
	return o.MaxQty, true
}

// HasMaxQty returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasMaxQty() bool {
	if o != nil && !common.IsNil(o.MaxQty) {
		return true
	}

	return false
}

// SetMaxQty gets a reference to the given string and assigns it to the MaxQty field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetMaxQty(v string) {
	o.MaxQty = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetInitialMargin() string {
	if o == nil || common.IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasInitialMargin() bool {
	if o != nil && !common.IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetMaintenanceMargin returns the MaintenanceMargin field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetMaintenanceMargin() string {
	if o == nil || common.IsNil(o.MaintenanceMargin) {
		var ret string
		return ret
	}
	return *o.MaintenanceMargin
}

// GetMaintenanceMarginOk returns a tuple with the MaintenanceMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetMaintenanceMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaintenanceMargin) {
		return nil, false
	}
	return o.MaintenanceMargin, true
}

// HasMaintenanceMargin returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasMaintenanceMargin() bool {
	if o != nil && !common.IsNil(o.MaintenanceMargin) {
		return true
	}

	return false
}

// SetMaintenanceMargin gets a reference to the given string and assigns it to the MaintenanceMargin field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetMaintenanceMargin(v string) {
	o.MaintenanceMargin = &v
}

// GetMinInitialMargin returns the MinInitialMargin field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetMinInitialMargin() string {
	if o == nil || common.IsNil(o.MinInitialMargin) {
		var ret string
		return ret
	}
	return *o.MinInitialMargin
}

// GetMinInitialMarginOk returns a tuple with the MinInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetMinInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinInitialMargin) {
		return nil, false
	}
	return o.MinInitialMargin, true
}

// HasMinInitialMargin returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasMinInitialMargin() bool {
	if o != nil && !common.IsNil(o.MinInitialMargin) {
		return true
	}

	return false
}

// SetMinInitialMargin gets a reference to the given string and assigns it to the MinInitialMargin field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetMinInitialMargin(v string) {
	o.MinInitialMargin = &v
}

// GetMinMaintenanceMargin returns the MinMaintenanceMargin field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetMinMaintenanceMargin() string {
	if o == nil || common.IsNil(o.MinMaintenanceMargin) {
		var ret string
		return ret
	}
	return *o.MinMaintenanceMargin
}

// GetMinMaintenanceMarginOk returns a tuple with the MinMaintenanceMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetMinMaintenanceMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinMaintenanceMargin) {
		return nil, false
	}
	return o.MinMaintenanceMargin, true
}

// HasMinMaintenanceMargin returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasMinMaintenanceMargin() bool {
	if o != nil && !common.IsNil(o.MinMaintenanceMargin) {
		return true
	}

	return false
}

// SetMinMaintenanceMargin gets a reference to the given string and assigns it to the MinMaintenanceMargin field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetMinMaintenanceMargin(v string) {
	o.MinMaintenanceMargin = &v
}

// GetPriceScale returns the PriceScale field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetPriceScale() int64 {
	if o == nil || common.IsNil(o.PriceScale) {
		var ret int64
		return ret
	}
	return *o.PriceScale
}

// GetPriceScaleOk returns a tuple with the PriceScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetPriceScaleOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PriceScale) {
		return nil, false
	}
	return o.PriceScale, true
}

// HasPriceScale returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasPriceScale() bool {
	if o != nil && !common.IsNil(o.PriceScale) {
		return true
	}

	return false
}

// SetPriceScale gets a reference to the given int64 and assigns it to the PriceScale field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetPriceScale(v int64) {
	o.PriceScale = &v
}

// GetQuantityScale returns the QuantityScale field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetQuantityScale() int64 {
	if o == nil || common.IsNil(o.QuantityScale) {
		var ret int64
		return ret
	}
	return *o.QuantityScale
}

// GetQuantityScaleOk returns a tuple with the QuantityScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetQuantityScaleOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QuantityScale) {
		return nil, false
	}
	return o.QuantityScale, true
}

// HasQuantityScale returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasQuantityScale() bool {
	if o != nil && !common.IsNil(o.QuantityScale) {
		return true
	}

	return false
}

// SetQuantityScale gets a reference to the given int64 and assigns it to the QuantityScale field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetQuantityScale(v int64) {
	o.QuantityScale = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetQuoteAsset() string {
	if o == nil || common.IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) GetQuoteAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionSymbolsInner) HasQuoteAsset() bool {
	if o != nil && !common.IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *ExchangeInformationResponseOptionSymbolsInner) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

func (o ExchangeInformationResponseOptionSymbolsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeInformationResponseOptionSymbolsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ExpiryDate) {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if !common.IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.StrikePrice) {
		toSerialize["strikePrice"] = o.StrikePrice
	}
	if !common.IsNil(o.Underlying) {
		toSerialize["underlying"] = o.Underlying
	}
	if !common.IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !common.IsNil(o.MakerFeeRate) {
		toSerialize["makerFeeRate"] = o.MakerFeeRate
	}
	if !common.IsNil(o.TakerFeeRate) {
		toSerialize["takerFeeRate"] = o.TakerFeeRate
	}
	if !common.IsNil(o.LiquidationFeeRate) {
		toSerialize["liquidationFeeRate"] = o.LiquidationFeeRate
	}
	if !common.IsNil(o.MinQty) {
		toSerialize["minQty"] = o.MinQty
	}
	if !common.IsNil(o.MaxQty) {
		toSerialize["maxQty"] = o.MaxQty
	}
	if !common.IsNil(o.InitialMargin) {
		toSerialize["initialMargin"] = o.InitialMargin
	}
	if !common.IsNil(o.MaintenanceMargin) {
		toSerialize["maintenanceMargin"] = o.MaintenanceMargin
	}
	if !common.IsNil(o.MinInitialMargin) {
		toSerialize["minInitialMargin"] = o.MinInitialMargin
	}
	if !common.IsNil(o.MinMaintenanceMargin) {
		toSerialize["minMaintenanceMargin"] = o.MinMaintenanceMargin
	}
	if !common.IsNil(o.PriceScale) {
		toSerialize["priceScale"] = o.PriceScale
	}
	if !common.IsNil(o.QuantityScale) {
		toSerialize["quantityScale"] = o.QuantityScale
	}
	if !common.IsNil(o.QuoteAsset) {
		toSerialize["quoteAsset"] = o.QuoteAsset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExchangeInformationResponseOptionSymbolsInner) UnmarshalJSON(data []byte) (err error) {
	varExchangeInformationResponseOptionSymbolsInner := _ExchangeInformationResponseOptionSymbolsInner{}

	err = json.Unmarshal(data, &varExchangeInformationResponseOptionSymbolsInner)

	if err != nil {
		return err
	}

	*o = ExchangeInformationResponseOptionSymbolsInner(varExchangeInformationResponseOptionSymbolsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expiryDate")
		delete(additionalProperties, "filters")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "side")
		delete(additionalProperties, "strikePrice")
		delete(additionalProperties, "underlying")
		delete(additionalProperties, "unit")
		delete(additionalProperties, "makerFeeRate")
		delete(additionalProperties, "takerFeeRate")
		delete(additionalProperties, "liquidationFeeRate")
		delete(additionalProperties, "minQty")
		delete(additionalProperties, "maxQty")
		delete(additionalProperties, "initialMargin")
		delete(additionalProperties, "maintenanceMargin")
		delete(additionalProperties, "minInitialMargin")
		delete(additionalProperties, "minMaintenanceMargin")
		delete(additionalProperties, "priceScale")
		delete(additionalProperties, "quantityScale")
		delete(additionalProperties, "quoteAsset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeInformationResponseOptionSymbolsInner struct {
	value *ExchangeInformationResponseOptionSymbolsInner
	isSet bool
}

func (v NullableExchangeInformationResponseOptionSymbolsInner) Get() *ExchangeInformationResponseOptionSymbolsInner {
	return v.value
}

func (v *NullableExchangeInformationResponseOptionSymbolsInner) Set(val *ExchangeInformationResponseOptionSymbolsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInformationResponseOptionSymbolsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInformationResponseOptionSymbolsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInformationResponseOptionSymbolsInner(val *ExchangeInformationResponseOptionSymbolsInner) *NullableExchangeInformationResponseOptionSymbolsInner {
	return &NullableExchangeInformationResponseOptionSymbolsInner{value: val, isSet: true}
}

func (v NullableExchangeInformationResponseOptionSymbolsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInformationResponseOptionSymbolsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
