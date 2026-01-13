/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetRwusdQuotaDetailsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetRwusdQuotaDetailsResponse{}

// GetRwusdQuotaDetailsResponse struct for GetRwusdQuotaDetailsResponse
type GetRwusdQuotaDetailsResponse struct {
	SubscriptionQuota       *GetRwusdQuotaDetailsResponseSubscriptionQuota       `json:"subscriptionQuota,omitempty"`
	FastRedemptionQuota     *GetRwusdQuotaDetailsResponseFastRedemptionQuota     `json:"fastRedemptionQuota,omitempty"`
	StandardRedemptionQuota *GetRwusdQuotaDetailsResponseStandardRedemptionQuota `json:"standardRedemptionQuota,omitempty"`
	SubscribeEnable         *bool                                                `json:"subscribeEnable,omitempty"`
	RedeemEnable            *bool                                                `json:"redeemEnable,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _GetRwusdQuotaDetailsResponse GetRwusdQuotaDetailsResponse

// NewGetRwusdQuotaDetailsResponse instantiates a new GetRwusdQuotaDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRwusdQuotaDetailsResponse() *GetRwusdQuotaDetailsResponse {
	this := GetRwusdQuotaDetailsResponse{}
	return &this
}

// NewGetRwusdQuotaDetailsResponseWithDefaults instantiates a new GetRwusdQuotaDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRwusdQuotaDetailsResponseWithDefaults() *GetRwusdQuotaDetailsResponse {
	this := GetRwusdQuotaDetailsResponse{}
	return &this
}

// GetSubscriptionQuota returns the SubscriptionQuota field value if set, zero value otherwise.
func (o *GetRwusdQuotaDetailsResponse) GetSubscriptionQuota() GetRwusdQuotaDetailsResponseSubscriptionQuota {
	if o == nil || common.IsNil(o.SubscriptionQuota) {
		var ret GetRwusdQuotaDetailsResponseSubscriptionQuota
		return ret
	}
	return *o.SubscriptionQuota
}

// GetSubscriptionQuotaOk returns a tuple with the SubscriptionQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdQuotaDetailsResponse) GetSubscriptionQuotaOk() (*GetRwusdQuotaDetailsResponseSubscriptionQuota, bool) {
	if o == nil || common.IsNil(o.SubscriptionQuota) {
		return nil, false
	}
	return o.SubscriptionQuota, true
}

// HasSubscriptionQuota returns a boolean if a field has been set.
func (o *GetRwusdQuotaDetailsResponse) HasSubscriptionQuota() bool {
	if o != nil && !common.IsNil(o.SubscriptionQuota) {
		return true
	}

	return false
}

// SetSubscriptionQuota gets a reference to the given GetRwusdQuotaDetailsResponseSubscriptionQuota and assigns it to the SubscriptionQuota field.
func (o *GetRwusdQuotaDetailsResponse) SetSubscriptionQuota(v GetRwusdQuotaDetailsResponseSubscriptionQuota) {
	o.SubscriptionQuota = &v
}

// GetFastRedemptionQuota returns the FastRedemptionQuota field value if set, zero value otherwise.
func (o *GetRwusdQuotaDetailsResponse) GetFastRedemptionQuota() GetRwusdQuotaDetailsResponseFastRedemptionQuota {
	if o == nil || common.IsNil(o.FastRedemptionQuota) {
		var ret GetRwusdQuotaDetailsResponseFastRedemptionQuota
		return ret
	}
	return *o.FastRedemptionQuota
}

// GetFastRedemptionQuotaOk returns a tuple with the FastRedemptionQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdQuotaDetailsResponse) GetFastRedemptionQuotaOk() (*GetRwusdQuotaDetailsResponseFastRedemptionQuota, bool) {
	if o == nil || common.IsNil(o.FastRedemptionQuota) {
		return nil, false
	}
	return o.FastRedemptionQuota, true
}

// HasFastRedemptionQuota returns a boolean if a field has been set.
func (o *GetRwusdQuotaDetailsResponse) HasFastRedemptionQuota() bool {
	if o != nil && !common.IsNil(o.FastRedemptionQuota) {
		return true
	}

	return false
}

// SetFastRedemptionQuota gets a reference to the given GetRwusdQuotaDetailsResponseFastRedemptionQuota and assigns it to the FastRedemptionQuota field.
func (o *GetRwusdQuotaDetailsResponse) SetFastRedemptionQuota(v GetRwusdQuotaDetailsResponseFastRedemptionQuota) {
	o.FastRedemptionQuota = &v
}

// GetStandardRedemptionQuota returns the StandardRedemptionQuota field value if set, zero value otherwise.
func (o *GetRwusdQuotaDetailsResponse) GetStandardRedemptionQuota() GetRwusdQuotaDetailsResponseStandardRedemptionQuota {
	if o == nil || common.IsNil(o.StandardRedemptionQuota) {
		var ret GetRwusdQuotaDetailsResponseStandardRedemptionQuota
		return ret
	}
	return *o.StandardRedemptionQuota
}

// GetStandardRedemptionQuotaOk returns a tuple with the StandardRedemptionQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdQuotaDetailsResponse) GetStandardRedemptionQuotaOk() (*GetRwusdQuotaDetailsResponseStandardRedemptionQuota, bool) {
	if o == nil || common.IsNil(o.StandardRedemptionQuota) {
		return nil, false
	}
	return o.StandardRedemptionQuota, true
}

// HasStandardRedemptionQuota returns a boolean if a field has been set.
func (o *GetRwusdQuotaDetailsResponse) HasStandardRedemptionQuota() bool {
	if o != nil && !common.IsNil(o.StandardRedemptionQuota) {
		return true
	}

	return false
}

// SetStandardRedemptionQuota gets a reference to the given GetRwusdQuotaDetailsResponseStandardRedemptionQuota and assigns it to the StandardRedemptionQuota field.
func (o *GetRwusdQuotaDetailsResponse) SetStandardRedemptionQuota(v GetRwusdQuotaDetailsResponseStandardRedemptionQuota) {
	o.StandardRedemptionQuota = &v
}

// GetSubscribeEnable returns the SubscribeEnable field value if set, zero value otherwise.
func (o *GetRwusdQuotaDetailsResponse) GetSubscribeEnable() bool {
	if o == nil || common.IsNil(o.SubscribeEnable) {
		var ret bool
		return ret
	}
	return *o.SubscribeEnable
}

// GetSubscribeEnableOk returns a tuple with the SubscribeEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdQuotaDetailsResponse) GetSubscribeEnableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.SubscribeEnable) {
		return nil, false
	}
	return o.SubscribeEnable, true
}

// HasSubscribeEnable returns a boolean if a field has been set.
func (o *GetRwusdQuotaDetailsResponse) HasSubscribeEnable() bool {
	if o != nil && !common.IsNil(o.SubscribeEnable) {
		return true
	}

	return false
}

// SetSubscribeEnable gets a reference to the given bool and assigns it to the SubscribeEnable field.
func (o *GetRwusdQuotaDetailsResponse) SetSubscribeEnable(v bool) {
	o.SubscribeEnable = &v
}

// GetRedeemEnable returns the RedeemEnable field value if set, zero value otherwise.
func (o *GetRwusdQuotaDetailsResponse) GetRedeemEnable() bool {
	if o == nil || common.IsNil(o.RedeemEnable) {
		var ret bool
		return ret
	}
	return *o.RedeemEnable
}

// GetRedeemEnableOk returns a tuple with the RedeemEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdQuotaDetailsResponse) GetRedeemEnableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.RedeemEnable) {
		return nil, false
	}
	return o.RedeemEnable, true
}

// HasRedeemEnable returns a boolean if a field has been set.
func (o *GetRwusdQuotaDetailsResponse) HasRedeemEnable() bool {
	if o != nil && !common.IsNil(o.RedeemEnable) {
		return true
	}

	return false
}

// SetRedeemEnable gets a reference to the given bool and assigns it to the RedeemEnable field.
func (o *GetRwusdQuotaDetailsResponse) SetRedeemEnable(v bool) {
	o.RedeemEnable = &v
}

func (o GetRwusdQuotaDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRwusdQuotaDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.SubscriptionQuota) {
		toSerialize["subscriptionQuota"] = o.SubscriptionQuota
	}
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

func (o *GetRwusdQuotaDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	varGetRwusdQuotaDetailsResponse := _GetRwusdQuotaDetailsResponse{}

	err = json.Unmarshal(data, &varGetRwusdQuotaDetailsResponse)

	if err != nil {
		return err
	}

	*o = GetRwusdQuotaDetailsResponse(varGetRwusdQuotaDetailsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "subscriptionQuota")
		delete(additionalProperties, "fastRedemptionQuota")
		delete(additionalProperties, "standardRedemptionQuota")
		delete(additionalProperties, "subscribeEnable")
		delete(additionalProperties, "redeemEnable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRwusdQuotaDetailsResponse struct {
	value *GetRwusdQuotaDetailsResponse
	isSet bool
}

func (v NullableGetRwusdQuotaDetailsResponse) Get() *GetRwusdQuotaDetailsResponse {
	return v.value
}

func (v *NullableGetRwusdQuotaDetailsResponse) Set(val *GetRwusdQuotaDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRwusdQuotaDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRwusdQuotaDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRwusdQuotaDetailsResponse(val *GetRwusdQuotaDetailsResponse) *NullableGetRwusdQuotaDetailsResponse {
	return &NullableGetRwusdQuotaDetailsResponse{value: val, isSet: true}
}

func (v NullableGetRwusdQuotaDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRwusdQuotaDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
