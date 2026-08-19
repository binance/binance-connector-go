/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExchangeInfoResponseSymbolsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeInfoResponseSymbolsInner{}

// ExchangeInfoResponseSymbolsInner struct for ExchangeInfoResponseSymbolsInner
type ExchangeInfoResponseSymbolsInner struct {
	// US-equity ticker, e.g. `AAPL`.
	Symbol *string `json:"symbol,omitempty"`
	// Trading direction allowed — one of `BUY_SELL` / `BUY` / `SELL` / `NONE`.
	Tradability *string `json:"tradability,omitempty"`
	// Last time the `tradability` value was updated (ms epoch).
	TradabilityUpdateTime *int64 `json:"tradabilityUpdateTime,omitempty"`
	// Whether the symbol supports overnight trading.
	OvernightSupported *bool `json:"overnightSupported,omitempty"`
	// Whether fractional shares are supported during the regular session.
	Fractionable *bool `json:"fractionable,omitempty"`
	// Whether fractional shares are supported during extended hours.
	FractionableEh *bool `json:"fractionableEh,omitempty"`
	// Whether extended-session trading is enabled.
	ExtendedSession *bool `json:"extendedSession,omitempty"`
	// Maximum number of open orders a user may have for this symbol.
	MaxNumOrders *int32 `json:"maxNumOrders,omitempty"`
	// Lot size — minimum increment for `quantity`.
	StepSize *string `json:"stepSize,omitempty"`
	// Upper price multiplier limit relative to reference.
	MultiplierUp *string `json:"multiplierUp,omitempty"`
	// Lower price multiplier limit relative to reference.
	MultiplierDown *string `json:"multiplierDown,omitempty"`
	// Minimum allowed `quantity`.
	MinQty *string `json:"minQty,omitempty"`
	// Maximum allowed `quantity`.
	MaxQty *string `json:"maxQty,omitempty"`
	// Minimum order notional (USD).
	MinNotional *string `json:"minNotional,omitempty"`
	// Maximum order notional (USD).
	MaxNotional *string `json:"maxNotional,omitempty"`
	// Listing timestamp (ms epoch).
	ListingTime *int64 `json:"listingTime,omitempty"`
	// Scheduled delisting timestamp (ms epoch); `null` if not scheduled for delisting.
	DelistingTime        *int64 `json:"delistingTime,omitempty"`
	AdditionalProperties map[string]interface{}
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

// GetTradability returns the Tradability field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetTradability() string {
	if o == nil || common.IsNil(o.Tradability) {
		var ret string
		return ret
	}
	return *o.Tradability
}

// GetTradabilityOk returns a tuple with the Tradability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetTradabilityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Tradability) {
		return nil, false
	}
	return o.Tradability, true
}

// HasTradability returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasTradability() bool {
	if o != nil && !common.IsNil(o.Tradability) {
		return true
	}

	return false
}

// SetTradability gets a reference to the given string and assigns it to the Tradability field.
func (o *ExchangeInfoResponseSymbolsInner) SetTradability(v string) {
	o.Tradability = &v
}

// GetTradabilityUpdateTime returns the TradabilityUpdateTime field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetTradabilityUpdateTime() int64 {
	if o == nil || common.IsNil(o.TradabilityUpdateTime) {
		var ret int64
		return ret
	}
	return *o.TradabilityUpdateTime
}

// GetTradabilityUpdateTimeOk returns a tuple with the TradabilityUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetTradabilityUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradabilityUpdateTime) {
		return nil, false
	}
	return o.TradabilityUpdateTime, true
}

// HasTradabilityUpdateTime returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasTradabilityUpdateTime() bool {
	if o != nil && !common.IsNil(o.TradabilityUpdateTime) {
		return true
	}

	return false
}

// SetTradabilityUpdateTime gets a reference to the given int64 and assigns it to the TradabilityUpdateTime field.
func (o *ExchangeInfoResponseSymbolsInner) SetTradabilityUpdateTime(v int64) {
	o.TradabilityUpdateTime = &v
}

// GetOvernightSupported returns the OvernightSupported field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetOvernightSupported() bool {
	if o == nil || common.IsNil(o.OvernightSupported) {
		var ret bool
		return ret
	}
	return *o.OvernightSupported
}

// GetOvernightSupportedOk returns a tuple with the OvernightSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetOvernightSupportedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.OvernightSupported) {
		return nil, false
	}
	return o.OvernightSupported, true
}

