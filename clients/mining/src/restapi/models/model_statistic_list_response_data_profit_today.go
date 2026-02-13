/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the StatisticListResponseDataProfitToday type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StatisticListResponseDataProfitToday{}

// StatisticListResponseDataProfitToday struct for StatisticListResponseDataProfitToday
type StatisticListResponseDataProfitToday struct {
	BTC                  *string `json:"BTC,omitempty"`
	BSV                  *string `json:"BSV,omitempty"`
	BCH                  *string `json:"BCH,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StatisticListResponseDataProfitToday StatisticListResponseDataProfitToday

// NewStatisticListResponseDataProfitToday instantiates a new StatisticListResponseDataProfitToday object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticListResponseDataProfitToday() *StatisticListResponseDataProfitToday {
	this := StatisticListResponseDataProfitToday{}
	return &this
}

// NewStatisticListResponseDataProfitTodayWithDefaults instantiates a new StatisticListResponseDataProfitToday object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticListResponseDataProfitTodayWithDefaults() *StatisticListResponseDataProfitToday {
	this := StatisticListResponseDataProfitToday{}
	return &this
}

// GetBTC returns the BTC field value if set, zero value otherwise.
func (o *StatisticListResponseDataProfitToday) GetBTC() string {
	if o == nil || common.IsNil(o.BTC) {
		var ret string
		return ret
	}
	return *o.BTC
}

// GetBTCOk returns a tuple with the BTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticListResponseDataProfitToday) GetBTCOk() (*string, bool) {
	if o == nil || common.IsNil(o.BTC) {
		return nil, false
	}
	return o.BTC, true
}

// HasBTC returns a boolean if a field has been set.
func (o *StatisticListResponseDataProfitToday) HasBTC() bool {
	if o != nil && !common.IsNil(o.BTC) {
		return true
	}

	return false
}

// SetBTC gets a reference to the given string and assigns it to the BTC field.
func (o *StatisticListResponseDataProfitToday) SetBTC(v string) {
	o.BTC = &v
}

// GetBSV returns the BSV field value if set, zero value otherwise.
func (o *StatisticListResponseDataProfitToday) GetBSV() string {
	if o == nil || common.IsNil(o.BSV) {
		var ret string
		return ret
	}
	return *o.BSV
}

// GetBSVOk returns a tuple with the BSV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticListResponseDataProfitToday) GetBSVOk() (*string, bool) {
	if o == nil || common.IsNil(o.BSV) {
		return nil, false
	}
	return o.BSV, true
}

// HasBSV returns a boolean if a field has been set.
func (o *StatisticListResponseDataProfitToday) HasBSV() bool {
	if o != nil && !common.IsNil(o.BSV) {
		return true
	}

	return false
}

// SetBSV gets a reference to the given string and assigns it to the BSV field.
func (o *StatisticListResponseDataProfitToday) SetBSV(v string) {
	o.BSV = &v
}

// GetBCH returns the BCH field value if set, zero value otherwise.
func (o *StatisticListResponseDataProfitToday) GetBCH() string {
	if o == nil || common.IsNil(o.BCH) {
		var ret string
		return ret
	}
	return *o.BCH
}

// GetBCHOk returns a tuple with the BCH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticListResponseDataProfitToday) GetBCHOk() (*string, bool) {
	if o == nil || common.IsNil(o.BCH) {
		return nil, false
	}
	return o.BCH, true
}

// HasBCH returns a boolean if a field has been set.
func (o *StatisticListResponseDataProfitToday) HasBCH() bool {
	if o != nil && !common.IsNil(o.BCH) {
		return true
	}

	return false
}

// SetBCH gets a reference to the given string and assigns it to the BCH field.
func (o *StatisticListResponseDataProfitToday) SetBCH(v string) {
	o.BCH = &v
}

func (o StatisticListResponseDataProfitToday) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatisticListResponseDataProfitToday) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BTC) {
		toSerialize["BTC"] = o.BTC
	}
	if !common.IsNil(o.BSV) {
		toSerialize["BSV"] = o.BSV
	}
	if !common.IsNil(o.BCH) {
		toSerialize["BCH"] = o.BCH
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StatisticListResponseDataProfitToday) UnmarshalJSON(data []byte) (err error) {
	varStatisticListResponseDataProfitToday := _StatisticListResponseDataProfitToday{}

	err = json.Unmarshal(data, &varStatisticListResponseDataProfitToday)

	if err != nil {
		return err
	}

	*o = StatisticListResponseDataProfitToday(varStatisticListResponseDataProfitToday)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "BTC")
		delete(additionalProperties, "BSV")
		delete(additionalProperties, "BCH")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStatisticListResponseDataProfitToday struct {
	value *StatisticListResponseDataProfitToday
	isSet bool
}

func (v NullableStatisticListResponseDataProfitToday) Get() *StatisticListResponseDataProfitToday {
	return v.value
}

func (v *NullableStatisticListResponseDataProfitToday) Set(val *StatisticListResponseDataProfitToday) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticListResponseDataProfitToday) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticListResponseDataProfitToday) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticListResponseDataProfitToday(val *StatisticListResponseDataProfitToday) *NullableStatisticListResponseDataProfitToday {
	return &NullableStatisticListResponseDataProfitToday{value: val, isSet: true}
}

func (v NullableStatisticListResponseDataProfitToday) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticListResponseDataProfitToday) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
