/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DustTransferResponseTransferResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DustTransferResponseTransferResultInner{}

// DustTransferResponseTransferResultInner struct for DustTransferResponseTransferResultInner
type DustTransferResponseTransferResultInner struct {
	Amount               *string `json:"amount,omitempty"`
	FromAsset            *string `json:"fromAsset,omitempty"`
	OperateTime          *int64  `json:"operateTime,omitempty"`
	ServiceChargeAmount  *string `json:"serviceChargeAmount,omitempty"`
	TranId               *int64  `json:"tranId,omitempty"`
	TransferedAmount     *string `json:"transferedAmount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DustTransferResponseTransferResultInner DustTransferResponseTransferResultInner

// NewDustTransferResponseTransferResultInner instantiates a new DustTransferResponseTransferResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDustTransferResponseTransferResultInner() *DustTransferResponseTransferResultInner {
	this := DustTransferResponseTransferResultInner{}
	return &this
}

// NewDustTransferResponseTransferResultInnerWithDefaults instantiates a new DustTransferResponseTransferResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDustTransferResponseTransferResultInnerWithDefaults() *DustTransferResponseTransferResultInner {
	this := DustTransferResponseTransferResultInner{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *DustTransferResponseTransferResultInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustTransferResponseTransferResultInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *DustTransferResponseTransferResultInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *DustTransferResponseTransferResultInner) SetAmount(v string) {
	o.Amount = &v
}

// GetFromAsset returns the FromAsset field value if set, zero value otherwise.
func (o *DustTransferResponseTransferResultInner) GetFromAsset() string {
	if o == nil || common.IsNil(o.FromAsset) {
		var ret string
		return ret
	}
	return *o.FromAsset
}

// GetFromAssetOk returns a tuple with the FromAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustTransferResponseTransferResultInner) GetFromAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAsset) {
		return nil, false
	}
	return o.FromAsset, true
}

// HasFromAsset returns a boolean if a field has been set.
func (o *DustTransferResponseTransferResultInner) HasFromAsset() bool {
	if o != nil && !common.IsNil(o.FromAsset) {
		return true
	}

	return false
}

// SetFromAsset gets a reference to the given string and assigns it to the FromAsset field.
func (o *DustTransferResponseTransferResultInner) SetFromAsset(v string) {
	o.FromAsset = &v
}

// GetOperateTime returns the OperateTime field value if set, zero value otherwise.
func (o *DustTransferResponseTransferResultInner) GetOperateTime() int64 {
	if o == nil || common.IsNil(o.OperateTime) {
		var ret int64
		return ret
	}
	return *o.OperateTime
}

// GetOperateTimeOk returns a tuple with the OperateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustTransferResponseTransferResultInner) GetOperateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OperateTime) {
		return nil, false
	}
	return o.OperateTime, true
}

// HasOperateTime returns a boolean if a field has been set.
func (o *DustTransferResponseTransferResultInner) HasOperateTime() bool {
	if o != nil && !common.IsNil(o.OperateTime) {
		return true
	}

	return false
}

// SetOperateTime gets a reference to the given int64 and assigns it to the OperateTime field.
func (o *DustTransferResponseTransferResultInner) SetOperateTime(v int64) {
	o.OperateTime = &v
}

// GetServiceChargeAmount returns the ServiceChargeAmount field value if set, zero value otherwise.
func (o *DustTransferResponseTransferResultInner) GetServiceChargeAmount() string {
	if o == nil || common.IsNil(o.ServiceChargeAmount) {
		var ret string
		return ret
	}
	return *o.ServiceChargeAmount
}

// GetServiceChargeAmountOk returns a tuple with the ServiceChargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustTransferResponseTransferResultInner) GetServiceChargeAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ServiceChargeAmount) {
		return nil, false
	}
	return o.ServiceChargeAmount, true
}

// HasServiceChargeAmount returns a boolean if a field has been set.
func (o *DustTransferResponseTransferResultInner) HasServiceChargeAmount() bool {
	if o != nil && !common.IsNil(o.ServiceChargeAmount) {
		return true
	}

	return false
}

// SetServiceChargeAmount gets a reference to the given string and assigns it to the ServiceChargeAmount field.
func (o *DustTransferResponseTransferResultInner) SetServiceChargeAmount(v string) {
	o.ServiceChargeAmount = &v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *DustTransferResponseTransferResultInner) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustTransferResponseTransferResultInner) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *DustTransferResponseTransferResultInner) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *DustTransferResponseTransferResultInner) SetTranId(v int64) {
	o.TranId = &v
}

// GetTransferedAmount returns the TransferedAmount field value if set, zero value otherwise.
func (o *DustTransferResponseTransferResultInner) GetTransferedAmount() string {
	if o == nil || common.IsNil(o.TransferedAmount) {
		var ret string
		return ret
	}
	return *o.TransferedAmount
}

// GetTransferedAmountOk returns a tuple with the TransferedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustTransferResponseTransferResultInner) GetTransferedAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferedAmount) {
		return nil, false
	}
	return o.TransferedAmount, true
}

// HasTransferedAmount returns a boolean if a field has been set.
func (o *DustTransferResponseTransferResultInner) HasTransferedAmount() bool {
	if o != nil && !common.IsNil(o.TransferedAmount) {
		return true
	}

	return false
}

// SetTransferedAmount gets a reference to the given string and assigns it to the TransferedAmount field.
func (o *DustTransferResponseTransferResultInner) SetTransferedAmount(v string) {
	o.TransferedAmount = &v
}

func (o DustTransferResponseTransferResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DustTransferResponseTransferResultInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.FromAsset) {
		toSerialize["fromAsset"] = o.FromAsset
	}
	if !common.IsNil(o.OperateTime) {
		toSerialize["operateTime"] = o.OperateTime
	}
	if !common.IsNil(o.ServiceChargeAmount) {
		toSerialize["serviceChargeAmount"] = o.ServiceChargeAmount
	}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	if !common.IsNil(o.TransferedAmount) {
		toSerialize["transferedAmount"] = o.TransferedAmount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DustTransferResponseTransferResultInner) UnmarshalJSON(data []byte) (err error) {
	varDustTransferResponseTransferResultInner := _DustTransferResponseTransferResultInner{}

	err = json.Unmarshal(data, &varDustTransferResponseTransferResultInner)

	if err != nil {
		return err
	}

	*o = DustTransferResponseTransferResultInner(varDustTransferResponseTransferResultInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "fromAsset")
		delete(additionalProperties, "operateTime")
		delete(additionalProperties, "serviceChargeAmount")
		delete(additionalProperties, "tranId")
		delete(additionalProperties, "transferedAmount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDustTransferResponseTransferResultInner struct {
	value *DustTransferResponseTransferResultInner
	isSet bool
}

func (v NullableDustTransferResponseTransferResultInner) Get() *DustTransferResponseTransferResultInner {
	return v.value
}

func (v *NullableDustTransferResponseTransferResultInner) Set(val *DustTransferResponseTransferResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDustTransferResponseTransferResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDustTransferResponseTransferResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDustTransferResponseTransferResultInner(val *DustTransferResponseTransferResultInner) *NullableDustTransferResponseTransferResultInner {
	return &NullableDustTransferResponseTransferResultInner{value: val, isSet: true}
}

func (v NullableDustTransferResponseTransferResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDustTransferResponseTransferResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
