/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner{}

// QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner struct for QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner
type QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner struct {
	// symbol.
	Symbol *string `json:"symbol,omitempty"`
	// order Id. Returns `null` if the working order is not filled yet.
	OrderId *int64 `json:"orderId,omitempty"`
	// status.
	Status *string `json:"status,omitempty"`
	// client Order Id.
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner

// NewQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner instantiates a new QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner() *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner {
	this := QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner{}
	return &this
}

// NewQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInnerWithDefaults instantiates a new QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInnerWithDefaults() *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner {
	this := QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given NullableInt64 and assigns it to the OrderId field.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) SetStatus(v string) {
	o.Status = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

func (o QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) UnmarshalJSON(data []byte) (err error) {
	varQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner := _QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner{}

	err = json.Unmarshal(data, &varQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner)

	if err != nil {
		return err
	}

	*o = QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner(varQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "clientOrderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner struct {
	value *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner
	isSet bool
}

func (v NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) Get() *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner {
	return v.value
}

func (v *NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) Set(val *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner(val *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) *NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner {
	return &NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner{value: val, isSet: true}
}

func (v NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
