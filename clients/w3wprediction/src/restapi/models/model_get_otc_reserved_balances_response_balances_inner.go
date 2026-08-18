/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOtcReservedBalancesResponseBalancesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOtcReservedBalancesResponseBalancesInner{}

// GetOtcReservedBalancesResponseBalancesInner struct for GetOtcReservedBalancesResponseBalancesInner
type GetOtcReservedBalancesResponseBalancesInner struct {
	Type                 *string `json:"type,omitempty"`
	TokenId              *string `json:"tokenId,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOtcReservedBalancesResponseBalancesInner GetOtcReservedBalancesResponseBalancesInner

// NewGetOtcReservedBalancesResponseBalancesInner instantiates a new GetOtcReservedBalancesResponseBalancesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOtcReservedBalancesResponseBalancesInner() *GetOtcReservedBalancesResponseBalancesInner {
	this := GetOtcReservedBalancesResponseBalancesInner{}
	return &this
}

// NewGetOtcReservedBalancesResponseBalancesInnerWithDefaults instantiates a new GetOtcReservedBalancesResponseBalancesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOtcReservedBalancesResponseBalancesInnerWithDefaults() *GetOtcReservedBalancesResponseBalancesInner {
	this := GetOtcReservedBalancesResponseBalancesInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetOtcReservedBalancesResponseBalancesInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcReservedBalancesResponseBalancesInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetOtcReservedBalancesResponseBalancesInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetOtcReservedBalancesResponseBalancesInner) SetType(v string) {
	o.Type = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetOtcReservedBalancesResponseBalancesInner) GetTokenId() string {
	if o == nil || common.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetOtcReservedBalancesResponseBalancesInner) GetTokenIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *GetOtcReservedBalancesResponseBalancesInner) HasTokenId() bool {
	if o != nil && !common.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given NullableString and assigns it to the TokenId field.
func (o *GetOtcReservedBalancesResponseBalancesInner) SetTokenId(v string) {
	o.TokenId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetOtcReservedBalancesResponseBalancesInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcReservedBalancesResponseBalancesInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetOtcReservedBalancesResponseBalancesInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetOtcReservedBalancesResponseBalancesInner) SetAmount(v string) {
	o.Amount = &v
}

func (o GetOtcReservedBalancesResponseBalancesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOtcReservedBalancesResponseBalancesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOtcReservedBalancesResponseBalancesInner) UnmarshalJSON(data []byte) (err error) {
	varGetOtcReservedBalancesResponseBalancesInner := _GetOtcReservedBalancesResponseBalancesInner{}

	err = json.Unmarshal(data, &varGetOtcReservedBalancesResponseBalancesInner)

	if err != nil {
		return err
	}

	*o = GetOtcReservedBalancesResponseBalancesInner(varGetOtcReservedBalancesResponseBalancesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "tokenId")
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOtcReservedBalancesResponseBalancesInner struct {
	value *GetOtcReservedBalancesResponseBalancesInner
	isSet bool
}

func (v NullableGetOtcReservedBalancesResponseBalancesInner) Get() *GetOtcReservedBalancesResponseBalancesInner {
	return v.value
}

func (v *NullableGetOtcReservedBalancesResponseBalancesInner) Set(val *GetOtcReservedBalancesResponseBalancesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOtcReservedBalancesResponseBalancesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOtcReservedBalancesResponseBalancesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOtcReservedBalancesResponseBalancesInner(val *GetOtcReservedBalancesResponseBalancesInner) *NullableGetOtcReservedBalancesResponseBalancesInner {
	return &NullableGetOtcReservedBalancesResponseBalancesInner{value: val, isSet: true}
}

func (v NullableGetOtcReservedBalancesResponseBalancesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOtcReservedBalancesResponseBalancesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
