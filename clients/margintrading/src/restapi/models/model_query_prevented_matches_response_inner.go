/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryPreventedMatchesResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPreventedMatchesResponseInner{}

// QueryPreventedMatchesResponseInner struct for QueryPreventedMatchesResponseInner
type QueryPreventedMatchesResponseInner struct {
	Symbol                  *string `json:"symbol,omitempty"`
	PreventedMatchId        *int64  `json:"preventedMatchId,omitempty"`
	TakerOrderId            *int64  `json:"takerOrderId,omitempty"`
	MakerSymbol             *string `json:"makerSymbol,omitempty"`
	MakerOrderId            *int64  `json:"makerOrderId,omitempty"`
	TradeGroupId            *int64  `json:"tradeGroupId,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	Price                   *string `json:"price,omitempty"`
	MakerPreventedQuantity  *string `json:"makerPreventedQuantity,omitempty"`
	TransactTime            *int64  `json:"transactTime,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _QueryPreventedMatchesResponseInner QueryPreventedMatchesResponseInner

// NewQueryPreventedMatchesResponseInner instantiates a new QueryPreventedMatchesResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPreventedMatchesResponseInner() *QueryPreventedMatchesResponseInner {
	this := QueryPreventedMatchesResponseInner{}
	return &this
}

// NewQueryPreventedMatchesResponseInnerWithDefaults instantiates a new QueryPreventedMatchesResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPreventedMatchesResponseInnerWithDefaults() *QueryPreventedMatchesResponseInner {
	this := QueryPreventedMatchesResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryPreventedMatchesResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPreventedMatchesResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryPreventedMatchesResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryPreventedMatchesResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPreventedMatchId returns the PreventedMatchId field value if set, zero value otherwise.
func (o *QueryPreventedMatchesResponseInner) GetPreventedMatchId() int64 {
	if o == nil || common.IsNil(o.PreventedMatchId) {
		var ret int64
		return ret
	}
	return *o.PreventedMatchId
}

// GetPreventedMatchIdOk returns a tuple with the PreventedMatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPreventedMatchesResponseInner) GetPreventedMatchIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PreventedMatchId) {
		return nil, false
	}
	return o.PreventedMatchId, true
}

// HasPreventedMatchId returns a boolean if a field has been set.
func (o *QueryPreventedMatchesResponseInner) HasPreventedMatchId() bool {
	if o != nil && !common.IsNil(o.PreventedMatchId) {
		return true
	}

	return false
}

// SetPreventedMatchId gets a reference to the given int64 and assigns it to the PreventedMatchId field.
func (o *QueryPreventedMatchesResponseInner) SetPreventedMatchId(v int64) {
	o.PreventedMatchId = &v
}

// GetTakerOrderId returns the TakerOrderId field value if set, zero value otherwise.
func (o *QueryPreventedMatchesResponseInner) GetTakerOrderId() int64 {
	if o == nil || common.IsNil(o.TakerOrderId) {
		var ret int64
		return ret
	}
	return *o.TakerOrderId
}

// GetTakerOrderIdOk returns a tuple with the TakerOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPreventedMatchesResponseInner) GetTakerOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TakerOrderId) {
		return nil, false
	}
	return o.TakerOrderId, true
}

// HasTakerOrderId returns a boolean if a field has been set.
func (o *QueryPreventedMatchesResponseInner) HasTakerOrderId() bool {
	if o != nil && !common.IsNil(o.TakerOrderId) {
		return true
	}

	return false
}

// SetTakerOrderId gets a reference to the given int64 and assigns it to the TakerOrderId field.
func (o *QueryPreventedMatchesResponseInner) SetTakerOrderId(v int64) {
	o.TakerOrderId = &v
}

// GetMakerSymbol returns the MakerSymbol field value if set, zero value otherwise.
func (o *QueryPreventedMatchesResponseInner) GetMakerSymbol() string {
	if o == nil || common.IsNil(o.MakerSymbol) {
		var ret string
		return ret
	}
	return *o.MakerSymbol
}

