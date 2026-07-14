/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ContractKlineStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ContractKlineStreamResponse{}

// ContractKlineStreamResponse struct for ContractKlineStreamResponse
type ContractKlineStreamResponse struct {
	// Contract address@Chain ID
	Ca *string `json:"ca,omitempty"`
	// Event type
	E                    *string                       `json:"e,omitempty"`
	K                    *ContractKlineStreamResponseK `json:"k,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContractKlineStreamResponse ContractKlineStreamResponse

// NewContractKlineStreamResponse instantiates a new ContractKlineStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractKlineStreamResponse() *ContractKlineStreamResponse {
	this := ContractKlineStreamResponse{}
	return &this
}

// NewContractKlineStreamResponseWithDefaults instantiates a new ContractKlineStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractKlineStreamResponseWithDefaults() *ContractKlineStreamResponse {
	this := ContractKlineStreamResponse{}
	return &this
}

// GetCa returns the Ca field value if set, zero value otherwise.
func (o *ContractKlineStreamResponse) GetCa() string {
	if o == nil || common.IsNil(o.Ca) {
		var ret string
		return ret
	}
	return *o.Ca
}

// GetCaOk returns a tuple with the Ca field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractKlineStreamResponse) GetCaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Ca) {
		return nil, false
	}
	return o.Ca, true
}

// HasCa returns a boolean if a field has been set.
func (o *ContractKlineStreamResponse) HasCa() bool {
	if o != nil && !common.IsNil(o.Ca) {
		return true
	}

	return false
}

// SetCa gets a reference to the given string and assigns it to the Ca field.
func (o *ContractKlineStreamResponse) SetCa(v string) {
	o.Ca = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *ContractKlineStreamResponse) GetE() string {
	if o == nil || common.IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractKlineStreamResponse) GetEOk() (*string, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *ContractKlineStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *ContractKlineStreamResponse) SetE(v string) {
	o.E = &v
}

// GetK returns the K field value if set, zero value otherwise.
func (o *ContractKlineStreamResponse) GetK() ContractKlineStreamResponseK {
	if o == nil || common.IsNil(o.K) {
		var ret ContractKlineStreamResponseK
		return ret
	}
	return *o.K
}

// GetKOk returns a tuple with the K field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractKlineStreamResponse) GetKOk() (*ContractKlineStreamResponseK, bool) {
	if o == nil || common.IsNil(o.K) {
		return nil, false
	}
	return o.K, true
}

// HasK returns a boolean if a field has been set.
func (o *ContractKlineStreamResponse) HasK() bool {
	if o != nil && !common.IsNil(o.K) {
		return true
	}

	return false
}

// SetK gets a reference to the given ContractKlineStreamResponseK and assigns it to the K field.
func (o *ContractKlineStreamResponse) SetK(v ContractKlineStreamResponseK) {
	o.K = &v
}

func (o ContractKlineStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractKlineStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Ca) {
		toSerialize["ca"] = o.Ca
	}
	if !common.IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	if !common.IsNil(o.K) {
		toSerialize["k"] = o.K
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContractKlineStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varContractKlineStreamResponse := _ContractKlineStreamResponse{}

	err = json.Unmarshal(data, &varContractKlineStreamResponse)

	if err != nil {
		return err
	}

	*o = ContractKlineStreamResponse(varContractKlineStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ca")
		delete(additionalProperties, "e")
		delete(additionalProperties, "k")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContractKlineStreamResponse struct {
	value *ContractKlineStreamResponse
	isSet bool
}

func (v NullableContractKlineStreamResponse) Get() *ContractKlineStreamResponse {
	return v.value
}

func (v *NullableContractKlineStreamResponse) Set(val *ContractKlineStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContractKlineStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContractKlineStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractKlineStreamResponse(val *ContractKlineStreamResponse) *NullableContractKlineStreamResponse {
	return &NullableContractKlineStreamResponse{value: val, isSet: true}
}

func (v NullableContractKlineStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractKlineStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
