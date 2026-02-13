/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetPortfolioMarginProSpanAccountInfoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPortfolioMarginProSpanAccountInfoResponse{}

// GetPortfolioMarginProSpanAccountInfoResponse struct for GetPortfolioMarginProSpanAccountInfoResponse
type GetPortfolioMarginProSpanAccountInfoResponse struct {
	UniMMR               *string                                                           `json:"uniMMR,omitempty"`
	AccountEquity        *string                                                           `json:"accountEquity,omitempty"`
	ActualEquity         *string                                                           `json:"actualEquity,omitempty"`
	AccountMaintMargin   *string                                                           `json:"accountMaintMargin,omitempty"`
	RiskUnitMMList       []GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner `json:"riskUnitMMList,omitempty"`
	MarginMM             *string                                                           `json:"marginMM,omitempty"`
	OtherMM              *string                                                           `json:"otherMM,omitempty"`
	AccountStatus        *string                                                           `json:"accountStatus,omitempty"`
	AccountType          *string                                                           `json:"accountType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPortfolioMarginProSpanAccountInfoResponse GetPortfolioMarginProSpanAccountInfoResponse

// NewGetPortfolioMarginProSpanAccountInfoResponse instantiates a new GetPortfolioMarginProSpanAccountInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPortfolioMarginProSpanAccountInfoResponse() *GetPortfolioMarginProSpanAccountInfoResponse {
	this := GetPortfolioMarginProSpanAccountInfoResponse{}
	return &this
}

// NewGetPortfolioMarginProSpanAccountInfoResponseWithDefaults instantiates a new GetPortfolioMarginProSpanAccountInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPortfolioMarginProSpanAccountInfoResponseWithDefaults() *GetPortfolioMarginProSpanAccountInfoResponse {
	this := GetPortfolioMarginProSpanAccountInfoResponse{}
	return &this
}

// GetUniMMR returns the UniMMR field value if set, zero value otherwise.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetUniMMR() string {
	if o == nil || common.IsNil(o.UniMMR) {
		var ret string
		return ret
	}
	return *o.UniMMR
}

// GetUniMMROk returns a tuple with the UniMMR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetUniMMROk() (*string, bool) {
	if o == nil || common.IsNil(o.UniMMR) {
		return nil, false
	}
	return o.UniMMR, true
}

// HasUniMMR returns a boolean if a field has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasUniMMR() bool {
	if o != nil && !common.IsNil(o.UniMMR) {
		return true
	}

	return false
}

// SetUniMMR gets a reference to the given string and assigns it to the UniMMR field.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetUniMMR(v string) {
	o.UniMMR = &v
}

// GetAccountEquity returns the AccountEquity field value if set, zero value otherwise.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetAccountEquity() string {
	if o == nil || common.IsNil(o.AccountEquity) {
		var ret string
		return ret
	}
	return *o.AccountEquity
}

// GetAccountEquityOk returns a tuple with the AccountEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetAccountEquityOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountEquity) {
		return nil, false
	}
	return o.AccountEquity, true
}

// HasAccountEquity returns a boolean if a field has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasAccountEquity() bool {
	if o != nil && !common.IsNil(o.AccountEquity) {
		return true
	}

	return false
}

// SetAccountEquity gets a reference to the given string and assigns it to the AccountEquity field.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetAccountEquity(v string) {
	o.AccountEquity = &v
}

// GetActualEquity returns the ActualEquity field value if set, zero value otherwise.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetActualEquity() string {
	if o == nil || common.IsNil(o.ActualEquity) {
		var ret string
		return ret
	}
	return *o.ActualEquity
}

// GetActualEquityOk returns a tuple with the ActualEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetActualEquityOk() (*string, bool) {
	if o == nil || common.IsNil(o.ActualEquity) {
		return nil, false
	}
	return o.ActualEquity, true
}

// HasActualEquity returns a boolean if a field has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasActualEquity() bool {
	if o != nil && !common.IsNil(o.ActualEquity) {
		return true
	}

	return false
}

// SetActualEquity gets a reference to the given string and assigns it to the ActualEquity field.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetActualEquity(v string) {
	o.ActualEquity = &v
}

// GetAccountMaintMargin returns the AccountMaintMargin field value if set, zero value otherwise.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetAccountMaintMargin() string {
	if o == nil || common.IsNil(o.AccountMaintMargin) {
		var ret string
		return ret
	}
	return *o.AccountMaintMargin
}

// GetAccountMaintMarginOk returns a tuple with the AccountMaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetAccountMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountMaintMargin) {
		return nil, false
	}
	return o.AccountMaintMargin, true
}

// HasAccountMaintMargin returns a boolean if a field has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasAccountMaintMargin() bool {
	if o != nil && !common.IsNil(o.AccountMaintMargin) {
		return true
	}

	return false
}

// SetAccountMaintMargin gets a reference to the given string and assigns it to the AccountMaintMargin field.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetAccountMaintMargin(v string) {
	o.AccountMaintMargin = &v
}

// GetRiskUnitMMList returns the RiskUnitMMList field value if set, zero value otherwise.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetRiskUnitMMList() []GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner {
	if o == nil || common.IsNil(o.RiskUnitMMList) {
		var ret []GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner
		return ret
	}
	return o.RiskUnitMMList
}

// GetRiskUnitMMListOk returns a tuple with the RiskUnitMMList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetRiskUnitMMListOk() ([]GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner, bool) {
	if o == nil || common.IsNil(o.RiskUnitMMList) {
		return nil, false
	}
	return o.RiskUnitMMList, true
}

// HasRiskUnitMMList returns a boolean if a field has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasRiskUnitMMList() bool {
	if o != nil && !common.IsNil(o.RiskUnitMMList) {
		return true
	}

	return false
}

// SetRiskUnitMMList gets a reference to the given []GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner and assigns it to the RiskUnitMMList field.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetRiskUnitMMList(v []GetPortfolioMarginProSpanAccountInfoResponseRiskUnitMMListInner) {
	o.RiskUnitMMList = v
}

// GetMarginMM returns the MarginMM field value if set, zero value otherwise.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetMarginMM() string {
	if o == nil || common.IsNil(o.MarginMM) {
		var ret string
		return ret
	}
	return *o.MarginMM
}

// GetMarginMMOk returns a tuple with the MarginMM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetMarginMMOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginMM) {
		return nil, false
	}
	return o.MarginMM, true
}

// HasMarginMM returns a boolean if a field has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasMarginMM() bool {
	if o != nil && !common.IsNil(o.MarginMM) {
		return true
	}

	return false
}

// SetMarginMM gets a reference to the given string and assigns it to the MarginMM field.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetMarginMM(v string) {
	o.MarginMM = &v
}

// GetOtherMM returns the OtherMM field value if set, zero value otherwise.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetOtherMM() string {
	if o == nil || common.IsNil(o.OtherMM) {
		var ret string
		return ret
	}
	return *o.OtherMM
}

// GetOtherMMOk returns a tuple with the OtherMM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetOtherMMOk() (*string, bool) {
	if o == nil || common.IsNil(o.OtherMM) {
		return nil, false
	}
	return o.OtherMM, true
}

// HasOtherMM returns a boolean if a field has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasOtherMM() bool {
	if o != nil && !common.IsNil(o.OtherMM) {
		return true
	}

	return false
}

// SetOtherMM gets a reference to the given string and assigns it to the OtherMM field.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetOtherMM(v string) {
	o.OtherMM = &v
}

// GetAccountStatus returns the AccountStatus field value if set, zero value otherwise.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetAccountStatus() string {
	if o == nil || common.IsNil(o.AccountStatus) {
		var ret string
		return ret
	}
	return *o.AccountStatus
}

// GetAccountStatusOk returns a tuple with the AccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetAccountStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountStatus) {
		return nil, false
	}
	return o.AccountStatus, true
}

// HasAccountStatus returns a boolean if a field has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasAccountStatus() bool {
	if o != nil && !common.IsNil(o.AccountStatus) {
		return true
	}

	return false
}

// SetAccountStatus gets a reference to the given string and assigns it to the AccountStatus field.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetAccountStatus(v string) {
	o.AccountStatus = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetAccountType() string {
	if o == nil || common.IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) GetAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) HasAccountType() bool {
	if o != nil && !common.IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *GetPortfolioMarginProSpanAccountInfoResponse) SetAccountType(v string) {
	o.AccountType = &v
}

func (o GetPortfolioMarginProSpanAccountInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPortfolioMarginProSpanAccountInfoResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.RiskUnitMMList) {
		toSerialize["riskUnitMMList"] = o.RiskUnitMMList
	}
	if !common.IsNil(o.MarginMM) {
		toSerialize["marginMM"] = o.MarginMM
	}
	if !common.IsNil(o.OtherMM) {
		toSerialize["otherMM"] = o.OtherMM
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

func (o *GetPortfolioMarginProSpanAccountInfoResponse) UnmarshalJSON(data []byte) (err error) {
	varGetPortfolioMarginProSpanAccountInfoResponse := _GetPortfolioMarginProSpanAccountInfoResponse{}

	err = json.Unmarshal(data, &varGetPortfolioMarginProSpanAccountInfoResponse)

	if err != nil {
		return err
	}

	*o = GetPortfolioMarginProSpanAccountInfoResponse(varGetPortfolioMarginProSpanAccountInfoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uniMMR")
		delete(additionalProperties, "accountEquity")
		delete(additionalProperties, "actualEquity")
		delete(additionalProperties, "accountMaintMargin")
		delete(additionalProperties, "riskUnitMMList")
		delete(additionalProperties, "marginMM")
		delete(additionalProperties, "otherMM")
		delete(additionalProperties, "accountStatus")
		delete(additionalProperties, "accountType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPortfolioMarginProSpanAccountInfoResponse struct {
	value *GetPortfolioMarginProSpanAccountInfoResponse
	isSet bool
}

func (v NullableGetPortfolioMarginProSpanAccountInfoResponse) Get() *GetPortfolioMarginProSpanAccountInfoResponse {
	return v.value
}

func (v *NullableGetPortfolioMarginProSpanAccountInfoResponse) Set(val *GetPortfolioMarginProSpanAccountInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPortfolioMarginProSpanAccountInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPortfolioMarginProSpanAccountInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPortfolioMarginProSpanAccountInfoResponse(val *GetPortfolioMarginProSpanAccountInfoResponse) *NullableGetPortfolioMarginProSpanAccountInfoResponse {
	return &NullableGetPortfolioMarginProSpanAccountInfoResponse{value: val, isSet: true}
}

func (v NullableGetPortfolioMarginProSpanAccountInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPortfolioMarginProSpanAccountInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
