/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountInformationV3Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountInformationV3Response{}

// AccountInformationV3Response struct for AccountInformationV3Response
type AccountInformationV3Response struct {
	TotalInitialMargin          *string                                      `json:"totalInitialMargin,omitempty"`
	TotalMaintMargin            *string                                      `json:"totalMaintMargin,omitempty"`
	TotalWalletBalance          *string                                      `json:"totalWalletBalance,omitempty"`
	TotalUnrealizedProfit       *string                                      `json:"totalUnrealizedProfit,omitempty"`
	TotalMarginBalance          *string                                      `json:"totalMarginBalance,omitempty"`
	TotalPositionInitialMargin  *string                                      `json:"totalPositionInitialMargin,omitempty"`
	TotalOpenOrderInitialMargin *string                                      `json:"totalOpenOrderInitialMargin,omitempty"`
	TotalCrossWalletBalance     *string                                      `json:"totalCrossWalletBalance,omitempty"`
	TotalCrossUnPnl             *string                                      `json:"totalCrossUnPnl,omitempty"`
	AvailableBalance            *string                                      `json:"availableBalance,omitempty"`
	MaxWithdrawAmount           *string                                      `json:"maxWithdrawAmount,omitempty"`
	Assets                      []AccountInformationV3ResponseAssetsInner    `json:"assets,omitempty"`
	Positions                   []AccountInformationV3ResponsePositionsInner `json:"positions,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _AccountInformationV3Response AccountInformationV3Response

// NewAccountInformationV3Response instantiates a new AccountInformationV3Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInformationV3Response() *AccountInformationV3Response {
	this := AccountInformationV3Response{}
	return &this
}

// NewAccountInformationV3ResponseWithDefaults instantiates a new AccountInformationV3Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInformationV3ResponseWithDefaults() *AccountInformationV3Response {
	this := AccountInformationV3Response{}
	return &this
}

// GetTotalInitialMargin returns the TotalInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationV3Response) GetTotalInitialMargin() string {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalInitialMargin
}

// GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3Response) GetTotalInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		return nil, false
	}
	return o.TotalInitialMargin, true
}

// HasTotalInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationV3Response) HasTotalInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalInitialMargin) {
		return true
	}

	return false
}

// SetTotalInitialMargin gets a reference to the given string and assigns it to the TotalInitialMargin field.
func (o *AccountInformationV3Response) SetTotalInitialMargin(v string) {
	o.TotalInitialMargin = &v
}

// GetTotalMaintMargin returns the TotalMaintMargin field value if set, zero value otherwise.
func (o *AccountInformationV3Response) GetTotalMaintMargin() string {
	if o == nil || common.IsNil(o.TotalMaintMargin) {
		var ret string
		return ret
	}
	return *o.TotalMaintMargin
}

// GetTotalMaintMarginOk returns a tuple with the TotalMaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3Response) GetTotalMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMaintMargin) {
		return nil, false
	}
	return o.TotalMaintMargin, true
}

// HasTotalMaintMargin returns a boolean if a field has been set.
func (o *AccountInformationV3Response) HasTotalMaintMargin() bool {
	if o != nil && !common.IsNil(o.TotalMaintMargin) {
		return true
	}

	return false
}

// SetTotalMaintMargin gets a reference to the given string and assigns it to the TotalMaintMargin field.
func (o *AccountInformationV3Response) SetTotalMaintMargin(v string) {
	o.TotalMaintMargin = &v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value if set, zero value otherwise.
func (o *AccountInformationV3Response) GetTotalWalletBalance() string {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3Response) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		return nil, false
	}
	return o.TotalWalletBalance, true
}

// HasTotalWalletBalance returns a boolean if a field has been set.
func (o *AccountInformationV3Response) HasTotalWalletBalance() bool {
	if o != nil && !common.IsNil(o.TotalWalletBalance) {
		return true
	}

	return false
}

// SetTotalWalletBalance gets a reference to the given string and assigns it to the TotalWalletBalance field.
func (o *AccountInformationV3Response) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = &v
}

// GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field value if set, zero value otherwise.
func (o *AccountInformationV3Response) GetTotalUnrealizedProfit() string {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.TotalUnrealizedProfit
}

// GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3Response) GetTotalUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		return nil, false
	}
	return o.TotalUnrealizedProfit, true
}

// HasTotalUnrealizedProfit returns a boolean if a field has been set.
func (o *AccountInformationV3Response) HasTotalUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.TotalUnrealizedProfit) {
		return true
	}

	return false
}

// SetTotalUnrealizedProfit gets a reference to the given string and assigns it to the TotalUnrealizedProfit field.
func (o *AccountInformationV3Response) SetTotalUnrealizedProfit(v string) {
	o.TotalUnrealizedProfit = &v
}

// GetTotalMarginBalance returns the TotalMarginBalance field value if set, zero value otherwise.
func (o *AccountInformationV3Response) GetTotalMarginBalance() string {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		var ret string
		return ret
	}
	return *o.TotalMarginBalance
}

// GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3Response) GetTotalMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		return nil, false
	}
	return o.TotalMarginBalance, true
}

// HasTotalMarginBalance returns a boolean if a field has been set.
func (o *AccountInformationV3Response) HasTotalMarginBalance() bool {
	if o != nil && !common.IsNil(o.TotalMarginBalance) {
		return true
	}

	return false
}

// SetTotalMarginBalance gets a reference to the given string and assigns it to the TotalMarginBalance field.
func (o *AccountInformationV3Response) SetTotalMarginBalance(v string) {
	o.TotalMarginBalance = &v
}

// GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationV3Response) GetTotalPositionInitialMargin() string {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalPositionInitialMargin
}

// GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3Response) GetTotalPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		return nil, false
	}
	return o.TotalPositionInitialMargin, true
}

// HasTotalPositionInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationV3Response) HasTotalPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalPositionInitialMargin) {
		return true
	}

	return false
}

// SetTotalPositionInitialMargin gets a reference to the given string and assigns it to the TotalPositionInitialMargin field.
func (o *AccountInformationV3Response) SetTotalPositionInitialMargin(v string) {
	o.TotalPositionInitialMargin = &v
}

// GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationV3Response) GetTotalOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalOpenOrderInitialMargin
}

// GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3Response) GetTotalOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		return nil, false
	}
	return o.TotalOpenOrderInitialMargin, true
}

// HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationV3Response) HasTotalOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalOpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetTotalOpenOrderInitialMargin gets a reference to the given string and assigns it to the TotalOpenOrderInitialMargin field.
func (o *AccountInformationV3Response) SetTotalOpenOrderInitialMargin(v string) {
	o.TotalOpenOrderInitialMargin = &v
}

// GetTotalCrossWalletBalance returns the TotalCrossWalletBalance field value if set, zero value otherwise.
func (o *AccountInformationV3Response) GetTotalCrossWalletBalance() string {
	if o == nil || common.IsNil(o.TotalCrossWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalCrossWalletBalance
}

// GetTotalCrossWalletBalanceOk returns a tuple with the TotalCrossWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3Response) GetTotalCrossWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalCrossWalletBalance) {
		return nil, false
	}
	return o.TotalCrossWalletBalance, true
}

// HasTotalCrossWalletBalance returns a boolean if a field has been set.
func (o *AccountInformationV3Response) HasTotalCrossWalletBalance() bool {
	if o != nil && !common.IsNil(o.TotalCrossWalletBalance) {
		return true
	}

	return false
}

// SetTotalCrossWalletBalance gets a reference to the given string and assigns it to the TotalCrossWalletBalance field.
func (o *AccountInformationV3Response) SetTotalCrossWalletBalance(v string) {
	o.TotalCrossWalletBalance = &v
}

// GetTotalCrossUnPnl returns the TotalCrossUnPnl field value if set, zero value otherwise.
func (o *AccountInformationV3Response) GetTotalCrossUnPnl() string {
	if o == nil || common.IsNil(o.TotalCrossUnPnl) {
		var ret string
		return ret
	}
	return *o.TotalCrossUnPnl
}

// GetTotalCrossUnPnlOk returns a tuple with the TotalCrossUnPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3Response) GetTotalCrossUnPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalCrossUnPnl) {
		return nil, false
	}
	return o.TotalCrossUnPnl, true
}

// HasTotalCrossUnPnl returns a boolean if a field has been set.
func (o *AccountInformationV3Response) HasTotalCrossUnPnl() bool {
	if o != nil && !common.IsNil(o.TotalCrossUnPnl) {
		return true
	}

	return false
}

// SetTotalCrossUnPnl gets a reference to the given string and assigns it to the TotalCrossUnPnl field.
func (o *AccountInformationV3Response) SetTotalCrossUnPnl(v string) {
	o.TotalCrossUnPnl = &v
}

// GetAvailableBalance returns the AvailableBalance field value if set, zero value otherwise.
func (o *AccountInformationV3Response) GetAvailableBalance() string {
	if o == nil || common.IsNil(o.AvailableBalance) {
		var ret string
		return ret
	}
	return *o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3Response) GetAvailableBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvailableBalance) {
		return nil, false
	}
	return o.AvailableBalance, true
}

// HasAvailableBalance returns a boolean if a field has been set.
func (o *AccountInformationV3Response) HasAvailableBalance() bool {
	if o != nil && !common.IsNil(o.AvailableBalance) {
		return true
	}

	return false
}

// SetAvailableBalance gets a reference to the given string and assigns it to the AvailableBalance field.
func (o *AccountInformationV3Response) SetAvailableBalance(v string) {
	o.AvailableBalance = &v
}

// GetMaxWithdrawAmount returns the MaxWithdrawAmount field value if set, zero value otherwise.
func (o *AccountInformationV3Response) GetMaxWithdrawAmount() string {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.MaxWithdrawAmount
}

// GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3Response) GetMaxWithdrawAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		return nil, false
	}
	return o.MaxWithdrawAmount, true
}

// HasMaxWithdrawAmount returns a boolean if a field has been set.
func (o *AccountInformationV3Response) HasMaxWithdrawAmount() bool {
	if o != nil && !common.IsNil(o.MaxWithdrawAmount) {
		return true
	}

	return false
}

// SetMaxWithdrawAmount gets a reference to the given string and assigns it to the MaxWithdrawAmount field.
func (o *AccountInformationV3Response) SetMaxWithdrawAmount(v string) {
	o.MaxWithdrawAmount = &v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *AccountInformationV3Response) GetAssets() []AccountInformationV3ResponseAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []AccountInformationV3ResponseAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3Response) GetAssetsOk() ([]AccountInformationV3ResponseAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *AccountInformationV3Response) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []AccountInformationV3ResponseAssetsInner and assigns it to the Assets field.
func (o *AccountInformationV3Response) SetAssets(v []AccountInformationV3ResponseAssetsInner) {
	o.Assets = v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *AccountInformationV3Response) GetPositions() []AccountInformationV3ResponsePositionsInner {
	if o == nil || common.IsNil(o.Positions) {
		var ret []AccountInformationV3ResponsePositionsInner
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV3Response) GetPositionsOk() ([]AccountInformationV3ResponsePositionsInner, bool) {
	if o == nil || common.IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *AccountInformationV3Response) HasPositions() bool {
	if o != nil && !common.IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given []AccountInformationV3ResponsePositionsInner and assigns it to the Positions field.
func (o *AccountInformationV3Response) SetPositions(v []AccountInformationV3ResponsePositionsInner) {
	o.Positions = v
}

func (o AccountInformationV3Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInformationV3Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TotalInitialMargin) {
		toSerialize["totalInitialMargin"] = o.TotalInitialMargin
	}
	if !common.IsNil(o.TotalMaintMargin) {
		toSerialize["totalMaintMargin"] = o.TotalMaintMargin
	}
	if !common.IsNil(o.TotalWalletBalance) {
		toSerialize["totalWalletBalance"] = o.TotalWalletBalance
	}
	if !common.IsNil(o.TotalUnrealizedProfit) {
		toSerialize["totalUnrealizedProfit"] = o.TotalUnrealizedProfit
	}
	if !common.IsNil(o.TotalMarginBalance) {
		toSerialize["totalMarginBalance"] = o.TotalMarginBalance
	}
	if !common.IsNil(o.TotalPositionInitialMargin) {
		toSerialize["totalPositionInitialMargin"] = o.TotalPositionInitialMargin
	}
	if !common.IsNil(o.TotalOpenOrderInitialMargin) {
		toSerialize["totalOpenOrderInitialMargin"] = o.TotalOpenOrderInitialMargin
	}
	if !common.IsNil(o.TotalCrossWalletBalance) {
		toSerialize["totalCrossWalletBalance"] = o.TotalCrossWalletBalance
	}
	if !common.IsNil(o.TotalCrossUnPnl) {
		toSerialize["totalCrossUnPnl"] = o.TotalCrossUnPnl
	}
	if !common.IsNil(o.AvailableBalance) {
		toSerialize["availableBalance"] = o.AvailableBalance
	}
	if !common.IsNil(o.MaxWithdrawAmount) {
		toSerialize["maxWithdrawAmount"] = o.MaxWithdrawAmount
	}
	if !common.IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !common.IsNil(o.Positions) {
		toSerialize["positions"] = o.Positions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountInformationV3Response) UnmarshalJSON(data []byte) (err error) {
	varAccountInformationV3Response := _AccountInformationV3Response{}

	err = json.Unmarshal(data, &varAccountInformationV3Response)

	if err != nil {
		return err
	}

	*o = AccountInformationV3Response(varAccountInformationV3Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "totalInitialMargin")
		delete(additionalProperties, "totalMaintMargin")
		delete(additionalProperties, "totalWalletBalance")
		delete(additionalProperties, "totalUnrealizedProfit")
		delete(additionalProperties, "totalMarginBalance")
		delete(additionalProperties, "totalPositionInitialMargin")
		delete(additionalProperties, "totalOpenOrderInitialMargin")
		delete(additionalProperties, "totalCrossWalletBalance")
		delete(additionalProperties, "totalCrossUnPnl")
		delete(additionalProperties, "availableBalance")
		delete(additionalProperties, "maxWithdrawAmount")
		delete(additionalProperties, "assets")
		delete(additionalProperties, "positions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountInformationV3Response struct {
	value *AccountInformationV3Response
	isSet bool
}

func (v NullableAccountInformationV3Response) Get() *AccountInformationV3Response {
	return v.value
}

func (v *NullableAccountInformationV3Response) Set(val *AccountInformationV3Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInformationV3Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInformationV3Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInformationV3Response(val *AccountInformationV3Response) *NullableAccountInformationV3Response {
	return &NullableAccountInformationV3Response{value: val, isSet: true}
}

func (v NullableAccountInformationV3Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInformationV3Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
