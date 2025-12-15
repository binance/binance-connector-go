/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ExchangeInfoResponseResultSymbolsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeInfoResponseResultSymbolsInner{}

// ExchangeInfoResponseResultSymbolsInner struct for ExchangeInfoResponseResultSymbolsInner
type ExchangeInfoResponseResultSymbolsInner struct {
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

type _ExchangeInfoResponseResultSymbolsInner ExchangeInfoResponseResultSymbolsInner

// NewExchangeInfoResponseResultSymbolsInner instantiates a new ExchangeInfoResponseResultSymbolsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeInfoResponseResultSymbolsInner() *ExchangeInfoResponseResultSymbolsInner {
	this := ExchangeInfoResponseResultSymbolsInner{}
	return &this
}

// NewExchangeInfoResponseResultSymbolsInnerWithDefaults instantiates a new ExchangeInfoResponseResultSymbolsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeInfoResponseResultSymbolsInnerWithDefaults() *ExchangeInfoResponseResultSymbolsInner {
	this := ExchangeInfoResponseResultSymbolsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetStatus(v string) {
	o.Status = &v
}

// GetBaseAsset returns the BaseAsset field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetBaseAsset() string {
	if o == nil || common.IsNil(o.BaseAsset) {
		var ret string
		return ret
	}
	return *o.BaseAsset
}

// GetBaseAssetOk returns a tuple with the BaseAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetBaseAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.BaseAsset) {
		return nil, false
	}
	return o.BaseAsset, true
}

// HasBaseAsset returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasBaseAsset() bool {
	if o != nil && !common.IsNil(o.BaseAsset) {
		return true
	}

	return false
}

// SetBaseAsset gets a reference to the given string and assigns it to the BaseAsset field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetBaseAsset(v string) {
	o.BaseAsset = &v
}

// GetBaseAssetPrecision returns the BaseAssetPrecision field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetBaseAssetPrecision() int64 {
	if o == nil || common.IsNil(o.BaseAssetPrecision) {
		var ret int64
		return ret
	}
	return *o.BaseAssetPrecision
}

// GetBaseAssetPrecisionOk returns a tuple with the BaseAssetPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetBaseAssetPrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BaseAssetPrecision) {
		return nil, false
	}
	return o.BaseAssetPrecision, true
}

// HasBaseAssetPrecision returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasBaseAssetPrecision() bool {
	if o != nil && !common.IsNil(o.BaseAssetPrecision) {
		return true
	}

	return false
}

// SetBaseAssetPrecision gets a reference to the given int64 and assigns it to the BaseAssetPrecision field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetBaseAssetPrecision(v int64) {
	o.BaseAssetPrecision = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetQuoteAsset() string {
	if o == nil || common.IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetQuoteAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasQuoteAsset() bool {
	if o != nil && !common.IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

// GetQuotePrecision returns the QuotePrecision field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetQuotePrecision() int64 {
	if o == nil || common.IsNil(o.QuotePrecision) {
		var ret int64
		return ret
	}
	return *o.QuotePrecision
}

// GetQuotePrecisionOk returns a tuple with the QuotePrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetQuotePrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QuotePrecision) {
		return nil, false
	}
	return o.QuotePrecision, true
}

// HasQuotePrecision returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasQuotePrecision() bool {
	if o != nil && !common.IsNil(o.QuotePrecision) {
		return true
	}

	return false
}

// SetQuotePrecision gets a reference to the given int64 and assigns it to the QuotePrecision field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetQuotePrecision(v int64) {
	o.QuotePrecision = &v
}

// GetQuoteAssetPrecision returns the QuoteAssetPrecision field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetQuoteAssetPrecision() int64 {
	if o == nil || common.IsNil(o.QuoteAssetPrecision) {
		var ret int64
		return ret
	}
	return *o.QuoteAssetPrecision
}

// GetQuoteAssetPrecisionOk returns a tuple with the QuoteAssetPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetQuoteAssetPrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QuoteAssetPrecision) {
		return nil, false
	}
	return o.QuoteAssetPrecision, true
}

// HasQuoteAssetPrecision returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasQuoteAssetPrecision() bool {
	if o != nil && !common.IsNil(o.QuoteAssetPrecision) {
		return true
	}

	return false
}

