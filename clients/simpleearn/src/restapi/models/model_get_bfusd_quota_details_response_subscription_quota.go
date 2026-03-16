/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetBfusdQuotaDetailsResponseSubscriptionQuota type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBfusdQuotaDetailsResponseSubscriptionQuota{}

// GetBfusdQuotaDetailsResponseSubscriptionQuota struct for GetBfusdQuotaDetailsResponseSubscriptionQuota
type GetBfusdQuotaDetailsResponseSubscriptionQuota struct {
	LeftQuota            *string `json:"leftQuota,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBfusdQuotaDetailsResponseSubscriptionQuota GetBfusdQuotaDetailsResponseSubscriptionQuota

// NewGetBfusdQuotaDetailsResponseSubscriptionQuota instantiates a new GetBfusdQuotaDetailsResponseSubscriptionQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBfusdQuotaDetailsResponseSubscriptionQuota() *GetBfusdQuotaDetailsResponseSubscriptionQuota {
	this := GetBfusdQuotaDetailsResponseSubscriptionQuota{}
	return &this
}

// NewGetBfusdQuotaDetailsResponseSubscriptionQuotaWithDefaults instantiates a new GetBfusdQuotaDetailsResponseSubscriptionQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBfusdQuotaDetailsResponseSubscriptionQuotaWithDefaults() *GetBfusdQuotaDetailsResponseSubscriptionQuota {
	this := GetBfusdQuotaDetailsResponseSubscriptionQuota{}
	return &this
}

// GetLeftQuota returns the LeftQuota field value if set, zero value otherwise.
func (o *GetBfusdQuotaDetailsResponseSubscriptionQuota) GetLeftQuota() string {
	if o == nil || common.IsNil(o.LeftQuota) {
		var ret string
		return ret
	}
	return *o.LeftQuota
}

// GetLeftQuotaOk returns a tuple with the LeftQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdQuotaDetailsResponseSubscriptionQuota) GetLeftQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.LeftQuota) {
		return nil, false
	}
	return o.LeftQuota, true
}

// HasLeftQuota returns a boolean if a field has been set.
func (o *GetBfusdQuotaDetailsResponseSubscriptionQuota) HasLeftQuota() bool {
	if o != nil && !common.IsNil(o.LeftQuota) {
		return true
	}

	return false
}

// SetLeftQuota gets a reference to the given string and assigns it to the LeftQuota field.
func (o *GetBfusdQuotaDetailsResponseSubscriptionQuota) SetLeftQuota(v string) {
	o.LeftQuota = &v
}

func (o GetBfusdQuotaDetailsResponseSubscriptionQuota) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBfusdQuotaDetailsResponseSubscriptionQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LeftQuota) {
		toSerialize["leftQuota"] = o.LeftQuota
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetBfusdQuotaDetailsResponseSubscriptionQuota) UnmarshalJSON(data []byte) (err error) {
	varGetBfusdQuotaDetailsResponseSubscriptionQuota := _GetBfusdQuotaDetailsResponseSubscriptionQuota{}

	err = json.Unmarshal(data, &varGetBfusdQuotaDetailsResponseSubscriptionQuota)

	if err != nil {
		return err
	}

	*o = GetBfusdQuotaDetailsResponseSubscriptionQuota(varGetBfusdQuotaDetailsResponseSubscriptionQuota)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "leftQuota")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBfusdQuotaDetailsResponseSubscriptionQuota struct {
	value *GetBfusdQuotaDetailsResponseSubscriptionQuota
	isSet bool
}

func (v NullableGetBfusdQuotaDetailsResponseSubscriptionQuota) Get() *GetBfusdQuotaDetailsResponseSubscriptionQuota {
	return v.value
}

func (v *NullableGetBfusdQuotaDetailsResponseSubscriptionQuota) Set(val *GetBfusdQuotaDetailsResponseSubscriptionQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBfusdQuotaDetailsResponseSubscriptionQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBfusdQuotaDetailsResponseSubscriptionQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBfusdQuotaDetailsResponseSubscriptionQuota(val *GetBfusdQuotaDetailsResponseSubscriptionQuota) *NullableGetBfusdQuotaDetailsResponseSubscriptionQuota {
	return &NullableGetBfusdQuotaDetailsResponseSubscriptionQuota{value: val, isSet: true}
}

func (v NullableGetBfusdQuotaDetailsResponseSubscriptionQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBfusdQuotaDetailsResponseSubscriptionQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
