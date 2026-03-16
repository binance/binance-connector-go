/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ReferencePriceResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReferencePriceResponseResult{}

// ReferencePriceResponseResult struct for ReferencePriceResponseResult
type ReferencePriceResponseResult struct {
	Symbol               *string `json:"symbol,omitempty"`
	ReferencePrice       *string `json:"referencePrice,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReferencePriceResponseResult ReferencePriceResponseResult

// NewReferencePriceResponseResult instantiates a new ReferencePriceResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferencePriceResponseResult() *ReferencePriceResponseResult {
	this := ReferencePriceResponseResult{}
	return &this
}

// NewReferencePriceResponseResultWithDefaults instantiates a new ReferencePriceResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferencePriceResponseResultWithDefaults() *ReferencePriceResponseResult {
	this := ReferencePriceResponseResult{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ReferencePriceResponseResult) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponseResult) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ReferencePriceResponseResult) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ReferencePriceResponseResult) SetSymbol(v string) {
	o.Symbol = &v
}

// GetReferencePrice returns the ReferencePrice field value if set, zero value otherwise.
func (o *ReferencePriceResponseResult) GetReferencePrice() string {
	if o == nil || common.IsNil(o.ReferencePrice) {
		var ret string
		return ret
	}
	return *o.ReferencePrice
}

// GetReferencePriceOk returns a tuple with the ReferencePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponseResult) GetReferencePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReferencePrice) {
		return nil, false
	}
	return o.ReferencePrice, true
}

// HasReferencePrice returns a boolean if a field has been set.
func (o *ReferencePriceResponseResult) HasReferencePrice() bool {
	if o != nil && !common.IsNil(o.ReferencePrice) {
		return true
	}

	return false
}

// SetReferencePrice gets a reference to the given string and assigns it to the ReferencePrice field.
func (o *ReferencePriceResponseResult) SetReferencePrice(v string) {
	o.ReferencePrice = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ReferencePriceResponseResult) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponseResult) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ReferencePriceResponseResult) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ReferencePriceResponseResult) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o ReferencePriceResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferencePriceResponseResult) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReferencePriceResponseResult) UnmarshalJSON(data []byte) (err error) {
	varReferencePriceResponseResult := _ReferencePriceResponseResult{}

	err = json.Unmarshal(data, &varReferencePriceResponseResult)

	if err != nil {
		return err
	}

	*o = ReferencePriceResponseResult(varReferencePriceResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "referencePrice")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReferencePriceResponseResult struct {
	value *ReferencePriceResponseResult
	isSet bool
}

func (v NullableReferencePriceResponseResult) Get() *ReferencePriceResponseResult {
	return v.value
}

func (v *NullableReferencePriceResponseResult) Set(val *ReferencePriceResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableReferencePriceResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableReferencePriceResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferencePriceResponseResult(val *ReferencePriceResponseResult) *NullableReferencePriceResponseResult {
	return &NullableReferencePriceResponseResult{value: val, isSet: true}
}

func (v NullableReferencePriceResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferencePriceResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
