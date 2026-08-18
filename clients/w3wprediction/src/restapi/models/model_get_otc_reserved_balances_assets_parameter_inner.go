/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOtcReservedBalancesAssetsParameterInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOtcReservedBalancesAssetsParameterInner{}

// GetOtcReservedBalancesAssetsParameterInner struct for GetOtcReservedBalancesAssetsParameterInner
type GetOtcReservedBalancesAssetsParameterInner struct {
	Type                 *GetOtcReservedBalancesAssetsParameterInnerType `json:"type,omitempty"`
	TokenId              *string                                         `json:"tokenId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOtcReservedBalancesAssetsParameterInner GetOtcReservedBalancesAssetsParameterInner

// NewGetOtcReservedBalancesAssetsParameterInner instantiates a new GetOtcReservedBalancesAssetsParameterInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOtcReservedBalancesAssetsParameterInner() *GetOtcReservedBalancesAssetsParameterInner {
	this := GetOtcReservedBalancesAssetsParameterInner{}
	var type_ = GetOtcReservedBalancesAssetsParameterInnerTypeUsdt
	this.Type = &type_
	return &this
}

// NewGetOtcReservedBalancesAssetsParameterInnerWithDefaults instantiates a new GetOtcReservedBalancesAssetsParameterInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOtcReservedBalancesAssetsParameterInnerWithDefaults() *GetOtcReservedBalancesAssetsParameterInner {
	this := GetOtcReservedBalancesAssetsParameterInner{}
	var type_ = GetOtcReservedBalancesAssetsParameterInnerTypeUsdt
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetOtcReservedBalancesAssetsParameterInner) GetType() GetOtcReservedBalancesAssetsParameterInnerType {
	if o == nil || common.IsNil(o.Type) {
		var ret GetOtcReservedBalancesAssetsParameterInnerType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcReservedBalancesAssetsParameterInner) GetTypeOk() (*GetOtcReservedBalancesAssetsParameterInnerType, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetOtcReservedBalancesAssetsParameterInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given GetOtcReservedBalancesAssetsParameterInnerType and assigns it to the Type field.
func (o *GetOtcReservedBalancesAssetsParameterInner) SetType(v GetOtcReservedBalancesAssetsParameterInnerType) {
	o.Type = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *GetOtcReservedBalancesAssetsParameterInner) GetTokenId() string {
	if o == nil || common.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcReservedBalancesAssetsParameterInner) GetTokenIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *GetOtcReservedBalancesAssetsParameterInner) HasTokenId() bool {
	if o != nil && !common.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *GetOtcReservedBalancesAssetsParameterInner) SetTokenId(v string) {
	o.TokenId = &v
}

func (o GetOtcReservedBalancesAssetsParameterInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOtcReservedBalancesAssetsParameterInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOtcReservedBalancesAssetsParameterInner) UnmarshalJSON(data []byte) (err error) {
	varGetOtcReservedBalancesAssetsParameterInner := _GetOtcReservedBalancesAssetsParameterInner{}

	err = json.Unmarshal(data, &varGetOtcReservedBalancesAssetsParameterInner)

	if err != nil {
		return err
	}

	*o = GetOtcReservedBalancesAssetsParameterInner(varGetOtcReservedBalancesAssetsParameterInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "tokenId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOtcReservedBalancesAssetsParameterInner struct {
	value *GetOtcReservedBalancesAssetsParameterInner
	isSet bool
}

func (v NullableGetOtcReservedBalancesAssetsParameterInner) Get() *GetOtcReservedBalancesAssetsParameterInner {
	return v.value
}

func (v *NullableGetOtcReservedBalancesAssetsParameterInner) Set(val *GetOtcReservedBalancesAssetsParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOtcReservedBalancesAssetsParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOtcReservedBalancesAssetsParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOtcReservedBalancesAssetsParameterInner(val *GetOtcReservedBalancesAssetsParameterInner) *NullableGetOtcReservedBalancesAssetsParameterInner {
	return &NullableGetOtcReservedBalancesAssetsParameterInner{value: val, isSet: true}
}

func (v NullableGetOtcReservedBalancesAssetsParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOtcReservedBalancesAssetsParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
