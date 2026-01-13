/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MyTradesResponseResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MyTradesResponseResultInner{}

// MyTradesResponseResultInner struct for MyTradesResponseResultInner
type MyTradesResponseResultInner struct {
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

type _MyTradesResponseResultInner MyTradesResponseResultInner

// NewMyTradesResponseResultInner instantiates a new MyTradesResponseResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyTradesResponseResultInner() *MyTradesResponseResultInner {
	this := MyTradesResponseResultInner{}
	return &this
}

// NewMyTradesResponseResultInnerWithDefaults instantiates a new MyTradesResponseResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyTradesResponseResultInnerWithDefaults() *MyTradesResponseResultInner {
	this := MyTradesResponseResultInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MyTradesResponseResultInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseResultInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MyTradesResponseResultInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MyTradesResponseResultInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MyTradesResponseResultInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseResultInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MyTradesResponseResultInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *MyTradesResponseResultInner) SetId(v int64) {
	o.Id = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *MyTradesResponseResultInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseResultInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *MyTradesResponseResultInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *MyTradesResponseResultInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *MyTradesResponseResultInner) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseResultInner) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *MyTradesResponseResultInner) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *MyTradesResponseResultInner) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *MyTradesResponseResultInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseResultInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *MyTradesResponseResultInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *MyTradesResponseResultInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *MyTradesResponseResultInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseResultInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *MyTradesResponseResultInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *MyTradesResponseResultInner) SetQty(v string) {
	o.Qty = &v
}

// GetQuoteQty returns the QuoteQty field value if set, zero value otherwise.
func (o *MyTradesResponseResultInner) GetQuoteQty() string {
	if o == nil || common.IsNil(o.QuoteQty) {
		var ret string
		return ret
	}
	return *o.QuoteQty
}

// GetQuoteQtyOk returns a tuple with the QuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseResultInner) GetQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteQty) {
		return nil, false
	}
	return o.QuoteQty, true
}

// HasQuoteQty returns a boolean if a field has been set.
func (o *MyTradesResponseResultInner) HasQuoteQty() bool {
	if o != nil && !common.IsNil(o.QuoteQty) {
		return true
	}

	return false
}

// SetQuoteQty gets a reference to the given string and assigns it to the QuoteQty field.
func (o *MyTradesResponseResultInner) SetQuoteQty(v string) {
	o.QuoteQty = &v
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *MyTradesResponseResultInner) GetCommission() string {
	if o == nil || common.IsNil(o.Commission) {
		var ret string
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseResultInner) GetCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *MyTradesResponseResultInner) HasCommission() bool {
	if o != nil && !common.IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given string and assigns it to the Commission field.
func (o *MyTradesResponseResultInner) SetCommission(v string) {
	o.Commission = &v
}

// GetCommissionAsset returns the CommissionAsset field value if set, zero value otherwise.
func (o *MyTradesResponseResultInner) GetCommissionAsset() string {
	if o == nil || common.IsNil(o.CommissionAsset) {
		var ret string
		return ret
	}
	return *o.CommissionAsset
}

// GetCommissionAssetOk returns a tuple with the CommissionAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseResultInner) GetCommissionAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.CommissionAsset) {
		return nil, false
	}
	return o.CommissionAsset, true
}

// HasCommissionAsset returns a boolean if a field has been set.
func (o *MyTradesResponseResultInner) HasCommissionAsset() bool {
	if o != nil && !common.IsNil(o.CommissionAsset) {
		return true
	}

	return false
}

// SetCommissionAsset gets a reference to the given string and assigns it to the CommissionAsset field.
func (o *MyTradesResponseResultInner) SetCommissionAsset(v string) {
	o.CommissionAsset = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *MyTradesResponseResultInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseResultInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *MyTradesResponseResultInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *MyTradesResponseResultInner) SetTime(v int64) {
	o.Time = &v
}

// GetIsBuyer returns the IsBuyer field value if set, zero value otherwise.
func (o *MyTradesResponseResultInner) GetIsBuyer() bool {
	if o == nil || common.IsNil(o.IsBuyer) {
		var ret bool
		return ret
	}
	return *o.IsBuyer
}

// GetIsBuyerOk returns a tuple with the IsBuyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseResultInner) GetIsBuyerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBuyer) {
		return nil, false
	}
	return o.IsBuyer, true
}

// HasIsBuyer returns a boolean if a field has been set.
func (o *MyTradesResponseResultInner) HasIsBuyer() bool {
	if o != nil && !common.IsNil(o.IsBuyer) {
		return true
	}

	return false
}

// SetIsBuyer gets a reference to the given bool and assigns it to the IsBuyer field.
func (o *MyTradesResponseResultInner) SetIsBuyer(v bool) {
	o.IsBuyer = &v
}

// GetIsMaker returns the IsMaker field value if set, zero value otherwise.
func (o *MyTradesResponseResultInner) GetIsMaker() bool {
	if o == nil || common.IsNil(o.IsMaker) {
		var ret bool
		return ret
	}
	return *o.IsMaker
}

// GetIsMakerOk returns a tuple with the IsMaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseResultInner) GetIsMakerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsMaker) {
		return nil, false
	}
	return o.IsMaker, true
}

// HasIsMaker returns a boolean if a field has been set.
func (o *MyTradesResponseResultInner) HasIsMaker() bool {
	if o != nil && !common.IsNil(o.IsMaker) {
		return true
	}

	return false
}

// SetIsMaker gets a reference to the given bool and assigns it to the IsMaker field.
func (o *MyTradesResponseResultInner) SetIsMaker(v bool) {
	o.IsMaker = &v
}

// GetIsBestMatch returns the IsBestMatch field value if set, zero value otherwise.
func (o *MyTradesResponseResultInner) GetIsBestMatch() bool {
	if o == nil || common.IsNil(o.IsBestMatch) {
		var ret bool
		return ret
	}
	return *o.IsBestMatch
}

// GetIsBestMatchOk returns a tuple with the IsBestMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTradesResponseResultInner) GetIsBestMatchOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBestMatch) {
		return nil, false
	}
	return o.IsBestMatch, true
}

// HasIsBestMatch returns a boolean if a field has been set.
func (o *MyTradesResponseResultInner) HasIsBestMatch() bool {
	if o != nil && !common.IsNil(o.IsBestMatch) {
		return true
	}

	return false
}

// SetIsBestMatch gets a reference to the given bool and assigns it to the IsBestMatch field.
func (o *MyTradesResponseResultInner) SetIsBestMatch(v bool) {
	o.IsBestMatch = &v
}

func (o MyTradesResponseResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyTradesResponseResultInner) ToMap() (map[string]interface{}, error) {
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

func (o *MyTradesResponseResultInner) UnmarshalJSON(data []byte) (err error) {
	varMyTradesResponseResultInner := _MyTradesResponseResultInner{}

	err = json.Unmarshal(data, &varMyTradesResponseResultInner)

	if err != nil {
		return err
	}

	*o = MyTradesResponseResultInner(varMyTradesResponseResultInner)

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

type NullableMyTradesResponseResultInner struct {
	value *MyTradesResponseResultInner
	isSet bool
}

func (v NullableMyTradesResponseResultInner) Get() *MyTradesResponseResultInner {
	return v.value
}

func (v *NullableMyTradesResponseResultInner) Set(val *MyTradesResponseResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMyTradesResponseResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMyTradesResponseResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyTradesResponseResultInner(val *MyTradesResponseResultInner) *NullableMyTradesResponseResultInner {
	return &NullableMyTradesResponseResultInner{value: val, isSet: true}
}

func (v NullableMyTradesResponseResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyTradesResponseResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
