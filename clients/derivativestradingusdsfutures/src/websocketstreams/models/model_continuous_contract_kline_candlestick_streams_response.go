/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ContinuousContractKlineCandlestickStreamsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ContinuousContractKlineCandlestickStreamsResponse{}

// ContinuousContractKlineCandlestickStreamsResponse struct for ContinuousContractKlineCandlestickStreamsResponse
type ContinuousContractKlineCandlestickStreamsResponse struct {
	Smalle               *string                                             `json:"e,omitempty"`
	E                    *int64                                              `json:"E,omitempty"`
	Smallps              *string                                             `json:"ps,omitempty"`
	Smallct              *string                                             `json:"ct,omitempty"`
	Smallk               *ContinuousContractKlineCandlestickStreamsResponseK `json:"k,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContinuousContractKlineCandlestickStreamsResponse ContinuousContractKlineCandlestickStreamsResponse

// NewContinuousContractKlineCandlestickStreamsResponse instantiates a new ContinuousContractKlineCandlestickStreamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinuousContractKlineCandlestickStreamsResponse() *ContinuousContractKlineCandlestickStreamsResponse {
	this := ContinuousContractKlineCandlestickStreamsResponse{}
	return &this
}

// NewContinuousContractKlineCandlestickStreamsResponseWithDefaults instantiates a new ContinuousContractKlineCandlestickStreamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinuousContractKlineCandlestickStreamsResponseWithDefaults() *ContinuousContractKlineCandlestickStreamsResponse {
	this := ContinuousContractKlineCandlestickStreamsResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *ContinuousContractKlineCandlestickStreamsResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousContractKlineCandlestickStreamsResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *ContinuousContractKlineCandlestickStreamsResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *ContinuousContractKlineCandlestickStreamsResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *ContinuousContractKlineCandlestickStreamsResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousContractKlineCandlestickStreamsResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *ContinuousContractKlineCandlestickStreamsResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *ContinuousContractKlineCandlestickStreamsResponse) SetE(v int64) {
	o.E = &v
}

// GetPs returns the Ps field value if set, zero value otherwise.
func (o *ContinuousContractKlineCandlestickStreamsResponse) GetSmallps() string {
	if o == nil || common.IsNil(o.Smallps) {
		var ret string
		return ret
	}
	return *o.Smallps
}

// GetPsOk returns a tuple with the Ps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousContractKlineCandlestickStreamsResponse) GetSmallpsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallps) {
		return nil, false
	}
	return o.Smallps, true
}

// HasPs returns a boolean if a field has been set.
func (o *ContinuousContractKlineCandlestickStreamsResponse) HasSmallps() bool {
	if o != nil && !common.IsNil(o.Smallps) {
		return true
	}

	return false
}

// SetPs gets a reference to the given string and assigns it to the Ps field.
func (o *ContinuousContractKlineCandlestickStreamsResponse) SetSmallps(v string) {
	o.Smallps = &v
}

// GetCt returns the Ct field value if set, zero value otherwise.
func (o *ContinuousContractKlineCandlestickStreamsResponse) GetSmallct() string {
	if o == nil || common.IsNil(o.Smallct) {
		var ret string
		return ret
	}
	return *o.Smallct
}

// GetCtOk returns a tuple with the Ct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousContractKlineCandlestickStreamsResponse) GetSmallctOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallct) {
		return nil, false
	}
	return o.Smallct, true
}

// HasCt returns a boolean if a field has been set.
func (o *ContinuousContractKlineCandlestickStreamsResponse) HasSmallct() bool {
	if o != nil && !common.IsNil(o.Smallct) {
		return true
	}

	return false
}

// SetCt gets a reference to the given string and assigns it to the Ct field.
func (o *ContinuousContractKlineCandlestickStreamsResponse) SetSmallct(v string) {
	o.Smallct = &v
}

// GetK returns the K field value if set, zero value otherwise.
func (o *ContinuousContractKlineCandlestickStreamsResponse) GetSmallk() ContinuousContractKlineCandlestickStreamsResponseK {
	if o == nil || common.IsNil(o.Smallk) {
		var ret ContinuousContractKlineCandlestickStreamsResponseK
		return ret
	}
	return *o.Smallk
}

// GetKOk returns a tuple with the K field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousContractKlineCandlestickStreamsResponse) GetSmallkOk() (*ContinuousContractKlineCandlestickStreamsResponseK, bool) {
	if o == nil || common.IsNil(o.Smallk) {
		return nil, false
	}
	return o.Smallk, true
}

// HasK returns a boolean if a field has been set.
func (o *ContinuousContractKlineCandlestickStreamsResponse) HasSmallk() bool {
	if o != nil && !common.IsNil(o.Smallk) {
		return true
	}

	return false
}

// SetK gets a reference to the given ContinuousContractKlineCandlestickStreamsResponseK and assigns it to the K field.
func (o *ContinuousContractKlineCandlestickStreamsResponse) SetSmallk(v ContinuousContractKlineCandlestickStreamsResponseK) {
	o.Smallk = &v
}

func (o ContinuousContractKlineCandlestickStreamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContinuousContractKlineCandlestickStreamsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smallps) {
		toSerialize["ps"] = o.Smallps
	}
	if !common.IsNil(o.Smallct) {
		toSerialize["ct"] = o.Smallct
	}
	if !common.IsNil(o.Smallk) {
		toSerialize["k"] = o.Smallk
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContinuousContractKlineCandlestickStreamsResponse) UnmarshalJSON(data []byte) (err error) {
	varContinuousContractKlineCandlestickStreamsResponse := _ContinuousContractKlineCandlestickStreamsResponse{}

	err = json.Unmarshal(data, &varContinuousContractKlineCandlestickStreamsResponse)

	if err != nil {
		return err
	}

	*o = ContinuousContractKlineCandlestickStreamsResponse(varContinuousContractKlineCandlestickStreamsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "ps")
		delete(additionalProperties, "ct")
		delete(additionalProperties, "k")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContinuousContractKlineCandlestickStreamsResponse struct {
	value *ContinuousContractKlineCandlestickStreamsResponse
	isSet bool
}

func (v NullableContinuousContractKlineCandlestickStreamsResponse) Get() *ContinuousContractKlineCandlestickStreamsResponse {
	return v.value
}

func (v *NullableContinuousContractKlineCandlestickStreamsResponse) Set(val *ContinuousContractKlineCandlestickStreamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousContractKlineCandlestickStreamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousContractKlineCandlestickStreamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousContractKlineCandlestickStreamsResponse(val *ContinuousContractKlineCandlestickStreamsResponse) *NullableContinuousContractKlineCandlestickStreamsResponse {
	return &NullableContinuousContractKlineCandlestickStreamsResponse{value: val, isSet: true}
}

func (v NullableContinuousContractKlineCandlestickStreamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousContractKlineCandlestickStreamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
