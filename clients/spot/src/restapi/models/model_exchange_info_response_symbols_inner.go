/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ExchangeInfoResponseSymbolsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeInfoResponseSymbolsInner{}

// ExchangeInfoResponseSymbolsInner struct for ExchangeInfoResponseSymbolsInner
type ExchangeInfoResponseSymbolsInner struct {
	Symbol                          *string         `json:"symbol,omitempty"`
	Status                          *string         `json:"status,omitempty"`
	BaseAsset                       *string         `json:"baseAsset,omitempty"`
	BaseAssetPrecision              *int64          `json:"baseAssetPrecision,omitempty"`
	QuoteAsset                      *string         `json:"quoteAsset,omitempty"`
	QuotePrecision                  *int64          `json:"quotePrecision,omitempty"`
	QuoteAssetPrecision             *int64          `json:"quoteAssetPrecision,omitempty"`
	BaseCommissionPrecision         *int64          `json:"baseCommissionPrecision,omitempty"`
	QuoteCommissionPrecision        *int64          `json:"quoteCommissionPrecision,omitempty"`
	OrderTypes                      []string        `json:"orderTypes,omitempty"`
	IcebergAllowed                  *bool           `json:"icebergAllowed,omitempty"`
	OcoAllowed                      *bool           `json:"ocoAllowed,omitempty"`
	OtoAllowed                      *bool           `json:"otoAllowed,omitempty"`
	OpoAllowed                      *bool           `json:"opoAllowed,omitempty"`
	QuoteOrderQtyMarketAllowed      *bool           `json:"quoteOrderQtyMarketAllowed,omitempty"`
	AllowTrailingStop               *bool           `json:"allowTrailingStop,omitempty"`
	CancelReplaceAllowed            *bool           `json:"cancelReplaceAllowed,omitempty"`
	AmendAllowed                    *bool           `json:"amendAllowed,omitempty"`
	PegInstructionsAllowed          *bool           `json:"pegInstructionsAllowed,omitempty"`
	IsSpotTradingAllowed            *bool           `json:"isSpotTradingAllowed,omitempty"`
	IsMarginTradingAllowed          *bool           `json:"isMarginTradingAllowed,omitempty"`
	Filters                         []SymbolFilters `json:"filters,omitempty"`
	Permissions                     []string        `json:"permissions,omitempty"`
	PermissionSets                  [][]string      `json:"permissionSets,omitempty"`
	DefaultSelfTradePreventionMode  *string         `json:"defaultSelfTradePreventionMode,omitempty"`
	AllowedSelfTradePreventionModes []string        `json:"allowedSelfTradePreventionModes,omitempty"`
	AdditionalProperties            map[string]interface{}
}

type _ExchangeInfoResponseSymbolsInner ExchangeInfoResponseSymbolsInner

// NewExchangeInfoResponseSymbolsInner instantiates a new ExchangeInfoResponseSymbolsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeInfoResponseSymbolsInner() *ExchangeInfoResponseSymbolsInner {
	this := ExchangeInfoResponseSymbolsInner{}
	return &this
}

// NewExchangeInfoResponseSymbolsInnerWithDefaults instantiates a new ExchangeInfoResponseSymbolsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeInfoResponseSymbolsInnerWithDefaults() *ExchangeInfoResponseSymbolsInner {
	this := ExchangeInfoResponseSymbolsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ExchangeInfoResponseSymbolsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ExchangeInfoResponseSymbolsInner) SetStatus(v string) {
	o.Status = &v
}

// GetBaseAsset returns the BaseAsset field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetBaseAsset() string {
	if o == nil || common.IsNil(o.BaseAsset) {
		var ret string
		return ret
	}
	return *o.BaseAsset
}

// GetBaseAssetOk returns a tuple with the BaseAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetBaseAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.BaseAsset) {
		return nil, false
	}
	return o.BaseAsset, true
}

// HasBaseAsset returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasBaseAsset() bool {
	if o != nil && !common.IsNil(o.BaseAsset) {
		return true
	}

	return false
}

// SetBaseAsset gets a reference to the given string and assigns it to the BaseAsset field.
func (o *ExchangeInfoResponseSymbolsInner) SetBaseAsset(v string) {
	o.BaseAsset = &v
}

