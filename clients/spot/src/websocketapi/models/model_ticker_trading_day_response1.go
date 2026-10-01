/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TickerTradingDayResponse1 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TickerTradingDayResponse1{}

// TickerTradingDayResponse1 struct for TickerTradingDayResponse1
type TickerTradingDayResponse1 struct {
	Id                   *string                                  `json:"id,omitempty"`
	Status               *int64                                   `json:"status,omitempty"`
	Result               *TickerTradingDayResponse1Result         `json:"result,omitempty"`
	RateLimits           []OrderAmendmentsResponseRateLimitsInner `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TickerTradingDayResponse1 TickerTradingDayResponse1

// NewTickerTradingDayResponse1 instantiates a new TickerTradingDayResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTickerTradingDayResponse1() *TickerTradingDayResponse1 {
	this := TickerTradingDayResponse1{}
	return &this
}

// NewTickerTradingDayResponse1WithDefaults instantiates a new TickerTradingDayResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTickerTradingDayResponse1WithDefaults() *TickerTradingDayResponse1 {
	this := TickerTradingDayResponse1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TickerTradingDayResponse1) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerTradingDayResponse1) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TickerTradingDayResponse1) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TickerTradingDayResponse1) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TickerTradingDayResponse1) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerTradingDayResponse1) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TickerTradingDayResponse1) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *TickerTradingDayResponse1) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *TickerTradingDayResponse1) GetResult() TickerTradingDayResponse1Result {
	if o == nil || common.IsNil(o.Result) {
		var ret TickerTradingDayResponse1Result
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerTradingDayResponse1) GetResultOk() (*TickerTradingDayResponse1Result, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *TickerTradingDayResponse1) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given TickerTradingDayResponse1Result and assigns it to the Result field.
func (o *TickerTradingDayResponse1) SetResult(v TickerTradingDayResponse1Result) {
	o.Result = &v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *TickerTradingDayResponse1) GetRateLimits() []OrderAmendmentsResponseRateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []OrderAmendmentsResponseRateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerTradingDayResponse1) GetRateLimitsOk() ([]OrderAmendmentsResponseRateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *TickerTradingDayResponse1) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []OrderAmendmentsResponseRateLimitsInner and assigns it to the RateLimits field.
func (o *TickerTradingDayResponse1) SetRateLimits(v []OrderAmendmentsResponseRateLimitsInner) {
	o.RateLimits = v
}

func (o TickerTradingDayResponse1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TickerTradingDayResponse1) ToMap() (map[string]interface{}, error) {
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

func (o *TickerTradingDayResponse1) UnmarshalJSON(data []byte) (err error) {
	varTickerTradingDayResponse1 := _TickerTradingDayResponse1{}

	err = json.Unmarshal(data, &varTickerTradingDayResponse1)

	if err != nil {
		return err
	}

	*o = TickerTradingDayResponse1(varTickerTradingDayResponse1)

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

type NullableTickerTradingDayResponse1 struct {
	value *TickerTradingDayResponse1
	isSet bool
}

func (v NullableTickerTradingDayResponse1) Get() *TickerTradingDayResponse1 {
	return v.value
}

func (v *NullableTickerTradingDayResponse1) Set(val *TickerTradingDayResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerTradingDayResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerTradingDayResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerTradingDayResponse1(val *TickerTradingDayResponse1) *NullableTickerTradingDayResponse1 {
	return &NullableTickerTradingDayResponse1{value: val, isSet: true}
}

func (v NullableTickerTradingDayResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerTradingDayResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