// HasOvernightSupported returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasOvernightSupported() bool {
	if o != nil && !common.IsNil(o.OvernightSupported) {
		return true
	}

	return false
}

// SetOvernightSupported gets a reference to the given bool and assigns it to the OvernightSupported field.
func (o *ExchangeInfoResponseSymbolsInner) SetOvernightSupported(v bool) {
	o.OvernightSupported = &v
}

// GetFractionable returns the Fractionable field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetFractionable() bool {
	if o == nil || common.IsNil(o.Fractionable) {
		var ret bool
		return ret
	}
	return *o.Fractionable
}

// GetFractionableOk returns a tuple with the Fractionable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetFractionableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Fractionable) {
		return nil, false
	}
	return o.Fractionable, true
}

// HasFractionable returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasFractionable() bool {
	if o != nil && !common.IsNil(o.Fractionable) {
		return true
	}

	return false
}

// SetFractionable gets a reference to the given bool and assigns it to the Fractionable field.
func (o *ExchangeInfoResponseSymbolsInner) SetFractionable(v bool) {
	o.Fractionable = &v
}

// GetFractionableEh returns the FractionableEh field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetFractionableEh() bool {
	if o == nil || common.IsNil(o.FractionableEh) {
		var ret bool
		return ret
	}
	return *o.FractionableEh
}

// GetFractionableEhOk returns a tuple with the FractionableEh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetFractionableEhOk() (*bool, bool) {
	if o == nil || common.IsNil(o.FractionableEh) {
		return nil, false
	}
	return o.FractionableEh, true
}

// HasFractionableEh returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasFractionableEh() bool {
	if o != nil && !common.IsNil(o.FractionableEh) {
		return true
	}

	return false
}

// SetFractionableEh gets a reference to the given bool and assigns it to the FractionableEh field.
func (o *ExchangeInfoResponseSymbolsInner) SetFractionableEh(v bool) {
	o.FractionableEh = &v
}

// GetExtendedSession returns the ExtendedSession field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetExtendedSession() bool {
	if o == nil || common.IsNil(o.ExtendedSession) {
		var ret bool
		return ret
	}
	return *o.ExtendedSession
}

// GetExtendedSessionOk returns a tuple with the ExtendedSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetExtendedSessionOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ExtendedSession) {
		return nil, false
	}
	return o.ExtendedSession, true
}

// HasExtendedSession returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasExtendedSession() bool {
	if o != nil && !common.IsNil(o.ExtendedSession) {
		return true
	}

	return false
}

// SetExtendedSession gets a reference to the given bool and assigns it to the ExtendedSession field.
func (o *ExchangeInfoResponseSymbolsInner) SetExtendedSession(v bool) {
	o.ExtendedSession = &v
}

// GetMaxNumOrders returns the MaxNumOrders field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetMaxNumOrders() int32 {
	if o == nil || common.IsNil(o.MaxNumOrders) {
		var ret int32
		return ret
	}
	return *o.MaxNumOrders
}

// GetMaxNumOrdersOk returns a tuple with the MaxNumOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetMaxNumOrdersOk() (*int32, bool) {
	if o == nil || common.IsNil(o.MaxNumOrders) {
		return nil, false
	}
	return o.MaxNumOrders, true
}

// HasMaxNumOrders returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasMaxNumOrders() bool {
	if o != nil && !common.IsNil(o.MaxNumOrders) {
		return true
	}

	return false
}

// SetMaxNumOrders gets a reference to the given int32 and assigns it to the MaxNumOrders field.
func (o *ExchangeInfoResponseSymbolsInner) SetMaxNumOrders(v int32) {
	o.MaxNumOrders = &v
}

// GetStepSize returns the StepSize field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetStepSize() string {
	if o == nil || common.IsNil(o.StepSize) {
		var ret string
		return ret
	}
	return *o.StepSize
}

// GetStepSizeOk returns a tuple with the StepSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetStepSizeOk() (*string, bool) {
	if o == nil || common.IsNil(o.StepSize) {
		return nil, false
	}
	return o.StepSize, true
}

// HasStepSize returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasStepSize() bool {
	if o != nil && !common.IsNil(o.StepSize) {
		return true
	}

	return false
}

// SetStepSize gets a reference to the given string and assigns it to the StepSize field.
func (o *ExchangeInfoResponseSymbolsInner) SetStepSize(v string) {
	o.StepSize = &v
}

