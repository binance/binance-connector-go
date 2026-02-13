/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CloseUserDataStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CloseUserDataStreamResponse{}

// CloseUserDataStreamResponse struct for CloseUserDataStreamResponse
type CloseUserDataStreamResponse struct {
	Id                   *string                                         `json:"id,omitempty"`
	Status               *int64                                          `json:"status,omitempty"`
	Result               []interface{}                                   `json:"result,omitempty"`
	RateLimits           []SymbolOrderBookTickerResponse1RateLimitsInner `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloseUserDataStreamResponse CloseUserDataStreamResponse

// NewCloseUserDataStreamResponse instantiates a new CloseUserDataStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloseUserDataStreamResponse() *CloseUserDataStreamResponse {
	this := CloseUserDataStreamResponse{}
	return &this
}

// NewCloseUserDataStreamResponseWithDefaults instantiates a new CloseUserDataStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloseUserDataStreamResponseWithDefaults() *CloseUserDataStreamResponse {
	this := CloseUserDataStreamResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CloseUserDataStreamResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloseUserDataStreamResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloseUserDataStreamResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloseUserDataStreamResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CloseUserDataStreamResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloseUserDataStreamResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CloseUserDataStreamResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *CloseUserDataStreamResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CloseUserDataStreamResponse) GetResult() interface{} {
	return nil
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloseUserDataStreamResponse) GetResultOk() (interface{}, bool) {
	return nil, true
}

// HasResult returns a boolean if a field has been set.
func (o *CloseUserDataStreamResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given map[string]interface{} and assigns it to the Result field.
func (o *CloseUserDataStreamResponse) SetResult(v []interface{}) {
	o.Result = v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *CloseUserDataStreamResponse) GetRateLimits() []SymbolOrderBookTickerResponse1RateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []SymbolOrderBookTickerResponse1RateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloseUserDataStreamResponse) GetRateLimitsOk() ([]SymbolOrderBookTickerResponse1RateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *CloseUserDataStreamResponse) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []SymbolOrderBookTickerResponse1RateLimitsInner and assigns it to the RateLimits field.
func (o *CloseUserDataStreamResponse) SetRateLimits(v []SymbolOrderBookTickerResponse1RateLimitsInner) {
	o.RateLimits = v
}

func (o CloseUserDataStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloseUserDataStreamResponse) ToMap() (map[string]interface{}, error) {
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

func (o *CloseUserDataStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varCloseUserDataStreamResponse := _CloseUserDataStreamResponse{}

	err = json.Unmarshal(data, &varCloseUserDataStreamResponse)

	if err != nil {
		return err
	}

	*o = CloseUserDataStreamResponse(varCloseUserDataStreamResponse)

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

type NullableCloseUserDataStreamResponse struct {
	value *CloseUserDataStreamResponse
	isSet bool
}

func (v NullableCloseUserDataStreamResponse) Get() *CloseUserDataStreamResponse {
	return v.value
}

func (v *NullableCloseUserDataStreamResponse) Set(val *CloseUserDataStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloseUserDataStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloseUserDataStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloseUserDataStreamResponse(val *CloseUserDataStreamResponse) *NullableCloseUserDataStreamResponse {
	return &NullableCloseUserDataStreamResponse{value: val, isSet: true}
}

func (v NullableCloseUserDataStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloseUserDataStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
