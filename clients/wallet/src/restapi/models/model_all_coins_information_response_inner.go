/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AllCoinsInformationResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllCoinsInformationResponseInner{}

// AllCoinsInformationResponseInner struct for AllCoinsInformationResponseInner
type AllCoinsInformationResponseInner struct {
	Coin                 *string                                            `json:"coin,omitempty"`
	DepositAllEnable     *bool                                              `json:"depositAllEnable,omitempty"`
	WithdrawAllEnable    *bool                                              `json:"withdrawAllEnable,omitempty"`
	Name                 *string                                            `json:"name,omitempty"`
	Free                 *string                                            `json:"free,omitempty"`
	Locked               *string                                            `json:"locked,omitempty"`
	Freeze               *string                                            `json:"freeze,omitempty"`
	Withdrawing          *string                                            `json:"withdrawing,omitempty"`
	Ipoing               *string                                            `json:"ipoing,omitempty"`
	Ipoable              *string                                            `json:"ipoable,omitempty"`
	Storage              *string                                            `json:"storage,omitempty"`
	IsLegalMoney         *bool                                              `json:"isLegalMoney,omitempty"`
	Trading              *bool                                              `json:"trading,omitempty"`
	NetworkList          []AllCoinsInformationResponseInnerNetworkListInner `json:"networkList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AllCoinsInformationResponseInner AllCoinsInformationResponseInner

// NewAllCoinsInformationResponseInner instantiates a new AllCoinsInformationResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllCoinsInformationResponseInner() *AllCoinsInformationResponseInner {
	this := AllCoinsInformationResponseInner{}
	return &this
}

// NewAllCoinsInformationResponseInnerWithDefaults instantiates a new AllCoinsInformationResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllCoinsInformationResponseInnerWithDefaults() *AllCoinsInformationResponseInner {
	this := AllCoinsInformationResponseInner{}
	return &this
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInner) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInner) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInner) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *AllCoinsInformationResponseInner) SetCoin(v string) {
	o.Coin = &v
}

// GetDepositAllEnable returns the DepositAllEnable field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInner) GetDepositAllEnable() bool {
	if o == nil || common.IsNil(o.DepositAllEnable) {
		var ret bool
		return ret
	}
	return *o.DepositAllEnable
}

// GetDepositAllEnableOk returns a tuple with the DepositAllEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInner) GetDepositAllEnableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.DepositAllEnable) {
		return nil, false
	}
	return o.DepositAllEnable, true
}

// HasDepositAllEnable returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInner) HasDepositAllEnable() bool {
	if o != nil && !common.IsNil(o.DepositAllEnable) {
		return true
	}

	return false
}

// SetDepositAllEnable gets a reference to the given bool and assigns it to the DepositAllEnable field.
func (o *AllCoinsInformationResponseInner) SetDepositAllEnable(v bool) {
	o.DepositAllEnable = &v
}

// GetWithdrawAllEnable returns the WithdrawAllEnable field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInner) GetWithdrawAllEnable() bool {
	if o == nil || common.IsNil(o.WithdrawAllEnable) {
		var ret bool
		return ret
	}
	return *o.WithdrawAllEnable
}

// GetWithdrawAllEnableOk returns a tuple with the WithdrawAllEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInner) GetWithdrawAllEnableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.WithdrawAllEnable) {
		return nil, false
	}
	return o.WithdrawAllEnable, true
}

// HasWithdrawAllEnable returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInner) HasWithdrawAllEnable() bool {
	if o != nil && !common.IsNil(o.WithdrawAllEnable) {
		return true
	}

	return false
}

// SetWithdrawAllEnable gets a reference to the given bool and assigns it to the WithdrawAllEnable field.
func (o *AllCoinsInformationResponseInner) SetWithdrawAllEnable(v bool) {
	o.WithdrawAllEnable = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInner) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInner) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInner) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AllCoinsInformationResponseInner) SetName(v string) {
	o.Name = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInner) GetFree() string {
	if o == nil || common.IsNil(o.Free) {
		var ret string
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInner) GetFreeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInner) HasFree() bool {
	if o != nil && !common.IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given string and assigns it to the Free field.
func (o *AllCoinsInformationResponseInner) SetFree(v string) {
	o.Free = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInner) GetLocked() string {
	if o == nil || common.IsNil(o.Locked) {
		var ret string
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInner) GetLockedOk() (*string, bool) {
	if o == nil || common.IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInner) HasLocked() bool {
	if o != nil && !common.IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given string and assigns it to the Locked field.
func (o *AllCoinsInformationResponseInner) SetLocked(v string) {
	o.Locked = &v
}

// GetFreeze returns the Freeze field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInner) GetFreeze() string {
	if o == nil || common.IsNil(o.Freeze) {
		var ret string
		return ret
	}
	return *o.Freeze
}

// GetFreezeOk returns a tuple with the Freeze field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInner) GetFreezeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Freeze) {
		return nil, false
	}
	return o.Freeze, true
}

// HasFreeze returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInner) HasFreeze() bool {
	if o != nil && !common.IsNil(o.Freeze) {
		return true
	}

	return false
}

// SetFreeze gets a reference to the given string and assigns it to the Freeze field.
func (o *AllCoinsInformationResponseInner) SetFreeze(v string) {
	o.Freeze = &v
}

// GetWithdrawing returns the Withdrawing field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInner) GetWithdrawing() string {
	if o == nil || common.IsNil(o.Withdrawing) {
		var ret string
		return ret
	}
	return *o.Withdrawing
}

// GetWithdrawingOk returns a tuple with the Withdrawing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInner) GetWithdrawingOk() (*string, bool) {
	if o == nil || common.IsNil(o.Withdrawing) {
		return nil, false
	}
	return o.Withdrawing, true
}

// HasWithdrawing returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInner) HasWithdrawing() bool {
	if o != nil && !common.IsNil(o.Withdrawing) {
		return true
	}

	return false
}

// SetWithdrawing gets a reference to the given string and assigns it to the Withdrawing field.
func (o *AllCoinsInformationResponseInner) SetWithdrawing(v string) {
	o.Withdrawing = &v
}

// GetIpoing returns the Ipoing field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInner) GetIpoing() string {
	if o == nil || common.IsNil(o.Ipoing) {
		var ret string
		return ret
	}
	return *o.Ipoing
}

// GetIpoingOk returns a tuple with the Ipoing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInner) GetIpoingOk() (*string, bool) {
	if o == nil || common.IsNil(o.Ipoing) {
		return nil, false
	}
	return o.Ipoing, true
}

// HasIpoing returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInner) HasIpoing() bool {
	if o != nil && !common.IsNil(o.Ipoing) {
		return true
	}

	return false
}

// SetIpoing gets a reference to the given string and assigns it to the Ipoing field.
func (o *AllCoinsInformationResponseInner) SetIpoing(v string) {
	o.Ipoing = &v
}

// GetIpoable returns the Ipoable field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInner) GetIpoable() string {
	if o == nil || common.IsNil(o.Ipoable) {
		var ret string
		return ret
	}
	return *o.Ipoable
}

// GetIpoableOk returns a tuple with the Ipoable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInner) GetIpoableOk() (*string, bool) {
	if o == nil || common.IsNil(o.Ipoable) {
		return nil, false
	}
	return o.Ipoable, true
}

// HasIpoable returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInner) HasIpoable() bool {
	if o != nil && !common.IsNil(o.Ipoable) {
		return true
	}

	return false
}

// SetIpoable gets a reference to the given string and assigns it to the Ipoable field.
func (o *AllCoinsInformationResponseInner) SetIpoable(v string) {
	o.Ipoable = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInner) GetStorage() string {
	if o == nil || common.IsNil(o.Storage) {
		var ret string
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInner) GetStorageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInner) HasStorage() bool {
	if o != nil && !common.IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given string and assigns it to the Storage field.
func (o *AllCoinsInformationResponseInner) SetStorage(v string) {
	o.Storage = &v
}

// GetIsLegalMoney returns the IsLegalMoney field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInner) GetIsLegalMoney() bool {
	if o == nil || common.IsNil(o.IsLegalMoney) {
		var ret bool
		return ret
	}
	return *o.IsLegalMoney
}

// GetIsLegalMoneyOk returns a tuple with the IsLegalMoney field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInner) GetIsLegalMoneyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsLegalMoney) {
		return nil, false
	}
	return o.IsLegalMoney, true
}

// HasIsLegalMoney returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInner) HasIsLegalMoney() bool {
	if o != nil && !common.IsNil(o.IsLegalMoney) {
		return true
	}

	return false
}

// SetIsLegalMoney gets a reference to the given bool and assigns it to the IsLegalMoney field.
func (o *AllCoinsInformationResponseInner) SetIsLegalMoney(v bool) {
	o.IsLegalMoney = &v
}

// GetTrading returns the Trading field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInner) GetTrading() bool {
	if o == nil || common.IsNil(o.Trading) {
		var ret bool
		return ret
	}
	return *o.Trading
}

// GetTradingOk returns a tuple with the Trading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInner) GetTradingOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Trading) {
		return nil, false
	}
	return o.Trading, true
}

// HasTrading returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInner) HasTrading() bool {
	if o != nil && !common.IsNil(o.Trading) {
		return true
	}

	return false
}

// SetTrading gets a reference to the given bool and assigns it to the Trading field.
func (o *AllCoinsInformationResponseInner) SetTrading(v bool) {
	o.Trading = &v
}

// GetNetworkList returns the NetworkList field value if set, zero value otherwise.
func (o *AllCoinsInformationResponseInner) GetNetworkList() []AllCoinsInformationResponseInnerNetworkListInner {
	if o == nil || common.IsNil(o.NetworkList) {
		var ret []AllCoinsInformationResponseInnerNetworkListInner
		return ret
	}
	return o.NetworkList
}

// GetNetworkListOk returns a tuple with the NetworkList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllCoinsInformationResponseInner) GetNetworkListOk() ([]AllCoinsInformationResponseInnerNetworkListInner, bool) {
	if o == nil || common.IsNil(o.NetworkList) {
		return nil, false
	}
	return o.NetworkList, true
}

// HasNetworkList returns a boolean if a field has been set.
func (o *AllCoinsInformationResponseInner) HasNetworkList() bool {
	if o != nil && !common.IsNil(o.NetworkList) {
		return true
	}

	return false
}

// SetNetworkList gets a reference to the given []AllCoinsInformationResponseInnerNetworkListInner and assigns it to the NetworkList field.
func (o *AllCoinsInformationResponseInner) SetNetworkList(v []AllCoinsInformationResponseInnerNetworkListInner) {
	o.NetworkList = v
}

func (o AllCoinsInformationResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllCoinsInformationResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !common.IsNil(o.DepositAllEnable) {
		toSerialize["depositAllEnable"] = o.DepositAllEnable
	}
	if !common.IsNil(o.WithdrawAllEnable) {
		toSerialize["withdrawAllEnable"] = o.WithdrawAllEnable
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Free) {
		toSerialize["free"] = o.Free
	}
	if !common.IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !common.IsNil(o.Freeze) {
		toSerialize["freeze"] = o.Freeze
	}
	if !common.IsNil(o.Withdrawing) {
		toSerialize["withdrawing"] = o.Withdrawing
	}
	if !common.IsNil(o.Ipoing) {
		toSerialize["ipoing"] = o.Ipoing
	}
	if !common.IsNil(o.Ipoable) {
		toSerialize["ipoable"] = o.Ipoable
	}
	if !common.IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	if !common.IsNil(o.IsLegalMoney) {
		toSerialize["isLegalMoney"] = o.IsLegalMoney
	}
	if !common.IsNil(o.Trading) {
		toSerialize["trading"] = o.Trading
	}
	if !common.IsNil(o.NetworkList) {
		toSerialize["networkList"] = o.NetworkList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AllCoinsInformationResponseInner) UnmarshalJSON(data []byte) (err error) {
	varAllCoinsInformationResponseInner := _AllCoinsInformationResponseInner{}

	err = json.Unmarshal(data, &varAllCoinsInformationResponseInner)

	if err != nil {
		return err
	}

	*o = AllCoinsInformationResponseInner(varAllCoinsInformationResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "coin")
		delete(additionalProperties, "depositAllEnable")
		delete(additionalProperties, "withdrawAllEnable")
		delete(additionalProperties, "name")
		delete(additionalProperties, "free")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "freeze")
		delete(additionalProperties, "withdrawing")
		delete(additionalProperties, "ipoing")
		delete(additionalProperties, "ipoable")
		delete(additionalProperties, "storage")
		delete(additionalProperties, "isLegalMoney")
		delete(additionalProperties, "trading")
		delete(additionalProperties, "networkList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAllCoinsInformationResponseInner struct {
	value *AllCoinsInformationResponseInner
	isSet bool
}

func (v NullableAllCoinsInformationResponseInner) Get() *AllCoinsInformationResponseInner {
	return v.value
}

func (v *NullableAllCoinsInformationResponseInner) Set(val *AllCoinsInformationResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAllCoinsInformationResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAllCoinsInformationResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllCoinsInformationResponseInner(val *AllCoinsInformationResponseInner) *NullableAllCoinsInformationResponseInner {
	return &NullableAllCoinsInformationResponseInner{value: val, isSet: true}
}

func (v NullableAllCoinsInformationResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllCoinsInformationResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
