/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetLockedSubscriptionRecordResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLockedSubscriptionRecordResponseRowsInner{}

// GetLockedSubscriptionRecordResponseRowsInner struct for GetLockedSubscriptionRecordResponseRowsInner
type GetLockedSubscriptionRecordResponseRowsInner struct {
	PositionId           *int64  `json:"positionId,omitempty"`
	PurchaseId           *string `json:"purchaseId,omitempty"`
	ProjectId            *string `json:"projectId,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	LockPeriod           *string `json:"lockPeriod,omitempty"`
	Type                 *string `json:"type,omitempty"`
	SourceAccount        *string `json:"sourceAccount,omitempty"`
	AmtFromSpot          *string `json:"amtFromSpot,omitempty"`
	AmtFromFunding       *string `json:"amtFromFunding,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLockedSubscriptionRecordResponseRowsInner GetLockedSubscriptionRecordResponseRowsInner

// NewGetLockedSubscriptionRecordResponseRowsInner instantiates a new GetLockedSubscriptionRecordResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLockedSubscriptionRecordResponseRowsInner() *GetLockedSubscriptionRecordResponseRowsInner {
	this := GetLockedSubscriptionRecordResponseRowsInner{}
	return &this
}

// NewGetLockedSubscriptionRecordResponseRowsInnerWithDefaults instantiates a new GetLockedSubscriptionRecordResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLockedSubscriptionRecordResponseRowsInnerWithDefaults() *GetLockedSubscriptionRecordResponseRowsInner {
	this := GetLockedSubscriptionRecordResponseRowsInner{}
	return &this
}

// GetPositionId returns the PositionId field value if set, zero value otherwise.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetPositionId() int64 {
	if o == nil || common.IsNil(o.PositionId) {
		var ret int64
		return ret
	}
	return *o.PositionId
}

// GetPositionIdOk returns a tuple with the PositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetPositionIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PositionId) {
		return nil, false
	}
	return o.PositionId, true
}

// HasPositionId returns a boolean if a field has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) HasPositionId() bool {
	if o != nil && !common.IsNil(o.PositionId) {
		return true
	}

	return false
}

// SetPositionId gets a reference to the given int64 and assigns it to the PositionId field.
func (o *GetLockedSubscriptionRecordResponseRowsInner) SetPositionId(v int64) {
	o.PositionId = &v
}

// GetPurchaseId returns the PurchaseId field value if set, zero value otherwise.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetPurchaseId() string {
	if o == nil || common.IsNil(o.PurchaseId) {
		var ret string
		return ret
	}
	return *o.PurchaseId
}

// GetPurchaseIdOk returns a tuple with the PurchaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetPurchaseIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PurchaseId) {
		return nil, false
	}
	return o.PurchaseId, true
}

// HasPurchaseId returns a boolean if a field has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) HasPurchaseId() bool {
	if o != nil && !common.IsNil(o.PurchaseId) {
		return true
	}

	return false
}

// SetPurchaseId gets a reference to the given string and assigns it to the PurchaseId field.
func (o *GetLockedSubscriptionRecordResponseRowsInner) SetPurchaseId(v string) {
	o.PurchaseId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetProjectId() string {
	if o == nil || common.IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetProjectIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) HasProjectId() bool {
	if o != nil && !common.IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GetLockedSubscriptionRecordResponseRowsInner) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetLockedSubscriptionRecordResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetLockedSubscriptionRecordResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetLockedSubscriptionRecordResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetLockPeriod returns the LockPeriod field value if set, zero value otherwise.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetLockPeriod() string {
	if o == nil || common.IsNil(o.LockPeriod) {
		var ret string
		return ret
	}
	return *o.LockPeriod
}

// GetLockPeriodOk returns a tuple with the LockPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetLockPeriodOk() (*string, bool) {
	if o == nil || common.IsNil(o.LockPeriod) {
		return nil, false
	}
	return o.LockPeriod, true
}

// HasLockPeriod returns a boolean if a field has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) HasLockPeriod() bool {
	if o != nil && !common.IsNil(o.LockPeriod) {
		return true
	}

	return false
}

// SetLockPeriod gets a reference to the given string and assigns it to the LockPeriod field.
func (o *GetLockedSubscriptionRecordResponseRowsInner) SetLockPeriod(v string) {
	o.LockPeriod = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetLockedSubscriptionRecordResponseRowsInner) SetType(v string) {
	o.Type = &v
}

// GetSourceAccount returns the SourceAccount field value if set, zero value otherwise.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetSourceAccount() string {
	if o == nil || common.IsNil(o.SourceAccount) {
		var ret string
		return ret
	}
	return *o.SourceAccount
}

// GetSourceAccountOk returns a tuple with the SourceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetSourceAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.SourceAccount) {
		return nil, false
	}
	return o.SourceAccount, true
}

