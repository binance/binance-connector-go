/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner{}

// QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner struct for QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner
type QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner struct {
	Asset                *string `json:"asset,omitempty"`
	MarginBalance        *string `json:"marginBalance,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner

// NewQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner instantiates a new QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner() *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner {
	this := QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner{}
	return &this
}

// NewQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInnerWithDefaults instantiates a new QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInnerWithDefaults() *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner {
	this := QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetMarginBalance returns the MarginBalance field value if set, zero value otherwise.
func (o *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) GetMarginBalance() string {
	if o == nil || common.IsNil(o.MarginBalance) {
		var ret string
		return ret
	}
	return *o.MarginBalance
}

// GetMarginBalanceOk returns a tuple with the MarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) GetMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginBalance) {
		return nil, false
	}
	return o.MarginBalance, true
}

// HasMarginBalance returns a boolean if a field has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) HasMarginBalance() bool {
	if o != nil && !common.IsNil(o.MarginBalance) {
		return true
	}

	return false
}

// SetMarginBalance gets a reference to the given string and assigns it to the MarginBalance field.
func (o *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) SetMarginBalance(v string) {
	o.MarginBalance = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.MarginBalance) {
		toSerialize["marginBalance"] = o.MarginBalance
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner := _QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner{}

	err = json.Unmarshal(data, &varQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner)

	if err != nil {
		return err
	}

	*o = QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner(varQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "marginBalance")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner struct {
	value *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner
	isSet bool
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) Get() *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner {
	return v.value
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) Set(val *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner(val *QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) *NullableQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner {
	return &NullableQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner{value: val, isSet: true}
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
