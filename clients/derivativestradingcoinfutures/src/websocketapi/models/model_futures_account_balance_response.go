/*
Binance Derivatives Trading COIN Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FuturesAccountBalanceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FuturesAccountBalanceResponse{}

// FuturesAccountBalanceResponse struct for FuturesAccountBalanceResponse
type FuturesAccountBalanceResponse struct {
	Id                   *string                                     `json:"id,omitempty"`
	Status               *int64                                      `json:"status,omitempty"`
	Result               []FuturesAccountBalanceResponseResultInner  `json:"result,omitempty"`
	RateLimits           []AccountInformationResponseRateLimitsInner `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FuturesAccountBalanceResponse FuturesAccountBalanceResponse

// NewFuturesAccountBalanceResponse instantiates a new FuturesAccountBalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuturesAccountBalanceResponse() *FuturesAccountBalanceResponse {
	this := FuturesAccountBalanceResponse{}
	return &this
}

// NewFuturesAccountBalanceResponseWithDefaults instantiates a new FuturesAccountBalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFuturesAccountBalanceResponseWithDefaults() *FuturesAccountBalanceResponse {
	this := FuturesAccountBalanceResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FuturesAccountBalanceResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FuturesAccountBalanceResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FuturesAccountBalanceResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FuturesAccountBalanceResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FuturesAccountBalanceResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *FuturesAccountBalanceResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *FuturesAccountBalanceResponse) GetResult() []FuturesAccountBalanceResponseResultInner {
	if o == nil || common.IsNil(o.Result) {
		var ret []FuturesAccountBalanceResponseResultInner
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceResponse) GetResultOk() ([]FuturesAccountBalanceResponseResultInner, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *FuturesAccountBalanceResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []FuturesAccountBalanceResponseResultInner and assigns it to the Result field.
func (o *FuturesAccountBalanceResponse) SetResult(v []FuturesAccountBalanceResponseResultInner) {
	o.Result = v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *FuturesAccountBalanceResponse) GetRateLimits() []AccountInformationResponseRateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []AccountInformationResponseRateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceResponse) GetRateLimitsOk() ([]AccountInformationResponseRateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *FuturesAccountBalanceResponse) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []AccountInformationResponseRateLimitsInner and assigns it to the RateLimits field.
func (o *FuturesAccountBalanceResponse) SetRateLimits(v []AccountInformationResponseRateLimitsInner) {
	o.RateLimits = v
}

func (o FuturesAccountBalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FuturesAccountBalanceResponse) ToMap() (map[string]interface{}, error) {
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

func (o *FuturesAccountBalanceResponse) UnmarshalJSON(data []byte) (err error) {
	varFuturesAccountBalanceResponse := _FuturesAccountBalanceResponse{}

	err = json.Unmarshal(data, &varFuturesAccountBalanceResponse)

	if err != nil {
		return err
	}

	*o = FuturesAccountBalanceResponse(varFuturesAccountBalanceResponse)

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

type NullableFuturesAccountBalanceResponse struct {
	value *FuturesAccountBalanceResponse
	isSet bool
}

func (v NullableFuturesAccountBalanceResponse) Get() *FuturesAccountBalanceResponse {
	return v.value
}

func (v *NullableFuturesAccountBalanceResponse) Set(val *FuturesAccountBalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFuturesAccountBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFuturesAccountBalanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFuturesAccountBalanceResponse(val *FuturesAccountBalanceResponse) *NullableFuturesAccountBalanceResponse {
	return &NullableFuturesAccountBalanceResponse{value: val, isSet: true}
}

func (v NullableFuturesAccountBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuturesAccountBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
