/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the EarningsListResponseDataAccountProfitsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EarningsListResponseDataAccountProfitsInner{}

// EarningsListResponseDataAccountProfitsInner struct for EarningsListResponseDataAccountProfitsInner
type EarningsListResponseDataAccountProfitsInner struct {
	Time                 *int64   `json:"time,omitempty"`
	Type                 *int64   `json:"type,omitempty"`
	HashTransfer         *int64   `json:"hashTransfer,omitempty"`
	TransferAmount       *float32 `json:"transferAmount,omitempty"`
	DayHashRate          *int64   `json:"dayHashRate,omitempty"`
	ProfitAmount         *float32 `json:"profitAmount,omitempty"`
	CoinName             *string  `json:"coinName,omitempty"`
	Status               *int64   `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EarningsListResponseDataAccountProfitsInner EarningsListResponseDataAccountProfitsInner

// NewEarningsListResponseDataAccountProfitsInner instantiates a new EarningsListResponseDataAccountProfitsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEarningsListResponseDataAccountProfitsInner() *EarningsListResponseDataAccountProfitsInner {
	this := EarningsListResponseDataAccountProfitsInner{}
	return &this
}

// NewEarningsListResponseDataAccountProfitsInnerWithDefaults instantiates a new EarningsListResponseDataAccountProfitsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEarningsListResponseDataAccountProfitsInnerWithDefaults() *EarningsListResponseDataAccountProfitsInner {
	this := EarningsListResponseDataAccountProfitsInner{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *EarningsListResponseDataAccountProfitsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsListResponseDataAccountProfitsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *EarningsListResponseDataAccountProfitsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *EarningsListResponseDataAccountProfitsInner) SetTime(v int64) {
	o.Time = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EarningsListResponseDataAccountProfitsInner) GetType() int64 {
	if o == nil || common.IsNil(o.Type) {
		var ret int64
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsListResponseDataAccountProfitsInner) GetTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EarningsListResponseDataAccountProfitsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int64 and assigns it to the Type field.
func (o *EarningsListResponseDataAccountProfitsInner) SetType(v int64) {
	o.Type = &v
}

// GetHashTransfer returns the HashTransfer field value if set, zero value otherwise.
func (o *EarningsListResponseDataAccountProfitsInner) GetHashTransfer() int64 {
	if o == nil || common.IsNil(o.HashTransfer) {
		var ret int64
		return ret
	}
	return *o.HashTransfer
}

// GetHashTransferOk returns a tuple with the HashTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsListResponseDataAccountProfitsInner) GetHashTransferOk() (*int64, bool) {
	if o == nil || common.IsNil(o.HashTransfer) {
		return nil, false
	}
	return o.HashTransfer, true
}

// HasHashTransfer returns a boolean if a field has been set.
func (o *EarningsListResponseDataAccountProfitsInner) HasHashTransfer() bool {
	if o != nil && !common.IsNil(o.HashTransfer) {
		return true
	}

	return false
}

// SetHashTransfer gets a reference to the given int64 and assigns it to the HashTransfer field.
func (o *EarningsListResponseDataAccountProfitsInner) SetHashTransfer(v int64) {
	o.HashTransfer = &v
}

// GetTransferAmount returns the TransferAmount field value if set, zero value otherwise.
func (o *EarningsListResponseDataAccountProfitsInner) GetTransferAmount() float32 {
	if o == nil || common.IsNil(o.TransferAmount) {
		var ret float32
		return ret
	}
	return *o.TransferAmount
}

// GetTransferAmountOk returns a tuple with the TransferAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsListResponseDataAccountProfitsInner) GetTransferAmountOk() (*float32, bool) {
	if o == nil || common.IsNil(o.TransferAmount) {
		return nil, false
	}
	return o.TransferAmount, true
}

// HasTransferAmount returns a boolean if a field has been set.
func (o *EarningsListResponseDataAccountProfitsInner) HasTransferAmount() bool {
	if o != nil && !common.IsNil(o.TransferAmount) {
		return true
	}

	return false
}

// SetTransferAmount gets a reference to the given float32 and assigns it to the TransferAmount field.
func (o *EarningsListResponseDataAccountProfitsInner) SetTransferAmount(v float32) {
	o.TransferAmount = &v
}

// GetDayHashRate returns the DayHashRate field value if set, zero value otherwise.
func (o *EarningsListResponseDataAccountProfitsInner) GetDayHashRate() int64 {
	if o == nil || common.IsNil(o.DayHashRate) {
		var ret int64
		return ret
	}
	return *o.DayHashRate
}

// GetDayHashRateOk returns a tuple with the DayHashRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsListResponseDataAccountProfitsInner) GetDayHashRateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.DayHashRate) {
		return nil, false
	}
	return o.DayHashRate, true
}

// HasDayHashRate returns a boolean if a field has been set.
func (o *EarningsListResponseDataAccountProfitsInner) HasDayHashRate() bool {
	if o != nil && !common.IsNil(o.DayHashRate) {
		return true
	}

	return false
}

// SetDayHashRate gets a reference to the given int64 and assigns it to the DayHashRate field.
func (o *EarningsListResponseDataAccountProfitsInner) SetDayHashRate(v int64) {
	o.DayHashRate = &v
}

// GetProfitAmount returns the ProfitAmount field value if set, zero value otherwise.
func (o *EarningsListResponseDataAccountProfitsInner) GetProfitAmount() float32 {
	if o == nil || common.IsNil(o.ProfitAmount) {
		var ret float32
		return ret
	}
	return *o.ProfitAmount
}

// GetProfitAmountOk returns a tuple with the ProfitAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsListResponseDataAccountProfitsInner) GetProfitAmountOk() (*float32, bool) {
	if o == nil || common.IsNil(o.ProfitAmount) {
		return nil, false
	}
	return o.ProfitAmount, true
}

// HasProfitAmount returns a boolean if a field has been set.
func (o *EarningsListResponseDataAccountProfitsInner) HasProfitAmount() bool {
	if o != nil && !common.IsNil(o.ProfitAmount) {
		return true
	}

	return false
}

// SetProfitAmount gets a reference to the given float32 and assigns it to the ProfitAmount field.
func (o *EarningsListResponseDataAccountProfitsInner) SetProfitAmount(v float32) {
	o.ProfitAmount = &v
}

// GetCoinName returns the CoinName field value if set, zero value otherwise.
func (o *EarningsListResponseDataAccountProfitsInner) GetCoinName() string {
	if o == nil || common.IsNil(o.CoinName) {
		var ret string
		return ret
	}
	return *o.CoinName
}

// GetCoinNameOk returns a tuple with the CoinName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsListResponseDataAccountProfitsInner) GetCoinNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.CoinName) {
		return nil, false
	}
	return o.CoinName, true
}

// HasCoinName returns a boolean if a field has been set.
func (o *EarningsListResponseDataAccountProfitsInner) HasCoinName() bool {
	if o != nil && !common.IsNil(o.CoinName) {
		return true
	}

	return false
}

// SetCoinName gets a reference to the given string and assigns it to the CoinName field.
func (o *EarningsListResponseDataAccountProfitsInner) SetCoinName(v string) {
	o.CoinName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EarningsListResponseDataAccountProfitsInner) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsListResponseDataAccountProfitsInner) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EarningsListResponseDataAccountProfitsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *EarningsListResponseDataAccountProfitsInner) SetStatus(v int64) {
	o.Status = &v
}

func (o EarningsListResponseDataAccountProfitsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EarningsListResponseDataAccountProfitsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.HashTransfer) {
		toSerialize["hashTransfer"] = o.HashTransfer
	}
	if !common.IsNil(o.TransferAmount) {
		toSerialize["transferAmount"] = o.TransferAmount
	}
	if !common.IsNil(o.DayHashRate) {
		toSerialize["dayHashRate"] = o.DayHashRate
	}
	if !common.IsNil(o.ProfitAmount) {
		toSerialize["profitAmount"] = o.ProfitAmount
	}
	if !common.IsNil(o.CoinName) {
		toSerialize["coinName"] = o.CoinName
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EarningsListResponseDataAccountProfitsInner) UnmarshalJSON(data []byte) (err error) {
	varEarningsListResponseDataAccountProfitsInner := _EarningsListResponseDataAccountProfitsInner{}

	err = json.Unmarshal(data, &varEarningsListResponseDataAccountProfitsInner)

	if err != nil {
		return err
	}

	*o = EarningsListResponseDataAccountProfitsInner(varEarningsListResponseDataAccountProfitsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "type")
		delete(additionalProperties, "hashTransfer")
		delete(additionalProperties, "transferAmount")
		delete(additionalProperties, "dayHashRate")
		delete(additionalProperties, "profitAmount")
		delete(additionalProperties, "coinName")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEarningsListResponseDataAccountProfitsInner struct {
	value *EarningsListResponseDataAccountProfitsInner
	isSet bool
}

func (v NullableEarningsListResponseDataAccountProfitsInner) Get() *EarningsListResponseDataAccountProfitsInner {
	return v.value
}

func (v *NullableEarningsListResponseDataAccountProfitsInner) Set(val *EarningsListResponseDataAccountProfitsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEarningsListResponseDataAccountProfitsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEarningsListResponseDataAccountProfitsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEarningsListResponseDataAccountProfitsInner(val *EarningsListResponseDataAccountProfitsInner) *NullableEarningsListResponseDataAccountProfitsInner {
	return &NullableEarningsListResponseDataAccountProfitsInner{value: val, isSet: true}
}

func (v NullableEarningsListResponseDataAccountProfitsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEarningsListResponseDataAccountProfitsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
