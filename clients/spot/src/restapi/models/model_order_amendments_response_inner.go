/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OrderAmendmentsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderAmendmentsResponseInner{}

// OrderAmendmentsResponseInner struct for OrderAmendmentsResponseInner
type OrderAmendmentsResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	ExecutionId          *int64  `json:"executionId,omitempty"`
	OrigClientOrderId    *string `json:"origClientOrderId,omitempty"`
	NewClientOrderId     *string `json:"newClientOrderId,omitempty"`
	OrigQty              *string `json:"origQty,omitempty"`
	NewQty               *string `json:"newQty,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderAmendmentsResponseInner OrderAmendmentsResponseInner

// NewOrderAmendmentsResponseInner instantiates a new OrderAmendmentsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmendmentsResponseInner() *OrderAmendmentsResponseInner {
	this := OrderAmendmentsResponseInner{}
	return &this
}

// NewOrderAmendmentsResponseInnerWithDefaults instantiates a new OrderAmendmentsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmendmentsResponseInnerWithDefaults() *OrderAmendmentsResponseInner {
	this := OrderAmendmentsResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OrderAmendmentsResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendmentsResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OrderAmendmentsResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OrderAmendmentsResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderAmendmentsResponseInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendmentsResponseInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderAmendmentsResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *OrderAmendmentsResponseInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *OrderAmendmentsResponseInner) GetExecutionId() int64 {
	if o == nil || common.IsNil(o.ExecutionId) {
		var ret int64
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendmentsResponseInner) GetExecutionIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExecutionId) {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *OrderAmendmentsResponseInner) HasExecutionId() bool {
	if o != nil && !common.IsNil(o.ExecutionId) {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given int64 and assigns it to the ExecutionId field.
func (o *OrderAmendmentsResponseInner) SetExecutionId(v int64) {
	o.ExecutionId = &v
}

// GetOrigClientOrderId returns the OrigClientOrderId field value if set, zero value otherwise.
func (o *OrderAmendmentsResponseInner) GetOrigClientOrderId() string {
	if o == nil || common.IsNil(o.OrigClientOrderId) {
		var ret string
		return ret
	}
	return *o.OrigClientOrderId
}

// GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendmentsResponseInner) GetOrigClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigClientOrderId) {
		return nil, false
	}
	return o.OrigClientOrderId, true
}

// HasOrigClientOrderId returns a boolean if a field has been set.
func (o *OrderAmendmentsResponseInner) HasOrigClientOrderId() bool {
	if o != nil && !common.IsNil(o.OrigClientOrderId) {
		return true
	}

	return false
}

// SetOrigClientOrderId gets a reference to the given string and assigns it to the OrigClientOrderId field.
func (o *OrderAmendmentsResponseInner) SetOrigClientOrderId(v string) {
	o.OrigClientOrderId = &v
}

// GetNewClientOrderId returns the NewClientOrderId field value if set, zero value otherwise.
func (o *OrderAmendmentsResponseInner) GetNewClientOrderId() string {
	if o == nil || common.IsNil(o.NewClientOrderId) {
		var ret string
		return ret
	}
	return *o.NewClientOrderId
}

// GetNewClientOrderIdOk returns a tuple with the NewClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendmentsResponseInner) GetNewClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.NewClientOrderId) {
		return nil, false
	}
	return o.NewClientOrderId, true
}

// HasNewClientOrderId returns a boolean if a field has been set.
func (o *OrderAmendmentsResponseInner) HasNewClientOrderId() bool {
	if o != nil && !common.IsNil(o.NewClientOrderId) {
		return true
	}

	return false
}

// SetNewClientOrderId gets a reference to the given string and assigns it to the NewClientOrderId field.
func (o *OrderAmendmentsResponseInner) SetNewClientOrderId(v string) {
	o.NewClientOrderId = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *OrderAmendmentsResponseInner) GetOrigQty() string {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendmentsResponseInner) GetOrigQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *OrderAmendmentsResponseInner) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *OrderAmendmentsResponseInner) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetNewQty returns the NewQty field value if set, zero value otherwise.
func (o *OrderAmendmentsResponseInner) GetNewQty() string {
	if o == nil || common.IsNil(o.NewQty) {
		var ret string
		return ret
	}
	return *o.NewQty
}

// GetNewQtyOk returns a tuple with the NewQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendmentsResponseInner) GetNewQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.NewQty) {
		return nil, false
	}
	return o.NewQty, true
}

// HasNewQty returns a boolean if a field has been set.
func (o *OrderAmendmentsResponseInner) HasNewQty() bool {
	if o != nil && !common.IsNil(o.NewQty) {
		return true
	}

	return false
}

// SetNewQty gets a reference to the given string and assigns it to the NewQty field.
func (o *OrderAmendmentsResponseInner) SetNewQty(v string) {
	o.NewQty = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *OrderAmendmentsResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendmentsResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *OrderAmendmentsResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *OrderAmendmentsResponseInner) SetTime(v int64) {
	o.Time = &v
}

func (o OrderAmendmentsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderAmendmentsResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.ExecutionId) {
		toSerialize["executionId"] = o.ExecutionId
	}
	if !common.IsNil(o.OrigClientOrderId) {
		toSerialize["origClientOrderId"] = o.OrigClientOrderId
	}
	if !common.IsNil(o.NewClientOrderId) {
		toSerialize["newClientOrderId"] = o.NewClientOrderId
	}
	if !common.IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !common.IsNil(o.NewQty) {
		toSerialize["newQty"] = o.NewQty
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderAmendmentsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varOrderAmendmentsResponseInner := _OrderAmendmentsResponseInner{}

	err = json.Unmarshal(data, &varOrderAmendmentsResponseInner)

	if err != nil {
		return err
	}

	*o = OrderAmendmentsResponseInner(varOrderAmendmentsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "executionId")
		delete(additionalProperties, "origClientOrderId")
		delete(additionalProperties, "newClientOrderId")
		delete(additionalProperties, "origQty")
		delete(additionalProperties, "newQty")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderAmendmentsResponseInner struct {
	value *OrderAmendmentsResponseInner
	isSet bool
}

func (v NullableOrderAmendmentsResponseInner) Get() *OrderAmendmentsResponseInner {
	return v.value
}

func (v *NullableOrderAmendmentsResponseInner) Set(val *OrderAmendmentsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmendmentsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmendmentsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmendmentsResponseInner(val *OrderAmendmentsResponseInner) *NullableOrderAmendmentsResponseInner {
	return &NullableOrderAmendmentsResponseInner{value: val, isSet: true}
}

func (v NullableOrderAmendmentsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmendmentsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
