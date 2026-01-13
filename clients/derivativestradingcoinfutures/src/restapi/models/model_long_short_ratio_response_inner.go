/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the LongShortRatioResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &LongShortRatioResponseInner{}

// LongShortRatioResponseInner struct for LongShortRatioResponseInner
type LongShortRatioResponseInner struct {
	Pair                 *string `json:"pair,omitempty"`
	LongShortRatio       *string `json:"longShortRatio,omitempty"`
	LongAccount          *string `json:"longAccount,omitempty"`
	ShortAccount         *string `json:"shortAccount,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LongShortRatioResponseInner LongShortRatioResponseInner

// NewLongShortRatioResponseInner instantiates a new LongShortRatioResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLongShortRatioResponseInner() *LongShortRatioResponseInner {
	this := LongShortRatioResponseInner{}
	return &this
}

// NewLongShortRatioResponseInnerWithDefaults instantiates a new LongShortRatioResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLongShortRatioResponseInnerWithDefaults() *LongShortRatioResponseInner {
	this := LongShortRatioResponseInner{}
	return &this
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *LongShortRatioResponseInner) GetPair() string {
	if o == nil || common.IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongShortRatioResponseInner) GetPairOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *LongShortRatioResponseInner) HasPair() bool {
	if o != nil && !common.IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *LongShortRatioResponseInner) SetPair(v string) {
	o.Pair = &v
}

// GetLongShortRatio returns the LongShortRatio field value if set, zero value otherwise.
func (o *LongShortRatioResponseInner) GetLongShortRatio() string {
	if o == nil || common.IsNil(o.LongShortRatio) {
		var ret string
		return ret
	}
	return *o.LongShortRatio
}

// GetLongShortRatioOk returns a tuple with the LongShortRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongShortRatioResponseInner) GetLongShortRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.LongShortRatio) {
		return nil, false
	}
	return o.LongShortRatio, true
}

// HasLongShortRatio returns a boolean if a field has been set.
func (o *LongShortRatioResponseInner) HasLongShortRatio() bool {
	if o != nil && !common.IsNil(o.LongShortRatio) {
		return true
	}

	return false
}

// SetLongShortRatio gets a reference to the given string and assigns it to the LongShortRatio field.
func (o *LongShortRatioResponseInner) SetLongShortRatio(v string) {
	o.LongShortRatio = &v
}

// GetLongAccount returns the LongAccount field value if set, zero value otherwise.
func (o *LongShortRatioResponseInner) GetLongAccount() string {
	if o == nil || common.IsNil(o.LongAccount) {
		var ret string
		return ret
	}
	return *o.LongAccount
}

// GetLongAccountOk returns a tuple with the LongAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongShortRatioResponseInner) GetLongAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.LongAccount) {
		return nil, false
	}
	return o.LongAccount, true
}

// HasLongAccount returns a boolean if a field has been set.
func (o *LongShortRatioResponseInner) HasLongAccount() bool {
	if o != nil && !common.IsNil(o.LongAccount) {
		return true
	}

	return false
}

// SetLongAccount gets a reference to the given string and assigns it to the LongAccount field.
func (o *LongShortRatioResponseInner) SetLongAccount(v string) {
	o.LongAccount = &v
}

// GetShortAccount returns the ShortAccount field value if set, zero value otherwise.
func (o *LongShortRatioResponseInner) GetShortAccount() string {
	if o == nil || common.IsNil(o.ShortAccount) {
		var ret string
		return ret
	}
	return *o.ShortAccount
}

// GetShortAccountOk returns a tuple with the ShortAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongShortRatioResponseInner) GetShortAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShortAccount) {
		return nil, false
	}
	return o.ShortAccount, true
}

// HasShortAccount returns a boolean if a field has been set.
func (o *LongShortRatioResponseInner) HasShortAccount() bool {
	if o != nil && !common.IsNil(o.ShortAccount) {
		return true
	}

	return false
}

// SetShortAccount gets a reference to the given string and assigns it to the ShortAccount field.
func (o *LongShortRatioResponseInner) SetShortAccount(v string) {
	o.ShortAccount = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *LongShortRatioResponseInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongShortRatioResponseInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *LongShortRatioResponseInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *LongShortRatioResponseInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o LongShortRatioResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LongShortRatioResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !common.IsNil(o.LongShortRatio) {
		toSerialize["longShortRatio"] = o.LongShortRatio
	}
	if !common.IsNil(o.LongAccount) {
		toSerialize["longAccount"] = o.LongAccount
	}
	if !common.IsNil(o.ShortAccount) {
		toSerialize["shortAccount"] = o.ShortAccount
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LongShortRatioResponseInner) UnmarshalJSON(data []byte) (err error) {
	varLongShortRatioResponseInner := _LongShortRatioResponseInner{}

	err = json.Unmarshal(data, &varLongShortRatioResponseInner)

	if err != nil {
		return err
	}

	*o = LongShortRatioResponseInner(varLongShortRatioResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pair")
		delete(additionalProperties, "longShortRatio")
		delete(additionalProperties, "longAccount")
		delete(additionalProperties, "shortAccount")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLongShortRatioResponseInner struct {
	value *LongShortRatioResponseInner
	isSet bool
}

func (v NullableLongShortRatioResponseInner) Get() *LongShortRatioResponseInner {
	return v.value
}

func (v *NullableLongShortRatioResponseInner) Set(val *LongShortRatioResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableLongShortRatioResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableLongShortRatioResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLongShortRatioResponseInner(val *LongShortRatioResponseInner) *NullableLongShortRatioResponseInner {
	return &NullableLongShortRatioResponseInner{value: val, isSet: true}
}

func (v NullableLongShortRatioResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLongShortRatioResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