// SetQuoteAssetPrecision gets a reference to the given int64 and assigns it to the QuoteAssetPrecision field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetQuoteAssetPrecision(v int64) {
	o.QuoteAssetPrecision = &v
}

// GetBaseCommissionPrecision returns the BaseCommissionPrecision field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetBaseCommissionPrecision() int64 {
	if o == nil || common.IsNil(o.BaseCommissionPrecision) {
		var ret int64
		return ret
	}
	return *o.BaseCommissionPrecision
}

// GetBaseCommissionPrecisionOk returns a tuple with the BaseCommissionPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetBaseCommissionPrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BaseCommissionPrecision) {
		return nil, false
	}
	return o.BaseCommissionPrecision, true
}

// HasBaseCommissionPrecision returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasBaseCommissionPrecision() bool {
	if o != nil && !common.IsNil(o.BaseCommissionPrecision) {
		return true
	}

	return false
}

// SetBaseCommissionPrecision gets a reference to the given int64 and assigns it to the BaseCommissionPrecision field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetBaseCommissionPrecision(v int64) {
	o.BaseCommissionPrecision = &v
}

// GetQuoteCommissionPrecision returns the QuoteCommissionPrecision field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetQuoteCommissionPrecision() int64 {
	if o == nil || common.IsNil(o.QuoteCommissionPrecision) {
		var ret int64
		return ret
	}
	return *o.QuoteCommissionPrecision
}

// GetQuoteCommissionPrecisionOk returns a tuple with the QuoteCommissionPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetQuoteCommissionPrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QuoteCommissionPrecision) {
		return nil, false
	}
	return o.QuoteCommissionPrecision, true
}

// HasQuoteCommissionPrecision returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasQuoteCommissionPrecision() bool {
	if o != nil && !common.IsNil(o.QuoteCommissionPrecision) {
		return true
	}

	return false
}

// SetQuoteCommissionPrecision gets a reference to the given int64 and assigns it to the QuoteCommissionPrecision field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetQuoteCommissionPrecision(v int64) {
	o.QuoteCommissionPrecision = &v
}

// GetOrderTypes returns the OrderTypes field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetOrderTypes() []string {
	if o == nil || common.IsNil(o.OrderTypes) {
		var ret []string
		return ret
	}
	return o.OrderTypes
}

// GetOrderTypesOk returns a tuple with the OrderTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetOrderTypesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.OrderTypes) {
		return nil, false
	}
	return o.OrderTypes, true
}

// HasOrderTypes returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasOrderTypes() bool {
	if o != nil && !common.IsNil(o.OrderTypes) {
		return true
	}

	return false
}

// SetOrderTypes gets a reference to the given []string and assigns it to the OrderTypes field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetOrderTypes(v []string) {
	o.OrderTypes = v
}

// GetIcebergAllowed returns the IcebergAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetIcebergAllowed() bool {
	if o == nil || common.IsNil(o.IcebergAllowed) {
		var ret bool
		return ret
	}
	return *o.IcebergAllowed
}

// GetIcebergAllowedOk returns a tuple with the IcebergAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetIcebergAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IcebergAllowed) {
		return nil, false
	}
	return o.IcebergAllowed, true
}

// HasIcebergAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasIcebergAllowed() bool {
	if o != nil && !common.IsNil(o.IcebergAllowed) {
		return true
	}

	return false
}

// SetIcebergAllowed gets a reference to the given bool and assigns it to the IcebergAllowed field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetIcebergAllowed(v bool) {
	o.IcebergAllowed = &v
}

// GetOcoAllowed returns the OcoAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetOcoAllowed() bool {
	if o == nil || common.IsNil(o.OcoAllowed) {
		var ret bool
		return ret
	}
	return *o.OcoAllowed
}

// GetOcoAllowedOk returns a tuple with the OcoAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetOcoAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.OcoAllowed) {
		return nil, false
	}
	return o.OcoAllowed, true
}

// HasOcoAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasOcoAllowed() bool {
	if o != nil && !common.IsNil(o.OcoAllowed) {
		return true
	}

	return false
}

// SetOcoAllowed gets a reference to the given bool and assigns it to the OcoAllowed field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetOcoAllowed(v bool) {
	o.OcoAllowed = &v
}

