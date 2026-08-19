/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TokenizedConvertStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TokenizedConvertStatusResponse{}

// TokenizedConvertStatusResponse Empty object `{}` when no record matches.
type TokenizedConvertStatusResponse struct {
	// Underlying US-equity ticker, e.g. `AAPL`.
	UnderlyingAsset *string `json:"underlyingAsset,omitempty"`
	// Quantity of the underlying asset involved.
	UnderlyingAssetAmount *string `json:"underlyingAssetAmount,omitempty"`
	// Tokenized asset, e.g. `AAPLB`.
	TokenizedAsset *string `json:"tokenizedAsset,omitempty"`
	// Quantity of the tokenized asset involved.
	TokenizedAssetAmount *string `json:"tokenizedAssetAmount,omitempty"`
	// Echoes the requested id.
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

type _TokenizedConvertStatusResponse TokenizedConvertStatusResponse

// NewTokenizedConvertStatusResponse instantiates a new TokenizedConvertStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizedConvertStatusResponse() *TokenizedConvertStatusResponse {
	this := TokenizedConvertStatusResponse{}
	return &this
}

// NewTokenizedConvertStatusResponseWithDefaults instantiates a new TokenizedConvertStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizedConvertStatusResponseWithDefaults() *TokenizedConvertStatusResponse {
	this := TokenizedConvertStatusResponse{}
	return &this
}

// GetUnderlyingAsset returns the UnderlyingAsset field value if set, zero value otherwise.
func (o *TokenizedConvertStatusResponse) GetUnderlyingAsset() string {
	if o == nil || common.IsNil(o.UnderlyingAsset) {
		var ret string
		return ret
	}
	return *o.UnderlyingAsset
}

// GetUnderlyingAssetOk returns a tuple with the UnderlyingAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertStatusResponse) GetUnderlyingAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnderlyingAsset) {
		return nil, false
	}
	return o.UnderlyingAsset, true
}

// HasUnderlyingAsset returns a boolean if a field has been set.
func (o *TokenizedConvertStatusResponse) HasUnderlyingAsset() bool {
	if o != nil && !common.IsNil(o.UnderlyingAsset) {
		return true
	}

	return false
}

// SetUnderlyingAsset gets a reference to the given string and assigns it to the UnderlyingAsset field.
func (o *TokenizedConvertStatusResponse) SetUnderlyingAsset(v string) {
	o.UnderlyingAsset = &v
}

// GetUnderlyingAssetAmount returns the UnderlyingAssetAmount field value if set, zero value otherwise.
func (o *TokenizedConvertStatusResponse) GetUnderlyingAssetAmount() string {
	if o == nil || common.IsNil(o.UnderlyingAssetAmount) {
		var ret string
		return ret
	}
	return *o.UnderlyingAssetAmount
}

// GetUnderlyingAssetAmountOk returns a tuple with the UnderlyingAssetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertStatusResponse) GetUnderlyingAssetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnderlyingAssetAmount) {
		return nil, false
	}
	return o.UnderlyingAssetAmount, true
}

// HasUnderlyingAssetAmount returns a boolean if a field has been set.
func (o *TokenizedConvertStatusResponse) HasUnderlyingAssetAmount() bool {
	if o != nil && !common.IsNil(o.UnderlyingAssetAmount) {
		return true
	}

	return false
}

// SetUnderlyingAssetAmount gets a reference to the given string and assigns it to the UnderlyingAssetAmount field.
func (o *TokenizedConvertStatusResponse) SetUnderlyingAssetAmount(v string) {
	o.UnderlyingAssetAmount = &v
}

// GetTokenizedAsset returns the TokenizedAsset field value if set, zero value otherwise.
func (o *TokenizedConvertStatusResponse) GetTokenizedAsset() string {
	if o == nil || common.IsNil(o.TokenizedAsset) {
		var ret string
		return ret
	}
	return *o.TokenizedAsset
}

// GetTokenizedAssetOk returns a tuple with the TokenizedAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertStatusResponse) GetTokenizedAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenizedAsset) {
		return nil, false
	}
	return o.TokenizedAsset, true
}

// HasTokenizedAsset returns a boolean if a field has been set.
func (o *TokenizedConvertStatusResponse) HasTokenizedAsset() bool {
	if o != nil && !common.IsNil(o.TokenizedAsset) {
		return true
	}

	return false
}

// SetTokenizedAsset gets a reference to the given string and assigns it to the TokenizedAsset field.
func (o *TokenizedConvertStatusResponse) SetTokenizedAsset(v string) {
	o.TokenizedAsset = &v
}

