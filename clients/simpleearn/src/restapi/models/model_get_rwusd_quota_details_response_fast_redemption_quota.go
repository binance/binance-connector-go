/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetRwusdQuotaDetailsResponseFastRedemptionQuota type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetRwusdQuotaDetailsResponseFastRedemptionQuota{}

// GetRwusdQuotaDetailsResponseFastRedemptionQuota struct for GetRwusdQuotaDetailsResponseFastRedemptionQuota
type GetRwusdQuotaDetailsResponseFastRedemptionQuota struct {
	LeftQuota            *string `json:"leftQuota,omitempty"`
	Minimum              *string `json:"minimum,omitempty"`
	Fee                  *string `json:"fee,omitempty"`
	FreeQuota            *string `json:"freeQuota,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRwusdQuotaDetailsResponseFastRedemptionQuota GetRwusdQuotaDetailsResponseFastRedemptionQuota

// NewGetRwusdQuotaDetailsResponseFastRedemptionQuota instantiates a new GetRwusdQuotaDetailsResponseFastRedemptionQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRwusdQuotaDetailsResponseFastRedemptionQuota() *GetRwusdQuotaDetailsResponseFastRedemptionQuota {
	this := GetRwusdQuotaDetailsResponseFastRedemptionQuota{}
	return &this
}

// NewGetRwusdQuotaDetailsResponseFastRedemptionQuotaWithDefaults instantiates a new GetRwusdQuotaDetailsResponseFastRedemptionQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRwusdQuotaDetailsResponseFastRedemptionQuotaWithDefaults() *GetRwusdQuotaDetailsResponseFastRedemptionQuota {
	this := GetRwusdQuotaDetailsResponseFastRedemptionQuota{}
	return &this
}

// GetLeftQuota returns the LeftQuota field value if set, zero value otherwise.
func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) GetLeftQuota() string {
	if o == nil || common.IsNil(o.LeftQuota) {
		var ret string
		return ret
	}
	return *o.LeftQuota
}

// GetLeftQuotaOk returns a tuple with the LeftQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) GetLeftQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.LeftQuota) {
		return nil, false
	}
	return o.LeftQuota, true
}

// HasLeftQuota returns a boolean if a field has been set.
func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) HasLeftQuota() bool {
	if o != nil && !common.IsNil(o.LeftQuota) {
		return true
	}

	return false
}

// SetLeftQuota gets a reference to the given string and assigns it to the LeftQuota field.
func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) SetLeftQuota(v string) {
	o.LeftQuota = &v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) GetMinimum() string {
	if o == nil || common.IsNil(o.Minimum) {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) GetMinimumOk() (*string, bool) {
	if o == nil || common.IsNil(o.Minimum) {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) HasMinimum() bool {
	if o != nil && !common.IsNil(o.Minimum) {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) SetMinimum(v string) {
	o.Minimum = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) GetFee() string {
	if o == nil || common.IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) GetFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) SetFee(v string) {
	o.Fee = &v
}

// GetFreeQuota returns the FreeQuota field value if set, zero value otherwise.
func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) GetFreeQuota() string {
	if o == nil || common.IsNil(o.FreeQuota) {
		var ret string
		return ret
	}
	return *o.FreeQuota
}

// GetFreeQuotaOk returns a tuple with the FreeQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) GetFreeQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.FreeQuota) {
		return nil, false
	}
	return o.FreeQuota, true
}

// HasFreeQuota returns a boolean if a field has been set.
func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) HasFreeQuota() bool {
	if o != nil && !common.IsNil(o.FreeQuota) {
		return true
	}

	return false
}

// SetFreeQuota gets a reference to the given string and assigns it to the FreeQuota field.
func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) SetFreeQuota(v string) {
	o.FreeQuota = &v
}

func (o GetRwusdQuotaDetailsResponseFastRedemptionQuota) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRwusdQuotaDetailsResponseFastRedemptionQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LeftQuota) {
		toSerialize["leftQuota"] = o.LeftQuota
	}
	if !common.IsNil(o.Minimum) {
		toSerialize["minimum"] = o.Minimum
	}
	if !common.IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !common.IsNil(o.FreeQuota) {
		toSerialize["freeQuota"] = o.FreeQuota
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetRwusdQuotaDetailsResponseFastRedemptionQuota) UnmarshalJSON(data []byte) (err error) {
	varGetRwusdQuotaDetailsResponseFastRedemptionQuota := _GetRwusdQuotaDetailsResponseFastRedemptionQuota{}

	err = json.Unmarshal(data, &varGetRwusdQuotaDetailsResponseFastRedemptionQuota)

	if err != nil {
		return err
	}

	*o = GetRwusdQuotaDetailsResponseFastRedemptionQuota(varGetRwusdQuotaDetailsResponseFastRedemptionQuota)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "leftQuota")
		delete(additionalProperties, "minimum")
		delete(additionalProperties, "fee")
		delete(additionalProperties, "freeQuota")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRwusdQuotaDetailsResponseFastRedemptionQuota struct {
	value *GetRwusdQuotaDetailsResponseFastRedemptionQuota
	isSet bool
}

func (v NullableGetRwusdQuotaDetailsResponseFastRedemptionQuota) Get() *GetRwusdQuotaDetailsResponseFastRedemptionQuota {
	return v.value
}

func (v *NullableGetRwusdQuotaDetailsResponseFastRedemptionQuota) Set(val *GetRwusdQuotaDetailsResponseFastRedemptionQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRwusdQuotaDetailsResponseFastRedemptionQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRwusdQuotaDetailsResponseFastRedemptionQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRwusdQuotaDetailsResponseFastRedemptionQuota(val *GetRwusdQuotaDetailsResponseFastRedemptionQuota) *NullableGetRwusdQuotaDetailsResponseFastRedemptionQuota {
	return &NullableGetRwusdQuotaDetailsResponseFastRedemptionQuota{value: val, isSet: true}
}

func (v NullableGetRwusdQuotaDetailsResponseFastRedemptionQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRwusdQuotaDetailsResponseFastRedemptionQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
