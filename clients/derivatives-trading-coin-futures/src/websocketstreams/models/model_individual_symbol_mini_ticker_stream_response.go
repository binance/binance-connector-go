/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the IndividualSymbolMiniTickerStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IndividualSymbolMiniTickerStreamResponse{}

// IndividualSymbolMiniTickerStreamResponse struct for IndividualSymbolMiniTickerStreamResponse
type IndividualSymbolMiniTickerStreamResponse struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smallps              *string `json:"ps,omitempty"`
	Smallc               *string `json:"c,omitempty"`
	Smallo               *string `json:"o,omitempty"`
	Smallh               *string `json:"h,omitempty"`
	Smalll               *string `json:"l,omitempty"`
	Smallv               *string `json:"v,omitempty"`
	Smallq               *string `json:"q,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IndividualSymbolMiniTickerStreamResponse IndividualSymbolMiniTickerStreamResponse

// NewIndividualSymbolMiniTickerStreamResponse instantiates a new IndividualSymbolMiniTickerStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndividualSymbolMiniTickerStreamResponse() *IndividualSymbolMiniTickerStreamResponse {
	this := IndividualSymbolMiniTickerStreamResponse{}
	return &this
}

// NewIndividualSymbolMiniTickerStreamResponseWithDefaults instantiates a new IndividualSymbolMiniTickerStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndividualSymbolMiniTickerStreamResponseWithDefaults() *IndividualSymbolMiniTickerStreamResponse {
	this := IndividualSymbolMiniTickerStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *IndividualSymbolMiniTickerStreamResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *IndividualSymbolMiniTickerStreamResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *IndividualSymbolMiniTickerStreamResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *IndividualSymbolMiniTickerStreamResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetPs returns the Ps field value if set, zero value otherwise.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmallps() string {
	if o == nil || common.IsNil(o.Smallps) {
		var ret string
		return ret
	}
	return *o.Smallps
}

// GetPsOk returns a tuple with the Ps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmallpsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallps) {
		return nil, false
	}
	return o.Smallps, true
}

// HasPs returns a boolean if a field has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) HasSmallps() bool {
	if o != nil && !common.IsNil(o.Smallps) {
		return true
	}

	return false
}

// SetPs gets a reference to the given string and assigns it to the Ps field.
func (o *IndividualSymbolMiniTickerStreamResponse) SetSmallps(v string) {
	o.Smallps = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *IndividualSymbolMiniTickerStreamResponse) SetSmallc(v string) {
	o.Smallc = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *IndividualSymbolMiniTickerStreamResponse) SetSmallo(v string) {
	o.Smallo = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmallh() string {
	if o == nil || common.IsNil(o.Smallh) {
		var ret string
		return ret
	}
	return *o.Smallh
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmallhOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallh) {
		return nil, false
	}
	return o.Smallh, true
}

// HasH returns a boolean if a field has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) HasSmallh() bool {
	if o != nil && !common.IsNil(o.Smallh) {
		return true
	}

	return false
}

// SetH gets a reference to the given string and assigns it to the H field.
func (o *IndividualSymbolMiniTickerStreamResponse) SetSmallh(v string) {
	o.Smallh = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *IndividualSymbolMiniTickerStreamResponse) SetSmalll(v string) {
	o.Smalll = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmallv() string {
	if o == nil || common.IsNil(o.Smallv) {
		var ret string
		return ret
	}
	return *o.Smallv
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmallvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallv) {
		return nil, false
	}
	return o.Smallv, true
}

// HasV returns a boolean if a field has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) HasSmallv() bool {
	if o != nil && !common.IsNil(o.Smallv) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *IndividualSymbolMiniTickerStreamResponse) SetSmallv(v string) {
	o.Smallv = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *IndividualSymbolMiniTickerStreamResponse) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *IndividualSymbolMiniTickerStreamResponse) SetSmallq(v string) {
	o.Smallq = &v
}

func (o IndividualSymbolMiniTickerStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndividualSymbolMiniTickerStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallps) {
		toSerialize["ps"] = o.Smallps
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}
	if !common.IsNil(o.Smallh) {
		toSerialize["h"] = o.Smallh
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}
	if !common.IsNil(o.Smallv) {
		toSerialize["v"] = o.Smallv
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IndividualSymbolMiniTickerStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varIndividualSymbolMiniTickerStreamResponse := _IndividualSymbolMiniTickerStreamResponse{}

	err = json.Unmarshal(data, &varIndividualSymbolMiniTickerStreamResponse)

	if err != nil {
		return err
	}

	*o = IndividualSymbolMiniTickerStreamResponse(varIndividualSymbolMiniTickerStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "ps")
		delete(additionalProperties, "c")
		delete(additionalProperties, "o")
		delete(additionalProperties, "h")
		delete(additionalProperties, "l")
		delete(additionalProperties, "v")
		delete(additionalProperties, "q")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIndividualSymbolMiniTickerStreamResponse struct {
	value *IndividualSymbolMiniTickerStreamResponse
	isSet bool
}

func (v NullableIndividualSymbolMiniTickerStreamResponse) Get() *IndividualSymbolMiniTickerStreamResponse {
	return v.value
}

func (v *NullableIndividualSymbolMiniTickerStreamResponse) Set(val *IndividualSymbolMiniTickerStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIndividualSymbolMiniTickerStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIndividualSymbolMiniTickerStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndividualSymbolMiniTickerStreamResponse(val *IndividualSymbolMiniTickerStreamResponse) *NullableIndividualSymbolMiniTickerStreamResponse {
	return &NullableIndividualSymbolMiniTickerStreamResponse{value: val, isSet: true}
}

func (v NullableIndividualSymbolMiniTickerStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndividualSymbolMiniTickerStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
