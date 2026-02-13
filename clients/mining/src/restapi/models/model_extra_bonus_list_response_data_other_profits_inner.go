/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExtraBonusListResponseDataOtherProfitsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExtraBonusListResponseDataOtherProfitsInner{}

// ExtraBonusListResponseDataOtherProfitsInner struct for ExtraBonusListResponseDataOtherProfitsInner
type ExtraBonusListResponseDataOtherProfitsInner struct {
	Time                 *int64   `json:"time,omitempty"`
	CoinName             *string  `json:"coinName,omitempty"`
	Type                 *int64   `json:"type,omitempty"`
	ProfitAmount         *float32 `json:"profitAmount,omitempty"`
	Status               *int64   `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExtraBonusListResponseDataOtherProfitsInner ExtraBonusListResponseDataOtherProfitsInner

// NewExtraBonusListResponseDataOtherProfitsInner instantiates a new ExtraBonusListResponseDataOtherProfitsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtraBonusListResponseDataOtherProfitsInner() *ExtraBonusListResponseDataOtherProfitsInner {
	this := ExtraBonusListResponseDataOtherProfitsInner{}
	return &this
}

// NewExtraBonusListResponseDataOtherProfitsInnerWithDefaults instantiates a new ExtraBonusListResponseDataOtherProfitsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtraBonusListResponseDataOtherProfitsInnerWithDefaults() *ExtraBonusListResponseDataOtherProfitsInner {
	this := ExtraBonusListResponseDataOtherProfitsInner{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ExtraBonusListResponseDataOtherProfitsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraBonusListResponseDataOtherProfitsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ExtraBonusListResponseDataOtherProfitsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *ExtraBonusListResponseDataOtherProfitsInner) SetTime(v int64) {
	o.Time = &v
}

// GetCoinName returns the CoinName field value if set, zero value otherwise.
func (o *ExtraBonusListResponseDataOtherProfitsInner) GetCoinName() string {
	if o == nil || common.IsNil(o.CoinName) {
		var ret string
		return ret
	}
	return *o.CoinName
}

// GetCoinNameOk returns a tuple with the CoinName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraBonusListResponseDataOtherProfitsInner) GetCoinNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.CoinName) {
		return nil, false
	}
	return o.CoinName, true
}

// HasCoinName returns a boolean if a field has been set.
func (o *ExtraBonusListResponseDataOtherProfitsInner) HasCoinName() bool {
	if o != nil && !common.IsNil(o.CoinName) {
		return true
	}

	return false
}

// SetCoinName gets a reference to the given string and assigns it to the CoinName field.
func (o *ExtraBonusListResponseDataOtherProfitsInner) SetCoinName(v string) {
	o.CoinName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExtraBonusListResponseDataOtherProfitsInner) GetType() int64 {
	if o == nil || common.IsNil(o.Type) {
		var ret int64
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraBonusListResponseDataOtherProfitsInner) GetTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExtraBonusListResponseDataOtherProfitsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int64 and assigns it to the Type field.
func (o *ExtraBonusListResponseDataOtherProfitsInner) SetType(v int64) {
	o.Type = &v
}

// GetProfitAmount returns the ProfitAmount field value if set, zero value otherwise.
func (o *ExtraBonusListResponseDataOtherProfitsInner) GetProfitAmount() float32 {
	if o == nil || common.IsNil(o.ProfitAmount) {
		var ret float32
		return ret
	}
	return *o.ProfitAmount
}

// GetProfitAmountOk returns a tuple with the ProfitAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraBonusListResponseDataOtherProfitsInner) GetProfitAmountOk() (*float32, bool) {
	if o == nil || common.IsNil(o.ProfitAmount) {
		return nil, false
	}
	return o.ProfitAmount, true
}

// HasProfitAmount returns a boolean if a field has been set.
func (o *ExtraBonusListResponseDataOtherProfitsInner) HasProfitAmount() bool {
	if o != nil && !common.IsNil(o.ProfitAmount) {
		return true
	}

	return false
}

// SetProfitAmount gets a reference to the given float32 and assigns it to the ProfitAmount field.
func (o *ExtraBonusListResponseDataOtherProfitsInner) SetProfitAmount(v float32) {
	o.ProfitAmount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExtraBonusListResponseDataOtherProfitsInner) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraBonusListResponseDataOtherProfitsInner) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExtraBonusListResponseDataOtherProfitsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *ExtraBonusListResponseDataOtherProfitsInner) SetStatus(v int64) {
	o.Status = &v
}

func (o ExtraBonusListResponseDataOtherProfitsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtraBonusListResponseDataOtherProfitsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.CoinName) {
		toSerialize["coinName"] = o.CoinName
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.ProfitAmount) {
		toSerialize["profitAmount"] = o.ProfitAmount
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExtraBonusListResponseDataOtherProfitsInner) UnmarshalJSON(data []byte) (err error) {
	varExtraBonusListResponseDataOtherProfitsInner := _ExtraBonusListResponseDataOtherProfitsInner{}

	err = json.Unmarshal(data, &varExtraBonusListResponseDataOtherProfitsInner)

	if err != nil {
		return err
	}

	*o = ExtraBonusListResponseDataOtherProfitsInner(varExtraBonusListResponseDataOtherProfitsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "coinName")
		delete(additionalProperties, "type")
		delete(additionalProperties, "profitAmount")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExtraBonusListResponseDataOtherProfitsInner struct {
	value *ExtraBonusListResponseDataOtherProfitsInner
	isSet bool
}

func (v NullableExtraBonusListResponseDataOtherProfitsInner) Get() *ExtraBonusListResponseDataOtherProfitsInner {
	return v.value
}

func (v *NullableExtraBonusListResponseDataOtherProfitsInner) Set(val *ExtraBonusListResponseDataOtherProfitsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExtraBonusListResponseDataOtherProfitsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExtraBonusListResponseDataOtherProfitsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtraBonusListResponseDataOtherProfitsInner(val *ExtraBonusListResponseDataOtherProfitsInner) *NullableExtraBonusListResponseDataOtherProfitsInner {
	return &NullableExtraBonusListResponseDataOtherProfitsInner{value: val, isSet: true}
}

func (v NullableExtraBonusListResponseDataOtherProfitsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtraBonusListResponseDataOtherProfitsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
