/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetUmAccountDetailResponseAssetsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetUmAccountDetailResponseAssetsInner{}

// GetUmAccountDetailResponseAssetsInner struct for GetUmAccountDetailResponseAssetsInner
type GetUmAccountDetailResponseAssetsInner struct {
	// asset name
	Asset *string `json:"asset,omitempty"`
	// wallet balance
	CrossWalletBalance *string `json:"crossWalletBalance,omitempty"`
	// unrealized profit
	CrossUnPnl *string `json:"crossUnPnl,omitempty"`
	// maintenance margin required
	MaintMargin *string `json:"maintMargin,omitempty"`
	// total initial margin required with current mark price
	InitialMargin *string `json:"initialMargin,omitempty"`
	// initial margin required for positions with current mark price
	PositionInitialMargin *string `json:"positionInitialMargin,omitempty"`
	// initial margin required for open orders with current mark price
	OpenOrderInitialMargin *string `json:"openOrderInitialMargin,omitempty"`
	// last update time
	UpdateTime           *int64 `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUmAccountDetailResponseAssetsInner GetUmAccountDetailResponseAssetsInner

// NewGetUmAccountDetailResponseAssetsInner instantiates a new GetUmAccountDetailResponseAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUmAccountDetailResponseAssetsInner() *GetUmAccountDetailResponseAssetsInner {
	this := GetUmAccountDetailResponseAssetsInner{}
	return &this
}

// NewGetUmAccountDetailResponseAssetsInnerWithDefaults instantiates a new GetUmAccountDetailResponseAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUmAccountDetailResponseAssetsInnerWithDefaults() *GetUmAccountDetailResponseAssetsInner {
	this := GetUmAccountDetailResponseAssetsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponseAssetsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponseAssetsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponseAssetsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetUmAccountDetailResponseAssetsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetCrossWalletBalance returns the CrossWalletBalance field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponseAssetsInner) GetCrossWalletBalance() string {
	if o == nil || common.IsNil(o.CrossWalletBalance) {
		var ret string
		return ret
	}
	return *o.CrossWalletBalance
}

// GetCrossWalletBalanceOk returns a tuple with the CrossWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponseAssetsInner) GetCrossWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossWalletBalance) {
		return nil, false
	}
	return o.CrossWalletBalance, true
}

// HasCrossWalletBalance returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponseAssetsInner) HasCrossWalletBalance() bool {
	if o != nil && !common.IsNil(o.CrossWalletBalance) {
		return true
	}

	return false
}

// SetCrossWalletBalance gets a reference to the given string and assigns it to the CrossWalletBalance field.
func (o *GetUmAccountDetailResponseAssetsInner) SetCrossWalletBalance(v string) {
	o.CrossWalletBalance = &v
}

// GetCrossUnPnl returns the CrossUnPnl field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponseAssetsInner) GetCrossUnPnl() string {
	if o == nil || common.IsNil(o.CrossUnPnl) {
		var ret string
		return ret
	}
	return *o.CrossUnPnl
}

// GetCrossUnPnlOk returns a tuple with the CrossUnPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponseAssetsInner) GetCrossUnPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossUnPnl) {
		return nil, false
	}
	return o.CrossUnPnl, true
}

// HasCrossUnPnl returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponseAssetsInner) HasCrossUnPnl() bool {
	if o != nil && !common.IsNil(o.CrossUnPnl) {
		return true
	}

	return false
}

// SetCrossUnPnl gets a reference to the given string and assigns it to the CrossUnPnl field.
func (o *GetUmAccountDetailResponseAssetsInner) SetCrossUnPnl(v string) {
	o.CrossUnPnl = &v
}

// GetMaintMargin returns the MaintMargin field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponseAssetsInner) GetMaintMargin() string {
	if o == nil || common.IsNil(o.MaintMargin) {
		var ret string
		return ret
	}
	return *o.MaintMargin
}

// GetMaintMarginOk returns a tuple with the MaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponseAssetsInner) GetMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaintMargin) {
		return nil, false
	}
	return o.MaintMargin, true
}

// HasMaintMargin returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponseAssetsInner) HasMaintMargin() bool {
	if o != nil && !common.IsNil(o.MaintMargin) {
		return true
	}

	return false
}

// SetMaintMargin gets a reference to the given string and assigns it to the MaintMargin field.
func (o *GetUmAccountDetailResponseAssetsInner) SetMaintMargin(v string) {
	o.MaintMargin = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponseAssetsInner) GetInitialMargin() string {
	if o == nil || common.IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponseAssetsInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponseAssetsInner) HasInitialMargin() bool {
	if o != nil && !common.IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *GetUmAccountDetailResponseAssetsInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetPositionInitialMargin returns the PositionInitialMargin field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponseAssetsInner) GetPositionInitialMargin() string {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.PositionInitialMargin
}

// GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponseAssetsInner) GetPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		return nil, false
	}
	return o.PositionInitialMargin, true
}

// HasPositionInitialMargin returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponseAssetsInner) HasPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.PositionInitialMargin) {
		return true
	}

	return false
}

// SetPositionInitialMargin gets a reference to the given string and assigns it to the PositionInitialMargin field.
func (o *GetUmAccountDetailResponseAssetsInner) SetPositionInitialMargin(v string) {
	o.PositionInitialMargin = &v
}

// GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponseAssetsInner) GetOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.OpenOrderInitialMargin
}

// GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponseAssetsInner) GetOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		return nil, false
	}
	return o.OpenOrderInitialMargin, true
}

// HasOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponseAssetsInner) HasOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.OpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetOpenOrderInitialMargin gets a reference to the given string and assigns it to the OpenOrderInitialMargin field.
func (o *GetUmAccountDetailResponseAssetsInner) SetOpenOrderInitialMargin(v string) {
	o.OpenOrderInitialMargin = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponseAssetsInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponseAssetsInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponseAssetsInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetUmAccountDetailResponseAssetsInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o GetUmAccountDetailResponseAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUmAccountDetailResponseAssetsInner) ToMap() (map[string]interface{}, error) {
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

func (o *GetUmAccountDetailResponseAssetsInner) UnmarshalJSON(data []byte) (err error) {
	varGetUmAccountDetailResponseAssetsInner := _GetUmAccountDetailResponseAssetsInner{}

	err = json.Unmarshal(data, &varGetUmAccountDetailResponseAssetsInner)

	if err != nil {
		return err
	}

	*o = GetUmAccountDetailResponseAssetsInner(varGetUmAccountDetailResponseAssetsInner)

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

type NullableGetUmAccountDetailResponseAssetsInner struct {
	value *GetUmAccountDetailResponseAssetsInner
	isSet bool
}

func (v NullableGetUmAccountDetailResponseAssetsInner) Get() *GetUmAccountDetailResponseAssetsInner {
	return v.value
}

func (v *NullableGetUmAccountDetailResponseAssetsInner) Set(val *GetUmAccountDetailResponseAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUmAccountDetailResponseAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUmAccountDetailResponseAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUmAccountDetailResponseAssetsInner(val *GetUmAccountDetailResponseAssetsInner) *NullableGetUmAccountDetailResponseAssetsInner {
	return &NullableGetUmAccountDetailResponseAssetsInner{value: val, isSet: true}
}

func (v NullableGetUmAccountDetailResponseAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUmAccountDetailResponseAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
