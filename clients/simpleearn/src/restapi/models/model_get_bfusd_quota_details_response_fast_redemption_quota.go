/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetBfusdQuotaDetailsResponseFastRedemptionQuota type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBfusdQuotaDetailsResponseFastRedemptionQuota{}

// GetBfusdQuotaDetailsResponseFastRedemptionQuota struct for GetBfusdQuotaDetailsResponseFastRedemptionQuota
type GetBfusdQuotaDetailsResponseFastRedemptionQuota struct {
	LeftQuota            *string `json:"leftQuota,omitempty"`
	Minimum              *string `json:"minimum,omitempty"`
	Fee                  *string `json:"fee,omitempty"`
	FreeQuota            *string `json:"freeQuota,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBfusdQuotaDetailsResponseFastRedemptionQuota GetBfusdQuotaDetailsResponseFastRedemptionQuota

// NewGetBfusdQuotaDetailsResponseFastRedemptionQuota instantiates a new GetBfusdQuotaDetailsResponseFastRedemptionQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBfusdQuotaDetailsResponseFastRedemptionQuota() *GetBfusdQuotaDetailsResponseFastRedemptionQuota {
	this := GetBfusdQuotaDetailsResponseFastRedemptionQuota{}
	return &this
}

// NewGetBfusdQuotaDetailsResponseFastRedemptionQuotaWithDefaults instantiates a new GetBfusdQuotaDetailsResponseFastRedemptionQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBfusdQuotaDetailsResponseFastRedemptionQuotaWithDefaults() *GetBfusdQuotaDetailsResponseFastRedemptionQuota {
	this := GetBfusdQuotaDetailsResponseFastRedemptionQuota{}
	return &this
}

// GetLeftQuota returns the LeftQuota field value if set, zero value otherwise.
func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) GetLeftQuota() string {
	if o == nil || common.IsNil(o.LeftQuota) {
		var ret string
		return ret
	}
	return *o.LeftQuota
}

// GetLeftQuotaOk returns a tuple with the LeftQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) GetLeftQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.LeftQuota) {
		return nil, false
	}
	return o.LeftQuota, true
}

// HasLeftQuota returns a boolean if a field has been set.
func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) HasLeftQuota() bool {
	if o != nil && !common.IsNil(o.LeftQuota) {
		return true
	}

	return false
}

// SetLeftQuota gets a reference to the given string and assigns it to the LeftQuota field.
func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) SetLeftQuota(v string) {
	o.LeftQuota = &v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) GetMinimum() string {
	if o == nil || common.IsNil(o.Minimum) {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) GetMinimumOk() (*string, bool) {
	if o == nil || common.IsNil(o.Minimum) {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) HasMinimum() bool {
	if o != nil && !common.IsNil(o.Minimum) {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) SetMinimum(v string) {
	o.Minimum = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) GetFee() string {
	if o == nil || common.IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) GetFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) SetFee(v string) {
	o.Fee = &v
}

// GetFreeQuota returns the FreeQuota field value if set, zero value otherwise.
func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) GetFreeQuota() string {
	if o == nil || common.IsNil(o.FreeQuota) {
		var ret string
		return ret
	}
	return *o.FreeQuota
}

// GetFreeQuotaOk returns a tuple with the FreeQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) GetFreeQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.FreeQuota) {
		return nil, false
	}
	return o.FreeQuota, true
}

// HasFreeQuota returns a boolean if a field has been set.
func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) HasFreeQuota() bool {
	if o != nil && !common.IsNil(o.FreeQuota) {
		return true
	}

	return false
}

// SetFreeQuota gets a reference to the given string and assigns it to the FreeQuota field.
func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) SetFreeQuota(v string) {
	o.FreeQuota = &v
}

func (o GetBfusdQuotaDetailsResponseFastRedemptionQuota) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBfusdQuotaDetailsResponseFastRedemptionQuota) ToMap() (map[string]interface{}, error) {
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

func (o *GetBfusdQuotaDetailsResponseFastRedemptionQuota) UnmarshalJSON(data []byte) (err error) {
	varGetBfusdQuotaDetailsResponseFastRedemptionQuota := _GetBfusdQuotaDetailsResponseFastRedemptionQuota{}

	err = json.Unmarshal(data, &varGetBfusdQuotaDetailsResponseFastRedemptionQuota)

	if err != nil {
		return err
	}

	*o = GetBfusdQuotaDetailsResponseFastRedemptionQuota(varGetBfusdQuotaDetailsResponseFastRedemptionQuota)

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

type NullableGetBfusdQuotaDetailsResponseFastRedemptionQuota struct {
	value *GetBfusdQuotaDetailsResponseFastRedemptionQuota
	isSet bool
}

func (v NullableGetBfusdQuotaDetailsResponseFastRedemptionQuota) Get() *GetBfusdQuotaDetailsResponseFastRedemptionQuota {
	return v.value
}

func (v *NullableGetBfusdQuotaDetailsResponseFastRedemptionQuota) Set(val *GetBfusdQuotaDetailsResponseFastRedemptionQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBfusdQuotaDetailsResponseFastRedemptionQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBfusdQuotaDetailsResponseFastRedemptionQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBfusdQuotaDetailsResponseFastRedemptionQuota(val *GetBfusdQuotaDetailsResponseFastRedemptionQuota) *NullableGetBfusdQuotaDetailsResponseFastRedemptionQuota {
	return &NullableGetBfusdQuotaDetailsResponseFastRedemptionQuota{value: val, isSet: true}
}

func (v NullableGetBfusdQuotaDetailsResponseFastRedemptionQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBfusdQuotaDetailsResponseFastRedemptionQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
