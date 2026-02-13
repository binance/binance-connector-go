/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetWbethUnwrapHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetWbethUnwrapHistoryResponseRowsInner{}

// GetWbethUnwrapHistoryResponseRowsInner struct for GetWbethUnwrapHistoryResponseRowsInner
type GetWbethUnwrapHistoryResponseRowsInner struct {
	Time                 *int64  `json:"time,omitempty"`
	FromAsset            *string `json:"fromAsset,omitempty"`
	FromAmount           *string `json:"fromAmount,omitempty"`
	ToAsset              *string `json:"toAsset,omitempty"`
	ToAmount             *string `json:"toAmount,omitempty"`
	ExchangeRate         *string `json:"exchangeRate,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetWbethUnwrapHistoryResponseRowsInner GetWbethUnwrapHistoryResponseRowsInner

// NewGetWbethUnwrapHistoryResponseRowsInner instantiates a new GetWbethUnwrapHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWbethUnwrapHistoryResponseRowsInner() *GetWbethUnwrapHistoryResponseRowsInner {
	this := GetWbethUnwrapHistoryResponseRowsInner{}
	return &this
}

// NewGetWbethUnwrapHistoryResponseRowsInnerWithDefaults instantiates a new GetWbethUnwrapHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWbethUnwrapHistoryResponseRowsInnerWithDefaults() *GetWbethUnwrapHistoryResponseRowsInner {
	this := GetWbethUnwrapHistoryResponseRowsInner{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetWbethUnwrapHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethUnwrapHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetWbethUnwrapHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetWbethUnwrapHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetFromAsset returns the FromAsset field value if set, zero value otherwise.
func (o *GetWbethUnwrapHistoryResponseRowsInner) GetFromAsset() string {
	if o == nil || common.IsNil(o.FromAsset) {
		var ret string
		return ret
	}
	return *o.FromAsset
}

// GetFromAssetOk returns a tuple with the FromAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethUnwrapHistoryResponseRowsInner) GetFromAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAsset) {
		return nil, false
	}
	return o.FromAsset, true
}

// HasFromAsset returns a boolean if a field has been set.
func (o *GetWbethUnwrapHistoryResponseRowsInner) HasFromAsset() bool {
	if o != nil && !common.IsNil(o.FromAsset) {
		return true
	}

	return false
}

// SetFromAsset gets a reference to the given string and assigns it to the FromAsset field.
func (o *GetWbethUnwrapHistoryResponseRowsInner) SetFromAsset(v string) {
	o.FromAsset = &v
}

// GetFromAmount returns the FromAmount field value if set, zero value otherwise.
func (o *GetWbethUnwrapHistoryResponseRowsInner) GetFromAmount() string {
	if o == nil || common.IsNil(o.FromAmount) {
		var ret string
		return ret
	}
	return *o.FromAmount
}

// GetFromAmountOk returns a tuple with the FromAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethUnwrapHistoryResponseRowsInner) GetFromAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAmount) {
		return nil, false
	}
	return o.FromAmount, true
}

// HasFromAmount returns a boolean if a field has been set.
func (o *GetWbethUnwrapHistoryResponseRowsInner) HasFromAmount() bool {
	if o != nil && !common.IsNil(o.FromAmount) {
		return true
	}

	return false
}

// SetFromAmount gets a reference to the given string and assigns it to the FromAmount field.
func (o *GetWbethUnwrapHistoryResponseRowsInner) SetFromAmount(v string) {
	o.FromAmount = &v
}

// GetToAsset returns the ToAsset field value if set, zero value otherwise.
func (o *GetWbethUnwrapHistoryResponseRowsInner) GetToAsset() string {
	if o == nil || common.IsNil(o.ToAsset) {
		var ret string
		return ret
	}
	return *o.ToAsset
}

// GetToAssetOk returns a tuple with the ToAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethUnwrapHistoryResponseRowsInner) GetToAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAsset) {
		return nil, false
	}
	return o.ToAsset, true
}

// HasToAsset returns a boolean if a field has been set.
func (o *GetWbethUnwrapHistoryResponseRowsInner) HasToAsset() bool {
	if o != nil && !common.IsNil(o.ToAsset) {
		return true
	}

	return false
}

// SetToAsset gets a reference to the given string and assigns it to the ToAsset field.
func (o *GetWbethUnwrapHistoryResponseRowsInner) SetToAsset(v string) {
	o.ToAsset = &v
}

// GetToAmount returns the ToAmount field value if set, zero value otherwise.
func (o *GetWbethUnwrapHistoryResponseRowsInner) GetToAmount() string {
	if o == nil || common.IsNil(o.ToAmount) {
		var ret string
		return ret
	}
	return *o.ToAmount
}

// GetToAmountOk returns a tuple with the ToAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethUnwrapHistoryResponseRowsInner) GetToAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAmount) {
		return nil, false
	}
	return o.ToAmount, true
}

// HasToAmount returns a boolean if a field has been set.
func (o *GetWbethUnwrapHistoryResponseRowsInner) HasToAmount() bool {
	if o != nil && !common.IsNil(o.ToAmount) {
		return true
	}

	return false
}

// SetToAmount gets a reference to the given string and assigns it to the ToAmount field.
func (o *GetWbethUnwrapHistoryResponseRowsInner) SetToAmount(v string) {
	o.ToAmount = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *GetWbethUnwrapHistoryResponseRowsInner) GetExchangeRate() string {
	if o == nil || common.IsNil(o.ExchangeRate) {
		var ret string
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethUnwrapHistoryResponseRowsInner) GetExchangeRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *GetWbethUnwrapHistoryResponseRowsInner) HasExchangeRate() bool {
	if o != nil && !common.IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given string and assigns it to the ExchangeRate field.
func (o *GetWbethUnwrapHistoryResponseRowsInner) SetExchangeRate(v string) {
	o.ExchangeRate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetWbethUnwrapHistoryResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethUnwrapHistoryResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetWbethUnwrapHistoryResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetWbethUnwrapHistoryResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetWbethUnwrapHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWbethUnwrapHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
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

func (o *GetWbethUnwrapHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetWbethUnwrapHistoryResponseRowsInner := _GetWbethUnwrapHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetWbethUnwrapHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetWbethUnwrapHistoryResponseRowsInner(varGetWbethUnwrapHistoryResponseRowsInner)

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

type NullableGetWbethUnwrapHistoryResponseRowsInner struct {
	value *GetWbethUnwrapHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetWbethUnwrapHistoryResponseRowsInner) Get() *GetWbethUnwrapHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetWbethUnwrapHistoryResponseRowsInner) Set(val *GetWbethUnwrapHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWbethUnwrapHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWbethUnwrapHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWbethUnwrapHistoryResponseRowsInner(val *GetWbethUnwrapHistoryResponseRowsInner) *NullableGetWbethUnwrapHistoryResponseRowsInner {
	return &NullableGetWbethUnwrapHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetWbethUnwrapHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWbethUnwrapHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
