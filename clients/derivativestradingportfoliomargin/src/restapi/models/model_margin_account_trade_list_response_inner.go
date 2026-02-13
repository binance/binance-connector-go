/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MarginAccountTradeListResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginAccountTradeListResponseInner{}

// MarginAccountTradeListResponseInner struct for MarginAccountTradeListResponseInner
type MarginAccountTradeListResponseInner struct {
	Commission           *string `json:"commission,omitempty"`
	CommissionAsset      *string `json:"commissionAsset,omitempty"`
	Id                   *int64  `json:"id,omitempty"`
	IsBestMatch          *bool   `json:"isBestMatch,omitempty"`
	IsBuyer              *bool   `json:"isBuyer,omitempty"`
	IsMaker              *bool   `json:"isMaker,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	Price                *string `json:"price,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginAccountTradeListResponseInner MarginAccountTradeListResponseInner

// NewMarginAccountTradeListResponseInner instantiates a new MarginAccountTradeListResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginAccountTradeListResponseInner() *MarginAccountTradeListResponseInner {
	this := MarginAccountTradeListResponseInner{}
	return &this
}

// NewMarginAccountTradeListResponseInnerWithDefaults instantiates a new MarginAccountTradeListResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginAccountTradeListResponseInnerWithDefaults() *MarginAccountTradeListResponseInner {
	this := MarginAccountTradeListResponseInner{}
	return &this
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *MarginAccountTradeListResponseInner) GetCommission() string {
	if o == nil || common.IsNil(o.Commission) {
		var ret string
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountTradeListResponseInner) GetCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *MarginAccountTradeListResponseInner) HasCommission() bool {
	if o != nil && !common.IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given string and assigns it to the Commission field.
func (o *MarginAccountTradeListResponseInner) SetCommission(v string) {
	o.Commission = &v
}

// GetCommissionAsset returns the CommissionAsset field value if set, zero value otherwise.
func (o *MarginAccountTradeListResponseInner) GetCommissionAsset() string {
	if o == nil || common.IsNil(o.CommissionAsset) {
		var ret string
		return ret
	}
	return *o.CommissionAsset
}

// GetCommissionAssetOk returns a tuple with the CommissionAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountTradeListResponseInner) GetCommissionAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.CommissionAsset) {
		return nil, false
	}
	return o.CommissionAsset, true
}

// HasCommissionAsset returns a boolean if a field has been set.
func (o *MarginAccountTradeListResponseInner) HasCommissionAsset() bool {
	if o != nil && !common.IsNil(o.CommissionAsset) {
		return true
	}

	return false
}

// SetCommissionAsset gets a reference to the given string and assigns it to the CommissionAsset field.
func (o *MarginAccountTradeListResponseInner) SetCommissionAsset(v string) {
	o.CommissionAsset = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MarginAccountTradeListResponseInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountTradeListResponseInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MarginAccountTradeListResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *MarginAccountTradeListResponseInner) SetId(v int64) {
	o.Id = &v
}

// GetIsBestMatch returns the IsBestMatch field value if set, zero value otherwise.
func (o *MarginAccountTradeListResponseInner) GetIsBestMatch() bool {
	if o == nil || common.IsNil(o.IsBestMatch) {
		var ret bool
		return ret
	}
	return *o.IsBestMatch
}

// GetIsBestMatchOk returns a tuple with the IsBestMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountTradeListResponseInner) GetIsBestMatchOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBestMatch) {
		return nil, false
	}
	return o.IsBestMatch, true
}

// HasIsBestMatch returns a boolean if a field has been set.
func (o *MarginAccountTradeListResponseInner) HasIsBestMatch() bool {
	if o != nil && !common.IsNil(o.IsBestMatch) {
		return true
	}

	return false
}

// SetIsBestMatch gets a reference to the given bool and assigns it to the IsBestMatch field.
func (o *MarginAccountTradeListResponseInner) SetIsBestMatch(v bool) {
	o.IsBestMatch = &v
}

// GetIsBuyer returns the IsBuyer field value if set, zero value otherwise.
func (o *MarginAccountTradeListResponseInner) GetIsBuyer() bool {
	if o == nil || common.IsNil(o.IsBuyer) {
		var ret bool
		return ret
	}
	return *o.IsBuyer
}

