/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountInfoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountInfoResponse{}

// AccountInfoResponse struct for AccountInfoResponse
type AccountInfoResponse struct {
	VipLevel                       *int64 `json:"vipLevel,omitempty"`
	IsMarginEnabled                *bool  `json:"isMarginEnabled,omitempty"`
	IsFutureEnabled                *bool  `json:"isFutureEnabled,omitempty"`
	IsOptionsEnabled               *bool  `json:"isOptionsEnabled,omitempty"`
	IsPortfolioMarginRetailEnabled *bool  `json:"isPortfolioMarginRetailEnabled,omitempty"`
	AdditionalProperties           map[string]interface{}
}

type _AccountInfoResponse AccountInfoResponse

// NewAccountInfoResponse instantiates a new AccountInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInfoResponse() *AccountInfoResponse {
	this := AccountInfoResponse{}
	return &this
}

// NewAccountInfoResponseWithDefaults instantiates a new AccountInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInfoResponseWithDefaults() *AccountInfoResponse {
	this := AccountInfoResponse{}
	return &this
}

// GetVipLevel returns the VipLevel field value if set, zero value otherwise.
func (o *AccountInfoResponse) GetVipLevel() int64 {
	if o == nil || common.IsNil(o.VipLevel) {
		var ret int64
		return ret
	}
	return *o.VipLevel
}

// GetVipLevelOk returns a tuple with the VipLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfoResponse) GetVipLevelOk() (*int64, bool) {
	if o == nil || common.IsNil(o.VipLevel) {
		return nil, false
	}
	return o.VipLevel, true
}

// HasVipLevel returns a boolean if a field has been set.
func (o *AccountInfoResponse) HasVipLevel() bool {
	if o != nil && !common.IsNil(o.VipLevel) {
		return true
	}

	return false
}

// SetVipLevel gets a reference to the given int64 and assigns it to the VipLevel field.
func (o *AccountInfoResponse) SetVipLevel(v int64) {
	o.VipLevel = &v
}

// GetIsMarginEnabled returns the IsMarginEnabled field value if set, zero value otherwise.
func (o *AccountInfoResponse) GetIsMarginEnabled() bool {
	if o == nil || common.IsNil(o.IsMarginEnabled) {
		var ret bool
		return ret
	}
	return *o.IsMarginEnabled
}

// GetIsMarginEnabledOk returns a tuple with the IsMarginEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfoResponse) GetIsMarginEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsMarginEnabled) {
		return nil, false
	}
	return o.IsMarginEnabled, true
}

// HasIsMarginEnabled returns a boolean if a field has been set.
func (o *AccountInfoResponse) HasIsMarginEnabled() bool {
	if o != nil && !common.IsNil(o.IsMarginEnabled) {
		return true
	}

	return false
}

// SetIsMarginEnabled gets a reference to the given bool and assigns it to the IsMarginEnabled field.
func (o *AccountInfoResponse) SetIsMarginEnabled(v bool) {
	o.IsMarginEnabled = &v
}

// GetIsFutureEnabled returns the IsFutureEnabled field value if set, zero value otherwise.
func (o *AccountInfoResponse) GetIsFutureEnabled() bool {
	if o == nil || common.IsNil(o.IsFutureEnabled) {
		var ret bool
		return ret
	}
	return *o.IsFutureEnabled
}

// GetIsFutureEnabledOk returns a tuple with the IsFutureEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfoResponse) GetIsFutureEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsFutureEnabled) {
		return nil, false
	}
	return o.IsFutureEnabled, true
}

// HasIsFutureEnabled returns a boolean if a field has been set.
func (o *AccountInfoResponse) HasIsFutureEnabled() bool {
	if o != nil && !common.IsNil(o.IsFutureEnabled) {
		return true
	}

	return false
}

// SetIsFutureEnabled gets a reference to the given bool and assigns it to the IsFutureEnabled field.
func (o *AccountInfoResponse) SetIsFutureEnabled(v bool) {
	o.IsFutureEnabled = &v
}