// GetBaseAssetPrecision returns the BaseAssetPrecision field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetBaseAssetPrecision() int64 {
	if o == nil || common.IsNil(o.BaseAssetPrecision) {
		var ret int64
		return ret
	}
	return *o.BaseAssetPrecision
}

// GetBaseAssetPrecisionOk returns a tuple with the BaseAssetPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetBaseAssetPrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BaseAssetPrecision) {
		return nil, false
	}
	return o.BaseAssetPrecision, true
}

// HasBaseAssetPrecision returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasBaseAssetPrecision() bool {
	if o != nil && !common.IsNil(o.BaseAssetPrecision) {
		return true
	}

	return false
}

// SetBaseAssetPrecision gets a reference to the given int64 and assigns it to the BaseAssetPrecision field.
func (o *ExchangeInfoResponseSymbolsInner) SetBaseAssetPrecision(v int64) {
	o.BaseAssetPrecision = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetQuoteAsset() string {
	if o == nil || common.IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetQuoteAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasQuoteAsset() bool {
	if o != nil && !common.IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *ExchangeInfoResponseSymbolsInner) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

// GetQuotePrecision returns the QuotePrecision field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetQuotePrecision() int64 {
	if o == nil || common.IsNil(o.QuotePrecision) {
		var ret int64
		return ret
	}
	return *o.QuotePrecision
}

// GetQuotePrecisionOk returns a tuple with the QuotePrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetQuotePrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QuotePrecision) {
		return nil, false
	}
	return o.QuotePrecision, true
}

// HasQuotePrecision returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasQuotePrecision() bool {
	if o != nil && !common.IsNil(o.QuotePrecision) {
		return true
	}

	return false
}

// SetQuotePrecision gets a reference to the given int64 and assigns it to the QuotePrecision field.
func (o *ExchangeInfoResponseSymbolsInner) SetQuotePrecision(v int64) {
	o.QuotePrecision = &v
}

// GetQuoteAssetPrecision returns the QuoteAssetPrecision field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetQuoteAssetPrecision() int64 {
	if o == nil || common.IsNil(o.QuoteAssetPrecision) {
		var ret int64
		return ret
	}
	return *o.QuoteAssetPrecision
}

// GetQuoteAssetPrecisionOk returns a tuple with the QuoteAssetPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetQuoteAssetPrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QuoteAssetPrecision) {
		return nil, false
	}
	return o.QuoteAssetPrecision, true
}

// HasQuoteAssetPrecision returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasQuoteAssetPrecision() bool {
	if o != nil && !common.IsNil(o.QuoteAssetPrecision) {
		return true
	}

	return false
}

// SetQuoteAssetPrecision gets a reference to the given int64 and assigns it to the QuoteAssetPrecision field.
func (o *ExchangeInfoResponseSymbolsInner) SetQuoteAssetPrecision(v int64) {
	o.QuoteAssetPrecision = &v
}

// GetBaseCommissionPrecision returns the BaseCommissionPrecision field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetBaseCommissionPrecision() int64 {
	if o == nil || common.IsNil(o.BaseCommissionPrecision) {
		var ret int64
		return ret
	}
	return *o.BaseCommissionPrecision
}

// GetBaseCommissionPrecisionOk returns a tuple with the BaseCommissionPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetBaseCommissionPrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BaseCommissionPrecision) {
		return nil, false
	}
	return o.BaseCommissionPrecision, true
}

// HasBaseCommissionPrecision returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasBaseCommissionPrecision() bool {
	if o != nil && !common.IsNil(o.BaseCommissionPrecision) {
		return true
	}

	return false
}

// SetBaseCommissionPrecision gets a reference to the given int64 and assigns it to the BaseCommissionPrecision field.
func (o *ExchangeInfoResponseSymbolsInner) SetBaseCommissionPrecision(v int64) {
	o.BaseCommissionPrecision = &v
}

// GetQuoteCommissionPrecision returns the QuoteCommissionPrecision field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetQuoteCommissionPrecision() int64 {
	if o == nil || common.IsNil(o.QuoteCommissionPrecision) {
		var ret int64
		return ret
	}
	return *o.QuoteCommissionPrecision
}

