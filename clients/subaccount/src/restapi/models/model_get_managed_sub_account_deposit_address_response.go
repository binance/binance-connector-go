/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetManagedSubAccountDepositAddressResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetManagedSubAccountDepositAddressResponse{}

// GetManagedSubAccountDepositAddressResponse struct for GetManagedSubAccountDepositAddressResponse
type GetManagedSubAccountDepositAddressResponse struct {
	Coin                 *string `json:"coin,omitempty"`
	Address              *string `json:"address,omitempty"`
	Tag                  *string `json:"tag,omitempty"`
	Url                  *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetManagedSubAccountDepositAddressResponse GetManagedSubAccountDepositAddressResponse

// NewGetManagedSubAccountDepositAddressResponse instantiates a new GetManagedSubAccountDepositAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetManagedSubAccountDepositAddressResponse() *GetManagedSubAccountDepositAddressResponse {
	this := GetManagedSubAccountDepositAddressResponse{}
	return &this
}

// NewGetManagedSubAccountDepositAddressResponseWithDefaults instantiates a new GetManagedSubAccountDepositAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetManagedSubAccountDepositAddressResponseWithDefaults() *GetManagedSubAccountDepositAddressResponse {
	this := GetManagedSubAccountDepositAddressResponse{}
	return &this
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *GetManagedSubAccountDepositAddressResponse) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedSubAccountDepositAddressResponse) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *GetManagedSubAccountDepositAddressResponse) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *GetManagedSubAccountDepositAddressResponse) SetCoin(v string) {
	o.Coin = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GetManagedSubAccountDepositAddressResponse) GetAddress() string {
	if o == nil || common.IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedSubAccountDepositAddressResponse) GetAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GetManagedSubAccountDepositAddressResponse) HasAddress() bool {
	if o != nil && !common.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GetManagedSubAccountDepositAddressResponse) SetAddress(v string) {
	o.Address = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GetManagedSubAccountDepositAddressResponse) GetTag() string {
	if o == nil || common.IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedSubAccountDepositAddressResponse) GetTagOk() (*string, bool) {
	if o == nil || common.IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GetManagedSubAccountDepositAddressResponse) HasTag() bool {
	if o != nil && !common.IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GetManagedSubAccountDepositAddressResponse) SetTag(v string) {
	o.Tag = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetManagedSubAccountDepositAddressResponse) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedSubAccountDepositAddressResponse) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetManagedSubAccountDepositAddressResponse) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetManagedSubAccountDepositAddressResponse) SetUrl(v string) {
	o.Url = &v
}

func (o GetManagedSubAccountDepositAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetManagedSubAccountDepositAddressResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !common.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !common.IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !common.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetManagedSubAccountDepositAddressResponse) UnmarshalJSON(data []byte) (err error) {
	varGetManagedSubAccountDepositAddressResponse := _GetManagedSubAccountDepositAddressResponse{}

	err = json.Unmarshal(data, &varGetManagedSubAccountDepositAddressResponse)

	if err != nil {
		return err
	}

	*o = GetManagedSubAccountDepositAddressResponse(varGetManagedSubAccountDepositAddressResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "coin")
		delete(additionalProperties, "address")
		delete(additionalProperties, "tag")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetManagedSubAccountDepositAddressResponse struct {
	value *GetManagedSubAccountDepositAddressResponse
	isSet bool
}

func (v NullableGetManagedSubAccountDepositAddressResponse) Get() *GetManagedSubAccountDepositAddressResponse {
	return v.value
}

func (v *NullableGetManagedSubAccountDepositAddressResponse) Set(val *GetManagedSubAccountDepositAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetManagedSubAccountDepositAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetManagedSubAccountDepositAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetManagedSubAccountDepositAddressResponse(val *GetManagedSubAccountDepositAddressResponse) *NullableGetManagedSubAccountDepositAddressResponse {
	return &NullableGetManagedSubAccountDepositAddressResponse{value: val, isSet: true}
}

func (v NullableGetManagedSubAccountDepositAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetManagedSubAccountDepositAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
