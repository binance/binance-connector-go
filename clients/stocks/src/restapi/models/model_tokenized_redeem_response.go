/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TokenizedRedeemResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TokenizedRedeemResponse{}

// TokenizedRedeemResponse struct for TokenizedRedeemResponse
type TokenizedRedeemResponse struct {
	// Server-assigned convert request id. Use it to poll `/tokenized/convert-status`.
	IssuerRequestId *string `json:"issuerRequestId,omitempty"`
	// Convert status code: `P` = processing, `S` = success, `F` = failed.
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenizedRedeemResponse TokenizedRedeemResponse

// NewTokenizedRedeemResponse instantiates a new TokenizedRedeemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizedRedeemResponse() *TokenizedRedeemResponse {
	this := TokenizedRedeemResponse{}
	return &this
}

// NewTokenizedRedeemResponseWithDefaults instantiates a new TokenizedRedeemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizedRedeemResponseWithDefaults() *TokenizedRedeemResponse {
	this := TokenizedRedeemResponse{}
	return &this
}

// GetIssuerRequestId returns the IssuerRequestId field value if set, zero value otherwise.
func (o *TokenizedRedeemResponse) GetIssuerRequestId() string {
	if o == nil || common.IsNil(o.IssuerRequestId) {
		var ret string
		return ret
	}
	return *o.IssuerRequestId
}

// GetIssuerRequestIdOk returns a tuple with the IssuerRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedRedeemResponse) GetIssuerRequestIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuerRequestId) {
		return nil, false
	}
	return o.IssuerRequestId, true
}

// HasIssuerRequestId returns a boolean if a field has been set.
func (o *TokenizedRedeemResponse) HasIssuerRequestId() bool {
	if o != nil && !common.IsNil(o.IssuerRequestId) {
		return true
	}

	return false
}

// SetIssuerRequestId gets a reference to the given string and assigns it to the IssuerRequestId field.
func (o *TokenizedRedeemResponse) SetIssuerRequestId(v string) {
	o.IssuerRequestId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TokenizedRedeemResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedRedeemResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TokenizedRedeemResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TokenizedRedeemResponse) SetStatus(v string) {
	o.Status = &v
}

func (o TokenizedRedeemResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenizedRedeemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.IssuerRequestId) {
		toSerialize["issuerRequestId"] = o.IssuerRequestId
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenizedRedeemResponse) UnmarshalJSON(data []byte) (err error) {
	varTokenizedRedeemResponse := _TokenizedRedeemResponse{}

	err = json.Unmarshal(data, &varTokenizedRedeemResponse)

	if err != nil {
		return err
	}

	*o = TokenizedRedeemResponse(varTokenizedRedeemResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "issuerRequestId")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenizedRedeemResponse struct {
	value *TokenizedRedeemResponse
	isSet bool
}

func (v NullableTokenizedRedeemResponse) Get() *TokenizedRedeemResponse {
	return v.value
}

func (v *NullableTokenizedRedeemResponse) Set(val *TokenizedRedeemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizedRedeemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizedRedeemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizedRedeemResponse(val *TokenizedRedeemResponse) *NullableTokenizedRedeemResponse {
	return &NullableTokenizedRedeemResponse{value: val, isSet: true}
}

func (v NullableTokenizedRedeemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizedRedeemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
