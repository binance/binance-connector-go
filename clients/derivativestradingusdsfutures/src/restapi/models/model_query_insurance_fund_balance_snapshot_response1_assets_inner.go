/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryInsuranceFundBalanceSnapshotResponse1AssetsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryInsuranceFundBalanceSnapshotResponse1AssetsInner{}

// QueryInsuranceFundBalanceSnapshotResponse1AssetsInner struct for QueryInsuranceFundBalanceSnapshotResponse1AssetsInner
type QueryInsuranceFundBalanceSnapshotResponse1AssetsInner struct {
	Asset                *string `json:"asset,omitempty"`
	MarginBalance        *string `json:"marginBalance,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryInsuranceFundBalanceSnapshotResponse1AssetsInner QueryInsuranceFundBalanceSnapshotResponse1AssetsInner

// NewQueryInsuranceFundBalanceSnapshotResponse1AssetsInner instantiates a new QueryInsuranceFundBalanceSnapshotResponse1AssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryInsuranceFundBalanceSnapshotResponse1AssetsInner() *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner {
	this := QueryInsuranceFundBalanceSnapshotResponse1AssetsInner{}
	return &this
}

// NewQueryInsuranceFundBalanceSnapshotResponse1AssetsInnerWithDefaults instantiates a new QueryInsuranceFundBalanceSnapshotResponse1AssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryInsuranceFundBalanceSnapshotResponse1AssetsInnerWithDefaults() *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner {
	this := QueryInsuranceFundBalanceSnapshotResponse1AssetsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetMarginBalance returns the MarginBalance field value if set, zero value otherwise.
func (o *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) GetMarginBalance() string {
	if o == nil || common.IsNil(o.MarginBalance) {
		var ret string
		return ret
	}
	return *o.MarginBalance
}

// GetMarginBalanceOk returns a tuple with the MarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) GetMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginBalance) {
		return nil, false
	}
	return o.MarginBalance, true
}

// HasMarginBalance returns a boolean if a field has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) HasMarginBalance() bool {
	if o != nil && !common.IsNil(o.MarginBalance) {
		return true
	}

	return false
}

// SetMarginBalance gets a reference to the given string and assigns it to the MarginBalance field.
func (o *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) SetMarginBalance(v string) {
	o.MarginBalance = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) ToMap() (map[string]interface{}, error) {
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

func (o *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryInsuranceFundBalanceSnapshotResponse1AssetsInner := _QueryInsuranceFundBalanceSnapshotResponse1AssetsInner{}

	err = json.Unmarshal(data, &varQueryInsuranceFundBalanceSnapshotResponse1AssetsInner)

	if err != nil {
		return err
	}

	*o = QueryInsuranceFundBalanceSnapshotResponse1AssetsInner(varQueryInsuranceFundBalanceSnapshotResponse1AssetsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "marginBalance")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryInsuranceFundBalanceSnapshotResponse1AssetsInner struct {
	value *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner
	isSet bool
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse1AssetsInner) Get() *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner {
	return v.value
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse1AssetsInner) Set(val *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse1AssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse1AssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryInsuranceFundBalanceSnapshotResponse1AssetsInner(val *QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) *NullableQueryInsuranceFundBalanceSnapshotResponse1AssetsInner {
	return &NullableQueryInsuranceFundBalanceSnapshotResponse1AssetsInner{value: val, isSet: true}
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse1AssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse1AssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
