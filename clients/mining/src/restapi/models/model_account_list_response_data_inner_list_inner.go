/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountListResponseDataInnerListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountListResponseDataInnerListInner{}

// AccountListResponseDataInnerListInner struct for AccountListResponseDataInnerListInner
type AccountListResponseDataInnerListInner struct {
	Time                 *int64  `json:"time,omitempty"`
	Hashrate             *string `json:"hashrate,omitempty"`
	Reject               *string `json:"reject,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountListResponseDataInnerListInner AccountListResponseDataInnerListInner

// NewAccountListResponseDataInnerListInner instantiates a new AccountListResponseDataInnerListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountListResponseDataInnerListInner() *AccountListResponseDataInnerListInner {
	this := AccountListResponseDataInnerListInner{}
	return &this
}

// NewAccountListResponseDataInnerListInnerWithDefaults instantiates a new AccountListResponseDataInnerListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountListResponseDataInnerListInnerWithDefaults() *AccountListResponseDataInnerListInner {
	this := AccountListResponseDataInnerListInner{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *AccountListResponseDataInnerListInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountListResponseDataInnerListInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *AccountListResponseDataInnerListInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *AccountListResponseDataInnerListInner) SetTime(v int64) {
	o.Time = &v
}

// GetHashrate returns the Hashrate field value if set, zero value otherwise.
func (o *AccountListResponseDataInnerListInner) GetHashrate() string {
	if o == nil || common.IsNil(o.Hashrate) {
		var ret string
		return ret
	}
	return *o.Hashrate
}

// GetHashrateOk returns a tuple with the Hashrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountListResponseDataInnerListInner) GetHashrateOk() (*string, bool) {
	if o == nil || common.IsNil(o.Hashrate) {
		return nil, false
	}
	return o.Hashrate, true
}

// HasHashrate returns a boolean if a field has been set.
func (o *AccountListResponseDataInnerListInner) HasHashrate() bool {
	if o != nil && !common.IsNil(o.Hashrate) {
		return true
	}

	return false
}

// SetHashrate gets a reference to the given string and assigns it to the Hashrate field.
func (o *AccountListResponseDataInnerListInner) SetHashrate(v string) {
	o.Hashrate = &v
}

// GetReject returns the Reject field value if set, zero value otherwise.
func (o *AccountListResponseDataInnerListInner) GetReject() string {
	if o == nil || common.IsNil(o.Reject) {
		var ret string
		return ret
	}
	return *o.Reject
}

// GetRejectOk returns a tuple with the Reject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountListResponseDataInnerListInner) GetRejectOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reject) {
		return nil, false
	}
	return o.Reject, true
}

// HasReject returns a boolean if a field has been set.
func (o *AccountListResponseDataInnerListInner) HasReject() bool {
	if o != nil && !common.IsNil(o.Reject) {
		return true
	}

	return false
}

// SetReject gets a reference to the given string and assigns it to the Reject field.
func (o *AccountListResponseDataInnerListInner) SetReject(v string) {
	o.Reject = &v
}

func (o AccountListResponseDataInnerListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountListResponseDataInnerListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.Hashrate) {
		toSerialize["hashrate"] = o.Hashrate
	}
	if !common.IsNil(o.Reject) {
		toSerialize["reject"] = o.Reject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountListResponseDataInnerListInner) UnmarshalJSON(data []byte) (err error) {
	varAccountListResponseDataInnerListInner := _AccountListResponseDataInnerListInner{}

	err = json.Unmarshal(data, &varAccountListResponseDataInnerListInner)

	if err != nil {
		return err
	}

	*o = AccountListResponseDataInnerListInner(varAccountListResponseDataInnerListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "hashrate")
		delete(additionalProperties, "reject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountListResponseDataInnerListInner struct {
	value *AccountListResponseDataInnerListInner
	isSet bool
}

func (v NullableAccountListResponseDataInnerListInner) Get() *AccountListResponseDataInnerListInner {
	return v.value
}

func (v *NullableAccountListResponseDataInnerListInner) Set(val *AccountListResponseDataInnerListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountListResponseDataInnerListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountListResponseDataInnerListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountListResponseDataInnerListInner(val *AccountListResponseDataInnerListInner) *NullableAccountListResponseDataInnerListInner {
	return &NullableAccountListResponseDataInnerListInner{value: val, isSet: true}
}

func (v NullableAccountListResponseDataInnerListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountListResponseDataInnerListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
