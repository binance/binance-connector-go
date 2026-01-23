/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API

*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetExchangeInfoResponseDataSymbolsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetExchangeInfoResponseDataSymbolsInner{}

// GetExchangeInfoResponseDataSymbolsInner struct for GetExchangeInfoResponseDataSymbolsInner
type GetExchangeInfoResponseDataSymbolsInner struct {
	Symbol               *string                                               `json:"symbol,omitempty"`
	Status               *string                                               `json:"status,omitempty"`
	BaseAsset            *string                                               `json:"baseAsset,omitempty"`
	QuoteAsset           *string                                               `json:"quoteAsset,omitempty"`
	PricePrecision       *int64                                                `json:"pricePrecision,omitempty"`
	QuantityPrecision    *int64                                                `json:"quantityPrecision,omitempty"`
	BaseAssetPrecision   *int64                                                `json:"baseAssetPrecision,omitempty"`
	QuotePrecision       *int64                                                `json:"quotePrecision,omitempty"`
	Filters              []GetExchangeInfoResponseDataSymbolsInnerFiltersInner `json:"filters,omitempty"`
	OrderTypes           []string                                              `json:"orderTypes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetExchangeInfoResponseDataSymbolsInner GetExchangeInfoResponseDataSymbolsInner

// NewGetExchangeInfoResponseDataSymbolsInner instantiates a new GetExchangeInfoResponseDataSymbolsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExchangeInfoResponseDataSymbolsInner() *GetExchangeInfoResponseDataSymbolsInner {
	this := GetExchangeInfoResponseDataSymbolsInner{}
	return &this
}

// NewGetExchangeInfoResponseDataSymbolsInnerWithDefaults instantiates a new GetExchangeInfoResponseDataSymbolsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExchangeInfoResponseDataSymbolsInnerWithDefaults() *GetExchangeInfoResponseDataSymbolsInner {
	this := GetExchangeInfoResponseDataSymbolsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetExchangeInfoResponseDataSymbolsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetExchangeInfoResponseDataSymbolsInner) SetStatus(v string) {
	o.Status = &v
}

// GetBaseAsset returns the BaseAsset field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetBaseAsset() string {
	if o == nil || common.IsNil(o.BaseAsset) {
		var ret string
		return ret
	}
	return *o.BaseAsset
}

// GetBaseAssetOk returns a tuple with the BaseAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetBaseAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.BaseAsset) {
		return nil, false
	}
	return o.BaseAsset, true
}

// HasBaseAsset returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) HasBaseAsset() bool {
	if o != nil && !common.IsNil(o.BaseAsset) {
		return true
	}

	return false
}

// SetBaseAsset gets a reference to the given string and assigns it to the BaseAsset field.
func (o *GetExchangeInfoResponseDataSymbolsInner) SetBaseAsset(v string) {
	o.BaseAsset = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetQuoteAsset() string {
	if o == nil || common.IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetQuoteAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) HasQuoteAsset() bool {
	if o != nil && !common.IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *GetExchangeInfoResponseDataSymbolsInner) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

// GetPricePrecision returns the PricePrecision field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetPricePrecision() int64 {
	if o == nil || common.IsNil(o.PricePrecision) {
		var ret int64
		return ret
	}
	return *o.PricePrecision
}

// GetPricePrecisionOk returns a tuple with the PricePrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetPricePrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PricePrecision) {
		return nil, false
	}
	return o.PricePrecision, true
}

// HasPricePrecision returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) HasPricePrecision() bool {
	if o != nil && !common.IsNil(o.PricePrecision) {
		return true
	}

	return false
}

// SetPricePrecision gets a reference to the given int64 and assigns it to the PricePrecision field.
func (o *GetExchangeInfoResponseDataSymbolsInner) SetPricePrecision(v int64) {
	o.PricePrecision = &v
}

// GetQuantityPrecision returns the QuantityPrecision field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetQuantityPrecision() int64 {
	if o == nil || common.IsNil(o.QuantityPrecision) {
		var ret int64
		return ret
	}
	return *o.QuantityPrecision
}

// GetQuantityPrecisionOk returns a tuple with the QuantityPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetQuantityPrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QuantityPrecision) {
		return nil, false
	}
	return o.QuantityPrecision, true
}

// HasQuantityPrecision returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) HasQuantityPrecision() bool {
	if o != nil && !common.IsNil(o.QuantityPrecision) {
		return true
	}

	return false
}

// SetQuantityPrecision gets a reference to the given int64 and assigns it to the QuantityPrecision field.
func (o *GetExchangeInfoResponseDataSymbolsInner) SetQuantityPrecision(v int64) {
	o.QuantityPrecision = &v
}

// GetBaseAssetPrecision returns the BaseAssetPrecision field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetBaseAssetPrecision() int64 {
	if o == nil || common.IsNil(o.BaseAssetPrecision) {
		var ret int64
		return ret
	}
	return *o.BaseAssetPrecision
}

// GetBaseAssetPrecisionOk returns a tuple with the BaseAssetPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetBaseAssetPrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BaseAssetPrecision) {
		return nil, false
	}
	return o.BaseAssetPrecision, true
}

// HasBaseAssetPrecision returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) HasBaseAssetPrecision() bool {
	if o != nil && !common.IsNil(o.BaseAssetPrecision) {
		return true
	}

	return false
}

// SetBaseAssetPrecision gets a reference to the given int64 and assigns it to the BaseAssetPrecision field.
func (o *GetExchangeInfoResponseDataSymbolsInner) SetBaseAssetPrecision(v int64) {
	o.BaseAssetPrecision = &v
}

// GetQuotePrecision returns the QuotePrecision field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetQuotePrecision() int64 {
	if o == nil || common.IsNil(o.QuotePrecision) {
		var ret int64
		return ret
	}
	return *o.QuotePrecision
}

// GetQuotePrecisionOk returns a tuple with the QuotePrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetQuotePrecisionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QuotePrecision) {
		return nil, false
	}
	return o.QuotePrecision, true
}

// HasQuotePrecision returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) HasQuotePrecision() bool {
	if o != nil && !common.IsNil(o.QuotePrecision) {
		return true
	}

	return false
}

// SetQuotePrecision gets a reference to the given int64 and assigns it to the QuotePrecision field.
func (o *GetExchangeInfoResponseDataSymbolsInner) SetQuotePrecision(v int64) {
	o.QuotePrecision = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetFilters() []GetExchangeInfoResponseDataSymbolsInnerFiltersInner {
	if o == nil || common.IsNil(o.Filters) {
		var ret []GetExchangeInfoResponseDataSymbolsInnerFiltersInner
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetFiltersOk() ([]GetExchangeInfoResponseDataSymbolsInnerFiltersInner, bool) {
	if o == nil || common.IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) HasFilters() bool {
	if o != nil && !common.IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []GetExchangeInfoResponseDataSymbolsInnerFiltersInner and assigns it to the Filters field.
func (o *GetExchangeInfoResponseDataSymbolsInner) SetFilters(v []GetExchangeInfoResponseDataSymbolsInnerFiltersInner) {
	o.Filters = v
}

// GetOrderTypes returns the OrderTypes field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetOrderTypes() []string {
	if o == nil || common.IsNil(o.OrderTypes) {
		var ret []string
		return ret
	}
	return o.OrderTypes
}

// GetOrderTypesOk returns a tuple with the OrderTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) GetOrderTypesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.OrderTypes) {
		return nil, false
	}
	return o.OrderTypes, true
}

// HasOrderTypes returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInner) HasOrderTypes() bool {
	if o != nil && !common.IsNil(o.OrderTypes) {
		return true
	}

	return false
}

// SetOrderTypes gets a reference to the given []string and assigns it to the OrderTypes field.
func (o *GetExchangeInfoResponseDataSymbolsInner) SetOrderTypes(v []string) {
	o.OrderTypes = v
}

func (o GetExchangeInfoResponseDataSymbolsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExchangeInfoResponseDataSymbolsInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.QuoteAsset) {
		toSerialize["quoteAsset"] = o.QuoteAsset
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
	if !common.IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !common.IsNil(o.OrderTypes) {
		toSerialize["orderTypes"] = o.OrderTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetExchangeInfoResponseDataSymbolsInner) UnmarshalJSON(data []byte) (err error) {
	varGetExchangeInfoResponseDataSymbolsInner := _GetExchangeInfoResponseDataSymbolsInner{}

	err = json.Unmarshal(data, &varGetExchangeInfoResponseDataSymbolsInner)

	if err != nil {
		return err
	}

	*o = GetExchangeInfoResponseDataSymbolsInner(varGetExchangeInfoResponseDataSymbolsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "status")
		delete(additionalProperties, "baseAsset")
		delete(additionalProperties, "quoteAsset")
		delete(additionalProperties, "pricePrecision")
		delete(additionalProperties, "quantityPrecision")
		delete(additionalProperties, "baseAssetPrecision")
		delete(additionalProperties, "quotePrecision")
		delete(additionalProperties, "filters")
		delete(additionalProperties, "orderTypes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetExchangeInfoResponseDataSymbolsInner struct {
	value *GetExchangeInfoResponseDataSymbolsInner
	isSet bool
}

func (v NullableGetExchangeInfoResponseDataSymbolsInner) Get() *GetExchangeInfoResponseDataSymbolsInner {
	return v.value
}

func (v *NullableGetExchangeInfoResponseDataSymbolsInner) Set(val *GetExchangeInfoResponseDataSymbolsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExchangeInfoResponseDataSymbolsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExchangeInfoResponseDataSymbolsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExchangeInfoResponseDataSymbolsInner(val *GetExchangeInfoResponseDataSymbolsInner) *NullableGetExchangeInfoResponseDataSymbolsInner {
	return &NullableGetExchangeInfoResponseDataSymbolsInner{value: val, isSet: true}
}

func (v NullableGetExchangeInfoResponseDataSymbolsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExchangeInfoResponseDataSymbolsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
