/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountInformationV2Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountInformationV2Response{}

// AccountInformationV2Response struct for AccountInformationV2Response
type AccountInformationV2Response struct {
	Id                   *string                                       `json:"id,omitempty"`
	Status               *int64                                        `json:"status,omitempty"`
	Result               *AccountInformationV2ResponseResult           `json:"result,omitempty"`
	RateLimits           []AccountInformationV2ResponseRateLimitsInner `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountInformationV2Response AccountInformationV2Response

// NewAccountInformationV2Response instantiates a new AccountInformationV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInformationV2Response() *AccountInformationV2Response {
	this := AccountInformationV2Response{}
	return &this
}

// NewAccountInformationV2ResponseWithDefaults instantiates a new AccountInformationV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInformationV2ResponseWithDefaults() *AccountInformationV2Response {
	this := AccountInformationV2Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountInformationV2Response) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *AccountInformationV2Response) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetResult() AccountInformationV2ResponseResult {
	if o == nil || common.IsNil(o.Result) {
		var ret AccountInformationV2ResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetResultOk() (*AccountInformationV2ResponseResult, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given AccountInformationV2ResponseResult and assigns it to the Result field.
func (o *AccountInformationV2Response) SetResult(v AccountInformationV2ResponseResult) {
	o.Result = &v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *AccountInformationV2Response) GetRateLimits() []AccountInformationV2ResponseRateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []AccountInformationV2ResponseRateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2Response) GetRateLimitsOk() ([]AccountInformationV2ResponseRateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *AccountInformationV2Response) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []AccountInformationV2ResponseRateLimitsInner and assigns it to the RateLimits field.
func (o *AccountInformationV2Response) SetRateLimits(v []AccountInformationV2ResponseRateLimitsInner) {
	o.RateLimits = v
}

func (o AccountInformationV2Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInformationV2Response) ToMap() (map[string]interface{}, error) {
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

func (o *AccountInformationV2Response) UnmarshalJSON(data []byte) (err error) {
	varAccountInformationV2Response := _AccountInformationV2Response{}

	err = json.Unmarshal(data, &varAccountInformationV2Response)

	if err != nil {
		return err
	}

	*o = AccountInformationV2Response(varAccountInformationV2Response)

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

type NullableAccountInformationV2Response struct {
	value *AccountInformationV2Response
	isSet bool
}

func (v NullableAccountInformationV2Response) Get() *AccountInformationV2Response {
	return v.value
}

func (v *NullableAccountInformationV2Response) Set(val *AccountInformationV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInformationV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInformationV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInformationV2Response(val *AccountInformationV2Response) *NullableAccountInformationV2Response {
	return &NullableAccountInformationV2Response{value: val, isSet: true}
}

func (v NullableAccountInformationV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInformationV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
