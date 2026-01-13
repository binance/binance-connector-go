/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetUmAccountDetailV2ResponseAssetsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetUmAccountDetailV2ResponseAssetsInner{}

// GetUmAccountDetailV2ResponseAssetsInner struct for GetUmAccountDetailV2ResponseAssetsInner
type GetUmAccountDetailV2ResponseAssetsInner struct {
	Asset                  *string `json:"asset,omitempty"`
	CrossWalletBalance     *string `json:"crossWalletBalance,omitempty"`
	CrossUnPnl             *string `json:"crossUnPnl,omitempty"`
	MaintMargin            *string `json:"maintMargin,omitempty"`
	InitialMargin          *string `json:"initialMargin,omitempty"`
	PositionInitialMargin  *string `json:"positionInitialMargin,omitempty"`
	OpenOrderInitialMargin *string `json:"openOrderInitialMargin,omitempty"`
	UpdateTime             *int64  `json:"updateTime,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _GetUmAccountDetailV2ResponseAssetsInner GetUmAccountDetailV2ResponseAssetsInner

// NewGetUmAccountDetailV2ResponseAssetsInner instantiates a new GetUmAccountDetailV2ResponseAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUmAccountDetailV2ResponseAssetsInner() *GetUmAccountDetailV2ResponseAssetsInner {
	this := GetUmAccountDetailV2ResponseAssetsInner{}
	return &this
}

// NewGetUmAccountDetailV2ResponseAssetsInnerWithDefaults instantiates a new GetUmAccountDetailV2ResponseAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUmAccountDetailV2ResponseAssetsInnerWithDefaults() *GetUmAccountDetailV2ResponseAssetsInner {
	this := GetUmAccountDetailV2ResponseAssetsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetUmAccountDetailV2ResponseAssetsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailV2ResponseAssetsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetUmAccountDetailV2ResponseAssetsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetUmAccountDetailV2ResponseAssetsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetCrossWalletBalance returns the CrossWalletBalance field value if set, zero value otherwise.
func (o *GetUmAccountDetailV2ResponseAssetsInner) GetCrossWalletBalance() string {
	if o == nil || common.IsNil(o.CrossWalletBalance) {
		var ret string
		return ret
	}
	return *o.CrossWalletBalance
}

// GetCrossWalletBalanceOk returns a tuple with the CrossWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailV2ResponseAssetsInner) GetCrossWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossWalletBalance) {
		return nil, false
	}
	return o.CrossWalletBalance, true
}

// HasCrossWalletBalance returns a boolean if a field has been set.
func (o *GetUmAccountDetailV2ResponseAssetsInner) HasCrossWalletBalance() bool {
	if o != nil && !common.IsNil(o.CrossWalletBalance) {
		return true
	}

	return false
}

// SetCrossWalletBalance gets a reference to the given string and assigns it to the CrossWalletBalance field.
func (o *GetUmAccountDetailV2ResponseAssetsInner) SetCrossWalletBalance(v string) {
	o.CrossWalletBalance = &v
}

// GetCrossUnPnl returns the CrossUnPnl field value if set, zero value otherwise.
func (o *GetUmAccountDetailV2ResponseAssetsInner) GetCrossUnPnl() string {
	if o == nil || common.IsNil(o.CrossUnPnl) {
		var ret string
		return ret
	}
	return *o.CrossUnPnl
}

// GetCrossUnPnlOk returns a tuple with the CrossUnPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailV2ResponseAssetsInner) GetCrossUnPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossUnPnl) {
		return nil, false
	}
	return o.CrossUnPnl, true
}

// HasCrossUnPnl returns a boolean if a field has been set.
func (o *GetUmAccountDetailV2ResponseAssetsInner) HasCrossUnPnl() bool {
	if o != nil && !common.IsNil(o.CrossUnPnl) {
		return true
	}

	return false
}

// SetCrossUnPnl gets a reference to the given string and assigns it to the CrossUnPnl field.
func (o *GetUmAccountDetailV2ResponseAssetsInner) SetCrossUnPnl(v string) {
	o.CrossUnPnl = &v
}

// GetMaintMargin returns the MaintMargin field value if set, zero value otherwise.
func (o *GetUmAccountDetailV2ResponseAssetsInner) GetMaintMargin() string {
	if o == nil || common.IsNil(o.MaintMargin) {
		var ret string
		return ret
	}
	return *o.MaintMargin
}

// GetMaintMarginOk returns a tuple with the MaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailV2ResponseAssetsInner) GetMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaintMargin) {
		return nil, false
	}
	return o.MaintMargin, true
}

// HasMaintMargin returns a boolean if a field has been set.
func (o *GetUmAccountDetailV2ResponseAssetsInner) HasMaintMargin() bool {
	if o != nil && !common.IsNil(o.MaintMargin) {
		return true
	}

	return false
}

// SetMaintMargin gets a reference to the given string and assigns it to the MaintMargin field.
func (o *GetUmAccountDetailV2ResponseAssetsInner) SetMaintMargin(v string) {
	o.MaintMargin = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *GetUmAccountDetailV2ResponseAssetsInner) GetInitialMargin() string {
	if o == nil || common.IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailV2ResponseAssetsInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *GetUmAccountDetailV2ResponseAssetsInner) HasInitialMargin() bool {
	if o != nil && !common.IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *GetUmAccountDetailV2ResponseAssetsInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetPositionInitialMargin returns the PositionInitialMargin field value if set, zero value otherwise.
func (o *GetUmAccountDetailV2ResponseAssetsInner) GetPositionInitialMargin() string {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.PositionInitialMargin
}

// GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailV2ResponseAssetsInner) GetPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		return nil, false
	}
	return o.PositionInitialMargin, true
}

// HasPositionInitialMargin returns a boolean if a field has been set.
func (o *GetUmAccountDetailV2ResponseAssetsInner) HasPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.PositionInitialMargin) {
		return true
	}

	return false
}

// SetPositionInitialMargin gets a reference to the given string and assigns it to the PositionInitialMargin field.
func (o *GetUmAccountDetailV2ResponseAssetsInner) SetPositionInitialMargin(v string) {
	o.PositionInitialMargin = &v
}

// GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field value if set, zero value otherwise.
func (o *GetUmAccountDetailV2ResponseAssetsInner) GetOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.OpenOrderInitialMargin
}

// GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailV2ResponseAssetsInner) GetOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		return nil, false
	}
	return o.OpenOrderInitialMargin, true
}

// HasOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *GetUmAccountDetailV2ResponseAssetsInner) HasOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.OpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetOpenOrderInitialMargin gets a reference to the given string and assigns it to the OpenOrderInitialMargin field.
func (o *GetUmAccountDetailV2ResponseAssetsInner) SetOpenOrderInitialMargin(v string) {
	o.OpenOrderInitialMargin = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetUmAccountDetailV2ResponseAssetsInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailV2ResponseAssetsInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetUmAccountDetailV2ResponseAssetsInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetUmAccountDetailV2ResponseAssetsInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o GetUmAccountDetailV2ResponseAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUmAccountDetailV2ResponseAssetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.CrossWalletBalance) {
		toSerialize["crossWalletBalance"] = o.CrossWalletBalance
	}
	if !common.IsNil(o.CrossUnPnl) {
		toSerialize["crossUnPnl"] = o.CrossUnPnl
	}
	if !common.IsNil(o.MaintMargin) {
		toSerialize["maintMargin"] = o.MaintMargin
	}
	if !common.IsNil(o.InitialMargin) {
		toSerialize["initialMargin"] = o.InitialMargin
	}
	if !common.IsNil(o.PositionInitialMargin) {
		toSerialize["positionInitialMargin"] = o.PositionInitialMargin
	}
	if !common.IsNil(o.OpenOrderInitialMargin) {
		toSerialize["openOrderInitialMargin"] = o.OpenOrderInitialMargin
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUmAccountDetailV2ResponseAssetsInner) UnmarshalJSON(data []byte) (err error) {
	varGetUmAccountDetailV2ResponseAssetsInner := _GetUmAccountDetailV2ResponseAssetsInner{}

	err = json.Unmarshal(data, &varGetUmAccountDetailV2ResponseAssetsInner)

	if err != nil {
		return err
	}

	*o = GetUmAccountDetailV2ResponseAssetsInner(varGetUmAccountDetailV2ResponseAssetsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "crossWalletBalance")
		delete(additionalProperties, "crossUnPnl")
		delete(additionalProperties, "maintMargin")
		delete(additionalProperties, "initialMargin")
		delete(additionalProperties, "positionInitialMargin")
		delete(additionalProperties, "openOrderInitialMargin")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUmAccountDetailV2ResponseAssetsInner struct {
	value *GetUmAccountDetailV2ResponseAssetsInner
	isSet bool
}

func (v NullableGetUmAccountDetailV2ResponseAssetsInner) Get() *GetUmAccountDetailV2ResponseAssetsInner {
	return v.value
}

func (v *NullableGetUmAccountDetailV2ResponseAssetsInner) Set(val *GetUmAccountDetailV2ResponseAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUmAccountDetailV2ResponseAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUmAccountDetailV2ResponseAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUmAccountDetailV2ResponseAssetsInner(val *GetUmAccountDetailV2ResponseAssetsInner) *NullableGetUmAccountDetailV2ResponseAssetsInner {
	return &NullableGetUmAccountDetailV2ResponseAssetsInner{value: val, isSet: true}
}

func (v NullableGetUmAccountDetailV2ResponseAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUmAccountDetailV2ResponseAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
