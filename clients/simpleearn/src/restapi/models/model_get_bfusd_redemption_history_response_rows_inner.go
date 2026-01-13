/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetBfusdRedemptionHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBfusdRedemptionHistoryResponseRowsInner{}

// GetBfusdRedemptionHistoryResponseRowsInner struct for GetBfusdRedemptionHistoryResponseRowsInner
type GetBfusdRedemptionHistoryResponseRowsInner struct {
	Time                 *int64  `json:"time,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	ReceiveAsset         *string `json:"receiveAsset,omitempty"`
	ReceiveAmount        *string `json:"receiveAmount,omitempty"`
	Fee                  *string `json:"fee,omitempty"`
	ArrivalTime          *int64  `json:"arrivalTime,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBfusdRedemptionHistoryResponseRowsInner GetBfusdRedemptionHistoryResponseRowsInner

// NewGetBfusdRedemptionHistoryResponseRowsInner instantiates a new GetBfusdRedemptionHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBfusdRedemptionHistoryResponseRowsInner() *GetBfusdRedemptionHistoryResponseRowsInner {
	this := GetBfusdRedemptionHistoryResponseRowsInner{}
	return &this
}

// NewGetBfusdRedemptionHistoryResponseRowsInnerWithDefaults instantiates a new GetBfusdRedemptionHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBfusdRedemptionHistoryResponseRowsInnerWithDefaults() *GetBfusdRedemptionHistoryResponseRowsInner {
	this := GetBfusdRedemptionHistoryResponseRowsInner{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetReceiveAsset returns the ReceiveAsset field value if set, zero value otherwise.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) GetReceiveAsset() string {
	if o == nil || common.IsNil(o.ReceiveAsset) {
		var ret string
		return ret
	}
	return *o.ReceiveAsset
}

// GetReceiveAssetOk returns a tuple with the ReceiveAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) GetReceiveAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReceiveAsset) {
		return nil, false
	}
	return o.ReceiveAsset, true
}

// HasReceiveAsset returns a boolean if a field has been set.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) HasReceiveAsset() bool {
	if o != nil && !common.IsNil(o.ReceiveAsset) {
		return true
	}

	return false
}

// SetReceiveAsset gets a reference to the given string and assigns it to the ReceiveAsset field.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) SetReceiveAsset(v string) {
	o.ReceiveAsset = &v
}

// GetReceiveAmount returns the ReceiveAmount field value if set, zero value otherwise.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) GetReceiveAmount() string {
	if o == nil || common.IsNil(o.ReceiveAmount) {
		var ret string
		return ret
	}
	return *o.ReceiveAmount
}

// GetReceiveAmountOk returns a tuple with the ReceiveAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) GetReceiveAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReceiveAmount) {
		return nil, false
	}
	return o.ReceiveAmount, true
}

// HasReceiveAmount returns a boolean if a field has been set.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) HasReceiveAmount() bool {
	if o != nil && !common.IsNil(o.ReceiveAmount) {
		return true
	}

	return false
}

// SetReceiveAmount gets a reference to the given string and assigns it to the ReceiveAmount field.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) SetReceiveAmount(v string) {
	o.ReceiveAmount = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) GetFee() string {
	if o == nil || common.IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) GetFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) SetFee(v string) {
	o.Fee = &v
}

// GetArrivalTime returns the ArrivalTime field value if set, zero value otherwise.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) GetArrivalTime() int64 {
	if o == nil || common.IsNil(o.ArrivalTime) {
		var ret int64
		return ret
	}
	return *o.ArrivalTime
}

// GetArrivalTimeOk returns a tuple with the ArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) GetArrivalTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ArrivalTime) {
		return nil, false
	}
	return o.ArrivalTime, true
}

// HasArrivalTime returns a boolean if a field has been set.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) HasArrivalTime() bool {
	if o != nil && !common.IsNil(o.ArrivalTime) {
		return true
	}

	return false
}

// SetArrivalTime gets a reference to the given int64 and assigns it to the ArrivalTime field.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) SetArrivalTime(v int64) {
	o.ArrivalTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetBfusdRedemptionHistoryResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetBfusdRedemptionHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBfusdRedemptionHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !common.IsNil(o.ArrivalTime) {
		toSerialize["arrivalTime"] = o.ArrivalTime
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetBfusdRedemptionHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetBfusdRedemptionHistoryResponseRowsInner := _GetBfusdRedemptionHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetBfusdRedemptionHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetBfusdRedemptionHistoryResponseRowsInner(varGetBfusdRedemptionHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "receiveAsset")
		delete(additionalProperties, "receiveAmount")
		delete(additionalProperties, "fee")
		delete(additionalProperties, "arrivalTime")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBfusdRedemptionHistoryResponseRowsInner struct {
	value *GetBfusdRedemptionHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetBfusdRedemptionHistoryResponseRowsInner) Get() *GetBfusdRedemptionHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetBfusdRedemptionHistoryResponseRowsInner) Set(val *GetBfusdRedemptionHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBfusdRedemptionHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBfusdRedemptionHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBfusdRedemptionHistoryResponseRowsInner(val *GetBfusdRedemptionHistoryResponseRowsInner) *NullableGetBfusdRedemptionHistoryResponseRowsInner {
	return &NullableGetBfusdRedemptionHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetBfusdRedemptionHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBfusdRedemptionHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
