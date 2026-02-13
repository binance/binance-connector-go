/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AllCoinsInformationResponseInnerNetworkListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllCoinsInformationResponseInnerNetworkListInner{}

// AllCoinsInformationResponseInnerNetworkListInner struct for AllCoinsInformationResponseInnerNetworkListInner
type AllCoinsInformationResponseInnerNetworkListInner struct {
	Network                 *string `json:"network,omitempty"`
	Coin                    *string `json:"coin,omitempty"`
	WithdrawIntegerMultiple *string `json:"withdrawIntegerMultiple,omitempty"`
	IsDefault               *bool   `json:"isDefault,omitempty"`
	DepositEnable           *bool   `json:"depositEnable,omitempty"`
	WithdrawEnable          *bool   `json:"withdrawEnable,omitempty"`
	DepositDesc             *string `json:"depositDesc,omitempty"`
	WithdrawDesc            *string `json:"withdrawDesc,omitempty"`
	SpecialTips             *string `json:"specialTips,omitempty"`
	SpecialWithdrawTips     *string `json:"specialWithdrawTips,omitempty"`
	Name                    *string `json:"name,omitempty"`
	ResetAddressStatus      *bool   `json:"resetAddressStatus,omitempty"`
	AddressRegex            *string `json:"addressRegex,omitempty"`
	MemoRegex               *string `json:"memoRegex,omitempty"`
	WithdrawFee             *string `json:"withdrawFee,omitempty"`
	WithdrawMin             *string `json:"withdrawMin,omitempty"`
	WithdrawMax             *string `json:"withdrawMax,omitempty"`
	WithdrawInternalMin     *string `json:"withdrawInternalMin,omitempty"`
	DepositDust             *string `json:"depositDust,omitempty"`
	MinConfirm              *int64  `json:"minConfirm,omitempty"`
	UnLockConfirm           *int64  `json:"unLockConfirm,omitempty"`
	SameAddress             *bool   `json:"sameAddress,omitempty"`
	WithdrawTag             *bool   `json:"withdrawTag,omitempty"`
	EstimatedArrivalTime    *int64  `json:"estimatedArrivalTime,omitempty"`
	Busy                    *bool   `json:"busy,omitempty"`
	ContractAddressUrl      *string `json:"contractAddressUrl,omitempty"`
	ContractAddress         *string `json:"contractAddress,omitempty"`
	Denomination            *int64  `json:"denomination,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _AllCoinsInformationResponseInnerNetworkListInner AllCoinsInformationResponseInnerNetworkListInner

// NewAllCoinsInformationResponseInnerNetworkListInner instantiates a new AllCoinsInformationResponseInnerNetworkListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllCoinsInformationResponseInnerNetworkListInner() *AllCoinsInformationResponseInnerNetworkListInner {
	this := AllCoinsInformationResponseInnerNetworkListInner{}
	return &this
}

// NewAllCoinsInformationResponseInnerNetworkListInnerWithDefaults instantiates a new AllCoinsInformationResponseInnerNetworkListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllCoinsInformationResponseInnerNetworkListInnerWithDefaults() *AllCoinsInformationResponseInnerNetworkListInner {
	this := AllCoinsInformationResponseInnerNetworkListInner{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetNetwork() string {
	if o == nil || common.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetNetworkOk() (*string, bool) {
	if o == nil || common.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasNetwork() bool {
	if o != nil && !common.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetNetwork(v string) {
	o.Network = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetCoin(v string) {
	o.Coin = &v
}

// GetWithdrawIntegerMultiple returns the WithdrawIntegerMultiple field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetWithdrawIntegerMultiple() string {
	if o == nil || common.IsNil(o.WithdrawIntegerMultiple) {
		var ret string
		return ret
	}
	return *o.WithdrawIntegerMultiple
}

// GetWithdrawIntegerMultipleOk returns a tuple with the WithdrawIntegerMultiple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetWithdrawIntegerMultipleOk() (*string, bool) {
	if o == nil || common.IsNil(o.WithdrawIntegerMultiple) {
		return nil, false
	}
	return o.WithdrawIntegerMultiple, true
}

// HasWithdrawIntegerMultiple returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasWithdrawIntegerMultiple() bool {
	if o != nil && !common.IsNil(o.WithdrawIntegerMultiple) {
		return true
	}

	return false
}

// SetWithdrawIntegerMultiple gets a reference to the given string and assigns it to the WithdrawIntegerMultiple field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetWithdrawIntegerMultiple(v string) {
	o.WithdrawIntegerMultiple = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetIsDefault() bool {
	if o == nil || common.IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetIsDefaultOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasIsDefault() bool {
	if o != nil && !common.IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetDepositEnable returns the DepositEnable field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetDepositEnable() bool {
	if o == nil || common.IsNil(o.DepositEnable) {
		var ret bool
		return ret
	}
	return *o.DepositEnable
}

// GetDepositEnableOk returns a tuple with the DepositEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetDepositEnableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.DepositEnable) {
		return nil, false
	}
	return o.DepositEnable, true
}

// HasDepositEnable returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasDepositEnable() bool {
	if o != nil && !common.IsNil(o.DepositEnable) {
		return true
	}

	return false
}

// SetDepositEnable gets a reference to the given bool and assigns it to the DepositEnable field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetDepositEnable(v bool) {
	o.DepositEnable = &v
}

// GetWithdrawEnable returns the WithdrawEnable field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetWithdrawEnable() bool {
	if o == nil || common.IsNil(o.WithdrawEnable) {
		var ret bool
		return ret
	}
	return *o.WithdrawEnable
}

// GetWithdrawEnableOk returns a tuple with the WithdrawEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetWithdrawEnableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.WithdrawEnable) {
		return nil, false
	}
	return o.WithdrawEnable, true
}

// HasWithdrawEnable returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasWithdrawEnable() bool {
	if o != nil && !common.IsNil(o.WithdrawEnable) {
		return true
	}

	return false
}

// SetWithdrawEnable gets a reference to the given bool and assigns it to the WithdrawEnable field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetWithdrawEnable(v bool) {
	o.WithdrawEnable = &v
}

// GetDepositDesc returns the DepositDesc field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetDepositDesc() string {
	if o == nil || common.IsNil(o.DepositDesc) {
		var ret string
		return ret
	}
	return *o.DepositDesc
}

// GetDepositDescOk returns a tuple with the DepositDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetDepositDescOk() (*string, bool) {
	if o == nil || common.IsNil(o.DepositDesc) {
		return nil, false
	}
	return o.DepositDesc, true
}

// HasDepositDesc returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasDepositDesc() bool {
	if o != nil && !common.IsNil(o.DepositDesc) {
		return true
	}

	return false
}

// SetDepositDesc gets a reference to the given string and assigns it to the DepositDesc field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetDepositDesc(v string) {
	o.DepositDesc = &v
}

// GetWithdrawDesc returns the WithdrawDesc field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetWithdrawDesc() string {
	if o == nil || common.IsNil(o.WithdrawDesc) {
		var ret string
		return ret
	}
	return *o.WithdrawDesc
}

// GetWithdrawDescOk returns a tuple with the WithdrawDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetWithdrawDescOk() (*string, bool) {
	if o == nil || common.IsNil(o.WithdrawDesc) {
		return nil, false
	}
	return o.WithdrawDesc, true
}

// HasWithdrawDesc returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasWithdrawDesc() bool {
	if o != nil && !common.IsNil(o.WithdrawDesc) {
		return true
	}

	return false
}

// SetWithdrawDesc gets a reference to the given string and assigns it to the WithdrawDesc field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetWithdrawDesc(v string) {
	o.WithdrawDesc = &v
}

// GetSpecialTips returns the SpecialTips field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetSpecialTips() string {
	if o == nil || common.IsNil(o.SpecialTips) {
		var ret string
		return ret
	}
	return *o.SpecialTips
}

// GetSpecialTipsOk returns a tuple with the SpecialTips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetSpecialTipsOk() (*string, bool) {
	if o == nil || common.IsNil(o.SpecialTips) {
		return nil, false
	}
	return o.SpecialTips, true
}

// HasSpecialTips returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasSpecialTips() bool {
	if o != nil && !common.IsNil(o.SpecialTips) {
		return true
	}

	return false
}

// SetSpecialTips gets a reference to the given string and assigns it to the SpecialTips field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetSpecialTips(v string) {
	o.SpecialTips = &v
}

// GetSpecialWithdrawTips returns the SpecialWithdrawTips field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetSpecialWithdrawTips() string {
	if o == nil || common.IsNil(o.SpecialWithdrawTips) {
		var ret string
		return ret
	}
	return *o.SpecialWithdrawTips
}

// GetSpecialWithdrawTipsOk returns a tuple with the SpecialWithdrawTips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetSpecialWithdrawTipsOk() (*string, bool) {
	if o == nil || common.IsNil(o.SpecialWithdrawTips) {
		return nil, false
	}
	return o.SpecialWithdrawTips, true
}

// HasSpecialWithdrawTips returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasSpecialWithdrawTips() bool {
	if o != nil && !common.IsNil(o.SpecialWithdrawTips) {
		return true
	}

	return false
}

// SetSpecialWithdrawTips gets a reference to the given string and assigns it to the SpecialWithdrawTips field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetSpecialWithdrawTips(v string) {
	o.SpecialWithdrawTips = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetName(v string) {
	o.Name = &v
}

// GetResetAddressStatus returns the ResetAddressStatus field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetResetAddressStatus() bool {
	if o == nil || common.IsNil(o.ResetAddressStatus) {
		var ret bool
		return ret
	}
	return *o.ResetAddressStatus
}

// GetResetAddressStatusOk returns a tuple with the ResetAddressStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetResetAddressStatusOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ResetAddressStatus) {
		return nil, false
	}
	return o.ResetAddressStatus, true
}

// HasResetAddressStatus returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasResetAddressStatus() bool {
	if o != nil && !common.IsNil(o.ResetAddressStatus) {
		return true
	}

	return false
}

// SetResetAddressStatus gets a reference to the given bool and assigns it to the ResetAddressStatus field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetResetAddressStatus(v bool) {
	o.ResetAddressStatus = &v
}

// GetAddressRegex returns the AddressRegex field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetAddressRegex() string {
	if o == nil || common.IsNil(o.AddressRegex) {
		var ret string
		return ret
	}
	return *o.AddressRegex
}

// GetAddressRegexOk returns a tuple with the AddressRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetAddressRegexOk() (*string, bool) {
	if o == nil || common.IsNil(o.AddressRegex) {
		return nil, false
	}
	return o.AddressRegex, true
}

// HasAddressRegex returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasAddressRegex() bool {
	if o != nil && !common.IsNil(o.AddressRegex) {
		return true
	}

	return false
}

// SetAddressRegex gets a reference to the given string and assigns it to the AddressRegex field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetAddressRegex(v string) {
	o.AddressRegex = &v
}

// GetMemoRegex returns the MemoRegex field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetMemoRegex() string {
	if o == nil || common.IsNil(o.MemoRegex) {
		var ret string
		return ret
	}
	return *o.MemoRegex
}

// GetMemoRegexOk returns a tuple with the MemoRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetMemoRegexOk() (*string, bool) {
	if o == nil || common.IsNil(o.MemoRegex) {
		return nil, false
	}
	return o.MemoRegex, true
}

// HasMemoRegex returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasMemoRegex() bool {
	if o != nil && !common.IsNil(o.MemoRegex) {
		return true
	}

	return false
}

// SetMemoRegex gets a reference to the given string and assigns it to the MemoRegex field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetMemoRegex(v string) {
	o.MemoRegex = &v
}

// GetWithdrawFee returns the WithdrawFee field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetWithdrawFee() string {
	if o == nil || common.IsNil(o.WithdrawFee) {
		var ret string
		return ret
	}
	return *o.WithdrawFee
}

// GetWithdrawFeeOk returns a tuple with the WithdrawFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetWithdrawFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.WithdrawFee) {
		return nil, false
	}
	return o.WithdrawFee, true
}

// HasWithdrawFee returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasWithdrawFee() bool {
	if o != nil && !common.IsNil(o.WithdrawFee) {
		return true
	}

	return false
}

// SetWithdrawFee gets a reference to the given string and assigns it to the WithdrawFee field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetWithdrawFee(v string) {
	o.WithdrawFee = &v
}

// GetWithdrawMin returns the WithdrawMin field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetWithdrawMin() string {
	if o == nil || common.IsNil(o.WithdrawMin) {
		var ret string
		return ret
	}
	return *o.WithdrawMin
}

// GetWithdrawMinOk returns a tuple with the WithdrawMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetWithdrawMinOk() (*string, bool) {
	if o == nil || common.IsNil(o.WithdrawMin) {
		return nil, false
	}
	return o.WithdrawMin, true
}

// HasWithdrawMin returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasWithdrawMin() bool {
	if o != nil && !common.IsNil(o.WithdrawMin) {
		return true
	}

	return false
}

// SetWithdrawMin gets a reference to the given string and assigns it to the WithdrawMin field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetWithdrawMin(v string) {
	o.WithdrawMin = &v
}

// GetWithdrawMax returns the WithdrawMax field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetWithdrawMax() string {
	if o == nil || common.IsNil(o.WithdrawMax) {
		var ret string
		return ret
	}
	return *o.WithdrawMax
}

// GetWithdrawMaxOk returns a tuple with the WithdrawMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetWithdrawMaxOk() (*string, bool) {
	if o == nil || common.IsNil(o.WithdrawMax) {
		return nil, false
	}
	return o.WithdrawMax, true
}

// HasWithdrawMax returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasWithdrawMax() bool {
	if o != nil && !common.IsNil(o.WithdrawMax) {
		return true
	}

	return false
}

// SetWithdrawMax gets a reference to the given string and assigns it to the WithdrawMax field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetWithdrawMax(v string) {
	o.WithdrawMax = &v
}

// GetWithdrawInternalMin returns the WithdrawInternalMin field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetWithdrawInternalMin() string {
	if o == nil || common.IsNil(o.WithdrawInternalMin) {
		var ret string
		return ret
	}
	return *o.WithdrawInternalMin
}

// GetWithdrawInternalMinOk returns a tuple with the WithdrawInternalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetWithdrawInternalMinOk() (*string, bool) {
	if o == nil || common.IsNil(o.WithdrawInternalMin) {
		return nil, false
	}
	return o.WithdrawInternalMin, true
}

// HasWithdrawInternalMin returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasWithdrawInternalMin() bool {
	if o != nil && !common.IsNil(o.WithdrawInternalMin) {
		return true
	}

	return false
}

// SetWithdrawInternalMin gets a reference to the given string and assigns it to the WithdrawInternalMin field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetWithdrawInternalMin(v string) {
	o.WithdrawInternalMin = &v
}

// GetDepositDust returns the DepositDust field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetDepositDust() string {
	if o == nil || common.IsNil(o.DepositDust) {
		var ret string
		return ret
	}
	return *o.DepositDust
}

// GetDepositDustOk returns a tuple with the DepositDust field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetDepositDustOk() (*string, bool) {
	if o == nil || common.IsNil(o.DepositDust) {
		return nil, false
	}
	return o.DepositDust, true
}

// HasDepositDust returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasDepositDust() bool {
	if o != nil && !common.IsNil(o.DepositDust) {
		return true
	}

	return false
}

// SetDepositDust gets a reference to the given string and assigns it to the DepositDust field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetDepositDust(v string) {
	o.DepositDust = &v
}

// GetMinConfirm returns the MinConfirm field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetMinConfirm() int64 {
	if o == nil || common.IsNil(o.MinConfirm) {
		var ret int64
		return ret
	}
	return *o.MinConfirm
}

// GetMinConfirmOk returns a tuple with the MinConfirm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetMinConfirmOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MinConfirm) {
		return nil, false
	}
	return o.MinConfirm, true
}

// HasMinConfirm returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasMinConfirm() bool {
	if o != nil && !common.IsNil(o.MinConfirm) {
		return true
	}

	return false
}

// SetMinConfirm gets a reference to the given int64 and assigns it to the MinConfirm field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetMinConfirm(v int64) {
	o.MinConfirm = &v
}

// GetUnLockConfirm returns the UnLockConfirm field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetUnLockConfirm() int64 {
	if o == nil || common.IsNil(o.UnLockConfirm) {
		var ret int64
		return ret
	}
	return *o.UnLockConfirm
}

// GetUnLockConfirmOk returns a tuple with the UnLockConfirm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetUnLockConfirmOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UnLockConfirm) {
		return nil, false
	}
	return o.UnLockConfirm, true
}

// HasUnLockConfirm returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasUnLockConfirm() bool {
	if o != nil && !common.IsNil(o.UnLockConfirm) {
		return true
	}

	return false
}

// SetUnLockConfirm gets a reference to the given int64 and assigns it to the UnLockConfirm field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetUnLockConfirm(v int64) {
	o.UnLockConfirm = &v
}

// GetSameAddress returns the SameAddress field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetSameAddress() bool {
	if o == nil || common.IsNil(o.SameAddress) {
		var ret bool
		return ret
	}
	return *o.SameAddress
}

// GetSameAddressOk returns a tuple with the SameAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetSameAddressOk() (*bool, bool) {
	if o == nil || common.IsNil(o.SameAddress) {
		return nil, false
	}
	return o.SameAddress, true
}

// HasSameAddress returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasSameAddress() bool {
	if o != nil && !common.IsNil(o.SameAddress) {
		return true
	}

	return false
}

// SetSameAddress gets a reference to the given bool and assigns it to the SameAddress field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetSameAddress(v bool) {
	o.SameAddress = &v
}

// GetWithdrawTag returns the WithdrawTag field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetWithdrawTag() bool {
	if o == nil || common.IsNil(o.WithdrawTag) {
		var ret bool
		return ret
	}
	return *o.WithdrawTag
}

// GetWithdrawTagOk returns a tuple with the WithdrawTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetWithdrawTagOk() (*bool, bool) {
	if o == nil || common.IsNil(o.WithdrawTag) {
		return nil, false
	}
	return o.WithdrawTag, true
}

// HasWithdrawTag returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasWithdrawTag() bool {
	if o != nil && !common.IsNil(o.WithdrawTag) {
		return true
	}

	return false
}

// SetWithdrawTag gets a reference to the given bool and assigns it to the WithdrawTag field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetWithdrawTag(v bool) {
	o.WithdrawTag = &v
}

// GetEstimatedArrivalTime returns the EstimatedArrivalTime field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetEstimatedArrivalTime() int64 {
	if o == nil || common.IsNil(o.EstimatedArrivalTime) {
		var ret int64
		return ret
	}
	return *o.EstimatedArrivalTime
}

// GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetEstimatedArrivalTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EstimatedArrivalTime) {
		return nil, false
	}
	return o.EstimatedArrivalTime, true
}

// HasEstimatedArrivalTime returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasEstimatedArrivalTime() bool {
	if o != nil && !common.IsNil(o.EstimatedArrivalTime) {
		return true
	}

	return false
}

// SetEstimatedArrivalTime gets a reference to the given int64 and assigns it to the EstimatedArrivalTime field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetEstimatedArrivalTime(v int64) {
	o.EstimatedArrivalTime = &v
}

// GetBusy returns the Busy field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetBusy() bool {
	if o == nil || common.IsNil(o.Busy) {
		var ret bool
		return ret
	}
	return *o.Busy
}

// GetBusyOk returns a tuple with the Busy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetBusyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Busy) {
		return nil, false
	}
	return o.Busy, true
}

// HasBusy returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasBusy() bool {
	if o != nil && !common.IsNil(o.Busy) {
		return true
	}

	return false
}

// SetBusy gets a reference to the given bool and assigns it to the Busy field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetBusy(v bool) {
	o.Busy = &v
}

// GetContractAddressUrl returns the ContractAddressUrl field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetContractAddressUrl() string {
	if o == nil || common.IsNil(o.ContractAddressUrl) {
		var ret string
		return ret
	}
	return *o.ContractAddressUrl
}

// GetContractAddressUrlOk returns a tuple with the ContractAddressUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetContractAddressUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContractAddressUrl) {
		return nil, false
	}
	return o.ContractAddressUrl, true
}

// HasContractAddressUrl returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasContractAddressUrl() bool {
	if o != nil && !common.IsNil(o.ContractAddressUrl) {
		return true
	}

	return false
}

// SetContractAddressUrl gets a reference to the given string and assigns it to the ContractAddressUrl field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetContractAddressUrl(v string) {
	o.ContractAddressUrl = &v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetContractAddress() string {
	if o == nil || common.IsNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetContractAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContractAddress) {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasContractAddress() bool {
	if o != nil && !common.IsNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetDenomination returns the Denomination field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetDenomination() int64 {
	if o == nil || common.IsNil(o.Denomination) {
		var ret int64
		return ret
	}
	return *o.Denomination
}

// GetDenominationOk returns a tuple with the Denomination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) GetDenominationOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Denomination) {
		return nil, false
	}
	return o.Denomination, true
}

// HasDenomination returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInnerNetworkListInner) HasDenomination() bool {
	if o != nil && !common.IsNil(o.Denomination) {
		return true
	}

	return false
}

// SetDenomination gets a reference to the given int64 and assigns it to the Denomination field.
func (o *AllCoinsInformationResponseInnerNetworkListInner) SetDenomination(v int64) {
	o.Denomination = &v
}

func (o AllCoinsInformationResponseInnerNetworkListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllCoinsInformationResponseInnerNetworkListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !common.IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !common.IsNil(o.WithdrawIntegerMultiple) {
		toSerialize["withdrawIntegerMultiple"] = o.WithdrawIntegerMultiple
	}
	if !common.IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !common.IsNil(o.DepositEnable) {
		toSerialize["depositEnable"] = o.DepositEnable
	}
	if !common.IsNil(o.WithdrawEnable) {
		toSerialize["withdrawEnable"] = o.WithdrawEnable
	}
	if !common.IsNil(o.DepositDesc) {
		toSerialize["depositDesc"] = o.DepositDesc
	}
	if !common.IsNil(o.WithdrawDesc) {
		toSerialize["withdrawDesc"] = o.WithdrawDesc
	}
	if !common.IsNil(o.SpecialTips) {
		toSerialize["specialTips"] = o.SpecialTips
	}
	if !common.IsNil(o.SpecialWithdrawTips) {
		toSerialize["specialWithdrawTips"] = o.SpecialWithdrawTips
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.ResetAddressStatus) {
		toSerialize["resetAddressStatus"] = o.ResetAddressStatus
	}
	if !common.IsNil(o.AddressRegex) {
		toSerialize["addressRegex"] = o.AddressRegex
	}
	if !common.IsNil(o.MemoRegex) {
		toSerialize["memoRegex"] = o.MemoRegex
	}
	if !common.IsNil(o.WithdrawFee) {
		toSerialize["withdrawFee"] = o.WithdrawFee
	}
	if !common.IsNil(o.WithdrawMin) {
		toSerialize["withdrawMin"] = o.WithdrawMin
	}
	if !common.IsNil(o.WithdrawMax) {
		toSerialize["withdrawMax"] = o.WithdrawMax
	}
	if !common.IsNil(o.WithdrawInternalMin) {
		toSerialize["withdrawInternalMin"] = o.WithdrawInternalMin
	}
	if !common.IsNil(o.DepositDust) {
		toSerialize["depositDust"] = o.DepositDust
	}
	if !common.IsNil(o.MinConfirm) {
		toSerialize["minConfirm"] = o.MinConfirm
	}
	if !common.IsNil(o.UnLockConfirm) {
		toSerialize["unLockConfirm"] = o.UnLockConfirm
	}
	if !common.IsNil(o.SameAddress) {
		toSerialize["sameAddress"] = o.SameAddress
	}
	if !common.IsNil(o.WithdrawTag) {
		toSerialize["withdrawTag"] = o.WithdrawTag
	}
	if !common.IsNil(o.EstimatedArrivalTime) {
		toSerialize["estimatedArrivalTime"] = o.EstimatedArrivalTime
	}
	if !common.IsNil(o.Busy) {
		toSerialize["busy"] = o.Busy
	}
	if !common.IsNil(o.ContractAddressUrl) {
		toSerialize["contractAddressUrl"] = o.ContractAddressUrl
	}
	if !common.IsNil(o.ContractAddress) {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	if !common.IsNil(o.Denomination) {
		toSerialize["denomination"] = o.Denomination
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AllCoinsInformationResponseInnerNetworkListInner) UnmarshalJSON(data []byte) (err error) {
	varAllCoinsInformationResponseInnerNetworkListInner := _AllCoinsInformationResponseInnerNetworkListInner{}

	err = json.Unmarshal(data, &varAllCoinsInformationResponseInnerNetworkListInner)

	if err != nil {
		return err
	}

	*o = AllCoinsInformationResponseInnerNetworkListInner(varAllCoinsInformationResponseInnerNetworkListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "network")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "withdrawIntegerMultiple")
		delete(additionalProperties, "isDefault")
		delete(additionalProperties, "depositEnable")
		delete(additionalProperties, "withdrawEnable")
		delete(additionalProperties, "depositDesc")
		delete(additionalProperties, "withdrawDesc")
		delete(additionalProperties, "specialTips")
		delete(additionalProperties, "specialWithdrawTips")
		delete(additionalProperties, "name")
		delete(additionalProperties, "resetAddressStatus")
		delete(additionalProperties, "addressRegex")
		delete(additionalProperties, "memoRegex")
		delete(additionalProperties, "withdrawFee")
		delete(additionalProperties, "withdrawMin")
		delete(additionalProperties, "withdrawMax")
		delete(additionalProperties, "withdrawInternalMin")
		delete(additionalProperties, "depositDust")
		delete(additionalProperties, "minConfirm")
		delete(additionalProperties, "unLockConfirm")
		delete(additionalProperties, "sameAddress")
		delete(additionalProperties, "withdrawTag")
		delete(additionalProperties, "estimatedArrivalTime")
		delete(additionalProperties, "busy")
		delete(additionalProperties, "contractAddressUrl")
		delete(additionalProperties, "contractAddress")
		delete(additionalProperties, "denomination")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAllCoinsInformationResponseInnerNetworkListInner struct {
	value *AllCoinsInformationResponseInnerNetworkListInner
	isSet bool
}

func (v NullableAllCoinsInformationResponseInnerNetworkListInner) Get() *AllCoinsInformationResponseInnerNetworkListInner {
	return v.value
}

func (v *NullableAllCoinsInformationResponseInnerNetworkListInner) Set(val *AllCoinsInformationResponseInnerNetworkListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAllCoinsInformationResponseInnerNetworkListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAllCoinsInformationResponseInnerNetworkListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllCoinsInformationResponseInnerNetworkListInner(val *AllCoinsInformationResponseInnerNetworkListInner) *NullableAllCoinsInformationResponseInnerNetworkListInner {
	return &NullableAllCoinsInformationResponseInnerNetworkListInner{value: val, isSet: true}
}

func (v NullableAllCoinsInformationResponseInnerNetworkListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllCoinsInformationResponseInnerNetworkListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
