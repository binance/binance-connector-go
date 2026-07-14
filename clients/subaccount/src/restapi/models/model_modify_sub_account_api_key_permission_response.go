/*
Sub Account REST API

Create and manage sub-accounts, control permissions, and transfer assets via the Sub Account API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ModifySubAccountApiKeyPermissionResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ModifySubAccountApiKeyPermissionResponse{}

// ModifySubAccountApiKeyPermissionResponse struct for ModifySubAccountApiKeyPermissionResponse
type ModifySubAccountApiKeyPermissionResponse struct {
	ApiName              *string `json:"apiName,omitempty"`
	Apikey               *string `json:"apikey,omitempty"`
	CanTrade             *bool   `json:"canTrade,omitempty"`
	CanMarginLoanRepay   *bool   `json:"canMarginLoanRepay,omitempty"`
	CanFuturesTrade      *bool   `json:"canFuturesTrade,omitempty"`
	CanUniversalTransfer *bool   `json:"canUniversalTransfer,omitempty"`
	CanVanillaOptions    *bool   `json:"canVanillaOptions,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModifySubAccountApiKeyPermissionResponse ModifySubAccountApiKeyPermissionResponse

// NewModifySubAccountApiKeyPermissionResponse instantiates a new ModifySubAccountApiKeyPermissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifySubAccountApiKeyPermissionResponse() *ModifySubAccountApiKeyPermissionResponse {
	this := ModifySubAccountApiKeyPermissionResponse{}
	return &this
}

// NewModifySubAccountApiKeyPermissionResponseWithDefaults instantiates a new ModifySubAccountApiKeyPermissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifySubAccountApiKeyPermissionResponseWithDefaults() *ModifySubAccountApiKeyPermissionResponse {
	this := ModifySubAccountApiKeyPermissionResponse{}
	return &this
}

// GetApiName returns the ApiName field value if set, zero value otherwise.
func (o *ModifySubAccountApiKeyPermissionResponse) GetApiName() string {
	if o == nil || common.IsNil(o.ApiName) {
		var ret string
		return ret
	}
	return *o.ApiName
}

// GetApiNameOk returns a tuple with the ApiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifySubAccountApiKeyPermissionResponse) GetApiNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.ApiName) {
		return nil, false
	}
	return o.ApiName, true
}

// HasApiName returns a boolean if a field has been set.
func (o *ModifySubAccountApiKeyPermissionResponse) HasApiName() bool {
	if o != nil && !common.IsNil(o.ApiName) {
		return true
	}

	return false
}

// SetApiName gets a reference to the given string and assigns it to the ApiName field.
func (o *ModifySubAccountApiKeyPermissionResponse) SetApiName(v string) {
	o.ApiName = &v
}

// GetApikey returns the Apikey field value if set, zero value otherwise.
func (o *ModifySubAccountApiKeyPermissionResponse) GetApikey() string {
	if o == nil || common.IsNil(o.Apikey) {
		var ret string
		return ret
	}
	return *o.Apikey
}

// GetApikeyOk returns a tuple with the Apikey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifySubAccountApiKeyPermissionResponse) GetApikeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Apikey) {
		return nil, false
	}
	return o.Apikey, true
}

// HasApikey returns a boolean if a field has been set.
func (o *ModifySubAccountApiKeyPermissionResponse) HasApikey() bool {
	if o != nil && !common.IsNil(o.Apikey) {
		return true
	}

	return false
}

// SetApikey gets a reference to the given string and assigns it to the Apikey field.
func (o *ModifySubAccountApiKeyPermissionResponse) SetApikey(v string) {
	o.Apikey = &v
}

// GetCanTrade returns the CanTrade field value if set, zero value otherwise.
func (o *ModifySubAccountApiKeyPermissionResponse) GetCanTrade() bool {
	if o == nil || common.IsNil(o.CanTrade) {
		var ret bool
		return ret
	}
	return *o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifySubAccountApiKeyPermissionResponse) GetCanTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanTrade) {
		return nil, false
	}
	return o.CanTrade, true
}

// HasCanTrade returns a boolean if a field has been set.
func (o *ModifySubAccountApiKeyPermissionResponse) HasCanTrade() bool {
	if o != nil && !common.IsNil(o.CanTrade) {
		return true
	}

	return false
}

// SetCanTrade gets a reference to the given bool and assigns it to the CanTrade field.
func (o *ModifySubAccountApiKeyPermissionResponse) SetCanTrade(v bool) {
	o.CanTrade = &v
}

// GetCanMarginLoanRepay returns the CanMarginLoanRepay field value if set, zero value otherwise.
func (o *ModifySubAccountApiKeyPermissionResponse) GetCanMarginLoanRepay() bool {
	if o == nil || common.IsNil(o.CanMarginLoanRepay) {
		var ret bool
		return ret
	}
	return *o.CanMarginLoanRepay
}

// GetCanMarginLoanRepayOk returns a tuple with the CanMarginLoanRepay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifySubAccountApiKeyPermissionResponse) GetCanMarginLoanRepayOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanMarginLoanRepay) {
		return nil, false
	}
	return o.CanMarginLoanRepay, true
}

// HasCanMarginLoanRepay returns a boolean if a field has been set.
func (o *ModifySubAccountApiKeyPermissionResponse) HasCanMarginLoanRepay() bool {
	if o != nil && !common.IsNil(o.CanMarginLoanRepay) {
		return true
	}

	return false
}

// SetCanMarginLoanRepay gets a reference to the given bool and assigns it to the CanMarginLoanRepay field.
func (o *ModifySubAccountApiKeyPermissionResponse) SetCanMarginLoanRepay(v bool) {
	o.CanMarginLoanRepay = &v
}

// GetCanFuturesTrade returns the CanFuturesTrade field value if set, zero value otherwise.
func (o *ModifySubAccountApiKeyPermissionResponse) GetCanFuturesTrade() bool {
	if o == nil || common.IsNil(o.CanFuturesTrade) {
		var ret bool
		return ret
	}
	return *o.CanFuturesTrade
}

// GetCanFuturesTradeOk returns a tuple with the CanFuturesTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifySubAccountApiKeyPermissionResponse) GetCanFuturesTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanFuturesTrade) {
		return nil, false
	}
	return o.CanFuturesTrade, true
}

// HasCanFuturesTrade returns a boolean if a field has been set.
func (o *ModifySubAccountApiKeyPermissionResponse) HasCanFuturesTrade() bool {
	if o != nil && !common.IsNil(o.CanFuturesTrade) {
		return true
	}

	return false
}

// SetCanFuturesTrade gets a reference to the given bool and assigns it to the CanFuturesTrade field.
func (o *ModifySubAccountApiKeyPermissionResponse) SetCanFuturesTrade(v bool) {
	o.CanFuturesTrade = &v
}

// GetCanUniversalTransfer returns the CanUniversalTransfer field value if set, zero value otherwise.
func (o *ModifySubAccountApiKeyPermissionResponse) GetCanUniversalTransfer() bool {
	if o == nil || common.IsNil(o.CanUniversalTransfer) {
		var ret bool
		return ret
	}
	return *o.CanUniversalTransfer
}

// GetCanUniversalTransferOk returns a tuple with the CanUniversalTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifySubAccountApiKeyPermissionResponse) GetCanUniversalTransferOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanUniversalTransfer) {
		return nil, false
	}
	return o.CanUniversalTransfer, true
}

// HasCanUniversalTransfer returns a boolean if a field has been set.
func (o *ModifySubAccountApiKeyPermissionResponse) HasCanUniversalTransfer() bool {
	if o != nil && !common.IsNil(o.CanUniversalTransfer) {
		return true
	}

	return false
}

// SetCanUniversalTransfer gets a reference to the given bool and assigns it to the CanUniversalTransfer field.
func (o *ModifySubAccountApiKeyPermissionResponse) SetCanUniversalTransfer(v bool) {
	o.CanUniversalTransfer = &v
}

// GetCanVanillaOptions returns the CanVanillaOptions field value if set, zero value otherwise.
func (o *ModifySubAccountApiKeyPermissionResponse) GetCanVanillaOptions() bool {
	if o == nil || common.IsNil(o.CanVanillaOptions) {
		var ret bool
		return ret
	}
	return *o.CanVanillaOptions
}

// GetCanVanillaOptionsOk returns a tuple with the CanVanillaOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifySubAccountApiKeyPermissionResponse) GetCanVanillaOptionsOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanVanillaOptions) {
		return nil, false
	}
	return o.CanVanillaOptions, true
}

// HasCanVanillaOptions returns a boolean if a field has been set.
func (o *ModifySubAccountApiKeyPermissionResponse) HasCanVanillaOptions() bool {
	if o != nil && !common.IsNil(o.CanVanillaOptions) {
		return true
	}

	return false
}

// SetCanVanillaOptions gets a reference to the given bool and assigns it to the CanVanillaOptions field.
func (o *ModifySubAccountApiKeyPermissionResponse) SetCanVanillaOptions(v bool) {
	o.CanVanillaOptions = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ModifySubAccountApiKeyPermissionResponse) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifySubAccountApiKeyPermissionResponse) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ModifySubAccountApiKeyPermissionResponse) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ModifySubAccountApiKeyPermissionResponse) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o ModifySubAccountApiKeyPermissionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifySubAccountApiKeyPermissionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ApiName) {
		toSerialize["apiName"] = o.ApiName
	}
	if !common.IsNil(o.Apikey) {
		toSerialize["apikey"] = o.Apikey
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
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModifySubAccountApiKeyPermissionResponse) UnmarshalJSON(data []byte) (err error) {
	varModifySubAccountApiKeyPermissionResponse := _ModifySubAccountApiKeyPermissionResponse{}

	err = json.Unmarshal(data, &varModifySubAccountApiKeyPermissionResponse)

	if err != nil {
		return err
	}

	*o = ModifySubAccountApiKeyPermissionResponse(varModifySubAccountApiKeyPermissionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apiName")
		delete(additionalProperties, "apikey")
		delete(additionalProperties, "canTrade")
		delete(additionalProperties, "canMarginLoanRepay")
		delete(additionalProperties, "canFuturesTrade")
		delete(additionalProperties, "canUniversalTransfer")
		delete(additionalProperties, "canVanillaOptions")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModifySubAccountApiKeyPermissionResponse struct {
	value *ModifySubAccountApiKeyPermissionResponse
	isSet bool
}

func (v NullableModifySubAccountApiKeyPermissionResponse) Get() *ModifySubAccountApiKeyPermissionResponse {
	return v.value
}

func (v *NullableModifySubAccountApiKeyPermissionResponse) Set(val *ModifySubAccountApiKeyPermissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifySubAccountApiKeyPermissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifySubAccountApiKeyPermissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifySubAccountApiKeyPermissionResponse(val *ModifySubAccountApiKeyPermissionResponse) *NullableModifySubAccountApiKeyPermissionResponse {
	return &NullableModifySubAccountApiKeyPermissionResponse{value: val, isSet: true}
}

func (v NullableModifySubAccountApiKeyPermissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifySubAccountApiKeyPermissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
