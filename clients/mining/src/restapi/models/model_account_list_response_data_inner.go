/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountListResponseDataInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountListResponseDataInner{}

// AccountListResponseDataInner struct for AccountListResponseDataInner
type AccountListResponseDataInner struct {
	Type                 *string                                 `json:"type,omitempty"`
	UserName             *string                                 `json:"userName,omitempty"`
	List                 []AccountListResponseDataInnerListInner `json:"list,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountListResponseDataInner AccountListResponseDataInner

// NewAccountListResponseDataInner instantiates a new AccountListResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountListResponseDataInner() *AccountListResponseDataInner {
	this := AccountListResponseDataInner{}
	return &this
}

// NewAccountListResponseDataInnerWithDefaults instantiates a new AccountListResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountListResponseDataInnerWithDefaults() *AccountListResponseDataInner {
	this := AccountListResponseDataInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccountListResponseDataInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountListResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccountListResponseDataInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccountListResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *AccountListResponseDataInner) GetUserName() string {
	if o == nil || common.IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountListResponseDataInner) GetUserNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *AccountListResponseDataInner) HasUserName() bool {
	if o != nil && !common.IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *AccountListResponseDataInner) SetUserName(v string) {
	o.UserName = &v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *AccountListResponseDataInner) GetList() []AccountListResponseDataInnerListInner {
	if o == nil || common.IsNil(o.List) {
		var ret []AccountListResponseDataInnerListInner
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountListResponseDataInner) GetListOk() ([]AccountListResponseDataInnerListInner, bool) {
	if o == nil || common.IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *AccountListResponseDataInner) HasList() bool {
	if o != nil && !common.IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []AccountListResponseDataInnerListInner and assigns it to the List field.
func (o *AccountListResponseDataInner) SetList(v []AccountListResponseDataInnerListInner) {
	o.List = v
}

func (o AccountListResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountListResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !common.IsNil(o.List) {
		toSerialize["list"] = o.List
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountListResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	varAccountListResponseDataInner := _AccountListResponseDataInner{}

	err = json.Unmarshal(data, &varAccountListResponseDataInner)

	if err != nil {
		return err
	}

	*o = AccountListResponseDataInner(varAccountListResponseDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "userName")
		delete(additionalProperties, "list")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountListResponseDataInner struct {
	value *AccountListResponseDataInner
	isSet bool
}

func (v NullableAccountListResponseDataInner) Get() *AccountListResponseDataInner {
	return v.value
}

func (v *NullableAccountListResponseDataInner) Set(val *AccountListResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountListResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountListResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountListResponseDataInner(val *AccountListResponseDataInner) *NullableAccountListResponseDataInner {
	return &NullableAccountListResponseDataInner{value: val, isSet: true}
}

func (v NullableAccountListResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountListResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
