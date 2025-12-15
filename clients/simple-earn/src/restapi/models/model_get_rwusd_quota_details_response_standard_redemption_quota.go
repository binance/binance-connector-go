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

// checks if the GetRwusdQuotaDetailsResponseStandardRedemptionQuota type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetRwusdQuotaDetailsResponseStandardRedemptionQuota{}

// GetRwusdQuotaDetailsResponseStandardRedemptionQuota struct for GetRwusdQuotaDetailsResponseStandardRedemptionQuota
type GetRwusdQuotaDetailsResponseStandardRedemptionQuota struct {
	LeftQuota            *string `json:"leftQuota,omitempty"`
	Minimum              *string `json:"minimum,omitempty"`
	Fee                  *string `json:"fee,omitempty"`
	RedeemPeriod         *int64  `json:"redeemPeriod,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRwusdQuotaDetailsResponseStandardRedemptionQuota GetRwusdQuotaDetailsResponseStandardRedemptionQuota

// NewGetRwusdQuotaDetailsResponseStandardRedemptionQuota instantiates a new GetRwusdQuotaDetailsResponseStandardRedemptionQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRwusdQuotaDetailsResponseStandardRedemptionQuota() *GetRwusdQuotaDetailsResponseStandardRedemptionQuota {
	this := GetRwusdQuotaDetailsResponseStandardRedemptionQuota{}
	return &this
}

// NewGetRwusdQuotaDetailsResponseStandardRedemptionQuotaWithDefaults instantiates a new GetRwusdQuotaDetailsResponseStandardRedemptionQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRwusdQuotaDetailsResponseStandardRedemptionQuotaWithDefaults() *GetRwusdQuotaDetailsResponseStandardRedemptionQuota {
	this := GetRwusdQuotaDetailsResponseStandardRedemptionQuota{}
	return &this
}

// GetLeftQuota returns the LeftQuota field value if set, zero value otherwise.
func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) GetLeftQuota() string {
	if o == nil || common.IsNil(o.LeftQuota) {
		var ret string
		return ret
	}
	return *o.LeftQuota
}

// GetLeftQuotaOk returns a tuple with the LeftQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) GetLeftQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.LeftQuota) {
		return nil, false
	}
	return o.LeftQuota, true
}

// HasLeftQuota returns a boolean if a field has been set.
func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) HasLeftQuota() bool {
	if o != nil && !common.IsNil(o.LeftQuota) {
		return true
	}

	return false
}

// SetLeftQuota gets a reference to the given string and assigns it to the LeftQuota field.
func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) SetLeftQuota(v string) {
	o.LeftQuota = &v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) GetMinimum() string {
	if o == nil || common.IsNil(o.Minimum) {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) GetMinimumOk() (*string, bool) {
	if o == nil || common.IsNil(o.Minimum) {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) HasMinimum() bool {
	if o != nil && !common.IsNil(o.Minimum) {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) SetMinimum(v string) {
	o.Minimum = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) GetFee() string {
	if o == nil || common.IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) GetFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) SetFee(v string) {
	o.Fee = &v
}

// GetRedeemPeriod returns the RedeemPeriod field value if set, zero value otherwise.
func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) GetRedeemPeriod() int64 {
	if o == nil || common.IsNil(o.RedeemPeriod) {
		var ret int64
		return ret
	}
	return *o.RedeemPeriod
}

// GetRedeemPeriodOk returns a tuple with the RedeemPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) GetRedeemPeriodOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RedeemPeriod) {
		return nil, false
	}
	return o.RedeemPeriod, true
}

// HasRedeemPeriod returns a boolean if a field has been set.
func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) HasRedeemPeriod() bool {
	if o != nil && !common.IsNil(o.RedeemPeriod) {
		return true
	}

	return false
}

// SetRedeemPeriod gets a reference to the given int64 and assigns it to the RedeemPeriod field.
func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) SetRedeemPeriod(v int64) {
	o.RedeemPeriod = &v
}

func (o GetRwusdQuotaDetailsResponseStandardRedemptionQuota) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRwusdQuotaDetailsResponseStandardRedemptionQuota) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.RedeemPeriod) {
		toSerialize["redeemPeriod"] = o.RedeemPeriod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) UnmarshalJSON(data []byte) (err error) {
	varGetRwusdQuotaDetailsResponseStandardRedemptionQuota := _GetRwusdQuotaDetailsResponseStandardRedemptionQuota{}

	err = json.Unmarshal(data, &varGetRwusdQuotaDetailsResponseStandardRedemptionQuota)

	if err != nil {
		return err
	}

	*o = GetRwusdQuotaDetailsResponseStandardRedemptionQuota(varGetRwusdQuotaDetailsResponseStandardRedemptionQuota)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "leftQuota")
		delete(additionalProperties, "minimum")
		delete(additionalProperties, "fee")
		delete(additionalProperties, "redeemPeriod")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRwusdQuotaDetailsResponseStandardRedemptionQuota struct {
	value *GetRwusdQuotaDetailsResponseStandardRedemptionQuota
	isSet bool
}

func (v NullableGetRwusdQuotaDetailsResponseStandardRedemptionQuota) Get() *GetRwusdQuotaDetailsResponseStandardRedemptionQuota {
	return v.value
}

func (v *NullableGetRwusdQuotaDetailsResponseStandardRedemptionQuota) Set(val *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRwusdQuotaDetailsResponseStandardRedemptionQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRwusdQuotaDetailsResponseStandardRedemptionQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRwusdQuotaDetailsResponseStandardRedemptionQuota(val *GetRwusdQuotaDetailsResponseStandardRedemptionQuota) *NullableGetRwusdQuotaDetailsResponseStandardRedemptionQuota {
	return &NullableGetRwusdQuotaDetailsResponseStandardRedemptionQuota{value: val, isSet: true}
}

func (v NullableGetRwusdQuotaDetailsResponseStandardRedemptionQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRwusdQuotaDetailsResponseStandardRedemptionQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