// GetTokenizedAssetAmount returns the TokenizedAssetAmount field value if set, zero value otherwise.
func (o *TokenizedConvertStatusResponse) GetTokenizedAssetAmount() string {
	if o == nil || common.IsNil(o.TokenizedAssetAmount) {
		var ret string
		return ret
	}
	return *o.TokenizedAssetAmount
}

// GetTokenizedAssetAmountOk returns a tuple with the TokenizedAssetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertStatusResponse) GetTokenizedAssetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenizedAssetAmount) {
		return nil, false
	}
	return o.TokenizedAssetAmount, true
}

// HasTokenizedAssetAmount returns a boolean if a field has been set.
func (o *TokenizedConvertStatusResponse) HasTokenizedAssetAmount() bool {
	if o != nil && !common.IsNil(o.TokenizedAssetAmount) {
		return true
	}

	return false
}

// SetTokenizedAssetAmount gets a reference to the given string and assigns it to the TokenizedAssetAmount field.
func (o *TokenizedConvertStatusResponse) SetTokenizedAssetAmount(v string) {
	o.TokenizedAssetAmount = &v
}

// GetIssuerRequestId returns the IssuerRequestId field value if set, zero value otherwise.
func (o *TokenizedConvertStatusResponse) GetIssuerRequestId() string {
	if o == nil || common.IsNil(o.IssuerRequestId) {
		var ret string
		return ret
	}
	return *o.IssuerRequestId
}

// GetIssuerRequestIdOk returns a tuple with the IssuerRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertStatusResponse) GetIssuerRequestIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuerRequestId) {
		return nil, false
	}
	return o.IssuerRequestId, true
}

// HasIssuerRequestId returns a boolean if a field has been set.
func (o *TokenizedConvertStatusResponse) HasIssuerRequestId() bool {
	if o != nil && !common.IsNil(o.IssuerRequestId) {
		return true
	}

	return false
}

// SetIssuerRequestId gets a reference to the given string and assigns it to the IssuerRequestId field.
func (o *TokenizedConvertStatusResponse) SetIssuerRequestId(v string) {
	o.IssuerRequestId = &v
}

// GetConvertType returns the ConvertType field value if set, zero value otherwise.
func (o *TokenizedConvertStatusResponse) GetConvertType() string {
	if o == nil || common.IsNil(o.ConvertType) {
		var ret string
		return ret
	}
	return *o.ConvertType
}

// GetConvertTypeOk returns a tuple with the ConvertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertStatusResponse) GetConvertTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ConvertType) {
		return nil, false
	}
	return o.ConvertType, true
}

// HasConvertType returns a boolean if a field has been set.
func (o *TokenizedConvertStatusResponse) HasConvertType() bool {
	if o != nil && !common.IsNil(o.ConvertType) {
		return true
	}

	return false
}

// SetConvertType gets a reference to the given string and assigns it to the ConvertType field.
func (o *TokenizedConvertStatusResponse) SetConvertType(v string) {
	o.ConvertType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TokenizedConvertStatusResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertStatusResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TokenizedConvertStatusResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TokenizedConvertStatusResponse) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TokenizedConvertStatusResponse) GetCreatedAt() int64 {
	if o == nil || common.IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertStatusResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TokenizedConvertStatusResponse) HasCreatedAt() bool {
	if o != nil && !common.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *TokenizedConvertStatusResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TokenizedConvertStatusResponse) GetUpdatedAt() int64 {
	if o == nil || common.IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedConvertStatusResponse) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TokenizedConvertStatusResponse) HasUpdatedAt() bool {
	if o != nil && !common.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *TokenizedConvertStatusResponse) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

func (o TokenizedConvertStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenizedConvertStatusResponse) ToMap() (map[string]interface{}, error) {
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

func (o *TokenizedConvertStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varTokenizedConvertStatusResponse := _TokenizedConvertStatusResponse{}

	err = json.Unmarshal(data, &varTokenizedConvertStatusResponse)

	if err != nil {
		return err
	}

	*o = TokenizedConvertStatusResponse(varTokenizedConvertStatusResponse)

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

type NullableTokenizedConvertStatusResponse struct {
	value *TokenizedConvertStatusResponse
	isSet bool
}

func (v NullableTokenizedConvertStatusResponse) Get() *TokenizedConvertStatusResponse {
	return v.value
}

func (v *NullableTokenizedConvertStatusResponse) Set(val *TokenizedConvertStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizedConvertStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizedConvertStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizedConvertStatusResponse(val *TokenizedConvertStatusResponse) *NullableTokenizedConvertStatusResponse {
	return &NullableTokenizedConvertStatusResponse{value: val, isSet: true}
}

func (v NullableTokenizedConvertStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizedConvertStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
