/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the FetchAddressVerificationListResponseInnerAddressQuestionnaire type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FetchAddressVerificationListResponseInnerAddressQuestionnaire{}

// FetchAddressVerificationListResponseInnerAddressQuestionnaire struct for FetchAddressVerificationListResponseInnerAddressQuestionnaire
type FetchAddressVerificationListResponseInnerAddressQuestionnaire struct {
	SendTo               *int64  `json:"sendTo,omitempty"`
	SatoshiToken         *string `json:"satoshiToken,omitempty"`
	IsAddressOwner       *int64  `json:"isAddressOwner,omitempty"`
	VerifyMethod         *int64  `json:"verifyMethod,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FetchAddressVerificationListResponseInnerAddressQuestionnaire FetchAddressVerificationListResponseInnerAddressQuestionnaire

// NewFetchAddressVerificationListResponseInnerAddressQuestionnaire instantiates a new FetchAddressVerificationListResponseInnerAddressQuestionnaire object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchAddressVerificationListResponseInnerAddressQuestionnaire() *FetchAddressVerificationListResponseInnerAddressQuestionnaire {
	this := FetchAddressVerificationListResponseInnerAddressQuestionnaire{}
	return &this
}

// NewFetchAddressVerificationListResponseInnerAddressQuestionnaireWithDefaults instantiates a new FetchAddressVerificationListResponseInnerAddressQuestionnaire object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchAddressVerificationListResponseInnerAddressQuestionnaireWithDefaults() *FetchAddressVerificationListResponseInnerAddressQuestionnaire {
	this := FetchAddressVerificationListResponseInnerAddressQuestionnaire{}
	return &this
}

// GetSendTo returns the SendTo field value if set, zero value otherwise.
func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) GetSendTo() int64 {
	if o == nil || common.IsNil(o.SendTo) {
		var ret int64
		return ret
	}
	return *o.SendTo
}

// GetSendToOk returns a tuple with the SendTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) GetSendToOk() (*int64, bool) {
	if o == nil || common.IsNil(o.SendTo) {
		return nil, false
	}
	return o.SendTo, true
}

// HasSendTo returns a boolean if a field has been set.
func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) HasSendTo() bool {
	if o != nil && !common.IsNil(o.SendTo) {
		return true
	}

	return false
}

// SetSendTo gets a reference to the given int64 and assigns it to the SendTo field.
func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) SetSendTo(v int64) {
	o.SendTo = &v
}

// GetSatoshiToken returns the SatoshiToken field value if set, zero value otherwise.
func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) GetSatoshiToken() string {
	if o == nil || common.IsNil(o.SatoshiToken) {
		var ret string
		return ret
	}
	return *o.SatoshiToken
}

// GetSatoshiTokenOk returns a tuple with the SatoshiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) GetSatoshiTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.SatoshiToken) {
		return nil, false
	}
	return o.SatoshiToken, true
}

// HasSatoshiToken returns a boolean if a field has been set.
func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) HasSatoshiToken() bool {
	if o != nil && !common.IsNil(o.SatoshiToken) {
		return true
	}

	return false
}

// SetSatoshiToken gets a reference to the given string and assigns it to the SatoshiToken field.
func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) SetSatoshiToken(v string) {
	o.SatoshiToken = &v
}

// GetIsAddressOwner returns the IsAddressOwner field value if set, zero value otherwise.
func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) GetIsAddressOwner() int64 {
	if o == nil || common.IsNil(o.IsAddressOwner) {
		var ret int64
		return ret
	}
	return *o.IsAddressOwner
}

// GetIsAddressOwnerOk returns a tuple with the IsAddressOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) GetIsAddressOwnerOk() (*int64, bool) {
	if o == nil || common.IsNil(o.IsAddressOwner) {
		return nil, false
	}
	return o.IsAddressOwner, true
}

// HasIsAddressOwner returns a boolean if a field has been set.
func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) HasIsAddressOwner() bool {
	if o != nil && !common.IsNil(o.IsAddressOwner) {
		return true
	}

	return false
}

// SetIsAddressOwner gets a reference to the given int64 and assigns it to the IsAddressOwner field.
func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) SetIsAddressOwner(v int64) {
	o.IsAddressOwner = &v
}

// GetVerifyMethod returns the VerifyMethod field value if set, zero value otherwise.
func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) GetVerifyMethod() int64 {
	if o == nil || common.IsNil(o.VerifyMethod) {
		var ret int64
		return ret
	}
	return *o.VerifyMethod
}

// GetVerifyMethodOk returns a tuple with the VerifyMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) GetVerifyMethodOk() (*int64, bool) {
	if o == nil || common.IsNil(o.VerifyMethod) {
		return nil, false
	}
	return o.VerifyMethod, true
}

// HasVerifyMethod returns a boolean if a field has been set.
func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) HasVerifyMethod() bool {
	if o != nil && !common.IsNil(o.VerifyMethod) {
		return true
	}

	return false
}

// SetVerifyMethod gets a reference to the given int64 and assigns it to the VerifyMethod field.
func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) SetVerifyMethod(v int64) {
	o.VerifyMethod = &v
}

func (o FetchAddressVerificationListResponseInnerAddressQuestionnaire) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchAddressVerificationListResponseInnerAddressQuestionnaire) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.SendTo) {
		toSerialize["sendTo"] = o.SendTo
	}
	if !common.IsNil(o.SatoshiToken) {
		toSerialize["satoshiToken"] = o.SatoshiToken
	}
	if !common.IsNil(o.IsAddressOwner) {
		toSerialize["isAddressOwner"] = o.IsAddressOwner
	}
	if !common.IsNil(o.VerifyMethod) {
		toSerialize["verifyMethod"] = o.VerifyMethod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FetchAddressVerificationListResponseInnerAddressQuestionnaire) UnmarshalJSON(data []byte) (err error) {
	varFetchAddressVerificationListResponseInnerAddressQuestionnaire := _FetchAddressVerificationListResponseInnerAddressQuestionnaire{}

	err = json.Unmarshal(data, &varFetchAddressVerificationListResponseInnerAddressQuestionnaire)

	if err != nil {
		return err
	}

	*o = FetchAddressVerificationListResponseInnerAddressQuestionnaire(varFetchAddressVerificationListResponseInnerAddressQuestionnaire)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sendTo")
		delete(additionalProperties, "satoshiToken")
		delete(additionalProperties, "isAddressOwner")
		delete(additionalProperties, "verifyMethod")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFetchAddressVerificationListResponseInnerAddressQuestionnaire struct {
	value *FetchAddressVerificationListResponseInnerAddressQuestionnaire
	isSet bool
}

func (v NullableFetchAddressVerificationListResponseInnerAddressQuestionnaire) Get() *FetchAddressVerificationListResponseInnerAddressQuestionnaire {
	return v.value
}

func (v *NullableFetchAddressVerificationListResponseInnerAddressQuestionnaire) Set(val *FetchAddressVerificationListResponseInnerAddressQuestionnaire) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchAddressVerificationListResponseInnerAddressQuestionnaire) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchAddressVerificationListResponseInnerAddressQuestionnaire) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchAddressVerificationListResponseInnerAddressQuestionnaire(val *FetchAddressVerificationListResponseInnerAddressQuestionnaire) *NullableFetchAddressVerificationListResponseInnerAddressQuestionnaire {
	return &NullableFetchAddressVerificationListResponseInnerAddressQuestionnaire{value: val, isSet: true}
}

func (v NullableFetchAddressVerificationListResponseInnerAddressQuestionnaire) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchAddressVerificationListResponseInnerAddressQuestionnaire) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
