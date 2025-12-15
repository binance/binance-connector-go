/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryIsolatedMarginTierDataResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryIsolatedMarginTierDataResponseInner{}

// QueryIsolatedMarginTierDataResponseInner struct for QueryIsolatedMarginTierDataResponseInner
type QueryIsolatedMarginTierDataResponseInner struct {
	Symbol                  *string `json:"symbol,omitempty"`
	Tier                    *int64  `json:"tier,omitempty"`
	EffectiveMultiple       *string `json:"effectiveMultiple,omitempty"`
	InitialRiskRatio        *string `json:"initialRiskRatio,omitempty"`
	LiquidationRiskRatio    *string `json:"liquidationRiskRatio,omitempty"`
	BaseAssetMaxBorrowable  *string `json:"baseAssetMaxBorrowable,omitempty"`
	QuoteAssetMaxBorrowable *string `json:"quoteAssetMaxBorrowable,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _QueryIsolatedMarginTierDataResponseInner QueryIsolatedMarginTierDataResponseInner

// NewQueryIsolatedMarginTierDataResponseInner instantiates a new QueryIsolatedMarginTierDataResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryIsolatedMarginTierDataResponseInner() *QueryIsolatedMarginTierDataResponseInner {
	this := QueryIsolatedMarginTierDataResponseInner{}
	return &this
}

// NewQueryIsolatedMarginTierDataResponseInnerWithDefaults instantiates a new QueryIsolatedMarginTierDataResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryIsolatedMarginTierDataResponseInnerWithDefaults() *QueryIsolatedMarginTierDataResponseInner {
	this := QueryIsolatedMarginTierDataResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryIsolatedMarginTierDataResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginTierDataResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryIsolatedMarginTierDataResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryIsolatedMarginTierDataResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *QueryIsolatedMarginTierDataResponseInner) GetTier() int64 {
	if o == nil || common.IsNil(o.Tier) {
		var ret int64
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginTierDataResponseInner) GetTierOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *QueryIsolatedMarginTierDataResponseInner) HasTier() bool {
	if o != nil && !common.IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given int64 and assigns it to the Tier field.
func (o *QueryIsolatedMarginTierDataResponseInner) SetTier(v int64) {
	o.Tier = &v
}

// GetEffectiveMultiple returns the EffectiveMultiple field value if set, zero value otherwise.
func (o *QueryIsolatedMarginTierDataResponseInner) GetEffectiveMultiple() string {
	if o == nil || common.IsNil(o.EffectiveMultiple) {
		var ret string
		return ret
	}
	return *o.EffectiveMultiple
}

// GetEffectiveMultipleOk returns a tuple with the EffectiveMultiple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginTierDataResponseInner) GetEffectiveMultipleOk() (*string, bool) {
	if o == nil || common.IsNil(o.EffectiveMultiple) {
		return nil, false
	}
	return o.EffectiveMultiple, true
}

// HasEffectiveMultiple returns a boolean if a field has been set.
func (o *QueryIsolatedMarginTierDataResponseInner) HasEffectiveMultiple() bool {
	if o != nil && !common.IsNil(o.EffectiveMultiple) {
		return true
	}

	return false
}

// SetEffectiveMultiple gets a reference to the given string and assigns it to the EffectiveMultiple field.
func (o *QueryIsolatedMarginTierDataResponseInner) SetEffectiveMultiple(v string) {
	o.EffectiveMultiple = &v
}

// GetInitialRiskRatio returns the InitialRiskRatio field value if set, zero value otherwise.
func (o *QueryIsolatedMarginTierDataResponseInner) GetInitialRiskRatio() string {
	if o == nil || common.IsNil(o.InitialRiskRatio) {
		var ret string
		return ret
	}
	return *o.InitialRiskRatio
}

// GetInitialRiskRatioOk returns a tuple with the InitialRiskRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginTierDataResponseInner) GetInitialRiskRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialRiskRatio) {
		return nil, false
	}
	return o.InitialRiskRatio, true
}

// HasInitialRiskRatio returns a boolean if a field has been set.
func (o *QueryIsolatedMarginTierDataResponseInner) HasInitialRiskRatio() bool {
	if o != nil && !common.IsNil(o.InitialRiskRatio) {
		return true
	}

	return false
}

// SetInitialRiskRatio gets a reference to the given string and assigns it to the InitialRiskRatio field.
func (o *QueryIsolatedMarginTierDataResponseInner) SetInitialRiskRatio(v string) {
	o.InitialRiskRatio = &v
}

// GetLiquidationRiskRatio returns the LiquidationRiskRatio field value if set, zero value otherwise.
func (o *QueryIsolatedMarginTierDataResponseInner) GetLiquidationRiskRatio() string {
	if o == nil || common.IsNil(o.LiquidationRiskRatio) {
		var ret string
		return ret
	}
	return *o.LiquidationRiskRatio
}

// GetLiquidationRiskRatioOk returns a tuple with the LiquidationRiskRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginTierDataResponseInner) GetLiquidationRiskRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiquidationRiskRatio) {
		return nil, false
	}
	return o.LiquidationRiskRatio, true
}

// HasLiquidationRiskRatio returns a boolean if a field has been set.
func (o *QueryIsolatedMarginTierDataResponseInner) HasLiquidationRiskRatio() bool {
	if o != nil && !common.IsNil(o.LiquidationRiskRatio) {
		return true
	}

	return false
}

// SetLiquidationRiskRatio gets a reference to the given string and assigns it to the LiquidationRiskRatio field.
func (o *QueryIsolatedMarginTierDataResponseInner) SetLiquidationRiskRatio(v string) {
	o.LiquidationRiskRatio = &v
}

// GetBaseAssetMaxBorrowable returns the BaseAssetMaxBorrowable field value if set, zero value otherwise.
func (o *QueryIsolatedMarginTierDataResponseInner) GetBaseAssetMaxBorrowable() string {
	if o == nil || common.IsNil(o.BaseAssetMaxBorrowable) {
		var ret string
		return ret
	}
	return *o.BaseAssetMaxBorrowable
}

// GetBaseAssetMaxBorrowableOk returns a tuple with the BaseAssetMaxBorrowable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginTierDataResponseInner) GetBaseAssetMaxBorrowableOk() (*string, bool) {
	if o == nil || common.IsNil(o.BaseAssetMaxBorrowable) {
		return nil, false
	}
	return o.BaseAssetMaxBorrowable, true
}

// HasBaseAssetMaxBorrowable returns a boolean if a field has been set.
func (o *QueryIsolatedMarginTierDataResponseInner) HasBaseAssetMaxBorrowable() bool {
	if o != nil && !common.IsNil(o.BaseAssetMaxBorrowable) {
		return true
	}

	return false
}

// SetBaseAssetMaxBorrowable gets a reference to the given string and assigns it to the BaseAssetMaxBorrowable field.
func (o *QueryIsolatedMarginTierDataResponseInner) SetBaseAssetMaxBorrowable(v string) {
	o.BaseAssetMaxBorrowable = &v
}

// GetQuoteAssetMaxBorrowable returns the QuoteAssetMaxBorrowable field value if set, zero value otherwise.
func (o *QueryIsolatedMarginTierDataResponseInner) GetQuoteAssetMaxBorrowable() string {
	if o == nil || common.IsNil(o.QuoteAssetMaxBorrowable) {
		var ret string
		return ret
	}
	return *o.QuoteAssetMaxBorrowable
}

// GetQuoteAssetMaxBorrowableOk returns a tuple with the QuoteAssetMaxBorrowable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginTierDataResponseInner) GetQuoteAssetMaxBorrowableOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteAssetMaxBorrowable) {
		return nil, false
	}
	return o.QuoteAssetMaxBorrowable, true
}

// HasQuoteAssetMaxBorrowable returns a boolean if a field has been set.
func (o *QueryIsolatedMarginTierDataResponseInner) HasQuoteAssetMaxBorrowable() bool {
	if o != nil && !common.IsNil(o.QuoteAssetMaxBorrowable) {
		return true
	}

	return false
}

// SetQuoteAssetMaxBorrowable gets a reference to the given string and assigns it to the QuoteAssetMaxBorrowable field.
func (o *QueryIsolatedMarginTierDataResponseInner) SetQuoteAssetMaxBorrowable(v string) {
	o.QuoteAssetMaxBorrowable = &v
}

func (o QueryIsolatedMarginTierDataResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryIsolatedMarginTierDataResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}
	if !common.IsNil(o.EffectiveMultiple) {
		toSerialize["effectiveMultiple"] = o.EffectiveMultiple
	}
	if !common.IsNil(o.InitialRiskRatio) {
		toSerialize["initialRiskRatio"] = o.InitialRiskRatio
	}
	if !common.IsNil(o.LiquidationRiskRatio) {
		toSerialize["liquidationRiskRatio"] = o.LiquidationRiskRatio
	}
	if !common.IsNil(o.BaseAssetMaxBorrowable) {
		toSerialize["baseAssetMaxBorrowable"] = o.BaseAssetMaxBorrowable
	}
	if !common.IsNil(o.QuoteAssetMaxBorrowable) {
		toSerialize["quoteAssetMaxBorrowable"] = o.QuoteAssetMaxBorrowable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryIsolatedMarginTierDataResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryIsolatedMarginTierDataResponseInner := _QueryIsolatedMarginTierDataResponseInner{}

	err = json.Unmarshal(data, &varQueryIsolatedMarginTierDataResponseInner)

	if err != nil {
		return err
	}

	*o = QueryIsolatedMarginTierDataResponseInner(varQueryIsolatedMarginTierDataResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "tier")
		delete(additionalProperties, "effectiveMultiple")
		delete(additionalProperties, "initialRiskRatio")
		delete(additionalProperties, "liquidationRiskRatio")
		delete(additionalProperties, "baseAssetMaxBorrowable")
		delete(additionalProperties, "quoteAssetMaxBorrowable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryIsolatedMarginTierDataResponseInner struct {
	value *QueryIsolatedMarginTierDataResponseInner
	isSet bool
}

func (v NullableQueryIsolatedMarginTierDataResponseInner) Get() *QueryIsolatedMarginTierDataResponseInner {
	return v.value
}

func (v *NullableQueryIsolatedMarginTierDataResponseInner) Set(val *QueryIsolatedMarginTierDataResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryIsolatedMarginTierDataResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryIsolatedMarginTierDataResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryIsolatedMarginTierDataResponseInner(val *QueryIsolatedMarginTierDataResponseInner) *NullableQueryIsolatedMarginTierDataResponseInner {
	return &NullableQueryIsolatedMarginTierDataResponseInner{value: val, isSet: true}
}

func (v NullableQueryIsolatedMarginTierDataResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryIsolatedMarginTierDataResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
