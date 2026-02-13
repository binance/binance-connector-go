/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the SubAccountTransferHistoryResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubAccountTransferHistoryResponseInner{}

// SubAccountTransferHistoryResponseInner struct for SubAccountTransferHistoryResponseInner
type SubAccountTransferHistoryResponseInner struct {
	CounterParty         *string `json:"counterParty,omitempty"`
	Email                *string `json:"email,omitempty"`
	Type                 *int64  `json:"type,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	FromAccountType      *string `json:"fromAccountType,omitempty"`
	ToAccountType        *string `json:"toAccountType,omitempty"`
	Status               *string `json:"status,omitempty"`
	TranId               *int64  `json:"tranId,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubAccountTransferHistoryResponseInner SubAccountTransferHistoryResponseInner

// NewSubAccountTransferHistoryResponseInner instantiates a new SubAccountTransferHistoryResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubAccountTransferHistoryResponseInner() *SubAccountTransferHistoryResponseInner {
	this := SubAccountTransferHistoryResponseInner{}
	return &this
}

// NewSubAccountTransferHistoryResponseInnerWithDefaults instantiates a new SubAccountTransferHistoryResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubAccountTransferHistoryResponseInnerWithDefaults() *SubAccountTransferHistoryResponseInner {
	this := SubAccountTransferHistoryResponseInner{}
	return &this
}

// GetCounterParty returns the CounterParty field value if set, zero value otherwise.
func (o *SubAccountTransferHistoryResponseInner) GetCounterParty() string {
	if o == nil || common.IsNil(o.CounterParty) {
		var ret string
		return ret
	}
	return *o.CounterParty
}

// GetCounterPartyOk returns a tuple with the CounterParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountTransferHistoryResponseInner) GetCounterPartyOk() (*string, bool) {
	if o == nil || common.IsNil(o.CounterParty) {
		return nil, false
	}
	return o.CounterParty, true
}

// HasCounterParty returns a boolean if a field has been set.
func (o *SubAccountTransferHistoryResponseInner) HasCounterParty() bool {
	if o != nil && !common.IsNil(o.CounterParty) {
		return true
	}

	return false
}

// SetCounterParty gets a reference to the given string and assigns it to the CounterParty field.
func (o *SubAccountTransferHistoryResponseInner) SetCounterParty(v string) {
	o.CounterParty = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SubAccountTransferHistoryResponseInner) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountTransferHistoryResponseInner) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SubAccountTransferHistoryResponseInner) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SubAccountTransferHistoryResponseInner) SetEmail(v string) {
	o.Email = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SubAccountTransferHistoryResponseInner) GetType() int64 {
	if o == nil || common.IsNil(o.Type) {
		var ret int64
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountTransferHistoryResponseInner) GetTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SubAccountTransferHistoryResponseInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int64 and assigns it to the Type field.
func (o *SubAccountTransferHistoryResponseInner) SetType(v int64) {
	o.Type = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *SubAccountTransferHistoryResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountTransferHistoryResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *SubAccountTransferHistoryResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *SubAccountTransferHistoryResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *SubAccountTransferHistoryResponseInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountTransferHistoryResponseInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *SubAccountTransferHistoryResponseInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *SubAccountTransferHistoryResponseInner) SetQty(v string) {
	o.Qty = &v
}

// GetFromAccountType returns the FromAccountType field value if set, zero value otherwise.
func (o *SubAccountTransferHistoryResponseInner) GetFromAccountType() string {
	if o == nil || common.IsNil(o.FromAccountType) {
		var ret string
		return ret
	}
	return *o.FromAccountType
}

// GetFromAccountTypeOk returns a tuple with the FromAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountTransferHistoryResponseInner) GetFromAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAccountType) {
		return nil, false
	}
	return o.FromAccountType, true
}

// HasFromAccountType returns a boolean if a field has been set.
func (o *SubAccountTransferHistoryResponseInner) HasFromAccountType() bool {
	if o != nil && !common.IsNil(o.FromAccountType) {
		return true
	}

	return false
}

// SetFromAccountType gets a reference to the given string and assigns it to the FromAccountType field.
func (o *SubAccountTransferHistoryResponseInner) SetFromAccountType(v string) {
	o.FromAccountType = &v
}

// GetToAccountType returns the ToAccountType field value if set, zero value otherwise.
func (o *SubAccountTransferHistoryResponseInner) GetToAccountType() string {
	if o == nil || common.IsNil(o.ToAccountType) {
		var ret string
		return ret
	}
	return *o.ToAccountType
}

// GetToAccountTypeOk returns a tuple with the ToAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountTransferHistoryResponseInner) GetToAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAccountType) {
		return nil, false
	}
	return o.ToAccountType, true
}

// HasToAccountType returns a boolean if a field has been set.
func (o *SubAccountTransferHistoryResponseInner) HasToAccountType() bool {
	if o != nil && !common.IsNil(o.ToAccountType) {
		return true
	}

	return false
}

// SetToAccountType gets a reference to the given string and assigns it to the ToAccountType field.
func (o *SubAccountTransferHistoryResponseInner) SetToAccountType(v string) {
	o.ToAccountType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SubAccountTransferHistoryResponseInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountTransferHistoryResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SubAccountTransferHistoryResponseInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SubAccountTransferHistoryResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *SubAccountTransferHistoryResponseInner) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountTransferHistoryResponseInner) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *SubAccountTransferHistoryResponseInner) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *SubAccountTransferHistoryResponseInner) SetTranId(v int64) {
	o.TranId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SubAccountTransferHistoryResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountTransferHistoryResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SubAccountTransferHistoryResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *SubAccountTransferHistoryResponseInner) SetTime(v int64) {
	o.Time = &v
}

func (o SubAccountTransferHistoryResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubAccountTransferHistoryResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CounterParty) {
		toSerialize["counterParty"] = o.CounterParty
	}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !common.IsNil(o.FromAccountType) {
		toSerialize["fromAccountType"] = o.FromAccountType
	}
	if !common.IsNil(o.ToAccountType) {
		toSerialize["toAccountType"] = o.ToAccountType
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubAccountTransferHistoryResponseInner) UnmarshalJSON(data []byte) (err error) {
	varSubAccountTransferHistoryResponseInner := _SubAccountTransferHistoryResponseInner{}

	err = json.Unmarshal(data, &varSubAccountTransferHistoryResponseInner)

	if err != nil {
		return err
	}

	*o = SubAccountTransferHistoryResponseInner(varSubAccountTransferHistoryResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "counterParty")
		delete(additionalProperties, "email")
		delete(additionalProperties, "type")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "fromAccountType")
		delete(additionalProperties, "toAccountType")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tranId")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubAccountTransferHistoryResponseInner struct {
	value *SubAccountTransferHistoryResponseInner
	isSet bool
}

func (v NullableSubAccountTransferHistoryResponseInner) Get() *SubAccountTransferHistoryResponseInner {
	return v.value
}

func (v *NullableSubAccountTransferHistoryResponseInner) Set(val *SubAccountTransferHistoryResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSubAccountTransferHistoryResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSubAccountTransferHistoryResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubAccountTransferHistoryResponseInner(val *SubAccountTransferHistoryResponseInner) *NullableSubAccountTransferHistoryResponseInner {
	return &NullableSubAccountTransferHistoryResponseInner{value: val, isSet: true}
}

func (v NullableSubAccountTransferHistoryResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubAccountTransferHistoryResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
