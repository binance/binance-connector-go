/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryUmModifyOrderHistoryResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUmModifyOrderHistoryResponseInner{}

// QueryUmModifyOrderHistoryResponseInner struct for QueryUmModifyOrderHistoryResponseInner
type QueryUmModifyOrderHistoryResponseInner struct {
	AmendmentId          *int64                                           `json:"amendmentId,omitempty"`
	Symbol               *string                                          `json:"symbol,omitempty"`
	Pair                 *string                                          `json:"pair,omitempty"`
	OrderId              *int64                                           `json:"orderId,omitempty"`
	ClientOrderId        *string                                          `json:"clientOrderId,omitempty"`
	Time                 *int64                                           `json:"time,omitempty"`
	Amendment            *QueryCmModifyOrderHistoryResponseInnerAmendment `json:"amendment,omitempty"`
	PriceMatch           *string                                          `json:"priceMatch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryUmModifyOrderHistoryResponseInner QueryUmModifyOrderHistoryResponseInner

// NewQueryUmModifyOrderHistoryResponseInner instantiates a new QueryUmModifyOrderHistoryResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUmModifyOrderHistoryResponseInner() *QueryUmModifyOrderHistoryResponseInner {
	this := QueryUmModifyOrderHistoryResponseInner{}
	return &this
}

// NewQueryUmModifyOrderHistoryResponseInnerWithDefaults instantiates a new QueryUmModifyOrderHistoryResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUmModifyOrderHistoryResponseInnerWithDefaults() *QueryUmModifyOrderHistoryResponseInner {
	this := QueryUmModifyOrderHistoryResponseInner{}
	return &this
}

// GetAmendmentId returns the AmendmentId field value if set, zero value otherwise.
func (o *QueryUmModifyOrderHistoryResponseInner) GetAmendmentId() int64 {
	if o == nil || common.IsNil(o.AmendmentId) {
		var ret int64
		return ret
	}
	return *o.AmendmentId
}

// GetAmendmentIdOk returns a tuple with the AmendmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmModifyOrderHistoryResponseInner) GetAmendmentIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AmendmentId) {
		return nil, false
	}
	return o.AmendmentId, true
}

// HasAmendmentId returns a boolean if a field has been set.
func (o *QueryUmModifyOrderHistoryResponseInner) HasAmendmentId() bool {
	if o != nil && !common.IsNil(o.AmendmentId) {
		return true
	}

	return false
}

// SetAmendmentId gets a reference to the given int64 and assigns it to the AmendmentId field.
func (o *QueryUmModifyOrderHistoryResponseInner) SetAmendmentId(v int64) {
	o.AmendmentId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryUmModifyOrderHistoryResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmModifyOrderHistoryResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryUmModifyOrderHistoryResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryUmModifyOrderHistoryResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *QueryUmModifyOrderHistoryResponseInner) GetPair() string {
	if o == nil || common.IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmModifyOrderHistoryResponseInner) GetPairOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *QueryUmModifyOrderHistoryResponseInner) HasPair() bool {
	if o != nil && !common.IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *QueryUmModifyOrderHistoryResponseInner) SetPair(v string) {
	o.Pair = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *QueryUmModifyOrderHistoryResponseInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmModifyOrderHistoryResponseInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *QueryUmModifyOrderHistoryResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *QueryUmModifyOrderHistoryResponseInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *QueryUmModifyOrderHistoryResponseInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmModifyOrderHistoryResponseInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *QueryUmModifyOrderHistoryResponseInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *QueryUmModifyOrderHistoryResponseInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *QueryUmModifyOrderHistoryResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmModifyOrderHistoryResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *QueryUmModifyOrderHistoryResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *QueryUmModifyOrderHistoryResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetAmendment returns the Amendment field value if set, zero value otherwise.
func (o *QueryUmModifyOrderHistoryResponseInner) GetAmendment() QueryCmModifyOrderHistoryResponseInnerAmendment {
	if o == nil || common.IsNil(o.Amendment) {
		var ret QueryCmModifyOrderHistoryResponseInnerAmendment
		return ret
	}
	return *o.Amendment
}

// GetAmendmentOk returns a tuple with the Amendment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmModifyOrderHistoryResponseInner) GetAmendmentOk() (*QueryCmModifyOrderHistoryResponseInnerAmendment, bool) {
	if o == nil || common.IsNil(o.Amendment) {
		return nil, false
	}
	return o.Amendment, true
}

// HasAmendment returns a boolean if a field has been set.
func (o *QueryUmModifyOrderHistoryResponseInner) HasAmendment() bool {
	if o != nil && !common.IsNil(o.Amendment) {
		return true
	}

	return false
}

// SetAmendment gets a reference to the given QueryCmModifyOrderHistoryResponseInnerAmendment and assigns it to the Amendment field.
func (o *QueryUmModifyOrderHistoryResponseInner) SetAmendment(v QueryCmModifyOrderHistoryResponseInnerAmendment) {
	o.Amendment = &v
}

// GetPriceMatch returns the PriceMatch field value if set, zero value otherwise.
func (o *QueryUmModifyOrderHistoryResponseInner) GetPriceMatch() string {
	if o == nil || common.IsNil(o.PriceMatch) {
		var ret string
		return ret
	}
	return *o.PriceMatch
}

// GetPriceMatchOk returns a tuple with the PriceMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmModifyOrderHistoryResponseInner) GetPriceMatchOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceMatch) {
		return nil, false
	}
	return o.PriceMatch, true
}

// HasPriceMatch returns a boolean if a field has been set.
func (o *QueryUmModifyOrderHistoryResponseInner) HasPriceMatch() bool {
	if o != nil && !common.IsNil(o.PriceMatch) {
		return true
	}

	return false
}

// SetPriceMatch gets a reference to the given string and assigns it to the PriceMatch field.
func (o *QueryUmModifyOrderHistoryResponseInner) SetPriceMatch(v string) {
	o.PriceMatch = &v
}

func (o QueryUmModifyOrderHistoryResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUmModifyOrderHistoryResponseInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.PriceMatch) {
		toSerialize["priceMatch"] = o.PriceMatch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryUmModifyOrderHistoryResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryUmModifyOrderHistoryResponseInner := _QueryUmModifyOrderHistoryResponseInner{}

	err = json.Unmarshal(data, &varQueryUmModifyOrderHistoryResponseInner)

	if err != nil {
		return err
	}

	*o = QueryUmModifyOrderHistoryResponseInner(varQueryUmModifyOrderHistoryResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amendmentId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "pair")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "clientOrderId")
		delete(additionalProperties, "time")
		delete(additionalProperties, "amendment")
		delete(additionalProperties, "priceMatch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryUmModifyOrderHistoryResponseInner struct {
	value *QueryUmModifyOrderHistoryResponseInner
	isSet bool
}

func (v NullableQueryUmModifyOrderHistoryResponseInner) Get() *QueryUmModifyOrderHistoryResponseInner {
	return v.value
}

func (v *NullableQueryUmModifyOrderHistoryResponseInner) Set(val *QueryUmModifyOrderHistoryResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUmModifyOrderHistoryResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUmModifyOrderHistoryResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUmModifyOrderHistoryResponseInner(val *QueryUmModifyOrderHistoryResponseInner) *NullableQueryUmModifyOrderHistoryResponseInner {
	return &NullableQueryUmModifyOrderHistoryResponseInner{value: val, isSet: true}
}

func (v NullableQueryUmModifyOrderHistoryResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUmModifyOrderHistoryResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
