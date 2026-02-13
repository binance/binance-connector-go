/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the SetMarketMakerProtectionConfigResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SetMarketMakerProtectionConfigResponse{}

// SetMarketMakerProtectionConfigResponse struct for SetMarketMakerProtectionConfigResponse
type SetMarketMakerProtectionConfigResponse struct {
	UnderlyingId             *int64  `json:"underlyingId,omitempty"`
	Underlying               *string `json:"underlying,omitempty"`
	WindowTimeInMilliseconds *int64  `json:"windowTimeInMilliseconds,omitempty"`
	FrozenTimeInMilliseconds *int64  `json:"frozenTimeInMilliseconds,omitempty"`
	QtyLimit                 *string `json:"qtyLimit,omitempty"`
	DeltaLimit               *string `json:"deltaLimit,omitempty"`
	LastTriggerTime          *int64  `json:"lastTriggerTime,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _SetMarketMakerProtectionConfigResponse SetMarketMakerProtectionConfigResponse

// NewSetMarketMakerProtectionConfigResponse instantiates a new SetMarketMakerProtectionConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetMarketMakerProtectionConfigResponse() *SetMarketMakerProtectionConfigResponse {
	this := SetMarketMakerProtectionConfigResponse{}
	return &this
}

// NewSetMarketMakerProtectionConfigResponseWithDefaults instantiates a new SetMarketMakerProtectionConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetMarketMakerProtectionConfigResponseWithDefaults() *SetMarketMakerProtectionConfigResponse {
	this := SetMarketMakerProtectionConfigResponse{}
	return &this
}

// GetUnderlyingId returns the UnderlyingId field value if set, zero value otherwise.
func (o *SetMarketMakerProtectionConfigResponse) GetUnderlyingId() int64 {
	if o == nil || common.IsNil(o.UnderlyingId) {
		var ret int64
		return ret
	}
	return *o.UnderlyingId
}

// GetUnderlyingIdOk returns a tuple with the UnderlyingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetMarketMakerProtectionConfigResponse) GetUnderlyingIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UnderlyingId) {
		return nil, false
	}
	return o.UnderlyingId, true
}

// HasUnderlyingId returns a boolean if a field has been set.
func (o *SetMarketMakerProtectionConfigResponse) HasUnderlyingId() bool {
	if o != nil && !common.IsNil(o.UnderlyingId) {
		return true
	}

	return false
}

// SetUnderlyingId gets a reference to the given int64 and assigns it to the UnderlyingId field.
func (o *SetMarketMakerProtectionConfigResponse) SetUnderlyingId(v int64) {
	o.UnderlyingId = &v
}

// GetUnderlying returns the Underlying field value if set, zero value otherwise.
func (o *SetMarketMakerProtectionConfigResponse) GetUnderlying() string {
	if o == nil || common.IsNil(o.Underlying) {
		var ret string
		return ret
	}
	return *o.Underlying
}

// GetUnderlyingOk returns a tuple with the Underlying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetMarketMakerProtectionConfigResponse) GetUnderlyingOk() (*string, bool) {
	if o == nil || common.IsNil(o.Underlying) {
		return nil, false
	}
	return o.Underlying, true
}

// HasUnderlying returns a boolean if a field has been set.
func (o *SetMarketMakerProtectionConfigResponse) HasUnderlying() bool {
	if o != nil && !common.IsNil(o.Underlying) {
		return true
	}

	return false
}

// SetUnderlying gets a reference to the given string and assigns it to the Underlying field.
func (o *SetMarketMakerProtectionConfigResponse) SetUnderlying(v string) {
	o.Underlying = &v
}

// GetWindowTimeInMilliseconds returns the WindowTimeInMilliseconds field value if set, zero value otherwise.
func (o *SetMarketMakerProtectionConfigResponse) GetWindowTimeInMilliseconds() int64 {
	if o == nil || common.IsNil(o.WindowTimeInMilliseconds) {
		var ret int64
		return ret
	}
	return *o.WindowTimeInMilliseconds
}

// GetWindowTimeInMillisecondsOk returns a tuple with the WindowTimeInMilliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetMarketMakerProtectionConfigResponse) GetWindowTimeInMillisecondsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.WindowTimeInMilliseconds) {
		return nil, false
	}
	return o.WindowTimeInMilliseconds, true
}

// HasWindowTimeInMilliseconds returns a boolean if a field has been set.
func (o *SetMarketMakerProtectionConfigResponse) HasWindowTimeInMilliseconds() bool {
	if o != nil && !common.IsNil(o.WindowTimeInMilliseconds) {
		return true
	}

	return false
}

// SetWindowTimeInMilliseconds gets a reference to the given int64 and assigns it to the WindowTimeInMilliseconds field.
func (o *SetMarketMakerProtectionConfigResponse) SetWindowTimeInMilliseconds(v int64) {
	o.WindowTimeInMilliseconds = &v
}

// GetFrozenTimeInMilliseconds returns the FrozenTimeInMilliseconds field value if set, zero value otherwise.
func (o *SetMarketMakerProtectionConfigResponse) GetFrozenTimeInMilliseconds() int64 {
	if o == nil || common.IsNil(o.FrozenTimeInMilliseconds) {
		var ret int64
		return ret
	}
	return *o.FrozenTimeInMilliseconds
}

// GetFrozenTimeInMillisecondsOk returns a tuple with the FrozenTimeInMilliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetMarketMakerProtectionConfigResponse) GetFrozenTimeInMillisecondsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FrozenTimeInMilliseconds) {
		return nil, false
	}
	return o.FrozenTimeInMilliseconds, true
}

// HasFrozenTimeInMilliseconds returns a boolean if a field has been set.
func (o *SetMarketMakerProtectionConfigResponse) HasFrozenTimeInMilliseconds() bool {
	if o != nil && !common.IsNil(o.FrozenTimeInMilliseconds) {
		return true
	}

	return false
}

// SetFrozenTimeInMilliseconds gets a reference to the given int64 and assigns it to the FrozenTimeInMilliseconds field.
func (o *SetMarketMakerProtectionConfigResponse) SetFrozenTimeInMilliseconds(v int64) {
	o.FrozenTimeInMilliseconds = &v
}

// GetQtyLimit returns the QtyLimit field value if set, zero value otherwise.
func (o *SetMarketMakerProtectionConfigResponse) GetQtyLimit() string {
	if o == nil || common.IsNil(o.QtyLimit) {
		var ret string
		return ret
	}
	return *o.QtyLimit
}

// GetQtyLimitOk returns a tuple with the QtyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetMarketMakerProtectionConfigResponse) GetQtyLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.QtyLimit) {
		return nil, false
	}
	return o.QtyLimit, true
}

// HasQtyLimit returns a boolean if a field has been set.
func (o *SetMarketMakerProtectionConfigResponse) HasQtyLimit() bool {
	if o != nil && !common.IsNil(o.QtyLimit) {
		return true
	}

	return false
}

// SetQtyLimit gets a reference to the given string and assigns it to the QtyLimit field.
func (o *SetMarketMakerProtectionConfigResponse) SetQtyLimit(v string) {
	o.QtyLimit = &v
}

// GetDeltaLimit returns the DeltaLimit field value if set, zero value otherwise.
func (o *SetMarketMakerProtectionConfigResponse) GetDeltaLimit() string {
	if o == nil || common.IsNil(o.DeltaLimit) {
		var ret string
		return ret
	}
	return *o.DeltaLimit
}

// GetDeltaLimitOk returns a tuple with the DeltaLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetMarketMakerProtectionConfigResponse) GetDeltaLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeltaLimit) {
		return nil, false
	}
	return o.DeltaLimit, true
}

// HasDeltaLimit returns a boolean if a field has been set.
func (o *SetMarketMakerProtectionConfigResponse) HasDeltaLimit() bool {
	if o != nil && !common.IsNil(o.DeltaLimit) {
		return true
	}

	return false
}

// SetDeltaLimit gets a reference to the given string and assigns it to the DeltaLimit field.
func (o *SetMarketMakerProtectionConfigResponse) SetDeltaLimit(v string) {
	o.DeltaLimit = &v
}

// GetLastTriggerTime returns the LastTriggerTime field value if set, zero value otherwise.
func (o *SetMarketMakerProtectionConfigResponse) GetLastTriggerTime() int64 {
	if o == nil || common.IsNil(o.LastTriggerTime) {
		var ret int64
		return ret
	}
	return *o.LastTriggerTime
}

// GetLastTriggerTimeOk returns a tuple with the LastTriggerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetMarketMakerProtectionConfigResponse) GetLastTriggerTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastTriggerTime) {
		return nil, false
	}
	return o.LastTriggerTime, true
}

// HasLastTriggerTime returns a boolean if a field has been set.
func (o *SetMarketMakerProtectionConfigResponse) HasLastTriggerTime() bool {
	if o != nil && !common.IsNil(o.LastTriggerTime) {
		return true
	}

	return false
}

// SetLastTriggerTime gets a reference to the given int64 and assigns it to the LastTriggerTime field.
func (o *SetMarketMakerProtectionConfigResponse) SetLastTriggerTime(v int64) {
	o.LastTriggerTime = &v
}

func (o SetMarketMakerProtectionConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetMarketMakerProtectionConfigResponse) ToMap() (map[string]interface{}, error) {
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

func (o *SetMarketMakerProtectionConfigResponse) UnmarshalJSON(data []byte) (err error) {
	varSetMarketMakerProtectionConfigResponse := _SetMarketMakerProtectionConfigResponse{}

	err = json.Unmarshal(data, &varSetMarketMakerProtectionConfigResponse)

	if err != nil {
		return err
	}

	*o = SetMarketMakerProtectionConfigResponse(varSetMarketMakerProtectionConfigResponse)

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

type NullableSetMarketMakerProtectionConfigResponse struct {
	value *SetMarketMakerProtectionConfigResponse
	isSet bool
}

func (v NullableSetMarketMakerProtectionConfigResponse) Get() *SetMarketMakerProtectionConfigResponse {
	return v.value
}

func (v *NullableSetMarketMakerProtectionConfigResponse) Set(val *SetMarketMakerProtectionConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSetMarketMakerProtectionConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSetMarketMakerProtectionConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetMarketMakerProtectionConfigResponse(val *SetMarketMakerProtectionConfigResponse) *NullableSetMarketMakerProtectionConfigResponse {
	return &NullableSetMarketMakerProtectionConfigResponse{value: val, isSet: true}
}

func (v NullableSetMarketMakerProtectionConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetMarketMakerProtectionConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
