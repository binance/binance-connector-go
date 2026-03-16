/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ReferencePriceCalculationResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReferencePriceCalculationResponseResult{}

// ReferencePriceCalculationResponseResult struct for ReferencePriceCalculationResponseResult
type ReferencePriceCalculationResponseResult struct {
	Symbol                *string `json:"symbol,omitempty"`
	CalculationType       *string `json:"calculationType,omitempty"`
	BucketCount           *int64  `json:"bucketCount,omitempty"`
	BucketWidthMs         *int64  `json:"bucketWidthMs,omitempty"`
	ExternalCalculationId *int64  `json:"externalCalculationId,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _ReferencePriceCalculationResponseResult ReferencePriceCalculationResponseResult

// NewReferencePriceCalculationResponseResult instantiates a new ReferencePriceCalculationResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferencePriceCalculationResponseResult() *ReferencePriceCalculationResponseResult {
	this := ReferencePriceCalculationResponseResult{}
	return &this
}

// NewReferencePriceCalculationResponseResultWithDefaults instantiates a new ReferencePriceCalculationResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferencePriceCalculationResponseResultWithDefaults() *ReferencePriceCalculationResponseResult {
	this := ReferencePriceCalculationResponseResult{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ReferencePriceCalculationResponseResult) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceCalculationResponseResult) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ReferencePriceCalculationResponseResult) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ReferencePriceCalculationResponseResult) SetSymbol(v string) {
	o.Symbol = &v
}

// GetCalculationType returns the CalculationType field value if set, zero value otherwise.
func (o *ReferencePriceCalculationResponseResult) GetCalculationType() string {
	if o == nil || common.IsNil(o.CalculationType) {
		var ret string
		return ret
	}
	return *o.CalculationType
}

// GetCalculationTypeOk returns a tuple with the CalculationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceCalculationResponseResult) GetCalculationTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CalculationType) {
		return nil, false
	}
	return o.CalculationType, true
}

// HasCalculationType returns a boolean if a field has been set.
func (o *ReferencePriceCalculationResponseResult) HasCalculationType() bool {
	if o != nil && !common.IsNil(o.CalculationType) {
		return true
	}

	return false
}

// SetCalculationType gets a reference to the given string and assigns it to the CalculationType field.
func (o *ReferencePriceCalculationResponseResult) SetCalculationType(v string) {
	o.CalculationType = &v
}

// GetBucketCount returns the BucketCount field value if set, zero value otherwise.
func (o *ReferencePriceCalculationResponseResult) GetBucketCount() int64 {
	if o == nil || common.IsNil(o.BucketCount) {
		var ret int64
		return ret
	}
	return *o.BucketCount
}

// GetBucketCountOk returns a tuple with the BucketCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceCalculationResponseResult) GetBucketCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BucketCount) {
		return nil, false
	}
	return o.BucketCount, true
}

// HasBucketCount returns a boolean if a field has been set.
func (o *ReferencePriceCalculationResponseResult) HasBucketCount() bool {
	if o != nil && !common.IsNil(o.BucketCount) {
		return true
	}

	return false
}

// SetBucketCount gets a reference to the given int64 and assigns it to the BucketCount field.
func (o *ReferencePriceCalculationResponseResult) SetBucketCount(v int64) {
	o.BucketCount = &v
}

// GetBucketWidthMs returns the BucketWidthMs field value if set, zero value otherwise.
func (o *ReferencePriceCalculationResponseResult) GetBucketWidthMs() int64 {
	if o == nil || common.IsNil(o.BucketWidthMs) {
		var ret int64
		return ret
	}
	return *o.BucketWidthMs
}

// GetBucketWidthMsOk returns a tuple with the BucketWidthMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceCalculationResponseResult) GetBucketWidthMsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BucketWidthMs) {
		return nil, false
	}
	return o.BucketWidthMs, true
}

// HasBucketWidthMs returns a boolean if a field has been set.
func (o *ReferencePriceCalculationResponseResult) HasBucketWidthMs() bool {
	if o != nil && !common.IsNil(o.BucketWidthMs) {
		return true
	}

	return false
}

// SetBucketWidthMs gets a reference to the given int64 and assigns it to the BucketWidthMs field.
func (o *ReferencePriceCalculationResponseResult) SetBucketWidthMs(v int64) {
	o.BucketWidthMs = &v
}

// GetExternalCalculationId returns the ExternalCalculationId field value if set, zero value otherwise.
func (o *ReferencePriceCalculationResponseResult) GetExternalCalculationId() int64 {
	if o == nil || common.IsNil(o.ExternalCalculationId) {
		var ret int64
		return ret
	}
	return *o.ExternalCalculationId
}

// GetExternalCalculationIdOk returns a tuple with the ExternalCalculationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceCalculationResponseResult) GetExternalCalculationIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExternalCalculationId) {
		return nil, false
	}
	return o.ExternalCalculationId, true
}

// HasExternalCalculationId returns a boolean if a field has been set.
func (o *ReferencePriceCalculationResponseResult) HasExternalCalculationId() bool {
	if o != nil && !common.IsNil(o.ExternalCalculationId) {
		return true
	}

	return false
}

// SetExternalCalculationId gets a reference to the given int64 and assigns it to the ExternalCalculationId field.
func (o *ReferencePriceCalculationResponseResult) SetExternalCalculationId(v int64) {
	o.ExternalCalculationId = &v
}

func (o ReferencePriceCalculationResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferencePriceCalculationResponseResult) ToMap() (map[string]interface{}, error) {
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

func (o *ReferencePriceCalculationResponseResult) UnmarshalJSON(data []byte) (err error) {
	varReferencePriceCalculationResponseResult := _ReferencePriceCalculationResponseResult{}

	err = json.Unmarshal(data, &varReferencePriceCalculationResponseResult)

	if err != nil {
		return err
	}

	*o = ReferencePriceCalculationResponseResult(varReferencePriceCalculationResponseResult)

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

type NullableReferencePriceCalculationResponseResult struct {
	value *ReferencePriceCalculationResponseResult
	isSet bool
}

func (v NullableReferencePriceCalculationResponseResult) Get() *ReferencePriceCalculationResponseResult {
	return v.value
}

func (v *NullableReferencePriceCalculationResponseResult) Set(val *ReferencePriceCalculationResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableReferencePriceCalculationResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableReferencePriceCalculationResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferencePriceCalculationResponseResult(val *ReferencePriceCalculationResponseResult) *NullableReferencePriceCalculationResponseResult {
	return &NullableReferencePriceCalculationResponseResult{value: val, isSet: true}
}

func (v NullableReferencePriceCalculationResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferencePriceCalculationResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
