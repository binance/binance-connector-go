/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner{}

// QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner struct for QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner
type QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner struct {
	Asset                *string `json:"asset,omitempty"`
	Interest             *string `json:"interest,omitempty"`
	InterestAccruedTime  *int64  `json:"interestAccruedTime,omitempty"`
	InterestRate         *string `json:"interestRate,omitempty"`
	Principal            *string `json:"principal,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner

// NewQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner instantiates a new QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner() *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner {
	this := QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner{}
	return &this
}

// NewQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInnerWithDefaults instantiates a new QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInnerWithDefaults() *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner {
	this := QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) GetInterest() string {
	if o == nil || common.IsNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) GetInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) HasInterest() bool {
	if o != nil && !common.IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) SetInterest(v string) {
	o.Interest = &v
}

// GetInterestAccruedTime returns the InterestAccruedTime field value if set, zero value otherwise.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) GetInterestAccruedTime() int64 {
	if o == nil || common.IsNil(o.InterestAccruedTime) {
		var ret int64
		return ret
	}
	return *o.InterestAccruedTime
}

// GetInterestAccruedTimeOk returns a tuple with the InterestAccruedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) GetInterestAccruedTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.InterestAccruedTime) {
		return nil, false
	}
	return o.InterestAccruedTime, true
}

// HasInterestAccruedTime returns a boolean if a field has been set.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) HasInterestAccruedTime() bool {
	if o != nil && !common.IsNil(o.InterestAccruedTime) {
		return true
	}

	return false
}

// SetInterestAccruedTime gets a reference to the given int64 and assigns it to the InterestAccruedTime field.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) SetInterestAccruedTime(v int64) {
	o.InterestAccruedTime = &v
}

// GetInterestRate returns the InterestRate field value if set, zero value otherwise.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) GetInterestRate() string {
	if o == nil || common.IsNil(o.InterestRate) {
		var ret string
		return ret
	}
	return *o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) GetInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.InterestRate) {
		return nil, false
	}
	return o.InterestRate, true
}

// HasInterestRate returns a boolean if a field has been set.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) HasInterestRate() bool {
	if o != nil && !common.IsNil(o.InterestRate) {
		return true
	}

	return false
}

// SetInterestRate gets a reference to the given string and assigns it to the InterestRate field.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) SetInterestRate(v string) {
	o.InterestRate = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) GetPrincipal() string {
	if o == nil || common.IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) GetPrincipalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) HasPrincipal() bool {
	if o != nil && !common.IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) SetPrincipal(v string) {
	o.Principal = &v
}

func (o QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !common.IsNil(o.InterestAccruedTime) {
		toSerialize["interestAccruedTime"] = o.InterestAccruedTime
	}
	if !common.IsNil(o.InterestRate) {
		toSerialize["interestRate"] = o.InterestRate
	}
	if !common.IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner := _QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner{}

	err = json.Unmarshal(data, &varQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner)

	if err != nil {
		return err
	}

	*o = QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner(varQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "interest")
		delete(additionalProperties, "interestAccruedTime")
		delete(additionalProperties, "interestRate")
		delete(additionalProperties, "principal")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner struct {
	value *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner
	isSet bool
}

func (v NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) Get() *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner {
	return v.value
}

func (v *NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) Set(val *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner(val *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) *NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner {
	return &NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner{value: val, isSet: true}
}

func (v NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
