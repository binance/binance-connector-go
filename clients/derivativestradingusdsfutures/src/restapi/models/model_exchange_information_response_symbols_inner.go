/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ExchangeInformationResponseSymbolsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeInformationResponseSymbolsInner{}

// ExchangeInformationResponseSymbolsInner struct for ExchangeInformationResponseSymbolsInner
type ExchangeInformationResponseSymbolsInner struct {
	Symbol                *string                                               `json:"symbol,omitempty"`
	Pair                  *string                                               `json:"pair,omitempty"`
	ContractType          *string                                               `json:"contractType,omitempty"`
	DeliveryDate          *int64                                                `json:"deliveryDate,omitempty"`
	OnboardDate           *int64                                                `json:"onboardDate,omitempty"`
	Status                *string                                               `json:"status,omitempty"`
	MaintMarginPercent    *string                                               `json:"maintMarginPercent,omitempty"`
	RequiredMarginPercent *string                                               `json:"requiredMarginPercent,omitempty"`
	BaseAsset             *string                                               `json:"baseAsset,omitempty"`
	QuoteAsset            *string                                               `json:"quoteAsset,omitempty"`
	MarginAsset           *string                                               `json:"marginAsset,omitempty"`
	PricePrecision        *int64                                                `json:"pricePrecision,omitempty"`
	QuantityPrecision     *int64                                                `json:"quantityPrecision,omitempty"`
	BaseAssetPrecision    *int64                                                `json:"baseAssetPrecision,omitempty"`
	QuotePrecision        *int64                                                `json:"quotePrecision,omitempty"`
	UnderlyingType        *string                                               `json:"underlyingType,omitempty"`
	UnderlyingSubType     []string                                              `json:"underlyingSubType,omitempty"`
	SettlePlan            *int64                                                `json:"settlePlan,omitempty"`
	TriggerProtect        *string                                               `json:"triggerProtect,omitempty"`
	Filters               []ExchangeInformationResponseSymbolsInnerFiltersInner `json:"filters,omitempty"`
	OrderType             []string                                              `json:"OrderType,omitempty"`
	TimeInForce           []string                                              `json:"timeInForce,omitempty"`
	LiquidationFee        *string                                               `json:"liquidationFee,omitempty"`
	MarketTakeBound       *string                                               `json:"marketTakeBound,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _ExchangeInformationResponseSymbolsInner ExchangeInformationResponseSymbolsInner

// NewExchangeInformationResponseSymbolsInner instantiates a new ExchangeInformationResponseSymbolsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeInformationResponseSymbolsInner() *ExchangeInformationResponseSymbolsInner {
	this := ExchangeInformationResponseSymbolsInner{}
	return &this
}

// NewExchangeInformationResponseSymbolsInnerWithDefaults instantiates a new ExchangeInformationResponseSymbolsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeInformationResponseSymbolsInnerWithDefaults() *ExchangeInformationResponseSymbolsInner {
	this := ExchangeInformationResponseSymbolsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ExchangeInformationResponseSymbolsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetPair() string {
	if o == nil || common.IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetPairOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasPair() bool {
	if o != nil && !common.IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *ExchangeInformationResponseSymbolsInner) SetPair(v string) {
	o.Pair = &v
}

// GetContractType returns the ContractType field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetContractType() string {
	if o == nil || common.IsNil(o.ContractType) {
		var ret string
		return ret
	}
	return *o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetContractTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContractType) {
		return nil, false
	}
	return o.ContractType, true
}

// HasContractType returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasContractType() bool {
	if o != nil && !common.IsNil(o.ContractType) {
		return true
	}

	return false
}

// SetContractType gets a reference to the given string and assigns it to the ContractType field.
func (o *ExchangeInformationResponseSymbolsInner) SetContractType(v string) {
	o.ContractType = &v
}

// GetDeliveryDate returns the DeliveryDate field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetDeliveryDate() int64 {
	if o == nil || common.IsNil(o.DeliveryDate) {
		var ret int64
		return ret
	}
	return *o.DeliveryDate
}

// GetDeliveryDateOk returns a tuple with the DeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetDeliveryDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.DeliveryDate) {
		return nil, false
	}
	return o.DeliveryDate, true
}

// HasDeliveryDate returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasDeliveryDate() bool {
	if o != nil && !common.IsNil(o.DeliveryDate) {
		return true
	}

	return false
}

// SetDeliveryDate gets a reference to the given int64 and assigns it to the DeliveryDate field.
func (o *ExchangeInformationResponseSymbolsInner) SetDeliveryDate(v int64) {
	o.DeliveryDate = &v
}

// GetOnboardDate returns the OnboardDate field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetOnboardDate() int64 {
	if o == nil || common.IsNil(o.OnboardDate) {
		var ret int64
		return ret
	}
	return *o.OnboardDate
}

// GetOnboardDateOk returns a tuple with the OnboardDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetOnboardDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OnboardDate) {
		return nil, false
	}
	return o.OnboardDate, true
}

// HasOnboardDate returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasOnboardDate() bool {
	if o != nil && !common.IsNil(o.OnboardDate) {
		return true
	}

	return false
}

// SetOnboardDate gets a reference to the given int64 and assigns it to the OnboardDate field.
func (o *ExchangeInformationResponseSymbolsInner) SetOnboardDate(v int64) {
	o.OnboardDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ExchangeInformationResponseSymbolsInner) SetStatus(v string) {
	o.Status = &v
}

// GetMaintMarginPercent returns the MaintMarginPercent field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetMaintMarginPercent() string {
	if o == nil || common.IsNil(o.MaintMarginPercent) {
		var ret string
		return ret
	}
	return *o.MaintMarginPercent
}

// GetMaintMarginPercentOk returns a tuple with the MaintMarginPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetMaintMarginPercentOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaintMarginPercent) {
		return nil, false
	}
	return o.MaintMarginPercent, true
}

// HasMaintMarginPercent returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasMaintMarginPercent() bool {
	if o != nil && !common.IsNil(o.MaintMarginPercent) {
		return true
	}

	return false
}

// SetMaintMarginPercent gets a reference to the given string and assigns it to the MaintMarginPercent field.
func (o *ExchangeInformationResponseSymbolsInner) SetMaintMarginPercent(v string) {
	o.MaintMarginPercent = &v
}

// GetRequiredMarginPercent returns the RequiredMarginPercent field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetRequiredMarginPercent() string {
	if o == nil || common.IsNil(o.RequiredMarginPercent) {
		var ret string
		return ret
	}
	return *o.RequiredMarginPercent
}

// GetRequiredMarginPercentOk returns a tuple with the RequiredMarginPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetRequiredMarginPercentOk() (*string, bool) {
	if o == nil || common.IsNil(o.RequiredMarginPercent) {
		return nil, false
	}
	return o.RequiredMarginPercent, true
}

// HasRequiredMarginPercent returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasRequiredMarginPercent() bool {
	if o != nil && !common.IsNil(o.RequiredMarginPercent) {
		return true
	}

	return false
}

// SetRequiredMarginPercent gets a reference to the given string and assigns it to the RequiredMarginPercent field.
func (o *ExchangeInformationResponseSymbolsInner) SetRequiredMarginPercent(v string) {
	o.RequiredMarginPercent = &v
}

// GetBaseAsset returns the BaseAsset field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetBaseAsset() string {
	if o == nil || common.IsNil(o.BaseAsset) {
		var ret string
		return ret
	}
	return *o.BaseAsset
}

// GetBaseAssetOk returns a tuple with the BaseAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetBaseAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.BaseAsset) {
		return nil, false
	}
	return o.BaseAsset, true
}

// HasBaseAsset returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasBaseAsset() bool {
	if o != nil && !common.IsNil(o.BaseAsset) {
		return true
	}

	return false
}

// SetBaseAsset gets a reference to the given string and assigns it to the BaseAsset field.
func (o *ExchangeInformationResponseSymbolsInner) SetBaseAsset(v string) {
	o.BaseAsset = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetQuoteAsset() string {
	if o == nil || common.IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetQuoteAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasQuoteAsset() bool {
	if o != nil && !common.IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *ExchangeInformationResponseSymbolsInner) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

// GetMarginAsset returns the MarginAsset field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetMarginAsset() string {
	if o == nil || common.IsNil(o.MarginAsset) {
		var ret string
		return ret
	}
	return *o.MarginAsset
}

// GetMarginAssetOk returns a tuple with the MarginAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetMarginAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginAsset) {
		return nil, false
	}
	return o.MarginAsset, true
}

// HasMarginAsset returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasMarginAsset() bool {
	if o != nil && !common.IsNil(o.MarginAsset) {
		return true
	}

	return false
}

// SetMarginAsset gets a reference to the given string and assigns it to the MarginAsset field.
func (o *ExchangeInformationResponseSymbolsInner) SetMarginAsset(v string) {
	o.MarginAsset = &v
}

// GetPricePrecision returns the PricePrecision field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetPricePrecision() int64 {
	if o == nil || common.IsNil(o.PricePrecision) {
		var ret int64
		return ret
	}
	return *o.PricePrecision
}

// GetPricePrecisionOk returns a tuple with the PricePrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetPricePrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PricePrecision) {
		return nil, false
	}
	return o.PricePrecision, true
}

// HasPricePrecision returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasPricePrecision() bool {
	if o != nil && !common.IsNil(o.PricePrecision) {
		return true
	}

	return false
}

// SetPricePrecision gets a reference to the given int64 and assigns it to the PricePrecision field.
func (o *ExchangeInformationResponseSymbolsInner) SetPricePrecision(v int64) {
	o.PricePrecision = &v
}

// GetQuantityPrecision returns the QuantityPrecision field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetQuantityPrecision() int64 {
	if o == nil || common.IsNil(o.QuantityPrecision) {
		var ret int64
		return ret
	}
	return *o.QuantityPrecision
}

// GetQuantityPrecisionOk returns a tuple with the QuantityPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetQuantityPrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QuantityPrecision) {
		return nil, false
	}
	return o.QuantityPrecision, true
}

// HasQuantityPrecision returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasQuantityPrecision() bool {
	if o != nil && !common.IsNil(o.QuantityPrecision) {
		return true
	}

	return false
}

// SetQuantityPrecision gets a reference to the given int64 and assigns it to the QuantityPrecision field.
func (o *ExchangeInformationResponseSymbolsInner) SetQuantityPrecision(v int64) {
	o.QuantityPrecision = &v
}

// GetBaseAssetPrecision returns the BaseAssetPrecision field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetBaseAssetPrecision() int64 {
	if o == nil || common.IsNil(o.BaseAssetPrecision) {
		var ret int64
		return ret
	}
	return *o.BaseAssetPrecision
}

// GetBaseAssetPrecisionOk returns a tuple with the BaseAssetPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetBaseAssetPrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BaseAssetPrecision) {
		return nil, false
	}
	return o.BaseAssetPrecision, true
}

// HasBaseAssetPrecision returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasBaseAssetPrecision() bool {
	if o != nil && !common.IsNil(o.BaseAssetPrecision) {
		return true
	}

	return false
}

// SetBaseAssetPrecision gets a reference to the given int64 and assigns it to the BaseAssetPrecision field.
func (o *ExchangeInformationResponseSymbolsInner) SetBaseAssetPrecision(v int64) {
	o.BaseAssetPrecision = &v
}

// GetQuotePrecision returns the QuotePrecision field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetQuotePrecision() int64 {
	if o == nil || common.IsNil(o.QuotePrecision) {
		var ret int64
		return ret
	}
	return *o.QuotePrecision
}

// GetQuotePrecisionOk returns a tuple with the QuotePrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetQuotePrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QuotePrecision) {
		return nil, false
	}
	return o.QuotePrecision, true
}

// HasQuotePrecision returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasQuotePrecision() bool {
	if o != nil && !common.IsNil(o.QuotePrecision) {
		return true
	}

	return false
}

// SetQuotePrecision gets a reference to the given int64 and assigns it to the QuotePrecision field.
func (o *ExchangeInformationResponseSymbolsInner) SetQuotePrecision(v int64) {
	o.QuotePrecision = &v
}

// GetUnderlyingType returns the UnderlyingType field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetUnderlyingType() string {
	if o == nil || common.IsNil(o.UnderlyingType) {
		var ret string
		return ret
	}
	return *o.UnderlyingType
}

// GetUnderlyingTypeOk returns a tuple with the UnderlyingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetUnderlyingTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnderlyingType) {
		return nil, false
	}
	return o.UnderlyingType, true
}

// HasUnderlyingType returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasUnderlyingType() bool {
	if o != nil && !common.IsNil(o.UnderlyingType) {
		return true
	}

	return false
}

// SetUnderlyingType gets a reference to the given string and assigns it to the UnderlyingType field.
func (o *ExchangeInformationResponseSymbolsInner) SetUnderlyingType(v string) {
	o.UnderlyingType = &v
}

// GetUnderlyingSubType returns the UnderlyingSubType field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetUnderlyingSubType() []string {
	if o == nil || common.IsNil(o.UnderlyingSubType) {
		var ret []string
		return ret
	}
	return o.UnderlyingSubType
}

// GetUnderlyingSubTypeOk returns a tuple with the UnderlyingSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetUnderlyingSubTypeOk() ([]string, bool) {
	if o == nil || common.IsNil(o.UnderlyingSubType) {
		return nil, false
	}
	return o.UnderlyingSubType, true
}

// HasUnderlyingSubType returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasUnderlyingSubType() bool {
	if o != nil && !common.IsNil(o.UnderlyingSubType) {
		return true
	}

	return false
}

// SetUnderlyingSubType gets a reference to the given []string and assigns it to the UnderlyingSubType field.
func (o *ExchangeInformationResponseSymbolsInner) SetUnderlyingSubType(v []string) {
	o.UnderlyingSubType = v
}

// GetSettlePlan returns the SettlePlan field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetSettlePlan() int64 {
	if o == nil || common.IsNil(o.SettlePlan) {
		var ret int64
		return ret
	}
	return *o.SettlePlan
}

// GetSettlePlanOk returns a tuple with the SettlePlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetSettlePlanOk() (*int64, bool) {
	if o == nil || common.IsNil(o.SettlePlan) {
		return nil, false
	}
	return o.SettlePlan, true
}

// HasSettlePlan returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasSettlePlan() bool {
	if o != nil && !common.IsNil(o.SettlePlan) {
		return true
	}

	return false
}

// SetSettlePlan gets a reference to the given int64 and assigns it to the SettlePlan field.
func (o *ExchangeInformationResponseSymbolsInner) SetSettlePlan(v int64) {
	o.SettlePlan = &v
}

// GetTriggerProtect returns the TriggerProtect field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetTriggerProtect() string {
	if o == nil || common.IsNil(o.TriggerProtect) {
		var ret string
		return ret
	}
	return *o.TriggerProtect
}

// GetTriggerProtectOk returns a tuple with the TriggerProtect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetTriggerProtectOk() (*string, bool) {
	if o == nil || common.IsNil(o.TriggerProtect) {
		return nil, false
	}
	return o.TriggerProtect, true
}

// HasTriggerProtect returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasTriggerProtect() bool {
	if o != nil && !common.IsNil(o.TriggerProtect) {
		return true
	}

	return false
}

// SetTriggerProtect gets a reference to the given string and assigns it to the TriggerProtect field.
func (o *ExchangeInformationResponseSymbolsInner) SetTriggerProtect(v string) {
	o.TriggerProtect = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetFilters() []ExchangeInformationResponseSymbolsInnerFiltersInner {
	if o == nil || common.IsNil(o.Filters) {
		var ret []ExchangeInformationResponseSymbolsInnerFiltersInner
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetFiltersOk() ([]ExchangeInformationResponseSymbolsInnerFiltersInner, bool) {
	if o == nil || common.IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasFilters() bool {
	if o != nil && !common.IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []ExchangeInformationResponseSymbolsInnerFiltersInner and assigns it to the Filters field.
func (o *ExchangeInformationResponseSymbolsInner) SetFilters(v []ExchangeInformationResponseSymbolsInnerFiltersInner) {
	o.Filters = v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetOrderType() []string {
	if o == nil || common.IsNil(o.OrderType) {
		var ret []string
		return ret
	}
	return o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetOrderTypeOk() ([]string, bool) {
	if o == nil || common.IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasOrderType() bool {
	if o != nil && !common.IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given []string and assigns it to the OrderType field.
func (o *ExchangeInformationResponseSymbolsInner) SetOrderType(v []string) {
	o.OrderType = v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetTimeInForce() []string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret []string
		return ret
	}
	return o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetTimeInForceOk() ([]string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given []string and assigns it to the TimeInForce field.
func (o *ExchangeInformationResponseSymbolsInner) SetTimeInForce(v []string) {
	o.TimeInForce = v
}

// GetLiquidationFee returns the LiquidationFee field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetLiquidationFee() string {
	if o == nil || common.IsNil(o.LiquidationFee) {
		var ret string
		return ret
	}
	return *o.LiquidationFee
}

// GetLiquidationFeeOk returns a tuple with the LiquidationFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetLiquidationFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiquidationFee) {
		return nil, false
	}
	return o.LiquidationFee, true
}

// HasLiquidationFee returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasLiquidationFee() bool {
	if o != nil && !common.IsNil(o.LiquidationFee) {
		return true
	}

	return false
}

// SetLiquidationFee gets a reference to the given string and assigns it to the LiquidationFee field.
func (o *ExchangeInformationResponseSymbolsInner) SetLiquidationFee(v string) {
	o.LiquidationFee = &v
}

// GetMarketTakeBound returns the MarketTakeBound field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInner) GetMarketTakeBound() string {
	if o == nil || common.IsNil(o.MarketTakeBound) {
		var ret string
		return ret
	}
	return *o.MarketTakeBound
}

// GetMarketTakeBoundOk returns a tuple with the MarketTakeBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInner) GetMarketTakeBoundOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarketTakeBound) {
		return nil, false
	}
	return o.MarketTakeBound, true
}

// HasMarketTakeBound returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInner) HasMarketTakeBound() bool {
	if o != nil && !common.IsNil(o.MarketTakeBound) {
		return true
	}

	return false
}

// SetMarketTakeBound gets a reference to the given string and assigns it to the MarketTakeBound field.
func (o *ExchangeInformationResponseSymbolsInner) SetMarketTakeBound(v string) {
	o.MarketTakeBound = &v
}

func (o ExchangeInformationResponseSymbolsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeInformationResponseSymbolsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !common.IsNil(o.ContractType) {
		toSerialize["contractType"] = o.ContractType
	}
	if !common.IsNil(o.DeliveryDate) {
		toSerialize["deliveryDate"] = o.DeliveryDate
	}
	if !common.IsNil(o.OnboardDate) {
		toSerialize["onboardDate"] = o.OnboardDate
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.MaintMarginPercent) {
		toSerialize["maintMarginPercent"] = o.MaintMarginPercent
	}
	if !common.IsNil(o.RequiredMarginPercent) {
		toSerialize["requiredMarginPercent"] = o.RequiredMarginPercent
	}
	if !common.IsNil(o.BaseAsset) {
		toSerialize["baseAsset"] = o.BaseAsset
	}
	if !common.IsNil(o.QuoteAsset) {
		toSerialize["quoteAsset"] = o.QuoteAsset
	}
	if !common.IsNil(o.MarginAsset) {
		toSerialize["marginAsset"] = o.MarginAsset
	}
	if !common.IsNil(o.PricePrecision) {
		toSerialize["pricePrecision"] = o.PricePrecision
	}
	if !common.IsNil(o.QuantityPrecision) {
		toSerialize["quantityPrecision"] = o.QuantityPrecision
	}
	if !common.IsNil(o.BaseAssetPrecision) {
		toSerialize["baseAssetPrecision"] = o.BaseAssetPrecision
	}
	if !common.IsNil(o.QuotePrecision) {
		toSerialize["quotePrecision"] = o.QuotePrecision
	}
	if !common.IsNil(o.UnderlyingType) {
		toSerialize["underlyingType"] = o.UnderlyingType
	}
	if !common.IsNil(o.UnderlyingSubType) {
		toSerialize["underlyingSubType"] = o.UnderlyingSubType
	}
	if !common.IsNil(o.SettlePlan) {
		toSerialize["settlePlan"] = o.SettlePlan
	}
	if !common.IsNil(o.TriggerProtect) {
		toSerialize["triggerProtect"] = o.TriggerProtect
	}
	if !common.IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !common.IsNil(o.OrderType) {
		toSerialize["OrderType"] = o.OrderType
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.LiquidationFee) {
		toSerialize["liquidationFee"] = o.LiquidationFee
	}
	if !common.IsNil(o.MarketTakeBound) {
		toSerialize["marketTakeBound"] = o.MarketTakeBound
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExchangeInformationResponseSymbolsInner) UnmarshalJSON(data []byte) (err error) {
	varExchangeInformationResponseSymbolsInner := _ExchangeInformationResponseSymbolsInner{}

	err = json.Unmarshal(data, &varExchangeInformationResponseSymbolsInner)

	if err != nil {
		return err
	}

	*o = ExchangeInformationResponseSymbolsInner(varExchangeInformationResponseSymbolsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "pair")
		delete(additionalProperties, "contractType")
		delete(additionalProperties, "deliveryDate")
		delete(additionalProperties, "onboardDate")
		delete(additionalProperties, "status")
		delete(additionalProperties, "maintMarginPercent")
		delete(additionalProperties, "requiredMarginPercent")
		delete(additionalProperties, "baseAsset")
		delete(additionalProperties, "quoteAsset")
		delete(additionalProperties, "marginAsset")
		delete(additionalProperties, "pricePrecision")
		delete(additionalProperties, "quantityPrecision")
		delete(additionalProperties, "baseAssetPrecision")
		delete(additionalProperties, "quotePrecision")
		delete(additionalProperties, "underlyingType")
		delete(additionalProperties, "underlyingSubType")
		delete(additionalProperties, "settlePlan")
		delete(additionalProperties, "triggerProtect")
		delete(additionalProperties, "filters")
		delete(additionalProperties, "OrderType")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "liquidationFee")
		delete(additionalProperties, "marketTakeBound")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeInformationResponseSymbolsInner struct {
	value *ExchangeInformationResponseSymbolsInner
	isSet bool
}

func (v NullableExchangeInformationResponseSymbolsInner) Get() *ExchangeInformationResponseSymbolsInner {
	return v.value
}

func (v *NullableExchangeInformationResponseSymbolsInner) Set(val *ExchangeInformationResponseSymbolsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInformationResponseSymbolsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInformationResponseSymbolsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInformationResponseSymbolsInner(val *ExchangeInformationResponseSymbolsInner) *NullableExchangeInformationResponseSymbolsInner {
	return &NullableExchangeInformationResponseSymbolsInner{value: val, isSet: true}
}

func (v NullableExchangeInformationResponseSymbolsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInformationResponseSymbolsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