// HasSourceAccount returns a boolean if a field has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) HasSourceAccount() bool {
	if o != nil && !common.IsNil(o.SourceAccount) {
		return true
	}

	return false
}

// SetSourceAccount gets a reference to the given string and assigns it to the SourceAccount field.
func (o *GetLockedSubscriptionRecordResponseRowsInner) SetSourceAccount(v string) {
	o.SourceAccount = &v
}

// GetAmtFromSpot returns the AmtFromSpot field value if set, zero value otherwise.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetAmtFromSpot() string {
	if o == nil || common.IsNil(o.AmtFromSpot) {
		var ret string
		return ret
	}
	return *o.AmtFromSpot
}

// GetAmtFromSpotOk returns a tuple with the AmtFromSpot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetAmtFromSpotOk() (*string, bool) {
	if o == nil || common.IsNil(o.AmtFromSpot) {
		return nil, false
	}
	return o.AmtFromSpot, true
}

// HasAmtFromSpot returns a boolean if a field has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) HasAmtFromSpot() bool {
	if o != nil && !common.IsNil(o.AmtFromSpot) {
		return true
	}

	return false
}

// SetAmtFromSpot gets a reference to the given string and assigns it to the AmtFromSpot field.
func (o *GetLockedSubscriptionRecordResponseRowsInner) SetAmtFromSpot(v string) {
	o.AmtFromSpot = &v
}

// GetAmtFromFunding returns the AmtFromFunding field value if set, zero value otherwise.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetAmtFromFunding() string {
	if o == nil || common.IsNil(o.AmtFromFunding) {
		var ret string
		return ret
	}
	return *o.AmtFromFunding
}

// GetAmtFromFundingOk returns a tuple with the AmtFromFunding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetAmtFromFundingOk() (*string, bool) {
	if o == nil || common.IsNil(o.AmtFromFunding) {
		return nil, false
	}
	return o.AmtFromFunding, true
}

// HasAmtFromFunding returns a boolean if a field has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) HasAmtFromFunding() bool {
	if o != nil && !common.IsNil(o.AmtFromFunding) {
		return true
	}

	return false
}

// SetAmtFromFunding gets a reference to the given string and assigns it to the AmtFromFunding field.
func (o *GetLockedSubscriptionRecordResponseRowsInner) SetAmtFromFunding(v string) {
	o.AmtFromFunding = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetLockedSubscriptionRecordResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetLockedSubscriptionRecordResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetLockedSubscriptionRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLockedSubscriptionRecordResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PositionId) {
		toSerialize["positionId"] = o.PositionId
	}
	if !common.IsNil(o.PurchaseId) {
		toSerialize["purchaseId"] = o.PurchaseId
	}
	if !common.IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.LockPeriod) {
		toSerialize["lockPeriod"] = o.LockPeriod
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.SourceAccount) {
		toSerialize["sourceAccount"] = o.SourceAccount
	}
	if !common.IsNil(o.AmtFromSpot) {
		toSerialize["amtFromSpot"] = o.AmtFromSpot
	}
	if !common.IsNil(o.AmtFromFunding) {
		toSerialize["amtFromFunding"] = o.AmtFromFunding
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetLockedSubscriptionRecordResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetLockedSubscriptionRecordResponseRowsInner := _GetLockedSubscriptionRecordResponseRowsInner{}

	err = json.Unmarshal(data, &varGetLockedSubscriptionRecordResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetLockedSubscriptionRecordResponseRowsInner(varGetLockedSubscriptionRecordResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "positionId")
		delete(additionalProperties, "purchaseId")
		delete(additionalProperties, "projectId")
		delete(additionalProperties, "time")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "lockPeriod")
		delete(additionalProperties, "type")
		delete(additionalProperties, "sourceAccount")
		delete(additionalProperties, "amtFromSpot")
		delete(additionalProperties, "amtFromFunding")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLockedSubscriptionRecordResponseRowsInner struct {
	value *GetLockedSubscriptionRecordResponseRowsInner
	isSet bool
}

func (v NullableGetLockedSubscriptionRecordResponseRowsInner) Get() *GetLockedSubscriptionRecordResponseRowsInner {
	return v.value
}

func (v *NullableGetLockedSubscriptionRecordResponseRowsInner) Set(val *GetLockedSubscriptionRecordResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLockedSubscriptionRecordResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLockedSubscriptionRecordResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLockedSubscriptionRecordResponseRowsInner(val *GetLockedSubscriptionRecordResponseRowsInner) *NullableGetLockedSubscriptionRecordResponseRowsInner {
	return &NullableGetLockedSubscriptionRecordResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetLockedSubscriptionRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLockedSubscriptionRecordResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
