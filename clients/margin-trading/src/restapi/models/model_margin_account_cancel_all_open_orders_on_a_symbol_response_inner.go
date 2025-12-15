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

// checks if the MarginAccountCancelAllOpenOrdersOnASymbolResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginAccountCancelAllOpenOrdersOnASymbolResponseInner{}

// MarginAccountCancelAllOpenOrdersOnASymbolResponseInner struct for MarginAccountCancelAllOpenOrdersOnASymbolResponseInner
type MarginAccountCancelAllOpenOrdersOnASymbolResponseInner struct {
	Symbol                  *string                                                                   `json:"symbol,omitempty"`
	IsIsolated              *bool                                                                     `json:"isIsolated,omitempty"`
	OrigClientOrderId       *string                                                                   `json:"origClientOrderId,omitempty"`
	OrderId                 *int64                                                                    `json:"orderId,omitempty"`
	OrderListId             *int64                                                                    `json:"orderListId,omitempty"`
	ClientOrderId           *string                                                                   `json:"clientOrderId,omitempty"`
	Price                   *string                                                                   `json:"price,omitempty"`
	OrigQty                 *string                                                                   `json:"origQty,omitempty"`
	ExecutedQty             *string                                                                   `json:"executedQty,omitempty"`
	CummulativeQuoteQty     *string                                                                   `json:"cummulativeQuoteQty,omitempty"`
	Status                  *string                                                                   `json:"status,omitempty"`
	TimeInForce             *string                                                                   `json:"timeInForce,omitempty"`
	Type                    *string                                                                   `json:"type,omitempty"`
	Side                    *string                                                                   `json:"side,omitempty"`
	SelfTradePreventionMode *string                                                                   `json:"selfTradePreventionMode,omitempty"`
	ContingencyType         *string                                                                   `json:"contingencyType,omitempty"`
	ListStatusType          *string                                                                   `json:"listStatusType,omitempty"`
	ListOrderStatus         *string                                                                   `json:"listOrderStatus,omitempty"`
	ListClientOrderId       *string                                                                   `json:"listClientOrderId,omitempty"`
	TransactionTime         *int64                                                                    `json:"transactionTime,omitempty"`
	Orders                  []MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrdersInner       `json:"orders,omitempty"`
	OrderReports            []MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner `json:"orderReports,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _MarginAccountCancelAllOpenOrdersOnASymbolResponseInner MarginAccountCancelAllOpenOrdersOnASymbolResponseInner

// NewMarginAccountCancelAllOpenOrdersOnASymbolResponseInner instantiates a new MarginAccountCancelAllOpenOrdersOnASymbolResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginAccountCancelAllOpenOrdersOnASymbolResponseInner() *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner {
	this := MarginAccountCancelAllOpenOrdersOnASymbolResponseInner{}
	return &this
}

// NewMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerWithDefaults instantiates a new MarginAccountCancelAllOpenOrdersOnASymbolResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerWithDefaults() *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner {
	this := MarginAccountCancelAllOpenOrdersOnASymbolResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetIsIsolated returns the IsIsolated field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetIsIsolated() bool {
	if o == nil || common.IsNil(o.IsIsolated) {
		var ret bool
		return ret
	}
	return *o.IsIsolated
}

// GetIsIsolatedOk returns a tuple with the IsIsolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetIsIsolatedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsIsolated) {
		return nil, false
	}
	return o.IsIsolated, true
}

// HasIsIsolated returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasIsIsolated() bool {
	if o != nil && !common.IsNil(o.IsIsolated) {
		return true
	}

	return false
}

// SetIsIsolated gets a reference to the given bool and assigns it to the IsIsolated field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetIsIsolated(v bool) {
	o.IsIsolated = &v
}

// GetOrigClientOrderId returns the OrigClientOrderId field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrigClientOrderId() string {
	if o == nil || common.IsNil(o.OrigClientOrderId) {
		var ret string
		return ret
	}
	return *o.OrigClientOrderId
}

// GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrigClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigClientOrderId) {
		return nil, false
	}
	return o.OrigClientOrderId, true
}

// HasOrigClientOrderId returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasOrigClientOrderId() bool {
	if o != nil && !common.IsNil(o.OrigClientOrderId) {
		return true
	}

	return false
}

// SetOrigClientOrderId gets a reference to the given string and assigns it to the OrigClientOrderId field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetOrigClientOrderId(v string) {
	o.OrigClientOrderId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrigQty() string {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrigQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetCummulativeQuoteQty returns the CummulativeQuoteQty field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetCummulativeQuoteQty() string {
	if o == nil || common.IsNil(o.CummulativeQuoteQty) {
		var ret string
		return ret
	}
	return *o.CummulativeQuoteQty
}

// GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetCummulativeQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.CummulativeQuoteQty) {
		return nil, false
	}
	return o.CummulativeQuoteQty, true
}

// HasCummulativeQuoteQty returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasCummulativeQuoteQty() bool {
	if o != nil && !common.IsNil(o.CummulativeQuoteQty) {
		return true
	}

	return false
}

// SetCummulativeQuoteQty gets a reference to the given string and assigns it to the CummulativeQuoteQty field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetCummulativeQuoteQty(v string) {
	o.CummulativeQuoteQty = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetType(v string) {
	o.Type = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetSide(v string) {
	o.Side = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetContingencyType returns the ContingencyType field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetContingencyType() string {
	if o == nil || common.IsNil(o.ContingencyType) {
		var ret string
		return ret
	}
	return *o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetContingencyTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContingencyType) {
		return nil, false
	}
	return o.ContingencyType, true
}

// HasContingencyType returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasContingencyType() bool {
	if o != nil && !common.IsNil(o.ContingencyType) {
		return true
	}

	return false
}

// SetContingencyType gets a reference to the given string and assigns it to the ContingencyType field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetContingencyType(v string) {
	o.ContingencyType = &v
}

// GetListStatusType returns the ListStatusType field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetListStatusType() string {
	if o == nil || common.IsNil(o.ListStatusType) {
		var ret string
		return ret
	}
	return *o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetListStatusTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListStatusType) {
		return nil, false
	}
	return o.ListStatusType, true
}

// HasListStatusType returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasListStatusType() bool {
	if o != nil && !common.IsNil(o.ListStatusType) {
		return true
	}

	return false
}

// SetListStatusType gets a reference to the given string and assigns it to the ListStatusType field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetListStatusType(v string) {
	o.ListStatusType = &v
}

// GetListOrderStatus returns the ListOrderStatus field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetListOrderStatus() string {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		var ret string
		return ret
	}
	return *o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetListOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		return nil, false
	}
	return o.ListOrderStatus, true
}

// HasListOrderStatus returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasListOrderStatus() bool {
	if o != nil && !common.IsNil(o.ListOrderStatus) {
		return true
	}

	return false
}

// SetListOrderStatus gets a reference to the given string and assigns it to the ListOrderStatus field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetListOrderStatus(v string) {
	o.ListOrderStatus = &v
}

// GetListClientOrderId returns the ListClientOrderId field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetListClientOrderId() string {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		var ret string
		return ret
	}
	return *o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetListClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		return nil, false
	}
	return o.ListClientOrderId, true
}

// HasListClientOrderId returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasListClientOrderId() bool {
	if o != nil && !common.IsNil(o.ListClientOrderId) {
		return true
	}

	return false
}

// SetListClientOrderId gets a reference to the given string and assigns it to the ListClientOrderId field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetListClientOrderId(v string) {
	o.ListClientOrderId = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetTransactionTime() int64 {
	if o == nil || common.IsNil(o.TransactionTime) {
		var ret int64
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetTransactionTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasTransactionTime() bool {
	if o != nil && !common.IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given int64 and assigns it to the TransactionTime field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetTransactionTime(v int64) {
	o.TransactionTime = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrders() []MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrdersOk() ([]MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrdersInner and assigns it to the Orders field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetOrders(v []MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrdersInner) {
	o.Orders = v
}

// GetOrderReports returns the OrderReports field value if set, zero value otherwise.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrderReports() []MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner {
	if o == nil || common.IsNil(o.OrderReports) {
		var ret []MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner
		return ret
	}
	return o.OrderReports
}

// GetOrderReportsOk returns a tuple with the OrderReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrderReportsOk() ([]MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner, bool) {
	if o == nil || common.IsNil(o.OrderReports) {
		return nil, false
	}
	return o.OrderReports, true
}

// HasOrderReports returns a boolean if a field has been set.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasOrderReports() bool {
	if o != nil && !common.IsNil(o.OrderReports) {
		return true
	}

	return false
}

// SetOrderReports gets a reference to the given []MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner and assigns it to the OrderReports field.
func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetOrderReports(v []MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner) {
	o.OrderReports = v
}

func (o MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.IsIsolated) {
		toSerialize["isIsolated"] = o.IsIsolated
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
	if !common.IsNil(o.SelfTradePreventionMode) {
		toSerialize["selfTradePreventionMode"] = o.SelfTradePreventionMode
	}
	if !common.IsNil(o.ContingencyType) {
		toSerialize["contingencyType"] = o.ContingencyType
	}
	if !common.IsNil(o.ListStatusType) {
		toSerialize["listStatusType"] = o.ListStatusType
	}
	if !common.IsNil(o.ListOrderStatus) {
		toSerialize["listOrderStatus"] = o.ListOrderStatus
	}
	if !common.IsNil(o.ListClientOrderId) {
		toSerialize["listClientOrderId"] = o.ListClientOrderId
	}
	if !common.IsNil(o.TransactionTime) {
		toSerialize["transactionTime"] = o.TransactionTime
	}
	if !common.IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	if !common.IsNil(o.OrderReports) {
		toSerialize["orderReports"] = o.OrderReports
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) UnmarshalJSON(data []byte) (err error) {
	varMarginAccountCancelAllOpenOrdersOnASymbolResponseInner := _MarginAccountCancelAllOpenOrdersOnASymbolResponseInner{}

	err = json.Unmarshal(data, &varMarginAccountCancelAllOpenOrdersOnASymbolResponseInner)

	if err != nil {
		return err
	}

	*o = MarginAccountCancelAllOpenOrdersOnASymbolResponseInner(varMarginAccountCancelAllOpenOrdersOnASymbolResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "isIsolated")
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
		delete(additionalProperties, "selfTradePreventionMode")
		delete(additionalProperties, "contingencyType")
		delete(additionalProperties, "listStatusType")
		delete(additionalProperties, "listOrderStatus")
		delete(additionalProperties, "listClientOrderId")
		delete(additionalProperties, "transactionTime")
		delete(additionalProperties, "orders")
		delete(additionalProperties, "orderReports")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInner struct {
	value *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner
	isSet bool
}

func (v NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInner) Get() *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner {
	return v.value
}

func (v *NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInner) Set(val *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInner(val *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) *NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInner {
	return &NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInner{value: val, isSet: true}
}

func (v NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountCancelAllOpenOrdersOnASymbolResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
