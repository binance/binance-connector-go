/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MarginAccountNewOcoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginAccountNewOcoResponse{}

// MarginAccountNewOcoResponse struct for MarginAccountNewOcoResponse
type MarginAccountNewOcoResponse struct {
	OrderListId           *int64                                         `json:"orderListId,omitempty"`
	ContingencyType       *string                                        `json:"contingencyType,omitempty"`
	ListStatusType        *string                                        `json:"listStatusType,omitempty"`
	ListOrderStatus       *string                                        `json:"listOrderStatus,omitempty"`
	ListClientOrderId     *string                                        `json:"listClientOrderId,omitempty"`
	TransactionTime       *int64                                         `json:"transactionTime,omitempty"`
	Symbol                *string                                        `json:"symbol,omitempty"`
	MarginBuyBorrowAmount *string                                        `json:"marginBuyBorrowAmount,omitempty"`
	MarginBuyBorrowAsset  *string                                        `json:"marginBuyBorrowAsset,omitempty"`
	Orders                []MarginAccountNewOcoResponseOrdersInner       `json:"orders,omitempty"`
	OrderReports          []MarginAccountNewOcoResponseOrderReportsInner `json:"orderReports,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _MarginAccountNewOcoResponse MarginAccountNewOcoResponse

// NewMarginAccountNewOcoResponse instantiates a new MarginAccountNewOcoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginAccountNewOcoResponse() *MarginAccountNewOcoResponse {
	this := MarginAccountNewOcoResponse{}
	return &this
}

// NewMarginAccountNewOcoResponseWithDefaults instantiates a new MarginAccountNewOcoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginAccountNewOcoResponseWithDefaults() *MarginAccountNewOcoResponse {
	this := MarginAccountNewOcoResponse{}
	return &this
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *MarginAccountNewOcoResponse) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOcoResponse) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *MarginAccountNewOcoResponse) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *MarginAccountNewOcoResponse) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetContingencyType returns the ContingencyType field value if set, zero value otherwise.
func (o *MarginAccountNewOcoResponse) GetContingencyType() string {
	if o == nil || common.IsNil(o.ContingencyType) {
		var ret string
		return ret
	}
	return *o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOcoResponse) GetContingencyTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContingencyType) {
		return nil, false
	}
	return o.ContingencyType, true
}

// HasContingencyType returns a boolean if a field has been set.
func (o *MarginAccountNewOcoResponse) HasContingencyType() bool {
	if o != nil && !common.IsNil(o.ContingencyType) {
		return true
	}

	return false
}

// SetContingencyType gets a reference to the given string and assigns it to the ContingencyType field.
func (o *MarginAccountNewOcoResponse) SetContingencyType(v string) {
	o.ContingencyType = &v
}

// GetListStatusType returns the ListStatusType field value if set, zero value otherwise.
func (o *MarginAccountNewOcoResponse) GetListStatusType() string {
	if o == nil || common.IsNil(o.ListStatusType) {
		var ret string
		return ret
	}
	return *o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOcoResponse) GetListStatusTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListStatusType) {
		return nil, false
	}
	return o.ListStatusType, true
}

// HasListStatusType returns a boolean if a field has been set.
func (o *MarginAccountNewOcoResponse) HasListStatusType() bool {
	if o != nil && !common.IsNil(o.ListStatusType) {
		return true
	}

	return false
}

// SetListStatusType gets a reference to the given string and assigns it to the ListStatusType field.
func (o *MarginAccountNewOcoResponse) SetListStatusType(v string) {
	o.ListStatusType = &v
}

// GetListOrderStatus returns the ListOrderStatus field value if set, zero value otherwise.
func (o *MarginAccountNewOcoResponse) GetListOrderStatus() string {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		var ret string
		return ret
	}
	return *o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOcoResponse) GetListOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		return nil, false
	}
	return o.ListOrderStatus, true
}

// HasListOrderStatus returns a boolean if a field has been set.
func (o *MarginAccountNewOcoResponse) HasListOrderStatus() bool {
	if o != nil && !common.IsNil(o.ListOrderStatus) {
		return true
	}

	return false
}

// SetListOrderStatus gets a reference to the given string and assigns it to the ListOrderStatus field.
func (o *MarginAccountNewOcoResponse) SetListOrderStatus(v string) {
	o.ListOrderStatus = &v
}

// GetListClientOrderId returns the ListClientOrderId field value if set, zero value otherwise.
func (o *MarginAccountNewOcoResponse) GetListClientOrderId() string {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		var ret string
		return ret
	}
	return *o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOcoResponse) GetListClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		return nil, false
	}
	return o.ListClientOrderId, true
}

// HasListClientOrderId returns a boolean if a field has been set.
func (o *MarginAccountNewOcoResponse) HasListClientOrderId() bool {
	if o != nil && !common.IsNil(o.ListClientOrderId) {
		return true
	}

	return false
}

// SetListClientOrderId gets a reference to the given string and assigns it to the ListClientOrderId field.
func (o *MarginAccountNewOcoResponse) SetListClientOrderId(v string) {
	o.ListClientOrderId = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *MarginAccountNewOcoResponse) GetTransactionTime() int64 {
	if o == nil || common.IsNil(o.TransactionTime) {
		var ret int64
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOcoResponse) GetTransactionTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *MarginAccountNewOcoResponse) HasTransactionTime() bool {
	if o != nil && !common.IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given int64 and assigns it to the TransactionTime field.
func (o *MarginAccountNewOcoResponse) SetTransactionTime(v int64) {
	o.TransactionTime = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarginAccountNewOcoResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOcoResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarginAccountNewOcoResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarginAccountNewOcoResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetMarginBuyBorrowAmount returns the MarginBuyBorrowAmount field value if set, zero value otherwise.
func (o *MarginAccountNewOcoResponse) GetMarginBuyBorrowAmount() string {
	if o == nil || common.IsNil(o.MarginBuyBorrowAmount) {
		var ret string
		return ret
	}
	return *o.MarginBuyBorrowAmount
}

// GetMarginBuyBorrowAmountOk returns a tuple with the MarginBuyBorrowAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOcoResponse) GetMarginBuyBorrowAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginBuyBorrowAmount) {
		return nil, false
	}
	return o.MarginBuyBorrowAmount, true
}

// HasMarginBuyBorrowAmount returns a boolean if a field has been set.
func (o *MarginAccountNewOcoResponse) HasMarginBuyBorrowAmount() bool {
	if o != nil && !common.IsNil(o.MarginBuyBorrowAmount) {
		return true
	}

	return false
}

// SetMarginBuyBorrowAmount gets a reference to the given string and assigns it to the MarginBuyBorrowAmount field.
func (o *MarginAccountNewOcoResponse) SetMarginBuyBorrowAmount(v string) {
	o.MarginBuyBorrowAmount = &v
}

// GetMarginBuyBorrowAsset returns the MarginBuyBorrowAsset field value if set, zero value otherwise.
func (o *MarginAccountNewOcoResponse) GetMarginBuyBorrowAsset() string {
	if o == nil || common.IsNil(o.MarginBuyBorrowAsset) {
		var ret string
		return ret
	}
	return *o.MarginBuyBorrowAsset
}

// GetMarginBuyBorrowAssetOk returns a tuple with the MarginBuyBorrowAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOcoResponse) GetMarginBuyBorrowAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginBuyBorrowAsset) {
		return nil, false
	}
	return o.MarginBuyBorrowAsset, true
}

// HasMarginBuyBorrowAsset returns a boolean if a field has been set.
func (o *MarginAccountNewOcoResponse) HasMarginBuyBorrowAsset() bool {
	if o != nil && !common.IsNil(o.MarginBuyBorrowAsset) {
		return true
	}

	return false
}

// SetMarginBuyBorrowAsset gets a reference to the given string and assigns it to the MarginBuyBorrowAsset field.
func (o *MarginAccountNewOcoResponse) SetMarginBuyBorrowAsset(v string) {
	o.MarginBuyBorrowAsset = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *MarginAccountNewOcoResponse) GetOrders() []MarginAccountNewOcoResponseOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []MarginAccountNewOcoResponseOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOcoResponse) GetOrdersOk() ([]MarginAccountNewOcoResponseOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *MarginAccountNewOcoResponse) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []MarginAccountNewOcoResponseOrdersInner and assigns it to the Orders field.
func (o *MarginAccountNewOcoResponse) SetOrders(v []MarginAccountNewOcoResponseOrdersInner) {
	o.Orders = v
}

// GetOrderReports returns the OrderReports field value if set, zero value otherwise.
func (o *MarginAccountNewOcoResponse) GetOrderReports() []MarginAccountNewOcoResponseOrderReportsInner {
	if o == nil || common.IsNil(o.OrderReports) {
		var ret []MarginAccountNewOcoResponseOrderReportsInner
		return ret
	}
	return o.OrderReports
}

// GetOrderReportsOk returns a tuple with the OrderReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOcoResponse) GetOrderReportsOk() ([]MarginAccountNewOcoResponseOrderReportsInner, bool) {
	if o == nil || common.IsNil(o.OrderReports) {
		return nil, false
	}
	return o.OrderReports, true
}

// HasOrderReports returns a boolean if a field has been set.
func (o *MarginAccountNewOcoResponse) HasOrderReports() bool {
	if o != nil && !common.IsNil(o.OrderReports) {
		return true
	}

	return false
}

// SetOrderReports gets a reference to the given []MarginAccountNewOcoResponseOrderReportsInner and assigns it to the OrderReports field.
func (o *MarginAccountNewOcoResponse) SetOrderReports(v []MarginAccountNewOcoResponseOrderReportsInner) {
	o.OrderReports = v
}

func (o MarginAccountNewOcoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginAccountNewOcoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderListId) {
		toSerialize["orderListId"] = o.OrderListId
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
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.MarginBuyBorrowAmount) {
		toSerialize["marginBuyBorrowAmount"] = o.MarginBuyBorrowAmount
	}
	if !common.IsNil(o.MarginBuyBorrowAsset) {
		toSerialize["marginBuyBorrowAsset"] = o.MarginBuyBorrowAsset
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

func (o *MarginAccountNewOcoResponse) UnmarshalJSON(data []byte) (err error) {
	varMarginAccountNewOcoResponse := _MarginAccountNewOcoResponse{}

	err = json.Unmarshal(data, &varMarginAccountNewOcoResponse)

	if err != nil {
		return err
	}

	*o = MarginAccountNewOcoResponse(varMarginAccountNewOcoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderListId")
		delete(additionalProperties, "contingencyType")
		delete(additionalProperties, "listStatusType")
		delete(additionalProperties, "listOrderStatus")
		delete(additionalProperties, "listClientOrderId")
		delete(additionalProperties, "transactionTime")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "marginBuyBorrowAmount")
		delete(additionalProperties, "marginBuyBorrowAsset")
		delete(additionalProperties, "orders")
		delete(additionalProperties, "orderReports")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginAccountNewOcoResponse struct {
	value *MarginAccountNewOcoResponse
	isSet bool
}

func (v NullableMarginAccountNewOcoResponse) Get() *MarginAccountNewOcoResponse {
	return v.value
}

func (v *NullableMarginAccountNewOcoResponse) Set(val *MarginAccountNewOcoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountNewOcoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountNewOcoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountNewOcoResponse(val *MarginAccountNewOcoResponse) *NullableMarginAccountNewOcoResponse {
	return &NullableMarginAccountNewOcoResponse{value: val, isSet: true}
}

func (v NullableMarginAccountNewOcoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountNewOcoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
