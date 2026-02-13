/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the SymbolOrderBookTickerResponse1 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolOrderBookTickerResponse1{}

// SymbolOrderBookTickerResponse1 struct for SymbolOrderBookTickerResponse1
type SymbolOrderBookTickerResponse1 struct {
	Id                   *string                                         `json:"id,omitempty"`
	Status               *int64                                          `json:"status,omitempty"`
	Result               *SymbolOrderBookTickerResponse1Result           `json:"result,omitempty"`
	RateLimits           []SymbolOrderBookTickerResponse1RateLimitsInner `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SymbolOrderBookTickerResponse1 SymbolOrderBookTickerResponse1

// NewSymbolOrderBookTickerResponse1 instantiates a new SymbolOrderBookTickerResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolOrderBookTickerResponse1() *SymbolOrderBookTickerResponse1 {
	this := SymbolOrderBookTickerResponse1{}
	return &this
}

// NewSymbolOrderBookTickerResponse1WithDefaults instantiates a new SymbolOrderBookTickerResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolOrderBookTickerResponse1WithDefaults() *SymbolOrderBookTickerResponse1 {
	this := SymbolOrderBookTickerResponse1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse1) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse1) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse1) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SymbolOrderBookTickerResponse1) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse1) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse1) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse1) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *SymbolOrderBookTickerResponse1) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse1) GetResult() SymbolOrderBookTickerResponse1Result {
	if o == nil || common.IsNil(o.Result) {
		var ret SymbolOrderBookTickerResponse1Result
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse1) GetResultOk() (*SymbolOrderBookTickerResponse1Result, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse1) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given SymbolOrderBookTickerResponse1Result and assigns it to the Result field.
func (o *SymbolOrderBookTickerResponse1) SetResult(v SymbolOrderBookTickerResponse1Result) {
	o.Result = &v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse1) GetRateLimits() []SymbolOrderBookTickerResponse1RateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []SymbolOrderBookTickerResponse1RateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse1) GetRateLimitsOk() ([]SymbolOrderBookTickerResponse1RateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse1) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []SymbolOrderBookTickerResponse1RateLimitsInner and assigns it to the RateLimits field.
func (o *SymbolOrderBookTickerResponse1) SetRateLimits(v []SymbolOrderBookTickerResponse1RateLimitsInner) {
	o.RateLimits = v
}

func (o SymbolOrderBookTickerResponse1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolOrderBookTickerResponse1) ToMap() (map[string]interface{}, error) {
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

func (o *SymbolOrderBookTickerResponse1) UnmarshalJSON(data []byte) (err error) {
	varSymbolOrderBookTickerResponse1 := _SymbolOrderBookTickerResponse1{}

	err = json.Unmarshal(data, &varSymbolOrderBookTickerResponse1)

	if err != nil {
		return err
	}

	*o = SymbolOrderBookTickerResponse1(varSymbolOrderBookTickerResponse1)

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

type NullableSymbolOrderBookTickerResponse1 struct {
	value *SymbolOrderBookTickerResponse1
	isSet bool
}

func (v NullableSymbolOrderBookTickerResponse1) Get() *SymbolOrderBookTickerResponse1 {
	return v.value
}

func (v *NullableSymbolOrderBookTickerResponse1) Set(val *SymbolOrderBookTickerResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolOrderBookTickerResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolOrderBookTickerResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolOrderBookTickerResponse1(val *SymbolOrderBookTickerResponse1) *NullableSymbolOrderBookTickerResponse1 {
	return &NullableSymbolOrderBookTickerResponse1{value: val, isSet: true}
}

func (v NullableSymbolOrderBookTickerResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolOrderBookTickerResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
