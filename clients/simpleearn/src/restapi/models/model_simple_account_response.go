/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SimpleAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SimpleAccountResponse{}

// SimpleAccountResponse struct for SimpleAccountResponse
type SimpleAccountResponse struct {
	TotalAmountInBTC          *string `json:"totalAmountInBTC,omitempty"`
	TotalAmountInUSDT         *string `json:"totalAmountInUSDT,omitempty"`
	TotalFlexibleAmountInBTC  *string `json:"totalFlexibleAmountInBTC,omitempty"`
	TotalFlexibleAmountInUSDT *string `json:"totalFlexibleAmountInUSDT,omitempty"`
	TotalLockedInBTC          *string `json:"totalLockedInBTC,omitempty"`
	TotalLockedInUSDT         *string `json:"totalLockedInUSDT,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _SimpleAccountResponse SimpleAccountResponse

// NewSimpleAccountResponse instantiates a new SimpleAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleAccountResponse() *SimpleAccountResponse {
	this := SimpleAccountResponse{}
	return &this
}

// NewSimpleAccountResponseWithDefaults instantiates a new SimpleAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleAccountResponseWithDefaults() *SimpleAccountResponse {
	this := SimpleAccountResponse{}
	return &this
}

// GetTotalAmountInBTC returns the TotalAmountInBTC field value if set, zero value otherwise.
func (o *SimpleAccountResponse) GetTotalAmountInBTC() string {
	if o == nil || common.IsNil(o.TotalAmountInBTC) {
		var ret string
		return ret
	}
	return *o.TotalAmountInBTC
}

// GetTotalAmountInBTCOk returns a tuple with the TotalAmountInBTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleAccountResponse) GetTotalAmountInBTCOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAmountInBTC) {
		return nil, false
	}
	return o.TotalAmountInBTC, true
}

// HasTotalAmountInBTC returns a boolean if a field has been set.
func (o *SimpleAccountResponse) HasTotalAmountInBTC() bool {
	if o != nil && !common.IsNil(o.TotalAmountInBTC) {
		return true
	}

	return false
}

// SetTotalAmountInBTC gets a reference to the given string and assigns it to the TotalAmountInBTC field.
func (o *SimpleAccountResponse) SetTotalAmountInBTC(v string) {
	o.TotalAmountInBTC = &v
}

// GetTotalAmountInUSDT returns the TotalAmountInUSDT field value if set, zero value otherwise.
func (o *SimpleAccountResponse) GetTotalAmountInUSDT() string {
	if o == nil || common.IsNil(o.TotalAmountInUSDT) {
		var ret string
		return ret
	}
	return *o.TotalAmountInUSDT
}

// GetTotalAmountInUSDTOk returns a tuple with the TotalAmountInUSDT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleAccountResponse) GetTotalAmountInUSDTOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAmountInUSDT) {
		return nil, false
	}
	return o.TotalAmountInUSDT, true
}

// HasTotalAmountInUSDT returns a boolean if a field has been set.
func (o *SimpleAccountResponse) HasTotalAmountInUSDT() bool {
	if o != nil && !common.IsNil(o.TotalAmountInUSDT) {
		return true
	}

	return false
}

// SetTotalAmountInUSDT gets a reference to the given string and assigns it to the TotalAmountInUSDT field.
func (o *SimpleAccountResponse) SetTotalAmountInUSDT(v string) {
	o.TotalAmountInUSDT = &v
}

// GetTotalFlexibleAmountInBTC returns the TotalFlexibleAmountInBTC field value if set, zero value otherwise.
func (o *SimpleAccountResponse) GetTotalFlexibleAmountInBTC() string {
	if o == nil || common.IsNil(o.TotalFlexibleAmountInBTC) {
		var ret string
		return ret
	}
	return *o.TotalFlexibleAmountInBTC
}

// GetTotalFlexibleAmountInBTCOk returns a tuple with the TotalFlexibleAmountInBTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleAccountResponse) GetTotalFlexibleAmountInBTCOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalFlexibleAmountInBTC) {
		return nil, false
	}
	return o.TotalFlexibleAmountInBTC, true
}

// HasTotalFlexibleAmountInBTC returns a boolean if a field has been set.
func (o *SimpleAccountResponse) HasTotalFlexibleAmountInBTC() bool {
	if o != nil && !common.IsNil(o.TotalFlexibleAmountInBTC) {
		return true
	}

	return false
}

// SetTotalFlexibleAmountInBTC gets a reference to the given string and assigns it to the TotalFlexibleAmountInBTC field.
func (o *SimpleAccountResponse) SetTotalFlexibleAmountInBTC(v string) {
	o.TotalFlexibleAmountInBTC = &v
}

// GetTotalFlexibleAmountInUSDT returns the TotalFlexibleAmountInUSDT field value if set, zero value otherwise.
func (o *SimpleAccountResponse) GetTotalFlexibleAmountInUSDT() string {
	if o == nil || common.IsNil(o.TotalFlexibleAmountInUSDT) {
		var ret string
		return ret
	}
	return *o.TotalFlexibleAmountInUSDT
}

// GetTotalFlexibleAmountInUSDTOk returns a tuple with the TotalFlexibleAmountInUSDT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleAccountResponse) GetTotalFlexibleAmountInUSDTOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalFlexibleAmountInUSDT) {
		return nil, false
	}
	return o.TotalFlexibleAmountInUSDT, true
}

// HasTotalFlexibleAmountInUSDT returns a boolean if a field has been set.
func (o *SimpleAccountResponse) HasTotalFlexibleAmountInUSDT() bool {
	if o != nil && !common.IsNil(o.TotalFlexibleAmountInUSDT) {
		return true
	}

	return false
}

// SetTotalFlexibleAmountInUSDT gets a reference to the given string and assigns it to the TotalFlexibleAmountInUSDT field.
func (o *SimpleAccountResponse) SetTotalFlexibleAmountInUSDT(v string) {
	o.TotalFlexibleAmountInUSDT = &v
}

// GetTotalLockedInBTC returns the TotalLockedInBTC field value if set, zero value otherwise.
func (o *SimpleAccountResponse) GetTotalLockedInBTC() string {
	if o == nil || common.IsNil(o.TotalLockedInBTC) {
		var ret string
		return ret
	}
	return *o.TotalLockedInBTC
}

// GetTotalLockedInBTCOk returns a tuple with the TotalLockedInBTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleAccountResponse) GetTotalLockedInBTCOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalLockedInBTC) {
		return nil, false
	}
	return o.TotalLockedInBTC, true
}

// HasTotalLockedInBTC returns a boolean if a field has been set.
func (o *SimpleAccountResponse) HasTotalLockedInBTC() bool {
	if o != nil && !common.IsNil(o.TotalLockedInBTC) {
		return true
	}

	return false
}

// SetTotalLockedInBTC gets a reference to the given string and assigns it to the TotalLockedInBTC field.
func (o *SimpleAccountResponse) SetTotalLockedInBTC(v string) {
	o.TotalLockedInBTC = &v
}

// GetTotalLockedInUSDT returns the TotalLockedInUSDT field value if set, zero value otherwise.
func (o *SimpleAccountResponse) GetTotalLockedInUSDT() string {
	if o == nil || common.IsNil(o.TotalLockedInUSDT) {
		var ret string
		return ret
	}
	return *o.TotalLockedInUSDT
}

// GetTotalLockedInUSDTOk returns a tuple with the TotalLockedInUSDT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleAccountResponse) GetTotalLockedInUSDTOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalLockedInUSDT) {
		return nil, false
	}
	return o.TotalLockedInUSDT, true
}

// HasTotalLockedInUSDT returns a boolean if a field has been set.
func (o *SimpleAccountResponse) HasTotalLockedInUSDT() bool {
	if o != nil && !common.IsNil(o.TotalLockedInUSDT) {
		return true
	}

	return false
}

// SetTotalLockedInUSDT gets a reference to the given string and assigns it to the TotalLockedInUSDT field.
func (o *SimpleAccountResponse) SetTotalLockedInUSDT(v string) {
	o.TotalLockedInUSDT = &v
}

func (o SimpleAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimpleAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TotalAmountInBTC) {
		toSerialize["totalAmountInBTC"] = o.TotalAmountInBTC
	}
	if !common.IsNil(o.TotalAmountInUSDT) {
		toSerialize["totalAmountInUSDT"] = o.TotalAmountInUSDT
	}
	if !common.IsNil(o.TotalFlexibleAmountInBTC) {
		toSerialize["totalFlexibleAmountInBTC"] = o.TotalFlexibleAmountInBTC
	}
	if !common.IsNil(o.TotalFlexibleAmountInUSDT) {
		toSerialize["totalFlexibleAmountInUSDT"] = o.TotalFlexibleAmountInUSDT
	}
	if !common.IsNil(o.TotalLockedInBTC) {
		toSerialize["totalLockedInBTC"] = o.TotalLockedInBTC
	}
	if !common.IsNil(o.TotalLockedInUSDT) {
		toSerialize["totalLockedInUSDT"] = o.TotalLockedInUSDT
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SimpleAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varSimpleAccountResponse := _SimpleAccountResponse{}

	err = json.Unmarshal(data, &varSimpleAccountResponse)

	if err != nil {
		return err
	}

	*o = SimpleAccountResponse(varSimpleAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "totalAmountInBTC")
		delete(additionalProperties, "totalAmountInUSDT")
		delete(additionalProperties, "totalFlexibleAmountInBTC")
		delete(additionalProperties, "totalFlexibleAmountInUSDT")
		delete(additionalProperties, "totalLockedInBTC")
		delete(additionalProperties, "totalLockedInUSDT")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimpleAccountResponse struct {
	value *SimpleAccountResponse
	isSet bool
}

func (v NullableSimpleAccountResponse) Get() *SimpleAccountResponse {
	return v.value
}

func (v *NullableSimpleAccountResponse) Set(val *SimpleAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleAccountResponse(val *SimpleAccountResponse) *NullableSimpleAccountResponse {
	return &NullableSimpleAccountResponse{value: val, isSet: true}
}

func (v NullableSimpleAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
