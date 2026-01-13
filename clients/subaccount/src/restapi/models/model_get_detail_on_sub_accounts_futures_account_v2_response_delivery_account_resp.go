/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp{}

// GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp struct for GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp
type GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp struct {
	Email                *string                                                                        `json:"email,omitempty"`
	Assets               []GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountRespAssetsInner `json:"assets,omitempty"`
	CanDeposit           *bool                                                                          `json:"canDeposit,omitempty"`
	CanTrade             *bool                                                                          `json:"canTrade,omitempty"`
	CanWithdraw          *bool                                                                          `json:"canWithdraw,omitempty"`
	FeeTier              *int64                                                                         `json:"feeTier,omitempty"`
	UpdateTime           *int64                                                                         `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp

// NewGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp instantiates a new GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp() *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp {
	this := GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp{}
	return &this
}

// NewGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountRespWithDefaults instantiates a new GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountRespWithDefaults() *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp {
	this := GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) SetEmail(v string) {
	o.Email = &v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetAssets() []GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountRespAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountRespAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetAssetsOk() ([]GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountRespAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountRespAssetsInner and assigns it to the Assets field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) SetAssets(v []GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountRespAssetsInner) {
	o.Assets = v
}

// GetCanDeposit returns the CanDeposit field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetCanDeposit() bool {
	if o == nil || common.IsNil(o.CanDeposit) {
		var ret bool
		return ret
	}
	return *o.CanDeposit
}

// GetCanDepositOk returns a tuple with the CanDeposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetCanDepositOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanDeposit) {
		return nil, false
	}
	return o.CanDeposit, true
}

// HasCanDeposit returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) HasCanDeposit() bool {
	if o != nil && !common.IsNil(o.CanDeposit) {
		return true
	}

	return false
}

// SetCanDeposit gets a reference to the given bool and assigns it to the CanDeposit field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) SetCanDeposit(v bool) {
	o.CanDeposit = &v
}

// GetCanTrade returns the CanTrade field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetCanTrade() bool {
	if o == nil || common.IsNil(o.CanTrade) {
		var ret bool
		return ret
	}
	return *o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetCanTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanTrade) {
		return nil, false
	}
	return o.CanTrade, true
}

// HasCanTrade returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) HasCanTrade() bool {
	if o != nil && !common.IsNil(o.CanTrade) {
		return true
	}

	return false
}

// SetCanTrade gets a reference to the given bool and assigns it to the CanTrade field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) SetCanTrade(v bool) {
	o.CanTrade = &v
}

// GetCanWithdraw returns the CanWithdraw field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetCanWithdraw() bool {
	if o == nil || common.IsNil(o.CanWithdraw) {
		var ret bool
		return ret
	}
	return *o.CanWithdraw
}

// GetCanWithdrawOk returns a tuple with the CanWithdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetCanWithdrawOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanWithdraw) {
		return nil, false
	}
	return o.CanWithdraw, true
}

// HasCanWithdraw returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) HasCanWithdraw() bool {
	if o != nil && !common.IsNil(o.CanWithdraw) {
		return true
	}

	return false
}

// SetCanWithdraw gets a reference to the given bool and assigns it to the CanWithdraw field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) SetCanWithdraw(v bool) {
	o.CanWithdraw = &v
}

// GetFeeTier returns the FeeTier field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetFeeTier() int64 {
	if o == nil || common.IsNil(o.FeeTier) {
		var ret int64
		return ret
	}
	return *o.FeeTier
}

// GetFeeTierOk returns a tuple with the FeeTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetFeeTierOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FeeTier) {
		return nil, false
	}
	return o.FeeTier, true
}

// HasFeeTier returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) HasFeeTier() bool {
	if o != nil && !common.IsNil(o.FeeTier) {
		return true
	}

	return false
}

// SetFeeTier gets a reference to the given int64 and assigns it to the FeeTier field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) SetFeeTier(v int64) {
	o.FeeTier = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !common.IsNil(o.CanDeposit) {
		toSerialize["canDeposit"] = o.CanDeposit
	}
	if !common.IsNil(o.CanTrade) {
		toSerialize["canTrade"] = o.CanTrade
	}
	if !common.IsNil(o.CanWithdraw) {
		toSerialize["canWithdraw"] = o.CanWithdraw
	}
	if !common.IsNil(o.FeeTier) {
		toSerialize["feeTier"] = o.FeeTier
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) UnmarshalJSON(data []byte) (err error) {
	varGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp := _GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp{}

	err = json.Unmarshal(data, &varGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp)

	if err != nil {
		return err
	}

	*o = GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp(varGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "assets")
		delete(additionalProperties, "canDeposit")
		delete(additionalProperties, "canTrade")
		delete(additionalProperties, "canWithdraw")
		delete(additionalProperties, "feeTier")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp struct {
	value *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp
	isSet bool
}

func (v NullableGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) Get() *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp {
	return v.value
}

func (v *NullableGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) Set(val *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp(val *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) *NullableGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp {
	return &NullableGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp{value: val, isSet: true}
}

func (v NullableGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
