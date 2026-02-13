/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OpenOrdersCancelAllResponseResultInnerOrdersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OpenOrdersCancelAllResponseResultInnerOrdersInner{}

// OpenOrdersCancelAllResponseResultInnerOrdersInner struct for OpenOrdersCancelAllResponseResultInnerOrdersInner
type OpenOrdersCancelAllResponseResultInnerOrdersInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenOrdersCancelAllResponseResultInnerOrdersInner OpenOrdersCancelAllResponseResultInnerOrdersInner

// NewOpenOrdersCancelAllResponseResultInnerOrdersInner instantiates a new OpenOrdersCancelAllResponseResultInnerOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenOrdersCancelAllResponseResultInnerOrdersInner() *OpenOrdersCancelAllResponseResultInnerOrdersInner {
	this := OpenOrdersCancelAllResponseResultInnerOrdersInner{}
	return &this
}

// NewOpenOrdersCancelAllResponseResultInnerOrdersInnerWithDefaults instantiates a new OpenOrdersCancelAllResponseResultInnerOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenOrdersCancelAllResponseResultInnerOrdersInnerWithDefaults() *OpenOrdersCancelAllResponseResultInnerOrdersInner {
	this := OpenOrdersCancelAllResponseResultInnerOrdersInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OpenOrdersCancelAllResponseResultInnerOrdersInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenOrdersCancelAllResponseResultInnerOrdersInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OpenOrdersCancelAllResponseResultInnerOrdersInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OpenOrdersCancelAllResponseResultInnerOrdersInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OpenOrdersCancelAllResponseResultInnerOrdersInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenOrdersCancelAllResponseResultInnerOrdersInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OpenOrdersCancelAllResponseResultInnerOrdersInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *OpenOrdersCancelAllResponseResultInnerOrdersInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *OpenOrdersCancelAllResponseResultInnerOrdersInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenOrdersCancelAllResponseResultInnerOrdersInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *OpenOrdersCancelAllResponseResultInnerOrdersInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *OpenOrdersCancelAllResponseResultInnerOrdersInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

func (o OpenOrdersCancelAllResponseResultInnerOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenOrdersCancelAllResponseResultInnerOrdersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenOrdersCancelAllResponseResultInnerOrdersInner) UnmarshalJSON(data []byte) (err error) {
	varOpenOrdersCancelAllResponseResultInnerOrdersInner := _OpenOrdersCancelAllResponseResultInnerOrdersInner{}

	err = json.Unmarshal(data, &varOpenOrdersCancelAllResponseResultInnerOrdersInner)

	if err != nil {
		return err
	}

	*o = OpenOrdersCancelAllResponseResultInnerOrdersInner(varOpenOrdersCancelAllResponseResultInnerOrdersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "clientOrderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenOrdersCancelAllResponseResultInnerOrdersInner struct {
	value *OpenOrdersCancelAllResponseResultInnerOrdersInner
	isSet bool
}

func (v NullableOpenOrdersCancelAllResponseResultInnerOrdersInner) Get() *OpenOrdersCancelAllResponseResultInnerOrdersInner {
	return v.value
}

func (v *NullableOpenOrdersCancelAllResponseResultInnerOrdersInner) Set(val *OpenOrdersCancelAllResponseResultInnerOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenOrdersCancelAllResponseResultInnerOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenOrdersCancelAllResponseResultInnerOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenOrdersCancelAllResponseResultInnerOrdersInner(val *OpenOrdersCancelAllResponseResultInnerOrdersInner) *NullableOpenOrdersCancelAllResponseResultInnerOrdersInner {
	return &NullableOpenOrdersCancelAllResponseResultInnerOrdersInner{value: val, isSet: true}
}

func (v NullableOpenOrdersCancelAllResponseResultInnerOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenOrdersCancelAllResponseResultInnerOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
