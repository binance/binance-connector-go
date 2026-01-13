/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MarginAccountRepayDebtResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginAccountRepayDebtResponse{}

// MarginAccountRepayDebtResponse struct for MarginAccountRepayDebtResponse
type MarginAccountRepayDebtResponse struct {
	Amount               *string  `json:"amount,omitempty"`
	Asset                *string  `json:"asset,omitempty"`
	SpecifyRepayAssets   []string `json:"specifyRepayAssets,omitempty"`
	UpdateTime           *int64   `json:"updateTime,omitempty"`
	Success              *bool    `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginAccountRepayDebtResponse MarginAccountRepayDebtResponse

// NewMarginAccountRepayDebtResponse instantiates a new MarginAccountRepayDebtResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginAccountRepayDebtResponse() *MarginAccountRepayDebtResponse {
	this := MarginAccountRepayDebtResponse{}
	return &this
}

// NewMarginAccountRepayDebtResponseWithDefaults instantiates a new MarginAccountRepayDebtResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginAccountRepayDebtResponseWithDefaults() *MarginAccountRepayDebtResponse {
	this := MarginAccountRepayDebtResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *MarginAccountRepayDebtResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountRepayDebtResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *MarginAccountRepayDebtResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *MarginAccountRepayDebtResponse) SetAmount(v string) {
	o.Amount = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *MarginAccountRepayDebtResponse) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountRepayDebtResponse) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *MarginAccountRepayDebtResponse) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *MarginAccountRepayDebtResponse) SetAsset(v string) {
	o.Asset = &v
}

// GetSpecifyRepayAssets returns the SpecifyRepayAssets field value if set, zero value otherwise.
func (o *MarginAccountRepayDebtResponse) GetSpecifyRepayAssets() []string {
	if o == nil || common.IsNil(o.SpecifyRepayAssets) {
		var ret []string
		return ret
	}
	return o.SpecifyRepayAssets
}

// GetSpecifyRepayAssetsOk returns a tuple with the SpecifyRepayAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountRepayDebtResponse) GetSpecifyRepayAssetsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.SpecifyRepayAssets) {
		return nil, false
	}
	return o.SpecifyRepayAssets, true
}

// HasSpecifyRepayAssets returns a boolean if a field has been set.
func (o *MarginAccountRepayDebtResponse) HasSpecifyRepayAssets() bool {
	if o != nil && !common.IsNil(o.SpecifyRepayAssets) {
		return true
	}

	return false
}

// SetSpecifyRepayAssets gets a reference to the given []string and assigns it to the SpecifyRepayAssets field.
func (o *MarginAccountRepayDebtResponse) SetSpecifyRepayAssets(v []string) {
	o.SpecifyRepayAssets = v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *MarginAccountRepayDebtResponse) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountRepayDebtResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *MarginAccountRepayDebtResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *MarginAccountRepayDebtResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *MarginAccountRepayDebtResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountRepayDebtResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *MarginAccountRepayDebtResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *MarginAccountRepayDebtResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o MarginAccountRepayDebtResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginAccountRepayDebtResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.SpecifyRepayAssets) {
		toSerialize["specifyRepayAssets"] = o.SpecifyRepayAssets
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarginAccountRepayDebtResponse) UnmarshalJSON(data []byte) (err error) {
	varMarginAccountRepayDebtResponse := _MarginAccountRepayDebtResponse{}

	err = json.Unmarshal(data, &varMarginAccountRepayDebtResponse)

	if err != nil {
		return err
	}

	*o = MarginAccountRepayDebtResponse(varMarginAccountRepayDebtResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "specifyRepayAssets")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginAccountRepayDebtResponse struct {
	value *MarginAccountRepayDebtResponse
	isSet bool
}

func (v NullableMarginAccountRepayDebtResponse) Get() *MarginAccountRepayDebtResponse {
	return v.value
}

func (v *NullableMarginAccountRepayDebtResponse) Set(val *MarginAccountRepayDebtResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountRepayDebtResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountRepayDebtResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountRepayDebtResponse(val *MarginAccountRepayDebtResponse) *NullableMarginAccountRepayDebtResponse {
	return &NullableMarginAccountRepayDebtResponse{value: val, isSet: true}
}

func (v NullableMarginAccountRepayDebtResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountRepayDebtResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
