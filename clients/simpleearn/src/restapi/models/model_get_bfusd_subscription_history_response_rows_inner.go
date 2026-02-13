/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetBfusdSubscriptionHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBfusdSubscriptionHistoryResponseRowsInner{}

// GetBfusdSubscriptionHistoryResponseRowsInner struct for GetBfusdSubscriptionHistoryResponseRowsInner
type GetBfusdSubscriptionHistoryResponseRowsInner struct {
	Time                 *int64  `json:"time,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	ReceiveAsset         *string `json:"receiveAsset,omitempty"`
	ReceiveAmount        *string `json:"receiveAmount,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBfusdSubscriptionHistoryResponseRowsInner GetBfusdSubscriptionHistoryResponseRowsInner

// NewGetBfusdSubscriptionHistoryResponseRowsInner instantiates a new GetBfusdSubscriptionHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBfusdSubscriptionHistoryResponseRowsInner() *GetBfusdSubscriptionHistoryResponseRowsInner {
	this := GetBfusdSubscriptionHistoryResponseRowsInner{}
	return &this
}

// NewGetBfusdSubscriptionHistoryResponseRowsInnerWithDefaults instantiates a new GetBfusdSubscriptionHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBfusdSubscriptionHistoryResponseRowsInnerWithDefaults() *GetBfusdSubscriptionHistoryResponseRowsInner {
	this := GetBfusdSubscriptionHistoryResponseRowsInner{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetReceiveAsset returns the ReceiveAsset field value if set, zero value otherwise.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) GetReceiveAsset() string {
	if o == nil || common.IsNil(o.ReceiveAsset) {
		var ret string
		return ret
	}
	return *o.ReceiveAsset
}

// GetReceiveAssetOk returns a tuple with the ReceiveAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) GetReceiveAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReceiveAsset) {
		return nil, false
	}
	return o.ReceiveAsset, true
}

// HasReceiveAsset returns a boolean if a field has been set.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) HasReceiveAsset() bool {
	if o != nil && !common.IsNil(o.ReceiveAsset) {
		return true
	}

	return false
}

// SetReceiveAsset gets a reference to the given string and assigns it to the ReceiveAsset field.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) SetReceiveAsset(v string) {
	o.ReceiveAsset = &v
}

// GetReceiveAmount returns the ReceiveAmount field value if set, zero value otherwise.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) GetReceiveAmount() string {
	if o == nil || common.IsNil(o.ReceiveAmount) {
		var ret string
		return ret
	}
	return *o.ReceiveAmount
}

// GetReceiveAmountOk returns a tuple with the ReceiveAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) GetReceiveAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReceiveAmount) {
		return nil, false
	}
	return o.ReceiveAmount, true
}

// HasReceiveAmount returns a boolean if a field has been set.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) HasReceiveAmount() bool {
	if o != nil && !common.IsNil(o.ReceiveAmount) {
		return true
	}

	return false
}

// SetReceiveAmount gets a reference to the given string and assigns it to the ReceiveAmount field.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) SetReceiveAmount(v string) {
	o.ReceiveAmount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetBfusdSubscriptionHistoryResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetBfusdSubscriptionHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBfusdSubscriptionHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.ReceiveAsset) {
		toSerialize["receiveAsset"] = o.ReceiveAsset
	}
	if !common.IsNil(o.ReceiveAmount) {
		toSerialize["receiveAmount"] = o.ReceiveAmount
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetBfusdSubscriptionHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetBfusdSubscriptionHistoryResponseRowsInner := _GetBfusdSubscriptionHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetBfusdSubscriptionHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetBfusdSubscriptionHistoryResponseRowsInner(varGetBfusdSubscriptionHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "receiveAsset")
		delete(additionalProperties, "receiveAmount")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBfusdSubscriptionHistoryResponseRowsInner struct {
	value *GetBfusdSubscriptionHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetBfusdSubscriptionHistoryResponseRowsInner) Get() *GetBfusdSubscriptionHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetBfusdSubscriptionHistoryResponseRowsInner) Set(val *GetBfusdSubscriptionHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBfusdSubscriptionHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBfusdSubscriptionHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBfusdSubscriptionHistoryResponseRowsInner(val *GetBfusdSubscriptionHistoryResponseRowsInner) *NullableGetBfusdSubscriptionHistoryResponseRowsInner {
	return &NullableGetBfusdSubscriptionHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetBfusdSubscriptionHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBfusdSubscriptionHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