// GetOtoAllowed returns the OtoAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetOtoAllowed() bool {
	if o == nil || common.IsNil(o.OtoAllowed) {
		var ret bool
		return ret
	}
	return *o.OtoAllowed
}

// GetOtoAllowedOk returns a tuple with the OtoAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetOtoAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.OtoAllowed) {
		return nil, false
	}
	return o.OtoAllowed, true
}

// HasOtoAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasOtoAllowed() bool {
	if o != nil && !common.IsNil(o.OtoAllowed) {
		return true
	}

	return false
}

// SetOtoAllowed gets a reference to the given bool and assigns it to the OtoAllowed field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetOtoAllowed(v bool) {
	o.OtoAllowed = &v
}

// GetOpoAllowed returns the OpoAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetOpoAllowed() bool {
	if o == nil || common.IsNil(o.OpoAllowed) {
		var ret bool
		return ret
	}
	return *o.OpoAllowed
}

// GetOpoAllowedOk returns a tuple with the OpoAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetOpoAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.OpoAllowed) {
		return nil, false
	}
	return o.OpoAllowed, true
}

// HasOpoAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasOpoAllowed() bool {
	if o != nil && !common.IsNil(o.OpoAllowed) {
		return true
	}

	return false
}

// SetOpoAllowed gets a reference to the given bool and assigns it to the OpoAllowed field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetOpoAllowed(v bool) {
	o.OpoAllowed = &v
}

// GetQuoteOrderQtyMarketAllowed returns the QuoteOrderQtyMarketAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetQuoteOrderQtyMarketAllowed() bool {
	if o == nil || common.IsNil(o.QuoteOrderQtyMarketAllowed) {
		var ret bool
		return ret
	}
	return *o.QuoteOrderQtyMarketAllowed
}

// GetQuoteOrderQtyMarketAllowedOk returns a tuple with the QuoteOrderQtyMarketAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetQuoteOrderQtyMarketAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.QuoteOrderQtyMarketAllowed) {
		return nil, false
	}
	return o.QuoteOrderQtyMarketAllowed, true
}

// HasQuoteOrderQtyMarketAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasQuoteOrderQtyMarketAllowed() bool {
	if o != nil && !common.IsNil(o.QuoteOrderQtyMarketAllowed) {
		return true
	}

	return false
}

// SetQuoteOrderQtyMarketAllowed gets a reference to the given bool and assigns it to the QuoteOrderQtyMarketAllowed field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetQuoteOrderQtyMarketAllowed(v bool) {
	o.QuoteOrderQtyMarketAllowed = &v
}

// GetAllowTrailingStop returns the AllowTrailingStop field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetAllowTrailingStop() bool {
	if o == nil || common.IsNil(o.AllowTrailingStop) {
		var ret bool
		return ret
	}
	return *o.AllowTrailingStop
}

// GetAllowTrailingStopOk returns a tuple with the AllowTrailingStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetAllowTrailingStopOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AllowTrailingStop) {
		return nil, false
	}
	return o.AllowTrailingStop, true
}

// HasAllowTrailingStop returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasAllowTrailingStop() bool {
	if o != nil && !common.IsNil(o.AllowTrailingStop) {
		return true
	}

	return false
}

// SetAllowTrailingStop gets a reference to the given bool and assigns it to the AllowTrailingStop field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetAllowTrailingStop(v bool) {
	o.AllowTrailingStop = &v
}

// GetCancelReplaceAllowed returns the CancelReplaceAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetCancelReplaceAllowed() bool {
	if o == nil || common.IsNil(o.CancelReplaceAllowed) {
		var ret bool
		return ret
	}
	return *o.CancelReplaceAllowed
}

// GetCancelReplaceAllowedOk returns a tuple with the CancelReplaceAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetCancelReplaceAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CancelReplaceAllowed) {
		return nil, false
	}
	return o.CancelReplaceAllowed, true
}

// HasCancelReplaceAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasCancelReplaceAllowed() bool {
	if o != nil && !common.IsNil(o.CancelReplaceAllowed) {
		return true
	}

	return false
}

