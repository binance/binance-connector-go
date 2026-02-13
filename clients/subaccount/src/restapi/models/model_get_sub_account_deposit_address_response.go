/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetSubAccountDepositAddressResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSubAccountDepositAddressResponse{}

// GetSubAccountDepositAddressResponse struct for GetSubAccountDepositAddressResponse
type GetSubAccountDepositAddressResponse struct {
	Address              *string `json:"address,omitempty"`
	Coin                 *string `json:"coin,omitempty"`
	Tag                  *string `json:"tag,omitempty"`
	Url                  *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSubAccountDepositAddressResponse GetSubAccountDepositAddressResponse

// NewGetSubAccountDepositAddressResponse instantiates a new GetSubAccountDepositAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubAccountDepositAddressResponse() *GetSubAccountDepositAddressResponse {
	this := GetSubAccountDepositAddressResponse{}
	return &this
}

// NewGetSubAccountDepositAddressResponseWithDefaults instantiates a new GetSubAccountDepositAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubAccountDepositAddressResponseWithDefaults() *GetSubAccountDepositAddressResponse {
	this := GetSubAccountDepositAddressResponse{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GetSubAccountDepositAddressResponse) GetAddress() string {
	if o == nil || common.IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositAddressResponse) GetAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GetSubAccountDepositAddressResponse) HasAddress() bool {
	if o != nil && !common.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GetSubAccountDepositAddressResponse) SetAddress(v string) {
	o.Address = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *GetSubAccountDepositAddressResponse) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositAddressResponse) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *GetSubAccountDepositAddressResponse) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *GetSubAccountDepositAddressResponse) SetCoin(v string) {
	o.Coin = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GetSubAccountDepositAddressResponse) GetTag() string {
	if o == nil || common.IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositAddressResponse) GetTagOk() (*string, bool) {
	if o == nil || common.IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GetSubAccountDepositAddressResponse) HasTag() bool {
	if o != nil && !common.IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GetSubAccountDepositAddressResponse) SetTag(v string) {
	o.Tag = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetSubAccountDepositAddressResponse) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositAddressResponse) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetSubAccountDepositAddressResponse) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetSubAccountDepositAddressResponse) SetUrl(v string) {
	o.Url = &v
}

func (o GetSubAccountDepositAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSubAccountDepositAddressResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !common.IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
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

func (o *GetSubAccountDepositAddressResponse) UnmarshalJSON(data []byte) (err error) {
	varGetSubAccountDepositAddressResponse := _GetSubAccountDepositAddressResponse{}

	err = json.Unmarshal(data, &varGetSubAccountDepositAddressResponse)

	if err != nil {
		return err
	}

	*o = GetSubAccountDepositAddressResponse(varGetSubAccountDepositAddressResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "tag")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSubAccountDepositAddressResponse struct {
	value *GetSubAccountDepositAddressResponse
	isSet bool
}

func (v NullableGetSubAccountDepositAddressResponse) Get() *GetSubAccountDepositAddressResponse {
	return v.value
}

func (v *NullableGetSubAccountDepositAddressResponse) Set(val *GetSubAccountDepositAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubAccountDepositAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubAccountDepositAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubAccountDepositAddressResponse(val *GetSubAccountDepositAddressResponse) *NullableGetSubAccountDepositAddressResponse {
	return &NullableGetSubAccountDepositAddressResponse{value: val, isSet: true}
}

func (v NullableGetSubAccountDepositAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubAccountDepositAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
