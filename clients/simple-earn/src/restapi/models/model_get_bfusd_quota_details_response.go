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

// checks if the GetBfusdQuotaDetailsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBfusdQuotaDetailsResponse{}

// GetBfusdQuotaDetailsResponse struct for GetBfusdQuotaDetailsResponse
type GetBfusdQuotaDetailsResponse struct {
	FastRedemptionQuota     *GetBfusdQuotaDetailsResponseFastRedemptionQuota     `json:"fastRedemptionQuota,omitempty"`
	StandardRedemptionQuota *GetBfusdQuotaDetailsResponseStandardRedemptionQuota `json:"standardRedemptionQuota,omitempty"`
	SubscribeEnable         *bool                                                `json:"subscribeEnable,omitempty"`
	RedeemEnable            *bool                                                `json:"redeemEnable,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _GetBfusdQuotaDetailsResponse GetBfusdQuotaDetailsResponse

// NewGetBfusdQuotaDetailsResponse instantiates a new GetBfusdQuotaDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBfusdQuotaDetailsResponse() *GetBfusdQuotaDetailsResponse {
	this := GetBfusdQuotaDetailsResponse{}
	return &this
}

// NewGetBfusdQuotaDetailsResponseWithDefaults instantiates a new GetBfusdQuotaDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBfusdQuotaDetailsResponseWithDefaults() *GetBfusdQuotaDetailsResponse {
	this := GetBfusdQuotaDetailsResponse{}
	return &this
}

// GetFastRedemptionQuota returns the FastRedemptionQuota field value if set, zero value otherwise.
func (o *GetBfusdQuotaDetailsResponse) GetFastRedemptionQuota() GetBfusdQuotaDetailsResponseFastRedemptionQuota {
	if o == nil || common.IsNil(o.FastRedemptionQuota) {
		var ret GetBfusdQuotaDetailsResponseFastRedemptionQuota
		return ret
	}
	return *o.FastRedemptionQuota
}

// GetFastRedemptionQuotaOk returns a tuple with the FastRedemptionQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdQuotaDetailsResponse) GetFastRedemptionQuotaOk() (*GetBfusdQuotaDetailsResponseFastRedemptionQuota, bool) {
	if o == nil || common.IsNil(o.FastRedemptionQuota) {
		return nil, false
	}
	return o.FastRedemptionQuota, true
}

// HasFastRedemptionQuota returns a boolean if a field has been set.
func (o *GetBfusdQuotaDetailsResponse) HasFastRedemptionQuota() bool {
	if o != nil && !common.IsNil(o.FastRedemptionQuota) {
		return true
	}

	return false
}

// SetFastRedemptionQuota gets a reference to the given GetBfusdQuotaDetailsResponseFastRedemptionQuota and assigns it to the FastRedemptionQuota field.
func (o *GetBfusdQuotaDetailsResponse) SetFastRedemptionQuota(v GetBfusdQuotaDetailsResponseFastRedemptionQuota) {
	o.FastRedemptionQuota = &v
}

// GetStandardRedemptionQuota returns the StandardRedemptionQuota field value if set, zero value otherwise.
func (o *GetBfusdQuotaDetailsResponse) GetStandardRedemptionQuota() GetBfusdQuotaDetailsResponseStandardRedemptionQuota {
	if o == nil || common.IsNil(o.StandardRedemptionQuota) {
		var ret GetBfusdQuotaDetailsResponseStandardRedemptionQuota
		return ret
	}
	return *o.StandardRedemptionQuota
}

// GetStandardRedemptionQuotaOk returns a tuple with the StandardRedemptionQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdQuotaDetailsResponse) GetStandardRedemptionQuotaOk() (*GetBfusdQuotaDetailsResponseStandardRedemptionQuota, bool) {
	if o == nil || common.IsNil(o.StandardRedemptionQuota) {
		return nil, false
	}
	return o.StandardRedemptionQuota, true
}

// HasStandardRedemptionQuota returns a boolean if a field has been set.
func (o *GetBfusdQuotaDetailsResponse) HasStandardRedemptionQuota() bool {
	if o != nil && !common.IsNil(o.StandardRedemptionQuota) {
		return true
	}

	return false
}

// SetStandardRedemptionQuota gets a reference to the given GetBfusdQuotaDetailsResponseStandardRedemptionQuota and assigns it to the StandardRedemptionQuota field.
func (o *GetBfusdQuotaDetailsResponse) SetStandardRedemptionQuota(v GetBfusdQuotaDetailsResponseStandardRedemptionQuota) {
	o.StandardRedemptionQuota = &v
}

// GetSubscribeEnable returns the SubscribeEnable field value if set, zero value otherwise.
func (o *GetBfusdQuotaDetailsResponse) GetSubscribeEnable() bool {
	if o == nil || common.IsNil(o.SubscribeEnable) {
		var ret bool
		return ret
	}
	return *o.SubscribeEnable
}

// GetSubscribeEnableOk returns a tuple with the SubscribeEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdQuotaDetailsResponse) GetSubscribeEnableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.SubscribeEnable) {
		return nil, false
	}
	return o.SubscribeEnable, true
}

// HasSubscribeEnable returns a boolean if a field has been set.
func (o *GetBfusdQuotaDetailsResponse) HasSubscribeEnable() bool {
	if o != nil && !common.IsNil(o.SubscribeEnable) {
		return true
	}

	return false
}

// SetSubscribeEnable gets a reference to the given bool and assigns it to the SubscribeEnable field.
func (o *GetBfusdQuotaDetailsResponse) SetSubscribeEnable(v bool) {
	o.SubscribeEnable = &v
}

// GetRedeemEnable returns the RedeemEnable field value if set, zero value otherwise.
func (o *GetBfusdQuotaDetailsResponse) GetRedeemEnable() bool {
	if o == nil || common.IsNil(o.RedeemEnable) {
		var ret bool
		return ret
	}
	return *o.RedeemEnable
}

// GetRedeemEnableOk returns a tuple with the RedeemEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdQuotaDetailsResponse) GetRedeemEnableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.RedeemEnable) {
		return nil, false
	}
	return o.RedeemEnable, true
}

// HasRedeemEnable returns a boolean if a field has been set.
func (o *GetBfusdQuotaDetailsResponse) HasRedeemEnable() bool {
	if o != nil && !common.IsNil(o.RedeemEnable) {
		return true
	}

	return false
}

// SetRedeemEnable gets a reference to the given bool and assigns it to the RedeemEnable field.
func (o *GetBfusdQuotaDetailsResponse) SetRedeemEnable(v bool) {
	o.RedeemEnable = &v
}

func (o GetBfusdQuotaDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBfusdQuotaDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FastRedemptionQuota) {
		toSerialize["fastRedemptionQuota"] = o.FastRedemptionQuota
	}
	if !common.IsNil(o.StandardRedemptionQuota) {
		toSerialize["standardRedemptionQuota"] = o.StandardRedemptionQuota
	}
	if !common.IsNil(o.SubscribeEnable) {
		toSerialize["subscribeEnable"] = o.SubscribeEnable
	}
	if !common.IsNil(o.RedeemEnable) {
		toSerialize["redeemEnable"] = o.RedeemEnable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetBfusdQuotaDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	varGetBfusdQuotaDetailsResponse := _GetBfusdQuotaDetailsResponse{}

	err = json.Unmarshal(data, &varGetBfusdQuotaDetailsResponse)

	if err != nil {
		return err
	}

	*o = GetBfusdQuotaDetailsResponse(varGetBfusdQuotaDetailsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fastRedemptionQuota")
		delete(additionalProperties, "standardRedemptionQuota")
		delete(additionalProperties, "subscribeEnable")
		delete(additionalProperties, "redeemEnable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBfusdQuotaDetailsResponse struct {
	value *GetBfusdQuotaDetailsResponse
	isSet bool
}

func (v NullableGetBfusdQuotaDetailsResponse) Get() *GetBfusdQuotaDetailsResponse {
	return v.value
}

func (v *NullableGetBfusdQuotaDetailsResponse) Set(val *GetBfusdQuotaDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBfusdQuotaDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBfusdQuotaDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBfusdQuotaDetailsResponse(val *GetBfusdQuotaDetailsResponse) *NullableGetBfusdQuotaDetailsResponse {
	return &NullableGetBfusdQuotaDetailsResponse{value: val, isSet: true}
}

func (v NullableGetBfusdQuotaDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBfusdQuotaDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
