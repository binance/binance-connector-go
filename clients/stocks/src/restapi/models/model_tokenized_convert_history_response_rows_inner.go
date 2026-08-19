/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TokenizedConvertHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TokenizedConvertHistoryResponseRowsInner{}

// TokenizedConvertHistoryResponseRowsInner struct for TokenizedConvertHistoryResponseRowsInner
type TokenizedConvertHistoryResponseRowsInner struct {
	// Underlying US-equity ticker, e.g. `AAPL`.
	UnderlyingAsset *string `json:"underlyingAsset,omitempty"`
	// Quantity of the underlying asset involved.
	UnderlyingAssetAmount *string `json:"underlyingAssetAmount,omitempty"`
	// Tokenized asset, e.g. `AAPLB`.
	TokenizedAsset *string `json:"tokenizedAsset,omitempty"`
	// Quantity of the tokenized asset involved.
	TokenizedAssetAmount *string `json:"tokenizedAssetAmount,omitempty"`
	// Convert request id.
	IssuerRequestId *string `json:"issuerRequestId,omitempty"`
	// `MINT` or `REDEEM`.
	ConvertType *string `json:"convertType,omitempty"`
	// Convert status: `P` = processing, `S` = success, `F` = failed.
	Status *string `json:"status,omitempty"`
	// Creation time (ms epoch).
	CreatedAt *int64 `json:"createdAt,omitempty"`
	// Last update time (ms epoch).
	UpdatedAt            *int64 `json:"updatedAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenizedConvertHistoryResponseRowsInner TokenizedConvertHistoryResponseRowsInner

// NewTokenizedConvertHistoryResponseRowsInner instantiates a new TokenizedConvertHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizedConvertHistoryResponseRowsInner() *TokenizedConvertHistoryResponseRowsInner {
	this := TokenizedConvertHistoryResponseRowsInner{}
	return &this
}

// NewTokenizedConvertHistoryResponseRowsInnerWithDefaults instantiates a new TokenizedConvertHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizedConvertHistoryResponseRowsInnerWithDefaults() *TokenizedConvertHistoryResponseRowsInner {
	this := TokenizedConvertHistoryResponseRowsInner{}
	return &this
}

// GetUnderlyingAsset returns the UnderlyingAsset field value if set, zero value otherwise.
func (o *TokenizedConvertHistoryResponseRowsInner) GetUnderlyingAsset() string {
	if o == nil || common.IsNil(o.UnderlyingAsset) {
		var ret string
		return ret
	}
	return *o.UnderlyingAsset
}

// GetUnderlyingAssetOk returns a tuple with the UnderlyingAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) GetUnderlyingAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnderlyingAsset) {
		return nil, false
	}
	return o.UnderlyingAsset, true
}

// HasUnderlyingAsset returns a boolean if a field has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) HasUnderlyingAsset() bool {
	if o != nil && !common.IsNil(o.UnderlyingAsset) {
		return true
	}

	return false
}

// SetUnderlyingAsset gets a reference to the given string and assigns it to the UnderlyingAsset field.
func (o *TokenizedConvertHistoryResponseRowsInner) SetUnderlyingAsset(v string) {
	o.UnderlyingAsset = &v
}

// GetUnderlyingAssetAmount returns the UnderlyingAssetAmount field value if set, zero value otherwise.
func (o *TokenizedConvertHistoryResponseRowsInner) GetUnderlyingAssetAmount() string {
	if o == nil || common.IsNil(o.UnderlyingAssetAmount) {
		var ret string
		return ret
	}
	return *o.UnderlyingAssetAmount
}

// GetUnderlyingAssetAmountOk returns a tuple with the UnderlyingAssetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) GetUnderlyingAssetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnderlyingAssetAmount) {
		return nil, false
	}
	return o.UnderlyingAssetAmount, true
}

// HasUnderlyingAssetAmount returns a boolean if a field has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) HasUnderlyingAssetAmount() bool {
	if o != nil && !common.IsNil(o.UnderlyingAssetAmount) {
		return true
	}

	return false
}

// SetUnderlyingAssetAmount gets a reference to the given string and assigns it to the UnderlyingAssetAmount field.
func (o *TokenizedConvertHistoryResponseRowsInner) SetUnderlyingAssetAmount(v string) {
	o.UnderlyingAssetAmount = &v
}

// GetTokenizedAsset returns the TokenizedAsset field value if set, zero value otherwise.
func (o *TokenizedConvertHistoryResponseRowsInner) GetTokenizedAsset() string {
	if o == nil || common.IsNil(o.TokenizedAsset) {
		var ret string
		return ret
	}
	return *o.TokenizedAsset
}

// GetTokenizedAssetOk returns a tuple with the TokenizedAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) GetTokenizedAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenizedAsset) {
		return nil, false
	}
	return o.TokenizedAsset, true
}

// HasTokenizedAsset returns a boolean if a field has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) HasTokenizedAsset() bool {
	if o != nil && !common.IsNil(o.TokenizedAsset) {
		return true
	}

	return false
}

// SetTokenizedAsset gets a reference to the given string and assigns it to the TokenizedAsset field.
func (o *TokenizedConvertHistoryResponseRowsInner) SetTokenizedAsset(v string) {
	o.TokenizedAsset = &v
}

// GetTokenizedAssetAmount returns the TokenizedAssetAmount field value if set, zero value otherwise.
func (o *TokenizedConvertHistoryResponseRowsInner) GetTokenizedAssetAmount() string {
	if o == nil || common.IsNil(o.TokenizedAssetAmount) {
		var ret string
		return ret
	}
	return *o.TokenizedAssetAmount
}

// GetTokenizedAssetAmountOk returns a tuple with the TokenizedAssetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) GetTokenizedAssetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenizedAssetAmount) {
		return nil, false
	}
	return o.TokenizedAssetAmount, true
}

// HasTokenizedAssetAmount returns a boolean if a field has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) HasTokenizedAssetAmount() bool {
	if o != nil && !common.IsNil(o.TokenizedAssetAmount) {
		return true
	}

	return false
}

// SetTokenizedAssetAmount gets a reference to the given string and assigns it to the TokenizedAssetAmount field.
func (o *TokenizedConvertHistoryResponseRowsInner) SetTokenizedAssetAmount(v string) {
	o.TokenizedAssetAmount = &v
}

// GetIssuerRequestId returns the IssuerRequestId field value if set, zero value otherwise.
func (o *TokenizedConvertHistoryResponseRowsInner) GetIssuerRequestId() string {
	if o == nil || common.IsNil(o.IssuerRequestId) {
		var ret string
		return ret
	}
	return *o.IssuerRequestId
}

// GetIssuerRequestIdOk returns a tuple with the IssuerRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) GetIssuerRequestIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuerRequestId) {
		return nil, false
	}
	return o.IssuerRequestId, true
}

// HasIssuerRequestId returns a boolean if a field has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) HasIssuerRequestId() bool {
	if o != nil && !common.IsNil(o.IssuerRequestId) {
		return true
	}

	return false
}

// SetIssuerRequestId gets a reference to the given string and assigns it to the IssuerRequestId field.
func (o *TokenizedConvertHistoryResponseRowsInner) SetIssuerRequestId(v string) {
	o.IssuerRequestId = &v
}

// GetConvertType returns the ConvertType field value if set, zero value otherwise.
func (o *TokenizedConvertHistoryResponseRowsInner) GetConvertType() string {
	if o == nil || common.IsNil(o.ConvertType) {
		var ret string
		return ret
	}
	return *o.ConvertType
}

// GetConvertTypeOk returns a tuple with the ConvertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) GetConvertTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ConvertType) {
		return nil, false
	}
	return o.ConvertType, true
}

// HasConvertType returns a boolean if a field has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) HasConvertType() bool {
	if o != nil && !common.IsNil(o.ConvertType) {
		return true
	}

	return false
}

// SetConvertType gets a reference to the given string and assigns it to the ConvertType field.
func (o *TokenizedConvertHistoryResponseRowsInner) SetConvertType(v string) {
	o.ConvertType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TokenizedConvertHistoryResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TokenizedConvertHistoryResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TokenizedConvertHistoryResponseRowsInner) GetCreatedAt() int64 {
	if o == nil || common.IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) GetCreatedAtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) HasCreatedAt() bool {
	if o != nil && !common.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *TokenizedConvertHistoryResponseRowsInner) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TokenizedConvertHistoryResponseRowsInner) GetUpdatedAt() int64 {
	if o == nil || common.IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TokenizedConvertHistoryResponseRowsInner) HasUpdatedAt() bool {
	if o != nil && !common.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *TokenizedConvertHistoryResponseRowsInner) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

func (o TokenizedConvertHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenizedConvertHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.UnderlyingAsset) {
		toSerialize["underlyingAsset"] = o.UnderlyingAsset
	}
	if !common.IsNil(o.UnderlyingAssetAmount) {
		toSerialize["underlyingAssetAmount"] = o.UnderlyingAssetAmount
	}
	if !common.IsNil(o.TokenizedAsset) {
		toSerialize["tokenizedAsset"] = o.TokenizedAsset
	}
	if !common.IsNil(o.TokenizedAssetAmount) {
		toSerialize["tokenizedAssetAmount"] = o.TokenizedAssetAmount
	}
	if !common.IsNil(o.IssuerRequestId) {
		toSerialize["issuerRequestId"] = o.IssuerRequestId
	}
	if !common.IsNil(o.ConvertType) {
		toSerialize["convertType"] = o.ConvertType
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !common.IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenizedConvertHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varTokenizedConvertHistoryResponseRowsInner := _TokenizedConvertHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varTokenizedConvertHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = TokenizedConvertHistoryResponseRowsInner(varTokenizedConvertHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "underlyingAsset")
		delete(additionalProperties, "underlyingAssetAmount")
		delete(additionalProperties, "tokenizedAsset")
		delete(additionalProperties, "tokenizedAssetAmount")
		delete(additionalProperties, "issuerRequestId")
		delete(additionalProperties, "convertType")
		delete(additionalProperties, "status")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenizedConvertHistoryResponseRowsInner struct {
	value *TokenizedConvertHistoryResponseRowsInner
	isSet bool
}

func (v NullableTokenizedConvertHistoryResponseRowsInner) Get() *TokenizedConvertHistoryResponseRowsInner {
	return v.value
}

func (v *NullableTokenizedConvertHistoryResponseRowsInner) Set(val *TokenizedConvertHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizedConvertHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizedConvertHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizedConvertHistoryResponseRowsInner(val *TokenizedConvertHistoryResponseRowsInner) *NullableTokenizedConvertHistoryResponseRowsInner {
	return &NullableTokenizedConvertHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableTokenizedConvertHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizedConvertHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