// GetQuoteCommissionPrecisionOk returns a tuple with the QuoteCommissionPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetQuoteCommissionPrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QuoteCommissionPrecision) {
		return nil, false
	}
	return o.QuoteCommissionPrecision, true
}

// HasQuoteCommissionPrecision returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasQuoteCommissionPrecision() bool {
	if o != nil && !common.IsNil(o.QuoteCommissionPrecision) {
		return true
	}

	return false
}

// SetQuoteCommissionPrecision gets a reference to the given int64 and assigns it to the QuoteCommissionPrecision field.
func (o *ExchangeInfoResponseSymbolsInner) SetQuoteCommissionPrecision(v int64) {
	o.QuoteCommissionPrecision = &v
}

// GetOrderTypes returns the OrderTypes field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetOrderTypes() []string {
	if o == nil || common.IsNil(o.OrderTypes) {
		var ret []string
		return ret
	}
	return o.OrderTypes
}

// GetOrderTypesOk returns a tuple with the OrderTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetOrderTypesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.OrderTypes) {
		return nil, false
	}
	return o.OrderTypes, true
}

// HasOrderTypes returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasOrderTypes() bool {
	if o != nil && !common.IsNil(o.OrderTypes) {
		return true
	}

	return false
}

// SetOrderTypes gets a reference to the given []string and assigns it to the OrderTypes field.
func (o *ExchangeInfoResponseSymbolsInner) SetOrderTypes(v []string) {
	o.OrderTypes = v
}

// GetIcebergAllowed returns the IcebergAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetIcebergAllowed() bool {
	if o == nil || common.IsNil(o.IcebergAllowed) {
		var ret bool
		return ret
	}
	return *o.IcebergAllowed
}

// GetIcebergAllowedOk returns a tuple with the IcebergAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetIcebergAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IcebergAllowed) {
		return nil, false
	}
	return o.IcebergAllowed, true
}

// HasIcebergAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasIcebergAllowed() bool {
	if o != nil && !common.IsNil(o.IcebergAllowed) {
		return true
	}

	return false
}

// SetIcebergAllowed gets a reference to the given bool and assigns it to the IcebergAllowed field.
func (o *ExchangeInfoResponseSymbolsInner) SetIcebergAllowed(v bool) {
	o.IcebergAllowed = &v
}

// GetOcoAllowed returns the OcoAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetOcoAllowed() bool {
	if o == nil || common.IsNil(o.OcoAllowed) {
		var ret bool
		return ret
	}
	return *o.OcoAllowed
}

// GetOcoAllowedOk returns a tuple with the OcoAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetOcoAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.OcoAllowed) {
		return nil, false
	}
	return o.OcoAllowed, true
}

// HasOcoAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasOcoAllowed() bool {
	if o != nil && !common.IsNil(o.OcoAllowed) {
		return true
	}

	return false
}

// SetOcoAllowed gets a reference to the given bool and assigns it to the OcoAllowed field.
func (o *ExchangeInfoResponseSymbolsInner) SetOcoAllowed(v bool) {
	o.OcoAllowed = &v
}

// GetOtoAllowed returns the OtoAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetOtoAllowed() bool {
	if o == nil || common.IsNil(o.OtoAllowed) {
		var ret bool
		return ret
	}
	return *o.OtoAllowed
}

// GetOtoAllowedOk returns a tuple with the OtoAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetOtoAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.OtoAllowed) {
		return nil, false
	}
	return o.OtoAllowed, true
}

// HasOtoAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasOtoAllowed() bool {
	if o != nil && !common.IsNil(o.OtoAllowed) {
		return true
	}

	return false
}

// SetOtoAllowed gets a reference to the given bool and assigns it to the OtoAllowed field.
func (o *ExchangeInfoResponseSymbolsInner) SetOtoAllowed(v bool) {
	o.OtoAllowed = &v
}

// GetOpoAllowed returns the OpoAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetOpoAllowed() bool {
	if o == nil || common.IsNil(o.OpoAllowed) {
		var ret bool
		return ret
	}
	return *o.OpoAllowed
}

