/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountFundingFlowResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountFundingFlowResponseInner{}

// AccountFundingFlowResponseInner struct for AccountFundingFlowResponseInner
type AccountFundingFlowResponseInner struct {
	Id                   *int64  `json:"id,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	Type                 *string `json:"type,omitempty"`
	CreateDate           *int64  `json:"createDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountFundingFlowResponseInner AccountFundingFlowResponseInner

// NewAccountFundingFlowResponseInner instantiates a new AccountFundingFlowResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountFundingFlowResponseInner() *AccountFundingFlowResponseInner {
	this := AccountFundingFlowResponseInner{}
	return &this
}

// NewAccountFundingFlowResponseInnerWithDefaults instantiates a new AccountFundingFlowResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountFundingFlowResponseInnerWithDefaults() *AccountFundingFlowResponseInner {
	this := AccountFundingFlowResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountFundingFlowResponseInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFundingFlowResponseInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountFundingFlowResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AccountFundingFlowResponseInner) SetId(v int64) {
	o.Id = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *AccountFundingFlowResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFundingFlowResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *AccountFundingFlowResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *AccountFundingFlowResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AccountFundingFlowResponseInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFundingFlowResponseInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AccountFundingFlowResponseInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AccountFundingFlowResponseInner) SetAmount(v string) {
	o.Amount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccountFundingFlowResponseInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFundingFlowResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccountFundingFlowResponseInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccountFundingFlowResponseInner) SetType(v string) {
	o.Type = &v
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise.
func (o *AccountFundingFlowResponseInner) GetCreateDate() int64 {
	if o == nil || common.IsNil(o.CreateDate) {
		var ret int64
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFundingFlowResponseInner) GetCreateDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateDate) {
		return nil, false
	}
	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *AccountFundingFlowResponseInner) HasCreateDate() bool {
	if o != nil && !common.IsNil(o.CreateDate) {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given int64 and assigns it to the CreateDate field.
func (o *AccountFundingFlowResponseInner) SetCreateDate(v int64) {
	o.CreateDate = &v
}

func (o AccountFundingFlowResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountFundingFlowResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.CreateDate) {
		toSerialize["createDate"] = o.CreateDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountFundingFlowResponseInner) UnmarshalJSON(data []byte) (err error) {
	varAccountFundingFlowResponseInner := _AccountFundingFlowResponseInner{}

	err = json.Unmarshal(data, &varAccountFundingFlowResponseInner)

	if err != nil {
		return err
	}

	*o = AccountFundingFlowResponseInner(varAccountFundingFlowResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "type")
		delete(additionalProperties, "createDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountFundingFlowResponseInner struct {
	value *AccountFundingFlowResponseInner
	isSet bool
}

func (v NullableAccountFundingFlowResponseInner) Get() *AccountFundingFlowResponseInner {
	return v.value
}

func (v *NullableAccountFundingFlowResponseInner) Set(val *AccountFundingFlowResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountFundingFlowResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountFundingFlowResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountFundingFlowResponseInner(val *AccountFundingFlowResponseInner) *NullableAccountFundingFlowResponseInner {
	return &NullableAccountFundingFlowResponseInner{value: val, isSet: true}
}

func (v NullableAccountFundingFlowResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountFundingFlowResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
