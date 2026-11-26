/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CorporateActionRecordsResponseItemsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CorporateActionRecordsResponseItemsInner{}

// CorporateActionRecordsResponseItemsInner struct for CorporateActionRecordsResponseItemsInner
type CorporateActionRecordsResponseItemsInner struct {
	// Unique record ID.
	Id *int64 `json:"id,omitempty"`
	// User ID.
	UserId *int64 `json:"userId,omitempty"`
	// Asset code, the unique identifier of the stock in the system, for example `EQ_AAPL`. Currently all asset codes carry the `EQ_` prefix.
	AssetCode *string `json:"assetCode,omitempty"`
	// Corporate action event type. Values include `CASH_DIVIDEND`, `STOCK_DIVIDEND`, `FORWARD_SPLIT`, `REVERSE_SPLIT`, `SPIN_OFF`, `CASH_MERGER`, `STOCK_MERGER`, `STOCK_AND_CASH_MERGER`, `WORTHLESS_REMOVAL`, `NAME_CHANGE` and `RIGHTS_DISTRIBUTION`.
	EventType *string `json:"eventType,omitempty"`
	// Signed change of the asset quantity, including free and frozen balances.
	Amount *string `json:"amount,omitempty"`
	// Last modification time (ms epoch).
	DbModifyTime         *int64 `json:"dbModifyTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CorporateActionRecordsResponseItemsInner CorporateActionRecordsResponseItemsInner

// NewCorporateActionRecordsResponseItemsInner instantiates a new CorporateActionRecordsResponseItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorporateActionRecordsResponseItemsInner() *CorporateActionRecordsResponseItemsInner {
	this := CorporateActionRecordsResponseItemsInner{}
	return &this
}

// NewCorporateActionRecordsResponseItemsInnerWithDefaults instantiates a new CorporateActionRecordsResponseItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorporateActionRecordsResponseItemsInnerWithDefaults() *CorporateActionRecordsResponseItemsInner {
	this := CorporateActionRecordsResponseItemsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CorporateActionRecordsResponseItemsInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporateActionRecordsResponseItemsInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CorporateActionRecordsResponseItemsInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CorporateActionRecordsResponseItemsInner) SetId(v int64) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *CorporateActionRecordsResponseItemsInner) GetUserId() int64 {
	if o == nil || common.IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporateActionRecordsResponseItemsInner) GetUserIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *CorporateActionRecordsResponseItemsInner) HasUserId() bool {
	if o != nil && !common.IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *CorporateActionRecordsResponseItemsInner) SetUserId(v int64) {
	o.UserId = &v
}

// GetAssetCode returns the AssetCode field value if set, zero value otherwise.
func (o *CorporateActionRecordsResponseItemsInner) GetAssetCode() string {
	if o == nil || common.IsNil(o.AssetCode) {
		var ret string
		return ret
	}
	return *o.AssetCode
}

// GetAssetCodeOk returns a tuple with the AssetCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporateActionRecordsResponseItemsInner) GetAssetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AssetCode) {
		return nil, false
	}
	return o.AssetCode, true
}

// HasAssetCode returns a boolean if a field has been set.
func (o *CorporateActionRecordsResponseItemsInner) HasAssetCode() bool {
	if o != nil && !common.IsNil(o.AssetCode) {
		return true
	}

	return false
}

// SetAssetCode gets a reference to the given string and assigns it to the AssetCode field.
func (o *CorporateActionRecordsResponseItemsInner) SetAssetCode(v string) {
	o.AssetCode = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CorporateActionRecordsResponseItemsInner) GetEventType() string {
	if o == nil || common.IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporateActionRecordsResponseItemsInner) GetEventTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CorporateActionRecordsResponseItemsInner) HasEventType() bool {
	if o != nil && !common.IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CorporateActionRecordsResponseItemsInner) SetEventType(v string) {
	o.EventType = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CorporateActionRecordsResponseItemsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporateActionRecordsResponseItemsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CorporateActionRecordsResponseItemsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *CorporateActionRecordsResponseItemsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetDbModifyTime returns the DbModifyTime field value if set, zero value otherwise.
func (o *CorporateActionRecordsResponseItemsInner) GetDbModifyTime() int64 {
	if o == nil || common.IsNil(o.DbModifyTime) {
		var ret int64
		return ret
	}
	return *o.DbModifyTime
}

// GetDbModifyTimeOk returns a tuple with the DbModifyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporateActionRecordsResponseItemsInner) GetDbModifyTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.DbModifyTime) {
		return nil, false
	}
	return o.DbModifyTime, true
}

// HasDbModifyTime returns a boolean if a field has been set.
func (o *CorporateActionRecordsResponseItemsInner) HasDbModifyTime() bool {
	if o != nil && !common.IsNil(o.DbModifyTime) {
		return true
	}

	return false
}

// SetDbModifyTime gets a reference to the given int64 and assigns it to the DbModifyTime field.
func (o *CorporateActionRecordsResponseItemsInner) SetDbModifyTime(v int64) {
	o.DbModifyTime = &v
}

func (o CorporateActionRecordsResponseItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CorporateActionRecordsResponseItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !common.IsNil(o.AssetCode) {
		toSerialize["assetCode"] = o.AssetCode
	}
	if !common.IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.DbModifyTime) {
		toSerialize["dbModifyTime"] = o.DbModifyTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CorporateActionRecordsResponseItemsInner) UnmarshalJSON(data []byte) (err error) {
	varCorporateActionRecordsResponseItemsInner := _CorporateActionRecordsResponseItemsInner{}

	err = json.Unmarshal(data, &varCorporateActionRecordsResponseItemsInner)

	if err != nil {
		return err
	}

	*o = CorporateActionRecordsResponseItemsInner(varCorporateActionRecordsResponseItemsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "assetCode")
		delete(additionalProperties, "eventType")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "dbModifyTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCorporateActionRecordsResponseItemsInner struct {
	value *CorporateActionRecordsResponseItemsInner
	isSet bool
}

func (v NullableCorporateActionRecordsResponseItemsInner) Get() *CorporateActionRecordsResponseItemsInner {
	return v.value
}

func (v *NullableCorporateActionRecordsResponseItemsInner) Set(val *CorporateActionRecordsResponseItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCorporateActionRecordsResponseItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCorporateActionRecordsResponseItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorporateActionRecordsResponseItemsInner(val *CorporateActionRecordsResponseItemsInner) *NullableCorporateActionRecordsResponseItemsInner {
	return &NullableCorporateActionRecordsResponseItemsInner{value: val, isSet: true}
}

func (v NullableCorporateActionRecordsResponseItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorporateActionRecordsResponseItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
