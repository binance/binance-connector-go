/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the HashrateResaleListResponseDataConfigDetailsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &HashrateResaleListResponseDataConfigDetailsInner{}

// HashrateResaleListResponseDataConfigDetailsInner struct for HashrateResaleListResponseDataConfigDetailsInner
type HashrateResaleListResponseDataConfigDetailsInner struct {
	ConfigId             *int64  `json:"configId,omitempty"`
	PoolUsername         *string `json:"poolUsername,omitempty"`
	ToPoolUsername       *string `json:"toPoolUsername,omitempty"`
	AlgoName             *string `json:"algoName,omitempty"`
	HashRate             *int64  `json:"hashRate,omitempty"`
	StartDay             *int64  `json:"startDay,omitempty"`
	EndDay               *int64  `json:"endDay,omitempty"`
	Status               *int64  `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HashrateResaleListResponseDataConfigDetailsInner HashrateResaleListResponseDataConfigDetailsInner

// NewHashrateResaleListResponseDataConfigDetailsInner instantiates a new HashrateResaleListResponseDataConfigDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHashrateResaleListResponseDataConfigDetailsInner() *HashrateResaleListResponseDataConfigDetailsInner {
	this := HashrateResaleListResponseDataConfigDetailsInner{}
	return &this
}

// NewHashrateResaleListResponseDataConfigDetailsInnerWithDefaults instantiates a new HashrateResaleListResponseDataConfigDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHashrateResaleListResponseDataConfigDetailsInnerWithDefaults() *HashrateResaleListResponseDataConfigDetailsInner {
	this := HashrateResaleListResponseDataConfigDetailsInner{}
	return &this
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise.
func (o *HashrateResaleListResponseDataConfigDetailsInner) GetConfigId() int64 {
	if o == nil || common.IsNil(o.ConfigId) {
		var ret int64
		return ret
	}
	return *o.ConfigId
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleListResponseDataConfigDetailsInner) GetConfigIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ConfigId) {
		return nil, false
	}
	return o.ConfigId, true
}

// HasConfigId returns a boolean if a field has been set.
func (o *HashrateResaleListResponseDataConfigDetailsInner) HasConfigId() bool {
	if o != nil && !common.IsNil(o.ConfigId) {
		return true
	}

	return false
}

// SetConfigId gets a reference to the given int64 and assigns it to the ConfigId field.
func (o *HashrateResaleListResponseDataConfigDetailsInner) SetConfigId(v int64) {
	o.ConfigId = &v
}

// GetPoolUsername returns the PoolUsername field value if set, zero value otherwise.
func (o *HashrateResaleListResponseDataConfigDetailsInner) GetPoolUsername() string {
	if o == nil || common.IsNil(o.PoolUsername) {
		var ret string
		return ret
	}
	return *o.PoolUsername
}

// GetPoolUsernameOk returns a tuple with the PoolUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleListResponseDataConfigDetailsInner) GetPoolUsernameOk() (*string, bool) {
	if o == nil || common.IsNil(o.PoolUsername) {
		return nil, false
	}
	return o.PoolUsername, true
}

// HasPoolUsername returns a boolean if a field has been set.
func (o *HashrateResaleListResponseDataConfigDetailsInner) HasPoolUsername() bool {
	if o != nil && !common.IsNil(o.PoolUsername) {
		return true
	}

	return false
}

// SetPoolUsername gets a reference to the given string and assigns it to the PoolUsername field.
func (o *HashrateResaleListResponseDataConfigDetailsInner) SetPoolUsername(v string) {
	o.PoolUsername = &v
}

// GetToPoolUsername returns the ToPoolUsername field value if set, zero value otherwise.
func (o *HashrateResaleListResponseDataConfigDetailsInner) GetToPoolUsername() string {
	if o == nil || common.IsNil(o.ToPoolUsername) {
		var ret string
		return ret
	}
	return *o.ToPoolUsername
}

// GetToPoolUsernameOk returns a tuple with the ToPoolUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleListResponseDataConfigDetailsInner) GetToPoolUsernameOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToPoolUsername) {
		return nil, false
	}
	return o.ToPoolUsername, true
}

// HasToPoolUsername returns a boolean if a field has been set.
func (o *HashrateResaleListResponseDataConfigDetailsInner) HasToPoolUsername() bool {
	if o != nil && !common.IsNil(o.ToPoolUsername) {
		return true
	}

	return false
}

// SetToPoolUsername gets a reference to the given string and assigns it to the ToPoolUsername field.
func (o *HashrateResaleListResponseDataConfigDetailsInner) SetToPoolUsername(v string) {
	o.ToPoolUsername = &v
}

// GetAlgoName returns the AlgoName field value if set, zero value otherwise.
func (o *HashrateResaleListResponseDataConfigDetailsInner) GetAlgoName() string {
	if o == nil || common.IsNil(o.AlgoName) {
		var ret string
		return ret
	}
	return *o.AlgoName
}

// GetAlgoNameOk returns a tuple with the AlgoName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleListResponseDataConfigDetailsInner) GetAlgoNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.AlgoName) {
		return nil, false
	}
	return o.AlgoName, true
}

// HasAlgoName returns a boolean if a field has been set.
func (o *HashrateResaleListResponseDataConfigDetailsInner) HasAlgoName() bool {
	if o != nil && !common.IsNil(o.AlgoName) {
		return true
	}

	return false
}

// SetAlgoName gets a reference to the given string and assigns it to the AlgoName field.
func (o *HashrateResaleListResponseDataConfigDetailsInner) SetAlgoName(v string) {
	o.AlgoName = &v
}

// GetHashRate returns the HashRate field value if set, zero value otherwise.
func (o *HashrateResaleListResponseDataConfigDetailsInner) GetHashRate() int64 {
	if o == nil || common.IsNil(o.HashRate) {
		var ret int64
		return ret
	}
	return *o.HashRate
}

// GetHashRateOk returns a tuple with the HashRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleListResponseDataConfigDetailsInner) GetHashRateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.HashRate) {
		return nil, false
	}
	return o.HashRate, true
}

// HasHashRate returns a boolean if a field has been set.
func (o *HashrateResaleListResponseDataConfigDetailsInner) HasHashRate() bool {
	if o != nil && !common.IsNil(o.HashRate) {
		return true
	}

	return false
}

// SetHashRate gets a reference to the given int64 and assigns it to the HashRate field.
func (o *HashrateResaleListResponseDataConfigDetailsInner) SetHashRate(v int64) {
	o.HashRate = &v
}

// GetStartDay returns the StartDay field value if set, zero value otherwise.
func (o *HashrateResaleListResponseDataConfigDetailsInner) GetStartDay() int64 {
	if o == nil || common.IsNil(o.StartDay) {
		var ret int64
		return ret
	}
	return *o.StartDay
}

// GetStartDayOk returns a tuple with the StartDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleListResponseDataConfigDetailsInner) GetStartDayOk() (*int64, bool) {
	if o == nil || common.IsNil(o.StartDay) {
		return nil, false
	}
	return o.StartDay, true
}

// HasStartDay returns a boolean if a field has been set.
func (o *HashrateResaleListResponseDataConfigDetailsInner) HasStartDay() bool {
	if o != nil && !common.IsNil(o.StartDay) {
		return true
	}

	return false
}

// SetStartDay gets a reference to the given int64 and assigns it to the StartDay field.
func (o *HashrateResaleListResponseDataConfigDetailsInner) SetStartDay(v int64) {
	o.StartDay = &v
}

// GetEndDay returns the EndDay field value if set, zero value otherwise.
func (o *HashrateResaleListResponseDataConfigDetailsInner) GetEndDay() int64 {
	if o == nil || common.IsNil(o.EndDay) {
		var ret int64
		return ret
	}
	return *o.EndDay
}

// GetEndDayOk returns a tuple with the EndDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleListResponseDataConfigDetailsInner) GetEndDayOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EndDay) {
		return nil, false
	}
	return o.EndDay, true
}

// HasEndDay returns a boolean if a field has been set.
func (o *HashrateResaleListResponseDataConfigDetailsInner) HasEndDay() bool {
	if o != nil && !common.IsNil(o.EndDay) {
		return true
	}

	return false
}

// SetEndDay gets a reference to the given int64 and assigns it to the EndDay field.
func (o *HashrateResaleListResponseDataConfigDetailsInner) SetEndDay(v int64) {
	o.EndDay = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HashrateResaleListResponseDataConfigDetailsInner) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleListResponseDataConfigDetailsInner) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HashrateResaleListResponseDataConfigDetailsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *HashrateResaleListResponseDataConfigDetailsInner) SetStatus(v int64) {
	o.Status = &v
}

func (o HashrateResaleListResponseDataConfigDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HashrateResaleListResponseDataConfigDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ConfigId) {
		toSerialize["configId"] = o.ConfigId
	}
	if !common.IsNil(o.PoolUsername) {
		toSerialize["poolUsername"] = o.PoolUsername
	}
	if !common.IsNil(o.ToPoolUsername) {
		toSerialize["toPoolUsername"] = o.ToPoolUsername
	}
	if !common.IsNil(o.AlgoName) {
		toSerialize["algoName"] = o.AlgoName
	}
	if !common.IsNil(o.HashRate) {
		toSerialize["hashRate"] = o.HashRate
	}
	if !common.IsNil(o.StartDay) {
		toSerialize["startDay"] = o.StartDay
	}
	if !common.IsNil(o.EndDay) {
		toSerialize["endDay"] = o.EndDay
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HashrateResaleListResponseDataConfigDetailsInner) UnmarshalJSON(data []byte) (err error) {
	varHashrateResaleListResponseDataConfigDetailsInner := _HashrateResaleListResponseDataConfigDetailsInner{}

	err = json.Unmarshal(data, &varHashrateResaleListResponseDataConfigDetailsInner)

	if err != nil {
		return err
	}

	*o = HashrateResaleListResponseDataConfigDetailsInner(varHashrateResaleListResponseDataConfigDetailsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "configId")
		delete(additionalProperties, "poolUsername")
		delete(additionalProperties, "toPoolUsername")
		delete(additionalProperties, "algoName")
		delete(additionalProperties, "hashRate")
		delete(additionalProperties, "startDay")
		delete(additionalProperties, "endDay")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHashrateResaleListResponseDataConfigDetailsInner struct {
	value *HashrateResaleListResponseDataConfigDetailsInner
	isSet bool
}

func (v NullableHashrateResaleListResponseDataConfigDetailsInner) Get() *HashrateResaleListResponseDataConfigDetailsInner {
	return v.value
}

func (v *NullableHashrateResaleListResponseDataConfigDetailsInner) Set(val *HashrateResaleListResponseDataConfigDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableHashrateResaleListResponseDataConfigDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableHashrateResaleListResponseDataConfigDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHashrateResaleListResponseDataConfigDetailsInner(val *HashrateResaleListResponseDataConfigDetailsInner) *NullableHashrateResaleListResponseDataConfigDetailsInner {
	return &NullableHashrateResaleListResponseDataConfigDetailsInner{value: val, isSet: true}
}

func (v NullableHashrateResaleListResponseDataConfigDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHashrateResaleListResponseDataConfigDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
