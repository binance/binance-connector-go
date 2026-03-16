/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ReferencePriceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReferencePriceResponse{}

// ReferencePriceResponse struct for ReferencePriceResponse
type ReferencePriceResponse struct {
	E                    *string `json:"e,omitempty"`
	S                    *string `json:"s,omitempty"`
	R                    *string `json:"r,omitempty"`
	T                    *int64  `json:"t,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReferencePriceResponse ReferencePriceResponse

// NewReferencePriceResponse instantiates a new ReferencePriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferencePriceResponse() *ReferencePriceResponse {
	this := ReferencePriceResponse{}
	return &this
}

// NewReferencePriceResponseWithDefaults instantiates a new ReferencePriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferencePriceResponseWithDefaults() *ReferencePriceResponse {
	this := ReferencePriceResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *ReferencePriceResponse) GetE() string {
	if o == nil || common.IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse) GetEOk() (*string, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *ReferencePriceResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *ReferencePriceResponse) SetE(v string) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *ReferencePriceResponse) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *ReferencePriceResponse) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *ReferencePriceResponse) SetS(v string) {
	o.S = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *ReferencePriceResponse) GetR() string {
	if o == nil || common.IsNil(o.R) {
		var ret string
		return ret
	}
	return *o.R
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse) GetROk() (*string, bool) {
	if o == nil || common.IsNil(o.R) {
		return nil, false
	}
	return o.R, true
}

// HasR returns a boolean if a field has been set.
func (o *ReferencePriceResponse) HasR() bool {
	if o != nil && !common.IsNil(o.R) {
		return true
	}

	return false
}

// SetR gets a reference to the given string and assigns it to the R field.
func (o *ReferencePriceResponse) SetR(v string) {
	o.R = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *ReferencePriceResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *ReferencePriceResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *ReferencePriceResponse) SetT(v int64) {
	o.T = &v
}

func (o ReferencePriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferencePriceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	if !common.IsNil(o.S) {
		toSerialize["s"] = o.S
	}
	if !common.IsNil(o.R) {
		toSerialize["r"] = o.R
	}
	if !common.IsNil(o.T) {
		toSerialize["t"] = o.T
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReferencePriceResponse) UnmarshalJSON(data []byte) (err error) {
	varReferencePriceResponse := _ReferencePriceResponse{}

	err = json.Unmarshal(data, &varReferencePriceResponse)

	if err != nil {
		return err
	}

	*o = ReferencePriceResponse(varReferencePriceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "s")
		delete(additionalProperties, "r")
		delete(additionalProperties, "t")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReferencePriceResponse struct {
	value *ReferencePriceResponse
	isSet bool
}

func (v NullableReferencePriceResponse) Get() *ReferencePriceResponse {
	return v.value
}

func (v *NullableReferencePriceResponse) Set(val *ReferencePriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReferencePriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReferencePriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferencePriceResponse(val *ReferencePriceResponse) *NullableReferencePriceResponse {
	return &NullableReferencePriceResponse{value: val, isSet: true}
}

func (v NullableReferencePriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferencePriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
