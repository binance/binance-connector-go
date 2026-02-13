/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MyPreventedMatchesResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MyPreventedMatchesResponseInner{}

// MyPreventedMatchesResponseInner struct for MyPreventedMatchesResponseInner
type MyPreventedMatchesResponseInner struct {
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

type _MyPreventedMatchesResponseInner MyPreventedMatchesResponseInner

// NewMyPreventedMatchesResponseInner instantiates a new MyPreventedMatchesResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyPreventedMatchesResponseInner() *MyPreventedMatchesResponseInner {
	this := MyPreventedMatchesResponseInner{}
	return &this
}

// NewMyPreventedMatchesResponseInnerWithDefaults instantiates a new MyPreventedMatchesResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyPreventedMatchesResponseInnerWithDefaults() *MyPreventedMatchesResponseInner {
	this := MyPreventedMatchesResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MyPreventedMatchesResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPreventedMatchId returns the PreventedMatchId field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseInner) GetPreventedMatchId() int64 {
	if o == nil || common.IsNil(o.PreventedMatchId) {
		var ret int64
		return ret
	}
	return *o.PreventedMatchId
}

// GetPreventedMatchIdOk returns a tuple with the PreventedMatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseInner) GetPreventedMatchIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PreventedMatchId) {
		return nil, false
	}
	return o.PreventedMatchId, true
}

// HasPreventedMatchId returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseInner) HasPreventedMatchId() bool {
	if o != nil && !common.IsNil(o.PreventedMatchId) {
		return true
	}

	return false
}

// SetPreventedMatchId gets a reference to the given int64 and assigns it to the PreventedMatchId field.
func (o *MyPreventedMatchesResponseInner) SetPreventedMatchId(v int64) {
	o.PreventedMatchId = &v
}

// GetTakerOrderId returns the TakerOrderId field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseInner) GetTakerOrderId() int64 {
	if o == nil || common.IsNil(o.TakerOrderId) {
		var ret int64
		return ret
	}
	return *o.TakerOrderId
}

// GetTakerOrderIdOk returns a tuple with the TakerOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseInner) GetTakerOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TakerOrderId) {
		return nil, false
	}
	return o.TakerOrderId, true
}

// HasTakerOrderId returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseInner) HasTakerOrderId() bool {
	if o != nil && !common.IsNil(o.TakerOrderId) {
		return true
	}

	return false
}

// SetTakerOrderId gets a reference to the given int64 and assigns it to the TakerOrderId field.
func (o *MyPreventedMatchesResponseInner) SetTakerOrderId(v int64) {
	o.TakerOrderId = &v
}

// GetMakerSymbol returns the MakerSymbol field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseInner) GetMakerSymbol() string {
	if o == nil || common.IsNil(o.MakerSymbol) {
		var ret string
		return ret
	}
	return *o.MakerSymbol
}

// GetMakerSymbolOk returns a tuple with the MakerSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseInner) GetMakerSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.MakerSymbol) {
		return nil, false
	}
	return o.MakerSymbol, true
}

// HasMakerSymbol returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseInner) HasMakerSymbol() bool {
	if o != nil && !common.IsNil(o.MakerSymbol) {
		return true
	}

	return false
}

// SetMakerSymbol gets a reference to the given string and assigns it to the MakerSymbol field.
func (o *MyPreventedMatchesResponseInner) SetMakerSymbol(v string) {
	o.MakerSymbol = &v
}

// GetMakerOrderId returns the MakerOrderId field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseInner) GetMakerOrderId() int64 {
	if o == nil || common.IsNil(o.MakerOrderId) {
		var ret int64
		return ret
	}
	return *o.MakerOrderId
}

// GetMakerOrderIdOk returns a tuple with the MakerOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseInner) GetMakerOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MakerOrderId) {
		return nil, false
	}
	return o.MakerOrderId, true
}

// HasMakerOrderId returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseInner) HasMakerOrderId() bool {
	if o != nil && !common.IsNil(o.MakerOrderId) {
		return true
	}

	return false
}

