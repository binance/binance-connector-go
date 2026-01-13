/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetOrderModifyHistoryResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOrderModifyHistoryResponseInner{}

// GetOrderModifyHistoryResponseInner struct for GetOrderModifyHistoryResponseInner
type GetOrderModifyHistoryResponseInner struct {
	AmendmentId          *int64                                       `json:"amendmentId,omitempty"`
	Symbol               *string                                      `json:"symbol,omitempty"`
	Pair                 *string                                      `json:"pair,omitempty"`
	OrderId              *int64                                       `json:"orderId,omitempty"`
	ClientOrderId        *string                                      `json:"clientOrderId,omitempty"`
	Time                 *int64                                       `json:"time,omitempty"`
	Amendment            *GetOrderModifyHistoryResponseInnerAmendment `json:"amendment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOrderModifyHistoryResponseInner GetOrderModifyHistoryResponseInner

// NewGetOrderModifyHistoryResponseInner instantiates a new GetOrderModifyHistoryResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderModifyHistoryResponseInner() *GetOrderModifyHistoryResponseInner {
	this := GetOrderModifyHistoryResponseInner{}
	return &this
}

// NewGetOrderModifyHistoryResponseInnerWithDefaults instantiates a new GetOrderModifyHistoryResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderModifyHistoryResponseInnerWithDefaults() *GetOrderModifyHistoryResponseInner {
	this := GetOrderModifyHistoryResponseInner{}
	return &this
}

// GetAmendmentId returns the AmendmentId field value if set, zero value otherwise.
func (o *GetOrderModifyHistoryResponseInner) GetAmendmentId() int64 {
	if o == nil || common.IsNil(o.AmendmentId) {
		var ret int64
		return ret
	}
	return *o.AmendmentId
}

// GetAmendmentIdOk returns a tuple with the AmendmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderModifyHistoryResponseInner) GetAmendmentIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AmendmentId) {
		return nil, false
	}
	return o.AmendmentId, true
}

// HasAmendmentId returns a boolean if a field has been set.
func (o *GetOrderModifyHistoryResponseInner) HasAmendmentId() bool {
	if o != nil && !common.IsNil(o.AmendmentId) {
		return true
	}

	return false
}

// SetAmendmentId gets a reference to the given int64 and assigns it to the AmendmentId field.
func (o *GetOrderModifyHistoryResponseInner) SetAmendmentId(v int64) {
	o.AmendmentId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetOrderModifyHistoryResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderModifyHistoryResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetOrderModifyHistoryResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetOrderModifyHistoryResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *GetOrderModifyHistoryResponseInner) GetPair() string {
	if o == nil || common.IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderModifyHistoryResponseInner) GetPairOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *GetOrderModifyHistoryResponseInner) HasPair() bool {
	if o != nil && !common.IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *GetOrderModifyHistoryResponseInner) SetPair(v string) {
	o.Pair = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetOrderModifyHistoryResponseInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderModifyHistoryResponseInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetOrderModifyHistoryResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetOrderModifyHistoryResponseInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *GetOrderModifyHistoryResponseInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderModifyHistoryResponseInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *GetOrderModifyHistoryResponseInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *GetOrderModifyHistoryResponseInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetOrderModifyHistoryResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderModifyHistoryResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetOrderModifyHistoryResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetOrderModifyHistoryResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetAmendment returns the Amendment field value if set, zero value otherwise.
func (o *GetOrderModifyHistoryResponseInner) GetAmendment() GetOrderModifyHistoryResponseInnerAmendment {
	if o == nil || common.IsNil(o.Amendment) {
		var ret GetOrderModifyHistoryResponseInnerAmendment
		return ret
	}
	return *o.Amendment
}

// GetAmendmentOk returns a tuple with the Amendment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderModifyHistoryResponseInner) GetAmendmentOk() (*GetOrderModifyHistoryResponseInnerAmendment, bool) {
	if o == nil || common.IsNil(o.Amendment) {
		return nil, false
	}
	return o.Amendment, true
}

// HasAmendment returns a boolean if a field has been set.
func (o *GetOrderModifyHistoryResponseInner) HasAmendment() bool {
	if o != nil && !common.IsNil(o.Amendment) {
		return true
	}

	return false
}

// SetAmendment gets a reference to the given GetOrderModifyHistoryResponseInnerAmendment and assigns it to the Amendment field.
func (o *GetOrderModifyHistoryResponseInner) SetAmendment(v GetOrderModifyHistoryResponseInnerAmendment) {
	o.Amendment = &v
}

func (o GetOrderModifyHistoryResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrderModifyHistoryResponseInner) ToMap() (map[string]interface{}, error) {
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

func (o *GetOrderModifyHistoryResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetOrderModifyHistoryResponseInner := _GetOrderModifyHistoryResponseInner{}

	err = json.Unmarshal(data, &varGetOrderModifyHistoryResponseInner)

	if err != nil {
		return err
	}

	*o = GetOrderModifyHistoryResponseInner(varGetOrderModifyHistoryResponseInner)

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

type NullableGetOrderModifyHistoryResponseInner struct {
	value *GetOrderModifyHistoryResponseInner
	isSet bool
}

func (v NullableGetOrderModifyHistoryResponseInner) Get() *GetOrderModifyHistoryResponseInner {
	return v.value
}

func (v *NullableGetOrderModifyHistoryResponseInner) Set(val *GetOrderModifyHistoryResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderModifyHistoryResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderModifyHistoryResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderModifyHistoryResponseInner(val *GetOrderModifyHistoryResponseInner) *NullableGetOrderModifyHistoryResponseInner {
	return &NullableGetOrderModifyHistoryResponseInner{value: val, isSet: true}
}

func (v NullableGetOrderModifyHistoryResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderModifyHistoryResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
