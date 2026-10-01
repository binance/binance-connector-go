/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ReferencePriceResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReferencePriceResponse2{}

// ReferencePriceResponse2 If no reference price is set:
type ReferencePriceResponse2 struct {
	Id                   *string                           `json:"id,omitempty"`
	Status               *int64                            `json:"status,omitempty"`
	Result               *ReferencePriceResponse2Result    `json:"result,omitempty"`
	RateLimits           []AvgPriceResponseRateLimitsInner `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReferencePriceResponse2 ReferencePriceResponse2

// NewReferencePriceResponse2 instantiates a new ReferencePriceResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferencePriceResponse2() *ReferencePriceResponse2 {
	this := ReferencePriceResponse2{}
	return &this
}

// NewReferencePriceResponse2WithDefaults instantiates a new ReferencePriceResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferencePriceResponse2WithDefaults() *ReferencePriceResponse2 {
	this := ReferencePriceResponse2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReferencePriceResponse2) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse2) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReferencePriceResponse2) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReferencePriceResponse2) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReferencePriceResponse2) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse2) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReferencePriceResponse2) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *ReferencePriceResponse2) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ReferencePriceResponse2) GetResult() ReferencePriceResponse2Result {
	if o == nil || common.IsNil(o.Result) {
		var ret ReferencePriceResponse2Result
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse2) GetResultOk() (*ReferencePriceResponse2Result, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ReferencePriceResponse2) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ReferencePriceResponse2Result and assigns it to the Result field.
func (o *ReferencePriceResponse2) SetResult(v ReferencePriceResponse2Result) {
	o.Result = &v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *ReferencePriceResponse2) GetRateLimits() []AvgPriceResponseRateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []AvgPriceResponseRateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse2) GetRateLimitsOk() ([]AvgPriceResponseRateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *ReferencePriceResponse2) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []AvgPriceResponseRateLimitsInner and assigns it to the RateLimits field.
func (o *ReferencePriceResponse2) SetRateLimits(v []AvgPriceResponseRateLimitsInner) {
	o.RateLimits = v
}

func (o ReferencePriceResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferencePriceResponse2) ToMap() (map[string]interface{}, error) {
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

func (o *ReferencePriceResponse2) UnmarshalJSON(data []byte) (err error) {
	varReferencePriceResponse2 := _ReferencePriceResponse2{}

	err = json.Unmarshal(data, &varReferencePriceResponse2)

	if err != nil {
		return err
	}

	*o = ReferencePriceResponse2(varReferencePriceResponse2)

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

type NullableReferencePriceResponse2 struct {
	value *ReferencePriceResponse2
	isSet bool
}

func (v NullableReferencePriceResponse2) Get() *ReferencePriceResponse2 {
	return v.value
}

func (v *NullableReferencePriceResponse2) Set(val *ReferencePriceResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableReferencePriceResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableReferencePriceResponse2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferencePriceResponse2(val *ReferencePriceResponse2) *NullableReferencePriceResponse2 {
	return &NullableReferencePriceResponse2{value: val, isSet: true}
}

func (v NullableReferencePriceResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferencePriceResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
