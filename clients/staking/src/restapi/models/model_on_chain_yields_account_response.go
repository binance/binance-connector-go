/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OnChainYieldsAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OnChainYieldsAccountResponse{}

// OnChainYieldsAccountResponse struct for OnChainYieldsAccountResponse
type OnChainYieldsAccountResponse struct {
	TotalAmountInBTC          *string `json:"totalAmountInBTC,omitempty"`
	TotalAmountInUSDT         *string `json:"totalAmountInUSDT,omitempty"`
	TotalFlexibleAmountInBTC  *string `json:"totalFlexibleAmountInBTC,omitempty"`
	TotalFlexibleAmountInUSDT *string `json:"totalFlexibleAmountInUSDT,omitempty"`
	TotalLockedInBTC          *string `json:"totalLockedInBTC,omitempty"`
	TotalLockedInUSDT         *string `json:"totalLockedInUSDT,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _OnChainYieldsAccountResponse OnChainYieldsAccountResponse

// NewOnChainYieldsAccountResponse instantiates a new OnChainYieldsAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnChainYieldsAccountResponse() *OnChainYieldsAccountResponse {
	this := OnChainYieldsAccountResponse{}
	return &this
}

// NewOnChainYieldsAccountResponseWithDefaults instantiates a new OnChainYieldsAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnChainYieldsAccountResponseWithDefaults() *OnChainYieldsAccountResponse {
	this := OnChainYieldsAccountResponse{}
	return &this
}

// GetTotalAmountInBTC returns the TotalAmountInBTC field value if set, zero value otherwise.
func (o *OnChainYieldsAccountResponse) GetTotalAmountInBTC() string {
	if o == nil || common.IsNil(o.TotalAmountInBTC) {
		var ret string
		return ret
	}
	return *o.TotalAmountInBTC
}

// GetTotalAmountInBTCOk returns a tuple with the TotalAmountInBTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnChainYieldsAccountResponse) GetTotalAmountInBTCOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAmountInBTC) {
		return nil, false
	}
	return o.TotalAmountInBTC, true
}

// HasTotalAmountInBTC returns a boolean if a field has been set.
func (o *OnChainYieldsAccountResponse) HasTotalAmountInBTC() bool {
	if o != nil && !common.IsNil(o.TotalAmountInBTC) {
		return true
	}

	return false
}

// SetTotalAmountInBTC gets a reference to the given string and assigns it to the TotalAmountInBTC field.
func (o *OnChainYieldsAccountResponse) SetTotalAmountInBTC(v string) {
	o.TotalAmountInBTC = &v
}

// GetTotalAmountInUSDT returns the TotalAmountInUSDT field value if set, zero value otherwise.
func (o *OnChainYieldsAccountResponse) GetTotalAmountInUSDT() string {
	if o == nil || common.IsNil(o.TotalAmountInUSDT) {
		var ret string
		return ret
	}
	return *o.TotalAmountInUSDT
}

// GetTotalAmountInUSDTOk returns a tuple with the TotalAmountInUSDT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnChainYieldsAccountResponse) GetTotalAmountInUSDTOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAmountInUSDT) {
		return nil, false
	}
	return o.TotalAmountInUSDT, true
}

// HasTotalAmountInUSDT returns a boolean if a field has been set.
func (o *OnChainYieldsAccountResponse) HasTotalAmountInUSDT() bool {
	if o != nil && !common.IsNil(o.TotalAmountInUSDT) {
		return true
	}

	return false
}

// SetTotalAmountInUSDT gets a reference to the given string and assigns it to the TotalAmountInUSDT field.
func (o *OnChainYieldsAccountResponse) SetTotalAmountInUSDT(v string) {
	o.TotalAmountInUSDT = &v
}

// GetTotalFlexibleAmountInBTC returns the TotalFlexibleAmountInBTC field value if set, zero value otherwise.
func (o *OnChainYieldsAccountResponse) GetTotalFlexibleAmountInBTC() string {
	if o == nil || common.IsNil(o.TotalFlexibleAmountInBTC) {
		var ret string
		return ret
	}
	return *o.TotalFlexibleAmountInBTC
}

// GetTotalFlexibleAmountInBTCOk returns a tuple with the TotalFlexibleAmountInBTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnChainYieldsAccountResponse) GetTotalFlexibleAmountInBTCOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalFlexibleAmountInBTC) {
		return nil, false
	}
	return o.TotalFlexibleAmountInBTC, true
}

// HasTotalFlexibleAmountInBTC returns a boolean if a field has been set.
func (o *OnChainYieldsAccountResponse) HasTotalFlexibleAmountInBTC() bool {
	if o != nil && !common.IsNil(o.TotalFlexibleAmountInBTC) {
		return true
	}

	return false
}

// SetTotalFlexibleAmountInBTC gets a reference to the given string and assigns it to the TotalFlexibleAmountInBTC field.
func (o *OnChainYieldsAccountResponse) SetTotalFlexibleAmountInBTC(v string) {
	o.TotalFlexibleAmountInBTC = &v
}

// GetTotalFlexibleAmountInUSDT returns the TotalFlexibleAmountInUSDT field value if set, zero value otherwise.
func (o *OnChainYieldsAccountResponse) GetTotalFlexibleAmountInUSDT() string {
	if o == nil || common.IsNil(o.TotalFlexibleAmountInUSDT) {
		var ret string
		return ret
	}
	return *o.TotalFlexibleAmountInUSDT
}

// GetTotalFlexibleAmountInUSDTOk returns a tuple with the TotalFlexibleAmountInUSDT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnChainYieldsAccountResponse) GetTotalFlexibleAmountInUSDTOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalFlexibleAmountInUSDT) {
		return nil, false
	}
	return o.TotalFlexibleAmountInUSDT, true
}

// HasTotalFlexibleAmountInUSDT returns a boolean if a field has been set.
func (o *OnChainYieldsAccountResponse) HasTotalFlexibleAmountInUSDT() bool {
	if o != nil && !common.IsNil(o.TotalFlexibleAmountInUSDT) {
		return true
	}

	return false
}

// SetTotalFlexibleAmountInUSDT gets a reference to the given string and assigns it to the TotalFlexibleAmountInUSDT field.
func (o *OnChainYieldsAccountResponse) SetTotalFlexibleAmountInUSDT(v string) {
	o.TotalFlexibleAmountInUSDT = &v
}

// GetTotalLockedInBTC returns the TotalLockedInBTC field value if set, zero value otherwise.
func (o *OnChainYieldsAccountResponse) GetTotalLockedInBTC() string {
	if o == nil || common.IsNil(o.TotalLockedInBTC) {
		var ret string
		return ret
	}
	return *o.TotalLockedInBTC
}

// GetTotalLockedInBTCOk returns a tuple with the TotalLockedInBTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnChainYieldsAccountResponse) GetTotalLockedInBTCOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalLockedInBTC) {
		return nil, false
	}
	return o.TotalLockedInBTC, true
}

// HasTotalLockedInBTC returns a boolean if a field has been set.
func (o *OnChainYieldsAccountResponse) HasTotalLockedInBTC() bool {
	if o != nil && !common.IsNil(o.TotalLockedInBTC) {
		return true
	}

	return false
}

// SetTotalLockedInBTC gets a reference to the given string and assigns it to the TotalLockedInBTC field.
func (o *OnChainYieldsAccountResponse) SetTotalLockedInBTC(v string) {
	o.TotalLockedInBTC = &v
}

// GetTotalLockedInUSDT returns the TotalLockedInUSDT field value if set, zero value otherwise.
func (o *OnChainYieldsAccountResponse) GetTotalLockedInUSDT() string {
	if o == nil || common.IsNil(o.TotalLockedInUSDT) {
		var ret string
		return ret
	}
	return *o.TotalLockedInUSDT
}

// GetTotalLockedInUSDTOk returns a tuple with the TotalLockedInUSDT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnChainYieldsAccountResponse) GetTotalLockedInUSDTOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalLockedInUSDT) {
		return nil, false
	}
	return o.TotalLockedInUSDT, true
}

// HasTotalLockedInUSDT returns a boolean if a field has been set.
func (o *OnChainYieldsAccountResponse) HasTotalLockedInUSDT() bool {
	if o != nil && !common.IsNil(o.TotalLockedInUSDT) {
		return true
	}

	return false
}

// SetTotalLockedInUSDT gets a reference to the given string and assigns it to the TotalLockedInUSDT field.
func (o *OnChainYieldsAccountResponse) SetTotalLockedInUSDT(v string) {
	o.TotalLockedInUSDT = &v
}

func (o OnChainYieldsAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnChainYieldsAccountResponse) ToMap() (map[string]interface{}, error) {
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

func (o *OnChainYieldsAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varOnChainYieldsAccountResponse := _OnChainYieldsAccountResponse{}

	err = json.Unmarshal(data, &varOnChainYieldsAccountResponse)

	if err != nil {
		return err
	}

	*o = OnChainYieldsAccountResponse(varOnChainYieldsAccountResponse)

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

type NullableOnChainYieldsAccountResponse struct {
	value *OnChainYieldsAccountResponse
	isSet bool
}

func (v NullableOnChainYieldsAccountResponse) Get() *OnChainYieldsAccountResponse {
	return v.value
}

func (v *NullableOnChainYieldsAccountResponse) Set(val *OnChainYieldsAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOnChainYieldsAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOnChainYieldsAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnChainYieldsAccountResponse(val *OnChainYieldsAccountResponse) *NullableOnChainYieldsAccountResponse {
	return &NullableOnChainYieldsAccountResponse{value: val, isSet: true}
}

func (v NullableOnChainYieldsAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnChainYieldsAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