// GetMultiplierUp returns the MultiplierUp field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetMultiplierUp() string {
	if o == nil || common.IsNil(o.MultiplierUp) {
		var ret string
		return ret
	}
	return *o.MultiplierUp
}

// GetMultiplierUpOk returns a tuple with the MultiplierUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetMultiplierUpOk() (*string, bool) {
	if o == nil || common.IsNil(o.MultiplierUp) {
		return nil, false
	}
	return o.MultiplierUp, true
}

// HasMultiplierUp returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasMultiplierUp() bool {
	if o != nil && !common.IsNil(o.MultiplierUp) {
		return true
	}

	return false
}

// SetMultiplierUp gets a reference to the given string and assigns it to the MultiplierUp field.
func (o *ExchangeInfoResponseSymbolsInner) SetMultiplierUp(v string) {
	o.MultiplierUp = &v
}

// GetMultiplierDown returns the MultiplierDown field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetMultiplierDown() string {
	if o == nil || common.IsNil(o.MultiplierDown) {
		var ret string
		return ret
	}
	return *o.MultiplierDown
}

// GetMultiplierDownOk returns a tuple with the MultiplierDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetMultiplierDownOk() (*string, bool) {
	if o == nil || common.IsNil(o.MultiplierDown) {
		return nil, false
	}
	return o.MultiplierDown, true
}

// HasMultiplierDown returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasMultiplierDown() bool {
	if o != nil && !common.IsNil(o.MultiplierDown) {
		return true
	}

	return false
}

// SetMultiplierDown gets a reference to the given string and assigns it to the MultiplierDown field.
func (o *ExchangeInfoResponseSymbolsInner) SetMultiplierDown(v string) {
	o.MultiplierDown = &v
}

// GetMinQty returns the MinQty field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetMinQty() string {
	if o == nil || common.IsNil(o.MinQty) {
		var ret string
		return ret
	}
	return *o.MinQty
}

// GetMinQtyOk returns a tuple with the MinQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetMinQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinQty) {
		return nil, false
	}
	return o.MinQty, true
}

// HasMinQty returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasMinQty() bool {
	if o != nil && !common.IsNil(o.MinQty) {
		return true
	}

	return false
}

// SetMinQty gets a reference to the given string and assigns it to the MinQty field.
func (o *ExchangeInfoResponseSymbolsInner) SetMinQty(v string) {
	o.MinQty = &v
}

// GetMaxQty returns the MaxQty field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetMaxQty() string {
	if o == nil || common.IsNil(o.MaxQty) {
		var ret string
		return ret
	}
	return *o.MaxQty
}

// GetMaxQtyOk returns a tuple with the MaxQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetMaxQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxQty) {
		return nil, false
	}
	return o.MaxQty, true
}

// HasMaxQty returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasMaxQty() bool {
	if o != nil && !common.IsNil(o.MaxQty) {
		return true
	}

	return false
}

// SetMaxQty gets a reference to the given string and assigns it to the MaxQty field.
func (o *ExchangeInfoResponseSymbolsInner) SetMaxQty(v string) {
	o.MaxQty = &v
}

// GetMinNotional returns the MinNotional field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetMinNotional() string {
	if o == nil || common.IsNil(o.MinNotional) {
		var ret string
		return ret
	}
	return *o.MinNotional
}

// GetMinNotionalOk returns a tuple with the MinNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetMinNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinNotional) {
		return nil, false
	}
	return o.MinNotional, true
}

// HasMinNotional returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasMinNotional() bool {
	if o != nil && !common.IsNil(o.MinNotional) {
		return true
	}

	return false
}

// SetMinNotional gets a reference to the given string and assigns it to the MinNotional field.
func (o *ExchangeInfoResponseSymbolsInner) SetMinNotional(v string) {
	o.MinNotional = &v
}

// GetMaxNotional returns the MaxNotional field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetMaxNotional() string {
	if o == nil || common.IsNil(o.MaxNotional) {
		var ret string
		return ret
	}
	return *o.MaxNotional
}

// GetMaxNotionalOk returns a tuple with the MaxNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetMaxNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxNotional) {
		return nil, false
	}
	return o.MaxNotional, true
}

// HasMaxNotional returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasMaxNotional() bool {
	if o != nil && !common.IsNil(o.MaxNotional) {
		return true
	}

	return false
}

// SetMaxNotional gets a reference to the given string and assigns it to the MaxNotional field.
func (o *ExchangeInfoResponseSymbolsInner) SetMaxNotional(v string) {
	o.MaxNotional = &v
}

