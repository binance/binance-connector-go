/*
Sub Account REST API

Create and manage sub-accounts, control permissions, and transfer assets via the Sub Account API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CreateSubAccountApiKeyResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateSubAccountApiKeyResponse{}

// CreateSubAccountApiKeyResponse struct for CreateSubAccountApiKeyResponse
type CreateSubAccountApiKeyResponse struct {
	ApiName *string `json:"apiName,omitempty"`
	ApiKey  *string `json:"apiKey,omitempty"`
	// Secret Key. Returned only once on creation, please keep it safe.
	SecretKey            *string  `json:"secretKey,omitempty"`
	CanTrade             *bool    `json:"canTrade,omitempty"`
	CanMarginLoanRepay   *bool    `json:"canMarginLoanRepay,omitempty"`
	CanFuturesTrade      *bool    `json:"canFuturesTrade,omitempty"`
	CanUniversalTransfer *bool    `json:"canUniversalTransfer,omitempty"`
	CanVanillaOptions    *bool    `json:"canVanillaOptions,omitempty"`
	Status               *int64   `json:"status,omitempty"`
	IpList               []string `json:"ipList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateSubAccountApiKeyResponse CreateSubAccountApiKeyResponse

// NewCreateSubAccountApiKeyResponse instantiates a new CreateSubAccountApiKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubAccountApiKeyResponse() *CreateSubAccountApiKeyResponse {
	this := CreateSubAccountApiKeyResponse{}
	return &this
}

// NewCreateSubAccountApiKeyResponseWithDefaults instantiates a new CreateSubAccountApiKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubAccountApiKeyResponseWithDefaults() *CreateSubAccountApiKeyResponse {
	this := CreateSubAccountApiKeyResponse{}
	return &this
}

// GetApiName returns the ApiName field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyResponse) GetApiName() string {
	if o == nil || common.IsNil(o.ApiName) {
		var ret string
		return ret
	}
	return *o.ApiName
}

// GetApiNameOk returns a tuple with the ApiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyResponse) GetApiNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.ApiName) {
		return nil, false
	}
	return o.ApiName, true
}

// HasApiName returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyResponse) HasApiName() bool {
	if o != nil && !common.IsNil(o.ApiName) {
		return true
	}

	return false
}

// SetApiName gets a reference to the given string and assigns it to the ApiName field.
func (o *CreateSubAccountApiKeyResponse) SetApiName(v string) {
	o.ApiName = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyResponse) GetApiKey() string {
	if o == nil || common.IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyResponse) GetApiKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyResponse) HasApiKey() bool {
	if o != nil && !common.IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *CreateSubAccountApiKeyResponse) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyResponse) GetSecretKey() string {
	if o == nil || common.IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyResponse) GetSecretKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyResponse) HasSecretKey() bool {
	if o != nil && !common.IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *CreateSubAccountApiKeyResponse) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetCanTrade returns the CanTrade field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyResponse) GetCanTrade() bool {
	if o == nil || common.IsNil(o.CanTrade) {
		var ret bool
		return ret
	}
	return *o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyResponse) GetCanTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanTrade) {
		return nil, false
	}
	return o.CanTrade, true
}

// HasCanTrade returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyResponse) HasCanTrade() bool {
	if o != nil && !common.IsNil(o.CanTrade) {
		return true
	}

	return false
}

// SetCanTrade gets a reference to the given bool and assigns it to the CanTrade field.
func (o *CreateSubAccountApiKeyResponse) SetCanTrade(v bool) {
	o.CanTrade = &v
}

// GetCanMarginLoanRepay returns the CanMarginLoanRepay field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyResponse) GetCanMarginLoanRepay() bool {
	if o == nil || common.IsNil(o.CanMarginLoanRepay) {
		var ret bool
		return ret
	}
	return *o.CanMarginLoanRepay
}

// GetCanMarginLoanRepayOk returns a tuple with the CanMarginLoanRepay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyResponse) GetCanMarginLoanRepayOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanMarginLoanRepay) {
		return nil, false
	}
	return o.CanMarginLoanRepay, true
}

// HasCanMarginLoanRepay returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyResponse) HasCanMarginLoanRepay() bool {
	if o != nil && !common.IsNil(o.CanMarginLoanRepay) {
		return true
	}

	return false
}

// SetCanMarginLoanRepay gets a reference to the given bool and assigns it to the CanMarginLoanRepay field.
func (o *CreateSubAccountApiKeyResponse) SetCanMarginLoanRepay(v bool) {
	o.CanMarginLoanRepay = &v
}

// GetCanFuturesTrade returns the CanFuturesTrade field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyResponse) GetCanFuturesTrade() bool {
	if o == nil || common.IsNil(o.CanFuturesTrade) {
		var ret bool
		return ret
	}
	return *o.CanFuturesTrade
}

// GetCanFuturesTradeOk returns a tuple with the CanFuturesTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyResponse) GetCanFuturesTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanFuturesTrade) {
		return nil, false
	}
	return o.CanFuturesTrade, true
}

// HasCanFuturesTrade returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyResponse) HasCanFuturesTrade() bool {
	if o != nil && !common.IsNil(o.CanFuturesTrade) {
		return true
	}

	return false
}

// SetCanFuturesTrade gets a reference to the given bool and assigns it to the CanFuturesTrade field.
func (o *CreateSubAccountApiKeyResponse) SetCanFuturesTrade(v bool) {
	o.CanFuturesTrade = &v
}

// GetCanUniversalTransfer returns the CanUniversalTransfer field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyResponse) GetCanUniversalTransfer() bool {
	if o == nil || common.IsNil(o.CanUniversalTransfer) {
		var ret bool
		return ret
	}
	return *o.CanUniversalTransfer
}

// GetCanUniversalTransferOk returns a tuple with the CanUniversalTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyResponse) GetCanUniversalTransferOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanUniversalTransfer) {
		return nil, false
	}
	return o.CanUniversalTransfer, true
}

// HasCanUniversalTransfer returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyResponse) HasCanUniversalTransfer() bool {
	if o != nil && !common.IsNil(o.CanUniversalTransfer) {
		return true
	}

	return false
}

// SetCanUniversalTransfer gets a reference to the given bool and assigns it to the CanUniversalTransfer field.
func (o *CreateSubAccountApiKeyResponse) SetCanUniversalTransfer(v bool) {
	o.CanUniversalTransfer = &v
}

// GetCanVanillaOptions returns the CanVanillaOptions field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyResponse) GetCanVanillaOptions() bool {
	if o == nil || common.IsNil(o.CanVanillaOptions) {
		var ret bool
		return ret
	}
	return *o.CanVanillaOptions
}

// GetCanVanillaOptionsOk returns a tuple with the CanVanillaOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyResponse) GetCanVanillaOptionsOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanVanillaOptions) {
		return nil, false
	}
	return o.CanVanillaOptions, true
}

// HasCanVanillaOptions returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyResponse) HasCanVanillaOptions() bool {
	if o != nil && !common.IsNil(o.CanVanillaOptions) {
		return true
	}

	return false
}

// SetCanVanillaOptions gets a reference to the given bool and assigns it to the CanVanillaOptions field.
func (o *CreateSubAccountApiKeyResponse) SetCanVanillaOptions(v bool) {
	o.CanVanillaOptions = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *CreateSubAccountApiKeyResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetIpList returns the IpList field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyResponse) GetIpList() []string {
	if o == nil || common.IsNil(o.IpList) {
		var ret []string
		return ret
	}
	return o.IpList
}

// GetIpListOk returns a tuple with the IpList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyResponse) GetIpListOk() ([]string, bool) {
	if o == nil || common.IsNil(o.IpList) {
		return nil, false
	}
	return o.IpList, true
}

// HasIpList returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyResponse) HasIpList() bool {
	if o != nil && !common.IsNil(o.IpList) {
		return true
	}

	return false
}

// SetIpList gets a reference to the given []string and assigns it to the IpList field.
func (o *CreateSubAccountApiKeyResponse) SetIpList(v []string) {
	o.IpList = v
}

func (o CreateSubAccountApiKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubAccountApiKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ApiName) {
		toSerialize["apiName"] = o.ApiName
	}
	if !common.IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	if !common.IsNil(o.SecretKey) {
		toSerialize["secretKey"] = o.SecretKey
	}
	if !common.IsNil(o.CanTrade) {
		toSerialize["canTrade"] = o.CanTrade
	}
	if !common.IsNil(o.CanMarginLoanRepay) {
		toSerialize["canMarginLoanRepay"] = o.CanMarginLoanRepay
	}
	if !common.IsNil(o.CanFuturesTrade) {
		toSerialize["canFuturesTrade"] = o.CanFuturesTrade
	}
	if !common.IsNil(o.CanUniversalTransfer) {
		toSerialize["canUniversalTransfer"] = o.CanUniversalTransfer
	}
	if !common.IsNil(o.CanVanillaOptions) {
		toSerialize["canVanillaOptions"] = o.CanVanillaOptions
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.IpList) {
		toSerialize["ipList"] = o.IpList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateSubAccountApiKeyResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateSubAccountApiKeyResponse := _CreateSubAccountApiKeyResponse{}

	err = json.Unmarshal(data, &varCreateSubAccountApiKeyResponse)

	if err != nil {
		return err
	}

	*o = CreateSubAccountApiKeyResponse(varCreateSubAccountApiKeyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apiName")
		delete(additionalProperties, "apiKey")
		delete(additionalProperties, "secretKey")
		delete(additionalProperties, "canTrade")
		delete(additionalProperties, "canMarginLoanRepay")
		delete(additionalProperties, "canFuturesTrade")
		delete(additionalProperties, "canUniversalTransfer")
		delete(additionalProperties, "canVanillaOptions")
		delete(additionalProperties, "status")
		delete(additionalProperties, "ipList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSubAccountApiKeyResponse struct {
	value *CreateSubAccountApiKeyResponse
	isSet bool
}

func (v NullableCreateSubAccountApiKeyResponse) Get() *CreateSubAccountApiKeyResponse {
	return v.value
}

func (v *NullableCreateSubAccountApiKeyResponse) Set(val *CreateSubAccountApiKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubAccountApiKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubAccountApiKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubAccountApiKeyResponse(val *CreateSubAccountApiKeyResponse) *NullableCreateSubAccountApiKeyResponse {
	return &NullableCreateSubAccountApiKeyResponse{value: val, isSet: true}
}

func (v NullableCreateSubAccountApiKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSubAccountApiKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