// SetCancelReplaceAllowed gets a reference to the given bool and assigns it to the CancelReplaceAllowed field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetCancelReplaceAllowed(v bool) {
	o.CancelReplaceAllowed = &v
}

// GetAmendAllowed returns the AmendAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetAmendAllowed() bool {
	if o == nil || common.IsNil(o.AmendAllowed) {
		var ret bool
		return ret
	}
	return *o.AmendAllowed
}

// GetAmendAllowedOk returns a tuple with the AmendAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetAmendAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AmendAllowed) {
		return nil, false
	}
	return o.AmendAllowed, true
}

// HasAmendAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasAmendAllowed() bool {
	if o != nil && !common.IsNil(o.AmendAllowed) {
		return true
	}

	return false
}

// SetAmendAllowed gets a reference to the given bool and assigns it to the AmendAllowed field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetAmendAllowed(v bool) {
	o.AmendAllowed = &v
}

// GetPegInstructionsAllowed returns the PegInstructionsAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetPegInstructionsAllowed() bool {
	if o == nil || common.IsNil(o.PegInstructionsAllowed) {
		var ret bool
		return ret
	}
	return *o.PegInstructionsAllowed
}

// GetPegInstructionsAllowedOk returns a tuple with the PegInstructionsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetPegInstructionsAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PegInstructionsAllowed) {
		return nil, false
	}
	return o.PegInstructionsAllowed, true
}

// HasPegInstructionsAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasPegInstructionsAllowed() bool {
	if o != nil && !common.IsNil(o.PegInstructionsAllowed) {
		return true
	}

	return false
}

// SetPegInstructionsAllowed gets a reference to the given bool and assigns it to the PegInstructionsAllowed field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetPegInstructionsAllowed(v bool) {
	o.PegInstructionsAllowed = &v
}

// GetIsSpotTradingAllowed returns the IsSpotTradingAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetIsSpotTradingAllowed() bool {
	if o == nil || common.IsNil(o.IsSpotTradingAllowed) {
		var ret bool
		return ret
	}
	return *o.IsSpotTradingAllowed
}

// GetIsSpotTradingAllowedOk returns a tuple with the IsSpotTradingAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetIsSpotTradingAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsSpotTradingAllowed) {
		return nil, false
	}
	return o.IsSpotTradingAllowed, true
}

// HasIsSpotTradingAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasIsSpotTradingAllowed() bool {
	if o != nil && !common.IsNil(o.IsSpotTradingAllowed) {
		return true
	}

	return false
}

// SetIsSpotTradingAllowed gets a reference to the given bool and assigns it to the IsSpotTradingAllowed field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetIsSpotTradingAllowed(v bool) {
	o.IsSpotTradingAllowed = &v
}

// GetIsMarginTradingAllowed returns the IsMarginTradingAllowed field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetIsMarginTradingAllowed() bool {
	if o == nil || common.IsNil(o.IsMarginTradingAllowed) {
		var ret bool
		return ret
	}
	return *o.IsMarginTradingAllowed
}

// GetIsMarginTradingAllowedOk returns a tuple with the IsMarginTradingAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetIsMarginTradingAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsMarginTradingAllowed) {
		return nil, false
	}
	return o.IsMarginTradingAllowed, true
}

// HasIsMarginTradingAllowed returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasIsMarginTradingAllowed() bool {
	if o != nil && !common.IsNil(o.IsMarginTradingAllowed) {
		return true
	}

	return false
}

