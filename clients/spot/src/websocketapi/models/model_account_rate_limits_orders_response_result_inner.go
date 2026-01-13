/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountRateLimitsOrdersResponseResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountRateLimitsOrdersResponseResultInner{}

// AccountRateLimitsOrdersResponseResultInner struct for AccountRateLimitsOrdersResponseResultInner
type AccountRateLimitsOrdersResponseResultInner struct {
	RateLimitType        *string `json:"rateLimitType,omitempty"`
	Interval             *string `json:"interval,omitempty"`
	IntervalNum          *int64  `json:"intervalNum,omitempty"`
	Limit                *int64  `json:"limit,omitempty"`
	Count                *int64  `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountRateLimitsOrdersResponseResultInner AccountRateLimitsOrdersResponseResultInner

// NewAccountRateLimitsOrdersResponseResultInner instantiates a new AccountRateLimitsOrdersResponseResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountRateLimitsOrdersResponseResultInner() *AccountRateLimitsOrdersResponseResultInner {
	this := AccountRateLimitsOrdersResponseResultInner{}
	return &this
}

// NewAccountRateLimitsOrdersResponseResultInnerWithDefaults instantiates a new AccountRateLimitsOrdersResponseResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountRateLimitsOrdersResponseResultInnerWithDefaults() *AccountRateLimitsOrdersResponseResultInner {
	this := AccountRateLimitsOrdersResponseResultInner{}
	return &this
}

// GetRateLimitType returns the RateLimitType field value if set, zero value otherwise.
func (o *AccountRateLimitsOrdersResponseResultInner) GetRateLimitType() string {
	if o == nil || common.IsNil(o.RateLimitType) {
		var ret string
		return ret
	}
	return *o.RateLimitType
}

// GetRateLimitTypeOk returns a tuple with the RateLimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRateLimitsOrdersResponseResultInner) GetRateLimitTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RateLimitType) {
		return nil, false
	}
	return o.RateLimitType, true
}

// HasRateLimitType returns a boolean if a field has been set.
func (o *AccountRateLimitsOrdersResponseResultInner) HasRateLimitType() bool {
	if o != nil && !common.IsNil(o.RateLimitType) {
		return true
	}

	return false
}

// SetRateLimitType gets a reference to the given string and assigns it to the RateLimitType field.
func (o *AccountRateLimitsOrdersResponseResultInner) SetRateLimitType(v string) {
	o.RateLimitType = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *AccountRateLimitsOrdersResponseResultInner) GetInterval() string {
	if o == nil || common.IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRateLimitsOrdersResponseResultInner) GetIntervalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *AccountRateLimitsOrdersResponseResultInner) HasInterval() bool {
	if o != nil && !common.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *AccountRateLimitsOrdersResponseResultInner) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalNum returns the IntervalNum field value if set, zero value otherwise.
func (o *AccountRateLimitsOrdersResponseResultInner) GetIntervalNum() int64 {
	if o == nil || common.IsNil(o.IntervalNum) {
		var ret int64
		return ret
	}
	return *o.IntervalNum
}

// GetIntervalNumOk returns a tuple with the IntervalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRateLimitsOrdersResponseResultInner) GetIntervalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.IntervalNum) {
		return nil, false
	}
	return o.IntervalNum, true
}

// HasIntervalNum returns a boolean if a field has been set.
func (o *AccountRateLimitsOrdersResponseResultInner) HasIntervalNum() bool {
	if o != nil && !common.IsNil(o.IntervalNum) {
		return true
	}

	return false
}

// SetIntervalNum gets a reference to the given int64 and assigns it to the IntervalNum field.
func (o *AccountRateLimitsOrdersResponseResultInner) SetIntervalNum(v int64) {
	o.IntervalNum = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *AccountRateLimitsOrdersResponseResultInner) GetLimit() int64 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRateLimitsOrdersResponseResultInner) GetLimitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *AccountRateLimitsOrdersResponseResultInner) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *AccountRateLimitsOrdersResponseResultInner) SetLimit(v int64) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AccountRateLimitsOrdersResponseResultInner) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRateLimitsOrdersResponseResultInner) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AccountRateLimitsOrdersResponseResultInner) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *AccountRateLimitsOrdersResponseResultInner) SetCount(v int64) {
	o.Count = &v
}

func (o AccountRateLimitsOrdersResponseResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountRateLimitsOrdersResponseResultInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RateLimitType) {
		toSerialize["rateLimitType"] = o.RateLimitType
	}
	if !common.IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !common.IsNil(o.IntervalNum) {
		toSerialize["intervalNum"] = o.IntervalNum
	}
	if !common.IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !common.IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountRateLimitsOrdersResponseResultInner) UnmarshalJSON(data []byte) (err error) {
	varAccountRateLimitsOrdersResponseResultInner := _AccountRateLimitsOrdersResponseResultInner{}

	err = json.Unmarshal(data, &varAccountRateLimitsOrdersResponseResultInner)

	if err != nil {
		return err
	}

	*o = AccountRateLimitsOrdersResponseResultInner(varAccountRateLimitsOrdersResponseResultInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rateLimitType")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "intervalNum")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountRateLimitsOrdersResponseResultInner struct {
	value *AccountRateLimitsOrdersResponseResultInner
	isSet bool
}

func (v NullableAccountRateLimitsOrdersResponseResultInner) Get() *AccountRateLimitsOrdersResponseResultInner {
	return v.value
}

func (v *NullableAccountRateLimitsOrdersResponseResultInner) Set(val *AccountRateLimitsOrdersResponseResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountRateLimitsOrdersResponseResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountRateLimitsOrdersResponseResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountRateLimitsOrdersResponseResultInner(val *AccountRateLimitsOrdersResponseResultInner) *NullableAccountRateLimitsOrdersResponseResultInner {
	return &NullableAccountRateLimitsOrdersResponseResultInner{value: val, isSet: true}
}

func (v NullableAccountRateLimitsOrdersResponseResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountRateLimitsOrdersResponseResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
