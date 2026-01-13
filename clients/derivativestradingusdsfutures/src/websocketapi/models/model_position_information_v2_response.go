/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PositionInformationV2Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PositionInformationV2Response{}

// PositionInformationV2Response struct for PositionInformationV2Response
type PositionInformationV2Response struct {
	Id                   *string                                       `json:"id,omitempty"`
	Status               *int64                                        `json:"status,omitempty"`
	Result               []PositionInformationV2ResponseResultInner    `json:"result,omitempty"`
	RateLimits           []AccountInformationV2ResponseRateLimitsInner `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PositionInformationV2Response PositionInformationV2Response

// NewPositionInformationV2Response instantiates a new PositionInformationV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPositionInformationV2Response() *PositionInformationV2Response {
	this := PositionInformationV2Response{}
	return &this
}

// NewPositionInformationV2ResponseWithDefaults instantiates a new PositionInformationV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositionInformationV2ResponseWithDefaults() *PositionInformationV2Response {
	this := PositionInformationV2Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PositionInformationV2Response) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV2Response) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PositionInformationV2Response) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PositionInformationV2Response) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PositionInformationV2Response) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV2Response) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PositionInformationV2Response) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *PositionInformationV2Response) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *PositionInformationV2Response) GetResult() []PositionInformationV2ResponseResultInner {
	if o == nil || common.IsNil(o.Result) {
		var ret []PositionInformationV2ResponseResultInner
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV2Response) GetResultOk() ([]PositionInformationV2ResponseResultInner, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *PositionInformationV2Response) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []PositionInformationV2ResponseResultInner and assigns it to the Result field.
func (o *PositionInformationV2Response) SetResult(v []PositionInformationV2ResponseResultInner) {
	o.Result = v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *PositionInformationV2Response) GetRateLimits() []AccountInformationV2ResponseRateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []AccountInformationV2ResponseRateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationV2Response) GetRateLimitsOk() ([]AccountInformationV2ResponseRateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *PositionInformationV2Response) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []AccountInformationV2ResponseRateLimitsInner and assigns it to the RateLimits field.
func (o *PositionInformationV2Response) SetRateLimits(v []AccountInformationV2ResponseRateLimitsInner) {
	o.RateLimits = v
}

func (o PositionInformationV2Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PositionInformationV2Response) ToMap() (map[string]interface{}, error) {
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

func (o *PositionInformationV2Response) UnmarshalJSON(data []byte) (err error) {
	varPositionInformationV2Response := _PositionInformationV2Response{}

	err = json.Unmarshal(data, &varPositionInformationV2Response)

	if err != nil {
		return err
	}

	*o = PositionInformationV2Response(varPositionInformationV2Response)

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

type NullablePositionInformationV2Response struct {
	value *PositionInformationV2Response
	isSet bool
}

func (v NullablePositionInformationV2Response) Get() *PositionInformationV2Response {
	return v.value
}

func (v *NullablePositionInformationV2Response) Set(val *PositionInformationV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePositionInformationV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePositionInformationV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositionInformationV2Response(val *PositionInformationV2Response) *NullablePositionInformationV2Response {
	return &NullablePositionInformationV2Response{value: val, isSet: true}
}

func (v NullablePositionInformationV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositionInformationV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