// GetIsOptionsEnabled returns the IsOptionsEnabled field value if set, zero value otherwise.
func (o *AccountInfoResponse) GetIsOptionsEnabled() bool {
	if o == nil || common.IsNil(o.IsOptionsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsOptionsEnabled
}

// GetIsOptionsEnabledOk returns a tuple with the IsOptionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfoResponse) GetIsOptionsEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsOptionsEnabled) {
		return nil, false
	}
	return o.IsOptionsEnabled, true
}

// HasIsOptionsEnabled returns a boolean if a field has been set.
func (o *AccountInfoResponse) HasIsOptionsEnabled() bool {
	if o != nil && !common.IsNil(o.IsOptionsEnabled) {
		return true
	}

	return false
}

// SetIsOptionsEnabled gets a reference to the given bool and assigns it to the IsOptionsEnabled field.
func (o *AccountInfoResponse) SetIsOptionsEnabled(v bool) {
	o.IsOptionsEnabled = &v
}

// GetIsPortfolioMarginRetailEnabled returns the IsPortfolioMarginRetailEnabled field value if set, zero value otherwise.
func (o *AccountInfoResponse) GetIsPortfolioMarginRetailEnabled() bool {
	if o == nil || common.IsNil(o.IsPortfolioMarginRetailEnabled) {
		var ret bool
		return ret
	}
	return *o.IsPortfolioMarginRetailEnabled
}

// GetIsPortfolioMarginRetailEnabledOk returns a tuple with the IsPortfolioMarginRetailEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfoResponse) GetIsPortfolioMarginRetailEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsPortfolioMarginRetailEnabled) {
		return nil, false
	}
	return o.IsPortfolioMarginRetailEnabled, true
}

// HasIsPortfolioMarginRetailEnabled returns a boolean if a field has been set.
func (o *AccountInfoResponse) HasIsPortfolioMarginRetailEnabled() bool {
	if o != nil && !common.IsNil(o.IsPortfolioMarginRetailEnabled) {
		return true
	}

	return false
}

// SetIsPortfolioMarginRetailEnabled gets a reference to the given bool and assigns it to the IsPortfolioMarginRetailEnabled field.
func (o *AccountInfoResponse) SetIsPortfolioMarginRetailEnabled(v bool) {
	o.IsPortfolioMarginRetailEnabled = &v
}

func (o AccountInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.VipLevel) {
		toSerialize["vipLevel"] = o.VipLevel
	}
	if !common.IsNil(o.IsMarginEnabled) {
		toSerialize["isMarginEnabled"] = o.IsMarginEnabled
	}
	if !common.IsNil(o.IsFutureEnabled) {
		toSerialize["isFutureEnabled"] = o.IsFutureEnabled
	}
	if !common.IsNil(o.IsOptionsEnabled) {
		toSerialize["isOptionsEnabled"] = o.IsOptionsEnabled
	}
	if !common.IsNil(o.IsPortfolioMarginRetailEnabled) {
		toSerialize["isPortfolioMarginRetailEnabled"] = o.IsPortfolioMarginRetailEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountInfoResponse) UnmarshalJSON(data []byte) (err error) {
	varAccountInfoResponse := _AccountInfoResponse{}

	err = json.Unmarshal(data, &varAccountInfoResponse)

	if err != nil {
		return err
	}

	*o = AccountInfoResponse(varAccountInfoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vipLevel")
		delete(additionalProperties, "isMarginEnabled")
		delete(additionalProperties, "isFutureEnabled")
		delete(additionalProperties, "isOptionsEnabled")
		delete(additionalProperties, "isPortfolioMarginRetailEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountInfoResponse struct {
	value *AccountInfoResponse
	isSet bool
}

func (v NullableAccountInfoResponse) Get() *AccountInfoResponse {
	return v.value
}

func (v *NullableAccountInfoResponse) Set(val *AccountInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInfoResponse(val *AccountInfoResponse) *NullableAccountInfoResponse {
	return &NullableAccountInfoResponse{value: val, isSet: true}
}

func (v NullableAccountInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
