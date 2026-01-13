/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PositionInformationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PositionInformationResponse{}

// PositionInformationResponse struct for PositionInformationResponse
type PositionInformationResponse struct {
	Id                   *string                                       `json:"id,omitempty"`
	Status               *int64                                        `json:"status,omitempty"`
	Result               []PositionInformationResponseResultInner      `json:"result,omitempty"`
	RateLimits           []AccountInformationV2ResponseRateLimitsInner `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PositionInformationResponse PositionInformationResponse

// NewPositionInformationResponse instantiates a new PositionInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPositionInformationResponse() *PositionInformationResponse {
	this := PositionInformationResponse{}
	return &this
}

// NewPositionInformationResponseWithDefaults instantiates a new PositionInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositionInformationResponseWithDefaults() *PositionInformationResponse {
	this := PositionInformationResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PositionInformationResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PositionInformationResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PositionInformationResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PositionInformationResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PositionInformationResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *PositionInformationResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *PositionInformationResponse) GetResult() []PositionInformationResponseResultInner {
	if o == nil || common.IsNil(o.Result) {
		var ret []PositionInformationResponseResultInner
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponse) GetResultOk() ([]PositionInformationResponseResultInner, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *PositionInformationResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []PositionInformationResponseResultInner and assigns it to the Result field.
func (o *PositionInformationResponse) SetResult(v []PositionInformationResponseResultInner) {
	o.Result = v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *PositionInformationResponse) GetRateLimits() []AccountInformationV2ResponseRateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []AccountInformationV2ResponseRateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponse) GetRateLimitsOk() ([]AccountInformationV2ResponseRateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *PositionInformationResponse) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []AccountInformationV2ResponseRateLimitsInner and assigns it to the RateLimits field.
func (o *PositionInformationResponse) SetRateLimits(v []AccountInformationV2ResponseRateLimitsInner) {
	o.RateLimits = v
}

func (o PositionInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PositionInformationResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.RateLimits) {
		toSerialize["rateLimits"] = o.RateLimits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PositionInformationResponse) UnmarshalJSON(data []byte) (err error) {
	varPositionInformationResponse := _PositionInformationResponse{}

	err = json.Unmarshal(data, &varPositionInformationResponse)

	if err != nil {
		return err
	}

	*o = PositionInformationResponse(varPositionInformationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "result")
		delete(additionalProperties, "rateLimits")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePositionInformationResponse struct {
	value *PositionInformationResponse
	isSet bool
}

func (v NullablePositionInformationResponse) Get() *PositionInformationResponse {
	return v.value
}

func (v *NullablePositionInformationResponse) Set(val *PositionInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePositionInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePositionInformationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositionInformationResponse(val *PositionInformationResponse) *NullablePositionInformationResponse {
	return &NullablePositionInformationResponse{value: val, isSet: true}
}

func (v NullablePositionInformationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositionInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
