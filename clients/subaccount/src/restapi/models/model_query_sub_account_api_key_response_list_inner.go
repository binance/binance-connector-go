/*
Sub Account REST API

Create and manage sub-accounts, control permissions, and transfer assets via the Sub Account API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QuerySubAccountApiKeyResponseListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubAccountApiKeyResponseListInner{}

// QuerySubAccountApiKeyResponseListInner struct for QuerySubAccountApiKeyResponseListInner
type QuerySubAccountApiKeyResponseListInner struct {
	Email                *string `json:"email,omitempty"`
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

type _QuerySubAccountApiKeyResponseListInner QuerySubAccountApiKeyResponseListInner

// NewQuerySubAccountApiKeyResponseListInner instantiates a new QuerySubAccountApiKeyResponseListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubAccountApiKeyResponseListInner() *QuerySubAccountApiKeyResponseListInner {
	this := QuerySubAccountApiKeyResponseListInner{}
	return &this
}

// NewQuerySubAccountApiKeyResponseListInnerWithDefaults instantiates a new QuerySubAccountApiKeyResponseListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubAccountApiKeyResponseListInnerWithDefaults() *QuerySubAccountApiKeyResponseListInner {
	this := QuerySubAccountApiKeyResponseListInner{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *QuerySubAccountApiKeyResponseListInner) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountApiKeyResponseListInner) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *QuerySubAccountApiKeyResponseListInner) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *QuerySubAccountApiKeyResponseListInner) SetEmail(v string) {
	o.Email = &v
}

// GetApiName returns the ApiName field value if set, zero value otherwise.
func (o *QuerySubAccountApiKeyResponseListInner) GetApiName() string {
	if o == nil || common.IsNil(o.ApiName) {
		var ret string
		return ret
	}
	return *o.ApiName
}

// GetApiNameOk returns a tuple with the ApiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountApiKeyResponseListInner) GetApiNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.ApiName) {
		return nil, false
	}
	return o.ApiName, true
}

// HasApiName returns a boolean if a field has been set.
func (o *QuerySubAccountApiKeyResponseListInner) HasApiName() bool {
	if o != nil && !common.IsNil(o.ApiName) {
		return true
	}

	return false
}

// SetApiName gets a reference to the given string and assigns it to the ApiName field.
func (o *QuerySubAccountApiKeyResponseListInner) SetApiName(v string) {
	o.ApiName = &v
}

// GetApikey returns the Apikey field value if set, zero value otherwise.
func (o *QuerySubAccountApiKeyResponseListInner) GetApikey() string {
	if o == nil || common.IsNil(o.Apikey) {
		var ret string
		return ret
	}
	return *o.Apikey
}

// GetApikeyOk returns a tuple with the Apikey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountApiKeyResponseListInner) GetApikeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Apikey) {
		return nil, false
	}
	return o.Apikey, true
}

// HasApikey returns a boolean if a field has been set.
func (o *QuerySubAccountApiKeyResponseListInner) HasApikey() bool {
	if o != nil && !common.IsNil(o.Apikey) {
		return true
	}

	return false
}

// SetApikey gets a reference to the given string and assigns it to the Apikey field.
func (o *QuerySubAccountApiKeyResponseListInner) SetApikey(v string) {
	o.Apikey = &v
}

// GetCanTrade returns the CanTrade field value if set, zero value otherwise.
func (o *QuerySubAccountApiKeyResponseListInner) GetCanTrade() bool {
	if o == nil || common.IsNil(o.CanTrade) {
		var ret bool
		return ret
	}
	return *o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountApiKeyResponseListInner) GetCanTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanTrade) {
		return nil, false
	}
	return o.CanTrade, true
}

// HasCanTrade returns a boolean if a field has been set.
func (o *QuerySubAccountApiKeyResponseListInner) HasCanTrade() bool {
	if o != nil && !common.IsNil(o.CanTrade) {
		return true
	}

	return false
}

// SetCanTrade gets a reference to the given bool and assigns it to the CanTrade field.
func (o *QuerySubAccountApiKeyResponseListInner) SetCanTrade(v bool) {
	o.CanTrade = &v
}

// GetCanMarginLoanRepay returns the CanMarginLoanRepay field value if set, zero value otherwise.
func (o *QuerySubAccountApiKeyResponseListInner) GetCanMarginLoanRepay() bool {
	if o == nil || common.IsNil(o.CanMarginLoanRepay) {
		var ret bool
		return ret
	}
	return *o.CanMarginLoanRepay
}

// GetCanMarginLoanRepayOk returns a tuple with the CanMarginLoanRepay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountApiKeyResponseListInner) GetCanMarginLoanRepayOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanMarginLoanRepay) {
		return nil, false
	}
	return o.CanMarginLoanRepay, true
}

// HasCanMarginLoanRepay returns a boolean if a field has been set.
func (o *QuerySubAccountApiKeyResponseListInner) HasCanMarginLoanRepay() bool {
	if o != nil && !common.IsNil(o.CanMarginLoanRepay) {
		return true
	}

	return false
}

// SetCanMarginLoanRepay gets a reference to the given bool and assigns it to the CanMarginLoanRepay field.
func (o *QuerySubAccountApiKeyResponseListInner) SetCanMarginLoanRepay(v bool) {
	o.CanMarginLoanRepay = &v
}

// GetCanFuturesTrade returns the CanFuturesTrade field value if set, zero value otherwise.
func (o *QuerySubAccountApiKeyResponseListInner) GetCanFuturesTrade() bool {
	if o == nil || common.IsNil(o.CanFuturesTrade) {
		var ret bool
		return ret
	}
	return *o.CanFuturesTrade
}

// GetCanFuturesTradeOk returns a tuple with the CanFuturesTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountApiKeyResponseListInner) GetCanFuturesTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanFuturesTrade) {
		return nil, false
	}
	return o.CanFuturesTrade, true
}

// HasCanFuturesTrade returns a boolean if a field has been set.
func (o *QuerySubAccountApiKeyResponseListInner) HasCanFuturesTrade() bool {
	if o != nil && !common.IsNil(o.CanFuturesTrade) {
		return true
	}

	return false
}

// SetCanFuturesTrade gets a reference to the given bool and assigns it to the CanFuturesTrade field.
func (o *QuerySubAccountApiKeyResponseListInner) SetCanFuturesTrade(v bool) {
	o.CanFuturesTrade = &v
}

// GetCanUniversalTransfer returns the CanUniversalTransfer field value if set, zero value otherwise.
func (o *QuerySubAccountApiKeyResponseListInner) GetCanUniversalTransfer() bool {
	if o == nil || common.IsNil(o.CanUniversalTransfer) {
		var ret bool
		return ret
	}
	return *o.CanUniversalTransfer
}

// GetCanUniversalTransferOk returns a tuple with the CanUniversalTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountApiKeyResponseListInner) GetCanUniversalTransferOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanUniversalTransfer) {
		return nil, false
	}
	return o.CanUniversalTransfer, true
}

// HasCanUniversalTransfer returns a boolean if a field has been set.
func (o *QuerySubAccountApiKeyResponseListInner) HasCanUniversalTransfer() bool {
	if o != nil && !common.IsNil(o.CanUniversalTransfer) {
		return true
	}

	return false
}

// SetCanUniversalTransfer gets a reference to the given bool and assigns it to the CanUniversalTransfer field.
func (o *QuerySubAccountApiKeyResponseListInner) SetCanUniversalTransfer(v bool) {
	o.CanUniversalTransfer = &v
}

// GetCanVanillaOptions returns the CanVanillaOptions field value if set, zero value otherwise.
func (o *QuerySubAccountApiKeyResponseListInner) GetCanVanillaOptions() bool {
	if o == nil || common.IsNil(o.CanVanillaOptions) {
		var ret bool
		return ret
	}
	return *o.CanVanillaOptions
}

// GetCanVanillaOptionsOk returns a tuple with the CanVanillaOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountApiKeyResponseListInner) GetCanVanillaOptionsOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanVanillaOptions) {
		return nil, false
	}
	return o.CanVanillaOptions, true
}

// HasCanVanillaOptions returns a boolean if a field has been set.
func (o *QuerySubAccountApiKeyResponseListInner) HasCanVanillaOptions() bool {
	if o != nil && !common.IsNil(o.CanVanillaOptions) {
		return true
	}

	return false
}

// SetCanVanillaOptions gets a reference to the given bool and assigns it to the CanVanillaOptions field.
func (o *QuerySubAccountApiKeyResponseListInner) SetCanVanillaOptions(v bool) {
	o.CanVanillaOptions = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *QuerySubAccountApiKeyResponseListInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountApiKeyResponseListInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *QuerySubAccountApiKeyResponseListInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *QuerySubAccountApiKeyResponseListInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o QuerySubAccountApiKeyResponseListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubAccountApiKeyResponseListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
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

func (o *QuerySubAccountApiKeyResponseListInner) UnmarshalJSON(data []byte) (err error) {
	varQuerySubAccountApiKeyResponseListInner := _QuerySubAccountApiKeyResponseListInner{}

	err = json.Unmarshal(data, &varQuerySubAccountApiKeyResponseListInner)

	if err != nil {
		return err
	}

	*o = QuerySubAccountApiKeyResponseListInner(varQuerySubAccountApiKeyResponseListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
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

type NullableQuerySubAccountApiKeyResponseListInner struct {
	value *QuerySubAccountApiKeyResponseListInner
	isSet bool
}

func (v NullableQuerySubAccountApiKeyResponseListInner) Get() *QuerySubAccountApiKeyResponseListInner {
	return v.value
}

func (v *NullableQuerySubAccountApiKeyResponseListInner) Set(val *QuerySubAccountApiKeyResponseListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubAccountApiKeyResponseListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubAccountApiKeyResponseListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubAccountApiKeyResponseListInner(val *QuerySubAccountApiKeyResponseListInner) *NullableQuerySubAccountApiKeyResponseListInner {
	return &NullableQuerySubAccountApiKeyResponseListInner{value: val, isSet: true}
}

func (v NullableQuerySubAccountApiKeyResponseListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubAccountApiKeyResponseListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
