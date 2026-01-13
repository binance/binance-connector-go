/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MyPreventedMatchesResponseResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MyPreventedMatchesResponseResultInner{}

// MyPreventedMatchesResponseResultInner struct for MyPreventedMatchesResponseResultInner
type MyPreventedMatchesResponseResultInner struct {
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

type _MyPreventedMatchesResponseResultInner MyPreventedMatchesResponseResultInner

// NewMyPreventedMatchesResponseResultInner instantiates a new MyPreventedMatchesResponseResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyPreventedMatchesResponseResultInner() *MyPreventedMatchesResponseResultInner {
	this := MyPreventedMatchesResponseResultInner{}
	return &this
}

// NewMyPreventedMatchesResponseResultInnerWithDefaults instantiates a new MyPreventedMatchesResponseResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyPreventedMatchesResponseResultInnerWithDefaults() *MyPreventedMatchesResponseResultInner {
	this := MyPreventedMatchesResponseResultInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseResultInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseResultInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseResultInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MyPreventedMatchesResponseResultInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPreventedMatchId returns the PreventedMatchId field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseResultInner) GetPreventedMatchId() int64 {
	if o == nil || common.IsNil(o.PreventedMatchId) {
		var ret int64
		return ret
	}
	return *o.PreventedMatchId
}

// GetPreventedMatchIdOk returns a tuple with the PreventedMatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseResultInner) GetPreventedMatchIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PreventedMatchId) {
		return nil, false
	}
	return o.PreventedMatchId, true
}

// HasPreventedMatchId returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseResultInner) HasPreventedMatchId() bool {
	if o != nil && !common.IsNil(o.PreventedMatchId) {
		return true
	}

	return false
}

// SetPreventedMatchId gets a reference to the given int64 and assigns it to the PreventedMatchId field.
func (o *MyPreventedMatchesResponseResultInner) SetPreventedMatchId(v int64) {
	o.PreventedMatchId = &v
}

// GetTakerOrderId returns the TakerOrderId field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseResultInner) GetTakerOrderId() int64 {
	if o == nil || common.IsNil(o.TakerOrderId) {
		var ret int64
		return ret
	}
	return *o.TakerOrderId
}

// GetTakerOrderIdOk returns a tuple with the TakerOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseResultInner) GetTakerOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TakerOrderId) {
		return nil, false
	}
	return o.TakerOrderId, true
}

// HasTakerOrderId returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseResultInner) HasTakerOrderId() bool {
	if o != nil && !common.IsNil(o.TakerOrderId) {
		return true
	}

	return false
}

// SetTakerOrderId gets a reference to the given int64 and assigns it to the TakerOrderId field.
func (o *MyPreventedMatchesResponseResultInner) SetTakerOrderId(v int64) {
	o.TakerOrderId = &v
}

// GetMakerSymbol returns the MakerSymbol field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseResultInner) GetMakerSymbol() string {
	if o == nil || common.IsNil(o.MakerSymbol) {
		var ret string
		return ret
	}
	return *o.MakerSymbol
}

// GetMakerSymbolOk returns a tuple with the MakerSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseResultInner) GetMakerSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.MakerSymbol) {
		return nil, false
	}
	return o.MakerSymbol, true
}

// HasMakerSymbol returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseResultInner) HasMakerSymbol() bool {
	if o != nil && !common.IsNil(o.MakerSymbol) {
		return true
	}

	return false
}

// SetMakerSymbol gets a reference to the given string and assigns it to the MakerSymbol field.
func (o *MyPreventedMatchesResponseResultInner) SetMakerSymbol(v string) {
	o.MakerSymbol = &v
}

// GetMakerOrderId returns the MakerOrderId field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseResultInner) GetMakerOrderId() int64 {
	if o == nil || common.IsNil(o.MakerOrderId) {
		var ret int64
		return ret
	}
	return *o.MakerOrderId
}

// GetMakerOrderIdOk returns a tuple with the MakerOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseResultInner) GetMakerOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MakerOrderId) {
		return nil, false
	}
	return o.MakerOrderId, true
}

// HasMakerOrderId returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseResultInner) HasMakerOrderId() bool {
	if o != nil && !common.IsNil(o.MakerOrderId) {
		return true
	}

	return false
}

// SetMakerOrderId gets a reference to the given int64 and assigns it to the MakerOrderId field.
func (o *MyPreventedMatchesResponseResultInner) SetMakerOrderId(v int64) {
	o.MakerOrderId = &v
}

