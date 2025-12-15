/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner{}

// QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner struct for QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner
type QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner struct {
	FromEmail            *string `json:"fromEmail,omitempty"`
	FromAccountType      *string `json:"fromAccountType,omitempty"`
	ToEmail              *string `json:"toEmail,omitempty"`
	ToAccountType        *string `json:"toAccountType,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	ScheduledData        *int64  `json:"scheduledData,omitempty"`
	CreateTime           *int64  `json:"createTime,omitempty"`
	Status               *string `json:"status,omitempty"`
	TranId               *int64  `json:"tranId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner

// NewQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner instantiates a new QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner() *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner {
	this := QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner{}
	return &this
}

// NewQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInnerWithDefaults instantiates a new QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInnerWithDefaults() *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner {
	this := QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner{}
	return &this
}

// GetFromEmail returns the FromEmail field value if set, zero value otherwise.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetFromEmail() string {
	if o == nil || common.IsNil(o.FromEmail) {
		var ret string
		return ret
	}
	return *o.FromEmail
}

// GetFromEmailOk returns a tuple with the FromEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetFromEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromEmail) {
		return nil, false
	}
	return o.FromEmail, true
}

// HasFromEmail returns a boolean if a field has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) HasFromEmail() bool {
	if o != nil && !common.IsNil(o.FromEmail) {
		return true
	}

	return false
}

// SetFromEmail gets a reference to the given string and assigns it to the FromEmail field.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) SetFromEmail(v string) {
	o.FromEmail = &v
}

// GetFromAccountType returns the FromAccountType field value if set, zero value otherwise.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetFromAccountType() string {
	if o == nil || common.IsNil(o.FromAccountType) {
		var ret string
		return ret
	}
	return *o.FromAccountType
}

// GetFromAccountTypeOk returns a tuple with the FromAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetFromAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAccountType) {
		return nil, false
	}
	return o.FromAccountType, true
}

// HasFromAccountType returns a boolean if a field has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) HasFromAccountType() bool {
	if o != nil && !common.IsNil(o.FromAccountType) {
		return true
	}

	return false
}

// SetFromAccountType gets a reference to the given string and assigns it to the FromAccountType field.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) SetFromAccountType(v string) {
	o.FromAccountType = &v
}

// GetToEmail returns the ToEmail field value if set, zero value otherwise.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetToEmail() string {
	if o == nil || common.IsNil(o.ToEmail) {
		var ret string
		return ret
	}
	return *o.ToEmail
}

// GetToEmailOk returns a tuple with the ToEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetToEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToEmail) {
		return nil, false
	}
	return o.ToEmail, true
}

// HasToEmail returns a boolean if a field has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) HasToEmail() bool {
	if o != nil && !common.IsNil(o.ToEmail) {
		return true
	}

	return false
}

// SetToEmail gets a reference to the given string and assigns it to the ToEmail field.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) SetToEmail(v string) {
	o.ToEmail = &v
}

// GetToAccountType returns the ToAccountType field value if set, zero value otherwise.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetToAccountType() string {
	if o == nil || common.IsNil(o.ToAccountType) {
		var ret string
		return ret
	}
	return *o.ToAccountType
}

// GetToAccountTypeOk returns a tuple with the ToAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetToAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAccountType) {
		return nil, false
	}
	return o.ToAccountType, true
}

// HasToAccountType returns a boolean if a field has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) HasToAccountType() bool {
	if o != nil && !common.IsNil(o.ToAccountType) {
		return true
	}

	return false
}

// SetToAccountType gets a reference to the given string and assigns it to the ToAccountType field.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) SetToAccountType(v string) {
	o.ToAccountType = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) SetAmount(v string) {
	o.Amount = &v
}

// GetScheduledData returns the ScheduledData field value if set, zero value otherwise.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetScheduledData() int64 {
	if o == nil || common.IsNil(o.ScheduledData) {
		var ret int64
		return ret
	}
	return *o.ScheduledData
}

// GetScheduledDataOk returns a tuple with the ScheduledData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetScheduledDataOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ScheduledData) {
		return nil, false
	}
	return o.ScheduledData, true
}

// HasScheduledData returns a boolean if a field has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) HasScheduledData() bool {
	if o != nil && !common.IsNil(o.ScheduledData) {
		return true
	}

	return false
}

// SetScheduledData gets a reference to the given int64 and assigns it to the ScheduledData field.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) SetScheduledData(v int64) {
	o.ScheduledData = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) SetStatus(v string) {
	o.Status = &v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) SetTranId(v int64) {
	o.TranId = &v
}

func (o QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FromEmail) {
		toSerialize["fromEmail"] = o.FromEmail
	}
	if !common.IsNil(o.FromAccountType) {
		toSerialize["fromAccountType"] = o.FromAccountType
	}
	if !common.IsNil(o.ToEmail) {
		toSerialize["toEmail"] = o.ToEmail
	}
	if !common.IsNil(o.ToAccountType) {
		toSerialize["toAccountType"] = o.ToAccountType
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.ScheduledData) {
		toSerialize["scheduledData"] = o.ScheduledData
	}
	if !common.IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner := _QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner(varQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fromEmail")
		delete(additionalProperties, "fromAccountType")
		delete(additionalProperties, "toEmail")
		delete(additionalProperties, "toAccountType")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "scheduledData")
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tranId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner struct {
	value *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner
	isSet bool
}

func (v NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) Get() *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner {
	return v.value
}

func (v *NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) Set(val *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner(val *QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) *NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner {
	return &NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
