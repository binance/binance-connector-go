/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ReferencePriceResponse1Result type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReferencePriceResponse1Result{}

// ReferencePriceResponse1Result struct for ReferencePriceResponse1Result
type ReferencePriceResponse1Result struct {
	Symbol         *string `json:"symbol,omitempty"`
	ReferencePrice *string `json:"referencePrice,omitempty"`
	// Timestamp when the reference price was valid
	Timestamp            *int64  `json:"timestamp,omitempty"`
	Code                 *int64  `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReferencePriceResponse1Result ReferencePriceResponse1Result

// NewReferencePriceResponse1Result instantiates a new ReferencePriceResponse1Result object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferencePriceResponse1Result() *ReferencePriceResponse1Result {
	this := ReferencePriceResponse1Result{}
	return &this
}

// NewReferencePriceResponse1ResultWithDefaults instantiates a new ReferencePriceResponse1Result object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferencePriceResponse1ResultWithDefaults() *ReferencePriceResponse1Result {
	this := ReferencePriceResponse1Result{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ReferencePriceResponse1Result) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse1Result) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ReferencePriceResponse1Result) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ReferencePriceResponse1Result) SetSymbol(v string) {
	o.Symbol = &v
}

// GetReferencePrice returns the ReferencePrice field value if set, zero value otherwise.
func (o *ReferencePriceResponse1Result) GetReferencePrice() string {
	if o == nil || common.IsNil(o.ReferencePrice) {
		var ret string
		return ret
	}
	return *o.ReferencePrice
}

// GetReferencePriceOk returns a tuple with the ReferencePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse1Result) GetReferencePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReferencePrice) {
		return nil, false
	}
	return o.ReferencePrice, true
}

// HasReferencePrice returns a boolean if a field has been set.
func (o *ReferencePriceResponse1Result) HasReferencePrice() bool {
	if o != nil && !common.IsNil(o.ReferencePrice) {
		return true
	}

	return false
}

// SetReferencePrice gets a reference to the given string and assigns it to the ReferencePrice field.
func (o *ReferencePriceResponse1Result) SetReferencePrice(v string) {
	o.ReferencePrice = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ReferencePriceResponse1Result) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse1Result) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ReferencePriceResponse1Result) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ReferencePriceResponse1Result) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ReferencePriceResponse1Result) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse1Result) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ReferencePriceResponse1Result) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *ReferencePriceResponse1Result) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *ReferencePriceResponse1Result) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse1Result) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *ReferencePriceResponse1Result) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *ReferencePriceResponse1Result) SetMsg(v string) {
	o.Msg = &v
}

func (o ReferencePriceResponse1Result) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferencePriceResponse1Result) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.ReferencePrice) {
		toSerialize["referencePrice"] = o.ReferencePrice
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReferencePriceResponse1Result) UnmarshalJSON(data []byte) (err error) {
	varReferencePriceResponse1Result := _ReferencePriceResponse1Result{}

	err = json.Unmarshal(data, &varReferencePriceResponse1Result)

	if err != nil {
		return err
	}

	*o = ReferencePriceResponse1Result(varReferencePriceResponse1Result)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "referencePrice")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReferencePriceResponse1Result struct {
	value *ReferencePriceResponse1Result
	isSet bool
}

func (v NullableReferencePriceResponse1Result) Get() *ReferencePriceResponse1Result {
	return v.value
}

func (v *NullableReferencePriceResponse1Result) Set(val *ReferencePriceResponse1Result) {
	v.value = val
	v.isSet = true
}

func (v NullableReferencePriceResponse1Result) IsSet() bool {
	return v.isSet
}

func (v *NullableReferencePriceResponse1Result) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferencePriceResponse1Result(val *ReferencePriceResponse1Result) *NullableReferencePriceResponse1Result {
	return &NullableReferencePriceResponse1Result{value: val, isSet: true}
}

func (v NullableReferencePriceResponse1Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferencePriceResponse1Result) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
