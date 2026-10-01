/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TickerTradingDayResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TickerTradingDayResponse2{}

// TickerTradingDayResponse2 struct for TickerTradingDayResponse2
type TickerTradingDayResponse2 struct {
	Id                   *string                                `json:"id,omitempty"`
	Status               *int64                                 `json:"status,omitempty"`
	Result               []TickerTradingDayResponse2ResultInner `json:"result,omitempty"`
	RateLimits           []TickerResponse2RateLimitsInner       `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TickerTradingDayResponse2 TickerTradingDayResponse2

// NewTickerTradingDayResponse2 instantiates a new TickerTradingDayResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTickerTradingDayResponse2() *TickerTradingDayResponse2 {
	this := TickerTradingDayResponse2{}
	return &this
}

// NewTickerTradingDayResponse2WithDefaults instantiates a new TickerTradingDayResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTickerTradingDayResponse2WithDefaults() *TickerTradingDayResponse2 {
	this := TickerTradingDayResponse2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TickerTradingDayResponse2) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerTradingDayResponse2) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TickerTradingDayResponse2) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TickerTradingDayResponse2) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TickerTradingDayResponse2) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerTradingDayResponse2) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TickerTradingDayResponse2) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *TickerTradingDayResponse2) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *TickerTradingDayResponse2) GetResult() []TickerTradingDayResponse2ResultInner {
	if o == nil || common.IsNil(o.Result) {
		var ret []TickerTradingDayResponse2ResultInner
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerTradingDayResponse2) GetResultOk() ([]TickerTradingDayResponse2ResultInner, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *TickerTradingDayResponse2) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []TickerTradingDayResponse2ResultInner and assigns it to the Result field.
func (o *TickerTradingDayResponse2) SetResult(v []TickerTradingDayResponse2ResultInner) {
	o.Result = v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *TickerTradingDayResponse2) GetRateLimits() []TickerResponse2RateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []TickerResponse2RateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerTradingDayResponse2) GetRateLimitsOk() ([]TickerResponse2RateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *TickerTradingDayResponse2) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []TickerResponse2RateLimitsInner and assigns it to the RateLimits field.
func (o *TickerTradingDayResponse2) SetRateLimits(v []TickerResponse2RateLimitsInner) {
	o.RateLimits = v
}

func (o TickerTradingDayResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TickerTradingDayResponse2) ToMap() (map[string]interface{}, error) {
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

func (o *TickerTradingDayResponse2) UnmarshalJSON(data []byte) (err error) {
	varTickerTradingDayResponse2 := _TickerTradingDayResponse2{}

	err = json.Unmarshal(data, &varTickerTradingDayResponse2)

	if err != nil {
		return err
	}

	*o = TickerTradingDayResponse2(varTickerTradingDayResponse2)

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

type NullableTickerTradingDayResponse2 struct {
	value *TickerTradingDayResponse2
	isSet bool
}

func (v NullableTickerTradingDayResponse2) Get() *TickerTradingDayResponse2 {
	return v.value
}

func (v *NullableTickerTradingDayResponse2) Set(val *TickerTradingDayResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerTradingDayResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerTradingDayResponse2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerTradingDayResponse2(val *TickerTradingDayResponse2) *NullableTickerTradingDayResponse2 {
	return &NullableTickerTradingDayResponse2{value: val, isSet: true}
}

func (v NullableTickerTradingDayResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerTradingDayResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
