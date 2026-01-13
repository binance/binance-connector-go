/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SymbolPriceTickerResponse1 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolPriceTickerResponse1{}

// SymbolPriceTickerResponse1 struct for SymbolPriceTickerResponse1
type SymbolPriceTickerResponse1 struct {
	Id                   *string                                         `json:"id,omitempty"`
	Status               *int64                                          `json:"status,omitempty"`
	Result               *SymbolPriceTickerResponse1Result               `json:"result,omitempty"`
	RateLimits           []SymbolOrderBookTickerResponse1RateLimitsInner `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SymbolPriceTickerResponse1 SymbolPriceTickerResponse1

// NewSymbolPriceTickerResponse1 instantiates a new SymbolPriceTickerResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolPriceTickerResponse1() *SymbolPriceTickerResponse1 {
	this := SymbolPriceTickerResponse1{}
	return &this
}

// NewSymbolPriceTickerResponse1WithDefaults instantiates a new SymbolPriceTickerResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolPriceTickerResponse1WithDefaults() *SymbolPriceTickerResponse1 {
	this := SymbolPriceTickerResponse1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponse1) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponse1) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponse1) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SymbolPriceTickerResponse1) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponse1) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponse1) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponse1) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *SymbolPriceTickerResponse1) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponse1) GetResult() SymbolPriceTickerResponse1Result {
	if o == nil || common.IsNil(o.Result) {
		var ret SymbolPriceTickerResponse1Result
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponse1) GetResultOk() (*SymbolPriceTickerResponse1Result, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponse1) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given SymbolPriceTickerResponse1Result and assigns it to the Result field.
func (o *SymbolPriceTickerResponse1) SetResult(v SymbolPriceTickerResponse1Result) {
	o.Result = &v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponse1) GetRateLimits() []SymbolOrderBookTickerResponse1RateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []SymbolOrderBookTickerResponse1RateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponse1) GetRateLimitsOk() ([]SymbolOrderBookTickerResponse1RateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponse1) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []SymbolOrderBookTickerResponse1RateLimitsInner and assigns it to the RateLimits field.
func (o *SymbolPriceTickerResponse1) SetRateLimits(v []SymbolOrderBookTickerResponse1RateLimitsInner) {
	o.RateLimits = v
}

func (o SymbolPriceTickerResponse1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolPriceTickerResponse1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !common.IsNil(o.RateLimits) {
		toSerialize["rateLimits"] = o.RateLimits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SymbolPriceTickerResponse1) UnmarshalJSON(data []byte) (err error) {
	varSymbolPriceTickerResponse1 := _SymbolPriceTickerResponse1{}

	err = json.Unmarshal(data, &varSymbolPriceTickerResponse1)

	if err != nil {
		return err
	}

	*o = SymbolPriceTickerResponse1(varSymbolPriceTickerResponse1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "result")
		delete(additionalProperties, "rateLimits")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSymbolPriceTickerResponse1 struct {
	value *SymbolPriceTickerResponse1
	isSet bool
}

func (v NullableSymbolPriceTickerResponse1) Get() *SymbolPriceTickerResponse1 {
	return v.value
}

func (v *NullableSymbolPriceTickerResponse1) Set(val *SymbolPriceTickerResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolPriceTickerResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolPriceTickerResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolPriceTickerResponse1(val *SymbolPriceTickerResponse1) *NullableSymbolPriceTickerResponse1 {
	return &NullableSymbolPriceTickerResponse1{value: val, isSet: true}
}

func (v NullableSymbolPriceTickerResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolPriceTickerResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