// GetMakerSymbolOk returns a tuple with the MakerSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPreventedMatchesResponseInner) GetMakerSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.MakerSymbol) {
		return nil, false
	}
	return o.MakerSymbol, true
}

// HasMakerSymbol returns a boolean if a field has been set.
func (o *QueryPreventedMatchesResponseInner) HasMakerSymbol() bool {
	if o != nil && !common.IsNil(o.MakerSymbol) {
		return true
	}

	return false
}

// SetMakerSymbol gets a reference to the given string and assigns it to the MakerSymbol field.
func (o *QueryPreventedMatchesResponseInner) SetMakerSymbol(v string) {
	o.MakerSymbol = &v
}

// GetMakerOrderId returns the MakerOrderId field value if set, zero value otherwise.
func (o *QueryPreventedMatchesResponseInner) GetMakerOrderId() int64 {
	if o == nil || common.IsNil(o.MakerOrderId) {
		var ret int64
		return ret
	}
	return *o.MakerOrderId
}

// GetMakerOrderIdOk returns a tuple with the MakerOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPreventedMatchesResponseInner) GetMakerOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MakerOrderId) {
		return nil, false
	}
	return o.MakerOrderId, true
}

// HasMakerOrderId returns a boolean if a field has been set.
func (o *QueryPreventedMatchesResponseInner) HasMakerOrderId() bool {
	if o != nil && !common.IsNil(o.MakerOrderId) {
		return true
	}

	return false
}

// SetMakerOrderId gets a reference to the given int64 and assigns it to the MakerOrderId field.
func (o *QueryPreventedMatchesResponseInner) SetMakerOrderId(v int64) {
	o.MakerOrderId = &v
}

// GetTradeGroupId returns the TradeGroupId field value if set, zero value otherwise.
func (o *QueryPreventedMatchesResponseInner) GetTradeGroupId() int64 {
	if o == nil || common.IsNil(o.TradeGroupId) {
		var ret int64
		return ret
	}
	return *o.TradeGroupId
}

// GetTradeGroupIdOk returns a tuple with the TradeGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPreventedMatchesResponseInner) GetTradeGroupIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeGroupId) {
		return nil, false
	}
	return o.TradeGroupId, true
}

// HasTradeGroupId returns a boolean if a field has been set.
func (o *QueryPreventedMatchesResponseInner) HasTradeGroupId() bool {
	if o != nil && !common.IsNil(o.TradeGroupId) {
		return true
	}

	return false
}

// SetTradeGroupId gets a reference to the given int64 and assigns it to the TradeGroupId field.
func (o *QueryPreventedMatchesResponseInner) SetTradeGroupId(v int64) {
	o.TradeGroupId = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *QueryPreventedMatchesResponseInner) GetSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPreventedMatchesResponseInner) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *QueryPreventedMatchesResponseInner) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *QueryPreventedMatchesResponseInner) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QueryPreventedMatchesResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPreventedMatchesResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QueryPreventedMatchesResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *QueryPreventedMatchesResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetMakerPreventedQuantity returns the MakerPreventedQuantity field value if set, zero value otherwise.
func (o *QueryPreventedMatchesResponseInner) GetMakerPreventedQuantity() string {
	if o == nil || common.IsNil(o.MakerPreventedQuantity) {
		var ret string
		return ret
	}
	return *o.MakerPreventedQuantity
}

// GetMakerPreventedQuantityOk returns a tuple with the MakerPreventedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPreventedMatchesResponseInner) GetMakerPreventedQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.MakerPreventedQuantity) {
		return nil, false
	}
	return o.MakerPreventedQuantity, true
}

// HasMakerPreventedQuantity returns a boolean if a field has been set.
func (o *QueryPreventedMatchesResponseInner) HasMakerPreventedQuantity() bool {
	if o != nil && !common.IsNil(o.MakerPreventedQuantity) {
		return true
	}

	return false
}

