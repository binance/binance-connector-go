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

// checks if the GetRwusdQuotaDetailsResponseSubscriptionQuota type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetRwusdQuotaDetailsResponseSubscriptionQuota{}

// GetRwusdQuotaDetailsResponseSubscriptionQuota struct for GetRwusdQuotaDetailsResponseSubscriptionQuota
type GetRwusdQuotaDetailsResponseSubscriptionQuota struct {
	Assets               []string `json:"assets,omitempty"`
	LeftQuota            *string  `json:"leftQuota,omitempty"`
	Minimum              *string  `json:"minimum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRwusdQuotaDetailsResponseSubscriptionQuota GetRwusdQuotaDetailsResponseSubscriptionQuota

// NewGetRwusdQuotaDetailsResponseSubscriptionQuota instantiates a new GetRwusdQuotaDetailsResponseSubscriptionQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRwusdQuotaDetailsResponseSubscriptionQuota() *GetRwusdQuotaDetailsResponseSubscriptionQuota {
	this := GetRwusdQuotaDetailsResponseSubscriptionQuota{}
	return &this
}

// NewGetRwusdQuotaDetailsResponseSubscriptionQuotaWithDefaults instantiates a new GetRwusdQuotaDetailsResponseSubscriptionQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRwusdQuotaDetailsResponseSubscriptionQuotaWithDefaults() *GetRwusdQuotaDetailsResponseSubscriptionQuota {
	this := GetRwusdQuotaDetailsResponseSubscriptionQuota{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *GetRwusdQuotaDetailsResponseSubscriptionQuota) GetAssets() []string {
	if o == nil || common.IsNil(o.Assets) {
		var ret []string
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdQuotaDetailsResponseSubscriptionQuota) GetAssetsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *GetRwusdQuotaDetailsResponseSubscriptionQuota) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []string and assigns it to the Assets field.
func (o *GetRwusdQuotaDetailsResponseSubscriptionQuota) SetAssets(v []string) {
	o.Assets = v
}

// GetLeftQuota returns the LeftQuota field value if set, zero value otherwise.
func (o *GetRwusdQuotaDetailsResponseSubscriptionQuota) GetLeftQuota() string {
	if o == nil || common.IsNil(o.LeftQuota) {
		var ret string
		return ret
	}
	return *o.LeftQuota
}

// GetLeftQuotaOk returns a tuple with the LeftQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdQuotaDetailsResponseSubscriptionQuota) GetLeftQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.LeftQuota) {
		return nil, false
	}
	return o.LeftQuota, true
}

// HasLeftQuota returns a boolean if a field has been set.
func (o *GetRwusdQuotaDetailsResponseSubscriptionQuota) HasLeftQuota() bool {
	if o != nil && !common.IsNil(o.LeftQuota) {
		return true
	}

	return false
}

// SetLeftQuota gets a reference to the given string and assigns it to the LeftQuota field.
func (o *GetRwusdQuotaDetailsResponseSubscriptionQuota) SetLeftQuota(v string) {
	o.LeftQuota = &v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *GetRwusdQuotaDetailsResponseSubscriptionQuota) GetMinimum() string {
	if o == nil || common.IsNil(o.Minimum) {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdQuotaDetailsResponseSubscriptionQuota) GetMinimumOk() (*string, bool) {
	if o == nil || common.IsNil(o.Minimum) {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *GetRwusdQuotaDetailsResponseSubscriptionQuota) HasMinimum() bool {
	if o != nil && !common.IsNil(o.Minimum) {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *GetRwusdQuotaDetailsResponseSubscriptionQuota) SetMinimum(v string) {
	o.Minimum = &v
}

func (o GetRwusdQuotaDetailsResponseSubscriptionQuota) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRwusdQuotaDetailsResponseSubscriptionQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !common.IsNil(o.LeftQuota) {
		toSerialize["leftQuota"] = o.LeftQuota
	}
	if !common.IsNil(o.Minimum) {
		toSerialize["minimum"] = o.Minimum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetRwusdQuotaDetailsResponseSubscriptionQuota) UnmarshalJSON(data []byte) (err error) {
	varGetRwusdQuotaDetailsResponseSubscriptionQuota := _GetRwusdQuotaDetailsResponseSubscriptionQuota{}

	err = json.Unmarshal(data, &varGetRwusdQuotaDetailsResponseSubscriptionQuota)

	if err != nil {
		return err
	}

	*o = GetRwusdQuotaDetailsResponseSubscriptionQuota(varGetRwusdQuotaDetailsResponseSubscriptionQuota)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assets")
		delete(additionalProperties, "leftQuota")
		delete(additionalProperties, "minimum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRwusdQuotaDetailsResponseSubscriptionQuota struct {
	value *GetRwusdQuotaDetailsResponseSubscriptionQuota
	isSet bool
}

func (v NullableGetRwusdQuotaDetailsResponseSubscriptionQuota) Get() *GetRwusdQuotaDetailsResponseSubscriptionQuota {
	return v.value
}

func (v *NullableGetRwusdQuotaDetailsResponseSubscriptionQuota) Set(val *GetRwusdQuotaDetailsResponseSubscriptionQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRwusdQuotaDetailsResponseSubscriptionQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRwusdQuotaDetailsResponseSubscriptionQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRwusdQuotaDetailsResponseSubscriptionQuota(val *GetRwusdQuotaDetailsResponseSubscriptionQuota) *NullableGetRwusdQuotaDetailsResponseSubscriptionQuota {
	return &NullableGetRwusdQuotaDetailsResponseSubscriptionQuota{value: val, isSet: true}
}

func (v NullableGetRwusdQuotaDetailsResponseSubscriptionQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRwusdQuotaDetailsResponseSubscriptionQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
