/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the FetchWithdrawAddressListResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FetchWithdrawAddressListResponseInner{}

// FetchWithdrawAddressListResponseInner struct for FetchWithdrawAddressListResponseInner
type FetchWithdrawAddressListResponseInner struct {
	Address              *string `json:"address,omitempty"`
	AddressTag           *string `json:"addressTag,omitempty"`
	Coin                 *string `json:"coin,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Network              *string `json:"network,omitempty"`
	Origin               *string `json:"origin,omitempty"`
	OriginType           *string `json:"originType,omitempty"`
	WhiteStatus          *bool   `json:"whiteStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FetchWithdrawAddressListResponseInner FetchWithdrawAddressListResponseInner

// NewFetchWithdrawAddressListResponseInner instantiates a new FetchWithdrawAddressListResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchWithdrawAddressListResponseInner() *FetchWithdrawAddressListResponseInner {
	this := FetchWithdrawAddressListResponseInner{}
	return &this
}

// NewFetchWithdrawAddressListResponseInnerWithDefaults instantiates a new FetchWithdrawAddressListResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchWithdrawAddressListResponseInnerWithDefaults() *FetchWithdrawAddressListResponseInner {
	this := FetchWithdrawAddressListResponseInner{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *FetchWithdrawAddressListResponseInner) GetAddress() string {
	if o == nil || common.IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchWithdrawAddressListResponseInner) GetAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *FetchWithdrawAddressListResponseInner) HasAddress() bool {
	if o != nil && !common.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *FetchWithdrawAddressListResponseInner) SetAddress(v string) {
	o.Address = &v
}

// GetAddressTag returns the AddressTag field value if set, zero value otherwise.
func (o *FetchWithdrawAddressListResponseInner) GetAddressTag() string {
	if o == nil || common.IsNil(o.AddressTag) {
		var ret string
		return ret
	}
	return *o.AddressTag
}

// GetAddressTagOk returns a tuple with the AddressTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchWithdrawAddressListResponseInner) GetAddressTagOk() (*string, bool) {
	if o == nil || common.IsNil(o.AddressTag) {
		return nil, false
	}
	return o.AddressTag, true
}

// HasAddressTag returns a boolean if a field has been set.
func (o *FetchWithdrawAddressListResponseInner) HasAddressTag() bool {
	if o != nil && !common.IsNil(o.AddressTag) {
		return true
	}

	return false
}

// SetAddressTag gets a reference to the given string and assigns it to the AddressTag field.
func (o *FetchWithdrawAddressListResponseInner) SetAddressTag(v string) {
	o.AddressTag = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *FetchWithdrawAddressListResponseInner) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchWithdrawAddressListResponseInner) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *FetchWithdrawAddressListResponseInner) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *FetchWithdrawAddressListResponseInner) SetCoin(v string) {
	o.Coin = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FetchWithdrawAddressListResponseInner) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchWithdrawAddressListResponseInner) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FetchWithdrawAddressListResponseInner) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FetchWithdrawAddressListResponseInner) SetName(v string) {
	o.Name = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *FetchWithdrawAddressListResponseInner) GetNetwork() string {
	if o == nil || common.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchWithdrawAddressListResponseInner) GetNetworkOk() (*string, bool) {
	if o == nil || common.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *FetchWithdrawAddressListResponseInner) HasNetwork() bool {
	if o != nil && !common.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *FetchWithdrawAddressListResponseInner) SetNetwork(v string) {
	o.Network = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *FetchWithdrawAddressListResponseInner) GetOrigin() string {
	if o == nil || common.IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchWithdrawAddressListResponseInner) GetOriginOk() (*string, bool) {
	if o == nil || common.IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *FetchWithdrawAddressListResponseInner) HasOrigin() bool {
	if o != nil && !common.IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *FetchWithdrawAddressListResponseInner) SetOrigin(v string) {
	o.Origin = &v
}

// GetOriginType returns the OriginType field value if set, zero value otherwise.
func (o *FetchWithdrawAddressListResponseInner) GetOriginType() string {
	if o == nil || common.IsNil(o.OriginType) {
		var ret string
		return ret
	}
	return *o.OriginType
}

// GetOriginTypeOk returns a tuple with the OriginType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchWithdrawAddressListResponseInner) GetOriginTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OriginType) {
		return nil, false
	}
	return o.OriginType, true
}

// HasOriginType returns a boolean if a field has been set.
func (o *FetchWithdrawAddressListResponseInner) HasOriginType() bool {
	if o != nil && !common.IsNil(o.OriginType) {
		return true
	}

	return false
}

// SetOriginType gets a reference to the given string and assigns it to the OriginType field.
func (o *FetchWithdrawAddressListResponseInner) SetOriginType(v string) {
	o.OriginType = &v
}

// GetWhiteStatus returns the WhiteStatus field value if set, zero value otherwise.
func (o *FetchWithdrawAddressListResponseInner) GetWhiteStatus() bool {
	if o == nil || common.IsNil(o.WhiteStatus) {
		var ret bool
		return ret
	}
	return *o.WhiteStatus
}

// GetWhiteStatusOk returns a tuple with the WhiteStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchWithdrawAddressListResponseInner) GetWhiteStatusOk() (*bool, bool) {
	if o == nil || common.IsNil(o.WhiteStatus) {
		return nil, false
	}
	return o.WhiteStatus, true
}

// HasWhiteStatus returns a boolean if a field has been set.
func (o *FetchWithdrawAddressListResponseInner) HasWhiteStatus() bool {
	if o != nil && !common.IsNil(o.WhiteStatus) {
		return true
	}

	return false
}

// SetWhiteStatus gets a reference to the given bool and assigns it to the WhiteStatus field.
func (o *FetchWithdrawAddressListResponseInner) SetWhiteStatus(v bool) {
	o.WhiteStatus = &v
}

func (o FetchWithdrawAddressListResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchWithdrawAddressListResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !common.IsNil(o.AddressTag) {
		toSerialize["addressTag"] = o.AddressTag
	}
	if !common.IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !common.IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !common.IsNil(o.OriginType) {
		toSerialize["originType"] = o.OriginType
	}
	if !common.IsNil(o.WhiteStatus) {
		toSerialize["whiteStatus"] = o.WhiteStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FetchWithdrawAddressListResponseInner) UnmarshalJSON(data []byte) (err error) {
	varFetchWithdrawAddressListResponseInner := _FetchWithdrawAddressListResponseInner{}

	err = json.Unmarshal(data, &varFetchWithdrawAddressListResponseInner)

	if err != nil {
		return err
	}

	*o = FetchWithdrawAddressListResponseInner(varFetchWithdrawAddressListResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "addressTag")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "name")
		delete(additionalProperties, "network")
		delete(additionalProperties, "origin")
		delete(additionalProperties, "originType")
		delete(additionalProperties, "whiteStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFetchWithdrawAddressListResponseInner struct {
	value *FetchWithdrawAddressListResponseInner
	isSet bool
}

func (v NullableFetchWithdrawAddressListResponseInner) Get() *FetchWithdrawAddressListResponseInner {
	return v.value
}

func (v *NullableFetchWithdrawAddressListResponseInner) Set(val *FetchWithdrawAddressListResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchWithdrawAddressListResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchWithdrawAddressListResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchWithdrawAddressListResponseInner(val *FetchWithdrawAddressListResponseInner) *NullableFetchWithdrawAddressListResponseInner {
	return &NullableFetchWithdrawAddressListResponseInner{value: val, isSet: true}
}

func (v NullableFetchWithdrawAddressListResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchWithdrawAddressListResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
