/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MyAllocationsResponseResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MyAllocationsResponseResultInner{}

// MyAllocationsResponseResultInner struct for MyAllocationsResponseResultInner
type MyAllocationsResponseResultInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	AllocationId         *int64  `json:"allocationId,omitempty"`
	AllocationType       *string `json:"allocationType,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	OrderListId          *int64  `json:"orderListId,omitempty"`
	Price                *string `json:"price,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	QuoteQty             *string `json:"quoteQty,omitempty"`
	Commission           *string `json:"commission,omitempty"`
	CommissionAsset      *string `json:"commissionAsset,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	IsBuyer              *bool   `json:"isBuyer,omitempty"`
	IsMaker              *bool   `json:"isMaker,omitempty"`
	IsAllocator          *bool   `json:"isAllocator,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MyAllocationsResponseResultInner MyAllocationsResponseResultInner

// NewMyAllocationsResponseResultInner instantiates a new MyAllocationsResponseResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyAllocationsResponseResultInner() *MyAllocationsResponseResultInner {
	this := MyAllocationsResponseResultInner{}
	return &this
}

// NewMyAllocationsResponseResultInnerWithDefaults instantiates a new MyAllocationsResponseResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyAllocationsResponseResultInnerWithDefaults() *MyAllocationsResponseResultInner {
	this := MyAllocationsResponseResultInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MyAllocationsResponseResultInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyAllocationsResponseResultInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MyAllocationsResponseResultInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MyAllocationsResponseResultInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetAllocationId returns the AllocationId field value if set, zero value otherwise.
func (o *MyAllocationsResponseResultInner) GetAllocationId() int64 {
	if o == nil || common.IsNil(o.AllocationId) {
		var ret int64
		return ret
	}
	return *o.AllocationId
}

// GetAllocationIdOk returns a tuple with the AllocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyAllocationsResponseResultInner) GetAllocationIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AllocationId) {
		return nil, false
	}
	return o.AllocationId, true
}

// HasAllocationId returns a boolean if a field has been set.
func (o *MyAllocationsResponseResultInner) HasAllocationId() bool {
	if o != nil && !common.IsNil(o.AllocationId) {
		return true
	}

	return false
}

// SetAllocationId gets a reference to the given int64 and assigns it to the AllocationId field.
func (o *MyAllocationsResponseResultInner) SetAllocationId(v int64) {
	o.AllocationId = &v
}

// GetAllocationType returns the AllocationType field value if set, zero value otherwise.
func (o *MyAllocationsResponseResultInner) GetAllocationType() string {
	if o == nil || common.IsNil(o.AllocationType) {
		var ret string
		return ret
	}
	return *o.AllocationType
}

// GetAllocationTypeOk returns a tuple with the AllocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyAllocationsResponseResultInner) GetAllocationTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AllocationType) {
		return nil, false
	}
	return o.AllocationType, true
}

// HasAllocationType returns a boolean if a field has been set.
func (o *MyAllocationsResponseResultInner) HasAllocationType() bool {
	if o != nil && !common.IsNil(o.AllocationType) {
		return true
	}

	return false
}

// SetAllocationType gets a reference to the given string and assigns it to the AllocationType field.
func (o *MyAllocationsResponseResultInner) SetAllocationType(v string) {
	o.AllocationType = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *MyAllocationsResponseResultInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyAllocationsResponseResultInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *MyAllocationsResponseResultInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *MyAllocationsResponseResultInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *MyAllocationsResponseResultInner) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyAllocationsResponseResultInner) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *MyAllocationsResponseResultInner) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *MyAllocationsResponseResultInner) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *MyAllocationsResponseResultInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyAllocationsResponseResultInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *MyAllocationsResponseResultInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *MyAllocationsResponseResultInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *MyAllocationsResponseResultInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyAllocationsResponseResultInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *MyAllocationsResponseResultInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *MyAllocationsResponseResultInner) SetQty(v string) {
	o.Qty = &v
}

