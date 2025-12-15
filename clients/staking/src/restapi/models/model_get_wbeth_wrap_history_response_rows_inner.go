/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetWbethWrapHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetWbethWrapHistoryResponseRowsInner{}

// GetWbethWrapHistoryResponseRowsInner struct for GetWbethWrapHistoryResponseRowsInner
type GetWbethWrapHistoryResponseRowsInner struct {
	Time                 *int64  `json:"time,omitempty"`
	FromAsset            *string `json:"fromAsset,omitempty"`
	FromAmount           *string `json:"fromAmount,omitempty"`
	ToAsset              *string `json:"toAsset,omitempty"`
	ToAmount             *string `json:"toAmount,omitempty"`
	ExchangeRate         *string `json:"exchangeRate,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetWbethWrapHistoryResponseRowsInner GetWbethWrapHistoryResponseRowsInner

// NewGetWbethWrapHistoryResponseRowsInner instantiates a new GetWbethWrapHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWbethWrapHistoryResponseRowsInner() *GetWbethWrapHistoryResponseRowsInner {
	this := GetWbethWrapHistoryResponseRowsInner{}
	return &this
}

// NewGetWbethWrapHistoryResponseRowsInnerWithDefaults instantiates a new GetWbethWrapHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWbethWrapHistoryResponseRowsInnerWithDefaults() *GetWbethWrapHistoryResponseRowsInner {
	this := GetWbethWrapHistoryResponseRowsInner{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetWbethWrapHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethWrapHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetWbethWrapHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetWbethWrapHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetFromAsset returns the FromAsset field value if set, zero value otherwise.
func (o *GetWbethWrapHistoryResponseRowsInner) GetFromAsset() string {
	if o == nil || common.IsNil(o.FromAsset) {
		var ret string
		return ret
	}
	return *o.FromAsset
}

// GetFromAssetOk returns a tuple with the FromAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethWrapHistoryResponseRowsInner) GetFromAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAsset) {
		return nil, false
	}
	return o.FromAsset, true
}

// HasFromAsset returns a boolean if a field has been set.
func (o *GetWbethWrapHistoryResponseRowsInner) HasFromAsset() bool {
	if o != nil && !common.IsNil(o.FromAsset) {
		return true
	}

	return false
}

// SetFromAsset gets a reference to the given string and assigns it to the FromAsset field.
func (o *GetWbethWrapHistoryResponseRowsInner) SetFromAsset(v string) {
	o.FromAsset = &v
}

// GetFromAmount returns the FromAmount field value if set, zero value otherwise.
func (o *GetWbethWrapHistoryResponseRowsInner) GetFromAmount() string {
	if o == nil || common.IsNil(o.FromAmount) {
		var ret string
		return ret
	}
	return *o.FromAmount
}

// GetFromAmountOk returns a tuple with the FromAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethWrapHistoryResponseRowsInner) GetFromAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAmount) {
		return nil, false
	}
	return o.FromAmount, true
}

// HasFromAmount returns a boolean if a field has been set.
func (o *GetWbethWrapHistoryResponseRowsInner) HasFromAmount() bool {
	if o != nil && !common.IsNil(o.FromAmount) {
		return true
	}

	return false
}

// SetFromAmount gets a reference to the given string and assigns it to the FromAmount field.
func (o *GetWbethWrapHistoryResponseRowsInner) SetFromAmount(v string) {
	o.FromAmount = &v
}

// GetToAsset returns the ToAsset field value if set, zero value otherwise.
func (o *GetWbethWrapHistoryResponseRowsInner) GetToAsset() string {
	if o == nil || common.IsNil(o.ToAsset) {
		var ret string
		return ret
	}
	return *o.ToAsset
}

// GetToAssetOk returns a tuple with the ToAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethWrapHistoryResponseRowsInner) GetToAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAsset) {
		return nil, false
	}
	return o.ToAsset, true
}

// HasToAsset returns a boolean if a field has been set.
func (o *GetWbethWrapHistoryResponseRowsInner) HasToAsset() bool {
	if o != nil && !common.IsNil(o.ToAsset) {
		return true
	}

	return false
}

// SetToAsset gets a reference to the given string and assigns it to the ToAsset field.
func (o *GetWbethWrapHistoryResponseRowsInner) SetToAsset(v string) {
	o.ToAsset = &v
}

// GetToAmount returns the ToAmount field value if set, zero value otherwise.
func (o *GetWbethWrapHistoryResponseRowsInner) GetToAmount() string {
	if o == nil || common.IsNil(o.ToAmount) {
		var ret string
		return ret
	}
	return *o.ToAmount
}

// GetToAmountOk returns a tuple with the ToAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethWrapHistoryResponseRowsInner) GetToAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAmount) {
		return nil, false
	}
	return o.ToAmount, true
}

// HasToAmount returns a boolean if a field has been set.
func (o *GetWbethWrapHistoryResponseRowsInner) HasToAmount() bool {
	if o != nil && !common.IsNil(o.ToAmount) {
		return true
	}

	return false
}

// SetToAmount gets a reference to the given string and assigns it to the ToAmount field.
func (o *GetWbethWrapHistoryResponseRowsInner) SetToAmount(v string) {
	o.ToAmount = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *GetWbethWrapHistoryResponseRowsInner) GetExchangeRate() string {
	if o == nil || common.IsNil(o.ExchangeRate) {
		var ret string
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethWrapHistoryResponseRowsInner) GetExchangeRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *GetWbethWrapHistoryResponseRowsInner) HasExchangeRate() bool {
	if o != nil && !common.IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given string and assigns it to the ExchangeRate field.
func (o *GetWbethWrapHistoryResponseRowsInner) SetExchangeRate(v string) {
	o.ExchangeRate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetWbethWrapHistoryResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethWrapHistoryResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetWbethWrapHistoryResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetWbethWrapHistoryResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetWbethWrapHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWbethWrapHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.FromAsset) {
		toSerialize["fromAsset"] = o.FromAsset
	}
	if !common.IsNil(o.FromAmount) {
		toSerialize["fromAmount"] = o.FromAmount
	}
	if !common.IsNil(o.ToAsset) {
		toSerialize["toAsset"] = o.ToAsset
	}
	if !common.IsNil(o.ToAmount) {
		toSerialize["toAmount"] = o.ToAmount
	}
	if !common.IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetWbethWrapHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetWbethWrapHistoryResponseRowsInner := _GetWbethWrapHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetWbethWrapHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetWbethWrapHistoryResponseRowsInner(varGetWbethWrapHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "fromAsset")
		delete(additionalProperties, "fromAmount")
		delete(additionalProperties, "toAsset")
		delete(additionalProperties, "toAmount")
		delete(additionalProperties, "exchangeRate")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetWbethWrapHistoryResponseRowsInner struct {
	value *GetWbethWrapHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetWbethWrapHistoryResponseRowsInner) Get() *GetWbethWrapHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetWbethWrapHistoryResponseRowsInner) Set(val *GetWbethWrapHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWbethWrapHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWbethWrapHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWbethWrapHistoryResponseRowsInner(val *GetWbethWrapHistoryResponseRowsInner) *NullableGetWbethWrapHistoryResponseRowsInner {
	return &NullableGetWbethWrapHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetWbethWrapHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWbethWrapHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
