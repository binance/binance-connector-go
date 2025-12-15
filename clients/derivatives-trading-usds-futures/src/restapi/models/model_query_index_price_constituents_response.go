/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryIndexPriceConstituentsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryIndexPriceConstituentsResponse{}

// QueryIndexPriceConstituentsResponse struct for QueryIndexPriceConstituentsResponse
type QueryIndexPriceConstituentsResponse struct {
	Symbol               *string                                                `json:"symbol,omitempty"`
	Time                 *int64                                                 `json:"time,omitempty"`
	Constituents         []QueryIndexPriceConstituentsResponseConstituentsInner `json:"constituents,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryIndexPriceConstituentsResponse QueryIndexPriceConstituentsResponse

// NewQueryIndexPriceConstituentsResponse instantiates a new QueryIndexPriceConstituentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryIndexPriceConstituentsResponse() *QueryIndexPriceConstituentsResponse {
	this := QueryIndexPriceConstituentsResponse{}
	return &this
}

// NewQueryIndexPriceConstituentsResponseWithDefaults instantiates a new QueryIndexPriceConstituentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryIndexPriceConstituentsResponseWithDefaults() *QueryIndexPriceConstituentsResponse {
	this := QueryIndexPriceConstituentsResponse{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryIndexPriceConstituentsResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIndexPriceConstituentsResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryIndexPriceConstituentsResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryIndexPriceConstituentsResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *QueryIndexPriceConstituentsResponse) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIndexPriceConstituentsResponse) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *QueryIndexPriceConstituentsResponse) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *QueryIndexPriceConstituentsResponse) SetTime(v int64) {
	o.Time = &v
}

// GetConstituents returns the Constituents field value if set, zero value otherwise.
func (o *QueryIndexPriceConstituentsResponse) GetConstituents() []QueryIndexPriceConstituentsResponseConstituentsInner {
	if o == nil || common.IsNil(o.Constituents) {
		var ret []QueryIndexPriceConstituentsResponseConstituentsInner
		return ret
	}
	return o.Constituents
}

// GetConstituentsOk returns a tuple with the Constituents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIndexPriceConstituentsResponse) GetConstituentsOk() ([]QueryIndexPriceConstituentsResponseConstituentsInner, bool) {
	if o == nil || common.IsNil(o.Constituents) {
		return nil, false
	}
	return o.Constituents, true
}

// HasConstituents returns a boolean if a field has been set.
func (o *QueryIndexPriceConstituentsResponse) HasConstituents() bool {
	if o != nil && !common.IsNil(o.Constituents) {
		return true
	}

	return false
}

// SetConstituents gets a reference to the given []QueryIndexPriceConstituentsResponseConstituentsInner and assigns it to the Constituents field.
func (o *QueryIndexPriceConstituentsResponse) SetConstituents(v []QueryIndexPriceConstituentsResponseConstituentsInner) {
	o.Constituents = v
}

func (o QueryIndexPriceConstituentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryIndexPriceConstituentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.Constituents) {
		toSerialize["constituents"] = o.Constituents
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryIndexPriceConstituentsResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryIndexPriceConstituentsResponse := _QueryIndexPriceConstituentsResponse{}

	err = json.Unmarshal(data, &varQueryIndexPriceConstituentsResponse)

	if err != nil {
		return err
	}

	*o = QueryIndexPriceConstituentsResponse(varQueryIndexPriceConstituentsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "time")
		delete(additionalProperties, "constituents")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryIndexPriceConstituentsResponse struct {
	value *QueryIndexPriceConstituentsResponse
	isSet bool
}

func (v NullableQueryIndexPriceConstituentsResponse) Get() *QueryIndexPriceConstituentsResponse {
	return v.value
}

func (v *NullableQueryIndexPriceConstituentsResponse) Set(val *QueryIndexPriceConstituentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryIndexPriceConstituentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryIndexPriceConstituentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryIndexPriceConstituentsResponse(val *QueryIndexPriceConstituentsResponse) *NullableQueryIndexPriceConstituentsResponse {
	return &NullableQueryIndexPriceConstituentsResponse{value: val, isSet: true}
}

func (v NullableQueryIndexPriceConstituentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryIndexPriceConstituentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
