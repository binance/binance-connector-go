/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountBlockTradeListResponseInnerLegsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountBlockTradeListResponseInnerLegsInner{}

// AccountBlockTradeListResponseInnerLegsInner struct for AccountBlockTradeListResponseInnerLegsInner
type AccountBlockTradeListResponseInnerLegsInner struct {
	CreateTime           *int64   `json:"createTime,omitempty"`
	UpdateTime           *int64   `json:"updateTime,omitempty"`
	Symbol               *string  `json:"symbol,omitempty"`
	OrderId              *string  `json:"orderId,omitempty"`
	OrderPrice           *float32 `json:"orderPrice,omitempty"`
	OrderQuantity        *float32 `json:"orderQuantity,omitempty"`
	OrderStatus          *string  `json:"orderStatus,omitempty"`
	ExecutedQty          *float32 `json:"executedQty,omitempty"`
	ExecutedAmount       *float32 `json:"executedAmount,omitempty"`
	Fee                  *float32 `json:"fee,omitempty"`
	OrderType            *string  `json:"orderType,omitempty"`
	OrderSide            *string  `json:"orderSide,omitempty"`
	Id                   *string  `json:"id,omitempty"`
	TradeId              *int64   `json:"tradeId,omitempty"`
	TradePrice           *float32 `json:"tradePrice,omitempty"`
	TradeQty             *float32 `json:"tradeQty,omitempty"`
	TradeTime            *int64   `json:"tradeTime,omitempty"`
	Liquidity            *string  `json:"liquidity,omitempty"`
	Commission           *float32 `json:"commission,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountBlockTradeListResponseInnerLegsInner AccountBlockTradeListResponseInnerLegsInner

// NewAccountBlockTradeListResponseInnerLegsInner instantiates a new AccountBlockTradeListResponseInnerLegsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBlockTradeListResponseInnerLegsInner() *AccountBlockTradeListResponseInnerLegsInner {
	this := AccountBlockTradeListResponseInnerLegsInner{}
	return &this
}

// NewAccountBlockTradeListResponseInnerLegsInnerWithDefaults instantiates a new AccountBlockTradeListResponseInnerLegsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBlockTradeListResponseInnerLegsInnerWithDefaults() *AccountBlockTradeListResponseInnerLegsInner {
	this := AccountBlockTradeListResponseInnerLegsInner{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOrderPrice returns the OrderPrice field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetOrderPrice() float32 {
	if o == nil || common.IsNil(o.OrderPrice) {
		var ret float32
		return ret
	}
	return *o.OrderPrice
}

// GetOrderPriceOk returns a tuple with the OrderPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetOrderPriceOk() (*float32, bool) {
	if o == nil || common.IsNil(o.OrderPrice) {
		return nil, false
	}
	return o.OrderPrice, true
}

// HasOrderPrice returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasOrderPrice() bool {
	if o != nil && !common.IsNil(o.OrderPrice) {
		return true
	}

	return false
}

// SetOrderPrice gets a reference to the given float32 and assigns it to the OrderPrice field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetOrderPrice(v float32) {
	o.OrderPrice = &v
}

// GetOrderQuantity returns the OrderQuantity field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetOrderQuantity() float32 {
	if o == nil || common.IsNil(o.OrderQuantity) {
		var ret float32
		return ret
	}
	return *o.OrderQuantity
}

// GetOrderQuantityOk returns a tuple with the OrderQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetOrderQuantityOk() (*float32, bool) {
	if o == nil || common.IsNil(o.OrderQuantity) {
		return nil, false
	}
	return o.OrderQuantity, true
}

// HasOrderQuantity returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasOrderQuantity() bool {
	if o != nil && !common.IsNil(o.OrderQuantity) {
		return true
	}

	return false
}

// SetOrderQuantity gets a reference to the given float32 and assigns it to the OrderQuantity field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetOrderQuantity(v float32) {
	o.OrderQuantity = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetOrderStatus() string {
	if o == nil || common.IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasOrderStatus() bool {
	if o != nil && !common.IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetExecutedQty() float32 {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret float32
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetExecutedQtyOk() (*float32, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given float32 and assigns it to the ExecutedQty field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetExecutedQty(v float32) {
	o.ExecutedQty = &v
}

// GetExecutedAmount returns the ExecutedAmount field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetExecutedAmount() float32 {
	if o == nil || common.IsNil(o.ExecutedAmount) {
		var ret float32
		return ret
	}
	return *o.ExecutedAmount
}

// GetExecutedAmountOk returns a tuple with the ExecutedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetExecutedAmountOk() (*float32, bool) {
	if o == nil || common.IsNil(o.ExecutedAmount) {
		return nil, false
	}
	return o.ExecutedAmount, true
}

// HasExecutedAmount returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasExecutedAmount() bool {
	if o != nil && !common.IsNil(o.ExecutedAmount) {
		return true
	}

	return false
}

// SetExecutedAmount gets a reference to the given float32 and assigns it to the ExecutedAmount field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetExecutedAmount(v float32) {
	o.ExecutedAmount = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetFee() float32 {
	if o == nil || common.IsNil(o.Fee) {
		var ret float32
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetFeeOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given float32 and assigns it to the Fee field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetFee(v float32) {
	o.Fee = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetOrderType() string {
	if o == nil || common.IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetOrderTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasOrderType() bool {
	if o != nil && !common.IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetOrderType(v string) {
	o.OrderType = &v
}

// GetOrderSide returns the OrderSide field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetOrderSide() string {
	if o == nil || common.IsNil(o.OrderSide) {
		var ret string
		return ret
	}
	return *o.OrderSide
}

// GetOrderSideOk returns a tuple with the OrderSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetOrderSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderSide) {
		return nil, false
	}
	return o.OrderSide, true
}

// HasOrderSide returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasOrderSide() bool {
	if o != nil && !common.IsNil(o.OrderSide) {
		return true
	}

	return false
}

// SetOrderSide gets a reference to the given string and assigns it to the OrderSide field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetOrderSide(v string) {
	o.OrderSide = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetId(v string) {
	o.Id = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetTradeId() int64 {
	if o == nil || common.IsNil(o.TradeId) {
		var ret int64
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetTradeIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasTradeId() bool {
	if o != nil && !common.IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given int64 and assigns it to the TradeId field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetTradeId(v int64) {
	o.TradeId = &v
}

// GetTradePrice returns the TradePrice field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetTradePrice() float32 {
	if o == nil || common.IsNil(o.TradePrice) {
		var ret float32
		return ret
	}
	return *o.TradePrice
}

// GetTradePriceOk returns a tuple with the TradePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetTradePriceOk() (*float32, bool) {
	if o == nil || common.IsNil(o.TradePrice) {
		return nil, false
	}
	return o.TradePrice, true
}

// HasTradePrice returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasTradePrice() bool {
	if o != nil && !common.IsNil(o.TradePrice) {
		return true
	}

	return false
}

// SetTradePrice gets a reference to the given float32 and assigns it to the TradePrice field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetTradePrice(v float32) {
	o.TradePrice = &v
}

// GetTradeQty returns the TradeQty field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetTradeQty() float32 {
	if o == nil || common.IsNil(o.TradeQty) {
		var ret float32
		return ret
	}
	return *o.TradeQty
}

// GetTradeQtyOk returns a tuple with the TradeQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetTradeQtyOk() (*float32, bool) {
	if o == nil || common.IsNil(o.TradeQty) {
		return nil, false
	}
	return o.TradeQty, true
}

// HasTradeQty returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasTradeQty() bool {
	if o != nil && !common.IsNil(o.TradeQty) {
		return true
	}

	return false
}

// SetTradeQty gets a reference to the given float32 and assigns it to the TradeQty field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetTradeQty(v float32) {
	o.TradeQty = &v
}

// GetTradeTime returns the TradeTime field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetTradeTime() int64 {
	if o == nil || common.IsNil(o.TradeTime) {
		var ret int64
		return ret
	}
	return *o.TradeTime
}

// GetTradeTimeOk returns a tuple with the TradeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetTradeTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeTime) {
		return nil, false
	}
	return o.TradeTime, true
}

// HasTradeTime returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasTradeTime() bool {
	if o != nil && !common.IsNil(o.TradeTime) {
		return true
	}

	return false
}

// SetTradeTime gets a reference to the given int64 and assigns it to the TradeTime field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetTradeTime(v int64) {
	o.TradeTime = &v
}

// GetLiquidity returns the Liquidity field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetLiquidity() string {
	if o == nil || common.IsNil(o.Liquidity) {
		var ret string
		return ret
	}
	return *o.Liquidity
}

// GetLiquidityOk returns a tuple with the Liquidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetLiquidityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Liquidity) {
		return nil, false
	}
	return o.Liquidity, true
}

// HasLiquidity returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasLiquidity() bool {
	if o != nil && !common.IsNil(o.Liquidity) {
		return true
	}

	return false
}

// SetLiquidity gets a reference to the given string and assigns it to the Liquidity field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetLiquidity(v string) {
	o.Liquidity = &v
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetCommission() float32 {
	if o == nil || common.IsNil(o.Commission) {
		var ret float32
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) GetCommissionOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInnerLegsInner) HasCommission() bool {
	if o != nil && !common.IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given float32 and assigns it to the Commission field.
func (o *AccountBlockTradeListResponseInnerLegsInner) SetCommission(v float32) {
	o.Commission = &v
}

func (o AccountBlockTradeListResponseInnerLegsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountBlockTradeListResponseInnerLegsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.OrderPrice) {
		toSerialize["orderPrice"] = o.OrderPrice
	}
	if !common.IsNil(o.OrderQuantity) {
		toSerialize["orderQuantity"] = o.OrderQuantity
	}
	if !common.IsNil(o.OrderStatus) {
		toSerialize["orderStatus"] = o.OrderStatus
	}
	if !common.IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !common.IsNil(o.ExecutedAmount) {
		toSerialize["executedAmount"] = o.ExecutedAmount
	}
	if !common.IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !common.IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	if !common.IsNil(o.OrderSide) {
		toSerialize["orderSide"] = o.OrderSide
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
	}
	if !common.IsNil(o.TradePrice) {
		toSerialize["tradePrice"] = o.TradePrice
	}
	if !common.IsNil(o.TradeQty) {
		toSerialize["tradeQty"] = o.TradeQty
	}
	if !common.IsNil(o.TradeTime) {
		toSerialize["tradeTime"] = o.TradeTime
	}
	if !common.IsNil(o.Liquidity) {
		toSerialize["liquidity"] = o.Liquidity
	}
	if !common.IsNil(o.Commission) {
		toSerialize["commission"] = o.Commission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountBlockTradeListResponseInnerLegsInner) UnmarshalJSON(data []byte) (err error) {
	varAccountBlockTradeListResponseInnerLegsInner := _AccountBlockTradeListResponseInnerLegsInner{}

	err = json.Unmarshal(data, &varAccountBlockTradeListResponseInnerLegsInner)

	if err != nil {
		return err
	}

	*o = AccountBlockTradeListResponseInnerLegsInner(varAccountBlockTradeListResponseInnerLegsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "orderPrice")
		delete(additionalProperties, "orderQuantity")
		delete(additionalProperties, "orderStatus")
		delete(additionalProperties, "executedQty")
		delete(additionalProperties, "executedAmount")
		delete(additionalProperties, "fee")
		delete(additionalProperties, "orderType")
		delete(additionalProperties, "orderSide")
		delete(additionalProperties, "id")
		delete(additionalProperties, "tradeId")
		delete(additionalProperties, "tradePrice")
		delete(additionalProperties, "tradeQty")
		delete(additionalProperties, "tradeTime")
		delete(additionalProperties, "liquidity")
		delete(additionalProperties, "commission")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountBlockTradeListResponseInnerLegsInner struct {
	value *AccountBlockTradeListResponseInnerLegsInner
	isSet bool
}

func (v NullableAccountBlockTradeListResponseInnerLegsInner) Get() *AccountBlockTradeListResponseInnerLegsInner {
	return v.value
}

func (v *NullableAccountBlockTradeListResponseInnerLegsInner) Set(val *AccountBlockTradeListResponseInnerLegsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBlockTradeListResponseInnerLegsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBlockTradeListResponseInnerLegsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBlockTradeListResponseInnerLegsInner(val *AccountBlockTradeListResponseInnerLegsInner) *NullableAccountBlockTradeListResponseInnerLegsInner {
	return &NullableAccountBlockTradeListResponseInnerLegsInner{value: val, isSet: true}
}

func (v NullableAccountBlockTradeListResponseInnerLegsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBlockTradeListResponseInnerLegsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