// GetQuoteQty returns the QuoteQty field value if set, zero value otherwise.
func (o *MyAllocationsResponseResultInner) GetQuoteQty() string {
	if o == nil || common.IsNil(o.QuoteQty) {
		var ret string
		return ret
	}
	return *o.QuoteQty
}

// GetQuoteQtyOk returns a tuple with the QuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyAllocationsResponseResultInner) GetQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteQty) {
		return nil, false
	}
	return o.QuoteQty, true
}

// HasQuoteQty returns a boolean if a field has been set.
func (o *MyAllocationsResponseResultInner) HasQuoteQty() bool {
	if o != nil && !common.IsNil(o.QuoteQty) {
		return true
	}

	return false
}

// SetQuoteQty gets a reference to the given string and assigns it to the QuoteQty field.
func (o *MyAllocationsResponseResultInner) SetQuoteQty(v string) {
	o.QuoteQty = &v
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *MyAllocationsResponseResultInner) GetCommission() string {
	if o == nil || common.IsNil(o.Commission) {
		var ret string
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyAllocationsResponseResultInner) GetCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *MyAllocationsResponseResultInner) HasCommission() bool {
	if o != nil && !common.IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given string and assigns it to the Commission field.
func (o *MyAllocationsResponseResultInner) SetCommission(v string) {
	o.Commission = &v
}

// GetCommissionAsset returns the CommissionAsset field value if set, zero value otherwise.
func (o *MyAllocationsResponseResultInner) GetCommissionAsset() string {
	if o == nil || common.IsNil(o.CommissionAsset) {
		var ret string
		return ret
	}
	return *o.CommissionAsset
}

// GetCommissionAssetOk returns a tuple with the CommissionAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyAllocationsResponseResultInner) GetCommissionAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.CommissionAsset) {
		return nil, false
	}
	return o.CommissionAsset, true
}

// HasCommissionAsset returns a boolean if a field has been set.
func (o *MyAllocationsResponseResultInner) HasCommissionAsset() bool {
	if o != nil && !common.IsNil(o.CommissionAsset) {
		return true
	}

	return false
}

// SetCommissionAsset gets a reference to the given string and assigns it to the CommissionAsset field.
func (o *MyAllocationsResponseResultInner) SetCommissionAsset(v string) {
	o.CommissionAsset = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *MyAllocationsResponseResultInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyAllocationsResponseResultInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *MyAllocationsResponseResultInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *MyAllocationsResponseResultInner) SetTime(v int64) {
	o.Time = &v
}

// GetIsBuyer returns the IsBuyer field value if set, zero value otherwise.
func (o *MyAllocationsResponseResultInner) GetIsBuyer() bool {
	if o == nil || common.IsNil(o.IsBuyer) {
		var ret bool
		return ret
	}
	return *o.IsBuyer
}

// GetIsBuyerOk returns a tuple with the IsBuyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyAllocationsResponseResultInner) GetIsBuyerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBuyer) {
		return nil, false
	}
	return o.IsBuyer, true
}

// HasIsBuyer returns a boolean if a field has been set.
func (o *MyAllocationsResponseResultInner) HasIsBuyer() bool {
	if o != nil && !common.IsNil(o.IsBuyer) {
		return true
	}

	return false
}

// SetIsBuyer gets a reference to the given bool and assigns it to the IsBuyer field.
func (o *MyAllocationsResponseResultInner) SetIsBuyer(v bool) {
	o.IsBuyer = &v
}

// GetIsMaker returns the IsMaker field value if set, zero value otherwise.
func (o *MyAllocationsResponseResultInner) GetIsMaker() bool {
	if o == nil || common.IsNil(o.IsMaker) {
		var ret bool
		return ret
	}
	return *o.IsMaker
}

// GetIsMakerOk returns a tuple with the IsMaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyAllocationsResponseResultInner) GetIsMakerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsMaker) {
		return nil, false
	}
	return o.IsMaker, true
}

