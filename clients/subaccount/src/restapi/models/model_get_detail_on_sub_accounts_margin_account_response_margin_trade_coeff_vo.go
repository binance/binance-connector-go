/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo{}

// GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo struct for GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo
type GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo struct {
	ForceLiquidationBar  *string `json:"forceLiquidationBar,omitempty"`
	MarginCallBar        *string `json:"marginCallBar,omitempty"`
	NormalBar            *string `json:"normalBar,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo

// NewGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo instantiates a new GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo() *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo {
	this := GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo{}
	return &this
}

// NewGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVoWithDefaults instantiates a new GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVoWithDefaults() *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo {
	this := GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo{}
	return &this
}

// GetForceLiquidationBar returns the ForceLiquidationBar field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) GetForceLiquidationBar() string {
	if o == nil || common.IsNil(o.ForceLiquidationBar) {
		var ret string
		return ret
	}
	return *o.ForceLiquidationBar
}

// GetForceLiquidationBarOk returns a tuple with the ForceLiquidationBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) GetForceLiquidationBarOk() (*string, bool) {
	if o == nil || common.IsNil(o.ForceLiquidationBar) {
		return nil, false
	}
	return o.ForceLiquidationBar, true
}

// HasForceLiquidationBar returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) HasForceLiquidationBar() bool {
	if o != nil && !common.IsNil(o.ForceLiquidationBar) {
		return true
	}

	return false
}

// SetForceLiquidationBar gets a reference to the given string and assigns it to the ForceLiquidationBar field.
func (o *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) SetForceLiquidationBar(v string) {
	o.ForceLiquidationBar = &v
}

// GetMarginCallBar returns the MarginCallBar field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) GetMarginCallBar() string {
	if o == nil || common.IsNil(o.MarginCallBar) {
		var ret string
		return ret
	}
	return *o.MarginCallBar
}

// GetMarginCallBarOk returns a tuple with the MarginCallBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) GetMarginCallBarOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginCallBar) {
		return nil, false
	}
	return o.MarginCallBar, true
}

// HasMarginCallBar returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) HasMarginCallBar() bool {
	if o != nil && !common.IsNil(o.MarginCallBar) {
		return true
	}

	return false
}

// SetMarginCallBar gets a reference to the given string and assigns it to the MarginCallBar field.
func (o *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) SetMarginCallBar(v string) {
	o.MarginCallBar = &v
}

// GetNormalBar returns the NormalBar field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) GetNormalBar() string {
	if o == nil || common.IsNil(o.NormalBar) {
		var ret string
		return ret
	}
	return *o.NormalBar
}

// GetNormalBarOk returns a tuple with the NormalBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) GetNormalBarOk() (*string, bool) {
	if o == nil || common.IsNil(o.NormalBar) {
		return nil, false
	}
	return o.NormalBar, true
}

// HasNormalBar returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) HasNormalBar() bool {
	if o != nil && !common.IsNil(o.NormalBar) {
		return true
	}

	return false
}

// SetNormalBar gets a reference to the given string and assigns it to the NormalBar field.
func (o *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) SetNormalBar(v string) {
	o.NormalBar = &v
}

func (o GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ForceLiquidationBar) {
		toSerialize["forceLiquidationBar"] = o.ForceLiquidationBar
	}
	if !common.IsNil(o.MarginCallBar) {
		toSerialize["marginCallBar"] = o.MarginCallBar
	}
	if !common.IsNil(o.NormalBar) {
		toSerialize["normalBar"] = o.NormalBar
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) UnmarshalJSON(data []byte) (err error) {
	varGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo := _GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo{}

	err = json.Unmarshal(data, &varGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo)

	if err != nil {
		return err
	}

	*o = GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo(varGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "forceLiquidationBar")
		delete(additionalProperties, "marginCallBar")
		delete(additionalProperties, "normalBar")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo struct {
	value *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo
	isSet bool
}

func (v NullableGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) Get() *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo {
	return v.value
}

func (v *NullableGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) Set(val *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo(val *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) *NullableGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo {
	return &NullableGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo{value: val, isSet: true}
}

func (v NullableGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
