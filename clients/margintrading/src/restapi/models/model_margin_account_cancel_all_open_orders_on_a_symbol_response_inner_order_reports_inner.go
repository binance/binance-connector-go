/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner{}

// MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner struct for MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner
type MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	OrigClientOrderId    *string `json:"origClientOrderId,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	OrderListId          *int64  `json:"orderListId,omitempty"`
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	Price                *string `json:"price,omitempty"`
	OrigQty              *string `json:"origQty,omitempty"`
	ExecutedQty          *string `json:"executedQty,omitempty"`
	CummulativeQuoteQty  *string `json:"cummulativeQuoteQty,omitempty"`
	Status               *string `json:"status,omitempty"`
	TimeInForce          *string `json:"timeInForce,omitempty"`
	Type                 *string `json:"type,omitempty"`
	Side                 *string `json:"side,omitempty"`
	StopPrice            *string `json:"stopPrice,omitempty"`
	IcebergQty           *string `json:"icebergQty,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner

// NewMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner instantiates a new MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner() *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner {
	this := MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner{}
	return &this
}

// NewMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInnerWithDefaults instantiates a new MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInnerWithDefaults() *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner {
	this := MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrigClientOrderId returns the OrigClientOrderId field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetOrigClientOrderId() string {
	if o == nil || common.IsNil(o.OrigClientOrderId) {
		var ret string
		return ret
	}
	return *o.OrigClientOrderId
}

// GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetOrigClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigClientOrderId) {
		return nil, false
	}
	return o.OrigClientOrderId, true
}

// HasOrigClientOrderId returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) HasOrigClientOrderId() bool {
	if o != nil && !common.IsNil(o.OrigClientOrderId) {
		return true
	}

	return false
}

// SetOrigClientOrderId gets a reference to the given string and assigns it to the OrigClientOrderId field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) SetOrigClientOrderId(v string) {
	o.OrigClientOrderId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) SetPrice(v string) {
	o.Price = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetOrigQty() string {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetOrigQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetCummulativeQuoteQty returns the CummulativeQuoteQty field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetCummulativeQuoteQty() string {
	if o == nil || common.IsNil(o.CummulativeQuoteQty) {
		var ret string
		return ret
	}
	return *o.CummulativeQuoteQty
}

// GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetCummulativeQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.CummulativeQuoteQty) {
		return nil, false
	}
	return o.CummulativeQuoteQty, true
}

// HasCummulativeQuoteQty returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) HasCummulativeQuoteQty() bool {
	if o != nil && !common.IsNil(o.CummulativeQuoteQty) {
		return true
	}

	return false
}

// SetCummulativeQuoteQty gets a reference to the given string and assigns it to the CummulativeQuoteQty field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) SetCummulativeQuoteQty(v string) {
	o.CummulativeQuoteQty = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) SetStatus(v string) {
	o.Status = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) SetType(v string) {
	o.Type = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) SetSide(v string) {
	o.Side = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetStopPrice() string {
	if o == nil || common.IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetStopPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) HasStopPrice() bool {
	if o != nil && !common.IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetIcebergQty returns the IcebergQty field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetIcebergQty() string {
	if o == nil || common.IsNil(o.IcebergQty) {
		var ret string
		return ret
	}
	return *o.IcebergQty
}

// GetIcebergQtyOk returns a tuple with the IcebergQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) GetIcebergQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.IcebergQty) {
		return nil, false
	}
	return o.IcebergQty, true
}

// HasIcebergQty returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) HasIcebergQty() bool {
	if o != nil && !common.IsNil(o.IcebergQty) {
		return true
	}

	return false
}

// SetIcebergQty gets a reference to the given string and assigns it to the IcebergQty field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) SetIcebergQty(v string) {
	o.IcebergQty = &v
}

func (o MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.OrigClientOrderId) {
		toSerialize["origClientOrderId"] = o.OrigClientOrderId
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.OrderListId) {
		toSerialize["orderListId"] = o.OrderListId
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !common.IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !common.IsNil(o.CummulativeQuoteQty) {
		toSerialize["cummulativeQuoteQty"] = o.CummulativeQuoteQty
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.StopPrice) {
		toSerialize["stopPrice"] = o.StopPrice
	}
	if !common.IsNil(o.IcebergQty) {
		toSerialize["icebergQty"] = o.IcebergQty
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) UnmarshalJSON(data []byte) (err error) {
	varMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner := _MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner{}

	err = json.Unmarshal(data, &varMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner)

	if err != nil {
		return err
	}

	*o = MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner(varMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "origClientOrderId")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "orderListId")
		delete(additionalProperties, "clientOrderId")
		delete(additionalProperties, "price")
		delete(additionalProperties, "origQty")
		delete(additionalProperties, "executedQty")
		delete(additionalProperties, "cummulativeQuoteQty")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "type")
		delete(additionalProperties, "side")
		delete(additionalProperties, "stopPrice")
		delete(additionalProperties, "icebergQty")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner struct {
	value *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner
	isSet bool
}

func (v NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) Get() *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner {
	return v.value
}

func (v *NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) Set(val *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner(val *MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) *NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner {
	return &NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner{value: val, isSet: true}
}

func (v NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