// GetTradeGroupId returns the TradeGroupId field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseResultInner) GetTradeGroupId() int64 {
	if o == nil || common.IsNil(o.TradeGroupId) {
		var ret int64
		return ret
	}
	return *o.TradeGroupId
}

// GetTradeGroupIdOk returns a tuple with the TradeGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseResultInner) GetTradeGroupIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeGroupId) {
		return nil, false
	}
	return o.TradeGroupId, true
}

// HasTradeGroupId returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseResultInner) HasTradeGroupId() bool {
	if o != nil && !common.IsNil(o.TradeGroupId) {
		return true
	}

	return false
}

// SetTradeGroupId gets a reference to the given int64 and assigns it to the TradeGroupId field.
func (o *MyPreventedMatchesResponseResultInner) SetTradeGroupId(v int64) {
	o.TradeGroupId = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseResultInner) GetSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseResultInner) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseResultInner) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *MyPreventedMatchesResponseResultInner) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseResultInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseResultInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseResultInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *MyPreventedMatchesResponseResultInner) SetPrice(v string) {
	o.Price = &v
}

// GetMakerPreventedQuantity returns the MakerPreventedQuantity field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseResultInner) GetMakerPreventedQuantity() string {
	if o == nil || common.IsNil(o.MakerPreventedQuantity) {
		var ret string
		return ret
	}
	return *o.MakerPreventedQuantity
}

// GetMakerPreventedQuantityOk returns a tuple with the MakerPreventedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseResultInner) GetMakerPreventedQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.MakerPreventedQuantity) {
		return nil, false
	}
	return o.MakerPreventedQuantity, true
}

// HasMakerPreventedQuantity returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseResultInner) HasMakerPreventedQuantity() bool {
	if o != nil && !common.IsNil(o.MakerPreventedQuantity) {
		return true
	}

	return false
}

// SetMakerPreventedQuantity gets a reference to the given string and assigns it to the MakerPreventedQuantity field.
func (o *MyPreventedMatchesResponseResultInner) SetMakerPreventedQuantity(v string) {
	o.MakerPreventedQuantity = &v
}

// GetTransactTime returns the TransactTime field value if set, zero value otherwise.
func (o *MyPreventedMatchesResponseResultInner) GetTransactTime() int64 {
	if o == nil || common.IsNil(o.TransactTime) {
		var ret int64
		return ret
	}
	return *o.TransactTime
}

// GetTransactTimeOk returns a tuple with the TransactTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyPreventedMatchesResponseResultInner) GetTransactTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactTime) {
		return nil, false
	}
	return o.TransactTime, true
}

// HasTransactTime returns a boolean if a field has been set.
func (o *MyPreventedMatchesResponseResultInner) HasTransactTime() bool {
	if o != nil && !common.IsNil(o.TransactTime) {
		return true
	}

	return false
}

// SetTransactTime gets a reference to the given int64 and assigns it to the TransactTime field.
func (o *MyPreventedMatchesResponseResultInner) SetTransactTime(v int64) {
	o.TransactTime = &v
}

func (o MyPreventedMatchesResponseResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyPreventedMatchesResponseResultInner) ToMap() (map[string]interface{}, error) {
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

func (o *MyPreventedMatchesResponseResultInner) UnmarshalJSON(data []byte) (err error) {
	varMyPreventedMatchesResponseResultInner := _MyPreventedMatchesResponseResultInner{}

	err = json.Unmarshal(data, &varMyPreventedMatchesResponseResultInner)

	if err != nil {
		return err
	}

	*o = MyPreventedMatchesResponseResultInner(varMyPreventedMatchesResponseResultInner)

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

type NullableMyPreventedMatchesResponseResultInner struct {
	value *MyPreventedMatchesResponseResultInner
	isSet bool
}

func (v NullableMyPreventedMatchesResponseResultInner) Get() *MyPreventedMatchesResponseResultInner {
	return v.value
}

func (v *NullableMyPreventedMatchesResponseResultInner) Set(val *MyPreventedMatchesResponseResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMyPreventedMatchesResponseResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMyPreventedMatchesResponseResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyPreventedMatchesResponseResultInner(val *MyPreventedMatchesResponseResultInner) *NullableMyPreventedMatchesResponseResultInner {
	return &NullableMyPreventedMatchesResponseResultInner{value: val, isSet: true}
}

func (v NullableMyPreventedMatchesResponseResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyPreventedMatchesResponseResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
