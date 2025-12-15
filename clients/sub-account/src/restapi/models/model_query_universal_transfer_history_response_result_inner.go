/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryUniversalTransferHistoryResponseResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUniversalTransferHistoryResponseResultInner{}

// QueryUniversalTransferHistoryResponseResultInner struct for QueryUniversalTransferHistoryResponseResultInner
type QueryUniversalTransferHistoryResponseResultInner struct {
	TranId               *int64  `json:"tranId,omitempty"`
	FromEmail            *string `json:"fromEmail,omitempty"`
	ToEmail              *string `json:"toEmail,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	CreateTimeStamp      *int64  `json:"createTimeStamp,omitempty"`
	FromAccountType      *string `json:"fromAccountType,omitempty"`
	ToAccountType        *string `json:"toAccountType,omitempty"`
	Status               *string `json:"status,omitempty"`
	ClientTranId         *string `json:"clientTranId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryUniversalTransferHistoryResponseResultInner QueryUniversalTransferHistoryResponseResultInner

// NewQueryUniversalTransferHistoryResponseResultInner instantiates a new QueryUniversalTransferHistoryResponseResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUniversalTransferHistoryResponseResultInner() *QueryUniversalTransferHistoryResponseResultInner {
	this := QueryUniversalTransferHistoryResponseResultInner{}
	return &this
}

// NewQueryUniversalTransferHistoryResponseResultInnerWithDefaults instantiates a new QueryUniversalTransferHistoryResponseResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUniversalTransferHistoryResponseResultInnerWithDefaults() *QueryUniversalTransferHistoryResponseResultInner {
	this := QueryUniversalTransferHistoryResponseResultInner{}
	return &this
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *QueryUniversalTransferHistoryResponseResultInner) SetTranId(v int64) {
	o.TranId = &v
}

// GetFromEmail returns the FromEmail field value if set, zero value otherwise.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetFromEmail() string {
	if o == nil || common.IsNil(o.FromEmail) {
		var ret string
		return ret
	}
	return *o.FromEmail
}

// GetFromEmailOk returns a tuple with the FromEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetFromEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromEmail) {
		return nil, false
	}
	return o.FromEmail, true
}

// HasFromEmail returns a boolean if a field has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) HasFromEmail() bool {
	if o != nil && !common.IsNil(o.FromEmail) {
		return true
	}

	return false
}

// SetFromEmail gets a reference to the given string and assigns it to the FromEmail field.
func (o *QueryUniversalTransferHistoryResponseResultInner) SetFromEmail(v string) {
	o.FromEmail = &v
}

// GetToEmail returns the ToEmail field value if set, zero value otherwise.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetToEmail() string {
	if o == nil || common.IsNil(o.ToEmail) {
		var ret string
		return ret
	}
	return *o.ToEmail
}

// GetToEmailOk returns a tuple with the ToEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetToEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToEmail) {
		return nil, false
	}
	return o.ToEmail, true
}

// HasToEmail returns a boolean if a field has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) HasToEmail() bool {
	if o != nil && !common.IsNil(o.ToEmail) {
		return true
	}

	return false
}

// SetToEmail gets a reference to the given string and assigns it to the ToEmail field.
func (o *QueryUniversalTransferHistoryResponseResultInner) SetToEmail(v string) {
	o.ToEmail = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryUniversalTransferHistoryResponseResultInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QueryUniversalTransferHistoryResponseResultInner) SetAmount(v string) {
	o.Amount = &v
}

// GetCreateTimeStamp returns the CreateTimeStamp field value if set, zero value otherwise.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetCreateTimeStamp() int64 {
	if o == nil || common.IsNil(o.CreateTimeStamp) {
		var ret int64
		return ret
	}
	return *o.CreateTimeStamp
}

// GetCreateTimeStampOk returns a tuple with the CreateTimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetCreateTimeStampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTimeStamp) {
		return nil, false
	}
	return o.CreateTimeStamp, true
}

// HasCreateTimeStamp returns a boolean if a field has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) HasCreateTimeStamp() bool {
	if o != nil && !common.IsNil(o.CreateTimeStamp) {
		return true
	}

	return false
}

// SetCreateTimeStamp gets a reference to the given int64 and assigns it to the CreateTimeStamp field.
func (o *QueryUniversalTransferHistoryResponseResultInner) SetCreateTimeStamp(v int64) {
	o.CreateTimeStamp = &v
}

// GetFromAccountType returns the FromAccountType field value if set, zero value otherwise.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetFromAccountType() string {
	if o == nil || common.IsNil(o.FromAccountType) {
		var ret string
		return ret
	}
	return *o.FromAccountType
}

