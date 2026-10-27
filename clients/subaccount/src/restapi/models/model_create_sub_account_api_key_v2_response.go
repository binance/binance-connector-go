/*
Sub Account REST API

Create and manage sub-accounts, control permissions, and transfer assets via the Sub Account API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CreateSubAccountApiKeyV2Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateSubAccountApiKeyV2Response{}

// CreateSubAccountApiKeyV2Response struct for CreateSubAccountApiKeyV2Response
type CreateSubAccountApiKeyV2Response struct {
	ApiName *string `json:"apiName,omitempty"`
	ApiKey  *string `json:"apiKey,omitempty"`
	// Secret Key. Returned only once on creation, please keep it safe.
	SecretKey            *string  `json:"secretKey,omitempty"`
	CanTrade             *bool    `json:"canTrade,omitempty"`
	CanMarginLoanRepay   *bool    `json:"canMarginLoanRepay,omitempty"`
	CanFuturesTrade      *bool    `json:"canFuturesTrade,omitempty"`
	CanUniversalTransfer *bool    `json:"canUniversalTransfer,omitempty"`
	CanVanillaOptions    *bool    `json:"canVanillaOptions,omitempty"`
	CanFixApiTrading     *bool    `json:"canFixApiTrading,omitempty"`
	CanFixApiReading     *bool    `json:"canFixApiReading,omitempty"`
	Status               *int64   `json:"status,omitempty"`
	IpList               []string `json:"ipList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateSubAccountApiKeyV2Response CreateSubAccountApiKeyV2Response

// NewCreateSubAccountApiKeyV2Response instantiates a new CreateSubAccountApiKeyV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubAccountApiKeyV2Response() *CreateSubAccountApiKeyV2Response {
	this := CreateSubAccountApiKeyV2Response{}
	return &this
}

// NewCreateSubAccountApiKeyV2ResponseWithDefaults instantiates a new CreateSubAccountApiKeyV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubAccountApiKeyV2ResponseWithDefaults() *CreateSubAccountApiKeyV2Response {
	this := CreateSubAccountApiKeyV2Response{}
	return &this
}

// GetApiName returns the ApiName field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyV2Response) GetApiName() string {
	if o == nil || common.IsNil(o.ApiName) {
		var ret string
		return ret
	}
	return *o.ApiName
}

// GetApiNameOk returns a tuple with the ApiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyV2Response) GetApiNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.ApiName) {
		return nil, false
	}
	return o.ApiName, true
}

// HasApiName returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyV2Response) HasApiName() bool {
	if o != nil && !common.IsNil(o.ApiName) {
		return true
	}

	return false
}

// SetApiName gets a reference to the given string and assigns it to the ApiName field.
func (o *CreateSubAccountApiKeyV2Response) SetApiName(v string) {
	o.ApiName = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyV2Response) GetApiKey() string {
	if o == nil || common.IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyV2Response) GetApiKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyV2Response) HasApiKey() bool {
	if o != nil && !common.IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *CreateSubAccountApiKeyV2Response) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyV2Response) GetSecretKey() string {
	if o == nil || common.IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyV2Response) GetSecretKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyV2Response) HasSecretKey() bool {
	if o != nil && !common.IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *CreateSubAccountApiKeyV2Response) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetCanTrade returns the CanTrade field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyV2Response) GetCanTrade() bool {
	if o == nil || common.IsNil(o.CanTrade) {
		var ret bool
		return ret
	}
	return *o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyV2Response) GetCanTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanTrade) {
		return nil, false
	}
	return o.CanTrade, true
}

// HasCanTrade returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyV2Response) HasCanTrade() bool {
	if o != nil && !common.IsNil(o.CanTrade) {
		return true
	}

	return false
}

// SetCanTrade gets a reference to the given bool and assigns it to the CanTrade field.
func (o *CreateSubAccountApiKeyV2Response) SetCanTrade(v bool) {
	o.CanTrade = &v
}

// GetCanMarginLoanRepay returns the CanMarginLoanRepay field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyV2Response) GetCanMarginLoanRepay() bool {
	if o == nil || common.IsNil(o.CanMarginLoanRepay) {
		var ret bool
		return ret
	}
	return *o.CanMarginLoanRepay
}

// GetCanMarginLoanRepayOk returns a tuple with the CanMarginLoanRepay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyV2Response) GetCanMarginLoanRepayOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanMarginLoanRepay) {
		return nil, false
	}
	return o.CanMarginLoanRepay, true
}

// HasCanMarginLoanRepay returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyV2Response) HasCanMarginLoanRepay() bool {
	if o != nil && !common.IsNil(o.CanMarginLoanRepay) {
		return true
	}

	return false
}

// SetCanMarginLoanRepay gets a reference to the given bool and assigns it to the CanMarginLoanRepay field.
func (o *CreateSubAccountApiKeyV2Response) SetCanMarginLoanRepay(v bool) {
	o.CanMarginLoanRepay = &v
}

// GetCanFuturesTrade returns the CanFuturesTrade field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyV2Response) GetCanFuturesTrade() bool {
	if o == nil || common.IsNil(o.CanFuturesTrade) {
		var ret bool
		return ret
	}
	return *o.CanFuturesTrade
}

// GetCanFuturesTradeOk returns a tuple with the CanFuturesTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyV2Response) GetCanFuturesTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanFuturesTrade) {
		return nil, false
	}
	return o.CanFuturesTrade, true
}

// HasCanFuturesTrade returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyV2Response) HasCanFuturesTrade() bool {
	if o != nil && !common.IsNil(o.CanFuturesTrade) {
		return true
	}

	return false
}

// SetCanFuturesTrade gets a reference to the given bool and assigns it to the CanFuturesTrade field.
func (o *CreateSubAccountApiKeyV2Response) SetCanFuturesTrade(v bool) {
	o.CanFuturesTrade = &v
}

// GetCanUniversalTransfer returns the CanUniversalTransfer field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyV2Response) GetCanUniversalTransfer() bool {
	if o == nil || common.IsNil(o.CanUniversalTransfer) {
		var ret bool
		return ret
	}
	return *o.CanUniversalTransfer
}

// GetCanUniversalTransferOk returns a tuple with the CanUniversalTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyV2Response) GetCanUniversalTransferOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanUniversalTransfer) {
		return nil, false
	}
	return o.CanUniversalTransfer, true
}

// HasCanUniversalTransfer returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyV2Response) HasCanUniversalTransfer() bool {
	if o != nil && !common.IsNil(o.CanUniversalTransfer) {
		return true
	}

	return false
}

// SetCanUniversalTransfer gets a reference to the given bool and assigns it to the CanUniversalTransfer field.
func (o *CreateSubAccountApiKeyV2Response) SetCanUniversalTransfer(v bool) {
	o.CanUniversalTransfer = &v
}

// GetCanVanillaOptions returns the CanVanillaOptions field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyV2Response) GetCanVanillaOptions() bool {
	if o == nil || common.IsNil(o.CanVanillaOptions) {
		var ret bool
		return ret
	}
	return *o.CanVanillaOptions
}

// GetCanVanillaOptionsOk returns a tuple with the CanVanillaOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyV2Response) GetCanVanillaOptionsOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanVanillaOptions) {
		return nil, false
	}
	return o.CanVanillaOptions, true
}

// HasCanVanillaOptions returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyV2Response) HasCanVanillaOptions() bool {
	if o != nil && !common.IsNil(o.CanVanillaOptions) {
		return true
	}

	return false
}

// SetCanVanillaOptions gets a reference to the given bool and assigns it to the CanVanillaOptions field.
func (o *CreateSubAccountApiKeyV2Response) SetCanVanillaOptions(v bool) {
	o.CanVanillaOptions = &v
}

// GetCanFixApiTrading returns the CanFixApiTrading field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyV2Response) GetCanFixApiTrading() bool {
	if o == nil || common.IsNil(o.CanFixApiTrading) {
		var ret bool
		return ret
	}
	return *o.CanFixApiTrading
}

// GetCanFixApiTradingOk returns a tuple with the CanFixApiTrading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyV2Response) GetCanFixApiTradingOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanFixApiTrading) {
		return nil, false
	}
	return o.CanFixApiTrading, true
}

// HasCanFixApiTrading returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyV2Response) HasCanFixApiTrading() bool {
	if o != nil && !common.IsNil(o.CanFixApiTrading) {
		return true
	}

	return false
}

// SetCanFixApiTrading gets a reference to the given bool and assigns it to the CanFixApiTrading field.
func (o *CreateSubAccountApiKeyV2Response) SetCanFixApiTrading(v bool) {
	o.CanFixApiTrading = &v
}

// GetCanFixApiReading returns the CanFixApiReading field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyV2Response) GetCanFixApiReading() bool {
	if o == nil || common.IsNil(o.CanFixApiReading) {
		var ret bool
		return ret
	}
	return *o.CanFixApiReading
}

// GetCanFixApiReadingOk returns a tuple with the CanFixApiReading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyV2Response) GetCanFixApiReadingOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanFixApiReading) {
		return nil, false
	}
	return o.CanFixApiReading, true
}

// HasCanFixApiReading returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyV2Response) HasCanFixApiReading() bool {
	if o != nil && !common.IsNil(o.CanFixApiReading) {
		return true
	}

	return false
}

// SetCanFixApiReading gets a reference to the given bool and assigns it to the CanFixApiReading field.
func (o *CreateSubAccountApiKeyV2Response) SetCanFixApiReading(v bool) {
	o.CanFixApiReading = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyV2Response) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyV2Response) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyV2Response) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *CreateSubAccountApiKeyV2Response) SetStatus(v int64) {
	o.Status = &v
}

// GetIpList returns the IpList field value if set, zero value otherwise.
func (o *CreateSubAccountApiKeyV2Response) GetIpList() []string {
	if o == nil || common.IsNil(o.IpList) {
		var ret []string
		return ret
	}
	return o.IpList
}

// GetIpListOk returns a tuple with the IpList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountApiKeyV2Response) GetIpListOk() ([]string, bool) {
	if o == nil || common.IsNil(o.IpList) {
		return nil, false
	}
	return o.IpList, true
}

// HasIpList returns a boolean if a field has been set.
func (o *CreateSubAccountApiKeyV2Response) HasIpList() bool {
	if o != nil && !common.IsNil(o.IpList) {
		return true
	}

	return false
}

// SetIpList gets a reference to the given []string and assigns it to the IpList field.
func (o *CreateSubAccountApiKeyV2Response) SetIpList(v []string) {
	o.IpList = v
}

func (o CreateSubAccountApiKeyV2Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubAccountApiKeyV2Response) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.CanFixApiTrading) {
		toSerialize["canFixApiTrading"] = o.CanFixApiTrading
	}
	if !common.IsNil(o.CanFixApiReading) {
		toSerialize["canFixApiReading"] = o.CanFixApiReading
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

func (o *CreateSubAccountApiKeyV2Response) UnmarshalJSON(data []byte) (err error) {
	varCreateSubAccountApiKeyV2Response := _CreateSubAccountApiKeyV2Response{}

	err = json.Unmarshal(data, &varCreateSubAccountApiKeyV2Response)

	if err != nil {
		return err
	}

	*o = CreateSubAccountApiKeyV2Response(varCreateSubAccountApiKeyV2Response)

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
		delete(additionalProperties, "canFixApiTrading")
		delete(additionalProperties, "canFixApiReading")
		delete(additionalProperties, "status")
		delete(additionalProperties, "ipList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSubAccountApiKeyV2Response struct {
	value *CreateSubAccountApiKeyV2Response
	isSet bool
}

func (v NullableCreateSubAccountApiKeyV2Response) Get() *CreateSubAccountApiKeyV2Response {
	return v.value
}

func (v *NullableCreateSubAccountApiKeyV2Response) Set(val *CreateSubAccountApiKeyV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubAccountApiKeyV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubAccountApiKeyV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubAccountApiKeyV2Response(val *CreateSubAccountApiKeyV2Response) *NullableCreateSubAccountApiKeyV2Response {
	return &NullableCreateSubAccountApiKeyV2Response{value: val, isSet: true}
}

func (v NullableCreateSubAccountApiKeyV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSubAccountApiKeyV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
