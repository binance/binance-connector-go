/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryCrossMarginAccountDetailsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCrossMarginAccountDetailsResponse{}

// QueryCrossMarginAccountDetailsResponse struct for QueryCrossMarginAccountDetailsResponse
type QueryCrossMarginAccountDetailsResponse struct {
	Created                    *bool                                                   `json:"created,omitempty"`
	BorrowEnabled              *bool                                                   `json:"borrowEnabled,omitempty"`
	MarginLevel                *string                                                 `json:"marginLevel,omitempty"`
	CollateralMarginLevel      *string                                                 `json:"collateralMarginLevel,omitempty"`
	TotalAssetOfBtc            *string                                                 `json:"totalAssetOfBtc,omitempty"`
	TotalLiabilityOfBtc        *string                                                 `json:"totalLiabilityOfBtc,omitempty"`
	TotalNetAssetOfBtc         *string                                                 `json:"totalNetAssetOfBtc,omitempty"`
	TotalCollateralValueInUSDT *string                                                 `json:"TotalCollateralValueInUSDT,omitempty"`
	TotalOpenOrderLossInUSDT   *string                                                 `json:"totalOpenOrderLossInUSDT,omitempty"`
	TradeEnabled               *bool                                                   `json:"tradeEnabled,omitempty"`
	TransferInEnabled          *bool                                                   `json:"transferInEnabled,omitempty"`
	TransferOutEnabled         *bool                                                   `json:"transferOutEnabled,omitempty"`
	AccountType                *string                                                 `json:"accountType,omitempty"`
	UserAssets                 []QueryCrossMarginAccountDetailsResponseUserAssetsInner `json:"userAssets,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _QueryCrossMarginAccountDetailsResponse QueryCrossMarginAccountDetailsResponse

// NewQueryCrossMarginAccountDetailsResponse instantiates a new QueryCrossMarginAccountDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCrossMarginAccountDetailsResponse() *QueryCrossMarginAccountDetailsResponse {
	this := QueryCrossMarginAccountDetailsResponse{}
	return &this
}

// NewQueryCrossMarginAccountDetailsResponseWithDefaults instantiates a new QueryCrossMarginAccountDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCrossMarginAccountDetailsResponseWithDefaults() *QueryCrossMarginAccountDetailsResponse {
	this := QueryCrossMarginAccountDetailsResponse{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *QueryCrossMarginAccountDetailsResponse) GetCreated() bool {
	if o == nil || common.IsNil(o.Created) {
		var ret bool
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginAccountDetailsResponse) GetCreatedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *QueryCrossMarginAccountDetailsResponse) HasCreated() bool {
	if o != nil && !common.IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given bool and assigns it to the Created field.
func (o *QueryCrossMarginAccountDetailsResponse) SetCreated(v bool) {
	o.Created = &v
}

// GetBorrowEnabled returns the BorrowEnabled field value if set, zero value otherwise.
func (o *QueryCrossMarginAccountDetailsResponse) GetBorrowEnabled() bool {
	if o == nil || common.IsNil(o.BorrowEnabled) {
		var ret bool
		return ret
	}
	return *o.BorrowEnabled
}

// GetBorrowEnabledOk returns a tuple with the BorrowEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginAccountDetailsResponse) GetBorrowEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.BorrowEnabled) {
		return nil, false
	}
	return o.BorrowEnabled, true
}

// HasBorrowEnabled returns a boolean if a field has been set.
func (o *QueryCrossMarginAccountDetailsResponse) HasBorrowEnabled() bool {
	if o != nil && !common.IsNil(o.BorrowEnabled) {
		return true
	}

	return false
}

// SetBorrowEnabled gets a reference to the given bool and assigns it to the BorrowEnabled field.
func (o *QueryCrossMarginAccountDetailsResponse) SetBorrowEnabled(v bool) {
	o.BorrowEnabled = &v
}

// GetMarginLevel returns the MarginLevel field value if set, zero value otherwise.
func (o *QueryCrossMarginAccountDetailsResponse) GetMarginLevel() string {
	if o == nil || common.IsNil(o.MarginLevel) {
		var ret string
		return ret
	}
	return *o.MarginLevel
}

// GetMarginLevelOk returns a tuple with the MarginLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginAccountDetailsResponse) GetMarginLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginLevel) {
		return nil, false
	}
	return o.MarginLevel, true
}

// HasMarginLevel returns a boolean if a field has been set.
func (o *QueryCrossMarginAccountDetailsResponse) HasMarginLevel() bool {
	if o != nil && !common.IsNil(o.MarginLevel) {
		return true
	}

	return false
}

// SetMarginLevel gets a reference to the given string and assigns it to the MarginLevel field.
func (o *QueryCrossMarginAccountDetailsResponse) SetMarginLevel(v string) {
	o.MarginLevel = &v
}

// GetCollateralMarginLevel returns the CollateralMarginLevel field value if set, zero value otherwise.
func (o *QueryCrossMarginAccountDetailsResponse) GetCollateralMarginLevel() string {
	if o == nil || common.IsNil(o.CollateralMarginLevel) {
		var ret string
		return ret
	}
	return *o.CollateralMarginLevel
}

// GetCollateralMarginLevelOk returns a tuple with the CollateralMarginLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginAccountDetailsResponse) GetCollateralMarginLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralMarginLevel) {
		return nil, false
	}
	return o.CollateralMarginLevel, true
}

// HasCollateralMarginLevel returns a boolean if a field has been set.
func (o *QueryCrossMarginAccountDetailsResponse) HasCollateralMarginLevel() bool {
	if o != nil && !common.IsNil(o.CollateralMarginLevel) {
		return true
	}

	return false
}

// SetCollateralMarginLevel gets a reference to the given string and assigns it to the CollateralMarginLevel field.
func (o *QueryCrossMarginAccountDetailsResponse) SetCollateralMarginLevel(v string) {
	o.CollateralMarginLevel = &v
}

// GetTotalAssetOfBtc returns the TotalAssetOfBtc field value if set, zero value otherwise.
func (o *QueryCrossMarginAccountDetailsResponse) GetTotalAssetOfBtc() string {
	if o == nil || common.IsNil(o.TotalAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalAssetOfBtc
}

// GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginAccountDetailsResponse) GetTotalAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAssetOfBtc) {
		return nil, false
	}
	return o.TotalAssetOfBtc, true
}

// HasTotalAssetOfBtc returns a boolean if a field has been set.
func (o *QueryCrossMarginAccountDetailsResponse) HasTotalAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalAssetOfBtc gets a reference to the given string and assigns it to the TotalAssetOfBtc field.
func (o *QueryCrossMarginAccountDetailsResponse) SetTotalAssetOfBtc(v string) {
	o.TotalAssetOfBtc = &v
}

// GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field value if set, zero value otherwise.
func (o *QueryCrossMarginAccountDetailsResponse) GetTotalLiabilityOfBtc() string {
	if o == nil || common.IsNil(o.TotalLiabilityOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalLiabilityOfBtc
}

// GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginAccountDetailsResponse) GetTotalLiabilityOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalLiabilityOfBtc) {
		return nil, false
	}
	return o.TotalLiabilityOfBtc, true
}

// HasTotalLiabilityOfBtc returns a boolean if a field has been set.
func (o *QueryCrossMarginAccountDetailsResponse) HasTotalLiabilityOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalLiabilityOfBtc) {
		return true
	}

	return false
}

// SetTotalLiabilityOfBtc gets a reference to the given string and assigns it to the TotalLiabilityOfBtc field.
func (o *QueryCrossMarginAccountDetailsResponse) SetTotalLiabilityOfBtc(v string) {
	o.TotalLiabilityOfBtc = &v
}

// GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field value if set, zero value otherwise.
func (o *QueryCrossMarginAccountDetailsResponse) GetTotalNetAssetOfBtc() string {
	if o == nil || common.IsNil(o.TotalNetAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalNetAssetOfBtc
}

// GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginAccountDetailsResponse) GetTotalNetAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalNetAssetOfBtc) {
		return nil, false
	}
	return o.TotalNetAssetOfBtc, true
}

// HasTotalNetAssetOfBtc returns a boolean if a field has been set.
func (o *QueryCrossMarginAccountDetailsResponse) HasTotalNetAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalNetAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalNetAssetOfBtc gets a reference to the given string and assigns it to the TotalNetAssetOfBtc field.
func (o *QueryCrossMarginAccountDetailsResponse) SetTotalNetAssetOfBtc(v string) {
	o.TotalNetAssetOfBtc = &v
}

// GetTotalCollateralValueInUSDT returns the TotalCollateralValueInUSDT field value if set, zero value otherwise.
func (o *QueryCrossMarginAccountDetailsResponse) GetTotalCollateralValueInUSDT() string {
	if o == nil || common.IsNil(o.TotalCollateralValueInUSDT) {
		var ret string
		return ret
	}
	return *o.TotalCollateralValueInUSDT
}

// GetTotalCollateralValueInUSDTOk returns a tuple with the TotalCollateralValueInUSDT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginAccountDetailsResponse) GetTotalCollateralValueInUSDTOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalCollateralValueInUSDT) {
		return nil, false
	}
	return o.TotalCollateralValueInUSDT, true
}

// HasTotalCollateralValueInUSDT returns a boolean if a field has been set.
func (o *QueryCrossMarginAccountDetailsResponse) HasTotalCollateralValueInUSDT() bool {
	if o != nil && !common.IsNil(o.TotalCollateralValueInUSDT) {
		return true
	}

	return false
}

// SetTotalCollateralValueInUSDT gets a reference to the given string and assigns it to the TotalCollateralValueInUSDT field.
func (o *QueryCrossMarginAccountDetailsResponse) SetTotalCollateralValueInUSDT(v string) {
	o.TotalCollateralValueInUSDT = &v
}

// GetTotalOpenOrderLossInUSDT returns the TotalOpenOrderLossInUSDT field value if set, zero value otherwise.
func (o *QueryCrossMarginAccountDetailsResponse) GetTotalOpenOrderLossInUSDT() string {
	if o == nil || common.IsNil(o.TotalOpenOrderLossInUSDT) {
		var ret string
		return ret
	}
	return *o.TotalOpenOrderLossInUSDT
}

// GetTotalOpenOrderLossInUSDTOk returns a tuple with the TotalOpenOrderLossInUSDT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginAccountDetailsResponse) GetTotalOpenOrderLossInUSDTOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalOpenOrderLossInUSDT) {
		return nil, false
	}
	return o.TotalOpenOrderLossInUSDT, true
}

// HasTotalOpenOrderLossInUSDT returns a boolean if a field has been set.
func (o *QueryCrossMarginAccountDetailsResponse) HasTotalOpenOrderLossInUSDT() bool {
	if o != nil && !common.IsNil(o.TotalOpenOrderLossInUSDT) {
		return true
	}

	return false
}

// SetTotalOpenOrderLossInUSDT gets a reference to the given string and assigns it to the TotalOpenOrderLossInUSDT field.
func (o *QueryCrossMarginAccountDetailsResponse) SetTotalOpenOrderLossInUSDT(v string) {
	o.TotalOpenOrderLossInUSDT = &v
}

// GetTradeEnabled returns the TradeEnabled field value if set, zero value otherwise.
func (o *QueryCrossMarginAccountDetailsResponse) GetTradeEnabled() bool {
	if o == nil || common.IsNil(o.TradeEnabled) {
		var ret bool
		return ret
	}
	return *o.TradeEnabled
}

// GetTradeEnabledOk returns a tuple with the TradeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginAccountDetailsResponse) GetTradeEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.TradeEnabled) {
		return nil, false
	}
	return o.TradeEnabled, true
}

// HasTradeEnabled returns a boolean if a field has been set.
func (o *QueryCrossMarginAccountDetailsResponse) HasTradeEnabled() bool {
	if o != nil && !common.IsNil(o.TradeEnabled) {
		return true
	}

	return false
}

// SetTradeEnabled gets a reference to the given bool and assigns it to the TradeEnabled field.
func (o *QueryCrossMarginAccountDetailsResponse) SetTradeEnabled(v bool) {
	o.TradeEnabled = &v
}

// GetTransferInEnabled returns the TransferInEnabled field value if set, zero value otherwise.
func (o *QueryCrossMarginAccountDetailsResponse) GetTransferInEnabled() bool {
	if o == nil || common.IsNil(o.TransferInEnabled) {
		var ret bool
		return ret
	}
	return *o.TransferInEnabled
}

// GetTransferInEnabledOk returns a tuple with the TransferInEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginAccountDetailsResponse) GetTransferInEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.TransferInEnabled) {
		return nil, false
	}
	return o.TransferInEnabled, true
}

// HasTransferInEnabled returns a boolean if a field has been set.
func (o *QueryCrossMarginAccountDetailsResponse) HasTransferInEnabled() bool {
	if o != nil && !common.IsNil(o.TransferInEnabled) {
		return true
	}

	return false
}

// SetTransferInEnabled gets a reference to the given bool and assigns it to the TransferInEnabled field.
func (o *QueryCrossMarginAccountDetailsResponse) SetTransferInEnabled(v bool) {
	o.TransferInEnabled = &v
}

// GetTransferOutEnabled returns the TransferOutEnabled field value if set, zero value otherwise.
func (o *QueryCrossMarginAccountDetailsResponse) GetTransferOutEnabled() bool {
	if o == nil || common.IsNil(o.TransferOutEnabled) {
		var ret bool
		return ret
	}
	return *o.TransferOutEnabled
}

// GetTransferOutEnabledOk returns a tuple with the TransferOutEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginAccountDetailsResponse) GetTransferOutEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.TransferOutEnabled) {
		return nil, false
	}
	return o.TransferOutEnabled, true
}

// HasTransferOutEnabled returns a boolean if a field has been set.
func (o *QueryCrossMarginAccountDetailsResponse) HasTransferOutEnabled() bool {
	if o != nil && !common.IsNil(o.TransferOutEnabled) {
		return true
	}

	return false
}

// SetTransferOutEnabled gets a reference to the given bool and assigns it to the TransferOutEnabled field.
func (o *QueryCrossMarginAccountDetailsResponse) SetTransferOutEnabled(v bool) {
	o.TransferOutEnabled = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *QueryCrossMarginAccountDetailsResponse) GetAccountType() string {
	if o == nil || common.IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginAccountDetailsResponse) GetAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *QueryCrossMarginAccountDetailsResponse) HasAccountType() bool {
	if o != nil && !common.IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *QueryCrossMarginAccountDetailsResponse) SetAccountType(v string) {
	o.AccountType = &v
}

// GetUserAssets returns the UserAssets field value if set, zero value otherwise.
func (o *QueryCrossMarginAccountDetailsResponse) GetUserAssets() []QueryCrossMarginAccountDetailsResponseUserAssetsInner {
	if o == nil || common.IsNil(o.UserAssets) {
		var ret []QueryCrossMarginAccountDetailsResponseUserAssetsInner
		return ret
	}
	return o.UserAssets
}

// GetUserAssetsOk returns a tuple with the UserAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossMarginAccountDetailsResponse) GetUserAssetsOk() ([]QueryCrossMarginAccountDetailsResponseUserAssetsInner, bool) {
	if o == nil || common.IsNil(o.UserAssets) {
		return nil, false
	}
	return o.UserAssets, true
}

// HasUserAssets returns a boolean if a field has been set.
func (o *QueryCrossMarginAccountDetailsResponse) HasUserAssets() bool {
	if o != nil && !common.IsNil(o.UserAssets) {
		return true
	}

	return false
}

// SetUserAssets gets a reference to the given []QueryCrossMarginAccountDetailsResponseUserAssetsInner and assigns it to the UserAssets field.
func (o *QueryCrossMarginAccountDetailsResponse) SetUserAssets(v []QueryCrossMarginAccountDetailsResponseUserAssetsInner) {
	o.UserAssets = v
}

func (o QueryCrossMarginAccountDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCrossMarginAccountDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !common.IsNil(o.BorrowEnabled) {
		toSerialize["borrowEnabled"] = o.BorrowEnabled
	}
	if !common.IsNil(o.MarginLevel) {
		toSerialize["marginLevel"] = o.MarginLevel
	}
	if !common.IsNil(o.CollateralMarginLevel) {
		toSerialize["collateralMarginLevel"] = o.CollateralMarginLevel
	}
	if !common.IsNil(o.TotalAssetOfBtc) {
		toSerialize["totalAssetOfBtc"] = o.TotalAssetOfBtc
	}
	if !common.IsNil(o.TotalLiabilityOfBtc) {
		toSerialize["totalLiabilityOfBtc"] = o.TotalLiabilityOfBtc
	}
	if !common.IsNil(o.TotalNetAssetOfBtc) {
		toSerialize["totalNetAssetOfBtc"] = o.TotalNetAssetOfBtc
	}
	if !common.IsNil(o.TotalCollateralValueInUSDT) {
		toSerialize["TotalCollateralValueInUSDT"] = o.TotalCollateralValueInUSDT
	}
	if !common.IsNil(o.TotalOpenOrderLossInUSDT) {
		toSerialize["totalOpenOrderLossInUSDT"] = o.TotalOpenOrderLossInUSDT
	}
	if !common.IsNil(o.TradeEnabled) {
		toSerialize["tradeEnabled"] = o.TradeEnabled
	}
	if !common.IsNil(o.TransferInEnabled) {
		toSerialize["transferInEnabled"] = o.TransferInEnabled
	}
	if !common.IsNil(o.TransferOutEnabled) {
		toSerialize["transferOutEnabled"] = o.TransferOutEnabled
	}
	if !common.IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !common.IsNil(o.UserAssets) {
		toSerialize["userAssets"] = o.UserAssets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryCrossMarginAccountDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryCrossMarginAccountDetailsResponse := _QueryCrossMarginAccountDetailsResponse{}

	err = json.Unmarshal(data, &varQueryCrossMarginAccountDetailsResponse)

	if err != nil {
		return err
	}

	*o = QueryCrossMarginAccountDetailsResponse(varQueryCrossMarginAccountDetailsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "borrowEnabled")
		delete(additionalProperties, "marginLevel")
		delete(additionalProperties, "collateralMarginLevel")
		delete(additionalProperties, "totalAssetOfBtc")
		delete(additionalProperties, "totalLiabilityOfBtc")
		delete(additionalProperties, "totalNetAssetOfBtc")
		delete(additionalProperties, "TotalCollateralValueInUSDT")
		delete(additionalProperties, "totalOpenOrderLossInUSDT")
		delete(additionalProperties, "tradeEnabled")
		delete(additionalProperties, "transferInEnabled")
		delete(additionalProperties, "transferOutEnabled")
		delete(additionalProperties, "accountType")
		delete(additionalProperties, "userAssets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryCrossMarginAccountDetailsResponse struct {
	value *QueryCrossMarginAccountDetailsResponse
	isSet bool
}

func (v NullableQueryCrossMarginAccountDetailsResponse) Get() *QueryCrossMarginAccountDetailsResponse {
	return v.value
}

func (v *NullableQueryCrossMarginAccountDetailsResponse) Set(val *QueryCrossMarginAccountDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCrossMarginAccountDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCrossMarginAccountDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCrossMarginAccountDetailsResponse(val *QueryCrossMarginAccountDetailsResponse) *NullableQueryCrossMarginAccountDetailsResponse {
	return &NullableQueryCrossMarginAccountDetailsResponse{value: val, isSet: true}
}

func (v NullableQueryCrossMarginAccountDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCrossMarginAccountDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
