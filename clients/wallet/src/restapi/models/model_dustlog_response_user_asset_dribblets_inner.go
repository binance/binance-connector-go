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

// checks if the DustlogResponseUserAssetDribbletsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DustlogResponseUserAssetDribbletsInner{}

// DustlogResponseUserAssetDribbletsInner struct for DustlogResponseUserAssetDribbletsInner
type DustlogResponseUserAssetDribbletsInner struct {
	OperateTime              *int64                                                                `json:"operateTime,omitempty"`
	TotalTransferedAmount    *string                                                               `json:"totalTransferedAmount,omitempty"`
	TotalServiceChargeAmount *string                                                               `json:"totalServiceChargeAmount,omitempty"`
	TransId                  *int64                                                                `json:"transId,omitempty"`
	UserAssetDribbletDetails []DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner `json:"userAssetDribbletDetails,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _DustlogResponseUserAssetDribbletsInner DustlogResponseUserAssetDribbletsInner

// NewDustlogResponseUserAssetDribbletsInner instantiates a new DustlogResponseUserAssetDribbletsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDustlogResponseUserAssetDribbletsInner() *DustlogResponseUserAssetDribbletsInner {
	this := DustlogResponseUserAssetDribbletsInner{}
	return &this
}

// NewDustlogResponseUserAssetDribbletsInnerWithDefaults instantiates a new DustlogResponseUserAssetDribbletsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDustlogResponseUserAssetDribbletsInnerWithDefaults() *DustlogResponseUserAssetDribbletsInner {
	this := DustlogResponseUserAssetDribbletsInner{}
	return &this
}

// GetOperateTime returns the OperateTime field value if set, zero value otherwise.
func (o *DustlogResponseUserAssetDribbletsInner) GetOperateTime() int64 {
	if o == nil || common.IsNil(o.OperateTime) {
		var ret int64
		return ret
	}
	return *o.OperateTime
}

// GetOperateTimeOk returns a tuple with the OperateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustlogResponseUserAssetDribbletsInner) GetOperateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OperateTime) {
		return nil, false
	}
	return o.OperateTime, true
}

// HasOperateTime returns a boolean if a field has been set.
func (o *DustlogResponseUserAssetDribbletsInner) HasOperateTime() bool {
	if o != nil && !common.IsNil(o.OperateTime) {
		return true
	}

	return false
}

// SetOperateTime gets a reference to the given int64 and assigns it to the OperateTime field.
func (o *DustlogResponseUserAssetDribbletsInner) SetOperateTime(v int64) {
	o.OperateTime = &v
}

// GetTotalTransferedAmount returns the TotalTransferedAmount field value if set, zero value otherwise.
func (o *DustlogResponseUserAssetDribbletsInner) GetTotalTransferedAmount() string {
	if o == nil || common.IsNil(o.TotalTransferedAmount) {
		var ret string
		return ret
	}
	return *o.TotalTransferedAmount
}

// GetTotalTransferedAmountOk returns a tuple with the TotalTransferedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustlogResponseUserAssetDribbletsInner) GetTotalTransferedAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalTransferedAmount) {
		return nil, false
	}
	return o.TotalTransferedAmount, true
}

// HasTotalTransferedAmount returns a boolean if a field has been set.
func (o *DustlogResponseUserAssetDribbletsInner) HasTotalTransferedAmount() bool {
	if o != nil && !common.IsNil(o.TotalTransferedAmount) {
		return true
	}

	return false
}

// SetTotalTransferedAmount gets a reference to the given string and assigns it to the TotalTransferedAmount field.
func (o *DustlogResponseUserAssetDribbletsInner) SetTotalTransferedAmount(v string) {
	o.TotalTransferedAmount = &v
}

// GetTotalServiceChargeAmount returns the TotalServiceChargeAmount field value if set, zero value otherwise.
func (o *DustlogResponseUserAssetDribbletsInner) GetTotalServiceChargeAmount() string {
	if o == nil || common.IsNil(o.TotalServiceChargeAmount) {
		var ret string
		return ret
	}
	return *o.TotalServiceChargeAmount
}

// GetTotalServiceChargeAmountOk returns a tuple with the TotalServiceChargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustlogResponseUserAssetDribbletsInner) GetTotalServiceChargeAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalServiceChargeAmount) {
		return nil, false
	}
	return o.TotalServiceChargeAmount, true
}

// HasTotalServiceChargeAmount returns a boolean if a field has been set.
func (o *DustlogResponseUserAssetDribbletsInner) HasTotalServiceChargeAmount() bool {
	if o != nil && !common.IsNil(o.TotalServiceChargeAmount) {
		return true
	}

	return false
}

// SetTotalServiceChargeAmount gets a reference to the given string and assigns it to the TotalServiceChargeAmount field.
func (o *DustlogResponseUserAssetDribbletsInner) SetTotalServiceChargeAmount(v string) {
	o.TotalServiceChargeAmount = &v
}

// GetTransId returns the TransId field value if set, zero value otherwise.
func (o *DustlogResponseUserAssetDribbletsInner) GetTransId() int64 {
	if o == nil || common.IsNil(o.TransId) {
		var ret int64
		return ret
	}
	return *o.TransId
}

// GetTransIdOk returns a tuple with the TransId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustlogResponseUserAssetDribbletsInner) GetTransIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransId) {
		return nil, false
	}
	return o.TransId, true
}

// HasTransId returns a boolean if a field has been set.
func (o *DustlogResponseUserAssetDribbletsInner) HasTransId() bool {
	if o != nil && !common.IsNil(o.TransId) {
		return true
	}

	return false
}

// SetTransId gets a reference to the given int64 and assigns it to the TransId field.
func (o *DustlogResponseUserAssetDribbletsInner) SetTransId(v int64) {
	o.TransId = &v
}

// GetUserAssetDribbletDetails returns the UserAssetDribbletDetails field value if set, zero value otherwise.
func (o *DustlogResponseUserAssetDribbletsInner) GetUserAssetDribbletDetails() []DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner {
	if o == nil || common.IsNil(o.UserAssetDribbletDetails) {
		var ret []DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner
		return ret
	}
	return o.UserAssetDribbletDetails
}

// GetUserAssetDribbletDetailsOk returns a tuple with the UserAssetDribbletDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustlogResponseUserAssetDribbletsInner) GetUserAssetDribbletDetailsOk() ([]DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner, bool) {
	if o == nil || common.IsNil(o.UserAssetDribbletDetails) {
		return nil, false
	}
	return o.UserAssetDribbletDetails, true
}

// HasUserAssetDribbletDetails returns a boolean if a field has been set.
func (o *DustlogResponseUserAssetDribbletsInner) HasUserAssetDribbletDetails() bool {
	if o != nil && !common.IsNil(o.UserAssetDribbletDetails) {
		return true
	}

	return false
}

// SetUserAssetDribbletDetails gets a reference to the given []DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner and assigns it to the UserAssetDribbletDetails field.
func (o *DustlogResponseUserAssetDribbletsInner) SetUserAssetDribbletDetails(v []DustlogResponseUserAssetDribbletsInnerUserAssetDribbletDetailsInner) {
	o.UserAssetDribbletDetails = v
}

func (o DustlogResponseUserAssetDribbletsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DustlogResponseUserAssetDribbletsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OperateTime) {
		toSerialize["operateTime"] = o.OperateTime
	}
	if !common.IsNil(o.TotalTransferedAmount) {
		toSerialize["totalTransferedAmount"] = o.TotalTransferedAmount
	}
	if !common.IsNil(o.TotalServiceChargeAmount) {
		toSerialize["totalServiceChargeAmount"] = o.TotalServiceChargeAmount
	}
	if !common.IsNil(o.TransId) {
		toSerialize["transId"] = o.TransId
	}
	if !common.IsNil(o.UserAssetDribbletDetails) {
		toSerialize["userAssetDribbletDetails"] = o.UserAssetDribbletDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DustlogResponseUserAssetDribbletsInner) UnmarshalJSON(data []byte) (err error) {
	varDustlogResponseUserAssetDribbletsInner := _DustlogResponseUserAssetDribbletsInner{}

	err = json.Unmarshal(data, &varDustlogResponseUserAssetDribbletsInner)

	if err != nil {
		return err
	}

	*o = DustlogResponseUserAssetDribbletsInner(varDustlogResponseUserAssetDribbletsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "operateTime")
		delete(additionalProperties, "totalTransferedAmount")
		delete(additionalProperties, "totalServiceChargeAmount")
		delete(additionalProperties, "transId")
		delete(additionalProperties, "userAssetDribbletDetails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDustlogResponseUserAssetDribbletsInner struct {
	value *DustlogResponseUserAssetDribbletsInner
	isSet bool
}

func (v NullableDustlogResponseUserAssetDribbletsInner) Get() *DustlogResponseUserAssetDribbletsInner {
	return v.value
}

func (v *NullableDustlogResponseUserAssetDribbletsInner) Set(val *DustlogResponseUserAssetDribbletsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDustlogResponseUserAssetDribbletsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDustlogResponseUserAssetDribbletsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDustlogResponseUserAssetDribbletsInner(val *DustlogResponseUserAssetDribbletsInner) *NullableDustlogResponseUserAssetDribbletsInner {
	return &NullableDustlogResponseUserAssetDribbletsInner{value: val, isSet: true}
}

func (v NullableDustlogResponseUserAssetDribbletsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDustlogResponseUserAssetDribbletsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
