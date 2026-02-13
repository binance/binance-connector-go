/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryManagedSubAccountAssetDetailsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountAssetDetailsResponseInner{}

// QueryManagedSubAccountAssetDetailsResponseInner struct for QueryManagedSubAccountAssetDetailsResponseInner
type QueryManagedSubAccountAssetDetailsResponseInner struct {
	Coin                 *string `json:"coin,omitempty"`
	Name                 *string `json:"name,omitempty"`
	TotalBalance         *string `json:"totalBalance,omitempty"`
	AvailableBalance     *string `json:"availableBalance,omitempty"`
	InOrder              *string `json:"inOrder,omitempty"`
	BtcValue             *string `json:"btcValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryManagedSubAccountAssetDetailsResponseInner QueryManagedSubAccountAssetDetailsResponseInner

// NewQueryManagedSubAccountAssetDetailsResponseInner instantiates a new QueryManagedSubAccountAssetDetailsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountAssetDetailsResponseInner() *QueryManagedSubAccountAssetDetailsResponseInner {
	this := QueryManagedSubAccountAssetDetailsResponseInner{}
	return &this
}

// NewQueryManagedSubAccountAssetDetailsResponseInnerWithDefaults instantiates a new QueryManagedSubAccountAssetDetailsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountAssetDetailsResponseInnerWithDefaults() *QueryManagedSubAccountAssetDetailsResponseInner {
	this := QueryManagedSubAccountAssetDetailsResponseInner{}
	return &this
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) SetCoin(v string) {
	o.Coin = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) SetName(v string) {
	o.Name = &v
}

// GetTotalBalance returns the TotalBalance field value if set, zero value otherwise.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) GetTotalBalance() string {
	if o == nil || common.IsNil(o.TotalBalance) {
		var ret string
		return ret
	}
	return *o.TotalBalance
}

// GetTotalBalanceOk returns a tuple with the TotalBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) GetTotalBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalBalance) {
		return nil, false
	}
	return o.TotalBalance, true
}

// HasTotalBalance returns a boolean if a field has been set.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) HasTotalBalance() bool {
	if o != nil && !common.IsNil(o.TotalBalance) {
		return true
	}

	return false
}

// SetTotalBalance gets a reference to the given string and assigns it to the TotalBalance field.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) SetTotalBalance(v string) {
	o.TotalBalance = &v
}

// GetAvailableBalance returns the AvailableBalance field value if set, zero value otherwise.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) GetAvailableBalance() string {
	if o == nil || common.IsNil(o.AvailableBalance) {
		var ret string
		return ret
	}
	return *o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) GetAvailableBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvailableBalance) {
		return nil, false
	}
	return o.AvailableBalance, true
}

// HasAvailableBalance returns a boolean if a field has been set.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) HasAvailableBalance() bool {
	if o != nil && !common.IsNil(o.AvailableBalance) {
		return true
	}

	return false
}

// SetAvailableBalance gets a reference to the given string and assigns it to the AvailableBalance field.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) SetAvailableBalance(v string) {
	o.AvailableBalance = &v
}

// GetInOrder returns the InOrder field value if set, zero value otherwise.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) GetInOrder() string {
	if o == nil || common.IsNil(o.InOrder) {
		var ret string
		return ret
	}
	return *o.InOrder
}

// GetInOrderOk returns a tuple with the InOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) GetInOrderOk() (*string, bool) {
	if o == nil || common.IsNil(o.InOrder) {
		return nil, false
	}
	return o.InOrder, true
}

// HasInOrder returns a boolean if a field has been set.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) HasInOrder() bool {
	if o != nil && !common.IsNil(o.InOrder) {
		return true
	}

	return false
}

// SetInOrder gets a reference to the given string and assigns it to the InOrder field.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) SetInOrder(v string) {
	o.InOrder = &v
}

// GetBtcValue returns the BtcValue field value if set, zero value otherwise.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) GetBtcValue() string {
	if o == nil || common.IsNil(o.BtcValue) {
		var ret string
		return ret
	}
	return *o.BtcValue
}

// GetBtcValueOk returns a tuple with the BtcValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) GetBtcValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.BtcValue) {
		return nil, false
	}
	return o.BtcValue, true
}

// HasBtcValue returns a boolean if a field has been set.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) HasBtcValue() bool {
	if o != nil && !common.IsNil(o.BtcValue) {
		return true
	}

	return false
}

// SetBtcValue gets a reference to the given string and assigns it to the BtcValue field.
func (o *QueryManagedSubAccountAssetDetailsResponseInner) SetBtcValue(v string) {
	o.BtcValue = &v
}

func (o QueryManagedSubAccountAssetDetailsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountAssetDetailsResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.TotalBalance) {
		toSerialize["totalBalance"] = o.TotalBalance
	}
	if !common.IsNil(o.AvailableBalance) {
		toSerialize["availableBalance"] = o.AvailableBalance
	}
	if !common.IsNil(o.InOrder) {
		toSerialize["inOrder"] = o.InOrder
	}
	if !common.IsNil(o.BtcValue) {
		toSerialize["btcValue"] = o.BtcValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryManagedSubAccountAssetDetailsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountAssetDetailsResponseInner := _QueryManagedSubAccountAssetDetailsResponseInner{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountAssetDetailsResponseInner)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountAssetDetailsResponseInner(varQueryManagedSubAccountAssetDetailsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "coin")
		delete(additionalProperties, "name")
		delete(additionalProperties, "totalBalance")
		delete(additionalProperties, "availableBalance")
		delete(additionalProperties, "inOrder")
		delete(additionalProperties, "btcValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountAssetDetailsResponseInner struct {
	value *QueryManagedSubAccountAssetDetailsResponseInner
	isSet bool
}

func (v NullableQueryManagedSubAccountAssetDetailsResponseInner) Get() *QueryManagedSubAccountAssetDetailsResponseInner {
	return v.value
}

func (v *NullableQueryManagedSubAccountAssetDetailsResponseInner) Set(val *QueryManagedSubAccountAssetDetailsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountAssetDetailsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountAssetDetailsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountAssetDetailsResponseInner(val *QueryManagedSubAccountAssetDetailsResponseInner) *NullableQueryManagedSubAccountAssetDetailsResponseInner {
	return &NullableQueryManagedSubAccountAssetDetailsResponseInner{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountAssetDetailsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountAssetDetailsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
