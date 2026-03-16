/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
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
	Id                   *string                                  `json:"id,omitempty"`
	Status               *int64                                   `json:"status,omitempty"`
	Result               *ReferencePriceCalculationResponseResult `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
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

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReferencePriceCalculationResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceCalculationResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReferencePriceCalculationResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReferencePriceCalculationResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReferencePriceCalculationResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceCalculationResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReferencePriceCalculationResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *ReferencePriceCalculationResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ReferencePriceCalculationResponse) GetResult() ReferencePriceCalculationResponseResult {
	if o == nil || common.IsNil(o.Result) {
		var ret ReferencePriceCalculationResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceCalculationResponse) GetResultOk() (*ReferencePriceCalculationResponseResult, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ReferencePriceCalculationResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ReferencePriceCalculationResponseResult and assigns it to the Result field.
func (o *ReferencePriceCalculationResponse) SetResult(v ReferencePriceCalculationResponseResult) {
	o.Result = &v
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
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Result) {
		toSerialize["result"] = o.Result
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
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "result")
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
