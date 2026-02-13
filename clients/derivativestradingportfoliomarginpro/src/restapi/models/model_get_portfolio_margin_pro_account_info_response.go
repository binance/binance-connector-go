/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetPortfolioMarginProAccountInfoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPortfolioMarginProAccountInfoResponse{}

// GetPortfolioMarginProAccountInfoResponse struct for GetPortfolioMarginProAccountInfoResponse
type GetPortfolioMarginProAccountInfoResponse struct {
	UniMMR                *string `json:"uniMMR,omitempty"`
	AccountEquity         *string `json:"accountEquity,omitempty"`
	ActualEquity          *string `json:"actualEquity,omitempty"`
	AccountMaintMargin    *string `json:"accountMaintMargin,omitempty"`
	AccountInitialMargin  *string `json:"accountInitialMargin,omitempty"`
	TotalAvailableBalance *string `json:"totalAvailableBalance,omitempty"`
	AccountStatus         *string `json:"accountStatus,omitempty"`
	AccountType           *string `json:"accountType,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _GetPortfolioMarginProAccountInfoResponse GetPortfolioMarginProAccountInfoResponse

// NewGetPortfolioMarginProAccountInfoResponse instantiates a new GetPortfolioMarginProAccountInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPortfolioMarginProAccountInfoResponse() *GetPortfolioMarginProAccountInfoResponse {
	this := GetPortfolioMarginProAccountInfoResponse{}
	return &this
}

// NewGetPortfolioMarginProAccountInfoResponseWithDefaults instantiates a new GetPortfolioMarginProAccountInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPortfolioMarginProAccountInfoResponseWithDefaults() *GetPortfolioMarginProAccountInfoResponse {
	this := GetPortfolioMarginProAccountInfoResponse{}
	return &this
}

// GetUniMMR returns the UniMMR field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountInfoResponse) GetUniMMR() string {
	if o == nil || common.IsNil(o.UniMMR) {
		var ret string
		return ret
	}
	return *o.UniMMR
}

// GetUniMMROk returns a tuple with the UniMMR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountInfoResponse) GetUniMMROk() (*string, bool) {
	if o == nil || common.IsNil(o.UniMMR) {
		return nil, false
	}
	return o.UniMMR, true
}

// HasUniMMR returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountInfoResponse) HasUniMMR() bool {
	if o != nil && !common.IsNil(o.UniMMR) {
		return true
	}

	return false
}

// SetUniMMR gets a reference to the given string and assigns it to the UniMMR field.
func (o *GetPortfolioMarginProAccountInfoResponse) SetUniMMR(v string) {
	o.UniMMR = &v
}

// GetAccountEquity returns the AccountEquity field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountEquity() string {
	if o == nil || common.IsNil(o.AccountEquity) {
		var ret string
		return ret
	}
	return *o.AccountEquity
}

// GetAccountEquityOk returns a tuple with the AccountEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountEquityOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountEquity) {
		return nil, false
	}
	return o.AccountEquity, true
}

// HasAccountEquity returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountInfoResponse) HasAccountEquity() bool {
	if o != nil && !common.IsNil(o.AccountEquity) {
		return true
	}

	return false
}

// SetAccountEquity gets a reference to the given string and assigns it to the AccountEquity field.
func (o *GetPortfolioMarginProAccountInfoResponse) SetAccountEquity(v string) {
	o.AccountEquity = &v
}

// GetActualEquity returns the ActualEquity field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountInfoResponse) GetActualEquity() string {
	if o == nil || common.IsNil(o.ActualEquity) {
		var ret string
		return ret
	}
	return *o.ActualEquity
}

// GetActualEquityOk returns a tuple with the ActualEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountInfoResponse) GetActualEquityOk() (*string, bool) {
	if o == nil || common.IsNil(o.ActualEquity) {
		return nil, false
	}
	return o.ActualEquity, true
}

// HasActualEquity returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountInfoResponse) HasActualEquity() bool {
	if o != nil && !common.IsNil(o.ActualEquity) {
		return true
	}

	return false
}

// SetActualEquity gets a reference to the given string and assigns it to the ActualEquity field.
func (o *GetPortfolioMarginProAccountInfoResponse) SetActualEquity(v string) {
	o.ActualEquity = &v
}

// GetAccountMaintMargin returns the AccountMaintMargin field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountMaintMargin() string {
	if o == nil || common.IsNil(o.AccountMaintMargin) {
		var ret string
		return ret
	}
	return *o.AccountMaintMargin
}

// GetAccountMaintMarginOk returns a tuple with the AccountMaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountMaintMargin) {
		return nil, false
	}
	return o.AccountMaintMargin, true
}

// HasAccountMaintMargin returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountInfoResponse) HasAccountMaintMargin() bool {
	if o != nil && !common.IsNil(o.AccountMaintMargin) {
		return true
	}

	return false
}

// SetAccountMaintMargin gets a reference to the given string and assigns it to the AccountMaintMargin field.
func (o *GetPortfolioMarginProAccountInfoResponse) SetAccountMaintMargin(v string) {
	o.AccountMaintMargin = &v
}

// GetAccountInitialMargin returns the AccountInitialMargin field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountInitialMargin() string {
	if o == nil || common.IsNil(o.AccountInitialMargin) {
		var ret string
		return ret
	}
	return *o.AccountInitialMargin
}

// GetAccountInitialMarginOk returns a tuple with the AccountInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountInitialMargin) {
		return nil, false
	}
	return o.AccountInitialMargin, true
}

// HasAccountInitialMargin returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountInfoResponse) HasAccountInitialMargin() bool {
	if o != nil && !common.IsNil(o.AccountInitialMargin) {
		return true
	}

	return false
}

// SetAccountInitialMargin gets a reference to the given string and assigns it to the AccountInitialMargin field.
func (o *GetPortfolioMarginProAccountInfoResponse) SetAccountInitialMargin(v string) {
	o.AccountInitialMargin = &v
}

// GetTotalAvailableBalance returns the TotalAvailableBalance field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountInfoResponse) GetTotalAvailableBalance() string {
	if o == nil || common.IsNil(o.TotalAvailableBalance) {
		var ret string
		return ret
	}
	return *o.TotalAvailableBalance
}

// GetTotalAvailableBalanceOk returns a tuple with the TotalAvailableBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountInfoResponse) GetTotalAvailableBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAvailableBalance) {
		return nil, false
	}
	return o.TotalAvailableBalance, true
}

// HasTotalAvailableBalance returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountInfoResponse) HasTotalAvailableBalance() bool {
	if o != nil && !common.IsNil(o.TotalAvailableBalance) {
		return true
	}

	return false
}

// SetTotalAvailableBalance gets a reference to the given string and assigns it to the TotalAvailableBalance field.
func (o *GetPortfolioMarginProAccountInfoResponse) SetTotalAvailableBalance(v string) {
	o.TotalAvailableBalance = &v
}

// GetAccountStatus returns the AccountStatus field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountStatus() string {
	if o == nil || common.IsNil(o.AccountStatus) {
		var ret string
		return ret
	}
	return *o.AccountStatus
}

// GetAccountStatusOk returns a tuple with the AccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountStatus) {
		return nil, false
	}
	return o.AccountStatus, true
}

// HasAccountStatus returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountInfoResponse) HasAccountStatus() bool {
	if o != nil && !common.IsNil(o.AccountStatus) {
		return true
	}

	return false
}

// SetAccountStatus gets a reference to the given string and assigns it to the AccountStatus field.
func (o *GetPortfolioMarginProAccountInfoResponse) SetAccountStatus(v string) {
	o.AccountStatus = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountType() string {
	if o == nil || common.IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountInfoResponse) HasAccountType() bool {
	if o != nil && !common.IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *GetPortfolioMarginProAccountInfoResponse) SetAccountType(v string) {
	o.AccountType = &v
}

func (o GetPortfolioMarginProAccountInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPortfolioMarginProAccountInfoResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.AccountMaintMargin) {
		toSerialize["accountMaintMargin"] = o.AccountMaintMargin
	}
	if !common.IsNil(o.AccountInitialMargin) {
		toSerialize["accountInitialMargin"] = o.AccountInitialMargin
	}
	if !common.IsNil(o.TotalAvailableBalance) {
		toSerialize["totalAvailableBalance"] = o.TotalAvailableBalance
	}
	if !common.IsNil(o.AccountStatus) {
		toSerialize["accountStatus"] = o.AccountStatus
	}
	if !common.IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPortfolioMarginProAccountInfoResponse) UnmarshalJSON(data []byte) (err error) {
	varGetPortfolioMarginProAccountInfoResponse := _GetPortfolioMarginProAccountInfoResponse{}

	err = json.Unmarshal(data, &varGetPortfolioMarginProAccountInfoResponse)

	if err != nil {
		return err
	}

	*o = GetPortfolioMarginProAccountInfoResponse(varGetPortfolioMarginProAccountInfoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uniMMR")
		delete(additionalProperties, "accountEquity")
		delete(additionalProperties, "actualEquity")
		delete(additionalProperties, "accountMaintMargin")
		delete(additionalProperties, "accountInitialMargin")
		delete(additionalProperties, "totalAvailableBalance")
		delete(additionalProperties, "accountStatus")
		delete(additionalProperties, "accountType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPortfolioMarginProAccountInfoResponse struct {
	value *GetPortfolioMarginProAccountInfoResponse
	isSet bool
}

func (v NullableGetPortfolioMarginProAccountInfoResponse) Get() *GetPortfolioMarginProAccountInfoResponse {
	return v.value
}

func (v *NullableGetPortfolioMarginProAccountInfoResponse) Set(val *GetPortfolioMarginProAccountInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPortfolioMarginProAccountInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPortfolioMarginProAccountInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPortfolioMarginProAccountInfoResponse(val *GetPortfolioMarginProAccountInfoResponse) *NullableGetPortfolioMarginProAccountInfoResponse {
	return &NullableGetPortfolioMarginProAccountInfoResponse{value: val, isSet: true}
}

func (v NullableGetPortfolioMarginProAccountInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPortfolioMarginProAccountInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
