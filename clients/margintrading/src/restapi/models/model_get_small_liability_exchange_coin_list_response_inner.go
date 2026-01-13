/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSmallLiabilityExchangeCoinListResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSmallLiabilityExchangeCoinListResponseInner{}

// GetSmallLiabilityExchangeCoinListResponseInner struct for GetSmallLiabilityExchangeCoinListResponseInner
type GetSmallLiabilityExchangeCoinListResponseInner struct {
	Asset                *string  `json:"asset,omitempty"`
	Interest             *string  `json:"interest,omitempty"`
	Principal            *string  `json:"principal,omitempty"`
	LiabilityAsset       *string  `json:"liabilityAsset,omitempty"`
	LiabilityQty         *float32 `json:"liabilityQty,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSmallLiabilityExchangeCoinListResponseInner GetSmallLiabilityExchangeCoinListResponseInner

// NewGetSmallLiabilityExchangeCoinListResponseInner instantiates a new GetSmallLiabilityExchangeCoinListResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSmallLiabilityExchangeCoinListResponseInner() *GetSmallLiabilityExchangeCoinListResponseInner {
	this := GetSmallLiabilityExchangeCoinListResponseInner{}
	return &this
}

// NewGetSmallLiabilityExchangeCoinListResponseInnerWithDefaults instantiates a new GetSmallLiabilityExchangeCoinListResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSmallLiabilityExchangeCoinListResponseInnerWithDefaults() *GetSmallLiabilityExchangeCoinListResponseInner {
	this := GetSmallLiabilityExchangeCoinListResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) GetInterest() string {
	if o == nil || common.IsNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) GetInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) HasInterest() bool {
	if o != nil && !common.IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) SetInterest(v string) {
	o.Interest = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) GetPrincipal() string {
	if o == nil || common.IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) GetPrincipalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) HasPrincipal() bool {
	if o != nil && !common.IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) SetPrincipal(v string) {
	o.Principal = &v
}

// GetLiabilityAsset returns the LiabilityAsset field value if set, zero value otherwise.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) GetLiabilityAsset() string {
	if o == nil || common.IsNil(o.LiabilityAsset) {
		var ret string
		return ret
	}
	return *o.LiabilityAsset
}

// GetLiabilityAssetOk returns a tuple with the LiabilityAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) GetLiabilityAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiabilityAsset) {
		return nil, false
	}
	return o.LiabilityAsset, true
}

// HasLiabilityAsset returns a boolean if a field has been set.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) HasLiabilityAsset() bool {
	if o != nil && !common.IsNil(o.LiabilityAsset) {
		return true
	}

	return false
}

// SetLiabilityAsset gets a reference to the given string and assigns it to the LiabilityAsset field.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) SetLiabilityAsset(v string) {
	o.LiabilityAsset = &v
}

// GetLiabilityQty returns the LiabilityQty field value if set, zero value otherwise.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) GetLiabilityQty() float32 {
	if o == nil || common.IsNil(o.LiabilityQty) {
		var ret float32
		return ret
	}
	return *o.LiabilityQty
}

// GetLiabilityQtyOk returns a tuple with the LiabilityQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) GetLiabilityQtyOk() (*float32, bool) {
	if o == nil || common.IsNil(o.LiabilityQty) {
		return nil, false
	}
	return o.LiabilityQty, true
}

// HasLiabilityQty returns a boolean if a field has been set.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) HasLiabilityQty() bool {
	if o != nil && !common.IsNil(o.LiabilityQty) {
		return true
	}

	return false
}

// SetLiabilityQty gets a reference to the given float32 and assigns it to the LiabilityQty field.
func (o *GetSmallLiabilityExchangeCoinListResponseInner) SetLiabilityQty(v float32) {
	o.LiabilityQty = &v
}

func (o GetSmallLiabilityExchangeCoinListResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSmallLiabilityExchangeCoinListResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !common.IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}
	if !common.IsNil(o.LiabilityAsset) {
		toSerialize["liabilityAsset"] = o.LiabilityAsset
	}
	if !common.IsNil(o.LiabilityQty) {
		toSerialize["liabilityQty"] = o.LiabilityQty
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSmallLiabilityExchangeCoinListResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetSmallLiabilityExchangeCoinListResponseInner := _GetSmallLiabilityExchangeCoinListResponseInner{}

	err = json.Unmarshal(data, &varGetSmallLiabilityExchangeCoinListResponseInner)

	if err != nil {
		return err
	}

	*o = GetSmallLiabilityExchangeCoinListResponseInner(varGetSmallLiabilityExchangeCoinListResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "interest")
		delete(additionalProperties, "principal")
		delete(additionalProperties, "liabilityAsset")
		delete(additionalProperties, "liabilityQty")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSmallLiabilityExchangeCoinListResponseInner struct {
	value *GetSmallLiabilityExchangeCoinListResponseInner
	isSet bool
}

func (v NullableGetSmallLiabilityExchangeCoinListResponseInner) Get() *GetSmallLiabilityExchangeCoinListResponseInner {
	return v.value
}

func (v *NullableGetSmallLiabilityExchangeCoinListResponseInner) Set(val *GetSmallLiabilityExchangeCoinListResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSmallLiabilityExchangeCoinListResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSmallLiabilityExchangeCoinListResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSmallLiabilityExchangeCoinListResponseInner(val *GetSmallLiabilityExchangeCoinListResponseInner) *NullableGetSmallLiabilityExchangeCoinListResponseInner {
	return &NullableGetSmallLiabilityExchangeCoinListResponseInner{value: val, isSet: true}
}

func (v NullableGetSmallLiabilityExchangeCoinListResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSmallLiabilityExchangeCoinListResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
