/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OptionPositionInformationResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OptionPositionInformationResponseInner{}

// OptionPositionInformationResponseInner struct for OptionPositionInformationResponseInner
type OptionPositionInformationResponseInner struct {
	EntryPrice           *string `json:"entryPrice,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	Side                 *string `json:"side,omitempty"`
	Quantity             *string `json:"quantity,omitempty"`
	MarkValue            *string `json:"markValue,omitempty"`
	UnrealizedPNL        *string `json:"unrealizedPNL,omitempty"`
	MarkPrice            *string `json:"markPrice,omitempty"`
	StrikePrice          *string `json:"strikePrice,omitempty"`
	ExpiryDate           *int64  `json:"expiryDate,omitempty"`
	PriceScale           *int64  `json:"priceScale,omitempty"`
	QuantityScale        *int64  `json:"quantityScale,omitempty"`
	OptionSide           *string `json:"optionSide,omitempty"`
	QuoteAsset           *string `json:"quoteAsset,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	BidQuantity          *string `json:"bidQuantity,omitempty"`
	AskQuantity          *string `json:"askQuantity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OptionPositionInformationResponseInner OptionPositionInformationResponseInner

// NewOptionPositionInformationResponseInner instantiates a new OptionPositionInformationResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionPositionInformationResponseInner() *OptionPositionInformationResponseInner {
	this := OptionPositionInformationResponseInner{}
	return &this
}

// NewOptionPositionInformationResponseInnerWithDefaults instantiates a new OptionPositionInformationResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionPositionInformationResponseInnerWithDefaults() *OptionPositionInformationResponseInner {
	this := OptionPositionInformationResponseInner{}
	return &this
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *OptionPositionInformationResponseInner) GetEntryPrice() string {
	if o == nil || common.IsNil(o.EntryPrice) {
		var ret string
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionPositionInformationResponseInner) GetEntryPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *OptionPositionInformationResponseInner) HasEntryPrice() bool {
	if o != nil && !common.IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given string and assigns it to the EntryPrice field.
func (o *OptionPositionInformationResponseInner) SetEntryPrice(v string) {
	o.EntryPrice = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OptionPositionInformationResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionPositionInformationResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OptionPositionInformationResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OptionPositionInformationResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *OptionPositionInformationResponseInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionPositionInformationResponseInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *OptionPositionInformationResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *OptionPositionInformationResponseInner) SetSide(v string) {
	o.Side = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OptionPositionInformationResponseInner) GetQuantity() string {
	if o == nil || common.IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionPositionInformationResponseInner) GetQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OptionPositionInformationResponseInner) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *OptionPositionInformationResponseInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetMarkValue returns the MarkValue field value if set, zero value otherwise.
func (o *OptionPositionInformationResponseInner) GetMarkValue() string {
	if o == nil || common.IsNil(o.MarkValue) {
		var ret string
		return ret
	}
	return *o.MarkValue
}

// GetMarkValueOk returns a tuple with the MarkValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionPositionInformationResponseInner) GetMarkValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarkValue) {
		return nil, false
	}
	return o.MarkValue, true
}

// HasMarkValue returns a boolean if a field has been set.
func (o *OptionPositionInformationResponseInner) HasMarkValue() bool {
	if o != nil && !common.IsNil(o.MarkValue) {
		return true
	}

	return false
}

// SetMarkValue gets a reference to the given string and assigns it to the MarkValue field.
func (o *OptionPositionInformationResponseInner) SetMarkValue(v string) {
	o.MarkValue = &v
}

// GetUnrealizedPNL returns the UnrealizedPNL field value if set, zero value otherwise.
func (o *OptionPositionInformationResponseInner) GetUnrealizedPNL() string {
	if o == nil || common.IsNil(o.UnrealizedPNL) {
		var ret string
		return ret
	}
	return *o.UnrealizedPNL
}

// GetUnrealizedPNLOk returns a tuple with the UnrealizedPNL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionPositionInformationResponseInner) GetUnrealizedPNLOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnrealizedPNL) {
		return nil, false
	}
	return o.UnrealizedPNL, true
}

// HasUnrealizedPNL returns a boolean if a field has been set.
func (o *OptionPositionInformationResponseInner) HasUnrealizedPNL() bool {
	if o != nil && !common.IsNil(o.UnrealizedPNL) {
		return true
	}

	return false
}

// SetUnrealizedPNL gets a reference to the given string and assigns it to the UnrealizedPNL field.
func (o *OptionPositionInformationResponseInner) SetUnrealizedPNL(v string) {
	o.UnrealizedPNL = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *OptionPositionInformationResponseInner) GetMarkPrice() string {
	if o == nil || common.IsNil(o.MarkPrice) {
		var ret string
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionPositionInformationResponseInner) GetMarkPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *OptionPositionInformationResponseInner) HasMarkPrice() bool {
	if o != nil && !common.IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given string and assigns it to the MarkPrice field.
func (o *OptionPositionInformationResponseInner) SetMarkPrice(v string) {
	o.MarkPrice = &v
}

// GetStrikePrice returns the StrikePrice field value if set, zero value otherwise.
func (o *OptionPositionInformationResponseInner) GetStrikePrice() string {
	if o == nil || common.IsNil(o.StrikePrice) {
		var ret string
		return ret
	}
	return *o.StrikePrice
}

// GetStrikePriceOk returns a tuple with the StrikePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionPositionInformationResponseInner) GetStrikePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StrikePrice) {
		return nil, false
	}
	return o.StrikePrice, true
}

// HasStrikePrice returns a boolean if a field has been set.
func (o *OptionPositionInformationResponseInner) HasStrikePrice() bool {
	if o != nil && !common.IsNil(o.StrikePrice) {
		return true
	}

	return false
}

// SetStrikePrice gets a reference to the given string and assigns it to the StrikePrice field.
func (o *OptionPositionInformationResponseInner) SetStrikePrice(v string) {
	o.StrikePrice = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *OptionPositionInformationResponseInner) GetExpiryDate() int64 {
	if o == nil || common.IsNil(o.ExpiryDate) {
		var ret int64
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionPositionInformationResponseInner) GetExpiryDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExpiryDate) {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *OptionPositionInformationResponseInner) HasExpiryDate() bool {
	if o != nil && !common.IsNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given int64 and assigns it to the ExpiryDate field.
func (o *OptionPositionInformationResponseInner) SetExpiryDate(v int64) {
	o.ExpiryDate = &v
}

// GetPriceScale returns the PriceScale field value if set, zero value otherwise.
func (o *OptionPositionInformationResponseInner) GetPriceScale() int64 {
	if o == nil || common.IsNil(o.PriceScale) {
		var ret int64
		return ret
	}
	return *o.PriceScale
}

// GetPriceScaleOk returns a tuple with the PriceScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionPositionInformationResponseInner) GetPriceScaleOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PriceScale) {
		return nil, false
	}
	return o.PriceScale, true
}

// HasPriceScale returns a boolean if a field has been set.
func (o *OptionPositionInformationResponseInner) HasPriceScale() bool {
	if o != nil && !common.IsNil(o.PriceScale) {
		return true
	}

	return false
}

// SetPriceScale gets a reference to the given int64 and assigns it to the PriceScale field.
func (o *OptionPositionInformationResponseInner) SetPriceScale(v int64) {
	o.PriceScale = &v
}

// GetQuantityScale returns the QuantityScale field value if set, zero value otherwise.
func (o *OptionPositionInformationResponseInner) GetQuantityScale() int64 {
	if o == nil || common.IsNil(o.QuantityScale) {
		var ret int64
		return ret
	}
	return *o.QuantityScale
}

// GetQuantityScaleOk returns a tuple with the QuantityScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionPositionInformationResponseInner) GetQuantityScaleOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QuantityScale) {
		return nil, false
	}
	return o.QuantityScale, true
}

// HasQuantityScale returns a boolean if a field has been set.
func (o *OptionPositionInformationResponseInner) HasQuantityScale() bool {
	if o != nil && !common.IsNil(o.QuantityScale) {
		return true
	}

	return false
}

// SetQuantityScale gets a reference to the given int64 and assigns it to the QuantityScale field.
func (o *OptionPositionInformationResponseInner) SetQuantityScale(v int64) {
	o.QuantityScale = &v
}

// GetOptionSide returns the OptionSide field value if set, zero value otherwise.
func (o *OptionPositionInformationResponseInner) GetOptionSide() string {
	if o == nil || common.IsNil(o.OptionSide) {
		var ret string
		return ret
	}
	return *o.OptionSide
}

// GetOptionSideOk returns a tuple with the OptionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionPositionInformationResponseInner) GetOptionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.OptionSide) {
		return nil, false
	}
	return o.OptionSide, true
}

// HasOptionSide returns a boolean if a field has been set.
func (o *OptionPositionInformationResponseInner) HasOptionSide() bool {
	if o != nil && !common.IsNil(o.OptionSide) {
		return true
	}

	return false
}

// SetOptionSide gets a reference to the given string and assigns it to the OptionSide field.
func (o *OptionPositionInformationResponseInner) SetOptionSide(v string) {
	o.OptionSide = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *OptionPositionInformationResponseInner) GetQuoteAsset() string {
	if o == nil || common.IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionPositionInformationResponseInner) GetQuoteAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *OptionPositionInformationResponseInner) HasQuoteAsset() bool {
	if o != nil && !common.IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *OptionPositionInformationResponseInner) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *OptionPositionInformationResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionPositionInformationResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *OptionPositionInformationResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *OptionPositionInformationResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetBidQuantity returns the BidQuantity field value if set, zero value otherwise.
func (o *OptionPositionInformationResponseInner) GetBidQuantity() string {
	if o == nil || common.IsNil(o.BidQuantity) {
		var ret string
		return ret
	}
	return *o.BidQuantity
}

// GetBidQuantityOk returns a tuple with the BidQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionPositionInformationResponseInner) GetBidQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidQuantity) {
		return nil, false
	}
	return o.BidQuantity, true
}

// HasBidQuantity returns a boolean if a field has been set.
func (o *OptionPositionInformationResponseInner) HasBidQuantity() bool {
	if o != nil && !common.IsNil(o.BidQuantity) {
		return true
	}

	return false
}

// SetBidQuantity gets a reference to the given string and assigns it to the BidQuantity field.
func (o *OptionPositionInformationResponseInner) SetBidQuantity(v string) {
	o.BidQuantity = &v
}

// GetAskQuantity returns the AskQuantity field value if set, zero value otherwise.
func (o *OptionPositionInformationResponseInner) GetAskQuantity() string {
	if o == nil || common.IsNil(o.AskQuantity) {
		var ret string
		return ret
	}
	return *o.AskQuantity
}

// GetAskQuantityOk returns a tuple with the AskQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionPositionInformationResponseInner) GetAskQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskQuantity) {
		return nil, false
	}
	return o.AskQuantity, true
}

// HasAskQuantity returns a boolean if a field has been set.
func (o *OptionPositionInformationResponseInner) HasAskQuantity() bool {
	if o != nil && !common.IsNil(o.AskQuantity) {
		return true
	}

	return false
}

// SetAskQuantity gets a reference to the given string and assigns it to the AskQuantity field.
func (o *OptionPositionInformationResponseInner) SetAskQuantity(v string) {
	o.AskQuantity = &v
}

func (o OptionPositionInformationResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionPositionInformationResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EntryPrice) {
		toSerialize["entryPrice"] = o.EntryPrice
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !common.IsNil(o.MarkValue) {
		toSerialize["markValue"] = o.MarkValue
	}
	if !common.IsNil(o.UnrealizedPNL) {
		toSerialize["unrealizedPNL"] = o.UnrealizedPNL
	}
	if !common.IsNil(o.MarkPrice) {
		toSerialize["markPrice"] = o.MarkPrice
	}
	if !common.IsNil(o.StrikePrice) {
		toSerialize["strikePrice"] = o.StrikePrice
	}
	if !common.IsNil(o.ExpiryDate) {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if !common.IsNil(o.PriceScale) {
		toSerialize["priceScale"] = o.PriceScale
	}
	if !common.IsNil(o.QuantityScale) {
		toSerialize["quantityScale"] = o.QuantityScale
	}
	if !common.IsNil(o.OptionSide) {
		toSerialize["optionSide"] = o.OptionSide
	}
	if !common.IsNil(o.QuoteAsset) {
		toSerialize["quoteAsset"] = o.QuoteAsset
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.BidQuantity) {
		toSerialize["bidQuantity"] = o.BidQuantity
	}
	if !common.IsNil(o.AskQuantity) {
		toSerialize["askQuantity"] = o.AskQuantity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OptionPositionInformationResponseInner) UnmarshalJSON(data []byte) (err error) {
	varOptionPositionInformationResponseInner := _OptionPositionInformationResponseInner{}

	err = json.Unmarshal(data, &varOptionPositionInformationResponseInner)

	if err != nil {
		return err
	}

	*o = OptionPositionInformationResponseInner(varOptionPositionInformationResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entryPrice")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "side")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "markValue")
		delete(additionalProperties, "unrealizedPNL")
		delete(additionalProperties, "markPrice")
		delete(additionalProperties, "strikePrice")
		delete(additionalProperties, "expiryDate")
		delete(additionalProperties, "priceScale")
		delete(additionalProperties, "quantityScale")
		delete(additionalProperties, "optionSide")
		delete(additionalProperties, "quoteAsset")
		delete(additionalProperties, "time")
		delete(additionalProperties, "bidQuantity")
		delete(additionalProperties, "askQuantity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOptionPositionInformationResponseInner struct {
	value *OptionPositionInformationResponseInner
	isSet bool
}

func (v NullableOptionPositionInformationResponseInner) Get() *OptionPositionInformationResponseInner {
	return v.value
}

func (v *NullableOptionPositionInformationResponseInner) Set(val *OptionPositionInformationResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionPositionInformationResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionPositionInformationResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionPositionInformationResponseInner(val *OptionPositionInformationResponseInner) *NullableOptionPositionInformationResponseInner {
	return &NullableOptionPositionInformationResponseInner{value: val, isSet: true}
}

func (v NullableOptionPositionInformationResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionPositionInformationResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