// GetOpoAllowedOk returns a tuple with the OpoAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetOpoAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.OpoAllowed) {
		return nil, false
	}
	return o.OpoAllowed, true
}

// HasOpoAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasOpoAllowed() bool {
	if o != nil && !common.IsNil(o.OpoAllowed) {
		return true
	}

	return false
}

// SetOpoAllowed gets a reference to the given bool and assigns it to the OpoAllowed field.
func (o *ExchangeInfoResponseSymbolsInner) SetOpoAllowed(v bool) {
	o.OpoAllowed = &v
}

// GetQuoteOrderQtyMarketAllowed returns the QuoteOrderQtyMarketAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetQuoteOrderQtyMarketAllowed() bool {
	if o == nil || common.IsNil(o.QuoteOrderQtyMarketAllowed) {
		var ret bool
		return ret
	}
	return *o.QuoteOrderQtyMarketAllowed
}

// GetQuoteOrderQtyMarketAllowedOk returns a tuple with the QuoteOrderQtyMarketAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetQuoteOrderQtyMarketAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.QuoteOrderQtyMarketAllowed) {
		return nil, false
	}
	return o.QuoteOrderQtyMarketAllowed, true
}

// HasQuoteOrderQtyMarketAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasQuoteOrderQtyMarketAllowed() bool {
	if o != nil && !common.IsNil(o.QuoteOrderQtyMarketAllowed) {
		return true
	}

	return false
}

// SetQuoteOrderQtyMarketAllowed gets a reference to the given bool and assigns it to the QuoteOrderQtyMarketAllowed field.
func (o *ExchangeInfoResponseSymbolsInner) SetQuoteOrderQtyMarketAllowed(v bool) {
	o.QuoteOrderQtyMarketAllowed = &v
}

// GetAllowTrailingStop returns the AllowTrailingStop field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetAllowTrailingStop() bool {
	if o == nil || common.IsNil(o.AllowTrailingStop) {
		var ret bool
		return ret
	}
	return *o.AllowTrailingStop
}

// GetAllowTrailingStopOk returns a tuple with the AllowTrailingStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetAllowTrailingStopOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AllowTrailingStop) {
		return nil, false
	}
	return o.AllowTrailingStop, true
}

// HasAllowTrailingStop returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasAllowTrailingStop() bool {
	if o != nil && !common.IsNil(o.AllowTrailingStop) {
		return true
	}

	return false
}

// SetAllowTrailingStop gets a reference to the given bool and assigns it to the AllowTrailingStop field.
func (o *ExchangeInfoResponseSymbolsInner) SetAllowTrailingStop(v bool) {
	o.AllowTrailingStop = &v
}

// GetCancelReplaceAllowed returns the CancelReplaceAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetCancelReplaceAllowed() bool {
	if o == nil || common.IsNil(o.CancelReplaceAllowed) {
		var ret bool
		return ret
	}
	return *o.CancelReplaceAllowed
}

// GetCancelReplaceAllowedOk returns a tuple with the CancelReplaceAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetCancelReplaceAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CancelReplaceAllowed) {
		return nil, false
	}
	return o.CancelReplaceAllowed, true
}

// HasCancelReplaceAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasCancelReplaceAllowed() bool {
	if o != nil && !common.IsNil(o.CancelReplaceAllowed) {
		return true
	}

	return false
}

// SetCancelReplaceAllowed gets a reference to the given bool and assigns it to the CancelReplaceAllowed field.
func (o *ExchangeInfoResponseSymbolsInner) SetCancelReplaceAllowed(v bool) {
	o.CancelReplaceAllowed = &v
}

// GetAmendAllowed returns the AmendAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetAmendAllowed() bool {
	if o == nil || common.IsNil(o.AmendAllowed) {
		var ret bool
		return ret
	}
	return *o.AmendAllowed
}

// GetAmendAllowedOk returns a tuple with the AmendAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetAmendAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AmendAllowed) {
		return nil, false
	}
	return o.AmendAllowed, true
}

// HasAmendAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasAmendAllowed() bool {
	if o != nil && !common.IsNil(o.AmendAllowed) {
		return true
	}

	return false
}

