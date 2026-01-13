/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountBlockTradeListResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountBlockTradeListResponseInner{}

// AccountBlockTradeListResponseInner struct for AccountBlockTradeListResponseInner
type AccountBlockTradeListResponseInner struct {
	ParentOrderId           *string                                       `json:"parentOrderId,omitempty"`
	CrossType               *string                                       `json:"crossType,omitempty"`
	Legs                    []AccountBlockTradeListResponseInnerLegsInner `json:"legs,omitempty"`
	BlockTradeSettlementKey *string                                       `json:"blockTradeSettlementKey,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _AccountBlockTradeListResponseInner AccountBlockTradeListResponseInner

// NewAccountBlockTradeListResponseInner instantiates a new AccountBlockTradeListResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBlockTradeListResponseInner() *AccountBlockTradeListResponseInner {
	this := AccountBlockTradeListResponseInner{}
	return &this
}

// NewAccountBlockTradeListResponseInnerWithDefaults instantiates a new AccountBlockTradeListResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBlockTradeListResponseInnerWithDefaults() *AccountBlockTradeListResponseInner {
	this := AccountBlockTradeListResponseInner{}
	return &this
}

// GetParentOrderId returns the ParentOrderId field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInner) GetParentOrderId() string {
	if o == nil || common.IsNil(o.ParentOrderId) {
		var ret string
		return ret
	}
	return *o.ParentOrderId
}

// GetParentOrderIdOk returns a tuple with the ParentOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInner) GetParentOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ParentOrderId) {
		return nil, false
	}
	return o.ParentOrderId, true
}

// HasParentOrderId returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInner) HasParentOrderId() bool {
	if o != nil && !common.IsNil(o.ParentOrderId) {
		return true
	}

	return false
}

// SetParentOrderId gets a reference to the given string and assigns it to the ParentOrderId field.
func (o *AccountBlockTradeListResponseInner) SetParentOrderId(v string) {
	o.ParentOrderId = &v
}

// GetCrossType returns the CrossType field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInner) GetCrossType() string {
	if o == nil || common.IsNil(o.CrossType) {
		var ret string
		return ret
	}
	return *o.CrossType
}

// GetCrossTypeOk returns a tuple with the CrossType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInner) GetCrossTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossType) {
		return nil, false
	}
	return o.CrossType, true
}

// HasCrossType returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInner) HasCrossType() bool {
	if o != nil && !common.IsNil(o.CrossType) {
		return true
	}

	return false
}

// SetCrossType gets a reference to the given string and assigns it to the CrossType field.
func (o *AccountBlockTradeListResponseInner) SetCrossType(v string) {
	o.CrossType = &v
}

// GetLegs returns the Legs field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInner) GetLegs() []AccountBlockTradeListResponseInnerLegsInner {
	if o == nil || common.IsNil(o.Legs) {
		var ret []AccountBlockTradeListResponseInnerLegsInner
		return ret
	}
	return o.Legs
}

// GetLegsOk returns a tuple with the Legs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInner) GetLegsOk() ([]AccountBlockTradeListResponseInnerLegsInner, bool) {
	if o == nil || common.IsNil(o.Legs) {
		return nil, false
	}
	return o.Legs, true
}

// HasLegs returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInner) HasLegs() bool {
	if o != nil && !common.IsNil(o.Legs) {
		return true
	}

	return false
}

// SetLegs gets a reference to the given []AccountBlockTradeListResponseInnerLegsInner and assigns it to the Legs field.
func (o *AccountBlockTradeListResponseInner) SetLegs(v []AccountBlockTradeListResponseInnerLegsInner) {
	o.Legs = v
}

// GetBlockTradeSettlementKey returns the BlockTradeSettlementKey field value if set, zero value otherwise.
func (o *AccountBlockTradeListResponseInner) GetBlockTradeSettlementKey() string {
	if o == nil || common.IsNil(o.BlockTradeSettlementKey) {
		var ret string
		return ret
	}
	return *o.BlockTradeSettlementKey
}

// GetBlockTradeSettlementKeyOk returns a tuple with the BlockTradeSettlementKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBlockTradeListResponseInner) GetBlockTradeSettlementKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.BlockTradeSettlementKey) {
		return nil, false
	}
	return o.BlockTradeSettlementKey, true
}

// HasBlockTradeSettlementKey returns a boolean if a field has been set.
func (o *AccountBlockTradeListResponseInner) HasBlockTradeSettlementKey() bool {
	if o != nil && !common.IsNil(o.BlockTradeSettlementKey) {
		return true
	}

	return false
}

// SetBlockTradeSettlementKey gets a reference to the given string and assigns it to the BlockTradeSettlementKey field.
func (o *AccountBlockTradeListResponseInner) SetBlockTradeSettlementKey(v string) {
	o.BlockTradeSettlementKey = &v
}

func (o AccountBlockTradeListResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountBlockTradeListResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ParentOrderId) {
		toSerialize["parentOrderId"] = o.ParentOrderId
	}
	if !common.IsNil(o.CrossType) {
		toSerialize["crossType"] = o.CrossType
	}
	if !common.IsNil(o.Legs) {
		toSerialize["legs"] = o.Legs
	}
	if !common.IsNil(o.BlockTradeSettlementKey) {
		toSerialize["blockTradeSettlementKey"] = o.BlockTradeSettlementKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountBlockTradeListResponseInner) UnmarshalJSON(data []byte) (err error) {
	varAccountBlockTradeListResponseInner := _AccountBlockTradeListResponseInner{}

	err = json.Unmarshal(data, &varAccountBlockTradeListResponseInner)

	if err != nil {
		return err
	}

	*o = AccountBlockTradeListResponseInner(varAccountBlockTradeListResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "parentOrderId")
		delete(additionalProperties, "crossType")
		delete(additionalProperties, "legs")
		delete(additionalProperties, "blockTradeSettlementKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountBlockTradeListResponseInner struct {
	value *AccountBlockTradeListResponseInner
	isSet bool
}

func (v NullableAccountBlockTradeListResponseInner) Get() *AccountBlockTradeListResponseInner {
	return v.value
}

func (v *NullableAccountBlockTradeListResponseInner) Set(val *AccountBlockTradeListResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBlockTradeListResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBlockTradeListResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBlockTradeListResponseInner(val *AccountBlockTradeListResponseInner) *NullableAccountBlockTradeListResponseInner {
	return &NullableAccountBlockTradeListResponseInner{value: val, isSet: true}
}

func (v NullableAccountBlockTradeListResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBlockTradeListResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