// SetMakerPreventedQuantity gets a reference to the given string and assigns it to the MakerPreventedQuantity field.
func (o *QueryPreventedMatchesResponseInner) SetMakerPreventedQuantity(v string) {
	o.MakerPreventedQuantity = &v
}

// GetTransactTime returns the TransactTime field value if set, zero value otherwise.
func (o *QueryPreventedMatchesResponseInner) GetTransactTime() int64 {
	if o == nil || common.IsNil(o.TransactTime) {
		var ret int64
		return ret
	}
	return *o.TransactTime
}

// GetTransactTimeOk returns a tuple with the TransactTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPreventedMatchesResponseInner) GetTransactTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactTime) {
		return nil, false
	}
	return o.TransactTime, true
}

// HasTransactTime returns a boolean if a field has been set.
func (o *QueryPreventedMatchesResponseInner) HasTransactTime() bool {
	if o != nil && !common.IsNil(o.TransactTime) {
		return true
	}

	return false
}

// SetTransactTime gets a reference to the given int64 and assigns it to the TransactTime field.
func (o *QueryPreventedMatchesResponseInner) SetTransactTime(v int64) {
	o.TransactTime = &v
}

func (o QueryPreventedMatchesResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPreventedMatchesResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.PreventedMatchId) {
		toSerialize["preventedMatchId"] = o.PreventedMatchId
	}
	if !common.IsNil(o.TakerOrderId) {
		toSerialize["takerOrderId"] = o.TakerOrderId
	}
	if !common.IsNil(o.MakerSymbol) {
		toSerialize["makerSymbol"] = o.MakerSymbol
	}
	if !common.IsNil(o.MakerOrderId) {
		toSerialize["makerOrderId"] = o.MakerOrderId
	}
	if !common.IsNil(o.TradeGroupId) {
		toSerialize["tradeGroupId"] = o.TradeGroupId
	}
	if !common.IsNil(o.SelfTradePreventionMode) {
		toSerialize["selfTradePreventionMode"] = o.SelfTradePreventionMode
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.MakerPreventedQuantity) {
		toSerialize["makerPreventedQuantity"] = o.MakerPreventedQuantity
	}
	if !common.IsNil(o.TransactTime) {
		toSerialize["transactTime"] = o.TransactTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryPreventedMatchesResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryPreventedMatchesResponseInner := _QueryPreventedMatchesResponseInner{}

	err = json.Unmarshal(data, &varQueryPreventedMatchesResponseInner)

	if err != nil {
		return err
	}

	*o = QueryPreventedMatchesResponseInner(varQueryPreventedMatchesResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "preventedMatchId")
		delete(additionalProperties, "takerOrderId")
		delete(additionalProperties, "makerSymbol")
		delete(additionalProperties, "makerOrderId")
		delete(additionalProperties, "tradeGroupId")
		delete(additionalProperties, "selfTradePreventionMode")
		delete(additionalProperties, "price")
		delete(additionalProperties, "makerPreventedQuantity")
		delete(additionalProperties, "transactTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryPreventedMatchesResponseInner struct {
	value *QueryPreventedMatchesResponseInner
	isSet bool
}

func (v NullableQueryPreventedMatchesResponseInner) Get() *QueryPreventedMatchesResponseInner {
	return v.value
}

func (v *NullableQueryPreventedMatchesResponseInner) Set(val *QueryPreventedMatchesResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPreventedMatchesResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPreventedMatchesResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPreventedMatchesResponseInner(val *QueryPreventedMatchesResponseInner) *NullableQueryPreventedMatchesResponseInner {
	return &NullableQueryPreventedMatchesResponseInner{value: val, isSet: true}
}

func (v NullableQueryPreventedMatchesResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPreventedMatchesResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