// SetMakerOrderId gets a reference to the given int64 and assigns it to the MakerOrderId field.
func (o *MyPreventedMatchesResponseInner) SetMakerOrderId(v int64) {
	o.MakerOrderId = &v
}

// GetTradeGroupId returns the TradeGroupId field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseInner) GetTradeGroupId() int64 {
	if o == nil || common.IsNil(o.TradeGroupId) {
		var ret int64
		return ret
	}
	return *o.TradeGroupId
}

// GetTradeGroupIdOk returns a tuple with the TradeGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseInner) GetTradeGroupIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeGroupId) {
		return nil, false
	}
	return o.TradeGroupId, true
}

// HasTradeGroupId returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseInner) HasTradeGroupId() bool {
	if o != nil && !common.IsNil(o.TradeGroupId) {
		return true
	}

	return false
}

// SetTradeGroupId gets a reference to the given int64 and assigns it to the TradeGroupId field.
func (o *MyPreventedMatchesResponseInner) SetTradeGroupId(v int64) {
	o.TradeGroupId = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseInner) GetSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseInner) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseInner) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *MyPreventedMatchesResponseInner) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *MyPreventedMatchesResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetMakerPreventedQuantity returns the MakerPreventedQuantity field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseInner) GetMakerPreventedQuantity() string {
	if o == nil || common.IsNil(o.MakerPreventedQuantity) {
		var ret string
		return ret
	}
	return *o.MakerPreventedQuantity
}

// GetMakerPreventedQuantityOk returns a tuple with the MakerPreventedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseInner) GetMakerPreventedQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.MakerPreventedQuantity) {
		return nil, false
	}
	return o.MakerPreventedQuantity, true
}

// HasMakerPreventedQuantity returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseInner) HasMakerPreventedQuantity() bool {
	if o != nil && !common.IsNil(o.MakerPreventedQuantity) {
		return true
	}

	return false
}

// SetMakerPreventedQuantity gets a reference to the given string and assigns it to the MakerPreventedQuantity field.
func (o *MyPreventedMatchesResponseInner) SetMakerPreventedQuantity(v string) {
	o.MakerPreventedQuantity = &v
}

// GetTransactTime returns the TransactTime field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseInner) GetTransactTime() int64 {
	if o == nil || common.IsNil(o.TransactTime) {
		var ret int64
		return ret
	}
	return *o.TransactTime
}

// GetTransactTimeOk returns a tuple with the TransactTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseInner) GetTransactTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactTime) {
		return nil, false
	}
	return o.TransactTime, true
}

// HasTransactTime returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseInner) HasTransactTime() bool {
	if o != nil && !common.IsNil(o.TransactTime) {
		return true
	}

	return false
}

// SetTransactTime gets a reference to the given int64 and assigns it to the TransactTime field.
func (o *MyPreventedMatchesResponseInner) SetTransactTime(v int64) {
	o.TransactTime = &v
}

func (o MyPreventedMatchesResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyPreventedMatchesResponseInner) ToMap() (map[string]interface{}, error) {
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

func (o *MyPreventedMatchesResponseInner) UnmarshalJSON(data []byte) (err error) {
	varMyPreventedMatchesResponseInner := _MyPreventedMatchesResponseInner{}

	err = json.Unmarshal(data, &varMyPreventedMatchesResponseInner)

	if err != nil {
		return err
	}

	*o = MyPreventedMatchesResponseInner(varMyPreventedMatchesResponseInner)

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

type NullableMyPreventedMatchesResponseInner struct {
	value *MyPreventedMatchesResponseInner
	isSet bool
}

func (v NullableMyPreventedMatchesResponseInner) Get() *MyPreventedMatchesResponseInner {
	return v.value
}

func (v *NullableMyPreventedMatchesResponseInner) Set(val *MyPreventedMatchesResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMyPreventedMatchesResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMyPreventedMatchesResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyPreventedMatchesResponseInner(val *MyPreventedMatchesResponseInner) *NullableMyPreventedMatchesResponseInner {
	return &NullableMyPreventedMatchesResponseInner{value: val, isSet: true}
}

func (v NullableMyPreventedMatchesResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyPreventedMatchesResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
