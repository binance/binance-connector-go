/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetVIPLoanOngoingOrdersResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetVIPLoanOngoingOrdersResponseRowsInner{}

// GetVIPLoanOngoingOrdersResponseRowsInner struct for GetVIPLoanOngoingOrdersResponseRowsInner
type GetVIPLoanOngoingOrdersResponseRowsInner struct {
	OrderId                          *int64  `json:"orderId,omitempty"`
	LoanCoin                         *string `json:"loanCoin,omitempty"`
	TotalDebt                        *string `json:"totalDebt,omitempty"`
	LoanRate                         *string `json:"loanRate,omitempty"`
	ResidualInterest                 *string `json:"residualInterest,omitempty"`
	CollateralAccountId              *string `json:"collateralAccountId,omitempty"`
	CollateralCoin                   *string `json:"collateralCoin,omitempty"`
	TotalCollateralValueAfterHaircut *string `json:"totalCollateralValueAfterHaircut,omitempty"`
	LockedCollateralValue            *string `json:"lockedCollateralValue,omitempty"`
	CurrentLTV                       *string `json:"currentLTV,omitempty"`
	ExpirationTime                   *int64  `json:"expirationTime,omitempty"`
	LoanDate                         *string `json:"loanDate,omitempty"`
	LoanTerm                         *string `json:"loanTerm,omitempty"`
	AdditionalProperties             map[string]interface{}
}

type _GetVIPLoanOngoingOrdersResponseRowsInner GetVIPLoanOngoingOrdersResponseRowsInner

// NewGetVIPLoanOngoingOrdersResponseRowsInner instantiates a new GetVIPLoanOngoingOrdersResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVIPLoanOngoingOrdersResponseRowsInner() *GetVIPLoanOngoingOrdersResponseRowsInner {
	this := GetVIPLoanOngoingOrdersResponseRowsInner{}
	return &this
}

// NewGetVIPLoanOngoingOrdersResponseRowsInnerWithDefaults instantiates a new GetVIPLoanOngoingOrdersResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVIPLoanOngoingOrdersResponseRowsInnerWithDefaults() *GetVIPLoanOngoingOrdersResponseRowsInner {
	this := GetVIPLoanOngoingOrdersResponseRowsInner{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetTotalDebt returns the TotalDebt field value if set, zero value otherwise.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetTotalDebt() string {
	if o == nil || common.IsNil(o.TotalDebt) {
		var ret string
		return ret
	}
	return *o.TotalDebt
}

// GetTotalDebtOk returns a tuple with the TotalDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetTotalDebtOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalDebt) {
		return nil, false
	}
	return o.TotalDebt, true
}

// HasTotalDebt returns a boolean if a field has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasTotalDebt() bool {
	if o != nil && !common.IsNil(o.TotalDebt) {
		return true
	}

	return false
}

// SetTotalDebt gets a reference to the given string and assigns it to the TotalDebt field.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetTotalDebt(v string) {
	o.TotalDebt = &v
}

// GetLoanRate returns the LoanRate field value if set, zero value otherwise.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLoanRate() string {
	if o == nil || common.IsNil(o.LoanRate) {
		var ret string
		return ret
	}
	return *o.LoanRate
}

// GetLoanRateOk returns a tuple with the LoanRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLoanRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanRate) {
		return nil, false
	}
	return o.LoanRate, true
}

// HasLoanRate returns a boolean if a field has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasLoanRate() bool {
	if o != nil && !common.IsNil(o.LoanRate) {
		return true
	}

	return false
}

// SetLoanRate gets a reference to the given string and assigns it to the LoanRate field.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetLoanRate(v string) {
	o.LoanRate = &v
}

// GetResidualInterest returns the ResidualInterest field value if set, zero value otherwise.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetResidualInterest() string {
	if o == nil || common.IsNil(o.ResidualInterest) {
		var ret string
		return ret
	}
	return *o.ResidualInterest
}

// GetResidualInterestOk returns a tuple with the ResidualInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetResidualInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.ResidualInterest) {
		return nil, false
	}
	return o.ResidualInterest, true
}

// HasResidualInterest returns a boolean if a field has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasResidualInterest() bool {
	if o != nil && !common.IsNil(o.ResidualInterest) {
		return true
	}

	return false
}

// SetResidualInterest gets a reference to the given string and assigns it to the ResidualInterest field.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetResidualInterest(v string) {
	o.ResidualInterest = &v
}

