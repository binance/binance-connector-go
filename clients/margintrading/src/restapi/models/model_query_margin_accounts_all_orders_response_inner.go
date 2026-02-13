/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryMarginAccountsAllOrdersResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginAccountsAllOrdersResponseInner{}

// QueryMarginAccountsAllOrdersResponseInner struct for QueryMarginAccountsAllOrdersResponseInner
type QueryMarginAccountsAllOrdersResponseInner struct {
	ClientOrderId           *string `json:"clientOrderId,omitempty"`
	CummulativeQuoteQty     *string `json:"cummulativeQuoteQty,omitempty"`
	ExecutedQty             *string `json:"executedQty,omitempty"`
	IcebergQty              *string `json:"icebergQty,omitempty"`
	IsWorking               *bool   `json:"isWorking,omitempty"`
	OrderId                 *int64  `json:"orderId,omitempty"`
	OrigQty                 *string `json:"origQty,omitempty"`
	Price                   *string `json:"price,omitempty"`
	Side                    *string `json:"side,omitempty"`
	Status                  *string `json:"status,omitempty"`
	StopPrice               *string `json:"stopPrice,omitempty"`
	Symbol                  *string `json:"symbol,omitempty"`
	IsIsolated              *bool   `json:"isIsolated,omitempty"`
	Time                    *int64  `json:"time,omitempty"`
	TimeInForce             *string `json:"timeInForce,omitempty"`
	Type                    *string `json:"type,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	UpdateTime              *int64  `json:"updateTime,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _QueryMarginAccountsAllOrdersResponseInner QueryMarginAccountsAllOrdersResponseInner

// NewQueryMarginAccountsAllOrdersResponseInner instantiates a new QueryMarginAccountsAllOrdersResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginAccountsAllOrdersResponseInner() *QueryMarginAccountsAllOrdersResponseInner {
	this := QueryMarginAccountsAllOrdersResponseInner{}
	return &this
}

// NewQueryMarginAccountsAllOrdersResponseInnerWithDefaults instantiates a new QueryMarginAccountsAllOrdersResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginAccountsAllOrdersResponseInnerWithDefaults() *QueryMarginAccountsAllOrdersResponseInner {
	this := QueryMarginAccountsAllOrdersResponseInner{}
	return &this
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetCummulativeQuoteQty returns the CummulativeQuoteQty field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetCummulativeQuoteQty() string {
	if o == nil || common.IsNil(o.CummulativeQuoteQty) {
		var ret string
		return ret
	}
	return *o.CummulativeQuoteQty
}

// GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetCummulativeQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.CummulativeQuoteQty) {
		return nil, false
	}
	return o.CummulativeQuoteQty, true
}

// HasCummulativeQuoteQty returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasCummulativeQuoteQty() bool {
	if o != nil && !common.IsNil(o.CummulativeQuoteQty) {
		return true
	}

	return false
}

// SetCummulativeQuoteQty gets a reference to the given string and assigns it to the CummulativeQuoteQty field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetCummulativeQuoteQty(v string) {
	o.CummulativeQuoteQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetIcebergQty returns the IcebergQty field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetIcebergQty() string {
	if o == nil || common.IsNil(o.IcebergQty) {
		var ret string
		return ret
	}
	return *o.IcebergQty
}

// GetIcebergQtyOk returns a tuple with the IcebergQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetIcebergQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.IcebergQty) {
		return nil, false
	}
	return o.IcebergQty, true
}

// HasIcebergQty returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasIcebergQty() bool {
	if o != nil && !common.IsNil(o.IcebergQty) {
		return true
	}

	return false
}

// SetIcebergQty gets a reference to the given string and assigns it to the IcebergQty field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetIcebergQty(v string) {
	o.IcebergQty = &v
}

// GetIsWorking returns the IsWorking field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetIsWorking() bool {
	if o == nil || common.IsNil(o.IsWorking) {
		var ret bool
		return ret
	}
	return *o.IsWorking
}

