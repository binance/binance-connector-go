/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetDetailOnSubAccountsMarginAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDetailOnSubAccountsMarginAccountResponse{}

// GetDetailOnSubAccountsMarginAccountResponse struct for GetDetailOnSubAccountsMarginAccountResponse
type GetDetailOnSubAccountsMarginAccountResponse struct {
	Email                 *string                                                                 `json:"email,omitempty"`
	MarginLevel           *string                                                                 `json:"marginLevel,omitempty"`
	TotalAssetOfBtc       *string                                                                 `json:"totalAssetOfBtc,omitempty"`
	TotalLiabilityOfBtc   *string                                                                 `json:"totalLiabilityOfBtc,omitempty"`
	TotalNetAssetOfBtc    *string                                                                 `json:"totalNetAssetOfBtc,omitempty"`
	MarginTradeCoeffVo    *GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo          `json:"marginTradeCoeffVo,omitempty"`
	MarginUserAssetVoList []GetDetailOnSubAccountsMarginAccountResponseMarginUserAssetVoListInner `json:"marginUserAssetVoList,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _GetDetailOnSubAccountsMarginAccountResponse GetDetailOnSubAccountsMarginAccountResponse

// NewGetDetailOnSubAccountsMarginAccountResponse instantiates a new GetDetailOnSubAccountsMarginAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDetailOnSubAccountsMarginAccountResponse() *GetDetailOnSubAccountsMarginAccountResponse {
	this := GetDetailOnSubAccountsMarginAccountResponse{}
	return &this
}

// NewGetDetailOnSubAccountsMarginAccountResponseWithDefaults instantiates a new GetDetailOnSubAccountsMarginAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDetailOnSubAccountsMarginAccountResponseWithDefaults() *GetDetailOnSubAccountsMarginAccountResponse {
	this := GetDetailOnSubAccountsMarginAccountResponse{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsMarginAccountResponse) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponse) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponse) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetDetailOnSubAccountsMarginAccountResponse) SetEmail(v string) {
	o.Email = &v
}

// GetMarginLevel returns the MarginLevel field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsMarginAccountResponse) GetMarginLevel() string {
	if o == nil || common.IsNil(o.MarginLevel) {
		var ret string
		return ret
	}
	return *o.MarginLevel
}

// GetMarginLevelOk returns a tuple with the MarginLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponse) GetMarginLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginLevel) {
		return nil, false
	}
	return o.MarginLevel, true
}

// HasMarginLevel returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponse) HasMarginLevel() bool {
	if o != nil && !common.IsNil(o.MarginLevel) {
		return true
	}

	return false
}

// SetMarginLevel gets a reference to the given string and assigns it to the MarginLevel field.
func (o *GetDetailOnSubAccountsMarginAccountResponse) SetMarginLevel(v string) {
	o.MarginLevel = &v
}

// GetTotalAssetOfBtc returns the TotalAssetOfBtc field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsMarginAccountResponse) GetTotalAssetOfBtc() string {
	if o == nil || common.IsNil(o.TotalAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalAssetOfBtc
}

// GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponse) GetTotalAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAssetOfBtc) {
		return nil, false
	}
	return o.TotalAssetOfBtc, true
}

// HasTotalAssetOfBtc returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponse) HasTotalAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalAssetOfBtc gets a reference to the given string and assigns it to the TotalAssetOfBtc field.
func (o *GetDetailOnSubAccountsMarginAccountResponse) SetTotalAssetOfBtc(v string) {
	o.TotalAssetOfBtc = &v
}

// GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsMarginAccountResponse) GetTotalLiabilityOfBtc() string {
	if o == nil || common.IsNil(o.TotalLiabilityOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalLiabilityOfBtc
}

// GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponse) GetTotalLiabilityOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalLiabilityOfBtc) {
		return nil, false
	}
	return o.TotalLiabilityOfBtc, true
}

// HasTotalLiabilityOfBtc returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponse) HasTotalLiabilityOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalLiabilityOfBtc) {
		return true
	}

	return false
}

// SetTotalLiabilityOfBtc gets a reference to the given string and assigns it to the TotalLiabilityOfBtc field.
func (o *GetDetailOnSubAccountsMarginAccountResponse) SetTotalLiabilityOfBtc(v string) {
	o.TotalLiabilityOfBtc = &v
}

// GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsMarginAccountResponse) GetTotalNetAssetOfBtc() string {
	if o == nil || common.IsNil(o.TotalNetAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalNetAssetOfBtc
}

// GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponse) GetTotalNetAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalNetAssetOfBtc) {
		return nil, false
	}
	return o.TotalNetAssetOfBtc, true
}

// HasTotalNetAssetOfBtc returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponse) HasTotalNetAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalNetAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalNetAssetOfBtc gets a reference to the given string and assigns it to the TotalNetAssetOfBtc field.
func (o *GetDetailOnSubAccountsMarginAccountResponse) SetTotalNetAssetOfBtc(v string) {
	o.TotalNetAssetOfBtc = &v
}

// GetMarginTradeCoeffVo returns the MarginTradeCoeffVo field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsMarginAccountResponse) GetMarginTradeCoeffVo() GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo {
	if o == nil || common.IsNil(o.MarginTradeCoeffVo) {
		var ret GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo
		return ret
	}
	return *o.MarginTradeCoeffVo
}

// GetMarginTradeCoeffVoOk returns a tuple with the MarginTradeCoeffVo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponse) GetMarginTradeCoeffVoOk() (*GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo, bool) {
	if o == nil || common.IsNil(o.MarginTradeCoeffVo) {
		return nil, false
	}
	return o.MarginTradeCoeffVo, true
}

// HasMarginTradeCoeffVo returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponse) HasMarginTradeCoeffVo() bool {
	if o != nil && !common.IsNil(o.MarginTradeCoeffVo) {
		return true
	}

	return false
}

// SetMarginTradeCoeffVo gets a reference to the given GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo and assigns it to the MarginTradeCoeffVo field.
func (o *GetDetailOnSubAccountsMarginAccountResponse) SetMarginTradeCoeffVo(v GetDetailOnSubAccountsMarginAccountResponseMarginTradeCoeffVo) {
	o.MarginTradeCoeffVo = &v
}

// GetMarginUserAssetVoList returns the MarginUserAssetVoList field value if set, zero value otherwise.
func (o *GetDetailOnSubAccountsMarginAccountResponse) GetMarginUserAssetVoList() []GetDetailOnSubAccountsMarginAccountResponseMarginUserAssetVoListInner {
	if o == nil || common.IsNil(o.MarginUserAssetVoList) {
		var ret []GetDetailOnSubAccountsMarginAccountResponseMarginUserAssetVoListInner
		return ret
	}
	return o.MarginUserAssetVoList
}

// GetMarginUserAssetVoListOk returns a tuple with the MarginUserAssetVoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponse) GetMarginUserAssetVoListOk() ([]GetDetailOnSubAccountsMarginAccountResponseMarginUserAssetVoListInner, bool) {
	if o == nil || common.IsNil(o.MarginUserAssetVoList) {
		return nil, false
	}
	return o.MarginUserAssetVoList, true
}

// HasMarginUserAssetVoList returns a boolean if a field has been set.
func (o *GetDetailOnSubAccountsMarginAccountResponse) HasMarginUserAssetVoList() bool {
	if o != nil && !common.IsNil(o.MarginUserAssetVoList) {
		return true
	}

	return false
}

// SetMarginUserAssetVoList gets a reference to the given []GetDetailOnSubAccountsMarginAccountResponseMarginUserAssetVoListInner and assigns it to the MarginUserAssetVoList field.
func (o *GetDetailOnSubAccountsMarginAccountResponse) SetMarginUserAssetVoList(v []GetDetailOnSubAccountsMarginAccountResponseMarginUserAssetVoListInner) {
	o.MarginUserAssetVoList = v
}

func (o GetDetailOnSubAccountsMarginAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDetailOnSubAccountsMarginAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.MarginLevel) {
		toSerialize["marginLevel"] = o.MarginLevel
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
	if !common.IsNil(o.MarginTradeCoeffVo) {
		toSerialize["marginTradeCoeffVo"] = o.MarginTradeCoeffVo
	}
	if !common.IsNil(o.MarginUserAssetVoList) {
		toSerialize["marginUserAssetVoList"] = o.MarginUserAssetVoList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDetailOnSubAccountsMarginAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varGetDetailOnSubAccountsMarginAccountResponse := _GetDetailOnSubAccountsMarginAccountResponse{}

	err = json.Unmarshal(data, &varGetDetailOnSubAccountsMarginAccountResponse)

	if err != nil {
		return err
	}

	*o = GetDetailOnSubAccountsMarginAccountResponse(varGetDetailOnSubAccountsMarginAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "marginLevel")
		delete(additionalProperties, "totalAssetOfBtc")
		delete(additionalProperties, "totalLiabilityOfBtc")
		delete(additionalProperties, "totalNetAssetOfBtc")
		delete(additionalProperties, "marginTradeCoeffVo")
		delete(additionalProperties, "marginUserAssetVoList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDetailOnSubAccountsMarginAccountResponse struct {
	value *GetDetailOnSubAccountsMarginAccountResponse
	isSet bool
}

func (v NullableGetDetailOnSubAccountsMarginAccountResponse) Get() *GetDetailOnSubAccountsMarginAccountResponse {
	return v.value
}

func (v *NullableGetDetailOnSubAccountsMarginAccountResponse) Set(val *GetDetailOnSubAccountsMarginAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDetailOnSubAccountsMarginAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDetailOnSubAccountsMarginAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDetailOnSubAccountsMarginAccountResponse(val *GetDetailOnSubAccountsMarginAccountResponse) *NullableGetDetailOnSubAccountsMarginAccountResponse {
	return &NullableGetDetailOnSubAccountsMarginAccountResponse{value: val, isSet: true}
}

func (v NullableGetDetailOnSubAccountsMarginAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDetailOnSubAccountsMarginAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
