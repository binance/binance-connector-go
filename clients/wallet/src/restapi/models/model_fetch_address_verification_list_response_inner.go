/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FetchAddressVerificationListResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FetchAddressVerificationListResponseInner{}

// FetchAddressVerificationListResponseInner struct for FetchAddressVerificationListResponseInner
type FetchAddressVerificationListResponseInner struct {
	Status               *string                                                        `json:"status,omitempty"`
	Token                *string                                                        `json:"token,omitempty"`
	Network              *string                                                        `json:"network,omitempty"`
	WalletAddress        *string                                                        `json:"walletAddress,omitempty"`
	AddressQuestionnaire *FetchAddressVerificationListResponseInnerAddressQuestionnaire `json:"addressQuestionnaire,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FetchAddressVerificationListResponseInner FetchAddressVerificationListResponseInner

// NewFetchAddressVerificationListResponseInner instantiates a new FetchAddressVerificationListResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchAddressVerificationListResponseInner() *FetchAddressVerificationListResponseInner {
	this := FetchAddressVerificationListResponseInner{}
	return &this
}

// NewFetchAddressVerificationListResponseInnerWithDefaults instantiates a new FetchAddressVerificationListResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchAddressVerificationListResponseInnerWithDefaults() *FetchAddressVerificationListResponseInner {
	this := FetchAddressVerificationListResponseInner{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FetchAddressVerificationListResponseInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchAddressVerificationListResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FetchAddressVerificationListResponseInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FetchAddressVerificationListResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *FetchAddressVerificationListResponseInner) GetToken() string {
	if o == nil || common.IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchAddressVerificationListResponseInner) GetTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *FetchAddressVerificationListResponseInner) HasToken() bool {
	if o != nil && !common.IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *FetchAddressVerificationListResponseInner) SetToken(v string) {
	o.Token = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *FetchAddressVerificationListResponseInner) GetNetwork() string {
	if o == nil || common.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchAddressVerificationListResponseInner) GetNetworkOk() (*string, bool) {
	if o == nil || common.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *FetchAddressVerificationListResponseInner) HasNetwork() bool {
	if o != nil && !common.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *FetchAddressVerificationListResponseInner) SetNetwork(v string) {
	o.Network = &v
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise.
func (o *FetchAddressVerificationListResponseInner) GetWalletAddress() string {
	if o == nil || common.IsNil(o.WalletAddress) {
		var ret string
		return ret
	}
	return *o.WalletAddress
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchAddressVerificationListResponseInner) GetWalletAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletAddress) {
		return nil, false
	}
	return o.WalletAddress, true
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *FetchAddressVerificationListResponseInner) HasWalletAddress() bool {
	if o != nil && !common.IsNil(o.WalletAddress) {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given string and assigns it to the WalletAddress field.
func (o *FetchAddressVerificationListResponseInner) SetWalletAddress(v string) {
	o.WalletAddress = &v
}

// GetAddressQuestionnaire returns the AddressQuestionnaire field value if set, zero value otherwise.
func (o *FetchAddressVerificationListResponseInner) GetAddressQuestionnaire() FetchAddressVerificationListResponseInnerAddressQuestionnaire {
	if o == nil || common.IsNil(o.AddressQuestionnaire) {
		var ret FetchAddressVerificationListResponseInnerAddressQuestionnaire
		return ret
	}
	return *o.AddressQuestionnaire
}

// GetAddressQuestionnaireOk returns a tuple with the AddressQuestionnaire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchAddressVerificationListResponseInner) GetAddressQuestionnaireOk() (*FetchAddressVerificationListResponseInnerAddressQuestionnaire, bool) {
	if o == nil || common.IsNil(o.AddressQuestionnaire) {
		return nil, false
	}
	return o.AddressQuestionnaire, true
}

// HasAddressQuestionnaire returns a boolean if a field has been set.
func (o *FetchAddressVerificationListResponseInner) HasAddressQuestionnaire() bool {
	if o != nil && !common.IsNil(o.AddressQuestionnaire) {
		return true
	}

	return false
}

// SetAddressQuestionnaire gets a reference to the given FetchAddressVerificationListResponseInnerAddressQuestionnaire and assigns it to the AddressQuestionnaire field.
func (o *FetchAddressVerificationListResponseInner) SetAddressQuestionnaire(v FetchAddressVerificationListResponseInnerAddressQuestionnaire) {
	o.AddressQuestionnaire = &v
}

func (o FetchAddressVerificationListResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchAddressVerificationListResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !common.IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !common.IsNil(o.WalletAddress) {
		toSerialize["walletAddress"] = o.WalletAddress
	}
	if !common.IsNil(o.AddressQuestionnaire) {
		toSerialize["addressQuestionnaire"] = o.AddressQuestionnaire
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FetchAddressVerificationListResponseInner) UnmarshalJSON(data []byte) (err error) {
	varFetchAddressVerificationListResponseInner := _FetchAddressVerificationListResponseInner{}

	err = json.Unmarshal(data, &varFetchAddressVerificationListResponseInner)

	if err != nil {
		return err
	}

	*o = FetchAddressVerificationListResponseInner(varFetchAddressVerificationListResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "token")
		delete(additionalProperties, "network")
		delete(additionalProperties, "walletAddress")
		delete(additionalProperties, "addressQuestionnaire")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFetchAddressVerificationListResponseInner struct {
	value *FetchAddressVerificationListResponseInner
	isSet bool
}

func (v NullableFetchAddressVerificationListResponseInner) Get() *FetchAddressVerificationListResponseInner {
	return v.value
}

func (v *NullableFetchAddressVerificationListResponseInner) Set(val *FetchAddressVerificationListResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchAddressVerificationListResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchAddressVerificationListResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchAddressVerificationListResponseInner(val *FetchAddressVerificationListResponseInner) *NullableFetchAddressVerificationListResponseInner {
	return &NullableFetchAddressVerificationListResponseInner{value: val, isSet: true}
}

func (v NullableFetchAddressVerificationListResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchAddressVerificationListResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