// GetIsWorkingOk returns a tuple with the IsWorking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetIsWorkingOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsWorking) {
		return nil, false
	}
	return o.IsWorking, true
}

// HasIsWorking returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasIsWorking() bool {
	if o != nil && !common.IsNil(o.IsWorking) {
		return true
	}

	return false
}

// SetIsWorking gets a reference to the given bool and assigns it to the IsWorking field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetIsWorking(v bool) {
	o.IsWorking = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetOrigQty() string {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetOrigQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetSide(v string) {
	o.Side = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetStopPrice() string {
	if o == nil || common.IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetStopPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasStopPrice() bool {
	if o != nil && !common.IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetIsIsolated returns the IsIsolated field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetIsIsolated() bool {
	if o == nil || common.IsNil(o.IsIsolated) {
		var ret bool
		return ret
	}
	return *o.IsIsolated
}

// GetIsIsolatedOk returns a tuple with the IsIsolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetIsIsolatedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsIsolated) {
		return nil, false
	}
	return o.IsIsolated, true
}

// HasIsIsolated returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasIsIsolated() bool {
	if o != nil && !common.IsNil(o.IsIsolated) {
		return true
	}

	return false
}

// SetIsIsolated gets a reference to the given bool and assigns it to the IsIsolated field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetIsIsolated(v bool) {
	o.IsIsolated = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetType(v string) {
	o.Type = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOrdersResponseInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *QueryMarginAccountsAllOrdersResponseInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o QueryMarginAccountsAllOrdersResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginAccountsAllOrdersResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !common.IsNil(o.CummulativeQuoteQty) {
		toSerialize["cummulativeQuoteQty"] = o.CummulativeQuoteQty
	}
	if !common.IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !common.IsNil(o.IcebergQty) {
		toSerialize["icebergQty"] = o.IcebergQty
	}
	if !common.IsNil(o.IsWorking) {
		toSerialize["isWorking"] = o.IsWorking
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.StopPrice) {
		toSerialize["stopPrice"] = o.StopPrice
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.IsIsolated) {
		toSerialize["isIsolated"] = o.IsIsolated
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.SelfTradePreventionMode) {
		toSerialize["selfTradePreventionMode"] = o.SelfTradePreventionMode
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryMarginAccountsAllOrdersResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryMarginAccountsAllOrdersResponseInner := _QueryMarginAccountsAllOrdersResponseInner{}

	err = json.Unmarshal(data, &varQueryMarginAccountsAllOrdersResponseInner)

	if err != nil {
		return err
	}

	*o = QueryMarginAccountsAllOrdersResponseInner(varQueryMarginAccountsAllOrdersResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "clientOrderId")
		delete(additionalProperties, "cummulativeQuoteQty")
		delete(additionalProperties, "executedQty")
		delete(additionalProperties, "icebergQty")
		delete(additionalProperties, "isWorking")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "origQty")
		delete(additionalProperties, "price")
		delete(additionalProperties, "side")
		delete(additionalProperties, "status")
		delete(additionalProperties, "stopPrice")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "isIsolated")
		delete(additionalProperties, "time")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "type")
		delete(additionalProperties, "selfTradePreventionMode")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryMarginAccountsAllOrdersResponseInner struct {
	value *QueryMarginAccountsAllOrdersResponseInner
	isSet bool
}

func (v NullableQueryMarginAccountsAllOrdersResponseInner) Get() *QueryMarginAccountsAllOrdersResponseInner {
	return v.value
}

func (v *NullableQueryMarginAccountsAllOrdersResponseInner) Set(val *QueryMarginAccountsAllOrdersResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginAccountsAllOrdersResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginAccountsAllOrdersResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginAccountsAllOrdersResponseInner(val *QueryMarginAccountsAllOrdersResponseInner) *NullableQueryMarginAccountsAllOrdersResponseInner {
	return &NullableQueryMarginAccountsAllOrdersResponseInner{value: val, isSet: true}
}

func (v NullableQueryMarginAccountsAllOrdersResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginAccountsAllOrdersResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