// GetIsBuyerOk returns a tuple with the IsBuyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountTradeListResponseInner) GetIsBuyerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBuyer) {
		return nil, false
	}
	return o.IsBuyer, true
}

// HasIsBuyer returns a boolean if a field has been set.
func (o *MarginAccountTradeListResponseInner) HasIsBuyer() bool {
	if o != nil && !common.IsNil(o.IsBuyer) {
		return true
	}

	return false
}

// SetIsBuyer gets a reference to the given bool and assigns it to the IsBuyer field.
func (o *MarginAccountTradeListResponseInner) SetIsBuyer(v bool) {
	o.IsBuyer = &v
}

// GetIsMaker returns the IsMaker field value if set, zero value otherwise.
func (o *MarginAccountTradeListResponseInner) GetIsMaker() bool {
	if o == nil || common.IsNil(o.IsMaker) {
		var ret bool
		return ret
	}
	return *o.IsMaker
}

// GetIsMakerOk returns a tuple with the IsMaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountTradeListResponseInner) GetIsMakerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsMaker) {
		return nil, false
	}
	return o.IsMaker, true
}

// HasIsMaker returns a boolean if a field has been set.
func (o *MarginAccountTradeListResponseInner) HasIsMaker() bool {
	if o != nil && !common.IsNil(o.IsMaker) {
		return true
	}

	return false
}

// SetIsMaker gets a reference to the given bool and assigns it to the IsMaker field.
func (o *MarginAccountTradeListResponseInner) SetIsMaker(v bool) {
	o.IsMaker = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *MarginAccountTradeListResponseInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountTradeListResponseInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *MarginAccountTradeListResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *MarginAccountTradeListResponseInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *MarginAccountTradeListResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountTradeListResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *MarginAccountTradeListResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *MarginAccountTradeListResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *MarginAccountTradeListResponseInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountTradeListResponseInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *MarginAccountTradeListResponseInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *MarginAccountTradeListResponseInner) SetQty(v string) {
	o.Qty = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarginAccountTradeListResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountTradeListResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarginAccountTradeListResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarginAccountTradeListResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *MarginAccountTradeListResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountTradeListResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *MarginAccountTradeListResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *MarginAccountTradeListResponseInner) SetTime(v int64) {
	o.Time = &v
}

func (o MarginAccountTradeListResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginAccountTradeListResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Commission) {
		toSerialize["commission"] = o.Commission
	}
	if !common.IsNil(o.CommissionAsset) {
		toSerialize["commissionAsset"] = o.CommissionAsset
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.IsBestMatch) {
		toSerialize["isBestMatch"] = o.IsBestMatch
	}
	if !common.IsNil(o.IsBuyer) {
		toSerialize["isBuyer"] = o.IsBuyer
	}
	if !common.IsNil(o.IsMaker) {
		toSerialize["isMaker"] = o.IsMaker
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarginAccountTradeListResponseInner) UnmarshalJSON(data []byte) (err error) {
	varMarginAccountTradeListResponseInner := _MarginAccountTradeListResponseInner{}

	err = json.Unmarshal(data, &varMarginAccountTradeListResponseInner)

	if err != nil {
		return err
	}

	*o = MarginAccountTradeListResponseInner(varMarginAccountTradeListResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "commission")
		delete(additionalProperties, "commissionAsset")
		delete(additionalProperties, "id")
		delete(additionalProperties, "isBestMatch")
		delete(additionalProperties, "isBuyer")
		delete(additionalProperties, "isMaker")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "price")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginAccountTradeListResponseInner struct {
	value *MarginAccountTradeListResponseInner
	isSet bool
}

func (v NullableMarginAccountTradeListResponseInner) Get() *MarginAccountTradeListResponseInner {
	return v.value
}

func (v *NullableMarginAccountTradeListResponseInner) Set(val *MarginAccountTradeListResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountTradeListResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountTradeListResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountTradeListResponseInner(val *MarginAccountTradeListResponseInner) *NullableMarginAccountTradeListResponseInner {
	return &NullableMarginAccountTradeListResponseInner{value: val, isSet: true}
}

func (v NullableMarginAccountTradeListResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountTradeListResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
