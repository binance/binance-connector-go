/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetBfusdQuotaDetailsResponseStandardRedemptionQuota type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBfusdQuotaDetailsResponseStandardRedemptionQuota{}

// GetBfusdQuotaDetailsResponseStandardRedemptionQuota struct for GetBfusdQuotaDetailsResponseStandardRedemptionQuota
type GetBfusdQuotaDetailsResponseStandardRedemptionQuota struct {
	LeftQuota            *string `json:"leftQuota,omitempty"`
	Minimum              *string `json:"minimum,omitempty"`
	Fee                  *string `json:"fee,omitempty"`
	RedeemPeriod         *int64  `json:"redeemPeriod,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBfusdQuotaDetailsResponseStandardRedemptionQuota GetBfusdQuotaDetailsResponseStandardRedemptionQuota

// NewGetBfusdQuotaDetailsResponseStandardRedemptionQuota instantiates a new GetBfusdQuotaDetailsResponseStandardRedemptionQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBfusdQuotaDetailsResponseStandardRedemptionQuota() *GetBfusdQuotaDetailsResponseStandardRedemptionQuota {
	this := GetBfusdQuotaDetailsResponseStandardRedemptionQuota{}
	return &this
}

// NewGetBfusdQuotaDetailsResponseStandardRedemptionQuotaWithDefaults instantiates a new GetBfusdQuotaDetailsResponseStandardRedemptionQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBfusdQuotaDetailsResponseStandardRedemptionQuotaWithDefaults() *GetBfusdQuotaDetailsResponseStandardRedemptionQuota {
	this := GetBfusdQuotaDetailsResponseStandardRedemptionQuota{}
	return &this
}

// GetLeftQuota returns the LeftQuota field value if set, zero value otherwise.
func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) GetLeftQuota() string {
	if o == nil || common.IsNil(o.LeftQuota) {
		var ret string
		return ret
	}
	return *o.LeftQuota
}

// GetLeftQuotaOk returns a tuple with the LeftQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) GetLeftQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.LeftQuota) {
		return nil, false
	}
	return o.LeftQuota, true
}

// HasLeftQuota returns a boolean if a field has been set.
func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) HasLeftQuota() bool {
	if o != nil && !common.IsNil(o.LeftQuota) {
		return true
	}

	return false
}

// SetLeftQuota gets a reference to the given string and assigns it to the LeftQuota field.
func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) SetLeftQuota(v string) {
	o.LeftQuota = &v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) GetMinimum() string {
	if o == nil || common.IsNil(o.Minimum) {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) GetMinimumOk() (*string, bool) {
	if o == nil || common.IsNil(o.Minimum) {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) HasMinimum() bool {
	if o != nil && !common.IsNil(o.Minimum) {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) SetMinimum(v string) {
	o.Minimum = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) GetFee() string {
	if o == nil || common.IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) GetFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) SetFee(v string) {
	o.Fee = &v
}

// GetRedeemPeriod returns the RedeemPeriod field value if set, zero value otherwise.
func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) GetRedeemPeriod() int64 {
	if o == nil || common.IsNil(o.RedeemPeriod) {
		var ret int64
		return ret
	}
	return *o.RedeemPeriod
}

// GetRedeemPeriodOk returns a tuple with the RedeemPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) GetRedeemPeriodOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RedeemPeriod) {
		return nil, false
	}
	return o.RedeemPeriod, true
}

// HasRedeemPeriod returns a boolean if a field has been set.
func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) HasRedeemPeriod() bool {
	if o != nil && !common.IsNil(o.RedeemPeriod) {
		return true
	}

	return false
}

// SetRedeemPeriod gets a reference to the given int64 and assigns it to the RedeemPeriod field.
func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) SetRedeemPeriod(v int64) {
	o.RedeemPeriod = &v
}

func (o GetBfusdQuotaDetailsResponseStandardRedemptionQuota) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBfusdQuotaDetailsResponseStandardRedemptionQuota) ToMap() (map[string]interface{}, error) {
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

func (o *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) UnmarshalJSON(data []byte) (err error) {
	varGetBfusdQuotaDetailsResponseStandardRedemptionQuota := _GetBfusdQuotaDetailsResponseStandardRedemptionQuota{}

	err = json.Unmarshal(data, &varGetBfusdQuotaDetailsResponseStandardRedemptionQuota)

	if err != nil {
		return err
	}

	*o = GetBfusdQuotaDetailsResponseStandardRedemptionQuota(varGetBfusdQuotaDetailsResponseStandardRedemptionQuota)

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

type NullableGetBfusdQuotaDetailsResponseStandardRedemptionQuota struct {
	value *GetBfusdQuotaDetailsResponseStandardRedemptionQuota
	isSet bool
}

func (v NullableGetBfusdQuotaDetailsResponseStandardRedemptionQuota) Get() *GetBfusdQuotaDetailsResponseStandardRedemptionQuota {
	return v.value
}

func (v *NullableGetBfusdQuotaDetailsResponseStandardRedemptionQuota) Set(val *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBfusdQuotaDetailsResponseStandardRedemptionQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBfusdQuotaDetailsResponseStandardRedemptionQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBfusdQuotaDetailsResponseStandardRedemptionQuota(val *GetBfusdQuotaDetailsResponseStandardRedemptionQuota) *NullableGetBfusdQuotaDetailsResponseStandardRedemptionQuota {
	return &NullableGetBfusdQuotaDetailsResponseStandardRedemptionQuota{value: val, isSet: true}
}

func (v NullableGetBfusdQuotaDetailsResponseStandardRedemptionQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBfusdQuotaDetailsResponseStandardRedemptionQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