// GetFromAccountTypeOk returns a tuple with the FromAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetFromAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromAccountType) {
		return nil, false
	}
	return o.FromAccountType, true
}

// HasFromAccountType returns a boolean if a field has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) HasFromAccountType() bool {
	if o != nil && !common.IsNil(o.FromAccountType) {
		return true
	}

	return false
}

// SetFromAccountType gets a reference to the given string and assigns it to the FromAccountType field.
func (o *QueryUniversalTransferHistoryResponseResultInner) SetFromAccountType(v string) {
	o.FromAccountType = &v
}

// GetToAccountType returns the ToAccountType field value if set, zero value otherwise.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetToAccountType() string {
	if o == nil || common.IsNil(o.ToAccountType) {
		var ret string
		return ret
	}
	return *o.ToAccountType
}

// GetToAccountTypeOk returns a tuple with the ToAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetToAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToAccountType) {
		return nil, false
	}
	return o.ToAccountType, true
}

// HasToAccountType returns a boolean if a field has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) HasToAccountType() bool {
	if o != nil && !common.IsNil(o.ToAccountType) {
		return true
	}

	return false
}

// SetToAccountType gets a reference to the given string and assigns it to the ToAccountType field.
func (o *QueryUniversalTransferHistoryResponseResultInner) SetToAccountType(v string) {
	o.ToAccountType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryUniversalTransferHistoryResponseResultInner) SetStatus(v string) {
	o.Status = &v
}

// GetClientTranId returns the ClientTranId field value if set, zero value otherwise.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetClientTranId() string {
	if o == nil || common.IsNil(o.ClientTranId) {
		var ret string
		return ret
	}
	return *o.ClientTranId
}

// GetClientTranIdOk returns a tuple with the ClientTranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) GetClientTranIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientTranId) {
		return nil, false
	}
	return o.ClientTranId, true
}

// HasClientTranId returns a boolean if a field has been set.
func (o *QueryUniversalTransferHistoryResponseResultInner) HasClientTranId() bool {
	if o != nil && !common.IsNil(o.ClientTranId) {
		return true
	}

	return false
}

// SetClientTranId gets a reference to the given string and assigns it to the ClientTranId field.
func (o *QueryUniversalTransferHistoryResponseResultInner) SetClientTranId(v string) {
	o.ClientTranId = &v
}

func (o QueryUniversalTransferHistoryResponseResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUniversalTransferHistoryResponseResultInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	if !common.IsNil(o.FromEmail) {
		toSerialize["fromEmail"] = o.FromEmail
	}
	if !common.IsNil(o.ToEmail) {
		toSerialize["toEmail"] = o.ToEmail
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.CreateTimeStamp) {
		toSerialize["createTimeStamp"] = o.CreateTimeStamp
	}
	if !common.IsNil(o.FromAccountType) {
		toSerialize["fromAccountType"] = o.FromAccountType
	}
	if !common.IsNil(o.ToAccountType) {
		toSerialize["toAccountType"] = o.ToAccountType
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.ClientTranId) {
		toSerialize["clientTranId"] = o.ClientTranId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryUniversalTransferHistoryResponseResultInner) UnmarshalJSON(data []byte) (err error) {
	varQueryUniversalTransferHistoryResponseResultInner := _QueryUniversalTransferHistoryResponseResultInner{}

	err = json.Unmarshal(data, &varQueryUniversalTransferHistoryResponseResultInner)

	if err != nil {
		return err
	}

	*o = QueryUniversalTransferHistoryResponseResultInner(varQueryUniversalTransferHistoryResponseResultInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tranId")
		delete(additionalProperties, "fromEmail")
		delete(additionalProperties, "toEmail")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "createTimeStamp")
		delete(additionalProperties, "fromAccountType")
		delete(additionalProperties, "toAccountType")
		delete(additionalProperties, "status")
		delete(additionalProperties, "clientTranId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryUniversalTransferHistoryResponseResultInner struct {
	value *QueryUniversalTransferHistoryResponseResultInner
	isSet bool
}

func (v NullableQueryUniversalTransferHistoryResponseResultInner) Get() *QueryUniversalTransferHistoryResponseResultInner {
	return v.value
}

func (v *NullableQueryUniversalTransferHistoryResponseResultInner) Set(val *QueryUniversalTransferHistoryResponseResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUniversalTransferHistoryResponseResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUniversalTransferHistoryResponseResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUniversalTransferHistoryResponseResultInner(val *QueryUniversalTransferHistoryResponseResultInner) *NullableQueryUniversalTransferHistoryResponseResultInner {
	return &NullableQueryUniversalTransferHistoryResponseResultInner{value: val, isSet: true}
}

func (v NullableQueryUniversalTransferHistoryResponseResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUniversalTransferHistoryResponseResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
