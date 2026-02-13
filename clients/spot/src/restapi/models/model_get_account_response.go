/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetAccountResponse{}

// GetAccountResponse struct for GetAccountResponse
type GetAccountResponse struct {
	MakerCommission            *int64                             `json:"makerCommission,omitempty"`
	TakerCommission            *int64                             `json:"takerCommission,omitempty"`
	BuyerCommission            *int64                             `json:"buyerCommission,omitempty"`
	SellerCommission           *int64                             `json:"sellerCommission,omitempty"`
	CommissionRates            *GetAccountResponseCommissionRates `json:"commissionRates,omitempty"`
	CanTrade                   *bool                              `json:"canTrade,omitempty"`
	CanWithdraw                *bool                              `json:"canWithdraw,omitempty"`
	CanDeposit                 *bool                              `json:"canDeposit,omitempty"`
	Brokered                   *bool                              `json:"brokered,omitempty"`
	RequireSelfTradePrevention *bool                              `json:"requireSelfTradePrevention,omitempty"`
	PreventSor                 *bool                              `json:"preventSor,omitempty"`
	UpdateTime                 *int64                             `json:"updateTime,omitempty"`
	AccountType                *string                            `json:"accountType,omitempty"`
	Balances                   []GetAccountResponseBalancesInner  `json:"balances,omitempty"`
	Permissions                []string                           `json:"permissions,omitempty"`
	Uid                        *int64                             `json:"uid,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _GetAccountResponse GetAccountResponse

// NewGetAccountResponse instantiates a new GetAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountResponse() *GetAccountResponse {
	this := GetAccountResponse{}
	return &this
}

// NewGetAccountResponseWithDefaults instantiates a new GetAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountResponseWithDefaults() *GetAccountResponse {
	this := GetAccountResponse{}
	return &this
}

// GetMakerCommission returns the MakerCommission field value if set, zero value otherwise.
func (o *GetAccountResponse) GetMakerCommission() int64 {
	if o == nil || common.IsNil(o.MakerCommission) {
		var ret int64
		return ret
	}
	return *o.MakerCommission
}

// GetMakerCommissionOk returns a tuple with the MakerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetMakerCommissionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MakerCommission) {
		return nil, false
	}
	return o.MakerCommission, true
}

// HasMakerCommission returns a boolean if a field has been set.
func (o *GetAccountResponse) HasMakerCommission() bool {
	if o != nil && !common.IsNil(o.MakerCommission) {
		return true
	}

	return false
}

// SetMakerCommission gets a reference to the given int64 and assigns it to the MakerCommission field.
func (o *GetAccountResponse) SetMakerCommission(v int64) {
	o.MakerCommission = &v
}

// GetTakerCommission returns the TakerCommission field value if set, zero value otherwise.
func (o *GetAccountResponse) GetTakerCommission() int64 {
	if o == nil || common.IsNil(o.TakerCommission) {
		var ret int64
		return ret
	}
	return *o.TakerCommission
}

// GetTakerCommissionOk returns a tuple with the TakerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetTakerCommissionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TakerCommission) {
		return nil, false
	}
	return o.TakerCommission, true
}

// HasTakerCommission returns a boolean if a field has been set.
func (o *GetAccountResponse) HasTakerCommission() bool {
	if o != nil && !common.IsNil(o.TakerCommission) {
		return true
	}

	return false
}

// SetTakerCommission gets a reference to the given int64 and assigns it to the TakerCommission field.
func (o *GetAccountResponse) SetTakerCommission(v int64) {
	o.TakerCommission = &v
}

// GetBuyerCommission returns the BuyerCommission field value if set, zero value otherwise.
func (o *GetAccountResponse) GetBuyerCommission() int64 {
	if o == nil || common.IsNil(o.BuyerCommission) {
		var ret int64
		return ret
	}
	return *o.BuyerCommission
}

// GetBuyerCommissionOk returns a tuple with the BuyerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetBuyerCommissionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BuyerCommission) {
		return nil, false
	}
	return o.BuyerCommission, true
}

// HasBuyerCommission returns a boolean if a field has been set.
func (o *GetAccountResponse) HasBuyerCommission() bool {
	if o != nil && !common.IsNil(o.BuyerCommission) {
		return true
	}

	return false
}

// SetBuyerCommission gets a reference to the given int64 and assigns it to the BuyerCommission field.
func (o *GetAccountResponse) SetBuyerCommission(v int64) {
	o.BuyerCommission = &v
}

// GetSellerCommission returns the SellerCommission field value if set, zero value otherwise.
func (o *GetAccountResponse) GetSellerCommission() int64 {
	if o == nil || common.IsNil(o.SellerCommission) {
		var ret int64
		return ret
	}
	return *o.SellerCommission
}

// GetSellerCommissionOk returns a tuple with the SellerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetSellerCommissionOk() (*int64, bool) {
	if o == nil || common.IsNil(o.SellerCommission) {
		return nil, false
	}
	return o.SellerCommission, true
}

// HasSellerCommission returns a boolean if a field has been set.
func (o *GetAccountResponse) HasSellerCommission() bool {
	if o != nil && !common.IsNil(o.SellerCommission) {
		return true
	}

	return false
}

// SetSellerCommission gets a reference to the given int64 and assigns it to the SellerCommission field.
func (o *GetAccountResponse) SetSellerCommission(v int64) {
	o.SellerCommission = &v
}

// GetCommissionRates returns the CommissionRates field value if set, zero value otherwise.
func (o *GetAccountResponse) GetCommissionRates() GetAccountResponseCommissionRates {
	if o == nil || common.IsNil(o.CommissionRates) {
		var ret GetAccountResponseCommissionRates
		return ret
	}
	return *o.CommissionRates
}

// GetCommissionRatesOk returns a tuple with the CommissionRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetCommissionRatesOk() (*GetAccountResponseCommissionRates, bool) {
	if o == nil || common.IsNil(o.CommissionRates) {
		return nil, false
	}
	return o.CommissionRates, true
}

// HasCommissionRates returns a boolean if a field has been set.
func (o *GetAccountResponse) HasCommissionRates() bool {
	if o != nil && !common.IsNil(o.CommissionRates) {
		return true
	}

	return false
}

// SetCommissionRates gets a reference to the given GetAccountResponseCommissionRates and assigns it to the CommissionRates field.
func (o *GetAccountResponse) SetCommissionRates(v GetAccountResponseCommissionRates) {
	o.CommissionRates = &v
}

// GetCanTrade returns the CanTrade field value if set, zero value otherwise.
func (o *GetAccountResponse) GetCanTrade() bool {
	if o == nil || common.IsNil(o.CanTrade) {
		var ret bool
		return ret
	}
	return *o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetCanTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanTrade) {
		return nil, false
	}
	return o.CanTrade, true
}

// HasCanTrade returns a boolean if a field has been set.
func (o *GetAccountResponse) HasCanTrade() bool {
	if o != nil && !common.IsNil(o.CanTrade) {
		return true
	}

	return false
}

// SetCanTrade gets a reference to the given bool and assigns it to the CanTrade field.
func (o *GetAccountResponse) SetCanTrade(v bool) {
	o.CanTrade = &v
}

// GetCanWithdraw returns the CanWithdraw field value if set, zero value otherwise.
func (o *GetAccountResponse) GetCanWithdraw() bool {
	if o == nil || common.IsNil(o.CanWithdraw) {
		var ret bool
		return ret
	}
	return *o.CanWithdraw
}

// GetCanWithdrawOk returns a tuple with the CanWithdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetCanWithdrawOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanWithdraw) {
		return nil, false
	}
	return o.CanWithdraw, true
}

// HasCanWithdraw returns a boolean if a field has been set.
func (o *GetAccountResponse) HasCanWithdraw() bool {
	if o != nil && !common.IsNil(o.CanWithdraw) {
		return true
	}

	return false
}

// SetCanWithdraw gets a reference to the given bool and assigns it to the CanWithdraw field.
func (o *GetAccountResponse) SetCanWithdraw(v bool) {
	o.CanWithdraw = &v
}

// GetCanDeposit returns the CanDeposit field value if set, zero value otherwise.
func (o *GetAccountResponse) GetCanDeposit() bool {
	if o == nil || common.IsNil(o.CanDeposit) {
		var ret bool
		return ret
	}
	return *o.CanDeposit
}

// GetCanDepositOk returns a tuple with the CanDeposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetCanDepositOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanDeposit) {
		return nil, false
	}
	return o.CanDeposit, true
}

// HasCanDeposit returns a boolean if a field has been set.
func (o *GetAccountResponse) HasCanDeposit() bool {
	if o != nil && !common.IsNil(o.CanDeposit) {
		return true
	}

	return false
}

// SetCanDeposit gets a reference to the given bool and assigns it to the CanDeposit field.
func (o *GetAccountResponse) SetCanDeposit(v bool) {
	o.CanDeposit = &v
}

// GetBrokered returns the Brokered field value if set, zero value otherwise.
func (o *GetAccountResponse) GetBrokered() bool {
	if o == nil || common.IsNil(o.Brokered) {
		var ret bool
		return ret
	}
	return *o.Brokered
}

// GetBrokeredOk returns a tuple with the Brokered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetBrokeredOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Brokered) {
		return nil, false
	}
	return o.Brokered, true
}

// HasBrokered returns a boolean if a field has been set.
func (o *GetAccountResponse) HasBrokered() bool {
	if o != nil && !common.IsNil(o.Brokered) {
		return true
	}

	return false
}

// SetBrokered gets a reference to the given bool and assigns it to the Brokered field.
func (o *GetAccountResponse) SetBrokered(v bool) {
	o.Brokered = &v
}

// GetRequireSelfTradePrevention returns the RequireSelfTradePrevention field value if set, zero value otherwise.
func (o *GetAccountResponse) GetRequireSelfTradePrevention() bool {
	if o == nil || common.IsNil(o.RequireSelfTradePrevention) {
		var ret bool
		return ret
	}
	return *o.RequireSelfTradePrevention
}

// GetRequireSelfTradePreventionOk returns a tuple with the RequireSelfTradePrevention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetRequireSelfTradePreventionOk() (*bool, bool) {
	if o == nil || common.IsNil(o.RequireSelfTradePrevention) {
		return nil, false
	}
	return o.RequireSelfTradePrevention, true
}

// HasRequireSelfTradePrevention returns a boolean if a field has been set.
func (o *GetAccountResponse) HasRequireSelfTradePrevention() bool {
	if o != nil && !common.IsNil(o.RequireSelfTradePrevention) {
		return true
	}

	return false
}

// SetRequireSelfTradePrevention gets a reference to the given bool and assigns it to the RequireSelfTradePrevention field.
func (o *GetAccountResponse) SetRequireSelfTradePrevention(v bool) {
	o.RequireSelfTradePrevention = &v
}

// GetPreventSor returns the PreventSor field value if set, zero value otherwise.
func (o *GetAccountResponse) GetPreventSor() bool {
	if o == nil || common.IsNil(o.PreventSor) {
		var ret bool
		return ret
	}
	return *o.PreventSor
}

// GetPreventSorOk returns a tuple with the PreventSor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetPreventSorOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PreventSor) {
		return nil, false
	}
	return o.PreventSor, true
}

// HasPreventSor returns a boolean if a field has been set.
func (o *GetAccountResponse) HasPreventSor() bool {
	if o != nil && !common.IsNil(o.PreventSor) {
		return true
	}

	return false
}

// SetPreventSor gets a reference to the given bool and assigns it to the PreventSor field.
func (o *GetAccountResponse) SetPreventSor(v bool) {
	o.PreventSor = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetAccountResponse) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetAccountResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetAccountResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *GetAccountResponse) GetAccountType() string {
	if o == nil || common.IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *GetAccountResponse) HasAccountType() bool {
	if o != nil && !common.IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *GetAccountResponse) SetAccountType(v string) {
	o.AccountType = &v
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *GetAccountResponse) GetBalances() []GetAccountResponseBalancesInner {
	if o == nil || common.IsNil(o.Balances) {
		var ret []GetAccountResponseBalancesInner
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetBalancesOk() ([]GetAccountResponseBalancesInner, bool) {
	if o == nil || common.IsNil(o.Balances) {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *GetAccountResponse) HasBalances() bool {
	if o != nil && !common.IsNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []GetAccountResponseBalancesInner and assigns it to the Balances field.
func (o *GetAccountResponse) SetBalances(v []GetAccountResponseBalancesInner) {
	o.Balances = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *GetAccountResponse) GetPermissions() []string {
	if o == nil || common.IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetPermissionsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *GetAccountResponse) HasPermissions() bool {
	if o != nil && !common.IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *GetAccountResponse) SetPermissions(v []string) {
	o.Permissions = v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *GetAccountResponse) GetUid() int64 {
	if o == nil || common.IsNil(o.Uid) {
		var ret int64
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetUidOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *GetAccountResponse) HasUid() bool {
	if o != nil && !common.IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int64 and assigns it to the Uid field.
func (o *GetAccountResponse) SetUid(v int64) {
	o.Uid = &v
}

func (o GetAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MakerCommission) {
		toSerialize["makerCommission"] = o.MakerCommission
	}
	if !common.IsNil(o.TakerCommission) {
		toSerialize["takerCommission"] = o.TakerCommission
	}
	if !common.IsNil(o.BuyerCommission) {
		toSerialize["buyerCommission"] = o.BuyerCommission
	}
	if !common.IsNil(o.SellerCommission) {
		toSerialize["sellerCommission"] = o.SellerCommission
	}
	if !common.IsNil(o.CommissionRates) {
		toSerialize["commissionRates"] = o.CommissionRates
	}
	if !common.IsNil(o.CanTrade) {
		toSerialize["canTrade"] = o.CanTrade
	}
	if !common.IsNil(o.CanWithdraw) {
		toSerialize["canWithdraw"] = o.CanWithdraw
	}
	if !common.IsNil(o.CanDeposit) {
		toSerialize["canDeposit"] = o.CanDeposit
	}
	if !common.IsNil(o.Brokered) {
		toSerialize["brokered"] = o.Brokered
	}
	if !common.IsNil(o.RequireSelfTradePrevention) {
		toSerialize["requireSelfTradePrevention"] = o.RequireSelfTradePrevention
	}
	if !common.IsNil(o.PreventSor) {
		toSerialize["preventSor"] = o.PreventSor
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !common.IsNil(o.Balances) {
		toSerialize["balances"] = o.Balances
	}
	if !common.IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !common.IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varGetAccountResponse := _GetAccountResponse{}

	err = json.Unmarshal(data, &varGetAccountResponse)

	if err != nil {
		return err
	}

	*o = GetAccountResponse(varGetAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "makerCommission")
		delete(additionalProperties, "takerCommission")
		delete(additionalProperties, "buyerCommission")
		delete(additionalProperties, "sellerCommission")
		delete(additionalProperties, "commissionRates")
		delete(additionalProperties, "canTrade")
		delete(additionalProperties, "canWithdraw")
		delete(additionalProperties, "canDeposit")
		delete(additionalProperties, "brokered")
		delete(additionalProperties, "requireSelfTradePrevention")
		delete(additionalProperties, "preventSor")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "accountType")
		delete(additionalProperties, "balances")
		delete(additionalProperties, "permissions")
		delete(additionalProperties, "uid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAccountResponse struct {
	value *GetAccountResponse
	isSet bool
}

func (v NullableGetAccountResponse) Get() *GetAccountResponse {
	return v.value
}

func (v *NullableGetAccountResponse) Set(val *GetAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountResponse(val *GetAccountResponse) *NullableGetAccountResponse {
	return &NullableGetAccountResponse{value: val, isSet: true}
}

func (v NullableGetAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