// HasIsMaker returns a boolean if a field has been set.
func (o *MyAllocationsResponseResultInner) HasIsMaker() bool {
	if o != nil && !common.IsNil(o.IsMaker) {
		return true
	}

	return false
}

// SetIsMaker gets a reference to the given bool and assigns it to the IsMaker field.
func (o *MyAllocationsResponseResultInner) SetIsMaker(v bool) {
	o.IsMaker = &v
}

// GetIsAllocator returns the IsAllocator field value if set, zero value otherwise.
func (o *MyAllocationsResponseResultInner) GetIsAllocator() bool {
	if o == nil || common.IsNil(o.IsAllocator) {
		var ret bool
		return ret
	}
	return *o.IsAllocator
}

// GetIsAllocatorOk returns a tuple with the IsAllocator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyAllocationsResponseResultInner) GetIsAllocatorOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsAllocator) {
		return nil, false
	}
	return o.IsAllocator, true
}

// HasIsAllocator returns a boolean if a field has been set.
func (o *MyAllocationsResponseResultInner) HasIsAllocator() bool {
	if o != nil && !common.IsNil(o.IsAllocator) {
		return true
	}

	return false
}

// SetIsAllocator gets a reference to the given bool and assigns it to the IsAllocator field.
func (o *MyAllocationsResponseResultInner) SetIsAllocator(v bool) {
	o.IsAllocator = &v
}

func (o MyAllocationsResponseResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyAllocationsResponseResultInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.AllocationId) {
		toSerialize["allocationId"] = o.AllocationId
	}
	if !common.IsNil(o.AllocationType) {
		toSerialize["allocationType"] = o.AllocationType
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.OrderListId) {
		toSerialize["orderListId"] = o.OrderListId
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !common.IsNil(o.QuoteQty) {
		toSerialize["quoteQty"] = o.QuoteQty
	}
	if !common.IsNil(o.Commission) {
		toSerialize["commission"] = o.Commission
	}
	if !common.IsNil(o.CommissionAsset) {
		toSerialize["commissionAsset"] = o.CommissionAsset
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.IsBuyer) {
		toSerialize["isBuyer"] = o.IsBuyer
	}
	if !common.IsNil(o.IsMaker) {
		toSerialize["isMaker"] = o.IsMaker
	}
	if !common.IsNil(o.IsAllocator) {
		toSerialize["isAllocator"] = o.IsAllocator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MyAllocationsResponseResultInner) UnmarshalJSON(data []byte) (err error) {
	varMyAllocationsResponseResultInner := _MyAllocationsResponseResultInner{}

	err = json.Unmarshal(data, &varMyAllocationsResponseResultInner)

	if err != nil {
		return err
	}

	*o = MyAllocationsResponseResultInner(varMyAllocationsResponseResultInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "allocationId")
		delete(additionalProperties, "allocationType")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "orderListId")
		delete(additionalProperties, "price")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "quoteQty")
		delete(additionalProperties, "commission")
		delete(additionalProperties, "commissionAsset")
		delete(additionalProperties, "time")
		delete(additionalProperties, "isBuyer")
		delete(additionalProperties, "isMaker")
		delete(additionalProperties, "isAllocator")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMyAllocationsResponseResultInner struct {
	value *MyAllocationsResponseResultInner
	isSet bool
}

func (v NullableMyAllocationsResponseResultInner) Get() *MyAllocationsResponseResultInner {
	return v.value
}

func (v *NullableMyAllocationsResponseResultInner) Set(val *MyAllocationsResponseResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMyAllocationsResponseResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMyAllocationsResponseResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyAllocationsResponseResultInner(val *MyAllocationsResponseResultInner) *NullableMyAllocationsResponseResultInner {
	return &NullableMyAllocationsResponseResultInner{value: val, isSet: true}
}

func (v NullableMyAllocationsResponseResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyAllocationsResponseResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
