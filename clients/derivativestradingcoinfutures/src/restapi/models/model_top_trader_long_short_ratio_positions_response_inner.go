/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TopTraderLongShortRatioPositionsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TopTraderLongShortRatioPositionsResponseInner{}

// TopTraderLongShortRatioPositionsResponseInner struct for TopTraderLongShortRatioPositionsResponseInner
type TopTraderLongShortRatioPositionsResponseInner struct {
	Pair                 *string `json:"pair,omitempty"`
	LongShortRatio       *string `json:"longShortRatio,omitempty"`
	LongPosition         *string `json:"longPosition,omitempty"`
	ShortPosition        *string `json:"shortPosition,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TopTraderLongShortRatioPositionsResponseInner TopTraderLongShortRatioPositionsResponseInner

// NewTopTraderLongShortRatioPositionsResponseInner instantiates a new TopTraderLongShortRatioPositionsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopTraderLongShortRatioPositionsResponseInner() *TopTraderLongShortRatioPositionsResponseInner {
	this := TopTraderLongShortRatioPositionsResponseInner{}
	return &this
}

// NewTopTraderLongShortRatioPositionsResponseInnerWithDefaults instantiates a new TopTraderLongShortRatioPositionsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopTraderLongShortRatioPositionsResponseInnerWithDefaults() *TopTraderLongShortRatioPositionsResponseInner {
	this := TopTraderLongShortRatioPositionsResponseInner{}
	return &this
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetPair() string {
	if o == nil || common.IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetPairOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) HasPair() bool {
	if o != nil && !common.IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *TopTraderLongShortRatioPositionsResponseInner) SetPair(v string) {
	o.Pair = &v
}

// GetLongShortRatio returns the LongShortRatio field value if set, zero value otherwise.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetLongShortRatio() string {
	if o == nil || common.IsNil(o.LongShortRatio) {
		var ret string
		return ret
	}
	return *o.LongShortRatio
}

// GetLongShortRatioOk returns a tuple with the LongShortRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetLongShortRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.LongShortRatio) {
		return nil, false
	}
	return o.LongShortRatio, true
}

// HasLongShortRatio returns a boolean if a field has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) HasLongShortRatio() bool {
	if o != nil && !common.IsNil(o.LongShortRatio) {
		return true
	}

	return false
}

// SetLongShortRatio gets a reference to the given string and assigns it to the LongShortRatio field.
func (o *TopTraderLongShortRatioPositionsResponseInner) SetLongShortRatio(v string) {
	o.LongShortRatio = &v
}

// GetLongPosition returns the LongPosition field value if set, zero value otherwise.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetLongPosition() string {
	if o == nil || common.IsNil(o.LongPosition) {
		var ret string
		return ret
	}
	return *o.LongPosition
}

// GetLongPositionOk returns a tuple with the LongPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetLongPositionOk() (*string, bool) {
	if o == nil || common.IsNil(o.LongPosition) {
		return nil, false
	}
	return o.LongPosition, true
}

// HasLongPosition returns a boolean if a field has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) HasLongPosition() bool {
	if o != nil && !common.IsNil(o.LongPosition) {
		return true
	}

	return false
}

// SetLongPosition gets a reference to the given string and assigns it to the LongPosition field.
func (o *TopTraderLongShortRatioPositionsResponseInner) SetLongPosition(v string) {
	o.LongPosition = &v
}

// GetShortPosition returns the ShortPosition field value if set, zero value otherwise.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetShortPosition() string {
	if o == nil || common.IsNil(o.ShortPosition) {
		var ret string
		return ret
	}
	return *o.ShortPosition
}

// GetShortPositionOk returns a tuple with the ShortPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetShortPositionOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShortPosition) {
		return nil, false
	}
	return o.ShortPosition, true
}

// HasShortPosition returns a boolean if a field has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) HasShortPosition() bool {
	if o != nil && !common.IsNil(o.ShortPosition) {
		return true
	}

	return false
}

// SetShortPosition gets a reference to the given string and assigns it to the ShortPosition field.
func (o *TopTraderLongShortRatioPositionsResponseInner) SetShortPosition(v string) {
	o.ShortPosition = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TopTraderLongShortRatioPositionsResponseInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *TopTraderLongShortRatioPositionsResponseInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o TopTraderLongShortRatioPositionsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopTraderLongShortRatioPositionsResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !common.IsNil(o.LongShortRatio) {
		toSerialize["longShortRatio"] = o.LongShortRatio
	}
	if !common.IsNil(o.LongPosition) {
		toSerialize["longPosition"] = o.LongPosition
	}
	if !common.IsNil(o.ShortPosition) {
		toSerialize["shortPosition"] = o.ShortPosition
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TopTraderLongShortRatioPositionsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varTopTraderLongShortRatioPositionsResponseInner := _TopTraderLongShortRatioPositionsResponseInner{}

	err = json.Unmarshal(data, &varTopTraderLongShortRatioPositionsResponseInner)

	if err != nil {
		return err
	}

	*o = TopTraderLongShortRatioPositionsResponseInner(varTopTraderLongShortRatioPositionsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pair")
		delete(additionalProperties, "longShortRatio")
		delete(additionalProperties, "longPosition")
		delete(additionalProperties, "shortPosition")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTopTraderLongShortRatioPositionsResponseInner struct {
	value *TopTraderLongShortRatioPositionsResponseInner
	isSet bool
}

func (v NullableTopTraderLongShortRatioPositionsResponseInner) Get() *TopTraderLongShortRatioPositionsResponseInner {
	return v.value
}

func (v *NullableTopTraderLongShortRatioPositionsResponseInner) Set(val *TopTraderLongShortRatioPositionsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTopTraderLongShortRatioPositionsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTopTraderLongShortRatioPositionsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopTraderLongShortRatioPositionsResponseInner(val *TopTraderLongShortRatioPositionsResponseInner) *NullableTopTraderLongShortRatioPositionsResponseInner {
	return &NullableTopTraderLongShortRatioPositionsResponseInner{value: val, isSet: true}
}

func (v NullableTopTraderLongShortRatioPositionsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopTraderLongShortRatioPositionsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
