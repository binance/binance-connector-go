/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DustConvertResponseTransferResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DustConvertResponseTransferResultInner{}

// DustConvertResponseTransferResultInner struct for DustConvertResponseTransferResultInner
type DustConvertResponseTransferResultInner struct {
	TranId               *int64  `json:"tranId,omitempty"`
	FromAsset            *string `json:"fromAsset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	TransferedAmount     *string `json:"transferedAmount,omitempty"`
	ServiceChargeAmount  *string `json:"serviceChargeAmount,omitempty"`
	OperateTime          *int64  `json:"operateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DustConvertResponseTransferResultInner DustConvertResponseTransferResultInner

// NewDustConvertResponseTransferResultInner instantiates a new DustConvertResponseTransferResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDustConvertResponseTransferResultInner() *DustConvertResponseTransferResultInner {
	this := DustConvertResponseTransferResultInner{}
	return &this
}

// NewDustConvertResponseTransferResultInnerWithDefaults instantiates a new DustConvertResponseTransferResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDustConvertResponseTransferResultInnerWithDefaults() *DustConvertResponseTransferResultInner {
	this := DustConvertResponseTransferResultInner{}
	return &this
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *DustConvertResponseTransferResultInner) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertResponseTransferResultInner) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *DustConvertResponseTransferResultInner) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *DustConvertResponseTransferResultInner) SetTranId(v int64) {
	o.TranId = &v
}

// GetFromAsset returns the FromAsset field value if set, zero value otherwise.
func (o *DustConvertResponseTransferResultInner) GetFromAsset() string {
	if o == nil || common.IsNil(o.FromAsset) {
		var ret string
		return ret
	}
	return *o.FromAsset
}

// GetFromAssetOk returns a tuple with the FromAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertResponseTransferResultInner) GetFromAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAsset) {
		return nil, false
	}
	return o.FromAsset, true
}

// HasFromAsset returns a boolean if a field has been set.
func (o *DustConvertResponseTransferResultInner) HasFromAsset() bool {
	if o != nil && !common.IsNil(o.FromAsset) {
		return true
	}

	return false
}

// SetFromAsset gets a reference to the given string and assigns it to the FromAsset field.
func (o *DustConvertResponseTransferResultInner) SetFromAsset(v string) {
	o.FromAsset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *DustConvertResponseTransferResultInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertResponseTransferResultInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *DustConvertResponseTransferResultInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *DustConvertResponseTransferResultInner) SetAmount(v string) {
	o.Amount = &v
}

// GetTransferedAmount returns the TransferedAmount field value if set, zero value otherwise.
func (o *DustConvertResponseTransferResultInner) GetTransferedAmount() string {
	if o == nil || common.IsNil(o.TransferedAmount) {
		var ret string
		return ret
	}
	return *o.TransferedAmount
}

// GetTransferedAmountOk returns a tuple with the TransferedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertResponseTransferResultInner) GetTransferedAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferedAmount) {
		return nil, false
	}
	return o.TransferedAmount, true
}

// HasTransferedAmount returns a boolean if a field has been set.
func (o *DustConvertResponseTransferResultInner) HasTransferedAmount() bool {
	if o != nil && !common.IsNil(o.TransferedAmount) {
		return true
	}

	return false
}

// SetTransferedAmount gets a reference to the given string and assigns it to the TransferedAmount field.
func (o *DustConvertResponseTransferResultInner) SetTransferedAmount(v string) {
	o.TransferedAmount = &v
}

// GetServiceChargeAmount returns the ServiceChargeAmount field value if set, zero value otherwise.
func (o *DustConvertResponseTransferResultInner) GetServiceChargeAmount() string {
	if o == nil || common.IsNil(o.ServiceChargeAmount) {
		var ret string
		return ret
	}
	return *o.ServiceChargeAmount
}

// GetServiceChargeAmountOk returns a tuple with the ServiceChargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertResponseTransferResultInner) GetServiceChargeAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ServiceChargeAmount) {
		return nil, false
	}
	return o.ServiceChargeAmount, true
}

// HasServiceChargeAmount returns a boolean if a field has been set.
func (o *DustConvertResponseTransferResultInner) HasServiceChargeAmount() bool {
	if o != nil && !common.IsNil(o.ServiceChargeAmount) {
		return true
	}

	return false
}

// SetServiceChargeAmount gets a reference to the given string and assigns it to the ServiceChargeAmount field.
func (o *DustConvertResponseTransferResultInner) SetServiceChargeAmount(v string) {
	o.ServiceChargeAmount = &v
}

// GetOperateTime returns the OperateTime field value if set, zero value otherwise.
func (o *DustConvertResponseTransferResultInner) GetOperateTime() int64 {
	if o == nil || common.IsNil(o.OperateTime) {
		var ret int64
		return ret
	}
	return *o.OperateTime
}

// GetOperateTimeOk returns a tuple with the OperateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertResponseTransferResultInner) GetOperateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OperateTime) {
		return nil, false
	}
	return o.OperateTime, true
}

// HasOperateTime returns a boolean if a field has been set.
func (o *DustConvertResponseTransferResultInner) HasOperateTime() bool {
	if o != nil && !common.IsNil(o.OperateTime) {
		return true
	}

	return false
}

// SetOperateTime gets a reference to the given int64 and assigns it to the OperateTime field.
func (o *DustConvertResponseTransferResultInner) SetOperateTime(v int64) {
	o.OperateTime = &v
}

func (o DustConvertResponseTransferResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DustConvertResponseTransferResultInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	if !common.IsNil(o.FromAsset) {
		toSerialize["fromAsset"] = o.FromAsset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.TransferedAmount) {
		toSerialize["transferedAmount"] = o.TransferedAmount
	}
	if !common.IsNil(o.ServiceChargeAmount) {
		toSerialize["serviceChargeAmount"] = o.ServiceChargeAmount
	}
	if !common.IsNil(o.OperateTime) {
		toSerialize["operateTime"] = o.OperateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DustConvertResponseTransferResultInner) UnmarshalJSON(data []byte) (err error) {
	varDustConvertResponseTransferResultInner := _DustConvertResponseTransferResultInner{}

	err = json.Unmarshal(data, &varDustConvertResponseTransferResultInner)

	if err != nil {
		return err
	}

	*o = DustConvertResponseTransferResultInner(varDustConvertResponseTransferResultInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tranId")
		delete(additionalProperties, "fromAsset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "transferedAmount")
		delete(additionalProperties, "serviceChargeAmount")
		delete(additionalProperties, "operateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDustConvertResponseTransferResultInner struct {
	value *DustConvertResponseTransferResultInner
	isSet bool
}

func (v NullableDustConvertResponseTransferResultInner) Get() *DustConvertResponseTransferResultInner {
	return v.value
}

func (v *NullableDustConvertResponseTransferResultInner) Set(val *DustConvertResponseTransferResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDustConvertResponseTransferResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDustConvertResponseTransferResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDustConvertResponseTransferResultInner(val *DustConvertResponseTransferResultInner) *NullableDustConvertResponseTransferResultInner {
	return &NullableDustConvertResponseTransferResultInner{value: val, isSet: true}
}

func (v NullableDustConvertResponseTransferResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDustConvertResponseTransferResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
