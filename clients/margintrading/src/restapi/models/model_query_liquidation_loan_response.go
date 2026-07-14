/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryLiquidationLoanResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryLiquidationLoanResponse{}

// QueryLiquidationLoanResponse struct for QueryLiquidationLoanResponse
type QueryLiquidationLoanResponse struct {
	// The asset of the liquidation loan (USDC by default)
	Asset *string `json:"asset,omitempty"`
	// Total liquidation loan amount
	Amount *string `json:"amount,omitempty"`
	// Amount that has been repaid
	RepaidAmount *string `json:"repaidAmount,omitempty"`
	// Outstanding amount remaining to be repaid
	RemainingAmount      *string `json:"remainingAmount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryLiquidationLoanResponse QueryLiquidationLoanResponse

// NewQueryLiquidationLoanResponse instantiates a new QueryLiquidationLoanResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryLiquidationLoanResponse() *QueryLiquidationLoanResponse {
	this := QueryLiquidationLoanResponse{}
	return &this
}

// NewQueryLiquidationLoanResponseWithDefaults instantiates a new QueryLiquidationLoanResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryLiquidationLoanResponseWithDefaults() *QueryLiquidationLoanResponse {
	this := QueryLiquidationLoanResponse{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryLiquidationLoanResponse) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiquidationLoanResponse) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryLiquidationLoanResponse) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryLiquidationLoanResponse) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QueryLiquidationLoanResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiquidationLoanResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QueryLiquidationLoanResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QueryLiquidationLoanResponse) SetAmount(v string) {
	o.Amount = &v
}

// GetRepaidAmount returns the RepaidAmount field value if set, zero value otherwise.
func (o *QueryLiquidationLoanResponse) GetRepaidAmount() string {
	if o == nil || common.IsNil(o.RepaidAmount) {
		var ret string
		return ret
	}
	return *o.RepaidAmount
}

// GetRepaidAmountOk returns a tuple with the RepaidAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiquidationLoanResponse) GetRepaidAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.RepaidAmount) {
		return nil, false
	}
	return o.RepaidAmount, true
}

// HasRepaidAmount returns a boolean if a field has been set.
func (o *QueryLiquidationLoanResponse) HasRepaidAmount() bool {
	if o != nil && !common.IsNil(o.RepaidAmount) {
		return true
	}

	return false
}

// SetRepaidAmount gets a reference to the given string and assigns it to the RepaidAmount field.
func (o *QueryLiquidationLoanResponse) SetRepaidAmount(v string) {
	o.RepaidAmount = &v
}

// GetRemainingAmount returns the RemainingAmount field value if set, zero value otherwise.
func (o *QueryLiquidationLoanResponse) GetRemainingAmount() string {
	if o == nil || common.IsNil(o.RemainingAmount) {
		var ret string
		return ret
	}
	return *o.RemainingAmount
}

// GetRemainingAmountOk returns a tuple with the RemainingAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLiquidationLoanResponse) GetRemainingAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.RemainingAmount) {
		return nil, false
	}
	return o.RemainingAmount, true
}

// HasRemainingAmount returns a boolean if a field has been set.
func (o *QueryLiquidationLoanResponse) HasRemainingAmount() bool {
	if o != nil && !common.IsNil(o.RemainingAmount) {
		return true
	}

	return false
}

// SetRemainingAmount gets a reference to the given string and assigns it to the RemainingAmount field.
func (o *QueryLiquidationLoanResponse) SetRemainingAmount(v string) {
	o.RemainingAmount = &v
}

func (o QueryLiquidationLoanResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryLiquidationLoanResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.RepaidAmount) {
		toSerialize["repaidAmount"] = o.RepaidAmount
	}
	if !common.IsNil(o.RemainingAmount) {
		toSerialize["remainingAmount"] = o.RemainingAmount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryLiquidationLoanResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryLiquidationLoanResponse := _QueryLiquidationLoanResponse{}

	err = json.Unmarshal(data, &varQueryLiquidationLoanResponse)

	if err != nil {
		return err
	}

	*o = QueryLiquidationLoanResponse(varQueryLiquidationLoanResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "repaidAmount")
		delete(additionalProperties, "remainingAmount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryLiquidationLoanResponse struct {
	value *QueryLiquidationLoanResponse
	isSet bool
}

func (v NullableQueryLiquidationLoanResponse) Get() *QueryLiquidationLoanResponse {
	return v.value
}

func (v *NullableQueryLiquidationLoanResponse) Set(val *QueryLiquidationLoanResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryLiquidationLoanResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryLiquidationLoanResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryLiquidationLoanResponse(val *QueryLiquidationLoanResponse) *NullableQueryLiquidationLoanResponse {
	return &NullableQueryLiquidationLoanResponse{value: val, isSet: true}
}

func (v NullableQueryLiquidationLoanResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryLiquidationLoanResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
