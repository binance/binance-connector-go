/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryCmModifyOrderHistoryResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCmModifyOrderHistoryResponseInner{}

// QueryCmModifyOrderHistoryResponseInner struct for QueryCmModifyOrderHistoryResponseInner
type QueryCmModifyOrderHistoryResponseInner struct {
	AmendmentId          *int64                                           `json:"amendmentId,omitempty"`
	Symbol               *string                                          `json:"symbol,omitempty"`
	Pair                 *string                                          `json:"pair,omitempty"`
	OrderId              *int64                                           `json:"orderId,omitempty"`
	ClientOrderId        *string                                          `json:"clientOrderId,omitempty"`
	Time                 *int64                                           `json:"time,omitempty"`
	Amendment            *QueryCmModifyOrderHistoryResponseInnerAmendment `json:"amendment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryCmModifyOrderHistoryResponseInner QueryCmModifyOrderHistoryResponseInner

// NewQueryCmModifyOrderHistoryResponseInner instantiates a new QueryCmModifyOrderHistoryResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCmModifyOrderHistoryResponseInner() *QueryCmModifyOrderHistoryResponseInner {
	this := QueryCmModifyOrderHistoryResponseInner{}
	return &this
}

// NewQueryCmModifyOrderHistoryResponseInnerWithDefaults instantiates a new QueryCmModifyOrderHistoryResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCmModifyOrderHistoryResponseInnerWithDefaults() *QueryCmModifyOrderHistoryResponseInner {
	this := QueryCmModifyOrderHistoryResponseInner{}
	return &this
}

// GetAmendmentId returns the AmendmentId field value if set, zero value otherwise.
func (o *QueryCmModifyOrderHistoryResponseInner) GetAmendmentId() int64 {
	if o == nil || common.IsNil(o.AmendmentId) {
		var ret int64
		return ret
	}
	return *o.AmendmentId
}

// GetAmendmentIdOk returns a tuple with the AmendmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmModifyOrderHistoryResponseInner) GetAmendmentIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AmendmentId) {
		return nil, false
	}
	return o.AmendmentId, true
}

// HasAmendmentId returns a boolean if a field has been set.
func (o *QueryCmModifyOrderHistoryResponseInner) HasAmendmentId() bool {
	if o != nil && !common.IsNil(o.AmendmentId) {
		return true
	}

	return false
}

// SetAmendmentId gets a reference to the given int64 and assigns it to the AmendmentId field.
func (o *QueryCmModifyOrderHistoryResponseInner) SetAmendmentId(v int64) {
	o.AmendmentId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryCmModifyOrderHistoryResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmModifyOrderHistoryResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryCmModifyOrderHistoryResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryCmModifyOrderHistoryResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *QueryCmModifyOrderHistoryResponseInner) GetPair() string {
	if o == nil || common.IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmModifyOrderHistoryResponseInner) GetPairOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *QueryCmModifyOrderHistoryResponseInner) HasPair() bool {
	if o != nil && !common.IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *QueryCmModifyOrderHistoryResponseInner) SetPair(v string) {
	o.Pair = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *QueryCmModifyOrderHistoryResponseInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmModifyOrderHistoryResponseInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *QueryCmModifyOrderHistoryResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *QueryCmModifyOrderHistoryResponseInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *QueryCmModifyOrderHistoryResponseInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmModifyOrderHistoryResponseInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *QueryCmModifyOrderHistoryResponseInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *QueryCmModifyOrderHistoryResponseInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *QueryCmModifyOrderHistoryResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmModifyOrderHistoryResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *QueryCmModifyOrderHistoryResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *QueryCmModifyOrderHistoryResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetAmendment returns the Amendment field value if set, zero value otherwise.
func (o *QueryCmModifyOrderHistoryResponseInner) GetAmendment() QueryCmModifyOrderHistoryResponseInnerAmendment {
	if o == nil || common.IsNil(o.Amendment) {
		var ret QueryCmModifyOrderHistoryResponseInnerAmendment
		return ret
	}
	return *o.Amendment
}

// GetAmendmentOk returns a tuple with the Amendment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmModifyOrderHistoryResponseInner) GetAmendmentOk() (*QueryCmModifyOrderHistoryResponseInnerAmendment, bool) {
	if o == nil || common.IsNil(o.Amendment) {
		return nil, false
	}
	return o.Amendment, true
}

// HasAmendment returns a boolean if a field has been set.
func (o *QueryCmModifyOrderHistoryResponseInner) HasAmendment() bool {
	if o != nil && !common.IsNil(o.Amendment) {
		return true
	}

	return false
}

// SetAmendment gets a reference to the given QueryCmModifyOrderHistoryResponseInnerAmendment and assigns it to the Amendment field.
func (o *QueryCmModifyOrderHistoryResponseInner) SetAmendment(v QueryCmModifyOrderHistoryResponseInnerAmendment) {
	o.Amendment = &v
}

func (o QueryCmModifyOrderHistoryResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCmModifyOrderHistoryResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AmendmentId) {
		toSerialize["amendmentId"] = o.AmendmentId
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.Amendment) {
		toSerialize["amendment"] = o.Amendment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryCmModifyOrderHistoryResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryCmModifyOrderHistoryResponseInner := _QueryCmModifyOrderHistoryResponseInner{}

	err = json.Unmarshal(data, &varQueryCmModifyOrderHistoryResponseInner)

	if err != nil {
		return err
	}

	*o = QueryCmModifyOrderHistoryResponseInner(varQueryCmModifyOrderHistoryResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amendmentId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "pair")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "clientOrderId")
		delete(additionalProperties, "time")
		delete(additionalProperties, "amendment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryCmModifyOrderHistoryResponseInner struct {
	value *QueryCmModifyOrderHistoryResponseInner
	isSet bool
}

func (v NullableQueryCmModifyOrderHistoryResponseInner) Get() *QueryCmModifyOrderHistoryResponseInner {
	return v.value
}

func (v *NullableQueryCmModifyOrderHistoryResponseInner) Set(val *QueryCmModifyOrderHistoryResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCmModifyOrderHistoryResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCmModifyOrderHistoryResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCmModifyOrderHistoryResponseInner(val *QueryCmModifyOrderHistoryResponseInner) *NullableQueryCmModifyOrderHistoryResponseInner {
	return &NullableQueryCmModifyOrderHistoryResponseInner{value: val, isSet: true}
}

func (v NullableQueryCmModifyOrderHistoryResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCmModifyOrderHistoryResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
