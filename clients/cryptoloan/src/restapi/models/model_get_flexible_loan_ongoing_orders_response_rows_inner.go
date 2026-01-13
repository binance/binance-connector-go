/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFlexibleLoanOngoingOrdersResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleLoanOngoingOrdersResponseRowsInner{}

// GetFlexibleLoanOngoingOrdersResponseRowsInner struct for GetFlexibleLoanOngoingOrdersResponseRowsInner
type GetFlexibleLoanOngoingOrdersResponseRowsInner struct {
	LoanCoin             *string `json:"loanCoin,omitempty"`
	TotalDebt            *string `json:"totalDebt,omitempty"`
	CollateralCoin       *string `json:"collateralCoin,omitempty"`
	CollateralAmount     *string `json:"collateralAmount,omitempty"`
	CurrentLTV           *string `json:"currentLTV,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleLoanOngoingOrdersResponseRowsInner GetFlexibleLoanOngoingOrdersResponseRowsInner

// NewGetFlexibleLoanOngoingOrdersResponseRowsInner instantiates a new GetFlexibleLoanOngoingOrdersResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleLoanOngoingOrdersResponseRowsInner() *GetFlexibleLoanOngoingOrdersResponseRowsInner {
	this := GetFlexibleLoanOngoingOrdersResponseRowsInner{}
	return &this
}

// NewGetFlexibleLoanOngoingOrdersResponseRowsInnerWithDefaults instantiates a new GetFlexibleLoanOngoingOrdersResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleLoanOngoingOrdersResponseRowsInnerWithDefaults() *GetFlexibleLoanOngoingOrdersResponseRowsInner {
	this := GetFlexibleLoanOngoingOrdersResponseRowsInner{}
	return &this
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetTotalDebt returns the TotalDebt field value if set, zero value otherwise.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) GetTotalDebt() string {
	if o == nil || common.IsNil(o.TotalDebt) {
		var ret string
		return ret
	}
	return *o.TotalDebt
}

// GetTotalDebtOk returns a tuple with the TotalDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) GetTotalDebtOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalDebt) {
		return nil, false
	}
	return o.TotalDebt, true
}

// HasTotalDebt returns a boolean if a field has been set.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) HasTotalDebt() bool {
	if o != nil && !common.IsNil(o.TotalDebt) {
		return true
	}

	return false
}

// SetTotalDebt gets a reference to the given string and assigns it to the TotalDebt field.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) SetTotalDebt(v string) {
	o.TotalDebt = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetCollateralAmount returns the CollateralAmount field value if set, zero value otherwise.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) GetCollateralAmount() string {
	if o == nil || common.IsNil(o.CollateralAmount) {
		var ret string
		return ret
	}
	return *o.CollateralAmount
}

// GetCollateralAmountOk returns a tuple with the CollateralAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) GetCollateralAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralAmount) {
		return nil, false
	}
	return o.CollateralAmount, true
}

// HasCollateralAmount returns a boolean if a field has been set.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) HasCollateralAmount() bool {
	if o != nil && !common.IsNil(o.CollateralAmount) {
		return true
	}

	return false
}

// SetCollateralAmount gets a reference to the given string and assigns it to the CollateralAmount field.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) SetCollateralAmount(v string) {
	o.CollateralAmount = &v
}

// GetCurrentLTV returns the CurrentLTV field value if set, zero value otherwise.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) GetCurrentLTV() string {
	if o == nil || common.IsNil(o.CurrentLTV) {
		var ret string
		return ret
	}
	return *o.CurrentLTV
}

// GetCurrentLTVOk returns a tuple with the CurrentLTV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) GetCurrentLTVOk() (*string, bool) {
	if o == nil || common.IsNil(o.CurrentLTV) {
		return nil, false
	}
	return o.CurrentLTV, true
}

// HasCurrentLTV returns a boolean if a field has been set.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) HasCurrentLTV() bool {
	if o != nil && !common.IsNil(o.CurrentLTV) {
		return true
	}

	return false
}

// SetCurrentLTV gets a reference to the given string and assigns it to the CurrentLTV field.
func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) SetCurrentLTV(v string) {
	o.CurrentLTV = &v
}

func (o GetFlexibleLoanOngoingOrdersResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleLoanOngoingOrdersResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.TotalDebt) {
		toSerialize["totalDebt"] = o.TotalDebt
	}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !common.IsNil(o.CollateralAmount) {
		toSerialize["collateralAmount"] = o.CollateralAmount
	}
	if !common.IsNil(o.CurrentLTV) {
		toSerialize["currentLTV"] = o.CurrentLTV
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFlexibleLoanOngoingOrdersResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleLoanOngoingOrdersResponseRowsInner := _GetFlexibleLoanOngoingOrdersResponseRowsInner{}

	err = json.Unmarshal(data, &varGetFlexibleLoanOngoingOrdersResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetFlexibleLoanOngoingOrdersResponseRowsInner(varGetFlexibleLoanOngoingOrdersResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "totalDebt")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "collateralAmount")
		delete(additionalProperties, "currentLTV")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleLoanOngoingOrdersResponseRowsInner struct {
	value *GetFlexibleLoanOngoingOrdersResponseRowsInner
	isSet bool
}

func (v NullableGetFlexibleLoanOngoingOrdersResponseRowsInner) Get() *GetFlexibleLoanOngoingOrdersResponseRowsInner {
	return v.value
}

func (v *NullableGetFlexibleLoanOngoingOrdersResponseRowsInner) Set(val *GetFlexibleLoanOngoingOrdersResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleLoanOngoingOrdersResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleLoanOngoingOrdersResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleLoanOngoingOrdersResponseRowsInner(val *GetFlexibleLoanOngoingOrdersResponseRowsInner) *NullableGetFlexibleLoanOngoingOrdersResponseRowsInner {
	return &NullableGetFlexibleLoanOngoingOrdersResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetFlexibleLoanOngoingOrdersResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleLoanOngoingOrdersResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