// GetCollateralAccountId returns the CollateralAccountId field value if set, zero value otherwise.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetCollateralAccountId() string {
	if o == nil || common.IsNil(o.CollateralAccountId) {
		var ret string
		return ret
	}
	return *o.CollateralAccountId
}

// GetCollateralAccountIdOk returns a tuple with the CollateralAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetCollateralAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralAccountId) {
		return nil, false
	}
	return o.CollateralAccountId, true
}

// HasCollateralAccountId returns a boolean if a field has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasCollateralAccountId() bool {
	if o != nil && !common.IsNil(o.CollateralAccountId) {
		return true
	}

	return false
}

// SetCollateralAccountId gets a reference to the given string and assigns it to the CollateralAccountId field.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetCollateralAccountId(v string) {
	o.CollateralAccountId = &v
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetCollateralCoin() string {
	if o == nil || common.IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetCollateralCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasCollateralCoin() bool {
	if o != nil && !common.IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetTotalCollateralValueAfterHaircut returns the TotalCollateralValueAfterHaircut field value if set, zero value otherwise.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetTotalCollateralValueAfterHaircut() string {
	if o == nil || common.IsNil(o.TotalCollateralValueAfterHaircut) {
		var ret string
		return ret
	}
	return *o.TotalCollateralValueAfterHaircut
}

// GetTotalCollateralValueAfterHaircutOk returns a tuple with the TotalCollateralValueAfterHaircut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetTotalCollateralValueAfterHaircutOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalCollateralValueAfterHaircut) {
		return nil, false
	}
	return o.TotalCollateralValueAfterHaircut, true
}

// HasTotalCollateralValueAfterHaircut returns a boolean if a field has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasTotalCollateralValueAfterHaircut() bool {
	if o != nil && !common.IsNil(o.TotalCollateralValueAfterHaircut) {
		return true
	}

	return false
}

// SetTotalCollateralValueAfterHaircut gets a reference to the given string and assigns it to the TotalCollateralValueAfterHaircut field.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetTotalCollateralValueAfterHaircut(v string) {
	o.TotalCollateralValueAfterHaircut = &v
}

// GetLockedCollateralValue returns the LockedCollateralValue field value if set, zero value otherwise.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLockedCollateralValue() string {
	if o == nil || common.IsNil(o.LockedCollateralValue) {
		var ret string
		return ret
	}
	return *o.LockedCollateralValue
}

// GetLockedCollateralValueOk returns a tuple with the LockedCollateralValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLockedCollateralValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.LockedCollateralValue) {
		return nil, false
	}
	return o.LockedCollateralValue, true
}

// HasLockedCollateralValue returns a boolean if a field has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasLockedCollateralValue() bool {
	if o != nil && !common.IsNil(o.LockedCollateralValue) {
		return true
	}

	return false
}

// SetLockedCollateralValue gets a reference to the given string and assigns it to the LockedCollateralValue field.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetLockedCollateralValue(v string) {
	o.LockedCollateralValue = &v
}

// GetCurrentLTV returns the CurrentLTV field value if set, zero value otherwise.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetCurrentLTV() string {
	if o == nil || common.IsNil(o.CurrentLTV) {
		var ret string
		return ret
	}
	return *o.CurrentLTV
}

// GetCurrentLTVOk returns a tuple with the CurrentLTV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetCurrentLTVOk() (*string, bool) {
	if o == nil || common.IsNil(o.CurrentLTV) {
		return nil, false
	}
	return o.CurrentLTV, true
}

// HasCurrentLTV returns a boolean if a field has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasCurrentLTV() bool {
	if o != nil && !common.IsNil(o.CurrentLTV) {
		return true
	}

	return false
}

// SetCurrentLTV gets a reference to the given string and assigns it to the CurrentLTV field.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetCurrentLTV(v string) {
	o.CurrentLTV = &v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetExpirationTime() int64 {
	if o == nil || common.IsNil(o.ExpirationTime) {
		var ret int64
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetExpirationTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExpirationTime) {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasExpirationTime() bool {
	if o != nil && !common.IsNil(o.ExpirationTime) {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given int64 and assigns it to the ExpirationTime field.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetExpirationTime(v int64) {
	o.ExpirationTime = &v
}

// GetLoanDate returns the LoanDate field value if set, zero value otherwise.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLoanDate() string {
	if o == nil || common.IsNil(o.LoanDate) {
		var ret string
		return ret
	}
	return *o.LoanDate
}

// GetLoanDateOk returns a tuple with the LoanDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLoanDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanDate) {
		return nil, false
	}
	return o.LoanDate, true
}

// HasLoanDate returns a boolean if a field has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasLoanDate() bool {
	if o != nil && !common.IsNil(o.LoanDate) {
		return true
	}

	return false
}

