/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the UniversalTransferResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UniversalTransferResponse{}

// UniversalTransferResponse struct for UniversalTransferResponse
type UniversalTransferResponse struct {
	TranId               *int64  `json:"tranId,omitempty"`
	ClientTranId         *string `json:"clientTranId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UniversalTransferResponse UniversalTransferResponse

// NewUniversalTransferResponse instantiates a new UniversalTransferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniversalTransferResponse() *UniversalTransferResponse {
	this := UniversalTransferResponse{}
	return &this
}

// NewUniversalTransferResponseWithDefaults instantiates a new UniversalTransferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniversalTransferResponseWithDefaults() *UniversalTransferResponse {
	this := UniversalTransferResponse{}
	return &this
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *UniversalTransferResponse) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniversalTransferResponse) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *UniversalTransferResponse) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *UniversalTransferResponse) SetTranId(v int64) {
	o.TranId = &v
}

// GetClientTranId returns the ClientTranId field value if set, zero value otherwise.
func (o *UniversalTransferResponse) GetClientTranId() string {
	if o == nil || common.IsNil(o.ClientTranId) {
		var ret string
		return ret
	}
	return *o.ClientTranId
}

// GetClientTranIdOk returns a tuple with the ClientTranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniversalTransferResponse) GetClientTranIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientTranId) {
		return nil, false
	}
	return o.ClientTranId, true
}

// HasClientTranId returns a boolean if a field has been set.
func (o *UniversalTransferResponse) HasClientTranId() bool {
	if o != nil && !common.IsNil(o.ClientTranId) {
		return true
	}

	return false
}

// SetClientTranId gets a reference to the given string and assigns it to the ClientTranId field.
func (o *UniversalTransferResponse) SetClientTranId(v string) {
	o.ClientTranId = &v
}

func (o UniversalTransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UniversalTransferResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	if !common.IsNil(o.ClientTranId) {
		toSerialize["clientTranId"] = o.ClientTranId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UniversalTransferResponse) UnmarshalJSON(data []byte) (err error) {
	varUniversalTransferResponse := _UniversalTransferResponse{}

	err = json.Unmarshal(data, &varUniversalTransferResponse)

	if err != nil {
		return err
	}

	*o = UniversalTransferResponse(varUniversalTransferResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tranId")
		delete(additionalProperties, "clientTranId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUniversalTransferResponse struct {
	value *UniversalTransferResponse
	isSet bool
}

func (v NullableUniversalTransferResponse) Get() *UniversalTransferResponse {
	return v.value
}

func (v *NullableUniversalTransferResponse) Set(val *UniversalTransferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUniversalTransferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUniversalTransferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniversalTransferResponse(val *UniversalTransferResponse) *NullableUniversalTransferResponse {
	return &NullableUniversalTransferResponse{value: val, isSet: true}
}

func (v NullableUniversalTransferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniversalTransferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
