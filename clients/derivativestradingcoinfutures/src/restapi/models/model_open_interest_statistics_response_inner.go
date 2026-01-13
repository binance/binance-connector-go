/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OpenInterestStatisticsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OpenInterestStatisticsResponseInner{}

// OpenInterestStatisticsResponseInner struct for OpenInterestStatisticsResponseInner
type OpenInterestStatisticsResponseInner struct {
	Pair                 *string `json:"pair,omitempty"`
	ContractType         *string `json:"contractType,omitempty"`
	SumOpenInterest      *string `json:"sumOpenInterest,omitempty"`
	SumOpenInterestValue *string `json:"sumOpenInterestValue,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenInterestStatisticsResponseInner OpenInterestStatisticsResponseInner

// NewOpenInterestStatisticsResponseInner instantiates a new OpenInterestStatisticsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenInterestStatisticsResponseInner() *OpenInterestStatisticsResponseInner {
	this := OpenInterestStatisticsResponseInner{}
	return &this
}

// NewOpenInterestStatisticsResponseInnerWithDefaults instantiates a new OpenInterestStatisticsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenInterestStatisticsResponseInnerWithDefaults() *OpenInterestStatisticsResponseInner {
	this := OpenInterestStatisticsResponseInner{}
	return &this
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *OpenInterestStatisticsResponseInner) GetPair() string {
	if o == nil || common.IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestStatisticsResponseInner) GetPairOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *OpenInterestStatisticsResponseInner) HasPair() bool {
	if o != nil && !common.IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *OpenInterestStatisticsResponseInner) SetPair(v string) {
	o.Pair = &v
}

// GetContractType returns the ContractType field value if set, zero value otherwise.
func (o *OpenInterestStatisticsResponseInner) GetContractType() string {
	if o == nil || common.IsNil(o.ContractType) {
		var ret string
		return ret
	}
	return *o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestStatisticsResponseInner) GetContractTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContractType) {
		return nil, false
	}
	return o.ContractType, true
}

// HasContractType returns a boolean if a field has been set.
func (o *OpenInterestStatisticsResponseInner) HasContractType() bool {
	if o != nil && !common.IsNil(o.ContractType) {
		return true
	}

	return false
}

// SetContractType gets a reference to the given string and assigns it to the ContractType field.
func (o *OpenInterestStatisticsResponseInner) SetContractType(v string) {
	o.ContractType = &v
}

// GetSumOpenInterest returns the SumOpenInterest field value if set, zero value otherwise.
func (o *OpenInterestStatisticsResponseInner) GetSumOpenInterest() string {
	if o == nil || common.IsNil(o.SumOpenInterest) {
		var ret string
		return ret
	}
	return *o.SumOpenInterest
}

// GetSumOpenInterestOk returns a tuple with the SumOpenInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestStatisticsResponseInner) GetSumOpenInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.SumOpenInterest) {
		return nil, false
	}
	return o.SumOpenInterest, true
}

// HasSumOpenInterest returns a boolean if a field has been set.
func (o *OpenInterestStatisticsResponseInner) HasSumOpenInterest() bool {
	if o != nil && !common.IsNil(o.SumOpenInterest) {
		return true
	}

	return false
}

// SetSumOpenInterest gets a reference to the given string and assigns it to the SumOpenInterest field.
func (o *OpenInterestStatisticsResponseInner) SetSumOpenInterest(v string) {
	o.SumOpenInterest = &v
}

// GetSumOpenInterestValue returns the SumOpenInterestValue field value if set, zero value otherwise.
func (o *OpenInterestStatisticsResponseInner) GetSumOpenInterestValue() string {
	if o == nil || common.IsNil(o.SumOpenInterestValue) {
		var ret string
		return ret
	}
	return *o.SumOpenInterestValue
}

// GetSumOpenInterestValueOk returns a tuple with the SumOpenInterestValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestStatisticsResponseInner) GetSumOpenInterestValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.SumOpenInterestValue) {
		return nil, false
	}
	return o.SumOpenInterestValue, true
}

// HasSumOpenInterestValue returns a boolean if a field has been set.
func (o *OpenInterestStatisticsResponseInner) HasSumOpenInterestValue() bool {
	if o != nil && !common.IsNil(o.SumOpenInterestValue) {
		return true
	}

	return false
}

// SetSumOpenInterestValue gets a reference to the given string and assigns it to the SumOpenInterestValue field.
func (o *OpenInterestStatisticsResponseInner) SetSumOpenInterestValue(v string) {
	o.SumOpenInterestValue = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *OpenInterestStatisticsResponseInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestStatisticsResponseInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *OpenInterestStatisticsResponseInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *OpenInterestStatisticsResponseInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o OpenInterestStatisticsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenInterestStatisticsResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !common.IsNil(o.ContractType) {
		toSerialize["contractType"] = o.ContractType
	}
	if !common.IsNil(o.SumOpenInterest) {
		toSerialize["sumOpenInterest"] = o.SumOpenInterest
	}
	if !common.IsNil(o.SumOpenInterestValue) {
		toSerialize["sumOpenInterestValue"] = o.SumOpenInterestValue
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenInterestStatisticsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varOpenInterestStatisticsResponseInner := _OpenInterestStatisticsResponseInner{}

	err = json.Unmarshal(data, &varOpenInterestStatisticsResponseInner)

	if err != nil {
		return err
	}

	*o = OpenInterestStatisticsResponseInner(varOpenInterestStatisticsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pair")
		delete(additionalProperties, "contractType")
		delete(additionalProperties, "sumOpenInterest")
		delete(additionalProperties, "sumOpenInterestValue")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenInterestStatisticsResponseInner struct {
	value *OpenInterestStatisticsResponseInner
	isSet bool
}

func (v NullableOpenInterestStatisticsResponseInner) Get() *OpenInterestStatisticsResponseInner {
	return v.value
}

func (v *NullableOpenInterestStatisticsResponseInner) Set(val *OpenInterestStatisticsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenInterestStatisticsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenInterestStatisticsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenInterestStatisticsResponseInner(val *OpenInterestStatisticsResponseInner) *NullableOpenInterestStatisticsResponseInner {
	return &NullableOpenInterestStatisticsResponseInner{value: val, isSet: true}
}

func (v NullableOpenInterestStatisticsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenInterestStatisticsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
