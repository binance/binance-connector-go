/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ResetMarketMakerProtectionConfigResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ResetMarketMakerProtectionConfigResponse{}

// ResetMarketMakerProtectionConfigResponse struct for ResetMarketMakerProtectionConfigResponse
type ResetMarketMakerProtectionConfigResponse struct {
	UnderlyingId             *int64  `json:"underlyingId,omitempty"`
	Underlying               *string `json:"underlying,omitempty"`
	WindowTimeInMilliseconds *int64  `json:"windowTimeInMilliseconds,omitempty"`
	FrozenTimeInMilliseconds *int64  `json:"frozenTimeInMilliseconds,omitempty"`
	QtyLimit                 *string `json:"qtyLimit,omitempty"`
	DeltaLimit               *string `json:"deltaLimit,omitempty"`
	LastTriggerTime          *int64  `json:"lastTriggerTime,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _ResetMarketMakerProtectionConfigResponse ResetMarketMakerProtectionConfigResponse

// NewResetMarketMakerProtectionConfigResponse instantiates a new ResetMarketMakerProtectionConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetMarketMakerProtectionConfigResponse() *ResetMarketMakerProtectionConfigResponse {
	this := ResetMarketMakerProtectionConfigResponse{}
	return &this
}

// NewResetMarketMakerProtectionConfigResponseWithDefaults instantiates a new ResetMarketMakerProtectionConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetMarketMakerProtectionConfigResponseWithDefaults() *ResetMarketMakerProtectionConfigResponse {
	this := ResetMarketMakerProtectionConfigResponse{}
	return &this
}

// GetUnderlyingId returns the UnderlyingId field value if set, zero value otherwise.
func (o *ResetMarketMakerProtectionConfigResponse) GetUnderlyingId() int64 {
	if o == nil || common.IsNil(o.UnderlyingId) {
		var ret int64
		return ret
	}
	return *o.UnderlyingId
}

// GetUnderlyingIdOk returns a tuple with the UnderlyingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetMarketMakerProtectionConfigResponse) GetUnderlyingIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UnderlyingId) {
		return nil, false
	}
	return o.UnderlyingId, true
}

// HasUnderlyingId returns a boolean if a field has been set.
func (o *ResetMarketMakerProtectionConfigResponse) HasUnderlyingId() bool {
	if o != nil && !common.IsNil(o.UnderlyingId) {
		return true
	}

	return false
}

// SetUnderlyingId gets a reference to the given int64 and assigns it to the UnderlyingId field.
func (o *ResetMarketMakerProtectionConfigResponse) SetUnderlyingId(v int64) {
	o.UnderlyingId = &v
}

// GetUnderlying returns the Underlying field value if set, zero value otherwise.
func (o *ResetMarketMakerProtectionConfigResponse) GetUnderlying() string {
	if o == nil || common.IsNil(o.Underlying) {
		var ret string
		return ret
	}
	return *o.Underlying
}

// GetUnderlyingOk returns a tuple with the Underlying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetMarketMakerProtectionConfigResponse) GetUnderlyingOk() (*string, bool) {
	if o == nil || common.IsNil(o.Underlying) {
		return nil, false
	}
	return o.Underlying, true
}

// HasUnderlying returns a boolean if a field has been set.
func (o *ResetMarketMakerProtectionConfigResponse) HasUnderlying() bool {
	if o != nil && !common.IsNil(o.Underlying) {
		return true
	}

	return false
}

// SetUnderlying gets a reference to the given string and assigns it to the Underlying field.
func (o *ResetMarketMakerProtectionConfigResponse) SetUnderlying(v string) {
	o.Underlying = &v
}

// GetWindowTimeInMilliseconds returns the WindowTimeInMilliseconds field value if set, zero value otherwise.
func (o *ResetMarketMakerProtectionConfigResponse) GetWindowTimeInMilliseconds() int64 {
	if o == nil || common.IsNil(o.WindowTimeInMilliseconds) {
		var ret int64
		return ret
	}
	return *o.WindowTimeInMilliseconds
}

// GetWindowTimeInMillisecondsOk returns a tuple with the WindowTimeInMilliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetMarketMakerProtectionConfigResponse) GetWindowTimeInMillisecondsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.WindowTimeInMilliseconds) {
		return nil, false
	}
	return o.WindowTimeInMilliseconds, true
}

// HasWindowTimeInMilliseconds returns a boolean if a field has been set.
func (o *ResetMarketMakerProtectionConfigResponse) HasWindowTimeInMilliseconds() bool {
	if o != nil && !common.IsNil(o.WindowTimeInMilliseconds) {
		return true
	}

	return false
}

// SetWindowTimeInMilliseconds gets a reference to the given int64 and assigns it to the WindowTimeInMilliseconds field.
func (o *ResetMarketMakerProtectionConfigResponse) SetWindowTimeInMilliseconds(v int64) {
	o.WindowTimeInMilliseconds = &v
}

// GetFrozenTimeInMilliseconds returns the FrozenTimeInMilliseconds field value if set, zero value otherwise.
func (o *ResetMarketMakerProtectionConfigResponse) GetFrozenTimeInMilliseconds() int64 {
	if o == nil || common.IsNil(o.FrozenTimeInMilliseconds) {
		var ret int64
		return ret
	}
	return *o.FrozenTimeInMilliseconds
}

// GetFrozenTimeInMillisecondsOk returns a tuple with the FrozenTimeInMilliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetMarketMakerProtectionConfigResponse) GetFrozenTimeInMillisecondsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FrozenTimeInMilliseconds) {
		return nil, false
	}
	return o.FrozenTimeInMilliseconds, true
}

// HasFrozenTimeInMilliseconds returns a boolean if a field has been set.
func (o *ResetMarketMakerProtectionConfigResponse) HasFrozenTimeInMilliseconds() bool {
	if o != nil && !common.IsNil(o.FrozenTimeInMilliseconds) {
		return true
	}

	return false
}

// SetFrozenTimeInMilliseconds gets a reference to the given int64 and assigns it to the FrozenTimeInMilliseconds field.
func (o *ResetMarketMakerProtectionConfigResponse) SetFrozenTimeInMilliseconds(v int64) {
	o.FrozenTimeInMilliseconds = &v
}

// GetQtyLimit returns the QtyLimit field value if set, zero value otherwise.
func (o *ResetMarketMakerProtectionConfigResponse) GetQtyLimit() string {
	if o == nil || common.IsNil(o.QtyLimit) {
		var ret string
		return ret
	}
	return *o.QtyLimit
}

// GetQtyLimitOk returns a tuple with the QtyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetMarketMakerProtectionConfigResponse) GetQtyLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.QtyLimit) {
		return nil, false
	}
	return o.QtyLimit, true
}

// HasQtyLimit returns a boolean if a field has been set.
func (o *ResetMarketMakerProtectionConfigResponse) HasQtyLimit() bool {
	if o != nil && !common.IsNil(o.QtyLimit) {
		return true
	}

	return false
}

// SetQtyLimit gets a reference to the given string and assigns it to the QtyLimit field.
func (o *ResetMarketMakerProtectionConfigResponse) SetQtyLimit(v string) {
	o.QtyLimit = &v
}

// GetDeltaLimit returns the DeltaLimit field value if set, zero value otherwise.
func (o *ResetMarketMakerProtectionConfigResponse) GetDeltaLimit() string {
	if o == nil || common.IsNil(o.DeltaLimit) {
		var ret string
		return ret
	}
	return *o.DeltaLimit
}

// GetDeltaLimitOk returns a tuple with the DeltaLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetMarketMakerProtectionConfigResponse) GetDeltaLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeltaLimit) {
		return nil, false
	}
	return o.DeltaLimit, true
}

// HasDeltaLimit returns a boolean if a field has been set.
func (o *ResetMarketMakerProtectionConfigResponse) HasDeltaLimit() bool {
	if o != nil && !common.IsNil(o.DeltaLimit) {
		return true
	}

	return false
}

// SetDeltaLimit gets a reference to the given string and assigns it to the DeltaLimit field.
func (o *ResetMarketMakerProtectionConfigResponse) SetDeltaLimit(v string) {
	o.DeltaLimit = &v
}

// GetLastTriggerTime returns the LastTriggerTime field value if set, zero value otherwise.
func (o *ResetMarketMakerProtectionConfigResponse) GetLastTriggerTime() int64 {
	if o == nil || common.IsNil(o.LastTriggerTime) {
		var ret int64
		return ret
	}
	return *o.LastTriggerTime
}

// GetLastTriggerTimeOk returns a tuple with the LastTriggerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetMarketMakerProtectionConfigResponse) GetLastTriggerTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastTriggerTime) {
		return nil, false
	}
	return o.LastTriggerTime, true
}

// HasLastTriggerTime returns a boolean if a field has been set.
func (o *ResetMarketMakerProtectionConfigResponse) HasLastTriggerTime() bool {
	if o != nil && !common.IsNil(o.LastTriggerTime) {
		return true
	}

	return false
}

// SetLastTriggerTime gets a reference to the given int64 and assigns it to the LastTriggerTime field.
func (o *ResetMarketMakerProtectionConfigResponse) SetLastTriggerTime(v int64) {
	o.LastTriggerTime = &v
}

func (o ResetMarketMakerProtectionConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResetMarketMakerProtectionConfigResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.UnderlyingId) {
		toSerialize["underlyingId"] = o.UnderlyingId
	}
	if !common.IsNil(o.Underlying) {
		toSerialize["underlying"] = o.Underlying
	}
	if !common.IsNil(o.WindowTimeInMilliseconds) {
		toSerialize["windowTimeInMilliseconds"] = o.WindowTimeInMilliseconds
	}
	if !common.IsNil(o.FrozenTimeInMilliseconds) {
		toSerialize["frozenTimeInMilliseconds"] = o.FrozenTimeInMilliseconds
	}
	if !common.IsNil(o.QtyLimit) {
		toSerialize["qtyLimit"] = o.QtyLimit
	}
	if !common.IsNil(o.DeltaLimit) {
		toSerialize["deltaLimit"] = o.DeltaLimit
	}
	if !common.IsNil(o.LastTriggerTime) {
		toSerialize["lastTriggerTime"] = o.LastTriggerTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResetMarketMakerProtectionConfigResponse) UnmarshalJSON(data []byte) (err error) {
	varResetMarketMakerProtectionConfigResponse := _ResetMarketMakerProtectionConfigResponse{}

	err = json.Unmarshal(data, &varResetMarketMakerProtectionConfigResponse)

	if err != nil {
		return err
	}

	*o = ResetMarketMakerProtectionConfigResponse(varResetMarketMakerProtectionConfigResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "underlyingId")
		delete(additionalProperties, "underlying")
		delete(additionalProperties, "windowTimeInMilliseconds")
		delete(additionalProperties, "frozenTimeInMilliseconds")
		delete(additionalProperties, "qtyLimit")
		delete(additionalProperties, "deltaLimit")
		delete(additionalProperties, "lastTriggerTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResetMarketMakerProtectionConfigResponse struct {
	value *ResetMarketMakerProtectionConfigResponse
	isSet bool
}

func (v NullableResetMarketMakerProtectionConfigResponse) Get() *ResetMarketMakerProtectionConfigResponse {
	return v.value
}

func (v *NullableResetMarketMakerProtectionConfigResponse) Set(val *ResetMarketMakerProtectionConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResetMarketMakerProtectionConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResetMarketMakerProtectionConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetMarketMakerProtectionConfigResponse(val *ResetMarketMakerProtectionConfigResponse) *NullableResetMarketMakerProtectionConfigResponse {
	return &NullableResetMarketMakerProtectionConfigResponse{value: val, isSet: true}
}

func (v NullableResetMarketMakerProtectionConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetMarketMakerProtectionConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
