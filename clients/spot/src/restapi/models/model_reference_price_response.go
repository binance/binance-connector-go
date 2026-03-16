/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ReferencePriceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReferencePriceResponse{}

// ReferencePriceResponse struct for ReferencePriceResponse
type ReferencePriceResponse struct {
	Symbol               *string `json:"symbol,omitempty"`
	ReferencePrice       *string `json:"referencePrice,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReferencePriceResponse ReferencePriceResponse

// NewReferencePriceResponse instantiates a new ReferencePriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferencePriceResponse() *ReferencePriceResponse {
	this := ReferencePriceResponse{}
	return &this
}

// NewReferencePriceResponseWithDefaults instantiates a new ReferencePriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferencePriceResponseWithDefaults() *ReferencePriceResponse {
	this := ReferencePriceResponse{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ReferencePriceResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ReferencePriceResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ReferencePriceResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetReferencePrice returns the ReferencePrice field value if set, zero value otherwise.
func (o *ReferencePriceResponse) GetReferencePrice() string {
	if o == nil || common.IsNil(o.ReferencePrice) {
		var ret string
		return ret
	}
	return *o.ReferencePrice
}

// GetReferencePriceOk returns a tuple with the ReferencePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse) GetReferencePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReferencePrice) {
		return nil, false
	}
	return o.ReferencePrice, true
}

// HasReferencePrice returns a boolean if a field has been set.
func (o *ReferencePriceResponse) HasReferencePrice() bool {
	if o != nil && !common.IsNil(o.ReferencePrice) {
		return true
	}

	return false
}

// SetReferencePrice gets a reference to the given string and assigns it to the ReferencePrice field.
func (o *ReferencePriceResponse) SetReferencePrice(v string) {
	o.ReferencePrice = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ReferencePriceResponse) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ReferencePriceResponse) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ReferencePriceResponse) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o ReferencePriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferencePriceResponse) ToMap() (map[string]interface{}, error) {
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

func (o *ReferencePriceResponse) UnmarshalJSON(data []byte) (err error) {
	varReferencePriceResponse := _ReferencePriceResponse{}

	err = json.Unmarshal(data, &varReferencePriceResponse)

	if err != nil {
		return err
	}

	*o = ReferencePriceResponse(varReferencePriceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "referencePrice")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReferencePriceResponse struct {
	value *ReferencePriceResponse
	isSet bool
}

func (v NullableReferencePriceResponse) Get() *ReferencePriceResponse {
	return v.value
}

func (v *NullableReferencePriceResponse) Set(val *ReferencePriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReferencePriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReferencePriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferencePriceResponse(val *ReferencePriceResponse) *NullableReferencePriceResponse {
	return &NullableReferencePriceResponse{value: val, isSet: true}
}

func (v NullableReferencePriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferencePriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