// SetIsMarginTradingAllowed gets a reference to the given bool and assigns it to the IsMarginTradingAllowed field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetIsMarginTradingAllowed(v bool) {
	o.IsMarginTradingAllowed = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetFilters() []SymbolFilters {
	if o == nil || common.IsNil(o.Filters) {
		var ret []SymbolFilters
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetFiltersOk() ([]SymbolFilters, bool) {
	if o == nil || common.IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasFilters() bool {
	if o != nil && !common.IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []SymbolFilters and assigns it to the Filters field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetFilters(v []SymbolFilters) {
	o.Filters = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetPermissions() []string {
	if o == nil || common.IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetPermissionsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasPermissions() bool {
	if o != nil && !common.IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetPermissions(v []string) {
	o.Permissions = v
}

// GetPermissionSets returns the PermissionSets field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetPermissionSets() [][]string {
	if o == nil || common.IsNil(o.PermissionSets) {
		var ret [][]string
		return ret
	}
	return o.PermissionSets
}

// GetPermissionSetsOk returns a tuple with the PermissionSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetPermissionSetsOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.PermissionSets) {
		return nil, false
	}
	return o.PermissionSets, true
}

// HasPermissionSets returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasPermissionSets() bool {
	if o != nil && !common.IsNil(o.PermissionSets) {
		return true
	}

	return false
}

// SetPermissionSets gets a reference to the given [][]string and assigns it to the PermissionSets field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetPermissionSets(v [][]string) {
	o.PermissionSets = v
}

// GetDefaultSelfTradePreventionMode returns the DefaultSelfTradePreventionMode field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetDefaultSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.DefaultSelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.DefaultSelfTradePreventionMode
}

// GetDefaultSelfTradePreventionModeOk returns a tuple with the DefaultSelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetDefaultSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.DefaultSelfTradePreventionMode) {
		return nil, false
	}
	return o.DefaultSelfTradePreventionMode, true
}

// HasDefaultSelfTradePreventionMode returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasDefaultSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.DefaultSelfTradePreventionMode) {
		return true
	}

	return false
}

// SetDefaultSelfTradePreventionMode gets a reference to the given string and assigns it to the DefaultSelfTradePreventionMode field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetDefaultSelfTradePreventionMode(v string) {
	o.DefaultSelfTradePreventionMode = &v
}

// GetAllowedSelfTradePreventionModes returns the AllowedSelfTradePreventionModes field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSymbolsInner) GetAllowedSelfTradePreventionModes() []string {
	if o == nil || common.IsNil(o.AllowedSelfTradePreventionModes) {
		var ret []string
		return ret
	}
	return o.AllowedSelfTradePreventionModes
}

// GetAllowedSelfTradePreventionModesOk returns a tuple with the AllowedSelfTradePreventionModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) GetAllowedSelfTradePreventionModesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AllowedSelfTradePreventionModes) {
		return nil, false
	}
	return o.AllowedSelfTradePreventionModes, true
}

// HasAllowedSelfTradePreventionModes returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSymbolsInner) HasAllowedSelfTradePreventionModes() bool {
	if o != nil && !common.IsNil(o.AllowedSelfTradePreventionModes) {
		return true
	}

	return false
}

// SetAllowedSelfTradePreventionModes gets a reference to the given []string and assigns it to the AllowedSelfTradePreventionModes field.
func (o *ExchangeInfoResponseResultSymbolsInner) SetAllowedSelfTradePreventionModes(v []string) {
	o.AllowedSelfTradePreventionModes = v
}

func (o ExchangeInfoResponseResultSymbolsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeInfoResponseResultSymbolsInner) ToMap() (map[string]interface{}, error) {
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

func (o *ExchangeInfoResponseResultSymbolsInner) UnmarshalJSON(data []byte) (err error) {
	varExchangeInfoResponseResultSymbolsInner := _ExchangeInfoResponseResultSymbolsInner{}

	err = json.Unmarshal(data, &varExchangeInfoResponseResultSymbolsInner)

	if err != nil {
		return err
	}

	*o = ExchangeInfoResponseResultSymbolsInner(varExchangeInfoResponseResultSymbolsInner)

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

type NullableExchangeInfoResponseResultSymbolsInner struct {
	value *ExchangeInfoResponseResultSymbolsInner
	isSet bool
}

func (v NullableExchangeInfoResponseResultSymbolsInner) Get() *ExchangeInfoResponseResultSymbolsInner {
	return v.value
}

func (v *NullableExchangeInfoResponseResultSymbolsInner) Set(val *ExchangeInfoResponseResultSymbolsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInfoResponseResultSymbolsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInfoResponseResultSymbolsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInfoResponseResultSymbolsInner(val *ExchangeInfoResponseResultSymbolsInner) *NullableExchangeInfoResponseResultSymbolsInner {
	return &NullableExchangeInfoResponseResultSymbolsInner{value: val, isSet: true}
}

func (v NullableExchangeInfoResponseResultSymbolsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInfoResponseResultSymbolsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
