/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountInformationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountInformationResponse{}

// AccountInformationResponse struct for AccountInformationResponse
type AccountInformationResponse struct {
	UniMMR                   *string `json:"uniMMR,omitempty"`
	AccountEquity            *string `json:"accountEquity,omitempty"`
	ActualEquity             *string `json:"actualEquity,omitempty"`
	AccountInitialMargin     *string `json:"accountInitialMargin,omitempty"`
	AccountMaintMargin       *string `json:"accountMaintMargin,omitempty"`
	AccountStatus            *string `json:"accountStatus,omitempty"`
	VirtualMaxWithdrawAmount *string `json:"virtualMaxWithdrawAmount,omitempty"`
	TotalAvailableBalance    *string `json:"totalAvailableBalance,omitempty"`
	TotalMarginOpenLoss      *string `json:"totalMarginOpenLoss,omitempty"`
	UpdateTime               *int64  `json:"updateTime,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _AccountInformationResponse AccountInformationResponse

// NewAccountInformationResponse instantiates a new AccountInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInformationResponse() *AccountInformationResponse {
	this := AccountInformationResponse{}
	return &this
}

// NewAccountInformationResponseWithDefaults instantiates a new AccountInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInformationResponseWithDefaults() *AccountInformationResponse {
	this := AccountInformationResponse{}
	return &this
}

// GetUniMMR returns the UniMMR field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetUniMMR() string {
	if o == nil || common.IsNil(o.UniMMR) {
		var ret string
		return ret
	}
	return *o.UniMMR
}

// GetUniMMROk returns a tuple with the UniMMR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetUniMMROk() (*string, bool) {
	if o == nil || common.IsNil(o.UniMMR) {
		return nil, false
	}
	return o.UniMMR, true
}

// HasUniMMR returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasUniMMR() bool {
	if o != nil && !common.IsNil(o.UniMMR) {
		return true
	}

	return false
}

// SetUniMMR gets a reference to the given string and assigns it to the UniMMR field.
func (o *AccountInformationResponse) SetUniMMR(v string) {
	o.UniMMR = &v
}

// GetAccountEquity returns the AccountEquity field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetAccountEquity() string {
	if o == nil || common.IsNil(o.AccountEquity) {
		var ret string
		return ret
	}
	return *o.AccountEquity
}

// GetAccountEquityOk returns a tuple with the AccountEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetAccountEquityOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountEquity) {
		return nil, false
	}
	return o.AccountEquity, true
}

// HasAccountEquity returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasAccountEquity() bool {
	if o != nil && !common.IsNil(o.AccountEquity) {
		return true
	}

	return false
}

// SetAccountEquity gets a reference to the given string and assigns it to the AccountEquity field.
func (o *AccountInformationResponse) SetAccountEquity(v string) {
	o.AccountEquity = &v
}

// GetActualEquity returns the ActualEquity field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetActualEquity() string {
	if o == nil || common.IsNil(o.ActualEquity) {
		var ret string
		return ret
	}
	return *o.ActualEquity
}

// GetActualEquityOk returns a tuple with the ActualEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetActualEquityOk() (*string, bool) {
	if o == nil || common.IsNil(o.ActualEquity) {
		return nil, false
	}
	return o.ActualEquity, true
}

// HasActualEquity returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasActualEquity() bool {
	if o != nil && !common.IsNil(o.ActualEquity) {
		return true
	}

	return false
}

// SetActualEquity gets a reference to the given string and assigns it to the ActualEquity field.
func (o *AccountInformationResponse) SetActualEquity(v string) {
	o.ActualEquity = &v
}

// GetAccountInitialMargin returns the AccountInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetAccountInitialMargin() string {
	if o == nil || common.IsNil(o.AccountInitialMargin) {
		var ret string
		return ret
	}
	return *o.AccountInitialMargin
}

// GetAccountInitialMarginOk returns a tuple with the AccountInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetAccountInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountInitialMargin) {
		return nil, false
	}
	return o.AccountInitialMargin, true
}

// HasAccountInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasAccountInitialMargin() bool {
	if o != nil && !common.IsNil(o.AccountInitialMargin) {
		return true
	}

	return false
}

// SetAccountInitialMargin gets a reference to the given string and assigns it to the AccountInitialMargin field.
func (o *AccountInformationResponse) SetAccountInitialMargin(v string) {
	o.AccountInitialMargin = &v
}

// GetAccountMaintMargin returns the AccountMaintMargin field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetAccountMaintMargin() string {
	if o == nil || common.IsNil(o.AccountMaintMargin) {
		var ret string
		return ret
	}
	return *o.AccountMaintMargin
}

// GetAccountMaintMarginOk returns a tuple with the AccountMaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetAccountMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountMaintMargin) {
		return nil, false
	}
	return o.AccountMaintMargin, true
}

// HasAccountMaintMargin returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasAccountMaintMargin() bool {
	if o != nil && !common.IsNil(o.AccountMaintMargin) {
		return true
	}

	return false
}

// SetAccountMaintMargin gets a reference to the given string and assigns it to the AccountMaintMargin field.
func (o *AccountInformationResponse) SetAccountMaintMargin(v string) {
	o.AccountMaintMargin = &v
}

// GetAccountStatus returns the AccountStatus field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetAccountStatus() string {
	if o == nil || common.IsNil(o.AccountStatus) {
		var ret string
		return ret
	}
	return *o.AccountStatus
}

// GetAccountStatusOk returns a tuple with the AccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetAccountStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountStatus) {
		return nil, false
	}
	return o.AccountStatus, true
}

// HasAccountStatus returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasAccountStatus() bool {
	if o != nil && !common.IsNil(o.AccountStatus) {
		return true
	}

	return false
}

// SetAccountStatus gets a reference to the given string and assigns it to the AccountStatus field.
func (o *AccountInformationResponse) SetAccountStatus(v string) {
	o.AccountStatus = &v
}

// GetVirtualMaxWithdrawAmount returns the VirtualMaxWithdrawAmount field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetVirtualMaxWithdrawAmount() string {
	if o == nil || common.IsNil(o.VirtualMaxWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.VirtualMaxWithdrawAmount
}

// GetVirtualMaxWithdrawAmountOk returns a tuple with the VirtualMaxWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetVirtualMaxWithdrawAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.VirtualMaxWithdrawAmount) {
		return nil, false
	}
	return o.VirtualMaxWithdrawAmount, true
}

// HasVirtualMaxWithdrawAmount returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasVirtualMaxWithdrawAmount() bool {
	if o != nil && !common.IsNil(o.VirtualMaxWithdrawAmount) {
		return true
	}

	return false
}

// SetVirtualMaxWithdrawAmount gets a reference to the given string and assigns it to the VirtualMaxWithdrawAmount field.
func (o *AccountInformationResponse) SetVirtualMaxWithdrawAmount(v string) {
	o.VirtualMaxWithdrawAmount = &v
}

// GetTotalAvailableBalance returns the TotalAvailableBalance field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetTotalAvailableBalance() string {
	if o == nil || common.IsNil(o.TotalAvailableBalance) {
		var ret string
		return ret
	}
	return *o.TotalAvailableBalance
}

// GetTotalAvailableBalanceOk returns a tuple with the TotalAvailableBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetTotalAvailableBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAvailableBalance) {
		return nil, false
	}
	return o.TotalAvailableBalance, true
}

// HasTotalAvailableBalance returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasTotalAvailableBalance() bool {
	if o != nil && !common.IsNil(o.TotalAvailableBalance) {
		return true
	}

	return false
}

// SetTotalAvailableBalance gets a reference to the given string and assigns it to the TotalAvailableBalance field.
func (o *AccountInformationResponse) SetTotalAvailableBalance(v string) {
	o.TotalAvailableBalance = &v
}

// GetTotalMarginOpenLoss returns the TotalMarginOpenLoss field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetTotalMarginOpenLoss() string {
	if o == nil || common.IsNil(o.TotalMarginOpenLoss) {
		var ret string
		return ret
	}
	return *o.TotalMarginOpenLoss
}

// GetTotalMarginOpenLossOk returns a tuple with the TotalMarginOpenLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetTotalMarginOpenLossOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMarginOpenLoss) {
		return nil, false
	}
	return o.TotalMarginOpenLoss, true
}

// HasTotalMarginOpenLoss returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasTotalMarginOpenLoss() bool {
	if o != nil && !common.IsNil(o.TotalMarginOpenLoss) {
		return true
	}

	return false
}

// SetTotalMarginOpenLoss gets a reference to the given string and assigns it to the TotalMarginOpenLoss field.
func (o *AccountInformationResponse) SetTotalMarginOpenLoss(v string) {
	o.TotalMarginOpenLoss = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AccountInformationResponse) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AccountInformationResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AccountInformationResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o AccountInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.UniMMR) {
		toSerialize["uniMMR"] = o.UniMMR
	}
	if !common.IsNil(o.AccountEquity) {
		toSerialize["accountEquity"] = o.AccountEquity
	}
	if !common.IsNil(o.ActualEquity) {
		toSerialize["actualEquity"] = o.ActualEquity
	}
	if !common.IsNil(o.AccountInitialMargin) {
		toSerialize["accountInitialMargin"] = o.AccountInitialMargin
	}
	if !common.IsNil(o.AccountMaintMargin) {
		toSerialize["accountMaintMargin"] = o.AccountMaintMargin
	}
	if !common.IsNil(o.AccountStatus) {
		toSerialize["accountStatus"] = o.AccountStatus
	}
	if !common.IsNil(o.VirtualMaxWithdrawAmount) {
		toSerialize["virtualMaxWithdrawAmount"] = o.VirtualMaxWithdrawAmount
	}
	if !common.IsNil(o.TotalAvailableBalance) {
		toSerialize["totalAvailableBalance"] = o.TotalAvailableBalance
	}
	if !common.IsNil(o.TotalMarginOpenLoss) {
		toSerialize["totalMarginOpenLoss"] = o.TotalMarginOpenLoss
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountInformationResponse) UnmarshalJSON(data []byte) (err error) {
	varAccountInformationResponse := _AccountInformationResponse{}

	err = json.Unmarshal(data, &varAccountInformationResponse)

	if err != nil {
		return err
	}

	*o = AccountInformationResponse(varAccountInformationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uniMMR")
		delete(additionalProperties, "accountEquity")
		delete(additionalProperties, "actualEquity")
		delete(additionalProperties, "accountInitialMargin")
		delete(additionalProperties, "accountMaintMargin")
		delete(additionalProperties, "accountStatus")
		delete(additionalProperties, "virtualMaxWithdrawAmount")
		delete(additionalProperties, "totalAvailableBalance")
		delete(additionalProperties, "totalMarginOpenLoss")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountInformationResponse struct {
	value *AccountInformationResponse
	isSet bool
}

func (v NullableAccountInformationResponse) Get() *AccountInformationResponse {
	return v.value
}

func (v *NullableAccountInformationResponse) Set(val *AccountInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInformationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInformationResponse(val *AccountInformationResponse) *NullableAccountInformationResponse {
	return &NullableAccountInformationResponse{value: val, isSet: true}
}

func (v NullableAccountInformationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
