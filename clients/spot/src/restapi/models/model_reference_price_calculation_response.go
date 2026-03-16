/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ReferencePriceCalculationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReferencePriceCalculationResponse{}

// ReferencePriceCalculationResponse struct for ReferencePriceCalculationResponse
type ReferencePriceCalculationResponse struct {
	Symbol                *string `json:"symbol,omitempty"`
	CalculationType       *string `json:"calculationType,omitempty"`
	BucketCount           *int64  `json:"bucketCount,omitempty"`
	BucketWidthMs         *int64  `json:"bucketWidthMs,omitempty"`
	ExternalCalculationId *int64  `json:"externalCalculationId,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _ReferencePriceCalculationResponse ReferencePriceCalculationResponse

// NewReferencePriceCalculationResponse instantiates a new ReferencePriceCalculationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferencePriceCalculationResponse() *ReferencePriceCalculationResponse {
	this := ReferencePriceCalculationResponse{}
	return &this
}

// NewReferencePriceCalculationResponseWithDefaults instantiates a new ReferencePriceCalculationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferencePriceCalculationResponseWithDefaults() *ReferencePriceCalculationResponse {
	this := ReferencePriceCalculationResponse{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ReferencePriceCalculationResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceCalculationResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ReferencePriceCalculationResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ReferencePriceCalculationResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetCalculationType returns the CalculationType field value if set, zero value otherwise.
func (o *ReferencePriceCalculationResponse) GetCalculationType() string {
	if o == nil || common.IsNil(o.CalculationType) {
		var ret string
		return ret
	}
	return *o.CalculationType
}

// GetCalculationTypeOk returns a tuple with the CalculationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceCalculationResponse) GetCalculationTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CalculationType) {
		return nil, false
	}
	return o.CalculationType, true
}

// HasCalculationType returns a boolean if a field has been set.
func (o *ReferencePriceCalculationResponse) HasCalculationType() bool {
	if o != nil && !common.IsNil(o.CalculationType) {
		return true
	}

	return false
}

// SetCalculationType gets a reference to the given string and assigns it to the CalculationType field.
func (o *ReferencePriceCalculationResponse) SetCalculationType(v string) {
	o.CalculationType = &v
}

// GetBucketCount returns the BucketCount field value if set, zero value otherwise.
func (o *ReferencePriceCalculationResponse) GetBucketCount() int64 {
	if o == nil || common.IsNil(o.BucketCount) {
		var ret int64
		return ret
	}
	return *o.BucketCount
}

// GetBucketCountOk returns a tuple with the BucketCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceCalculationResponse) GetBucketCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BucketCount) {
		return nil, false
	}
	return o.BucketCount, true
}

// HasBucketCount returns a boolean if a field has been set.
func (o *ReferencePriceCalculationResponse) HasBucketCount() bool {
	if o != nil && !common.IsNil(o.BucketCount) {
		return true
	}

	return false
}

// SetBucketCount gets a reference to the given int64 and assigns it to the BucketCount field.
func (o *ReferencePriceCalculationResponse) SetBucketCount(v int64) {
	o.BucketCount = &v
}

// GetBucketWidthMs returns the BucketWidthMs field value if set, zero value otherwise.
func (o *ReferencePriceCalculationResponse) GetBucketWidthMs() int64 {
	if o == nil || common.IsNil(o.BucketWidthMs) {
		var ret int64
		return ret
	}
	return *o.BucketWidthMs
}

// GetBucketWidthMsOk returns a tuple with the BucketWidthMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceCalculationResponse) GetBucketWidthMsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BucketWidthMs) {
		return nil, false
	}
	return o.BucketWidthMs, true
}

// HasBucketWidthMs returns a boolean if a field has been set.
func (o *ReferencePriceCalculationResponse) HasBucketWidthMs() bool {
	if o != nil && !common.IsNil(o.BucketWidthMs) {
		return true
	}

	return false
}

// SetBucketWidthMs gets a reference to the given int64 and assigns it to the BucketWidthMs field.
func (o *ReferencePriceCalculationResponse) SetBucketWidthMs(v int64) {
	o.BucketWidthMs = &v
}

// GetExternalCalculationId returns the ExternalCalculationId field value if set, zero value otherwise.
func (o *ReferencePriceCalculationResponse) GetExternalCalculationId() int64 {
	if o == nil || common.IsNil(o.ExternalCalculationId) {
		var ret int64
		return ret
	}
	return *o.ExternalCalculationId
}

// GetExternalCalculationIdOk returns a tuple with the ExternalCalculationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceCalculationResponse) GetExternalCalculationIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExternalCalculationId) {
		return nil, false
	}
	return o.ExternalCalculationId, true
}

// HasExternalCalculationId returns a boolean if a field has been set.
func (o *ReferencePriceCalculationResponse) HasExternalCalculationId() bool {
	if o != nil && !common.IsNil(o.ExternalCalculationId) {
		return true
	}

	return false
}

// SetExternalCalculationId gets a reference to the given int64 and assigns it to the ExternalCalculationId field.
func (o *ReferencePriceCalculationResponse) SetExternalCalculationId(v int64) {
	o.ExternalCalculationId = &v
}

func (o ReferencePriceCalculationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferencePriceCalculationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.CalculationType) {
		toSerialize["calculationType"] = o.CalculationType
	}
	if !common.IsNil(o.BucketCount) {
		toSerialize["bucketCount"] = o.BucketCount
	}
	if !common.IsNil(o.BucketWidthMs) {
		toSerialize["bucketWidthMs"] = o.BucketWidthMs
	}
	if !common.IsNil(o.ExternalCalculationId) {
		toSerialize["externalCalculationId"] = o.ExternalCalculationId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReferencePriceCalculationResponse) UnmarshalJSON(data []byte) (err error) {
	varReferencePriceCalculationResponse := _ReferencePriceCalculationResponse{}

	err = json.Unmarshal(data, &varReferencePriceCalculationResponse)

	if err != nil {
		return err
	}

	*o = ReferencePriceCalculationResponse(varReferencePriceCalculationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "calculationType")
		delete(additionalProperties, "bucketCount")
		delete(additionalProperties, "bucketWidthMs")
		delete(additionalProperties, "externalCalculationId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReferencePriceCalculationResponse struct {
	value *ReferencePriceCalculationResponse
	isSet bool
}

func (v NullableReferencePriceCalculationResponse) Get() *ReferencePriceCalculationResponse {
	return v.value
}

func (v *NullableReferencePriceCalculationResponse) Set(val *ReferencePriceCalculationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReferencePriceCalculationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReferencePriceCalculationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferencePriceCalculationResponse(val *ReferencePriceCalculationResponse) *NullableReferencePriceCalculationResponse {
	return &NullableReferencePriceCalculationResponse{value: val, isSet: true}
}

func (v NullableReferencePriceCalculationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferencePriceCalculationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
