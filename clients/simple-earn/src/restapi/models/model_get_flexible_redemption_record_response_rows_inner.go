/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFlexibleRedemptionRecordResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleRedemptionRecordResponseRowsInner{}

// GetFlexibleRedemptionRecordResponseRowsInner struct for GetFlexibleRedemptionRecordResponseRowsInner
type GetFlexibleRedemptionRecordResponseRowsInner struct {
	Amount               *string `json:"amount,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	ProjectId            *string `json:"projectId,omitempty"`
	RedeemId             *int64  `json:"redeemId,omitempty"`
	DestAccount          *string `json:"destAccount,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleRedemptionRecordResponseRowsInner GetFlexibleRedemptionRecordResponseRowsInner

// NewGetFlexibleRedemptionRecordResponseRowsInner instantiates a new GetFlexibleRedemptionRecordResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleRedemptionRecordResponseRowsInner() *GetFlexibleRedemptionRecordResponseRowsInner {
	this := GetFlexibleRedemptionRecordResponseRowsInner{}
	return &this
}

// NewGetFlexibleRedemptionRecordResponseRowsInnerWithDefaults instantiates a new GetFlexibleRedemptionRecordResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleRedemptionRecordResponseRowsInnerWithDefaults() *GetFlexibleRedemptionRecordResponseRowsInner {
	this := GetFlexibleRedemptionRecordResponseRowsInner{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) GetProjectId() string {
	if o == nil || common.IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) GetProjectIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) HasProjectId() bool {
	if o != nil && !common.IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetRedeemId returns the RedeemId field value if set, zero value otherwise.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) GetRedeemId() int64 {
	if o == nil || common.IsNil(o.RedeemId) {
		var ret int64
		return ret
	}
	return *o.RedeemId
}

// GetRedeemIdOk returns a tuple with the RedeemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) GetRedeemIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RedeemId) {
		return nil, false
	}
	return o.RedeemId, true
}

// HasRedeemId returns a boolean if a field has been set.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) HasRedeemId() bool {
	if o != nil && !common.IsNil(o.RedeemId) {
		return true
	}

	return false
}

// SetRedeemId gets a reference to the given int64 and assigns it to the RedeemId field.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) SetRedeemId(v int64) {
	o.RedeemId = &v
}

// GetDestAccount returns the DestAccount field value if set, zero value otherwise.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) GetDestAccount() string {
	if o == nil || common.IsNil(o.DestAccount) {
		var ret string
		return ret
	}
	return *o.DestAccount
}

// GetDestAccountOk returns a tuple with the DestAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) GetDestAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.DestAccount) {
		return nil, false
	}
	return o.DestAccount, true
}

// HasDestAccount returns a boolean if a field has been set.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) HasDestAccount() bool {
	if o != nil && !common.IsNil(o.DestAccount) {
		return true
	}

	return false
}

// SetDestAccount gets a reference to the given string and assigns it to the DestAccount field.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) SetDestAccount(v string) {
	o.DestAccount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetFlexibleRedemptionRecordResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetFlexibleRedemptionRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleRedemptionRecordResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !common.IsNil(o.RedeemId) {
		toSerialize["redeemId"] = o.RedeemId
	}
	if !common.IsNil(o.DestAccount) {
		toSerialize["destAccount"] = o.DestAccount
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFlexibleRedemptionRecordResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleRedemptionRecordResponseRowsInner := _GetFlexibleRedemptionRecordResponseRowsInner{}

	err = json.Unmarshal(data, &varGetFlexibleRedemptionRecordResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetFlexibleRedemptionRecordResponseRowsInner(varGetFlexibleRedemptionRecordResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "time")
		delete(additionalProperties, "projectId")
		delete(additionalProperties, "redeemId")
		delete(additionalProperties, "destAccount")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleRedemptionRecordResponseRowsInner struct {
	value *GetFlexibleRedemptionRecordResponseRowsInner
	isSet bool
}

func (v NullableGetFlexibleRedemptionRecordResponseRowsInner) Get() *GetFlexibleRedemptionRecordResponseRowsInner {
	return v.value
}

func (v *NullableGetFlexibleRedemptionRecordResponseRowsInner) Set(val *GetFlexibleRedemptionRecordResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleRedemptionRecordResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleRedemptionRecordResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleRedemptionRecordResponseRowsInner(val *GetFlexibleRedemptionRecordResponseRowsInner) *NullableGetFlexibleRedemptionRecordResponseRowsInner {
	return &NullableGetFlexibleRedemptionRecordResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetFlexibleRedemptionRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleRedemptionRecordResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