// GetListingTime returns the ListingTime field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSymbolsInner) GetListingTime() int64 {
	if o == nil || common.IsNil(o.ListingTime) {
		var ret int64
		return ret
	}
	return *o.ListingTime
}

// GetListingTimeOk returns a tuple with the ListingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSymbolsInner) GetListingTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ListingTime) {
		return nil, false
	}
	return o.ListingTime, true
}

// HasListingTime returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasListingTime() bool {
	if o != nil && !common.IsNil(o.ListingTime) {
		return true
	}

	return false
}

// SetListingTime gets a reference to the given int64 and assigns it to the ListingTime field.
func (o *ExchangeInfoResponseSymbolsInner) SetListingTime(v int64) {
	o.ListingTime = &v
}

// GetDelistingTime returns the DelistingTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExchangeInfoResponseSymbolsInner) GetDelistingTime() int64 {
	if o == nil || common.IsNil(o.DelistingTime) {
		var ret int64
		return ret
	}
	return *o.DelistingTime
}

// GetDelistingTimeOk returns a tuple with the DelistingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExchangeInfoResponseSymbolsInner) GetDelistingTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.DelistingTime) {
		return nil, false
	}
	return o.DelistingTime, true
}

// HasDelistingTime returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSymbolsInner) HasDelistingTime() bool {
	if o != nil && !common.IsNil(o.DelistingTime) {
		return true
	}

	return false
}

// SetDelistingTime gets a reference to the given NullableInt64 and assigns it to the DelistingTime field.
func (o *ExchangeInfoResponseSymbolsInner) SetDelistingTime(v int64) {
	o.DelistingTime = &v
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
	if !common.IsNil(o.Tradability) {
		toSerialize["tradability"] = o.Tradability
	}
	if !common.IsNil(o.TradabilityUpdateTime) {
		toSerialize["tradabilityUpdateTime"] = o.TradabilityUpdateTime
	}
	if !common.IsNil(o.OvernightSupported) {
		toSerialize["overnightSupported"] = o.OvernightSupported
	}
	if !common.IsNil(o.Fractionable) {
		toSerialize["fractionable"] = o.Fractionable
	}
	if !common.IsNil(o.FractionableEh) {
		toSerialize["fractionableEh"] = o.FractionableEh
	}
	if !common.IsNil(o.ExtendedSession) {
		toSerialize["extendedSession"] = o.ExtendedSession
	}
	if !common.IsNil(o.MaxNumOrders) {
		toSerialize["maxNumOrders"] = o.MaxNumOrders
	}
	if !common.IsNil(o.StepSize) {
		toSerialize["stepSize"] = o.StepSize
	}
	if !common.IsNil(o.MultiplierUp) {
		toSerialize["multiplierUp"] = o.MultiplierUp
	}
	if !common.IsNil(o.MultiplierDown) {
		toSerialize["multiplierDown"] = o.MultiplierDown
	}
	if !common.IsNil(o.MinQty) {
		toSerialize["minQty"] = o.MinQty
	}
	if !common.IsNil(o.MaxQty) {
		toSerialize["maxQty"] = o.MaxQty
	}
	if !common.IsNil(o.MinNotional) {
		toSerialize["minNotional"] = o.MinNotional
	}
	if !common.IsNil(o.MaxNotional) {
		toSerialize["maxNotional"] = o.MaxNotional
	}
	if !common.IsNil(o.ListingTime) {
		toSerialize["listingTime"] = o.ListingTime
	}
	if !common.IsNil(o.DelistingTime) {
		toSerialize["delistingTime"] = o.DelistingTime
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
		delete(additionalProperties, "tradability")
		delete(additionalProperties, "tradabilityUpdateTime")
		delete(additionalProperties, "overnightSupported")
		delete(additionalProperties, "fractionable")
		delete(additionalProperties, "fractionableEh")
		delete(additionalProperties, "extendedSession")
		delete(additionalProperties, "maxNumOrders")
		delete(additionalProperties, "stepSize")
		delete(additionalProperties, "multiplierUp")
		delete(additionalProperties, "multiplierDown")
		delete(additionalProperties, "minQty")
		delete(additionalProperties, "maxQty")
		delete(additionalProperties, "minNotional")
		delete(additionalProperties, "maxNotional")
		delete(additionalProperties, "listingTime")
		delete(additionalProperties, "delistingTime")
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
