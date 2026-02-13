/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetRateHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetRateHistoryResponseRowsInner{}

// GetRateHistoryResponseRowsInner struct for GetRateHistoryResponseRowsInner
type GetRateHistoryResponseRowsInner struct {
	ProductId            *string `json:"productId,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	AnnualPercentageRate *string `json:"annualPercentageRate,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRateHistoryResponseRowsInner GetRateHistoryResponseRowsInner

// NewGetRateHistoryResponseRowsInner instantiates a new GetRateHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRateHistoryResponseRowsInner() *GetRateHistoryResponseRowsInner {
	this := GetRateHistoryResponseRowsInner{}
	return &this
}

// NewGetRateHistoryResponseRowsInnerWithDefaults instantiates a new GetRateHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRateHistoryResponseRowsInnerWithDefaults() *GetRateHistoryResponseRowsInner {
	this := GetRateHistoryResponseRowsInner{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *GetRateHistoryResponseRowsInner) GetProductId() string {
	if o == nil || common.IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRateHistoryResponseRowsInner) GetProductIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *GetRateHistoryResponseRowsInner) HasProductId() bool {
	if o != nil && !common.IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *GetRateHistoryResponseRowsInner) SetProductId(v string) {
	o.ProductId = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetRateHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRateHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetRateHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetRateHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAnnualPercentageRate returns the AnnualPercentageRate field value if set, zero value otherwise.
func (o *GetRateHistoryResponseRowsInner) GetAnnualPercentageRate() string {
	if o == nil || common.IsNil(o.AnnualPercentageRate) {
		var ret string
		return ret
	}
	return *o.AnnualPercentageRate
}

// GetAnnualPercentageRateOk returns a tuple with the AnnualPercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRateHistoryResponseRowsInner) GetAnnualPercentageRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.AnnualPercentageRate) {
		return nil, false
	}
	return o.AnnualPercentageRate, true
}

// HasAnnualPercentageRate returns a boolean if a field has been set.
func (o *GetRateHistoryResponseRowsInner) HasAnnualPercentageRate() bool {
	if o != nil && !common.IsNil(o.AnnualPercentageRate) {
		return true
	}

	return false
}

// SetAnnualPercentageRate gets a reference to the given string and assigns it to the AnnualPercentageRate field.
func (o *GetRateHistoryResponseRowsInner) SetAnnualPercentageRate(v string) {
	o.AnnualPercentageRate = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetRateHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRateHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetRateHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetRateHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

func (o GetRateHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRateHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.AnnualPercentageRate) {
		toSerialize["annualPercentageRate"] = o.AnnualPercentageRate
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetRateHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetRateHistoryResponseRowsInner := _GetRateHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetRateHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetRateHistoryResponseRowsInner(varGetRateHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "productId")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "annualPercentageRate")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRateHistoryResponseRowsInner struct {
	value *GetRateHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetRateHistoryResponseRowsInner) Get() *GetRateHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetRateHistoryResponseRowsInner) Set(val *GetRateHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRateHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRateHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRateHistoryResponseRowsInner(val *GetRateHistoryResponseRowsInner) *NullableGetRateHistoryResponseRowsInner {
	return &NullableGetRateHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetRateHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRateHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
