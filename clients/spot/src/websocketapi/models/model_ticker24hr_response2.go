/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the Ticker24hrResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Ticker24hrResponse2{}

// Ticker24hrResponse2 struct for Ticker24hrResponse2
type Ticker24hrResponse2 struct {
	Id                   *string                          `json:"id,omitempty"`
	Status               *int64                           `json:"status,omitempty"`
	Result               []Ticker24hrResponse2ResultInner `json:"result,omitempty"`
	RateLimits           []RateLimits                     `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Ticker24hrResponse2 Ticker24hrResponse2

// NewTicker24hrResponse2 instantiates a new Ticker24hrResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicker24hrResponse2() *Ticker24hrResponse2 {
	this := Ticker24hrResponse2{}
	return &this
}

// NewTicker24hrResponse2WithDefaults instantiates a new Ticker24hrResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicker24hrResponse2WithDefaults() *Ticker24hrResponse2 {
	this := Ticker24hrResponse2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Ticker24hrResponse2) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse2) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ticker24hrResponse2) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ticker24hrResponse2) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Ticker24hrResponse2) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse2) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Ticker24hrResponse2) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *Ticker24hrResponse2) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *Ticker24hrResponse2) GetResult() []Ticker24hrResponse2ResultInner {
	if o == nil || common.IsNil(o.Result) {
		var ret []Ticker24hrResponse2ResultInner
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse2) GetResultOk() ([]Ticker24hrResponse2ResultInner, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *Ticker24hrResponse2) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []Ticker24hrResponse2ResultInner and assigns it to the Result field.
func (o *Ticker24hrResponse2) SetResult(v []Ticker24hrResponse2ResultInner) {
	o.Result = v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *Ticker24hrResponse2) GetRateLimits() []RateLimits {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []RateLimits
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24hrResponse2) GetRateLimitsOk() ([]RateLimits, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *Ticker24hrResponse2) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []RateLimits and assigns it to the RateLimits field.
func (o *Ticker24hrResponse2) SetRateLimits(v []RateLimits) {
	o.RateLimits = v
}

func (o Ticker24hrResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ticker24hrResponse2) ToMap() (map[string]interface{}, error) {
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

func (o *Ticker24hrResponse2) UnmarshalJSON(data []byte) (err error) {
	varTicker24hrResponse2 := _Ticker24hrResponse2{}

	err = json.Unmarshal(data, &varTicker24hrResponse2)

	if err != nil {
		return err
	}

	*o = Ticker24hrResponse2(varTicker24hrResponse2)

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

type NullableTicker24hrResponse2 struct {
	value *Ticker24hrResponse2
	isSet bool
}

func (v NullableTicker24hrResponse2) Get() *Ticker24hrResponse2 {
	return v.value
}

func (v *NullableTicker24hrResponse2) Set(val *Ticker24hrResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableTicker24hrResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableTicker24hrResponse2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicker24hrResponse2(val *Ticker24hrResponse2) *NullableTicker24hrResponse2 {
	return &NullableTicker24hrResponse2{value: val, isSet: true}
}

func (v NullableTicker24hrResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicker24hrResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
