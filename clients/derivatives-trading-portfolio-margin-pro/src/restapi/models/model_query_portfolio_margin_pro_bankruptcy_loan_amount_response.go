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

// checks if the QueryPortfolioMarginProBankruptcyLoanAmountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPortfolioMarginProBankruptcyLoanAmountResponse{}

// QueryPortfolioMarginProBankruptcyLoanAmountResponse struct for QueryPortfolioMarginProBankruptcyLoanAmountResponse
type QueryPortfolioMarginProBankruptcyLoanAmountResponse struct {
	Asset                *string `json:"asset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryPortfolioMarginProBankruptcyLoanAmountResponse QueryPortfolioMarginProBankruptcyLoanAmountResponse

// NewQueryPortfolioMarginProBankruptcyLoanAmountResponse instantiates a new QueryPortfolioMarginProBankruptcyLoanAmountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPortfolioMarginProBankruptcyLoanAmountResponse() *QueryPortfolioMarginProBankruptcyLoanAmountResponse {
	this := QueryPortfolioMarginProBankruptcyLoanAmountResponse{}
	return &this
}

// NewQueryPortfolioMarginProBankruptcyLoanAmountResponseWithDefaults instantiates a new QueryPortfolioMarginProBankruptcyLoanAmountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPortfolioMarginProBankruptcyLoanAmountResponseWithDefaults() *QueryPortfolioMarginProBankruptcyLoanAmountResponse {
	this := QueryPortfolioMarginProBankruptcyLoanAmountResponse{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryPortfolioMarginProBankruptcyLoanAmountResponse) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginProBankruptcyLoanAmountResponse) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryPortfolioMarginProBankruptcyLoanAmountResponse) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryPortfolioMarginProBankruptcyLoanAmountResponse) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QueryPortfolioMarginProBankruptcyLoanAmountResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPortfolioMarginProBankruptcyLoanAmountResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QueryPortfolioMarginProBankruptcyLoanAmountResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QueryPortfolioMarginProBankruptcyLoanAmountResponse) SetAmount(v string) {
	o.Amount = &v
}

func (o QueryPortfolioMarginProBankruptcyLoanAmountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPortfolioMarginProBankruptcyLoanAmountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryPortfolioMarginProBankruptcyLoanAmountResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryPortfolioMarginProBankruptcyLoanAmountResponse := _QueryPortfolioMarginProBankruptcyLoanAmountResponse{}

	err = json.Unmarshal(data, &varQueryPortfolioMarginProBankruptcyLoanAmountResponse)

	if err != nil {
		return err
	}

	*o = QueryPortfolioMarginProBankruptcyLoanAmountResponse(varQueryPortfolioMarginProBankruptcyLoanAmountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryPortfolioMarginProBankruptcyLoanAmountResponse struct {
	value *QueryPortfolioMarginProBankruptcyLoanAmountResponse
	isSet bool
}

func (v NullableQueryPortfolioMarginProBankruptcyLoanAmountResponse) Get() *QueryPortfolioMarginProBankruptcyLoanAmountResponse {
	return v.value
}

func (v *NullableQueryPortfolioMarginProBankruptcyLoanAmountResponse) Set(val *QueryPortfolioMarginProBankruptcyLoanAmountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPortfolioMarginProBankruptcyLoanAmountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPortfolioMarginProBankruptcyLoanAmountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPortfolioMarginProBankruptcyLoanAmountResponse(val *QueryPortfolioMarginProBankruptcyLoanAmountResponse) *NullableQueryPortfolioMarginProBankruptcyLoanAmountResponse {
	return &NullableQueryPortfolioMarginProBankruptcyLoanAmountResponse{value: val, isSet: true}
}

func (v NullableQueryPortfolioMarginProBankruptcyLoanAmountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPortfolioMarginProBankruptcyLoanAmountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
