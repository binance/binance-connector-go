/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MyTradesResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MyTradesResponseInner{}

// MyTradesResponseInner struct for MyTradesResponseInner
type MyTradesResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	Id                   *int64  `json:"id,omitempty"`
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
	IsBestMatch          *bool   `json:"isBestMatch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MyTradesResponseInner MyTradesResponseInner

// NewMyTradesResponseInner instantiates a new MyTradesResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyTradesResponseInner() *MyTradesResponseInner {
	this := MyTradesResponseInner{}
	return &this
}

// NewMyTradesResponseInnerWithDefaults instantiates a new MyTradesResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyTradesResponseInnerWithDefaults() *MyTradesResponseInner {
	this := MyTradesResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MyTradesResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MyTradesResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MyTradesResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MyTradesResponseInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MyTradesResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *MyTradesResponseInner) SetId(v int64) {
	o.Id = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *MyTradesResponseInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *MyTradesResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *MyTradesResponseInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *MyTradesResponseInner) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseInner) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *MyTradesResponseInner) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *MyTradesResponseInner) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *MyTradesResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *MyTradesResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *MyTradesResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *MyTradesResponseInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *MyTradesResponseInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *MyTradesResponseInner) SetQty(v string) {
	o.Qty = &v
}

// GetQuoteQty returns the QuoteQty field value if set, zero value otherwise.
func (o *MyTradesResponseInner) GetQuoteQty() string {
	if o == nil || common.IsNil(o.QuoteQty) {
		var ret string
		return ret
	}
	return *o.QuoteQty
}

// GetQuoteQtyOk returns a tuple with the QuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseInner) GetQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteQty) {
		return nil, false
	}
	return o.QuoteQty, true
}

// HasQuoteQty returns a boolean if a field has been set.
func (o *MyTradesResponseInner) HasQuoteQty() bool {
	if o != nil && !common.IsNil(o.QuoteQty) {
		return true
	}

	return false
}

// SetQuoteQty gets a reference to the given string and assigns it to the QuoteQty field.
func (o *MyTradesResponseInner) SetQuoteQty(v string) {
	o.QuoteQty = &v
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *MyTradesResponseInner) GetCommission() string {
	if o == nil || common.IsNil(o.Commission) {
		var ret string
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseInner) GetCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *MyTradesResponseInner) HasCommission() bool {
	if o != nil && !common.IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given string and assigns it to the Commission field.
func (o *MyTradesResponseInner) SetCommission(v string) {
	o.Commission = &v
}

// GetCommissionAsset returns the CommissionAsset field value if set, zero value otherwise.
func (o *MyTradesResponseInner) GetCommissionAsset() string {
	if o == nil || common.IsNil(o.CommissionAsset) {
		var ret string
		return ret
	}
	return *o.CommissionAsset
}

// GetCommissionAssetOk returns a tuple with the CommissionAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseInner) GetCommissionAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.CommissionAsset) {
		return nil, false
	}
	return o.CommissionAsset, true
}

// HasCommissionAsset returns a boolean if a field has been set.
func (o *MyTradesResponseInner) HasCommissionAsset() bool {
	if o != nil && !common.IsNil(o.CommissionAsset) {
		return true
	}

	return false
}

// SetCommissionAsset gets a reference to the given string and assigns it to the CommissionAsset field.
func (o *MyTradesResponseInner) SetCommissionAsset(v string) {
	o.CommissionAsset = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *MyTradesResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *MyTradesResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *MyTradesResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetIsBuyer returns the IsBuyer field value if set, zero value otherwise.
func (o *MyTradesResponseInner) GetIsBuyer() bool {
	if o == nil || common.IsNil(o.IsBuyer) {
		var ret bool
		return ret
	}
	return *o.IsBuyer
}

// GetIsBuyerOk returns a tuple with the IsBuyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseInner) GetIsBuyerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBuyer) {
		return nil, false
	}
	return o.IsBuyer, true
}

// HasIsBuyer returns a boolean if a field has been set.
func (o *MyTradesResponseInner) HasIsBuyer() bool {
	if o != nil && !common.IsNil(o.IsBuyer) {
		return true
	}

	return false
}

// SetIsBuyer gets a reference to the given bool and assigns it to the IsBuyer field.
func (o *MyTradesResponseInner) SetIsBuyer(v bool) {
	o.IsBuyer = &v
}

// GetIsMaker returns the IsMaker field value if set, zero value otherwise.
func (o *MyTradesResponseInner) GetIsMaker() bool {
	if o == nil || common.IsNil(o.IsMaker) {
		var ret bool
		return ret
	}
	return *o.IsMaker
}

// GetIsMakerOk returns a tuple with the IsMaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseInner) GetIsMakerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsMaker) {
		return nil, false
	}
	return o.IsMaker, true
}

// HasIsMaker returns a boolean if a field has been set.
func (o *MyTradesResponseInner) HasIsMaker() bool {
	if o != nil && !common.IsNil(o.IsMaker) {
		return true
	}

	return false
}

// SetIsMaker gets a reference to the given bool and assigns it to the IsMaker field.
func (o *MyTradesResponseInner) SetIsMaker(v bool) {
	o.IsMaker = &v
}

// GetIsBestMatch returns the IsBestMatch field value if set, zero value otherwise.
func (o *MyTradesResponseInner) GetIsBestMatch() bool {
	if o == nil || common.IsNil(o.IsBestMatch) {
		var ret bool
		return ret
	}
	return *o.IsBestMatch
}

// GetIsBestMatchOk returns a tuple with the IsBestMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseInner) GetIsBestMatchOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBestMatch) {
		return nil, false
	}
	return o.IsBestMatch, true
}

// HasIsBestMatch returns a boolean if a field has been set.
func (o *MyTradesResponseInner) HasIsBestMatch() bool {
	if o != nil && !common.IsNil(o.IsBestMatch) {
		return true
	}

	return false
}

// SetIsBestMatch gets a reference to the given bool and assigns it to the IsBestMatch field.
func (o *MyTradesResponseInner) SetIsBestMatch(v bool) {
	o.IsBestMatch = &v
}

func (o MyTradesResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyTradesResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
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
	if !common.IsNil(o.IsBestMatch) {
		toSerialize["isBestMatch"] = o.IsBestMatch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MyTradesResponseInner) UnmarshalJSON(data []byte) (err error) {
	varMyTradesResponseInner := _MyTradesResponseInner{}

	err = json.Unmarshal(data, &varMyTradesResponseInner)

	if err != nil {
		return err
	}

	*o = MyTradesResponseInner(varMyTradesResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "id")
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
		delete(additionalProperties, "isBestMatch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMyTradesResponseInner struct {
	value *MyTradesResponseInner
	isSet bool
}

func (v NullableMyTradesResponseInner) Get() *MyTradesResponseInner {
	return v.value
}

func (v *NullableMyTradesResponseInner) Set(val *MyTradesResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMyTradesResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMyTradesResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyTradesResponseInner(val *MyTradesResponseInner) *NullableMyTradesResponseInner {
	return &NullableMyTradesResponseInner{value: val, isSet: true}
}

func (v NullableMyTradesResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyTradesResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
