/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner{}

// DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner struct for DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner
type DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner struct {
	TransId              *int64  `json:"transId,omitempty"`
	ServiceChargeAmount  *string `json:"serviceChargeAmount,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	OperateTime          *int64  `json:"operateTime,omitempty"`
	TransferedAmount     *string `json:"transferedAmount,omitempty"`
	FromAsset            *string `json:"fromAsset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner

// NewDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner instantiates a new DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner() *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner {
	this := DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner{}
	return &this
}

// NewDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInnerWithDefaults instantiates a new DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInnerWithDefaults() *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner {
	this := DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner{}
	return &this
}

// GetTransId returns the TransId field value if set, zero value otherwise.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) GetTransId() int64 {
	if o == nil || common.IsNil(o.TransId) {
		var ret int64
		return ret
	}
	return *o.TransId
}

// GetTransIdOk returns a tuple with the TransId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) GetTransIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransId) {
		return nil, false
	}
	return o.TransId, true
}

// HasTransId returns a boolean if a field has been set.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) HasTransId() bool {
	if o != nil && !common.IsNil(o.TransId) {
		return true
	}

	return false
}

// SetTransId gets a reference to the given int64 and assigns it to the TransId field.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) SetTransId(v int64) {
	o.TransId = &v
}

// GetServiceChargeAmount returns the ServiceChargeAmount field value if set, zero value otherwise.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) GetServiceChargeAmount() string {
	if o == nil || common.IsNil(o.ServiceChargeAmount) {
		var ret string
		return ret
	}
	return *o.ServiceChargeAmount
}

// GetServiceChargeAmountOk returns a tuple with the ServiceChargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) GetServiceChargeAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ServiceChargeAmount) {
		return nil, false
	}
	return o.ServiceChargeAmount, true
}

// HasServiceChargeAmount returns a boolean if a field has been set.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) HasServiceChargeAmount() bool {
	if o != nil && !common.IsNil(o.ServiceChargeAmount) {
		return true
	}

	return false
}

// SetServiceChargeAmount gets a reference to the given string and assigns it to the ServiceChargeAmount field.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) SetServiceChargeAmount(v string) {
	o.ServiceChargeAmount = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetOperateTime returns the OperateTime field value if set, zero value otherwise.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) GetOperateTime() int64 {
	if o == nil || common.IsNil(o.OperateTime) {
		var ret int64
		return ret
	}
	return *o.OperateTime
}

// GetOperateTimeOk returns a tuple with the OperateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) GetOperateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OperateTime) {
		return nil, false
	}
	return o.OperateTime, true
}

// HasOperateTime returns a boolean if a field has been set.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) HasOperateTime() bool {
	if o != nil && !common.IsNil(o.OperateTime) {
		return true
	}

	return false
}

// SetOperateTime gets a reference to the given int64 and assigns it to the OperateTime field.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) SetOperateTime(v int64) {
	o.OperateTime = &v
}

// GetTransferedAmount returns the TransferedAmount field value if set, zero value otherwise.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) GetTransferedAmount() string {
	if o == nil || common.IsNil(o.TransferedAmount) {
		var ret string
		return ret
	}
	return *o.TransferedAmount
}

// GetTransferedAmountOk returns a tuple with the TransferedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) GetTransferedAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferedAmount) {
		return nil, false
	}
	return o.TransferedAmount, true
}

// HasTransferedAmount returns a boolean if a field has been set.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) HasTransferedAmount() bool {
	if o != nil && !common.IsNil(o.TransferedAmount) {
		return true
	}

	return false
}

// SetTransferedAmount gets a reference to the given string and assigns it to the TransferedAmount field.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) SetTransferedAmount(v string) {
	o.TransferedAmount = &v
}

// GetFromAsset returns the FromAsset field value if set, zero value otherwise.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) GetFromAsset() string {
	if o == nil || common.IsNil(o.FromAsset) {
		var ret string
		return ret
	}
	return *o.FromAsset
}

// GetFromAssetOk returns a tuple with the FromAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) GetFromAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAsset) {
		return nil, false
	}
	return o.FromAsset, true
}

// HasFromAsset returns a boolean if a field has been set.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) HasFromAsset() bool {
	if o != nil && !common.IsNil(o.FromAsset) {
		return true
	}

	return false
}

// SetFromAsset gets a reference to the given string and assigns it to the FromAsset field.
func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) SetFromAsset(v string) {
	o.FromAsset = &v
}

func (o DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TransId) {
		toSerialize["transId"] = o.TransId
	}
	if !common.IsNil(o.ServiceChargeAmount) {
		toSerialize["serviceChargeAmount"] = o.ServiceChargeAmount
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.OperateTime) {
		toSerialize["operateTime"] = o.OperateTime
	}
	if !common.IsNil(o.TransferedAmount) {
		toSerialize["transferedAmount"] = o.TransferedAmount
	}
	if !common.IsNil(o.FromAsset) {
		toSerialize["fromAsset"] = o.FromAsset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) UnmarshalJSON(data []byte) (err error) {
	varDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner := _DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner{}

	err = json.Unmarshal(data, &varDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner)

	if err != nil {
		return err
	}

	*o = DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner(varDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "transId")
		delete(additionalProperties, "serviceChargeAmount")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "operateTime")
		delete(additionalProperties, "transferedAmount")
		delete(additionalProperties, "fromAsset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner struct {
	value *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner
	isSet bool
}

func (v NullableDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) Get() *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner {
	return v.value
}

func (v *NullableDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) Set(val *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner(val *DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) *NullableDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner {
	return &NullableDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner{value: val, isSet: true}
}

func (v NullableDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
