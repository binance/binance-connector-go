/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner{}

// QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner struct for QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner
type QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner struct {
	Asset                *string `json:"asset,omitempty"`
	Interest             *string `json:"interest,omitempty"`
	InterestAccuredTime  *int64  `json:"interestAccuredTime,omitempty"`
	InterestRate         *string `json:"interestRate,omitempty"`
	Principal            *string `json:"principal,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner

// NewQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner instantiates a new QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner() *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner {
	this := QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner{}
	return &this
}

// NewQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInnerWithDefaults instantiates a new QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInnerWithDefaults() *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner {
	this := QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) GetInterest() string {
	if o == nil || common.IsNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) GetInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) HasInterest() bool {
	if o != nil && !common.IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) SetInterest(v string) {
	o.Interest = &v
}

// GetInterestAccuredTime returns the InterestAccuredTime field value if set, zero value otherwise.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) GetInterestAccuredTime() int64 {
	if o == nil || common.IsNil(o.InterestAccuredTime) {
		var ret int64
		return ret
	}
	return *o.InterestAccuredTime
}

// GetInterestAccuredTimeOk returns a tuple with the InterestAccuredTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) GetInterestAccuredTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.InterestAccuredTime) {
		return nil, false
	}
	return o.InterestAccuredTime, true
}

// HasInterestAccuredTime returns a boolean if a field has been set.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) HasInterestAccuredTime() bool {
	if o != nil && !common.IsNil(o.InterestAccuredTime) {
		return true
	}

	return false
}

// SetInterestAccuredTime gets a reference to the given int64 and assigns it to the InterestAccuredTime field.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) SetInterestAccuredTime(v int64) {
	o.InterestAccuredTime = &v
}

// GetInterestRate returns the InterestRate field value if set, zero value otherwise.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) GetInterestRate() string {
	if o == nil || common.IsNil(o.InterestRate) {
		var ret string
		return ret
	}
	return *o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) GetInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.InterestRate) {
		return nil, false
	}
	return o.InterestRate, true
}

// HasInterestRate returns a boolean if a field has been set.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) HasInterestRate() bool {
	if o != nil && !common.IsNil(o.InterestRate) {
		return true
	}

	return false
}

// SetInterestRate gets a reference to the given string and assigns it to the InterestRate field.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) SetInterestRate(v string) {
	o.InterestRate = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) GetPrincipal() string {
	if o == nil || common.IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) GetPrincipalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) HasPrincipal() bool {
	if o != nil && !common.IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) SetPrincipal(v string) {
	o.Principal = &v
}

func (o QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !common.IsNil(o.InterestAccuredTime) {
		toSerialize["interestAccuredTime"] = o.InterestAccuredTime
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

func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner := _QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner{}

	err = json.Unmarshal(data, &varQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner)

	if err != nil {
		return err
	}

	*o = QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner(varQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "interest")
		delete(additionalProperties, "interestAccuredTime")
		delete(additionalProperties, "interestRate")
		delete(additionalProperties, "principal")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner struct {
	value *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner
	isSet bool
}

func (v NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) Get() *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner {
	return v.value
}

func (v *NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) Set(val *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner(val *QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) *NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner {
	return &NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner{value: val, isSet: true}
}

func (v NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