// SetAmendAllowed gets a reference to the given bool and assigns it to the AmendAllowed field.
func (o *ExchangeInfoResponseSymbolsInner) SetAmendAllowed(v bool) {
	o.AmendAllowed = &v
}

// GetPegInstructionsAllowed returns the PegInstructionsAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetPegInstructionsAllowed() bool {
	if o == nil || common.IsNil(o.PegInstructionsAllowed) {
		var ret bool
		return ret
	}
	return *o.PegInstructionsAllowed
}

// GetPegInstructionsAllowedOk returns a tuple with the PegInstructionsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetPegInstructionsAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PegInstructionsAllowed) {
		return nil, false
	}
	return o.PegInstructionsAllowed, true
}

// HasPegInstructionsAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasPegInstructionsAllowed() bool {
	if o != nil && !common.IsNil(o.PegInstructionsAllowed) {
		return true
	}

	return false
}

// SetPegInstructionsAllowed gets a reference to the given bool and assigns it to the PegInstructionsAllowed field.
func (o *ExchangeInfoResponseSymbolsInner) SetPegInstructionsAllowed(v bool) {
	o.PegInstructionsAllowed = &v
}

// GetIsSpotTradingAllowed returns the IsSpotTradingAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetIsSpotTradingAllowed() bool {
	if o == nil || common.IsNil(o.IsSpotTradingAllowed) {
		var ret bool
		return ret
	}
	return *o.IsSpotTradingAllowed
}

// GetIsSpotTradingAllowedOk returns a tuple with the IsSpotTradingAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetIsSpotTradingAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsSpotTradingAllowed) {
		return nil, false
	}
	return o.IsSpotTradingAllowed, true
}

// HasIsSpotTradingAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasIsSpotTradingAllowed() bool {
	if o != nil && !common.IsNil(o.IsSpotTradingAllowed) {
		return true
	}

	return false
}

// SetIsSpotTradingAllowed gets a reference to the given bool and assigns it to the IsSpotTradingAllowed field.
func (o *ExchangeInfoResponseSymbolsInner) SetIsSpotTradingAllowed(v bool) {
	o.IsSpotTradingAllowed = &v
}

// GetIsMarginTradingAllowed returns the IsMarginTradingAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetIsMarginTradingAllowed() bool {
	if o == nil || common.IsNil(o.IsMarginTradingAllowed) {
		var ret bool
		return ret
	}
	return *o.IsMarginTradingAllowed
}

// GetIsMarginTradingAllowedOk returns a tuple with the IsMarginTradingAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetIsMarginTradingAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsMarginTradingAllowed) {
		return nil, false
	}
	return o.IsMarginTradingAllowed, true
}

// HasIsMarginTradingAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasIsMarginTradingAllowed() bool {
	if o != nil && !common.IsNil(o.IsMarginTradingAllowed) {
		return true
	}

	return false
}

// SetIsMarginTradingAllowed gets a reference to the given bool and assigns it to the IsMarginTradingAllowed field.
func (o *ExchangeInfoResponseSymbolsInner) SetIsMarginTradingAllowed(v bool) {
	o.IsMarginTradingAllowed = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetFilters() []SymbolFilters {
	if o == nil || common.IsNil(o.Filters) {
		var ret []SymbolFilters
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetFiltersOk() ([]SymbolFilters, bool) {
	if o == nil || common.IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasFilters() bool {
	if o != nil && !common.IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []SymbolFilters and assigns it to the Filters field.
func (o *ExchangeInfoResponseSymbolsInner) SetFilters(v []SymbolFilters) {
	o.Filters = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetPermissions() []string {
	if o == nil || common.IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetPermissionsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasPermissions() bool {
	if o != nil && !common.IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *ExchangeInfoResponseSymbolsInner) SetPermissions(v []string) {
	o.Permissions = v
}

// GetPermissionSets returns the PermissionSets field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetPermissionSets() [][]string {
	if o == nil || common.IsNil(o.PermissionSets) {
		var ret [][]string
		return ret
	}
	return o.PermissionSets
}

// GetPermissionSetsOk returns a tuple with the PermissionSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetPermissionSetsOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.PermissionSets) {
		return nil, false
	}
	return o.PermissionSets, true
}

// HasPermissionSets returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasPermissionSets() bool {
	if o != nil && !common.IsNil(o.PermissionSets) {
		return true
	}

	return false
}

// SetPermissionSets gets a reference to the given [][]string and assigns it to the PermissionSets field.
func (o *ExchangeInfoResponseSymbolsInner) SetPermissionSets(v [][]string) {
	o.PermissionSets = v
}

// GetDefaultSelfTradePreventionMode returns the DefaultSelfTradePreventionMode field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetDefaultSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.DefaultSelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.DefaultSelfTradePreventionMode
}

// GetDefaultSelfTradePreventionModeOk returns a tuple with the DefaultSelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetDefaultSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.DefaultSelfTradePreventionMode) {
		return nil, false
	}
	return o.DefaultSelfTradePreventionMode, true
}