// SetLoanDate gets a reference to the given string and assigns it to the LoanDate field.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetLoanDate(v string) {
	o.LoanDate = &v
}

// GetLoanTerm returns the LoanTerm field value if set, zero value otherwise.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLoanTerm() string {
	if o == nil || common.IsNil(o.LoanTerm) {
		var ret string
		return ret
	}
	return *o.LoanTerm
}

// GetLoanTermOk returns a tuple with the LoanTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLoanTermOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanTerm) {
		return nil, false
	}
	return o.LoanTerm, true
}

// HasLoanTerm returns a boolean if a field has been set.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasLoanTerm() bool {
	if o != nil && !common.IsNil(o.LoanTerm) {
		return true
	}

	return false
}

// SetLoanTerm gets a reference to the given string and assigns it to the LoanTerm field.
func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetLoanTerm(v string) {
	o.LoanTerm = &v
}

func (o GetVIPLoanOngoingOrdersResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVIPLoanOngoingOrdersResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.TotalDebt) {
		toSerialize["totalDebt"] = o.TotalDebt
	}
	if !common.IsNil(o.LoanRate) {
		toSerialize["loanRate"] = o.LoanRate
	}
	if !common.IsNil(o.ResidualInterest) {
		toSerialize["residualInterest"] = o.ResidualInterest
	}
	if !common.IsNil(o.CollateralAccountId) {
		toSerialize["collateralAccountId"] = o.CollateralAccountId
	}
	if !common.IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !common.IsNil(o.TotalCollateralValueAfterHaircut) {
		toSerialize["totalCollateralValueAfterHaircut"] = o.TotalCollateralValueAfterHaircut
	}
	if !common.IsNil(o.LockedCollateralValue) {
		toSerialize["lockedCollateralValue"] = o.LockedCollateralValue
	}
	if !common.IsNil(o.CurrentLTV) {
		toSerialize["currentLTV"] = o.CurrentLTV
	}
	if !common.IsNil(o.ExpirationTime) {
		toSerialize["expirationTime"] = o.ExpirationTime
	}
	if !common.IsNil(o.LoanDate) {
		toSerialize["loanDate"] = o.LoanDate
	}
	if !common.IsNil(o.LoanTerm) {
		toSerialize["loanTerm"] = o.LoanTerm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetVIPLoanOngoingOrdersResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetVIPLoanOngoingOrdersResponseRowsInner := _GetVIPLoanOngoingOrdersResponseRowsInner{}

	err = json.Unmarshal(data, &varGetVIPLoanOngoingOrdersResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetVIPLoanOngoingOrdersResponseRowsInner(varGetVIPLoanOngoingOrdersResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "totalDebt")
		delete(additionalProperties, "loanRate")
		delete(additionalProperties, "residualInterest")
		delete(additionalProperties, "collateralAccountId")
		delete(additionalProperties, "collateralCoin")
		delete(additionalProperties, "totalCollateralValueAfterHaircut")
		delete(additionalProperties, "lockedCollateralValue")
		delete(additionalProperties, "currentLTV")
		delete(additionalProperties, "expirationTime")
		delete(additionalProperties, "loanDate")
		delete(additionalProperties, "loanTerm")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetVIPLoanOngoingOrdersResponseRowsInner struct {
	value *GetVIPLoanOngoingOrdersResponseRowsInner
	isSet bool
}

func (v NullableGetVIPLoanOngoingOrdersResponseRowsInner) Get() *GetVIPLoanOngoingOrdersResponseRowsInner {
	return v.value
}

func (v *NullableGetVIPLoanOngoingOrdersResponseRowsInner) Set(val *GetVIPLoanOngoingOrdersResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVIPLoanOngoingOrdersResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVIPLoanOngoingOrdersResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVIPLoanOngoingOrdersResponseRowsInner(val *GetVIPLoanOngoingOrdersResponseRowsInner) *NullableGetVIPLoanOngoingOrdersResponseRowsInner {
	return &NullableGetVIPLoanOngoingOrdersResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetVIPLoanOngoingOrdersResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVIPLoanOngoingOrdersResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
