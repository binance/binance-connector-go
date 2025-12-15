/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSoftStakingProductListResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSoftStakingProductListResponseRowsInner{}

// GetSoftStakingProductListResponseRowsInner struct for GetSoftStakingProductListResponseRowsInner
type GetSoftStakingProductListResponseRowsInner struct {
	Asset                *string `json:"asset,omitempty"`
	MinAmount            *string `json:"minAmount,omitempty"`
	MaxCap               *string `json:"maxCap,omitempty"`
	Apr                  *string `json:"apr,omitempty"`
	StakedAmount         *string `json:"stakedAmount,omitempty"`
	TotalProfit          *string `json:"totalProfit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSoftStakingProductListResponseRowsInner GetSoftStakingProductListResponseRowsInner

// NewGetSoftStakingProductListResponseRowsInner instantiates a new GetSoftStakingProductListResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSoftStakingProductListResponseRowsInner() *GetSoftStakingProductListResponseRowsInner {
	this := GetSoftStakingProductListResponseRowsInner{}
	return &this
}

// NewGetSoftStakingProductListResponseRowsInnerWithDefaults instantiates a new GetSoftStakingProductListResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSoftStakingProductListResponseRowsInnerWithDefaults() *GetSoftStakingProductListResponseRowsInner {
	this := GetSoftStakingProductListResponseRowsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSoftStakingProductListResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingProductListResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSoftStakingProductListResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSoftStakingProductListResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetMinAmount returns the MinAmount field value if set, zero value otherwise.
func (o *GetSoftStakingProductListResponseRowsInner) GetMinAmount() string {
	if o == nil || common.IsNil(o.MinAmount) {
		var ret string
		return ret
	}
	return *o.MinAmount
}

// GetMinAmountOk returns a tuple with the MinAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingProductListResponseRowsInner) GetMinAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinAmount) {
		return nil, false
	}
	return o.MinAmount, true
}

// HasMinAmount returns a boolean if a field has been set.
func (o *GetSoftStakingProductListResponseRowsInner) HasMinAmount() bool {
	if o != nil && !common.IsNil(o.MinAmount) {
		return true
	}

	return false
}

// SetMinAmount gets a reference to the given string and assigns it to the MinAmount field.
func (o *GetSoftStakingProductListResponseRowsInner) SetMinAmount(v string) {
	o.MinAmount = &v
}

// GetMaxCap returns the MaxCap field value if set, zero value otherwise.
func (o *GetSoftStakingProductListResponseRowsInner) GetMaxCap() string {
	if o == nil || common.IsNil(o.MaxCap) {
		var ret string
		return ret
	}
	return *o.MaxCap
}

// GetMaxCapOk returns a tuple with the MaxCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingProductListResponseRowsInner) GetMaxCapOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxCap) {
		return nil, false
	}
	return o.MaxCap, true
}

// HasMaxCap returns a boolean if a field has been set.
func (o *GetSoftStakingProductListResponseRowsInner) HasMaxCap() bool {
	if o != nil && !common.IsNil(o.MaxCap) {
		return true
	}

	return false
}

// SetMaxCap gets a reference to the given string and assigns it to the MaxCap field.
func (o *GetSoftStakingProductListResponseRowsInner) SetMaxCap(v string) {
	o.MaxCap = &v
}

// GetApr returns the Apr field value if set, zero value otherwise.
func (o *GetSoftStakingProductListResponseRowsInner) GetApr() string {
	if o == nil || common.IsNil(o.Apr) {
		var ret string
		return ret
	}
	return *o.Apr
}

// GetAprOk returns a tuple with the Apr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingProductListResponseRowsInner) GetAprOk() (*string, bool) {
	if o == nil || common.IsNil(o.Apr) {
		return nil, false
	}
	return o.Apr, true
}

// HasApr returns a boolean if a field has been set.
func (o *GetSoftStakingProductListResponseRowsInner) HasApr() bool {
	if o != nil && !common.IsNil(o.Apr) {
		return true
	}

	return false
}

// SetApr gets a reference to the given string and assigns it to the Apr field.
func (o *GetSoftStakingProductListResponseRowsInner) SetApr(v string) {
	o.Apr = &v
}

// GetStakedAmount returns the StakedAmount field value if set, zero value otherwise.
func (o *GetSoftStakingProductListResponseRowsInner) GetStakedAmount() string {
	if o == nil || common.IsNil(o.StakedAmount) {
		var ret string
		return ret
	}
	return *o.StakedAmount
}

// GetStakedAmountOk returns a tuple with the StakedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingProductListResponseRowsInner) GetStakedAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.StakedAmount) {
		return nil, false
	}
	return o.StakedAmount, true
}

// HasStakedAmount returns a boolean if a field has been set.
func (o *GetSoftStakingProductListResponseRowsInner) HasStakedAmount() bool {
	if o != nil && !common.IsNil(o.StakedAmount) {
		return true
	}

	return false
}

// SetStakedAmount gets a reference to the given string and assigns it to the StakedAmount field.
func (o *GetSoftStakingProductListResponseRowsInner) SetStakedAmount(v string) {
	o.StakedAmount = &v
}

// GetTotalProfit returns the TotalProfit field value if set, zero value otherwise.
func (o *GetSoftStakingProductListResponseRowsInner) GetTotalProfit() string {
	if o == nil || common.IsNil(o.TotalProfit) {
		var ret string
		return ret
	}
	return *o.TotalProfit
}

// GetTotalProfitOk returns a tuple with the TotalProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingProductListResponseRowsInner) GetTotalProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalProfit) {
		return nil, false
	}
	return o.TotalProfit, true
}

// HasTotalProfit returns a boolean if a field has been set.
func (o *GetSoftStakingProductListResponseRowsInner) HasTotalProfit() bool {
	if o != nil && !common.IsNil(o.TotalProfit) {
		return true
	}

	return false
}

// SetTotalProfit gets a reference to the given string and assigns it to the TotalProfit field.
func (o *GetSoftStakingProductListResponseRowsInner) SetTotalProfit(v string) {
	o.TotalProfit = &v
}

func (o GetSoftStakingProductListResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSoftStakingProductListResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.MinAmount) {
		toSerialize["minAmount"] = o.MinAmount
	}
	if !common.IsNil(o.MaxCap) {
		toSerialize["maxCap"] = o.MaxCap
	}
	if !common.IsNil(o.Apr) {
		toSerialize["apr"] = o.Apr
	}
	if !common.IsNil(o.StakedAmount) {
		toSerialize["stakedAmount"] = o.StakedAmount
	}
	if !common.IsNil(o.TotalProfit) {
		toSerialize["totalProfit"] = o.TotalProfit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSoftStakingProductListResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetSoftStakingProductListResponseRowsInner := _GetSoftStakingProductListResponseRowsInner{}

	err = json.Unmarshal(data, &varGetSoftStakingProductListResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetSoftStakingProductListResponseRowsInner(varGetSoftStakingProductListResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "minAmount")
		delete(additionalProperties, "maxCap")
		delete(additionalProperties, "apr")
		delete(additionalProperties, "stakedAmount")
		delete(additionalProperties, "totalProfit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSoftStakingProductListResponseRowsInner struct {
	value *GetSoftStakingProductListResponseRowsInner
	isSet bool
}

func (v NullableGetSoftStakingProductListResponseRowsInner) Get() *GetSoftStakingProductListResponseRowsInner {
	return v.value
}

func (v *NullableGetSoftStakingProductListResponseRowsInner) Set(val *GetSoftStakingProductListResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSoftStakingProductListResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSoftStakingProductListResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSoftStakingProductListResponseRowsInner(val *GetSoftStakingProductListResponseRowsInner) *NullableGetSoftStakingProductListResponseRowsInner {
	return &NullableGetSoftStakingProductListResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetSoftStakingProductListResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSoftStakingProductListResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