// HasDefaultSelfTradePreventionMode returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasDefaultSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.DefaultSelfTradePreventionMode) {
		return true
	}

	return false
}

// SetDefaultSelfTradePreventionMode gets a reference to the given string and assigns it to the DefaultSelfTradePreventionMode field.
func (o *ExchangeInfoResponseSymbolsInner) SetDefaultSelfTradePreventionMode(v string) {
	o.DefaultSelfTradePreventionMode = &v
}

// GetAllowedSelfTradePreventionModes returns the AllowedSelfTradePreventionModes field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetAllowedSelfTradePreventionModes() []string {
	if o == nil || common.IsNil(o.AllowedSelfTradePreventionModes) {
		var ret []string
		return ret
	}
	return o.AllowedSelfTradePreventionModes
}

// GetAllowedSelfTradePreventionModesOk returns a tuple with the AllowedSelfTradePreventionModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetAllowedSelfTradePreventionModesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AllowedSelfTradePreventionModes) {
		return nil, false
	}
	return o.AllowedSelfTradePreventionModes, true
}

// HasAllowedSelfTradePreventionModes returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasAllowedSelfTradePreventionModes() bool {
	if o != nil && !common.IsNil(o.AllowedSelfTradePreventionModes) {
		return true
	}

	return false
}

// SetAllowedSelfTradePreventionModes gets a reference to the given []string and assigns it to the AllowedSelfTradePreventionModes field.
func (o *ExchangeInfoResponseSymbolsInner) SetAllowedSelfTradePreventionModes(v []string) {
	o.AllowedSelfTradePreventionModes = v
}

func (o ExchangeInfoResponseSymbolsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeInfoResponseSymbolsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.BaseAsset) {
		toSerialize["baseAsset"] = o.BaseAsset
	}
	if !common.IsNil(o.BaseAssetPrecision) {
		toSerialize["baseAssetPrecision"] = o.BaseAssetPrecision
	}
	if !common.IsNil(o.QuoteAsset) {
		toSerialize["quoteAsset"] = o.QuoteAsset
	}
	if !common.IsNil(o.QuotePrecision) {
		toSerialize["quotePrecision"] = o.QuotePrecision
	}
	if !common.IsNil(o.QuoteAssetPrecision) {
		toSerialize["quoteAssetPrecision"] = o.QuoteAssetPrecision
	}
	if !common.IsNil(o.BaseCommissionPrecision) {
		toSerialize["baseCommissionPrecision"] = o.BaseCommissionPrecision
	}
	if !common.IsNil(o.QuoteCommissionPrecision) {
		toSerialize["quoteCommissionPrecision"] = o.QuoteCommissionPrecision
	}
	if !common.IsNil(o.OrderTypes) {
		toSerialize["orderTypes"] = o.OrderTypes
	}
	if !common.IsNil(o.IcebergAllowed) {
		toSerialize["icebergAllowed"] = o.IcebergAllowed
	}
	if !common.IsNil(o.OcoAllowed) {
		toSerialize["ocoAllowed"] = o.OcoAllowed
	}
	if !common.IsNil(o.OtoAllowed) {
		toSerialize["otoAllowed"] = o.OtoAllowed
	}
	if !common.IsNil(o.OpoAllowed) {
		toSerialize["opoAllowed"] = o.OpoAllowed
	}
	if !common.IsNil(o.QuoteOrderQtyMarketAllowed) {
		toSerialize["quoteOrderQtyMarketAllowed"] = o.QuoteOrderQtyMarketAllowed
	}
	if !common.IsNil(o.AllowTrailingStop) {
		toSerialize["allowTrailingStop"] = o.AllowTrailingStop
	}
	if !common.IsNil(o.CancelReplaceAllowed) {
		toSerialize["cancelReplaceAllowed"] = o.CancelReplaceAllowed
	}
	if !common.IsNil(o.AmendAllowed) {
		toSerialize["amendAllowed"] = o.AmendAllowed
	}
	if !common.IsNil(o.PegInstructionsAllowed) {
		toSerialize["pegInstructionsAllowed"] = o.PegInstructionsAllowed
	}
	if !common.IsNil(o.IsSpotTradingAllowed) {
		toSerialize["isSpotTradingAllowed"] = o.IsSpotTradingAllowed
	}
	if !common.IsNil(o.IsMarginTradingAllowed) {
		toSerialize["isMarginTradingAllowed"] = o.IsMarginTradingAllowed
	}
	if !common.IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !common.IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !common.IsNil(o.PermissionSets) {
		toSerialize["permissionSets"] = o.PermissionSets
	}
	if !common.IsNil(o.DefaultSelfTradePreventionMode) {
		toSerialize["defaultSelfTradePreventionMode"] = o.DefaultSelfTradePreventionMode
	}
	if !common.IsNil(o.AllowedSelfTradePreventionModes) {
		toSerialize["allowedSelfTradePreventionModes"] = o.AllowedSelfTradePreventionModes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExchangeInfoResponseSymbolsInner) UnmarshalJSON(data []byte) (err error) {
	varExchangeInfoResponseSymbolsInner := _ExchangeInfoResponseSymbolsInner{}

	err = json.Unmarshal(data, &varExchangeInfoResponseSymbolsInner)

	if err != nil {
		return err
	}

	*o = ExchangeInfoResponseSymbolsInner(varExchangeInfoResponseSymbolsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "status")
		delete(additionalProperties, "baseAsset")
		delete(additionalProperties, "baseAssetPrecision")
		delete(additionalProperties, "quoteAsset")
		delete(additionalProperties, "quotePrecision")
		delete(additionalProperties, "quoteAssetPrecision")
		delete(additionalProperties, "baseCommissionPrecision")
		delete(additionalProperties, "quoteCommissionPrecision")
		delete(additionalProperties, "orderTypes")
		delete(additionalProperties, "icebergAllowed")
		delete(additionalProperties, "ocoAllowed")
		delete(additionalProperties, "otoAllowed")
		delete(additionalProperties, "opoAllowed")
		delete(additionalProperties, "quoteOrderQtyMarketAllowed")
		delete(additionalProperties, "allowTrailingStop")
		delete(additionalProperties, "cancelReplaceAllowed")
		delete(additionalProperties, "amendAllowed")
		delete(additionalProperties, "pegInstructionsAllowed")
		delete(additionalProperties, "isSpotTradingAllowed")
		delete(additionalProperties, "isMarginTradingAllowed")
		delete(additionalProperties, "filters")
		delete(additionalProperties, "permissions")
		delete(additionalProperties, "permissionSets")
		delete(additionalProperties, "defaultSelfTradePreventionMode")
		delete(additionalProperties, "allowedSelfTradePreventionModes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeInfoResponseSymbolsInner struct {
	value *ExchangeInfoResponseSymbolsInner
	isSet bool
}

func (v NullableExchangeInfoResponseSymbolsInner) Get() *ExchangeInfoResponseSymbolsInner {
	return v.value
}

func (v *NullableExchangeInfoResponseSymbolsInner) Set(val *ExchangeInfoResponseSymbolsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInfoResponseSymbolsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInfoResponseSymbolsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInfoResponseSymbolsInner(val *ExchangeInfoResponseSymbolsInner) *NullableExchangeInfoResponseSymbolsInner {
	return &NullableExchangeInfoResponseSymbolsInner{value: val, isSet: true}
}

func (v NullableExchangeInfoResponseSymbolsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInfoResponseSymbolsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
